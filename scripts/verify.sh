#!/bin/bash

npx hardhat verify --contract contracts/accounts/Paymaster.sol:Paymaster --network polygon 0x60CbEf32D940408f29eBAFbE07Aa1A9e05C46569
npx hardhat verify --contract contracts/accounts/TokenEntryPoint.sol:TokenEntryPoint --network polygon 0x4f0DCB82aB0787d81ee64E27dFb55a1DA90b1BFf
npx hardhat verify --contract contracts/accounts/AccountFactory.sol:AccountFactory --network polygon 0x3e8Ab3C2AB0351718E77B8a98e564faB62AbD0b3 "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789" "0x4f0DCB82aB0787d81ee64E27dFb55a1DA90b1BFf"