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
	Bin: "0x60a06040523060805234801561001457600080fd5b5060805161176b61004c6000396000818161023b01528181610284015281816103230152818161036301526103f6015261176b6000f3fe60806040526004361061009c5760003560e01c8063715018a611610064578063715018a61461016a5780638da5cb5b1461017f578063c4d66de8146101b1578063f2fde38b146101d1578063fae0d7dc146101f1578063fc0c546a1461021157600080fd5b80633659cfe6146100a15780633d8b9e07146100c35780634f1ef2861461011457806352d1902d14610127578063536f3fff1461014a575b600080fd5b3480156100ad57600080fd5b506100c16100bc36600461126a565b610231565b005b3480156100cf57600080fd5b506100f86100de366004611287565b60ca6020526000908152604090205465ffffffffffff1681565b60405165ffffffffffff90911681526020015b60405180910390f35b6100c16101223660046112b6565b610319565b34801561013357600080fd5b5061013c6103e9565b60405190815260200161010b565b34801561015657600080fd5b506100c1610165366004611395565b61049c565b34801561017657600080fd5b506100c161084c565b34801561018b57600080fd5b506033546001600160a01b03165b6040516001600160a01b03909116815260200161010b565b3480156101bd57600080fd5b506100c16101cc36600461126a565b610860565b3480156101dd57600080fd5b506100c16101ec36600461126a565b610994565b3480156101fd57600080fd5b5061013c61020c36600461144f565b610a0a565b34801561021d57600080fd5b5060c954610199906001600160a01b031681565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036102825760405162461bcd60e51b8152600401610279906114a6565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166102cb6000805160206116cf833981519152546001600160a01b031690565b6001600160a01b0316146102f15760405162461bcd60e51b8152600401610279906114f2565b6102fa81610a88565b6040805160008082526020820190925261031691839190610a90565b50565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036103615760405162461bcd60e51b8152600401610279906114a6565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166103aa6000805160206116cf833981519152546001600160a01b031690565b6001600160a01b0316146103d05760405162461bcd60e51b8152600401610279906114f2565b6103d982610a88565b6103e582826001610a90565b5050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104895760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610279565b506000805160206116cf83398151915290565b4265ffffffffffff808616908216108015906104c757508565ffffffffffff168165ffffffffffff16105b61051e5760405162461bcd60e51b815260206004820152602260248201527f4552433230494f553a2065787069726564206f72206e6f742076616c69642079604482015261195d60f21b6064820152608401610279565b600061052d8989898989610a0a565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c810191909152605c0160408051601f198184030181529181528151602092830120600081815260ca90935291205490915065ffffffffffff16156105de5760405162461bcd60e51b815260206004820152601a60248201527f4552433230494f553a20616c72656164792072656465656d65640000000000006044820152606401610279565b883b156106bd57604051630b135d3f60e11b815289906001600160a01b03821690631626ba7e906106179085908990899060040161153e565b602060405180830381865afa158015610634573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106589190611574565b6001600160e01b031916631626ba7e60e01b146106b75760405162461bcd60e51b815260206004820152601b60248201527f4552433230494f553a20696e76616c6964207369676e617475726500000000006044820152606401610279565b5061075d565b886001600160a01b03166107078286868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610c0092505050565b6001600160a01b03161461075d5760405162461bcd60e51b815260206004820152601b60248201527f4552433230494f553a20696e76616c6964207369676e617475726500000000006044820152606401610279565b60c9546040516323b872dd60e01b81526001600160a01b038b81166004830152336024830152604482018b9052909116906323b872dd906064016020604051808303816000875af11580156107b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107da919061159e565b50600081815260ca6020908152604091829020805465ffffffffffff191665ffffffffffff8616179055905189815233916001600160a01b038c16917fd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9910160405180910390a3505050505050505050565b610854610e67565b61085e6000610ec1565b565b600054610100900460ff16158080156108805750600054600160ff909116105b8061089a5750303b15801561089a575060005460ff166001145b6108fd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610279565b6000805460ff191660011790558015610920576000805461ff0019166101001790555b610928610f13565b610930610f42565b60c980546001600160a01b0319166001600160a01b03841617905580156103e5576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b61099c610e67565b6001600160a01b038116610a015760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610279565b61031681610ec1565b6040516bffffffffffffffffffffffff19606087811b82166020840152603483018790526001600160d01b031960d087811b8216605486015286901b16605a84015280830184905246608084015230901b1660a082015260009060b40160405160208183030381529060405280519060200120905095945050505050565b610316610e67565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610ac857610ac383610f69565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610b22575060408051601f3d908101601f19168201909252610b1f918101906115c0565b60015b610b855760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610279565b6000805160206116cf8339815191528114610bf45760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610279565b50610ac3838383611005565b60008151604114610c675760405162461bcd60e51b815260206004820152603a602482015260008051602061171683398151915260448201527f3a20696e76616c6964207369676e6174757265206c656e6774680000000000006064820152608401610279565b600082604081518110610c7c57610c7c6115d9565b016020015160f81c90506000610c928482611030565b90506000610ca1856020611030565b90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0811115610d275760405162461bcd60e51b815260206004820152603d602482015260008051602061171683398151915260448201527f3a20696e76616c6964207369676e6174757265202773272076616c75650000006064820152608401610279565b8260ff16601b14158015610d3f57508260ff16601c14155b15610da05760405162461bcd60e51b815260206004820152603d602482015260008051602061171683398151915260448201527f3a20696e76616c6964207369676e6174757265202776272076616c75650000006064820152608401610279565b60408051600081526020810180835288905260ff851691810191909152606081018390526080810182905260019060a0016020604051602081039080840390855afa158015610df3573d6000803e3d6000fd5b5050604051601f1901519450506001600160a01b038416610e5d5760405162461bcd60e51b8152602060048201526030602482015260008051602061171683398151915260448201526f1d1024a72b20a624a22fa9a4a3a722a960811b6064820152608401610279565b5050505b92915050565b6033546001600160a01b0316331461085e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610279565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16610f3a5760405162461bcd60e51b8152600401610279906115ef565b61085e611096565b600054610100900460ff1661085e5760405162461bcd60e51b8152600401610279906115ef565b6001600160a01b0381163b610fd65760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610279565b6000805160206116cf83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b61100e836110c6565b60008251118061101b5750805b15610ac35761102a8383611106565b50505050565b600061103d82602061163a565b8351101561108d5760405162461bcd60e51b815260206004820181905260248201527f72656164427974657333323a20696e76616c69642064617461206c656e6774686044820152606401610279565b50016020015190565b600054610100900460ff166110bd5760405162461bcd60e51b8152600401610279906115ef565b61085e33610ec1565b6110cf81610f69565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b606061112b83836040518060600160405280602781526020016116ef60279139611132565b9392505050565b6060600080856001600160a01b03168560405161114f919061167f565b600060405180830381855af49150503d806000811461118a576040519150601f19603f3d011682016040523d82523d6000602084013e61118f565b606091505b50915091506111a0868383876111aa565b9695505050505050565b60608315611219578251600003611212576001600160a01b0385163b6112125760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610279565b5081611223565b611223838361122b565b949350505050565b81511561123b5781518083602001fd5b8060405162461bcd60e51b8152600401610279919061169b565b6001600160a01b038116811461031657600080fd5b60006020828403121561127c57600080fd5b813561112b81611255565b60006020828403121561129957600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156112c957600080fd5b82356112d481611255565b9150602083013567ffffffffffffffff808211156112f157600080fd5b818501915085601f83011261130557600080fd5b813581811115611317576113176112a0565b604051601f8201601f19908116603f0116810190838211818310171561133f5761133f6112a0565b8160405282815288602084870101111561135857600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b803565ffffffffffff8116811461139057600080fd5b919050565b600080600080600080600060c0888a0312156113b057600080fd5b87356113bb81611255565b9650602088013595506113d06040890161137a565b94506113de6060890161137a565b93506080880135925060a088013567ffffffffffffffff8082111561140257600080fd5b818a0191508a601f83011261141657600080fd5b81358181111561142557600080fd5b8b602082850101111561143757600080fd5b60208301945080935050505092959891949750929550565b600080600080600060a0868803121561146757600080fd5b853561147281611255565b9450602086013593506114876040870161137a565b92506114956060870161137a565b949793965091946080013592915050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b83815260406020820152816040820152818360608301376000818301606090810191909152601f909201601f1916010192915050565b60006020828403121561158657600080fd5b81516001600160e01b03198116811461112b57600080fd5b6000602082840312156115b057600080fd5b8151801515811461112b57600080fd5b6000602082840312156115d257600080fd5b5051919050565b634e487b7160e01b600052603260045260246000fd5b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b80820180821115610e6157634e487b7160e01b600052601160045260246000fd5b60005b8381101561167657818101518382015260200161165e565b50506000910152565b6000825161169181846020870161165b565b9190910192915050565b60208152600082518060208401526116ba81604085016020870161165b565b601f01601f1916919091016040019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645369676e617475726556616c696461746f72237265636f7665725369676e6572a26469706673582212202729ddc80b66869a87af20e296b0f74502697c8a3d05aa4b47cee159d33d2e1464736f6c63430008140033",
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
