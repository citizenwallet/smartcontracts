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
	Bin: "0x60c06040523060805234801562000014575f80fd5b5060405162001e5b38038062001e5b833981016040819052620000379162000114565b6001600160a01b03811660a0526200004e62000055565b5062000143565b5f54610100900460ff1615620000c15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116101562000112575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f6020828403121562000125575f80fd5b81516001600160a01b03811681146200013c575f80fd5b9392505050565b60805160a051611ccc6200018f5f395f81816103dd01528181610c1801526110c201525f818161052d015281816105760152818161065101528181610691015261071e0152611ccc5ff3fe60806040526004361061013c575f3560e01c806352d1902d116100b357806395d89b411161006d57806395d89b411461037d578063a457c2d714610391578063a9059cbb146103b0578063b0d691fe146103cf578063c4d66de814610401578063dd62ed3e14610420575f80fd5b806352d1902d146102a157806370a08231146102b557806379cc6790146102e95780638a0b0360146103085780638da5cb5b14610327578063925247251461035e575f80fd5b8063313ce56711610104578063313ce567146101f75780633659cfe614610212578063395093511461023157806340c10f191461025057806342966c681461026f5780634f1ef2861461028e575f80fd5b806306fdde0314610140578063095ea7b31461016a57806318160ddd14610199578063203abe34146101b757806323b872dd146101d8575b5f80fd5b34801561014b575f80fd5b5061015461043f565b60405161016191906115ef565b60405180910390f35b348015610175575f80fd5b5061018961018436600461163c565b6104cf565b6040519015158152602001610161565b3480156101a4575f80fd5b506035545b604051908152602001610161565b3480156101c2575f80fd5b506101d66101d136600461173b565b6104e8565b005b3480156101e3575f80fd5b506101896101f2366004611775565b610505565b348015610202575f80fd5b5060405160128152602001610161565b34801561021d575f80fd5b506101d661022c3660046117ae565b610523565b34801561023c575f80fd5b5061018961024b36600461163c565b610606565b34801561025b575f80fd5b506101d661026a36600461163c565b610627565b34801561027a575f80fd5b506101d66102893660046117c7565b61063d565b6101d661029c3660046117de565b610647565b3480156102ac575f80fd5b506101a9610712565b3480156102c0575f80fd5b506101a96102cf3660046117ae565b6001600160a01b03165f9081526033602052604090205490565b3480156102f4575f80fd5b506101d661030336600461163c565b6107c3565b348015610313575f80fd5b506101d661032236600461187e565b6107d8565b348015610332575f80fd5b5060fb54610346906001600160a01b031681565b6040516001600160a01b039091168152602001610161565b348015610369575f80fd5b506101d66103783660046118c0565b61082a565b348015610388575f80fd5b5061015461090d565b34801561039c575f80fd5b506101896103ab36600461163c565b61091c565b3480156103bb575f80fd5b506101896103ca36600461163c565b6109a1565b3480156103da575f80fd5b507f0000000000000000000000000000000000000000000000000000000000000000610346565b34801561040c575f80fd5b506101d661041b3660046117ae565b6109b4565b34801561042b575f80fd5b506101a961043a366004611975565b610ac0565b60606036805461044e906119a6565b80601f016020809104026020016040519081016040528092919081815260200182805461047a906119a6565b80156104c55780601f1061049c576101008083540402835291602001916104c5565b820191905f5260205f20905b8154815290600101906020018083116104a857829003601f168201915b5050505050905090565b5f336104dc818585610aea565b60019150505b92915050565b6104f0610c0d565b61050281670de0b6b3a76400006107d8565b50565b5f61050e610c0d565b610519848484610c9c565b90505b9392505050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036105745760405162461bcd60e51b815260040161056b906119de565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166105bc5f80516020611c50833981519152546001600160a01b031690565b6001600160a01b0316146105e25760405162461bcd60e51b815260040161056b90611a2a565b6105eb81610cb4565b604080515f8082526020820190925261050291839190610cbc565b5f336104dc8185856106188383610ac0565b6106229190611a8a565b610aea565b61062f610c0d565b6106398282610e26565b5050565b6105023382610ee5565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361068f5760405162461bcd60e51b815260040161056b906119de565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166106d75f80516020611c50833981519152546001600160a01b031690565b6001600160a01b0316146106fd5760405162461bcd60e51b815260040161056b90611a2a565b61070682610cb4565b61063982826001610cbc565b5f306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107b15760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000606482015260840161056b565b505f80516020611c5083398151915290565b6107ce823383611017565b6106398282610ee5565b6107e0610c0d565b5f5b8251811015610825575f8382815181106107fe576107fe611a9d565b602002602001015190506108128184610e26565b508061081d81611ab1565b9150506107e2565b505050565b610832610c0d565b80518251146108a95760405162461bcd60e51b815260206004820152603760248201527f726563697069656e747320616e6420616d6f756e747320617272617973206d7560448201527f73742068617665207468652073616d65206c656e677468000000000000000000606482015260840161056b565b5f5b8251811015610825575f8382815181106108c7576108c7611a9d565b602002602001015190505f8383815181106108e4576108e4611a9d565b602002602001015190506108f88282610e26565b5050808061090590611ab1565b9150506108ab565b60606037805461044e906119a6565b5f33816109298286610ac0565b9050838110156109895760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b606482015260840161056b565b6109968286868403610aea565b506001949350505050565b5f6109aa610c0d565b61051c838361108f565b5f54610100900460ff16158080156109d257505f54600160ff909116105b806109eb5750303b1580156109eb57505f5460ff166001145b610a4e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161056b565b5f805460ff191660011790558015610a6f575f805461ff0019166101001790555b610a788261109c565b8015610639575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6001600160a01b039182165f90815260346020908152604080832093909416825291909152205490565b6001600160a01b038316610b4c5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b606482015260840161056b565b6001600160a01b038216610bad5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b606482015260840161056b565b6001600160a01b038381165f8181526034602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610c4e575060fb546001600160a01b031633145b610c9a5760405162461bcd60e51b815260206004820181905260248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e74604482015260640161056b565b565b5f33610ca9858285611017565b610996858585611163565b610502610c0d565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610cef576108258361130c565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610d49575060408051601f3d908101601f19168201909252610d4691810190611ac9565b60015b610dac5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b606482015260840161056b565b5f80516020611c508339815191528114610e1a5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b606482015260840161056b565b506108258383836113a7565b6001600160a01b038216610e7c5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161056b565b8060355f828254610e8d9190611a8a565b90915550506001600160a01b0382165f818152603360209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6001600160a01b038216610f455760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b606482015260840161056b565b6001600160a01b0382165f9081526033602052604090205481811015610fb85760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b606482015260840161056b565b6001600160a01b0383165f8181526033602090815260408083208686039055603580548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b5f6110228484610ac0565b90505f198114611089578181101561107c5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161056b565b6110898484848403610aea565b50505050565b5f336104dc818585611163565b60fb80546001600160a01b0319166001600160a01b0383811691821790925560405190917f000000000000000000000000000000000000000000000000000000000000000016907fe90b83f24696d649f6cdd7a0b35764256ae51f4998f6b266afb6fc6f2d50c602905f90a36111536040518060400160405280600e81526020016d23b930ba34ba3ab232aa37b5b2b760911b8152506040518060400160405280600381526020016211d51560ea1b8152506113cb565b61115b6113fb565b6105026113fb565b6001600160a01b0383166111c75760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161056b565b6001600160a01b0382166112295760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b606482015260840161056b565b6001600160a01b0383165f90815260336020526040902054818110156112a05760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b606482015260840161056b565b6001600160a01b038085165f8181526033602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906112ff9086815260200190565b60405180910390a3611089565b6001600160a01b0381163b6113795760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b606482015260840161056b565b5f80516020611c5083398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6113b083611421565b5f825111806113bc5750805b15610825576110898383611460565b5f54610100900460ff166113f15760405162461bcd60e51b815260040161056b90611ae0565b6106398282611550565b5f54610100900460ff16610c9a5760405162461bcd60e51b815260040161056b90611ae0565b61142a8161130c565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606001600160a01b0383163b6114c85760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b606482015260840161056b565b5f80846001600160a01b0316846040516114e29190611b2b565b5f60405180830381855af49150503d805f811461151a576040519150601f19603f3d011682016040523d82523d5f602084013e61151f565b606091505b50915091506115478282604051806060016040528060278152602001611c706027913961158f565b95945050505050565b5f54610100900460ff166115765760405162461bcd60e51b815260040161056b90611ae0565b60366115828382611b93565b5060376108258282611b93565b6060831561159e57508161051c565b61051c83838151156115b35781518083602001fd5b8060405162461bcd60e51b815260040161056b91906115ef565b5f5b838110156115e75781810151838201526020016115cf565b50505f910152565b602081525f825180602084015261160d8160408501602087016115cd565b601f01601f19169190910160400192915050565b80356001600160a01b0381168114611637575f80fd5b919050565b5f806040838503121561164d575f80fd5b61165683611621565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff811182821017156116a1576116a1611664565b604052919050565b5f67ffffffffffffffff8211156116c2576116c2611664565b5060051b60200190565b5f82601f8301126116db575f80fd5b813560206116f06116eb836116a9565b611678565b82815260059290921b8401810191818101908684111561170e575f80fd5b8286015b848110156117305761172381611621565b8352918301918301611712565b509695505050505050565b5f6020828403121561174b575f80fd5b813567ffffffffffffffff811115611761575f80fd5b61176d848285016116cc565b949350505050565b5f805f60608486031215611787575f80fd5b61179084611621565b925061179e60208501611621565b9150604084013590509250925092565b5f602082840312156117be575f80fd5b61051c82611621565b5f602082840312156117d7575f80fd5b5035919050565b5f80604083850312156117ef575f80fd5b6117f883611621565b915060208084013567ffffffffffffffff80821115611815575f80fd5b818601915086601f830112611828575f80fd5b81358181111561183a5761183a611664565b61184c601f8201601f19168501611678565b91508082528784828501011115611861575f80fd5b80848401858401375f848284010152508093505050509250929050565b5f806040838503121561188f575f80fd5b823567ffffffffffffffff8111156118a5575f80fd5b6118b1858286016116cc565b95602094909401359450505050565b5f80604083850312156118d1575f80fd5b823567ffffffffffffffff808211156118e8575f80fd5b6118f4868387016116cc565b935060209150818501358181111561190a575f80fd5b85019050601f8101861361191c575f80fd5b803561192a6116eb826116a9565b81815260059190911b82018301908381019088831115611948575f80fd5b928401925b828410156119665783358252928401929084019061194d565b80955050505050509250929050565b5f8060408385031215611986575f80fd5b61198f83611621565b915061199d60208401611621565b90509250929050565b600181811c908216806119ba57607f821691505b6020821081036119d857634e487b7160e01b5f52602260045260245ffd5b50919050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b808201808211156104e2576104e2611a76565b634e487b7160e01b5f52603260045260245ffd5b5f60018201611ac257611ac2611a76565b5060010190565b5f60208284031215611ad9575f80fd5b5051919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b5f8251611b3c8184602087016115cd565b9190910192915050565b601f821115610825575f81815260208120601f850160051c81016020861015611b6c5750805b601f850160051c820191505b81811015611b8b57828155600101611b78565b505050505050565b815167ffffffffffffffff811115611bad57611bad611664565b611bc181611bbb84546119a6565b84611b46565b602080601f831160018114611bf4575f8415611bdd5750858301515b5f19600386901b1c1916600185901b178555611b8b565b5f85815260208120601f198616915b82811015611c2257888601518255948401946001909101908401611c03565b5085821015611c3f57878501515f19600388901b60f8161c191681555b5050505050600190811b0190555056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220c0d2ea3be259e09a065a9ec8410726c4220dca87ae915559a5709bf2233d1cca64736f6c63430008140033",
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
