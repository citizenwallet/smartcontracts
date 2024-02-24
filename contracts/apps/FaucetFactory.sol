// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";

import "./SimpleFaucet.sol";
import "./RedeemCodeFaucet.sol";

contract FaucetFactory {
    SimpleFaucet public immutable simpleFaucetImplementation;
    RedeemCodeFaucet public immutable redeemCodeFaucetImplementation;

    event SimpleFaucetCreated(address indexed faucet, address token, uint256 amount, uint48 redeemInterval);
    event RedeemCodeFaucetCreated(address indexed faucet, address token, uint256 amount, uint48 redeemInterval);

    constructor() {
        simpleFaucetImplementation = new SimpleFaucet();
        redeemCodeFaucetImplementation = new RedeemCodeFaucet();
    }


    function createSimpleFaucet(
        address owner,
        uint256 salt,
        IERC20Upgradeable _token, 
        uint256 _amount, 
        uint48 _redeemInterval, 
        address _redeemAdmin
    ) public returns (SimpleFaucet ret) {
        address addr = getSimpleFaucetAddress(owner, salt, _token, _amount, _redeemInterval, _redeemAdmin);

        emit SimpleFaucetCreated(addr, address(_token), _amount, _redeemInterval);

        uint codeSize = addr.code.length;
        if (codeSize > 0) {
            return SimpleFaucet(payable(addr));
        }
        ret = SimpleFaucet(
            payable(
                new ERC1967Proxy{salt: bytes32(salt)}(
                    address(simpleFaucetImplementation),
                    abi.encodeCall(SimpleFaucet.initialize, (owner, _token, _amount, _redeemInterval, _redeemAdmin))
                )
            )
        );
    }

    function getSimpleFaucetAddress(
        address owner,
        uint256 salt,
        IERC20Upgradeable _token, 
        uint256 _amount, 
        uint48 _redeemInterval, 
        address _redeemAdmin
    ) public view returns (address) {
        return
            Create2.computeAddress(
                bytes32(salt),
                keccak256(
                    abi.encodePacked(
                        type(ERC1967Proxy).creationCode,
                        abi.encode(
                            address(simpleFaucetImplementation),
                            abi.encodeCall(SimpleFaucet.initialize, (owner, _token, _amount, _redeemInterval, _redeemAdmin))
                        )
                    )
                )
            );
    }

    function createRedeemCodeFaucet(
        address owner,
        uint256 salt,
        IERC20Upgradeable _token, 
        uint48 _redeemInterval, 
        address _codeCreator
    ) public returns (RedeemCodeFaucet ret) {
        address addr = getRedeemCodeFaucetAddress(owner, salt, _token, _redeemInterval, _codeCreator);

        emit RedeemCodeFaucetCreated(addr, address(_token), 0, _redeemInterval);

        uint codeSize = addr.code.length;
        if (codeSize > 0) {
            return RedeemCodeFaucet(payable(addr));
        }
        ret = RedeemCodeFaucet(
            payable(
                new ERC1967Proxy{salt: bytes32(salt)}(
                    address(redeemCodeFaucetImplementation),
                    abi.encodeCall(RedeemCodeFaucet.initialize, (owner, _token, _redeemInterval, _codeCreator))
                )
            )
        );
    }

    function getRedeemCodeFaucetAddress(
        address owner,
        uint256 salt,
        IERC20Upgradeable _token, 
        uint48 _redeemInterval, 
        address _codeCreator
    ) public view returns (address) {
        return
            Create2.computeAddress(
                bytes32(salt),
                keccak256(
                    abi.encodePacked(
                        type(ERC1967Proxy).creationCode,
                        abi.encode(
                            address(redeemCodeFaucetImplementation),
                            abi.encodeCall(RedeemCodeFaucet.initialize, (owner, _token, _redeemInterval, _codeCreator))
                        )
                    )
                )
            );
    }
}
