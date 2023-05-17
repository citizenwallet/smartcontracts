// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";

import "./GratitudeToken.sol";

/**
 * A sample factory contract for GratitudeToken
 * A UserOperations "initCode" holds the address of the factory, and a method call (to createGratitudeToken, in this sample factory).
 * The factory's createGratitudeToken returns the target account address even if it is already installed.
 * This way, the entryPoint.getSenderAddress() can be called either before or after the account is created.
 */
contract AccountFactory {
    GratitudeToken public immutable accountImplementation;

    constructor(IEntryPoint _entryPoint) {
        accountImplementation = new GratitudeToken(_entryPoint);
    }

    /**
     * create an account, and return its address.
     * returns the address even if the account is already deployed.
     * Note that during UserOperation execution, this method is called only if the account is not deployed.
     * This method returns an existing account address so that entryPoint.getSenderAddress() would work even after account creation
     */
    function createGratitudeToken(
        address owner,
        uint256 salt
    ) public returns (GratitudeToken ret) {
        address addr = getGratitudeTokenAddress(owner, salt);
        uint codeSize = addr.code.length;
        if (codeSize > 0) {
            return GratitudeToken(payable(addr));
        }
        ret = GratitudeToken(
            payable(
                new ERC1967Proxy{salt: bytes32(salt)}(
                    address(accountImplementation),
                    abi.encodeCall(GratitudeToken.initialize, (owner))
                )
            )
        );
    }

    /**
     * calculate the counterfactual address of this account as it would be returned by createGratitudeToken()
     */
    function getGratitudeTokenAddress(
        address owner,
        uint256 salt
    ) public view returns (address) {
        return
            Create2.computeAddress(
                bytes32(salt),
                keccak256(
                    abi.encodePacked(
                        type(ERC1967Proxy).creationCode,
                        abi.encode(
                            address(accountImplementation),
                            abi.encodeCall(GratitudeToken.initialize, (owner))
                        )
                    )
                )
            );
    }
}
