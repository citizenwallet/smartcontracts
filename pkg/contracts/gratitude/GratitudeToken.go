// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gratitude

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

// GratitudeMetaData contains all meta data concerning the Gratitude contract.
var GratitudeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"GratitudeTokenInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"}],\"name\":\"mintToMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintToMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"mintToMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60c0604052306080523480156200001557600080fd5b5060405162001eec38038062001eec833981016040819052620000389162000117565b6001600160a01b03811660a0526200004f62000056565b5062000149565b600054610100900460ff1615620000c35760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161462000115576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6000602082840312156200012a57600080fd5b81516001600160a01b03811681146200014257600080fd5b9392505050565b60805160a051611d5362000199600039600081816103f701528181610c52015261110501526000818161054b0152818161059401528181610672015281816106b201526107410152611d536000f3fe6080604052600436106101405760003560e01c806352d1902d116100b657806395d89b411161006f57806395d89b4114610393578063a457c2d7146103a8578063a9059cbb146103c8578063b0d691fe146103e8578063c4d66de81461041b578063dd62ed3e1461043b57600080fd5b806352d1902d146102b057806370a08231146102c557806379cc6790146102fb5780638a0b03601461031b5780638da5cb5b1461033b578063925247251461037357600080fd5b8063313ce56711610108578063313ce567146102015780633659cfe61461021d578063395093511461023d57806340c10f191461025d57806342966c681461027d5780634f1ef2861461029d57600080fd5b806306fdde0314610145578063095ea7b31461017057806318160ddd146101a0578063203abe34146101bf57806323b872dd146101e1575b600080fd5b34801561015157600080fd5b5061015a61045b565b604051610167919061164a565b60405180910390f35b34801561017c57600080fd5b5061019061018b366004611699565b6104ed565b6040519015158152602001610167565b3480156101ac57600080fd5b506035545b604051908152602001610167565b3480156101cb57600080fd5b506101df6101da3660046117a0565b610507565b005b3480156101ed57600080fd5b506101906101fc3660046117d5565b610524565b34801561020d57600080fd5b5060405160128152602001610167565b34801561022957600080fd5b506101df610238366004611811565b610541565b34801561024957600080fd5b50610190610258366004611699565b610626565b34801561026957600080fd5b506101df610278366004611699565b610648565b34801561028957600080fd5b506101df61029836600461182c565b61065e565b6101df6102ab366004611845565b610668565b3480156102bc57600080fd5b506101b1610734565b3480156102d157600080fd5b506101b16102e0366004611811565b6001600160a01b031660009081526033602052604090205490565b34801561030757600080fd5b506101df610316366004611699565b6107e7565b34801561032757600080fd5b506101df6103363660046118eb565b6107fc565b34801561034757600080fd5b5060fb5461035b906001600160a01b031681565b6040516001600160a01b039091168152602001610167565b34801561037f57600080fd5b506101df61038e366004611930565b610850565b34801561039f57600080fd5b5061015a610936565b3480156103b457600080fd5b506101906103c3366004611699565b610945565b3480156103d457600080fd5b506101906103e3366004611699565b6109cb565b3480156103f457600080fd5b507f000000000000000000000000000000000000000000000000000000000000000061035b565b34801561042757600080fd5b506101df610436366004611811565b6109e6565b34801561044757600080fd5b506101b16104563660046119eb565b610af8565b60606036805461046a90611a1e565b80601f016020809104026020016040519081016040528092919081815260200182805461049690611a1e565b80156104e35780601f106104b8576101008083540402835291602001916104e3565b820191906000526020600020905b8154815290600101906020018083116104c657829003601f168201915b5050505050905090565b6000336104fb818585610b23565b60019150505b92915050565b61050f610c47565b61052181670de0b6b3a76400006107fc565b50565b600061052e610c47565b610539848484610cd6565b949350505050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036105925760405162461bcd60e51b815260040161058990611a58565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166105db600080516020611cd7833981519152546001600160a01b031690565b6001600160a01b0316146106015760405162461bcd60e51b815260040161058990611aa4565b61060a81610cef565b6040805160008082526020820190925261052191839190610cf7565b6000336104fb8185856106398383610af8565b6106439190611b06565b610b23565b610650610c47565b61065a8282610e62565b5050565b6105213382610f23565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036106b05760405162461bcd60e51b815260040161058990611a58565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166106f9600080516020611cd7833981519152546001600160a01b031690565b6001600160a01b03161461071f5760405162461bcd60e51b815260040161058990611aa4565b61072882610cef565b61065a82826001610cf7565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107d45760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610589565b50600080516020611cd783398151915290565b6107f2823383611057565b61065a8282610f23565b610804610c47565b60005b825181101561084b57600083828151811061082457610824611b19565b602002602001015190506108388184610e62565b508061084381611b2f565b915050610807565b505050565b610858610c47565b80518251146108cf5760405162461bcd60e51b815260206004820152603760248201527f726563697069656e747320616e6420616d6f756e747320617272617973206d7560448201527f73742068617665207468652073616d65206c656e6774680000000000000000006064820152608401610589565b60005b825181101561084b5760008382815181106108ef576108ef611b19565b60200260200101519050600083838151811061090d5761090d611b19565b602002602001015190506109218282610e62565b5050808061092e90611b2f565b9150506108d2565b60606037805461046a90611a1e565b600033816109538286610af8565b9050838110156109b35760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608401610589565b6109c08286868403610b23565b506001949350505050565b60006109d5610c47565b6109df83836110d1565b9392505050565b600054610100900460ff1615808015610a065750600054600160ff909116105b80610a205750303b158015610a20575060005460ff166001145b610a835760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610589565b6000805460ff191660011790558015610aa6576000805461ff0019166101001790555b610aaf826110df565b801561065a576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6001600160a01b03918216600090815260346020908152604080832093909416825291909152205490565b6001600160a01b038316610b855760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610589565b6001600160a01b038216610be65760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610589565b6001600160a01b0383811660008181526034602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610c88575060fb546001600160a01b031633145b610cd45760405162461bcd60e51b815260206004820181905260248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152606401610589565b565b600033610ce4858285611057565b6109c08585856111a7565b610521610c47565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610d2a5761084b83611352565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610d84575060408051601f3d908101601f19168201909252610d8191810190611b48565b60015b610de75760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610589565b600080516020611cd78339815191528114610e565760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610589565b5061084b8383836113ee565b6001600160a01b038216610eb85760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610589565b8060356000828254610eca9190611b06565b90915550506001600160a01b0382166000818152603360209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6001600160a01b038216610f835760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610589565b6001600160a01b03821660009081526033602052604090205481811015610ff75760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610589565b6001600160a01b03831660008181526033602090815260408083208686039055603580548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b60006110638484610af8565b905060001981146110cb57818110156110be5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610589565b6110cb8484848403610b23565b50505050565b6000336104fb8185856111a7565b60fb80546001600160a01b0319166001600160a01b0383811691821790925560405190917f000000000000000000000000000000000000000000000000000000000000000016907fe90b83f24696d649f6cdd7a0b35764256ae51f4998f6b266afb6fc6f2d50c60290600090a36111976040518060400160405280600e81526020016d23b930ba34ba3ab232aa37b5b2b760911b8152506040518060400160405280600381526020016211d51560ea1b815250611413565b61119f611444565b610521611444565b6001600160a01b03831661120b5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610589565b6001600160a01b03821661126d5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610589565b6001600160a01b038316600090815260336020526040902054818110156112e55760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610589565b6001600160a01b0380851660008181526033602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906113459086815260200190565b60405180910390a36110cb565b6001600160a01b0381163b6113bf5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610589565b600080516020611cd783398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6113f78361146b565b6000825111806114045750805b1561084b576110cb83836114ab565b600054610100900460ff1661143a5760405162461bcd60e51b815260040161058990611b61565b61065a82826114d0565b600054610100900460ff16610cd45760405162461bcd60e51b815260040161058990611b61565b61147481611352565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606109df8383604051806060016040528060278152602001611cf760279139611510565b600054610100900460ff166114f75760405162461bcd60e51b815260040161058990611b61565b60366115038382611bfa565b50603761084b8282611bfa565b6060600080856001600160a01b03168560405161152d9190611cba565b600060405180830381855af49150503d8060008114611568576040519150601f19603f3d011682016040523d82523d6000602084013e61156d565b606091505b509150915061157e86838387611588565b9695505050505050565b606083156115f75782516000036115f0576001600160a01b0385163b6115f05760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610589565b5081610539565b610539838381511561160c5781518083602001fd5b8060405162461bcd60e51b8152600401610589919061164a565b60005b83811015611641578181015183820152602001611629565b50506000910152565b6020815260008251806020840152611669816040850160208701611626565b601f01601f19169190910160400192915050565b80356001600160a01b038116811461169457600080fd5b919050565b600080604083850312156116ac57600080fd5b6116b58361167d565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611702576117026116c3565b604052919050565b600067ffffffffffffffff821115611724576117246116c3565b5060051b60200190565b600082601f83011261173f57600080fd5b8135602061175461174f8361170a565b6116d9565b82815260059290921b8401810191818101908684111561177357600080fd5b8286015b84811015611795576117888161167d565b8352918301918301611777565b509695505050505050565b6000602082840312156117b257600080fd5b813567ffffffffffffffff8111156117c957600080fd5b6105398482850161172e565b6000806000606084860312156117ea57600080fd5b6117f38461167d565b92506118016020850161167d565b9150604084013590509250925092565b60006020828403121561182357600080fd5b6109df8261167d565b60006020828403121561183e57600080fd5b5035919050565b6000806040838503121561185857600080fd5b6118618361167d565b915060208084013567ffffffffffffffff8082111561187f57600080fd5b818601915086601f83011261189357600080fd5b8135818111156118a5576118a56116c3565b6118b7601f8201601f191685016116d9565b915080825287848285010111156118cd57600080fd5b80848401858401376000848284010152508093505050509250929050565b600080604083850312156118fe57600080fd5b823567ffffffffffffffff81111561191557600080fd5b6119218582860161172e565b95602094909401359450505050565b6000806040838503121561194357600080fd5b823567ffffffffffffffff8082111561195b57600080fd5b6119678683870161172e565b935060209150818501358181111561197e57600080fd5b85019050601f8101861361199157600080fd5b803561199f61174f8261170a565b81815260059190911b820183019083810190888311156119be57600080fd5b928401925b828410156119dc578335825292840192908401906119c3565b80955050505050509250929050565b600080604083850312156119fe57600080fd5b611a078361167d565b9150611a156020840161167d565b90509250929050565b600181811c90821680611a3257607f821691505b602082108103611a5257634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b8082018082111561050157610501611af0565b634e487b7160e01b600052603260045260246000fd5b600060018201611b4157611b41611af0565b5060010190565b600060208284031215611b5a57600080fd5b5051919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b601f82111561084b57600081815260208120601f850160051c81016020861015611bd35750805b601f850160051c820191505b81811015611bf257828155600101611bdf565b505050505050565b815167ffffffffffffffff811115611c1457611c146116c3565b611c2881611c228454611a1e565b84611bac565b602080601f831160018114611c5d5760008415611c455750858301515b600019600386901b1c1916600185901b178555611bf2565b600085815260208120601f198616915b82811015611c8c57888601518255948401946001909101908401611c6d565b5085821015611caa5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008251611ccc818460208701611626565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220bdf83fb6744cb1acb86923b9c456ce1c8d8bd023a107266efdb69997ea8351bf64736f6c63430008140033",
}

