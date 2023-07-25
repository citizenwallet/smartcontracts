import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers, upgrades } from "hardhat";

const friend1UsernameA = ethers.utils.formatBytes32String("friend1a");
const friend1UsernameB = ethers.utils.formatBytes32String("friend1b");
const friend2Username = ethers.utils.formatBytes32String("friend2");

// sould bound ERC721 with profile metadata uri
describe("Profile", function () {
  async function deployProfileFixture() {
    const [owner, friend1, friend2] = await ethers.getSigners();

    const ProfileContract = await ethers.getContractFactory("Profile");
    const profile = await upgrades.deployProxy(ProfileContract, [], {
      kind: "uups",
      initializer: "initialize",
    });

    return { profile, owner, friend1, friend2 };
  }

  describe("Deployment", function () {
    it("Should set the right owner", async function () {
      const { profile, owner } = await loadFixture(deployProfileFixture);

      expect(await profile.owner()).to.equal(owner.address);
    });
  });

  describe("Setting", function () {
    it("Should set the profile metadata uri", async function () {
      const { profile, owner, friend1, friend2 } = await loadFixture(
        deployProfileFixture
      );
      await profile
        .connect(friend1)
        .set(friend1.address, friend1UsernameA, "https://test.com");
      expect(await profile.get(friend1.address)).to.equal("https://test.com");
      expect(await profile.getFromUsername(friend1UsernameA)).to.equal(
        "https://test.com"
      );

      await profile
        .connect(friend2)
        .set(friend2.address, friend2Username, "https://test2.com");
      expect(await profile.get(friend2.address)).to.equal("https://test2.com");
      expect(await profile.getFromUsername(friend2Username)).to.equal(
        "https://test2.com"
      );

      await profile
        .connect(friend1)
        .set(friend1.address, friend1UsernameA, "https://testing.com");
      expect(await profile.get(friend1.address)).to.equal(
        "https://testing.com"
      );
      expect(await profile.getFromUsername(friend1UsernameA)).to.equal(
        "https://testing.com"
      );
    });

    it("Should only allow the profile owner to set", async function () {
      const { profile, owner, friend1, friend2 } = await loadFixture(
        deployProfileFixture
      );
      await expect(
        profile
          .connect(friend2)
          .set(friend1.address, friend1UsernameA, "https://test.com")
      ).to.be.revertedWith("Only the profile owner can set it.");

      await expect(
        profile
          .connect(owner)
          .set(friend1.address, friend1UsernameA, "https://test.com")
      ).to.be.revertedWith("Only the profile owner can set it.");
    });

    it("Should set the profile metadata uri and mint if not already minted", async function () {
      const { profile, owner, friend1, friend2 } = await loadFixture(
        deployProfileFixture
      );
      expect(await profile.balanceOf(friend1.address)).to.equal(0);
      await profile
        .connect(friend1)
        .set(friend1.address, friend1UsernameA, "https://test.com");
      expect(await profile.balanceOf(friend1.address)).to.equal(1);

      expect(await profile.get(friend1.address)).to.equal("https://test.com");

      expect(await profile.balanceOf(friend2.address)).to.equal(0);
      await profile
        .connect(friend2)
        .set(friend2.address, friend2Username, "https://test2.com");

      expect(await profile.balanceOf(friend2.address)).to.equal(1);
      expect(await profile.get(friend2.address)).to.equal("https://test2.com");
    });

    it("Should not allow transfers", async function () {
      const { profile, owner, friend1, friend2 } = await loadFixture(
        deployProfileFixture
      );
      await profile
        .connect(friend1)
        .set(friend1.address, friend1UsernameA, "https://test.com");
      await expect(
        profile
          .connect(friend1)
          .transferFrom(
            friend1.address,
            friend2.address,
            ethers.BigNumber.from(friend1.address)
          )
      ).to.be.revertedWith(
        "This a Soulbound token. It cannot be transferred. It can only be burned by the token owner."
      );
    });

    it("Should only allow the token owner to burn", async function () {
      const { profile, owner, friend1, friend2 } = await loadFixture(
        deployProfileFixture
      );
      await profile
        .connect(friend1)
        .set(friend1.address, friend1UsernameA, "https://test.com");
      await expect(
        profile.connect(friend2).burn(ethers.BigNumber.from(friend1.address))
      ).to.be.revertedWith(
        "Only the owner of the token or profile can burn it."
      );

      await profile.connect(owner).burn(ethers.BigNumber.from(friend1.address));
      expect(await profile.balanceOf(friend1.address)).to.equal(0);
    });

    it("Should only allow the profile owner to burn", async function () {
      const { profile, owner, friend1, friend2 } = await loadFixture(
        deployProfileFixture
      );
      await profile
        .connect(friend1)
        .set(friend1.address, friend1UsernameA, "https://test.com");
      await expect(
        profile.connect(friend2).burn(ethers.BigNumber.from(friend1.address))
      ).to.be.revertedWith(
        "Only the owner of the token or profile can burn it."
      );

      await profile
        .connect(friend1)
        .burn(ethers.BigNumber.from(friend1.address));
      expect(await profile.balanceOf(friend1.address)).to.equal(0);
    });

    it("Should allow setting the profile again after burning", async function () {
      const { profile, owner, friend1, friend2 } = await loadFixture(
        deployProfileFixture
      );
      expect(await profile.balanceOf(friend1.address)).to.equal(0);
      await profile
        .connect(friend1)
        .set(friend1.address, friend1UsernameA, "https://test.com");
      expect(await profile.balanceOf(friend1.address)).to.equal(1);
      await profile
        .connect(friend1)
        .burn(ethers.BigNumber.from(friend1.address));
      expect(await profile.balanceOf(friend1.address)).to.equal(0);
      await profile
        .connect(friend1)
        .set(friend1.address, friend1UsernameA, "https://test.com");
      expect(await profile.balanceOf(friend1.address)).to.equal(1);
    });

    it("Should not allow setting a username that belongs to someone else", async function () {
      const { profile, owner, friend1, friend2 } = await loadFixture(
        deployProfileFixture
      );

      await profile
        .connect(friend1)
        .set(friend1.address, friend1UsernameA, "https://test.com");

      await expect(
        profile
          .connect(friend2)
          .set(friend2.address, friend1UsernameA, "https://test2.com")
      ).to.be.revertedWith("This username is already taken.");
    });

    // set a username, then set another one, the value should be the new one
    it("Should replace the username to the new one when set a second time", async function () {
      const { profile, owner, friend1, friend2 } = await loadFixture(
        deployProfileFixture
      );

      await profile
        .connect(friend1)
        .set(friend1.address, friend1UsernameA, "https://test.com");

      await profile
        .connect(friend1)
        .set(friend1.address, friend1UsernameB, "https://test.com");

      expect(await profile.getFromUsername(friend1UsernameB)).to.equal(
        "https://test.com"
      );

      await expect(
        profile.getFromUsername(friend1UsernameA)
      ).to.be.revertedWith("This username does not exist.");
    });

    // set a username, then set another one, the old username should be available again
    it("Should replace the username to the new one and the old one should be available again", async function () {
      const { profile, owner, friend1, friend2 } = await loadFixture(
        deployProfileFixture
      );

      await profile
        .connect(friend1)
        .set(friend1.address, friend1UsernameA, "https://test.com");

      await expect(
        profile
          .connect(friend2)
          .set(friend2.address, friend1UsernameA, "https://test2.com")
      ).to.be.revertedWith("This username is already taken.");

      await profile
        .connect(friend1)
        .set(friend1.address, friend1UsernameB, "https://test.com");

      await profile
        .connect(friend2)
        .set(friend2.address, friend1UsernameA, "https://test2.com");

      expect(await profile.getFromUsername(friend1UsernameB)).to.equal(
        "https://test.com"
      );

      expect(await profile.getFromUsername(friend1UsernameA)).to.equal(
        "https://test2.com"
      );
    });
  });
});
