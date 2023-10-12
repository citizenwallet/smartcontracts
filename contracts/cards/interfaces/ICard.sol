// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@account-abstraction/contracts/interfaces/IAccount.sol";
import "@openzeppelin/contracts/interfaces/IERC1271.sol";

import "@openzeppelin/contracts/token/ERC777/IERC777Recipient.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155Receiver.sol";

interface ICard is
    IAccount,
    IERC1271,
    IERC777Recipient,
    IERC721Receiver,
    IERC1155Receiver
{
    function hasExpired() external view returns (bool);

    function expiresAt() external view returns (uint256);
}
