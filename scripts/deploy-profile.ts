import { ethers } from "hardhat";
import { config } from "dotenv";

async function main() {
  console.log("reading config...");
  config();

  console.log("deploying...");
  const profile = await ethers.deployContract("Profile");

  console.log("request sent...");
  await profile.deployed();

  console.log(`Profile deployed to ${profile.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
