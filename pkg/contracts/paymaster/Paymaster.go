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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aSponsor\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sponsor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSponsor\",\"type\":\"address\"}],\"name\":\"updateSponsor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a080604052346100dc57306080526000549060ff8260081c1661008a575060ff8082160361004f575b60405161149990816100e2823960805181818161058601528181610683015261092d0152f35b60ff90811916176000557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160ff8152a138610029565b62461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608490fd5b600080fdfe6040608081526004908136101561001557600080fd5b600091823560e01c80633659cfe6146109065780634f1ef2861461063357806352d1902d14610571578063715018a61461051457806377c93662146104eb5780638da5cb5b146104c257806394e1fc1914610458578063a9a23409146103ff578063c4d66de8146101e4578063c9a54e2b1461019c578063f2fde38b1461010b5763f465c77e146100a557600080fd5b3461010757600319906060368301126100ff5780359167ffffffffffffffff8311610103576101609083360301126100ff576100e19101610ddc565b82518381529283926100f591840190610bbd565b9060208301520390f35b8380fd5b8480fd5b8280fd5b503461010757602036600319011261010757610125610acc565b9161012e610fd2565b6001600160a01b0383161561014a57836101478461102a565b80f35b906020608492519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152fd5b83346101e15760203660031901126101e1576101b6610acc565b6101be610fd2565b60018060a01b03166bffffffffffffffffffffffff60a01b60c954161760c95580f35b80fd5b503461010757602090816003193601126100ff57610200610acc565b9084549160ff8360081c1615928380946103f2575b80156103db575b156103815760ff198116600117875583610370575b5061024b60ff875460081c1661024681610f72565b610f72565b6102543361102a565b60c954918551908682019082821067ffffffffffffffff83111761035d57508652601e81527f6578656375746528616464726573732c75696e743235362c6279746573290000908501528451916102aa83610b19565b60298084527f65786563757465426174636828616464726573735b5d2c75696e743235365b5d868501908152682c62797465735b5d2960b81b948801949094529092206001600160e01b03199092166001600160a01b039091161790831c63ffffffff60c01b1617635b0e93fb60a11b1760c955610326578280f35b7f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989161ff001984541684555160018152a138808280f35b634e487b7160e01b895260419052602488fd5b61ffff191661010117865538610231565b855162461bcd60e51b8152808401869052602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b50303b15801561021c5750600160ff82161461021c565b50600160ff821610610215565b83823461045457606036600319011261045457600381351015610454576024359067ffffffffffffffff908183116100ff57366023840112156100ff5782013590811161010757369101602401116101e15780f35b5080fd5b5034610107576003196060368201126100ff5781359067ffffffffffffffff8211610103576101609082360301126100ff576024359165ffffffffffff9485841684036101e15760443595861686036101e1575091602094916104bb9301610c7c565b9051908152f35b50503461045457816003193601126104545760335490516001600160a01b039091168152602090f35b50503461045457816003193601126104545760c95490516001600160a01b039091168152602090f35b83346101e157806003193601126101e15761052d610fd2565b603380546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b5082346101e157806003193601126101e157507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036105cb57602082516000805160206114448339815191528152f35b6020608492519162461bcd60e51b8352820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152fd5b50908060031936011261010757610648610acc565b9060243567ffffffffffffffff81116101035736602382011215610103576106799036906024818701359101610b73565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169291906106b330851415611073565b6106d06000805160206114448339815191529482865416146110d4565b6106d8610fd2565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff161561071057505050506101479150611316565b82516352d1902d60e01b81526020959394918416919086818981865afa8991816108d3575b5061079457855162461bcd60e51b8152808901889052602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608490fd5b969192959493960361087e57506107aa85611316565b7fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8680a2825115801590610876575b6107e5575b5050505080f35b6108629385808351956107f787610b19565b602787527f416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c86880152660819985a5b195960ca1b858801528581519101845af4913d1561086c573d61085461084b82610b57565b92519283610b35565b81528681943d92013e6113a6565b50388080806107de565b50606092506113a6565b5060016107d9565b825162461bcd60e51b8152908101849052602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608490fd5b9091508781813d83116108ff575b6108eb8183610b35565b810103126108fb57519038610735565b8980fd5b503d6108e1565b509034610107576020806003193601126100ff57610922610acc565b916001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811661095a30821415611073565b6109776000805160206114448339815191529183835416146110d4565b61097f610fd2565b82519161098b83610ae7565b8783527f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff16156109c75750505050506101479150611316565b85939495169084516352d1902d60e01b815286818981865afa899181610a9d575b50610a4757855162461bcd60e51b8152808901889052602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608490fd5b969192959493960361087e5750610a5d85611316565b7fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8680a2825115801590610a96576107e5575050505080f35b50846107d9565b9091508781813d8311610ac5575b610ab58183610b35565b810103126108fb575190386109e8565b503d610aab565b600435906001600160a01b0382168203610ae257565b600080fd5b6020810190811067ffffffffffffffff821117610b0357604052565b634e487b7160e01b600052604160045260246000fd5b6060810190811067ffffffffffffffff821117610b0357604052565b90601f8019910116810190811067ffffffffffffffff821117610b0357604052565b67ffffffffffffffff8111610b0357601f01601f191660200190565b929192610b7f82610b57565b91610b8d6040519384610b35565b829481845281830111610ae2578281602093846000960137010152565b359065ffffffffffff82168203610ae257565b919082519283825260005b848110610be9575050826000602080949584010152601f8019910116010190565b602081830181015184830182015201610bc8565b903590601e1981360301821215610ae2570180359067ffffffffffffffff8211610ae257602001918136038313610ae257565b15610c3757565b60405162461bcd60e51b815260206004820152601e60248201527f4141323720696e76616c69642066756e6374696f6e2073656c6563746f7200006044820152606490fd5b9160608301610c8b8185610bfd565b600492919211610ae2576001600160e01b03199135821690610cd590610cb2831515610c30565b610ccf60c95493858560401b168114908115610d81575b50610c30565b86610bfd565b600493919311610ae25760405192602060018060a01b03978881351682870152013560408501523516606083015260608252608082019382851067ffffffffffffffff861117610b0357610d7b938560405260a084019660c08852610d3e610160860186610bbd565b934660c08701523060e08701521661010085015265ffffffffffff8092166101208501521661014083015203609f1981018352607f190182610b35565b51902090565b9050858560201b161438610cc9565b15610d9757565b60405162461bcd60e51b815260206004820152601760248201527f414133322065787069726564206f72206e6f74206475650000000000000000006044820152606490fd5b90610deb610120830183610bfd565b9283605411610ae257610e0060148301610baa565b91610e0d60348201610baa565b946053190191604083148015610f68575b15610f23578584610e5292610e4d65ffffffffffff80421690610e45818716831015610d90565b841611610d90565b610c7c565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52610ead610ea5603c600020610e9f60018060a01b03948560c95416966054369201610b73565b9061124f565b919091611135565b1603610ee957604051610ebf81610ae7565b600081529260a09190911b65ffffffffffff60a01b1660d09190911b6001600160d01b0319161790565b600190604051610ef881610ae7565b600081529360a09190911b65ffffffffffff60a01b1660d09190911b6001600160d01b031916171790565b60405162461bcd60e51b815260206004820152601d60248201527f4141333520696e76616c6964207369676e6174757265206c656e6774680000006044820152606490fd5b5060418314610e1e565b15610f7957565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b6033546001600160a01b03163303610fe657565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b603380546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b1561107a57565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608490fd5b156110db57565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b6064820152608490fd5b600581101561123957806111465750565b600181036111935760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606490fd5b600281036111e05760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606490fd5b6003146111e957565b60405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608490fd5b634e487b7160e01b600052602160045260246000fd5b90604181511460001461127d57611279916020820151906060604084015193015160001a90611287565b9091565b5050600090600290565b9291907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831161130a5791608094939160ff602094604051948552168484015260408301526060820152600093849182805260015afa156112fd5781516001600160a01b038116156112f7579190565b50600190565b50604051903d90823e3d90fd5b50505050600090600390565b803b1561134b5760008051602061144483398151915280546001600160a01b0319166001600160a01b03909216919091179055565b60405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608490fd5b9192901561140857508151156113ba575090565b3b156113c35790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b82519091501561141b5750805190602001fd5b60405162461bcd60e51b81526020600482015290819061143f906024830190610bbd565b0390fdfe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca26469706673582212200fd6da9b00c35d05fa40f96be6d38b506ab95bb6865ebaf0a2940d7299586cc064736f6c63430008140033",
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

