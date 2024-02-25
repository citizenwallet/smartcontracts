// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package redeemCodeFaucet

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RedeemCodeFaucetMetaData contains all meta data concerning the RedeemCodeFaucet contract.
var RedeemCodeFaucetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REDEEM_CODE_CREATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"redeemableAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"}],\"name\":\"addRedeemCode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"codeCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_codeCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"_redeemInterval\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"_codeCreator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"lastRedeem\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"time\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemInterval\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"}],\"name\":\"redeemed\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"time\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b50608051611e9a61004c60003960008181610795015281816107d501528181610976015281816109b60152610a450152611e9a6000f3fe60806040526004361061014b5760003560e01c806391d14854116100b6578063d547741f1161006f578063d547741f146103eb578063db006a751461040b578063e93950961461042b578063ed05582b1461044b578063f2fde38b1461046b578063fc0c546a1461048b57600080fd5b806391d1485414610316578063931409081461033657806393c0a4aa14610358578063a217fddf14610379578063a5ff3f801461038e578063c0a9444c146103b557600080fd5b80633ccfd60b116101085780633ccfd60b146102455780633d8b9e071461025a5780634f1ef286146102a757806352d1902d146102ba578063715018a6146102cf5780638da5cb5b146102e457600080fd5b806301ffc9a714610150578063163df9e814610185578063248a9ca3146101a75780632f2ff15d146101e557806336568abe146102055780633659cfe614610225575b600080fd5b34801561015c57600080fd5b5061017061016b36600461190f565b6104ac565b60405190151581526020015b60405180910390f35b34801561019157600080fd5b506101a56101a0366004611969565b6104e3565b005b3480156101b357600080fd5b506101d76101c23660046119c3565b60009081526097602052604090206001015490565b60405190815260200161017c565b3480156101f157600080fd5b506101a56102003660046119dc565b610671565b34801561021157600080fd5b506101a56102203660046119dc565b61070d565b34801561023157600080fd5b506101a5610240366004611a0c565b61078b565b34801561025157600080fd5b506101a561086a565b34801561026657600080fd5b506102906102753660046119c3565b6101316020526000908152604090205465ffffffffffff1681565b60405165ffffffffffff909116815260200161017c565b6101a56102b5366004611a3f565b61096c565b3480156102c657600080fd5b506101d7610a38565b3480156102db57600080fd5b506101a5610aeb565b3480156102f057600080fd5b506033546001600160a01b03165b6040516001600160a01b03909116815260200161017c565b34801561032257600080fd5b506101706103313660046119dc565b610aff565b34801561034257600080fd5b506101d7600080516020611dfe83398151915281565b34801561036457600080fd5b5061012e546102fe906001600160a01b031681565b34801561038557600080fd5b506101d7600081565b34801561039a57600080fd5b5061012d5461029090600160a01b900465ffffffffffff1681565b3480156103c157600080fd5b506102906103d0366004611a0c565b6101326020526000908152604090205465ffffffffffff1681565b3480156103f757600080fd5b506101a56104063660046119dc565b610b2a565b34801561041757600080fd5b506101a56104263660046119c3565b610b4f565b34801561043757600080fd5b506101a5610446366004611b03565b610e94565b34801561045757600080fd5b506101d7610466366004611b38565b611044565b34801561047757600080fd5b506101a5610486366004611a0c565b61109b565b34801561049757600080fd5b5061012d546102fe906001600160a01b031681565b60006001600160e01b03198216637965db0b60e01b14806104dd57506301ffc9a760e01b6001600160e01b03198316145b92915050565b600054610100900460ff16158080156105035750600054600160ff909116105b8061051d5750303b15801561051d575060005460ff166001145b6105855760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156105a8576000805461ff0019166101001790555b6105b0611111565b6105b8611140565b61012d80546001600160a01b038681166001600160d01b031990921691909117600160a01b65ffffffffffff8716021790915561012e80546001600160a01b03191691841691909117905561061b600080516020611dfe83398151915283611167565b6106248561109b565b801561066a576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b60008281526097602052604090206001015461068c81611171565b600080516020611dfe83398151915283036106fe5760405162461bcd60e51b815260206004820152602c60248201527f416363657373436f6e74726f6c3a2063616e6e6f74206772616e7420636f646560448201526b43726561746f7220726f6c6560a01b606482015260840161057c565b610708838361117b565b505050565b6001600160a01b038116331461077d5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b606482015260840161057c565b6107878282611201565b5050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036107d35760405162461bcd60e51b815260040161057c90611b64565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661081c600080516020611e1e833981519152546001600160a01b031690565b6001600160a01b0316146108425760405162461bcd60e51b815260040161057c90611bb0565b61084b81611268565b6040805160008082526020820190925261086791839190611270565b50565b600080516020611dfe83398151915261088281611171565b61012d5461012e546040516370a0823160e01b81523060048201526001600160a01b039283169263a9059cbb92169083906370a0823190602401602060405180830381865afa1580156108d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fd9190611bfc565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af1158015610948573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107879190611c15565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036109b45760405162461bcd60e51b815260040161057c90611b64565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166109fd600080516020611e1e833981519152546001600160a01b031690565b6001600160a01b031614610a235760405162461bcd60e51b815260040161057c90611bb0565b610a2c82611268565b61078782826001611270565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ad85760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000606482015260840161057c565b50600080516020611e1e83398151915290565b610af36113db565b610afd6000611435565b565b60009182526097602090815260408084206001600160a01b0393909316845291905290205460ff1690565b600082815260976020526040902060010154610b4581611171565b6107088383611201565b61012d543360009081526101326020526040902054429165ffffffffffff80841692610b8592600160a01b900482169116611c4d565b65ffffffffffff1610610bef5760405162461bcd60e51b815260206004820152602c60248201527f52656465656d436f64654661756365743a2072656465656d20696e746572766160448201526b1b081b9bdd081c185cdcd95960a21b606482015260840161057c565b61012e54600090610c09906001600160a01b031684611044565b6000818152610131602052604090205490915065ffffffffffff1615610c815760405162461bcd60e51b815260206004820152602760248201527f52656465656d436f64654661756365743a20636f646520616c72656164792072604482015266195919595b595960ca1b606482015260840161057c565b6000818152610130602052604090205465ffffffffffff90811690831610610ceb5760405162461bcd60e51b815260206004820152601e60248201527f52656465656d436f64654661756365743a20636f646520657870697265640000604482015260640161057c565b600081815261012f6020526040908190205461012d5491516370a0823160e01b815230600482015290916001600160a01b0316906370a0823190602401602060405180830381865afa158015610d45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d699190611bfc565b1015610dc65760405162461bcd60e51b815260206004820152602660248201527f52656465656d436f64654661756365743a20696e73756666696369656e742062604482015265616c616e636560d01b606482015260840161057c565b600081815261012f6020526040908190205461012d54915163a9059cbb60e01b81523360048201526024810182905290916001600160a01b03169063a9059cbb906044016020604051808303816000875af1158015610e29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e4d9190611c15565b5050600090815261013160209081526040808320805465ffffffffffff90951665ffffffffffff199586168117909155338452610132909252909120805490921617905550565b600080516020611dfe833981519152610eac81611171565b60008311610f225760405162461bcd60e51b815260206004820152603c60248201527f52656465656d436f64654661756365743a2072656465656d61626c65416d6f7560448201527f6e74206d7573742062652067726561746572207468616e207a65726f00000000606482015260840161057c565b4265ffffffffffff80821690841611610f985760405162461bcd60e51b815260206004820152603260248201527f52656465656d436f64654661756365743a2076616c6964556e74696c206d75736044820152717420626520696e207468652066757475726560701b606482015260840161057c565b6000858152610131602052604090205465ffffffffffff16156110085760405162461bcd60e51b815260206004820152602260248201527f52656465656d436f64654661756365743a20616c72656164792072656465656d604482015261195960f21b606482015260840161057c565b5050600092835261012f602090815260408085209390935561013090529120805465ffffffffffff191665ffffffffffff909216919091179055565b6040516bffffffffffffffffffffffff19606084811b821660208401526034830184905246605484015230901b16607482015260009060880160405160208183030381529060405280519060200120905092915050565b6110a36113db565b6001600160a01b0381166111085760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161057c565b61086781611435565b600054610100900460ff166111385760405162461bcd60e51b815260040161057c90611c73565b610afd611487565b600054610100900460ff16610afd5760405162461bcd60e51b815260040161057c90611c73565b610787828261117b565b61086781336114b7565b6111858282610aff565b6107875760008281526097602090815260408083206001600160a01b03851684529091529020805460ff191660011790556111bd3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b61120b8282610aff565b156107875760008281526097602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6108676113db565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff16156112a35761070883611510565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156112fd575060408051601f3d908101601f191682019092526112fa91810190611bfc565b60015b6113605760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b606482015260840161057c565b600080516020611e1e83398151915281146113cf5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b606482015260840161057c565b506107088383836115ac565b6033546001600160a01b03163314610afd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161057c565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff166114ae5760405162461bcd60e51b815260040161057c90611c73565b610afd33611435565b6114c18282610aff565b610787576114ce816115d7565b6114d98360206115e9565b6040516020016114ea929190611ce2565b60408051601f198184030181529082905262461bcd60e51b825261057c91600401611d57565b6001600160a01b0381163b61157d5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b606482015260840161057c565b600080516020611e1e83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6115b58361178c565b6000825111806115c25750805b15610708576115d183836117cc565b50505050565b60606104dd6001600160a01b03831660145b606060006115f8836002611d8a565b611603906002611da1565b67ffffffffffffffff81111561161b5761161b611a29565b6040519080825280601f01601f191660200182016040528015611645576020820181803683370190505b509050600360fc1b8160008151811061166057611660611db4565b60200101906001600160f81b031916908160001a905350600f60fb1b8160018151811061168f5761168f611db4565b60200101906001600160f81b031916908160001a90535060006116b3846002611d8a565b6116be906001611da1565b90505b6001811115611736576f181899199a1a9b1b9c1cb0b131b232b360811b85600f16601081106116f2576116f2611db4565b1a60f81b82828151811061170857611708611db4565b60200101906001600160f81b031916908160001a90535060049490941c9361172f81611dca565b90506116c1565b5083156117855760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161057c565b9392505050565b61179581611510565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606117858383604051806060016040528060278152602001611e3e602791396060600080856001600160a01b0316856040516118099190611de1565b600060405180830381855af49150503d8060008114611844576040519150601f19603f3d011682016040523d82523d6000602084013e611849565b606091505b509150915061185a86838387611864565b9695505050505050565b606083156118d35782516000036118cc576001600160a01b0385163b6118cc5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161057c565b50816118dd565b6118dd83836118e5565b949350505050565b8151156118f55781518083602001fd5b8060405162461bcd60e51b815260040161057c9190611d57565b60006020828403121561192157600080fd5b81356001600160e01b03198116811461178557600080fd5b6001600160a01b038116811461086757600080fd5b803565ffffffffffff8116811461196457600080fd5b919050565b6000806000806080858703121561197f57600080fd5b843561198a81611939565b9350602085013561199a81611939565b92506119a86040860161194e565b915060608501356119b881611939565b939692955090935050565b6000602082840312156119d557600080fd5b5035919050565b600080604083850312156119ef57600080fd5b823591506020830135611a0181611939565b809150509250929050565b600060208284031215611a1e57600080fd5b813561178581611939565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215611a5257600080fd5b8235611a5d81611939565b9150602083013567ffffffffffffffff80821115611a7a57600080fd5b818501915085601f830112611a8e57600080fd5b813581811115611aa057611aa0611a29565b604051601f8201601f19908116603f01168101908382118183101715611ac857611ac8611a29565b81604052828152886020848701011115611ae157600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b600080600060608486031215611b1857600080fd5b8335925060208401359150611b2f6040850161194e565b90509250925092565b60008060408385031215611b4b57600080fd5b8235611b5681611939565b946020939093013593505050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b600060208284031215611c0e57600080fd5b5051919050565b600060208284031215611c2757600080fd5b8151801515811461178557600080fd5b634e487b7160e01b600052601160045260246000fd5b65ffffffffffff818116838216019080821115611c6c57611c6c611c37565b5092915050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60005b83811015611cd9578181015183820152602001611cc1565b50506000910152565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351611d1a816017850160208801611cbe565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351611d4b816028840160208801611cbe565b01602801949350505050565b6020815260008251806020840152611d76816040850160208701611cbe565b601f01601f19169190910160400192915050565b80820281158282048414176104dd576104dd611c37565b808201808211156104dd576104dd611c37565b634e487b7160e01b600052603260045260246000fd5b600081611dd957611dd9611c37565b506000190190565b60008251611df3818460208701611cbe565b919091019291505056fe707a390ecff002f977980b7598d697e51a8a18e7708b8f7ada4c1e67e5ea4808360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220cbabb29dee83815abe0d0ddaefd573d4979e60ef6019204ef1c2fc60633c50f664736f6c63430008140033",
}

