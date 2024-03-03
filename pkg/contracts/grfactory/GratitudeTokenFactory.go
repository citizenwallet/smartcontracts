// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package grfactory

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

// GrfactoryMetaData contains all meta data concerning the Grfactory contract.
var GrfactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"GratitudeTokenCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"createGratitudeToken\",\"outputs\":[{\"internalType\":\"contractGratitudeToken\",\"name\":\"ret\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getGratitudeTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gratitudeImplementation\",\"outputs\":[{\"internalType\":\"contractGratitudeToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161287938038061287983398101604081905261002f91610088565b8060405161003c9061007b565b6001600160a01b039091168152602001604051809103906000f080158015610068573d6000803e3d6000fd5b506001600160a01b0316608052506100b8565b611eec8061098d83390190565b60006020828403121561009a57600080fd5b81516001600160a01b03811681146100b157600080fd5b9392505050565b6080516108ad6100e060003960008181607a0152818161012101526101ee01526108ad6000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806301dfecf014610046578063c56ce58814610075578063d87c2ee51461009c575b600080fd5b6100596100543660046102c9565b6100af565b6040516001600160a01b03909116815260200160405180910390f35b6100597f000000000000000000000000000000000000000000000000000000000000000081565b6100596100aa3660046102c9565b6101af565b6000806100bc84846101af565b90506001600160a01b0381163b80156100d7575090506101a9565b6040516001600160a01b038616907f632ffdb1f5dbeb99afea447aa1f5f6cc4b23d0f61f074784c8c03a4134d7c15a90600090a26040516001600160a01b038616602482015284907f00000000000000000000000000000000000000000000000000000000000000009060440160408051601f198184030181529181526020820180516001600160e01b031663189acdbd60e31b17905251610178906102bc565b610183929190610325565b8190604051809103906000f59050801580156101a3573d6000803e3d6000fd5b50925050505b92915050565b60006102838260001b604051806020016101c8906102bc565b601f1982820381018352601f9091011660408190526001600160a01b03871660248201527f00000000000000000000000000000000000000000000000000000000000000009060440160408051601f19818403018152918152602080830180516001600160e01b031663189acdbd60e31b179052905161024a93929101610325565b60408051601f19818403018152908290526102689291602001610367565b6040516020818303038152906040528051906020012061028a565b9392505050565b60006102838383306000604051836040820152846020820152828152600b8101905060ff815360559020949350505050565b6104e18061039783390190565b600080604083850312156102dc57600080fd5b82356001600160a01b03811681146102f357600080fd5b946020939093013593505050565b60005b8381101561031c578181015183820152602001610304565b50506000910152565b60018060a01b03831681526040602082015260008251806040840152610352816060850160208701610301565b601f01601f1916919091016060019392505050565b60008351610379818460208801610301565b83519083019061038d818360208801610301565b0194935050505056fe60806040526040516104e13803806104e1833981016040819052610022916102de565b61002e82826000610035565b50506103fb565b61003e83610061565b60008251118061004b5750805b1561005c5761005a83836100a1565b505b505050565b61006a816100cd565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606100c683836040518060600160405280602781526020016104ba60279139610180565b9392505050565b6001600160a01b0381163b61013f5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080856001600160a01b03168560405161019d91906103ac565b600060405180830381855af49150503d80600081146101d8576040519150601f19603f3d011682016040523d82523d6000602084013e6101dd565b606091505b5090925090506101ef868383876101f9565b9695505050505050565b60608315610268578251600003610261576001600160a01b0385163b6102615760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610136565b5081610272565b610272838361027a565b949350505050565b81511561028a5781518083602001fd5b8060405162461bcd60e51b815260040161013691906103c8565b634e487b7160e01b600052604160045260246000fd5b60005b838110156102d55781810151838201526020016102bd565b50506000910152565b600080604083850312156102f157600080fd5b82516001600160a01b038116811461030857600080fd5b60208401519092506001600160401b038082111561032557600080fd5b818501915085601f83011261033957600080fd5b81518181111561034b5761034b6102a4565b604051601f8201601f19908116603f01168101908382118183101715610373576103736102a4565b8160405282815288602084870101111561038c57600080fd5b61039d8360208301602088016102ba565b80955050505050509250929050565b600082516103be8184602087016102ba565b9190910192915050565b60208152600082518060208401526103e78160408501602087016102ba565b601f01601f19169190910160400192915050565b60b1806104096000396000f3fe608060405236601057600e6013565b005b600e5b601f601b6021565b6058565b565b600060537f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b3660008037600080366000845af43d6000803e8080156076573d6000f35b3d6000fdfea264697066735822122004eafc12009ddd1bac09a0fbcdbbc0e38fdd09caf7f585983e4d324fd8f99a7364736f6c63430008140033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220e1aa4b48adb87b21fcb47bb450f472cb72d334e2cd0e21c7b21404ab71e71c8364736f6c6343000814003360c0604052306080523480156200001557600080fd5b5060405162001eec38038062001eec833981016040819052620000389162000117565b6001600160a01b03811660a0526200004f62000056565b5062000149565b600054610100900460ff1615620000c35760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161462000115576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6000602082840312156200012a57600080fd5b81516001600160a01b03811681146200014257600080fd5b9392505050565b60805160a051611d5362000199600039600081816103f701528181610c52015261110501526000818161054b0152818161059401528181610672015281816106b201526107410152611d536000f3fe6080604052600436106101405760003560e01c806352d1902d116100b657806395d89b411161006f57806395d89b4114610393578063a457c2d7146103a8578063a9059cbb146103c8578063b0d691fe146103e8578063c4d66de81461041b578063dd62ed3e1461043b57600080fd5b806352d1902d146102b057806370a08231146102c557806379cc6790146102fb5780638a0b03601461031b5780638da5cb5b1461033b578063925247251461037357600080fd5b8063313ce56711610108578063313ce567146102015780633659cfe61461021d578063395093511461023d57806340c10f191461025d57806342966c681461027d5780634f1ef2861461029d57600080fd5b806306fdde0314610145578063095ea7b31461017057806318160ddd146101a0578063203abe34146101bf57806323b872dd146101e1575b600080fd5b34801561015157600080fd5b5061015a61045b565b604051610167919061164a565b60405180910390f35b34801561017c57600080fd5b5061019061018b366004611699565b6104ed565b6040519015158152602001610167565b3480156101ac57600080fd5b506035545b604051908152602001610167565b3480156101cb57600080fd5b506101df6101da3660046117a0565b610507565b005b3480156101ed57600080fd5b506101906101fc3660046117d5565b610524565b34801561020d57600080fd5b5060405160128152602001610167565b34801561022957600080fd5b506101df610238366004611811565b610541565b34801561024957600080fd5b50610190610258366004611699565b610626565b34801561026957600080fd5b506101df610278366004611699565b610648565b34801561028957600080fd5b506101df61029836600461182c565b61065e565b6101df6102ab366004611845565b610668565b3480156102bc57600080fd5b506101b1610734565b3480156102d157600080fd5b506101b16102e0366004611811565b6001600160a01b031660009081526033602052604090205490565b34801561030757600080fd5b506101df610316366004611699565b6107e7565b34801561032757600080fd5b506101df6103363660046118eb565b6107fc565b34801561034757600080fd5b5060fb5461035b906001600160a01b031681565b6040516001600160a01b039091168152602001610167565b34801561037f57600080fd5b506101df61038e366004611930565b610850565b34801561039f57600080fd5b5061015a610936565b3480156103b457600080fd5b506101906103c3366004611699565b610945565b3480156103d457600080fd5b506101906103e3366004611699565b6109cb565b3480156103f457600080fd5b507f000000000000000000000000000000000000000000000000000000000000000061035b565b34801561042757600080fd5b506101df610436366004611811565b6109e6565b34801561044757600080fd5b506101b16104563660046119eb565b610af8565b60606036805461046a90611a1e565b80601f016020809104026020016040519081016040528092919081815260200182805461049690611a1e565b80156104e35780601f106104b8576101008083540402835291602001916104e3565b820191906000526020600020905b8154815290600101906020018083116104c657829003601f168201915b5050505050905090565b6000336104fb818585610b23565b60019150505b92915050565b61050f610c47565b61052181670de0b6b3a76400006107fc565b50565b600061052e610c47565b610539848484610cd6565b949350505050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036105925760405162461bcd60e51b815260040161058990611a58565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166105db600080516020611cd7833981519152546001600160a01b031690565b6001600160a01b0316146106015760405162461bcd60e51b815260040161058990611aa4565b61060a81610cef565b6040805160008082526020820190925261052191839190610cf7565b6000336104fb8185856106398383610af8565b6106439190611b06565b610b23565b610650610c47565b61065a8282610e62565b5050565b6105213382610f23565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036106b05760405162461bcd60e51b815260040161058990611a58565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166106f9600080516020611cd7833981519152546001600160a01b031690565b6001600160a01b03161461071f5760405162461bcd60e51b815260040161058990611aa4565b61072882610cef565b61065a82826001610cf7565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107d45760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610589565b50600080516020611cd783398151915290565b6107f2823383611057565b61065a8282610f23565b610804610c47565b60005b825181101561084b57600083828151811061082457610824611b19565b602002602001015190506108388184610e62565b508061084381611b2f565b915050610807565b505050565b610858610c47565b80518251146108cf5760405162461bcd60e51b815260206004820152603760248201527f726563697069656e747320616e6420616d6f756e747320617272617973206d7560448201527f73742068617665207468652073616d65206c656e6774680000000000000000006064820152608401610589565b60005b825181101561084b5760008382815181106108ef576108ef611b19565b60200260200101519050600083838151811061090d5761090d611b19565b602002602001015190506109218282610e62565b5050808061092e90611b2f565b9150506108d2565b60606037805461046a90611a1e565b600033816109538286610af8565b9050838110156109b35760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608401610589565b6109c08286868403610b23565b506001949350505050565b60006109d5610c47565b6109df83836110d1565b9392505050565b600054610100900460ff1615808015610a065750600054600160ff909116105b80610a205750303b158015610a20575060005460ff166001145b610a835760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610589565b6000805460ff191660011790558015610aa6576000805461ff0019166101001790555b610aaf826110df565b801561065a576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6001600160a01b03918216600090815260346020908152604080832093909416825291909152205490565b6001600160a01b038316610b855760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610589565b6001600160a01b038216610be65760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610589565b6001600160a01b0383811660008181526034602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610c88575060fb546001600160a01b031633145b610cd45760405162461bcd60e51b815260206004820181905260248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152606401610589565b565b600033610ce4858285611057565b6109c08585856111a7565b610521610c47565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610d2a5761084b83611352565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610d84575060408051601f3d908101601f19168201909252610d8191810190611b48565b60015b610de75760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610589565b600080516020611cd78339815191528114610e565760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610589565b5061084b8383836113ee565b6001600160a01b038216610eb85760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610589565b8060356000828254610eca9190611b06565b90915550506001600160a01b0382166000818152603360209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6001600160a01b038216610f835760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610589565b6001600160a01b03821660009081526033602052604090205481811015610ff75760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610589565b6001600160a01b03831660008181526033602090815260408083208686039055603580548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b60006110638484610af8565b905060001981146110cb57818110156110be5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610589565b6110cb8484848403610b23565b50505050565b6000336104fb8185856111a7565b60fb80546001600160a01b0319166001600160a01b0383811691821790925560405190917f000000000000000000000000000000000000000000000000000000000000000016907fe90b83f24696d649f6cdd7a0b35764256ae51f4998f6b266afb6fc6f2d50c60290600090a36111976040518060400160405280600e81526020016d23b930ba34ba3ab232aa37b5b2b760911b8152506040518060400160405280600381526020016211d51560ea1b815250611413565b61119f611444565b610521611444565b6001600160a01b03831661120b5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610589565b6001600160a01b03821661126d5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610589565b6001600160a01b038316600090815260336020526040902054818110156112e55760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610589565b6001600160a01b0380851660008181526033602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906113459086815260200190565b60405180910390a36110cb565b6001600160a01b0381163b6113bf5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610589565b600080516020611cd783398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6113f78361146b565b6000825111806114045750805b1561084b576110cb83836114ab565b600054610100900460ff1661143a5760405162461bcd60e51b815260040161058990611b61565b61065a82826114d0565b600054610100900460ff16610cd45760405162461bcd60e51b815260040161058990611b61565b61147481611352565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606109df8383604051806060016040528060278152602001611cf760279139611510565b600054610100900460ff166114f75760405162461bcd60e51b815260040161058990611b61565b60366115038382611bfa565b50603761084b8282611bfa565b6060600080856001600160a01b03168560405161152d9190611cba565b600060405180830381855af49150503d8060008114611568576040519150601f19603f3d011682016040523d82523d6000602084013e61156d565b606091505b509150915061157e86838387611588565b9695505050505050565b606083156115f75782516000036115f0576001600160a01b0385163b6115f05760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610589565b5081610539565b610539838381511561160c5781518083602001fd5b8060405162461bcd60e51b8152600401610589919061164a565b60005b83811015611641578181015183820152602001611629565b50506000910152565b6020815260008251806020840152611669816040850160208701611626565b601f01601f19169190910160400192915050565b80356001600160a01b038116811461169457600080fd5b919050565b600080604083850312156116ac57600080fd5b6116b58361167d565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611702576117026116c3565b604052919050565b600067ffffffffffffffff821115611724576117246116c3565b5060051b60200190565b600082601f83011261173f57600080fd5b8135602061175461174f8361170a565b6116d9565b82815260059290921b8401810191818101908684111561177357600080fd5b8286015b84811015611795576117888161167d565b8352918301918301611777565b509695505050505050565b6000602082840312156117b257600080fd5b813567ffffffffffffffff8111156117c957600080fd5b6105398482850161172e565b6000806000606084860312156117ea57600080fd5b6117f38461167d565b92506118016020850161167d565b9150604084013590509250925092565b60006020828403121561182357600080fd5b6109df8261167d565b60006020828403121561183e57600080fd5b5035919050565b6000806040838503121561185857600080fd5b6118618361167d565b915060208084013567ffffffffffffffff8082111561187f57600080fd5b818601915086601f83011261189357600080fd5b8135818111156118a5576118a56116c3565b6118b7601f8201601f191685016116d9565b915080825287848285010111156118cd57600080fd5b80848401858401376000848284010152508093505050509250929050565b600080604083850312156118fe57600080fd5b823567ffffffffffffffff81111561191557600080fd5b6119218582860161172e565b95602094909401359450505050565b6000806040838503121561194357600080fd5b823567ffffffffffffffff8082111561195b57600080fd5b6119678683870161172e565b935060209150818501358181111561197e57600080fd5b85019050601f8101861361199157600080fd5b803561199f61174f8261170a565b81815260059190911b820183019083810190888311156119be57600080fd5b928401925b828410156119dc578335825292840192908401906119c3565b80955050505050509250929050565b600080604083850312156119fe57600080fd5b611a078361167d565b9150611a156020840161167d565b90509250929050565b600181811c90821680611a3257607f821691505b602082108103611a5257634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b8082018082111561050157610501611af0565b634e487b7160e01b600052603260045260246000fd5b600060018201611b4157611b41611af0565b5060010190565b600060208284031215611b5a57600080fd5b5051919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b601f82111561084b57600081815260208120601f850160051c81016020861015611bd35750805b601f850160051c820191505b81811015611bf257828155600101611bdf565b505050505050565b815167ffffffffffffffff811115611c1457611c146116c3565b611c2881611c228454611a1e565b84611bac565b602080601f831160018114611c5d5760008415611c455750858301515b600019600386901b1c1916600185901b178555611bf2565b600085815260208120601f198616915b82811015611c8c57888601518255948401946001909101908401611c6d565b5085821015611caa5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008251611ccc818460208701611626565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220c831ef426c96c6d9aa5816c5089cc3674064281c88b2a456acee5d178285e42c64736f6c63430008140033",
}

// GrfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use GrfactoryMetaData.ABI instead.
var GrfactoryABI = GrfactoryMetaData.ABI

// GrfactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GrfactoryMetaData.Bin instead.
var GrfactoryBin = GrfactoryMetaData.Bin

// DeployGrfactory deploys a new Ethereum contract, binding an instance of Grfactory to it.
func DeployGrfactory(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *Grfactory, error) {
	parsed, err := GrfactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GrfactoryBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Grfactory{GrfactoryCaller: GrfactoryCaller{contract: contract}, GrfactoryTransactor: GrfactoryTransactor{contract: contract}, GrfactoryFilterer: GrfactoryFilterer{contract: contract}}, nil
}

// Grfactory is an auto generated Go binding around an Ethereum contract.
type Grfactory struct {
	GrfactoryCaller     // Read-only binding to the contract
	GrfactoryTransactor // Write-only binding to the contract
	GrfactoryFilterer   // Log filterer for contract events
}

// GrfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type GrfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GrfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GrfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GrfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GrfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GrfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GrfactorySession struct {
	Contract     *Grfactory        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GrfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GrfactoryCallerSession struct {
	Contract *GrfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GrfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GrfactoryTransactorSession struct {
	Contract     *GrfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GrfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type GrfactoryRaw struct {
	Contract *Grfactory // Generic contract binding to access the raw methods on
}

// GrfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GrfactoryCallerRaw struct {
	Contract *GrfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// GrfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GrfactoryTransactorRaw struct {
	Contract *GrfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGrfactory creates a new instance of Grfactory, bound to a specific deployed contract.
func NewGrfactory(address common.Address, backend bind.ContractBackend) (*Grfactory, error) {
	contract, err := bindGrfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Grfactory{GrfactoryCaller: GrfactoryCaller{contract: contract}, GrfactoryTransactor: GrfactoryTransactor{contract: contract}, GrfactoryFilterer: GrfactoryFilterer{contract: contract}}, nil
}

// NewGrfactoryCaller creates a new read-only instance of Grfactory, bound to a specific deployed contract.
func NewGrfactoryCaller(address common.Address, caller bind.ContractCaller) (*GrfactoryCaller, error) {
	contract, err := bindGrfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GrfactoryCaller{contract: contract}, nil
}

// NewGrfactoryTransactor creates a new write-only instance of Grfactory, bound to a specific deployed contract.
func NewGrfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*GrfactoryTransactor, error) {
	contract, err := bindGrfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GrfactoryTransactor{contract: contract}, nil
}

// NewGrfactoryFilterer creates a new log filterer instance of Grfactory, bound to a specific deployed contract.
func NewGrfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*GrfactoryFilterer, error) {
	contract, err := bindGrfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GrfactoryFilterer{contract: contract}, nil
}

