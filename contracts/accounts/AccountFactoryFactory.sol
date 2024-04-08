// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";

import "./AccountFactory.sol";

import "./interfaces/ITokenEntryPoint.sol";

/**
 * @title AccountFactoryFactory
 * @dev Contract for creating new accounts and calculating their counterfactual addresses.
 *
 * https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/samples/SimpleAccountFactory.sol
 */
contract AccountFactoryFactory {
    address public immutable entryPoint;
    AccountFactory public immutable accountFactoryImplementation;

    event AccountFactoryCreated(address indexed owner, address accountFactory);

    constructor(IEntryPoint _entryPoint) {
        entryPoint = address(_entryPoint);
        accountFactoryImplementation = new AccountFactory();
    }

    /**
     * create an account, and return its address.
     * returns the address even if the account is already deployed.
     * Note that during UserOperation execution, this method is called only if the account is not deployed.
     * This method returns an existing account address so that entryPoint.getSenderAddress() would work even after account creation
     */
    function create(
        address owner,
        address tokenEntryPoint,
        uint256 salt
    ) public returns (AccountFactory accountFactory) {
        address _accountFactory = get(owner, tokenEntryPoint, salt);

        if (_accountFactory.code.length > 0) {
            return (AccountFactory(address(_accountFactory)));
        }

        bytes32 derivedSalt = keccak256(
            abi.encodePacked(owner, tokenEntryPoint, salt)
        );

        accountFactory = AccountFactory(
            address(
                new ERC1967Proxy{salt: derivedSalt}(
                    address(accountFactoryImplementation),
                    abi.encodeCall(
                        AccountFactory.initialize,
                        (
                            IEntryPoint(entryPoint),
                            ITokenEntryPoint(tokenEntryPoint),
                            owner
                        )
                    )
                )
            )
        );

        emit AccountFactoryCreated(owner, address(accountFactory));
    }

    /**
     * calculate the counterfactual address of this account as it would be returned by createAccount()
     */
    function get(
        address owner,
        address tokenEntryPoint,
        uint256 salt
    ) public view returns (address) {
        bytes32 derivedSalt = keccak256(
            abi.encodePacked(owner, tokenEntryPoint, salt)
        );

        address accountFactory = Create2.computeAddress(
            derivedSalt,
            keccak256(
                abi.encodePacked(
                    type(ERC1967Proxy).creationCode,
                    abi.encode(
                        address(accountFactoryImplementation),
                        abi.encodeCall(
                            AccountFactory.initialize,
                            (
                                IEntryPoint(entryPoint),
                                ITokenEntryPoint(tokenEntryPoint),
                                owner
                            )
                        )
                    )
                )
            )
        );

        return (accountFactory);
    }
}
