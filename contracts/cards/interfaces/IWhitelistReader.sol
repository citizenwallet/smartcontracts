// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IWhitelistReader {
    function isAllowed(address from) external view returns (bool);

    function maxPersonal() external view returns (uint256);

    function updatePersonalList(
        address owner,
        address[] calldata authorized
    ) external returns (address[] memory);
}
