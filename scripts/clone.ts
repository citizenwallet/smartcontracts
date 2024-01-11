import { artifacts, ethers } from "hardhat";
import fs from "fs";
import path from "path";
import type { BigNumber } from 'ethers';

import tokenHolders from "../ZinneTokenHolders.json";

const profileContractAddress = "0xcA1B9EC1117340818C1c1fdd1B48Ea79E57C140F";
const sourceTokenAddress = "0x0785d720279f42326846d5396b5f44b97d0bfecd";
const newTokenAddress = "0x230542eda83346929e4E54f4a98e1ca1A4BFc0c3";

const deployerPrivateKey = process.env.DEPLOYER_PRIVATE_KEY;
if (!deployerPrivateKey) {
    console.log('DEPLOYER_PRIVATE_KEY is not set in .env file');
    process.exit(1);
}

const networkName = hre.network.name;
if (!networkName) {
  console.log('HARDHAT_NETWORK is not set in .env file');
  process.exit(1);
}

function getABI(jsonPath: string): any {
  const filePath = path.join(__dirname, jsonPath);
  const fileContents = fs.readFileSync(filePath, 'utf8');
  return JSON.parse(fileContents).abi;
}

const tokenABI = getABI("../artifacts/contracts/tokens/UpgradeableCommunityToken.sol/UpgradeableCommunityToken.json");
const profileABI = getABI("../artifacts/contracts/apps/Profile.sol/Profile.json");
const contractInterface = new ethers.utils.Interface(tokenABI);
type TokenHolder = {
  address: string,
  balance: BigNumber,
  username?: string
};

type AddressMapping = { [key: string]: string };

const skip = {"0xaa0f7EddCF6CaA17d2DBC919D0Fd6232D1b345c4": true, "0x4053289aa2e0E8556c6e1c6F6CD7D1183e2232ac": true};
// const skip = { 
//   "0xaa0f7EddCF6CaA17d2DBC919D0Fd6232D1b345c4": true, 
//   "0x3E483246e3178e9834549f2be9cC6547db750777": true, 
//   "0xCF1A8e25982E8B8e9C43B30D5F104940Aa5b13F2": true, 
//   "0xCc7142DE40Ed1c2D86308145DcC6836eFb370D6c": true, 
//   "0x2b4ACB6305fEF10be4Bc34a1FEF1Cf912ac39638": true, 
//   "0xB771864AC27Cd3e4F26a307DB05517bf10fF42A9": true, 
//   "0xe4f2633eb55A66473B50c9630EC8550945268874": true, 
//   "0x4053289aa2e0E8556c6e1c6F6CD7D1183e2232ac": true };

const mapping: AddressMapping = {
  "0xadDCA8E4b04AA119168b37DFa11C97827493ABe4": "0x39F2638B50DED59f9385BC76C96E44374f0A2119", // @bank
  "0xb1F9B11D8D5A0C69c0a4c337c5FdA823617cAF44": "0xaa34c2d6C923c176eCa418E681Aa7520c0e7BCFD", // @bank2
  "0x323b57416eb47aEeCE53F35E207228Eed765105F": "0x508846a71989ea4e07da68F4a838C9A5442dA617", // @kevex91
  "0x5EBfd5359D819AbB9c067b1d57f28597f8e949da": "0x34C4360bd268a9d615Bc3382f35792e03C026Ca7", // @bar
  "0xCF1A8e25982E8B8e9C43B30D5F104940Aa5b13F2": "0xabf962E0DaFa36126DB6C38F7599fe889837c320" // Xavier
};

// const tokenHolders: string[] = [
//   "0xaa0f7EddCF6CaA17d2DBC919D0Fd6232D1b345c4",
//   "0x4053289aa2e0E8556c6e1c6F6CD7D1183e2232ac",
//   "0xe0aB8F24c97d0BF1c1AB7aAbD76f861C7ca1e70A",
//   "0x1f16E4425144cE4A9C7f55c74756Df793A3AE32f",
//   "0x5EBfd5359D819AbB9c067b1d57f28597f8e949da"];

const options = {
  maxFeePerGas: ethers.utils.parseUnits("150", 'gwei'),
  maxPriorityFeePerGas: ethers.utils.parseUnits("31", 'gwei')
};

console.log(">>> using", options);

