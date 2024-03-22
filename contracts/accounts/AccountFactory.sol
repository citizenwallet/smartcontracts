// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";

import "./Account.sol";
import "./interfaces/ITokenEntryPoint.sol";

/**
 * @title AccountFactory
 * @dev Contract for creating new accounts and calculating their counterfactual addresses.
 *
 * https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/samples/SimpleAccountFactory.sol
 */
contract AccountFactory is Initializable, OwnableUpgradeable, UUPSUpgradeable {
    /// @custom:oz-upgrades-unsafe-allow state-variable-immutable
    Account public accountImplementation;

    event AccountCreated(address indexed account);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(
        IEntryPoint _entryPoint,
        ITokenEntryPoint _tokenEntryPoint,
        address anOwner
    ) public virtual initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();

        transferOwnership(anOwner);

        accountImplementation = new Account(_entryPoint, _tokenEntryPoint);
    }

    /**
     * create an account, and return its address.
     * returns the address even if the account is already deployed.
     * Note that during UserOperation execution, this method is called only if the account is not deployed.
     * This method returns an existing account address so that entryPoint.getSenderAddress() would work even after account creation
     */
    function createAccount(
        address owner,
        uint256 salt
    ) public returns (Account ret) {
        address addr = getAddress(owner, salt);

        emit AccountCreated(addr);

        uint codeSize = addr.code.length;
        if (codeSize > 0) {
            return Account(payable(addr));
        }
        ret = Account(
            payable(
                new ERC1967Proxy{salt: bytes32(salt)}(
                    address(accountImplementation),
                    abi.encodeCall(Account.initialize, (owner))
                )
            )
        );
    }

    /**
     * calculate the counterfactual address of this account as it would be returned by createAccount()
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
                            address(accountImplementation),
                            abi.encodeCall(Account.initialize, (owner))
                        )
                    )
                )
            );
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal view override onlyOwner {
        (newImplementation);
    }
}
