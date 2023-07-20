import "@openzeppelin/hardhat-upgrades";
import { HardhatUserConfig, task, types } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import { config } from "dotenv";

config();

const hhconfig: HardhatUserConfig = {
  solidity: {
    version: "0.8.20",
    settings: {
      evmVersion: "paris",
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
  networks: {
    chainstack: {
      url: process.env.CHAINSTACK_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
    },
  },
  etherscan: {
    apiKey: {
      polygonMumbai: process.env.POLYGON_MUMBAI_ETHERSCAN_API_KEY || "",
    },
  },
};

export default hhconfig;
