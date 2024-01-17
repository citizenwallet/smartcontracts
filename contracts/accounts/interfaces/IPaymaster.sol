// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@account-abstraction/contracts/interfaces/UserOperation.sol";

interface IPaymaster {
    function validatePaymasterUserOp(
        UserOperation calldata userOp
    ) external view returns (bool);
}
