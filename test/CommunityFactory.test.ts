import { expect } from "chai";
import EntryPointArtifact from "@account-abstraction/contracts/artifacts/EntryPoint.json";
import "@nomicfoundation/hardhat-toolbox";
import { loadFixture, time } from "@nomicfoundation/hardhat-network-helpers";
import { config } from "dotenv";
import { ethers, upgrades } from "hardhat";
import { randomUint256 } from "../scripts/functions/randomNumber";

describe("TokenEntryPointFactory", function () {
  config();

  async function deployTokenEntryPointFactoryFixture() {
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

    const TokenEntryPointFactoryContract = await ethers.getContractFactory(
      "TokenEntryPointFactory",
      {
        signer: owner,
      }
    );

    const tokenEntryPointFactory = await TokenEntryPointFactoryContract.deploy(
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
      tokenEntryPointFactory,
      token,
    };
  }

  async function deployCommunityFixture() {
    const {
      owner,
      tokenEntryPointFactory,
      token,
      friend1,
      friend2,
      vendor1,
      vendor2,
      sponsor,
      sponsor2,
    } = await loadFixture(deployTokenEntryPointFactoryFixture);

    const salt = randomUint256();

    const tx = await tokenEntryPointFactory.create(
      friend1.address,
      sponsor.address,
      [token.address],
      salt
    );

    await tx.wait();

    const [communityAddress, paymasterAddress] =
      await tokenEntryPointFactory.get(
        friend1.address,
        sponsor.address,
        [token.address],
        salt
      );

    const community = await ethers.getContractAt(
      "TokenEntryPoint",
      communityAddress
    );

    const paymaster = await ethers.getContractAt("Paymaster", paymasterAddress);

    // const accountFactory = await ethers.getContractAt(
    //   "AccountFactory",
    //   accountFactoryAddress
    // );

    // const profile = await ethers.getContractAt("Profile", profileAddress);

    return {
      owner,
      token,
      tokenEntryPointFactory,
      community,
      paymaster,
      //   accountFactory,
      //   profile,
      friend1,
      friend2,
      vendor1,
      vendor2,
      sponsor,
      sponsor2,
      salt,
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
      const { sponsor, paymaster } = await loadFixture(deployCommunityFixture);

      expect(await paymaster.sponsor()).to.equal(sponsor.address);
    });

    it("Owner should be able to modify the sponsor", async function () {
      const { friend1, paymaster, sponsor, sponsor2 } = await loadFixture(
        deployCommunityFixture
      );

      expect(await paymaster.sponsor()).to.equal(sponsor.address);

      await paymaster.connect(friend1).updateSponsor(sponsor2.address);

      expect(await paymaster.sponsor()).to.equal(sponsor2.address);
    });

    it("Random account should not be able to modify the sponsor", async function () {
      const { sponsor, friend2, paymaster } = await loadFixture(
        deployCommunityFixture
      );

      expect(await paymaster.sponsor()).to.equal(sponsor.address);

      await expect(
        paymaster.connect(friend2).updateSponsor(friend2.address)
      ).to.be.revertedWith("Ownable: caller is not the owner");

      expect(await paymaster.sponsor()).to.equal(sponsor.address);
    });

    it("Should create different addresses for different salts", async function () {
      const {
        friend1,
        token,
        tokenEntryPointFactory,
        community,
        paymaster,
        // accountFactory,
        // profile,
        sponsor,
        salt,
      } = await loadFixture(deployCommunityFixture);

      const [
        communityAddress2,
        paymasterAddress2,
        // accountFactoryAddress2,
        // profileAddress2,
      ] = await tokenEntryPointFactory.get(
        friend1.address,
        sponsor.address,
        [token.address],
        salt
      );

      expect(community.address).to.equal(communityAddress2);
      expect(paymaster.address).to.equal(paymasterAddress2);
      //   expect(accountFactory.address).to.equal(accountFactoryAddress2);
      //   expect(profile.address).to.equal(profileAddress2);

      const [
        communityAddress3,
        paymasterAddress3,
        // accountFactoryAddress3,
        // profileAddress3,
      ] = await tokenEntryPointFactory.get(
        friend1.address,
        sponsor.address,
        [token.address],
        1
      );

      expect(community.address).to.not.equal(communityAddress3);
      expect(paymaster.address).to.not.equal(paymasterAddress3);
      //   expect(accountFactory.address).to.not.equal(accountFactoryAddress3);
      //   expect(profile.address).to.not.equal(profileAddress3);
    });

    it("Should return different addresses for different owners", async function () {
      const {
        friend2,
        token,
        tokenEntryPointFactory,
        community,
        paymaster,
        // accountFactory,
        // profile,
        sponsor,
        salt,
      } = await loadFixture(deployCommunityFixture);

      const [
        communityAddress2,
        paymasterAddress2,
        // accountFactoryAddress2,
        // profileAddress2,
      ] = await tokenEntryPointFactory.get(
        friend2.address,
        sponsor.address,
        [token.address],
        salt
      );

      expect(community.address).to.not.equal(communityAddress2);
      expect(paymaster.address).to.not.equal(paymasterAddress2);
      //   expect(accountFactory.address).to.not.equal(accountFactoryAddress2);
      //   expect(profile.address).to.not.equal(profileAddress2);

      const [
        communityAddress3,
        paymasterAddress3,
        // accountFactoryAddress3,
        // profileAddress3,
      ] = await tokenEntryPointFactory.get(
        friend2.address,
        sponsor.address,
        [token.address],
        salt
      );

      expect(communityAddress2).to.equal(communityAddress3);
      expect(paymasterAddress2).to.equal(paymasterAddress3);
      //   expect(accountFactoryAddress2).to.equal(accountFactoryAddress3);
      //   expect(profileAddress2).to.equal(profileAddress3);
    });

    it("Creating multiple times should not affect other communities", async function () {
      const {
        friend1,
        friend2,
        token,
        tokenEntryPointFactory,
        community,
        paymaster,
        // accountFactory,
        // profile,
        sponsor,
        salt,
      } = await loadFixture(deployCommunityFixture);

      let tx = await tokenEntryPointFactory.create(
        friend2.address,
        sponsor.address,
        [token.address],
        salt
      );

      await tx.wait();

      tx = await tokenEntryPointFactory.create(
        friend2.address,
        sponsor.address,
        [token.address],
        salt
      );

      await tx.wait();

      tx = await tokenEntryPointFactory.create(
        friend2.address,
        sponsor.address,
        [token.address],
        1
      );

      await tx.wait();

      const [
        communityAddress3,
        paymasterAddress3,
        // accountFactoryAddress3,
        // profileAddress3,
      ] = await tokenEntryPointFactory.get(
        friend1.address,
        sponsor.address,
        [token.address],
        1
      );

      expect(community.address).to.not.equal(communityAddress3);
      expect(paymaster.address).to.not.equal(paymasterAddress3);
      //   expect(accountFactory.address).to.not.equal(accountFactoryAddress3);
      //   expect(profile.address).to.not.equal(profileAddress3);
    });
  });
});
