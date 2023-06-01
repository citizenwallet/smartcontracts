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
	Bin: "0x60a060405234801561000f575f80fd5b5060405161181e38038061181e83398101604081905261002e916100af565b8061003833610060565b6001600160a01b031660805261005a335f908152600360205260409020439055565b506100dc565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602082840312156100bf575f80fd5b81516001600160a01b03811681146100d5575f80fd5b9392505050565b6080516116f96101255f395f81816103290152818161047d01528181610510015281816106f00152818161077f015281816107f0015281816109480152610b4c01526116f95ff3fe608060405260043610610131575f3560e01c8063addd5099116100a8578063c399ec881161006d578063c399ec88146103b4578063cc9c837c146103c8578063cd8f80c2146103e7578063d0e30db014610407578063f2fde38b1461040f578063f465c77e1461042e575f80fd5b8063addd5099146102e4578063b0d691fe14610318578063bb9fe6bf1461034b578063c23a5cea1461035f578063c23f001f1461037e575f80fd5b80635476bd72116100f95780635476bd721461022b578063715018a61461024a578063796d43711461025e5780638da5cb5b146102735780639ed0fb68146102a3578063a9a23409146102c5575f80fd5b80630396cb6014610135578063205c28781461014a578063382edd9e14610169578063493b0170146101885780634a6f84cf146101f2575b5f80fd5b6101486101433660046112b4565b61045b565b005b348015610155575f80fd5b506101486101643660046112eb565b6104e2565b348015610174575f80fd5b50610148610183366004611315565b610551565b348015610193575f80fd5b506101d86101a2366004611353565b6001600160a01b039182165f90815260026020908152604080832093909416825291825282812054600390925291909120549091565b604080519283526020830191909152015b60405180910390f35b3480156101fd575f80fd5b5061021d61020c36600461138a565b60036020525f908152604090205481565b6040519081526020016101e9565b348015610236575f80fd5b50610148610245366004611353565b610629565b348015610255575f80fd5b506101486106b9565b348015610269575f80fd5b5061021d6188b881565b34801561027e575f80fd5b505f546001600160a01b03165b6040516001600160a01b0390911681526020016101e9565b3480156102ae575f80fd5b50610148335f908152600360205260409020439055565b3480156102d0575f80fd5b506101486102df3660046113a5565b6106cc565b3480156102ef575f80fd5b5061028b6102fe36600461138a565b60016020525f90815260409020546001600160a01b031681565b348015610323575f80fd5b5061028b7f000000000000000000000000000000000000000000000000000000000000000081565b348015610356575f80fd5b506101486106e6565b34801561036a575f80fd5b5061014861037936600461138a565b610758565b348015610389575f80fd5b5061021d610398366004611353565b600260209081525f928352604080842090915290825290205481565b3480156103bf575f80fd5b5061021d6107d9565b3480156103d3575f80fd5b506101486103e2366004611315565b610866565b3480156103f2575f80fd5b50610148335f90815260036020526040812055565b610148610933565b34801561041a575f80fd5b5061014861042936600461138a565b610993565b348015610439575f80fd5b5061044d61044836600461142c565b610a0c565b6040516101e99291906114c8565b610463610a2e565b604051621cb65b60e51b815263ffffffff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690630396cb609034906024015f604051808303818588803b1580156104c8575f80fd5b505af11580156104da573d5f803e3d5ffd5b505050505050565b6104ea610a2e565b60405163040b850f60e31b81526001600160a01b038381166004830152602482018390527f0000000000000000000000000000000000000000000000000000000000000000169063205c2878906044015f604051808303815f87803b1580156104c8575f80fd5b6105666001600160a01b038416333084610a87565b6001600160a01b038381165f90815260016020526040902054166105c55760405162461bcd60e51b81526020600482015260116024820152703ab739bab83837b93a32b2103a37b5b2b760791b60448201526064015b60405180910390fd5b6001600160a01b038084165f908152600260209081526040808320938616835292905290812080548392906105fb9084906114fd565b90915550506001600160a01b038216330361062457610624335f90815260036020526040812055565b505050565b610631610a2e565b6001600160a01b038281165f90815260016020526040902054161561068c5760405162461bcd60e51b8152602060048201526011602482015270151bdad95b88185b1c9958591e481cd95d607a1b60448201526064016105bc565b6001600160a01b039182165f90815260016020526040902080546001600160a01b03191691909216179055565b6106c1610a2e565b6106ca5f610af2565b565b6106d4610b41565b6106e084848484610bb1565b50505050565b6106ee610a2e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bb9fe6bf6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610746575f80fd5b505af11580156106e0573d5f803e3d5ffd5b610760610a2e565b60405163611d2e7560e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c23a5cea906024015f604051808303815f87803b1580156107c0575f80fd5b505af11580156107d2573d5f803e3d5ffd5b5050505050565b6040516370a0823160e01b81523060048201525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa15801561083d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108619190611510565b905090565b335f90815260036020526040902054158015906108905750335f9081526003602052604090205443115b6108e75760405162461bcd60e51b815260206004820152602260248201527f5061796d61737465723a206d75737420756e6c6f636b546f6b656e4465706f736044820152611a5d60f21b60648201526084016105bc565b6001600160a01b0383165f90815260026020908152604080832033845290915281208054839290610919908490611527565b9091555061062490506001600160a01b0384168383610cd4565b60405163b760faf960e01b81523060048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063b760faf99034906024015f604051808303818588803b1580156107c0575f80fd5b61099b610a2e565b6001600160a01b038116610a005760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016105bc565b610a0981610af2565b50565b60605f610a17610b41565b610a22858585610d04565b91509150935093915050565b5f546001600160a01b031633146106ca5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105bc565b6040516001600160a01b03808516602483015283166044820152606481018290526106e09085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610f40565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106ca5760405162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08115b9d1c9e541bda5b9d605a1b60448201526064016105bc565b5f80808080610bc28789018961153a565b945094509450945094505f8183856188b8610bdd9190611587565b610be7908a6114fd565b610bf19190611587565b610bfb919061159e565b905060028a6002811115610c1157610c116115bd565b14610c3057610c2b6001600160a01b038616873084610a87565b610c6c565b6001600160a01b038086165f908152600260209081526040808320938a1683529290529081208054839290610c66908490611527565b90915550505b6001600160a01b0385165f9081526002602052604081208291610c965f546001600160a01b031690565b6001600160a01b03166001600160a01b031681526020019081526020015f205f828254610cc391906114fd565b909155505050505050505050505050565b6040516001600160a01b03831660248201526044810182905261062490849063a9059cbb60e01b90606401610abb565b60605f6188b88560a0013511610d665760405162461bcd60e51b815260206004820152602160248201527f5061796d61737465723a2067617320746f6f206c6f7720666f7220706f73744f6044820152600760fc1b60648201526084016105bc565b365f610d766101208801886115d1565b909250905060288114610de25760405162461bcd60e51b815260206004820152602e60248201527f5061796d61737465723a207061796d6173746572416e6444617461206d75737460448201526d1039b832b1b4b33c903a37b5b2b760911b60648201526084016105bc565b5f610df0826014818661161b565b610df991611642565b60601c905087355f610e0b8389611013565b90505f610e178b6110ed565b6001600160a01b0384165f9081526003602052604090205490915015610e7f5760405162461bcd60e51b815260206004820152601d60248201527f5061796d61737465723a206465706f736974206e6f74206c6f636b656400000060448201526064016105bc565b6001600160a01b038085165f90815260026020908152604080832093871683529290522054821115610ef35760405162461bcd60e51b815260206004820152601a60248201527f5061796d61737465723a206465706f73697420746f6f206c6f7700000000000060448201526064016105bc565b604080516001600160a01b03948516602082015294909316848401526060840152608083015260a0808301979097528051808303909701875260c090910190525092955f95509350505050565b5f610f94826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661111b9092919063ffffffff16565b905080515f1480610fb4575080806020019051810190610fb49190611677565b6106245760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016105bc565b6001600160a01b038083165f908152600160205260408120549091168061107c5760405162461bcd60e51b815260206004820152601c60248201527f5061796d61737465723a20756e737570706f7274656420746f6b656e0000000060448201526064016105bc565b60405163d1eca9cf60e01b8152600481018490526001600160a01b0382169063d1eca9cf90602401602060405180830381865afa1580156110bf573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110e39190611510565b9150505b92915050565b5f60e0820135610100830135808203611107575092915050565b61111382488301611129565b949350505050565b606061111384845f85611140565b5f8183106111375781611139565b825b9392505050565b6060824710156111a15760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016105bc565b5f80866001600160a01b031685876040516111bc9190611696565b5f6040518083038185875af1925050503d805f81146111f6576040519150601f19603f3d011682016040523d82523d5f602084013e6111fb565b606091505b509150915061120c87838387611217565b979650505050505050565b606083156112855782515f0361127e576001600160a01b0385163b61127e5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016105bc565b5081611113565b611113838381511561129a5781518083602001fd5b8060405162461bcd60e51b81526004016105bc91906116b1565b5f602082840312156112c4575f80fd5b813563ffffffff81168114611139575f80fd5b6001600160a01b0381168114610a09575f80fd5b5f80604083850312156112fc575f80fd5b8235611307816112d7565b946020939093013593505050565b5f805f60608486031215611327575f80fd5b8335611332816112d7565b92506020840135611342816112d7565b929592945050506040919091013590565b5f8060408385031215611364575f80fd5b823561136f816112d7565b9150602083013561137f816112d7565b809150509250929050565b5f6020828403121561139a575f80fd5b8135611139816112d7565b5f805f80606085870312156113b8575f80fd5b8435600381106113c6575f80fd5b9350602085013567ffffffffffffffff808211156113e2575f80fd5b818701915087601f8301126113f5575f80fd5b813581811115611403575f80fd5b886020828501011115611414575f80fd5b95986020929092019750949560400135945092505050565b5f805f6060848603121561143e575f80fd5b833567ffffffffffffffff811115611454575f80fd5b84016101608187031215611466575f80fd5b95602085013595506040909401359392505050565b5f5b8381101561149557818101518382015260200161147d565b50505f910152565b5f81518084526114b481602086016020860161147b565b601f01601f19169290920160200192915050565b604081525f6114da604083018561149d565b90508260208301529392505050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156110e7576110e76114e9565b5f60208284031215611520575f80fd5b5051919050565b818103818111156110e7576110e76114e9565b5f805f805f60a0868803121561154e575f80fd5b8535611559816112d7565b94506020860135611569816112d7565b94979496505050506040830135926060810135926080909101359150565b80820281158282048414176110e7576110e76114e9565b5f826115b857634e487b7160e01b5f52601260045260245ffd5b500490565b634e487b7160e01b5f52602160045260245ffd5b5f808335601e198436030181126115e6575f80fd5b83018035915067ffffffffffffffff821115611600575f80fd5b602001915036819003821315611614575f80fd5b9250929050565b5f8085851115611629575f80fd5b83861115611635575f80fd5b5050820193919092039150565b6bffffffffffffffffffffffff19813581811691601485101561166f5780818660140360031b1b83161692505b505092915050565b5f60208284031215611687575f80fd5b81518015158114611139575f80fd5b5f82516116a781846020870161147b565b9190910192915050565b602081525f611139602083018461149d56fea26469706673582212202b651b475ac2876c856aac67c137e282cec8cd3f6b12fd71b1269c6ae2a9b42764736f6c63430008140033",
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