// GratitudeABI is the input ABI used to generate the binding from.
// Deprecated: Use GratitudeMetaData.ABI instead.
var GratitudeABI = GratitudeMetaData.ABI

// GratitudeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GratitudeMetaData.Bin instead.
var GratitudeBin = GratitudeMetaData.Bin

// DeployGratitude deploys a new Ethereum contract, binding an instance of Gratitude to it.
func DeployGratitude(auth *bind.TransactOpts, backend bind.ContractBackend, anEntryPoint common.Address) (common.Address, *types.Transaction, *Gratitude, error) {
	parsed, err := GratitudeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GratitudeBin), backend, anEntryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gratitude{GratitudeCaller: GratitudeCaller{contract: contract}, GratitudeTransactor: GratitudeTransactor{contract: contract}, GratitudeFilterer: GratitudeFilterer{contract: contract}}, nil
}

// Gratitude is an auto generated Go binding around an Ethereum contract.
type Gratitude struct {
	GratitudeCaller     // Read-only binding to the contract
	GratitudeTransactor // Write-only binding to the contract
	GratitudeFilterer   // Log filterer for contract events
}

// GratitudeCaller is an auto generated read-only Go binding around an Ethereum contract.
type GratitudeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GratitudeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GratitudeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GratitudeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GratitudeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GratitudeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GratitudeSession struct {
	Contract     *Gratitude        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GratitudeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GratitudeCallerSession struct {
	Contract *GratitudeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GratitudeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GratitudeTransactorSession struct {
	Contract     *GratitudeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GratitudeRaw is an auto generated low-level Go binding around an Ethereum contract.
type GratitudeRaw struct {
	Contract *Gratitude // Generic contract binding to access the raw methods on
}

// GratitudeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GratitudeCallerRaw struct {
	Contract *GratitudeCaller // Generic read-only contract binding to access the raw methods on
}

// GratitudeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GratitudeTransactorRaw struct {
	Contract *GratitudeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGratitude creates a new instance of Gratitude, bound to a specific deployed contract.
func NewGratitude(address common.Address, backend bind.ContractBackend) (*Gratitude, error) {
	contract, err := bindGratitude(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gratitude{GratitudeCaller: GratitudeCaller{contract: contract}, GratitudeTransactor: GratitudeTransactor{contract: contract}, GratitudeFilterer: GratitudeFilterer{contract: contract}}, nil
}

// NewGratitudeCaller creates a new read-only instance of Gratitude, bound to a specific deployed contract.
func NewGratitudeCaller(address common.Address, caller bind.ContractCaller) (*GratitudeCaller, error) {
	contract, err := bindGratitude(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GratitudeCaller{contract: contract}, nil
}

// NewGratitudeTransactor creates a new write-only instance of Gratitude, bound to a specific deployed contract.
func NewGratitudeTransactor(address common.Address, transactor bind.ContractTransactor) (*GratitudeTransactor, error) {
	contract, err := bindGratitude(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GratitudeTransactor{contract: contract}, nil
}

// NewGratitudeFilterer creates a new log filterer instance of Gratitude, bound to a specific deployed contract.
func NewGratitudeFilterer(address common.Address, filterer bind.ContractFilterer) (*GratitudeFilterer, error) {
	contract, err := bindGratitude(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GratitudeFilterer{contract: contract}, nil
}

// bindGratitude binds a generic wrapper to an already deployed contract.
func bindGratitude(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GratitudeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gratitude *GratitudeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gratitude.Contract.GratitudeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gratitude *GratitudeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gratitude.Contract.GratitudeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gratitude *GratitudeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gratitude.Contract.GratitudeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gratitude *GratitudeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gratitude.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gratitude *GratitudeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gratitude.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gratitude *GratitudeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gratitude.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Gratitude *GratitudeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gratitude.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Gratitude *GratitudeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Gratitude.Contract.Allowance(&_Gratitude.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Gratitude *GratitudeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Gratitude.Contract.Allowance(&_Gratitude.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gratitude *GratitudeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gratitude.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gratitude *GratitudeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Gratitude.Contract.BalanceOf(&_Gratitude.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gratitude *GratitudeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Gratitude.Contract.BalanceOf(&_Gratitude.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Gratitude *GratitudeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Gratitude.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Gratitude *GratitudeSession) Decimals() (uint8, error) {
	return _Gratitude.Contract.Decimals(&_Gratitude.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Gratitude *GratitudeCallerSession) Decimals() (uint8, error) {
	return _Gratitude.Contract.Decimals(&_Gratitude.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Gratitude *GratitudeCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gratitude.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Gratitude *GratitudeSession) EntryPoint() (common.Address, error) {
	return _Gratitude.Contract.EntryPoint(&_Gratitude.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Gratitude *GratitudeCallerSession) EntryPoint() (common.Address, error) {
	return _Gratitude.Contract.EntryPoint(&_Gratitude.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Gratitude *GratitudeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Gratitude.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Gratitude *GratitudeSession) Name() (string, error) {
	return _Gratitude.Contract.Name(&_Gratitude.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Gratitude *GratitudeCallerSession) Name() (string, error) {
	return _Gratitude.Contract.Name(&_Gratitude.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gratitude *GratitudeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gratitude.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gratitude *GratitudeSession) Owner() (common.Address, error) {
	return _Gratitude.Contract.Owner(&_Gratitude.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gratitude *GratitudeCallerSession) Owner() (common.Address, error) {
	return _Gratitude.Contract.Owner(&_Gratitude.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Gratitude *GratitudeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Gratitude.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Gratitude *GratitudeSession) ProxiableUUID() ([32]byte, error) {
	return _Gratitude.Contract.ProxiableUUID(&_Gratitude.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Gratitude *GratitudeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Gratitude.Contract.ProxiableUUID(&_Gratitude.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Gratitude *GratitudeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Gratitude.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Gratitude *GratitudeSession) Symbol() (string, error) {
	return _Gratitude.Contract.Symbol(&_Gratitude.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Gratitude *GratitudeCallerSession) Symbol() (string, error) {
	return _Gratitude.Contract.Symbol(&_Gratitude.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gratitude *GratitudeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gratitude.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gratitude *GratitudeSession) TotalSupply() (*big.Int, error) {
	return _Gratitude.Contract.TotalSupply(&_Gratitude.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gratitude *GratitudeCallerSession) TotalSupply() (*big.Int, error) {
	return _Gratitude.Contract.TotalSupply(&_Gratitude.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Gratitude *GratitudeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Gratitude *GratitudeSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.Approve(&_Gratitude.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Gratitude *GratitudeTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.Approve(&_Gratitude.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Gratitude *GratitudeTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Gratitude *GratitudeSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.Burn(&_Gratitude.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Gratitude *GratitudeTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.Burn(&_Gratitude.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Gratitude *GratitudeTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Gratitude *GratitudeSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.BurnFrom(&_Gratitude.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Gratitude *GratitudeTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.BurnFrom(&_Gratitude.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Gratitude *GratitudeTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Gratitude.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Gratitude *GratitudeSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.DecreaseAllowance(&_Gratitude.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Gratitude *GratitudeTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.DecreaseAllowance(&_Gratitude.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Gratitude *GratitudeTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Gratitude.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Gratitude *GratitudeSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.IncreaseAllowance(&_Gratitude.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Gratitude *GratitudeTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.IncreaseAllowance(&_Gratitude.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_Gratitude *GratitudeTransactor) Initialize(opts *bind.TransactOpts, anOwner common.Address) (*types.Transaction, error) {
	return _Gratitude.contract.Transact(opts, "initialize", anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_Gratitude *GratitudeSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _Gratitude.Contract.Initialize(&_Gratitude.TransactOpts, anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_Gratitude *GratitudeTransactorSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _Gratitude.Contract.Initialize(&_Gratitude.TransactOpts, anOwner)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Gratitude *GratitudeTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Gratitude *GratitudeSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.Mint(&_Gratitude.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Gratitude *GratitudeTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.Mint(&_Gratitude.TransactOpts, to, amount)
}

// MintToMany is a paid mutator transaction binding the contract method 0x203abe34.
//
// Solidity: function mintToMany(address[] recipients) returns()
func (_Gratitude *GratitudeTransactor) MintToMany(opts *bind.TransactOpts, recipients []common.Address) (*types.Transaction, error) {
	return _Gratitude.contract.Transact(opts, "mintToMany", recipients)
}

// MintToMany is a paid mutator transaction binding the contract method 0x203abe34.
//
// Solidity: function mintToMany(address[] recipients) returns()
func (_Gratitude *GratitudeSession) MintToMany(recipients []common.Address) (*types.Transaction, error) {
	return _Gratitude.Contract.MintToMany(&_Gratitude.TransactOpts, recipients)
}

// MintToMany is a paid mutator transaction binding the contract method 0x203abe34.
//
// Solidity: function mintToMany(address[] recipients) returns()
func (_Gratitude *GratitudeTransactorSession) MintToMany(recipients []common.Address) (*types.Transaction, error) {
	return _Gratitude.Contract.MintToMany(&_Gratitude.TransactOpts, recipients)
}

// MintToMany0 is a paid mutator transaction binding the contract method 0x8a0b0360.
//
// Solidity: function mintToMany(address[] recipients, uint256 amount) returns()
func (_Gratitude *GratitudeTransactor) MintToMany0(opts *bind.TransactOpts, recipients []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.contract.Transact(opts, "mintToMany0", recipients, amount)
}

// MintToMany0 is a paid mutator transaction binding the contract method 0x8a0b0360.
//
// Solidity: function mintToMany(address[] recipients, uint256 amount) returns()
func (_Gratitude *GratitudeSession) MintToMany0(recipients []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.MintToMany0(&_Gratitude.TransactOpts, recipients, amount)
}

// MintToMany0 is a paid mutator transaction binding the contract method 0x8a0b0360.
//
// Solidity: function mintToMany(address[] recipients, uint256 amount) returns()
func (_Gratitude *GratitudeTransactorSession) MintToMany0(recipients []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.MintToMany0(&_Gratitude.TransactOpts, recipients, amount)
}

// MintToMany1 is a paid mutator transaction binding the contract method 0x92524725.
//
// Solidity: function mintToMany(address[] recipients, uint256[] amounts) returns()
func (_Gratitude *GratitudeTransactor) MintToMany1(opts *bind.TransactOpts, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Gratitude.contract.Transact(opts, "mintToMany1", recipients, amounts)
}

// MintToMany1 is a paid mutator transaction binding the contract method 0x92524725.
//
// Solidity: function mintToMany(address[] recipients, uint256[] amounts) returns()
func (_Gratitude *GratitudeSession) MintToMany1(recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.MintToMany1(&_Gratitude.TransactOpts, recipients, amounts)
}

// MintToMany1 is a paid mutator transaction binding the contract method 0x92524725.
//
// Solidity: function mintToMany(address[] recipients, uint256[] amounts) returns()
func (_Gratitude *GratitudeTransactorSession) MintToMany1(recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.MintToMany1(&_Gratitude.TransactOpts, recipients, amounts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Gratitude *GratitudeTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Gratitude *GratitudeSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.Transfer(&_Gratitude.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Gratitude *GratitudeTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.Transfer(&_Gratitude.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Gratitude *GratitudeTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Gratitude *GratitudeSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.TransferFrom(&_Gratitude.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Gratitude *GratitudeTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gratitude.Contract.TransferFrom(&_Gratitude.TransactOpts, sender, recipient, amount)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Gratitude *GratitudeTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Gratitude.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Gratitude *GratitudeSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Gratitude.Contract.UpgradeTo(&_Gratitude.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Gratitude *GratitudeTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Gratitude.Contract.UpgradeTo(&_Gratitude.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Gratitude *GratitudeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Gratitude.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Gratitude *GratitudeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Gratitude.Contract.UpgradeToAndCall(&_Gratitude.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Gratitude *GratitudeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Gratitude.Contract.UpgradeToAndCall(&_Gratitude.TransactOpts, newImplementation, data)
}

// GratitudeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Gratitude contract.
type GratitudeAdminChangedIterator struct {
	Event *GratitudeAdminChanged // Event containing the contract specifics and raw log

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
func (it *GratitudeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GratitudeAdminChanged)
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
		it.Event = new(GratitudeAdminChanged)
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
func (it *GratitudeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GratitudeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GratitudeAdminChanged represents a AdminChanged event raised by the Gratitude contract.
type GratitudeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Gratitude *GratitudeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*GratitudeAdminChangedIterator, error) {

	logs, sub, err := _Gratitude.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &GratitudeAdminChangedIterator{contract: _Gratitude.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Gratitude *GratitudeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *GratitudeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Gratitude.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GratitudeAdminChanged)
				if err := _Gratitude.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Gratitude *GratitudeFilterer) ParseAdminChanged(log types.Log) (*GratitudeAdminChanged, error) {
	event := new(GratitudeAdminChanged)
	if err := _Gratitude.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GratitudeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Gratitude contract.
type GratitudeApprovalIterator struct {
	Event *GratitudeApproval // Event containing the contract specifics and raw log

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
func (it *GratitudeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GratitudeApproval)
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
		it.Event = new(GratitudeApproval)
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
func (it *GratitudeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GratitudeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GratitudeApproval represents a Approval event raised by the Gratitude contract.
type GratitudeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Gratitude *GratitudeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GratitudeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Gratitude.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GratitudeApprovalIterator{contract: _Gratitude.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Gratitude *GratitudeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GratitudeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Gratitude.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GratitudeApproval)
				if err := _Gratitude.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Gratitude *GratitudeFilterer) ParseApproval(log types.Log) (*GratitudeApproval, error) {
	event := new(GratitudeApproval)
	if err := _Gratitude.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GratitudeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Gratitude contract.
type GratitudeBeaconUpgradedIterator struct {
	Event *GratitudeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *GratitudeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GratitudeBeaconUpgraded)
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
		it.Event = new(GratitudeBeaconUpgraded)
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
func (it *GratitudeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GratitudeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GratitudeBeaconUpgraded represents a BeaconUpgraded event raised by the Gratitude contract.
type GratitudeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Gratitude *GratitudeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*GratitudeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Gratitude.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &GratitudeBeaconUpgradedIterator{contract: _Gratitude.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Gratitude *GratitudeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *GratitudeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Gratitude.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GratitudeBeaconUpgraded)
				if err := _Gratitude.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Gratitude *GratitudeFilterer) ParseBeaconUpgraded(log types.Log) (*GratitudeBeaconUpgraded, error) {
	event := new(GratitudeBeaconUpgraded)
	if err := _Gratitude.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GratitudeGratitudeTokenInitializedIterator is returned from FilterGratitudeTokenInitialized and is used to iterate over the raw logs and unpacked data for GratitudeTokenInitialized events raised by the Gratitude contract.
type GratitudeGratitudeTokenInitializedIterator struct {
	Event *GratitudeGratitudeTokenInitialized // Event containing the contract specifics and raw log

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
func (it *GratitudeGratitudeTokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GratitudeGratitudeTokenInitialized)
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
		it.Event = new(GratitudeGratitudeTokenInitialized)
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
func (it *GratitudeGratitudeTokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GratitudeGratitudeTokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GratitudeGratitudeTokenInitialized represents a GratitudeTokenInitialized event raised by the Gratitude contract.
type GratitudeGratitudeTokenInitialized struct {
	EntryPoint common.Address
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGratitudeTokenInitialized is a free log retrieval operation binding the contract event 0xe90b83f24696d649f6cdd7a0b35764256ae51f4998f6b266afb6fc6f2d50c602.
//
// Solidity: event GratitudeTokenInitialized(address indexed entryPoint, address indexed owner)
func (_Gratitude *GratitudeFilterer) FilterGratitudeTokenInitialized(opts *bind.FilterOpts, entryPoint []common.Address, owner []common.Address) (*GratitudeGratitudeTokenInitializedIterator, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Gratitude.contract.FilterLogs(opts, "GratitudeTokenInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &GratitudeGratitudeTokenInitializedIterator{contract: _Gratitude.contract, event: "GratitudeTokenInitialized", logs: logs, sub: sub}, nil
}

// WatchGratitudeTokenInitialized is a free log subscription operation binding the contract event 0xe90b83f24696d649f6cdd7a0b35764256ae51f4998f6b266afb6fc6f2d50c602.
//
// Solidity: event GratitudeTokenInitialized(address indexed entryPoint, address indexed owner)
func (_Gratitude *GratitudeFilterer) WatchGratitudeTokenInitialized(opts *bind.WatchOpts, sink chan<- *GratitudeGratitudeTokenInitialized, entryPoint []common.Address, owner []common.Address) (event.Subscription, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Gratitude.contract.WatchLogs(opts, "GratitudeTokenInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GratitudeGratitudeTokenInitialized)
				if err := _Gratitude.contract.UnpackLog(event, "GratitudeTokenInitialized", log); err != nil {
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

// ParseGratitudeTokenInitialized is a log parse operation binding the contract event 0xe90b83f24696d649f6cdd7a0b35764256ae51f4998f6b266afb6fc6f2d50c602.
//
// Solidity: event GratitudeTokenInitialized(address indexed entryPoint, address indexed owner)
func (_Gratitude *GratitudeFilterer) ParseGratitudeTokenInitialized(log types.Log) (*GratitudeGratitudeTokenInitialized, error) {
	event := new(GratitudeGratitudeTokenInitialized)
	if err := _Gratitude.contract.UnpackLog(event, "GratitudeTokenInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GratitudeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Gratitude contract.
type GratitudeInitializedIterator struct {
	Event *GratitudeInitialized // Event containing the contract specifics and raw log

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
func (it *GratitudeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GratitudeInitialized)
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
		it.Event = new(GratitudeInitialized)
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
func (it *GratitudeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GratitudeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GratitudeInitialized represents a Initialized event raised by the Gratitude contract.
type GratitudeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Gratitude *GratitudeFilterer) FilterInitialized(opts *bind.FilterOpts) (*GratitudeInitializedIterator, error) {

	logs, sub, err := _Gratitude.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GratitudeInitializedIterator{contract: _Gratitude.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Gratitude *GratitudeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GratitudeInitialized) (event.Subscription, error) {

	logs, sub, err := _Gratitude.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GratitudeInitialized)
				if err := _Gratitude.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Gratitude *GratitudeFilterer) ParseInitialized(log types.Log) (*GratitudeInitialized, error) {
	event := new(GratitudeInitialized)
	if err := _Gratitude.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GratitudeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Gratitude contract.
type GratitudeTransferIterator struct {
	Event *GratitudeTransfer // Event containing the contract specifics and raw log

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
func (it *GratitudeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GratitudeTransfer)
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
		it.Event = new(GratitudeTransfer)
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
func (it *GratitudeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GratitudeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GratitudeTransfer represents a Transfer event raised by the Gratitude contract.
type GratitudeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Gratitude *GratitudeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GratitudeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Gratitude.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GratitudeTransferIterator{contract: _Gratitude.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Gratitude *GratitudeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GratitudeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Gratitude.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GratitudeTransfer)
				if err := _Gratitude.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Gratitude *GratitudeFilterer) ParseTransfer(log types.Log) (*GratitudeTransfer, error) {
	event := new(GratitudeTransfer)
	if err := _Gratitude.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GratitudeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Gratitude contract.
type GratitudeUpgradedIterator struct {
	Event *GratitudeUpgraded // Event containing the contract specifics and raw log

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
func (it *GratitudeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GratitudeUpgraded)
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
		it.Event = new(GratitudeUpgraded)
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
func (it *GratitudeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GratitudeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GratitudeUpgraded represents a Upgraded event raised by the Gratitude contract.
type GratitudeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Gratitude *GratitudeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GratitudeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Gratitude.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GratitudeUpgradedIterator{contract: _Gratitude.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Gratitude *GratitudeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GratitudeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Gratitude.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GratitudeUpgraded)
				if err := _Gratitude.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Gratitude *GratitudeFilterer) ParseUpgraded(log types.Log) (*GratitudeUpgraded, error) {
	event := new(GratitudeUpgraded)
	if err := _Gratitude.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
