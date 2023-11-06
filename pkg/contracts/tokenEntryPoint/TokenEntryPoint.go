// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenEntryPoint

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

// TokenEntryPointMetaData contains all meta data concerning the TokenEntryPoint contract.
var TokenEntryPointMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aPaymaster\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPaymaster\",\"type\":\"address\"}],\"name\":\"updatePaymaster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100e1565b600054610100900460ff161561008e5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146100df576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b608051611f20610118600039600081816103a7015281816103e7015281816105aa015281816105ea015261067d0152611f206000f3fe60806040526004361061009c5760003560e01c806352d1902d1161006457806352d1902d14610149578063715018a61461015e5780638da5cb5b14610173578063a61935311461019b578063daf6a3f6146101bb578063f2fde38b146101db57600080fd5b80631fad948c146100a157806335567e1a146100c35780633659cfe6146100f6578063485cc955146101165780634f1ef28614610136575b600080fd5b3480156100ad57600080fd5b506100c16100bc366004611770565b6101fb565b005b3480156100cf57600080fd5b506100e36100de366004611806565b610318565b6040519081526020015b60405180910390f35b34801561010257600080fd5b506100c161011136600461184b565b61039d565b34801561012257600080fd5b506100c1610131366004611868565b61047c565b6100c1610144366004611905565b6105a0565b34801561015557600080fd5b506100e3610670565b34801561016a57600080fd5b506100c1610723565b34801561017f57600080fd5b506065546040516001600160a01b0390911681526020016100ed565b3480156101a757600080fd5b506100e36101b6366004611998565b610737565b3480156101c757600080fd5b506100c16101d636600461184b565b610779565b3480156101e757600080fd5b506100c16101f636600461184b565b6107e5565b61020361085b565b8161025f5760405162461bcd60e51b815260206004820152602160248201527f6e65656473206174206c65617374206f6e652075736572206f7065726174696f6044820152603760f91b60648201526084015b60405180910390fd5b8160005b81811015610308573685858381811061027e5761027e6119d4565b905060200281019061029091906119ea565b9050803561029e82826108b4565b6102a88282610906565b6102b182610a42565b6102fe8160006102c46060860186611a0b565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610ac892505050565b5050600101610263565b505061031360018055565b505050565b60fb54604051631aab3f0d60e11b81526001600160a01b0384811660048301526001600160c01b038416602483015260009216906335567e1a90604401602060405180830381865afa158015610372573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103969190611a52565b9392505050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036103e55760405162461bcd60e51b815260040161025690611a6b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661042e600080516020611ea4833981519152546001600160a01b031690565b6001600160a01b0316146104545760405162461bcd60e51b815260040161025690611ab7565b61045d81610b45565b6040805160008082526020820190925261047991839190610b4d565b50565b600054610100900460ff161580801561049c5750600054600160ff909116105b806104b65750303b1580156104b6575060005460ff166001145b6105195760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610256565b6000805460ff19166001179055801561053c576000805461ff0019166101001790555b610544610cb8565b61054c610ce7565b6105568383610d16565b8015610313576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036105e85760405162461bcd60e51b815260040161025690611a6b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610631600080516020611ea4833981519152546001600160a01b031690565b6001600160a01b0316146106575760405162461bcd60e51b815260040161025690611ab7565b61066082610b45565b61066c82826001610b4d565b5050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107105760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610256565b50600080516020611ea483398151915290565b61072b610d42565b6107356000610d9c565b565b600061074282610dee565b6040805160208101929092523090820152466060820152608001604051602081830303815290604052805190602001209050919050565b610781610d42565b803b6107c35760405162461bcd60e51b815260206004820152601160248201527034b73b30b634b2103830bcb6b0b9ba32b960791b6044820152606401610256565b60fb80546001600160a01b0319166001600160a01b0392909216919091179055565b6107ed610d42565b6001600160a01b0381166108525760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610256565b61047981610d9c565b6002600154036108ad5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610256565b6002600155565b60006108c1826000610318565b9050808360200135146103135760405162461bcd60e51b815260206004820152600d60248201526c696e76616c6964206e6f6e636560981b6044820152606401610256565b600061094761091484610737565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c91909152603c902090565b9050826020013560000361095f5761095f8382610e07565b81630b135d3f60e11b6001600160a01b038216631626ba7e84610986610140890189611a0b565b6040518463ffffffff1660e01b81526004016109a493929190611b2c565b602060405180830381865afa1580156109c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e59190611b4f565b6001600160e01b03191614610a3c5760405162461bcd60e51b815260206004820152601960248201527f696e76616c6964206163636f756e74207369676e6174757265000000000000006044820152606401610256565b50505050565b6000610a4d8261108d565b604051637a32e3bf60e11b81529091506001600160a01b0382169063f465c77e90610a819085906000908190600401611bbf565b6000604051808303816000875af1158015610aa0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a3c9190810190611d00565b600080846001600160a01b03168484604051610ae49190611d81565b60006040518083038185875af1925050503d8060008114610b21576040519150601f19603f3d011682016040523d82523d6000602084013e610b26565b606091505b509150915081610b3857805160208201fd5b5050505050565b60018055565b610479610d42565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610b80576103138361116b565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610bda575060408051601f3d908101601f19168201909252610bd791810190611a52565b60015b610c3d5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610256565b600080516020611ea48339815191528114610cac5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610256565b50610313838383611207565b600054610100900460ff16610cdf5760405162461bcd60e51b815260040161025690611d93565b61073561122c565b600054610100900460ff16610d0e5760405162461bcd60e51b815260040161025690611d93565b61073561125c565b610d1f826107e5565b60fb80546001600160a01b0319166001600160a01b039290921691909117905550565b6065546001600160a01b031633146107355760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610256565b606580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000610df982611283565b805190602001209050919050565b366000610e176040850185611a0b565b90925090506014811015610e605760405162461bcd60e51b815260206004820152601060248201526f696e76616c696420696e6974436f646560801b6044820152606401610256565b6000610e6f6014828486611dde565b610e7891611e08565b60601c9050803b610ebd5760405162461bcd60e51b815260206004820152600f60248201526e696e76616c696420666163746f727960881b6044820152606401610256565b6000610f0d610ed0610140880188611a0b565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525089939250506113569050565b905081863560405163119709c360e31b81526001600160a01b0384811660048301526000602483015291821691831690638cb84e1890604401602060405180830381865afa158015610f63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f879190611e3d565b6001600160a01b031614610fdd5760405162461bcd60e51b815260206004820152601a60248201527f666163746f7279206d7573742072657475726e2073656e6465720000000000006044820152606401610256565b366000610fed866014818a611dde565b9150915061103385600084848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610ac892505050565b88353b6110825760405162461bcd60e51b815260206004820152601e60248201527f696e76616c6964206163636f756e7420696e697469616c697a6174696f6e00006044820152606401610256565b505050505050505050565b6000368161109f610120850185611a0b565b909250905060148110156110f55760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964207061796d6173746572416e644461746100000000000000006044820152606401610256565b60006111046014828486611dde565b61110d91611e08565b60fb5460609190911c91506001600160a01b031681146111635760405162461bcd60e51b815260206004820152601160248201527034b73b30b634b2103830bcb6b0b9ba32b960791b6044820152606401610256565b949350505050565b6001600160a01b0381163b6111d85760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610256565b600080516020611ea483398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6112108361137a565b60008251118061121d5750805b1561031357610a3c83836113ba565b600054610100900460ff166112535760405162461bcd60e51b815260040161025690611d93565b61073533610d9c565b600054610100900460ff16610b3f5760405162461bcd60e51b815260040161025690611d93565b60608135602083013560006112a361129e6040870187611a0b565b6113df565b905060006112b761129e6060880188611a0b565b9050608086013560a087013560c088013560e08901356101008a013560006112e661129e6101208e018e611a0b565b604080516001600160a01b039c909c1660208d01528b81019a909a5260608b019890985250608089019590955260a088019390935260c087019190915260e08601526101008501526101208401526101408084019190915281518084039091018152610160909201905292915050565b600080600061136585856113f2565b9150915061137281611437565b509392505050565b6113838161116b565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606103968383604051806060016040528060278152602001611ec460279139611581565b6000604051828085833790209392505050565b60008082516041036114285760208301516040840151606085015160001a61141c878285856115f9565b94509450505050611430565b506000905060025b9250929050565b600081600481111561144b5761144b611e5a565b036114535750565b600181600481111561146757611467611e5a565b036114b45760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610256565b60028160048111156114c8576114c8611e5a565b036115155760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610256565b600381600481111561152957611529611e5a565b036104795760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610256565b6060600080856001600160a01b03168560405161159e9190611d81565b600060405180830381855af49150503d80600081146115d9576040519150601f19603f3d011682016040523d82523d6000602084013e6115de565b606091505b50915091506115ef868383876116bd565b9695505050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561163057506000905060036116b4565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611684573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166116ad576000600192509250506116b4565b9150600090505b94509492505050565b6060831561172c578251600003611725576001600160a01b0385163b6117255760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610256565b5081611163565b61116383838151156117415781518083602001fd5b8060405162461bcd60e51b81526004016102569190611e70565b6001600160a01b038116811461047957600080fd5b60008060006040848603121561178557600080fd5b833567ffffffffffffffff8082111561179d57600080fd5b818601915086601f8301126117b157600080fd5b8135818111156117c057600080fd5b8760208260051b85010111156117d557600080fd5b602092830195509350508401356117eb8161175b565b809150509250925092565b80356118018161175b565b919050565b6000806040838503121561181957600080fd5b82356118248161175b565b915060208301356001600160c01b038116811461184057600080fd5b809150509250929050565b60006020828403121561185d57600080fd5b81356103968161175b565b6000806040838503121561187b57600080fd5b82356118868161175b565b915060208301356118408161175b565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156118d5576118d5611896565b604052919050565b600067ffffffffffffffff8211156118f7576118f7611896565b50601f01601f191660200190565b6000806040838503121561191857600080fd5b82356119238161175b565b9150602083013567ffffffffffffffff81111561193f57600080fd5b8301601f8101851361195057600080fd5b803561196361195e826118dd565b6118ac565b81815286602083850101111561197857600080fd5b816020840160208301376000602083830101528093505050509250929050565b6000602082840312156119aa57600080fd5b813567ffffffffffffffff8111156119c157600080fd5b8201610160818503121561039657600080fd5b634e487b7160e01b600052603260045260246000fd5b6000823561015e19833603018112611a0157600080fd5b9190910192915050565b6000808335601e19843603018112611a2257600080fd5b83018035915067ffffffffffffffff821115611a3d57600080fd5b60200191503681900382131561143057600080fd5b600060208284031215611a6457600080fd5b5051919050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b838152604060208201526000611b46604083018486611b03565b95945050505050565b600060208284031215611b6157600080fd5b81516001600160e01b03198116811461039657600080fd5b6000808335601e19843603018112611b9057600080fd5b830160208101925035905067ffffffffffffffff811115611bb057600080fd5b80360382131561143057600080fd5b60608152611be060608201611bd3866117f6565b6001600160a01b03169052565b602084013560808201526000611bf96040860186611b79565b6101608060a0860152611c116101c086018385611b03565b9250611c206060890189611b79565b9250605f19808786030160c0880152611c3a858584611b03565b945060808a013560e0880152610100935060a08a013584880152610120915060c08a01358288015261014060e08b013581890152848b013584890152611c82838c018c611b79565b955093508188870301610180890152611c9c868686611b03565b9550611caa818c018c611b79565b955093505080878603016101a08801525050611cc7838383611b03565b60208601979097525050505060400152919050565b60005b83811015611cf7578181015183820152602001611cdf565b50506000910152565b60008060408385031215611d1357600080fd5b825167ffffffffffffffff811115611d2a57600080fd5b8301601f81018513611d3b57600080fd5b8051611d4961195e826118dd565b818152866020838501011115611d5e57600080fd5b611d6f826020830160208601611cdc565b60209590950151949694955050505050565b60008251611a01818460208701611cdc565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60008085851115611dee57600080fd5b83861115611dfb57600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff198135818116916014851015611e355780818660140360031b1b83161692505b505092915050565b600060208284031215611e4f57600080fd5b81516103968161175b565b634e487b7160e01b600052602160045260246000fd5b6020815260008251806020840152611e8f816040850160208701611cdc565b601f01601f1916919091016040019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122044c9345a6886a41fb2be121e824f93112c78c366782c3ebbd8ba3bbd12f669e364736f6c63430008140033",
}

// TokenEntryPointABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenEntryPointMetaData.ABI instead.
var TokenEntryPointABI = TokenEntryPointMetaData.ABI

// TokenEntryPointBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenEntryPointMetaData.Bin instead.
var TokenEntryPointBin = TokenEntryPointMetaData.Bin

// DeployTokenEntryPoint deploys a new Ethereum contract, binding an instance of TokenEntryPoint to it.
func DeployTokenEntryPoint(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenEntryPoint, error) {
	parsed, err := TokenEntryPointMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenEntryPointBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenEntryPoint{TokenEntryPointCaller: TokenEntryPointCaller{contract: contract}, TokenEntryPointTransactor: TokenEntryPointTransactor{contract: contract}, TokenEntryPointFilterer: TokenEntryPointFilterer{contract: contract}}, nil
}

// TokenEntryPoint is an auto generated Go binding around an Ethereum contract.
type TokenEntryPoint struct {
	TokenEntryPointCaller     // Read-only binding to the contract
	TokenEntryPointTransactor // Write-only binding to the contract
	TokenEntryPointFilterer   // Log filterer for contract events
}

// TokenEntryPointCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenEntryPointCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenEntryPointTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenEntryPointTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenEntryPointFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenEntryPointFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenEntryPointSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenEntryPointSession struct {
	Contract     *TokenEntryPoint  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenEntryPointCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenEntryPointCallerSession struct {
	Contract *TokenEntryPointCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TokenEntryPointTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenEntryPointTransactorSession struct {
	Contract     *TokenEntryPointTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TokenEntryPointRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenEntryPointRaw struct {
	Contract *TokenEntryPoint // Generic contract binding to access the raw methods on
}

// TokenEntryPointCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenEntryPointCallerRaw struct {
	Contract *TokenEntryPointCaller // Generic read-only contract binding to access the raw methods on
}

// TokenEntryPointTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenEntryPointTransactorRaw struct {
	Contract *TokenEntryPointTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenEntryPoint creates a new instance of TokenEntryPoint, bound to a specific deployed contract.
func NewTokenEntryPoint(address common.Address, backend bind.ContractBackend) (*TokenEntryPoint, error) {
	contract, err := bindTokenEntryPoint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenEntryPoint{TokenEntryPointCaller: TokenEntryPointCaller{contract: contract}, TokenEntryPointTransactor: TokenEntryPointTransactor{contract: contract}, TokenEntryPointFilterer: TokenEntryPointFilterer{contract: contract}}, nil
}

// NewTokenEntryPointCaller creates a new read-only instance of TokenEntryPoint, bound to a specific deployed contract.
func NewTokenEntryPointCaller(address common.Address, caller bind.ContractCaller) (*TokenEntryPointCaller, error) {
	contract, err := bindTokenEntryPoint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenEntryPointCaller{contract: contract}, nil
}

// NewTokenEntryPointTransactor creates a new write-only instance of TokenEntryPoint, bound to a specific deployed contract.
func NewTokenEntryPointTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenEntryPointTransactor, error) {
	contract, err := bindTokenEntryPoint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenEntryPointTransactor{contract: contract}, nil
}