async function main() {
    // Replace with your contract addresses and ABIs

    const deployerWallet = new ethers.Wallet(deployerPrivateKey, ethers.provider);
    const currentNonce = await deployerWallet.getTransactionCount();
    console.log(">>> currentNonce", currentNonce);
    const [signer] = await ethers.getSigners();

    const sourceToken = new ethers.Contract(sourceTokenAddress, tokenABI, signer);
    const profileContract = new ethers.Contract(profileContractAddress, profileABI, signer);
    const newToken = new ethers.Contract(newTokenAddress, tokenABI, deployerWallet);
    // console.log("Granting MINTER_ROLE to", deployerWallet.address);
    // await newToken.grantRole("0x9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6", deployerWallet.address);
    const balanceHolders:TokenHolder[] = [];

    for (const tokenHolder of tokenHolders) {
      let username, balance, holder:TokenHolder;
      if (typeof tokenHolder === "string" && tokenHolder.startsWith("0x")) {
        const balance = await sourceToken.balanceOf(tokenHolder);
        holder = {address: tokenHolder, balance };
        try {
          const profileBytes = await profileContract.usernames(tokenHolder);
          if (profileBytes !== "0x0000000000000000000000000000000000000000000000000000000000000000") {
            holder.username = (ethers.utils.toUtf8String(profileBytes)).trim();
          }
        } catch (e) {
          console.log("!!!", e);
        }
        if (balance > 0) {
          balanceHolders.push(holder);
        }
      }
      else {
        holder = {
          address: tokenHolder.address,
          balance: ethers.BigNumber.from(tokenHolder.balance),
          username: tokenHolder.username
        };
        balanceHolders.push(holder);
      }
      console.log(`${holder.address}\t${ethers.utils.formatUnits(holder.balance, 6)}\t`, holder.username && `@${holder.username}`);
    }

    console.log("\n");
    let totalSupply = 0, i=0;
    for(const holder of balanceHolders) {
      let accountAddress: string = holder.address;
      if (skip[accountAddress]) {
        console.log("Skipping", accountAddress);
      } else {
        if (mapping[holder.address]) {
          accountAddress = mapping[holder.address];
          console.log(`Minting ${ethers.utils.formatUnits(holder.balance, 6)} tokens to new address ${accountAddress}`);
        } else {
        console.log(`Minting ${ethers.utils.formatUnits(holder.balance, 6)} tokens to ${accountAddress}`);
        }
        totalSupply += parseFloat(ethers.utils.formatUnits(holder.balance, 6));

        // ...

        let estimatedGasLimit: BigNumber;
        try {
          estimatedGasLimit = await newToken.estimateGas.mint(accountAddress, holder.balance, "recovered from old token");
          options.gasLimit = estimatedGasLimit.mul(110).div(100);
        } catch (error) {
          console.error("Error estimating gas:", error);
          // Handle the error appropriately

        }
        const data = contractInterface.encodeFunctionData('mint', [accountAddress, holder.balance, "recovered from old token"]);

        const tx = {
          to: newTokenAddress, // or another address if you prefer
          data,
          value: ethers.utils.parseEther("0"),
          // nonce: currentNonce + i, // Nonce of the stuck transaction
          maxFeePerGas: options.maxFeePerGas, // Higher gas price
          maxPriorityFeePerGas: options.maxPriorityFeePerGas, // Higher gas price
          gasLimit: options.gasLimit, // Minimum gas limit for a simple transaction
        };
        try {
          // const transactionResponse = await deployerWallet.sendTransaction(tx);
          // const transactionResponse = await newToken.mint(accountAddress, holder.balance, "recovered from old token", options);
          console.log(">>> hash", transactionResponse.hash);
          const transactionReceipt = await transactionResponse.wait();
          console.log(">>> block number", transactionReceipt.blockNumber);
          console.log(">>> gas limit / used", options.gasLimit.toString(), transactionReceipt.gasUsed.toString());
          console.log(">>> effective gas price", transactionReceipt.effectiveGasPrice.toString());
          i++;
        } catch (error) {
          console.error("Minting failed with error: ", error);
          process.exit(1);
      }
        // process.exit(1);
      }
    }

    console.log(">>> balanceHolders", balanceHolders.length, " - Total supply", totalSupply);
    if (balanceHolders.length != tokenHolders.length) {
      console.log(">>> To speed up next execution, import this JSON", JSON.stringify(balanceHolders, null, 2));
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
