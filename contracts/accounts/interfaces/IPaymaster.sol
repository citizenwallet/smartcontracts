// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@account-abstraction/contracts/interfaces/UserOperation.sol";

interface IPaymaster {
    function getNonce(
        address sender,
        uint192 key
    ) external view returns (uint256 nonce);

    function validatePaymasterUserOp(
        UserOperation calldata userOp
    ) external returns (bool);
}
