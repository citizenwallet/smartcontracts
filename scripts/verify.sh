#!/bin/bash

npx hardhat verify --contract contracts/accounts/Paymaster.sol:Paymaster --network polygon 0x3Ff9D40B8Ac651689Ded5a9BAe00486d00dCeCe3
npx hardhat verify --contract contracts/accounts/TokenEntryPoint.sol:TokenEntryPoint --network polygon 0xab0628e3405a3f79fB5ab03f5Aa9898282309e51
npx hardhat verify --contract contracts/accounts/AccountFactory.sol:AccountFactory --network polygon 0xD7BcA175B68E3E1E34d79B296797D5D9518F84Dc "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789" "0xab0628e3405a3f79fB5ab03f5Aa9898282309e51"