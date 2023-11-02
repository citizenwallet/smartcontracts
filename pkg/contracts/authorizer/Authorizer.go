// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package authorizer

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

// AuthorizerMetaData contains all meta data concerning the Authorizer contract.
var AuthorizerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"PaymasterSponsorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"_userOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"senderNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sponsor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSponsor\",\"type\":\"address\"}],\"name\":\"updatePaymasterSponsor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100e1565b600254610100900460ff161561008e5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60025460ff908116146100df576002805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6080516120d962000119600039600081816103f501528181610435015281816105160152818161055601526105e901526120d96000f3fe6080604052600436106100dd5760003560e01c806377c936621161007f5780639c90b443116100595780639c90b44314610240578063ba4fdba01461026d578063c4d66de81461028d578063f2fde38b146102ad57600080fd5b806377c93662146101ca5780638da5cb5b1461020257806394e1fc191461022057600080fd5b80634564cc2e116100bb5780634564cc2e1461016d5780634f1ef2861461018d57806352d1902d146101a0578063715018a6146101b557600080fd5b80631fad948c146100e25780632d0335ab146101045780633659cfe61461014d575b600080fd5b3480156100ee57600080fd5b506101026100fd366004611a8d565b6102cd565b005b34801561011057600080fd5b5061013a61011f366004611b13565b6001600160a01b031660009081526001602052604090205490565b6040519081526020015b60405180910390f35b34801561015957600080fd5b50610102610168366004611b13565b6103eb565b34801561017957600080fd5b5061013a610188366004611b49565b6104ca565b61010261019b366004611b94565b61050c565b3480156101ac57600080fd5b5061013a6105dc565b3480156101c157600080fd5b5061010261068f565b3480156101d657600080fd5b506000546101ea906001600160a01b031681565b6040516001600160a01b039091168152602001610144565b34801561020e57600080fd5b506067546001600160a01b03166101ea565b34801561022c57600080fd5b5061013a61023b366004611c73565b6106a3565b34801561024c57600080fd5b5061013a61025b366004611b13565b60016020526000908152604090205481565b34801561027957600080fd5b50610102610288366004611b13565b610700565b34801561029957600080fd5b506101026102a8366004611b13565b610750565b3480156102b957600080fd5b506101026102c8366004611b13565b610894565b6102d561090a565b816103315760405162461bcd60e51b815260206004820152602160248201527f6e65656473206174206c65617374206f6e652075736572206f7065726174696f6044820152603760f91b60648201526084015b60405180910390fd5b8160005b818110156103da573685858381811061035057610350611cd1565b90506020028101906103629190611ce7565b905080356103708282610963565b61037a82826109c1565b61038382610afd565b6103d08160006103966060860186611d08565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610c1992505050565b5050600101610335565b50506103e66001600355565b505050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036104335760405162461bcd60e51b815260040161032890611d4f565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661047c60008051602061205d833981519152546001600160a01b031690565b6001600160a01b0316146104a25760405162461bcd60e51b815260040161032890611d9b565b6104ab81610c97565b604080516000808252602082019092526104c791839190610c9f565b50565b60006104d582610e0a565b6040805160208101929092523090820152466060820152608001604051602081830303815290604052805190602001209050919050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036105545760405162461bcd60e51b815260040161032890611d4f565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661059d60008051602061205d833981519152546001600160a01b031690565b6001600160a01b0316146105c35760405162461bcd60e51b815260040161032890611d9b565b6105cc82610c97565b6105d882826001610c9f565b5050565b6000306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461067c5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610328565b5060008051602061205d83398151915290565b610697610e23565b6106a16000610e7d565b565b60006106ae84610ecf565b6001600160a01b038535166000908152600160209081526040918290205491516106e19392469230928991899101611e37565b6040516020818303038152906040528051906020012090509392505050565b610708610e23565b600080546001600160a01b0319166001600160a01b038316908117825560405190917fec30c97bb366e1114879be4964c925a5a9d29f5932231fdb30ef0d53de7ac04c91a250565b600254610100900460ff16158080156107705750600254600160ff909116105b8061078a5750303b15801561078a575060025460ff166001145b6107ed5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610328565b6002805460ff191660011790558015610810576002805461ff0019166101001790555b610818610f87565b61083a82600080546001600160a01b0319166001600160a01b03831617905550565b610842610fb6565b61084b82610fe5565b80156105d8576002805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b61089c610e23565b6001600160a01b0381166109015760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610328565b6104c781610e7d565b60026003540361095c5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610328565b6002600355565b6001600160a01b0381166000908152600160205260408120549050808360200135146103e65760405162461bcd60e51b815260206004820152600d60248201526c696e76616c6964206e6f6e636560981b6044820152606401610328565b6000610a026109cf846104ca565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c91909152603c902090565b90508260200135600003610a1a57610a1a8382610fee565b81630b135d3f60e11b6001600160a01b038216631626ba7e84610a41610140890189611d08565b6040518463ffffffff1660e01b8152600401610a5f93929190611e87565b602060405180830381865afa158015610a7c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa09190611ebd565b6001600160e01b03191614610af75760405162461bcd60e51b815260206004820152601960248201527f696e76616c6964206163636f756e74207369676e6174757265000000000000006044820152606401610328565b50505050565b366000610b0e610120840184611d08565b90925090506014811015610b645760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964207061796d6173746572416e644461746100000000000000006044820152606401610328565b6000610b736014828486611ee7565b610b7c91611f11565b60601c9050308114610bc45760405162461bcd60e51b815260206004820152601160248201527034b73b30b634b2103830bcb6b0b9ba32b960791b6044820152606401610328565b610bcd84611274565b610af75760405162461bcd60e51b815260206004820152601b60248201527f696e76616c6964207061796d6173746572207369676e617475726500000000006044820152606401610328565b600080846001600160a01b03168484604051610c359190611f46565b60006040518083038185875af1925050503d8060008114610c72576040519150601f19603f3d011682016040523d82523d6000602084013e610c77565b606091505b509150915081610c8957805160208201fd5b5050505050565b6001600355565b6104c7610e23565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610cd2576103e68361143c565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610d2c575060408051601f3d908101601f19168201909252610d2991810190611f58565b60015b610d8f5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610328565b60008051602061205d8339815191528114610dfe5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610328565b506103e68383836114d8565b6000610e15826114fd565b805190602001209050919050565b6067546001600160a01b031633146106a15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610328565b606780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6060813560208301356000610eef610eea6040870187611d08565b6115cb565b90506000610f03610eea6060880188611d08565b604080516001600160a01b039690961660208701528581019490945260608501929092525060808084019190915284013560a08084019190915284013560c08084019190915284013560e080840191909152840135610100808401919091529093013561012080830191909152835180830390910181526101409091019092525090565b600254610100900460ff16610fae5760405162461bcd60e51b815260040161032890611f71565b6106a16115de565b600254610100900460ff16610fdd5760405162461bcd60e51b815260040161032890611f71565b6106a161160e565b6104c781610894565b366000610ffe6040850185611d08565b909250905060148110156110475760405162461bcd60e51b815260206004820152601060248201526f696e76616c696420696e6974436f646560801b6044820152606401610328565b60006110566014828486611ee7565b61105f91611f11565b60601c9050803b6110a45760405162461bcd60e51b815260206004820152600f60248201526e696e76616c696420666163746f727960881b6044820152606401610328565b8060006110f56110b8610140890189611d08565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a939250506116359050565b9050863560405163119709c360e31b81526001600160a01b0383811660048301526000602483015291821691841690638cb84e1890604401602060405180830381865afa15801561114a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061116e9190611fbc565b6001600160a01b0316146111c45760405162461bcd60e51b815260206004820152601a60248201527f666163746f7279206d7573742072657475726e2073656e6465720000000000006044820152606401610328565b3660006111d4866014818a611ee7565b9150915061121a85600084848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610c1992505050565b88353b6112695760405162461bcd60e51b815260206004820152601e60248201527f696e76616c6964206163636f756e7420696e697469616c697a6174696f6e00006044820152606401610328565b505050505050505050565b60008080368161129061128b610120880188611d08565b611659565b9296509094509250905060408114806112a95750604181145b61131d576040805162461bcd60e51b81526020600482015260248101919091527f566572696679696e675061796d61737465723a20696e76616c6964207369676e60448201527f6174757265206c656e67746820696e207061796d6173746572416e64446174616064820152608401610328565b4265ffffffffffff80851690821611801561134857508465ffffffffffff168165ffffffffffff1611155b61138c5760405162461bcd60e51b81526020600482015260156024820152741cda59db985d1d5c99481a185cc8195e1c1a5c9959605a1b6044820152606401610328565b600061139c6109cf8988886106a3565b6001600160a01b0389351660009081526001602052604081208054929350906113c483611fd9565b919050555061140b84848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525085939250506116359050565b6000546001600160a01b0390811691161461142e57506000979650505050505050565b506001979650505050505050565b6001600160a01b0381163b6114a95760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610328565b60008051602061205d83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6114e183611696565b6000825111806114ee5750805b156103e657610af783836116d6565b6060813560208301356000611518610eea6040870187611d08565b9050600061152c610eea6060880188611d08565b9050608086013560a087013560c088013560e08901356101008a0135600061155b610eea6101208e018e611d08565b604080516001600160a01b039c909c1660208d01528b81019a909a5260608b019890985250608089019590955260a088019390935260c087019190915260e08601526101008501526101208401526101408084019190915281518084039091018152610160909201905292915050565b6000604051828085833790209392505050565b600254610100900460ff166116055760405162461bcd60e51b815260040161032890611f71565b6106a133610e7d565b600254610100900460ff16610c905760405162461bcd60e51b815260040161032890611f71565b60008060006116448585611702565b9150915061165181611747565b509392505050565b600080368161166c605460148789611ee7565b8101906116799190612000565b909450925061168b8560548189611ee7565b949793965094505050565b61169f8161143c565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606116fb838360405180606001604052806027815260200161207d60279139611891565b9392505050565b60008082516041036117385760208301516040840151606085015160001a61172c87828585611909565b94509450505050611740565b506000905060025b9250929050565b600081600481111561175b5761175b612033565b036117635750565b600181600481111561177757611777612033565b036117c45760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610328565b60028160048111156117d8576117d8612033565b036118255760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610328565b600381600481111561183957611839612033565b036104c75760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610328565b6060600080856001600160a01b0316856040516118ae9190611f46565b600060405180830381855af49150503d80600081146118e9576040519150601f19603f3d011682016040523d82523d6000602084013e6118ee565b606091505b50915091506118ff868383876119cd565b9695505050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561194057506000905060036119c4565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611994573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166119bd576000600192509250506119c4565b9150600090505b94509492505050565b60608315611a3c578251600003611a35576001600160a01b0385163b611a355760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610328565b5081611a46565b611a468383611a4e565b949350505050565b815115611a5e5781518083602001fd5b8060405162461bcd60e51b81526004016103289190612049565b6001600160a01b03811681146104c757600080fd5b600080600060408486031215611aa257600080fd5b833567ffffffffffffffff80821115611aba57600080fd5b818601915086601f830112611ace57600080fd5b813581811115611add57600080fd5b8760208260051b8501011115611af257600080fd5b60209283019550935050840135611b0881611a78565b809150509250925092565b600060208284031215611b2557600080fd5b81356116fb81611a78565b60006101608284031215611b4357600080fd5b50919050565b600060208284031215611b5b57600080fd5b813567ffffffffffffffff811115611b7257600080fd5b611a4684828501611b30565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215611ba757600080fd5b8235611bb281611a78565b9150602083013567ffffffffffffffff80821115611bcf57600080fd5b818501915085601f830112611be357600080fd5b813581811115611bf557611bf5611b7e565b604051601f8201601f19908116603f01168101908382118183101715611c1d57611c1d611b7e565b81604052828152886020848701011115611c3657600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b803565ffffffffffff81168114611c6e57600080fd5b919050565b600080600060608486031215611c8857600080fd5b833567ffffffffffffffff811115611c9f57600080fd5b611cab86828701611b30565b935050611cba60208501611c58565b9150611cc860408501611c58565b90509250925092565b634e487b7160e01b600052603260045260246000fd5b6000823561015e19833603018112611cfe57600080fd5b9190910192915050565b6000808335601e19843603018112611d1f57600080fd5b83018035915067ffffffffffffffff821115611d3a57600080fd5b60200191503681900382131561174057600080fd5b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b60005b83811015611e02578181015183820152602001611dea565b50506000910152565b60008151808452611e23816020860160208601611de7565b601f01601f19169290920160200192915050565b60c081526000611e4a60c0830189611e0b565b6020830197909752506001600160a01b03949094166040850152606084019290925265ffffffffffff90811660808401521660a090910152919050565b83815260406020820152816040820152818360608301376000818301606090810191909152601f909201601f1916010192915050565b600060208284031215611ecf57600080fd5b81516001600160e01b0319811681146116fb57600080fd5b60008085851115611ef757600080fd5b83861115611f0457600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff198135818116916014851015611f3e5780818660140360031b1b83161692505b505092915050565b60008251611cfe818460208701611de7565b600060208284031215611f6a57600080fd5b5051919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b600060208284031215611fce57600080fd5b81516116fb81611a78565b600060018201611ff957634e487b7160e01b600052601160045260246000fd5b5060010190565b6000806040838503121561201357600080fd5b61201c83611c58565b915061202a60208401611c58565b90509250929050565b634e487b7160e01b600052602160045260246000fd5b6020815260006116fb6020830184611e0b56fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122059fa27b99f7c55c4514ba8a074109bf27c280562622f9ab7a5b8ab3a00ac711f64736f6c63430008140033",
}

