// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@account-abstraction/contracts/core/BaseAccount.sol";
import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";

import "./interfaces/ICard.sol";
import "./interfaces/IWhitelistReader.sol";
import "./interfaces/ITimestamps.sol";
import "../accounts/callback/TokenCallbackHandler.sol";

// https://github.com/eth-infinitism/account-abstraction/blob/develop/contracts/samples/SimpleAccount.sol
// Card,
contract Card is
    BaseAccount,
    ICard,
    TokenCallbackHandler,
    Initializable,
    OwnableUpgradeable,
    UUPSUpgradeable
{
    constructor(
        IEntryPoint anEntryPoint,
        IWhitelistReader aWhitelist,
        ITimestamps aTimestamps
    ) {
        _entryPoint = anEntryPoint;
        _whitelist = aWhitelist;
        _timestamps = aTimestamps;

        _expiration = ITimestamps(timestamps()).getTimestampAfterXYears(3);
        _disableInitializers();
    }

    // ERC4337 implementation
    using ECDSA for bytes32;

    IEntryPoint private immutable _entryPoint;

    event SimpleAccountInitialized(
        IEntryPoint indexed entryPoint,
        address indexed owner
    );

    /// @inheritdoc BaseAccount
    function entryPoint() public view virtual override returns (IEntryPoint) {
        return _entryPoint;
    }

    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    /**
     * execute a transaction (called directly from owner, or by entryPoint)
     */
    function execute(
        address dest,
        uint256 value,
        bytes calldata func
    ) external {
        _requireFromEntryPointOrOwnerOrWhitelist();
        _call(dest, value, func);
    }

    /**
     * execute a sequence of transactions
     * @dev to reduce gas consumption for trivial case (no value), use a zero-length array to mean zero value
     */
    function executeBatch(
        address[] calldata dest,
        uint256[] calldata value,
        bytes[] calldata func
    ) external {
        _requireFromEntryPointOrOwnerOrWhitelist();
        require(
            dest.length == func.length &&
                (value.length == 0 || value.length == func.length),
            "wrong array lengths"
        );
        if (value.length == 0) {
            for (uint256 i = 0; i < dest.length; i++) {
                _call(dest[i], 0, func[i]);
            }
        } else {
            for (uint256 i = 0; i < dest.length; i++) {
                _call(dest[i], value[i], func[i]);
            }
        }
    }

    /**
     * @dev The _entryPoint member is immutable, to reduce gas consumption.  To upgrade EntryPoint,
     * a new implementation of SimpleAccount must be deployed with the new EntryPoint address, then upgrading
     * the implementation by calling `upgradeTo()`
     */
    function initialize(address anOwner) public virtual initializer {
        _initialize(anOwner);
    }

    function _initialize(address anOwner) internal virtual {
        __Ownable_init();
        __UUPSUpgradeable_init();

        transferOwnership(anOwner);

        emit SimpleAccountInitialized(_entryPoint, owner());
    }

    // Require the function call went through EntryPoint or owner
    function _requireFromEntryPointOrOwnerOrWhitelist() internal view {
        require(
            msg.sender == address(entryPoint()) ||
                msg.sender == owner() ||
                _isWhitelisted(),
            "card: not Whitelisted or Owner or EntryPoint"
        );

        require(hasExpired() == false, "card: expired");
    }

    /// implement template method of BaseAccount
    function _validateSignature(
        UserOperation calldata userOp,
        bytes32 userOpHash
    ) internal virtual override returns (uint256 validationData) {
        bytes32 hash = userOpHash.toEthSignedMessageHash();
        if (owner() != hash.recover(userOp.signature))
            return SIG_VALIDATION_FAILED;
        return 0;
    }

    function _call(address target, uint256 value, bytes memory data) internal {
        (bool success, bytes memory result) = target.call{value: value}(data);
        if (!success) {
            assembly {
                revert(add(result, 32), mload(result))
            }
        }
    }

    /**
     * check current account deposit in the entryPoint
     */
    function getDeposit() public view returns (uint256) {
        return entryPoint().balanceOf(address(this));
    }

    /**
     * deposit more funds for this account in the entryPoint
     */
    function addDeposit() public payable {
        entryPoint().depositTo{value: msg.value}(address(this));
    }

    /**
     * withdraw value from the account's deposit
     * @param withdrawAddress target to send to
     * @param amount to withdraw
     */
    function withdrawDepositTo(
        address payable withdrawAddress,
        uint256 amount
    ) public onlyOwner {
        entryPoint().withdrawTo(withdrawAddress, amount);
    }

    // ************************

    // Timestamps implementation
    ITimestamps private immutable _timestamps;

    function timestamps() public view virtual returns (ITimestamps) {
        return _timestamps;
    }

    // ************************

    // Card implementation
    uint256 private immutable _expiration;

    IWhitelistReader private immutable _whitelist;

    function whitelist() public view virtual returns (IWhitelistReader) {
        return _whitelist;
    }

    function hasExpired() public view override returns (bool) {
        uint256 currentTimestamp = ITimestamps(timestamps())
            .getCurrentTimestamp();

        return currentTimestamp > _expiration;
    }

    function expiresAt() external view returns (uint256) {
        return _expiration;
    }

    function _isWhitelisted() internal view virtual returns (bool) {
        return IWhitelistReader(whitelist()).isAllowed(address(this));
    }

    // ************************

    // ERC1271 implementation
    bytes4 internal constant MAGICVALUE = 0x1626ba7e;

    /**
     * @notice Verifies that the signer is the owner of the signing contract.
     */
    function isValidSignature(
        bytes32 _hash,
        bytes calldata _signature
    ) external view override returns (bytes4) {
        address signer = recoverSigner(_hash, _signature);

        // Validate signatures
        if (signer == owner()) {
            return MAGICVALUE;
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

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}
}
