import "@nomicfoundation/hardhat-toolbox";
import fs from "fs";
import path from "path";
import { terminal as term } from "terminal-kit";
import { execSync } from "child_process";
import { ethers, upgrades } from "hardhat";
import { config } from "dotenv";

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
    term("Enter the profile contract address to migrate FROM: ");
    fromContractAddress = await term.inputField({}).promise;
  }

  if (!fromContractAddress) {
    term("❌ invalid or missing from address");
    terminate();
  }

  term("\n");

  if (!toContractAddress) {
    term("Enter the profile contract address from migrate TO: ");
    toContractAddress = await term.inputField({}).promise;
  }

  if (!toContractAddress) {
    term("❌ invalid or missing to address");
    terminate();
  }

  // Read the file
  const filePath = path.join(__dirname, "./toMigrate.json");
  const data = fs.readFileSync(filePath, "utf8");

  // Parse the JSON into an array
  const addresses = JSON.parse(data);

  const factory = await ethers.getContractFactory("Profile");

  const fromContract = factory.attach(fromContractAddress!);

  const toContract = factory.attach(toContractAddress!);

  term("\n");

  for (const address of addresses) {
    term(`\n⚙️ migrating... ${address}...`);

    try {
      const ipfsHash = await fromContract.get(address);

      // fetch json from ipfs
      const ipfsUrl = `${process.env.IPFS_URL}/${ipfsHash}`;

      const response = await fetch(ipfsUrl);

      const data = await response.json();

      console.log(data);
      await toContract.set(
        address,
        ethers.utils.formatBytes32String(data.username),
        ipfsHash
      );

      term(`\n🚀 ${data.username} migrated to new contract`);
    } catch (error: any) {
      if (error.message.includes("ERC721: invalid token ID")) {
        term("\n❌ no profile at address");
        continue;
      }

      console.error(error);
    }
  }

  term(`Profiles migrated to ${toContractAddress}...`);

  process.exit();
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
