// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";

import "@openzeppelin/contracts/interfaces/IERC1271.sol";

/**
 * @title RedeemCodeFaucet
 * @dev A simple faucet contract that allows users to redeem tokens at a specified interval.
 */
contract RedeemCodeFaucet is Initializable, OwnableUpgradeable, AccessControlUpgradeable, UUPSUpgradeable {
    bytes32 public constant ISSUER_ROLE = keccak256("ISSUER_ROLE");

    IERC20Upgradeable public token;
    uint48 public redeemInterval;

    /**
     * @dev A mapping that keeps track of the last time each address redeemed from the faucet.
     * The key of the mapping is the address of the sender, and the value is the timestamp (in seconds) when the redemption occurred.
     */
    mapping(bytes32 codeHash => uint256 amount) private amounts;
    mapping(bytes32 codeHash => uint48 validUntil) private validity;
    mapping(bytes32 codeHash => uint48 time) public redeemed;

    mapping(address sender => uint48 time) public lastRedeem;

    function initialize(IERC20Upgradeable _token, uint48 _redeemInterval, address issuer) public initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();
        token = _token;
        redeemInterval = _redeemInterval;
        _setupRole(ISSUER_ROLE, issuer);
    }

    function _getHash(
        address from,
        uint256 code
    ) internal view returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    from,
                    code,
                    block.chainid,
                    address(this)
                )
            );
    }

    function addRedeemCode(uint256 code, uint256 amount, uint48 validUntil) public onlyRole(ISSUER_ROLE) {
        require(
            code != 0,
            "RedeemCodeFaucet: code cannot be zero"
        );

        require(
            amount > 0,
            "RedeemCodeFaucet: amount must be greater than zero"
        );


        uint48 currentTime = uint48(block.timestamp);

        require(
            validUntil > currentTime,
            "RedeemCodeFaucet: validUntil must be in the future"
        );

        bytes32 codeHash = _getHash(
                    msg.sender,
                    code
                );

        require(
            redeemed[codeHash] == 0,
            "RedeemCodeFaucet: already redeemed"
        );

        amounts[codeHash] = amount;
        validity[codeHash] = validUntil;
    }


    function redeem(address issuer, uint256 code) public {
        uint48 currentTime = uint48(block.timestamp);

        // stop a single address from redeeming too often
        require(
            lastRedeem[msg.sender] + redeemInterval < currentTime,
            "RedeemCodeFaucet: redeem interval not passed"
        );

        bytes32 codeHash = _getHash(
                    issuer,
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
            token.balanceOf(address(this)) >= amounts[codeHash],
            "RedeemCodeFaucet: insufficient balance"
        );

        uint256 amount = amounts[codeHash];

        token.transfer(msg.sender, amount);

        redeemed[codeHash] = currentTime;
        lastRedeem[msg.sender] = currentTime;
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal view override onlyOwner {
        (newImplementation);
    }
}
