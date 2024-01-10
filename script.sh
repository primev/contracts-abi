#!/bin/bash

# Define the base directory where the .json files are located
BASE_DIR="$HOME/.primev/contracts/out"

# Create the abi directory if it doesn't exist
ABI_DIR="./abi"
mkdir -p "$ABI_DIR"

# Function to extract and save the ABI
extract_and_save_abi() {
    local json_file="$1"
    local abi_file="$2"
    jq .abi "$json_file" > "$abi_file"
}

# Extract ABI for BidderRegistry.json
extract_and_save_abi "$BASE_DIR/BidderRegistry.sol/BidderRegistry.json" "$ABI_DIR/BidderRegistry.abi"

# Extract ABI for ProviderRegistry.json
extract_and_save_abi "$BASE_DIR/ProviderRegistry.sol/ProviderRegistry.json" "$ABI_DIR/ProviderRegistry.abi"

# Extract ABI for Oracle.json
extract_and_save_abi "$BASE_DIR/Oracle.sol/Oracle.json" "$ABI_DIR/Oracle.abi"

# Extract ABI for PreConfCommitmentStore.json
extract_and_save_abi "$BASE_DIR/PreConfirmations.sol/PreConfCommitmentStore.json" "$ABI_DIR/PreConfCommitmentStore.abi"

echo "ABI files extracted successfully."

