// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";

import "./TokenEntryPoint.sol";
import "./Paymaster.sol";

import "./interfaces/ITokenEntryPoint.sol";

/**
 * @title TokenEntryPointFactory
 * @dev Contract for creating new accounts and calculating their counterfactual addresses.
 *
 * https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/samples/SimpleAccountFactory.sol
 */
contract TokenEntryPointFactory {
    address public immutable entryPoint;
    TokenEntryPoint public immutable tokenEntryPointImplementation;
    Paymaster public immutable paymasterImplementation;

    event EntryPointCreated(
        address indexed owner,
        address tokenEntryPoint,
        address paymaster
    );

    constructor(IEntryPoint _entryPoint) {
        entryPoint = address(_entryPoint);
        tokenEntryPointImplementation = new TokenEntryPoint(_entryPoint);
        paymasterImplementation = new Paymaster();
    }

    /**
     * create an account, and return its address.
     * returns the address even if the account is already deployed.
     * Note that during UserOperation execution, this method is called only if the account is not deployed.
     * This method returns an existing account address so that entryPoint.getSenderAddress() would work even after account creation
     */
    function create(
        address owner,
        address sponsor,
        address[] memory whitelistedAddresses,
        uint256 salt
    ) public returns (TokenEntryPoint tokenEntryPoint, Paymaster paymaster) {
        (address _tokenEntryPoint, address _paymaster) = get(
            owner,
            sponsor,
            whitelistedAddresses,
            salt
        );

        if (_tokenEntryPoint.code.length > 0 && _paymaster.code.length > 0) {
            return (
                TokenEntryPoint(address(_tokenEntryPoint)),
                Paymaster(address(_paymaster))
            );
        }

        require(
            whitelistedAddresses.length > 0,
            "whitelistedAddresses is empty"
        );
        address firstWhitelistedAddress = whitelistedAddresses[0];

        bytes32 derivedSalt = keccak256(
            abi.encodePacked(owner, sponsor, firstWhitelistedAddress, salt)
        );

        paymaster = Paymaster(
            address(
                new ERC1967Proxy{salt: derivedSalt}(
                    address(paymasterImplementation),
                    abi.encodeCall(Paymaster.initialize, (sponsor))
                )
            )
        );

        paymaster.transferOwnership(owner);

        tokenEntryPoint = TokenEntryPoint(
            address(
                new ERC1967Proxy{salt: derivedSalt}(
                    address(tokenEntryPointImplementation),
                    abi.encodeCall(
                        TokenEntryPoint.initialize,
                        (owner, address(paymaster), whitelistedAddresses)
                    )
                )
            )
        );

        emit EntryPointCreated(
            owner,
            address(tokenEntryPoint),
            address(paymaster)
        );
    }

    /**
     * calculate the counterfactual address of this account as it would be returned by createAccount()
     */
    function get(
        address owner,
        address sponsor,
        address[] memory whitelistedAddresses,
        uint256 salt
    ) public view returns (address, address) {
        require(
            whitelistedAddresses.length > 0,
            "whitelistedAddresses is empty"
        );
        address firstWhitelistedAddress = whitelistedAddresses[0];

        bytes32 derivedSalt = keccak256(
            abi.encodePacked(owner, sponsor, firstWhitelistedAddress, salt)
        );

        address paymaster = Create2.computeAddress(
            derivedSalt,
            keccak256(
                abi.encodePacked(
                    type(ERC1967Proxy).creationCode,
                    abi.encode(
                        address(paymasterImplementation),
                        abi.encodeCall(Paymaster.initialize, (sponsor))
                    )
                )
            )
        );

        address tokenEntryPoint = Create2.computeAddress(
            derivedSalt,
            keccak256(
                abi.encodePacked(
                    type(ERC1967Proxy).creationCode,
                    abi.encode(
                        address(tokenEntryPointImplementation),
                        abi.encodeCall(
                            TokenEntryPoint.initialize,
                            (owner, address(paymaster), whitelistedAddresses)
                        )
                    )
                )
            )
        );

        return (tokenEntryPoint, paymaster);
    }
}
