// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface ITimestamps {
    function getCurrentTimestamp() external view returns (uint256);

    function getTimestampAfterXYears(
        uint256 amount
    ) external view returns (uint256);

    function getTimestampAfterXHours(
        uint256 amount
    ) external view returns (uint256);
}
