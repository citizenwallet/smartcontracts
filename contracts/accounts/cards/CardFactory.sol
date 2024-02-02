// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

import "./Card.sol";
import "../interfaces/IWhitelist.sol";

/**
 * @title CardFactory
 * @dev Contract for creating new cards and calculating their counterfactual addresses.
 *
 * https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/samples/SimpleCardFactory.sol
 */
contract CardFactory {
    Card public immutable cardImplementation;

    address public immutable tokenAddress;
    IWhitelist public immutable whitelist;

    event CardCreated(address indexed account);

    constructor(address _tokenAddress, address _whitelist) {
        cardImplementation = new Card(_tokenAddress);
        tokenAddress = _tokenAddress;
        whitelist = IWhitelist(_whitelist);
    }

    /**
     * create an account, and return its address.
     * returns the address even if the account is already deployed.
     * Note that during UserOperation execution, this method is called only if the account is not deployed.
     * This method returns an existing account address so that entryPoint.getSenderAddress() would work even after account creation
     */
    function createCard(address owner, uint256 salt) public returns (Card ret) {
        address addr = getAddress(owner, salt);

        emit CardCreated(addr);

        uint codeSize = addr.code.length;
        if (codeSize > 0) {
            return Card(payable(addr));
        }
        ret = Card(
            payable(
                new ERC1967Proxy{salt: bytes32(salt)}(
                    address(cardImplementation),
                    abi.encodeCall(Card.initialize, (owner, address(whitelist)))
                )
            )
        );
    }

    /**
     * calculate the counterfactual address of this account as it would be returned by createCard()
     */
    function getAddress(
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
                            address(cardImplementation),
                            abi.encodeCall(
                                Card.initialize,
                                (owner, address(whitelist))
                            )
                        )
                    )
                )
            );
    }
}
