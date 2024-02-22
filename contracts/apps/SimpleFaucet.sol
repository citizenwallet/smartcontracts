// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";

import "@openzeppelin/contracts/interfaces/IERC1271.sol";

contract SimpleFaucet is Initializable, OwnableUpgradeable, UUPSUpgradeable {
    IERC20Upgradeable public token;
    uint256 public amount;
    uint48 public redeemInterval;

    mapping(address sender => uint48 time) public redeemed;

    function initialize(IERC20Upgradeable _token, uint256 _amount, uint48 _redeemInterval) public initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();
        token = _token;
        amount = _amount;
        redeemInterval = _redeemInterval;
    }

    function redeem() public {
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

        token.transfer(msg.sender, amount);

        redeemed[msg.sender] = currentTime;
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal view override onlyOwner {
        (newImplementation);
    }
}
