import "@nomicfoundation/hardhat-toolbox";
import { ethers, upgrades } from "hardhat";
import fs from "fs";
import path from "path";
import { terminal as term } from "terminal-kit";
import { config } from "dotenv";
import { execSync } from "child_process";

const listFiles = (directory: string, pattern: string): string[] => {
  try {
    const files = fs.readdirSync(directory);
    const extension = pattern.startsWith("*.") ? pattern.slice(1) : pattern;

    return files.filter((file) => path.extname(file) === extension);
  } catch (error) {
    console.error("Error reading directory:", error);
    return []; // Return an empty array in case of an error
  }
};

async function verifyContract(
  networkName: string,
  contractName: string,
  deployedContractAddress: string
) {
  execSync(
    `HARDHAT_NETWORK=${networkName} npx hardhat verify --contract contracts/tokens/${contractName}.sol:${contractName} ${deployedContractAddress}`,
    { stdio: "inherit" }
  );
}

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

  term("Choose a contract to deploy: ");

  const files = listFiles("./contracts/tokens/", "*.sol");

  const response = await term.singleColumnMenu(files).promise;

  const contractName = response.selectedText.replace(".sol", "");
  term("\n");

  term("Enter the token name (My Token): ");
  const tokenName = (await term.inputField({}).promise) || "My Token";

  term("\n");

  term("Enter the token symbol (MT): ");
  const tokenSymbol = (await term.inputField({}).promise) || "MT";

  term("\n");

  const factory = await ethers.getContractFactory(contractName);

  console.log("⚙️ deploying ERC20 Token...");

  const token = await upgrades.deployProxy(
    factory,
    [
      [process.env.MINTER1_ADDRESS, process.env.MINTER2_ADDRESS],
      tokenName,
      tokenSymbol,
    ],
    {
      kind: "uups",
      initializer: "initialize",
      timeout: 60000,
    }
  );

  console.log("🚀 request sent...");
  await token.deployed();

  // wait 2 seconds for the transaction to be mined
  await new Promise((resolve) => setTimeout(resolve, 2000));

  try {
    await verifyContract(networkName, contractName, token.address);
  } catch (error: any) {
    console.log("Error verifying contract: %s\n", error && error.message);
  }

  console.log(`Token deployed to ${token.address}`);

  process.exit();
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