// AuthorizerABI is the input ABI used to generate the binding from.
// Deprecated: Use AuthorizerMetaData.ABI instead.
var AuthorizerABI = AuthorizerMetaData.ABI

// AuthorizerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuthorizerMetaData.Bin instead.
var AuthorizerBin = AuthorizerMetaData.Bin

// DeployAuthorizer deploys a new Ethereum contract, binding an instance of Authorizer to it.
func DeployAuthorizer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Authorizer, error) {
	parsed, err := AuthorizerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuthorizerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Authorizer{AuthorizerCaller: AuthorizerCaller{contract: contract}, AuthorizerTransactor: AuthorizerTransactor{contract: contract}, AuthorizerFilterer: AuthorizerFilterer{contract: contract}}, nil
}

// Authorizer is an auto generated Go binding around an Ethereum contract.
type Authorizer struct {
	AuthorizerCaller     // Read-only binding to the contract
	AuthorizerTransactor // Write-only binding to the contract
	AuthorizerFilterer   // Log filterer for contract events
}

// AuthorizerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthorizerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorizerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthorizerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorizerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthorizerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorizerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthorizerSession struct {
	Contract     *Authorizer       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthorizerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthorizerCallerSession struct {
	Contract *AuthorizerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AuthorizerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthorizerTransactorSession struct {
	Contract     *AuthorizerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AuthorizerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthorizerRaw struct {
	Contract *Authorizer // Generic contract binding to access the raw methods on
}

// AuthorizerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthorizerCallerRaw struct {
	Contract *AuthorizerCaller // Generic read-only contract binding to access the raw methods on
}

// AuthorizerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthorizerTransactorRaw struct {
	Contract *AuthorizerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthorizer creates a new instance of Authorizer, bound to a specific deployed contract.
func NewAuthorizer(address common.Address, backend bind.ContractBackend) (*Authorizer, error) {
	contract, err := bindAuthorizer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Authorizer{AuthorizerCaller: AuthorizerCaller{contract: contract}, AuthorizerTransactor: AuthorizerTransactor{contract: contract}, AuthorizerFilterer: AuthorizerFilterer{contract: contract}}, nil
}

// NewAuthorizerCaller creates a new read-only instance of Authorizer, bound to a specific deployed contract.
func NewAuthorizerCaller(address common.Address, caller bind.ContractCaller) (*AuthorizerCaller, error) {
	contract, err := bindAuthorizer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorizerCaller{contract: contract}, nil
}

// NewAuthorizerTransactor creates a new write-only instance of Authorizer, bound to a specific deployed contract.
func NewAuthorizerTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthorizerTransactor, error) {
	contract, err := bindAuthorizer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorizerTransactor{contract: contract}, nil
}

