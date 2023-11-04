#!/bin/bash

npx hardhat verify --contract contracts/accounts/TokenEntryPoint.sol:TokenEntryPoint --network polygon 0x4f7d321062B26B12caB3c1E6104eE269697932cF
npx hardhat verify --contract contracts/accounts/AccountFactory.sol:AccountFactory --network polygon 0xC5B479616ef0dcA586e45443c95F50dC7537DE1b "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789" "0x4f7d321062B26B12caB3c1E6104eE269697932cF"