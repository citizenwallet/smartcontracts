#!/bin/bash

npx hardhat verify --contract contracts/accounts/Paymaster.sol:Paymaster --network polygon 0xE13751E876D78F6e03941A7763d22291d53602B9
npx hardhat verify --contract contracts/accounts/TokenEntryPoint.sol:TokenEntryPoint --network polygon 0x5e107a280b2bE5fE80Ad9355D78a5686e770C569 "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789"
npx hardhat verify --contract contracts/accounts/AccountFactory.sol:AccountFactory --network polygon 0x5d57fab36C284b3Ab5Ba86d4A9a26EC367B8caa6 "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789" "0x5e107a280b2bE5fE80Ad9355D78a5686e770C569"