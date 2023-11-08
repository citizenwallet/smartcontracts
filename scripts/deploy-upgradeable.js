const { ethers, upgrades } = require("hardhat");
require("dotenv").config();

async function main() {
  const RegensUniteCoin = await ethers.getContractFactory("RegensUniteToken");
  const regensUniteCoin = await upgrades.deployProxy(
    RegensUniteCoin,
    [[process.env.MINTER1_ADDRESS, process.env.MINTER2_ADDRESS]],
    { initializer: "initialize" }
  );

  console.log("RegensUniteToken address:", regensUniteCoin.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
