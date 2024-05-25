// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

contract Profile is
    Initializable,
    ERC721URIStorageUpgradeable,
    OwnableUpgradeable,
    AccessControlUpgradeable,
    UUPSUpgradeable
{
    bytes32 constant NULL = "";
    bytes32 public constant PROFILE_ADMIN_ROLE =
        keccak256("PROFILE_ADMIN_ROLE");

    /// only a single username per address
    ///
    /// allows us to check which username was already reserved by a profile
    mapping(address profile => bytes32 username) public usernames;

    /// only a single username per profile
    ///
    /// allows us to easily fetch a profile from a username
    mapping(bytes32 username => address profile) public profiles;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __ERC721_init("Profile", "PRF");
        __Ownable_init();
        __AccessControl_init();
        __UUPSUpgradeable_init();

        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    function set(
        address profile,
        bytes32 _username,
        string memory _uri
    ) public returns (uint256) {
        require(
            owner() == msg.sender ||
                profile == msg.sender ||
                hasRole(PROFILE_ADMIN_ROLE, msg.sender),
            "Only the profile owner or contract owner can set it."
        );

        bytes32 currentUsername = usernames[profile];

        require(
            (currentUsername == NULL &&
                (profiles[_username] == address(0) ||
                    profiles[_username] == profile)) || // username is not set and is also not reserved
                (currentUsername != NULL &&
                    profiles[currentUsername] == profile), // username is set but belongs to the profile
            "This username is already taken."
        );

        uint256 newProfileId = fromAddressToId(profile);
        if (!_exists(newProfileId)) {
            _mint(profile, newProfileId);
        }
        if (currentUsername != _username) {
            _setUsername(profile, _username);
        }
        if (
            keccak256(abi.encodePacked(tokenURI(newProfileId))) !=
            keccak256(abi.encodePacked(_uri))
        ) {
            _setTokenURI(newProfileId, _uri);
        }

        return newProfileId;
    }

    function get(address profile) public view returns (string memory) {
        uint256 profileId = fromAddressToId(profile);
        return tokenURI(profileId);
    }

    function getFromUsername(
        bytes32 username
    ) public view returns (string memory) {
        address profile = profiles[username];
        require(profile != address(0), "This username does not exist.");

        uint256 profileId = fromAddressToId(profile);
        return tokenURI(profileId);
    }

    function burn(uint256 tokenId) external {
        address profileOwner = fromIdToAddress(tokenId);
        require(
            owner() == msg.sender ||
                profileOwner == msg.sender ||
                hasRole(PROFILE_ADMIN_ROLE, msg.sender),
            "Only the owner of the token or profile can burn it."
        );
        _burn(tokenId);
    }

    function fromIdToAddress(uint256 tokenId) public pure returns (address) {
        return address(uint160(uint256(tokenId)));
    }

    function fromAddressToId(address profile) public pure returns (uint256) {
        return uint256(uint160(profile));
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 firstTokenId,
        uint256 batchSize
    ) internal pure override {
        (firstTokenId);
        (batchSize);
        require(
            from == address(0) || to == address(0),
            "This a Soulbound token. It cannot be transferred. It can only be burned by the token owner."
        );
    }

    function _burn(
        uint256 tokenId
    ) internal override(ERC721URIStorageUpgradeable) {
        _unsetUsername(fromIdToAddress(tokenId));
        super._burn(tokenId);
    }

    function _setUsername(address profile, bytes32 username) internal {
        if (usernames[profile] != NULL) {
            // clean up old username if it exists
            delete profiles[usernames[profile]];
        }

        profiles[username] = profile;
        usernames[profile] = username;
    }

    function _unsetUsername(address _profile) internal {
        bytes32 username = usernames[_profile];

        delete profiles[username];
        delete usernames[_profile];
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}

    function supportsInterface(
        bytes4 interfaceId
    )
        public
        view
        override(ERC721URIStorageUpgradeable, AccessControlUpgradeable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}
