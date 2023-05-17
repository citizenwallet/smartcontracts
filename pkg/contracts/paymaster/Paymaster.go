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
	Bin: "0x60a060405234801561000f575f80fd5b5060405161180638038061180683398101604081905261002e916100af565b8061003833610060565b6001600160a01b031660805261005a335f908152600360205260409020439055565b506100dc565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602082840312156100bf575f80fd5b81516001600160a01b03811681146100d5575f80fd5b9392505050565b6080516116e16101255f395f81816102ed01528181610441015281816104c5015281816106c001528181610747015281816107c5015281816109220152610b2f01526116e15ff3fe608060405260043610610104575f3560e01c80630396cb6014610108578063205c28781461011d578063382edd9e1461013c578063493b01701461015b5780634a6f84cf146101c55780635476bd72146101fe578063715018a61461021d578063796d4371146102315780638da5cb5b146102465780639ed0fb6814610267578063a9a2340914610289578063addd5099146102a8578063b0d691fe146102dc578063bb9fe6bf1461030f578063c23a5cea14610323578063c23f001f14610342578063c399ec8814610378578063cc9c837c1461038c578063cd8f80c2146103ab578063d0e30db0146103cb578063f2fde38b146103d3578063f465c77e146103f2575b5f80fd5b61011b610116366004611277565b61041f565b005b348015610128575f80fd5b5061011b6101373660046112ae565b6104a6565b348015610147575f80fd5b5061011b6101563660046112d8565b610513565b348015610166575f80fd5b506101ab610175366004611316565b6001600160a01b039182165f90815260026020908152604080832093909416825291825282812054600390925291909120549091565b604080519283526020830191909152015b60405180910390f35b3480156101d0575f80fd5b506101f06101df36600461134d565b60036020525f908152604090205481565b6040519081526020016101bc565b348015610209575f80fd5b5061011b610218366004611316565b6105eb565b348015610228575f80fd5b5061011b61067b565b34801561023c575f80fd5b506101f06188b881565b348015610251575f80fd5b5061025a61068e565b6040516101bc9190611368565b348015610272575f80fd5b5061011b335f908152600360205260409020439055565b348015610294575f80fd5b5061011b6102a336600461137c565b61069c565b3480156102b3575f80fd5b5061025a6102c236600461134d565b60016020525f90815260409020546001600160a01b031681565b3480156102e7575f80fd5b5061025a7f000000000000000000000000000000000000000000000000000000000000000081565b34801561031a575f80fd5b5061011b6106b6565b34801561032e575f80fd5b5061011b61033d36600461134d565b610728565b34801561034d575f80fd5b506101f061035c366004611316565b600260209081525f928352604080842090915290825290205481565b348015610383575f80fd5b506101f06107ac565b348015610397575f80fd5b5061011b6103a63660046112d8565b61083e565b3480156103b6575f80fd5b5061011b335f90815260036020526040812055565b61011b61090b565b3480156103de575f80fd5b5061011b6103ed36600461134d565b610970565b3480156103fd575f80fd5b5061041161040c366004611402565b6109e9565b6040516101bc92919061149d565b610427610a0b565b604051621cb65b60e51b815263ffffffff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690630396cb609034906024015f604051808303818588803b15801561048c575f80fd5b505af115801561049e573d5f803e3d5ffd5b505050505050565b6104ae610a0b565b60405163040b850f60e31b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063205c2878906104fc90859085906004016114be565b5f604051808303815f87803b15801561048c575f80fd5b6105286001600160a01b038416333084610a6a565b6001600160a01b038381165f90815260016020526040902054166105875760405162461bcd60e51b81526020600482015260116024820152703ab739bab83837b93a32b2103a37b5b2b760791b60448201526064015b60405180910390fd5b6001600160a01b038084165f908152600260209081526040808320938616835292905290812080548392906105bd9084906114eb565b90915550506001600160a01b03821633036105e6576105e6335f90815260036020526040812055565b505050565b6105f3610a0b565b6001600160a01b038281165f90815260016020526040902054161561064e5760405162461bcd60e51b8152602060048201526011602482015270151bdad95b88185b1c9958591e481cd95d607a1b604482015260640161057e565b6001600160a01b039182165f90815260016020526040902080546001600160a01b03191691909216179055565b610683610a0b565b61068c5f610ad5565b565b5f546001600160a01b031690565b6106a4610b24565b6106b084848484610b94565b50505050565b6106be610a0b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bb9fe6bf6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610716575f80fd5b505af11580156106b0573d5f803e3d5ffd5b610730610a0b565b60405163611d2e7560e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c23a5cea9061077c908490600401611368565b5f604051808303815f87803b158015610793575f80fd5b505af11580156107a5573d5f803e3d5ffd5b5050505050565b6040516370a0823160e01b81525f906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a08231906107fa903090600401611368565b602060405180830381865afa158015610815573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061083991906114fe565b905090565b335f90815260036020526040902054158015906108685750335f9081526003602052604090205443115b6108bf5760405162461bcd60e51b815260206004820152602260248201527f5061796d61737465723a206d75737420756e6c6f636b546f6b656e4465706f736044820152611a5d60f21b606482015260840161057e565b6001600160a01b0383165f908152600260209081526040808320338452909152812080548392906108f1908490611515565b909155506105e690506001600160a01b0384168383610cae565b60405163b760faf960e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063b760faf9903490610959903090600401611368565b5f604051808303818588803b158015610793575f80fd5b610978610a0b565b6001600160a01b0381166109dd5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161057e565b6109e681610ad5565b50565b60605f6109f4610b24565b6109ff858585610ccd565b91509150935093915050565b33610a1461068e565b6001600160a01b03161461068c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161057e565b6040516001600160a01b03808516602483015283166044820152606481018290526106b09085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610f06565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461068c5760405162461bcd60e51b815260206004820152601560248201527414d95b99195c881b9bdd08115b9d1c9e541bda5b9d605a1b604482015260640161057e565b5f80808080610ba587890189611528565b945094509450945094505f8183856188b8610bc09190611575565b610bca908a6114eb565b610bd49190611575565b610bde919061158c565b905060028a6002811115610bf457610bf46115ab565b14610c1357610c0e6001600160a01b038616873084610a6a565b610c4f565b6001600160a01b038086165f908152600260209081526040808320938a1683529290529081208054839290610c49908490611515565b90915550505b6001600160a01b0385165f9081526002602052604081208291610c7061068e565b6001600160a01b03166001600160a01b031681526020019081526020015f205f828254610c9d91906114eb565b909155505050505050505050505050565b6105e68363a9059cbb60e01b8484604051602401610a9e9291906114be565b60605f6188b88560a0013511610d2f5760405162461bcd60e51b815260206004820152602160248201527f5061796d61737465723a2067617320746f6f206c6f7720666f7220706f73744f6044820152600760fc1b606482015260840161057e565b365f610d3f6101208801886115bf565b909250905060288114610dab5760405162461bcd60e51b815260206004820152602e60248201527f5061796d61737465723a207061796d6173746572416e6444617461206d75737460448201526d1039b832b1b4b33c903a37b5b2b760911b606482015260840161057e565b5f610db98260148186611608565b610dc29161162f565b60601c905087355f610dd48389610fd7565b90505f610de08b6110b0565b6001600160a01b0384165f9081526003602052604090205490915015610e485760405162461bcd60e51b815260206004820152601d60248201527f5061796d61737465723a206465706f736974206e6f74206c6f636b6564000000604482015260640161057e565b6001600160a01b038085165f90815260026020908152604080832093871683529290522054821115610eb95760405162461bcd60e51b815260206004820152601a6024820152795061796d61737465723a206465706f73697420746f6f206c6f7760301b604482015260640161057e565b604080516001600160a01b03948516602082015294909316848401526060840152608083015260a0808301979097528051808303909701875260c090910190525092955f95509350505050565b5f610f5a826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166110de9092919063ffffffff16565b8051909150156105e65780806020019051810190610f78919061165f565b6105e65760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161057e565b6001600160a01b038083165f908152600160205260408120549091168061103f5760405162461bcd60e51b815260206004820152601c60248201527b2830bcb6b0b9ba32b91d103ab739bab83837b93a32b2103a37b5b2b760211b604482015260640161057e565b60405163d1eca9cf60e01b8152600481018490526001600160a01b0382169063d1eca9cf90602401602060405180830381865afa158015611082573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110a691906114fe565b9150505b92915050565b5f60e08201356101008301358082036110ca575092915050565b6110d6824883016110ec565b949350505050565b60606110d684845f85611103565b5f8183106110fa57816110fc565b825b9392505050565b6060824710156111645760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161057e565b5f80866001600160a01b0316858760405161117f919061167e565b5f6040518083038185875af1925050503d805f81146111b9576040519150601f19603f3d011682016040523d82523d5f602084013e6111be565b606091505b50915091506111cf878383876111da565b979650505050505050565b606083156112485782515f03611241576001600160a01b0385163b6112415760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161057e565b50816110d6565b6110d6838381511561125d5781518083602001fd5b8060405162461bcd60e51b815260040161057e9190611699565b5f60208284031215611287575f80fd5b813563ffffffff811681146110fc575f80fd5b6001600160a01b03811681146109e6575f80fd5b5f80604083850312156112bf575f80fd5b82356112ca8161129a565b946020939093013593505050565b5f805f606084860312156112ea575f80fd5b83356112f58161129a565b925060208401356113058161129a565b929592945050506040919091013590565b5f8060408385031215611327575f80fd5b82356113328161129a565b915060208301356113428161129a565b809150509250929050565b5f6020828403121561135d575f80fd5b81356110fc8161129a565b6001600160a01b0391909116815260200190565b5f805f806060858703121561138f575f80fd5b84356003811061139d575f80fd5b935060208501356001600160401b03808211156113b8575f80fd5b818701915087601f8301126113cb575f80fd5b8135818111156113d9575f80fd5b8860208285010111156113ea575f80fd5b95986020929092019750949560400135945092505050565b5f805f60608486031215611414575f80fd5b83356001600160401b03811115611429575f80fd5b8401610160818703121561143b575f80fd5b95602085013595506040909401359392505050565b5f5b8381101561146a578181015183820152602001611452565b50505f910152565b5f8151808452611489816020860160208601611450565b601f01601f19169290920160200192915050565b604081525f6114af6040830185611472565b90508260208301529392505050565b6001600160a01b03929092168252602082015260400190565b634e487b7160e01b5f52601160045260245ffd5b808201808211156110aa576110aa6114d7565b5f6020828403121561150e575f80fd5b5051919050565b818103818111156110aa576110aa6114d7565b5f805f805f60a0868803121561153c575f80fd5b85356115478161129a565b945060208601356115578161129a565b94979496505050506040830135926060810135926080909101359150565b80820281158282048414176110aa576110aa6114d7565b5f826115a657634e487b7160e01b5f52601260045260245ffd5b500490565b634e487b7160e01b5f52602160045260245ffd5b5f808335601e198436030181126115d4575f80fd5b8301803591506001600160401b038211156115ed575f80fd5b602001915036819003821315611601575f80fd5b9250929050565b5f8085851115611616575f80fd5b83861115611622575f80fd5b5050820193919092039150565b6001600160601b031981358181169160148510156116575780818660140360031b1b83161692505b505092915050565b5f6020828403121561166f575f80fd5b815180151581146110fc575f80fd5b5f825161168f818460208701611450565b9190910192915050565b602081525f6110fc602083018461147256fea26469706673582212200ba589bb9b1e5c7bffaef595b36b8c698c3e1f4e8553082d9ed3efea3776fbe864736f6c63430008140033",
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
