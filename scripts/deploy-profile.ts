import "@nomicfoundation/hardhat-toolbox";
import { ethers, upgrades } from "hardhat";
import { terminal as term } from "terminal-kit";
import { config } from "dotenv";
import { execSync } from "child_process";

async function verifyContract(
  networkName: string,
  contractName: string,
  deployedContractAddress: string
) {
  execSync(
    `HARDHAT_NETWORK=${networkName} npx hardhat verify --contract contracts/apps/${contractName}.sol:${contractName} ${deployedContractAddress}`,
    { stdio: "inherit" }
  );
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

const networkName = process.env.HARDHAT_NETWORK || "";
if (!networkName) {
  term.red("HARDHAT_NETWORK missing in your environment");
  term("\n");
  process.exit();
}

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

  // wait 2 seconds for the transaction to be mined
  await new Promise((resolve) => setTimeout(resolve, 2000));

  try {
    await verifyContract(networkName, "Profile", profile.address);
  } catch (error: any) {
    console.log("Error verifying contract: %s\n", error && error.message);
  }

  console.log(`Profile deployed to ${profile.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
