const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("RegensUniteTokens", function () {
  let RegensUniteTokens, regensUniteTokens, admin1, admin2, user1, user2, minters;
  const amountRegenCoin = 1000;
  const amountGratitude = 1;
  const amountVoucher = 10;
  const ipfs_metadata = "ipfs://metadata.json";

  const MINTER_ROLE = ethers.utils.id("MINTER_ROLE");

  before(async function () {
    // Get the ContractFactory and Signers here.
    RegensUniteTokens = await ethers.getContractFactory("RegensUniteTokens");
    [admin1, admin2, user1, user2, ...minters] = await ethers.getSigners();

    // Deploy the contract and set minters
    regensUniteTokens = await RegensUniteTokens.deploy([admin1.address, admin2.address]);
  });

  describe("Deployment", function () {
    it("Should set the right admins", async function () {
      expect(await regensUniteTokens.hasRole(MINTER_ROLE, admin1.address)).to.equal(true);
      expect(await regensUniteTokens.hasRole(MINTER_ROLE, admin2.address)).to.equal(true);
      expect(await regensUniteTokens.hasRole(MINTER_ROLE, user1.address)).to.equal(false);
    });
  });

  describe("Minting", function () {
    it("Users cannot mint Regen Coins", async function () {
      await expect(regensUniteTokens.connect(user1).mintCoin(user1.address, amountRegenCoin, [])).to.be.reverted;
    });

    it("Minters can mint Regen Coins and transfer to user1", async function () {
      await regensUniteTokens.connect(admin1).mintCoin(user1.address, amountRegenCoin, []);
      expect(await regensUniteTokens.balanceOf(user1.address, 1)).to.equal(amountRegenCoin);
    });

    it("User1 can send 500 Regen Coins to user2", async function () {
      await regensUniteTokens.connect(user1).safeTransferFrom(user1.address, user2.address, 1, 500, []);
      expect(await regensUniteTokens.balanceOf(user1.address, 1)).to.equal(500);
      expect(await regensUniteTokens.balanceOf(user2.address, 1)).to.equal(500);
    });

    it("User1 can mint a gratitude token and transfer it to user2", async function () {
      await regensUniteTokens.connect(user1).mintGratitude(user2.address);
      expect(await regensUniteTokens.balanceOf(user2.address, 2)).to.equal(amountGratitude);
    });

    it("User1 can mint a custom voucher and send one to user2", async function () {
      // Assume the voucher has an ID of 10001
      await regensUniteTokens.connect(user1).mintCustomVoucher(user2.address, 10001, amountVoucher, [], ipfs_metadata);
      expect(await regensUniteTokens.balanceOf(user2.address, 10001)).to.equal(amountVoucher);
    });

    it("User2 can send custom voucher back to User1", async function () {
      await regensUniteTokens.connect(user2).safeTransferFrom(user2.address, user1.address, 10001, 1, []);
      expect(await regensUniteTokens.balanceOf(user1.address, 10001)).to.equal(1);
      expect(await regensUniteTokens.balanceOf(user2.address, 10001)).to.equal(amountVoucher - 1);
    });

    it("User1 can burn voucher", async function () {
      await regensUniteTokens.connect(user1).burn(user1.address, 10001, 1);
      expect(await regensUniteTokens.balanceOf(user1.address, 10001)).to.equal(0);
    });
  });
});
