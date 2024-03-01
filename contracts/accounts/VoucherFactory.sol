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
 * @dev A factory contract that creates Voucher contracts.
 */
contract VoucherFactory {
    Voucher public immutable voucherImplementation;

    event VoucherCreated(address indexed voucher);

    constructor(IEntryPoint _entryPoint, ITokenEntryPoint _tokenEntryPoint) {
        voucherImplementation = new Voucher(_entryPoint, _tokenEntryPoint);
    }

    /**
     * @dev Calculates the hash value for a given code.
     * This function should only be used to test hash values.
     *
     * @param code The code to be hashed.
     * @return The calculated hash value.
     */
    function getCodeHash(
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
     * claims a voucher, and return its address.
     * returns the address even if the voucher is already deployed.
     * Note that during UserOperation execution, this method is called only if the voucher is not deployed.
     * This method returns an existing voucher address so that entryPoint.getSenderAddress() would work even after voucher creation
     */
    function claimVoucher(
        address owner,
        uint256 code
    ) public returns (Voucher ret) {
        bytes32 codeHash = getCodeHash(code);

        address addr = getVoucherAddress(codeHash);

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
     * calculate the counterfactual address of this voucher as it would be returned by claimVoucher()
     */
    function getVoucherAddress(
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
