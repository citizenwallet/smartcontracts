// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paymaster

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

// PaymasterMetaData contains all meta data concerning the Paymaster contract.
var PaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COST_OF_POST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addDepositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIOracle\",\"name\":\"tokenPriceOracle\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unlockBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTokenDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oracles\",\"outputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unlockBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockTokenDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokensTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051620018bf380380620018bf833981016040819052610031916100b4565b8061003b33610064565b6001600160a01b031660805261005e336000908152600360205260409020439055565b506100e4565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100c657600080fd5b81516001600160a01b03811681146100dd57600080fd5b9392505050565b60805161178e620001316000396000818161033f0152818161049d015281816105340152818161071d015281816107b101528181610828015281816109850152610b8e015261178e6000f3fe6080604052600436106101355760003560e01c8063addd5099116100ab578063c399ec881161006f578063c399ec88146103ce578063cc9c837c146103e3578063cd8f80c214610403578063d0e30db014610425578063f2fde38b1461042d578063f465c77e1461044d57600080fd5b8063addd5099146102f7578063b0d691fe1461032d578063bb9fe6bf14610361578063c23a5cea14610376578063c23f001f1461039657600080fd5b80635476bd72116100fd5780635476bd7214610236578063715018a614610256578063796d43711461026b5780638da5cb5b146102815780639ed0fb68146102b3578063a9a23409146102d757600080fd5b80630396cb601461013a578063205c28781461014f578063382edd9e1461016f578063493b01701461018f5780634a6f84cf146101fb575b600080fd5b61014d610148366004611312565b61047b565b005b34801561015b57600080fd5b5061014d61016a36600461134d565b610506565b34801561017b57600080fd5b5061014d61018a366004611379565b610578565b34801561019b57600080fd5b506101e16101aa3660046113ba565b6001600160a01b03918216600090815260026020908152604080832093909416825291825282812054600390925291909120549091565b604080519283526020830191909152015b60405180910390f35b34801561020757600080fd5b506102286102163660046113f3565b60036020526000908152604090205481565b6040519081526020016101f2565b34801561024257600080fd5b5061014d6102513660046113ba565b610653565b34801561026257600080fd5b5061014d6106e5565b34801561027757600080fd5b506102286188b881565b34801561028d57600080fd5b506000546001600160a01b03165b6040516001600160a01b0390911681526020016101f2565b3480156102bf57600080fd5b5061014d336000908152600360205260409020439055565b3480156102e357600080fd5b5061014d6102f2366004611410565b6106f9565b34801561030357600080fd5b5061029b6103123660046113f3565b6001602052600090815260409020546001600160a01b031681565b34801561033957600080fd5b5061029b7f000000000000000000000000000000000000000000000000000000000000000081565b34801561036d57600080fd5b5061014d610713565b34801561038257600080fd5b5061014d6103913660046113f3565b61078a565b3480156103a257600080fd5b506102286103b13660046113ba565b600260209081526000928352604080842090915290825290205481565b3480156103da57600080fd5b50610228610810565b3480156103ef57600080fd5b5061014d6103fe366004611379565b6108a0565b34801561040f57600080fd5b5061014d33600090815260036020526040812055565b61014d610970565b34801561043957600080fd5b5061014d6104483660046113f3565b6109d2565b34801561045957600080fd5b5061046d61046836600461149f565b610a4b565b6040516101f2929190611543565b610483610a6e565b604051621cb65b60e51b815263ffffffff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690630396cb609034906024016000604051808303818588803b1580156104ea57600080fd5b505af11580156104fe573d6000803e3d6000fd5b505050505050565b61050e610a6e565b60405163040b850f60e31b81526001600160a01b038381166004830152602482018390527f0000000000000000000000000000000000000000000000000000000000000000169063205c287890604401600060405180830381600087803b1580156104ea57600080fd5b61058d6001600160a01b038416333084610ac8565b6001600160a01b03838116600090815260016020526040902054166105ed5760405162461bcd60e51b81526020600482015260116024820152703ab739bab83837b93a32b2103a37b5b2b760791b60448201526064015b60405180910390fd5b6001600160a01b0380841660009081526002602090815260408083209386168352929052908120805483929061062490849061157b565b90915550506001600160a01b038216330361064e5761064e33600090815260036020526040812055565b505050565b61065b610a6e565b6001600160a01b0382811660009081526001602052604090205416156106b75760405162461bcd60e51b8152602060048201526011602482015270151bdad95b88185b1c9958591e481cd95d607a1b60448201526064016105e4565b6001600160a01b03918216600090815260016020526040902080546001600160a01b03191691909216179055565b6106ed610a6e565b6106f76000610b33565b565b610701610b83565b61070d84848484610bf3565b50505050565b61071b610a6e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bb9fe6bf6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561077657600080fd5b505af115801561070d573d6000803e3d6000fd5b610792610a6e565b60405163611d2e7560e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c23a5cea90602401600060405180830381600087803b1580156107f557600080fd5b505af1158015610809573d6000803e3d6000fd5b5050505050565b6040516370a0823160e01b81523060048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa158015610877573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089b919061158e565b905090565b33600090815260036020526040902054158015906108cc57503360009081526003602052604090205443115b6109235760405162461bcd60e51b815260206004820152602260248201527f5061796d61737465723a206d75737420756e6c6f636b546f6b656e4465706f736044820152611a5d60f21b60648201526084016105e4565b6001600160a01b0383166000908152600260209081526040808320338452909152812080548392906109569084906115a7565b9091555061064e90506001600160a01b0384168383610d1d565b60405163b760faf960e01b81523060048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063b760faf99034906024016000604051808303818588803b1580156107f557600080fd5b6109da610a6e565b6001600160a01b038116610a3f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016105e4565b610a4881610b33565b50565b60606000610a57610b83565b610a62858585610d4d565b91509150935093915050565b6000546001600160a01b031633146106f75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105e4565b6040516001600160a01b038085166024830152831660448201526064810182905261070d9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610f91565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106f75760405162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08115b9d1c9e541bda5b9d605a1b60448201526064016105e4565b600080808080610c05878901896115ba565b9450945094509450945060008183856188b8610c21919061160b565b610c2b908a61157b565b610c35919061160b565b610c3f9190611622565b905060028a6002811115610c5557610c55611644565b14610c7457610c6f6001600160a01b038616873084610ac8565b610cb1565b6001600160a01b038086166000908152600260209081526040808320938a1683529290529081208054839290610cab9084906115a7565b90915550505b6001600160a01b03851660009081526002602052604081208291610cdd6000546001600160a01b031690565b6001600160a01b03166001600160a01b031681526020019081526020016000206000828254610d0c919061157b565b909155505050505050505050505050565b6040516001600160a01b03831660248201526044810182905261064e90849063a9059cbb60e01b90606401610afc565b606060006188b88560a0013511610db05760405162461bcd60e51b815260206004820152602160248201527f5061796d61737465723a2067617320746f6f206c6f7720666f7220706f73744f6044820152600760fc1b60648201526084016105e4565b366000610dc161012088018861165a565b909250905060288114610e2d5760405162461bcd60e51b815260206004820152602e60248201527f5061796d61737465723a207061796d6173746572416e6444617461206d75737460448201526d1039b832b1b4b33c903a37b5b2b760911b60648201526084016105e4565b6000610e3c82601481866116a8565b610e45916116d2565b60601c905087356000610e588389611066565b90506000610e658b611143565b6001600160a01b03841660009081526003602052604090205490915015610ece5760405162461bcd60e51b815260206004820152601d60248201527f5061796d61737465723a206465706f736974206e6f74206c6f636b656400000060448201526064016105e4565b6001600160a01b03808516600090815260026020908152604080832093871683529290522054821115610f435760405162461bcd60e51b815260206004820152601a60248201527f5061796d61737465723a206465706f73697420746f6f206c6f7700000000000060448201526064016105e4565b604080516001600160a01b03948516602082015294909316848401526060840152608083015260a0808301979097528051808303909701875260c09091019052509295600095509350505050565b6000610fe6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166111729092919063ffffffff16565b90508051600014806110075750808060200190518101906110079190611707565b61064e5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016105e4565b6001600160a01b03808316600090815260016020526040812054909116806110d05760405162461bcd60e51b815260206004820152601c60248201527f5061796d61737465723a20756e737570706f7274656420746f6b656e0000000060448201526064016105e4565b60405163d1eca9cf60e01b8152600481018490526001600160a01b0382169063d1eca9cf90602401602060405180830381865afa158015611115573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611139919061158e565b9150505b92915050565b600060e082013561010083013580820361115e575092915050565b61116a82488301611181565b949350505050565b606061116a8484600085611199565b60008183106111905781611192565b825b9392505050565b6060824710156111fa5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016105e4565b600080866001600160a01b031685876040516112169190611729565b60006040518083038185875af1925050503d8060008114611253576040519150601f19603f3d011682016040523d82523d6000602084013e611258565b606091505b509150915061126987838387611274565b979650505050505050565b606083156112e35782516000036112dc576001600160a01b0385163b6112dc5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016105e4565b508161116a565b61116a83838151156112f85781518083602001fd5b8060405162461bcd60e51b81526004016105e49190611745565b60006020828403121561132457600080fd5b813563ffffffff8116811461119257600080fd5b6001600160a01b0381168114610a4857600080fd5b6000806040838503121561136057600080fd5b823561136b81611338565b946020939093013593505050565b60008060006060848603121561138e57600080fd5b833561139981611338565b925060208401356113a981611338565b929592945050506040919091013590565b600080604083850312156113cd57600080fd5b82356113d881611338565b915060208301356113e881611338565b809150509250929050565b60006020828403121561140557600080fd5b813561119281611338565b6000806000806060858703121561142657600080fd5b84356003811061143557600080fd5b9350602085013567ffffffffffffffff8082111561145257600080fd5b818701915087601f83011261146657600080fd5b81358181111561147557600080fd5b88602082850101111561148757600080fd5b95986020929092019750949560400135945092505050565b6000806000606084860312156114b457600080fd5b833567ffffffffffffffff8111156114cb57600080fd5b840161016081870312156114de57600080fd5b95602085013595506040909401359392505050565b60005b8381101561150e5781810151838201526020016114f6565b50506000910152565b6000815180845261152f8160208601602086016114f3565b601f01601f19169290920160200192915050565b6040815260006115566040830185611517565b90508260208301529392505050565b634e487b7160e01b600052601160045260246000fd5b8082018082111561113d5761113d611565565b6000602082840312156115a057600080fd5b5051919050565b8181038181111561113d5761113d611565565b600080600080600060a086880312156115d257600080fd5b85356115dd81611338565b945060208601356115ed81611338565b94979496505050506040830135926060810135926080909101359150565b808202811582820484141761113d5761113d611565565b60008261163f57634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052602160045260246000fd5b6000808335601e1984360301811261167157600080fd5b83018035915067ffffffffffffffff82111561168c57600080fd5b6020019150368190038213156116a157600080fd5b9250929050565b600080858511156116b857600080fd5b838611156116c557600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff1981358181169160148510156116ff5780818660140360031b1b83161692505b505092915050565b60006020828403121561171957600080fd5b8151801515811461119257600080fd5b6000825161173b8184602087016114f3565b9190910192915050565b602081526000611192602083018461151756fea264697066735822122023136e1fb0d65826bd6a7d8e4fe731ecfd41362650adfb9a39c31d89f69fff3f64736f6c63430008140033",
}

// PaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymasterMetaData.ABI instead.
var PaymasterABI = PaymasterMetaData.ABI

// PaymasterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PaymasterMetaData.Bin instead.
var PaymasterBin = PaymasterMetaData.Bin

// DeployPaymaster deploys a new Ethereum contract, binding an instance of Paymaster to it.
func DeployPaymaster(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *Paymaster, error) {
	parsed, err := PaymasterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PaymasterBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Paymaster{PaymasterCaller: PaymasterCaller{contract: contract}, PaymasterTransactor: PaymasterTransactor{contract: contract}, PaymasterFilterer: PaymasterFilterer{contract: contract}}, nil
}

// Paymaster is an auto generated Go binding around an Ethereum contract.
type Paymaster struct {
	PaymasterCaller     // Read-only binding to the contract
	PaymasterTransactor // Write-only binding to the contract
	PaymasterFilterer   // Log filterer for contract events
}

// PaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymasterSession struct {
	Contract     *Paymaster        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymasterCallerSession struct {
	Contract *PaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymasterTransactorSession struct {
	Contract     *PaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymasterRaw struct {
	Contract *Paymaster // Generic contract binding to access the raw methods on
}

// PaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymasterCallerRaw struct {
	Contract *PaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// PaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymasterTransactorRaw struct {
	Contract *PaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymaster creates a new instance of Paymaster, bound to a specific deployed contract.
func NewPaymaster(address common.Address, backend bind.ContractBackend) (*Paymaster, error) {
	contract, err := bindPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Paymaster{PaymasterCaller: PaymasterCaller{contract: contract}, PaymasterTransactor: PaymasterTransactor{contract: contract}, PaymasterFilterer: PaymasterFilterer{contract: contract}}, nil
}

// NewPaymasterCaller creates a new read-only instance of Paymaster, bound to a specific deployed contract.
func NewPaymasterCaller(address common.Address, caller bind.ContractCaller) (*PaymasterCaller, error) {
	contract, err := bindPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymasterCaller{contract: contract}, nil
}

// NewPaymasterTransactor creates a new write-only instance of Paymaster, bound to a specific deployed contract.
func NewPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymasterTransactor, error) {
	contract, err := bindPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymasterTransactor{contract: contract}, nil
}

// NewPaymasterFilterer creates a new log filterer instance of Paymaster, bound to a specific deployed contract.
func NewPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymasterFilterer, error) {
	contract, err := bindPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymasterFilterer{contract: contract}, nil
}

