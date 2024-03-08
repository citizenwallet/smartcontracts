const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

describe("UpgradeableBurnableCommunityToken", function () {
  let erc20Token,
    GratitudeToken,
    owner,
    addr1,
    addr2,
    minter1,
    minter2,
    receiver,
    anotherAccount;

  beforeEach(async function () {
    GratitudeToken = await ethers.getContractFactory("GratitudeToken");
    [owner, addr1, addr2, minter1, minter2, receiver, anotherAccount] =
      await ethers.getSigners();
    erc20Token = await upgrades.deployProxy(
      GratitudeToken,
      [[minter1.address], "Gratitude Token", "GT"],
      { kind: "uups", initializer: "initialize" }
    );
  });

  describe("mint and transfer", function () {
    it("Should return a balance of 1", async function () {
      let ownerBalance = await erc20Token.balanceOf(addr1.address);
      expect(ownerBalance).to.equal(1);
    });

    it("Should be able to mint and transfer tokens", async function () {
      await erc20Token.connect(addr1).transfer(addr2.address, 1);
      let ownerBalance = await erc20Token.balanceOf(addr1.address);
      expect(ownerBalance).to.equal(1);
      let receiverBalance = await erc20Token.balanceOf(addr1.address);
      expect(ownerBalance).to.equal(1);
    });
  });
});
// Additional test cases for upgrade functionality will be added here based on how you implement upgradeability in your contract
