import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers, upgrades } from "hardhat";

const getHash = (
  from: string,
  amount: number,
  validUntil: number,
  validFrom: number,
  sequence: number,
  chainId: number,
  contract: string
) => {
  return ethers.utils.solidityKeccak256(
    ["address", "uint256", "uint48", "uint48", "uint256", "uint256", "address"],
    [from, amount, validUntil, validFrom, sequence, chainId, contract]
  );
};

// sould bound ERC721 with tokeniou metadata uri
describe("ERC20IOU", function () {
  async function deployTokenIOUFixture() {
    const [owner, friend1, friend2] = await ethers.getSigners();

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

    const TokenIOUContract = await ethers.getContractFactory("ERC20IOU");
    const tokeniou = await upgrades.deployProxy(
      TokenIOUContract,
      [token.address],
      {
        kind: "uups",
        initializer: "initialize",
      }
    );

    const tokeniou2 = await upgrades.deployProxy(
      TokenIOUContract,
      [token.address],
      {
        kind: "uups",
        initializer: "initialize",
      }
    );

    const network = await ethers.provider.getNetwork();

    const current = await time.latest();

    return {
      network,
      current,
      tokeniou,
      tokeniou2,
      token,
      owner,
      friend1,
      friend2,
    };
  }

  async function redeemFixture() {
    const { current, tokeniou, tokeniou2, token, friend1, friend2 } =
      await loadFixture(deployTokenIOUFixture);

    const validUntil = current + 2592000; // 30 days
    const validAfter = current;

    const hash = await tokeniou.getHash(
      friend2.address,
      1,
      validUntil,
      validAfter,
      0
    );

    await token.mint(friend2.address, 1, "hello");

    await token.connect(friend2).approve(tokeniou.address, 1);

    const signature = await friend2.signMessage(ethers.utils.arrayify(hash));

    console.log(
      [
        friend2.address,
        ethers.utils.hexValue(1),
        ethers.utils.hexValue(validUntil),
        ethers.utils.hexValue(validAfter),
        ethers.utils.hexValue(0),
        signature,
      ].join("|")
    );

    return {
      tokeniou,
      tokeniou2,
      token,
      friend1,
      friend2,
      current,
      validUntil,
      validAfter,
      signature,
    };
  }

  describe("Deployment", function () {
    it("Should set the right owner", async function () {
      const { tokeniou, owner } = await loadFixture(deployTokenIOUFixture);

      expect(await tokeniou.owner()).to.equal(owner.address);
    });
  });

  describe("Hash", function () {
    it("Should generate the right hash", async function () {
      const { network, current, tokeniou, friend1 } = await loadFixture(
        deployTokenIOUFixture
      );

      const validUntil = current + 2592000; // 30 days
      const validAfter = current;

      const hash = await tokeniou.getHash(
        friend1.address,
        1,
        validUntil,
        validAfter,
        0
      );

      const localHash = getHash(
        friend1.address,
        1,
        validUntil,
        validAfter,
        0,
        network.chainId,
        tokeniou.address
      );

      expect(hash).to.equal(localHash);
    });
  });

  describe("Redeem", async function () {
    it("Should allow redeeming a valid IOU", async function () {
      const {
        tokeniou,
        friend2,
        token,
        friend1,
        validUntil,
        validAfter,
        signature,
      } = await loadFixture(redeemFixture);

      expect(await token.balanceOf(friend1.address)).to.equal(0);

      await tokeniou
        .connect(friend1)
        .redeem(friend2.address, 1, validUntil, validAfter, 0, signature);

      expect(await token.balanceOf(friend1.address)).to.equal(1);
    });

    it("Should not allow redeeming again", async function () {
      const { tokeniou, friend1, friend2, validUntil, validAfter, signature } =
        await loadFixture(redeemFixture);

      await tokeniou
        .connect(friend1)
        .redeem(friend2.address, 1, validUntil, validAfter, 0, signature);

      await expect(
        tokeniou
          .connect(friend1)
          .redeem(friend2.address, 1, validUntil, validAfter, 0, signature)
      ).to.be.revertedWith("ERC20IOU: already redeemed");
    });

    it("Should not allow someone to redeem with someone else's signature", async function () {
      const { tokeniou, friend1, validUntil, validAfter, signature } =
        await loadFixture(redeemFixture);

      await expect(
        tokeniou
          .connect(friend1)
          .redeem(friend1.address, 1, validUntil, validAfter, 0, signature)
      ).to.be.revertedWith("ERC20IOU: invalid signature");
    });

    it("Should not allow tampering with validity", async function () {
      const { tokeniou, friend1, friend2, current, validUntil, signature } =
        await loadFixture(redeemFixture);

      const validAfter = current - 2;

      await expect(
        tokeniou
          .connect(friend1)
          .redeem(friend2.address, 1, validUntil, validAfter, 0, signature)
      ).to.be.revertedWith("ERC20IOU: invalid signature");
    });

    it("Should not allow redeeming before validity", async function () {
      const { current, tokeniou, friend1, friend2 } = await loadFixture(
        deployTokenIOUFixture
      );

      const validUntil = current + 2592000; // 30 days
      const validAfter = current + 100;

      const hash = await tokeniou.getHash(
        friend1.address,
        1,
        validUntil,
        validAfter,
        0
      );

      const signature = await friend1.signMessage(ethers.utils.arrayify(hash));

      await expect(
        tokeniou
          .connect(friend1)
          .redeem(friend2.address, 1, validUntil, validAfter, 0, signature)
      ).to.be.revertedWith("ERC20IOU: expired or not valid yet");
    });

    it("Should not allow redeeming after validity", async function () {
      const { current, tokeniou, friend1, friend2 } = await loadFixture(
        deployTokenIOUFixture
      );

      const validUntil = current - 2; // 30 days
      const validAfter = current - 100;

      const hash = await tokeniou.getHash(
        friend1.address,
        1,
        validUntil,
        validAfter,
        0
      );

      const signature = await friend1.signMessage(ethers.utils.arrayify(hash));

      await expect(
        tokeniou
          .connect(friend1)
          .redeem(friend2.address, 1, validUntil, validAfter, 0, signature)
      ).to.be.revertedWith("ERC20IOU: expired or not valid yet");
    });

    it("Should not allow redeeming on another contract", async function () {
      const { tokeniou2, friend1, friend2, validUntil, validAfter, signature } =
        await loadFixture(redeemFixture);

      await expect(
        tokeniou2
          .connect(friend1)
          .redeem(friend2.address, 1, validUntil, validAfter, 0, signature)
      ).to.be.revertedWith("ERC20IOU: invalid signature");
    });
  });
});
