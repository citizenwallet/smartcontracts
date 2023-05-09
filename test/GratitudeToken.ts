import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { GratitudeToken } from "../typechain-types/GratitudeToken";

const ONE_GWEI = 1_000_000_000;

describe("GratitudeToken", function () {
  // We define a fixture to reuse the same setup in every test.
  // We use loadFixture to run this setup once, snapshot that state,
  // and reset Hardhat Network to that snapshot in every test.
  async function deployGratitudeTokenFixture() {

    // Contracts are deployed using the first signer/account by default
    const [owner, friend1, friend2] = await ethers.getSigners();

    const GratitudeTokenContract = await ethers.getContractFactory("GratitudeToken");
    const gratitudeToken = await upgrades.deployProxy(GratitudeTokenContract, [], { kind: 'uups', initializer: 'initialize' }) as GratitudeToken;

    return { gratitudeToken, owner, friend1, friend2 };
  }

  describe("Deployment", function () {
    it("Should set the right owner", async function () {
      const { gratitudeToken, owner } = await loadFixture(deployGratitudeTokenFixture);

      expect(await gratitudeToken.owner()).to.equal(owner.address);
    });
  });

  describe("Minting", function () {
    describe("", function () {
      it("Should revert with the right error if called from another account", async function () {
        const { gratitudeToken, friend1, friend2 } = await loadFixture(
          deployGratitudeTokenFixture
        );

        // We use gratitudeToken.connect() to send a transaction from another account
        await expect(gratitudeToken.connect(friend1).transfer(friend2.address, ONE_GWEI)).to.be.revertedWith(
          "Ownable: caller is not the owner"
        );
      });

      it("Should mint and send the gratitude tokens to the recipient", async function () {
        const { gratitudeToken, friend1, owner } = await loadFixture(
          deployGratitudeTokenFixture
        );
        await gratitudeToken.mint(friend1.address, ONE_GWEI);
        expect(await gratitudeToken.balanceOf(friend1.address)).to.equal(ONE_GWEI);
      });

      it("Should mint and send the gratitude tokens to many recipients", async function () {
        const { gratitudeToken, friend1, friend2, owner } = await loadFixture(
          deployGratitudeTokenFixture
        );

        await gratitudeToken.functions["mintToMany(address[],uint256)"]([friend1.address, friend2.address],ONE_GWEI);
        expect(await gratitudeToken.balanceOf(friend1.address)).to.equal(ONE_GWEI);
        expect(await gratitudeToken.balanceOf(friend2.address)).to.equal(ONE_GWEI);
      });
    });

    describe("Events", function () {
      it("Should emit a Transfer event when minting", async function () {
        const { gratitudeToken, friend1 } = await loadFixture(
          deployGratitudeTokenFixture
        );

        await expect(gratitudeToken.mint(friend1.address, ONE_GWEI))
          .to.emit(gratitudeToken, "Transfer")
          .withArgs(ethers.constants.AddressZero, friend1.address, ONE_GWEI);
      });
    });

  });
});