// NewAuthorizerFilterer creates a new log filterer instance of Authorizer, bound to a specific deployed contract.
func NewAuthorizerFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthorizerFilterer, error) {
	contract, err := bindAuthorizer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthorizerFilterer{contract: contract}, nil
}

// bindAuthorizer binds a generic wrapper to an already deployed contract.
func bindAuthorizer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuthorizerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authorizer *AuthorizerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Authorizer.Contract.AuthorizerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authorizer *AuthorizerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authorizer.Contract.AuthorizerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authorizer *AuthorizerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authorizer.Contract.AuthorizerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authorizer *AuthorizerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Authorizer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authorizer *AuthorizerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authorizer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authorizer *AuthorizerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authorizer.Contract.contract.Transact(opts, method, params...)
}

// UserOpHash is a free data retrieval call binding the contract method 0x4564cc2e.
//
// Solidity: function _userOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_Authorizer *AuthorizerCaller) UserOpHash(opts *bind.CallOpts, userOp UserOperation) ([32]byte, error) {
	var out []interface{}
	err := _Authorizer.contract.Call(opts, &out, "_userOpHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UserOpHash is a free data retrieval call binding the contract method 0x4564cc2e.
//
// Solidity: function _userOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_Authorizer *AuthorizerSession) UserOpHash(userOp UserOperation) ([32]byte, error) {
	return _Authorizer.Contract.UserOpHash(&_Authorizer.CallOpts, userOp)
}

