const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

describe("RegensUniteToken", function () {
  let regensUniteToken,
    RegensUniteToken,
    owner,
    addr1,
    addr2,
    minter1,
    minter2,
    receiver,
    anotherAccount;

  beforeEach(async function () {
    RegensUniteToken = await ethers.getContractFactory("RegensUniteToken");
    [owner, addr1, addr2, minter1, minter2, receiver, anotherAccount] =
      await ethers.getSigners();
    regensUniteToken = await upgrades.deployProxy(
      RegensUniteToken,
      [[minter1.address]],
      { initializer: "initialize" }
    );
  });

  describe("Deployment", function () {
    it("Should set the right owner", async function () {
      expect(await regensUniteToken.owner()).to.equal(owner.address);
    });

    it("Should set initial minter correctly", async function () {
      expect(
        await regensUniteToken.hasRole(
          await regensUniteToken.MINTER_ROLE(),
          minter1.address
        )
      ).to.equal(true);
    });
  });

  describe("Minting", function () {
    it("Should allow minters to mint new tokens", async function () {
      await regensUniteToken
        .connect(minter1)
        .mint(receiver.address, 100, "For services rendered");
      const receiverBalance = await regensUniteToken.balanceOf(
        receiver.address
      );
      expect(receiverBalance).to.equal(100);
    });

    it("Should emit a Minted event", async function () {
      await expect(
        regensUniteToken
          .connect(minter1)
          .mint(receiver.address, 100, "For services rendered")
      )
        .to.emit(regensUniteToken, "Minted")
        .withArgs(receiver.address, 100, "For services rendered");
    });

    it("Should allow owner to assign minter role to another account", async function () {
      await regensUniteToken.grantRole(
        await regensUniteToken.MINTER_ROLE(),
        minter2.address
      );
      expect(
        await regensUniteToken.hasRole(
          await regensUniteToken.MINTER_ROLE(),
          minter2.address
        )
      ).to.equal(true);
    });

    it("Should allow new minter to mint tokens", async function () {
      await regensUniteToken.grantRole(
        await regensUniteToken.MINTER_ROLE(),
        minter2.address
      );
      await regensUniteToken
        .connect(minter2)
        .mint(receiver.address, 100, "For additional services rendered");
      const receiverBalance = await regensUniteToken.balanceOf(
        receiver.address
      );
      expect(receiverBalance).to.equal(100);
    });

    it("Should not allow non-minters to mint new tokens", async function () {
      await expect(
        regensUniteToken
          .connect(addr1)
          .mint(receiver.address, 100, "For services rendered")
      ).to.be.revertedWith("Must have minter role to mint");
    });
  });

  describe("Burning", function () {
    it("Should allow token holders to burn their tokens", async function () {
      await regensUniteToken
        .connect(minter1)
        .mint(addr1.address, 100, "Initial mint");
      await regensUniteToken.connect(addr1).burn(50);
      const addr1Balance = await regensUniteToken.balanceOf(addr1.address);
      expect(addr1Balance).to.equal(50);
    });
  });

  describe("Pausing", function () {
    it("Should allow pausers to pause and unpause the contract", async function () {
      await regensUniteToken.pause();
      expect(await regensUniteToken.paused()).to.equal(true);

      await regensUniteToken.unpause();
      expect(await regensUniteToken.paused()).to.equal(false);
    });

    it("Should not allow non-pausers to pause and unpause the contract", async function () {
      await expect(regensUniteToken.connect(addr1).pause()).to.be.reverted;
      await expect(regensUniteToken.connect(addr1).unpause()).to.be.reverted;
    });
  });

  // Additional test cases for upgrade functionality will be added here based on how you implement upgradeability in your contract
});
