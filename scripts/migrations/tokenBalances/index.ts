import "@nomicfoundation/hardhat-toolbox";
import fs from "fs";
import path from "path";
import { terminal as term } from "terminal-kit";
import { execSync } from "child_process";
import { ethers, upgrades } from "hardhat";
import { config } from "dotenv";

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

  let fromContractAddress = process.env.FROM_CONTRACT;
  let toContractAddress = process.env.TO_CONTRACT;

  if (!fromContractAddress) {
    term("Enter the token contract address to migrate FROM: ");
    fromContractAddress = await term.inputField({}).promise;
  }

  if (!fromContractAddress) {
    term("❌ invalid or missing from address");
    terminate();
  }

  term("\n");

  term("Choose a contract to read from: ");

  let files = listFiles("./contracts/tokens/", "*.sol");

  let response = await term.singleColumnMenu(files).promise;

  let contractName = response.selectedText.replace(".sol", "");
  term("\n");

  if (!toContractAddress) {
    term("Enter the token contract address from migrate TO: ");
    toContractAddress = await term.inputField({}).promise;
  }

  if (!toContractAddress) {
    term("❌ invalid or missing to address");
    terminate();
  }

  term("Choose a contract to write to: ");

  files = listFiles("./contracts/tokens/", "*.sol");

  response = await term.singleColumnMenu(files).promise;

  contractName = response.selectedText.replace(".sol", "");
  term("\n");

  // Read the file
  const filePath = path.join(__dirname, "./toMigrate.json");
  const data = fs.readFileSync(filePath, "utf8");

  // Parse the JSON into an array
  const addresses = JSON.parse(data);

  const fromFactory = await ethers.getContractFactory(contractName);

  const fromContract = fromFactory.attach(fromContractAddress!);

  const toFactory = await ethers.getContractFactory(contractName);

  const toContract = toFactory.attach(toContractAddress!);

  const [signer] = await ethers.getSigners();
  const signerAddress = await signer.getAddress();

  term("\n");

  // temporarily grant minter role
  const tx = await toContract.grantRole(
    ethers.utils.id("MINTER_ROLE"),
    signerAddress
  );

  await tx.wait();

  for (const address of addresses) {
    term(`\n⚙️ migrating... ${address}...`);

    try {
      const balance = await fromContract.balanceOf(address); // 0x2d900678a66df705D3F3184267eAf603d809d3c4 // 0x9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6

      if (balance == 0) {
        term("\n❌ nothing to migrate");
        continue;
      }

      const toBalance = await toContract.balanceOf(address);

      if (toBalance > 0) {
        term("\n❌ already migrated");
        continue;
      }

      await toContract.mint(
        address,
        balance,
        "migration from " + fromContractAddress + " to " + toContractAddress
      );

      term(`\n🚀 ${balance} for ${address} migrated to new contract`);
    } catch (error: any) {
      console.error(error);
    }
  }

  // revoke minter role
  await toContract.revokeRole(ethers.utils.id("MINTER_ROLE"), signerAddress);

  term(`\nToken balance migrated to ${toContractAddress}...`);

  process.exit();
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