// bindGrfactory binds a generic wrapper to an already deployed contract.
func bindGrfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GrfactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Grfactory *GrfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Grfactory.Contract.GrfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Grfactory *GrfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Grfactory.Contract.GrfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Grfactory *GrfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Grfactory.Contract.GrfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Grfactory *GrfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Grfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Grfactory *GrfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Grfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Grfactory *GrfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Grfactory.Contract.contract.Transact(opts, method, params...)
}

// GetGratitudeTokenAddress is a free data retrieval call binding the contract method 0xd87c2ee5.
//
// Solidity: function getGratitudeTokenAddress(address owner, uint256 salt) view returns(address)
func (_Grfactory *GrfactoryCaller) GetGratitudeTokenAddress(opts *bind.CallOpts, owner common.Address, salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Grfactory.contract.Call(opts, &out, "getGratitudeTokenAddress", owner, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGratitudeTokenAddress is a free data retrieval call binding the contract method 0xd87c2ee5.
//
// Solidity: function getGratitudeTokenAddress(address owner, uint256 salt) view returns(address)
func (_Grfactory *GrfactorySession) GetGratitudeTokenAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _Grfactory.Contract.GetGratitudeTokenAddress(&_Grfactory.CallOpts, owner, salt)
}

// GetGratitudeTokenAddress is a free data retrieval call binding the contract method 0xd87c2ee5.
//
// Solidity: function getGratitudeTokenAddress(address owner, uint256 salt) view returns(address)
func (_Grfactory *GrfactoryCallerSession) GetGratitudeTokenAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _Grfactory.Contract.GetGratitudeTokenAddress(&_Grfactory.CallOpts, owner, salt)
}

