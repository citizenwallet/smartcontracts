// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gratitudeToken

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

// GratitudeTokenMetaData contains all meta data concerning the GratitudeToken contract.
var GratitudeTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"}],\"name\":\"mintToMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintToMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"mintToMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a060405230608052348015610013575f80fd5b5061001c610021565b6100de565b5f54610100900460ff161561008c5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811610156100dc575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b608051611d4d6101125f395f818161051e015281816105670152818161064201528181610682015261070f0152611d4d5ff3fe608060405260043610610147575f3560e01c806370a08231116100b3578063925247251161006d578063925247251461038157806395d89b41146103a0578063a457c2d7146103b4578063a9059cbb146103d3578063dd62ed3e146103f2578063f2fde38b14610411575f80fd5b806370a08231146102c0578063715018a6146102f457806379cc6790146103085780638129fc1c146103275780638a0b03601461033b5780638da5cb5b1461035a575f80fd5b80633659cfe6116101045780633659cfe61461021d578063395093511461023c57806340c10f191461025b57806342966c681461027a5780634f1ef2861461029957806352d1902d146102ac575f80fd5b806306fdde031461014b578063095ea7b31461017557806318160ddd146101a4578063203abe34146101c257806323b872dd146101e3578063313ce56714610202575b5f80fd5b348015610156575f80fd5b5061015f610430565b60405161016c9190611670565b60405180910390f35b348015610180575f80fd5b5061019461018f3660046116bd565b6104c0565b604051901515815260200161016c565b3480156101af575f80fd5b506035545b60405190815260200161016c565b3480156101cd575f80fd5b506101e16101dc3660046117bc565b6104d9565b005b3480156101ee575f80fd5b506101946101fd3660046117f6565b6104f6565b34801561020d575f80fd5b506040516012815260200161016c565b348015610228575f80fd5b506101e161023736600461182f565b610514565b348015610247575f80fd5b506101946102563660046116bd565b6105f7565b348015610266575f80fd5b506101e16102753660046116bd565b610618565b348015610285575f80fd5b506101e1610294366004611848565b61062e565b6101e16102a736600461185f565b610638565b3480156102b7575f80fd5b506101b4610703565b3480156102cb575f80fd5b506101b46102da36600461182f565b6001600160a01b03165f9081526033602052604090205490565b3480156102ff575f80fd5b506101e16107b4565b348015610313575f80fd5b506101e16103223660046116bd565b6107c7565b348015610332575f80fd5b506101e16107dc565b348015610346575f80fd5b506101e16103553660046118ff565b610941565b348015610365575f80fd5b506097546040516001600160a01b03909116815260200161016c565b34801561038c575f80fd5b506101e161039b366004611941565b610993565b3480156103ab575f80fd5b5061015f610a76565b3480156103bf575f80fd5b506101946103ce3660046116bd565b610a85565b3480156103de575f80fd5b506101946103ed3660046116bd565b610b0a565b3480156103fd575f80fd5b506101b461040c3660046119f6565b610b1d565b34801561041c575f80fd5b506101e161042b36600461182f565b610b47565b60606036805461043f90611a27565b80601f016020809104026020016040519081016040528092919081815260200182805461046b90611a27565b80156104b65780601f1061048d576101008083540402835291602001916104b6565b820191905f5260205f20905b81548152906001019060200180831161049957829003601f168201915b5050505050905090565b5f336104cd818585610bbd565b60019150505b92915050565b6104e1610ce0565b6104f381670de0b6b3a7640000610941565b50565b5f6104ff610ce0565b61050a848484610d3a565b90505b9392505050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036105655760405162461bcd60e51b815260040161055c90611a5f565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166105ad5f80516020611cd1833981519152546001600160a01b031690565b6001600160a01b0316146105d35760405162461bcd60e51b815260040161055c90611aab565b6105dc81610d52565b604080515f808252602082019092526104f391839190610d5a565b5f336104cd8185856106098383610b1d565b6106139190611b0b565b610bbd565b610620610ce0565b61062a8282610ec4565b5050565b6104f33382610f83565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036106805760405162461bcd60e51b815260040161055c90611a5f565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166106c85f80516020611cd1833981519152546001600160a01b031690565b6001600160a01b0316146106ee5760405162461bcd60e51b815260040161055c90611aab565b6106f782610d52565b61062a82826001610d5a565b5f306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107a25760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000606482015260840161055c565b505f80516020611cd183398151915290565b6107bc610ce0565b6107c55f6110b5565b565b6107d2823383611106565b61062a8282610f83565b5f54610100900460ff16158080156107fa57505f54600160ff909116105b806108135750303b15801561081357505f5460ff166001145b6108765760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161055c565b5f805460ff191660011790558015610897575f805461ff0019166101001790555b6108e26040518060400160405280600e81526020016d23b930ba34ba3ab232aa37b5b2b760911b8152506040518060400160405280600381526020016211d51560ea1b81525061117e565b6108ea6111ae565b6108f26111d4565b6108fa6111ae565b80156104f3575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b610949610ce0565b5f5b825181101561098e575f83828151811061096757610967611b1e565b6020026020010151905061097b8184610ec4565b508061098681611b32565b91505061094b565b505050565b61099b610ce0565b8051825114610a125760405162461bcd60e51b815260206004820152603760248201527f726563697069656e747320616e6420616d6f756e747320617272617973206d7560448201527f73742068617665207468652073616d65206c656e677468000000000000000000606482015260840161055c565b5f5b825181101561098e575f838281518110610a3057610a30611b1e565b602002602001015190505f838381518110610a4d57610a4d611b1e565b60200260200101519050610a618282610ec4565b50508080610a6e90611b32565b915050610a14565b60606037805461043f90611a27565b5f3381610a928286610b1d565b905083811015610af25760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b606482015260840161055c565b610aff8286868403610bbd565b506001949350505050565b5f610b13610ce0565b61050d8383611202565b6001600160a01b039182165f90815260346020908152604080832093909416825291909152205490565b610b4f610ce0565b6001600160a01b038116610bb45760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161055c565b6104f3816110b5565b6001600160a01b038316610c1f5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b606482015260840161055c565b6001600160a01b038216610c805760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b606482015260840161055c565b6001600160a01b038381165f8181526034602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6097546001600160a01b031633146107c55760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161055c565b5f33610d47858285611106565b610aff85858561120b565b6104f3610ce0565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610d8d5761098e836113b4565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610de7575060408051601f3d908101601f19168201909252610de491810190611b4a565b60015b610e4a5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b606482015260840161055c565b5f80516020611cd18339815191528114610eb85760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b606482015260840161055c565b5061098e83838361144f565b6001600160a01b038216610f1a5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161055c565b8060355f828254610f2b9190611b0b565b90915550506001600160a01b0382165f818152603360209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6001600160a01b038216610fe35760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b606482015260840161055c565b6001600160a01b0382165f90815260336020526040902054818110156110565760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b606482015260840161055c565b6001600160a01b0383165f8181526033602090815260408083208686039055603580548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b609780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f6111118484610b1d565b90505f198114611178578181101561116b5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161055c565b6111788484848403610bbd565b50505050565b5f54610100900460ff166111a45760405162461bcd60e51b815260040161055c90611b61565b61062a8282611473565b5f54610100900460ff166107c55760405162461bcd60e51b815260040161055c90611b61565b5f54610100900460ff166111fa5760405162461bcd60e51b815260040161055c90611b61565b6107c56114b2565b5f336104cd8185855b6001600160a01b03831661126f5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161055c565b6001600160a01b0382166112d15760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b606482015260840161055c565b6001600160a01b0383165f90815260336020526040902054818110156113485760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b606482015260840161055c565b6001600160a01b038085165f8181526033602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906113a79086815260200190565b60405180910390a3611178565b6001600160a01b0381163b6114215760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b606482015260840161055c565b5f80516020611cd183398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b611458836114e1565b5f825111806114645750805b1561098e576111788383611520565b5f54610100900460ff166114995760405162461bcd60e51b815260040161055c90611b61565b60366114a58382611bf9565b50603761098e8282611bf9565b5f54610100900460ff166114d85760405162461bcd60e51b815260040161055c90611b61565b6107c5336110b5565b6114ea816113b4565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606001600160a01b0383163b6115885760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b606482015260840161055c565b5f80846001600160a01b0316846040516115a29190611cb5565b5f60405180830381855af49150503d805f81146115da576040519150601f19603f3d011682016040523d82523d5f602084013e6115df565b606091505b50915091506116078282604051806060016040528060278152602001611cf160279139611610565b95945050505050565b6060831561161f57508161050d565b61050d83838151156116345781518083602001fd5b8060405162461bcd60e51b815260040161055c9190611670565b5f5b83811015611668578181015183820152602001611650565b50505f910152565b602081525f825180602084015261168e81604085016020870161164e565b601f01601f19169190910160400192915050565b80356001600160a01b03811681146116b8575f80fd5b919050565b5f80604083850312156116ce575f80fd5b6116d7836116a2565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611722576117226116e5565b604052919050565b5f67ffffffffffffffff821115611743576117436116e5565b5060051b60200190565b5f82601f83011261175c575f80fd5b8135602061177161176c8361172a565b6116f9565b82815260059290921b8401810191818101908684111561178f575f80fd5b8286015b848110156117b1576117a4816116a2565b8352918301918301611793565b509695505050505050565b5f602082840312156117cc575f80fd5b813567ffffffffffffffff8111156117e2575f80fd5b6117ee8482850161174d565b949350505050565b5f805f60608486031215611808575f80fd5b611811846116a2565b925061181f602085016116a2565b9150604084013590509250925092565b5f6020828403121561183f575f80fd5b61050d826116a2565b5f60208284031215611858575f80fd5b5035919050565b5f8060408385031215611870575f80fd5b611879836116a2565b915060208084013567ffffffffffffffff80821115611896575f80fd5b818601915086601f8301126118a9575f80fd5b8135818111156118bb576118bb6116e5565b6118cd601f8201601f191685016116f9565b915080825287848285010111156118e2575f80fd5b80848401858401375f848284010152508093505050509250929050565b5f8060408385031215611910575f80fd5b823567ffffffffffffffff811115611926575f80fd5b6119328582860161174d565b95602094909401359450505050565b5f8060408385031215611952575f80fd5b823567ffffffffffffffff80821115611969575f80fd5b6119758683870161174d565b935060209150818501358181111561198b575f80fd5b85019050601f8101861361199d575f80fd5b80356119ab61176c8261172a565b81815260059190911b820183019083810190888311156119c9575f80fd5b928401925b828410156119e7578335825292840192908401906119ce565b80955050505050509250929050565b5f8060408385031215611a07575f80fd5b611a10836116a2565b9150611a1e602084016116a2565b90509250929050565b600181811c90821680611a3b57607f821691505b602082108103611a5957634e487b7160e01b5f52602260045260245ffd5b50919050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b808201808211156104d3576104d3611af7565b634e487b7160e01b5f52603260045260245ffd5b5f60018201611b4357611b43611af7565b5060010190565b5f60208284031215611b5a575f80fd5b5051919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b601f82111561098e575f81815260208120601f850160051c81016020861015611bd25750805b601f850160051c820191505b81811015611bf157828155600101611bde565b505050505050565b815167ffffffffffffffff811115611c1357611c136116e5565b611c2781611c218454611a27565b84611bac565b602080601f831160018114611c5a575f8415611c435750858301515b5f19600386901b1c1916600185901b178555611bf1565b5f85815260208120601f198616915b82811015611c8857888601518255948401946001909101908401611c69565b5085821015611ca557878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f8251611cc681846020870161164e565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212201079ee5f9cf0494aa73640fa8d326d79b0e7770dda5c2748bb4d523c4a0bcef164736f6c63430008140033",
}

// GratitudeTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use GratitudeTokenMetaData.ABI instead.
var GratitudeTokenABI = GratitudeTokenMetaData.ABI

// GratitudeTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GratitudeTokenMetaData.Bin instead.
var GratitudeTokenBin = GratitudeTokenMetaData.Bin

// DeployGratitudeToken deploys a new Ethereum contract, binding an instance of GratitudeToken to it.
func DeployGratitudeToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GratitudeToken, error) {
	parsed, err := GratitudeTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GratitudeTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GratitudeToken{GratitudeTokenCaller: GratitudeTokenCaller{contract: contract}, GratitudeTokenTransactor: GratitudeTokenTransactor{contract: contract}, GratitudeTokenFilterer: GratitudeTokenFilterer{contract: contract}}, nil
}

// GratitudeToken is an auto generated Go binding around an Ethereum contract.
type GratitudeToken struct {
	GratitudeTokenCaller     // Read-only binding to the contract
	GratitudeTokenTransactor // Write-only binding to the contract
	GratitudeTokenFilterer   // Log filterer for contract events
}

// GratitudeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type GratitudeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GratitudeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GratitudeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GratitudeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GratitudeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GratitudeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GratitudeTokenSession struct {
	Contract     *GratitudeToken   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GratitudeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GratitudeTokenCallerSession struct {
	Contract *GratitudeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GratitudeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GratitudeTokenTransactorSession struct {
	Contract     *GratitudeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GratitudeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type GratitudeTokenRaw struct {
	Contract *GratitudeToken // Generic contract binding to access the raw methods on
}

// GratitudeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GratitudeTokenCallerRaw struct {
	Contract *GratitudeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// GratitudeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GratitudeTokenTransactorRaw struct {
	Contract *GratitudeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGratitudeToken creates a new instance of GratitudeToken, bound to a specific deployed contract.
func NewGratitudeToken(address common.Address, backend bind.ContractBackend) (*GratitudeToken, error) {
	contract, err := bindGratitudeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GratitudeToken{GratitudeTokenCaller: GratitudeTokenCaller{contract: contract}, GratitudeTokenTransactor: GratitudeTokenTransactor{contract: contract}, GratitudeTokenFilterer: GratitudeTokenFilterer{contract: contract}}, nil
}

// NewGratitudeTokenCaller creates a new read-only instance of GratitudeToken, bound to a specific deployed contract.
func NewGratitudeTokenCaller(address common.Address, caller bind.ContractCaller) (*GratitudeTokenCaller, error) {
	contract, err := bindGratitudeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GratitudeTokenCaller{contract: contract}, nil
}

// NewGratitudeTokenTransactor creates a new write-only instance of GratitudeToken, bound to a specific deployed contract.
func NewGratitudeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*GratitudeTokenTransactor, error) {
	contract, err := bindGratitudeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GratitudeTokenTransactor{contract: contract}, nil
}

// NewGratitudeTokenFilterer creates a new log filterer instance of GratitudeToken, bound to a specific deployed contract.
func NewGratitudeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*GratitudeTokenFilterer, error) {
	contract, err := bindGratitudeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GratitudeTokenFilterer{contract: contract}, nil
}

