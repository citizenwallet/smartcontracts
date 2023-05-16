// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@account-abstraction/contracts/core/BaseAccount.sol";

import "./Account/TokenCallbackHandler.sol";

// https://github.com/guizostudios/ERC-4337/blob/main/contracts/SimpleAccount.sol
// Account,
contract Account is BaseAccount, TokenCallbackHandler, Ownable {
    using ECDSA for bytes32;

    IEntryPoint private immutable _entryPoint; // Private immutable variable to store the entry point contract address

    event AccountInitialized(
        IEntryPoint indexed entryPoint,
        address indexed owner
    ); // Event to indicate when the Account contract has been initialized

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor(IEntryPoint anEntryPoint, address _owner) {
        _entryPoint = anEntryPoint; // Constructor to set the entry point contract address during deployment
        _initialize(_owner); // Call the internal _initialize function with the provided input parameter
    }

    // Internal function to initialize the contract with an owner address
    function _initialize(address _owner) internal virtual {
        emit AccountInitialized(_entryPoint, _owner); // Emit an event to indicate that the Account contract has been initialized with the entry point contract address and the owner address
    }

    // Modifier to require the function call to be made from the entry point contract or the owner
    function _requireFromEntryPointOrOwner() internal view {
        require(
            msg.sender == address(entryPoint()) || msg.sender == owner(),
            "account: not Owner or EntryPoint"
        ); // Check if the caller is the entry point contract or the owner
    }

    // Override the template method from BaseAccount to validate the signature of a user operation
    function _validateSignature(
        UserOperation calldata userOp,
        bytes32 userOpHash
    ) internal virtual override returns (uint256 validationData) {
        bytes32 hash = userOpHash.toEthSignedMessageHash(); // Convert the user operation hash to an Ethereum signed message hash
        if (owner() != hash.recover(userOp.signature))
            // Check if the owner address matches the recovered address from the signature
            return SIG_VALIDATION_FAILED; // Return an error code if the signature validation fails
        return 0; // Return success code if the signature validation is successful
    }

    /// @inheritdoc BaseAccount
    function entryPoint() public view virtual override returns (IEntryPoint) {
        // Function to get the entry point contract address
        return _entryPoint;
    }

    // solhint-disable-next-line no-empty-blocks
    receive() external payable {} // Fallback function to receive Ether

    /**
     * execute a transaction (called directly from owner, or by entryPoint)
     */
    function execute(
        address dest,
        uint256 value,
        bytes calldata func
    ) external {
        _requireFromEntryPointOrOwner(); // Function modifier to check if the caller is the entry point contract or the owner
        _call(dest, value, func); // Internal function to call another contract with specified address, value, and function data
    }

    /**
     * execute a sequence of transactions
     */
    function executeBatch(
        address[] calldata dest,
        bytes[] calldata func
    ) external {
        _requireFromEntryPointOrOwner(); // Function modifier to check if the caller is the entry point contract or the owner
        require(dest.length == func.length, "wrong array lengths"); // Check if the lengths of destination and function data arrays are the same
        for (uint256 i = 0; i < dest.length; i++) {
            _call(dest[i], 0, func[i]); // Internal function to call multiple contracts in a batch with specified addresses and function data
        }
    }

    // Internal function to call another contract
    function _call(address target, uint256 value, bytes memory data) internal {
        (bool success, bytes memory result) = target.call{value: value}(data); // Call the contract with the provided address, value, and function data
        if (!success) {
            // Check if the call was unsuccessful
            assembly {
                // Use assembly to revert with the error message from the failed call
                revert(add(result, 32), mload(result))
            }
        }
    }

    // Get the current deposit balance in the entry point contract
    function getDeposit() public view returns (uint256) {
        return entryPoint().balanceOf(address(this)); // Call the entry point contract's balanceOf function to get the balance of this contract's address
    }

    // Deposit more funds into the entry point contract for this account
    function addDeposit() public payable {
        entryPoint().depositTo{value: msg.value}(address(this)); // Call the entry point contract's depositTo function with the provided value and this contract's address as the target
    }

    // Withdraw funds from the account's deposit in the entry point contract
    function withdrawDepositTo(
        address payable withdrawAddress,
        uint256 amount
    ) public onlyOwner {
        entryPoint().withdrawTo(withdrawAddress, amount); // Call the entry point contract's withdrawTo function with the provided withdrawal address and amount
    }

    function setOwner(address _owner) public onlyOwner {
        transferOwnership(_owner);
    }
}
