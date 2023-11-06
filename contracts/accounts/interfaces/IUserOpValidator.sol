// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@account-abstraction/contracts/interfaces/UserOperation.sol";

interface IUserOpValidator {
    function isValidUserOp(
        UserOperation calldata userOp,
        bytes32 userOpHash
    ) external returns (bool);
}
