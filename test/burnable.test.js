const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

describe("UpgradeableBurnableCommunityToken", function () {
  let erc20Token,
    UpgradeableCommunityToken,
    owner,
    addr1,
    addr2,
    minter1,
    minter2,
    receiver,
    anotherAccount;

  beforeEach(async function () {
    UpgradeableCommunityToken = await ethers.getContractFactory(
      "UpgradeableBurnableCommunityToken"
    );
    [owner, addr1, addr2, minter1, minter2, receiver, anotherAccount] =
      await ethers.getSigners();
    erc20Token = await upgrades.deployProxy(
      UpgradeableCommunityToken,
      [[minter1.address], "Regens Unite Token", "RGN"],
      { kind: "uups", initializer: "initialize", constructorArgs: [6] }
    );
  });

  describe("Burning", function () {
    it("Minter role should be able to burn any token", async function () {
      await erc20Token
        .connect(minter1)
        .mint(addr1.address, 100, "Initial mint");
      await erc20Token.connect(minter1).burnFrom(addr1.address, 50);
      const addr1Balance = await erc20Token.balanceOf(addr1.address);
      expect(addr1Balance).to.equal(50);
    });
    it("Non minter role cannot burn any token", async function () {
      await erc20Token
        .connect(minter1)
        .mint(addr1.address, 100, "Initial mint");
      await expect(erc20Token.connect(addr2).burnFrom(addr1.address, 50)).to.be
        .reverted;
      const addr1Balance = await erc20Token.balanceOf(addr1.address);
      expect(addr1Balance).to.equal(100);
    });
  });

  // Additional test cases for upgrade functionality will be added here based on how you implement upgradeability in your contract
});
