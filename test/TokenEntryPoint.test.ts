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
  signerAddress: string;
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
  signerAddress,
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

  // key is a uint192
  // let key = BigNumber.from(0);
  // if (receiverAddress) {
  //   // convert address to uint192
  //   const receiver = ethers.BigNumber.from(receiverAddress);
  //   key = receiver.shl(32);
  // }

  // // nonce is uint256
  // const nonce: BigNumber = await entrypoint.getNonce(sender, key);

  // get the uint64 seq from the nonce
  const seq = nonce.mask(64);

  // // combine the uint64 nonce and uint192 key
  // const combined = key.shl(64).add(seq);
  // userop.nonce = combined;

  // initCode
  if (seq.eq(ethers.constants.Zero)) {
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

  const current = await time.latest();

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

    const accountFactory = await upgrades.deployProxy(
      AccountFactoryContract,
      [entrypoint.address, tokenEntryPointContract.address, owner.address],
      {
        kind: "uups",
        initializer: "initialize",
        constructorArgs: [],
        salt: "0x01",
      }
    );

    await accountFactory.deployed();

    const accountFactory2 = await upgrades.deployProxy(
      AccountFactoryContract,
      [entrypoint.address, tokenEntryPointContract.address, owner.address],
      {
        kind: "uups",
        initializer: "initialize",
        constructorArgs: [],
        salt: "0x02",
      }
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

      await token.connect(owner).mint(account1.address, mintedAmount);

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

      await token.connect(owner).mint(account1.address, mintedAmount);

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
        signerAddress: friend3.address,
        accountFactoryAddress: accountFactory.address,
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

    it("TokenEntryPoint should be able to transfer ERC20 in any order (out of order)", async function () {
      const {
        owner,
        token,
        tokenEntryPointContract,
        paymasterContract,
        friend1,
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

      const accountAddress1 = await accountFactory.getAddress(
        friend1.address,
        ethers.BigNumber.from(0)
      );

      const mintedAmount = BigNumber.from(1000000000);

      await token.connect(owner).mint(address, mintedAmount);

      // balance should match what was minted
      expect(await token.balanceOf(address)).to.equal(mintedAmount);

      const transferAmount = BigNumber.from(100);

      const [hashData, nonce] = await getPaymasterOOSignature({
        sender: address,
        paymasterContract,
        entrypoint: tokenEntryPointContract,
        sponsor,
      });

      const userop = await createUserOp({
        sender: address,
        nonce,
        signerAddress: friend3.address,
        accountFactoryAddress: accountFactory.address,
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

      const [hashData2, nonce2] = await getPaymasterOOSignature({
        sender: address,
        paymasterContract,
        entrypoint: tokenEntryPointContract,
        sponsor,
      });

      const userop2 = await createUserOp({
        sender: address,
        nonce: nonce2,
        signerAddress: friend3.address,
        accountFactoryAddress: accountFactory.address,
        callData: transferCallData(
          token.address,
          accountAddress1,
          transferAmount
        ),
      });

      const paymasterAndData2 = await getPaymasterAndData({
        paymasterContract,
        hashData: hashData2,
      });

      userop2.paymasterAndData = paymasterAndData2;

      userop2.signature = await signUserOp(
        userop2,
        tokenEntryPointContract,
        friend3
      );

      const [hashData3, nonce3] = await getPaymasterOOSignature({
        sender: address,
        paymasterContract,
        entrypoint: tokenEntryPointContract,
        sponsor,
      });

      const userop3 = await createUserOp({
        sender: address,
        nonce: nonce3,
        signerAddress: friend3.address,
        accountFactoryAddress: accountFactory.address,
        callData: transferCallData(
          token.address,
          accountAddress2,
          transferAmount
        ),
      });

      const paymasterAndData3 = await getPaymasterAndData({
        paymasterContract,
        hashData: hashData3,
      });

      userop3.paymasterAndData = paymasterAndData3;

      userop3.signature = await signUserOp(
        userop3,
        tokenEntryPointContract,
        friend3
      );

      // op 3
      await tokenEntryPointContract.handleOps([userop3], sponsor.address);

      // op 2
      await tokenEntryPointContract.handleOps([userop2], sponsor.address);

      // op 1
      await tokenEntryPointContract.handleOps([userop], sponsor.address);

      // cannot replay transaction
      await expect(
        tokenEntryPointContract.handleOps([userop], sponsor.address)
      ).to.be.revertedWith("AA25 invalid account nonce");

      // check all balances are correct

      // balance should match what was sent
      expect(await token.balanceOf(accountAddress1)).to.equal(transferAmount);

      // balance should match what was sent
      expect(await token.balanceOf(address)).to.equal(
        mintedAmount.sub(transferAmount.mul(3))
      );

      // balance should match what was sent
      expect(await token.balanceOf(accountAddress2)).to.equal(
        transferAmount.mul(2)
      );
    });

    it("Updating the sponsor address of the paymaster should allow this new address to be used instead", async function () {
      const {
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

      await token.connect(owner).mint(address, mintedAmount);

      // balance should match what was minted
      expect(await token.balanceOf(address)).to.equal(mintedAmount);

      const transferAmount = 100n;

      const [hashData, nonce] = await getPaymasterOOSignature({
        sender: address,
        paymasterContract,
        entrypoint: tokenEntryPointContract,
        sponsor: sponsor2,
      });

      const userop = await createUserOp({
        sender: address,
        nonce,
        signerAddress: friend3.address,
        accountFactoryAddress: accountFactory.address,
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

      const [hashData2, nonce2] = await getPaymasterOOSignature({
        sender: address,
        paymasterContract,
        entrypoint: tokenEntryPointContract,
        sponsor: sponsor2,
      });

      const userop2 = await createUserOp({
        sender: address,
        nonce: nonce2,
        signerAddress: friend3.address,
        accountFactoryAddress: accountFactory.address,
        callData: transferCallData(
          token.address,
          accountAddress2,
          transferAmount
        ),
      });

      const paymasterAndData2 = await getPaymasterAndData({
        paymasterContract,
        hashData: hashData2,
      });

      userop2.paymasterAndData = paymasterAndData2;

      userop2.signature = await signUserOp(
        userop2,
        tokenEntryPointContract,
        friend3
      );

      await tokenEntryPointContract.handleOps([userop2], sponsor2.address);

      // balance should match what was sent
      expect(await token.balanceOf(accountAddress2)).to.equal(transferAmount);

      // cannot replay transaction
      await expect(
        tokenEntryPointContract.handleOps([userop2], sponsor.address)
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

      const [hashData, nonce] = await getPaymasterOOSignature({
        sender: account1.address,
        paymasterContract,
        entrypoint: tokenEntryPointContract,
        sponsor,
      });

      // upgrade account 1 to the new implementation by using a user op
      // await account1.connect(friend1).upgradeTo(implementation); << this but with a user op
      const userop = await createUserOp({
        sender: account1.address,
        nonce,
        signerAddress: friend1.address,
        accountFactoryAddress: accountFactory2.address,
        callData: upgradeCallData(account1.address, implementation),
      });

      const paymasterAndData = await getPaymasterAndData({
        paymasterContract,
        hashData,
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

  describe("Account Recovery", function () {
    it("Token EntryPoint owner should be able to recovery ownership for an account", async function () {
      const { owner, sponsor, friend1, friend2, account1 } = await loadFixture(
        deployAccountFactoryFixture
      );

      expect(await account1.owner()).equal(friend1.address);

      await expect(
        account1.connect(owner).transferOwnership(friend2.address)
      ).to.be.revertedWith("Ownable: caller is not the owner or the contract");

      await account1.connect(sponsor).recoverOwnership(friend2.address);

      expect(await account1.owner()).equal(friend2.address);
    });

    it("A random person shouldn't be able to recovery ownership for an account", async function () {
      const { owner, friend3, friend1, friend2, account1 } = await loadFixture(
        deployAccountFactoryFixture
      );

      expect(await account1.owner()).equal(friend1.address);

      await expect(
        account1.connect(friend3).recoverOwnership(friend2.address)
      ).to.be.revertedWith("Ownable: not TokenEntryPoint owner");

      expect(await account1.owner()).equal(friend1.address);
    });
  });
});
