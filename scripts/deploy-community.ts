import { ethers, upgrades, run } from "hardhat";
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

async function deployContract(
  contractName: string,
  minter1: string,
  minter2: string,
  tokenName: string,
  tokenSymbol: string
) {
  const Contract = await ethers.getContractFactory(contractName);
  const deployedContract = await upgrades.deployProxy(
    Contract,
    [[minter1, minter2], tokenName, tokenSymbol],
    {
      kind: "uups",
      initializer: "initialize",
    }
  );

  console.log("Deployed token address:", deployedContract.address);
  return deployedContract.address;
}

async function verifyContract(
  address: string,
  constructorArguments?: string[]
) {
  await run("verify:verify", {
    address,
    constructorArguments,
  });
}

async function deployCommunityEntrypoint(
  networkName: string,
  deployedContractAddress: string
) {
  execSync(
    `HARDHAT_NETWORK=${networkName} COMMUNITY_TOKEN_ADDRESS=${deployedContractAddress} npx hardhat run ./scripts/deploy-community-entrypoint.ts`,
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

const pk = process.env.DEPLOYER_PRIVATE_KEY;
if (!pk) {
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
if (!networkName) {
  term.red("HARDHAT_NETWORK missing in your environment");
  term("\n");
  process.exit();
}

async function getChainId() {
  const networkDetails = await ethers.provider.getNetwork();
  return networkDetails.chainId;
}

async function main() {
  // const spinner = term.spinner();
  // Fetch the chain ID from the provider
  let chainId;
  try {
    chainId = await getChainId();
  } catch (e) {
    term.red("Error fetching chain ID: %s\n", e && e.message);
    process.exit();
  }

  if (!chainId || !nativeCurrencySymbols[chainId]) {
    term.red(`Unsupported chain ID: ${chainId} (network: ${networkName})`);
    term("\n");
    process.exit();
  }
  const nativeCurrencySymbol =
    nativeCurrencySymbols[chainId] || `Unknown chainId: ${chainId}`;

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

  let tokenDecimals: number,
    tokenSymbol: string,
    deployedContractAddress: string;

  if (process.env.TOKEN_ADDRESS) {
    deployedContractAddress = process.env.TOKEN_ADDRESS;
    const contract = await ethers.getContractAt(
      contractName,
      process.env.TOKEN_ADDRESS
    );
    tokenDecimals = await contract.decimals();
    tokenSymbol = await contract.symbol();
    term.green("  Contract Name: %s\n", contractName);
    term.green("  Network: %s\n", networkName);
    // term.green("  Token Name: %s\n", tokenName);
    term.green("  Token Symbol: %s\n", tokenSymbol);
    term.green("  Token Decimals: %s\n", tokenDecimals);
    term.green("  Token Address: %s\n", process.env.TOKEN_ADDRESS);
    term("\n");
    term("Continue? [Y/n]");
  } else {
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

    tokenSymbol = (await term.inputField({}).promise) || "MTT";

    term("\n");
    term("You are about to deploy the following contract:\n");
    term.green("  Contract Name: %s\n", contractName);
    term.green("  Network: %s\n", networkName);
    term.green("  Minter 1: %s\n", minter1);
    term.green("  Minter 2: %s\n", minter2);
    term.green("  Token Name: %s\n", tokenName);
    term.green("  Token Symbol: %s\n", tokenSymbol);
    term("\n");
    term("Deploy contract? [Y/n]");
    const confirm = await term.yesOrNo({ yes: ["y", "ENTER"], no: ["n"] })
      .promise;
    if (!confirm) {
      term("\n");
      term.red("Exiting...\n");
      process.exit();
    } else {
      term("\n").eraseLineAfter.green(
        "Deploying %s on %s\n",
        contractName,
        networkName
      );
      deployedContractAddress = await deployContract(
        contractName,
        minter1,
        minter2,
        tokenName,
        tokenSymbol
      );
    }
  }

  term("\n");
  term("Do you want to verify this new contract on etherscan? [Y/n]");
  term("\n");

  const confirmVerify = await term.yesOrNo({ yes: ["y", "ENTER"], no: ["n"] })
    .promise;
  if (confirmVerify) {
    try {
      await verifyContract(deployedContractAddress);
    } catch (error) {
      term.red("Error verifying contract: %s\n", error && error.message);
    }
  }

  term("\n");
  term("Do you want to deploy a community entry point for this token? [Y/n]");
  const confirmDeployEntryPoint = await term.yesOrNo({
    yes: ["y", "ENTER"],
    no: ["n"],
  }).promise;
  if (confirmDeployEntryPoint) {
    await deployCommunityEntrypoint(networkName, deployedContractAddress);
  } else {
    term("\n");
    term.red("Exiting...\n");
    process.exit();
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