// Sponsor is a free data retrieval call binding the contract method 0x77c93662.
//
// Solidity: function sponsor() view returns(address)
func (_Paymaster *PaymasterCaller) Sponsor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "sponsor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sponsor is a free data retrieval call binding the contract method 0x77c93662.
//
// Solidity: function sponsor() view returns(address)
func (_Paymaster *PaymasterSession) Sponsor() (common.Address, error) {
	return _Paymaster.Contract.Sponsor(&_Paymaster.CallOpts)
}

// Sponsor is a free data retrieval call binding the contract method 0x77c93662.
//
// Solidity: function sponsor() view returns(address)
func (_Paymaster *PaymasterCallerSession) Sponsor() (common.Address, error) {
	return _Paymaster.Contract.Sponsor(&_Paymaster.CallOpts)
}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0xf465c77e.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) view returns(bytes context, uint256 validationData)
func (_Paymaster *PaymasterCaller) ValidatePaymasterUserOp(opts *bind.CallOpts, userOp UserOperation, userOpHash [32]byte, maxCost *big.Int) (struct {
	Context        []byte
	ValidationData *big.Int
}, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "validatePaymasterUserOp", userOp, userOpHash, maxCost)

	outstruct := new(struct {
		Context        []byte
		ValidationData *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Context = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ValidationData = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0xf465c77e.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) view returns(bytes context, uint256 validationData)
func (_Paymaster *PaymasterSession) ValidatePaymasterUserOp(userOp UserOperation, userOpHash [32]byte, maxCost *big.Int) (struct {
	Context        []byte
	ValidationData *big.Int
}, error) {
	return _Paymaster.Contract.ValidatePaymasterUserOp(&_Paymaster.CallOpts, userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0xf465c77e.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) view returns(bytes context, uint256 validationData)
func (_Paymaster *PaymasterCallerSession) ValidatePaymasterUserOp(userOp UserOperation, userOpHash [32]byte, maxCost *big.Int) (struct {
	Context        []byte
	ValidationData *big.Int
}, error) {
	return _Paymaster.Contract.ValidatePaymasterUserOp(&_Paymaster.CallOpts, userOp, userOpHash, maxCost)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address aSponsor) returns()
func (_Paymaster *PaymasterTransactor) Initialize(opts *bind.TransactOpts, aSponsor common.Address) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "initialize", aSponsor)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address aSponsor) returns()
func (_Paymaster *PaymasterSession) Initialize(aSponsor common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.Initialize(&_Paymaster.TransactOpts, aSponsor)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address aSponsor) returns()
func (_Paymaster *PaymasterTransactorSession) Initialize(aSponsor common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.Initialize(&_Paymaster.TransactOpts, aSponsor)
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

// UpdateSponsor is a paid mutator transaction binding the contract method 0xc9a54e2b.
//
// Solidity: function updateSponsor(address newSponsor) returns()
func (_Paymaster *PaymasterTransactor) UpdateSponsor(opts *bind.TransactOpts, newSponsor common.Address) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "updateSponsor", newSponsor)
}

// UpdateSponsor is a paid mutator transaction binding the contract method 0xc9a54e2b.
//
// Solidity: function updateSponsor(address newSponsor) returns()
func (_Paymaster *PaymasterSession) UpdateSponsor(newSponsor common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.UpdateSponsor(&_Paymaster.TransactOpts, newSponsor)
}

// UpdateSponsor is a paid mutator transaction binding the contract method 0xc9a54e2b.
//
// Solidity: function updateSponsor(address newSponsor) returns()
func (_Paymaster *PaymasterTransactorSession) UpdateSponsor(newSponsor common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.UpdateSponsor(&_Paymaster.TransactOpts, newSponsor)
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
