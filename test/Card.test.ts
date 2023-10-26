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
describe("Card", function () {
  config();

  async function deployCardFactoryFixture() {
    const [owner, friend1, friend2, friend3, paymasterSigner] =
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

    const paymasterBin = fs
      .readFileSync(path.join(__dirname, "data", "paymaster.bin"))
      .toString();
    const paymasterABI = JSON.parse(
      fs
        .readFileSync(path.join(__dirname, "data", "paymaster.abi.json"))
        .toString()
    );

    const VerifyingPaymasterContract = new ContractFactory(
      paymasterABI,
      paymasterBin,
      paymasterSigner
    );
    const paymaster = await VerifyingPaymasterContract.deploy(
      entrypoint.address,
      paymasterSigner.address
    );

    const WhitelistContract = await ethers.getContractFactory("Whitelist", {
      signer: owner,
    });
    const whitelist = await upgrades.deployProxy(WhitelistContract, [], {
      kind: "uups",
      initializer: "initialize",
    });

    await whitelist.updateWhitelist([friend2.address]);

    const CardFactoryContract = await ethers.getContractFactory("CardFactory", {
      signer: owner,
    });
    const cardFactory = await upgrades.deployProxy(
      CardFactoryContract,
      [entrypoint.address, whitelist.address],
      {
        kind: "uups",
        initializer: "initialize",
      }
    );

    await cardFactory.createAccount(friend1.address, ethers.BigNumber.from(0));
    await cardFactory.createAccount(friend2.address, ethers.BigNumber.from(0));

    const card1 = await ethers.getContractAt(
      "Card",
      await cardFactory.getAddress(friend1.address, ethers.BigNumber.from(0))
    );

    const card2 = await ethers.getContractAt(
      "Card",
      await cardFactory.getAddress(friend2.address, ethers.BigNumber.from(0))
    );

    await entrypoint.connect(owner).depositTo(owner.address, {
      value: ethers.BigNumber.from("1000000000000000000"),
    });
    await entrypoint.connect(owner).depositTo(paymaster.address, {
      value: ethers.BigNumber.from("1000000000000000000"),
    });
    await entrypoint.connect(paymasterSigner).addStake(1, {
      value: ethers.BigNumber.from("1000000000000000000"),
    });
    await entrypoint.connect(owner).depositTo(friend1.address, {
      value: ethers.BigNumber.from("1000000000000000000"),
    });
    await entrypoint.connect(owner).depositTo(card1.address, {
      value: ethers.BigNumber.from("1000000000000000000"),
    });

    // add funds
    await token.connect(owner).mint(card1.address, 1000000000, "owner");
    await token.connect(owner).mint(friend2.address, 1000000000, "friend1");

    return {
      entrypoint,
      paymaster,
      token,
      cardFactory,
      whitelist,
      owner,
      friend1,
      friend2,
      friend3,
      card1,
      card2,
      paymasterSigner,
    };
  }

  describe("Entrypoint", function () {
    it("Should be available and have an appropriate SIG_VALIDATION_FAILED", async function () {
      const { entrypoint } = await loadFixture(deployCardFactoryFixture);

      expect(await entrypoint.SIG_VALIDATION_FAILED()).to.equal(1);
    });

    it("Should be able to handle ops", async function () {
      const {
        entrypoint,
        paymaster,
        token,
        owner,
        friend1,
        friend2,
        friend3,
        card1,
        paymasterSigner,
        cardFactory,
      } = await loadFixture(deployCardFactoryFixture);

      const cardAddress = await cardFactory.getAddress(
        friend3.address,
        ethers.BigNumber.from(0)
      );

      // await cardFactory.createAccount(
      //   friend3.address,
      //   ethers.BigNumber.from(0)
      // );

      console.log(`card: ${cardAddress}`);

      await token.connect(owner).mint(cardAddress, 1000000000, "owner");

      console.log(await token.balanceOf(cardAddress));

      const addSender = async (ctx: any) => {
        ctx.op.sender = cardAddress;
      };

      const resolveAccount = async (ctx: any) => {
        // Fetch the latest nonce and initCode if required.
        const nonce = await entrypoint.getNonce(cardAddress, 0n);
        ctx.op.nonce = nonce;

        const abi = ethers.utils.defaultAbiCoder;

        const initCode = keccak256(
          abi.encode(["address", "uint256"], [cardAddress, nonce])
        );

        // const initCode = cardFactory.interface.encodeFunctionData(
        //   "createAccount",
        //   [cardAddress, nonce]
        // );

        console.log("initCode: ", initCode);

        ctx.op.initCode = "0x";

        // ctx.op.initCode = `${cardFactory.address}${initCode.replace("0x", "")}`;
      };

      const constructCallData = async (ctx: any) => {
        ctx.op.callData = account.encodeFunctionData("execute", [
          token.address,
          ethers.constants.Zero,
          erc20Token.encodeFunctionData("transfer", [friend2.address, 100n]),
        ]);
      };

      const addDefaultGasPrices = async (ctx: any) => {
        ctx.op.callGasLimit = "0x989680";
        ctx.op.verificationGasLimit = "0x989680";
        ctx.op.preVerificationGas = "0x1";
        ctx.op.maxFeePerGas = "0x1";
        ctx.op.maxPriorityFeePerGas = "0x1";
      };

      const constructPaymasterData = async (ctx: any) => {
        const current = (await time.latest()) + 86400;
        const hash = await paymaster.getHash(ctx.op, current, current + 86400);

        console.log(`hash: ${hash}`);

        const signature = await paymasterSigner.signMessage(hash);

        const parsedSignature = ethers.utils.splitSignature(signature);
        const joinedSignature = ethers.utils.joinSignature(parsedSignature);

        console.log(
          `signed hash: ${signature}`,
          `length: ${joinedSignature.length}`
        );

        const abi = ethers.utils.defaultAbiCoder;

        const encodedData = keccak256(
          abi.encode(["uint48", "uint48"], [current, current + 86400])
        );

        console.log(
          `encodedData: ${encodedData}`,
          `length: ${encodedData.length}`
        );

        const paymasterAndData = `${paymaster.address}${encodedData.replace(
          "0x",
          ""
        )}${signature.replace("0x", "")}`;

        console.log(
          `paymasterAndData: ${paymasterAndData}`,
          `length: ${paymasterAndData.length}`
        );

        ctx.op.paymasterAndData = paymasterAndData;
        // ctx.op.paymasterAndData = "0x";
      };

      const provider = new providers.JsonRpcProvider(
        process.env["PUBLIC_RPC_URL"]
      );

      const builder = new UserOperationBuilder()
        .useMiddleware(addSender)
        .useMiddleware(resolveAccount)
        .useMiddleware(constructCallData)
        // .useMiddleware(Presets.Middleware.estimateUserOperationGas(provider))
        // .useMiddleware(Presets.Middleware.getGasPrice(provider))
        .useMiddleware(addDefaultGasPrices)
        .useMiddleware(constructPaymasterData)
        .useMiddleware(Presets.Middleware.EOASignature(friend3));

      const userop = await builder.buildOp(entrypoint.address, 1);

      console.log("submitting user op");
      console.log(userop);

      await entrypoint
        .connect(owner)
        .handleOps([userop], ethers.constants.AddressZero, {
          gasLimit: 10000000,
        });
      // await entrypoint
      //   .connect(entrypoint.signer)
      //   .simulateHandleOp([userop], ethers.constants.AddressZero, "", {
      //     gasLimit: 10000000,
      //   });

      console.log("submitted user op!");

      // expect(await entrypoint.handleOps([userop], owner.address)).to.equal(1);

      builder.resetOp();
    });
  });

  describe("Card Factory Deployment", function () {
    it("Should be owned by the deployer", async function () {
      const { cardFactory, owner } = await loadFixture(
        deployCardFactoryFixture
      );

      expect(await cardFactory.owner()).to.equal(owner.address);
    });

    it("Should be possible to change the owner", async function () {
      const { cardFactory, owner, friend1 } = await loadFixture(
        deployCardFactoryFixture
      );

      await cardFactory.connect(owner).transferOwnership(friend1.address);

      expect(await cardFactory.owner()).to.equal(friend1.address);
    });
  });

  describe("Card Deployment", function () {
    it("Should be owned by the caller", async function () {
      const { cardFactory, friend1 } = await loadFixture(
        deployCardFactoryFixture
      );

      const card = await ethers.getContractAt(
        "Card",
        await cardFactory.getAddress(friend1.address, ethers.BigNumber.from(0))
      );

      expect(await card.owner()).to.equal(friend1.address);
    });

    it("Should be possible to change the owner", async function () {
      const { cardFactory, friend1, friend2 } = await loadFixture(
        deployCardFactoryFixture
      );

      const address = friend1.address;
      const address2 = friend2.address;

      const card = await ethers.getContractAt(
        "Card",
        await cardFactory.getAddress(address, ethers.BigNumber.from(0))
      );

      await card.connect(friend1).transferOwnership(address2);

      expect(await card.connect(friend2).owner()).to.equal(address2);
    });
  });

  describe("Address", function () {
    it("Should return the same address twice", async function () {
      const { cardFactory, owner, friend1, friend2 } = await loadFixture(
        deployCardFactoryFixture
      );

      expect(
        await cardFactory.getAddress(owner.address, ethers.BigNumber.from(0))
      ).to.equal(
        await cardFactory.getAddress(owner.address, ethers.BigNumber.from(0))
      );

      expect(
        await cardFactory.getAddress(owner.address, ethers.BigNumber.from(0))
      ).not.equal(
        await cardFactory.getAddress(owner.address, ethers.BigNumber.from(1))
      );

      expect(
        await cardFactory.getAddress(friend1.address, ethers.BigNumber.from(0))
      ).not.equal(
        await cardFactory.getAddress(friend2.address, ethers.BigNumber.from(0))
      );
    });

    it("Should return a different address for a different salt", async function () {
      const { cardFactory, owner } = await loadFixture(
        deployCardFactoryFixture
      );

      expect(
        await cardFactory.getAddress(owner.address, ethers.BigNumber.from(0))
      ).not.equal(
        await cardFactory.getAddress(owner.address, ethers.BigNumber.from(1))
      );
    });

    it("Should return a different address for a different user", async function () {
      const { cardFactory, friend1, friend2 } = await loadFixture(
        deployCardFactoryFixture
      );

      expect(
        await cardFactory.getAddress(friend1.address, ethers.BigNumber.from(0))
      ).not.equal(
        await cardFactory.getAddress(friend2.address, ethers.BigNumber.from(0))
      );
    });
  });

  describe("Card", function () {
    it("Should not be expired after 2 years", async function () {
      const { card1, friend1, card2, friend2 } = await loadFixture(
        deployCardFactoryFixture
      );

      const current = await time.latest();

      await time.increaseTo(current + 86400 * 365 * 2);

      expect(await card1.connect(friend1).hasExpired()).to.equal(false);
      expect(await card2.connect(friend2).hasExpired()).to.equal(false);
    });

    it("Should allow anyone to check the expiry", async function () {
      const { card1, friend1, card2, friend2 } = await loadFixture(
        deployCardFactoryFixture
      );

      expect(await card1.connect(friend2).hasExpired()).to.equal(false);
      expect(await card2.connect(friend1).hasExpired()).to.equal(false);
    });

    it("Should be expired after 3 years", async function () {
      const { card1, friend1, card2, friend2 } = await loadFixture(
        deployCardFactoryFixture
      );

      const current = await time.latest();

      await time.increaseTo(current + (86400 * 365 * 3 + 86400));

      expect(await card1.connect(friend1).hasExpired()).to.equal(true);
      expect(await card2.connect(friend2).hasExpired()).to.equal(true);
    });
  });

  describe("Card Whitelist", function () {});

  describe("Card Timestamps", function () {});
});
