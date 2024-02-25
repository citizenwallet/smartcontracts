import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers, upgrades } from "hardhat";
import { getRandomNumber } from "../scripts/functions/randomNumber";

const getHash = (
  from: string,
  code: number,
  chainId: number,
  contract: string
) => {
  return ethers.utils.solidityKeccak256(
    ["address", "uint256", "uint256", "address"],
    [from, code, chainId, contract]
  );
};

describe("RedeemCodeFaucet", function () {
  async function deployOnboardingFaucetFixture() {
    const [owner, codeCreator, friend1, friend2] = await ethers.getSigners();

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

    const FaucetFactoryContract = await ethers.getContractFactory(
      "FaucetFactory"
    );

    const faucetFactory = await FaucetFactoryContract.deploy();

    const callInterval = 10;

    await faucetFactory.createRedeemCodeFaucet(
      owner.address,
      0n,
      token.address,
      callInterval,
      codeCreator.address
    );

    const redeemCodeFaucet = await ethers.getContractAt(
      "RedeemCodeFaucet",
      await faucetFactory.getRedeemCodeFaucetAddress(
        owner.address,
        0n,
        token.address,
        callInterval,
        codeCreator.address
      )
    );

    await token.mint(redeemCodeFaucet.address, 20, "hello");

    const network = await ethers.provider.getNetwork();

    const current = await time.latest();

    return {
      network,
      current,
      redeemCodeFaucet,
      callInterval,
      token,
      owner,
      codeCreator,
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

    it("Should set the right codeCreator as codeCreator role", async function () {
      const { redeemCodeFaucet, owner, codeCreator } = await loadFixture(
        deployOnboardingFaucetFixture
      );

      expect(await redeemCodeFaucet.owner()).to.equal(owner.address);
      expect(
        await redeemCodeFaucet.hasRole(
          await redeemCodeFaucet.REDEEM_CODE_CREATOR_ROLE(),
          codeCreator.address
        )
      ).to.equal(true);
    });
  });

  describe("Withdraw", function () {
    it("Should allow codeCreator to withdraw", async function () {
      const { redeemCodeFaucet, token, codeCreator } = await loadFixture(
        deployOnboardingFaucetFixture
      );

      expect(await token.balanceOf(codeCreator.address)).to.equal(0);

      const faucetBalance = await token.balanceOf(redeemCodeFaucet.address);

      await redeemCodeFaucet.connect(codeCreator).withdraw();

      expect(await token.balanceOf(codeCreator.address)).to.equal(
        faucetBalance
      );

      expect(
        await redeemCodeFaucet.hasRole(
          await redeemCodeFaucet.REDEEM_CODE_CREATOR_ROLE(),
          codeCreator.address
        )
      ).to.equal(true);
    });

    it("Should only allow codeCreator to withdraw", async function () {
      const { redeemCodeFaucet, owner } = await loadFixture(
        deployOnboardingFaucetFixture
      );

      await expect(
        redeemCodeFaucet.connect(owner).withdraw()
      ).to.be.revertedWith(
        `AccessControl: account ${owner.address.toLowerCase()} is missing role ${await redeemCodeFaucet.REDEEM_CODE_CREATOR_ROLE()}`
      );
    });
  });

  describe("Redeem", async function () {
    it("Should generate matching local hash", async function () {
      const { friend1, token, redeemCodeFaucet, network, codeCreator } =
        await loadFixture(deployOnboardingFaucetFixture);

      expect(await token.balanceOf(friend1.address)).to.equal(0);

      const code = getRandomNumber(6);

      const localHash = getHash(
        codeCreator.address,
        code,
        network.chainId,
        redeemCodeFaucet.address
      );

      const codeHash = await redeemCodeFaucet.getHash(
        codeCreator.address,
        code
      );

      expect(localHash).to.equal(codeHash);
    });

    it("Should allow redeeming", async function () {
      const {
        friend1,
        token,
        current,
        redeemCodeFaucet,
        network,
        codeCreator,
      } = await loadFixture(deployOnboardingFaucetFixture);

      expect(await token.balanceOf(friend1.address)).to.equal(0);

      const code = getRandomNumber(6);

      const codeHash = getHash(
        codeCreator.address,
        code,
        network.chainId,
        redeemCodeFaucet.address
      );

      await redeemCodeFaucet
        .connect(codeCreator)
        .addRedeemCode(codeHash, 10, current + 100);

      const tx = await redeemCodeFaucet.connect(friend1).redeem(code);
      const receipt = await tx.wait();

      expect(receipt.events?.find((e) => e.event === "CodeRedeemed")).to.not.be
        .undefined;

      expect(await token.balanceOf(friend1.address)).to.equal(10);
    });

    it("Should not allow redeeming again", async function () {
      const {
        friend1,
        token,
        callInterval,
        current,
        redeemCodeFaucet,
        codeCreator,
        network,
      } = await loadFixture(deployOnboardingFaucetFixture);

      expect(await token.balanceOf(friend1.address)).to.equal(0);

      const code = getRandomNumber(6);

      const codeHash = getHash(
        codeCreator.address,
        code,
        network.chainId,
        redeemCodeFaucet.address
      );

      await redeemCodeFaucet
        .connect(codeCreator)
        .addRedeemCode(codeHash, 10, current + 100);

      const tx = await redeemCodeFaucet.connect(friend1).redeem(code);
      const receipt = await tx.wait();

      expect(receipt.events?.find((e) => e.event === "CodeRedeemed")).to.not.be
        .undefined;

      expect(await token.balanceOf(friend1.address)).to.equal(10);

      await time.increase(callInterval);

      await expect(
        redeemCodeFaucet.connect(friend1).redeem(code)
      ).to.be.revertedWith("RedeemCodeFaucet: code already redeemed");
    });

    it("Should not allow redeeming again immediately", async function () {
      const {
        friend1,
        token,
        current,
        redeemCodeFaucet,
        codeCreator,
        network,
      } = await loadFixture(deployOnboardingFaucetFixture);

      expect(await token.balanceOf(friend1.address)).to.equal(0);

      const code = getRandomNumber(6);

      const codeHash = getHash(
        codeCreator.address,
        code,
        network.chainId,
        redeemCodeFaucet.address
      );

      await redeemCodeFaucet
        .connect(codeCreator)
        .addRedeemCode(codeHash, 10, current + 100);

      let tx = await redeemCodeFaucet.connect(friend1).redeem(code);
      let receipt = await tx.wait();

      expect(receipt.events?.find((e) => e.event === "CodeRedeemed")).to.not.be
        .undefined;

      expect(await token.balanceOf(friend1.address)).to.equal(10);

      tx = await redeemCodeFaucet.connect(friend1).redeem(code);
      receipt = await tx.wait();

      expect(receipt.events?.find((e) => e.event === "CodeRedeemed")).to.be
        .undefined;
    });

    it("Should not allow calling the redeem function too often", async function () {
      const {
        friend1,
        token,
        current,
        redeemCodeFaucet,
        codeCreator,
        network,
        callInterval,
      } = await loadFixture(deployOnboardingFaucetFixture);

      expect(await token.balanceOf(friend1.address)).to.equal(0);

      const code = getRandomNumber(6);

      const codeHash = getHash(
        codeCreator.address,
        code,
        network.chainId,
        redeemCodeFaucet.address
      );

      await redeemCodeFaucet
        .connect(codeCreator)
        .addRedeemCode(codeHash, 10, current + 100);

      let tx = await redeemCodeFaucet.connect(friend1).redeem(code);
      let receipt = await tx.wait();

      expect(receipt.events?.find((e) => e.event === "CodeRedeemed")).to.not.be
        .undefined;

      expect(await token.balanceOf(friend1.address)).to.equal(10);

      const code2 = getRandomNumber(6);

      const codeHash2 = getHash(
        codeCreator.address,
        code2,
        network.chainId,
        redeemCodeFaucet.address
      );

      await redeemCodeFaucet
        .connect(codeCreator)
        .addRedeemCode(codeHash2, 10, current + 100);

      tx = await redeemCodeFaucet.connect(friend1).redeem(code2);
      receipt = await tx.wait();

      expect(receipt.events?.find((e) => e.event === "CodeRedeemed")).to.be
        .undefined;

      await time.increase(callInterval);

      tx = await redeemCodeFaucet.connect(friend1).redeem(code2);
      receipt = await tx.wait();

      expect(receipt.events?.find((e) => e.event === "CodeRedeemed")).to.not.be
        .undefined;
    });

    it("Should fail when insufficient balance", async function () {
      const {
        friend1,
        token,
        callInterval,
        current,
        redeemCodeFaucet,
        codeCreator,
        network,
      } = await loadFixture(deployOnboardingFaucetFixture);

      expect(await token.balanceOf(friend1.address)).to.equal(0);

      const code = getRandomNumber(6);

      const codeHash = getHash(
        codeCreator.address,
        code,
        network.chainId,
        redeemCodeFaucet.address
      );

      await redeemCodeFaucet
        .connect(codeCreator)
        .addRedeemCode(codeHash, 10, current + 100);

      let tx = await redeemCodeFaucet.connect(friend1).redeem(code);
      let receipt = await tx.wait();

      expect(receipt.events?.find((e) => e.event === "CodeRedeemed")).to.not.be
        .undefined;

      expect(await token.balanceOf(friend1.address)).to.equal(10);

      await time.increase(callInterval);

      const code2 = getRandomNumber(6);

      const codeHash2 = getHash(
        codeCreator.address,
        code2,
        network.chainId,
        redeemCodeFaucet.address
      );

      await redeemCodeFaucet
        .connect(codeCreator)
        .addRedeemCode(codeHash2, 10, current + 100);

      tx = await redeemCodeFaucet.connect(friend1).redeem(code2);
      receipt = await tx.wait();

      expect(receipt.events?.find((e) => e.event === "CodeRedeemed")).to.not.be
        .undefined;

      await time.increase(callInterval);

      const code3 = getRandomNumber(6);

      const codeHash3 = getHash(
        codeCreator.address,
        code3,
        network.chainId,
        redeemCodeFaucet.address
      );

      await redeemCodeFaucet
        .connect(codeCreator)
        .addRedeemCode(codeHash3, 10, current + 100);

      await expect(
        redeemCodeFaucet.connect(friend1).redeem(code3)
      ).to.be.revertedWith("RedeemCodeFaucet: insufficient balance");
    });
  });
});
