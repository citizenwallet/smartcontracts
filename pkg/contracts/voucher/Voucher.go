// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voucher

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

// VoucherMetaData contains all meta data concerning the Voucher contract.
var VoucherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"},{\"internalType\":\"contractITokenEntryPoint\",\"name\":\"aTokenEntryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"VoucherInitialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"func\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dest\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"func\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"recoverOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenEntryPoint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e0604052306080523480156200001557600080fd5b50604051620026b8380380620026b8833981016040819052620000389162000137565b6001600160a01b0380831660a052811660c052620000556200005d565b505062000176565b600054610100900460ff1615620000ca5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146200011c576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200013457600080fd5b50565b600080604083850312156200014b57600080fd5b825162000158816200011e565b60208401519092506200016b816200011e565b809150509250929050565b60805160a05160c0516124b062000208600039600081816103a3015281816104e301526110d50152600081816103500152818161084f015281816108d701528181610c0b01528181610ca301528181610d07015281816110880152818161137001526116a1015260008181610770015281816107b001528181610960015281816109a00152610a1801526124b06000f3fe6080604052600436106101435760003560e01c806352d1902d116100b6578063bc197c811161006f578063bc197c81146103c7578063c399ec88146103f6578063d087d2881461040b578063f23a6e6114610420578063f2fde38b1461044d578063fff35b721461046d57600080fd5b806352d1902d146102d95780638129fc1c146102ee5780638da5cb5b14610303578063b0d691fe14610341578063b61d27f614610374578063bb5032061461039457600080fd5b806318dfb3c71161010857806318dfb3c7146102305780633659cfe6146102505780633a871cdd146102705780634a58db191461029e5780634d44560d146102a65780634f1ef286146102c657600080fd5b806223de291461014f57806301ffc9a7146101765780630b699abe146101ab578063150b7a02146101cb5780631626ba7e1461021057600080fd5b3661014a57005b600080fd5b34801561015b57600080fd5b5061017461016a366004611ca5565b5050505050505050565b005b34801561018257600080fd5b50610196610191366004611d55565b61048d565b60405190151581526020015b60405180910390f35b3480156101b757600080fd5b506101746101c6366004611d7f565b6104df565b3480156101d757600080fd5b506101f76101e6366004611d9c565b630a85bd0160e11b95945050505050565b6040516001600160e01b031990911681526020016101a2565b34801561021c57600080fd5b506101f761022b366004611e0e565b6105e3565b34801561023c57600080fd5b5061017461024b366004611e9d565b610669565b34801561025c57600080fd5b5061017461026b366004611d7f565b610766565b34801561027c57600080fd5b5061029061028b366004611f21565b61082e565b6040519081526020016101a2565b61017461084d565b3480156102b257600080fd5b506101746102c1366004611f6e565b6108cd565b6101746102d4366004611fb0565b610956565b3480156102e557600080fd5b50610290610a0b565b3480156102fa57600080fd5b50610174610abe565b34801561030f57600080fd5b50600054610329906201000090046001600160a01b031681565b6040516001600160a01b0390911681526020016101a2565b34801561034d57600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610329565b34801561038057600080fd5b5061017461038f366004612073565b610bd9565b3480156103a057600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610329565b3480156103d357600080fd5b506101f76103e23660046120c2565b63bc197c8160e01b98975050505050505050565b34801561040257600080fd5b50610290610beb565b34801561041757600080fd5b50610290610c7c565b34801561042c57600080fd5b506101f761043b36600461215f565b63f23a6e6160e01b9695505050505050565b34801561045957600080fd5b50610174610468366004611d7f565b610cd2565b34801561047957600080fd5b506101966104883660046121da565b610ce3565b60006001600160e01b03198216630a85bd0160e11b14806104be57506001600160e01b03198216630271189760e51b145b806104d957506001600160e01b031982166301ffc9a760e01b145b92915050565b60007f00000000000000000000000000000000000000000000000000000000000000009050806001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610542573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610566919061221e565b6001600160a01b0316336001600160a01b0316146105d65760405162461bcd60e51b815260206004820152602260248201527f4f776e61626c653a206e6f7420546f6b656e456e747279506f696e74206f776e60448201526132b960f11b60648201526084015b60405180910390fd5b6105df82610d92565b5050565b6000806106268585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e2192505050565b6000549091506001600160a01b03620100009091048116908216036106555750630b135d3f60e11b9050610662565b506001600160e01b031990505b9392505050565b61067161107d565b808381146106b75760405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b60448201526064016105cd565b60005b8181101561075e5761074c8686838181106106d7576106d761223b565b90506020020160208101906106ec9190611d7f565b60008686858181106107005761070061223b565b90506020028101906107129190612251565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061116192505050565b80610756816122ad565b9150506106ba565b505050505050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036107ae5760405162461bcd60e51b81526004016105cd906122c6565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166107e06111d1565b6001600160a01b0316146108065760405162461bcd60e51b81526004016105cd90612312565b61080f816111ed565b6040805160008082526020820190925261082b918391906111f5565b50565b6000610838611365565b61084284846113dd565b905061066282611489565b7f000000000000000000000000000000000000000000000000000000000000000060405163b760faf960e01b81523060048201526001600160a01b03919091169063b760faf99034906024016000604051808303818588803b1580156108b257600080fd5b505af11580156108c6573d6000803e3d6000fd5b5050505050565b6108d56114d6565b7f000000000000000000000000000000000000000000000000000000000000000060405163040b850f60e31b81526001600160a01b03848116600483015260248201849052919091169063205c287890604401600060405180830381600087803b15801561094257600080fd5b505af115801561075e573d6000803e3d6000fd5b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361099e5760405162461bcd60e51b81526004016105cd906122c6565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166109d06111d1565b6001600160a01b0316146109f65760405162461bcd60e51b81526004016105cd90612312565b6109ff826111ed565b6105df828260016111f5565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610aab5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c000000000000000060648201526084016105cd565b5060008051602061241483398151915290565b600054610100900460ff1615808015610ade5750600054600160ff909116105b80610af85750303b158015610af8575060005460ff166001145b610b5b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016105cd565b6000805460ff191660011790558015610b7e576000805461ff0019166101001790555b610b886000611559565b610b9133611672565b801561082b576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b610be5848484846116ea565b50505050565b6040516370a0823160e01b81523060048201526000906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a08231906024015b602060405180830381865afa158015610c53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c77919061235e565b905090565b604051631aab3f0d60e11b8152306004820152600060248201819052906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906335567e1a90604401610c36565b610cda6114d6565b61082b81610d92565b600080610cf084846113dd565b1590508015610662576000602085013560401c90507f0000000000000000000000000000000000000000000000000000000000000000604051630bd28e3b60e01b81526001600160c01b03831660048201526001600160a01b039190911690630bd28e3b90602401600060405180830381600087803b158015610d7257600080fd5b505af1158015610d86573d6000803e3d6000fd5b50505050509392505050565b6001600160a01b038116610df75760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016105cd565b600080546001600160a01b03909216620100000262010000600160b01b0319909216919091179055565b60008151604114610e885760405162461bcd60e51b815260206004820152603a602482015260008051602061245b83398151915260448201527f3a20696e76616c6964207369676e6174757265206c656e67746800000000000060648201526084016105cd565b600082604081518110610e9d57610e9d61223b565b016020015160f81c90506000610eb38482611733565b90506000610ec2856020611733565b90506fa2a8918ca85bafe22016d0b997e4df60600160ff1b03811115610f3e5760405162461bcd60e51b815260206004820152603d602482015260008051602061245b83398151915260448201527f3a20696e76616c6964207369676e6174757265202773272076616c756500000060648201526084016105cd565b8260ff16601b14158015610f5657508260ff16601c14155b15610fb75760405162461bcd60e51b815260206004820152603d602482015260008051602061245b83398151915260448201527f3a20696e76616c6964207369676e6174757265202776272076616c756500000060648201526084016105cd565b60408051600081526020810180835288905260ff851691810191909152606081018390526080810182905260019060a0016020604051602081039080840390855afa15801561100a573d6000803e3d6000fd5b5050604051601f1901519450506001600160a01b0384166110745760405162461bcd60e51b8152602060048201526030602482015260008051602061245b83398151915260448201526f1d1024a72b20a624a22fa9a4a3a722a960811b60648201526084016105cd565b50505092915050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806110c457506000546201000090046001600160a01b031633145b806110f75750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b61115f5760405162461bcd60e51b815260206004820152603360248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e74604482015272081bdc88151bdad95b915b9d1c9e541bda5b9d606a1b60648201526084016105cd565b565b600080846001600160a01b0316848460405161117d919061239b565b60006040518083038185875af1925050503d80600081146111ba576040519150601f19603f3d011682016040523d82523d6000602084013e6111bf565b606091505b5091509150816108c657805160208201fd5b600080516020612414833981519152546001600160a01b031690565b61082b6114d6565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff161561122d5761122883611799565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611287575060408051601f3d908101601f191682019092526112849181019061235e565b60015b6112ea5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b60648201526084016105cd565b60008051602061241483398151915281146113595760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b60648201526084016105cd565b50611228838383611835565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461115f5760405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e740000000060448201526064016105cd565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c829052603c812061145a61141d610140860186612251565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250859392505061185a9050565b6000546201000090046001600160a01b0390811691161461147f5760019150506104d9565b5060009392505050565b801561082b57604051600090339060001990849084818181858888f193505050503d80600081146108c6576040519150601f19603f3d011682016040523d82523d6000602084013e6108c6565b6000546201000090046001600160a01b03163314806114f457503033145b61115f5760405162461bcd60e51b815260206004820152603060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526f081bdc881d1a194818dbdb9d1c9858dd60821b60648201526084016105cd565b600054610100900460ff166115c45760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b60648201526084016105cd565b600160006115d06111d1565b6001600160a01b0316815260208101919091526040016000205460ff161561163a5760405162461bcd60e51b815260206004820152601960248201527f566f75636865723a20616c7265616479206d696772617465640000000000000060448201526064016105cd565b60018060006116476111d1565b6001600160a01b031681526020810191909152604001600020805460ff191691151591909117905550565b6000805462010000600160b01b031916620100006001600160a01b0384811691820292909217835560405190927f0000000000000000000000000000000000000000000000000000000000000000909216917f706f6c731664f703c9b00eb42d1b09183bc2ac9772b13f96c2c0975df783d8eb91a350565b6116f261107d565b610be5848484848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061116192505050565b60006117408260206123b7565b835110156117905760405162461bcd60e51b815260206004820181905260248201527f72656164427974657333323a20696e76616c69642064617461206c656e67746860448201526064016105cd565b50016020015190565b6001600160a01b0381163b6118065760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016105cd565b60008051602061241483398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b61183e8361187e565b60008251118061184b5750805b1561122857610be583836118be565b600080600061186985856118e3565b9150915061187681611928565b509392505050565b61188781611799565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060610662838360405180606001604052806027815260200161243460279139611a72565b60008082516041036119195760208301516040840151606085015160001a61190d87828585611aea565b94509450505050611921565b506000905060025b9250929050565b600081600481111561193c5761193c6123ca565b036119445750565b6001816004811115611958576119586123ca565b036119a55760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016105cd565b60028160048111156119b9576119b96123ca565b03611a065760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016105cd565b6003816004811115611a1a57611a1a6123ca565b0361082b5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016105cd565b6060600080856001600160a01b031685604051611a8f919061239b565b600060405180830381855af49150503d8060008114611aca576040519150601f19603f3d011682016040523d82523d6000602084013e611acf565b606091505b5091509150611ae086838387611ba4565b9695505050505050565b6000806fa2a8918ca85bafe22016d0b997e4df60600160ff1b03831115611b175750600090506003611b9b565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611b6b573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611b9457600060019250925050611b9b565b9150600090505b94509492505050565b60608315611c13578251600003611c0c576001600160a01b0385163b611c0c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016105cd565b5081611c1d565b611c1d8383611c25565b949350505050565b815115611c355781518083602001fd5b8060405162461bcd60e51b81526004016105cd91906123e0565b6001600160a01b038116811461082b57600080fd5b60008083601f840112611c7657600080fd5b5081356001600160401b03811115611c8d57600080fd5b60208301915083602082850101111561192157600080fd5b60008060008060008060008060c0898b031215611cc157600080fd5b8835611ccc81611c4f565b97506020890135611cdc81611c4f565b96506040890135611cec81611c4f565b95506060890135945060808901356001600160401b0380821115611d0f57600080fd5b611d1b8c838d01611c64565b909650945060a08b0135915080821115611d3457600080fd5b50611d418b828c01611c64565b999c989b5096995094979396929594505050565b600060208284031215611d6757600080fd5b81356001600160e01b03198116811461066257600080fd5b600060208284031215611d9157600080fd5b813561066281611c4f565b600080600080600060808688031215611db457600080fd5b8535611dbf81611c4f565b94506020860135611dcf81611c4f565b93506040860135925060608601356001600160401b03811115611df157600080fd5b611dfd88828901611c64565b969995985093965092949392505050565b600080600060408486031215611e2357600080fd5b8335925060208401356001600160401b03811115611e4057600080fd5b611e4c86828701611c64565b9497909650939450505050565b60008083601f840112611e6b57600080fd5b5081356001600160401b03811115611e8257600080fd5b6020830191508360208260051b850101111561192157600080fd5b60008060008060408587031215611eb357600080fd5b84356001600160401b0380821115611eca57600080fd5b611ed688838901611e59565b90965094506020870135915080821115611eef57600080fd5b50611efc87828801611e59565b95989497509550505050565b60006101608284031215611f1b57600080fd5b50919050565b600080600060608486031215611f3657600080fd5b83356001600160401b03811115611f4c57600080fd5b611f5886828701611f08565b9660208601359650604090950135949350505050565b60008060408385031215611f8157600080fd5b8235611f8c81611c4f565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215611fc357600080fd5b8235611fce81611c4f565b915060208301356001600160401b0380821115611fea57600080fd5b818501915085601f830112611ffe57600080fd5b81358181111561201057612010611f9a565b604051601f8201601f19908116603f0116810190838211818310171561203857612038611f9a565b8160405282815288602084870101111561205157600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b6000806000806060858703121561208957600080fd5b843561209481611c4f565b93506020850135925060408501356001600160401b038111156120b657600080fd5b611efc87828801611c64565b60008060008060008060008060a0898b0312156120de57600080fd5b88356120e981611c4f565b975060208901356120f981611c4f565b965060408901356001600160401b038082111561211557600080fd5b6121218c838d01611e59565b909850965060608b013591508082111561213a57600080fd5b6121468c838d01611e59565b909650945060808b0135915080821115611d3457600080fd5b60008060008060008060a0878903121561217857600080fd5b863561218381611c4f565b9550602087013561219381611c4f565b9450604087013593506060870135925060808701356001600160401b038111156121bc57600080fd5b6121c889828a01611c64565b979a9699509497509295939492505050565b600080604083850312156121ed57600080fd5b82356001600160401b0381111561220357600080fd5b61220f85828601611f08565b95602094909401359450505050565b60006020828403121561223057600080fd5b815161066281611c4f565b634e487b7160e01b600052603260045260246000fd5b6000808335601e1984360301811261226857600080fd5b8301803591506001600160401b0382111561228257600080fd5b60200191503681900382131561192157600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016122bf576122bf612297565b5060010190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b60006020828403121561237057600080fd5b5051919050565b60005b8381101561239257818101518382015260200161237a565b50506000910152565b600082516123ad818460208701612377565b9190910192915050565b808201808211156104d9576104d9612297565b634e487b7160e01b600052602160045260246000fd5b60208152600082518060208401526123ff816040850160208701612377565b601f01601f1916919091016040019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645369676e617475726556616c696461746f72237265636f7665725369676e6572a26469706673582212203e3881ae2194bd4768cf3ed83856588bee9667eb98c28b3e992eb3127131e38a64736f6c63430008140033",
}