// bindGratitudeToken binds a generic wrapper to an already deployed contract.
func bindGratitudeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GratitudeTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GratitudeToken *GratitudeTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GratitudeToken.Contract.GratitudeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GratitudeToken *GratitudeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GratitudeToken.Contract.GratitudeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GratitudeToken *GratitudeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GratitudeToken.Contract.GratitudeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GratitudeToken *GratitudeTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GratitudeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GratitudeToken *GratitudeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GratitudeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GratitudeToken *GratitudeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GratitudeToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_GratitudeToken *GratitudeTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GratitudeToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_GratitudeToken *GratitudeTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _GratitudeToken.Contract.Allowance(&_GratitudeToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_GratitudeToken *GratitudeTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _GratitudeToken.Contract.Allowance(&_GratitudeToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_GratitudeToken *GratitudeTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GratitudeToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_GratitudeToken *GratitudeTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _GratitudeToken.Contract.BalanceOf(&_GratitudeToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_GratitudeToken *GratitudeTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _GratitudeToken.Contract.BalanceOf(&_GratitudeToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GratitudeToken *GratitudeTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GratitudeToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GratitudeToken *GratitudeTokenSession) Decimals() (uint8, error) {
	return _GratitudeToken.Contract.Decimals(&_GratitudeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GratitudeToken *GratitudeTokenCallerSession) Decimals() (uint8, error) {
	return _GratitudeToken.Contract.Decimals(&_GratitudeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GratitudeToken *GratitudeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GratitudeToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GratitudeToken *GratitudeTokenSession) Name() (string, error) {
	return _GratitudeToken.Contract.Name(&_GratitudeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GratitudeToken *GratitudeTokenCallerSession) Name() (string, error) {
	return _GratitudeToken.Contract.Name(&_GratitudeToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GratitudeToken *GratitudeTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GratitudeToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GratitudeToken *GratitudeTokenSession) Owner() (common.Address, error) {
	return _GratitudeToken.Contract.Owner(&_GratitudeToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GratitudeToken *GratitudeTokenCallerSession) Owner() (common.Address, error) {
	return _GratitudeToken.Contract.Owner(&_GratitudeToken.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GratitudeToken *GratitudeTokenCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GratitudeToken.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GratitudeToken *GratitudeTokenSession) ProxiableUUID() ([32]byte, error) {
	return _GratitudeToken.Contract.ProxiableUUID(&_GratitudeToken.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GratitudeToken *GratitudeTokenCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GratitudeToken.Contract.ProxiableUUID(&_GratitudeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GratitudeToken *GratitudeTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GratitudeToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GratitudeToken *GratitudeTokenSession) Symbol() (string, error) {
	return _GratitudeToken.Contract.Symbol(&_GratitudeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GratitudeToken *GratitudeTokenCallerSession) Symbol() (string, error) {
	return _GratitudeToken.Contract.Symbol(&_GratitudeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GratitudeToken *GratitudeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GratitudeToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GratitudeToken *GratitudeTokenSession) TotalSupply() (*big.Int, error) {
	return _GratitudeToken.Contract.TotalSupply(&_GratitudeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GratitudeToken *GratitudeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _GratitudeToken.Contract.TotalSupply(&_GratitudeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_GratitudeToken *GratitudeTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_GratitudeToken *GratitudeTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.Approve(&_GratitudeToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_GratitudeToken *GratitudeTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.Approve(&_GratitudeToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_GratitudeToken *GratitudeTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_GratitudeToken *GratitudeTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.Burn(&_GratitudeToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_GratitudeToken *GratitudeTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.Burn(&_GratitudeToken.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_GratitudeToken *GratitudeTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_GratitudeToken *GratitudeTokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.BurnFrom(&_GratitudeToken.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_GratitudeToken *GratitudeTokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.BurnFrom(&_GratitudeToken.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_GratitudeToken *GratitudeTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_GratitudeToken *GratitudeTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.DecreaseAllowance(&_GratitudeToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_GratitudeToken *GratitudeTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.DecreaseAllowance(&_GratitudeToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_GratitudeToken *GratitudeTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_GratitudeToken *GratitudeTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.IncreaseAllowance(&_GratitudeToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_GratitudeToken *GratitudeTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.IncreaseAllowance(&_GratitudeToken.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GratitudeToken *GratitudeTokenTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GratitudeToken *GratitudeTokenSession) Initialize() (*types.Transaction, error) {
	return _GratitudeToken.Contract.Initialize(&_GratitudeToken.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GratitudeToken *GratitudeTokenTransactorSession) Initialize() (*types.Transaction, error) {
	return _GratitudeToken.Contract.Initialize(&_GratitudeToken.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_GratitudeToken *GratitudeTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_GratitudeToken *GratitudeTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.Mint(&_GratitudeToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_GratitudeToken *GratitudeTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.Mint(&_GratitudeToken.TransactOpts, to, amount)
}

// MintToMany is a paid mutator transaction binding the contract method 0x203abe34.
//
// Solidity: function mintToMany(address[] recipients) returns()
func (_GratitudeToken *GratitudeTokenTransactor) MintToMany(opts *bind.TransactOpts, recipients []common.Address) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "mintToMany", recipients)
}

// MintToMany is a paid mutator transaction binding the contract method 0x203abe34.
//
// Solidity: function mintToMany(address[] recipients) returns()
func (_GratitudeToken *GratitudeTokenSession) MintToMany(recipients []common.Address) (*types.Transaction, error) {
	return _GratitudeToken.Contract.MintToMany(&_GratitudeToken.TransactOpts, recipients)
}

// MintToMany is a paid mutator transaction binding the contract method 0x203abe34.
//
// Solidity: function mintToMany(address[] recipients) returns()
func (_GratitudeToken *GratitudeTokenTransactorSession) MintToMany(recipients []common.Address) (*types.Transaction, error) {
	return _GratitudeToken.Contract.MintToMany(&_GratitudeToken.TransactOpts, recipients)
}

// MintToMany0 is a paid mutator transaction binding the contract method 0x8a0b0360.
//
// Solidity: function mintToMany(address[] recipients, uint256 amount) returns()
func (_GratitudeToken *GratitudeTokenTransactor) MintToMany0(opts *bind.TransactOpts, recipients []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "mintToMany0", recipients, amount)
}

// MintToMany0 is a paid mutator transaction binding the contract method 0x8a0b0360.
//
// Solidity: function mintToMany(address[] recipients, uint256 amount) returns()
func (_GratitudeToken *GratitudeTokenSession) MintToMany0(recipients []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.MintToMany0(&_GratitudeToken.TransactOpts, recipients, amount)
}

// MintToMany0 is a paid mutator transaction binding the contract method 0x8a0b0360.
//
// Solidity: function mintToMany(address[] recipients, uint256 amount) returns()
func (_GratitudeToken *GratitudeTokenTransactorSession) MintToMany0(recipients []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.MintToMany0(&_GratitudeToken.TransactOpts, recipients, amount)
}

// MintToMany1 is a paid mutator transaction binding the contract method 0x92524725.
//
// Solidity: function mintToMany(address[] recipients, uint256[] amounts) returns()
func (_GratitudeToken *GratitudeTokenTransactor) MintToMany1(opts *bind.TransactOpts, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "mintToMany1", recipients, amounts)
}

// MintToMany1 is a paid mutator transaction binding the contract method 0x92524725.
//
// Solidity: function mintToMany(address[] recipients, uint256[] amounts) returns()
func (_GratitudeToken *GratitudeTokenSession) MintToMany1(recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.MintToMany1(&_GratitudeToken.TransactOpts, recipients, amounts)
}

// MintToMany1 is a paid mutator transaction binding the contract method 0x92524725.
//
// Solidity: function mintToMany(address[] recipients, uint256[] amounts) returns()
func (_GratitudeToken *GratitudeTokenTransactorSession) MintToMany1(recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.MintToMany1(&_GratitudeToken.TransactOpts, recipients, amounts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GratitudeToken *GratitudeTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GratitudeToken *GratitudeTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _GratitudeToken.Contract.RenounceOwnership(&_GratitudeToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GratitudeToken *GratitudeTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GratitudeToken.Contract.RenounceOwnership(&_GratitudeToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_GratitudeToken *GratitudeTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_GratitudeToken *GratitudeTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.Transfer(&_GratitudeToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_GratitudeToken *GratitudeTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.Transfer(&_GratitudeToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_GratitudeToken *GratitudeTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_GratitudeToken *GratitudeTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.TransferFrom(&_GratitudeToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_GratitudeToken *GratitudeTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GratitudeToken.Contract.TransferFrom(&_GratitudeToken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GratitudeToken *GratitudeTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GratitudeToken *GratitudeTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GratitudeToken.Contract.TransferOwnership(&_GratitudeToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GratitudeToken *GratitudeTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GratitudeToken.Contract.TransferOwnership(&_GratitudeToken.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GratitudeToken *GratitudeTokenTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GratitudeToken *GratitudeTokenSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GratitudeToken.Contract.UpgradeTo(&_GratitudeToken.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GratitudeToken *GratitudeTokenTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GratitudeToken.Contract.UpgradeTo(&_GratitudeToken.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GratitudeToken *GratitudeTokenTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GratitudeToken.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GratitudeToken *GratitudeTokenSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GratitudeToken.Contract.UpgradeToAndCall(&_GratitudeToken.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GratitudeToken *GratitudeTokenTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GratitudeToken.Contract.UpgradeToAndCall(&_GratitudeToken.TransactOpts, newImplementation, data)
}

// GratitudeTokenAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the GratitudeToken contract.
type GratitudeTokenAdminChangedIterator struct {
	Event *GratitudeTokenAdminChanged // Event containing the contract specifics and raw log

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
func (it *GratitudeTokenAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GratitudeTokenAdminChanged)
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
		it.Event = new(GratitudeTokenAdminChanged)
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
func (it *GratitudeTokenAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GratitudeTokenAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GratitudeTokenAdminChanged represents a AdminChanged event raised by the GratitudeToken contract.
type GratitudeTokenAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GratitudeToken *GratitudeTokenFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*GratitudeTokenAdminChangedIterator, error) {

	logs, sub, err := _GratitudeToken.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &GratitudeTokenAdminChangedIterator{contract: _GratitudeToken.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GratitudeToken *GratitudeTokenFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *GratitudeTokenAdminChanged) (event.Subscription, error) {

	logs, sub, err := _GratitudeToken.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GratitudeTokenAdminChanged)
				if err := _GratitudeToken.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_GratitudeToken *GratitudeTokenFilterer) ParseAdminChanged(log types.Log) (*GratitudeTokenAdminChanged, error) {
	event := new(GratitudeTokenAdminChanged)
	if err := _GratitudeToken.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GratitudeTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GratitudeToken contract.
type GratitudeTokenApprovalIterator struct {
	Event *GratitudeTokenApproval // Event containing the contract specifics and raw log

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
func (it *GratitudeTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GratitudeTokenApproval)
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
		it.Event = new(GratitudeTokenApproval)
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
func (it *GratitudeTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GratitudeTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GratitudeTokenApproval represents a Approval event raised by the GratitudeToken contract.
type GratitudeTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GratitudeToken *GratitudeTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GratitudeTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GratitudeToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GratitudeTokenApprovalIterator{contract: _GratitudeToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GratitudeToken *GratitudeTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GratitudeTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GratitudeToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GratitudeTokenApproval)
				if err := _GratitudeToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GratitudeToken *GratitudeTokenFilterer) ParseApproval(log types.Log) (*GratitudeTokenApproval, error) {
	event := new(GratitudeTokenApproval)
	if err := _GratitudeToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GratitudeTokenBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the GratitudeToken contract.
type GratitudeTokenBeaconUpgradedIterator struct {
	Event *GratitudeTokenBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *GratitudeTokenBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GratitudeTokenBeaconUpgraded)
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
		it.Event = new(GratitudeTokenBeaconUpgraded)
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
func (it *GratitudeTokenBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GratitudeTokenBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GratitudeTokenBeaconUpgraded represents a BeaconUpgraded event raised by the GratitudeToken contract.
type GratitudeTokenBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GratitudeToken *GratitudeTokenFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*GratitudeTokenBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GratitudeToken.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &GratitudeTokenBeaconUpgradedIterator{contract: _GratitudeToken.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GratitudeToken *GratitudeTokenFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *GratitudeTokenBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GratitudeToken.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GratitudeTokenBeaconUpgraded)
				if err := _GratitudeToken.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_GratitudeToken *GratitudeTokenFilterer) ParseBeaconUpgraded(log types.Log) (*GratitudeTokenBeaconUpgraded, error) {
	event := new(GratitudeTokenBeaconUpgraded)
	if err := _GratitudeToken.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GratitudeTokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GratitudeToken contract.
type GratitudeTokenInitializedIterator struct {
	Event *GratitudeTokenInitialized // Event containing the contract specifics and raw log

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
func (it *GratitudeTokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GratitudeTokenInitialized)
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
		it.Event = new(GratitudeTokenInitialized)
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
func (it *GratitudeTokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GratitudeTokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GratitudeTokenInitialized represents a Initialized event raised by the GratitudeToken contract.
type GratitudeTokenInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GratitudeToken *GratitudeTokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*GratitudeTokenInitializedIterator, error) {

	logs, sub, err := _GratitudeToken.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GratitudeTokenInitializedIterator{contract: _GratitudeToken.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GratitudeToken *GratitudeTokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GratitudeTokenInitialized) (event.Subscription, error) {

	logs, sub, err := _GratitudeToken.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GratitudeTokenInitialized)
				if err := _GratitudeToken.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GratitudeToken *GratitudeTokenFilterer) ParseInitialized(log types.Log) (*GratitudeTokenInitialized, error) {
	event := new(GratitudeTokenInitialized)
	if err := _GratitudeToken.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GratitudeTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GratitudeToken contract.
type GratitudeTokenOwnershipTransferredIterator struct {
	Event *GratitudeTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GratitudeTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GratitudeTokenOwnershipTransferred)
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
		it.Event = new(GratitudeTokenOwnershipTransferred)
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
func (it *GratitudeTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GratitudeTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GratitudeTokenOwnershipTransferred represents a OwnershipTransferred event raised by the GratitudeToken contract.
type GratitudeTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GratitudeToken *GratitudeTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GratitudeTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GratitudeToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GratitudeTokenOwnershipTransferredIterator{contract: _GratitudeToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GratitudeToken *GratitudeTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GratitudeTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GratitudeToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GratitudeTokenOwnershipTransferred)
				if err := _GratitudeToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GratitudeToken *GratitudeTokenFilterer) ParseOwnershipTransferred(log types.Log) (*GratitudeTokenOwnershipTransferred, error) {
	event := new(GratitudeTokenOwnershipTransferred)
	if err := _GratitudeToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GratitudeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GratitudeToken contract.
type GratitudeTokenTransferIterator struct {
	Event *GratitudeTokenTransfer // Event containing the contract specifics and raw log

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
func (it *GratitudeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GratitudeTokenTransfer)
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
		it.Event = new(GratitudeTokenTransfer)
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
func (it *GratitudeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GratitudeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GratitudeTokenTransfer represents a Transfer event raised by the GratitudeToken contract.
type GratitudeTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GratitudeToken *GratitudeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GratitudeTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GratitudeToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GratitudeTokenTransferIterator{contract: _GratitudeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GratitudeToken *GratitudeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GratitudeTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GratitudeToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GratitudeTokenTransfer)
				if err := _GratitudeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GratitudeToken *GratitudeTokenFilterer) ParseTransfer(log types.Log) (*GratitudeTokenTransfer, error) {
	event := new(GratitudeTokenTransfer)
	if err := _GratitudeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GratitudeTokenUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GratitudeToken contract.
type GratitudeTokenUpgradedIterator struct {
	Event *GratitudeTokenUpgraded // Event containing the contract specifics and raw log

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
func (it *GratitudeTokenUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GratitudeTokenUpgraded)
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
		it.Event = new(GratitudeTokenUpgraded)
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
func (it *GratitudeTokenUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GratitudeTokenUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GratitudeTokenUpgraded represents a Upgraded event raised by the GratitudeToken contract.
type GratitudeTokenUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GratitudeToken *GratitudeTokenFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GratitudeTokenUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GratitudeToken.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GratitudeTokenUpgradedIterator{contract: _GratitudeToken.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GratitudeToken *GratitudeTokenFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GratitudeTokenUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GratitudeToken.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GratitudeTokenUpgraded)
				if err := _GratitudeToken.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_GratitudeToken *GratitudeTokenFilterer) ParseUpgraded(log types.Log) (*GratitudeTokenUpgraded, error) {
	event := new(GratitudeTokenUpgraded)
	if err := _GratitudeToken.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
