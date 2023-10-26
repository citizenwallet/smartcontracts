import { ethers } from "hardhat";
import { config } from "dotenv";

async function main() {
  console.log("reading config...");
  config();

  if (
    process.env.ENTRYPOINT_ADDR === undefined ||
    process.env.ENTRYPOINT_ADDR === ""
  ) {
    throw Error("ENTRYPOINT_ADDR is not set");
  }

  console.log("deploying...");
  const accFactory = await ethers.deployContract("AccountFactory", [
    process.env.ENTRYPOINT_ADDR,
  ]);

  console.log("request sent...");
  await accFactory.deployed();

  console.log(`Account Factory deployed to: ${accFactory.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
