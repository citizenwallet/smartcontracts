module.exports = {
  contracts_directory: "./contracts",
  contracts_build_directory: "./output",
  dashboard: {
    port: 24012,
    host: "localhost",
    verbose: false,
  },
  networks: {
    development: {
      host: "127.0.0.1",
      port: 8545,
      network_id: "1682515751360", // Match any network id
      gas: 30000000,
      gasPrice: 2000000000,
    },
  },
  compilers: {
    solc: {
      version: "^0.8.0",
      settings: {
        optimizer: {
          enabled: true,
          runs: 200,
        },
      },
    },
  },
};
