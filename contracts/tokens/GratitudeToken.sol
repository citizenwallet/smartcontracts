// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

contract GratitudeToken is
    Initializable,
    ERC20Upgradeable,
    ERC20BurnableUpgradeable,
    OwnableUpgradeable,
    AccessControlUpgradeable,
    PausableUpgradeable,
    UUPSUpgradeable
{
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    event Minted(address indexed to, uint256 amount);

    function initialize(
        address[] memory minters,
        string memory name,
        string memory symbol
    ) public initializer {
        __ERC20_init(name, symbol);
        __ERC20Burnable_init();
        __Ownable_init();
        __AccessControl_init();
        __Pausable_init();

        _setupRole(PAUSER_ROLE, _msgSender());
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        for (uint256 i = 0; i < minters.length; i++) {
            _setupRole(MINTER_ROLE, minters[i]);
        }
    }

    function decimals() public view virtual override returns (uint8) {
        return 0;
    }

    // You always have at least one token of gratitude to give away
    function balanceOf(
        address account
    ) public view virtual override returns (uint256) {
        uint256 balance = super.balanceOf(account);

        return balance > 1 ? balance : 1;
    }

    // Every time you send a token of gratitude, you mint a new one
    function transfer(
        address recipient,
        uint256 amount
    ) public virtual override returns (bool) {
        _mint(msg.sender, 1);
        return super.transfer(recipient, amount);
    }

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) public virtual override returns (bool) {
        _mint(sender, 1);
        (amount);
        return super.transferFrom(sender, recipient, 1);
    }

    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}
}