// RedeemCodeFaucetABI is the input ABI used to generate the binding from.
// Deprecated: Use RedeemCodeFaucetMetaData.ABI instead.
var RedeemCodeFaucetABI = RedeemCodeFaucetMetaData.ABI

// RedeemCodeFaucetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RedeemCodeFaucetMetaData.Bin instead.
var RedeemCodeFaucetBin = RedeemCodeFaucetMetaData.Bin

// DeployRedeemCodeFaucet deploys a new Ethereum contract, binding an instance of RedeemCodeFaucet to it.
func DeployRedeemCodeFaucet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RedeemCodeFaucet, error) {
	parsed, err := RedeemCodeFaucetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RedeemCodeFaucetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RedeemCodeFaucet{RedeemCodeFaucetCaller: RedeemCodeFaucetCaller{contract: contract}, RedeemCodeFaucetTransactor: RedeemCodeFaucetTransactor{contract: contract}, RedeemCodeFaucetFilterer: RedeemCodeFaucetFilterer{contract: contract}}, nil
}

// RedeemCodeFaucet is an auto generated Go binding around an Ethereum contract.
type RedeemCodeFaucet struct {
	RedeemCodeFaucetCaller     // Read-only binding to the contract
	RedeemCodeFaucetTransactor // Write-only binding to the contract
	RedeemCodeFaucetFilterer   // Log filterer for contract events
}

