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
# - Compiler version: 0.8.20
# - Optimization: Enabled
# - Optimization runs: 200

CW_CONTRACT_OUTPUT_PATH='build/contracts'
CW_CONTRACT_DART_ACCOUNT_OUTPUT_PATH='contracts/accounts'
CW_CONTRACT_DART_APP_OUTPUT_PATH='contracts/apps'
CW_CONTRACT_DART_STD_OUTPUT_PATH='contracts/standards'
CW_CONTRACT_DART_EXT_OUTPUT_PATH='contracts/external'
CW_CONTRACT_PKG_PATH='pkg/contracts'
CW_CONTRACT_ACCOUNT_PATH='contracts/accounts'
CW_CONTRACT_APP_PATH='contracts/apps'
CW_CONTRACT_STD_PATH='contracts/standards'
CW_CONTRACT_EXT_PATH='contracts/external'

# List of contracts to compile (package,file)
# package = golang package name
# file = solidity file name
# TODO: fix contract compilation to not have duplicate package names
CW_ACCOUNT_CONTRACTS=('account,Account' \ 
        'accfactory,AccountFactory' \
        'tokenEntryPoint,TokenEntryPoint')

CW_APP_CONTRACTS=('grfactory,GratitudeTokenFactory' \
        'gratitude,GratitudeToken' \
        'profile,Profile' \
        'regensToken,RegensUniteTokens')

CW_STD_CONTRACTS=('erc20,ERC20' \
        'erc721,ERC721' \
        'erc1155,ERC1155')

CW_EXTERNAL_CONTRACTS=('simpleaccountfactory,SimpleAccountFactory' \
        'derc20,DERC20' \
        'simpleaccount,SimpleAccount' \
        'entrypoint,EntryPoint')

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

rm -rf "lib/$CW_CONTRACT_DART_STD_OUTPUT_PATH"
mkdir "lib/$CW_CONTRACT_DART_STD_OUTPUT_PATH"

rm -rf "lib/$CW_CONTRACT_DART_EXT_OUTPUT_PATH"
mkdir "lib/$CW_CONTRACT_DART_EXT_OUTPUT_PATH"

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

rm -rf "lib/standards.dart"
touch "lib/standards.dart"
echo "/// Account contracts" >> "lib/standards.dart"
echo "///" >> "lib/standards.dart"
echo "/// This file is auto-generated, edit scripts/compile.sh to modify" >> "lib/standards.dart"
echo "library;" >> "lib/standards.dart"
echo "" >> "lib/standards.dart"

rm -rf "lib/external.dart"
touch "lib/external.dart"
echo "/// External contracts" >> "lib/external.dart"
echo "///" >> "lib/external.dart"
echo "/// This file is auto-generated, edit scripts/compile.sh to modify" >> "lib/external.dart"
echo "library;" >> "lib/external.dart"
echo "" >> "lib/external.dart"

rm -rf "lib/smartcontracts.dart"
touch "lib/smartcontracts.dart"
echo "/// ERC4337 Smart Contracts for Citizen Wallet App" >> "lib/smartcontracts.dart"
echo "///" >> "lib/smartcontracts.dart"
echo "/// This file is auto-generated, edit scripts/compile.sh to modify" >> "lib/smartcontracts.dart"
echo "library;" >> "lib/smartcontracts.dart"
echo "" >> "lib/smartcontracts.dart"
echo "export 'accounts.dart';" >> "lib/smartcontracts.dart"
echo "export 'apps.dart';" >> "lib/smartcontracts.dart"
echo "export 'standards.dart';" >> "lib/smartcontracts.dart"
echo "export 'external.dart';" >> "lib/smartcontracts.dart"

# Compile the contracts
echo "[account] Compiling the contracts..."

