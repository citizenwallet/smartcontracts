const { ethers } = require("ethers");

async function main() {
  const value = process.env.VALUE;

  if (!value) {
    throw new Error("VALUE is not set");
  }

  // Parse the address to a BigNumber
  const profileNumber = ethers.BigNumber.from(value, 16);

  console.log("done");

  console.log(profileNumber.toString());
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