// UserOpHash is a free data retrieval call binding the contract method 0x4564cc2e.
//
// Solidity: function _userOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_Authorizer *AuthorizerCallerSession) UserOpHash(userOp UserOperation) ([32]byte, error) {
	return _Authorizer.Contract.UserOpHash(&_Authorizer.CallOpts, userOp)
}

// GetHash is a free data retrieval call binding the contract method 0x94e1fc19.
//
// Solidity: function getHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, uint48 validUntil, uint48 validAfter) view returns(bytes32)
func (_Authorizer *AuthorizerCaller) GetHash(opts *bind.CallOpts, userOp UserOperation, validUntil *big.Int, validAfter *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Authorizer.contract.Call(opts, &out, "getHash", userOp, validUntil, validAfter)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0x94e1fc19.
//
// Solidity: function getHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, uint48 validUntil, uint48 validAfter) view returns(bytes32)
func (_Authorizer *AuthorizerSession) GetHash(userOp UserOperation, validUntil *big.Int, validAfter *big.Int) ([32]byte, error) {
	return _Authorizer.Contract.GetHash(&_Authorizer.CallOpts, userOp, validUntil, validAfter)
}

// GetHash is a free data retrieval call binding the contract method 0x94e1fc19.
//
// Solidity: function getHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, uint48 validUntil, uint48 validAfter) view returns(bytes32)
func (_Authorizer *AuthorizerCallerSession) GetHash(userOp UserOperation, validUntil *big.Int, validAfter *big.Int) ([32]byte, error) {
	return _Authorizer.Contract.GetHash(&_Authorizer.CallOpts, userOp, validUntil, validAfter)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address sender) view returns(uint256 nonce)
func (_Authorizer *AuthorizerCaller) GetNonce(opts *bind.CallOpts, sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Authorizer.contract.Call(opts, &out, "getNonce", sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address sender) view returns(uint256 nonce)
func (_Authorizer *AuthorizerSession) GetNonce(sender common.Address) (*big.Int, error) {
	return _Authorizer.Contract.GetNonce(&_Authorizer.CallOpts, sender)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address sender) view returns(uint256 nonce)
func (_Authorizer *AuthorizerCallerSession) GetNonce(sender common.Address) (*big.Int, error) {
	return _Authorizer.Contract.GetNonce(&_Authorizer.CallOpts, sender)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Authorizer *AuthorizerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Authorizer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Authorizer *AuthorizerSession) Owner() (common.Address, error) {
	return _Authorizer.Contract.Owner(&_Authorizer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Authorizer *AuthorizerCallerSession) Owner() (common.Address, error) {
	return _Authorizer.Contract.Owner(&_Authorizer.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Authorizer *AuthorizerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Authorizer.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Authorizer *AuthorizerSession) ProxiableUUID() ([32]byte, error) {
	return _Authorizer.Contract.ProxiableUUID(&_Authorizer.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Authorizer *AuthorizerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Authorizer.Contract.ProxiableUUID(&_Authorizer.CallOpts)
}

// SenderNonce is a free data retrieval call binding the contract method 0x9c90b443.
//
// Solidity: function senderNonce(address ) view returns(uint256)
func (_Authorizer *AuthorizerCaller) SenderNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Authorizer.contract.Call(opts, &out, "senderNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SenderNonce is a free data retrieval call binding the contract method 0x9c90b443.
//
// Solidity: function senderNonce(address ) view returns(uint256)
func (_Authorizer *AuthorizerSession) SenderNonce(arg0 common.Address) (*big.Int, error) {
	return _Authorizer.Contract.SenderNonce(&_Authorizer.CallOpts, arg0)
}

// SenderNonce is a free data retrieval call binding the contract method 0x9c90b443.
//
// Solidity: function senderNonce(address ) view returns(uint256)
func (_Authorizer *AuthorizerCallerSession) SenderNonce(arg0 common.Address) (*big.Int, error) {
	return _Authorizer.Contract.SenderNonce(&_Authorizer.CallOpts, arg0)
}

// Sponsor is a free data retrieval call binding the contract method 0x77c93662.
//
// Solidity: function sponsor() view returns(address)
func (_Authorizer *AuthorizerCaller) Sponsor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Authorizer.contract.Call(opts, &out, "sponsor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sponsor is a free data retrieval call binding the contract method 0x77c93662.
//
// Solidity: function sponsor() view returns(address)
func (_Authorizer *AuthorizerSession) Sponsor() (common.Address, error) {
	return _Authorizer.Contract.Sponsor(&_Authorizer.CallOpts)
}

// Sponsor is a free data retrieval call binding the contract method 0x77c93662.
//
// Solidity: function sponsor() view returns(address)
func (_Authorizer *AuthorizerCallerSession) Sponsor() (common.Address, error) {
	return _Authorizer.Contract.Sponsor(&_Authorizer.CallOpts)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_Authorizer *AuthorizerTransactor) HandleOps(opts *bind.TransactOpts, ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _Authorizer.contract.Transact(opts, "handleOps", ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_Authorizer *AuthorizerSession) HandleOps(ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _Authorizer.Contract.HandleOps(&_Authorizer.TransactOpts, ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_Authorizer *AuthorizerTransactorSession) HandleOps(ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _Authorizer.Contract.HandleOps(&_Authorizer.TransactOpts, ops, beneficiary)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_Authorizer *AuthorizerTransactor) Initialize(opts *bind.TransactOpts, anOwner common.Address) (*types.Transaction, error) {
	return _Authorizer.contract.Transact(opts, "initialize", anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_Authorizer *AuthorizerSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _Authorizer.Contract.Initialize(&_Authorizer.TransactOpts, anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_Authorizer *AuthorizerTransactorSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _Authorizer.Contract.Initialize(&_Authorizer.TransactOpts, anOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Authorizer *AuthorizerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authorizer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Authorizer *AuthorizerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Authorizer.Contract.RenounceOwnership(&_Authorizer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Authorizer *AuthorizerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Authorizer.Contract.RenounceOwnership(&_Authorizer.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Authorizer *AuthorizerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Authorizer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Authorizer *AuthorizerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Authorizer.Contract.TransferOwnership(&_Authorizer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Authorizer *AuthorizerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Authorizer.Contract.TransferOwnership(&_Authorizer.TransactOpts, newOwner)
}

// UpdatePaymasterSponsor is a paid mutator transaction binding the contract method 0xba4fdba0.
//
// Solidity: function updatePaymasterSponsor(address newSponsor) returns()
func (_Authorizer *AuthorizerTransactor) UpdatePaymasterSponsor(opts *bind.TransactOpts, newSponsor common.Address) (*types.Transaction, error) {
	return _Authorizer.contract.Transact(opts, "updatePaymasterSponsor", newSponsor)
}

// UpdatePaymasterSponsor is a paid mutator transaction binding the contract method 0xba4fdba0.
//
// Solidity: function updatePaymasterSponsor(address newSponsor) returns()
func (_Authorizer *AuthorizerSession) UpdatePaymasterSponsor(newSponsor common.Address) (*types.Transaction, error) {
	return _Authorizer.Contract.UpdatePaymasterSponsor(&_Authorizer.TransactOpts, newSponsor)
}

// UpdatePaymasterSponsor is a paid mutator transaction binding the contract method 0xba4fdba0.
//
// Solidity: function updatePaymasterSponsor(address newSponsor) returns()
func (_Authorizer *AuthorizerTransactorSession) UpdatePaymasterSponsor(newSponsor common.Address) (*types.Transaction, error) {
	return _Authorizer.Contract.UpdatePaymasterSponsor(&_Authorizer.TransactOpts, newSponsor)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Authorizer *AuthorizerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Authorizer.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Authorizer *AuthorizerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Authorizer.Contract.UpgradeTo(&_Authorizer.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Authorizer *AuthorizerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Authorizer.Contract.UpgradeTo(&_Authorizer.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Authorizer *AuthorizerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Authorizer.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Authorizer *AuthorizerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Authorizer.Contract.UpgradeToAndCall(&_Authorizer.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Authorizer *AuthorizerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Authorizer.Contract.UpgradeToAndCall(&_Authorizer.TransactOpts, newImplementation, data)
}

// AuthorizerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Authorizer contract.
type AuthorizerAdminChangedIterator struct {
	Event *AuthorizerAdminChanged // Event containing the contract specifics and raw log

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
func (it *AuthorizerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorizerAdminChanged)
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
		it.Event = new(AuthorizerAdminChanged)
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
func (it *AuthorizerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorizerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorizerAdminChanged represents a AdminChanged event raised by the Authorizer contract.
type AuthorizerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Authorizer *AuthorizerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AuthorizerAdminChangedIterator, error) {

	logs, sub, err := _Authorizer.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AuthorizerAdminChangedIterator{contract: _Authorizer.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Authorizer *AuthorizerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AuthorizerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Authorizer.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorizerAdminChanged)
				if err := _Authorizer.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Authorizer *AuthorizerFilterer) ParseAdminChanged(log types.Log) (*AuthorizerAdminChanged, error) {
	event := new(AuthorizerAdminChanged)
	if err := _Authorizer.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthorizerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Authorizer contract.
type AuthorizerBeaconUpgradedIterator struct {
	Event *AuthorizerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *AuthorizerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorizerBeaconUpgraded)
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
		it.Event = new(AuthorizerBeaconUpgraded)
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
func (it *AuthorizerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorizerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorizerBeaconUpgraded represents a BeaconUpgraded event raised by the Authorizer contract.
type AuthorizerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Authorizer *AuthorizerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AuthorizerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Authorizer.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AuthorizerBeaconUpgradedIterator{contract: _Authorizer.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Authorizer *AuthorizerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AuthorizerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Authorizer.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorizerBeaconUpgraded)
				if err := _Authorizer.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Authorizer *AuthorizerFilterer) ParseBeaconUpgraded(log types.Log) (*AuthorizerBeaconUpgraded, error) {
	event := new(AuthorizerBeaconUpgraded)
	if err := _Authorizer.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthorizerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Authorizer contract.
type AuthorizerInitializedIterator struct {
	Event *AuthorizerInitialized // Event containing the contract specifics and raw log

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
func (it *AuthorizerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorizerInitialized)
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
		it.Event = new(AuthorizerInitialized)
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
func (it *AuthorizerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorizerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorizerInitialized represents a Initialized event raised by the Authorizer contract.
type AuthorizerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Authorizer *AuthorizerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AuthorizerInitializedIterator, error) {

	logs, sub, err := _Authorizer.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AuthorizerInitializedIterator{contract: _Authorizer.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Authorizer *AuthorizerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AuthorizerInitialized) (event.Subscription, error) {

	logs, sub, err := _Authorizer.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorizerInitialized)
				if err := _Authorizer.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Authorizer *AuthorizerFilterer) ParseInitialized(log types.Log) (*AuthorizerInitialized, error) {
	event := new(AuthorizerInitialized)
	if err := _Authorizer.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthorizerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Authorizer contract.
type AuthorizerOwnershipTransferredIterator struct {
	Event *AuthorizerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuthorizerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorizerOwnershipTransferred)
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
		it.Event = new(AuthorizerOwnershipTransferred)
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
func (it *AuthorizerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorizerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorizerOwnershipTransferred represents a OwnershipTransferred event raised by the Authorizer contract.
type AuthorizerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Authorizer *AuthorizerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuthorizerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Authorizer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuthorizerOwnershipTransferredIterator{contract: _Authorizer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Authorizer *AuthorizerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuthorizerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Authorizer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorizerOwnershipTransferred)
				if err := _Authorizer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Authorizer *AuthorizerFilterer) ParseOwnershipTransferred(log types.Log) (*AuthorizerOwnershipTransferred, error) {
	event := new(AuthorizerOwnershipTransferred)
	if err := _Authorizer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthorizerPaymasterSponsorUpdatedIterator is returned from FilterPaymasterSponsorUpdated and is used to iterate over the raw logs and unpacked data for PaymasterSponsorUpdated events raised by the Authorizer contract.
type AuthorizerPaymasterSponsorUpdatedIterator struct {
	Event *AuthorizerPaymasterSponsorUpdated // Event containing the contract specifics and raw log

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
func (it *AuthorizerPaymasterSponsorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorizerPaymasterSponsorUpdated)
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
		it.Event = new(AuthorizerPaymasterSponsorUpdated)
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
func (it *AuthorizerPaymasterSponsorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorizerPaymasterSponsorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorizerPaymasterSponsorUpdated represents a PaymasterSponsorUpdated event raised by the Authorizer contract.
type AuthorizerPaymasterSponsorUpdated struct {
	NewVerifier common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPaymasterSponsorUpdated is a free log retrieval operation binding the contract event 0xec30c97bb366e1114879be4964c925a5a9d29f5932231fdb30ef0d53de7ac04c.
//
// Solidity: event PaymasterSponsorUpdated(address indexed newVerifier)
func (_Authorizer *AuthorizerFilterer) FilterPaymasterSponsorUpdated(opts *bind.FilterOpts, newVerifier []common.Address) (*AuthorizerPaymasterSponsorUpdatedIterator, error) {

	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _Authorizer.contract.FilterLogs(opts, "PaymasterSponsorUpdated", newVerifierRule)
	if err != nil {
		return nil, err
	}
	return &AuthorizerPaymasterSponsorUpdatedIterator{contract: _Authorizer.contract, event: "PaymasterSponsorUpdated", logs: logs, sub: sub}, nil
}

// WatchPaymasterSponsorUpdated is a free log subscription operation binding the contract event 0xec30c97bb366e1114879be4964c925a5a9d29f5932231fdb30ef0d53de7ac04c.
//
// Solidity: event PaymasterSponsorUpdated(address indexed newVerifier)
func (_Authorizer *AuthorizerFilterer) WatchPaymasterSponsorUpdated(opts *bind.WatchOpts, sink chan<- *AuthorizerPaymasterSponsorUpdated, newVerifier []common.Address) (event.Subscription, error) {

	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _Authorizer.contract.WatchLogs(opts, "PaymasterSponsorUpdated", newVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorizerPaymasterSponsorUpdated)
				if err := _Authorizer.contract.UnpackLog(event, "PaymasterSponsorUpdated", log); err != nil {
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

// ParsePaymasterSponsorUpdated is a log parse operation binding the contract event 0xec30c97bb366e1114879be4964c925a5a9d29f5932231fdb30ef0d53de7ac04c.
//
// Solidity: event PaymasterSponsorUpdated(address indexed newVerifier)
func (_Authorizer *AuthorizerFilterer) ParsePaymasterSponsorUpdated(log types.Log) (*AuthorizerPaymasterSponsorUpdated, error) {
	event := new(AuthorizerPaymasterSponsorUpdated)
	if err := _Authorizer.contract.UnpackLog(event, "PaymasterSponsorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthorizerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Authorizer contract.
type AuthorizerUpgradedIterator struct {
	Event *AuthorizerUpgraded // Event containing the contract specifics and raw log

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
func (it *AuthorizerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorizerUpgraded)
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
		it.Event = new(AuthorizerUpgraded)
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
func (it *AuthorizerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorizerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorizerUpgraded represents a Upgraded event raised by the Authorizer contract.
type AuthorizerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Authorizer *AuthorizerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AuthorizerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Authorizer.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AuthorizerUpgradedIterator{contract: _Authorizer.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Authorizer *AuthorizerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AuthorizerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Authorizer.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorizerUpgraded)
				if err := _Authorizer.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Authorizer *AuthorizerFilterer) ParseUpgraded(log types.Log) (*AuthorizerUpgraded, error) {
	event := new(AuthorizerUpgraded)
	if err := _Authorizer.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}