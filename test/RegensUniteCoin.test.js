const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

describe("RegensUniteCoin", function () {
  let regensUniteCoin,
    RegensUniteCoin,
    regensUniteCoinV2Test,
    RegensUniteCoinV2Test,
    owner,
    addr1,
    addr2,
    minter1,
    minter2,
    receiver,
    anotherAccount;

  beforeEach(async function () {
    RegensUniteCoin = await ethers.getContractFactory("RegensUniteCoin");
    [owner, addr1, addr2, minter1, minter2, receiver, anotherAccount] =
      await ethers.getSigners();
    regensUniteCoin = await upgrades.deployProxy(
      RegensUniteCoin,
      [[minter1.address]],
      { initializer: "initialize" }
    );
  });

  describe("Deployment", function () {
    it("Should set the right owner", async function () {
      expect(await regensUniteCoin.owner()).to.equal(owner.address);
    });

    it("Should set initial minter correctly", async function () {
      expect(
        await regensUniteCoin.hasRole(
          await regensUniteCoin.MINTER_ROLE(),
          minter1.address
        )
      ).to.equal(true);
    });

    it("Should be able to upgrade the contract", async function () {
      // Assume RegensUniteCoinV2 is a new version of your contract with a new function called `newFunction`
      RegensUniteCoinV2Test = await ethers.getContractFactory(
        "RegensUniteCoinV2Test"
      );

      // Upgrading the existing proxy to the new implementation
      regensUniteCoinV2Test = await upgrades.upgradeProxy(
        regensUniteCoin.address,
        RegensUniteCoinV2Test
      );

      // New function should be callable on the proxy address
      await expect(regensUniteCoinV2Test.newFunction()).to.be.ok;

      // Existing functionalities should still work
      await regensUniteCoinV2Test
        .connect(minter1)
        .mint(owner.address, 100, "For services rendered in V2");
      const ownerBalance = await regensUniteCoinV2Test.balanceOf(owner.address);
      expect(ownerBalance).to.equal(100);
    });
  });

  describe("Minting", function () {
    it("Should allow minters to mint new tokens", async function () {
      await regensUniteCoin
        .connect(minter1)
        .mint(receiver.address, 100, "For services rendered");
      const receiverBalance = await regensUniteCoin.balanceOf(receiver.address);
      expect(receiverBalance).to.equal(100);
    });

    it("Should emit a Minted event", async function () {
      await expect(
        regensUniteCoin
          .connect(minter1)
          .mint(receiver.address, 100, "For services rendered")
      )
        .to.emit(regensUniteCoin, "Minted")
        .withArgs(receiver.address, 100, "For services rendered");
    });

    it("Should allow owner to assign minter role to another account", async function () {
      await regensUniteCoin.grantRole(
        await regensUniteCoin.MINTER_ROLE(),
        minter2.address
      );
      expect(
        await regensUniteCoin.hasRole(
          await regensUniteCoin.MINTER_ROLE(),
          minter2.address
        )
      ).to.equal(true);
    });

    it("Should allow new minter to mint tokens", async function () {
      await regensUniteCoin.grantRole(
        await regensUniteCoin.MINTER_ROLE(),
        minter2.address
      );
      await regensUniteCoin
        .connect(minter2)
        .mint(receiver.address, 100, "For additional services rendered");
      const receiverBalance = await regensUniteCoin.balanceOf(receiver.address);
      expect(receiverBalance).to.equal(100);
    });

    it("Should not allow non-minters to mint new tokens", async function () {
      await expect(
        regensUniteCoin
          .connect(addr1)
          .mint(receiver.address, 100, "For services rendered")
      ).to.be.revertedWith("Must have minter role to mint");
    });
  });

  describe("Burning", function () {
    it("Should allow token holders to burn their tokens", async function () {
      await regensUniteCoin
        .connect(minter1)
        .mint(addr1.address, 100, "Initial mint");
      await regensUniteCoin.connect(addr1).burn(50);
      const addr1Balance = await regensUniteCoin.balanceOf(addr1.address);
      expect(addr1Balance).to.equal(50);
    });
  });

  describe("Pausing", function () {
    it("Should allow pausers to pause and unpause the contract", async function () {
      await regensUniteCoin.pause();
      expect(await regensUniteCoin.paused()).to.equal(true);

      await regensUniteCoin.unpause();
      expect(await regensUniteCoin.paused()).to.equal(false);
    });

    it("Should not allow non-pausers to pause and unpause the contract", async function () {
      await expect(regensUniteCoin.connect(addr1).pause()).to.be.reverted;
      await expect(regensUniteCoin.connect(addr1).unpause()).to.be.reverted;
    });
  });

  // Additional test cases for upgrade functionality will be added here based on how you implement upgradeability in your contract
});
