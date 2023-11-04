import { ethers, upgrades } from "hardhat";
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
    process.env.PAYMASTER_ADDR === undefined ||
    process.env.PAYMASTER_ADDR === ""
  ) {
    throw Error("PAYMASTER_ADDR is not set");
  }

  console.log("deploying Paymaster...");

  const paymasterFactory = await ethers.getContractFactory("Paymaster");
  const paymaster = await upgrades.deployProxy(
    paymasterFactory,
    [process.env.PAYMASTER_ADDR],
    {
      kind: "uups",
      initializer: "initialize",
    }
  );

  console.log("request sent...");

  await paymaster.deployed();
  console.log(`Paymaster deployed to: ${paymaster.address}`);

  console.log("deploying TokenEntryPoint...");

  const tokenEntryPointFactory = await ethers.getContractFactory(
    "TokenEntryPoint"
  );
  const tokenEntryPoint = await upgrades.deployProxy(
    tokenEntryPointFactory,
    [process.env.PAYMASTER_ADDR],
    {
      kind: "uups",
      initializer: "initialize",
    }
  );

  console.log("request sent...");

  await tokenEntryPoint.deployed();
  console.log(`TokenEntryPoint deployed to: ${tokenEntryPoint.address}`);

  console.log("deploying AccountFactory...");
  const accFactory = await ethers.deployContract("AccountFactory", [
    process.env.ENTRYPOINT_ADDR,
    tokenEntryPoint.address,
  ]);

  console.log("request sent...");

  await accFactory.deployed();

  console.log(`Account Factory deployed to: ${accFactory.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
