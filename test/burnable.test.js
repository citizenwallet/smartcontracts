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
      "UpgradeableCommunityToken"
    );
    [owner, addr1, addr2, minter1, minter2, receiver, anotherAccount] =
      await ethers.getSigners();
    erc20Token = await upgrades.deployProxy(
      UpgradeableCommunityToken,
      [[minter1.address], "Regens Unite Token", "RGN"],
      { kind: "uups", initializer: "initialize" }
    );
  });

  describe("upgrading", function () {
    it("Should be able to upgrade the contract to add the burnFrom function", async function () {
      const V2Factory = await ethers.getContractFactory(
        "UpgradeableBurnableCommunityToken"
      );

      // Upgrading the existing proxy to the new implementation
      await upgrades.upgradeProxy(erc20Token.address, V2Factory, {
        kind: "uups",
      });

      // New function should be callable on the proxy address
      await expect(erc20Token.burnFrom()).to.be.ok;
      // We deploy the first version of the contract
      expect(erc20Token.connect(minter1).burnFrom(addr1.address, 50)).to.be
        .reverted;

      // // Existing functionalities should still work
      await erc20Token.connect(minter1).mint(addr1.address, 100);

      let ownerBalance = await erc20Token.balanceOf(addr1.address);
      expect(ownerBalance).to.equal(100);

      await erc20Token.connect(minter1).burnFrom(addr1.address, 50);

      ownerBalance = await erc20Token.balanceOf(addr1.address);
      expect(ownerBalance).to.equal(50);

      expect(erc20Token.connect(addr1).burnFrom(addr1.address, 50)).to.be
        .reverted;
    });

    it("Should be able to upgrade the contract to keep the same storage values", async function () {
      expect(await erc20Token.name()).to.equal("Regens Unite Token");
      expect(await erc20Token.symbol()).to.equal("RGN");
      expect(await erc20Token.decimals()).to.equal(6);

      expect(await erc20Token.owner()).to.equal(owner.address);

      await erc20Token.connect(minter1).mint(addr1.address, 100);

      let ownerBalance = await erc20Token.balanceOf(addr1.address);
      expect(ownerBalance).to.equal(100);

      // We deploy the first version of the contract
      const V2Factory = await ethers.getContractFactory(
        "UpgradeableBurnableCommunityToken"
      );

      // Upgrading the existing proxy to the new implementation
      await upgrades.upgradeProxy(erc20Token.address, V2Factory, {
        kind: "uups",
      });

      // New function should be callable on the proxy address
      expect(await erc20Token.name()).to.equal("Regens Unite Token");
      expect(await erc20Token.symbol()).to.equal("RGN");
      expect(await erc20Token.decimals()).to.equal(6);
      expect(await erc20Token.owner()).to.equal(owner.address);

      ownerBalance = await erc20Token.balanceOf(addr1.address);
      expect(ownerBalance).to.equal(100);
    });
  });
  describe("Burning", function () {
    beforeEach(async function () {
      const V2Factory = await ethers.getContractFactory(
        "UpgradeableBurnableCommunityToken"
      );

      // Upgrading the existing proxy to the new implementation
      await upgrades.upgradeProxy(erc20Token.address, V2Factory, {
        kind: "uups",
      });
      // New function should be callable on the proxy address
      await expect(erc20Token.burnFrom()).to.be.ok;
    });

    it("Minter role should be able to burn any token", async function () {
      await erc20Token.connect(minter1).mint(addr1.address, 100);
      await erc20Token.connect(minter1).burnFrom(addr1.address, 50);
      const addr1Balance = await erc20Token.balanceOf(addr1.address);
      expect(addr1Balance).to.equal(50);
    });
    it("Non minter role cannot burn any token", async function () {
      await erc20Token.connect(minter1).mint(addr1.address, 100);
      await expect(erc20Token.connect(addr2).burnFrom(addr1.address, 50)).to.be
        .reverted;
      const addr1Balance = await erc20Token.balanceOf(addr1.address);
      expect(addr1Balance).to.equal(100);
    });
  });
});
// Additional test cases for upgrade functionality will be added here based on how you implement upgradeability in your contract
