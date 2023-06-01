// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accfactory

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

// AccfactoryMetaData contains all meta data concerning the Accfactory contract.
var AccfactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accountImplementation\",\"outputs\":[{\"internalType\":\"contractAccount\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"contractAccount\",\"name\":\"ret\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051611e5a380380611e5a83398101604081905261002f91610088565b8060405161003c9061007b565b6001600160a01b039091168152602001604051809103906000f080158015610068573d6000803e3d6000fd5b506001600160a01b0316608052506100b8565b6114cd8061098d83390190565b60006020828403121561009a57600080fd5b81516001600160a01b03811681146100b157600080fd5b9392505050565b6080516108ad6100e060003960008181604b0152818161012101526101ee01526108ad6000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806311464fbe146100465780635fbfb9cf146100895780638cb84e181461009c575b600080fd5b61006d7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b61006d6100973660046102c9565b6100af565b61006d6100aa3660046102c9565b6101af565b6000806100bc84846101af565b90506001600160a01b0381163b80156100d7575090506101a9565b6040516001600160a01b038616907f805996f252884581e2f74cf3d2b03564d5ec26ccc90850ae12653dc1b72d1fa290600090a26040516001600160a01b038616602482015284907f00000000000000000000000000000000000000000000000000000000000000009060440160408051601f198184030181529181526020820180516001600160e01b031663189acdbd60e31b17905251610178906102bc565b610183929190610325565b8190604051809103906000f59050801580156101a3573d6000803e3d6000fd5b50925050505b92915050565b60006102838260001b604051806020016101c8906102bc565b601f1982820381018352601f9091011660408190526001600160a01b03871660248201527f00000000000000000000000000000000000000000000000000000000000000009060440160408051601f19818403018152918152602080830180516001600160e01b031663189acdbd60e31b179052905161024a93929101610325565b60408051601f19818403018152908290526102689291602001610367565b6040516020818303038152906040528051906020012061028a565b9392505050565b60006102838383306000604051836040820152846020820152828152600b8101905060ff815360559020949350505050565b6104e18061039783390190565b600080604083850312156102dc57600080fd5b82356001600160a01b03811681146102f357600080fd5b946020939093013593505050565b60005b8381101561031c578181015183820152602001610304565b50506000910152565b60018060a01b03831681526040602082015260008251806040840152610352816060850160208701610301565b601f01601f1916919091016060019392505050565b60008351610379818460208801610301565b83519083019061038d818360208801610301565b0194935050505056fe60806040526040516104e13803806104e1833981016040819052610022916102de565b61002e82826000610035565b50506103fb565b61003e83610061565b60008251118061004b5750805b1561005c5761005a83836100a1565b505b505050565b61006a816100cd565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606100c683836040518060600160405280602781526020016104ba60279139610180565b9392505050565b6001600160a01b0381163b61013f5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080856001600160a01b03168560405161019d91906103ac565b600060405180830381855af49150503d80600081146101d8576040519150601f19603f3d011682016040523d82523d6000602084013e6101dd565b606091505b5090925090506101ef868383876101f9565b9695505050505050565b60608315610268578251600003610261576001600160a01b0385163b6102615760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610136565b5081610272565b610272838361027a565b949350505050565b81511561028a5781518083602001fd5b8060405162461bcd60e51b815260040161013691906103c8565b634e487b7160e01b600052604160045260246000fd5b60005b838110156102d55781810151838201526020016102bd565b50506000910152565b600080604083850312156102f157600080fd5b82516001600160a01b038116811461030857600080fd5b60208401519092506001600160401b038082111561032557600080fd5b818501915085601f83011261033957600080fd5b81518181111561034b5761034b6102a4565b604051601f8201601f19908116603f01168101908382118183101715610373576103736102a4565b8160405282815288602084870101111561038c57600080fd5b61039d8360208301602088016102ba565b80955050505050509250929050565b600082516103be8184602087016102ba565b9190910192915050565b60208152600082518060208401526103e78160408501602087016102ba565b601f01601f19169190910160400192915050565b60b1806104096000396000f3fe608060405236601057600e6013565b005b600e5b601f601b6021565b6058565b565b600060537f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b3660008037600080366000845af43d6000803e8080156076573d6000f35b3d6000fdfea264697066735822122004eafc12009ddd1bac09a0fbcdbbc0e38fdd09caf7f585983e4d324fd8f99a7364736f6c63430008140033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212202f427ba3ee1c0981d50865747918aa3383fda90413558d39907aed9a1c9fc2d864736f6c6343000814003360a060405234801561001057600080fd5b506040516114cd3803806114cd83398101604081905261002f91610109565b6001600160a01b03811660805261004461004a565b50610139565b600054610100900460ff16156100b65760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811614610107576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60006020828403121561011b57600080fd5b81516001600160a01b038116811461013257600080fd5b9392505050565b6080516113486101856000396000818161025b015281816104bf0152818161054001528181610636015281816107e10152818161081b015281816109200152610b1301526113486000f3fe6080604052600436106100eb5760003560e01c80638da5cb5b1161008a578063c399ec8811610059578063c399ec88146102ce578063c4d66de8146102e3578063d087d28814610303578063f23a6e611461031857600080fd5b80638da5cb5b1461020e578063b0d691fe1461024c578063b61d27f61461027f578063bc197c811461029f57600080fd5b806318dfb3c7116100c657806318dfb3c7146101985780633a871cdd146101b85780634a58db19146101e65780634d44560d146101ee57600080fd5b806223de29146100f757806301ffc9a71461011e578063150b7a021461015357600080fd5b366100f257005b600080fd5b34801561010357600080fd5b5061011c610112366004610e2a565b5050505050505050565b005b34801561012a57600080fd5b5061013e610139366004610edb565b610345565b60405190151581526020015b60405180910390f35b34801561015f57600080fd5b5061017f61016e366004610f05565b630a85bd0160e11b95945050505050565b6040516001600160e01b0319909116815260200161014a565b3480156101a457600080fd5b5061011c6101b3366004610fbd565b610397565b3480156101c457600080fd5b506101d86101d3366004611029565b610497565b60405190815260200161014a565b61011c6104bd565b3480156101fa57600080fd5b5061011c61020936600461107d565b610536565b34801561021a57600080fd5b50600054610234906201000090046001600160a01b031681565b6040516001600160a01b03909116815260200161014a565b34801561025857600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610234565b34801561028b57600080fd5b5061011c61029a3660046110a9565b6105c7565b3480156102ab57600080fd5b5061017f6102ba3660046110f9565b63bc197c8160e01b98975050505050505050565b3480156102da57600080fd5b506101d8610616565b3480156102ef57600080fd5b5061011c6102fe366004611197565b6106a7565b34801561030f57600080fd5b506101d86107ba565b34801561032457600080fd5b5061017f6103333660046111b4565b63f23a6e6160e01b9695505050505050565b60006001600160e01b03198216630a85bd0160e11b148061037657506001600160e01b03198216630271189760e51b145b8061039157506001600160e01b031982166301ffc9a760e01b145b92915050565b61039f610810565b8281146103e95760405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b60448201526064015b60405180910390fd5b60005b838110156104905761047e85858381811061040957610409611230565b905060200201602081019061041e9190611197565b600085858581811061043257610432611230565b90506020028101906104449190611246565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506108a592505050565b806104888161128d565b9150506103ec565b5050505050565b60006104a1610915565b6104ab848461098d565b90506104b682610a3c565b9392505050565b7f000000000000000000000000000000000000000000000000000000000000000060405163b760faf960e01b81523060048201526001600160a01b03919091169063b760faf99034906024016000604051808303818588803b15801561052257600080fd5b505af1158015610490573d6000803e3d6000fd5b61053e610a89565b7f000000000000000000000000000000000000000000000000000000000000000060405163040b850f60e31b81526001600160a01b03848116600483015260248201849052919091169063205c287890604401600060405180830381600087803b1580156105ab57600080fd5b505af11580156105bf573d6000803e3d6000fd5b505050505050565b6105cf610810565b610610848484848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506108a592505050565b50505050565b6040516370a0823160e01b81523060048201526000906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a08231906024015b602060405180830381865afa15801561067e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a291906112b4565b905090565b600054610100900460ff16158080156106c75750600054600160ff909116105b806106e15750303b1580156106e1575060005460ff166001145b6107445760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103e0565b6000805460ff191660011790558015610767576000805461ff0019166101001790555b61077082610ae0565b80156107b6576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b604051631aab3f0d60e11b8152306004820152600060248201819052906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906335567e1a90604401610661565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061085757506000546201000090046001600160a01b031633145b6108a35760405162461bcd60e51b815260206004820181905260248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e7460448201526064016103e0565b565b600080846001600160a01b031684846040516108c191906112cd565b60006040518083038185875af1925050503d80600081146108fe576040519150601f19603f3d011682016040523d82523d6000602084013e610903565b606091505b50915091508161049057805160208201fd5b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108a35760405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e740000000060448201526064016103e0565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c829052603c8120610a0a6109cd610140860186611246565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508593925050610b5c9050565b6000546201000090046001600160a01b03908116911614610a2f576001915050610391565b5060009392505050565b50565b8015610a3957604051600090339060001990849084818181858888f193505050503d8060008114610490576040519150601f19603f3d011682016040523d82523d6000602084013e610490565b6000546201000090046001600160a01b0316331480610aa757503330145b6108a35760405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b60448201526064016103e0565b6000805462010000600160b01b031916620100006001600160a01b038481168202929092178084556040519190048216927f0000000000000000000000000000000000000000000000000000000000000000909216917f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de91a350565b6000806000610b6b8585610b80565b91509150610b7881610bc5565b509392505050565b6000808251604103610bb65760208301516040840151606085015160001a610baa87828585610d0f565b94509450505050610bbe565b506000905060025b9250929050565b6000816004811115610bd957610bd96112fc565b03610be15750565b6001816004811115610bf557610bf56112fc565b03610c425760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016103e0565b6002816004811115610c5657610c566112fc565b03610ca35760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016103e0565b6003816004811115610cb757610cb76112fc565b03610a395760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016103e0565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610d465750600090506003610dca565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610d9a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610dc357600060019250925050610dca565b9150600090505b94509492505050565b6001600160a01b0381168114610a3957600080fd5b60008083601f840112610dfa57600080fd5b50813567ffffffffffffffff811115610e1257600080fd5b602083019150836020828501011115610bbe57600080fd5b60008060008060008060008060c0898b031215610e4657600080fd5b8835610e5181610dd3565b97506020890135610e6181610dd3565b96506040890135610e7181610dd3565b955060608901359450608089013567ffffffffffffffff80821115610e9557600080fd5b610ea18c838d01610de8565b909650945060a08b0135915080821115610eba57600080fd5b50610ec78b828c01610de8565b999c989b5096995094979396929594505050565b600060208284031215610eed57600080fd5b81356001600160e01b0319811681146104b657600080fd5b600080600080600060808688031215610f1d57600080fd5b8535610f2881610dd3565b94506020860135610f3881610dd3565b935060408601359250606086013567ffffffffffffffff811115610f5b57600080fd5b610f6788828901610de8565b969995985093965092949392505050565b60008083601f840112610f8a57600080fd5b50813567ffffffffffffffff811115610fa257600080fd5b6020830191508360208260051b8501011115610bbe57600080fd5b60008060008060408587031215610fd357600080fd5b843567ffffffffffffffff80821115610feb57600080fd5b610ff788838901610f78565b9096509450602087013591508082111561101057600080fd5b5061101d87828801610f78565b95989497509550505050565b60008060006060848603121561103e57600080fd5b833567ffffffffffffffff81111561105557600080fd5b8401610160818703121561106857600080fd5b95602085013595506040909401359392505050565b6000806040838503121561109057600080fd5b823561109b81610dd3565b946020939093013593505050565b600080600080606085870312156110bf57600080fd5b84356110ca81610dd3565b935060208501359250604085013567ffffffffffffffff8111156110ed57600080fd5b61101d87828801610de8565b60008060008060008060008060a0898b03121561111557600080fd5b883561112081610dd3565b9750602089013561113081610dd3565b9650604089013567ffffffffffffffff8082111561114d57600080fd5b6111598c838d01610f78565b909850965060608b013591508082111561117257600080fd5b61117e8c838d01610f78565b909650945060808b0135915080821115610eba57600080fd5b6000602082840312156111a957600080fd5b81356104b681610dd3565b60008060008060008060a087890312156111cd57600080fd5b86356111d881610dd3565b955060208701356111e881610dd3565b94506040870135935060608701359250608087013567ffffffffffffffff81111561121257600080fd5b61121e89828a01610de8565b979a9699509497509295939492505050565b634e487b7160e01b600052603260045260246000fd5b6000808335601e1984360301811261125d57600080fd5b83018035915067ffffffffffffffff82111561127857600080fd5b602001915036819003821315610bbe57600080fd5b6000600182016112ad57634e487b7160e01b600052601160045260246000fd5b5060010190565b6000602082840312156112c657600080fd5b5051919050565b6000825160005b818110156112ee57602081860181015185830152016112d4565b506000920191825250919050565b634e487b7160e01b600052602160045260246000fdfea26469706673582212204f30953d99988a405bf94ae9ebbf486f331d43c7b9f959cd185f40b054ea411664736f6c63430008140033",
}

// AccfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AccfactoryMetaData.ABI instead.
var AccfactoryABI = AccfactoryMetaData.ABI

// AccfactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccfactoryMetaData.Bin instead.
var AccfactoryBin = AccfactoryMetaData.Bin

// DeployAccfactory deploys a new Ethereum contract, binding an instance of Accfactory to it.
func DeployAccfactory(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *Accfactory, error) {
	parsed, err := AccfactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccfactoryBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Accfactory{AccfactoryCaller: AccfactoryCaller{contract: contract}, AccfactoryTransactor: AccfactoryTransactor{contract: contract}, AccfactoryFilterer: AccfactoryFilterer{contract: contract}}, nil
}

// Accfactory is an auto generated Go binding around an Ethereum contract.
type Accfactory struct {
	AccfactoryCaller     // Read-only binding to the contract
	AccfactoryTransactor // Write-only binding to the contract
	AccfactoryFilterer   // Log filterer for contract events
}

// AccfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccfactorySession struct {
	Contract     *Accfactory       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccfactoryCallerSession struct {
	Contract *AccfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AccfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccfactoryTransactorSession struct {
	Contract     *AccfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AccfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccfactoryRaw struct {
	Contract *Accfactory // Generic contract binding to access the raw methods on
}

// AccfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccfactoryCallerRaw struct {
	Contract *AccfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AccfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccfactoryTransactorRaw struct {
	Contract *AccfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccfactory creates a new instance of Accfactory, bound to a specific deployed contract.
func NewAccfactory(address common.Address, backend bind.ContractBackend) (*Accfactory, error) {
	contract, err := bindAccfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accfactory{AccfactoryCaller: AccfactoryCaller{contract: contract}, AccfactoryTransactor: AccfactoryTransactor{contract: contract}, AccfactoryFilterer: AccfactoryFilterer{contract: contract}}, nil
}

// NewAccfactoryCaller creates a new read-only instance of Accfactory, bound to a specific deployed contract.
func NewAccfactoryCaller(address common.Address, caller bind.ContractCaller) (*AccfactoryCaller, error) {
	contract, err := bindAccfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccfactoryCaller{contract: contract}, nil
}

// NewAccfactoryTransactor creates a new write-only instance of Accfactory, bound to a specific deployed contract.
func NewAccfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AccfactoryTransactor, error) {
	contract, err := bindAccfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccfactoryTransactor{contract: contract}, nil
}

// NewAccfactoryFilterer creates a new log filterer instance of Accfactory, bound to a specific deployed contract.
func NewAccfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AccfactoryFilterer, error) {
	contract, err := bindAccfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccfactoryFilterer{contract: contract}, nil
}

// bindAccfactory binds a generic wrapper to an already deployed contract.
func bindAccfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccfactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accfactory *AccfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accfactory.Contract.AccfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accfactory *AccfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accfactory.Contract.AccfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accfactory *AccfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accfactory.Contract.AccfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accfactory *AccfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accfactory *AccfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accfactory *AccfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accfactory.Contract.contract.Transact(opts, method, params...)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_Accfactory *AccfactoryCaller) AccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accfactory.contract.Call(opts, &out, "accountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_Accfactory *AccfactorySession) AccountImplementation() (common.Address, error) {
	return _Accfactory.Contract.AccountImplementation(&_Accfactory.CallOpts)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_Accfactory *AccfactoryCallerSession) AccountImplementation() (common.Address, error) {
	return _Accfactory.Contract.AccountImplementation(&_Accfactory.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_Accfactory *AccfactoryCaller) GetAddress(opts *bind.CallOpts, owner common.Address, salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Accfactory.contract.Call(opts, &out, "getAddress", owner, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_Accfactory *AccfactorySession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _Accfactory.Contract.GetAddress(&_Accfactory.CallOpts, owner, salt)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_Accfactory *AccfactoryCallerSession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _Accfactory.Contract.GetAddress(&_Accfactory.CallOpts, owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_Accfactory *AccfactoryTransactor) CreateAccount(opts *bind.TransactOpts, owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _Accfactory.contract.Transact(opts, "createAccount", owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_Accfactory *AccfactorySession) CreateAccount(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _Accfactory.Contract.CreateAccount(&_Accfactory.TransactOpts, owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_Accfactory *AccfactoryTransactorSession) CreateAccount(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _Accfactory.Contract.CreateAccount(&_Accfactory.TransactOpts, owner, salt)
}

// AccfactoryAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the Accfactory contract.
type AccfactoryAccountCreatedIterator struct {
	Event *AccfactoryAccountCreated // Event containing the contract specifics and raw log

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
func (it *AccfactoryAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccfactoryAccountCreated)
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
		it.Event = new(AccfactoryAccountCreated)
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
func (it *AccfactoryAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccfactoryAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccfactoryAccountCreated represents a AccountCreated event raised by the Accfactory contract.
type AccfactoryAccountCreated struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0x805996f252884581e2f74cf3d2b03564d5ec26ccc90850ae12653dc1b72d1fa2.
//
// Solidity: event AccountCreated(address indexed owner)
func (_Accfactory *AccfactoryFilterer) FilterAccountCreated(opts *bind.FilterOpts, owner []common.Address) (*AccfactoryAccountCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Accfactory.contract.FilterLogs(opts, "AccountCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AccfactoryAccountCreatedIterator{contract: _Accfactory.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0x805996f252884581e2f74cf3d2b03564d5ec26ccc90850ae12653dc1b72d1fa2.
//
// Solidity: event AccountCreated(address indexed owner)
func (_Accfactory *AccfactoryFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *AccfactoryAccountCreated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Accfactory.contract.WatchLogs(opts, "AccountCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccfactoryAccountCreated)
				if err := _Accfactory.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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

// ParseAccountCreated is a log parse operation binding the contract event 0x805996f252884581e2f74cf3d2b03564d5ec26ccc90850ae12653dc1b72d1fa2.
//
// Solidity: event AccountCreated(address indexed owner)
func (_Accfactory *AccfactoryFilterer) ParseAccountCreated(log types.Log) (*AccfactoryAccountCreated, error) {
	event := new(AccfactoryAccountCreated)
	if err := _Accfactory.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