// NewTokenEntryPointFilterer creates a new log filterer instance of TokenEntryPoint, bound to a specific deployed contract.
func NewTokenEntryPointFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenEntryPointFilterer, error) {
	contract, err := bindTokenEntryPoint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenEntryPointFilterer{contract: contract}, nil
}

// bindTokenEntryPoint binds a generic wrapper to an already deployed contract.
func bindTokenEntryPoint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenEntryPointMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenEntryPoint *TokenEntryPointRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenEntryPoint.Contract.TokenEntryPointCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenEntryPoint *TokenEntryPointRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.TokenEntryPointTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenEntryPoint *TokenEntryPointRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.TokenEntryPointTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenEntryPoint *TokenEntryPointCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenEntryPoint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenEntryPoint *TokenEntryPointTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenEntryPoint *TokenEntryPointTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.contract.Transact(opts, method, params...)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_TokenEntryPoint *TokenEntryPointCaller) GetNonce(opts *bind.CallOpts, sender common.Address, key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenEntryPoint.contract.Call(opts, &out, "getNonce", sender, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_TokenEntryPoint *TokenEntryPointSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _TokenEntryPoint.Contract.GetNonce(&_TokenEntryPoint.CallOpts, sender, key)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_TokenEntryPoint *TokenEntryPointCallerSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _TokenEntryPoint.Contract.GetNonce(&_TokenEntryPoint.CallOpts, sender, key)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0xa6193531.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_TokenEntryPoint *TokenEntryPointCaller) GetUserOpHash(opts *bind.CallOpts, userOp UserOperation) ([32]byte, error) {
	var out []interface{}
	err := _TokenEntryPoint.contract.Call(opts, &out, "getUserOpHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOpHash is a free data retrieval call binding the contract method 0xa6193531.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_TokenEntryPoint *TokenEntryPointSession) GetUserOpHash(userOp UserOperation) ([32]byte, error) {
	return _TokenEntryPoint.Contract.GetUserOpHash(&_TokenEntryPoint.CallOpts, userOp)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0xa6193531.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_TokenEntryPoint *TokenEntryPointCallerSession) GetUserOpHash(userOp UserOperation) ([32]byte, error) {
	return _TokenEntryPoint.Contract.GetUserOpHash(&_TokenEntryPoint.CallOpts, userOp)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenEntryPoint *TokenEntryPointCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenEntryPoint.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenEntryPoint *TokenEntryPointSession) Owner() (common.Address, error) {
	return _TokenEntryPoint.Contract.Owner(&_TokenEntryPoint.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenEntryPoint *TokenEntryPointCallerSession) Owner() (common.Address, error) {
	return _TokenEntryPoint.Contract.Owner(&_TokenEntryPoint.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TokenEntryPoint *TokenEntryPointCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenEntryPoint.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TokenEntryPoint *TokenEntryPointSession) ProxiableUUID() ([32]byte, error) {
	return _TokenEntryPoint.Contract.ProxiableUUID(&_TokenEntryPoint.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TokenEntryPoint *TokenEntryPointCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TokenEntryPoint.Contract.ProxiableUUID(&_TokenEntryPoint.CallOpts)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_TokenEntryPoint *TokenEntryPointTransactor) HandleOps(opts *bind.TransactOpts, ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.contract.Transact(opts, "handleOps", ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_TokenEntryPoint *TokenEntryPointSession) HandleOps(ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.HandleOps(&_TokenEntryPoint.TransactOpts, ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_TokenEntryPoint *TokenEntryPointTransactorSession) HandleOps(ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.HandleOps(&_TokenEntryPoint.TransactOpts, ops, beneficiary)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address anOwner, address aPaymaster) returns()
func (_TokenEntryPoint *TokenEntryPointTransactor) Initialize(opts *bind.TransactOpts, anOwner common.Address, aPaymaster common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.contract.Transact(opts, "initialize", anOwner, aPaymaster)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address anOwner, address aPaymaster) returns()
func (_TokenEntryPoint *TokenEntryPointSession) Initialize(anOwner common.Address, aPaymaster common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.Initialize(&_TokenEntryPoint.TransactOpts, anOwner, aPaymaster)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address anOwner, address aPaymaster) returns()
func (_TokenEntryPoint *TokenEntryPointTransactorSession) Initialize(anOwner common.Address, aPaymaster common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.Initialize(&_TokenEntryPoint.TransactOpts, anOwner, aPaymaster)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenEntryPoint *TokenEntryPointTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenEntryPoint.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenEntryPoint *TokenEntryPointSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.RenounceOwnership(&_TokenEntryPoint.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenEntryPoint *TokenEntryPointTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.RenounceOwnership(&_TokenEntryPoint.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenEntryPoint *TokenEntryPointTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenEntryPoint *TokenEntryPointSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.TransferOwnership(&_TokenEntryPoint.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenEntryPoint *TokenEntryPointTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.TransferOwnership(&_TokenEntryPoint.TransactOpts, newOwner)
}

// UpdatePaymaster is a paid mutator transaction binding the contract method 0xdaf6a3f6.
//
// Solidity: function updatePaymaster(address newPaymaster) returns()
func (_TokenEntryPoint *TokenEntryPointTransactor) UpdatePaymaster(opts *bind.TransactOpts, newPaymaster common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.contract.Transact(opts, "updatePaymaster", newPaymaster)
}

// UpdatePaymaster is a paid mutator transaction binding the contract method 0xdaf6a3f6.
//
// Solidity: function updatePaymaster(address newPaymaster) returns()
func (_TokenEntryPoint *TokenEntryPointSession) UpdatePaymaster(newPaymaster common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.UpdatePaymaster(&_TokenEntryPoint.TransactOpts, newPaymaster)
}

// UpdatePaymaster is a paid mutator transaction binding the contract method 0xdaf6a3f6.
//
// Solidity: function updatePaymaster(address newPaymaster) returns()
func (_TokenEntryPoint *TokenEntryPointTransactorSession) UpdatePaymaster(newPaymaster common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.UpdatePaymaster(&_TokenEntryPoint.TransactOpts, newPaymaster)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TokenEntryPoint *TokenEntryPointTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TokenEntryPoint *TokenEntryPointSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.UpgradeTo(&_TokenEntryPoint.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TokenEntryPoint *TokenEntryPointTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.UpgradeTo(&_TokenEntryPoint.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TokenEntryPoint *TokenEntryPointTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TokenEntryPoint.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TokenEntryPoint *TokenEntryPointSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.UpgradeToAndCall(&_TokenEntryPoint.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TokenEntryPoint *TokenEntryPointTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.UpgradeToAndCall(&_TokenEntryPoint.TransactOpts, newImplementation, data)
}

// TokenEntryPointAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the TokenEntryPoint contract.
type TokenEntryPointAdminChangedIterator struct {
	Event *TokenEntryPointAdminChanged // Event containing the contract specifics and raw log

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
func (it *TokenEntryPointAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEntryPointAdminChanged)
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
		it.Event = new(TokenEntryPointAdminChanged)
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
func (it *TokenEntryPointAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEntryPointAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEntryPointAdminChanged represents a AdminChanged event raised by the TokenEntryPoint contract.
type TokenEntryPointAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TokenEntryPoint *TokenEntryPointFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*TokenEntryPointAdminChangedIterator, error) {

	logs, sub, err := _TokenEntryPoint.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &TokenEntryPointAdminChangedIterator{contract: _TokenEntryPoint.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TokenEntryPoint *TokenEntryPointFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *TokenEntryPointAdminChanged) (event.Subscription, error) {

	logs, sub, err := _TokenEntryPoint.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEntryPointAdminChanged)
				if err := _TokenEntryPoint.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_TokenEntryPoint *TokenEntryPointFilterer) ParseAdminChanged(log types.Log) (*TokenEntryPointAdminChanged, error) {
	event := new(TokenEntryPointAdminChanged)
	if err := _TokenEntryPoint.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenEntryPointBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the TokenEntryPoint contract.
type TokenEntryPointBeaconUpgradedIterator struct {
	Event *TokenEntryPointBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *TokenEntryPointBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEntryPointBeaconUpgraded)
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
		it.Event = new(TokenEntryPointBeaconUpgraded)
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
func (it *TokenEntryPointBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEntryPointBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEntryPointBeaconUpgraded represents a BeaconUpgraded event raised by the TokenEntryPoint contract.
type TokenEntryPointBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TokenEntryPoint *TokenEntryPointFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*TokenEntryPointBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _TokenEntryPoint.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &TokenEntryPointBeaconUpgradedIterator{contract: _TokenEntryPoint.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TokenEntryPoint *TokenEntryPointFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *TokenEntryPointBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _TokenEntryPoint.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEntryPointBeaconUpgraded)
				if err := _TokenEntryPoint.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_TokenEntryPoint *TokenEntryPointFilterer) ParseBeaconUpgraded(log types.Log) (*TokenEntryPointBeaconUpgraded, error) {
	event := new(TokenEntryPointBeaconUpgraded)
	if err := _TokenEntryPoint.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenEntryPointInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TokenEntryPoint contract.
type TokenEntryPointInitializedIterator struct {
	Event *TokenEntryPointInitialized // Event containing the contract specifics and raw log

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
func (it *TokenEntryPointInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEntryPointInitialized)
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
		it.Event = new(TokenEntryPointInitialized)
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
func (it *TokenEntryPointInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEntryPointInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEntryPointInitialized represents a Initialized event raised by the TokenEntryPoint contract.
type TokenEntryPointInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenEntryPoint *TokenEntryPointFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenEntryPointInitializedIterator, error) {

	logs, sub, err := _TokenEntryPoint.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenEntryPointInitializedIterator{contract: _TokenEntryPoint.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenEntryPoint *TokenEntryPointFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenEntryPointInitialized) (event.Subscription, error) {

	logs, sub, err := _TokenEntryPoint.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEntryPointInitialized)
				if err := _TokenEntryPoint.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TokenEntryPoint *TokenEntryPointFilterer) ParseInitialized(log types.Log) (*TokenEntryPointInitialized, error) {
	event := new(TokenEntryPointInitialized)
	if err := _TokenEntryPoint.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenEntryPointOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenEntryPoint contract.
type TokenEntryPointOwnershipTransferredIterator struct {
	Event *TokenEntryPointOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenEntryPointOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEntryPointOwnershipTransferred)
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
		it.Event = new(TokenEntryPointOwnershipTransferred)
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
func (it *TokenEntryPointOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEntryPointOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEntryPointOwnershipTransferred represents a OwnershipTransferred event raised by the TokenEntryPoint contract.
type TokenEntryPointOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenEntryPoint *TokenEntryPointFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenEntryPointOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenEntryPoint.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenEntryPointOwnershipTransferredIterator{contract: _TokenEntryPoint.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenEntryPoint *TokenEntryPointFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenEntryPointOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenEntryPoint.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEntryPointOwnershipTransferred)
				if err := _TokenEntryPoint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenEntryPoint *TokenEntryPointFilterer) ParseOwnershipTransferred(log types.Log) (*TokenEntryPointOwnershipTransferred, error) {
	event := new(TokenEntryPointOwnershipTransferred)
	if err := _TokenEntryPoint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenEntryPointUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TokenEntryPoint contract.
type TokenEntryPointUpgradedIterator struct {
	Event *TokenEntryPointUpgraded // Event containing the contract specifics and raw log

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
func (it *TokenEntryPointUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEntryPointUpgraded)
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
		it.Event = new(TokenEntryPointUpgraded)
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
func (it *TokenEntryPointUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEntryPointUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEntryPointUpgraded represents a Upgraded event raised by the TokenEntryPoint contract.
type TokenEntryPointUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TokenEntryPoint *TokenEntryPointFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TokenEntryPointUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TokenEntryPoint.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TokenEntryPointUpgradedIterator{contract: _TokenEntryPoint.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TokenEntryPoint *TokenEntryPointFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TokenEntryPointUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TokenEntryPoint.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEntryPointUpgraded)
				if err := _TokenEntryPoint.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TokenEntryPoint *TokenEntryPointFilterer) ParseUpgraded(log types.Log) (*TokenEntryPointUpgraded, error) {
	event := new(TokenEntryPointUpgraded)
	if err := _TokenEntryPoint.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
