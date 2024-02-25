// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/Create2.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";
import "@openzeppelin/contracts/utils/introspection/IERC165.sol";

import "./SimpleFaucet.sol";
import "./RedeemCodeFaucet.sol";

/**
 * @title FaucetFactory
 * @dev A contract for creating different types of faucets.
 */
contract FaucetFactory {
    SimpleFaucet public immutable simpleFaucetImplementation;
    RedeemCodeFaucet public immutable redeemCodeFaucetImplementation;

    event SimpleFaucetCreated(address indexed faucet, address token, uint256 amount, uint48 redeemInterval);
    event RedeemCodeFaucetCreated(address indexed faucet, address token, uint256 amount, uint48 redeemInterval);

    /**
     * @dev Constructor function that initializes the faucet implementations.
     */
    constructor() {
        simpleFaucetImplementation = new SimpleFaucet();
        redeemCodeFaucetImplementation = new RedeemCodeFaucet();
    }

    /**
     * @dev Creates a new instance of a SimpleFaucet contract.
     * @param owner The address of the owner of the faucet.
     * @param salt The salt value used for contract address generation.
     * @param _token The address of the ERC20 token to be used by the faucet.
     * @param _amount The amount of tokens to be initially deposited into the faucet.
     * @param _redeemInterval The time interval (in seconds) between each redemption of tokens from the faucet.
     * @param _redeemAdmin The address of the redemption admin for the faucet.
     * @return ret The created SimpleFaucet contract instance.
     */
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

    /**
     * @dev Computes the address of a SimpleFaucet contract using the Create2 function.
     * @param owner The address of the owner of the SimpleFaucet contract.
     * @param salt The salt value used in the Create2 computation.
     * @param _token The address of the ERC20 token used by the SimpleFaucet contract.
     * @param _amount The amount of tokens to be deposited in the SimpleFaucet contract.
     * @param _redeemInterval The interval (in seconds) between each redemption of tokens from the SimpleFaucet contract.
     * @param _redeemAdmin The address of the admin account that can redeem tokens from the SimpleFaucet contract.
     * @return The computed address of the SimpleFaucet contract.
     */
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

    /**
     * @dev Creates a new RedeemCodeFaucet contract.
     * @param owner The address of the owner of the faucet.
     * @param salt The salt value used for contract deployment.
     * @param _token The ERC20 token contract address.
     * @param _redeemInterval The interval (in seconds) between each redemption.
     * @param _codeCreator The address of the code creator.
     * @return ret The created RedeemCodeFaucet contract.
     */
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

    /**
     * @dev Computes the address of a RedeemCodeFaucet contract using the Create2 function.
     * @param owner The address of the owner of the RedeemCodeFaucet contract.
     * @param salt The salt value used in the Create2 computation.
     * @param _token The address of the ERC20 token used by the RedeemCodeFaucet contract.
     * @param _redeemInterval The interval (in seconds) between each redemption in the RedeemCodeFaucet contract.
     * @param _codeCreator The address of the code creator for the RedeemCodeFaucet contract.
     * @return The computed address of the RedeemCodeFaucet contract.
     */
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
