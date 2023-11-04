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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"_userOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100e1565b600054610100900460ff161561008e5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116146100df576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b608051611d6e610118600039600081816102df0152818161031f015281816104000152818161044001526104d30152611d6e6000f3fe6080604052600436106100865760003560e01c806352d1902d1161005957806352d1902d14610113578063715018a6146101285780638da5cb5b1461013d578063c4d66de814610165578063f2fde38b1461018557600080fd5b80631fad948c1461008b5780633659cfe6146100ad5780634564cc2e146100cd5780634f1ef28614610100575b600080fd5b34801561009757600080fd5b506100ab6100a6366004611631565b6101a5565b005b3480156100b957600080fd5b506100ab6100c83660046116c7565b6102d5565b3480156100d957600080fd5b506100ed6100e83660046116e4565b6103b4565b6040519081526020015b60405180910390f35b6100ab61010e36600461178f565b6103f6565b34801561011f57600080fd5b506100ed6104c6565b34801561013457600080fd5b506100ab610579565b34801561014957600080fd5b506065546040516001600160a01b0390911681526020016100f7565b34801561017157600080fd5b506100ab6101803660046116c7565b61058d565b34801561019157600080fd5b506100ab6101a03660046116c7565b6106af565b6101ad610725565b816102095760405162461bcd60e51b815260206004820152602160248201527f6e65656473206174206c65617374206f6e652075736572206f7065726174696f6044820152603760f91b60648201526084015b60405180910390fd5b8160005b818110156102c5573685858381811061022857610228611822565b905060200281019061023a9190611838565b9050803560006102498361077e565b9050610256838383610808565b61026083836108c8565b61026a83826109fe565b6102b782600061027d6060870187611859565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610a7692505050565b83600101935050505061020d565b50506102d060018055565b505050565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361031d5760405162461bcd60e51b8152600401610200906118a0565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610366600080516020611cf2833981519152546001600160a01b031690565b6001600160a01b03161461038c5760405162461bcd60e51b8152600401610200906118ec565b61039581610af3565b604080516000808252602082019092526103b191839190610afb565b50565b60006103bf82610c66565b6040805160208101929092523090820152466060820152608001604051602081830303815290604052805190602001209050919050565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361043e5760405162461bcd60e51b8152600401610200906118a0565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610487600080516020611cf2833981519152546001600160a01b031690565b6001600160a01b0316146104ad5760405162461bcd60e51b8152600401610200906118ec565b6104b682610af3565b6104c282826001610afb565b5050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105665760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610200565b50600080516020611cf283398151915290565b610581610c7f565b61058b6000610cd9565b565b600054610100900460ff16158080156105ad5750600054600160ff909116105b806105c75750303b1580156105c7575060005460ff166001145b61062a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610200565b6000805460ff19166001179055801561064d576000805461ff0019166101001790555b610655610d2b565b61065d610d5a565b61066682610d89565b80156104c2576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6106b7610c7f565b6001600160a01b03811661071c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610200565b6103b181610cd9565b6002600154036107775760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610200565b6002600155565b60003681610790610120850185611859565b909250905060148110156107e65760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964207061796d6173746572416e644461746100000000000000006044820152606401610200565b6107f4601460008385611938565b6107fd91611962565b60601c949350505050565b604051631aab3f0d60e11b81526001600160a01b03838116600483015260006024830181905291908316906335567e1a90604401602060405180830381865afa158015610859573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061087d9190611997565b9050808460200135146108c25760405162461bcd60e51b815260206004820152600d60248201526c696e76616c6964206e6f6e636560981b6044820152606401610200565b50505050565b60006109096108d6846103b4565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c91909152603c902090565b90508260200135600003610921576109218382610d92565b81630b135d3f60e11b6001600160a01b038216631626ba7e84610948610140890189611859565b6040518463ffffffff1660e01b8152600401610966939291906119d9565b602060405180830381865afa158015610983573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109a791906119fc565b6001600160e01b031916146108c25760405162461bcd60e51b815260206004820152601960248201527f696e76616c6964206163636f756e74207369676e6174757265000000000000006044820152606401610200565b604051637a32e3bf60e11b81526001600160a01b0382169063f465c77e90610a2f9085906000908190600401611a6c565b6000604051808303816000875af1158015610a4e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526108c29190810190611bad565b600080846001600160a01b03168484604051610a929190611c2e565b60006040518083038185875af1925050503d8060008114610acf576040519150601f19603f3d011682016040523d82523d6000602084013e610ad4565b606091505b509150915081610ae657805160208201fd5b5050505050565b60018055565b6103b1610c7f565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610b2e576102d083611018565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610b88575060408051601f3d908101601f19168201909252610b8591810190611997565b60015b610beb5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610200565b600080516020611cf28339815191528114610c5a5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610200565b506102d08383836110b4565b6000610c71826110d9565b805190602001209050919050565b6065546001600160a01b0316331461058b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610200565b606580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16610d525760405162461bcd60e51b815260040161020090611c40565b61058b6111ac565b600054610100900460ff16610d815760405162461bcd60e51b815260040161020090611c40565b61058b6111dc565b6103b1816106af565b366000610da26040850185611859565b90925090506014811015610deb5760405162461bcd60e51b815260206004820152601060248201526f696e76616c696420696e6974436f646560801b6044820152606401610200565b6000610dfa6014828486611938565b610e0391611962565b60601c9050803b610e485760405162461bcd60e51b815260206004820152600f60248201526e696e76616c696420666163746f727960881b6044820152606401610200565b806000610e99610e5c610140890189611859565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a939250506112039050565b9050863560405163119709c360e31b81526001600160a01b0383811660048301526000602483015291821691841690638cb84e1890604401602060405180830381865afa158015610eee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f129190611c8b565b6001600160a01b031614610f685760405162461bcd60e51b815260206004820152601a60248201527f666163746f7279206d7573742072657475726e2073656e6465720000000000006044820152606401610200565b366000610f78866014818a611938565b91509150610fbe85600084848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610a7692505050565b88353b61100d5760405162461bcd60e51b815260206004820152601e60248201527f696e76616c6964206163636f756e7420696e697469616c697a6174696f6e00006044820152606401610200565b505050505050505050565b6001600160a01b0381163b6110855760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610200565b600080516020611cf283398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6110bd83611227565b6000825111806110ca5750805b156102d0576108c28383611267565b60608135602083013560006110f96110f46040870187611859565b611293565b9050600061110d6110f46060880188611859565b9050608086013560a087013560c088013560e08901356101008a0135600061113c6110f46101208e018e611859565b604080516001600160a01b039c909c1660208d01528b81019a909a5260608b019890985250608089019590955260a088019390935260c087019190915260e08601526101008501526101208401526101408084019190915281518084039091018152610160909201905292915050565b600054610100900460ff166111d35760405162461bcd60e51b815260040161020090611c40565b61058b33610cd9565b600054610100900460ff16610aed5760405162461bcd60e51b815260040161020090611c40565b600080600061121285856112a6565b9150915061121f816112eb565b509392505050565b61123081611018565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b606061128c8383604051806060016040528060278152602001611d1260279139611435565b9392505050565b6000604051828085833790209392505050565b60008082516041036112dc5760208301516040840151606085015160001a6112d0878285856114ad565b945094505050506112e4565b506000905060025b9250929050565b60008160048111156112ff576112ff611ca8565b036113075750565b600181600481111561131b5761131b611ca8565b036113685760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610200565b600281600481111561137c5761137c611ca8565b036113c95760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610200565b60038160048111156113dd576113dd611ca8565b036103b15760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610200565b6060600080856001600160a01b0316856040516114529190611c2e565b600060405180830381855af49150503d806000811461148d576040519150601f19603f3d011682016040523d82523d6000602084013e611492565b606091505b50915091506114a386838387611571565b9695505050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156114e45750600090506003611568565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611538573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661156157600060019250925050611568565b9150600090505b94509492505050565b606083156115e05782516000036115d9576001600160a01b0385163b6115d95760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610200565b50816115ea565b6115ea83836115f2565b949350505050565b8151156116025781518083602001fd5b8060405162461bcd60e51b81526004016102009190611cbe565b6001600160a01b03811681146103b157600080fd5b60008060006040848603121561164657600080fd5b833567ffffffffffffffff8082111561165e57600080fd5b818601915086601f83011261167257600080fd5b81358181111561168157600080fd5b8760208260051b850101111561169657600080fd5b602092830195509350508401356116ac8161161c565b809150509250925092565b80356116c28161161c565b919050565b6000602082840312156116d957600080fd5b813561128c8161161c565b6000602082840312156116f657600080fd5b813567ffffffffffffffff81111561170d57600080fd5b8201610160818503121561128c57600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561175f5761175f611720565b604052919050565b600067ffffffffffffffff82111561178157611781611720565b50601f01601f191660200190565b600080604083850312156117a257600080fd5b82356117ad8161161c565b9150602083013567ffffffffffffffff8111156117c957600080fd5b8301601f810185136117da57600080fd5b80356117ed6117e882611767565b611736565b81815286602083850101111561180257600080fd5b816020840160208301376000602083830101528093505050509250929050565b634e487b7160e01b600052603260045260246000fd5b6000823561015e1983360301811261184f57600080fd5b9190910192915050565b6000808335601e1984360301811261187057600080fd5b83018035915067ffffffffffffffff82111561188b57600080fd5b6020019150368190038213156112e457600080fd5b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b6000808585111561194857600080fd5b8386111561195557600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff19813581811691601485101561198f5780818660140360031b1b83161692505b505092915050565b6000602082840312156119a957600080fd5b5051919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8381526040602082015260006119f36040830184866119b0565b95945050505050565b600060208284031215611a0e57600080fd5b81516001600160e01b03198116811461128c57600080fd5b6000808335601e19843603018112611a3d57600080fd5b830160208101925035905067ffffffffffffffff811115611a5d57600080fd5b8036038213156112e457600080fd5b60608152611a8d60608201611a80866116b7565b6001600160a01b03169052565b602084013560808201526000611aa66040860186611a26565b6101608060a0860152611abe6101c0860183856119b0565b9250611acd6060890189611a26565b9250605f19808786030160c0880152611ae78585846119b0565b945060808a013560e0880152610100935060a08a013584880152610120915060c08a01358288015261014060e08b013581890152848b013584890152611b2f838c018c611a26565b955093508188870301610180890152611b498686866119b0565b9550611b57818c018c611a26565b955093505080878603016101a08801525050611b748383836119b0565b60208601979097525050505060400152919050565b60005b83811015611ba4578181015183820152602001611b8c565b50506000910152565b60008060408385031215611bc057600080fd5b825167ffffffffffffffff811115611bd757600080fd5b8301601f81018513611be857600080fd5b8051611bf66117e882611767565b818152866020838501011115611c0b57600080fd5b611c1c826020830160208601611b89565b60209590950151949694955050505050565b6000825161184f818460208701611b89565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b600060208284031215611c9d57600080fd5b815161128c8161161c565b634e487b7160e01b600052602160045260246000fd5b6020815260008251806020840152611cdd816040850160208701611b89565b601f01601f1916919091016040019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212202691dbe2fd999c7b0f689ef4b30298cf78a609567e7e7721dcc98a1e19635ee664736f6c63430008140033",
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

