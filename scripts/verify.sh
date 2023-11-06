#!/bin/bash

npx hardhat verify --contract contracts/accounts/Paymaster.sol:Paymaster --network polygon 0x1462532563ec07F240547fDd66FA317Ae9419269
npx hardhat verify --contract contracts/accounts/TokenEntryPoint.sol:TokenEntryPoint --network polygon 0x74d6490f420b7cDdC53F3e1CD11Bf064B389290a
npx hardhat verify --contract contracts/accounts/AccountFactory.sol:AccountFactory --network polygon 0x48d8b66E787F9253812A89907DD9e3FefA0dAd78 "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789" "0x74d6490f420b7cDdC53F3e1CD11Bf064B389290a"