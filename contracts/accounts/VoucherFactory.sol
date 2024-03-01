// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";

import "./Voucher.sol";
import "./interfaces/ITokenEntryPoint.sol";

/**
 * @title VoucherFactory
 * @dev Contract for creating new accounts and calculating their counterfactual addresses.
 *
 * https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/samples/SimpleAccountFactory.sol
 */
contract VoucherFactory {
    Voucher public immutable voucherImplementation;

    event VoucherCreated(address indexed voucher);

    constructor(IEntryPoint _entryPoint, ITokenEntryPoint _tokenEntryPoint) {
        voucherImplementation = new Voucher(_entryPoint, _tokenEntryPoint);
    }

    /**
     * @dev Calculates the hash value for a given address and code.
     * This function should only be used to test hash values.
     *
     * @param code The code to be hashed.
     * @return The calculated hash value.
     */
    function getHash(
        uint256 code
    ) public view returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    code,
                    block.chainid,
                    address(this)
                )
            );
    }

    /**
     * create an account, and return its address.
     * returns the address even if the account is already deployed.
     * Note that during UserOperation execution, this method is called only if the account is not deployed.
     * This method returns an existing account address so that entryPoint.getSenderAddress() would work even after account creation
     */
    function createVoucher(
        address owner,
        uint256 code
    ) public returns (Voucher ret) {
        bytes32 codeHash = getHash(code);

        address addr = getAddress(codeHash);

        emit VoucherCreated(addr);

        uint codeSize = addr.code.length;
        if (codeSize > 0) {
            return Voucher(payable(addr));
        }
        ret = Voucher(
            payable(
                new ERC1967Proxy{salt: codeHash}(
                    address(voucherImplementation),
                    abi.encodeCall(Voucher.initialize, ())
                )
            )
        );

        ret.transferOwnership(owner);
    }

    /**
     * calculate the counterfactual address of this account as it would be returned by createAccount()
     */
    function getAddress(
        bytes32 codeHash
    ) public view returns (address) {
        return
            Create2.computeAddress(
                codeHash,
                keccak256(
                    abi.encodePacked(
                        type(ERC1967Proxy).creationCode,
                        abi.encode(
                            address(voucherImplementation),
                            abi.encodeCall(Voucher.initialize, ())
                        )
                    )
                )
            );
    }
}
