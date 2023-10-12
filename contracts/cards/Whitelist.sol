// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

import "./interfaces/IWhitelistReader.sol";

contract Whitelist is
    IWhitelistReader,
    Initializable,
    OwnableUpgradeable,
    UUPSUpgradeable
{
    // maximum number of personal whitelist addresses allowed
    uint256 public constant MAX_PERSONAL = 20;

    // global list of addresses that are allowed to use the card
    address[] private _authorized;

    // personal list of addresses that are allowed to use the card
    mapping(address owner => address[] personal) private _personal;

    // freeze all transfers
    bool private _frozen = false;

    function frozen() public view returns (bool) {
        return _frozen;
    }

    function freeze() public onlyOwner {
        _frozen = true;
    }

    function unfreeze() public onlyOwner {
        _frozen = false;
    }

    // ********************

    function authorized() public view virtual returns (address[] memory) {
        return _authorized;
    }

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();
    }

    function isAllowed(address from) public view override returns (bool) {
        require(frozen() == false, "transfers are frozen");

        address to = msg.sender;

        // check if the address is in the global list
        if (_containsAddress(_authorized, to)) {
            return true;
        }

        // retrieve the relevant personal list
        address[] memory personalList = _personal[from];
        if (personalList.length == 0) {
            return false;
        }

        // check if the address is in the personal list
        return _containsAddress(personalList, to);
    }

    function maxPersonal() public pure override returns (uint256) {
        return MAX_PERSONAL;
    }

    function updatePersonalList(
        address owner,
        address[] calldata newAuthorized
    ) public override returns (address[] memory) {
        require(frozen() == false, "transfers are frozen");

        // limit who can use this function
        require(
            msg.sender == owner,
            "only the owner of the list can update their own list"
        );

        // limit the number of addresses that can be added
        require(newAuthorized.length <= MAX_PERSONAL, "too many addresses");

        // update the personal list
        _personal[owner] = newAuthorized;

        // return the new personal list
        return _personal[owner];
    }

    function _containsAddress(
        address[] memory list,
        address value
    ) internal pure returns (bool) {
        for (uint256 i = 0; i < list.length; i++) {
            if (list[i] == value) {
                return true; // Value found in the list
            }
        }
        return false; // Value not found in the list
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}
}
