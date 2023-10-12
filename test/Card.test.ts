import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { config } from "dotenv";

// sould bound ERC721 with profile metadata uri
describe("Card", function () {
  config();
  const entrypoint = process.env["ERC4337_ENTRYPOINT"]!;

  async function deployCardFactoryFixture() {
    const [owner, friend1, friend2, deployed1, deployed2] =
      await ethers.getSigners();

    const CardFactoryContract = await ethers.getContractFactory("CardFactory");
    const cardFactory = await upgrades.deployProxy(
      CardFactoryContract,
      [ethers.utils.getAddress(entrypoint)],
      {
        kind: "uups",
        initializer: "initialize",
      }
    );

    await cardFactory.createAccount(
      deployed1.address,
      ethers.BigNumber.from(0)
    );
    await cardFactory.createAccount(
      deployed2.address,
      ethers.BigNumber.from(0)
    );

    const WhitelistContract = await ethers.getContractFactory("Whitelist");
    const whitelist = await upgrades.deployProxy(WhitelistContract, [], {
      kind: "uups",
      initializer: "initialize",
    });

    return {
      cardFactory,
      whitelist,
      owner,
      friend1,
      friend2,
      deployed1,
      deployed2,
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
      const { cardFactory, deployed1, deployed2 } = await loadFixture(
        deployCardFactoryFixture
      );

      const card = await ethers.getContractAt(
        "Card",
        await cardFactory.getAddress(
          deployed1.address,
          ethers.BigNumber.from(0)
        )
      );

      expect(await card.owner()).to.equal(deployed1.address);
    });

    it("Should be possible to change the owner", async function () {
      const { cardFactory, deployed1, deployed2 } = await loadFixture(
        deployCardFactoryFixture
      );

      const address = deployed1.address;
      const address2 = deployed2.address;

      const card = await ethers.getContractAt(
        "Card",
        await cardFactory.getAddress(address, ethers.BigNumber.from(0))
      );

      await card.connect(deployed1).transferOwnership(address2);

      expect(await card.connect(deployed1).owner()).to.equal(address2);
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

  describe("Account", function () {
    it("Should return the same address twice", async function () {
      const { cardFactory, owner, friend1, friend2 } = await loadFixture(
        deployCardFactoryFixture
      );
    });
  });
});
