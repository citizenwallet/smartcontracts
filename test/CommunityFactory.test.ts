import { expect } from "chai";
import EntryPointArtifact from "@account-abstraction/contracts/artifacts/EntryPoint.json";
import "@nomicfoundation/hardhat-toolbox";
import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { config } from "dotenv";
import { ethers, upgrades } from "hardhat";

describe("CommunityFactory", function () {
  config();

  async function deployCommunityFactoryFixture() {
    const [
      owner,
      friend1,
      friend2,
      friend3,
      sponsor,
      sponsor2,
      vendor1,
      vendor2,
      vendor3,
    ] = await ethers.getSigners();

    const EntryPointContract = await ethers.getContractFactory(
      EntryPointArtifact.abi,
      EntryPointArtifact.bytecode,
      owner
    );
    const entrypoint = await EntryPointContract.deploy();

    const CommunityFactoryContract = await ethers.getContractFactory(
      "CommunityFactory",
      {
        signer: owner,
      }
    );

    const communityFactory = await CommunityFactoryContract.deploy(
      entrypoint.address
    );

    const TokenContract = await ethers.getContractFactory(
      "UpgradeableCommunityToken",
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

    return {
      owner,
      friend1,
      friend2,
      friend3,
      sponsor,
      sponsor2,
      vendor1,
      vendor2,
      vendor3,
      communityFactory,
      token,
    };
  }

  async function deployCommunityFixture() {
    const {
      owner,
      communityFactory,
      token,
      friend1,
      friend2,
      vendor1,
      vendor2,
    } = await loadFixture(deployCommunityFactoryFixture);

    const tx = await communityFactory.create(friend1.address, token.address, 0);

    await tx.wait();

    const [
      communityAddress,
      paymasterAddress,
      accountFactoryAddress,
      profileAddress,
    ] = await communityFactory.get(friend1.address, token.address, 0);

    const community = await ethers.getContractAt(
      "TokenEntryPoint",
      communityAddress
    );

    const paymaster = await ethers.getContractAt("Paymaster", paymasterAddress);

    const accountFactory = await ethers.getContractAt(
      "AccountFactory",
      accountFactoryAddress
    );

    const profile = await ethers.getContractAt("Profile", profileAddress);

    return {
      owner,
      token,
      communityFactory,
      community,
      paymaster,
      accountFactory,
      profile,
      friend1,
      friend2,
      vendor1,
      vendor2,
    };
  }

  describe("Create Community", function () {
    it("Should be owned by the owner that was set", async function () {
      const { friend1, community, paymaster } = await loadFixture(
        deployCommunityFixture
      );

      expect(await community.owner()).to.equal(friend1.address);
      expect(await paymaster.owner()).to.equal(friend1.address);
    });

    it("Should be sponsored by what was set", async function () {
      const { friend1, paymaster } = await loadFixture(deployCommunityFixture);

      expect(await paymaster.sponsor()).to.equal(friend1.address);
    });

    it("Owner should be able to modify the sponsor", async function () {
      const { friend1, friend2, paymaster } = await loadFixture(
        deployCommunityFixture
      );

      expect(await paymaster.sponsor()).to.equal(friend1.address);

      await paymaster.connect(friend1).updateSponsor(friend2.address);

      expect(await paymaster.sponsor()).to.equal(friend2.address);
    });

    it("Random account should not be able to modify the sponsor", async function () {
      const { friend1, friend2, paymaster } = await loadFixture(
        deployCommunityFixture
      );

      expect(await paymaster.sponsor()).to.equal(friend1.address);

      await expect(
        paymaster.connect(friend2).updateSponsor(friend2.address)
      ).to.be.revertedWith("Ownable: caller is not the owner");

      expect(await paymaster.sponsor()).to.equal(friend1.address);
    });

    it("Should create different addresses for different salts", async function () {
      const {
        friend1,
        token,
        communityFactory,
        community,
        paymaster,
        accountFactory,
        profile,
      } = await loadFixture(deployCommunityFixture);

      const [
        communityAddress2,
        paymasterAddress2,
        accountFactoryAddress2,
        profileAddress2,
      ] = await communityFactory.get(friend1.address, token.address, 0);

      expect(community.address).to.equal(communityAddress2);
      expect(paymaster.address).to.equal(paymasterAddress2);
      expect(accountFactory.address).to.equal(accountFactoryAddress2);
      expect(profile.address).to.equal(profileAddress2);

      const [
        communityAddress3,
        paymasterAddress3,
        accountFactoryAddress3,
        profileAddress3,
      ] = await communityFactory.get(friend1.address, token.address, 1);

      expect(community.address).to.not.equal(communityAddress3);
      expect(paymaster.address).to.not.equal(paymasterAddress3);
      expect(accountFactory.address).to.not.equal(accountFactoryAddress3);
      expect(profile.address).to.not.equal(profileAddress3);
    });

    it("Should return different addresses for different owners", async function () {
      const {
        friend2,
        token,
        communityFactory,
        community,
        paymaster,
        accountFactory,
        profile,
      } = await loadFixture(deployCommunityFixture);

      const [
        communityAddress2,
        paymasterAddress2,
        accountFactoryAddress2,
        profileAddress2,
      ] = await communityFactory.get(friend2.address, token.address, 0);

      expect(community.address).to.not.equal(communityAddress2);
      expect(paymaster.address).to.not.equal(paymasterAddress2);
      expect(accountFactory.address).to.not.equal(accountFactoryAddress2);
      expect(profile.address).to.not.equal(profileAddress2);

      const [
        communityAddress3,
        paymasterAddress3,
        accountFactoryAddress3,
        profileAddress3,
      ] = await communityFactory.get(friend2.address, token.address, 0);

      expect(communityAddress2).to.equal(communityAddress3);
      expect(paymasterAddress2).to.equal(paymasterAddress3);
      expect(accountFactoryAddress2).to.equal(accountFactoryAddress3);
      expect(profileAddress2).to.equal(profileAddress3);
    });

    it("Creating multiple times should not affect other communities", async function () {
      const {
        friend1,
        friend2,
        token,
        communityFactory,
        community,
        paymaster,
        accountFactory,
        profile,
      } = await loadFixture(deployCommunityFixture);

      let tx = await communityFactory.create(friend2.address, token.address, 0);

      await tx.wait();

      tx = await communityFactory.create(friend2.address, token.address, 0);

      await tx.wait();

      tx = await communityFactory.create(friend2.address, token.address, 1);

      await tx.wait();

      const [
        communityAddress3,
        paymasterAddress3,
        accountFactoryAddress3,
        profileAddress3,
      ] = await communityFactory.get(friend1.address, token.address, 1);

      expect(community.address).to.not.equal(communityAddress3);
      expect(paymaster.address).to.not.equal(paymasterAddress3);
      expect(accountFactory.address).to.not.equal(accountFactoryAddress3);
      expect(profile.address).to.not.equal(profileAddress3);
    });
  });
});
