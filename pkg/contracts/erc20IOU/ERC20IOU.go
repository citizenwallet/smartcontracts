// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20IOU

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

// Erc20IOUMetaData contains all meta data concerning the Erc20IOU contract.
var Erc20IOUMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"redeemHash\",\"type\":\"bytes32\"}],\"name\":\"redeemed\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"time\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a0806040523461003157306080526115d2908161003782396080518181816101090152818161038001526104950152f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c80633659cfe6146100c75780633d8b9e07146100c25780634f1ef286146100bd57806352d1902d146100b8578063536f3fff146100b3578063715018a6146100ae5780638da5cb5b146100a9578063c4d66de8146100a4578063f2fde38b1461009f578063fae0d7dc1461009a5763fc0c546a1461009557600080fd5b61083c565b6107f4565b610763565b610676565b61064d565b6105ec565b610571565b610482565b61032d565b610231565b6100e2565b6001600160a01b038116036100dd57565b600080fd5b346100dd5760203660031901126100dd576004356100ff816100cc565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116919061013830841415611136565b61015560008051602061155d833981519152938285541614611197565b61015d611095565b6040519061016a8261027b565b600082527f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff16156101a65750506101a491506112b4565b005b6020600491604094939451928380926352d1902d60e01b825286165afa60009181610201575b506101ee5760405162461bcd60e51b8152806101ea60048201611265565b0390fd5b6101a4936101fc9114611207565b611344565b61022391925060203d811161022a575b61021b81836102b8565b8101906111f8565b90386101cc565b503d610211565b346100dd5760203660031901126100dd5760043560005260ca602052602065ffffffffffff60406000205416604051908152f35b634e487b7160e01b600052604160045260246000fd5b6020810190811067ffffffffffffffff82111761029757604052565b610265565b6060810190811067ffffffffffffffff82111761029757604052565b90601f8019910116810190811067ffffffffffffffff82111761029757604052565b67ffffffffffffffff811161029757601f01601f191660200190565b929192610302826102da565b9161031060405193846102b8565b8294818452818301116100dd578281602093846000960137010152565b60403660031901126100dd57600435610345816100cc565b60243567ffffffffffffffff81116100dd57366023820112156100dd576103769036906024816004013591016102f6565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169291906103b030851415611136565b6103cd60008051602061155d833981519152948286541614611197565b6103d5611095565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff161561040b5750506101a491506112b4565b6020600491604094939451928380926352d1902d60e01b825286165afa60009181610462575b5061044f5760405162461bcd60e51b8152806101ea60048201611265565b6101a49361045d9114611207565b611429565b61047b91925060203d811161022a5761021b81836102b8565b9038610431565b346100dd5760003660031901126100dd577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036104dc5760405160008051602061155d8339815191528152602090f35b60405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608490fd5b6044359065ffffffffffff821682036100dd57565b6064359065ffffffffffff821682036100dd57565b346100dd5760c03660031901126100dd5760043561058e816100cc565b610596610547565b61059e61055c565b9060a4359267ffffffffffffffff928385116100dd57366023860112156100dd5784600401359384116100dd5736602485870101116100dd5760246101a49501926084359260243590610afc565b346100dd5760008060031936011261064a57610606611095565b603380546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b80fd5b346100dd5760003660031901126100dd576033546040516001600160a01b039091168152602090f35b346100dd5760203660031901126100dd57600435610693816100cc565b6106d7600054916106bb60ff8460081c161580948195610755575b8115610735575b50610865565b826106ce600160ff196000541617600055565b61071c576108c8565b6106dd57005b6106ed61ff001960005416600055565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a1005b61073061010061ff00196000541617600055565b6108c8565b303b15915081610747575b50386106b5565b6001915060ff161438610740565b600160ff82161091506106ae565b346100dd5760203660031901126100dd57600435610780816100cc565b610788611095565b6001600160a01b038116156107a0576101a4906110ed565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b346100dd5760a03660031901126100dd576020610834600435610816816100cc565b61081e610547565b61082661055c565b90608435926024359061091e565b604051908152f35b346100dd5760003660031901126100dd5760c9546040516001600160a01b039091168152602090f35b1561086c57565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b6108e260ff60005460081c166108dd81611035565b611035565b6108eb336110ed565b6108fc60ff60005460081c16611035565b60018060a01b03166bffffffffffffffffffffffff60a01b60c954161760c955565b939190926040519360208501956bffffffffffffffffffffffff199060601b168652603485015265ffffffffffff60d01b809260d01b16605485015260d01b16605a83015260608201524660808201523060601b60a08201526094815260c0810181811067ffffffffffffffff8211176102975760405251902090565b156109a257565b60405162461bcd60e51b815260206004820152602260248201527f4552433230494f553a2065787069726564206f72206e6f742076616c69642079604482015261195d60f21b6064820152608490fd5b156109f957565b60405162461bcd60e51b815260206004820152601a60248201527f4552433230494f553a20616c72656164792072656465656d65640000000000006044820152606490fd5b15610a4557565b60405162461bcd60e51b815260206004820152601b60248201527f4552433230494f553a20696e76616c6964207369676e617475726500000000006044820152606490fd5b908160209103126100dd57516001600160e01b0319811681036100dd5790565b91926060938192845260406020850152816040850152848401376000828201840152601f01601f1916010190565b6040513d6000823e3d90fd5b908160209103126100dd575180151581036100dd5790565b959094919293610b2e9065ffffffffffff958642169587821687101580610d59575b610b279061099b565b888a61091e565b91610bb060405193602094610b8281610b748882019485603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c8201520190565b03601f1981018352826102b8565b51902095610ba9610b9d8860005260ca602052604060002090565b5465ffffffffffff1690565b16156109f2565b82873b15610d2857610bd99260405180948192630b135d3f60e11b958684528a60048501610aaa565b03816001600160a01b038b165afa8015610cf657610c0a92600091610cfb575b506001600160e01b03191614610a3e565b60c9546001600160a01b03166040516323b872dd60e01b81526001600160a01b03871660048201523360248201526044810186905293908290859060649082906000905af1918215610cf657610c8e94610c7593610cc8575b505060005260ca602052604060002090565b9065ffffffffffff1665ffffffffffff19825416179055565b60405190815233916001600160a01b0316907fd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d990602090a3565b81610ce792903d10610cef575b610cdf81836102b8565b810190610ae4565b503880610c63565b503d610cd5565b610ad8565b610d1b9150843d8611610d21575b610d1381836102b8565b810190610a8a565b38610bf9565b503d610d09565b50610d3b610d4191610d549336916102f6565b85610e3d565b6001600160a01b03878116911614610a3e565b610c0a565b508088168710610b1e565b15610d6b57565b60405162461bcd60e51b815260206004820152603a602482015260008051602061157d83398151915260448201527f3a20696e76616c6964207369676e6174757265206c656e6774680000000000006064820152608490fd5b805160401015610dd45760600190565b634e487b7160e01b600052603260045260246000fd5b15610df157565b60405162461bcd60e51b8152602060048201526030602482015260008051602061157d83398151915260448201526f1d1024a72b20a624a22fa9a4a3a722a960811b6064820152608490fd5b610e4a6041835114610d64565b610e6d610e67610e5984610dc4565b516001600160f81b03191690565b60f81c90565b610e7f610e7984610fd1565b93611025565b917f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311610f785760ff8216601b8114159081610f6c575b50610f1357610eeb600093602095604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa15610cf657600051610f106001600160a01b0382161515610dea565b90565b60405162461bcd60e51b815260206004820152603d602482015260008051602061157d83398151915260448201527f3a20696e76616c6964207369676e6174757265202776272076616c75650000006064820152608490fd5b601c9150141538610eb7565b60405162461bcd60e51b815260206004820152603d602482015260008051602061157d83398151915260448201527f3a20696e76616c6964207369676e6174757265202773272076616c75650000006064820152608490fd5b6020815110610fe1576020015190565b606460405162461bcd60e51b815260206004820152602060248201527f72656164427974657333323a20696e76616c69642064617461206c656e6774686044820152fd5b6040815110610fe1576040015190565b1561103c57565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b6033546001600160a01b031633036110a957565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b603380546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b1561113d57565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608490fd5b1561119e57565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b6064820152608490fd5b908160209103126100dd575190565b1561120e57565b60405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608490fd5b60809060208152602e60208201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960408201526d6f6e206973206e6f74205555505360901b60608201520190565b803b156112e95760008051602061155d83398151915280546001600160a01b0319166001600160a01b03909216919091179055565b60405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608490fd5b9061134e826112b4565b60006001600160a01b0383167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8280a2815115801590611422575b61139257505050565b611416928180604051946113a58661029c565b602786527f416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c6020870152660819985a5b195960ca1b6040870152602081519101845af4903d15611419573d6113f9816102da565b9061140760405192836102b8565b8152809360203d92013e6114ca565b50565b606092506114ca565b5080611389565b90611433826112b4565b60006001600160a01b0383167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8280a28151158015906114765761139257505050565b506001611389565b1561148557565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b919290156114ea57508151156114de575090565b610f10903b151561147e565b8251909150156114fd5750805190602001fd5b6040519062461bcd60e51b82528160208060048301528251908160248401526000935b828510611543575050604492506000838284010152601f80199101168101030190fd5b848101820151868601604401529381019385935061152056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5369676e617475726556616c696461746f72237265636f7665725369676e6572a2646970667358221220373cbd18f9b8519da0713e8da0430a609d0d4975f86d73da72f4f47ac054bc8964736f6c63430008140033",
}