// GratitudeImplementation is a free data retrieval call binding the contract method 0xc56ce588.
//
// Solidity: function gratitudeImplementation() view returns(address)
func (_Grfactory *GrfactoryCaller) GratitudeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Grfactory.contract.Call(opts, &out, "gratitudeImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GratitudeImplementation is a free data retrieval call binding the contract method 0xc56ce588.
//
// Solidity: function gratitudeImplementation() view returns(address)
func (_Grfactory *GrfactorySession) GratitudeImplementation() (common.Address, error) {
	return _Grfactory.Contract.GratitudeImplementation(&_Grfactory.CallOpts)
}

// GratitudeImplementation is a free data retrieval call binding the contract method 0xc56ce588.
//
// Solidity: function gratitudeImplementation() view returns(address)
func (_Grfactory *GrfactoryCallerSession) GratitudeImplementation() (common.Address, error) {
	return _Grfactory.Contract.GratitudeImplementation(&_Grfactory.CallOpts)
}

// CreateGratitudeToken is a paid mutator transaction binding the contract method 0x01dfecf0.
//
// Solidity: function createGratitudeToken(address owner, uint256 salt) returns(address ret)
func (_Grfactory *GrfactoryTransactor) CreateGratitudeToken(opts *bind.TransactOpts, owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _Grfactory.contract.Transact(opts, "createGratitudeToken", owner, salt)
}

// CreateGratitudeToken is a paid mutator transaction binding the contract method 0x01dfecf0.
//
// Solidity: function createGratitudeToken(address owner, uint256 salt) returns(address ret)
func (_Grfactory *GrfactorySession) CreateGratitudeToken(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _Grfactory.Contract.CreateGratitudeToken(&_Grfactory.TransactOpts, owner, salt)
}

// CreateGratitudeToken is a paid mutator transaction binding the contract method 0x01dfecf0.
//
// Solidity: function createGratitudeToken(address owner, uint256 salt) returns(address ret)
func (_Grfactory *GrfactoryTransactorSession) CreateGratitudeToken(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _Grfactory.Contract.CreateGratitudeToken(&_Grfactory.TransactOpts, owner, salt)
}

// GrfactoryGratitudeTokenCreatedIterator is returned from FilterGratitudeTokenCreated and is used to iterate over the raw logs and unpacked data for GratitudeTokenCreated events raised by the Grfactory contract.
type GrfactoryGratitudeTokenCreatedIterator struct {
	Event *GrfactoryGratitudeTokenCreated // Event containing the contract specifics and raw log

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
func (it *GrfactoryGratitudeTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GrfactoryGratitudeTokenCreated)
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
		it.Event = new(GrfactoryGratitudeTokenCreated)
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
func (it *GrfactoryGratitudeTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GrfactoryGratitudeTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GrfactoryGratitudeTokenCreated represents a GratitudeTokenCreated event raised by the Grfactory contract.
type GrfactoryGratitudeTokenCreated struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGratitudeTokenCreated is a free log retrieval operation binding the contract event 0x632ffdb1f5dbeb99afea447aa1f5f6cc4b23d0f61f074784c8c03a4134d7c15a.
//
// Solidity: event GratitudeTokenCreated(address indexed owner)
func (_Grfactory *GrfactoryFilterer) FilterGratitudeTokenCreated(opts *bind.FilterOpts, owner []common.Address) (*GrfactoryGratitudeTokenCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Grfactory.contract.FilterLogs(opts, "GratitudeTokenCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &GrfactoryGratitudeTokenCreatedIterator{contract: _Grfactory.contract, event: "GratitudeTokenCreated", logs: logs, sub: sub}, nil
}

// WatchGratitudeTokenCreated is a free log subscription operation binding the contract event 0x632ffdb1f5dbeb99afea447aa1f5f6cc4b23d0f61f074784c8c03a4134d7c15a.
//
// Solidity: event GratitudeTokenCreated(address indexed owner)
func (_Grfactory *GrfactoryFilterer) WatchGratitudeTokenCreated(opts *bind.WatchOpts, sink chan<- *GrfactoryGratitudeTokenCreated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Grfactory.contract.WatchLogs(opts, "GratitudeTokenCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GrfactoryGratitudeTokenCreated)
				if err := _Grfactory.contract.UnpackLog(event, "GratitudeTokenCreated", log); err != nil {
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

// ParseGratitudeTokenCreated is a log parse operation binding the contract event 0x632ffdb1f5dbeb99afea447aa1f5f6cc4b23d0f61f074784c8c03a4134d7c15a.
//
// Solidity: event GratitudeTokenCreated(address indexed owner)
func (_Grfactory *GrfactoryFilterer) ParseGratitudeTokenCreated(log types.Log) (*GrfactoryGratitudeTokenCreated, error) {
	event := new(GrfactoryGratitudeTokenCreated)
	if err := _Grfactory.contract.UnpackLog(event, "GratitudeTokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
