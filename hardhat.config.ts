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
    ethereum: {
      url: process.env.ETHEREUM_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
      gasPrice: 9000000000, // this is 90 Gwei
    },
    polygon: {
      url: process.env.POLYGON_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
      gasPrice: 90000000000, // this is 90 Gwei
      timeout: 999999,
      throwOnTransactionFailures: true,
      throwOnCallFailures: true,
      allowUnlimitedContractSize: true,
    },
    polygon_testnet: {
      url: process.env.POLYGON_TESTNET_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
      gasPrice: 90000000000, // this is 90 Gwei
      timeout: 999999,
      gas: 12400000,
      throwOnTransactionFailures: true,
      throwOnCallFailures: true,
      allowUnlimitedContractSize: true,
    },
    base: {
      url: process.env.BASE_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
      gasPrice: 30000000, // this is 30 Gwei
    },
    base_testnet: {
      url: process.env.BASE_TESTNET_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
      gasPrice: 30000000, // this is 30 Gwei
    },
    celo: {
      url: process.env.CELO_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
      gasPrice: 5000000000, // this is 30 Gwei
    },
    celo_testnet: {
      url: process.env.CELO_TESTNET_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
      gasPrice: 5000000000, // this is 30 Gwei
    },
  },
  etherscan: {
    apiKey: {
      mainnet: process.env.ETEHREUM_MAINNET_ETHERSCAN_API_KEY || "",
      polygonMumbai: process.env.POLYGON_MUMBAI_ETHERSCAN_API_KEY || "",
      polygon: process.env.POLYGON_MAINNET_ETHERSCAN_API_KEY || "",
      base: process.env.BASE_MAINNET_ETHERSCAN_API_KEY || "",
      celo: process.env.CELO_MAINNET_ETHERSCAN_API_KEY || "",
    },
    customChains: [
      {
        network: "base",
        chainId: 8453,
        urls: {
          apiURL: "https://api.basescan.org/api",
          browserURL: "https://basescan.org",
        },
      },
      {
        network: "celo",
        chainId: 42220,
        urls: {
          apiURL: "https://api.celoscan.io/api",
          browserURL: "https://celoscan.io",
        },
      },
    ],
  },
};

export default hhconfig;
