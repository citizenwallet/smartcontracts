import "@openzeppelin/hardhat-upgrades";
import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import { config } from "dotenv";
import { ethers } from "ethers";

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
    ethereum_mainnet: {
      url: process.env.ETHEREUM_MAINNET_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
      gasPrice: ethers.utils.parseUnits("90", "gwei").toNumber(), // this is 90 Gwei
    },
    polygon_mainnet: {
      url: process.env.POLYGON_MAINNET_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
      gasPrice: ethers.utils.parseUnits("450", "gwei").toNumber(), // this is 150 Gwei
      timeout: 999999,
      throwOnTransactionFailures: true,
      throwOnCallFailures: true,
      allowUnlimitedContractSize: true,
    },
    polygon_testnet: {
      url: process.env.POLYGON_TESTNET_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
      gasPrice: ethers.utils.parseUnits("5", "gwei").toNumber(), // this is 90 Gwei
      timeout: 999999,
      gas: 12400000,
      throwOnTransactionFailures: true,
      throwOnCallFailures: true,
      allowUnlimitedContractSize: true,
    },
    gnosis_mainnet: {
      url: process.env.GNOSIS_MAINNET_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
    },
    gnosis_testnet: {
      url: process.env.GNOSIS_TESTNET_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
    },
    base_mainnet: {
      url: process.env.BASE_MAINNET_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
    },
    base_testnet: {
      url: process.env.BASE_TESTNET_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
    },
    celo_mainnet: {
      url: process.env.CELO_MAINNET_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
    },
    celo_testnet: {
      url: process.env.CELO_TESTNET_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
    },
  },
  etherscan: {
    apiKey: {
      mainnet: process.env.ETHEREUM_MAINNET_ETHERSCAN_API_KEY || "",
      sepolia: process.env.ETHEREUM_TESTNET_ETHERSCAN_API_KEY || "",
      polygonMumbai: process.env.POLYGON_TESTNET_ETHERSCAN_API_KEY || "",
      polygon: process.env.POLYGON_MAINNET_ETHERSCAN_API_KEY || "",
      base_mainnet: process.env.BASE_MAINNET_ETHERSCAN_API_KEY || "",
      base_testnet: process.env.BASE_TESTNET_ETHERSCAN_API_KEY || "",
      celo_mainnet: process.env.CELO_MAINNET_ETHERSCAN_API_KEY || "",
      celo_testnet: process.env.CELO_TESTNET_ETHERSCAN_API_KEY || "",
    },
    customChains: [
      {
        network: "base_mainnet",
        chainId: 8453,
        urls: {
          apiURL: "https://api.basescan.org/api",
          browserURL: "https://basescan.org",
        },
      },
      {
        network: "base_testnet",
        chainId: 84532,
        urls: {
          apiURL: "https://api-sepolia.basescan.org/api",
          browserURL: "https://sepolia.basescan.org/",
        },
      },
      {
        network: "celo_mainnet",
        chainId: 42220,
        urls: {
          apiURL: "https://api.celoscan.io/api",
          browserURL: "https://celoscan.io",
        },
      },
      {
        network: "celo_testnet",
        chainId: 44787,
        urls: {
          apiURL: "https://api-alfajores.celoscan.io/api",
          browserURL: "https://alfajores.celoscan.io",
        },
      },
    ],
  },
};

export default hhconfig;
