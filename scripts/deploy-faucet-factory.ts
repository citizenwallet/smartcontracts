import "@nomicfoundation/hardhat-toolbox";
import { ethers, upgrades, run } from "hardhat";
import { terminal as term } from "terminal-kit";
import { config } from "dotenv";

const networkName = process.env.HARDHAT_NETWORK || "";
if (!networkName) {
  term.red("HARDHAT_NETWORK missing in your environment");
  term("\n");
  process.exit();
}

function terminate() {
  term.grabInput(false);
  setTimeout(function () {
    process.exit();
  }, 100);
}

term.on("key", function (name: string, _: any, __: any) {
  if (name === "CTRL_C") {
    terminate();
  }
});

async function main() {
  config();

  const contractName = "FaucetFactory";
  term(`Preparing to deploy ${contractName}.sol\n`);
  term("This will allow creating faucets.\n");

  term("\n");

  const factory = await ethers.getContractFactory(contractName);

  console.log("⚙️ deploying FaucetFactory...");

  const deployedContract = await factory.deploy();

  console.log("🚀 request sent...");
  await deployedContract.deployed();

  // wait for the transaction to be mined
  await deployedContract.deployTransaction.wait(5);

  console.log("🧐 verifying...\n");

  try {
    await run("verify:verify", {
      address: deployedContract.address,
    });
  } catch (error: any) {
    console.log("Error verifying contract: %s\n", error && error.message);
  }

  console.log(`\nContract deployed at ${deployedContract.address}\n`);

  process.exit();
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