// RedeemCodeFaucetCaller is an auto generated read-only Go binding around an Ethereum contract.
type RedeemCodeFaucetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemCodeFaucetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RedeemCodeFaucetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemCodeFaucetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RedeemCodeFaucetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemCodeFaucetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RedeemCodeFaucetSession struct {
	Contract     *RedeemCodeFaucet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RedeemCodeFaucetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RedeemCodeFaucetCallerSession struct {
	Contract *RedeemCodeFaucetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// RedeemCodeFaucetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RedeemCodeFaucetTransactorSession struct {
	Contract     *RedeemCodeFaucetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RedeemCodeFaucetRaw is an auto generated low-level Go binding around an Ethereum contract.
type RedeemCodeFaucetRaw struct {
	Contract *RedeemCodeFaucet // Generic contract binding to access the raw methods on
}

// RedeemCodeFaucetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RedeemCodeFaucetCallerRaw struct {
	Contract *RedeemCodeFaucetCaller // Generic read-only contract binding to access the raw methods on
}

// RedeemCodeFaucetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RedeemCodeFaucetTransactorRaw struct {
	Contract *RedeemCodeFaucetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRedeemCodeFaucet creates a new instance of RedeemCodeFaucet, bound to a specific deployed contract.
func NewRedeemCodeFaucet(address common.Address, backend bind.ContractBackend) (*RedeemCodeFaucet, error) {
	contract, err := bindRedeemCodeFaucet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucet{RedeemCodeFaucetCaller: RedeemCodeFaucetCaller{contract: contract}, RedeemCodeFaucetTransactor: RedeemCodeFaucetTransactor{contract: contract}, RedeemCodeFaucetFilterer: RedeemCodeFaucetFilterer{contract: contract}}, nil
}

// NewRedeemCodeFaucetCaller creates a new read-only instance of RedeemCodeFaucet, bound to a specific deployed contract.
func NewRedeemCodeFaucetCaller(address common.Address, caller bind.ContractCaller) (*RedeemCodeFaucetCaller, error) {
	contract, err := bindRedeemCodeFaucet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetCaller{contract: contract}, nil
}

// NewRedeemCodeFaucetTransactor creates a new write-only instance of RedeemCodeFaucet, bound to a specific deployed contract.
func NewRedeemCodeFaucetTransactor(address common.Address, transactor bind.ContractTransactor) (*RedeemCodeFaucetTransactor, error) {
	contract, err := bindRedeemCodeFaucet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetTransactor{contract: contract}, nil
}

// NewRedeemCodeFaucetFilterer creates a new log filterer instance of RedeemCodeFaucet, bound to a specific deployed contract.
func NewRedeemCodeFaucetFilterer(address common.Address, filterer bind.ContractFilterer) (*RedeemCodeFaucetFilterer, error) {
	contract, err := bindRedeemCodeFaucet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetFilterer{contract: contract}, nil
}

// bindRedeemCodeFaucet binds a generic wrapper to an already deployed contract.
func bindRedeemCodeFaucet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RedeemCodeFaucetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RedeemCodeFaucet *RedeemCodeFaucetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RedeemCodeFaucet.Contract.RedeemCodeFaucetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RedeemCodeFaucet *RedeemCodeFaucetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RedeemCodeFaucetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RedeemCodeFaucet *RedeemCodeFaucetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RedeemCodeFaucetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RedeemCodeFaucet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.DEFAULTADMINROLE(&_RedeemCodeFaucet.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.DEFAULTADMINROLE(&_RedeemCodeFaucet.CallOpts)
}

// REDEEMCODECREATORROLE is a free data retrieval call binding the contract method 0x93140908.
//
// Solidity: function REDEEM_CODE_CREATOR_ROLE() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) REDEEMCODECREATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "REDEEM_CODE_CREATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REDEEMCODECREATORROLE is a free data retrieval call binding the contract method 0x93140908.
//
// Solidity: function REDEEM_CODE_CREATOR_ROLE() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) REDEEMCODECREATORROLE() ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.REDEEMCODECREATORROLE(&_RedeemCodeFaucet.CallOpts)
}

