// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";

import "../apps/Profile.sol";

import "./interfaces/ITokenEntryPoint.sol";

/**
 * @title ProfileFactory
 * @dev Contract for creating new accounts and calculating their counterfactual addresses.
 *
 * https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/samples/SimpleAccountFactory.sol
 */
contract ProfileFactory {
    Profile public immutable profileImplementation;

    event ProfileCreated(address indexed owner, address profile);

    constructor() {
        profileImplementation = new Profile();
    }

    /**
     * create an account, and return its address.
     * returns the address even if the account is already deployed.
     * Note that during UserOperation execution, this method is called only if the account is not deployed.
     * This method returns an existing account address so that entryPoint.getSenderAddress() would work even after account creation
     */
    function create(
        address owner,
        uint256 salt
    ) public returns (Profile profile) {
        address _profile = get(owner, salt);

        if (_profile.code.length > 0) {
            return (Profile(address(_profile)));
        }

        bytes32 derivedSalt = keccak256(abi.encodePacked(owner, salt));

        profile = Profile(
            address(
                new ERC1967Proxy{salt: derivedSalt}(
                    address(profileImplementation),
                    abi.encodeCall(Profile.initialize, ())
                )
            )
        );

        profile.transferOwnership(owner);

        emit ProfileCreated(owner, address(profile));
    }

    /**
     * calculate the counterfactual address of this account as it would be returned by createAccount()
     */
    function get(address owner, uint256 salt) public view returns (address) {
        bytes32 derivedSalt = keccak256(abi.encodePacked(owner, salt));

        address profile = Create2.computeAddress(
            derivedSalt,
            keccak256(
                abi.encodePacked(
                    type(ERC1967Proxy).creationCode,
                    abi.encode(
                        address(profileImplementation),
                        abi.encodeCall(Profile.initialize, ())
                    )
                )
            )
        );

        return (profile);
    }
}
