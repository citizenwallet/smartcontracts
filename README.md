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

1. Deploy a community token (or bring your own)

Make a copy of a template (e.g. `contracts/RegensUniteToken.sol`, a simple ERC20, no pre-mint, upgradeable), edit the name and the symbol (line 12 and line 32). You also need to define the initial addresses that will have the `MINTER` role in your `.env.local` file (`MINTER1_ADDRESS` and `MINTER2_ADDRESS`).

Then run:

```shell
npm run deploy:contract
```

Take note of the address of the newly deployed contract, you will need it later.

2. Deploy the community entry point and paymaster

```shell
npm run deploy:community
```

3. Generate the config file for your community

This is the json configuration file that describes your community (name, logo, token address, profile contract address, etc.)

See an example here:
https://github.com/citizenwallet/app/blob/4981f38d69b6243091801b3f788f22558bd6ffc6/assets/data/config2.json

Note: This needs some work to simplify the process.

## Tests

Try running some of the following tasks:

```shell
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat run scripts/deploy.ts
```
