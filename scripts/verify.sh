#!/bin/bash

npx hardhat verify --contract contracts/accounts/Gateway.sol:Gateway --network chainstack 0xb0f48da8710776e63139c4894d40aece571d1af7
npx hardhat verify --contract contracts/accounts/Paymaster.sol:Paymaster --network chainstack 0xc1cc87b84d787657840ff09a3ebacc4911e50a24 "0xb0f48da8710776e63139c4894d40aece571d1af7"
npx hardhat verify --contract contracts/accounts/AccountFactory.sol:AccountFactory --network chainstack 0xa3c096b6562eb926aaac421ffab592aa4918a408 "0xb0f48da8710776e63139c4894d40aece571d1af7"
npx hardhat verify --contract contracts/apps/ProfileFactory.sol:ProfileFactory --network chainstack 0x8d4b5168c9b5ff974c6b7dc8c8fdda18ffa20a06 "0xb0f48da8710776e63139c4894d40aece571d1af7"