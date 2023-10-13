import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers, upgrades, network } from "hardhat";
import { config } from "dotenv";
import { providers } from "ethers";
import { UserOperationBuilder, Presets } from "userop";

const accountABI = ["function execute(address to, uint256 value, bytes data)"];

// An ABI can be fragments and does not have to include the entire interface.
// As long as it includes the parts we want to use.
const partialERC20TokenABI = [
  "function transfer(address to, uint amount) returns (bool)",
];

const account = new ethers.utils.Interface(accountABI);
const erc20Token = new ethers.utils.Interface(partialERC20TokenABI);

// sould bound ERC721 with profile metadata uri
describe("Card", function () {
  config();

  async function deployCardFactoryFixture() {
    const [owner, friend1, friend2] = await ethers.getSigners();

    // console.log(await friend1.getBalance());

    // await network.provider.send("hardhat_setBalance", [
    //   friend1.address,
    //   "0x8AC7230489E80000",
    // ]);

    // await network.provider.send("hardhat_setBalance", [
    //   friend2.address,
    //   "0x8AC7230489E80000",
    // ]);

    // console.log(await friend1.getBalance());

    const TokenContract = await ethers.getContractFactory("RegensUniteToken", {
      signer: owner,
    });
    const token = await upgrades.deployProxy(TokenContract, [[owner.address]], {
      kind: "uups",
      initializer: "initialize",
    });

    const EntryPointContract = await ethers.getContractFactory(
      "EntryPointTest",
      { signer: owner }
    );
    const entrypoint = await EntryPointContract.deploy();

    const VerifyingPaymasterContract = await ethers.getContractFactory(
      "VerifyingPaymasterTest",
      { signer: owner }
    );
    const paymaster = await VerifyingPaymasterContract.deploy(
      entrypoint.address,
      owner.address
    );

    const WhitelistContract = await ethers.getContractFactory("Whitelist", {
      signer: owner,
    });
    const whitelist = await upgrades.deployProxy(WhitelistContract, [], {
      kind: "uups",
      initializer: "initialize",
    });

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

    await entrypoint.connect(owner).depositTo(paymaster.address, {
      value: ethers.BigNumber.from("1000000000000000000"),
    });
    await entrypoint.connect(owner).depositTo(friend1.address, {
      value: ethers.BigNumber.from("1000000000000000000"),
    });
    await entrypoint.connect(owner).depositTo(card1.address, {
      value: ethers.BigNumber.from("1000000000000000000"),
    });

    console.log(await entrypoint.deposits(paymaster.address));

    // add funds to paymaster
    await paymaster
      .connect(owner)
      .deposit({ value: ethers.BigNumber.from("1000000000000000000") });
    await paymaster
      .connect(friend1)
      .deposit({ value: ethers.BigNumber.from("1000000000000000000") });

    // add funds
    await token.connect(owner).mint(card1.address, 1000000000, "owner");
    await token.connect(owner).mint(friend2.address, 1000000000, "friend1");

    await network.provider.send("hardhat_setBalance", [
      card1.address,
      "0x8AC7230489E80000",
    ]);

    await card1
      .connect(friend1)
      .addDeposit({ value: ethers.BigNumber.from("1000000000000000000") });

    return {
      entrypoint,
      paymaster,
      token,
      cardFactory,
      whitelist,
      owner,
      friend1,
      friend2,
      card1,
      card2,
    };
  }

  describe("Entrypoint", function () {
    it("Should be available and have an appropriate SIG_VALIDATION_FAILED", async function () {
      const { entrypoint } = await loadFixture(deployCardFactoryFixture);

      expect(await entrypoint.SIG_VALIDATION_FAILED()).to.equal(1);
    });

    it("Should be able to handle ops", async function () {
      const { entrypoint, paymaster, token, owner, friend1, friend2, card1 } =
        await loadFixture(deployCardFactoryFixture);

      const addSender = async (ctx: any) => {
        ctx.op.sender = card1.address;
      };

      const resolveAccount = async (ctx: any) => {
        // Fetch the latest nonce and initCode if required.
        ctx.op.nonce = await card1.getNonce();
        ctx.op.initCode = "0x";
      };

      const constructCallData = async (ctx: any) => {
        ctx.op.callData = account.encodeFunctionData("execute", [
          token.address,
          ethers.constants.Zero,
          erc20Token.encodeFunctionData("transfer", [friend2.address, 0]),
        ]);
      };

      const addDefaultGasPrices = async (ctx: any) => {
        ctx.op.callGasLimit = "0x88b8";
        ctx.op.verificationGasLimit = "0x11170";
        ctx.op.preVerificationGas = "0x5208";
        ctx.op.maxFeePerGas = "0x47a4b601a";
        ctx.op.maxPriorityFeePerGas = "0xac6ca0";
      };

      const constructPaymasterData = async (ctx: any) => {
        const current = (await time.latest()) + 86400;
        const hash = await paymaster.getHash(ctx.op, current, current + 86400);

        ctx.op.paymasterAndData = hash;
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
        .useMiddleware(Presets.Middleware.EOASignature(friend1));

      const userop = await builder.buildOp(entrypoint.address, 1);

      console.log(userop);

      expect(await entrypoint.handleOps([userop], owner.address)).to.equal(1);

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
