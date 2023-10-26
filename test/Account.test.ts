import fs from "fs";
import path from "path";
import "@nomicfoundation/hardhat-toolbox";
import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers, upgrades, network } from "hardhat";
import { config } from "dotenv";
import { providers, ContractFactory } from "ethers";
import { UserOperationBuilder, Presets } from "userop";
import { keccak256 } from "ethers/lib/utils";

const accountABI = ["function execute(address to, uint256 value, bytes data)"];

// An ABI can be fragments and does not have to include the entire interface.
// As long as it includes the parts we want to use.
const partialERC20TokenABI = [
  "function transfer(address to, uint amount) returns (bool)",
];

const account = new ethers.utils.Interface(accountABI);
const erc20Token = new ethers.utils.Interface(partialERC20TokenABI);

// cards with paymaster and whitelist
describe("Account", function () {
  config();

  async function deployCardFactoryFixture() {
    const [owner, friend1, friend2, friend3, authorizer] =
      await ethers.getSigners();

    const TokenContract = await ethers.getContractFactory("RegensUniteToken", {
      signer: owner,
    });
    const token = await upgrades.deployProxy(TokenContract, [[owner.address]], {
      kind: "uups",
      initializer: "initialize",
    });

    const entrypointBin = fs
      .readFileSync(path.join(__dirname, "data", "entrypoint.bin"))
      .toString();
    const entrypointABI = JSON.parse(
      fs
        .readFileSync(path.join(__dirname, "data", "entrypoint.abi.json"))
        .toString()
    );

    const EntryPointContract = new ContractFactory(
      entrypointABI,
      entrypointBin,
      owner
    );
    const entrypoint = await EntryPointContract.deploy();

    const AccountFactoryContract = await ethers.getContractFactory(
      "AccountFactory",
      {
        signer: owner,
      }
    );
    const accountFactory = await AccountFactoryContract.deploy(
      entrypoint.address
    );
    // const accountFactory = await upgrades.deployProxy(
    //   AccountFactoryContract,
    //   [],
    //   {
    //     kind: "uups",
    //     initializer: "initialize",
    //     constructorArgs: [entrypoint.address],
    //   }
    // );

    await accountFactory.createAccount(
      friend1.address,
      ethers.BigNumber.from(0)
    );
    await accountFactory.createAccount(
      friend2.address,
      ethers.BigNumber.from(0)
    );

    const account1 = await ethers.getContractAt(
      "Account",
      await accountFactory.getAddress(friend1.address, ethers.BigNumber.from(0))
    );

    await account1.connect(friend1).updateAuthorizer(authorizer.address);
    await account1.connect(friend1).updateWhitelist([token.address]);

    const account2 = await ethers.getContractAt(
      "Account",
      await accountFactory.getAddress(friend2.address, ethers.BigNumber.from(0))
    );

    await account2.connect(friend2).updateAuthorizer(authorizer.address);
    await account2.connect(friend2).updateWhitelist([token.address]);

    await entrypoint.connect(owner).depositTo(owner.address, {
      value: ethers.BigNumber.from("1000000000000000000"),
    });

    await entrypoint.connect(owner).depositTo(friend1.address, {
      value: ethers.BigNumber.from("1000000000000000000"),
    });
    await entrypoint.connect(owner).depositTo(account1.address, {
      value: ethers.BigNumber.from("1000000000000000000"),
    });

    // add funds
    // await token.connect(owner).mint(account1.address, 1000000000, "owner");
    // await token.connect(owner).mint(friend2.address, 1000000000, "friend1");

    return {
      entrypoint,
      token,
      accountFactory,
      owner,
      friend1,
      friend2,
      friend3,
      account1,
      account2,
      authorizer,
    };
  }

  describe("Execute", function () {
    it("Should be able to transfer ERC20 directly (owner)", async function () {
      const { owner, token, friend1, account1, account2 } = await loadFixture(
        deployCardFactoryFixture
      );

      const mintedAmount = 1000000000;

      await token.connect(owner).mint(account1.address, mintedAmount, "owner");

      // balance should match what was minted
      expect(await token.balanceOf(account1.address)).to.equal(mintedAmount);

      // transfer to friend2
      await account1
        .connect(friend1)
        .execute(
          token.address,
          ethers.constants.Zero,
          erc20Token.encodeFunctionData("transfer", [account2.address, 100n])
        );

      // balance should match what was sent
      expect(await token.balanceOf(account2.address)).to.equal(100);
    });

    it("Someone else should not be able to transfer ERC20 directly (not owner)", async function () {
      const { owner, token, account1, friend2, account2 } = await loadFixture(
        deployCardFactoryFixture
      );

      const mintedAmount = 1000000000;

      await token.connect(owner).mint(account1.address, mintedAmount, "owner");

      // balance should match what was minted
      expect(await token.balanceOf(account1.address)).to.equal(mintedAmount);

      // transfer to friend2
      await expect(
        account1
          .connect(friend2)
          .execute(
            token.address,
            ethers.constants.Zero,
            erc20Token.encodeFunctionData("transfer", [account2.address, 100n])
          )
      ).to.be.revertedWith("account: not Owner or EntryPoint or Authorizer");

      // balance should match what was sent
      expect(await token.balanceOf(account2.address)).to.equal(0);
    });

    it("Authorizer should not be able to transfer ERC20 directly unless whitelisted (authorizer)", async function () {
      const { owner, token, friend2, friend3, authorizer, accountFactory } =
        await loadFixture(deployCardFactoryFixture);

      const address = await accountFactory.getAddress(
        friend3.address,
        ethers.BigNumber.from(0)
      );

      await accountFactory.createAccount(
        friend3.address,
        ethers.BigNumber.from(0)
      );

      const account = await ethers.getContractAt("Account", address);

      await account.connect(friend3).updateAuthorizer(authorizer.address);

      await accountFactory.createAccount(
        friend3.address,
        ethers.BigNumber.from(0)
      );

      const accountAddress2 = await accountFactory.getAddress(
        friend2.address,
        ethers.BigNumber.from(0)
      );

      const mintedAmount = 1000000000;

      await token.connect(owner).mint(account.address, mintedAmount, "owner");

      // balance should match what was minted
      expect(await token.balanceOf(account.address)).to.equal(mintedAmount);

      // transfer to friend2
      await expect(
        account
          .connect(authorizer)
          .execute(
            token.address,
            ethers.constants.Zero,
            erc20Token.encodeFunctionData("transfer", [accountAddress2, 100n])
          )
      ).to.be.revertedWith("account: address not whitelisted");

      await account.connect(friend3).updateWhitelist([token.address]);

      // transfer to friend2
      await account
        .connect(authorizer)
        .execute(
          token.address,
          ethers.constants.Zero,
          erc20Token.encodeFunctionData("transfer", [accountAddress2, 100n])
        );

      // balance should match what was sent
      expect(await token.balanceOf(accountAddress2)).to.equal(100n);
    });
  });

  // describe("Nonce", function () {
  //   it("Authorizer execution should increment nonce", async function () {
  //     const { owner, token, friend2, friend3, authorizer, accountFactory } =
  //       await loadFixture(deployCardFactoryFixture);

  //     const address = await accountFactory.getAddress(
  //       friend3.address,
  //       ethers.BigNumber.from(0)
  //     );

  //     await accountFactory.createAccount(
  //       friend3.address,
  //       ethers.BigNumber.from(0)
  //     );

  //     const account = await ethers.getContractAt("Account", address);

  //     await account.connect(friend3).updateAuthorizer(authorizer.address);

  //     await accountFactory.createAccount(
  //       friend3.address,
  //       ethers.BigNumber.from(0)
  //     );

  //     const accountAddress2 = await accountFactory.getAddress(
  //       friend2.address,
  //       ethers.BigNumber.from(0)
  //     );

  //     const mintedAmount = 1000000000;

  //     await token.connect(owner).mint(account.address, mintedAmount, "owner");

  //     // balance should match what was minted
  //     expect(await token.balanceOf(account.address)).to.equal(mintedAmount);

  //     // transfer to friend2
  //     await expect(
  //       account
  //         .connect(authorizer)
  //         .execute(
  //           token.address,
  //           ethers.constants.Zero,
  //           erc20Token.encodeFunctionData("transfer", [accountAddress2, 100n])
  //         )
  //     ).to.be.revertedWith("account: address not whitelisted");

  //     await account.connect(friend3).updateWhitelist([token.address]);

  //     // transfer to friend2
  //     await account
  //       .connect(authorizer)
  //       .execute(
  //         token.address,
  //         ethers.constants.Zero,
  //         erc20Token.encodeFunctionData("transfer", [accountAddress2, 100n])
  //       );

  //     // balance should match what was sent
  //     expect(await token.balanceOf(accountAddress2)).to.equal(100n);
  //   });
  // });

  describe("Factory", function () {
    it("Should be owned by the caller", async function () {
      const { accountFactory, friend1 } = await loadFixture(
        deployCardFactoryFixture
      );

      const account = await ethers.getContractAt(
        "Account",
        await accountFactory.getAddress(
          friend1.address,
          ethers.BigNumber.from(0)
        )
      );

      expect(await account.owner()).to.equal(friend1.address);
    });

    it("Should be possible to change the owner", async function () {
      const { accountFactory, friend1, friend2 } = await loadFixture(
        deployCardFactoryFixture
      );

      const address = friend1.address;
      const address2 = friend2.address;

      const account = await ethers.getContractAt(
        "Account",
        await accountFactory.getAddress(address, ethers.BigNumber.from(0))
      );

      await account.connect(friend1).transferOwnership(address2);

      expect(await account.connect(friend2).owner()).to.equal(address2);
    });
  });

  describe("Address", function () {
    it("Should return the same address twice", async function () {
      const { accountFactory, owner, friend1, friend2 } = await loadFixture(
        deployCardFactoryFixture
      );

      expect(
        await accountFactory.getAddress(owner.address, ethers.BigNumber.from(0))
      ).to.equal(
        await accountFactory.getAddress(owner.address, ethers.BigNumber.from(0))
      );

      expect(
        await accountFactory.getAddress(owner.address, ethers.BigNumber.from(0))
      ).not.equal(
        await accountFactory.getAddress(owner.address, ethers.BigNumber.from(1))
      );

      expect(
        await accountFactory.getAddress(
          friend1.address,
          ethers.BigNumber.from(0)
        )
      ).not.equal(
        await accountFactory.getAddress(
          friend2.address,
          ethers.BigNumber.from(0)
        )
      );
    });

    it("Should return a different address for a different salt", async function () {
      const { accountFactory, owner } = await loadFixture(
        deployCardFactoryFixture
      );

      expect(
        await accountFactory.getAddress(owner.address, ethers.BigNumber.from(0))
      ).not.equal(
        await accountFactory.getAddress(owner.address, ethers.BigNumber.from(1))
      );
    });

    it("Should return a different address for a different user", async function () {
      const { accountFactory, friend1, friend2 } = await loadFixture(
        deployCardFactoryFixture
      );

      expect(
        await accountFactory.getAddress(
          friend1.address,
          ethers.BigNumber.from(0)
        )
      ).not.equal(
        await accountFactory.getAddress(
          friend2.address,
          ethers.BigNumber.from(0)
        )
      );
    });
  });
});
