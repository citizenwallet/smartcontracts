// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simpleaccount

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

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

// SimpleaccountMetaData contains all meta data concerning the Simpleaccount contract.
var SimpleaccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SimpleAccountInitialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"func\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dest\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"func\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516114cd3803806114cd83398101604081905261002f91610109565b6001600160a01b03811660805261004461004a565b50610139565b600054610100900460ff16156100b65760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811614610107576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60006020828403121561011b57600080fd5b81516001600160a01b038116811461013257600080fd5b9392505050565b6080516113486101856000396000818161025b015281816104bf0152818161054001528181610636015281816107e10152818161081b015281816109200152610b1301526113486000f3fe6080604052600436106100eb5760003560e01c80638da5cb5b1161008a578063c399ec8811610059578063c399ec88146102ce578063c4d66de8146102e3578063d087d28814610303578063f23a6e611461031857600080fd5b80638da5cb5b1461020e578063b0d691fe1461024c578063b61d27f61461027f578063bc197c811461029f57600080fd5b806318dfb3c7116100c657806318dfb3c7146101985780633a871cdd146101b85780634a58db19146101e65780634d44560d146101ee57600080fd5b806223de29146100f757806301ffc9a71461011e578063150b7a021461015357600080fd5b366100f257005b600080fd5b34801561010357600080fd5b5061011c610112366004610e2a565b5050505050505050565b005b34801561012a57600080fd5b5061013e610139366004610edb565b610345565b60405190151581526020015b60405180910390f35b34801561015f57600080fd5b5061017f61016e366004610f05565b630a85bd0160e11b95945050505050565b6040516001600160e01b0319909116815260200161014a565b3480156101a457600080fd5b5061011c6101b3366004610fbd565b610397565b3480156101c457600080fd5b506101d86101d3366004611029565b610497565b60405190815260200161014a565b61011c6104bd565b3480156101fa57600080fd5b5061011c61020936600461107d565b610536565b34801561021a57600080fd5b50600054610234906201000090046001600160a01b031681565b6040516001600160a01b03909116815260200161014a565b34801561025857600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610234565b34801561028b57600080fd5b5061011c61029a3660046110a9565b6105c7565b3480156102ab57600080fd5b5061017f6102ba3660046110f9565b63bc197c8160e01b98975050505050505050565b3480156102da57600080fd5b506101d8610616565b3480156102ef57600080fd5b5061011c6102fe366004611197565b6106a7565b34801561030f57600080fd5b506101d86107ba565b34801561032457600080fd5b5061017f6103333660046111b4565b63f23a6e6160e01b9695505050505050565b60006001600160e01b03198216630a85bd0160e11b148061037657506001600160e01b03198216630271189760e51b145b8061039157506001600160e01b031982166301ffc9a760e01b145b92915050565b61039f610810565b8281146103e95760405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b60448201526064015b60405180910390fd5b60005b838110156104905761047e85858381811061040957610409611230565b905060200201602081019061041e9190611197565b600085858581811061043257610432611230565b90506020028101906104449190611246565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506108a592505050565b806104888161128d565b9150506103ec565b5050505050565b60006104a1610915565b6104ab848461098d565b90506104b682610a3c565b9392505050565b7f000000000000000000000000000000000000000000000000000000000000000060405163b760faf960e01b81523060048201526001600160a01b03919091169063b760faf99034906024016000604051808303818588803b15801561052257600080fd5b505af1158015610490573d6000803e3d6000fd5b61053e610a89565b7f000000000000000000000000000000000000000000000000000000000000000060405163040b850f60e31b81526001600160a01b03848116600483015260248201849052919091169063205c287890604401600060405180830381600087803b1580156105ab57600080fd5b505af11580156105bf573d6000803e3d6000fd5b505050505050565b6105cf610810565b610610848484848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506108a592505050565b50505050565b6040516370a0823160e01b81523060048201526000906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a08231906024015b602060405180830381865afa15801561067e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a291906112b4565b905090565b600054610100900460ff16158080156106c75750600054600160ff909116105b806106e15750303b1580156106e1575060005460ff166001145b6107445760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103e0565b6000805460ff191660011790558015610767576000805461ff0019166101001790555b61077082610ae0565b80156107b6576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b604051631aab3f0d60e11b8152306004820152600060248201819052906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906335567e1a90604401610661565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061085757506000546201000090046001600160a01b031633145b6108a35760405162461bcd60e51b815260206004820181905260248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e7460448201526064016103e0565b565b600080846001600160a01b031684846040516108c191906112cd565b60006040518083038185875af1925050503d80600081146108fe576040519150601f19603f3d011682016040523d82523d6000602084013e610903565b606091505b50915091508161049057805160208201fd5b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108a35760405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e740000000060448201526064016103e0565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c829052603c8120610a0a6109cd610140860186611246565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508593925050610b5c9050565b6000546201000090046001600160a01b03908116911614610a2f576001915050610391565b5060009392505050565b50565b8015610a3957604051600090339060001990849084818181858888f193505050503d8060008114610490576040519150601f19603f3d011682016040523d82523d6000602084013e610490565b6000546201000090046001600160a01b0316331480610aa757503330145b6108a35760405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b60448201526064016103e0565b6000805462010000600160b01b031916620100006001600160a01b038481168202929092178084556040519190048216927f0000000000000000000000000000000000000000000000000000000000000000909216917f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de91a350565b6000806000610b6b8585610b80565b91509150610b7881610bc5565b509392505050565b6000808251604103610bb65760208301516040840151606085015160001a610baa87828585610d0f565b94509450505050610bbe565b506000905060025b9250929050565b6000816004811115610bd957610bd96112fc565b03610be15750565b6001816004811115610bf557610bf56112fc565b03610c425760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016103e0565b6002816004811115610c5657610c566112fc565b03610ca35760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016103e0565b6003816004811115610cb757610cb76112fc565b03610a395760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016103e0565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610d465750600090506003610dca565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610d9a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610dc357600060019250925050610dca565b9150600090505b94509492505050565b6001600160a01b0381168114610a3957600080fd5b60008083601f840112610dfa57600080fd5b50813567ffffffffffffffff811115610e1257600080fd5b602083019150836020828501011115610bbe57600080fd5b60008060008060008060008060c0898b031215610e4657600080fd5b8835610e5181610dd3565b97506020890135610e6181610dd3565b96506040890135610e7181610dd3565b955060608901359450608089013567ffffffffffffffff80821115610e9557600080fd5b610ea18c838d01610de8565b909650945060a08b0135915080821115610eba57600080fd5b50610ec78b828c01610de8565b999c989b5096995094979396929594505050565b600060208284031215610eed57600080fd5b81356001600160e01b0319811681146104b657600080fd5b600080600080600060808688031215610f1d57600080fd5b8535610f2881610dd3565b94506020860135610f3881610dd3565b935060408601359250606086013567ffffffffffffffff811115610f5b57600080fd5b610f6788828901610de8565b969995985093965092949392505050565b60008083601f840112610f8a57600080fd5b50813567ffffffffffffffff811115610fa257600080fd5b6020830191508360208260051b8501011115610bbe57600080fd5b60008060008060408587031215610fd357600080fd5b843567ffffffffffffffff80821115610feb57600080fd5b610ff788838901610f78565b9096509450602087013591508082111561101057600080fd5b5061101d87828801610f78565b95989497509550505050565b60008060006060848603121561103e57600080fd5b833567ffffffffffffffff81111561105557600080fd5b8401610160818703121561106857600080fd5b95602085013595506040909401359392505050565b6000806040838503121561109057600080fd5b823561109b81610dd3565b946020939093013593505050565b600080600080606085870312156110bf57600080fd5b84356110ca81610dd3565b935060208501359250604085013567ffffffffffffffff8111156110ed57600080fd5b61101d87828801610de8565b60008060008060008060008060a0898b03121561111557600080fd5b883561112081610dd3565b9750602089013561113081610dd3565b9650604089013567ffffffffffffffff8082111561114d57600080fd5b6111598c838d01610f78565b909850965060608b013591508082111561117257600080fd5b61117e8c838d01610f78565b909650945060808b0135915080821115610eba57600080fd5b6000602082840312156111a957600080fd5b81356104b681610dd3565b60008060008060008060a087890312156111cd57600080fd5b86356111d881610dd3565b955060208701356111e881610dd3565b94506040870135935060608701359250608087013567ffffffffffffffff81111561121257600080fd5b61121e89828a01610de8565b979a9699509497509295939492505050565b634e487b7160e01b600052603260045260246000fd5b6000808335601e1984360301811261125d57600080fd5b83018035915067ffffffffffffffff82111561127857600080fd5b602001915036819003821315610bbe57600080fd5b6000600182016112ad57634e487b7160e01b600052601160045260246000fd5b5060010190565b6000602082840312156112c657600080fd5b5051919050565b6000825160005b818110156112ee57602081860181015185830152016112d4565b506000920191825250919050565b634e487b7160e01b600052602160045260246000fdfea26469706673582212204f30953d99988a405bf94ae9ebbf486f331d43c7b9f959cd185f40b054ea411664736f6c63430008140033",
}

// SimpleaccountABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleaccountMetaData.ABI instead.
var SimpleaccountABI = SimpleaccountMetaData.ABI

// SimpleaccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleaccountMetaData.Bin instead.
var SimpleaccountBin = SimpleaccountMetaData.Bin

// DeploySimpleaccount deploys a new Ethereum contract, binding an instance of Simpleaccount to it.
func DeploySimpleaccount(auth *bind.TransactOpts, backend bind.ContractBackend, anEntryPoint common.Address) (common.Address, *types.Transaction, *Simpleaccount, error) {
	parsed, err := SimpleaccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleaccountBin), backend, anEntryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Simpleaccount{SimpleaccountCaller: SimpleaccountCaller{contract: contract}, SimpleaccountTransactor: SimpleaccountTransactor{contract: contract}, SimpleaccountFilterer: SimpleaccountFilterer{contract: contract}}, nil
}

// Simpleaccount is an auto generated Go binding around an Ethereum contract.
type Simpleaccount struct {
	SimpleaccountCaller     // Read-only binding to the contract
	SimpleaccountTransactor // Write-only binding to the contract
	SimpleaccountFilterer   // Log filterer for contract events
}

// SimpleaccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleaccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleaccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleaccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleaccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleaccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleaccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleaccountSession struct {
	Contract     *Simpleaccount    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleaccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleaccountCallerSession struct {
	Contract *SimpleaccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleaccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleaccountTransactorSession struct {
	Contract     *SimpleaccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleaccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleaccountRaw struct {
	Contract *Simpleaccount // Generic contract binding to access the raw methods on
}

// SimpleaccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleaccountCallerRaw struct {
	Contract *SimpleaccountCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleaccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleaccountTransactorRaw struct {
	Contract *SimpleaccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleaccount creates a new instance of Simpleaccount, bound to a specific deployed contract.
func NewSimpleaccount(address common.Address, backend bind.ContractBackend) (*Simpleaccount, error) {
	contract, err := bindSimpleaccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simpleaccount{SimpleaccountCaller: SimpleaccountCaller{contract: contract}, SimpleaccountTransactor: SimpleaccountTransactor{contract: contract}, SimpleaccountFilterer: SimpleaccountFilterer{contract: contract}}, nil
}

// NewSimpleaccountCaller creates a new read-only instance of Simpleaccount, bound to a specific deployed contract.
func NewSimpleaccountCaller(address common.Address, caller bind.ContractCaller) (*SimpleaccountCaller, error) {
	contract, err := bindSimpleaccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleaccountCaller{contract: contract}, nil
}

// NewSimpleaccountTransactor creates a new write-only instance of Simpleaccount, bound to a specific deployed contract.
func NewSimpleaccountTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleaccountTransactor, error) {
	contract, err := bindSimpleaccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleaccountTransactor{contract: contract}, nil
}

// NewSimpleaccountFilterer creates a new log filterer instance of Simpleaccount, bound to a specific deployed contract.
func NewSimpleaccountFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleaccountFilterer, error) {
	contract, err := bindSimpleaccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleaccountFilterer{contract: contract}, nil
}

// bindSimpleaccount binds a generic wrapper to an already deployed contract.
func bindSimpleaccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleaccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simpleaccount *SimpleaccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simpleaccount.Contract.SimpleaccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simpleaccount *SimpleaccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simpleaccount.Contract.SimpleaccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simpleaccount *SimpleaccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simpleaccount.Contract.SimpleaccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simpleaccount *SimpleaccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simpleaccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simpleaccount *SimpleaccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simpleaccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simpleaccount *SimpleaccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simpleaccount.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Simpleaccount *SimpleaccountCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Simpleaccount.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Simpleaccount *SimpleaccountSession) EntryPoint() (common.Address, error) {
	return _Simpleaccount.Contract.EntryPoint(&_Simpleaccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Simpleaccount *SimpleaccountCallerSession) EntryPoint() (common.Address, error) {
	return _Simpleaccount.Contract.EntryPoint(&_Simpleaccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Simpleaccount *SimpleaccountCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Simpleaccount.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Simpleaccount *SimpleaccountSession) GetDeposit() (*big.Int, error) {
	return _Simpleaccount.Contract.GetDeposit(&_Simpleaccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Simpleaccount *SimpleaccountCallerSession) GetDeposit() (*big.Int, error) {
	return _Simpleaccount.Contract.GetDeposit(&_Simpleaccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Simpleaccount *SimpleaccountCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Simpleaccount.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Simpleaccount *SimpleaccountSession) GetNonce() (*big.Int, error) {
	return _Simpleaccount.Contract.GetNonce(&_Simpleaccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Simpleaccount *SimpleaccountCallerSession) GetNonce() (*big.Int, error) {
	return _Simpleaccount.Contract.GetNonce(&_Simpleaccount.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Simpleaccount *SimpleaccountCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Simpleaccount.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Simpleaccount *SimpleaccountSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Simpleaccount.Contract.OnERC1155BatchReceived(&_Simpleaccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Simpleaccount *SimpleaccountCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Simpleaccount.Contract.OnERC1155BatchReceived(&_Simpleaccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Simpleaccount *SimpleaccountCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Simpleaccount.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Simpleaccount *SimpleaccountSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Simpleaccount.Contract.OnERC1155Received(&_Simpleaccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Simpleaccount *SimpleaccountCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Simpleaccount.Contract.OnERC1155Received(&_Simpleaccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Simpleaccount *SimpleaccountCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Simpleaccount.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Simpleaccount *SimpleaccountSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Simpleaccount.Contract.OnERC721Received(&_Simpleaccount.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Simpleaccount *SimpleaccountCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Simpleaccount.Contract.OnERC721Received(&_Simpleaccount.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Simpleaccount *SimpleaccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Simpleaccount.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Simpleaccount *SimpleaccountSession) Owner() (common.Address, error) {
	return _Simpleaccount.Contract.Owner(&_Simpleaccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Simpleaccount *SimpleaccountCallerSession) Owner() (common.Address, error) {
	return _Simpleaccount.Contract.Owner(&_Simpleaccount.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Simpleaccount *SimpleaccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Simpleaccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Simpleaccount *SimpleaccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Simpleaccount.Contract.SupportsInterface(&_Simpleaccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Simpleaccount *SimpleaccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Simpleaccount.Contract.SupportsInterface(&_Simpleaccount.CallOpts, interfaceId)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Simpleaccount *SimpleaccountCaller) TokensReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	var out []interface{}
	err := _Simpleaccount.contract.Call(opts, &out, "tokensReceived", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Simpleaccount *SimpleaccountSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _Simpleaccount.Contract.TokensReceived(&_Simpleaccount.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Simpleaccount *SimpleaccountCallerSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _Simpleaccount.Contract.TokensReceived(&_Simpleaccount.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Simpleaccount *SimpleaccountTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simpleaccount.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Simpleaccount *SimpleaccountSession) AddDeposit() (*types.Transaction, error) {
	return _Simpleaccount.Contract.AddDeposit(&_Simpleaccount.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Simpleaccount *SimpleaccountTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _Simpleaccount.Contract.AddDeposit(&_Simpleaccount.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Simpleaccount *SimpleaccountTransactor) Execute(opts *bind.TransactOpts, dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Simpleaccount.contract.Transact(opts, "execute", dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Simpleaccount *SimpleaccountSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Simpleaccount.Contract.Execute(&_Simpleaccount.TransactOpts, dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Simpleaccount *SimpleaccountTransactorSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Simpleaccount.Contract.Execute(&_Simpleaccount.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x18dfb3c7.
//
// Solidity: function executeBatch(address[] dest, bytes[] func) returns()
func (_Simpleaccount *SimpleaccountTransactor) ExecuteBatch(opts *bind.TransactOpts, dest []common.Address, arg1 [][]byte) (*types.Transaction, error) {
	return _Simpleaccount.contract.Transact(opts, "executeBatch", dest, arg1)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x18dfb3c7.
//
// Solidity: function executeBatch(address[] dest, bytes[] func) returns()
func (_Simpleaccount *SimpleaccountSession) ExecuteBatch(dest []common.Address, arg1 [][]byte) (*types.Transaction, error) {
	return _Simpleaccount.Contract.ExecuteBatch(&_Simpleaccount.TransactOpts, dest, arg1)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x18dfb3c7.
//
// Solidity: function executeBatch(address[] dest, bytes[] func) returns()
func (_Simpleaccount *SimpleaccountTransactorSession) ExecuteBatch(dest []common.Address, arg1 [][]byte) (*types.Transaction, error) {
	return _Simpleaccount.Contract.ExecuteBatch(&_Simpleaccount.TransactOpts, dest, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_Simpleaccount *SimpleaccountTransactor) Initialize(opts *bind.TransactOpts, anOwner common.Address) (*types.Transaction, error) {
	return _Simpleaccount.contract.Transact(opts, "initialize", anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_Simpleaccount *SimpleaccountSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _Simpleaccount.Contract.Initialize(&_Simpleaccount.TransactOpts, anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_Simpleaccount *SimpleaccountTransactorSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _Simpleaccount.Contract.Initialize(&_Simpleaccount.TransactOpts, anOwner)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Simpleaccount *SimpleaccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Simpleaccount.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Simpleaccount *SimpleaccountSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Simpleaccount.Contract.ValidateUserOp(&_Simpleaccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Simpleaccount *SimpleaccountTransactorSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Simpleaccount.Contract.ValidateUserOp(&_Simpleaccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Simpleaccount *SimpleaccountTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Simpleaccount.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Simpleaccount *SimpleaccountSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Simpleaccount.Contract.WithdrawDepositTo(&_Simpleaccount.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Simpleaccount *SimpleaccountTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Simpleaccount.Contract.WithdrawDepositTo(&_Simpleaccount.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Simpleaccount *SimpleaccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simpleaccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Simpleaccount *SimpleaccountSession) Receive() (*types.Transaction, error) {
	return _Simpleaccount.Contract.Receive(&_Simpleaccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Simpleaccount *SimpleaccountTransactorSession) Receive() (*types.Transaction, error) {
	return _Simpleaccount.Contract.Receive(&_Simpleaccount.TransactOpts)
}

// SimpleaccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Simpleaccount contract.
type SimpleaccountInitializedIterator struct {
	Event *SimpleaccountInitialized // Event containing the contract specifics and raw log

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
func (it *SimpleaccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleaccountInitialized)
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
		it.Event = new(SimpleaccountInitialized)
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
func (it *SimpleaccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleaccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleaccountInitialized represents a Initialized event raised by the Simpleaccount contract.
type SimpleaccountInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Simpleaccount *SimpleaccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*SimpleaccountInitializedIterator, error) {

	logs, sub, err := _Simpleaccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SimpleaccountInitializedIterator{contract: _Simpleaccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Simpleaccount *SimpleaccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SimpleaccountInitialized) (event.Subscription, error) {

	logs, sub, err := _Simpleaccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleaccountInitialized)
				if err := _Simpleaccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Simpleaccount *SimpleaccountFilterer) ParseInitialized(log types.Log) (*SimpleaccountInitialized, error) {
	event := new(SimpleaccountInitialized)
	if err := _Simpleaccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleaccountSimpleAccountInitializedIterator is returned from FilterSimpleAccountInitialized and is used to iterate over the raw logs and unpacked data for SimpleAccountInitialized events raised by the Simpleaccount contract.
type SimpleaccountSimpleAccountInitializedIterator struct {
	Event *SimpleaccountSimpleAccountInitialized // Event containing the contract specifics and raw log

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
func (it *SimpleaccountSimpleAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleaccountSimpleAccountInitialized)
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
		it.Event = new(SimpleaccountSimpleAccountInitialized)
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
func (it *SimpleaccountSimpleAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleaccountSimpleAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleaccountSimpleAccountInitialized represents a SimpleAccountInitialized event raised by the Simpleaccount contract.
type SimpleaccountSimpleAccountInitialized struct {
	EntryPoint common.Address
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSimpleAccountInitialized is a free log retrieval operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_Simpleaccount *SimpleaccountFilterer) FilterSimpleAccountInitialized(opts *bind.FilterOpts, entryPoint []common.Address, owner []common.Address) (*SimpleaccountSimpleAccountInitializedIterator, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Simpleaccount.contract.FilterLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SimpleaccountSimpleAccountInitializedIterator{contract: _Simpleaccount.contract, event: "SimpleAccountInitialized", logs: logs, sub: sub}, nil
}

// WatchSimpleAccountInitialized is a free log subscription operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_Simpleaccount *SimpleaccountFilterer) WatchSimpleAccountInitialized(opts *bind.WatchOpts, sink chan<- *SimpleaccountSimpleAccountInitialized, entryPoint []common.Address, owner []common.Address) (event.Subscription, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Simpleaccount.contract.WatchLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleaccountSimpleAccountInitialized)
				if err := _Simpleaccount.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
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

// ParseSimpleAccountInitialized is a log parse operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_Simpleaccount *SimpleaccountFilterer) ParseSimpleAccountInitialized(log types.Log) (*SimpleaccountSimpleAccountInitialized, error) {
	event := new(SimpleaccountSimpleAccountInitialized)
	if err := _Simpleaccount.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
