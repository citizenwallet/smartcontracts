const ethers = require("ethers");

const roleName = process.argv[2] || "MINTER_ROLE"; // replace with the actual role name
const roleHash = ethers.utils.id(roleName);

console.log(roleName, roleHash);
