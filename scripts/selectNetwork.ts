import fs from 'fs';
console.log(fs.readFileSync('./assets/citizenwallet.ans', 'utf8'));

import dotenv from "dotenv";
dotenv.config();

var execSync = require("child_process").execSync;
var term = require("terminal-kit").terminal;

dotenv.config();


const rpcUrlEnvVariables = Object.keys(process.env)
  .filter(key => key.endsWith('_RPC_URL'));

const networks = rpcUrlEnvVariables.map(key => key.replace('_RPC_URL', '').toLowerCase());

async function main() {

  if (process.env.HARDHAT_NETWORK) {
    runHardhatScriptOnNetwork(`./scripts/${process.argv[2]}.ts`, process.env.HARDHAT_NETWORK);
    process.exit(1);
  }

  term(`Please select a network to use: \n`);
  const response = await term.singleColumnMenu(networks).promise; 
  term("\n");
  term.green("Using %s\n", response.selectedText);
  term("\n");
  runHardhatScriptOnNetwork(`./scripts/${process.argv[2]}.ts`, response.selectedText);
  process.exit(1);
}

function runHardhatScriptOnNetwork(scriptPath: string, networkName: string) {
  process.env.HARDHAT_NETWORK = networkName;
  execSync(`npx hardhat run ${scriptPath}`, { stdio: "inherit" });
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
