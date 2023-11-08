// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

contract RegensUniteToken is
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

    event Minted(address indexed to, uint256 amount, string description);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address[] memory minters) public initializer {
        __ERC20_init("Regens Unite Token", "RGNU");
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
        return 6;
    }

    function mint(
        address to,
        uint256 amount,
        string memory description
    ) public {
        require(
            hasRole(MINTER_ROLE, msg.sender),
            "Must have minter role to mint"
        );
        _mint(to, amount);
        emit Minted(to, amount, description);
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
