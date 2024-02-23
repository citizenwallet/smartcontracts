import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers, upgrades } from "hardhat";

describe("SimpleFaucet", function () {
  async function deployOnboardingFaucetFixture() {
    const [owner, redeemAdmin, friend1, friend2] = await ethers.getSigners();

    const TokenContract = await ethers.getContractFactory(
      "UpgradeableBurnableCommunityToken",
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

    const OnboardingFaucetContract = await ethers.getContractFactory(
      "SimpleFaucet"
    );
    const singleRedeemFaucet = await upgrades.deployProxy(
      OnboardingFaucetContract,
      [token.address, 10, 0, redeemAdmin.address],
      {
        kind: "uups",
        initializer: "initialize",
      }
    );

    const redeemInterval = 10;

    const intervalRedeemFaucet = await upgrades.deployProxy(
      OnboardingFaucetContract,
      [token.address, 10, redeemInterval, redeemAdmin.address],
      {
        kind: "uups",
        initializer: "initialize",
      }
    );

    await token.mint(singleRedeemFaucet.address, 20, "hello");
    await token.mint(intervalRedeemFaucet.address, 40, "hello");

    const network = await ethers.provider.getNetwork();

    const current = await time.latest();

    return {
      network,
      current,
      singleRedeemFaucet,
      intervalRedeemFaucet,
      redeemInterval,
      token,
      owner,
      friend1,
      friend2,
      redeemAdmin,
    };
  }

  describe("Deployment", function () {
    it("Should set the right owner", async function () {
      const { singleRedeemFaucet, intervalRedeemFaucet, owner } =
        await loadFixture(deployOnboardingFaucetFixture);

      expect(await singleRedeemFaucet.owner()).to.equal(owner.address);
      expect(await intervalRedeemFaucet.owner()).to.equal(owner.address);
    });
  });

  describe("Withdraw", function () {
    it("Should allow codeCreator to withdraw", async function () {
      const { singleRedeemFaucet, token, redeemAdmin } = await loadFixture(
        deployOnboardingFaucetFixture
      );

      expect(await token.balanceOf(redeemAdmin.address)).to.equal(0);

      const faucetBalance = await token.balanceOf(singleRedeemFaucet.address);

      await singleRedeemFaucet.connect(redeemAdmin).withdraw();

      expect(await token.balanceOf(redeemAdmin.address)).to.equal(
        faucetBalance
      );

      expect(
        await singleRedeemFaucet.hasRole(
          await singleRedeemFaucet.REDEEM_ADMIN_ROLE(),
          redeemAdmin.address
        )
      ).to.equal(true);
    });

    it("Should only allow codeCreator to withdraw", async function () {
      const { singleRedeemFaucet, owner } = await loadFixture(
        deployOnboardingFaucetFixture
      );

      await expect(
        singleRedeemFaucet.connect(owner).withdraw()
      ).to.be.revertedWith(
        `AccessControl: account ${owner.address.toLowerCase()} is missing role ${await singleRedeemFaucet.REDEEM_ADMIN_ROLE()}`
      );
    });
  });

  describe("Single Redeem", async function () {
    it("Should allow redeeming", async function () {
      const { friend1, friend2, token, singleRedeemFaucet } = await loadFixture(
        deployOnboardingFaucetFixture
      );

      expect(await token.balanceOf(friend1.address)).to.equal(0);
      expect(await token.balanceOf(friend2.address)).to.equal(0);

      await singleRedeemFaucet.connect(friend1).redeem();

      await singleRedeemFaucet.connect(friend2).redeem();

      expect(await token.balanceOf(friend1.address)).to.equal(10);
      expect(await token.balanceOf(friend2.address)).to.equal(10);
    });

    it("Should not allow redeeming again", async function () {
      const { friend1, friend2, token, singleRedeemFaucet } = await loadFixture(
        deployOnboardingFaucetFixture
      );

      await singleRedeemFaucet.connect(friend1).redeem();

      await singleRedeemFaucet.connect(friend2).redeem();

      expect(await token.balanceOf(friend1.address)).to.equal(10);
      expect(await token.balanceOf(friend2.address)).to.equal(10);

      await expect(
        singleRedeemFaucet.connect(friend1).redeem()
      ).to.be.revertedWith("SimpleFaucet: already redeemed");
      await expect(
        singleRedeemFaucet.connect(friend2).redeem()
      ).to.be.revertedWith("SimpleFaucet: already redeemed");
    });
  });

  describe("Interval Redeem", async function () {
    it("Should allow redeeming", async function () {
      const { friend1, friend2, token, intervalRedeemFaucet } =
        await loadFixture(deployOnboardingFaucetFixture);

      expect(await token.balanceOf(friend1.address)).to.equal(0);
      expect(await token.balanceOf(friend2.address)).to.equal(0);

      await intervalRedeemFaucet.connect(friend1).redeem();
      await intervalRedeemFaucet.connect(friend2).redeem();

      expect(await token.balanceOf(friend1.address)).to.equal(10);
      expect(await token.balanceOf(friend2.address)).to.equal(10);
    });

    it("Should not allow redeeming again immediately", async function () {
      const { friend1, friend2, token, intervalRedeemFaucet } =
        await loadFixture(deployOnboardingFaucetFixture);

      await intervalRedeemFaucet.connect(friend1).redeem();
      await intervalRedeemFaucet.connect(friend2).redeem();

      expect(await token.balanceOf(friend1.address)).to.equal(10);
      expect(await token.balanceOf(friend2.address)).to.equal(10);

      await expect(
        intervalRedeemFaucet.connect(friend1).redeem()
      ).to.be.revertedWith("SimpleFaucet: redeem interval not passed");
      await expect(
        intervalRedeemFaucet.connect(friend2).redeem()
      ).to.be.revertedWith("SimpleFaucet: redeem interval not passed");
    });

    it("Should not allow redeeming again after interval", async function () {
      const { friend1, friend2, token, intervalRedeemFaucet, redeemInterval } =
        await loadFixture(deployOnboardingFaucetFixture);

      await intervalRedeemFaucet.connect(friend1).redeem();

      await intervalRedeemFaucet.connect(friend2).redeem();

      expect(await token.balanceOf(friend1.address)).to.equal(10);
      expect(await token.balanceOf(friend2.address)).to.equal(10);

      await time.increase(redeemInterval);

      await intervalRedeemFaucet.connect(friend1).redeem();
      await intervalRedeemFaucet.connect(friend2).redeem();
    });

    it("Should not allow redeeming if there are no more tokens", async function () {
      const { friend1, friend2, token, intervalRedeemFaucet, redeemInterval } =
        await loadFixture(deployOnboardingFaucetFixture);

      await intervalRedeemFaucet.connect(friend1).redeem();
      await intervalRedeemFaucet.connect(friend2).redeem();

      expect(await token.balanceOf(friend1.address)).to.equal(10);
      expect(await token.balanceOf(friend2.address)).to.equal(10);

      await time.increase(redeemInterval);

      await intervalRedeemFaucet.connect(friend1).redeem();
      await intervalRedeemFaucet.connect(friend2).redeem();

      await time.increase(redeemInterval);

      await expect(
        intervalRedeemFaucet.connect(friend1).redeem()
      ).to.be.revertedWith("SimpleFaucet: insufficient balance");
      await expect(
        intervalRedeemFaucet.connect(friend2).redeem()
      ).to.be.revertedWith("SimpleFaucet: insufficient balance");
    });
  });
});
