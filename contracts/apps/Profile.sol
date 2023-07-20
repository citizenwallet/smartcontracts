// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

contract Profile is
    Initializable,
    ERC721URIStorageUpgradeable,
    OwnableUpgradeable,
    UUPSUpgradeable
{
    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __ERC721_init("Profile", "PRF");
        __Ownable_init();
        __UUPSUpgradeable_init();
    }

    function set(address profile, string memory _uri) public returns (uint256) {
        require(profile == msg.sender, "Only the profile owner can set it.");

        uint256 newProfileId = _fromAddressToId(profile);
        if (_exists(newProfileId)) {
            _setTokenURI(newProfileId, _uri);
            return newProfileId;
        }
        _mint(profile, newProfileId);
        _setTokenURI(newProfileId, _uri);

        return newProfileId;
    }

    function get(address profile) public view returns (string memory) {
        uint256 profileId = _fromAddressToId(profile);
        return tokenURI(profileId);
    }

    function burn(uint256 tokenId) external {
        address profileOwner = _fromIdToAddress(tokenId);
        require(
            owner() == msg.sender || profileOwner == msg.sender,
            "Only the owner of the token or profile can burn it."
        );
        _burn(tokenId);
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 firstTokenId,
        uint256 batchSize
    ) internal pure override {
        require(
            from == address(0) || to == address(0),
            "This a Soulbound token. It cannot be transferred. It can only be burned by the token owner."
        );
    }

    function _burn(
        uint256 tokenId
    ) internal override(ERC721URIStorageUpgradeable) {
        super._burn(tokenId);
    }

    function _fromIdToAddress(uint256 tokenId) internal pure returns (address) {
        return address(uint160(uint256(tokenId)));
    }

    function _fromAddressToId(address profile) internal pure returns (uint256) {
        return uint256(uint160(profile));
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}
}
