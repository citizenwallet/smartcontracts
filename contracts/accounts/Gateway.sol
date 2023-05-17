// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@account-abstraction/contracts/core/EntryPoint.sol";

// Gateway,
contract Gateway is EntryPoint, Initializable, Ownable {
    using ECDSA for bytes32;

    event GatewayInitialized(address indexed owner); // Event to indicate when the Account contract has been initialized

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    // Internal function to initialize the contract with an owner address
    function _initialize(address _owner) public initializer {
        transferOwnership(_owner); // Transfer the ownership of the contract to the specified address
        emit GatewayInitialized(_owner); // Emit an event to indicate that the Account contract has been initialized with the entry point contract address and the owner address
    }

    function setOwner(address _owner) public onlyOwner {
        transferOwnership(_owner);
    }
}
