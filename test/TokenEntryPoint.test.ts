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
  signerAddress: string;
  accountFactoryAddress: string;
  entrypoint: any; // using Contract type causes errors
  callData: BytesLike;
}

interface UserOpPaymasterData {
  userop: IUserOp;
  paymasterContract: any; // using Contract type causes errors
  sponsor: SignerWithAddress;
}

const transferCallData = (
  tokenAddress: string,
  receiver: string,
  amount: BigNumberish
) =>
  accountExecute.encodeFunctionData("execute", [
    tokenAddress,
    ethers.constants.Zero,
    erc20Token.encodeFunctionData("transfer", [receiver, amount]),
  ]);

const upgradeCallData = (accountAddress: string, implementation: string) =>
  accountExecute.encodeFunctionData("execute", [
    accountAddress,
    ethers.constants.Zero,
    accountUpgrade.encodeFunctionData("upgradeTo", [implementation]),
  ]);

const createUserOp = async ({
  sender,
  signerAddress,
  accountFactoryAddress,
  entrypoint,
  callData,
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
  const nonce: BigNumber = await entrypoint.getNonce(sender, 0);

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
  userop.callData = callData;

  return userop;
};

const getPaymasterAndData = async ({
  userop,
  paymasterContract,
  sponsor,
}: UserOpPaymasterData): Promise<string> => {
  // paymasterAndData

  const current = await time.latest();

  const validUntil = current + 86400;
  const validAfter = current;

  const hash = ethers.utils.arrayify(
    await paymasterContract.getHash(userop, validUntil, validAfter)
  );

  const signedHash = await sponsor.signMessage(hash);

  // Define the types of the values
  const types = ["uint48", "uint48"];

  // Define the values
  const values = [validUntil, validAfter];

  // ABI encode the values
  const encoded = ethers.utils.defaultAbiCoder.encode(types, values);

  return ethers.utils.hexConcat([
    paymasterContract.address,
    encoded,
    signedHash,
  ]);
};

const getUserOpHash = (userop: IUserOp, tokenEntryPointContract: any) => {
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
      tokenEntryPointContract.address,
      network.config.chainId,
    ]
  );
  return ethers.utils.keccak256(enc);
};

const signUserOp = async (
  userop: IUserOp,
  tokenEntryPointContract: any,
  signer: SignerWithAddress
) => {
  const userOpHash = getUserOpHash(userop, tokenEntryPointContract);

  return await signer.signMessage(ethers.utils.arrayify(userOpHash));
};

const accountExecuteABI = [
  "function execute(address to, uint256 value, bytes data)",
];
const accountCreateABI = ["function createAccount(address, uint256)"];
const accountUpgradeABI = ["function upgradeTo(address newImplementation)"];

// An ABI can be fragments and does not have to include the entire interface.
// As long as it includes the parts we want to use.
const partialERC20TokenABI = [
  "function transfer(address to, uint amount) returns (bool)",
];

const accountExecute = new ethers.utils.Interface(accountExecuteABI);
const accountCreate = new ethers.utils.Interface(accountCreateABI);
const accountUpgrade = new ethers.utils.Interface(accountUpgradeABI);
const erc20Token = new ethers.utils.Interface(partialERC20TokenABI);