// REDEEMCODECREATORROLE is a free data retrieval call binding the contract method 0x93140908.
//
// Solidity: function REDEEM_CODE_CREATOR_ROLE() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) REDEEMCODECREATORROLE() ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.REDEEMCODECREATORROLE(&_RedeemCodeFaucet.CallOpts)
}

// CodeCreator is a free data retrieval call binding the contract method 0x93c0a4aa.
//
// Solidity: function codeCreator() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) CodeCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "codeCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CodeCreator is a free data retrieval call binding the contract method 0x93c0a4aa.
//
// Solidity: function codeCreator() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) CodeCreator() (common.Address, error) {
	return _RedeemCodeFaucet.Contract.CodeCreator(&_RedeemCodeFaucet.CallOpts)
}

// CodeCreator is a free data retrieval call binding the contract method 0x93c0a4aa.
//
// Solidity: function codeCreator() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) CodeCreator() (common.Address, error) {
	return _RedeemCodeFaucet.Contract.CodeCreator(&_RedeemCodeFaucet.CallOpts)
}

// GetHash is a free data retrieval call binding the contract method 0xed05582b.
//
// Solidity: function getHash(address _codeCreator, uint256 code) view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) GetHash(opts *bind.CallOpts, _codeCreator common.Address, code *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "getHash", _codeCreator, code)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0xed05582b.
//
// Solidity: function getHash(address _codeCreator, uint256 code) view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) GetHash(_codeCreator common.Address, code *big.Int) ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.GetHash(&_RedeemCodeFaucet.CallOpts, _codeCreator, code)
}

