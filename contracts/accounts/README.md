# Hybrid ERC4337

[![Compile Contracts](https://github.com/daobrussels/smartcontracts/actions/workflows/compile.yml/badge.svg)](https://github.com/daobrussels/smartcontracts/actions/workflows/compile.yml)

## Issues with ERC4337

### Reliance on off-chain infrastructure (bundlers)

**Traditional ERC4337:** A bundler + access to a full node is necessary in order to process **User Operations**. This increases the cost of processing transactions and adds an extra layer that can potentially fail.

Our approach allows user operations to be submitted directly to the **TokenEntryPoint** contract and be processed on chain.

The only pre-requisite is for a sponsor to sign the user operation before submission.

### Flexibility

**Traditional ERC4337:** If there is no bundler provider on your chain, it's impossible to use 4337. Or you need to spend a considerable amount of time spinning up a full node + running a bundler against it. Not your core business!

Our approach is on-chain and EVM-compatible. We have also made sure that the Accounts are still ERC4337-compatible.

It only requires publishing Smart Contracts and dealing with signed User Operations.

### High gas fees

**Traditional ERC4337:** The **EntryPoint** contract runs a lot of simulations on the user operation in order to control gas fee consumption. This, among other things, increases gas fee consumption greatly compared to regular transactions.

Our approach simplifies execution and verification into a single contract and reduces gas fee consumption.

## TokenEntryPoint

The `TokenEntryPoint` contract is a Solidity contract used to authorize and execute user operations. It inherits from several OpenZeppelin contracts and uses ECDSA for signature verification and UserOperationLib for user operation handling.

### Key Features

- **User Operation Handling**: The contract uses the `UserOperationLib` library to handle user operations. It can execute a batch of user operations with the `handleOps` function.

- **Signature Verification**: The contract uses ECDSA for signature verification. It verifies the nonce, account, and paymaster signature of each user operation before execution.

- **Paymaster Sponsor**: The contract includes a paymaster sponsor, which can be updated with the `updatePaymasterSponsor` function.

### Usage

To execute a batch of user operations, call the `handleOps` function with an array of `UserOperation` structs and a beneficiary address. The function will execute each operation in the array.

To update the paymaster sponsor, call the `updatePaymasterSponsor` function with the new sponsor's address. The function will update the sponsor and emit a `PaymasterSponsorUpdated` event.

### Dependencies

This contract depends on several OpenZeppelin contracts:

- `Paymaster`: Used for paymaster functionality.
- `Initializable`: Used for initialization functionality.
- `ReentrancyGuardUpgradeable`: Used to prevent reentrancy attacks.
- `OwnableUpgradeable`: Used for owner functionality.
- `UUPSUpgradeable`: Used for upgradeability functionality.

It also uses ECDSA for signature verification and UserOperationLib for user operation handling.

### References

This contract is based on the [EntryPoint](https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/core/EntryPoint.sol) and [VerifyingPaymaster](https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/samples/VerifyingPaymaster.sol) contracts from the eth-infinitism/account-abstraction GitHub repository.

## AccountFactory

The AccountFactory contract is a Solidity contract used for creating new accounts and calculating their counterfactual addresses. It leverages the OpenZeppelin Create2 and ERC1967Proxy contracts, and is part of a larger system implementing account abstraction.

### Key Features

Immutable Account Implementation: The contract creates an immutable instance of the Account contract upon deployment. This instance serves as the implementation for all accounts created by the factory.

Account Creation: The createAccount function is used to create a new account. It emits an AccountCreated event with the address of the new account. If the account already exists, it simply returns the existing account's address.

Counterfactual Address Calculation: The getAddress function calculates the counterfactual address of an account. This is the address that the account would have if it were created with the createAccount function.

### Usage

To create a new account, call the createAccount function with the owner's address and a salt. The function will return the address of the new account. If the account already exists, it will return the existing account's address.

To calculate the counterfactual address of an account, call the getAddress function with the owner's address and a salt. The function will return the counterfactual address of the account.

### Dependencies

This contract depends on several OpenZeppelin contracts:

Create2: Used for calculating counterfactual addresses.
ERC1967Proxy: Used for creating new accounts.
Account: The implementation for the accounts created by the factory.
It also depends on the IEntryPoint and ITokenEntryPoint interfaces from the @account-abstraction/contracts package.

### References

This contract is based on the [SimpleAccountFactory](https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/samples/SimpleAccountFactory.sol) contract from the eth-infinitism/account-abstraction GitHub repository.

## Account

The `Account` contract is a Solidity contract that represents an account capable of executing transactions and storing funds in an entry point contract. It implements the ERC1271 standard for signature validation and is upgradeable using the UUPSUpgradeable pattern. The account owner can execute transactions directly or through the entry point contract, and can allow an authorizer contract to execute transactions on its behalf.

### Key Features

- **ERC1271 Standard**: The contract implements the ERC1271 standard for signature validation.

- **Upgradeable**: The contract is upgradeable using the UUPSUpgradeable pattern.

- **Transaction Execution**: The account owner can execute transactions directly or through the entry point contract. An authorizer contract can also execute transactions on behalf of the account owner.

- **Funds Storage**: The contract can store funds in an entry point contract.

### Usage

To execute a transaction, call the `execute` function with the destination address, value, and function data. The function can be called directly by the owner, or by the entry point or authorizer contract.

To execute a batch of transactions, call the `executeBatch` function with arrays of destination addresses, values, and function data. The function can be called directly by the owner, or by the entry point or authorizer contract.

To deposit more funds for the account in the entry point, call the `addDeposit` function with the amount to deposit.

To withdraw funds from the account's deposit, call the `withdrawDepositTo` function with the withdrawal address and amount.

### Dependencies

This contract depends on several OpenZeppelin contracts:

- `IERC1271`: Used for signature validation.
- `BaseAccount`: Used for base account functionality.
- `TokenCallbackHandler`: Used for token callback handling.
- `Initializable`: Used for initialization functionality.
- `OwnableUpgradeable`: Used for owner functionality.
- `UUPSUpgradeable`: Used for upgradeability functionality.

It also uses ECDSA for signature recovery and `IEntryPoint` and `ITokenEntryPoint` interfaces for entry point and authorizer functionality.

### References

This contract is based on the [SimpleAccount](https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/samples/SimpleAccount.sol) contract from the eth-infinitism/account-abstraction GitHub repository.

## Test

Try running some of the following tasks:

```shell
npx hardhat test ./test/Account.test.ts
```
