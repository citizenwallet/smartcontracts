# Smart Contracts for the Citizen Wallet

## Set up

Make sure you set up your environment variables in the `.env.local` file after cloning the repo.

```
git clone git@github.com:citizenwallet/smartcontracts.git
cd smartcontracts
npm install
cp .env.example .env.local
vim .env.local
```

## Deploy a new community token

1. Deploy a community token

Make a copy of a template (e.g. `contracts/RegensUniteToken.sol`, a simple ERC20, no pre-mint, upgradeable), edit the name and the symbol (line 32) and edit the script `scripts/deploy-upgradeable.js` to change the name of your contract (line 5). You may also change the initial addresses that have the `MINTER` role (by default, `MINTER1_ADDRESS` and `MINTER2_ADDRESS` defined in your `.env.local` file).

Then run:

```shell
npx hardhat run scripts/deploy-upgradeable.js --network polygon
```

Take note of the address of the newly deployed contract, you will need it later.
Now you can deploy the community entry point and paymaster:

```shell
npx hardhat run scripts/deploy-community.ts --network polygon
```

`--network` can be `polygon`, `polygonMumbai`, `celo`

## Tests

Try running some of the following tasks:

```shell
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat run scripts/deploy.ts
```
