// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";

import "@openzeppelin/contracts/interfaces/IERC1271.sol";

contract TokenIOU is Initializable, OwnableUpgradeable, UUPSUpgradeable {
    IERC20Upgradeable public token;

    /// only a single redeem per hash
    ///
    /// allows us to check which username was already reserved by a profile
    mapping(bytes32 redeemHash => uint48 time) public redeemed;

    // Redeem event
    event Redeem(address indexed from, address indexed to, uint256 amount);

    function initialize(IERC20Upgradeable _token) public initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();
        token = _token;
    }

    /**
        * @notice Get the hash of the redeem request
        * @param from       The address of the redeemer
        * @param amount     The amount of tokens to redeem
        * @param validUntil The timestamp until which the redeem is valid
        * @param validAfter  The timestamp from which the redeem is valid
        * @param nonce      A unique number to prevent replay attacks
        * @return The hash of the redeem request

        Can be used to easily generate a valid hash for the redeem function
    */
    function getHash(
        address from,
        uint256 amount,
        uint48 validUntil,
        uint48 validAfter,
        uint256 nonce
    ) public view returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    from,
                    amount,
                    validUntil,
                    validAfter,
                    nonce,
                    block.chainid,
                    address(this)
                )
            );
    }

    /**
     * @notice Redeem tokens
     * @param from       The address of the redeemer
     * @param amount     The amount of tokens to redeem
     * @param validUntil The timestamp until which the redeem is valid
     * @param validAfter  The timestamp from which the redeem is valid
     * @param nonce      A unique number to prevent replay attacks
     * @param signature  The signature of the redeem request
     */
    function redeem(
        address from,
        uint256 amount,
        uint48 validUntil,
        uint48 validAfter,
        uint256 nonce,
        bytes calldata signature
    ) public {
        uint48 currentTime = uint48(block.timestamp);
        require(currentTime >= validAfter, "TokenIOU: expired or not due");
        require(currentTime < validUntil, "TokenIOU: expired or not due");

        bytes32 redeemHash = keccak256(
            abi.encodePacked(
                "\x19Ethereum Signed Message:\n32",
                getHash(from, amount, validUntil, validAfter, nonce)
            )
        );

        require(redeemed[redeemHash] == 0, "TokenIOU: Already redeemed");

        // check the sender's signature
        address to = msg.sender;

        // check if the msg.sender is a smart contract
        if (_contractExists(to)) {
            // check ownership of the contract
            IERC1271 account = IERC1271(to);

            require(
                account.isValidSignature(redeemHash, signature) == 0x1626ba7e,
                "TokenIOU: Invalid signature"
            );
        } else {
            // check the signature of the sender
            require(
                _recoverSigner(redeemHash, signature) == to,
                "TokenIOU: Invalid signature"
            );
        }

        require(
            token.transferFrom(from, msg.sender, amount),
            "TokenIOU: Transfer failed"
        );

        redeemed[redeemHash] = currentTime;

        emit Redeem(from, to, amount);
    }

    // ************************

    // Recovering signatures

    function _readBytes32(
        bytes memory data,
        uint256 index
    ) internal pure returns (bytes32 result) {
        require(data.length >= index + 32, "readBytes32: invalid data length");
        assembly {
            result := mload(add(data, add(32, index)))
        }
    }

    /**
     * @notice Recover the signer of hash, assuming it's an EOA account
     * @dev Only for EthSign signatures
     * @param _hash       Hash of message that was signed
     * @param _signature  Signature encoded as (bytes32 r, bytes32 s, uint8 v)
     */
    function _recoverSigner(
        bytes32 _hash,
        bytes memory _signature
    ) internal pure returns (address signer) {
        require(
            _signature.length == 65,
            "SignatureValidator#recoverSigner: invalid signature length"
        );

        // Variables are not scoped in Solidity.
        uint8 v = uint8(_signature[64]);
        bytes32 r = _readBytes32(_signature, 0);
        bytes32 s = _readBytes32(_signature, 32);

        // EIP-2 still allows signature malleability for ecrecover(). Remove this possibility and make the signature
        // unique. Appendix F in the Ethereum Yellow paper (https://ethereum.github.io/yellowpaper/paper.pdf), defines
        // the valid range for s in (281): 0 < s < secp256k1n ÷ 2 + 1, and for v in (282): v ∈ {27, 28}. Most
        // signatures from current libraries generate a unique signature with an s-value in the lower half order.
        //
        // If your library generates malleable signatures, such as s-values in the upper range, calculate a new s-value
        // with 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141 - s1 and flip v from 27 to 28 or
        // vice versa. If your library also generates signatures with 0/1 for v instead 27/28, add 27 to v to accept
        // these malleable signatures as well.
        //
        // Source OpenZeppelin
        // https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/cryptography/ECDSA.sol

        if (
            uint256(s) >
            0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0
        ) {
            revert(
                "SignatureValidator#recoverSigner: invalid signature 's' value"
            );
        }

        if (v != 27 && v != 28) {
            revert(
                "SignatureValidator#recoverSigner: invalid signature 'v' value"
            );
        }

        // Recover ECDSA signer
        signer = ecrecover(_hash, v, r, s);

        // Prevent signer from being 0x0
        require(
            signer != address(0x0),
            "SignatureValidator#recoverSigner: INVALID_SIGNER"
        );

        return signer;
    }

    // ************************

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

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}
}