// UserOpHash is a free data retrieval call binding the contract method 0x4564cc2e.
//
// Solidity: function _userOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_TokenEntryPoint *TokenEntryPointCaller) UserOpHash(opts *bind.CallOpts, userOp UserOperation) ([32]byte, error) {
	var out []interface{}
	err := _TokenEntryPoint.contract.Call(opts, &out, "_userOpHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UserOpHash is a free data retrieval call binding the contract method 0x4564cc2e.
//
// Solidity: function _userOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_TokenEntryPoint *TokenEntryPointSession) UserOpHash(userOp UserOperation) ([32]byte, error) {
	return _TokenEntryPoint.Contract.UserOpHash(&_TokenEntryPoint.CallOpts, userOp)
}

// UserOpHash is a free data retrieval call binding the contract method 0x4564cc2e.
//
// Solidity: function _userOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_TokenEntryPoint *TokenEntryPointCallerSession) UserOpHash(userOp UserOperation) ([32]byte, error) {
	return _TokenEntryPoint.Contract.UserOpHash(&_TokenEntryPoint.CallOpts, userOp)
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

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_TokenEntryPoint *TokenEntryPointTransactor) Initialize(opts *bind.TransactOpts, anOwner common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.contract.Transact(opts, "initialize", anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_TokenEntryPoint *TokenEntryPointSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.Initialize(&_TokenEntryPoint.TransactOpts, anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_TokenEntryPoint *TokenEntryPointTransactorSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _TokenEntryPoint.Contract.Initialize(&_TokenEntryPoint.TransactOpts, anOwner)
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