for contract in ${CW_ACCOUNT_CONTRACTS[@]} ; 
    do
        while IFS=',' read -r pkg file; 
            do
                echo "Compiling $CW_CONTRACT_ACCOUNT_PATH/$file.sol..."

                mkdir $CW_CONTRACT_OUTPUT_PATH/$pkg

                # https://ethereum.stackexchange.com/a/84719
                docker run -v $(pwd):/root --platform linux/amd64  ethereum/solc:0.8.20 --evm-version paris --abi --bin --optimize --optimize-runs 200 --allow-paths . --include-path root/node_modules/ --base-path ./root --overwrite --output-dir "root/$CW_CONTRACT_OUTPUT_PATH/$pkg/" "root/$CW_CONTRACT_ACCOUNT_PATH/$file.sol"
                cp "$CW_CONTRACT_OUTPUT_PATH/$pkg/$file.abi" "lib/$CW_CONTRACT_DART_ACCOUNT_OUTPUT_PATH/$file.abi.json"
                echo "export '$CW_CONTRACT_DART_ACCOUNT_OUTPUT_PATH/$file.g.dart' hide AdminChanged, BeaconUpgraded, Initialized, Upgraded, OwnershipTransferred;" >> "lib/accounts.dart"
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

                # https://ethereum.stackexchange.com/a/84719
                docker run -v $(pwd):/root --platform linux/amd64 ethereum/solc:0.8.20 --evm-version paris --abi --bin --optimize --optimize-runs 200 --allow-paths . --include-path root/node_modules/ --base-path ./root --overwrite --output-dir "root/$CW_CONTRACT_OUTPUT_PATH/$pkg/" "root/$CW_CONTRACT_APP_PATH/$file.sol"
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

# Compile the contracts
echo "[standards] Compiling the contracts..."

for contract in ${CW_STD_CONTRACTS[@]} ; 
    do
        while IFS=',' read -r pkg file; 
            do
                echo "Compiling $CW_CONTRACT_STD_PATH/$file.sol..."

                mkdir $CW_CONTRACT_OUTPUT_PATH/$pkg

                # https://ethereum.stackexchange.com/a/84719
                docker run -v $(pwd):/root --platform linux/amd64 ethereum/solc:0.8.20 --evm-version paris --abi --bin --optimize --optimize-runs 200 --allow-paths . --include-path root/node_modules/ --base-path ./root --overwrite --output-dir "root/$CW_CONTRACT_OUTPUT_PATH/$pkg/" "root/$CW_CONTRACT_STD_PATH/$file.sol"
                cp "$CW_CONTRACT_OUTPUT_PATH/$pkg/$file.abi" "lib/$CW_CONTRACT_DART_STD_OUTPUT_PATH/$file.abi.json"
                echo "export '$CW_CONTRACT_DART_STD_OUTPUT_PATH/$file.g.dart';" >> "lib/standards.dart"
                echo "[.abi] $CW_CONTRACT_STD_PATH/$file ✅";
                echo "[.bin] $CW_CONTRACT_STD_PATH/$file ✅";

                mkdir $CW_CONTRACT_PKG_PATH/$pkg
                abigen --bin="$CW_CONTRACT_OUTPUT_PATH/$pkg/$file.bin" --abi="$CW_CONTRACT_OUTPUT_PATH/$pkg/$file.abi" --pkg="$pkg" --out="$CW_CONTRACT_PKG_PATH/$pkg/$file.go"
                echo "[.go] $CW_CONTRACT_STD_PATH/$file package $pkg ✅";
        done <<< "$contract"
    done

echo "[standards] Done compiling the contracts."

# Compile the contracts
echo "[external] Compiling the contracts..."

for contract in ${CW_EXTERNAL_CONTRACTS[@]} ; 
    do
        while IFS=',' read -r pkg file; 
            do
                echo "Compiling $CW_CONTRACT_EXT_PATH/$file.sol..."

                mkdir $CW_CONTRACT_OUTPUT_PATH/$pkg

                cp "$CW_CONTRACT_EXT_PATH/$file.abi" "lib/$CW_CONTRACT_DART_EXT_OUTPUT_PATH/$file.abi.json"
                echo "export '$CW_CONTRACT_DART_EXT_OUTPUT_PATH/$file.g.dart';" >> "lib/external.dart"
                echo "[.abi] $CW_CONTRACT_EXT_PATH/$file ✅";
                echo "[.bin] $CW_CONTRACT_EXT_PATH/$file ✅";

                mkdir $CW_CONTRACT_PKG_PATH/$pkg
                abigen --bin="$CW_CONTRACT_EXT_PATH/$file.bin" --abi="$CW_CONTRACT_EXT_PATH/$file.abi" --pkg="$pkg" --out="$CW_CONTRACT_PKG_PATH/$pkg/$file.go"
                echo "[.go] $CW_CONTRACT_EXT_PATH/$file package $pkg ✅";
        done <<< "$contract"
    done

echo "[external] Done compiling the contracts."

# Compile dart bindings for contracts
echo "Compiling dart bindings for contracts..."

dart run build_runner build --delete-conflicting-outputs

# Install dependencies
echo "Installing dependencies..."
go get ./...
echo "Done installing dependencies."

