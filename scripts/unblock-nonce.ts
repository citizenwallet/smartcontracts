import { ethers } from "hardhat";
import { config } from "dotenv";
import { terminal as term } from "terminal-kit";

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
  console.log("reading config...");
  config();

  term("\n");
  term("*************************************\n");
  term("Replace a pending transaction of an account with a new one.\n");
  term("This script will send 0 native currency to yourself.\n");
  term("Useful when gas fees are unpredictable and a transaction is stuck.\n");
  term("*************************************\n");
  term("\n");

  term("Enter the nonce you would like to replace (0): ");

  const strnonce: string = (await term.inputField({}).promise) || "0";

  const nonce = parseInt(strnonce);

  term("\n");

  // // The amount of Matic to send, in wei
  const amount = ethers.utils.parseUnits("0", "ether"); // 1 Matic

  // // Get the private key from environment variables
  const privateKey = process.env.DEPLOYER_PRIVATE_KEY;
  if (!privateKey) {
    throw new Error("Please set your DEPLOYER_PRIVATE_KEY in a .env file");
  }

  // // Get the provider
  const provider = ethers.provider;

  // // Create a new wallet instance
  const wallet = new ethers.Wallet(privateKey, provider);

  // // Send the Matic
  const tx = await wallet.sendTransaction({
    to: wallet.address,
    value: amount,
    nonce,
    maxPriorityFeePerGas: ethers.utils.parseUnits("950", "gwei").toNumber(),
    maxFeePerGas: ethers.utils.parseUnits("950", "gwei").toNumber(),
    gasLimit: 21000,
  });

  console.log("sending...");

  // // Wait for the transaction to be mined
  await tx.wait();

  console.log("sent!");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
