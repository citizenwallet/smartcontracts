import "@nomicfoundation/hardhat-toolbox";
import { ethers, upgrades, run } from "hardhat";
import { terminal as term } from "terminal-kit";
import { config } from "dotenv";
import { contractExists } from "./functions/contracts";

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

  const contractName = "ERC20IOU";
  term(`Preparing to deploy ${contractName}.sol\n`);
  term("This will allow keeping track of IOUs on an existing ERC20 token.\n");

  term("\n");

  term("Enter the token contract address: ");
  const tokenAddress = await term.inputField({}).promise;

  if (!tokenAddress) {
    term.red("Token address is required");
    term("\n");
    process.exit();
  }

  if (!contractExists(tokenAddress)) {
    term.red("No contract found at the provided address: %s\n", tokenAddress);
    term("\n");
    process.exit();
  }

  term("\n");

  const factory = await ethers.getContractFactory(contractName);

  console.log("⚙️ deploying ERC20IOU...");

  const deployedContract = await upgrades.deployProxy(factory, [tokenAddress], {
    kind: "uups",
    initializer: "initialize",
    timeout: 60000,
  });

  console.log("🚀 request sent...");
  await deployedContract.deployed();

  // wait 2 seconds for the transaction to be mined
  await new Promise((resolve) => setTimeout(resolve, 2000));

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
