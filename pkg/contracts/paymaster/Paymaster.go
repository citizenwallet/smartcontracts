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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"senderNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100e1565b600054610100900460ff161561008e5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146100df576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6080516117c6610118600039600081816102a9015281816102f20152818161038e015281816103ce015261046101526117c66000f3fe6080604052600436106100c25760003560e01c80638da5cb5b1161007f578063a9a2340911610059578063a9a2340914610205578063c4d66de814610226578063f2fde38b14610246578063f465c77e1461026657600080fd5b80638da5cb5b1461019057806394e1fc19146101b85780639c90b443146101d857600080fd5b80630bd28e3b146100c757806335567e1a146100e95780633659cfe6146101335780634f1ef2861461015357806352d1902d14610166578063715018a61461017b575b600080fd5b3480156100d357600080fd5b506100e76100e23660046111c2565b610294565b005b3480156100f557600080fd5b506101206101043660046111f4565b506001600160a01b0316600090815260c9602052604090205490565b6040519081526020015b60405180910390f35b34801561013f57600080fd5b506100e761014e366004611227565b61029f565b6100e7610161366004611258565b610384565b34801561017257600080fd5b50610120610454565b34801561018757600080fd5b506100e7610507565b34801561019c57600080fd5b506033546040516001600160a01b03909116815260200161012a565b3480156101c457600080fd5b506101206101d3366004611349565b61051b565b3480156101e457600080fd5b506101206101f3366004611227565b60c96020526000908152604090205481565b34801561021157600080fd5b506100e76102203660046113a7565b50505050565b34801561023257600080fd5b506100e7610241366004611227565b610578565b34801561025257600080fd5b506100e7610261366004611227565b610692565b34801561027257600080fd5b50610286610281366004611436565b610708565b60405161012a9291906114d4565b61029c610953565b50565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036102f05760405162461bcd60e51b81526004016102e7906114f6565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661033960008051602061174a833981519152546001600160a01b031690565b6001600160a01b03161461035f5760405162461bcd60e51b81526004016102e790611542565b61036881610294565b6040805160008082526020820190925261029c918391906109ad565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036103cc5760405162461bcd60e51b81526004016102e7906114f6565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661041560008051602061174a833981519152546001600160a01b031690565b6001600160a01b03161461043b5760405162461bcd60e51b81526004016102e790611542565b61044482610294565b610450828260016109ad565b5050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104f45760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c000000000000000060648201526084016102e7565b5060008051602061174a83398151915290565b61050f610953565b6105196000610b1d565b565b600061052684610b6f565b6001600160a01b03853516600090815260c96020908152604091829020549151610559939246923092899189910161158e565b6040516020818303038152906040528051906020012090509392505050565b600054610100900460ff16158080156105985750600054600160ff909116105b806105b25750303b1580156105b2575060005460ff166001145b6106155760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102e7565b6000805460ff191660011790558015610638576000805461ff0019166101001790555b610640610c27565b61064982610c56565b8015610450576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b61069a610953565b6001600160a01b0381166106ff5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016102e7565b61029c81610b1d565b60606000808036816107266107216101208b018b6115de565b610c5f565b92965090945092509050604081148061073f5750604181145b6107b3576040805162461bcd60e51b81526020600482015260248101919091527f566572696679696e675061796d61737465723a20696e76616c6964207369676e60448201527f6174757265206c656e67746820696e207061796d6173746572416e644461746160648201526084016102e7565b4265ffffffffffff8085169082161180156107de57508465ffffffffffff168165ffffffffffff1611155b6108225760405162461bcd60e51b81526020600482015260156024820152741cda59db985d1d5c99481a185cc8195e1c1a5c9959605a1b60448201526064016102e7565b60006108656108328c888861051b565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c91909152603c902090565b6001600160a01b038c3516600090815260c96020526040812080549293509061088d83611625565b91905055506108d484848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508593925050610c9c9050565b6001600160a01b03166108ef6033546001600160a01b031690565b6001600160a01b0316146109455760405162461bcd60e51b815260206004820152601b60248201527f696e76616c6964207061796d6173746572207369676e6174757265000000000060448201526064016102e7565b505050505050935093915050565b6033546001600160a01b031633146105195760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102e7565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff16156109e5576109e083610cc0565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610a3f575060408051601f3d908101601f19168201909252610a3c9181019061164c565b60015b610aa25760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b60648201526084016102e7565b60008051602061174a8339815191528114610b115760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b60648201526084016102e7565b506109e0838383610d5c565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6060813560208301356000610b8f610b8a60408701876115de565b610d81565b90506000610ba3610b8a60608801886115de565b604080516001600160a01b039690961660208701528581019490945260608501929092525060808084019190915284013560a08084019190915284013560c08084019190915284013560e080840191909152840135610100808401919091529093013561012080830191909152835180830390910181526101409091019092525090565b600054610100900460ff16610c4e5760405162461bcd60e51b81526004016102e790611665565b610519610d94565b61029c81610692565b6000803681610c726054601487896116b0565b810190610c7f91906116da565b9094509250610c9185605481896116b0565b949793965094505050565b6000806000610cab8585610dc4565b91509150610cb881610e09565b509392505050565b6001600160a01b0381163b610d2d5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016102e7565b60008051602061174a83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b610d6583610f53565b600082511180610d725750805b156109e0576102208383610f93565b6000604051828085833790209392505050565b600054610100900460ff16610dbb5760405162461bcd60e51b81526004016102e790611665565b61051933610b1d565b6000808251604103610dfa5760208301516040840151606085015160001a610dee87828585610fbf565b94509450505050610e02565b506000905060025b9250929050565b6000816004811115610e1d57610e1d611704565b03610e255750565b6001816004811115610e3957610e39611704565b03610e865760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016102e7565b6002816004811115610e9a57610e9a611704565b03610ee75760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016102e7565b6003816004811115610efb57610efb611704565b0361029c5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016102e7565b610f5c81610cc0565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060610fb8838360405180606001604052806027815260200161176a60279139611083565b9392505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610ff6575060009050600361107a565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561104a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166110735760006001925092505061107a565b9150600090505b94509492505050565b6060600080856001600160a01b0316856040516110a0919061171a565b600060405180830381855af49150503d80600081146110db576040519150601f19603f3d011682016040523d82523d6000602084013e6110e0565b606091505b50915091506110f1868383876110fb565b9695505050505050565b6060831561116a578251600003611163576001600160a01b0385163b6111635760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102e7565b5081611174565b611174838361117c565b949350505050565b81511561118c5781518083602001fd5b8060405162461bcd60e51b81526004016102e79190611736565b80356001600160c01b03811681146111bd57600080fd5b919050565b6000602082840312156111d457600080fd5b610fb8826111a6565b80356001600160a01b03811681146111bd57600080fd5b6000806040838503121561120757600080fd5b611210836111dd565b915061121e602084016111a6565b90509250929050565b60006020828403121561123957600080fd5b610fb8826111dd565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561126b57600080fd5b611274836111dd565b9150602083013567ffffffffffffffff8082111561129157600080fd5b818501915085601f8301126112a557600080fd5b8135818111156112b7576112b7611242565b604051601f8201601f19908116603f011681019083821181831017156112df576112df611242565b816040528281528860208487010111156112f857600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b6000610160828403121561132d57600080fd5b50919050565b803565ffffffffffff811681146111bd57600080fd5b60008060006060848603121561135e57600080fd5b833567ffffffffffffffff81111561137557600080fd5b6113818682870161131a565b93505061139060208501611333565b915061139e60408501611333565b90509250925092565b600080600080606085870312156113bd57600080fd5b8435600381106113cc57600080fd5b9350602085013567ffffffffffffffff808211156113e957600080fd5b818701915087601f8301126113fd57600080fd5b81358181111561140c57600080fd5b88602082850101111561141e57600080fd5b95986020929092019750949560400135945092505050565b60008060006060848603121561144b57600080fd5b833567ffffffffffffffff81111561146257600080fd5b61146e8682870161131a565b9660208601359650604090950135949350505050565b60005b8381101561149f578181015183820152602001611487565b50506000910152565b600081518084526114c0816020860160208601611484565b601f01601f19169290920160200192915050565b6040815260006114e760408301856114a8565b90508260208301529392505050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b60c0815260006115a160c08301896114a8565b6020830197909752506001600160a01b03949094166040850152606084019290925265ffffffffffff90811660808401521660a090910152919050565b6000808335601e198436030181126115f557600080fd5b83018035915067ffffffffffffffff82111561161057600080fd5b602001915036819003821315610e0257600080fd5b60006001820161164557634e487b7160e01b600052601160045260246000fd5b5060010190565b60006020828403121561165e57600080fd5b5051919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b600080858511156116c057600080fd5b838611156116cd57600080fd5b5050820193919092039150565b600080604083850312156116ed57600080fd5b6116f683611333565b915061121e60208401611333565b634e487b7160e01b600052602160045260246000fd5b6000825161172c818460208701611484565b9190910192915050565b602081526000610fb860208301846114a856fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220995ae698df02bf4ad8d5a1095d18019200a80b3ba9a40ac52fd640da873d12e064736f6c63430008140033",
}

// PaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymasterMetaData.ABI instead.
var PaymasterABI = PaymasterMetaData.ABI

// PaymasterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PaymasterMetaData.Bin instead.
var PaymasterBin = PaymasterMetaData.Bin

// DeployPaymaster deploys a new Ethereum contract, binding an instance of Paymaster to it.
func DeployPaymaster(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Paymaster, error) {
	parsed, err := PaymasterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PaymasterBin), backend)
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

// GetHash is a free data retrieval call binding the contract method 0x94e1fc19.
//
// Solidity: function getHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, uint48 validUntil, uint48 validAfter) view returns(bytes32)
func (_Paymaster *PaymasterCaller) GetHash(opts *bind.CallOpts, userOp UserOperation, validUntil *big.Int, validAfter *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "getHash", userOp, validUntil, validAfter)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0x94e1fc19.
//
// Solidity: function getHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, uint48 validUntil, uint48 validAfter) view returns(bytes32)
func (_Paymaster *PaymasterSession) GetHash(userOp UserOperation, validUntil *big.Int, validAfter *big.Int) ([32]byte, error) {
	return _Paymaster.Contract.GetHash(&_Paymaster.CallOpts, userOp, validUntil, validAfter)
}

// GetHash is a free data retrieval call binding the contract method 0x94e1fc19.
//
// Solidity: function getHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, uint48 validUntil, uint48 validAfter) view returns(bytes32)
func (_Paymaster *PaymasterCallerSession) GetHash(userOp UserOperation, validUntil *big.Int, validAfter *big.Int) ([32]byte, error) {
	return _Paymaster.Contract.GetHash(&_Paymaster.CallOpts, userOp, validUntil, validAfter)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_Paymaster *PaymasterCaller) GetNonce(opts *bind.CallOpts, sender common.Address, key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "getNonce", sender, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_Paymaster *PaymasterSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _Paymaster.Contract.GetNonce(&_Paymaster.CallOpts, sender, key)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_Paymaster *PaymasterCallerSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _Paymaster.Contract.GetNonce(&_Paymaster.CallOpts, sender, key)
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

// PostOp is a free data retrieval call binding the contract method 0xa9a23409.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost) view returns()
func (_Paymaster *PaymasterCaller) PostOp(opts *bind.CallOpts, mode uint8, context []byte, actualGasCost *big.Int) error {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "postOp", mode, context, actualGasCost)

	if err != nil {
		return err
	}

	return err

}

