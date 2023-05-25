#!/bin/bash

# This script compiles the contracts and generates the ABI and BIN files
# for the contracts. The ABI and BIN files are used by the deploy_contracts.sh
# script to deploy the contracts to the blockchain.

# The script assumes that the contracts are located in the contracts/ directory
# and that the contracts are named as follows:
# - Store.sol

# The script also assumes that the contracts are written in Solidity and that
# the Solidity compiler is installed on the system.

# The script also assumes that the contracts are compiled with the following
# compiler settings:
# - Compiler version: 0.8.19
# - Optimization: Enabled
# - Optimization runs: 200

CW_CONTRACT_OUTPUT_PATH='build/contracts'
CW_CONTRACT_DART_ACCOUNT_OUTPUT_PATH='contracts/accounts'
CW_CONTRACT_DART_APP_OUTPUT_PATH='contracts/apps'
CW_CONTRACT_PKG_PATH='pkg/contracts'
CW_CONTRACT_ACCOUNT_PATH='contracts/accounts'
CW_CONTRACT_APP_PATH='contracts/apps'

# List of contracts to compile (package,file)
# package = golang package name
# file = solidity file name
CW_ACCOUNT_CONTRACTS=('account,Account' \ 
        'accfactory,AccountFactory' \ 
        'gateway,Gateway' \
        'paymaster,Paymaster')

CW_APP_CONTRACTS=('grfactory,GratitudeTokenFactory' \
        'gratitude,GratitudeToken' \
        'profactory,ProfileFactory' \
        'profile,Profile' \
        'regensToken,RegensUniteTokens')

# Clean build folder
echo "Cleaning build folder..."

rm -rf "$CW_CONTRACT_OUTPUT_PATH"
mkdir "$CW_CONTRACT_OUTPUT_PATH"

# Clean pkg folder
echo "Cleaning pkg folder..."

rm -rf "$CW_CONTRACT_PKG_PATH"
mkdir "$CW_CONTRACT_PKG_PATH"

# Clean lib folder
echo "Cleaning lib folder..."

rm -rf "lib/$CW_CONTRACT_DART_ACCOUNT_OUTPUT_PATH"
mkdir "lib/$CW_CONTRACT_DART_ACCOUNT_OUTPUT_PATH"

rm -rf "lib/$CW_CONTRACT_DART_APP_OUTPUT_PATH"
mkdir "lib/$CW_CONTRACT_DART_APP_OUTPUT_PATH"

# Generate export file
echo "Generating dart library export file..."

rm -rf "lib/accounts.dart"
touch "lib/accounts.dart"
echo "/// Account contracts" >> "lib/accounts.dart"
echo "///" >> "lib/accounts.dart"
echo "/// This file is auto-generated, edit scripts/compile.sh to modify" >> "lib/accounts.dart"
echo "library;" >> "lib/accounts.dart"
echo "" >> "lib/accounts.dart"

rm -rf "lib/apps.dart"
touch "lib/apps.dart"
echo "/// Account contracts" >> "lib/apps.dart"
echo "///" >> "lib/apps.dart"
echo "/// This file is auto-generated, edit scripts/compile.sh to modify" >> "lib/apps.dart"
echo "library;" >> "lib/apps.dart"
echo "" >> "lib/apps.dart"

rm -rf "lib/smartcontracts.dart"
touch "lib/smartcontracts.dart"
echo "/// ERC4337 Smart Contracts for Citizen Wallet App" >> "lib/smartcontracts.dart"
echo "///" >> "lib/smartcontracts.dart"
echo "/// This file is auto-generated, edit scripts/compile.sh to modify" >> "lib/smartcontracts.dart"
echo "library;" >> "lib/smartcontracts.dart"
echo "" >> "lib/smartcontracts.dart"
echo "export 'accounts.dart';" >> "lib/smartcontracts.dart"
echo "export 'apps.dart';" >> "lib/smartcontracts.dart"

# Compile the contracts
echo "[account] Compiling the contracts..."

for contract in ${CW_ACCOUNT_CONTRACTS[@]} ; 
    do
        while IFS=',' read -r pkg file; 
            do
                echo "Compiling $CW_CONTRACT_ACCOUNT_PATH/$file.sol..."

                mkdir $CW_CONTRACT_OUTPUT_PATH/$pkg
                solc --evm-version paris --abi --bin --optimize --optimize-runs 200 --allow-paths . --include-path node_modules/ --base-path . --overwrite --output-dir "$CW_CONTRACT_OUTPUT_PATH/$pkg/" "$CW_CONTRACT_ACCOUNT_PATH/$file.sol"
                cp "$CW_CONTRACT_OUTPUT_PATH/$pkg/$file.abi" "lib/$CW_CONTRACT_DART_ACCOUNT_OUTPUT_PATH/$file.abi.json"
                echo "export '$CW_CONTRACT_DART_ACCOUNT_OUTPUT_PATH/$file.g.dart';" >> "lib/accounts.dart"
                echo "[.abi] $CW_CONTRACT_ACCOUNT_PATH/$file ✅";
                echo "[.bin] $CW_CONTRACT_ACCOUNT_PATH/$file ✅";

                mkdir $CW_CONTRACT_PKG_PATH/$pkg
                abigen --bin="$CW_CONTRACT_OUTPUT_PATH/$pkg/$file.bin" --abi="$CW_CONTRACT_OUTPUT_PATH/$pkg/$file.abi" --pkg="$pkg" --out="$CW_CONTRACT_PKG_PATH/$pkg/$file.go"
                echo "[.go] $CW_CONTRACT_ACCOUNT_PATH/$file package $pkg ✅";
        done <<< "$contract"
    done

echo "[account] Done compiling the contracts."

# Compile the contracts
echo "[app] Compiling the contracts..."

for contract in ${CW_APP_CONTRACTS[@]} ; 
    do
        while IFS=',' read -r pkg file; 
            do
                echo "Compiling $CW_CONTRACT_APP_PATH/$file.sol..."

                mkdir $CW_CONTRACT_OUTPUT_PATH/$pkg
                solc --evm-version paris --abi --bin --optimize --optimize-runs 200 --allow-paths . --include-path node_modules/ --base-path . --overwrite --output-dir "$CW_CONTRACT_OUTPUT_PATH/$pkg/" "$CW_CONTRACT_APP_PATH/$file.sol"
                cp "$CW_CONTRACT_OUTPUT_PATH/$pkg/$file.abi" "lib/$CW_CONTRACT_DART_APP_OUTPUT_PATH/$file.abi.json"
                echo "export '$CW_CONTRACT_DART_APP_OUTPUT_PATH/$file.g.dart';" >> "lib/apps.dart"
                echo "[.abi] $CW_CONTRACT_APP_PATH/$file ✅";
                echo "[.bin] $CW_CONTRACT_APP_PATH/$file ✅";

                mkdir $CW_CONTRACT_PKG_PATH/$pkg
                abigen --bin="$CW_CONTRACT_OUTPUT_PATH/$pkg/$file.bin" --abi="$CW_CONTRACT_OUTPUT_PATH/$pkg/$file.abi" --pkg="$pkg" --out="$CW_CONTRACT_PKG_PATH/$pkg/$file.go"
                echo "[.go] $CW_CONTRACT_APP_PATH/$file package $pkg ✅";
        done <<< "$contract"
    done

echo "[app] Done compiling the contracts."

# Compile dart bindings for contracts
echo "Compiling dart bindings for contracts..."

dart run build_runner build --delete-conflicting-outputs

# Install dependencies
echo "Installing dependencies..."
go get ./...
echo "Done installing dependencies."

