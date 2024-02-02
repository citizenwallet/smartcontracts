// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@account-abstraction/contracts/interfaces/UserOperation.sol";

interface IWhitelist {
    function clear() external view;

    function isWhitelisted(address account) external view returns (bool);

    function add(address account) external;

    function remove(address account) external;
}
