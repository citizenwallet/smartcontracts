import "@nomicfoundation/hardhat-toolbox";
import { ethers, run } from "hardhat";
import { terminal as term } from "terminal-kit";
import { config } from "dotenv";

const networkName = process.env.HARDHAT_NETWORK || "";
if (!networkName) {
  term.red("HARDHAT_NETWORK missing in your environment");
  term("\n");
  process.exit();
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

async function main() {
  config();

  if (
    process.env.ENTRYPOINT_ADDR === undefined ||
    process.env.ENTRYPOINT_ADDR === ""
  ) {
    throw Error("ENTRYPOINT_ADDR is not set");
  }

  const entryPoint: string = process.env.ENTRYPOINT_ADDR;

  const contractName = "TokenEntryPointFactory";
  term(`Preparing to deploy ${contractName}.sol\n`);

  term("\n");

  let factory = await ethers.getContractFactory(contractName);

  console.log(`⚙️ deploying ${contractName}...`);

  const deployedContract = await factory.deploy(entryPoint);

  console.log("🚀 request sent...");
  await deployedContract.deployTransaction.wait(5);
  console.log("deployed...");

  const contract2Name = "AccountFactoryFactory";
  const factory2 = await ethers.getContractFactory(contract2Name);

  console.log(`⚙️ deploying ${contract2Name}...`);

  const deployedContract2 = (await factory2.deploy(entryPoint)) as any;

  console.log("🚀 request sent...");
  await deployedContract2.deployTransaction.wait(5);
  console.log("deployed...");

  const contract3Name = "ProfileFactory";
  const factory3 = await ethers.getContractFactory(contract3Name);

  console.log(`⚙️ deploying ${contract3Name}...`);

  const deployedContract3 = (await factory3.deploy()) as any;

  console.log("🚀 request sent...");
  await deployedContract3.deployTransaction.wait(5);
  console.log("deployed...");

  console.log("🧐 verifying...\n");

  try {
    await run("verify:verify", {
      address: deployedContract.address,
      constructorArguments: [entryPoint],
    });

    await run("verify:verify", {
      address: deployedContract2.address,
      constructorArguments: [entryPoint],
    });
  } catch (error: any) {
    console.log("Error verifying contract: %s\n", error && error.message);
  }

  console.log(`\nContracts deployed:`);
  console.log(`\n${contractName} deployed at ${deployedContract.address}`);
  console.log(`\n${contract2Name} deployed at ${deployedContract2.address}`);
  console.log(`\n${contract3Name} deployed at ${deployedContract3.address}`);

  process.exit();
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
