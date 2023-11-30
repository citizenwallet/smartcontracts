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
      // timeoutBlocks: 200,
      gas: 12400000,
      // gasPrice: 900000000000,
      throwOnTransactionFailures: true,
      throwOnCallFailures: true,
      allowUnlimitedContractSize: true,
      // blockGasLimit: 0x1fffffffffffff,
      // timeout: 1800000,
    },
    polygon_testnet: {
      url: process.env.POLYGON_TESTNET_RPC_URL,
      accounts: [process.env.DEPLOYER_PRIVATE_KEY || "0x0"],
      gasPrice: 90000000000, // this is 90 Gwei
      timeout: 999999,
      // timeoutBlocks: 200,
      gas: 12400000,
      // gasPrice: 900000000000,
      throwOnTransactionFailures: true,
      throwOnCallFailures: true,
      allowUnlimitedContractSize: true,
      // blockGasLimit: 0x1fffffffffffff,
      // timeout: 1800000,
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
};

export default hhconfig;
