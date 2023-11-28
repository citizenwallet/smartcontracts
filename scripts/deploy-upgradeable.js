const { ethers, upgrades } = require("hardhat");
require("dotenv").config();

async function main() {
  const Contract = await ethers.getContractFactory("RegensUniteToken");
  const deployedContract = await upgrades.deployProxy(
    Contract,
    [[process.env.MINTER1_ADDRESS, process.env.MINTER2_ADDRESS]],
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
