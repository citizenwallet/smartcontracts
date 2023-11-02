import fs from "fs";
import path from "path";
import "@nomicfoundation/hardhat-toolbox";
import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { expect } from "chai";
import { ethers, upgrades, network } from "hardhat";
import { config } from "dotenv";
import { ContractFactory, BigNumber, BigNumberish, BytesLike } from "ethers";

interface IUserOp {
  sender: string;
  nonce: BigNumberish;
  initCode: BytesLike;
  callData: BytesLike;
  callGasLimit: BigNumberish;
  verificationGasLimit: BigNumberish;
  preVerificationGas: BigNumberish;
  maxFeePerGas: BigNumberish;
  maxPriorityFeePerGas: BigNumberish;
  paymasterAndData: BytesLike;
  signature?: BytesLike;
}

interface UserOpCreation {
  sender: string;
  receiver: string;
  amount: BigNumberish;
  tokenAddress: string;
  signerAddress: string;
  accountFactoryAddress: string;
  authorizerContract: any; // using Contract type causes errors
  authorizer: SignerWithAddress;
}

const createUserOp = async ({
  sender,
  receiver,
  amount,
  tokenAddress,
  signerAddress,
  accountFactoryAddress,
  authorizerContract,
  authorizer,
}: UserOpCreation) => {
  const userop: IUserOp = {
    sender,
    nonce: ethers.BigNumber.from("0"),
    initCode: ethers.utils.arrayify("0x"),
    callData: ethers.utils.arrayify("0x"),
    callGasLimit: ethers.BigNumber.from("0"),
    verificationGasLimit: ethers.BigNumber.from("0"),
    preVerificationGas: ethers.BigNumber.from("0"),
    maxFeePerGas: ethers.BigNumber.from("0"),
    maxPriorityFeePerGas: ethers.BigNumber.from("0"),
    paymasterAndData: ethers.utils.arrayify("0x"),
    signature: ethers.utils.arrayify("0x"),
  };

  // nonce
  const nonce: BigNumber = await authorizerContract.getNonce(sender);

  userop.nonce = nonce;

  // initCode
  if (userop.nonce.eq(ethers.constants.Zero)) {
    const accountCreationCode = accountCreate.encodeFunctionData(
      "createAccount",
      [signerAddress, ethers.BigNumber.from(0)]
    );

    userop.initCode = ethers.utils.hexConcat([
      accountFactoryAddress,
      accountCreationCode,
    ]);
  }

  // callData
  userop.callData = accountExecute.encodeFunctionData("execute", [
    tokenAddress,
    ethers.constants.Zero,
    erc20Token.encodeFunctionData("transfer", [receiver, amount]),
  ]);

  // paymasterAndData

  const current = await time.latest();

  const hash = ethers.utils.arrayify(
    await authorizerContract.getHash(userop, current + 86400, current)
  );

  const signedHash = await authorizer.signMessage(hash);

  // Define the types of the values
  const types = ["uint48", "uint48"];

  // Define the values
  const values = [current + 86400, current];

  // ABI encode the values
  const encoded = ethers.utils.defaultAbiCoder.encode(types, values);

  userop.paymasterAndData = ethers.utils.hexConcat([
    authorizerContract.address,
    encoded,
    signedHash,
  ]);

  return userop;
};

const getUserOpHash = (userop: IUserOp, authorizerContract: any) => {
  const packedData = ethers.utils.defaultAbiCoder.encode(
    [
      "address",
      "uint256",
      "bytes32",
      "bytes32",
      "uint256",
      "uint256",
      "uint256",
      "uint256",
      "uint256",
      "bytes32",
    ],
    [
      userop.sender,
      userop.nonce,
      ethers.utils.keccak256(userop.initCode),
      ethers.utils.keccak256(userop.callData),
      userop.callGasLimit,
      userop.verificationGasLimit,
      userop.preVerificationGas,
      userop.maxFeePerGas,
      userop.maxPriorityFeePerGas,
      ethers.utils.keccak256(userop.paymasterAndData),
    ]
  );

  const enc = ethers.utils.defaultAbiCoder.encode(
    ["bytes32", "address", "uint256"],
    [
      ethers.utils.keccak256(packedData),
      authorizerContract.address,
      network.config.chainId,
    ]
  );
  return ethers.utils.keccak256(enc);
};

