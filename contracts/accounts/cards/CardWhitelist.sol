// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/* solhint-disable avoid-low-level-calls */
/* solhint-disable no-inline-assembly */
/* solhint-disable reason-string */

import "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";

import "../../standards/ERC20.sol";

contract CardWhitelist is UUPSUpgradeable, Initializable {
    address public owner;

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyAccountOwner() {
        _checkAccountOwner();
        _;
    }

    function _checkAccountOwner() internal view virtual {
        require(
            owner == msg.sender || address(this) == msg.sender,
            "Ownable: caller is not the owner or the contract"
        );
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal view override onlyAccountOwner {
        (newImplementation);
    }
}
