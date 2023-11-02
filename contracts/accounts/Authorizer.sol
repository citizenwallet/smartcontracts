// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/interfaces/IERC1271.sol";

import "@account-abstraction/contracts/interfaces/UserOperation.sol";

import "./Paymaster.sol";
import "./interfaces/IAccountFactory.sol";

/**
 * A sample factory contract for Account
 * A UserOperations "initCode" holds the address of the factory, and a method call (to createAccount, in this sample factory).
 * The factory's createAccount returns the target account address even if it is already installed.
 * This way, the entryPoint.getSenderAddress() can be called either before or after the account is created.
 */
contract Authorizer is
    Paymaster,
    Initializable,
    ReentrancyGuardUpgradeable,
    OwnableUpgradeable,
    UUPSUpgradeable
{
    using ECDSA for bytes32;
    using UserOperationLib for UserOperation;

    event PaymasterVerifierUpdated(address indexed newVerifier);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    // we make the owner of also the verifier by default
    function initialize(address anOwner) public virtual initializer {
        __Ownable_init();
        __Paymaster_init(anOwner);
        __ReentrancyGuard_init();

        _initialize(anOwner);
    }

    function _initialize(address anOwner) internal virtual {
        transferOwnership(anOwner);
    }

    function updatePaymasterVerifier(
        address newVerifier
    ) public virtual onlyOwner {
        verifyingSigner = newVerifier;
        emit PaymasterVerifierUpdated(newVerifier);
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

            address sender = op.getSender();

            // verify nonce
            _validateNonce(op, sender);

            // verify account
            _validateAccount(op, sender);

            // verify paymaster signature
            _validatePaymaster(op);

            // execute the op
            _call(sender, 0, op.callData);

            unchecked {
                ++i;
            }
        }
    }

    function _validatePaymaster(UserOperation calldata op) internal virtual {
        bytes calldata paymasterAndData = op.paymasterAndData;

        // paymasterAndData must be at least 20 bytes long, and the first 20 bytes must be the paymaster address
        require(paymasterAndData.length >= 20, "invalid paymasterAndData");

        address paymaster = address(bytes20(paymasterAndData[:20]));

        // the paymaster should be this contract
        require(paymaster == address(this), "invalid paymaster");

        // verify paymasterAndData signature
        require(_validatePaymasterUserOp(op), "invalid paymaster signature");
    }

    function _validateNonce(
        UserOperation calldata op,
        address sender
    ) internal virtual {
        uint256 nonce = getNonce(sender);

        // the nonce in the user op must match the nonce in the account
        require(op.nonce == nonce, "invalid nonce");
    }

    bytes4 internal constant MAGICVALUE = 0x1626ba7e;

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

        IAccountFactory accountFactory = IAccountFactory(factory);

        address signer = ophash.recover(op.signature);

        // the factory must return the address of the sender it will create
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

    function _contractExists(
        address contractAddress
    ) internal virtual returns (bool) {
        uint size;
        assembly {
            size := extcodesize(contractAddress)
        }
        return size > 0;
    }

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
