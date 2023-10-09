import { ethers } from "hardhat";
import { config } from "dotenv";

async function main() {
  console.log("reading config...");
  config();

  // const overrides = {
  //   nonce: 0,
  // };

  console.log("deploying...");
  // const accf = await ethers.deployContract("AccountFactory", [
  //   "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789",
  // ]);
  const accf = await ethers.getContractFactory("AccountFactory");

  const contract = await accf.deploy(
    "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789"
    // overrides
  );

  console.log("request sent...");
  console.log(
    `Account Factory will be deployed to ${contract.address} tx: ${contract.deployTransaction.hash}`
  );
  await contract.deployTransaction.wait();

  console.log("Account Factory deployed");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
