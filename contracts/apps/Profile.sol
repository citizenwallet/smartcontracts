// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";

import "../abstract/Linkable.sol";

contract Profile is Initializable, Linkable {
    IEntryPoint private immutable _entryPoint;

    address public owner;

    struct ProfileData {
        string name;
        string description;
        string icon;
        string meta;
    }

    ProfileData public profile;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor(IEntryPoint anEntryPoint) {
        _entryPoint = anEntryPoint;
        profile = ProfileData("", "", "", "");
        _disableInitializers();
    }

    modifier onlyOwner() override {
        _onlyOwner();
        _;
    }

    modifier onlyOwnerOrEntryPoint() override {
        _requireFromEntryPointOrOwner();
        _;
    }

    function _onlyOwner() internal view {
        //directly from EOA owner, or through the account itself (which gets redirected through execute())
        require(
            msg.sender == owner || msg.sender == address(this),
            "only owner"
        );
    }

    // Require the function call went through EntryPoint or owner
    function _requireFromEntryPointOrOwner() internal view {
        require(
            msg.sender == address(_entryPoint) || msg.sender == owner,
            "account: not Owner or EntryPoint"
        );
    }

    /**
     * @dev The _entryPoint member is immutable, to reduce gas consumption.  To upgrade EntryPoint,
     * a new implementation of SimpleAccount must be deployed with the new EntryPoint address, then upgrading
     * the implementation by calling `upgradeTo()`
     */
    function initialize(address anOwner) public virtual initializer {
        _initialize(anOwner);
    }

    function _initialize(address anOwner) internal virtual {
        owner = anOwner;
    }

    function entryPoint() public view virtual returns (IEntryPoint) {
        return _entryPoint;
    }

    // updateProfile updates the profile data
    function updateProfile(
        string memory name,
        string memory description,
        string memory icon,
        string memory meta
    ) public onlyOwnerOrEntryPoint {
        profile = ProfileData(name, description, icon, meta);
    }

    // clearProfile clears the profile data
    function clearProfile() public onlyOwnerOrEntryPoint {
        profile = ProfileData("", "", "", "");
    }

    // getProfile returns the profile data
    function getProfile()
        public
        view
        returns (string memory, string memory, string icon, string memory)
    {
        return (profile.name, profile.description, profile.icon, profile.meta);
    }
}
