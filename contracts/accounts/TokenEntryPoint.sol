// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/interfaces/IERC1271.sol";

import "@account-abstraction/contracts/interfaces/UserOperation.sol";
import "@account-abstraction/contracts/interfaces/IPaymaster.sol";
import "@account-abstraction/contracts/interfaces/INonceManager.sol";

import "./interfaces/IAccountFactory.sol";

/**
 * @title TokenEntryPoint
 * @dev This contract is used to authorize user operations and execute them.
 * It inherits from Paymaster, Initializable, ReentrancyGuardUpgradeable, OwnableUpgradeable, and UUPSUpgradeable.
 * It also uses ECDSA for signature verification and UserOperationLib for user operation handling.
 *
 * It is a simplified entry point contract which includes the verifying paymaster.
 *
 * https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/core/EntryPoint.sol
 * https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/samples/VerifyingPaymaster.sol
 * https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/interfaces/UserOperation.sol
 */
contract TokenEntryPoint is
    Initializable,
    ReentrancyGuardUpgradeable,
    OwnableUpgradeable,
    UUPSUpgradeable
{
    using ECDSA for bytes32;
    using UserOperationLib for UserOperation;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    // we make the owner of also the sponsor by default
    function initialize(address anOwner) public virtual initializer {
        __Ownable_init();
        __ReentrancyGuard_init();

        _initialize(anOwner);
    }

    function _initialize(address anOwner) internal virtual {
        transferOwnership(anOwner);
    }

    /**
     * @dev Executes a batch of user operations.
     * @param ops Array of UserOperation structs containing the operations to execute.
     * @param beneficiary << kept to make sure we keep the same function signature.
     * @notice This function is non-reentrant and requires at least one user operation.
     * @notice Each operation is validated for nonce, account, and paymaster signature before execution.
     */
    function handleOps(
        UserOperation[] calldata ops,
        address payable beneficiary // kept to make sure we keep the same function signature
    ) public nonReentrant {
        require(ops.length > 0, "needs at least one user operation");

        uint len = ops.length;
        for (uint i = 0; i < len; ) {
            // handle each op
            UserOperation calldata op = ops[i];

            address sender = op.getSender();

            address paymaster = _getPaymaster(op);

            // verify nonce
            _validateNonce(op, sender, paymaster);

            // verify account
            _validateAccount(op, sender);

            // verify paymaster signature
            _validatePaymasterUserOp(op, paymaster);

            // execute the op
            _call(sender, 0, op.callData);

            unchecked {
                ++i;
            }
        }
    }

    function _getPaymaster(
        UserOperation calldata op
    ) internal virtual returns (address) {
        bytes calldata paymasterAndData = op.paymasterAndData;

        // paymasterAndData must be at least 20 bytes long, and the first 20 bytes must be the paymaster address
        require(paymasterAndData.length >= 20, "invalid paymasterAndData");

        return address(bytes20(paymasterAndData[:20]));
    }

    /**
     * @dev Validates the paymaster address and data of a user operation.
     * @param op The user operation to validate.
     */
    function _validatePaymasterUserOp(
        UserOperation calldata op,
        address paymaster
    ) internal virtual {
        // verify paymasterAndData signature
        IPaymaster(paymaster).validatePaymasterUserOp(op, bytes32(0), 0);
    }

    /**
     * @dev Validates the nonce of a user operation against the nonce stored in the account.
     * @param op The user operation to validate.
     * @param sender The address of the account to validate against.
     * Requirements:
     * - The nonce in the user operation must match the nonce in the account.
     */
    function _validateNonce(
        UserOperation calldata op,
        address sender,
        address paymaster
    ) internal virtual {
        uint256 nonce = INonceManager(paymaster).getNonce(sender, 0);

        // the nonce in the user op must match the nonce in the account
        require(op.nonce == nonce, "invalid nonce");
    }

    bytes4 internal constant MAGICVALUE = 0x1626ba7e;

    /**
     * @dev Validates a user operation and its signature.
     * @param op The user operation to validate.
     * @param sender The address of the sender of the user operation.
     */
    function _validateAccount(
        UserOperation calldata op,
        address sender
    ) internal virtual {
        bytes32 ophash = _userOpHash(op).toEthSignedMessageHash();

        // call the initCode
        if (op.nonce == 0) {
            _initAccount(op, ophash);
        }

        // verify the user op signature
        IERC1271 account = IERC1271(sender);

        // the account must return the magic value according to IERC1271
        require(
            account.isValidSignature(ophash, op.signature) == MAGICVALUE,
            "invalid account signature"
        );
    }

    /**
     * @dev Initializes a new account using the provided UserOperation and ophash.
     * @param op The UserOperation which contains the initCode.
     * @param ophash The ophash to use for signature verification.
     */
    function _initAccount(
        UserOperation calldata op,
        bytes32 ophash
    ) internal virtual {
        bytes calldata initCode = op.initCode;

        // initCode must be at least 20 bytes long, and the first 20 bytes must be the factory address
        require(initCode.length >= 20, "invalid initCode");

        address factory = address(bytes20(initCode[0:20]));

        // the factory in the init code must be deployed
        require(_contractExists(factory), "invalid factory");

        address signer = ophash.recover(op.signature);

        // the factory must return the address of the sender it will create
        IAccountFactory accountFactory = IAccountFactory(factory);

        require(
            accountFactory.getAddress(signer, 0) == op.getSender(),
            "factory must return sender"
        );

        // the rest of the initCode is the data to be passed to the factory
        bytes calldata data = initCode[20:];

        // call the factory
        _call(factory, 0, data);

        // the account must be created
        require(
            _contractExists(op.getSender()),
            "invalid account initialization"
        );
    }

    // generate a full hash of the user op in order to verify the signature
    function _userOpHash(
        UserOperation calldata userOp
    ) public view returns (bytes32) {
        return
            keccak256(abi.encode(userOp.hash(), address(this), block.chainid));
    }

    /**
     * @dev Checks if a contract exists at a given address.
     * @param contractAddress The address to check.
     * @return A boolean indicating whether a contract exists at the given address.
     */
    function _contractExists(
        address contractAddress
    ) internal virtual returns (bool) {
        uint size;
        assembly {
            size := extcodesize(contractAddress)
        }
        return size > 0;
    }

    /**
     * @dev Internal function to call a contract with value and data.
     * If the call fails, it reverts with the reason returned by the callee.
     * @param target The address of the contract to call.
     * @param value The amount of ether to send with the call.
     * @param data The data to send with the call.
     */
    function _call(address target, uint256 value, bytes memory data) internal {
        (bool success, bytes memory result) = target.call{value: value}(data);
        if (!success) {
            assembly {
                revert(add(result, 32), mload(result))
            }
        }
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal view override onlyOwner {
        (newImplementation);
    }
}