// Erc20IOUABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20IOUMetaData.ABI instead.
var Erc20IOUABI = Erc20IOUMetaData.ABI

// Erc20IOUBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20IOUMetaData.Bin instead.
var Erc20IOUBin = Erc20IOUMetaData.Bin

// DeployErc20IOU deploys a new Ethereum contract, binding an instance of Erc20IOU to it.
func DeployErc20IOU(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Erc20IOU, error) {
	parsed, err := Erc20IOUMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20IOUBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20IOU{Erc20IOUCaller: Erc20IOUCaller{contract: contract}, Erc20IOUTransactor: Erc20IOUTransactor{contract: contract}, Erc20IOUFilterer: Erc20IOUFilterer{contract: contract}}, nil
}

// Erc20IOU is an auto generated Go binding around an Ethereum contract.
type Erc20IOU struct {
	Erc20IOUCaller     // Read-only binding to the contract
	Erc20IOUTransactor // Write-only binding to the contract
	Erc20IOUFilterer   // Log filterer for contract events
}

// Erc20IOUCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20IOUCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20IOUTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20IOUTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20IOUFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20IOUFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20IOUSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20IOUSession struct {
	Contract     *Erc20IOU         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20IOUCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20IOUCallerSession struct {
	Contract *Erc20IOUCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Erc20IOUTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20IOUTransactorSession struct {
	Contract     *Erc20IOUTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Erc20IOURaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20IOURaw struct {
	Contract *Erc20IOU // Generic contract binding to access the raw methods on
}

// Erc20IOUCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20IOUCallerRaw struct {
	Contract *Erc20IOUCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20IOUTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20IOUTransactorRaw struct {
	Contract *Erc20IOUTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20IOU creates a new instance of Erc20IOU, bound to a specific deployed contract.
func NewErc20IOU(address common.Address, backend bind.ContractBackend) (*Erc20IOU, error) {
	contract, err := bindErc20IOU(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20IOU{Erc20IOUCaller: Erc20IOUCaller{contract: contract}, Erc20IOUTransactor: Erc20IOUTransactor{contract: contract}, Erc20IOUFilterer: Erc20IOUFilterer{contract: contract}}, nil
}

// NewErc20IOUCaller creates a new read-only instance of Erc20IOU, bound to a specific deployed contract.
func NewErc20IOUCaller(address common.Address, caller bind.ContractCaller) (*Erc20IOUCaller, error) {
	contract, err := bindErc20IOU(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20IOUCaller{contract: contract}, nil
}

// NewErc20IOUTransactor creates a new write-only instance of Erc20IOU, bound to a specific deployed contract.
func NewErc20IOUTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20IOUTransactor, error) {
	contract, err := bindErc20IOU(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20IOUTransactor{contract: contract}, nil
}

// NewErc20IOUFilterer creates a new log filterer instance of Erc20IOU, bound to a specific deployed contract.
func NewErc20IOUFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20IOUFilterer, error) {
	contract, err := bindErc20IOU(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20IOUFilterer{contract: contract}, nil
}

// bindErc20IOU binds a generic wrapper to an already deployed contract.
func bindErc20IOU(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20IOUMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20IOU *Erc20IOURaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20IOU.Contract.Erc20IOUCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20IOU *Erc20IOURaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20IOU.Contract.Erc20IOUTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20IOU *Erc20IOURaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20IOU.Contract.Erc20IOUTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20IOU *Erc20IOUCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20IOU.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20IOU *Erc20IOUTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20IOU.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20IOU *Erc20IOUTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20IOU.Contract.contract.Transact(opts, method, params...)
}

// GetHash is a free data retrieval call binding the contract method 0xfae0d7dc.
//
// Solidity: function getHash(address from, uint256 amount, uint48 validUntil, uint48 validAfter, uint256 sequence) view returns(bytes32)
func (_Erc20IOU *Erc20IOUCaller) GetHash(opts *bind.CallOpts, from common.Address, amount *big.Int, validUntil *big.Int, validAfter *big.Int, sequence *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Erc20IOU.contract.Call(opts, &out, "getHash", from, amount, validUntil, validAfter, sequence)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0xfae0d7dc.
//
// Solidity: function getHash(address from, uint256 amount, uint48 validUntil, uint48 validAfter, uint256 sequence) view returns(bytes32)
func (_Erc20IOU *Erc20IOUSession) GetHash(from common.Address, amount *big.Int, validUntil *big.Int, validAfter *big.Int, sequence *big.Int) ([32]byte, error) {
	return _Erc20IOU.Contract.GetHash(&_Erc20IOU.CallOpts, from, amount, validUntil, validAfter, sequence)
}

// GetHash is a free data retrieval call binding the contract method 0xfae0d7dc.
//
// Solidity: function getHash(address from, uint256 amount, uint48 validUntil, uint48 validAfter, uint256 sequence) view returns(bytes32)
func (_Erc20IOU *Erc20IOUCallerSession) GetHash(from common.Address, amount *big.Int, validUntil *big.Int, validAfter *big.Int, sequence *big.Int) ([32]byte, error) {
	return _Erc20IOU.Contract.GetHash(&_Erc20IOU.CallOpts, from, amount, validUntil, validAfter, sequence)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20IOU *Erc20IOUCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20IOU.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20IOU *Erc20IOUSession) Owner() (common.Address, error) {
	return _Erc20IOU.Contract.Owner(&_Erc20IOU.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20IOU *Erc20IOUCallerSession) Owner() (common.Address, error) {
	return _Erc20IOU.Contract.Owner(&_Erc20IOU.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Erc20IOU *Erc20IOUCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc20IOU.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Erc20IOU *Erc20IOUSession) ProxiableUUID() ([32]byte, error) {
	return _Erc20IOU.Contract.ProxiableUUID(&_Erc20IOU.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Erc20IOU *Erc20IOUCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Erc20IOU.Contract.ProxiableUUID(&_Erc20IOU.CallOpts)
}

// Redeemed is a free data retrieval call binding the contract method 0x3d8b9e07.
//
// Solidity: function redeemed(bytes32 redeemHash) view returns(uint48 time)
func (_Erc20IOU *Erc20IOUCaller) Redeemed(opts *bind.CallOpts, redeemHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Erc20IOU.contract.Call(opts, &out, "redeemed", redeemHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Redeemed is a free data retrieval call binding the contract method 0x3d8b9e07.
//
// Solidity: function redeemed(bytes32 redeemHash) view returns(uint48 time)
func (_Erc20IOU *Erc20IOUSession) Redeemed(redeemHash [32]byte) (*big.Int, error) {
	return _Erc20IOU.Contract.Redeemed(&_Erc20IOU.CallOpts, redeemHash)
}

// Redeemed is a free data retrieval call binding the contract method 0x3d8b9e07.
//
// Solidity: function redeemed(bytes32 redeemHash) view returns(uint48 time)
func (_Erc20IOU *Erc20IOUCallerSession) Redeemed(redeemHash [32]byte) (*big.Int, error) {
	return _Erc20IOU.Contract.Redeemed(&_Erc20IOU.CallOpts, redeemHash)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Erc20IOU *Erc20IOUCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20IOU.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Erc20IOU *Erc20IOUSession) Token() (common.Address, error) {
	return _Erc20IOU.Contract.Token(&_Erc20IOU.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Erc20IOU *Erc20IOUCallerSession) Token() (common.Address, error) {
	return _Erc20IOU.Contract.Token(&_Erc20IOU.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _token) returns()
func (_Erc20IOU *Erc20IOUTransactor) Initialize(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Erc20IOU.contract.Transact(opts, "initialize", _token)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _token) returns()
func (_Erc20IOU *Erc20IOUSession) Initialize(_token common.Address) (*types.Transaction, error) {
	return _Erc20IOU.Contract.Initialize(&_Erc20IOU.TransactOpts, _token)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _token) returns()
func (_Erc20IOU *Erc20IOUTransactorSession) Initialize(_token common.Address) (*types.Transaction, error) {
	return _Erc20IOU.Contract.Initialize(&_Erc20IOU.TransactOpts, _token)
}

// Redeem is a paid mutator transaction binding the contract method 0x536f3fff.
//
// Solidity: function redeem(address from, uint256 amount, uint48 validUntil, uint48 validAfter, uint256 sequence, bytes signature) returns()
func (_Erc20IOU *Erc20IOUTransactor) Redeem(opts *bind.TransactOpts, from common.Address, amount *big.Int, validUntil *big.Int, validAfter *big.Int, sequence *big.Int, signature []byte) (*types.Transaction, error) {
	return _Erc20IOU.contract.Transact(opts, "redeem", from, amount, validUntil, validAfter, sequence, signature)
}

// Redeem is a paid mutator transaction binding the contract method 0x536f3fff.
//
// Solidity: function redeem(address from, uint256 amount, uint48 validUntil, uint48 validAfter, uint256 sequence, bytes signature) returns()
func (_Erc20IOU *Erc20IOUSession) Redeem(from common.Address, amount *big.Int, validUntil *big.Int, validAfter *big.Int, sequence *big.Int, signature []byte) (*types.Transaction, error) {
	return _Erc20IOU.Contract.Redeem(&_Erc20IOU.TransactOpts, from, amount, validUntil, validAfter, sequence, signature)
}

// Redeem is a paid mutator transaction binding the contract method 0x536f3fff.
//
// Solidity: function redeem(address from, uint256 amount, uint48 validUntil, uint48 validAfter, uint256 sequence, bytes signature) returns()
func (_Erc20IOU *Erc20IOUTransactorSession) Redeem(from common.Address, amount *big.Int, validUntil *big.Int, validAfter *big.Int, sequence *big.Int, signature []byte) (*types.Transaction, error) {
	return _Erc20IOU.Contract.Redeem(&_Erc20IOU.TransactOpts, from, amount, validUntil, validAfter, sequence, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc20IOU *Erc20IOUTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20IOU.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc20IOU *Erc20IOUSession) RenounceOwnership() (*types.Transaction, error) {
	return _Erc20IOU.Contract.RenounceOwnership(&_Erc20IOU.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc20IOU *Erc20IOUTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Erc20IOU.Contract.RenounceOwnership(&_Erc20IOU.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20IOU *Erc20IOUTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Erc20IOU.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20IOU *Erc20IOUSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc20IOU.Contract.TransferOwnership(&_Erc20IOU.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20IOU *Erc20IOUTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc20IOU.Contract.TransferOwnership(&_Erc20IOU.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Erc20IOU *Erc20IOUTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Erc20IOU.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Erc20IOU *Erc20IOUSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Erc20IOU.Contract.UpgradeTo(&_Erc20IOU.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Erc20IOU *Erc20IOUTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Erc20IOU.Contract.UpgradeTo(&_Erc20IOU.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Erc20IOU *Erc20IOUTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Erc20IOU.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Erc20IOU *Erc20IOUSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Erc20IOU.Contract.UpgradeToAndCall(&_Erc20IOU.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Erc20IOU *Erc20IOUTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Erc20IOU.Contract.UpgradeToAndCall(&_Erc20IOU.TransactOpts, newImplementation, data)
}

// Erc20IOUAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Erc20IOU contract.
type Erc20IOUAdminChangedIterator struct {
	Event *Erc20IOUAdminChanged // Event containing the contract specifics and raw log

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
func (it *Erc20IOUAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20IOUAdminChanged)
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
		it.Event = new(Erc20IOUAdminChanged)
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
func (it *Erc20IOUAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20IOUAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20IOUAdminChanged represents a AdminChanged event raised by the Erc20IOU contract.
type Erc20IOUAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Erc20IOU *Erc20IOUFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*Erc20IOUAdminChangedIterator, error) {

	logs, sub, err := _Erc20IOU.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &Erc20IOUAdminChangedIterator{contract: _Erc20IOU.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Erc20IOU *Erc20IOUFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *Erc20IOUAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Erc20IOU.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20IOUAdminChanged)
				if err := _Erc20IOU.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Erc20IOU *Erc20IOUFilterer) ParseAdminChanged(log types.Log) (*Erc20IOUAdminChanged, error) {
	event := new(Erc20IOUAdminChanged)
	if err := _Erc20IOU.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20IOUBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Erc20IOU contract.
type Erc20IOUBeaconUpgradedIterator struct {
	Event *Erc20IOUBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *Erc20IOUBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20IOUBeaconUpgraded)
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
		it.Event = new(Erc20IOUBeaconUpgraded)
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
func (it *Erc20IOUBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20IOUBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20IOUBeaconUpgraded represents a BeaconUpgraded event raised by the Erc20IOU contract.
type Erc20IOUBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Erc20IOU *Erc20IOUFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*Erc20IOUBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Erc20IOU.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &Erc20IOUBeaconUpgradedIterator{contract: _Erc20IOU.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Erc20IOU *Erc20IOUFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *Erc20IOUBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Erc20IOU.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20IOUBeaconUpgraded)
				if err := _Erc20IOU.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Erc20IOU *Erc20IOUFilterer) ParseBeaconUpgraded(log types.Log) (*Erc20IOUBeaconUpgraded, error) {
	event := new(Erc20IOUBeaconUpgraded)
	if err := _Erc20IOU.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20IOUInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Erc20IOU contract.
type Erc20IOUInitializedIterator struct {
	Event *Erc20IOUInitialized // Event containing the contract specifics and raw log

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
func (it *Erc20IOUInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20IOUInitialized)
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
		it.Event = new(Erc20IOUInitialized)
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
func (it *Erc20IOUInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20IOUInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20IOUInitialized represents a Initialized event raised by the Erc20IOU contract.
type Erc20IOUInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Erc20IOU *Erc20IOUFilterer) FilterInitialized(opts *bind.FilterOpts) (*Erc20IOUInitializedIterator, error) {

	logs, sub, err := _Erc20IOU.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Erc20IOUInitializedIterator{contract: _Erc20IOU.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Erc20IOU *Erc20IOUFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Erc20IOUInitialized) (event.Subscription, error) {

	logs, sub, err := _Erc20IOU.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20IOUInitialized)
				if err := _Erc20IOU.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Erc20IOU *Erc20IOUFilterer) ParseInitialized(log types.Log) (*Erc20IOUInitialized, error) {
	event := new(Erc20IOUInitialized)
	if err := _Erc20IOU.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20IOUOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Erc20IOU contract.
type Erc20IOUOwnershipTransferredIterator struct {
	Event *Erc20IOUOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Erc20IOUOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20IOUOwnershipTransferred)
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
		it.Event = new(Erc20IOUOwnershipTransferred)
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
func (it *Erc20IOUOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20IOUOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20IOUOwnershipTransferred represents a OwnershipTransferred event raised by the Erc20IOU contract.
type Erc20IOUOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc20IOU *Erc20IOUFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Erc20IOUOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc20IOU.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Erc20IOUOwnershipTransferredIterator{contract: _Erc20IOU.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc20IOU *Erc20IOUFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Erc20IOUOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc20IOU.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20IOUOwnershipTransferred)
				if err := _Erc20IOU.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Erc20IOU *Erc20IOUFilterer) ParseOwnershipTransferred(log types.Log) (*Erc20IOUOwnershipTransferred, error) {
	event := new(Erc20IOUOwnershipTransferred)
	if err := _Erc20IOU.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20IOURedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Erc20IOU contract.
type Erc20IOURedeemIterator struct {
	Event *Erc20IOURedeem // Event containing the contract specifics and raw log

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
func (it *Erc20IOURedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20IOURedeem)
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
		it.Event = new(Erc20IOURedeem)
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
func (it *Erc20IOURedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20IOURedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20IOURedeem represents a Redeem event raised by the Erc20IOU contract.
type Erc20IOURedeem struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed from, address indexed to, uint256 amount)
func (_Erc20IOU *Erc20IOUFilterer) FilterRedeem(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20IOURedeemIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20IOU.contract.FilterLogs(opts, "Redeem", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20IOURedeemIterator{contract: _Erc20IOU.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed from, address indexed to, uint256 amount)
func (_Erc20IOU *Erc20IOUFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *Erc20IOURedeem, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20IOU.contract.WatchLogs(opts, "Redeem", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20IOURedeem)
				if err := _Erc20IOU.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed from, address indexed to, uint256 amount)
func (_Erc20IOU *Erc20IOUFilterer) ParseRedeem(log types.Log) (*Erc20IOURedeem, error) {
	event := new(Erc20IOURedeem)
	if err := _Erc20IOU.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20IOUUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Erc20IOU contract.
type Erc20IOUUpgradedIterator struct {
	Event *Erc20IOUUpgraded // Event containing the contract specifics and raw log

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
func (it *Erc20IOUUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20IOUUpgraded)
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
		it.Event = new(Erc20IOUUpgraded)
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
func (it *Erc20IOUUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20IOUUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20IOUUpgraded represents a Upgraded event raised by the Erc20IOU contract.
type Erc20IOUUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Erc20IOU *Erc20IOUFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*Erc20IOUUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Erc20IOU.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &Erc20IOUUpgradedIterator{contract: _Erc20IOU.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Erc20IOU *Erc20IOUFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *Erc20IOUUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Erc20IOU.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20IOUUpgraded)
				if err := _Erc20IOU.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Erc20IOU *Erc20IOUFilterer) ParseUpgraded(log types.Log) (*Erc20IOUUpgraded, error) {
	event := new(Erc20IOUUpgraded)
	if err := _Erc20IOU.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