// GetHash is a free data retrieval call binding the contract method 0xed05582b.
//
// Solidity: function getHash(address _codeCreator, uint256 code) view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) GetHash(_codeCreator common.Address, code *big.Int) ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.GetHash(&_RedeemCodeFaucet.CallOpts, _codeCreator, code)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.GetRoleAdmin(&_RedeemCodeFaucet.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.GetRoleAdmin(&_RedeemCodeFaucet.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RedeemCodeFaucet.Contract.HasRole(&_RedeemCodeFaucet.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RedeemCodeFaucet.Contract.HasRole(&_RedeemCodeFaucet.CallOpts, role, account)
}

// LastRedeem is a free data retrieval call binding the contract method 0xc0a9444c.
//
// Solidity: function lastRedeem(address sender) view returns(uint48 time)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) LastRedeem(opts *bind.CallOpts, sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "lastRedeem", sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRedeem is a free data retrieval call binding the contract method 0xc0a9444c.
//
// Solidity: function lastRedeem(address sender) view returns(uint48 time)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) LastRedeem(sender common.Address) (*big.Int, error) {
	return _RedeemCodeFaucet.Contract.LastRedeem(&_RedeemCodeFaucet.CallOpts, sender)
}

// LastRedeem is a free data retrieval call binding the contract method 0xc0a9444c.
//
// Solidity: function lastRedeem(address sender) view returns(uint48 time)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) LastRedeem(sender common.Address) (*big.Int, error) {
	return _RedeemCodeFaucet.Contract.LastRedeem(&_RedeemCodeFaucet.CallOpts, sender)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) Owner() (common.Address, error) {
	return _RedeemCodeFaucet.Contract.Owner(&_RedeemCodeFaucet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) Owner() (common.Address, error) {
	return _RedeemCodeFaucet.Contract.Owner(&_RedeemCodeFaucet.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) ProxiableUUID() ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.ProxiableUUID(&_RedeemCodeFaucet.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) ProxiableUUID() ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.ProxiableUUID(&_RedeemCodeFaucet.CallOpts)
}

// RedeemInterval is a free data retrieval call binding the contract method 0xa5ff3f80.
//
// Solidity: function redeemInterval() view returns(uint48)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) RedeemInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "redeemInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemInterval is a free data retrieval call binding the contract method 0xa5ff3f80.
//
// Solidity: function redeemInterval() view returns(uint48)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) RedeemInterval() (*big.Int, error) {
	return _RedeemCodeFaucet.Contract.RedeemInterval(&_RedeemCodeFaucet.CallOpts)
}

