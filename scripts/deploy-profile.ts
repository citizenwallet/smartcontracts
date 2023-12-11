import "@nomicfoundation/hardhat-toolbox";
import { ethers, upgrades } from "hardhat";
import { config } from "dotenv";

async function main() {
  console.log("reading config...");
  config();

  const profileFactory = await ethers.getContractFactory("Profile");

  console.log("⚙️ deploying Profile...");

  const profile = await upgrades.deployProxy(profileFactory, [], {
    kind: "uups",
    initializer: "initialize",
    timeout: 60000,
  });

  console.log("🚀 request sent...");
  await profile.deployed();

  console.log(`Profile deployed to ${profile.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
