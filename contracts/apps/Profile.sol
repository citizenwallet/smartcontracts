// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";

import "../abstract/Linkable.sol";

contract Profile is Initializable, Linkable {
    IEntryPoint private immutable _entryPoint;

    address public owner;

    struct ProfileData {
        string name;
        string description;
        string meta;
    }

    ProfileData public profile;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor(IEntryPoint entryPoint) {
        _entryPoint = entryPoint;
        profile = ProfileData("", "", "");
        _disableInitializers();
    }

    modifier onlyOwner() override {
        _onlyOwner();
        _;
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

    function _onlyOwner() internal view {
        //directly from EOA owner, or through the account itself (which gets redirected through execute())
        require(
            msg.sender == owner || msg.sender == address(this),
            "only owner"
        );
    }

    // updateProfile updates the profile data
    function updateProfile(
        string memory name,
        string memory description,
        string memory meta
    ) public onlyOwner {
        profile = ProfileData(name, description, meta);
    }

    // clearProfile clears the profile data
    function clearProfile() public onlyOwner {
        profile = ProfileData("", "", "");
    }

    // getProfile returns the profile data
    function getProfile()
        public
        view
        returns (string memory, string memory, string memory)
    {
        return (profile.name, profile.description, profile.meta);
    }
}
