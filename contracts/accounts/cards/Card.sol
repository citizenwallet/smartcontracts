// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/* solhint-disable avoid-low-level-calls */
/* solhint-disable no-inline-assembly */
/* solhint-disable reason-string */

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/interfaces/IERC1271.sol";

import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";
import "@account-abstraction/contracts/interfaces/INonceManager.sol";

import "../callback/TokenCallbackHandler.sol";
import "../interfaces/ITokenEntryPoint.sol";
import "../interfaces/IOwnable.sol";
import "../interfaces/IWhitelist.sol";

import "../../standards/ERC20.sol";

/**
 * @title Card
 * @dev This contract represents an account that can execute transactions and store funds in an entry point contract.
 * It implements the ERC1271 standard for signature validation and is upgradeable using the UUPSUpgradeable pattern.
 * The account owner can execute transactions directly or through the entry point contract, and can allow an authorizer contract to execute transactions on its behalf.
 *
 * https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/samples/SimpleAccount.sol
 */
contract Card is
    IERC1271,
    TokenCallbackHandler,
    UUPSUpgradeable,
    Initializable
{
    using ECDSA for bytes32;

    address public owner;

    ERC20 private immutable _token;

    IWhitelist private _whitelist;

    event CardInitialized(address indexed token, address indexed owner);

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyAccountOwner() {
        _checkAccountOwner();
        _;
    }

    function _checkAccountOwner() internal view virtual {
        require(
            owner == msg.sender || address(this) == msg.sender,
            "Ownable: caller is not the owner or the contract"
        );
    }

    function token() public view virtual returns (address) {
        return address(_token);
    }

    function whitelist() public view virtual returns (IWhitelist) {
        return _whitelist;
    }

    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    constructor(address aToken) {
        _token = ERC20(aToken);
        _disableInitializers();
    }

    /**
     * @dev The _entryPoint member is immutable, to reduce gas consumption.  To upgrade EntryPoint,
     * a new implementation of Card must be deployed with the new EntryPoint address, then upgrading
     * the implementation by calling `upgradeTo()`
     */
    function initialize(
        address anOwner,
        address aWhitelist
    ) public virtual initializer {
        migrateState(address(0));
        _initialize(anOwner, aWhitelist);
    }

    function _initialize(address anOwner, address aWhitelist) internal virtual {
        owner = anOwner;
        _whitelist = IWhitelist(aWhitelist);
        emit CardInitialized(address(_token), anOwner);
    }

    function _call(address target, uint256 value, bytes memory data) internal {
        (bool success, bytes memory result) = target.call{value: value}(data);
        if (!success) {
            assembly {
                revert(add(result, 32), mload(result))
            }
        }
    }

    // ************************

    // Card

    function withdraw(address to, uint256 amount) external {
        require(_whitelist.isWhitelisted(to), "Card: to is not whitelisted");

        _token.transfer(to, amount);
    }

    function destroy() external onlyAccountOwner {
        owner = address(0);
        _whitelist.clear();
    }

    // ************************

    // ERC1271 implementation
    /**
     * @notice Verifies that the signer is the owner of the signing contract.
     */
    function isValidSignature(
        bytes32 _hash,
        bytes calldata _signature
    ) external view override returns (bytes4) {
        address signer = recoverSigner(_hash, _signature);

        // Validate signatures
        if (signer == owner) {
            return 0x1626ba7e;
        } else {
            return 0xffffffff;
        }
    }

    function readBytes32(
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
    function recoverSigner(
        bytes32 _hash,
        bytes memory _signature
    ) internal pure returns (address signer) {
        require(
            _signature.length == 65,
            "SignatureValidator#recoverSigner: invalid signature length"
        );

        // Variables are not scoped in Solidity.
        uint8 v = uint8(_signature[64]);
        bytes32 r = readBytes32(_signature, 0);
        bytes32 s = readBytes32(_signature, 32);

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

    // related to upgrades

    mapping(address => bool) private _upgradedTo;

    function migrateState(address oldImplementation) internal onlyInitializing {
        // check that we are allowed to migrate

        require(
            _upgradedTo[_getImplementation()] == false,
            "Card: already migrated"
        );

        _upgradedTo[_getImplementation()] = true;
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal view override onlyAccountOwner {
        (newImplementation);
    }
}