// RedeemInterval is a free data retrieval call binding the contract method 0xa5ff3f80.
//
// Solidity: function redeemInterval() view returns(uint48)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) RedeemInterval() (*big.Int, error) {
	return _RedeemCodeFaucet.Contract.RedeemInterval(&_RedeemCodeFaucet.CallOpts)
}

// Redeemed is a free data retrieval call binding the contract method 0x3d8b9e07.
//
// Solidity: function redeemed(bytes32 codeHash) view returns(uint48 time)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) Redeemed(opts *bind.CallOpts, codeHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "redeemed", codeHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Redeemed is a free data retrieval call binding the contract method 0x3d8b9e07.
//
// Solidity: function redeemed(bytes32 codeHash) view returns(uint48 time)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) Redeemed(codeHash [32]byte) (*big.Int, error) {
	return _RedeemCodeFaucet.Contract.Redeemed(&_RedeemCodeFaucet.CallOpts, codeHash)
}

// Redeemed is a free data retrieval call binding the contract method 0x3d8b9e07.
//
// Solidity: function redeemed(bytes32 codeHash) view returns(uint48 time)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) Redeemed(codeHash [32]byte) (*big.Int, error) {
	return _RedeemCodeFaucet.Contract.Redeemed(&_RedeemCodeFaucet.CallOpts, codeHash)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RedeemCodeFaucet.Contract.SupportsInterface(&_RedeemCodeFaucet.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RedeemCodeFaucet.Contract.SupportsInterface(&_RedeemCodeFaucet.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) Token() (common.Address, error) {
	return _RedeemCodeFaucet.Contract.Token(&_RedeemCodeFaucet.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) Token() (common.Address, error) {
	return _RedeemCodeFaucet.Contract.Token(&_RedeemCodeFaucet.CallOpts)
}

// AddRedeemCode is a paid mutator transaction binding the contract method 0xe9395096.
//
// Solidity: function addRedeemCode(bytes32 codeHash, uint256 redeemableAmount, uint48 validUntil) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) AddRedeemCode(opts *bind.TransactOpts, codeHash [32]byte, redeemableAmount *big.Int, validUntil *big.Int) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "addRedeemCode", codeHash, redeemableAmount, validUntil)
}

// AddRedeemCode is a paid mutator transaction binding the contract method 0xe9395096.
//
// Solidity: function addRedeemCode(bytes32 codeHash, uint256 redeemableAmount, uint48 validUntil) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) AddRedeemCode(codeHash [32]byte, redeemableAmount *big.Int, validUntil *big.Int) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.AddRedeemCode(&_RedeemCodeFaucet.TransactOpts, codeHash, redeemableAmount, validUntil)
}

