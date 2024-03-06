import EntryPointArtifact from "@account-abstraction/contracts/artifacts/EntryPoint.json";
import "@nomicfoundation/hardhat-toolbox";
import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { config } from "dotenv";
import { BigNumber, BigNumberish, BytesLike } from "ethers";

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
  nonce: BigNumber;
  owner: string;
  code: BigNumber;
  accountFactoryAddress: string;
  callData: BytesLike;
}

interface PaymasterOOSignatureCreation {
  sender: string;
  paymasterContract: any; // using Contract type causes errors
  entrypoint: any; // using Contract type causes errors
  sponsor: SignerWithAddress;
}

interface UserOpPaymasterData {
  paymasterContract: any; // using Contract type causes errors
  hashData: string;
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

const emptyTransferCallData = () =>
  accountExecute.encodeFunctionData("execute", [
    ethers.constants.AddressZero,
    ethers.constants.Zero,
    erc20Token.encodeFunctionData("transfer", [
      ethers.constants.AddressZero,
      ethers.constants.Zero,
    ]),
  ]);

const upgradeCallData = (accountAddress: string, implementation: string) =>
  accountExecute.encodeFunctionData("execute", [
    accountAddress,
    ethers.constants.Zero,
    accountUpgrade.encodeFunctionData("upgradeTo", [implementation]),
  ]);

const createUserOp = async ({
  sender,
  nonce,
  owner,
  code,
  accountFactoryAddress,
  callData,
}: UserOpCreation) => {
  const userop: IUserOp = {
    sender,
    nonce: nonce,
    initCode: ethers.utils.arrayify("0x"),
    callData: ethers.utils.arrayify("0x"),
    callGasLimit: ethers.BigNumber.from("0"),
    verificationGasLimit: ethers.BigNumber.from("1500000"),
    preVerificationGas: ethers.BigNumber.from("21000"),
    maxFeePerGas: ethers.BigNumber.from("0"),
    maxPriorityFeePerGas: ethers.BigNumber.from("1000000000"),
    paymasterAndData: ethers.utils.arrayify("0x"),
    signature: ethers.utils.arrayify("0x"),
  };

  // get the uint64 seq from the nonce
  const seq = nonce.mask(64);

  // initCode
  if (seq.eq(ethers.constants.Zero)) {
    const accountCreationCode = voucherCreate.encodeFunctionData(
      "claimVoucher",
      [owner, code]
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

const getPaymasterOOSignature = async ({
  sender,
  paymasterContract,
  entrypoint,
  sponsor,
}: PaymasterOOSignatureCreation): Promise<[string, BigNumber]> => {
  // Generate a random 24-byte value
  const randomBytes = ethers.utils.randomBytes(24);

  // Interpret the random bytes as a uint192
  const key = ethers.BigNumber.from(randomBytes);

  const seq = ethers.BigNumber.from("0");

  const nonce = key.shl(64).add(seq);

  // nonce is uint256
  const epNonce: BigNumber = await entrypoint.getNonce(sender, key);
  if (!epNonce.eq(nonce)) {
    throw Error("nonce already used");
  }

  const userop: IUserOp = {
    sender,
    nonce: nonce,
    initCode: ethers.utils.arrayify("0x"),
    callData: emptyTransferCallData(),
    callGasLimit: ethers.BigNumber.from("0"),
    verificationGasLimit: ethers.BigNumber.from("1500000"),
    preVerificationGas: ethers.BigNumber.from("21000"),
    maxFeePerGas: ethers.BigNumber.from("0"),
    maxPriorityFeePerGas: ethers.BigNumber.from("1000000000"),
    paymasterAndData: ethers.utils.arrayify("0x"),
    signature: ethers.utils.arrayify("0x"),
  };

  const current = await time.latest();

  const validUntil = current + 2592000; // 30 days
  const validAfter = current;

  const hash = ethers.utils.arrayify(
    await paymasterContract.getHash(userop, validUntil, validAfter)
  );

  return [
    `${await sponsor.signMessage(hash)}|${validUntil}|${validAfter}`,
    nonce,
  ];
};

const getPaymasterAndData = async ({
  paymasterContract,
  hashData,
}: UserOpPaymasterData): Promise<string> => {
  // paymasterAndData

  const [signedHash, sValidUntil, sValidAfter] = hashData.split("|");

  const validUntil = ethers.BigNumber.from(sValidUntil);
  const validAfter = ethers.BigNumber.from(sValidAfter);

  // Define the types of the values
  const types = ["uint48", "uint48"];

  // Define the values
  const values: BigNumber[] = [validUntil, validAfter];

  // ABI encode the values
  const encoded = ethers.utils.defaultAbiCoder.encode(types, values);

  return ethers.utils.hexConcat([
    paymasterContract.address,
    encoded,
    signedHash,
  ]);
};

const signUserOp = async (
  userop: IUserOp,
  tokenEntryPointContract: any,
  signer: SignerWithAddress
) => {
  const userOpHash = ethers.utils.arrayify(
    await tokenEntryPointContract.getUserOpHash(userop)
  );

  return await signer.signMessage(userOpHash);
};

const accountExecuteABI = [
  "function execute(address to, uint256 value, bytes data)",
];
const voucherCreateABI = ["function claimVoucher(address, uint256)"];
const accountUpgradeABI = ["function upgradeTo(address newImplementation)"];

// An ABI can be fragments and does not have to include the entire interface.
// As long as it includes the parts we want to use.
const partialERC20TokenABI = [
  "function transfer(address to, uint amount) returns (bool)",
];

const accountExecute = new ethers.utils.Interface(accountExecuteABI);
const voucherCreate = new ethers.utils.Interface(voucherCreateABI);
const accountUpgrade = new ethers.utils.Interface(accountUpgradeABI);
const erc20Token = new ethers.utils.Interface(partialERC20TokenABI);

// cards with paymaster and whitelist
describe("Voucher", function () {
  config();

  async function deployAccountFactoryFixture() {
    const [owner, friend1, friend2, friend3, sponsor, sponsor2] =
      await ethers.getSigners();

    const TokenContract = await ethers.getContractFactory(
      "UpgradeableCommunityToken",
      {
        signer: owner,
      }
    );
    const token = await upgrades.deployProxy(
      TokenContract,
      [[owner.address], "My Token", "MT"],
      {
        kind: "uups",
        initializer: "initialize",
      }
    );

    const token2 = await upgrades.deployProxy(
      TokenContract,
      [[owner.address], "My Token 2", "MT2"],
      {
        kind: "uups",
        initializer: "initialize",
      }
    );

    const EntryPointContract = await ethers.getContractFactory(
      EntryPointArtifact.abi,
      EntryPointArtifact.bytecode,
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

    const VoucherFactoryContract = await ethers.getContractFactory(
      "VoucherFactory",
      {
        signer: owner,
      }
    );

    const voucherFactory = await VoucherFactoryContract.deploy(
      entrypoint.address,
      tokenEntryPointContract.address
    );

    await voucherFactory.deployed();

    await voucherFactory
      .connect(friend1)
      .claimVoucher(friend1.address, ethers.BigNumber.from(123));

    await voucherFactory
      .connect(friend2)
      .claimVoucher(friend2.address, ethers.BigNumber.from(456));

    const voucher1Hash = await voucherFactory.getCodeHash(
      ethers.BigNumber.from(123)
    );

    const voucher1 = await ethers.getContractAt(
      "Voucher",
      await voucherFactory.getVoucherAddress(voucher1Hash)
    );

    const voucher2Hash = await voucherFactory.getCodeHash(
      ethers.BigNumber.from(456)
    );

    const voucher2 = await ethers.getContractAt(
      "Voucher",
      await voucherFactory.getVoucherAddress(voucher2Hash)
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
      token2,
      tokenEntryPointContract,
      paymasterContract,
      voucherFactory,
      accountFactory,
      owner,
      friend1,
      friend2,
      friend3,
      account1,
      account2,
      voucher1,
      voucher2,
      sponsor,
      sponsor2,
    };
  }

  describe("Execute", function () {
    it("Owner should be able to transfer ERC20 directly", async function () {
      const { owner, token, friend1, voucher1, account2 } = await loadFixture(
        deployAccountFactoryFixture
      );

      const mintedAmount = 1000000000n;

      await token.connect(owner).mint(voucher1.address, mintedAmount);

      // balance should match what was minted
      expect(await token.balanceOf(voucher1.address)).to.equal(mintedAmount);

      // transfer to friend2
      await voucher1
        .connect(friend1)
        .execute(
          token.address,
          ethers.constants.Zero,
          erc20Token.encodeFunctionData("transfer", [account2.address, 100n])
        );

      // balance should match what was sent
      expect(await token.balanceOf(account2.address)).to.equal(100n);
    });

    it("Only owner should be able to transfer ERC20 directly", async function () {
      const { owner, token, voucher1, friend2, account2 } = await loadFixture(
        deployAccountFactoryFixture
      );

      const mintedAmount = 1000000000;

      await token.connect(owner).mint(voucher1.address, mintedAmount);

      // balance should match what was minted
      expect(await token.balanceOf(voucher1.address)).to.equal(mintedAmount);

      // transfer to friend2
      await expect(
        voucher1
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

    it("Should be able to transfer from a voucher with a userop in a single call", async function () {
      const {
        owner,
        token,
        tokenEntryPointContract,
        paymasterContract,
        friend2,
        friend3,
        sponsor,
        voucherFactory,
        accountFactory,
      } = await loadFixture(deployAccountFactoryFixture);

      const code = ethers.BigNumber.from(888);

      const codeHash1 = await voucherFactory.getCodeHash(code);

      const address = await voucherFactory.getVoucherAddress(codeHash1);

      const accountAddress2 = await accountFactory.getAddress(
        friend2.address,
        ethers.BigNumber.from(0)
      );

      const mintedAmount = 1000000000;

      await token.connect(owner).mint(address, mintedAmount);

      // balance should match what was minted
      expect(await token.balanceOf(address)).to.equal(mintedAmount);

      const transferAmount = 100n;

      const [hashData, nonce] = await getPaymasterOOSignature({
        sender: address,
        paymasterContract,
        entrypoint: tokenEntryPointContract,
        sponsor,
      });

      const userop = await createUserOp({
        sender: address,
        nonce,
        owner: friend3.address,
        code,
        accountFactoryAddress: voucherFactory.address,
        callData: transferCallData(
          token.address,
          accountAddress2,
          transferAmount
        ),
      });

      const paymasterAndData = await getPaymasterAndData({
        paymasterContract,
        hashData,
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
  });

  describe("Factory", function () {
    it("Should be owned by the caller", async function () {
      const { voucherFactory, token, friend2 } = await loadFixture(
        deployAccountFactoryFixture
      );

      const codeHash1 = await voucherFactory.getCodeHash(
        ethers.BigNumber.from(999)
      );

      await voucherFactory
        .connect(friend2)
        .claimVoucher(friend2.address, ethers.BigNumber.from(999));

      const voucher = await ethers.getContractAt(
        "Voucher",
        await voucherFactory.getVoucherAddress(codeHash1)
      );

      expect(await voucher.owner()).to.equal(friend2.address);
    });

    it("Should be possible to change the owner", async function () {
      const { friend1, friend2, voucher1 } = await loadFixture(
        deployAccountFactoryFixture
      );

      await voucher1.connect(friend1).transferOwnership(friend2.address);

      expect(await voucher1.connect(friend2).owner()).to.equal(friend2.address);
    });
  });

  describe("Address", function () {
    it("Should return the same address twice", async function () {
      const { voucherFactory } = await loadFixture(deployAccountFactoryFixture);

      const codeHash1 = await voucherFactory.getCodeHash(
        ethers.BigNumber.from(111)
      );

      const codeHash2 = await voucherFactory.getCodeHash(
        ethers.BigNumber.from(222)
      );

      expect(await voucherFactory.getVoucherAddress(codeHash1)).to.equal(
        await voucherFactory.getVoucherAddress(codeHash1)
      );

      expect(await voucherFactory.getVoucherAddress(codeHash1)).not.equal(
        await voucherFactory.getVoucherAddress(codeHash2)
      );
    });

    it("Should return a different address for a different codeHash", async function () {
      const { voucherFactory, token } = await loadFixture(
        deployAccountFactoryFixture
      );

      const codeHash1 = await voucherFactory.getCodeHash(
        ethers.BigNumber.from(111)
      );

      const codeHash2 = await voucherFactory.getCodeHash(
        ethers.BigNumber.from(222)
      );

      expect(await voucherFactory.getVoucherAddress(codeHash1)).not.equal(
        await voucherFactory.getVoucherAddress(codeHash2)
      );
    });
  });

  describe("Token EntryPoint", function () {
    it("Should return the matching token entrypoint with which the factory was deployed", async function () {
      const { tokenEntryPointContract, voucher1 } = await loadFixture(
        deployAccountFactoryFixture
      );

      expect(await voucher1.tokenEntryPoint()).equal(
        tokenEntryPointContract.address
      );
    });
  });

  describe("Account Recovery", function () {
    it("Token EntryPoint owner should be able to recover ownership for a voucher", async function () {
      const { owner, sponsor, friend1, friend2, voucher1 } = await loadFixture(
        deployAccountFactoryFixture
      );

      expect(await voucher1.owner()).equal(friend1.address);

      await expect(
        voucher1.connect(owner).transferOwnership(friend2.address)
      ).to.be.revertedWith("Ownable: caller is not the owner or the contract");

      await voucher1.connect(sponsor).recoverOwnership(friend2.address);

      expect(await voucher1.owner()).equal(friend2.address);
    });

    it("A random person shouldn't be able to recover ownership for a voucher", async function () {
      const { friend3, friend1, friend2, voucher1 } = await loadFixture(
        deployAccountFactoryFixture
      );

      expect(await voucher1.owner()).equal(friend1.address);

      await expect(
        voucher1.connect(friend3).recoverOwnership(friend2.address)
      ).to.be.revertedWith("Ownable: not TokenEntryPoint owner");

      expect(await voucher1.owner()).equal(friend1.address);
    });
  });
});
