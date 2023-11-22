import { ethers, upgrades, run } from "hardhat";
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

  if (
    process.env.TOKEN_ENTRYPOINT_ADDR === undefined ||
    process.env.TOKEN_ENTRYPOINT_ADDR === ""
  ) {
    throw Error("TOKEN_ENTRYPOINT_ADDR is not set");
  }

  const tokenEntryPointAddress = process.env.TOKEN_ENTRYPOINT_ADDR;

  console.log("deploying AccountFactory...");

  const accFactory = await ethers.deployContract("AccountFactory", [
    process.env.ENTRYPOINT_ADDR,
    tokenEntryPointAddress,
  ]);

  console.log("request sent...");

  await accFactory.deployed();

  console.log(`Account Factory deployed to: ${accFactory.address}`);

  console.log("verifying...");

  // wait for etherscan to index the contract
  await new Promise((resolve) => setTimeout(resolve, 10000));

  await run("verify:verify", {
    address: accFactory.address,
    constructorArguments: [process.env.ENTRYPOINT_ADDR, tokenEntryPointAddress],
  });

  console.log("verified!");

  console.log("*************************************");
  console.log("DEPLOYMENT COMPLETE");
  console.log(" ");
  console.log(" ");
  console.log("Account Factory: ", accFactory.address);
  console.log("*************************************");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
