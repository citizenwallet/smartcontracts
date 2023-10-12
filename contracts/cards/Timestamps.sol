// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

import "./interfaces/ITimestamps.sol";

contract Timestamps is
    ITimestamps,
    Initializable,
    OwnableUpgradeable,
    UUPSUpgradeable
{
    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();
    }

    function getCurrentTimestamp() external view override returns (uint256) {
        return block.timestamp;
    }

    function getTimestampAfterXYears(
        uint256 amount
    ) external view override returns (uint256) {
        // Number of seconds in a year (assuming an average year of 365.25 days)
        uint256 secondsInAYear = 365 days + 6 hours;

        // Calculate the number of seconds in three years
        uint256 secondsInYears = secondsInAYear * amount;

        return block.timestamp + secondsInYears;
    }

    function getTimestampAfterXHours(
        uint256 amount
    ) external view override returns (uint256) {
        // Number of seconds in a year (assuming an average year of 365.25 days)
        uint256 secondsInAnHour = 60 minutes;

        // Calculate the number of seconds in three years
        uint256 secondsInHours = secondsInAnHour * amount;

        return block.timestamp + secondsInHours;
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}
}