// AddRedeemCode is a paid mutator transaction binding the contract method 0xe9395096.
//
// Solidity: function addRedeemCode(bytes32 codeHash, uint256 redeemableAmount, uint48 validUntil) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) AddRedeemCode(codeHash [32]byte, redeemableAmount *big.Int, validUntil *big.Int) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.AddRedeemCode(&_RedeemCodeFaucet.TransactOpts, codeHash, redeemableAmount, validUntil)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.GrantRole(&_RedeemCodeFaucet.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.GrantRole(&_RedeemCodeFaucet.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x163df9e8.
//
// Solidity: function initialize(address owner, address _token, uint48 _redeemInterval, address _codeCreator) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, _token common.Address, _redeemInterval *big.Int, _codeCreator common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "initialize", owner, _token, _redeemInterval, _codeCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0x163df9e8.
//
// Solidity: function initialize(address owner, address _token, uint48 _redeemInterval, address _codeCreator) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) Initialize(owner common.Address, _token common.Address, _redeemInterval *big.Int, _codeCreator common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.Initialize(&_RedeemCodeFaucet.TransactOpts, owner, _token, _redeemInterval, _codeCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0x163df9e8.
//
// Solidity: function initialize(address owner, address _token, uint48 _redeemInterval, address _codeCreator) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) Initialize(owner common.Address, _token common.Address, _redeemInterval *big.Int, _codeCreator common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.Initialize(&_RedeemCodeFaucet.TransactOpts, owner, _token, _redeemInterval, _codeCreator)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 code) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) Redeem(opts *bind.TransactOpts, code *big.Int) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "redeem", code)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 code) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) Redeem(code *big.Int) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.Redeem(&_RedeemCodeFaucet.TransactOpts, code)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 code) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) Redeem(code *big.Int) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.Redeem(&_RedeemCodeFaucet.TransactOpts, code)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) RenounceOwnership() (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RenounceOwnership(&_RedeemCodeFaucet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RenounceOwnership(&_RedeemCodeFaucet.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RenounceRole(&_RedeemCodeFaucet.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RenounceRole(&_RedeemCodeFaucet.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RevokeRole(&_RedeemCodeFaucet.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RevokeRole(&_RedeemCodeFaucet.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.TransferOwnership(&_RedeemCodeFaucet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.TransferOwnership(&_RedeemCodeFaucet.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.UpgradeTo(&_RedeemCodeFaucet.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.UpgradeTo(&_RedeemCodeFaucet.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.UpgradeToAndCall(&_RedeemCodeFaucet.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.UpgradeToAndCall(&_RedeemCodeFaucet.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) Withdraw() (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.Withdraw(&_RedeemCodeFaucet.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) Withdraw() (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.Withdraw(&_RedeemCodeFaucet.TransactOpts)
}

// RedeemCodeFaucetAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetAdminChangedIterator struct {
	Event *RedeemCodeFaucetAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RedeemCodeFaucetAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RedeemCodeFaucetAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RedeemCodeFaucetAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetAdminChanged represents a AdminChanged event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*RedeemCodeFaucetAdminChangedIterator, error) {

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetAdminChangedIterator{contract: _RedeemCodeFaucet.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetAdminChanged) (event.Subscription, error) {

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetAdminChanged)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseAdminChanged(log types.Log) (*RedeemCodeFaucetAdminChanged, error) {
	event := new(RedeemCodeFaucetAdminChanged)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemCodeFaucetBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetBeaconUpgradedIterator struct {
	Event *RedeemCodeFaucetBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RedeemCodeFaucetBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RedeemCodeFaucetBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RedeemCodeFaucetBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetBeaconUpgraded represents a BeaconUpgraded event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*RedeemCodeFaucetBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetBeaconUpgradedIterator{contract: _RedeemCodeFaucet.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetBeaconUpgraded)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseBeaconUpgraded(log types.Log) (*RedeemCodeFaucetBeaconUpgraded, error) {
	event := new(RedeemCodeFaucetBeaconUpgraded)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemCodeFaucetInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetInitializedIterator struct {
	Event *RedeemCodeFaucetInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RedeemCodeFaucetInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RedeemCodeFaucetInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RedeemCodeFaucetInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetInitialized represents a Initialized event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterInitialized(opts *bind.FilterOpts) (*RedeemCodeFaucetInitializedIterator, error) {

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetInitializedIterator{contract: _RedeemCodeFaucet.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetInitialized) (event.Subscription, error) {

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetInitialized)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseInitialized(log types.Log) (*RedeemCodeFaucetInitialized, error) {
	event := new(RedeemCodeFaucetInitialized)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemCodeFaucetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetOwnershipTransferredIterator struct {
	Event *RedeemCodeFaucetOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RedeemCodeFaucetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RedeemCodeFaucetOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RedeemCodeFaucetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetOwnershipTransferred represents a OwnershipTransferred event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RedeemCodeFaucetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetOwnershipTransferredIterator{contract: _RedeemCodeFaucet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetOwnershipTransferred)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseOwnershipTransferred(log types.Log) (*RedeemCodeFaucetOwnershipTransferred, error) {
	event := new(RedeemCodeFaucetOwnershipTransferred)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemCodeFaucetRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetRoleAdminChangedIterator struct {
	Event *RedeemCodeFaucetRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RedeemCodeFaucetRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RedeemCodeFaucetRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RedeemCodeFaucetRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetRoleAdminChanged represents a RoleAdminChanged event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RedeemCodeFaucetRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetRoleAdminChangedIterator{contract: _RedeemCodeFaucet.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetRoleAdminChanged)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseRoleAdminChanged(log types.Log) (*RedeemCodeFaucetRoleAdminChanged, error) {
	event := new(RedeemCodeFaucetRoleAdminChanged)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemCodeFaucetRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetRoleGrantedIterator struct {
	Event *RedeemCodeFaucetRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RedeemCodeFaucetRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RedeemCodeFaucetRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RedeemCodeFaucetRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetRoleGranted represents a RoleGranted event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RedeemCodeFaucetRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetRoleGrantedIterator{contract: _RedeemCodeFaucet.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetRoleGranted)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseRoleGranted(log types.Log) (*RedeemCodeFaucetRoleGranted, error) {
	event := new(RedeemCodeFaucetRoleGranted)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemCodeFaucetRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetRoleRevokedIterator struct {
	Event *RedeemCodeFaucetRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RedeemCodeFaucetRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RedeemCodeFaucetRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RedeemCodeFaucetRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetRoleRevoked represents a RoleRevoked event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RedeemCodeFaucetRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetRoleRevokedIterator{contract: _RedeemCodeFaucet.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetRoleRevoked)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseRoleRevoked(log types.Log) (*RedeemCodeFaucetRoleRevoked, error) {
	event := new(RedeemCodeFaucetRoleRevoked)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemCodeFaucetUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetUpgradedIterator struct {
	Event *RedeemCodeFaucetUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RedeemCodeFaucetUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RedeemCodeFaucetUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RedeemCodeFaucetUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetUpgraded represents a Upgraded event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RedeemCodeFaucetUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetUpgradedIterator{contract: _RedeemCodeFaucet.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetUpgraded)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseUpgraded(log types.Log) (*RedeemCodeFaucetUpgraded, error) {
	event := new(RedeemCodeFaucetUpgraded)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
