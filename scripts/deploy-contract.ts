import { ethers, network, upgrades } from "hardhat";
import dotenv from "dotenv";
import { terminal as term } from "terminal-kit";
import fs from "fs";
import path from "path";
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

const networkName = process.env.HARDHAT_NETWORK;

async function main() {
  // const spinner = term.spinner();
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

  term("Choose a contract to deploy: ");

  const files = listFiles("./contracts/tokens/", "*.sol");

  const response = await term.singleColumnMenu(files).promise;

  const contractName = response.selectedText.replace(".sol", "");
  term("\n");
  term
    .green("I will now deploy ")
    .green.bold(contractName)
    .green(" on ")
    .green.bold(networkName)
    .green("\n");
  term.green("Using the minters defined in your .env.local: \n");
  const minter1 = process.env.MINTER1_ADDRESS;
  const minter2 = process.env.MINTER2_ADDRESS;
  if (!minter1 || !minter2) {
    term.red(
      "Please set your MINTER1_ADDRESS and MINTER2_ADDRESS in your .env.local file"
    );
    term("\n");
    process.exit();
  }

  term.green("  Minter 1: %s\n", minter1);
  term.green("  Minter 2: %s\n", minter2);

  term("\n");

  term("Enter a token name (My Test Token): ");

  const tokenName: string =
    (await term.inputField({}).promise) || "My Test Token";

  term("\n");

  term("Enter a token symbol (MTT): ");

  const tokenSymbol: string = (await term.inputField({}).promise) || "MTT";

  term("\n");

  term("How many decimals should it have (6): ");

  const tokenDecimalsInput: string = (await term.inputField({}).promise) || "6";
  const tokenDecimals = parseInt(tokenDecimalsInput);

  if (isNaN(tokenDecimals) || tokenDecimals < 0 || tokenDecimals > 18) {
    term.red("Decimals should be between 0 and 18\n");
    process.exit();
  }

  term("\n");
  term("You are about to deploy the following contract:\n");
  term.green("  Contract Name: %s\n", contractName);
  term.green("  Network: %s\n", networkName);
  term.green("  Minter 1: %s\n", minter1);
  term.green("  Minter 2: %s\n", minter2);
  term.green("  Token Name: %s\n", tokenName);
  term.green("  Token Symbol: %s\n", tokenSymbol);
  term.green("  Token Decimals: %s\n", tokenDecimals);
  term("\n");

  term("Continue? [Y/n]");
  const confirm = await term.yesOrNo({ yes: ["y", "ENTER"], no: ["n"] })
    .promise;
  if (confirm) {
    term("\n").eraseLineAfter.green(
      "Deploying %s on %s\n",
      contractName,
      networkName
    );
    deployContract(
      contractName,
      minter1,
      minter2,
      tokenName,
      tokenSymbol,
      tokenDecimals
    );
  } else {
    term.red("Exiting...\n");
    process.exit();
  }
}

async function deployContract(
  contractName: string,
  minter1: string,
  minter2: string,
  tokenName: string,
  tokenSymbol: string,
  tokenDecimals = 6
) {
  const Contract = await ethers.getContractFactory(contractName);
  const deployedContract = await upgrades.deployProxy(
    Contract,
    [[minter1, minter2], tokenName, tokenSymbol],
    {
      kind: "uups",
      initializer: "initialize",
      constructorArgs: [tokenDecimals],
    }
  );

  console.log("Deployed contract address:", deployedContract.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
