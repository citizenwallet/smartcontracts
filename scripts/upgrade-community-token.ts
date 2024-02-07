import { ethers, upgrades } from "hardhat";
import dotenv from "dotenv";
import { terminal as term } from "terminal-kit";
import fs from "fs";
import path from "path";
import { execSync } from "child_process";
dotenv.config();

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

async function verifyContract(networkName: string, contractName: string, deployedContractAddress: string, tokenDecimals: number) {
  execSync(
    `HARDHAT_NETWORK=${networkName} npx hardhat verify --contract contracts/tokens/${contractName}.sol:${contractName} ${deployedContractAddress} ${tokenDecimals}`,
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

if (!process.env.DEPLOYER_PRIVATE_KEY) {
  term.red("Please set your DEPLOYER_PRIVATE_KEY in your .env.local file");
  term("\n");
  process.exit();
}

// Define a mapping of chain IDs to their native currency symbols
const nativeCurrencySymbols: { [chainId: number]: string } = {
  1: "ETH", // Ethereum Mainnet
  3: "ETH", // Ropsten
  4: "ETH", // Rinkeby
  5: "ETH", // Goerli
  42: "ETH", // Kovan
  56: "BNB", // Binance Smart Chain Mainnet
  97: "BNB", // Binance Smart Chain Testnet
  137: "MATIC", // Polygon Mainnet
  8453: "ETH", // Base
  84531: "ETH", // Base
  80001: "MATIC", // Polygon Mumbai Testnet
  42220: "CELO", // Celo Mainnet
  44787: "CELO", // Alfajores Testnet (Celo)
  10: "ETH", // Optimism Mainnet
  69: "ETH", // Optimism Kovan Testnet
  // Add other networks and their symbols as needed
};

const networkName = process.env.HARDHAT_NETWORK || "";

async function main() {
  // Fetch the chain ID from the provider
  const networkDetails = await ethers.provider.getNetwork();
  const chainId = networkDetails.chainId;
  if (!chainId || !nativeCurrencySymbols[chainId]) {
    term.red(`Unsupported chain ID: ${chainId} (network: ${networkName})`);
    term("\n");
    process.exit();
  }
  const nativeCurrencySymbol =
    nativeCurrencySymbols[chainId] || `Unknown chainId: ${chainId}`;

  const pk = process.env.DEPLOYER_PRIVATE_KEY;
  if (!pk) {
    term.red("Please set your DEPLOYER_PRIVATE_KEY in your .env.local file");
    term("\n");
    process.exit();
  }

  const wallet = new ethers.Wallet(pk, ethers.provider);
  const balanceWei = await wallet.getBalance();
  term(
    `The balance of the deployer wallet on ${networkName} is: ${ethers.utils.formatEther(
      balanceWei
    )} ${nativeCurrencySymbol}\n\n`
  );

  // Check if the balance is greater than 50 gwei
  if (balanceWei.toBigInt() < BigInt(50 * 10 ** 9)) {
    term.red(
      `Insufficiant balance on ${wallet.address} to deploy a contract.\n`
    );
    term(
      `Please add some ${nativeCurrencySymbol} to your wallet and try again.\n`
    );
    process.exit();
  }

  term("Enter the address of your community token that you want to upgrade: ");
  const communityTokenAddress = await term.inputField({}).promise;

  term("Choose a contract to upgrade to: ");

  const files = listFiles("./contracts/tokens/", "*.sol");

  const response = await term.singleColumnMenu(files).promise;

  const contractName = response.selectedText.replace(".sol", "");
  term("\n");

  let tokenDecimals: number, tokenSymbol: string, tokenName: string, deployedContractAddress: string;

  if (!communityTokenAddress) {
    return;
  }

  const contract = await ethers.getContractAt(contractName, communityTokenAddress);
  tokenDecimals = await contract.decimals();
  tokenSymbol = await contract.symbol();
  tokenName = await contract.name();
  term.green("  Contract Name: %s\n", tokenName);
  term.green("  Network: %s\n", networkName);
  // term.green("  Token Name: %s\n", tokenName);
  term.green("  Token Symbol: %s\n", tokenSymbol);
  term.green("  Token Decimals: %s\n", tokenDecimals);
    term.green("  Token Address: %s\n", communityTokenAddress);
  term("\n");  

  term
      .green("I will now upgrade ")
      .green.bold(tokenName)
      .green(" to ")
      .green.bold(contractName)
      .green(" on ")
      .green.bold(networkName)
      .green("\n");

      term("\n");
      term("Continue? [Y/n]");
  const confirm = await term.yesOrNo({ yes: ["y", "ENTER"], no: ["n"] }).promise;
  if (!confirm) {
    term("\n");
    process.exit();
    return;
  }     

  const tokenFactory = await ethers.getContractFactory("UpgradeableCommunityToken");
  const V2Factory = await ethers.getContractFactory(contractName);

  await upgrades.forceImport(communityTokenAddress, tokenFactory, {
    kind: "uups"
  });
  // Upgrading the existing proxy to the new implementation
  await upgrades.upgradeProxy(communityTokenAddress, V2Factory, {
    kind: "uups"
  });


  term("\n");
  term("Do you want to verify this new contract on etherscan? [Y/n]");
  term("\n");

  const confirmVerify = await term.yesOrNo({ yes: ["y", "ENTER"], no: ["n"] });
  if (confirmVerify) {
    try {
      await verifyContract(networkName, contractName, communityTokenAddress, tokenDecimals);
    } catch (error) {
      term.red("Error verifying contract: %s\n", error);
    }
  }

  term("\n");
  process.exit();
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