// VoucherABI is the input ABI used to generate the binding from.
// Deprecated: Use VoucherMetaData.ABI instead.
var VoucherABI = VoucherMetaData.ABI

// VoucherBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VoucherMetaData.Bin instead.
var VoucherBin = VoucherMetaData.Bin

// DeployVoucher deploys a new Ethereum contract, binding an instance of Voucher to it.
func DeployVoucher(auth *bind.TransactOpts, backend bind.ContractBackend, anEntryPoint common.Address, aTokenEntryPoint common.Address) (common.Address, *types.Transaction, *Voucher, error) {
	parsed, err := VoucherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VoucherBin), backend, anEntryPoint, aTokenEntryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Voucher{VoucherCaller: VoucherCaller{contract: contract}, VoucherTransactor: VoucherTransactor{contract: contract}, VoucherFilterer: VoucherFilterer{contract: contract}}, nil
}

// Voucher is an auto generated Go binding around an Ethereum contract.
type Voucher struct {
	VoucherCaller     // Read-only binding to the contract
	VoucherTransactor // Write-only binding to the contract
	VoucherFilterer   // Log filterer for contract events
}

// VoucherCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoucherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoucherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoucherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoucherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoucherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoucherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoucherSession struct {
	Contract     *Voucher          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoucherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoucherCallerSession struct {
	Contract *VoucherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VoucherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoucherTransactorSession struct {
	Contract     *VoucherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VoucherRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoucherRaw struct {
	Contract *Voucher // Generic contract binding to access the raw methods on
}

// VoucherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoucherCallerRaw struct {
	Contract *VoucherCaller // Generic read-only contract binding to access the raw methods on
}

// VoucherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoucherTransactorRaw struct {
	Contract *VoucherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoucher creates a new instance of Voucher, bound to a specific deployed contract.
func NewVoucher(address common.Address, backend bind.ContractBackend) (*Voucher, error) {
	contract, err := bindVoucher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voucher{VoucherCaller: VoucherCaller{contract: contract}, VoucherTransactor: VoucherTransactor{contract: contract}, VoucherFilterer: VoucherFilterer{contract: contract}}, nil
}

// NewVoucherCaller creates a new read-only instance of Voucher, bound to a specific deployed contract.
func NewVoucherCaller(address common.Address, caller bind.ContractCaller) (*VoucherCaller, error) {
	contract, err := bindVoucher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoucherCaller{contract: contract}, nil
}

// NewVoucherTransactor creates a new write-only instance of Voucher, bound to a specific deployed contract.
func NewVoucherTransactor(address common.Address, transactor bind.ContractTransactor) (*VoucherTransactor, error) {
	contract, err := bindVoucher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoucherTransactor{contract: contract}, nil
}

// NewVoucherFilterer creates a new log filterer instance of Voucher, bound to a specific deployed contract.
func NewVoucherFilterer(address common.Address, filterer bind.ContractFilterer) (*VoucherFilterer, error) {
	contract, err := bindVoucher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoucherFilterer{contract: contract}, nil
}

// bindVoucher binds a generic wrapper to an already deployed contract.
func bindVoucher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoucherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voucher *VoucherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voucher.Contract.VoucherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voucher *VoucherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voucher.Contract.VoucherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voucher *VoucherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voucher.Contract.VoucherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voucher *VoucherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voucher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voucher *VoucherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voucher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voucher *VoucherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voucher.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Voucher *VoucherCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Voucher *VoucherSession) EntryPoint() (common.Address, error) {
	return _Voucher.Contract.EntryPoint(&_Voucher.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Voucher *VoucherCallerSession) EntryPoint() (common.Address, error) {
	return _Voucher.Contract.EntryPoint(&_Voucher.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Voucher *VoucherCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Voucher *VoucherSession) GetDeposit() (*big.Int, error) {
	return _Voucher.Contract.GetDeposit(&_Voucher.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Voucher *VoucherCallerSession) GetDeposit() (*big.Int, error) {
	return _Voucher.Contract.GetDeposit(&_Voucher.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Voucher *VoucherCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Voucher *VoucherSession) GetNonce() (*big.Int, error) {
	return _Voucher.Contract.GetNonce(&_Voucher.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Voucher *VoucherCallerSession) GetNonce() (*big.Int, error) {
	return _Voucher.Contract.GetNonce(&_Voucher.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_Voucher *VoucherCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_Voucher *VoucherSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _Voucher.Contract.IsValidSignature(&_Voucher.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_Voucher *VoucherCallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _Voucher.Contract.IsValidSignature(&_Voucher.CallOpts, _hash, _signature)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Voucher *VoucherCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Voucher *VoucherSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Voucher.Contract.OnERC1155BatchReceived(&_Voucher.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Voucher *VoucherCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Voucher.Contract.OnERC1155BatchReceived(&_Voucher.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Voucher *VoucherCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Voucher *VoucherSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Voucher.Contract.OnERC1155Received(&_Voucher.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Voucher *VoucherCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Voucher.Contract.OnERC1155Received(&_Voucher.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Voucher *VoucherCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Voucher *VoucherSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Voucher.Contract.OnERC721Received(&_Voucher.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Voucher *VoucherCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Voucher.Contract.OnERC721Received(&_Voucher.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voucher *VoucherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voucher *VoucherSession) Owner() (common.Address, error) {
	return _Voucher.Contract.Owner(&_Voucher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voucher *VoucherCallerSession) Owner() (common.Address, error) {
	return _Voucher.Contract.Owner(&_Voucher.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Voucher *VoucherCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Voucher *VoucherSession) ProxiableUUID() ([32]byte, error) {
	return _Voucher.Contract.ProxiableUUID(&_Voucher.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Voucher *VoucherCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Voucher.Contract.ProxiableUUID(&_Voucher.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Voucher *VoucherCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Voucher *VoucherSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Voucher.Contract.SupportsInterface(&_Voucher.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Voucher *VoucherCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Voucher.Contract.SupportsInterface(&_Voucher.CallOpts, interfaceId)
}

// TokenEntryPoint is a free data retrieval call binding the contract method 0xbb503206.
//
// Solidity: function tokenEntryPoint() view returns(address)
func (_Voucher *VoucherCaller) TokenEntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "tokenEntryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenEntryPoint is a free data retrieval call binding the contract method 0xbb503206.
//
// Solidity: function tokenEntryPoint() view returns(address)
func (_Voucher *VoucherSession) TokenEntryPoint() (common.Address, error) {
	return _Voucher.Contract.TokenEntryPoint(&_Voucher.CallOpts)
}

// TokenEntryPoint is a free data retrieval call binding the contract method 0xbb503206.
//
// Solidity: function tokenEntryPoint() view returns(address)
func (_Voucher *VoucherCallerSession) TokenEntryPoint() (common.Address, error) {
	return _Voucher.Contract.TokenEntryPoint(&_Voucher.CallOpts)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Voucher *VoucherCaller) TokensReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "tokensReceived", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Voucher *VoucherSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _Voucher.Contract.TokensReceived(&_Voucher.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Voucher *VoucherCallerSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _Voucher.Contract.TokensReceived(&_Voucher.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Voucher *VoucherTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Voucher *VoucherSession) AddDeposit() (*types.Transaction, error) {
	return _Voucher.Contract.AddDeposit(&_Voucher.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Voucher *VoucherTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _Voucher.Contract.AddDeposit(&_Voucher.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Voucher *VoucherTransactor) Execute(opts *bind.TransactOpts, dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "execute", dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Voucher *VoucherSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Voucher.Contract.Execute(&_Voucher.TransactOpts, dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Voucher *VoucherTransactorSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Voucher.Contract.Execute(&_Voucher.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x18dfb3c7.
//
// Solidity: function executeBatch(address[] dest, bytes[] func) returns()
func (_Voucher *VoucherTransactor) ExecuteBatch(opts *bind.TransactOpts, dest []common.Address, arg1 [][]byte) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "executeBatch", dest, arg1)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x18dfb3c7.
//
// Solidity: function executeBatch(address[] dest, bytes[] func) returns()
func (_Voucher *VoucherSession) ExecuteBatch(dest []common.Address, arg1 [][]byte) (*types.Transaction, error) {
	return _Voucher.Contract.ExecuteBatch(&_Voucher.TransactOpts, dest, arg1)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x18dfb3c7.
//
// Solidity: function executeBatch(address[] dest, bytes[] func) returns()
func (_Voucher *VoucherTransactorSession) ExecuteBatch(dest []common.Address, arg1 [][]byte) (*types.Transaction, error) {
	return _Voucher.Contract.ExecuteBatch(&_Voucher.TransactOpts, dest, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Voucher *VoucherTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Voucher *VoucherSession) Initialize() (*types.Transaction, error) {
	return _Voucher.Contract.Initialize(&_Voucher.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Voucher *VoucherTransactorSession) Initialize() (*types.Transaction, error) {
	return _Voucher.Contract.Initialize(&_Voucher.TransactOpts)
}

// RecoverOwnership is a paid mutator transaction binding the contract method 0x0b699abe.
//
// Solidity: function recoverOwnership(address newOwner) returns()
func (_Voucher *VoucherTransactor) RecoverOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "recoverOwnership", newOwner)
}

// RecoverOwnership is a paid mutator transaction binding the contract method 0x0b699abe.
//
// Solidity: function recoverOwnership(address newOwner) returns()
func (_Voucher *VoucherSession) RecoverOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Voucher.Contract.RecoverOwnership(&_Voucher.TransactOpts, newOwner)
}

// RecoverOwnership is a paid mutator transaction binding the contract method 0x0b699abe.
//
// Solidity: function recoverOwnership(address newOwner) returns()
func (_Voucher *VoucherTransactorSession) RecoverOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Voucher.Contract.RecoverOwnership(&_Voucher.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Voucher *VoucherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Voucher *VoucherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Voucher.Contract.TransferOwnership(&_Voucher.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Voucher *VoucherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Voucher.Contract.TransferOwnership(&_Voucher.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Voucher *VoucherTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Voucher *VoucherSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Voucher.Contract.UpgradeTo(&_Voucher.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Voucher *VoucherTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Voucher.Contract.UpgradeTo(&_Voucher.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Voucher *VoucherTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Voucher *VoucherSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Voucher.Contract.UpgradeToAndCall(&_Voucher.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Voucher *VoucherTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Voucher.Contract.UpgradeToAndCall(&_Voucher.TransactOpts, newImplementation, data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Voucher *VoucherTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Voucher *VoucherSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Voucher.Contract.ValidateUserOp(&_Voucher.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Voucher *VoucherTransactorSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Voucher.Contract.ValidateUserOp(&_Voucher.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp0 is a paid mutator transaction binding the contract method 0xfff35b72.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash) returns(bool)
func (_Voucher *VoucherTransactor) ValidateUserOp0(opts *bind.TransactOpts, userOp UserOperation, userOpHash [32]byte) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "validateUserOp0", userOp, userOpHash)
}

// ValidateUserOp0 is a paid mutator transaction binding the contract method 0xfff35b72.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash) returns(bool)
func (_Voucher *VoucherSession) ValidateUserOp0(userOp UserOperation, userOpHash [32]byte) (*types.Transaction, error) {
	return _Voucher.Contract.ValidateUserOp0(&_Voucher.TransactOpts, userOp, userOpHash)
}

// ValidateUserOp0 is a paid mutator transaction binding the contract method 0xfff35b72.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash) returns(bool)
func (_Voucher *VoucherTransactorSession) ValidateUserOp0(userOp UserOperation, userOpHash [32]byte) (*types.Transaction, error) {
	return _Voucher.Contract.ValidateUserOp0(&_Voucher.TransactOpts, userOp, userOpHash)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Voucher *VoucherTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Voucher *VoucherSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Voucher.Contract.WithdrawDepositTo(&_Voucher.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Voucher *VoucherTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Voucher.Contract.WithdrawDepositTo(&_Voucher.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Voucher *VoucherTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voucher.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Voucher *VoucherSession) Receive() (*types.Transaction, error) {
	return _Voucher.Contract.Receive(&_Voucher.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Voucher *VoucherTransactorSession) Receive() (*types.Transaction, error) {
	return _Voucher.Contract.Receive(&_Voucher.TransactOpts)
}

// VoucherAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Voucher contract.
type VoucherAdminChangedIterator struct {
	Event *VoucherAdminChanged // Event containing the contract specifics and raw log

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
func (it *VoucherAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoucherAdminChanged)
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
		it.Event = new(VoucherAdminChanged)
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
func (it *VoucherAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoucherAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoucherAdminChanged represents a AdminChanged event raised by the Voucher contract.
type VoucherAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Voucher *VoucherFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*VoucherAdminChangedIterator, error) {

	logs, sub, err := _Voucher.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &VoucherAdminChangedIterator{contract: _Voucher.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Voucher *VoucherFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *VoucherAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Voucher.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoucherAdminChanged)
				if err := _Voucher.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Voucher *VoucherFilterer) ParseAdminChanged(log types.Log) (*VoucherAdminChanged, error) {
	event := new(VoucherAdminChanged)
	if err := _Voucher.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoucherBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Voucher contract.
type VoucherBeaconUpgradedIterator struct {
	Event *VoucherBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *VoucherBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoucherBeaconUpgraded)
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
		it.Event = new(VoucherBeaconUpgraded)
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
func (it *VoucherBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoucherBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoucherBeaconUpgraded represents a BeaconUpgraded event raised by the Voucher contract.
type VoucherBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Voucher *VoucherFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*VoucherBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Voucher.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &VoucherBeaconUpgradedIterator{contract: _Voucher.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Voucher *VoucherFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *VoucherBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Voucher.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoucherBeaconUpgraded)
				if err := _Voucher.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Voucher *VoucherFilterer) ParseBeaconUpgraded(log types.Log) (*VoucherBeaconUpgraded, error) {
	event := new(VoucherBeaconUpgraded)
	if err := _Voucher.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoucherInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Voucher contract.
type VoucherInitializedIterator struct {
	Event *VoucherInitialized // Event containing the contract specifics and raw log

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
func (it *VoucherInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoucherInitialized)
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
		it.Event = new(VoucherInitialized)
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
func (it *VoucherInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoucherInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoucherInitialized represents a Initialized event raised by the Voucher contract.
type VoucherInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Voucher *VoucherFilterer) FilterInitialized(opts *bind.FilterOpts) (*VoucherInitializedIterator, error) {

	logs, sub, err := _Voucher.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VoucherInitializedIterator{contract: _Voucher.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Voucher *VoucherFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VoucherInitialized) (event.Subscription, error) {

	logs, sub, err := _Voucher.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoucherInitialized)
				if err := _Voucher.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Voucher *VoucherFilterer) ParseInitialized(log types.Log) (*VoucherInitialized, error) {
	event := new(VoucherInitialized)
	if err := _Voucher.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoucherUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Voucher contract.
type VoucherUpgradedIterator struct {
	Event *VoucherUpgraded // Event containing the contract specifics and raw log

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
func (it *VoucherUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoucherUpgraded)
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
		it.Event = new(VoucherUpgraded)
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
func (it *VoucherUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoucherUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoucherUpgraded represents a Upgraded event raised by the Voucher contract.
type VoucherUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Voucher *VoucherFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VoucherUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Voucher.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VoucherUpgradedIterator{contract: _Voucher.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Voucher *VoucherFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VoucherUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Voucher.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoucherUpgraded)
				if err := _Voucher.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Voucher *VoucherFilterer) ParseUpgraded(log types.Log) (*VoucherUpgraded, error) {
	event := new(VoucherUpgraded)
	if err := _Voucher.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoucherVoucherInitializedIterator is returned from FilterVoucherInitialized and is used to iterate over the raw logs and unpacked data for VoucherInitialized events raised by the Voucher contract.
type VoucherVoucherInitializedIterator struct {
	Event *VoucherVoucherInitialized // Event containing the contract specifics and raw log

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
func (it *VoucherVoucherInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoucherVoucherInitialized)
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
		it.Event = new(VoucherVoucherInitialized)
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
func (it *VoucherVoucherInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoucherVoucherInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoucherVoucherInitialized represents a VoucherInitialized event raised by the Voucher contract.
type VoucherVoucherInitialized struct {
	EntryPoint common.Address
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoucherInitialized is a free log retrieval operation binding the contract event 0x706f6c731664f703c9b00eb42d1b09183bc2ac9772b13f96c2c0975df783d8eb.
//
// Solidity: event VoucherInitialized(address indexed entryPoint, address indexed owner)
func (_Voucher *VoucherFilterer) FilterVoucherInitialized(opts *bind.FilterOpts, entryPoint []common.Address, owner []common.Address) (*VoucherVoucherInitializedIterator, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Voucher.contract.FilterLogs(opts, "VoucherInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &VoucherVoucherInitializedIterator{contract: _Voucher.contract, event: "VoucherInitialized", logs: logs, sub: sub}, nil
}

// WatchVoucherInitialized is a free log subscription operation binding the contract event 0x706f6c731664f703c9b00eb42d1b09183bc2ac9772b13f96c2c0975df783d8eb.
//
// Solidity: event VoucherInitialized(address indexed entryPoint, address indexed owner)
func (_Voucher *VoucherFilterer) WatchVoucherInitialized(opts *bind.WatchOpts, sink chan<- *VoucherVoucherInitialized, entryPoint []common.Address, owner []common.Address) (event.Subscription, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Voucher.contract.WatchLogs(opts, "VoucherInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoucherVoucherInitialized)
				if err := _Voucher.contract.UnpackLog(event, "VoucherInitialized", log); err != nil {
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

// ParseVoucherInitialized is a log parse operation binding the contract event 0x706f6c731664f703c9b00eb42d1b09183bc2ac9772b13f96c2c0975df783d8eb.
//
// Solidity: event VoucherInitialized(address indexed entryPoint, address indexed owner)
func (_Voucher *VoucherFilterer) ParseVoucherInitialized(log types.Log) (*VoucherVoucherInitialized, error) {
	event := new(VoucherVoucherInitialized)
	if err := _Voucher.contract.UnpackLog(event, "VoucherInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}