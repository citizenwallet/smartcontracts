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
      await entrypoint.address
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
    };
  }

  async function deployCommunityFixture() {
    const { owner, communityFactory, friend1, friend2, vendor1, vendor2 } =
      await loadFixture(deployCommunityFactoryFixture);

    const tx = await communityFactory.create(friend1.address, 0);

    await tx.wait();

    const communityAddress = await communityFactory.get(friend1.address, 0);

    const community = await ethers.getContractAt(
      "TokenEntryPoint",
      communityAddress
    );

    const paymaster = await ethers.getContractAt(
      "Paymaster",
      await community.paymaster()
    );

    return {
      owner,
      communityFactory,
      community,
      paymaster,
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
  });
});
