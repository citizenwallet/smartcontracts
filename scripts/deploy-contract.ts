import { ethers, network, upgrades } from "hardhat";
import dotenv from "dotenv";
import { terminal as term } from "terminal-kit";
import fs from 'fs';
import path from 'path';
dotenv.config();

const listFiles = (directory: string, pattern: string): string[] => {
  try {
      const files = fs.readdirSync(directory);
      const extension = pattern.startsWith('*.') ? pattern.slice(1) : pattern;

      return files.filter(file => path.extname(file) === extension);
  } catch (error) {
      console.error('Error reading directory:', error);
      return []; // Return an empty array in case of an error
  }
};

function terminate() {
	term.grabInput( false ) ;
	setTimeout( function() { process.exit() } , 100 ) ;
}

term.on( 'key' , function( name , matches , data ) {
	console.log( "'key' event:" , name ) ;
	if ( name === 'CTRL_C' ) { terminate() ; }
} ) ;

if (!process.env.DEPLOYER_PRIVATE_KEY) {
  term.red('Please set your DEPLOYER_PRIVATE_KEY in your .env.local file');
  term('\n');
  process.exit();
}

// Define a mapping of chain IDs to their native currency symbols
const nativeCurrencySymbols: { [chainId: number]: string } = {
  1: 'ETH',     // Ethereum Mainnet
  3: 'ETH',     // Ropsten
  4: 'ETH',     // Rinkeby
  5: 'ETH',     // Goerli
  42: 'ETH',    // Kovan
  56: 'BNB',    // Binance Smart Chain Mainnet
  97: 'BNB',    // Binance Smart Chain Testnet
  137: 'MATIC', // Polygon Mainnet
  8453: "ETH",  // Base
  84531: "ETH",  // Base
  80001: 'MATIC', // Polygon Mumbai Testnet
  42220: 'CELO',  // Celo Mainnet
  44787: 'CELO',  // Alfajores Testnet (Celo)
  10: 'ETH',    // Optimism Mainnet
  69: 'ETH',    // Optimism Kovan Testnet
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
    term('\n');
    process.exit();
  }
  const nativeCurrencySymbol = nativeCurrencySymbols[chainId] || `Unknown chainId: ${chainId}`;

  const wallet = new ethers.Wallet(process.env.DEPLOYER_PRIVATE_KEY, ethers.provider);
  const balanceWei = await wallet.getBalance();
  term(`The balance of the deployer wallet on ${networkName} is: ${ethers.utils.formatEther(balanceWei)} ${nativeCurrencySymbol}\n\n`);

  // Check if the balance is greater than 50 gwei
  if (BigInt(balanceWei) < BigInt(50 * 10**9)) {
    term.red(`Insufficiant balance on ${wallet.address} to deploy a contract.\n`);
    term(`Please add some ${nativeCurrencySymbol} to your wallet and try again.\n`);
    process.exit();
  }
  
  term( 'Choose a contract to deploy: ' ) ;

  const files = listFiles('./contracts/', '*.sol');

  const response = await term.singleColumnMenu(files).promise; 

  const contractName = response.selectedText.replace(".sol", "");
  term("\n");
  term.green("I will now deploy ").green.bold(contractName).green(" on ").green.bold(networkName).green("\n");
  term.green("Using the minters defined in your .env.local: \n");
  term.green("  Minter 1: %s\n", process.env.MINTER1_ADDRESS);
  term.green("  Minter 2: %s\n", process.env.MINTER2_ADDRESS);

  term("\n");
  term("Continue? [Y/n]");
	const confirm = await term.yesOrNo( { yes: [ 'y' , 'ENTER' ] , no: [ 'n' ] }).promise;
  if (confirm) {
      term( '\n' ).eraseLineAfter.green("Deploying %s on %s\n", contractName, networkName);
      deployContract(contractName, process.env.MINTER1_ADDRESS, process.env.MINTER2_ADDRESS);
    } else {
      term.red("Exiting...\n");
      process.exit();
    }
}

async function deployContract(contractName: string, minter1: string, minter2: string) {

  const Contract = await ethers.getContractFactory(contractName);
  const deployedContract = await upgrades.deployProxy(
    Contract,
    [[minter1, minter2]],
    { initializer: "initialize" }
  );

  console.log("Deployed contract address:", deployedContract.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });


