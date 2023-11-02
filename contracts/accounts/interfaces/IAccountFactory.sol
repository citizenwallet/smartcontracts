// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@account-abstraction/contracts/interfaces/UserOperation.sol";

import "../Account.sol";

interface IAccountFactory {
    function createAccount(
        address owner,
        uint256 salt
    ) external returns (Account ret);

    function getAddress(
        address owner,
        uint256 salt
    ) external view returns (address);
}
