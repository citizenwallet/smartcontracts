import { ethers, upgrades, run } from "hardhat";
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

  if (
    process.env.PAYMASTER_SPONSOR_ADDR === undefined ||
    process.env.PAYMASTER_SPONSOR_ADDR === ""
  ) {
    throw Error("PAYMASTER_SPONSOR_ADDR is not set");
  }

  console.log("deploying Paymaster...");

  const paymasterFactory = await ethers.getContractFactory("Paymaster");

  const paymaster = await upgrades.deployProxy(
    paymasterFactory,
    [process.env.PAYMASTER_SPONSOR_ADDR],
    {
      kind: "uups",
      initializer: "initialize",
      timeout: 999999,
    }
  );

  console.log("request sent...");

  await paymaster.deployed();

  console.log(`Paymaster deployed to: ${paymaster.address}`);

  console.log("verifying...");

  // wait for etherscan to index the contract
  await new Promise((resolve) => setTimeout(resolve, 10000));

  await run("verify:verify", {
    address: paymaster.address,
    constructorArguments: [],
  });

  console.log("verified!");

  console.log("deploying TokenEntryPoint...");

  const tokenEntryPointFactory = await ethers.getContractFactory(
    "TokenEntryPoint"
  );
  const tokenEntryPoint = await upgrades.deployProxy(
    tokenEntryPointFactory,
    [
      process.env.PAYMASTER_SPONSOR_ADDR,
      paymaster.address,
      process.env.ENTRYPOINT_ADDR,
    ],
    {
      kind: "uups",
      initializer: "initialize",
      timeout: 999999,
    }
  );

  console.log("request sent...");

  await tokenEntryPoint.deployed();

  console.log(`TokenEntryPoint deployed to: ${tokenEntryPoint.address}`);

  console.log("verifying...");

  // wait for etherscan to index the contract
  await new Promise((resolve) => setTimeout(resolve, 10000));

  await run("verify:verify", {
    address: tokenEntryPoint.address,
    constructorArguments: [],
  });

  console.log("verified!");

  console.log("deploying AccountFactory...");

  const accFactory = await ethers.deployContract("AccountFactory", [
    process.env.ENTRYPOINT_ADDR,
    tokenEntryPoint.address,
  ]);

  console.log("request sent...");

  await accFactory.deployed();

  console.log(`Account Factory deployed to: ${accFactory.address}`);

  console.log("verifying...");

  // wait for etherscan to index the contract
  await new Promise((resolve) => setTimeout(resolve, 10000));

  await run("verify:verify", {
    address: accFactory.address,
    constructorArguments: [
      process.env.ENTRYPOINT_ADDR,
      tokenEntryPoint.address,
    ],
  });

  console.log("verified!");

  console.log("*************************************");
  console.log("DEPLOYMENT COMPLETE");
  console.log(" ");
  console.log(" ");
  console.log("Paymaster: ", paymaster.address);
  console.log("Token Entry Point: ", tokenEntryPoint.address);
  console.log("Account Factory: ", accFactory.address);
  console.log("*************************************");
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
