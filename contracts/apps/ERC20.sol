// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract ERC20 {
    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(
        address indexed _owner,
        address indexed _spender,
        uint256 _value
    );
}
