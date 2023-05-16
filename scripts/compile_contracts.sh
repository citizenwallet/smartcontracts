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
CW_CONTRACT_PKG_PATH='pkg/contracts'
CW_CONTRACT_BASE_PATH='contracts'

# List of contracts to compile (package,file)
# package = golang package name
# file = solidity file name
CW_CONTRACTS=('account,Account' \ 
        'gateway,Gateway' \
        'gratitudeToken,GratitudeToken')

# Clean build folder
echo "Cleaning build folder..."

rm -rf "$CW_CONTRACT_OUTPUT_PATH/*"

# Clean pkg folder
echo "Cleaning pkg folder..."

rm -rf "$CW_CONTRACT_PKG_PATH/*"

# Compile the contracts
echo "Compiling the contracts..."

for contract in ${CW_CONTRACTS[@]} ; 
    do
        while IFS=',' read -r pkg file; 
            do
                mkdir $CW_CONTRACT_OUTPUT_PATH/$pkg
                solc --abi --bin --optimize --optimize-runs 200 --allow-paths . --include-path node_modules/ --base-path . --overwrite --output-dir "$CW_CONTRACT_OUTPUT_PATH/$pkg/" "$CW_CONTRACT_BASE_PATH/$file.sol"
                echo "$CW_CONTRACT_BASE_PATH/$file ✅";
        done <<< "$contract"
    done

echo "Done compiling the contracts."

# Generate go interfaces
echo "Generating go interfaces..."

for contract in ${CW_CONTRACTS[@]} ; 
    do
        while IFS=',' read -r pkg file; do
            mkdir $CW_CONTRACT_PKG_PATH/$pkg
            abigen --bin="$CW_CONTRACT_OUTPUT_PATH/$pkg/$file.bin" --abi="$CW_CONTRACT_OUTPUT_PATH/$pkg/$file.abi" --pkg="$pkg" --out="$CW_CONTRACT_PKG_PATH/$pkg/$file.go"
            echo "$CW_CONTRACT_BASE_PATH/$file package $pkg ✅";
        done <<< "$contract"
    done
echo "Done generating go interfaces."

# Install dependencies
echo "Generating go interfaces..."
go get ./...
echo "Done generating go interfaces."

