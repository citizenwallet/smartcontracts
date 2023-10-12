import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { config } from "dotenv";
import { Contract } from "ethers";

// sould bound ERC721 with profile metadata uri
describe("Card", function () {
  config();
  const entrypoint = process.env["ERC4337_ENTRYPOINT"]!;

  async function deployCardFactoryFixture() {
    const [owner, friend1, friend2] = await ethers.getSigners();

    const WhitelistContract = await ethers.getContractFactory("Whitelist");
    const whitelist = await upgrades.deployProxy(WhitelistContract, [], {
      kind: "uups",
      initializer: "initialize",
    });

    const TimestampsLibraryContract = await ethers.getContractFactory(
      "Timestamps"
    );
    const timestamps = await upgrades.deployProxy(
      TimestampsLibraryContract,
      [],
      {
        kind: "uups",
        initializer: "initialize",
      }
    );

    const CardFactoryContract = await ethers.getContractFactory("CardFactory");
    const cardFactory = await upgrades.deployProxy(
      CardFactoryContract,
      [
        ethers.utils.getAddress(entrypoint),
        whitelist.address,
        timestamps.address,
      ],
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

    return {
      cardFactory,
      whitelist,
      owner,
      friend1,
      friend2,
      // deployed1,
      // deployed2,
      card1,
      card2,
    };
  }

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
    it("Should not be expired", async function () {
      const { card1, friend1, card2, friend2 } = await loadFixture(
        deployCardFactoryFixture
      );

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
  });
});
