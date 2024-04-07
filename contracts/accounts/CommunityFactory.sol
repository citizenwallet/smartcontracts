// // SPDX-License-Identifier: MIT
// pragma solidity ^0.8.20;

// import "@openzeppelin/contracts/utils/Create2.sol";
// import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
// import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";

// import "./TokenEntryPoint.sol";
// import "./Paymaster.sol";
// import "./AccountFactory.sol";
// import "../apps/Profile.sol";

// import "./interfaces/ITokenEntryPoint.sol";

// /**
//  * @title CommunityFactory
//  * @dev Contract for creating new accounts and calculating their counterfactual addresses.
//  *
//  * https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/samples/SimpleAccountFactory.sol
//  */
// contract CommunityFactory {
//     address public immutable entryPoint;
//     TokenEntryPoint public immutable tokenEntryPointImplementation;
//     Paymaster public immutable paymasterImplementation;
//     AccountFactory public immutable accountFactoryImplementation;
//     Profile public immutable profileImplementation;

//     event CommunityCreated(
//         address indexed owner,
//         address tokenEntryPoint,
//         address paymaster,
//         address accountFactory,
//         address profile
//     );

//     constructor(IEntryPoint _entryPoint) {
//         entryPoint = address(_entryPoint);
//         tokenEntryPointImplementation = new TokenEntryPoint(_entryPoint);
//         paymasterImplementation = new Paymaster();
//         accountFactoryImplementation = new AccountFactory();
//         profileImplementation = new Profile();
//     }

//     /**
//      * create an account, and return its address.
//      * returns the address even if the account is already deployed.
//      * Note that during UserOperation execution, this method is called only if the account is not deployed.
//      * This method returns an existing account address so that entryPoint.getSenderAddress() would work even after account creation
//      */
//     function create(
//         address owner,
//         address sponsor,
//         address token,
//         uint256 salt
//     )
//         public
//         returns (
//             TokenEntryPoint tokenEntryPoint,
//             Paymaster paymaster,
//             AccountFactory accountFactory,
//             Profile profile
//         )
//     {
//         (
//             address _tokenEntryPoint,
//             address _paymaster,
//             address _accountFactory,
//             address _profile
//         ) = get(owner, sponsor, token, salt);

//         if (
//             _tokenEntryPoint.code.length > 0 &&
//             _paymaster.code.length > 0 &&
//             _accountFactory.code.length > 0 &&
//             _profile.code.length > 0
//         ) {
//             return (
//                 TokenEntryPoint(address(_tokenEntryPoint)),
//                 Paymaster(address(_paymaster)),
//                 AccountFactory(address(_accountFactory)),
//                 Profile(address(_profile))
//             );
//         }

//         bytes32 derivedSalt = keccak256(
//             abi.encodePacked(owner, sponsor, token, salt)
//         );

//         paymaster = Paymaster(
//             address(
//                 new ERC1967Proxy{salt: derivedSalt}(
//                     address(paymasterImplementation),
//                     abi.encodeCall(Paymaster.initialize, (sponsor))
//                 )
//             )
//         );

//         paymaster.transferOwnership(owner);

//         profile = Profile(
//             address(
//                 new ERC1967Proxy{salt: derivedSalt}(
//                     address(profileImplementation),
//                     abi.encodeCall(Profile.initialize, ())
//                 )
//             )
//         );

//         profile.transferOwnership(owner);

//         address[] memory whitelistedAddresses = new address[](2);
//         whitelistedAddresses[0] = address(profile);
//         whitelistedAddresses[1] = token;

//         tokenEntryPoint = TokenEntryPoint(
//             address(
//                 new ERC1967Proxy{salt: derivedSalt}(
//                     address(tokenEntryPointImplementation),
//                     abi.encodeCall(
//                         TokenEntryPoint.initialize,
//                         (owner, address(paymaster), whitelistedAddresses)
//                     )
//                 )
//             )
//         );

//         accountFactory = AccountFactory(
//             address(
//                 new ERC1967Proxy{salt: derivedSalt}(
//                     address(accountFactoryImplementation),
//                     abi.encodeCall(
//                         AccountFactory.initialize,
//                         (
//                             IEntryPoint(entryPoint),
//                             ITokenEntryPoint(tokenEntryPoint),
//                             owner
//                         )
//                     )
//                 )
//             )
//         );

//         emit CommunityCreated(
//             owner,
//             address(tokenEntryPoint),
//             address(paymaster),
//             address(accountFactory),
//             address(profile)
//         );
//     }

//     /**
//      * calculate the counterfactual address of this account as it would be returned by createAccount()
//      */
//     function get(
//         address owner,
//         address sponsor,
//         address token,
//         uint256 salt
//     ) public view returns (address, address, address, address) {
//         bytes32 derivedSalt = keccak256(
//             abi.encodePacked(owner, sponsor, token, salt)
//         );

//         address paymaster = Create2.computeAddress(
//             derivedSalt,
//             keccak256(
//                 abi.encodePacked(
//                     type(ERC1967Proxy).creationCode,
//                     abi.encode(
//                         address(paymasterImplementation),
//                         abi.encodeCall(Paymaster.initialize, (sponsor))
//                     )
//                 )
//             )
//         );

//         address profile = Create2.computeAddress(
//             derivedSalt,
//             keccak256(
//                 abi.encodePacked(
//                     type(ERC1967Proxy).creationCode,
//                     abi.encode(
//                         address(profileImplementation),
//                         abi.encodeCall(Profile.initialize, ())
//                     )
//                 )
//             )
//         );

//         address[] memory whitelistedAddresses = new address[](2);
//         whitelistedAddresses[0] = address(profile);
//         whitelistedAddresses[1] = token;

//         address tokenEntryPoint = Create2.computeAddress(
//             derivedSalt,
//             keccak256(
//                 abi.encodePacked(
//                     type(ERC1967Proxy).creationCode,
//                     abi.encode(
//                         address(tokenEntryPointImplementation),
//                         abi.encodeCall(
//                             TokenEntryPoint.initialize,
//                             (owner, address(paymaster), whitelistedAddresses)
//                         )
//                     )
//                 )
//             )
//         );

//         address accountFactory = Create2.computeAddress(
//             derivedSalt,
//             keccak256(
//                 abi.encodePacked(
//                     type(ERC1967Proxy).creationCode,
//                     abi.encode(
//                         address(accountFactoryImplementation),
//                         abi.encodeCall(
//                             AccountFactory.initialize,
//                             (
//                                 IEntryPoint(entryPoint),
//                                 ITokenEntryPoint(tokenEntryPoint),
//                                 owner
//                             )
//                         )
//                     )
//                 )
//             )
//         );

//         return (tokenEntryPoint, paymaster, accountFactory, profile);
//     }
// }
