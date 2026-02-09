// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";

/**
 * @title SimpleFaucet
 * @dev A simple faucet contract that allows users to redeem tokens at a specified interval.
 */
contract SimpleFaucet is
    Initializable,
    OwnableUpgradeable,
    AccessControlUpgradeable,
    UUPSUpgradeable
{
    bytes32 public constant REDEEM_ADMIN_ROLE = keccak256("REDEEM_ADMIN_ROLE");

    IERC20Upgradeable public token;
    uint256 public amount;
    uint48 public redeemInterval;
    address public redeemAdmin;

    /**
     * @dev A mapping that keeps track of the last time each address redeemed from the faucet.
     * The key of the mapping is the address of the sender, and the value is the timestamp (in seconds) when the redemption occurred.
     */
    mapping(address receiver => uint48 time) public redeemed;

    function initialize(
        address owner,
        IERC20Upgradeable _token,
        uint256 _amount,
        uint48 _redeemInterval,
        address _redeemAdmin
    ) public initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();
        token = _token;
        amount = _amount;
        redeemInterval = _redeemInterval;
        redeemAdmin = _redeemAdmin;
        _setupRole(REDEEM_ADMIN_ROLE, _redeemAdmin);
        transferOwnership(owner);
    }

    /**
     * @dev Allows users to redeem tokens from the SimpleFaucet contract.
     * Users can redeem tokens based on a redeem interval or if they have not redeemed before.
     * If a redeem interval is set, the user must wait until the interval has passed since their last redemption.
     * If no interval is set, the user must not have redeemed before.
     * Upon successful redemption, the specified amount of tokens is transferred to the user's address.
     * The redemption time is recorded for future interval checks.
     */
    function redeem(address recipient) public {
        uint48 currentTime = uint48(block.timestamp);

        if (redeemInterval > 0) {
            // if a redeem interval is set, check if the user can redeem based on interval
            uint48 allowedTime = redeemed[recipient] + redeemInterval;

            require(
                currentTime >= allowedTime,
                "SimpleFaucet: redeem interval not passed"
            );
        } else {
            // if no interval is set, check if the user has redeemed before
            require(redeemed[recipient] == 0, "SimpleFaucet: already redeemed");
        }

        require(
            token.balanceOf(address(this)) >= amount,
            "SimpleFaucet: insufficient balance"
        );

        token.transfer(recipient, amount);

        redeemed[recipient] = currentTime;
    }

    function redeem() public {
        return redeem(msg.sender);
    }

    /**
     * @dev Same as redeem except that you can choose who it should go to.
     */
    function redeemTo(address to) public {
        uint48 currentTime = uint48(block.timestamp);

        if (redeemInterval > 0) {
            // if a redeem interval is set, check if the user can redeem based on interval
            uint48 allowedTime = redeemed[msg.sender] + redeemInterval;

            require(
                currentTime >= allowedTime,
                "SimpleFaucet: redeem interval not passed"
            );
        } else {
            // if no interval is set, check if the user has redeemed before
            require(
                redeemed[msg.sender] == 0,
                "SimpleFaucet: already redeemed"
            );
        }

        require(
            token.balanceOf(address(this)) >= amount,
            "SimpleFaucet: insufficient balance"
        );

        token.transfer(to, amount);

        redeemed[msg.sender] = currentTime;
    }

    /**
     * @dev Allows the redeem code creator to withdraw all the tokens from the contract.
     */
    function withdraw() public onlyRole(REDEEM_ADMIN_ROLE) {
        token.transfer(redeemAdmin, token.balanceOf(address(this)));
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal view override onlyOwner {
        (newImplementation);
    }
}
