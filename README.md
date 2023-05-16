# Citizen Smart Contract

[![Compile Contracts](https://github.com/daobrussels/smartcontracts/actions/workflows/compile.yml/badge.svg)](https://github.com/daobrussels/smartcontracts/actions/workflows/compile.yml)

## GratitudeToken

UUPS upgradable ERC20 token to mint and send gratitudes tokens that are not transferrables by the recipient (only the owner can transfer them)

## Test

Try running some of the following tasks:

```shell
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat run scripts/deploy.ts
```
