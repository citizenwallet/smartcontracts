import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { getRandomNumber } from "../scripts/functions/randomNumber";

describe("RedeemCodeFaucet", function () {
  async function deployOnboardingFaucetFixture() {
    const [owner, issuer, friend1, friend2] = await ethers.getSigners();

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

    const redeemInterval = 10;

    const OnboardingFaucetContract = await ethers.getContractFactory(
      "RedeemCodeFaucet"
    );
    const redeemCodeFaucet = await upgrades.deployProxy(
      OnboardingFaucetContract,
      [token.address, redeemInterval, issuer.address],
      {
        kind: "uups",
        initializer: "initialize",
      }
    );

    await token.mint(redeemCodeFaucet.address, 20, "hello");

    const network = await ethers.provider.getNetwork();

    const current = await time.latest();

    return {
      network,
      current,
      redeemCodeFaucet,
      redeemInterval,
      token,
      owner,
      issuer,
      friend1,
      friend2,
    };
  }

  describe("Deployment", function () {
    it("Should set the right owner", async function () {
      const { redeemCodeFaucet, owner } = await loadFixture(
        deployOnboardingFaucetFixture
      );

      expect(await redeemCodeFaucet.owner()).to.equal(owner.address);
    });

    it("Should set the right issuer as issuer role", async function () {
      const { redeemCodeFaucet, owner, issuer } = await loadFixture(
        deployOnboardingFaucetFixture
      );

      expect(await redeemCodeFaucet.owner()).to.equal(owner.address);
      expect(
        await redeemCodeFaucet.hasRole(
          await redeemCodeFaucet.ISSUER_ROLE(),
          issuer.address
        )
      ).to.equal(true);
    });
  });

  describe("Redeem", async function () {
    it("Should allow redeeming", async function () {
      const { friend1, token, current, redeemCodeFaucet, issuer } =
        await loadFixture(deployOnboardingFaucetFixture);

      expect(await token.balanceOf(friend1.address)).to.equal(0);

      const code = getRandomNumber(6);

      await redeemCodeFaucet
        .connect(issuer)
        .addRedeemCode(code, 10, current + 100);

      await redeemCodeFaucet.connect(friend1).redeem(issuer.address, code);

      expect(await token.balanceOf(friend1.address)).to.equal(10);
    });

    it("Should not allow redeeming again", async function () {
      const {
        friend1,
        token,
        redeemInterval,
        current,
        redeemCodeFaucet,
        issuer,
      } = await loadFixture(deployOnboardingFaucetFixture);

      expect(await token.balanceOf(friend1.address)).to.equal(0);

      const code = getRandomNumber(6);

      await redeemCodeFaucet
        .connect(issuer)
        .addRedeemCode(code, 10, current + 100);

      await redeemCodeFaucet.connect(friend1).redeem(issuer.address, code);

      expect(await token.balanceOf(friend1.address)).to.equal(10);

      await time.increase(redeemInterval);

      await expect(
        redeemCodeFaucet.connect(friend1).redeem(issuer.address, code)
      ).to.be.revertedWith("RedeemCodeFaucet: code already redeemed");
    });

    it("Should not allow redeeming again immediately", async function () {
      const { friend1, token, current, redeemCodeFaucet, issuer } =
        await loadFixture(deployOnboardingFaucetFixture);

      expect(await token.balanceOf(friend1.address)).to.equal(0);

      const code = getRandomNumber(6);

      await redeemCodeFaucet
        .connect(issuer)
        .addRedeemCode(code, 10, current + 100);

      await redeemCodeFaucet.connect(friend1).redeem(issuer.address, code);

      expect(await token.balanceOf(friend1.address)).to.equal(10);

      await expect(
        redeemCodeFaucet.connect(friend1).redeem(issuer.address, code)
      ).to.be.revertedWith("RedeemCodeFaucet: redeem interval not passed");
    });

    it("Should fail when insufficient balance", async function () {
      const {
        friend1,
        token,
        redeemInterval,
        current,
        redeemCodeFaucet,
        issuer,
      } = await loadFixture(deployOnboardingFaucetFixture);

      expect(await token.balanceOf(friend1.address)).to.equal(0);

      const code = getRandomNumber(6);

      await redeemCodeFaucet
        .connect(issuer)
        .addRedeemCode(code, 10, current + 100);

      await redeemCodeFaucet.connect(friend1).redeem(issuer.address, code);

      expect(await token.balanceOf(friend1.address)).to.equal(10);

      await time.increase(redeemInterval);

      const code2 = getRandomNumber(6);

      await redeemCodeFaucet
        .connect(issuer)
        .addRedeemCode(code2, 10, current + 100);

      await redeemCodeFaucet.connect(friend1).redeem(issuer.address, code2);

      await time.increase(redeemInterval);

      const code3 = getRandomNumber(6);

      await redeemCodeFaucet
        .connect(issuer)
        .addRedeemCode(code3, 10, current + 100);

      await expect(
        redeemCodeFaucet.connect(friend1).redeem(issuer.address, code3)
      ).to.be.revertedWith("RedeemCodeFaucet: insufficient balance");
    });
  });
});