const signUserOp = async (
  userop: IUserOp,
  authorizerContract: any,
  signer: SignerWithAddress
) => {
  const userOpHash = getUserOpHash(userop, authorizerContract);

  return await signer.signMessage(ethers.utils.arrayify(userOpHash));
};

const accountExecuteABI = [
  "function execute(address to, uint256 value, bytes data)",
];
const accountCreateABI = ["function createAccount(address, uint256)"];
const accountUpdateAuthorizerABI = [
  "function updateAuthorizer(address newAuthorizer)",
];
const accountUpdateWhitelistABI = [
  "function updateWhitelist(address[] newWhitelist)",
];

// An ABI can be fragments and does not have to include the entire interface.
// As long as it includes the parts we want to use.
const partialERC20TokenABI = [
  "function transfer(address to, uint amount) returns (bool)",
];

const accountExecute = new ethers.utils.Interface(accountExecuteABI);
const accountCreate = new ethers.utils.Interface(accountCreateABI);
const accountUpdateAuthorizer = new ethers.utils.Interface(
  accountUpdateAuthorizerABI
);
const accountUpdateWhitelist = new ethers.utils.Interface(
  accountUpdateWhitelistABI
);
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

    const AuthorizerContract = await ethers.getContractFactory("Authorizer", {
      signer: owner,
    });

    const authorizerContract = await upgrades.deployProxy(
      AuthorizerContract,
      [authorizer.address],
      {
        kind: "uups",
        initializer: "initialize",
      }
    );

    const AccountFactoryContract = await ethers.getContractFactory(
      "AccountFactory",
      {
        signer: owner,
      }
    );

    const accountFactory = await AccountFactoryContract.deploy(
      entrypoint.address,
      authorizerContract.address
    );

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

    const account2 = await ethers.getContractAt(
      "Account",
      await accountFactory.getAddress(friend2.address, ethers.BigNumber.from(0))
    );

    await entrypoint.connect(owner).depositTo(owner.address, {
      value: ethers.BigNumber.from("1000000000000000000"),
    });

    await entrypoint.connect(owner).depositTo(friend1.address, {
      value: ethers.BigNumber.from("1000000000000000000"),
    });
    await entrypoint.connect(owner).depositTo(account1.address, {
      value: ethers.BigNumber.from("1000000000000000000"),
    });

    return {
      entrypoint,
      token,
      authorizerContract,
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

      const mintedAmount = 1000000000n;

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
      expect(await token.balanceOf(account2.address)).to.equal(100n);
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

    it("Authorizer should be able to transfer ERC20 if user op signed by user (authorizer)", async function () {
      const {
        owner,
        token,
        authorizerContract,
        friend2,
        friend3,
        authorizer,
        accountFactory,
      } = await loadFixture(deployCardFactoryFixture);

      const address = await accountFactory.getAddress(
        friend3.address,
        ethers.BigNumber.from(0)
      );

      const accountAddress2 = await accountFactory.getAddress(
        friend2.address,
        ethers.BigNumber.from(0)
      );

      const mintedAmount = 1000000000;

      await token.connect(owner).mint(address, mintedAmount, "owner");

      // balance should match what was minted
      expect(await token.balanceOf(address)).to.equal(mintedAmount);

      const transferAmount = 100n;

      const userop = await createUserOp({
        sender: address,
        receiver: accountAddress2,
        amount: transferAmount,
        tokenAddress: token.address,
        signerAddress: friend3.address,
        accountFactoryAddress: accountFactory.address,
        authorizerContract,
        authorizer,
      });

      userop.signature = await signUserOp(userop, authorizerContract, friend3);

      await authorizerContract.handleOps([userop], authorizer.address);

      // balance should match what was sent
      expect(await token.balanceOf(accountAddress2)).to.equal(transferAmount);

      // cannot replay transaction
      await expect(
        authorizerContract.handleOps([userop], authorizer.address)
      ).to.be.revertedWith("invalid nonce");
    });
  });

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