// PostOp is a free data retrieval call binding the contract method 0xa9a23409.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost) view returns()
func (_Paymaster *PaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int) error {
	return _Paymaster.Contract.PostOp(&_Paymaster.CallOpts, mode, context, actualGasCost)
}

// PostOp is a free data retrieval call binding the contract method 0xa9a23409.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost) view returns()
func (_Paymaster *PaymasterCallerSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int) error {
	return _Paymaster.Contract.PostOp(&_Paymaster.CallOpts, mode, context, actualGasCost)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Paymaster *PaymasterCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Paymaster *PaymasterSession) ProxiableUUID() ([32]byte, error) {
	return _Paymaster.Contract.ProxiableUUID(&_Paymaster.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Paymaster *PaymasterCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Paymaster.Contract.ProxiableUUID(&_Paymaster.CallOpts)
}

// SenderNonce is a free data retrieval call binding the contract method 0x9c90b443.
//
// Solidity: function senderNonce(address ) view returns(uint256)
func (_Paymaster *PaymasterCaller) SenderNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "senderNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SenderNonce is a free data retrieval call binding the contract method 0x9c90b443.
//
// Solidity: function senderNonce(address ) view returns(uint256)
func (_Paymaster *PaymasterSession) SenderNonce(arg0 common.Address) (*big.Int, error) {
	return _Paymaster.Contract.SenderNonce(&_Paymaster.CallOpts, arg0)
}

// SenderNonce is a free data retrieval call binding the contract method 0x9c90b443.
//
// Solidity: function senderNonce(address ) view returns(uint256)
func (_Paymaster *PaymasterCallerSession) SenderNonce(arg0 common.Address) (*big.Int, error) {
	return _Paymaster.Contract.SenderNonce(&_Paymaster.CallOpts, arg0)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_Paymaster *PaymasterTransactor) IncrementNonce(opts *bind.TransactOpts, key *big.Int) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "incrementNonce", key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_Paymaster *PaymasterSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _Paymaster.Contract.IncrementNonce(&_Paymaster.TransactOpts, key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_Paymaster *PaymasterTransactorSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _Paymaster.Contract.IncrementNonce(&_Paymaster.TransactOpts, key)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sponsor) returns()
func (_Paymaster *PaymasterTransactor) Initialize(opts *bind.TransactOpts, sponsor common.Address) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "initialize", sponsor)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sponsor) returns()
func (_Paymaster *PaymasterSession) Initialize(sponsor common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.Initialize(&_Paymaster.TransactOpts, sponsor)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sponsor) returns()
func (_Paymaster *PaymasterTransactorSession) Initialize(sponsor common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.Initialize(&_Paymaster.TransactOpts, sponsor)
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

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Paymaster *PaymasterTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Paymaster *PaymasterSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.UpgradeTo(&_Paymaster.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Paymaster *PaymasterTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.UpgradeTo(&_Paymaster.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Paymaster *PaymasterTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Paymaster *PaymasterSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Paymaster.Contract.UpgradeToAndCall(&_Paymaster.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Paymaster *PaymasterTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Paymaster.Contract.UpgradeToAndCall(&_Paymaster.TransactOpts, newImplementation, data)
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

// PaymasterAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Paymaster contract.
type PaymasterAdminChangedIterator struct {
	Event *PaymasterAdminChanged // Event containing the contract specifics and raw log

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
func (it *PaymasterAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymasterAdminChanged)
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
		it.Event = new(PaymasterAdminChanged)
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
func (it *PaymasterAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymasterAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymasterAdminChanged represents a AdminChanged event raised by the Paymaster contract.
type PaymasterAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Paymaster *PaymasterFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*PaymasterAdminChangedIterator, error) {

	logs, sub, err := _Paymaster.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &PaymasterAdminChangedIterator{contract: _Paymaster.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Paymaster *PaymasterFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *PaymasterAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Paymaster.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymasterAdminChanged)
				if err := _Paymaster.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Paymaster *PaymasterFilterer) ParseAdminChanged(log types.Log) (*PaymasterAdminChanged, error) {
	event := new(PaymasterAdminChanged)
	if err := _Paymaster.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymasterBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Paymaster contract.
type PaymasterBeaconUpgradedIterator struct {
	Event *PaymasterBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *PaymasterBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymasterBeaconUpgraded)
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
		it.Event = new(PaymasterBeaconUpgraded)
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
func (it *PaymasterBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymasterBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymasterBeaconUpgraded represents a BeaconUpgraded event raised by the Paymaster contract.
type PaymasterBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Paymaster *PaymasterFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*PaymasterBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Paymaster.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &PaymasterBeaconUpgradedIterator{contract: _Paymaster.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Paymaster *PaymasterFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *PaymasterBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Paymaster.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymasterBeaconUpgraded)
				if err := _Paymaster.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Paymaster *PaymasterFilterer) ParseBeaconUpgraded(log types.Log) (*PaymasterBeaconUpgraded, error) {
	event := new(PaymasterBeaconUpgraded)
	if err := _Paymaster.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymasterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Paymaster contract.
type PaymasterInitializedIterator struct {
	Event *PaymasterInitialized // Event containing the contract specifics and raw log

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
func (it *PaymasterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymasterInitialized)
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
		it.Event = new(PaymasterInitialized)
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
func (it *PaymasterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymasterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymasterInitialized represents a Initialized event raised by the Paymaster contract.
type PaymasterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Paymaster *PaymasterFilterer) FilterInitialized(opts *bind.FilterOpts) (*PaymasterInitializedIterator, error) {

	logs, sub, err := _Paymaster.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PaymasterInitializedIterator{contract: _Paymaster.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Paymaster *PaymasterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PaymasterInitialized) (event.Subscription, error) {

	logs, sub, err := _Paymaster.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymasterInitialized)
				if err := _Paymaster.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Paymaster *PaymasterFilterer) ParseInitialized(log types.Log) (*PaymasterInitialized, error) {
	event := new(PaymasterInitialized)
	if err := _Paymaster.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// PaymasterUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Paymaster contract.
type PaymasterUpgradedIterator struct {
	Event *PaymasterUpgraded // Event containing the contract specifics and raw log

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
func (it *PaymasterUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymasterUpgraded)
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
		it.Event = new(PaymasterUpgraded)
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
func (it *PaymasterUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymasterUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymasterUpgraded represents a Upgraded event raised by the Paymaster contract.
type PaymasterUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Paymaster *PaymasterFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PaymasterUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Paymaster.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PaymasterUpgradedIterator{contract: _Paymaster.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Paymaster *PaymasterFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PaymasterUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Paymaster.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymasterUpgraded)
				if err := _Paymaster.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Paymaster *PaymasterFilterer) ParseUpgraded(log types.Log) (*PaymasterUpgraded, error) {
	event := new(PaymasterUpgraded)
	if err := _Paymaster.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
