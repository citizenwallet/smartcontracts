// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

library TimestampLib {
    function getCurrentTimestamp() internal view returns (uint256) {
        return block.timestamp;
    }

    function getTimestampAfterXYears(
        uint256 amount
    ) internal view returns (uint256) {
        // Number of seconds in a year (assuming an average year of 365.25 days)
        uint256 secondsInAYear = 365 days + 6 hours;

        // Calculate the number of seconds in three years
        uint256 secondsInYears = secondsInAYear * amount;

        return block.timestamp + secondsInYears;
    }

    function getTimestampAfterXHours(
        uint256 amount
    ) internal view returns (uint256) {
        // Number of seconds in a year (assuming an average year of 365.25 days)
        uint256 secondsInAnHour = 60 minutes;

        // Calculate the number of seconds in three years
        uint256 secondsInHours = secondsInAnHour * amount;

        return block.timestamp + secondsInHours;
    }
}
