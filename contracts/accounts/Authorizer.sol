// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/interfaces/IERC1271.sol";

import "@account-abstraction/contracts/interfaces/UserOperation.sol";
import "@account-abstraction/contracts/core/NonceManager.sol";

import "./Paymaster.sol";

/**
 * A sample factory contract for Account
 * A UserOperations "initCode" holds the address of the factory, and a method call (to createAccount, in this sample factory).
 * The factory's createAccount returns the target account address even if it is already installed.
 * This way, the entryPoint.getSenderAddress() can be called either before or after the account is created.
 */
contract Authorizer is
    Paymaster,
    ReentrancyGuard,
    Initializable,
    OwnableUpgradeable,
    UUPSUpgradeable
{
    using UserOperationLib for UserOperation;

    constructor() {
        _disableInitializers();
    }

    function initialize(address anOwner) public virtual initializer {
        __Ownable_init();
        __Paymaster_init(anOwner);

        _initialize(anOwner);
    }

    function _initialize(address anOwner) internal virtual {
        transferOwnership(anOwner);
    }

    function handleOps(
        UserOperation[] calldata ops,
        address payable /*beneficiary*/
    ) public nonReentrant {
        require(ops.length > 0, "needs at least one user operation");

        uint len = ops.length;
        for (uint i = 0; i < len; ) {
            // handle each op
            UserOperation calldata op = ops[i];

            // verify paymaster signature
            _validatePaymaster(op);

            // verify nonce
            _validateNonce(op);

            // verify account
            _validateAccount(op);

            // execute the op
            // TODO: execute the op

            unchecked {
                ++i;
            }
        }
    }

    function _validatePaymaster(UserOperation calldata op) internal virtual {
        bytes calldata paymasterAndData = op.paymasterAndData;

        require(paymasterAndData.length >= 20, "invalid paymasterAndData");

        address paymaster = address(bytes20(paymasterAndData[:20]));

        require(paymaster == owner(), "invalid paymaster");

        // verify paymasterAndData signature
        require(
            _validatePaymasterUserOp(op),
            "invalid paymasterAndData signature"
        );
    }

    function _validateNonce(UserOperation calldata op) internal virtual {
        uint256 nonce = getNonce(op.sender);

        require(op.nonce == nonce, "invalid nonce");
    }

    function _validateAccount(UserOperation calldata op) internal virtual {
        // verify the user op signature
        IERC1271 account = IERC1271(op.sender);

        account.isValidSignature(op.hash(), op.signature);

        // call the initCode
        if (op.nonce > 0) {
            return;
        }

        _createAccount(op);
    }

    function _createAccount(UserOperation calldata op) internal virtual {
        bytes calldata initCode = op.initCode;

        require(initCode.length >= 20, "invalid initCode");

        address factory = address(bytes20(initCode[0:20]));
        bytes calldata data = initCode[20:];

        _call(factory, 0, data);
    }

    function _call(address target, uint256 value, bytes memory data) internal {
        (bool success, bytes memory result) = target.call{value: value}(data);
        if (!success) {
            assembly {
                revert(add(result, 32), mload(result))
            }
        }
    }

    // A memory copy of UserOp static fields only.
    // Excluding: callData, initCode and signature. Replacing paymasterAndData with paymaster.
    struct MemoryUserOp {
        address sender;
        uint256 nonce;
        uint256 callGasLimit;
        uint256 verificationGasLimit;
        uint256 preVerificationGas;
        address paymaster;
        uint256 maxFeePerGas;
        uint256 maxPriorityFeePerGas;
    }

    /**
     * copy general fields from userOp into the memory opInfo structure.
     */
    function _copyUserOpToMemory(
        UserOperation calldata userOp,
        MemoryUserOp memory mUserOp
    ) internal pure {
        mUserOp.sender = userOp.sender;
        mUserOp.nonce = userOp.nonce;
        mUserOp.callGasLimit = userOp.callGasLimit;
        mUserOp.verificationGasLimit = userOp.verificationGasLimit;
        mUserOp.preVerificationGas = userOp.preVerificationGas;
        mUserOp.maxFeePerGas = userOp.maxFeePerGas;
        mUserOp.maxPriorityFeePerGas = userOp.maxPriorityFeePerGas;
        bytes calldata paymasterAndData = userOp.paymasterAndData;
        if (paymasterAndData.length > 0) {
            require(
                paymasterAndData.length >= 20,
                "AA93 invalid paymasterAndData"
            );
            mUserOp.paymaster = address(bytes20(paymasterAndData[:20]));
        } else {
            mUserOp.paymaster = address(0);
        }
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal view override onlyOwner {
        (newImplementation);
    }
}