// cards with paymaster and whitelist
describe("Account", function () {
  config();

  async function deployAccountFactoryFixture() {
    const [owner, friend1, friend2, friend3, sponsor, sponsor2] =
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

    const EntryPointContract = await ethers.getContractFactory(
      entrypointABI,
      entrypointBin,
      owner
    );
    const entrypoint = await EntryPointContract.deploy();

    await entrypoint.deployed();

    const PaymasterContract = await ethers.getContractFactory("Paymaster", {
      signer: owner,
    });

    const paymasterContract = await upgrades.deployProxy(
      PaymasterContract,
      [sponsor.address],
      {
        kind: "uups",
        initializer: "initialize",
      }
    );

    await paymasterContract.deployed();

    const TokenEntryPointContract = await ethers.getContractFactory(
      "TokenEntryPoint",
      {
        signer: owner,
      }
    );

    const tokenEntryPointContract = await upgrades.deployProxy(
      TokenEntryPointContract,
      [sponsor.address, paymasterContract.address, []],
      {
        kind: "uups",
        initializer: "initialize",
        constructorArgs: [entrypoint.address],
      }
    );

    await tokenEntryPointContract.deployed();

    const AccountFactoryContract = await ethers.getContractFactory(
      "AccountFactory",
      {
        signer: owner,
      }
    );

    const accountFactory = await AccountFactoryContract.deploy(
      entrypoint.address,
      tokenEntryPointContract.address
    );

    await accountFactory.deployed();

    const accountFactory2 = await AccountFactoryContract.deploy(
      entrypoint.address,
      tokenEntryPointContract.address
    );

    await accountFactory2.deployed();

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
      tokenEntryPointContract,
      paymasterContract,
      accountFactory,
      accountFactory2,
      owner,
      friend1,
      friend2,
      friend3,
      account1,
      account2,
      sponsor,
      sponsor2,
    };
  }

  describe("Execute", function () {
    it("Should be able to transfer ERC20 directly (owner)", async function () {
      const { owner, token, friend1, account1, account2 } = await loadFixture(
        deployAccountFactoryFixture
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
        deployAccountFactoryFixture
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
      ).to.be.revertedWith(
        "account: not Owner or EntryPoint or TokenEntryPoint"
      );

      // balance should match what was sent
      expect(await token.balanceOf(account2.address)).to.equal(0);
    });

    it("TokenEntryPoint should be able to transfer ERC20 if user op signed by user (sponsor)", async function () {
      const {
        entrypoint,
        owner,
        token,
        tokenEntryPointContract,
        paymasterContract,
        friend2,
        friend3,
        sponsor,
        accountFactory,
      } = await loadFixture(deployAccountFactoryFixture);

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
        signerAddress: friend3.address,
        accountFactoryAddress: accountFactory.address,
        entrypoint,
        callData: transferCallData(
          token.address,
          accountAddress2,
          transferAmount
        ),
      });

      const paymasterAndData = await getPaymasterAndData({
        userop,
        paymasterContract,
        sponsor,
      });

      userop.paymasterAndData = paymasterAndData;

      userop.signature = await signUserOp(
        userop,
        tokenEntryPointContract,
        friend3
      );

      await expect(
        tokenEntryPointContract.handleOps([userop], sponsor.address)
      ).to.be.revertedWith("AA28 contract not whitelisted");

      await tokenEntryPointContract
        .connect(sponsor)
        .updateWhitelist([token.address]);

      await tokenEntryPointContract.handleOps([userop], sponsor.address);

      // balance should match what was sent
      expect(await token.balanceOf(accountAddress2)).to.equal(transferAmount);

      // cannot replay transaction
      await expect(
        tokenEntryPointContract.handleOps([userop], sponsor.address)
      ).to.be.revertedWith("AA25 invalid account nonce");
    });

    it("Updating the verifying address of the paymaster should allow this new address to be used instead", async function () {
      const {
        entrypoint,
        owner,
        token,
        tokenEntryPointContract,
        friend2,
        friend3,
        sponsor,
        sponsor2,
        paymasterContract,
        accountFactory,
      } = await loadFixture(deployAccountFactoryFixture);

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
        signerAddress: friend3.address,
        accountFactoryAddress: accountFactory.address,
        entrypoint,
        callData: transferCallData(
          token.address,
          accountAddress2,
          transferAmount
        ),
      });

      const paymasterAndData = await getPaymasterAndData({
        userop,
        paymasterContract,
        sponsor: sponsor2,
      });

      userop.paymasterAndData = paymasterAndData;

      userop.signature = await signUserOp(
        userop,
        tokenEntryPointContract,
        friend3
      );

      await tokenEntryPointContract
        .connect(sponsor)
        .updateWhitelist([token.address]);

      await expect(
        tokenEntryPointContract.handleOps([userop], sponsor2.address)
      ).to.be.revertedWith("AA34 signature error");

      await expect(
        paymasterContract.connect(sponsor).updateSponsor(sponsor2.address)
      ).to.be.revertedWith("Ownable: caller is not the owner");

      await paymasterContract.connect(owner).updateSponsor(sponsor2.address);

      await tokenEntryPointContract.handleOps([userop], sponsor2.address);

      // balance should match what was sent
      expect(await token.balanceOf(accountAddress2)).to.equal(transferAmount);

      // cannot replay transaction
      await expect(
        tokenEntryPointContract.handleOps([userop], sponsor.address)
      ).to.be.revertedWith("AA25 invalid account nonce");
    });
  });

  describe("Factory", function () {
    it("Should be owned by the caller", async function () {
      const { accountFactory, friend1 } = await loadFixture(
        deployAccountFactoryFixture
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
        deployAccountFactoryFixture
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

    it("Should allow account upgrades directly", async function () {
      const { accountFactory, accountFactory2, friend1, account1 } =
        await loadFixture(deployAccountFactoryFixture);

      await accountFactory2.createAccount(friend1.address, 0);

      const slot = ethers.BigNumber.from(
        "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc"
      );

      // bytecode in the account factory and implementation of account 1 should match
      const account1Impl = `0x${(
        await ethers.provider.getStorageAt(account1.address, slot)
      ).slice(26)}`;

      const account1Bytecode = await ethers.provider.getCode(account1Impl);

      const accountFactoryBytecode = await ethers.provider.getCode(
        await accountFactory.accountImplementation()
      );

      expect(account1Bytecode).to.equal(accountFactoryBytecode);

      const accountFactory2Bytecode = await ethers.provider.getCode(
        await accountFactory2.accountImplementation()
      );

      expect(account1Bytecode).to.not.equal(accountFactory2Bytecode);

      const newAddress = await accountFactory2.getAddress(friend1.address, 0);

      expect(account1.address).to.not.equal(newAddress);

      const newAccount = await ethers.getContractAt("Account", newAddress);

      const implementation = `0x${(
        await ethers.provider.getStorageAt(newAccount.address, slot)
      ).slice(26)}`;

      await account1.connect(friend1).upgradeTo(implementation);

      const account1NewImpl = `0x${(
        await ethers.provider.getStorageAt(account1.address, slot)
      ).slice(26)}`;

      const account1NewBytecode = await ethers.provider.getCode(
        account1NewImpl
      );

      expect(account1NewBytecode).to.equal(accountFactory2Bytecode);

      // still useable
      expect(await account1.owner()).to.equal(friend1.address);
    });

    it("Should allow account upgrades through a user op", async function () {
      const {
        entrypoint,
        accountFactory,
        accountFactory2,
        friend1,
        account1,
        paymasterContract,
        sponsor,
        tokenEntryPointContract,
      } = await loadFixture(deployAccountFactoryFixture);

      await accountFactory2.createAccount(friend1.address, 0);

      const slot = ethers.BigNumber.from(
        "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc"
      );

      // bytecode in the account factory and implementation of account 1 should match
      const account1Impl = `0x${(
        await ethers.provider.getStorageAt(account1.address, slot)
      ).slice(26)}`;

      const account1Bytecode = await ethers.provider.getCode(account1Impl);

      const accountFactoryBytecode = await ethers.provider.getCode(
        await accountFactory.accountImplementation()
      );

      expect(account1Bytecode).to.equal(accountFactoryBytecode);

      const accountFactory2Bytecode = await ethers.provider.getCode(
        await accountFactory2.accountImplementation()
      );

      expect(account1Bytecode).to.not.equal(accountFactory2Bytecode);

      const newAddress = await accountFactory2.getAddress(friend1.address, 0);

      expect(account1.address).to.not.equal(newAddress);

      const newAccount = await ethers.getContractAt("Account", newAddress);

      const implementation = `0x${(
        await ethers.provider.getStorageAt(newAccount.address, slot)
      ).slice(26)}`;

      // upgrade account 1 to the new implementation by using a user op
      // await account1.connect(friend1).upgradeTo(implementation); << this but with a user op
      const userop = await createUserOp({
        sender: account1.address,
        signerAddress: friend1.address,
        accountFactoryAddress: accountFactory2.address,
        entrypoint,
        callData: upgradeCallData(account1.address, implementation),
      });

      const paymasterAndData = await getPaymasterAndData({
        userop,
        paymasterContract,
        sponsor: sponsor,
      });

      userop.paymasterAndData = paymasterAndData;

      userop.signature = await signUserOp(
        userop,
        tokenEntryPointContract,
        friend1
      );

      await tokenEntryPointContract.handleOps([userop], sponsor.address);

      const account1NewImpl = `0x${(
        await ethers.provider.getStorageAt(account1.address, slot)
      ).slice(26)}`;

      const account1NewBytecode = await ethers.provider.getCode(
        account1NewImpl
      );

      expect(account1NewBytecode).to.equal(accountFactory2Bytecode);

      // still useable
      expect(await account1.owner()).to.equal(friend1.address);
    });
  });

  describe("Address", function () {
    it("Should return the same address twice", async function () {
      const { accountFactory, owner, friend1, friend2 } = await loadFixture(
        deployAccountFactoryFixture
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
        deployAccountFactoryFixture
      );

      expect(
        await accountFactory.getAddress(owner.address, ethers.BigNumber.from(0))
      ).not.equal(
        await accountFactory.getAddress(owner.address, ethers.BigNumber.from(1))
      );
    });

    it("Should return a different address for a different user", async function () {
      const { accountFactory, friend1, friend2 } = await loadFixture(
        deployAccountFactoryFixture
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

  describe("Token EntryPoint", function () {
    it("Should return the matching token entrypoint with which the factory was deployed", async function () {
      const { tokenEntryPointContract, account1 } = await loadFixture(
        deployAccountFactoryFixture
      );

      expect(await account1.tokenEntryPoint()).equal(
        tokenEntryPointContract.address
      );
    });
  });
});
