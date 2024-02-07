const { ethers } = require("ethers");

async function main() {
  const value = process.env.VALUE;

  if (!value) {
    throw new Error("VALUE is not set");
  }

  const bytes32 = ethers.utils.formatBytes32String(value);

  console.log(`Value: ${bytes32}`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
