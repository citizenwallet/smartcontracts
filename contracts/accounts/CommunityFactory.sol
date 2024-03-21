// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";

import "./TokenEntryPoint.sol";
import "./Paymaster.sol";

import "./interfaces/ITokenEntryPoint.sol";

/**
 * @title CommunityFactory
 * @dev Contract for creating new accounts and calculating their counterfactual addresses.
 *
 * https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/samples/SimpleAccountFactory.sol
 */
contract CommunityFactory {
    TokenEntryPoint public immutable tokenEntryPointImplementation;
    Paymaster public immutable paymasterImplementation;

    event CommunityCreated(address indexed account);

    constructor(IEntryPoint _entryPoint) {
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
        uint256 salt
    ) public returns (TokenEntryPoint ret) {
        address addr = get(owner, salt);

        emit CommunityCreated(addr);

        uint codeSize = addr.code.length;
        if (codeSize > 0) {
            return TokenEntryPoint(payable(addr));
        }

        Paymaster paymaster = Paymaster(
            payable(
                new ERC1967Proxy{salt: bytes32(salt)}(
                    address(paymasterImplementation),
                    abi.encodeCall(Paymaster.initialize, (owner))
                )
            )
        );

        paymaster.transferOwnership(owner);

        address[] memory whitelistedAddresses = new address[](0);

        ret = TokenEntryPoint(
            payable(
                new ERC1967Proxy{salt: bytes32(salt)}(
                    address(tokenEntryPointImplementation),
                    abi.encodeCall(
                        TokenEntryPoint.initialize,
                        (owner, address(paymaster), whitelistedAddresses)
                    )
                )
            )
        );
    }

    /**
     * calculate the counterfactual address of this account as it would be returned by createAccount()
     */
    function get(address owner, uint256 salt) public view returns (address) {
        address paymaster = Create2.computeAddress(
            bytes32(salt),
            keccak256(
                abi.encodePacked(
                    type(ERC1967Proxy).creationCode,
                    abi.encode(
                        address(paymasterImplementation),
                        abi.encodeCall(Paymaster.initialize, (owner))
                    )
                )
            )
        );

        address[] memory whitelistedAddresses = new address[](0);

        return
            Create2.computeAddress(
                bytes32(salt),
                keccak256(
                    abi.encodePacked(
                        type(ERC1967Proxy).creationCode,
                        abi.encode(
                            address(tokenEntryPointImplementation),
                            abi.encodeCall(
                                TokenEntryPoint.initialize,
                                (
                                    owner,
                                    address(paymaster),
                                    whitelistedAddresses
                                )
                            )
                        )
                    )
                )
            );
    }
}
