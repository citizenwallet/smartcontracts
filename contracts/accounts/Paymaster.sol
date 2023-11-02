// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.12;

/* solhint-disable reason-string */
/* solhint-disable no-inline-assembly */

import "@account-abstraction/contracts/core/BasePaymaster.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import "hardhat/console.sol";

/**
 * A sample paymaster that uses external service to decide whether to pay for the UserOp.
 * The paymaster trusts an external signer to sign the transaction.
 * The calling user must pass the UserOp to that external signer first, which performs
 * whatever off-chain verification before signing the UserOp.
 * Note that this signature is NOT a replacement for the account-specific signature:
 * - the paymaster checks a signature to agree to PAY for GAS.
 * - the account checks a signature to prove identity and account ownership.
 */
contract Paymaster {
    using ECDSA for bytes32;
    using UserOperationLib for UserOperation;

    address public verifyingSigner;

    uint256 private constant VALID_TIMESTAMP_OFFSET = 20;

    uint256 private constant SIGNATURE_OFFSET = 84;

    constructor() {}

    function __Paymaster_init(address _verifyingSigner) internal {
        __Paymaster_init_unchained(_verifyingSigner);
    }

    function __Paymaster_init_unchained(address _verifyingSigner) internal {
        verifyingSigner = _verifyingSigner;
    }

    mapping(address => uint256) public senderNonce;

    function getNonce(address sender) public view returns (uint256 nonce) {
        return senderNonce[sender] | (uint256(0) << 64);
    }

    function pack(
        UserOperation calldata userOp
    ) internal pure returns (bytes memory ret) {
        // lighter signature scheme. must match UserOp.ts#packUserOp
        address sender = userOp.getSender();
        uint256 nonce = userOp.nonce;
        bytes32 hashInitCode = calldataKeccak(userOp.initCode);
        bytes32 hashCallData = calldataKeccak(userOp.callData);
        uint256 callGasLimit = userOp.callGasLimit;
        uint256 verificationGasLimit = userOp.verificationGasLimit;
        uint256 preVerificationGas = userOp.preVerificationGas;
        uint256 maxFeePerGas = userOp.maxFeePerGas;
        uint256 maxPriorityFeePerGas = userOp.maxPriorityFeePerGas;

        return
            abi.encode(
                sender,
                nonce,
                hashInitCode,
                hashCallData,
                callGasLimit,
                verificationGasLimit,
                preVerificationGas,
                maxFeePerGas,
                maxPriorityFeePerGas
            );
    }

    /**
     * return the hash we're going to sign off-chain (and validate on-chain)
     * this method is called by the off-chain service, to sign the request.
     * it is called on-chain from the validatePaymasterUserOp, to validate the signature.
     * note that this signature covers all fields of the UserOperation, except the "paymasterAndData",
     * which will carry the signature itself.
     */
    function getHash(
        UserOperation calldata userOp,
        uint48 validUntil,
        uint48 validAfter
    ) public view returns (bytes32) {
        // can't use userOp.hash(), since it contains also the paymasterAndData itself.

        return
            keccak256(
                abi.encode(
                    pack(userOp),
                    block.chainid,
                    address(this),
                    senderNonce[userOp.getSender()],
                    validUntil,
                    validAfter
                )
            );
    }

    /**
     * verify our external signer signed this request.
     * the "paymasterAndData" is expected to be the paymaster and a signature over the entire request params
     * paymasterAndData[:20] : address(this)
     * paymasterAndData[20:84] : abi.encode(validUntil, validAfter)
     * paymasterAndData[84:] : signature
     */
    function _validatePaymasterUserOp(
        UserOperation calldata userOp
    ) internal returns (bool) {
        (
            uint48 validUntil,
            uint48 validAfter,
            bytes calldata signature
        ) = _parsePaymasterAndData(userOp.paymasterAndData);
        // ECDSA library supports both 64 and 65-byte long signatures.
        // we only "require" it here so that the revert reason on invalid signature will be of "VerifyingPaymaster", and not "ECDSA"
        require(
            signature.length == 64 || signature.length == 65,
            "VerifyingPaymaster: invalid signature length in paymasterAndData"
        );

        uint48 currentTime = uint48(block.timestamp);
        require(
            currentTime > validAfter && currentTime <= validUntil,
            "signature has expired"
        );

        bytes32 hash = getHash(userOp, validUntil, validAfter)
            .toEthSignedMessageHash();

        senderNonce[userOp.getSender()]++;

        // don't revert on signature failure
        if (verifyingSigner != hash.recover(signature)) {
            return false;
        }

        // no need for other on-chain validation: entire UserOp should have been checked
        // by the external service prior to signing it.
        return true;
    }

    function _parsePaymasterAndData(
        bytes calldata paymasterAndData
    )
        internal
        pure
        returns (uint48 validUntil, uint48 validAfter, bytes calldata signature)
    {
        (validUntil, validAfter) = abi.decode(
            paymasterAndData[VALID_TIMESTAMP_OFFSET:SIGNATURE_OFFSET],
            (uint48, uint48)
        );
        signature = paymasterAndData[SIGNATURE_OFFSET:];
    }
}