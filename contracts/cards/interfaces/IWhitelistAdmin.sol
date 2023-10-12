// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IWhitelistAdmin {
    function frozen() external view returns (bool);

    function freeze() external;

    function unfreeze() external;

    function updateMaxPersonal(uint256 maxPersonal) external;

    function updateWhitelist(
        address[] calldata authorized
    ) external returns (address[] memory);
}