// bindPaymaster binds a generic wrapper to an already deployed contract.
func bindPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paymaster *PaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Paymaster.Contract.PaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paymaster *PaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymaster.Contract.PaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paymaster *PaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paymaster.Contract.PaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paymaster *PaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Paymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paymaster *PaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paymaster *PaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paymaster.Contract.contract.Transact(opts, method, params...)
}

// COSTOFPOST is a free data retrieval call binding the contract method 0x796d4371.
//
// Solidity: function COST_OF_POST() view returns(uint256)
func (_Paymaster *PaymasterCaller) COSTOFPOST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "COST_OF_POST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COSTOFPOST is a free data retrieval call binding the contract method 0x796d4371.
//
// Solidity: function COST_OF_POST() view returns(uint256)
func (_Paymaster *PaymasterSession) COSTOFPOST() (*big.Int, error) {
	return _Paymaster.Contract.COSTOFPOST(&_Paymaster.CallOpts)
}

// COSTOFPOST is a free data retrieval call binding the contract method 0x796d4371.
//
// Solidity: function COST_OF_POST() view returns(uint256)
func (_Paymaster *PaymasterCallerSession) COSTOFPOST() (*big.Int, error) {
	return _Paymaster.Contract.COSTOFPOST(&_Paymaster.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances(address , address ) view returns(uint256)
func (_Paymaster *PaymasterCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "balances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances(address , address ) view returns(uint256)
func (_Paymaster *PaymasterSession) Balances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Paymaster.Contract.Balances(&_Paymaster.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances(address , address ) view returns(uint256)
func (_Paymaster *PaymasterCallerSession) Balances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Paymaster.Contract.Balances(&_Paymaster.CallOpts, arg0, arg1)
}

// DepositInfo is a free data retrieval call binding the contract method 0x493b0170.
//
// Solidity: function depositInfo(address token, address account) view returns(uint256 amount, uint256 _unlockBlock)
func (_Paymaster *PaymasterCaller) DepositInfo(opts *bind.CallOpts, token common.Address, account common.Address) (struct {
	Amount      *big.Int
	UnlockBlock *big.Int
}, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "depositInfo", token, account)

	outstruct := new(struct {
		Amount      *big.Int
		UnlockBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnlockBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DepositInfo is a free data retrieval call binding the contract method 0x493b0170.
//
// Solidity: function depositInfo(address token, address account) view returns(uint256 amount, uint256 _unlockBlock)
func (_Paymaster *PaymasterSession) DepositInfo(token common.Address, account common.Address) (struct {
	Amount      *big.Int
	UnlockBlock *big.Int
}, error) {
	return _Paymaster.Contract.DepositInfo(&_Paymaster.CallOpts, token, account)
}

// DepositInfo is a free data retrieval call binding the contract method 0x493b0170.
//
// Solidity: function depositInfo(address token, address account) view returns(uint256 amount, uint256 _unlockBlock)
func (_Paymaster *PaymasterCallerSession) DepositInfo(token common.Address, account common.Address) (struct {
	Amount      *big.Int
	UnlockBlock *big.Int
}, error) {
	return _Paymaster.Contract.DepositInfo(&_Paymaster.CallOpts, token, account)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Paymaster *PaymasterCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Paymaster *PaymasterSession) EntryPoint() (common.Address, error) {
	return _Paymaster.Contract.EntryPoint(&_Paymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Paymaster *PaymasterCallerSession) EntryPoint() (common.Address, error) {
	return _Paymaster.Contract.EntryPoint(&_Paymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Paymaster *PaymasterCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Paymaster *PaymasterSession) GetDeposit() (*big.Int, error) {
	return _Paymaster.Contract.GetDeposit(&_Paymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Paymaster *PaymasterCallerSession) GetDeposit() (*big.Int, error) {
	return _Paymaster.Contract.GetDeposit(&_Paymaster.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0xaddd5099.
//
// Solidity: function oracles(address ) view returns(address)
func (_Paymaster *PaymasterCaller) Oracles(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "oracles", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracles is a free data retrieval call binding the contract method 0xaddd5099.
//
// Solidity: function oracles(address ) view returns(address)
func (_Paymaster *PaymasterSession) Oracles(arg0 common.Address) (common.Address, error) {
	return _Paymaster.Contract.Oracles(&_Paymaster.CallOpts, arg0)
}

// Oracles is a free data retrieval call binding the contract method 0xaddd5099.
//
// Solidity: function oracles(address ) view returns(address)
func (_Paymaster *PaymasterCallerSession) Oracles(arg0 common.Address) (common.Address, error) {
	return _Paymaster.Contract.Oracles(&_Paymaster.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Paymaster *PaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Paymaster *PaymasterSession) Owner() (common.Address, error) {
	return _Paymaster.Contract.Owner(&_Paymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Paymaster *PaymasterCallerSession) Owner() (common.Address, error) {
	return _Paymaster.Contract.Owner(&_Paymaster.CallOpts)
}

// UnlockBlock is a free data retrieval call binding the contract method 0x4a6f84cf.
//
// Solidity: function unlockBlock(address ) view returns(uint256)
func (_Paymaster *PaymasterCaller) UnlockBlock(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "unlockBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockBlock is a free data retrieval call binding the contract method 0x4a6f84cf.
//
// Solidity: function unlockBlock(address ) view returns(uint256)
func (_Paymaster *PaymasterSession) UnlockBlock(arg0 common.Address) (*big.Int, error) {
	return _Paymaster.Contract.UnlockBlock(&_Paymaster.CallOpts, arg0)
}

// UnlockBlock is a free data retrieval call binding the contract method 0x4a6f84cf.
//
// Solidity: function unlockBlock(address ) view returns(uint256)
func (_Paymaster *PaymasterCallerSession) UnlockBlock(arg0 common.Address) (*big.Int, error) {
	return _Paymaster.Contract.UnlockBlock(&_Paymaster.CallOpts, arg0)
}

// AddDepositFor is a paid mutator transaction binding the contract method 0x382edd9e.
//
// Solidity: function addDepositFor(address token, address account, uint256 amount) returns()
func (_Paymaster *PaymasterTransactor) AddDepositFor(opts *bind.TransactOpts, token common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "addDepositFor", token, account, amount)
}

// AddDepositFor is a paid mutator transaction binding the contract method 0x382edd9e.
//
// Solidity: function addDepositFor(address token, address account, uint256 amount) returns()
func (_Paymaster *PaymasterSession) AddDepositFor(token common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Paymaster.Contract.AddDepositFor(&_Paymaster.TransactOpts, token, account, amount)
}

// AddDepositFor is a paid mutator transaction binding the contract method 0x382edd9e.
//
// Solidity: function addDepositFor(address token, address account, uint256 amount) returns()
func (_Paymaster *PaymasterTransactorSession) AddDepositFor(token common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Paymaster.Contract.AddDepositFor(&_Paymaster.TransactOpts, token, account, amount)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Paymaster *PaymasterTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Paymaster *PaymasterSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Paymaster.Contract.AddStake(&_Paymaster.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Paymaster *PaymasterTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Paymaster.Contract.AddStake(&_Paymaster.TransactOpts, unstakeDelaySec)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address token, address tokenPriceOracle) returns()
func (_Paymaster *PaymasterTransactor) AddToken(opts *bind.TransactOpts, token common.Address, tokenPriceOracle common.Address) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "addToken", token, tokenPriceOracle)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address token, address tokenPriceOracle) returns()
func (_Paymaster *PaymasterSession) AddToken(token common.Address, tokenPriceOracle common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.AddToken(&_Paymaster.TransactOpts, token, tokenPriceOracle)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address token, address tokenPriceOracle) returns()
func (_Paymaster *PaymasterTransactorSession) AddToken(token common.Address, tokenPriceOracle common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.AddToken(&_Paymaster.TransactOpts, token, tokenPriceOracle)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Paymaster *PaymasterTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Paymaster *PaymasterSession) Deposit() (*types.Transaction, error) {
	return _Paymaster.Contract.Deposit(&_Paymaster.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Paymaster *PaymasterTransactorSession) Deposit() (*types.Transaction, error) {
	return _Paymaster.Contract.Deposit(&_Paymaster.TransactOpts)
}

// LockTokenDeposit is a paid mutator transaction binding the contract method 0xcd8f80c2.
//
// Solidity: function lockTokenDeposit() returns()
func (_Paymaster *PaymasterTransactor) LockTokenDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "lockTokenDeposit")
}

// LockTokenDeposit is a paid mutator transaction binding the contract method 0xcd8f80c2.
//
// Solidity: function lockTokenDeposit() returns()
func (_Paymaster *PaymasterSession) LockTokenDeposit() (*types.Transaction, error) {
	return _Paymaster.Contract.LockTokenDeposit(&_Paymaster.TransactOpts)
}

// LockTokenDeposit is a paid mutator transaction binding the contract method 0xcd8f80c2.
//
// Solidity: function lockTokenDeposit() returns()
func (_Paymaster *PaymasterTransactorSession) LockTokenDeposit() (*types.Transaction, error) {
	return _Paymaster.Contract.LockTokenDeposit(&_Paymaster.TransactOpts)
}

// PostOp is a paid mutator transaction binding the contract method 0xa9a23409.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost) returns()
func (_Paymaster *PaymasterTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "postOp", mode, context, actualGasCost)
}

// PostOp is a paid mutator transaction binding the contract method 0xa9a23409.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost) returns()
func (_Paymaster *PaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int) (*types.Transaction, error) {
	return _Paymaster.Contract.PostOp(&_Paymaster.TransactOpts, mode, context, actualGasCost)
}

// PostOp is a paid mutator transaction binding the contract method 0xa9a23409.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost) returns()
func (_Paymaster *PaymasterTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int) (*types.Transaction, error) {
	return _Paymaster.Contract.PostOp(&_Paymaster.TransactOpts, mode, context, actualGasCost)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Paymaster *PaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Paymaster *PaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Paymaster.Contract.RenounceOwnership(&_Paymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Paymaster *PaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Paymaster.Contract.RenounceOwnership(&_Paymaster.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Paymaster *PaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Paymaster *PaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.TransferOwnership(&_Paymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Paymaster *PaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.TransferOwnership(&_Paymaster.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Paymaster *PaymasterTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Paymaster *PaymasterSession) UnlockStake() (*types.Transaction, error) {
	return _Paymaster.Contract.UnlockStake(&_Paymaster.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Paymaster *PaymasterTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _Paymaster.Contract.UnlockStake(&_Paymaster.TransactOpts)
}

// UnlockTokenDeposit is a paid mutator transaction binding the contract method 0x9ed0fb68.
//
// Solidity: function unlockTokenDeposit() returns()
func (_Paymaster *PaymasterTransactor) UnlockTokenDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "unlockTokenDeposit")
}

// UnlockTokenDeposit is a paid mutator transaction binding the contract method 0x9ed0fb68.
//
// Solidity: function unlockTokenDeposit() returns()
func (_Paymaster *PaymasterSession) UnlockTokenDeposit() (*types.Transaction, error) {
	return _Paymaster.Contract.UnlockTokenDeposit(&_Paymaster.TransactOpts)
}

// UnlockTokenDeposit is a paid mutator transaction binding the contract method 0x9ed0fb68.
//
// Solidity: function unlockTokenDeposit() returns()
func (_Paymaster *PaymasterTransactorSession) UnlockTokenDeposit() (*types.Transaction, error) {
	return _Paymaster.Contract.UnlockTokenDeposit(&_Paymaster.TransactOpts)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0xf465c77e.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_Paymaster *PaymasterTransactor) ValidatePaymasterUserOp(opts *bind.TransactOpts, userOp UserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "validatePaymasterUserOp", userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0xf465c77e.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_Paymaster *PaymasterSession) ValidatePaymasterUserOp(userOp UserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _Paymaster.Contract.ValidatePaymasterUserOp(&_Paymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0xf465c77e.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_Paymaster *PaymasterTransactorSession) ValidatePaymasterUserOp(userOp UserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _Paymaster.Contract.ValidatePaymasterUserOp(&_Paymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Paymaster *PaymasterTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Paymaster *PaymasterSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.WithdrawStake(&_Paymaster.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Paymaster *PaymasterTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.WithdrawStake(&_Paymaster.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_Paymaster *PaymasterTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "withdrawTo", withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_Paymaster *PaymasterSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Paymaster.Contract.WithdrawTo(&_Paymaster.TransactOpts, withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_Paymaster *PaymasterTransactorSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Paymaster.Contract.WithdrawTo(&_Paymaster.TransactOpts, withdrawAddress, amount)
}

// WithdrawTokensTo is a paid mutator transaction binding the contract method 0xcc9c837c.
//
// Solidity: function withdrawTokensTo(address token, address target, uint256 amount) returns()
func (_Paymaster *PaymasterTransactor) WithdrawTokensTo(opts *bind.TransactOpts, token common.Address, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "withdrawTokensTo", token, target, amount)
}

// WithdrawTokensTo is a paid mutator transaction binding the contract method 0xcc9c837c.
//
// Solidity: function withdrawTokensTo(address token, address target, uint256 amount) returns()
func (_Paymaster *PaymasterSession) WithdrawTokensTo(token common.Address, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Paymaster.Contract.WithdrawTokensTo(&_Paymaster.TransactOpts, token, target, amount)
}

// WithdrawTokensTo is a paid mutator transaction binding the contract method 0xcc9c837c.
//
// Solidity: function withdrawTokensTo(address token, address target, uint256 amount) returns()
func (_Paymaster *PaymasterTransactorSession) WithdrawTokensTo(token common.Address, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Paymaster.Contract.WithdrawTokensTo(&_Paymaster.TransactOpts, token, target, amount)
}

// PaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Paymaster contract.
type PaymasterOwnershipTransferredIterator struct {
	Event *PaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymasterOwnershipTransferred)
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
		it.Event = new(PaymasterOwnershipTransferred)
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
func (it *PaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the Paymaster contract.
type PaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Paymaster *PaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Paymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymasterOwnershipTransferredIterator{contract: _Paymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Paymaster *PaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Paymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymasterOwnershipTransferred)
				if err := _Paymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Paymaster *PaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*PaymasterOwnershipTransferred, error) {
	event := new(PaymasterOwnershipTransferred)
	if err := _Paymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
