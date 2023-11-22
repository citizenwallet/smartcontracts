import readline from "readline";
import { ethers, upgrades, run } from "hardhat";
import "@nomicfoundation/hardhat-toolbox";
import { config } from "dotenv";

async function main() {
  console.log("reading config...");
  config();

  if (
    process.env.ENTRYPOINT_ADDR === undefined ||
    process.env.ENTRYPOINT_ADDR === ""
  ) {
    throw Error("ENTRYPOINT_ADDR is not set");
  }

  const response = new Promise<string[]>((resolve) => {
    const rl = readline.createInterface({
      input: process.stdin,
      output: process.stdout,
    });

    rl.question(
      "Please enter the contracts to be whitelisted (comma separated) (ex: 0x123...,0x345...): ",
      (answer) => {
        resolve(answer.split(",").map((addr) => addr.trim()));
        rl.close();
      }
    );
  });

  const contractAddresses = await response;

  console.log("\n");

  if (contractAddresses.length > 0) {
    console.log(`whitelisting contracts:`);
    contractAddresses.forEach((addr) => console.log(addr));
    console.log("\n");
  }

  const sponsor = ethers.Wallet.createRandom();

  console.log("⚙️ deploying Paymaster...");

  const paymasterFactory = await ethers.getContractFactory("Paymaster");

  const paymaster = await upgrades.deployProxy(
    paymasterFactory,
    [sponsor.address],
    {
      kind: "uups",
      initializer: "initialize",
      timeout: 999999,
    }
  );

  console.log("🚀 request sent...");

  await paymaster.deployed();

  console.log(`✅ Paymaster deployed to: ${paymaster.address}`);

  console.log("\n");

  console.log("⚙️ deploying TokenEntryPoint...");

  const tokenEntryPointFactory = await ethers.getContractFactory(
    "TokenEntryPoint"
  );
  const tokenEntryPoint = await upgrades.deployProxy(
    tokenEntryPointFactory,
    [sponsor.address, paymaster.address, contractAddresses],
    {
      kind: "uups",
      initializer: "initialize",
      constructorArgs: [process.env.ENTRYPOINT_ADDR],
      timeout: 999999,
    }
  );

  console.log("🚀 request sent...");

  await tokenEntryPoint.deployed();

  console.log(`✅ TokenEntryPoint deployed to: ${tokenEntryPoint.address}`);

  console.log("\n");

  console.log("⚙️ deploying AccountFactory...");

  const accFactory = await ethers.deployContract("AccountFactory", [
    process.env.ENTRYPOINT_ADDR,
    tokenEntryPoint.address,
  ]);

  console.log("🚀 request sent...");

  await accFactory.deployed();

  console.log(`✅ Account Factory deployed to: ${accFactory.address}`);

  console.log("\n");

  console.log("⚙️ deploying Profile...");

  const profileFactory = await ethers.getContractFactory("Profile");

  const profile = await upgrades.deployProxy(profileFactory, [], {
    kind: "uups",
    initializer: "initialize",
    timeout: 999999,
  });

  console.log("🚀 request sent...");
  await profile.deployed();

  console.log(`✅ Profile deployed to ${profile.address}`);

  console.log("\n");

  console.log("🧐 verifying...");

  // wait for etherscan to index the contract
  await new Promise((resolve) => setTimeout(resolve, 10000));

  try {
    await run("verify:verify", {
      address: paymaster.address,
      constructorArguments: [],
    });

    console.log("verified!");
  } catch (error) {
    console.log(error);
  }

  try {
    await run("verify:verify", {
      address: tokenEntryPoint.address,
      constructorArguments: [process.env.ENTRYPOINT_ADDR],
    });

    console.log("verified!");
  } catch (error) {
    console.log(error);
  }

  try {
    await run("verify:verify", {
      address: accFactory.address,
      constructorArguments: [
        process.env.ENTRYPOINT_ADDR,
        tokenEntryPoint.address,
      ],
    });

    console.log("verified!");
  } catch (error) {
    console.log(error);
  }

  try {
    await run("verify:verify", {
      address: profile.address,
      constructorArguments: [],
    });

    console.log("verified!");
  } catch (error) {
    console.log(error);
  }

  console.log("\n");
  console.log("\n");
  console.log("\n");

  console.log("*************************************");
  console.log("DEPLOYMENT COMPLETE 🎉");
  console.log("\n");
  console.log("\n");
  console.log("Paymaster: ", paymaster.address);
  console.log("paymaster sponsor address: ", sponsor.address);
  console.log(
    "paymaster sponsor private key: ",
    sponsor.privateKey.replace("0x", "")
  );
  console.log("\n");
  console.log("\n");
  console.log("Token Entry Point: ", tokenEntryPoint.address);
  console.log("Account Factory: ", accFactory.address);
  console.log("Profile: ", profile.address);
  console.log("*************************************");

  console.log("\n");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
