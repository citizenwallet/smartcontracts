// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";

/**
 * @title RedeemCodeFaucet
 * @dev A contract that allows users to redeem codes
 * It is upgradeable and has access control functionality.
 */
contract RedeemCodeFaucet is Initializable, OwnableUpgradeable, AccessControlUpgradeable, UUPSUpgradeable {
    bytes32 public constant REDEEM_CODE_CREATOR_ROLE = keccak256("REDEEM_CODE_CREATOR_ROLE");

    IERC20Upgradeable public token;
    uint48 public callInterval;
    address public codeCreator;

    event CodeRedeemed(address indexed receiver, uint256 amount);

    /**
     * @dev Mapping to store the redeemableAmount for each redeemable code.
     */
    mapping(bytes32 codeHash => uint256 redeemableAmount) private redeemableAmounts;

    /**
     * @dev Mapping to store the validity period of each redeemable code.
     */
    mapping(bytes32 codeHash => uint48 validUntil) private validity;

    /**
     * @dev Mapping to store the time when each redeemable code was redeemed.
     */
    mapping(bytes32 codeHash => uint48 time) public redeemed;

    /**
     * @dev Mapping to store the time when each address called to redeem from the faucet.
     */
    mapping(address sender => uint48 time) private lastCall;

    /**
     * @dev Mapping to store the cooldown time when an address is allowed to call redeem from the faucet again.
     */
    mapping(address sender => uint48 time) private redeemCooldown;

    function initialize(address owner, IERC20Upgradeable _token, uint48 _callInterval, address _codeCreator) public initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();
        token = _token;
        callInterval = _callInterval;
        codeCreator = _codeCreator;
        _setupRole(REDEEM_CODE_CREATOR_ROLE, _codeCreator);
        transferOwnership(owner);
    }

    /**
     * @dev Calculates the hash value for a given address and code.
     * @param _codeCreator The address of the code creator.
     * @param code The code to be hashed.
     * @return The calculated hash value.
     */
    function getHash(
        address _codeCreator,
        uint256 code
    ) public view returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    _codeCreator,
                    code,
                    block.chainid,
                    address(this)
                )
            );
    }

    /**
     * @dev Adds a redeem code to the faucet.
     * @param codeHash The redeem code hash to add.
     * @param redeemableAmount The redeemableAmount associated with the redeem code.
     * @param validUntil The expiration timestamp for the redeem code.
     * Requirements:
     * - `redeemableAmount` must be greater than zero.
     * - `validUntil` must be in the future.
     * - The redeem code must not have been already redeemed.
     */
    function addRedeemCode(bytes32 codeHash, uint256 redeemableAmount, uint48 validUntil) public onlyRole(REDEEM_CODE_CREATOR_ROLE) {
        require(
            redeemableAmount > 0,
            "RedeemCodeFaucet: redeemableAmount must be greater than zero"
        );

        uint48 currentTime = uint48(block.timestamp);

        require(
            validUntil > currentTime,
            "RedeemCodeFaucet: validUntil must be in the future"
        );

        require(
            redeemed[codeHash] == 0,
            "RedeemCodeFaucet: already redeemed"
        );

        redeemableAmounts[codeHash] = redeemableAmount;
        validity[codeHash] = validUntil;
    }


    /**
     * @dev Allows an address to redeem a code and receive tokens from the faucet.
     * @param code The code to redeem.
     * Requirements:
     * - The redeem interval must have passed since the last redemption by the caller.
     * - The code must not have been previously redeemed.
     * - The code must not have expired.
     * - The faucet must have sufficient balance of tokens.
     */
    function redeem(uint256 code) public {
        uint48 currentTime = uint48(block.timestamp);

        redeemCooldown[msg.sender] = callInterval;
        uint48 previousCall = lastCall[msg.sender];
        lastCall[msg.sender] = currentTime;
    
        // stop a single address from calling this function too often
        if (previousCall + redeemCooldown[msg.sender] >= currentTime) {
            redeemCooldown[msg.sender] += callInterval; // add a penalty to the cooldown
            return;
        }

        bytes32 codeHash = getHash(
                    codeCreator,
                    code
                );

        require(
            redeemed[codeHash] == 0,
            "RedeemCodeFaucet: code already redeemed"
        );

        require(
            currentTime < validity[codeHash],
            "RedeemCodeFaucet: code expired"
        );

        require(
            token.balanceOf(address(this)) >= redeemableAmounts[codeHash],
            "RedeemCodeFaucet: insufficient balance"
        );

        uint256 redeemableAmount = redeemableAmounts[codeHash];

        token.transfer(msg.sender, redeemableAmount);

        redeemed[codeHash] = currentTime;
        
        emit CodeRedeemed(msg.sender, redeemableAmount);
    }

    /**
     * @dev Allows the redeem code creator to withdraw all the tokens from the contract.
     */
    function withdraw() public onlyRole(REDEEM_CODE_CREATOR_ROLE) {
        token.transfer(codeCreator, token.balanceOf(address(this)));
    }

    function grantRole(bytes32 role, address account) public virtual override onlyRole(getRoleAdmin(role)) {
        require(
            role != REDEEM_CODE_CREATOR_ROLE,
            "AccessControl: cannot grant codeCreator role"
        );

        _grantRole(role, account);
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal view override onlyOwner {
        (newImplementation);
    }
}
