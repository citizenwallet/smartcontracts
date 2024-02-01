import fs from "fs";
import path from "path";
import { execSync } from "child_process";
import dotenv from "dotenv";
dotenv.config();

var term = require("terminal-kit").terminal;

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

console.log(fs.readFileSync('./assets/citizenwallet.ans', 'utf8'));

const rpcUrlEnvVariables = Object.keys(process.env)
  .filter(key => key.endsWith('_RPC_URL'));

const networks = rpcUrlEnvVariables.map(key => key.replace('_RPC_URL', '').toLowerCase());

async function main() {
  term(`Please select a network to use: \n`);
  const networkChoice = await term.singleColumnMenu(networks).promise; 
  term("\n");
  term.green("Using %s\n", networkChoice.selectedText);
  const networkName = networkChoice.selectedText;
  term("\n");
  term(``)

  term("Choose a contract to verify: ");

  const files = listFiles("./contracts/tokens/", "*.sol");

  const contractChoice = await term.singleColumnMenu(files).promise;
  const contractName = contractChoice.selectedText.replace(".sol", "");
  term("\n");

  term(`Enter the contract address on ${networkName}: `);
  const contractAddress = (await term.inputField({}).promise).trim();

  term("\n");
  term("Enter the contract arguments: ");
  const contractArguments = (await term.inputField({}).promise).trim();

  term("\n\n");
  term
    .green("I will now verify the following contract:\n")
    .green("\n")
    .green.bold("Contract: \t")
    .green(contractName)
    .green("\n")
    .green.bold("Network: \t")
    .green(networkName)
    .green("\n")
    .green.bold("Address: \t")
    .green(contractAddress)
    .green("\n")
    .green.bold("Arguments: \t")
    .green(`${contractArguments}`);

  term("\n\n");
  term("Continue? (Y/n) ");
  const confirm = await term.yesOrNo({ yes: ["y", "ENTER"], no: ["n"] }).promise;
  if (confirm) {
    process.env.HARDHAT_NETWORK = networkName;
    const cmd = `HARDHAT_NETWORK=${networkName} npx hardhat verify ${contractAddress} ${contractArguments}`;
    term("\n\n");
    term.yellow(cmd);
    term("\n\n");
    execSync(cmd, { stdio: "inherit" });
  } else {
    term("\n");
    term.red("Aborted");
    term("\n");
    process.exit(1);
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
