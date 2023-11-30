const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

describe("ERC20TokenWithUpgrades", function () {
  let erc20Token,
    ERC20TokenWithUpgrades,
    owner,
    addr1,
    addr2,
    minter1,
    minter2,
    receiver,
    anotherAccount;

  beforeEach(async function () {
    ERC20TokenWithUpgrades = await ethers.getContractFactory(
      "ERC20TokenWithUpgrades"
    );
    [owner, addr1, addr2, minter1, minter2, receiver, anotherAccount] =
      await ethers.getSigners();
    erc20Token = await upgrades.deployProxy(
      ERC20TokenWithUpgrades,
      [[minter1.address], "Regens Unite Token", "RGN"],
      { kind: "uups", initializer: "initialize", constructorArgs: [6] }
    );
  });

  describe("Deployment", function () {
    it("Should set the right owner", async function () {
      expect(await erc20Token.owner()).to.equal(owner.address);
    });

    it("Should set initial minter correctly", async function () {
      expect(
        await erc20Token.hasRole(
          await erc20Token.MINTER_ROLE(),
          minter1.address
        )
      ).to.equal(true);
    });
  });

  describe("Minting", function () {
    it("Should allow minters to mint new tokens", async function () {
      await erc20Token
        .connect(minter1)
        .mint(receiver.address, 100, "For services rendered");
      const receiverBalance = await erc20Token.balanceOf(receiver.address);
      expect(receiverBalance).to.equal(100);
    });

    it("Should emit a Minted event", async function () {
      await expect(
        erc20Token
          .connect(minter1)
          .mint(receiver.address, 100, "For services rendered")
      )
        .to.emit(erc20Token, "Minted")
        .withArgs(receiver.address, 100, "For services rendered");
    });

    it("Should allow owner to assign minter role to another account", async function () {
      await erc20Token.grantRole(
        await erc20Token.MINTER_ROLE(),
        minter2.address
      );
      expect(
        await erc20Token.hasRole(
          await erc20Token.MINTER_ROLE(),
          minter2.address
        )
      ).to.equal(true);
    });

    it("Should allow new minter to mint tokens", async function () {
      await erc20Token.grantRole(
        await erc20Token.MINTER_ROLE(),
        minter2.address
      );
      await erc20Token
        .connect(minter2)
        .mint(receiver.address, 100, "For additional services rendered");
      const receiverBalance = await erc20Token.balanceOf(receiver.address);
      expect(receiverBalance).to.equal(100);
    });

    it("Should not allow non-minters to mint new tokens", async function () {
      await expect(
        erc20Token
          .connect(addr1)
          .mint(receiver.address, 100, "For services rendered")
      ).to.be.revertedWith("Must have minter role to mint");
    });
  });

  describe("Burning", function () {
    it("Should allow token holders to burn their tokens", async function () {
      await erc20Token
        .connect(minter1)
        .mint(addr1.address, 100, "Initial mint");
      await erc20Token.connect(addr1).burn(50);
      const addr1Balance = await erc20Token.balanceOf(addr1.address);
      expect(addr1Balance).to.equal(50);
    });
  });

  describe("Pausing", function () {
    it("Should allow pausers to pause and unpause the contract", async function () {
      await erc20Token.pause();
      expect(await erc20Token.paused()).to.equal(true);

      await erc20Token.unpause();
      expect(await erc20Token.paused()).to.equal(false);
    });

    it("Should not allow non-pausers to pause and unpause the contract", async function () {
      await expect(erc20Token.connect(addr1).pause()).to.be.reverted;
      await expect(erc20Token.connect(addr1).unpause()).to.be.reverted;
    });
  });

  // Additional test cases for upgrade functionality will be added here based on how you implement upgradeability in your contract
});
