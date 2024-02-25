// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simpleFaucet

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

// SimpleFaucetMetaData contains all meta data concerning the SimpleFaucet contract.
var SimpleFaucetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REDEEM_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"_redeemInterval\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"_redeemAdmin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemInterval\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"redeemed\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"time\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b50608051611ab161004c6000396000818161052c0152818161056c015281816107260152818161076601526107f50152611ab16000f3fe6080604052600436106101355760003560e01c80638da5cb5b116100ab578063a5ff3f801161006f578063a5ff3f8014610391578063aa8c217c146103b1578063be040fb0146103c8578063d547741f146103dd578063f2fde38b146103fd578063fc0c546a1461041d57600080fd5b80638da5cb5b146102b557806391d14854146102e75780639f4568ef14610307578063a217fddf14610354578063a2d2c71c1461036957600080fd5b80633ccfd60b116100fd5780633ccfd60b1461020f5780634f1ef286146102245780635018c2091461023757806352d1902d1461026b578063715018a61461028057806377202ba11461029557600080fd5b806301ffc9a71461013a578063248a9ca31461016f5780632f2ff15d146101ad57806336568abe146101cf5780633659cfe6146101ef575b600080fd5b34801561014657600080fd5b5061015a6101553660046115a9565b61043e565b60405190151581526020015b60405180910390f35b34801561017b57600080fd5b5061019f61018a3660046115d3565b60009081526097602052604090206001015490565b604051908152602001610166565b3480156101b957600080fd5b506101cd6101c8366004611601565b610475565b005b3480156101db57600080fd5b506101cd6101ea366004611601565b61049f565b3480156101fb57600080fd5b506101cd61020a366004611631565b610522565b34801561021b57600080fd5b506101cd610601565b6101cd610232366004611664565b61071c565b34801561024357600080fd5b5061019f7f281081d9b36b37208f0d8dfce5adc7e00d31ece09269aaa8d0bfa43e6840a33881565b34801561027757600080fd5b5061019f6107e8565b34801561028c57600080fd5b506101cd61089b565b3480156102a157600080fd5b506101cd6102b0366004611728565b6108af565b3480156102c157600080fd5b506033546001600160a01b03165b6040516001600160a01b039091168152602001610166565b3480156102f357600080fd5b5061015a610302366004611601565b610a53565b34801561031357600080fd5b5061033d610322366004611631565b6101306020526000908152604090205465ffffffffffff1681565b60405165ffffffffffff9091168152602001610166565b34801561036057600080fd5b5061019f600081565b34801561037557600080fd5b5061012f546102cf90600160301b90046001600160a01b031681565b34801561039d57600080fd5b5061012f5461033d9065ffffffffffff1681565b3480156103bd57600080fd5b5061019f61012e5481565b3480156103d457600080fd5b506101cd610a7e565b3480156103e957600080fd5b506101cd6103f8366004611601565b610d10565b34801561040957600080fd5b506101cd610418366004611631565b610d35565b34801561042957600080fd5b5061012d546102cf906001600160a01b031681565b60006001600160e01b03198216637965db0b60e01b148061046f57506301ffc9a760e01b6001600160e01b03198316145b92915050565b60008281526097602052604090206001015461049081610dab565b61049a8383610db5565b505050565b6001600160a01b03811633146105145760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084015b60405180910390fd5b61051e8282610e3b565b5050565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361056a5760405162461bcd60e51b815260040161050b9061179b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166105b3600080516020611a35833981519152546001600160a01b031690565b6001600160a01b0316146105d95760405162461bcd60e51b815260040161050b906117e7565b6105e281610ea2565b604080516000808252602082019092526105fe91839190610eaa565b50565b7f281081d9b36b37208f0d8dfce5adc7e00d31ece09269aaa8d0bfa43e6840a33861062b81610dab565b61012d5461012f546040516370a0823160e01b81523060048201526001600160a01b039283169263a9059cbb92600160301b9004169083906370a0823190602401602060405180830381865afa158015610689573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ad9190611833565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af11580156106f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061051e919061184c565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036107645760405162461bcd60e51b815260040161050b9061179b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166107ad600080516020611a35833981519152546001600160a01b031690565b6001600160a01b0316146107d35760405162461bcd60e51b815260040161050b906117e7565b6107dc82610ea2565b61051e82826001610eaa565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108885760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000606482015260840161050b565b50600080516020611a3583398151915290565b6108a3611015565b6108ad600061106f565b565b600054610100900460ff16158080156108cf5750600054600160ff909116105b806108e95750303b1580156108e9575060005460ff166001145b61094c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161050b565b6000805460ff19166001179055801561096f576000805461ff0019166101001790555b6109776110c1565b61097f6110f0565b61012d80546001600160a01b038088166001600160a01b03199092169190911790915561012e85905561012f8054918416600160301b026001600160d01b031990921665ffffffffffff8616179190911790556109fc7f281081d9b36b37208f0d8dfce5adc7e00d31ece09269aaa8d0bfa43e6840a33883611117565b610a0586610d35565b8015610a4b576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b60009182526097602090815260408084206001600160a01b0393909316845291905290205460ff1690565b61012f54429065ffffffffffff1615610b365761012f5433600090815261013060205260408120549091610abd9165ffffffffffff9182169116611884565b90508065ffffffffffff168265ffffffffffff161015610b305760405162461bcd60e51b815260206004820152602860248201527f53696d706c654661756365743a2072656465656d20696e74657276616c206e6f6044820152671d081c185cdcd95960c21b606482015260840161050b565b50610b9c565b336000908152610130602052604090205465ffffffffffff1615610b9c5760405162461bcd60e51b815260206004820152601e60248201527f53696d706c654661756365743a20616c72656164792072656465656d65640000604482015260640161050b565b61012e5461012d546040516370a0823160e01b81523060048201526001600160a01b03909116906370a0823190602401602060405180830381865afa158015610be9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0d9190611833565b1015610c665760405162461bcd60e51b815260206004820152602260248201527f53696d706c654661756365743a20696e73756666696369656e742062616c616e604482015261636560f01b606482015260840161050b565b61012d5461012e5460405163a9059cbb60e01b815233600482015260248101919091526001600160a01b039091169063a9059cbb906044016020604051808303816000875af1158015610cbd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce1919061184c565b5033600090815261013060205260409020805465ffffffffffff191665ffffffffffff92909216919091179055565b600082815260976020526040902060010154610d2b81610dab565b61049a8383610e3b565b610d3d611015565b6001600160a01b038116610da25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161050b565b6105fe8161106f565b6105fe8133611121565b610dbf8282610a53565b61051e5760008281526097602090815260408083206001600160a01b03851684529091529020805460ff19166001179055610df73390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b610e458282610a53565b1561051e5760008281526097602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6105fe611015565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610edd5761049a8361117a565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610f37575060408051601f3d908101601f19168201909252610f3491810190611833565b60015b610f9a5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b606482015260840161050b565b600080516020611a3583398151915281146110095760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b606482015260840161050b565b5061049a838383611216565b6033546001600160a01b031633146108ad5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161050b565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff166110e85760405162461bcd60e51b815260040161050b906118aa565b6108ad611241565b600054610100900460ff166108ad5760405162461bcd60e51b815260040161050b906118aa565b61051e8282610db5565b61112b8282610a53565b61051e5761113881611271565b611143836020611283565b604051602001611154929190611919565b60408051601f198184030181529082905262461bcd60e51b825261050b9160040161198e565b6001600160a01b0381163b6111e75760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b606482015260840161050b565b600080516020611a3583398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b61121f83611426565b60008251118061122c5750805b1561049a5761123b8383611466565b50505050565b600054610100900460ff166112685760405162461bcd60e51b815260040161050b906118aa565b6108ad3361106f565b606061046f6001600160a01b03831660145b606060006112928360026119c1565b61129d9060026119d8565b67ffffffffffffffff8111156112b5576112b561164e565b6040519080825280601f01601f1916602001820160405280156112df576020820181803683370190505b509050600360fc1b816000815181106112fa576112fa6119eb565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110611329576113296119eb565b60200101906001600160f81b031916908160001a905350600061134d8460026119c1565b6113589060016119d8565b90505b60018111156113d0576f181899199a1a9b1b9c1cb0b131b232b360811b85600f166010811061138c5761138c6119eb565b1a60f81b8282815181106113a2576113a26119eb565b60200101906001600160f81b031916908160001a90535060049490941c936113c981611a01565b905061135b565b50831561141f5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161050b565b9392505050565b61142f8161117a565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b606061141f8383604051806060016040528060278152602001611a55602791396060600080856001600160a01b0316856040516114a39190611a18565b600060405180830381855af49150503d80600081146114de576040519150601f19603f3d011682016040523d82523d6000602084013e6114e3565b606091505b50915091506114f4868383876114fe565b9695505050505050565b6060831561156d578251600003611566576001600160a01b0385163b6115665760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161050b565b5081611577565b611577838361157f565b949350505050565b81511561158f5781518083602001fd5b8060405162461bcd60e51b815260040161050b919061198e565b6000602082840312156115bb57600080fd5b81356001600160e01b03198116811461141f57600080fd5b6000602082840312156115e557600080fd5b5035919050565b6001600160a01b03811681146105fe57600080fd5b6000806040838503121561161457600080fd5b823591506020830135611626816115ec565b809150509250929050565b60006020828403121561164357600080fd5b813561141f816115ec565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561167757600080fd5b8235611682816115ec565b9150602083013567ffffffffffffffff8082111561169f57600080fd5b818501915085601f8301126116b357600080fd5b8135818111156116c5576116c561164e565b604051601f8201601f19908116603f011681019083821181831017156116ed576116ed61164e565b8160405282815288602084870101111561170657600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b600080600080600060a0868803121561174057600080fd5b853561174b816115ec565b9450602086013561175b816115ec565b935060408601359250606086013565ffffffffffff8116811461177d57600080fd5b9150608086013561178d816115ec565b809150509295509295909350565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b60006020828403121561184557600080fd5b5051919050565b60006020828403121561185e57600080fd5b8151801515811461141f57600080fd5b634e487b7160e01b600052601160045260246000fd5b65ffffffffffff8181168382160190808211156118a3576118a361186e565b5092915050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60005b838110156119105781810151838201526020016118f8565b50506000910152565b7f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008152600083516119518160178501602088016118f5565b7001034b99036b4b9b9b4b733903937b6329607d1b60179184019182015283516119828160288401602088016118f5565b01602801949350505050565b60208152600082518060208401526119ad8160408501602087016118f5565b601f01601f19169190910160400192915050565b808202811582820484141761046f5761046f61186e565b8082018082111561046f5761046f61186e565b634e487b7160e01b600052603260045260246000fd5b600081611a1057611a1061186e565b506000190190565b60008251611a2a8184602087016118f5565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212200cd8901215320b410f59b081f215378ffd1a3f4fb9af6748d82b5f490aea23e564736f6c63430008140033",
}

// SimpleFaucetABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleFaucetMetaData.ABI instead.
var SimpleFaucetABI = SimpleFaucetMetaData.ABI

// SimpleFaucetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleFaucetMetaData.Bin instead.
var SimpleFaucetBin = SimpleFaucetMetaData.Bin

// DeploySimpleFaucet deploys a new Ethereum contract, binding an instance of SimpleFaucet to it.
func DeploySimpleFaucet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleFaucet, error) {
	parsed, err := SimpleFaucetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleFaucetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleFaucet{SimpleFaucetCaller: SimpleFaucetCaller{contract: contract}, SimpleFaucetTransactor: SimpleFaucetTransactor{contract: contract}, SimpleFaucetFilterer: SimpleFaucetFilterer{contract: contract}}, nil
}

// SimpleFaucet is an auto generated Go binding around an Ethereum contract.
type SimpleFaucet struct {
	SimpleFaucetCaller     // Read-only binding to the contract
	SimpleFaucetTransactor // Write-only binding to the contract
	SimpleFaucetFilterer   // Log filterer for contract events
}

// SimpleFaucetCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleFaucetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleFaucetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleFaucetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleFaucetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleFaucetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleFaucetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleFaucetSession struct {
	Contract     *SimpleFaucet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleFaucetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleFaucetCallerSession struct {
	Contract *SimpleFaucetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SimpleFaucetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleFaucetTransactorSession struct {
	Contract     *SimpleFaucetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SimpleFaucetRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleFaucetRaw struct {
	Contract *SimpleFaucet // Generic contract binding to access the raw methods on
}

// SimpleFaucetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleFaucetCallerRaw struct {
	Contract *SimpleFaucetCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleFaucetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleFaucetTransactorRaw struct {
	Contract *SimpleFaucetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleFaucet creates a new instance of SimpleFaucet, bound to a specific deployed contract.
func NewSimpleFaucet(address common.Address, backend bind.ContractBackend) (*SimpleFaucet, error) {
	contract, err := bindSimpleFaucet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleFaucet{SimpleFaucetCaller: SimpleFaucetCaller{contract: contract}, SimpleFaucetTransactor: SimpleFaucetTransactor{contract: contract}, SimpleFaucetFilterer: SimpleFaucetFilterer{contract: contract}}, nil
}

// NewSimpleFaucetCaller creates a new read-only instance of SimpleFaucet, bound to a specific deployed contract.
func NewSimpleFaucetCaller(address common.Address, caller bind.ContractCaller) (*SimpleFaucetCaller, error) {
	contract, err := bindSimpleFaucet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleFaucetCaller{contract: contract}, nil
}

// NewSimpleFaucetTransactor creates a new write-only instance of SimpleFaucet, bound to a specific deployed contract.
func NewSimpleFaucetTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleFaucetTransactor, error) {
	contract, err := bindSimpleFaucet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleFaucetTransactor{contract: contract}, nil
}

// NewSimpleFaucetFilterer creates a new log filterer instance of SimpleFaucet, bound to a specific deployed contract.
func NewSimpleFaucetFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleFaucetFilterer, error) {
	contract, err := bindSimpleFaucet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleFaucetFilterer{contract: contract}, nil
}

// bindSimpleFaucet binds a generic wrapper to an already deployed contract.
func bindSimpleFaucet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleFaucetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleFaucet *SimpleFaucetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleFaucet.Contract.SimpleFaucetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleFaucet *SimpleFaucetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.SimpleFaucetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleFaucet *SimpleFaucetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.SimpleFaucetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleFaucet *SimpleFaucetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleFaucet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleFaucet *SimpleFaucetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleFaucet *SimpleFaucetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SimpleFaucet *SimpleFaucetCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SimpleFaucet.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SimpleFaucet *SimpleFaucetSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SimpleFaucet.Contract.DEFAULTADMINROLE(&_SimpleFaucet.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SimpleFaucet *SimpleFaucetCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SimpleFaucet.Contract.DEFAULTADMINROLE(&_SimpleFaucet.CallOpts)
}

// REDEEMADMINROLE is a free data retrieval call binding the contract method 0x5018c209.
//
// Solidity: function REDEEM_ADMIN_ROLE() view returns(bytes32)
func (_SimpleFaucet *SimpleFaucetCaller) REDEEMADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SimpleFaucet.contract.Call(opts, &out, "REDEEM_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REDEEMADMINROLE is a free data retrieval call binding the contract method 0x5018c209.
//
// Solidity: function REDEEM_ADMIN_ROLE() view returns(bytes32)
func (_SimpleFaucet *SimpleFaucetSession) REDEEMADMINROLE() ([32]byte, error) {
	return _SimpleFaucet.Contract.REDEEMADMINROLE(&_SimpleFaucet.CallOpts)
}

// REDEEMADMINROLE is a free data retrieval call binding the contract method 0x5018c209.
//
// Solidity: function REDEEM_ADMIN_ROLE() view returns(bytes32)
func (_SimpleFaucet *SimpleFaucetCallerSession) REDEEMADMINROLE() ([32]byte, error) {
	return _SimpleFaucet.Contract.REDEEMADMINROLE(&_SimpleFaucet.CallOpts)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_SimpleFaucet *SimpleFaucetCaller) Amount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleFaucet.contract.Call(opts, &out, "amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_SimpleFaucet *SimpleFaucetSession) Amount() (*big.Int, error) {
	return _SimpleFaucet.Contract.Amount(&_SimpleFaucet.CallOpts)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_SimpleFaucet *SimpleFaucetCallerSession) Amount() (*big.Int, error) {
	return _SimpleFaucet.Contract.Amount(&_SimpleFaucet.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SimpleFaucet *SimpleFaucetCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SimpleFaucet.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SimpleFaucet *SimpleFaucetSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SimpleFaucet.Contract.GetRoleAdmin(&_SimpleFaucet.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SimpleFaucet *SimpleFaucetCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SimpleFaucet.Contract.GetRoleAdmin(&_SimpleFaucet.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SimpleFaucet *SimpleFaucetCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SimpleFaucet.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SimpleFaucet *SimpleFaucetSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SimpleFaucet.Contract.HasRole(&_SimpleFaucet.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SimpleFaucet *SimpleFaucetCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SimpleFaucet.Contract.HasRole(&_SimpleFaucet.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SimpleFaucet *SimpleFaucetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleFaucet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SimpleFaucet *SimpleFaucetSession) Owner() (common.Address, error) {
	return _SimpleFaucet.Contract.Owner(&_SimpleFaucet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SimpleFaucet *SimpleFaucetCallerSession) Owner() (common.Address, error) {
	return _SimpleFaucet.Contract.Owner(&_SimpleFaucet.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SimpleFaucet *SimpleFaucetCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SimpleFaucet.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SimpleFaucet *SimpleFaucetSession) ProxiableUUID() ([32]byte, error) {
	return _SimpleFaucet.Contract.ProxiableUUID(&_SimpleFaucet.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SimpleFaucet *SimpleFaucetCallerSession) ProxiableUUID() ([32]byte, error) {
	return _SimpleFaucet.Contract.ProxiableUUID(&_SimpleFaucet.CallOpts)
}

// RedeemAdmin is a free data retrieval call binding the contract method 0xa2d2c71c.
//
// Solidity: function redeemAdmin() view returns(address)
func (_SimpleFaucet *SimpleFaucetCaller) RedeemAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleFaucet.contract.Call(opts, &out, "redeemAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RedeemAdmin is a free data retrieval call binding the contract method 0xa2d2c71c.
//
// Solidity: function redeemAdmin() view returns(address)
func (_SimpleFaucet *SimpleFaucetSession) RedeemAdmin() (common.Address, error) {
	return _SimpleFaucet.Contract.RedeemAdmin(&_SimpleFaucet.CallOpts)
}

// RedeemAdmin is a free data retrieval call binding the contract method 0xa2d2c71c.
//
// Solidity: function redeemAdmin() view returns(address)
func (_SimpleFaucet *SimpleFaucetCallerSession) RedeemAdmin() (common.Address, error) {
	return _SimpleFaucet.Contract.RedeemAdmin(&_SimpleFaucet.CallOpts)
}

// RedeemInterval is a free data retrieval call binding the contract method 0xa5ff3f80.
//
// Solidity: function redeemInterval() view returns(uint48)
func (_SimpleFaucet *SimpleFaucetCaller) RedeemInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleFaucet.contract.Call(opts, &out, "redeemInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemInterval is a free data retrieval call binding the contract method 0xa5ff3f80.
//
// Solidity: function redeemInterval() view returns(uint48)
func (_SimpleFaucet *SimpleFaucetSession) RedeemInterval() (*big.Int, error) {
	return _SimpleFaucet.Contract.RedeemInterval(&_SimpleFaucet.CallOpts)
}

// RedeemInterval is a free data retrieval call binding the contract method 0xa5ff3f80.
//
// Solidity: function redeemInterval() view returns(uint48)
func (_SimpleFaucet *SimpleFaucetCallerSession) RedeemInterval() (*big.Int, error) {
	return _SimpleFaucet.Contract.RedeemInterval(&_SimpleFaucet.CallOpts)
}

// Redeemed is a free data retrieval call binding the contract method 0x9f4568ef.
//
// Solidity: function redeemed(address receiver) view returns(uint48 time)
func (_SimpleFaucet *SimpleFaucetCaller) Redeemed(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SimpleFaucet.contract.Call(opts, &out, "redeemed", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Redeemed is a free data retrieval call binding the contract method 0x9f4568ef.
//
// Solidity: function redeemed(address receiver) view returns(uint48 time)
func (_SimpleFaucet *SimpleFaucetSession) Redeemed(receiver common.Address) (*big.Int, error) {
	return _SimpleFaucet.Contract.Redeemed(&_SimpleFaucet.CallOpts, receiver)
}

// Redeemed is a free data retrieval call binding the contract method 0x9f4568ef.
//
// Solidity: function redeemed(address receiver) view returns(uint48 time)
func (_SimpleFaucet *SimpleFaucetCallerSession) Redeemed(receiver common.Address) (*big.Int, error) {
	return _SimpleFaucet.Contract.Redeemed(&_SimpleFaucet.CallOpts, receiver)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SimpleFaucet *SimpleFaucetCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SimpleFaucet.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SimpleFaucet *SimpleFaucetSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SimpleFaucet.Contract.SupportsInterface(&_SimpleFaucet.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SimpleFaucet *SimpleFaucetCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SimpleFaucet.Contract.SupportsInterface(&_SimpleFaucet.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_SimpleFaucet *SimpleFaucetCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleFaucet.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_SimpleFaucet *SimpleFaucetSession) Token() (common.Address, error) {
	return _SimpleFaucet.Contract.Token(&_SimpleFaucet.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_SimpleFaucet *SimpleFaucetCallerSession) Token() (common.Address, error) {
	return _SimpleFaucet.Contract.Token(&_SimpleFaucet.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SimpleFaucet *SimpleFaucetTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SimpleFaucet *SimpleFaucetSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.GrantRole(&_SimpleFaucet.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SimpleFaucet *SimpleFaucetTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.GrantRole(&_SimpleFaucet.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x77202ba1.
//
// Solidity: function initialize(address owner, address _token, uint256 _amount, uint48 _redeemInterval, address _redeemAdmin) returns()
func (_SimpleFaucet *SimpleFaucetTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, _token common.Address, _amount *big.Int, _redeemInterval *big.Int, _redeemAdmin common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.contract.Transact(opts, "initialize", owner, _token, _amount, _redeemInterval, _redeemAdmin)
}

// Initialize is a paid mutator transaction binding the contract method 0x77202ba1.
//
// Solidity: function initialize(address owner, address _token, uint256 _amount, uint48 _redeemInterval, address _redeemAdmin) returns()
func (_SimpleFaucet *SimpleFaucetSession) Initialize(owner common.Address, _token common.Address, _amount *big.Int, _redeemInterval *big.Int, _redeemAdmin common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.Initialize(&_SimpleFaucet.TransactOpts, owner, _token, _amount, _redeemInterval, _redeemAdmin)
}

// Initialize is a paid mutator transaction binding the contract method 0x77202ba1.
//
// Solidity: function initialize(address owner, address _token, uint256 _amount, uint48 _redeemInterval, address _redeemAdmin) returns()
func (_SimpleFaucet *SimpleFaucetTransactorSession) Initialize(owner common.Address, _token common.Address, _amount *big.Int, _redeemInterval *big.Int, _redeemAdmin common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.Initialize(&_SimpleFaucet.TransactOpts, owner, _token, _amount, _redeemInterval, _redeemAdmin)
}

// Redeem is a paid mutator transaction binding the contract method 0xbe040fb0.
//
// Solidity: function redeem() returns()
func (_SimpleFaucet *SimpleFaucetTransactor) Redeem(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleFaucet.contract.Transact(opts, "redeem")
}

// Redeem is a paid mutator transaction binding the contract method 0xbe040fb0.
//
// Solidity: function redeem() returns()
func (_SimpleFaucet *SimpleFaucetSession) Redeem() (*types.Transaction, error) {
	return _SimpleFaucet.Contract.Redeem(&_SimpleFaucet.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xbe040fb0.
//
// Solidity: function redeem() returns()
func (_SimpleFaucet *SimpleFaucetTransactorSession) Redeem() (*types.Transaction, error) {
	return _SimpleFaucet.Contract.Redeem(&_SimpleFaucet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SimpleFaucet *SimpleFaucetTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleFaucet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SimpleFaucet *SimpleFaucetSession) RenounceOwnership() (*types.Transaction, error) {
	return _SimpleFaucet.Contract.RenounceOwnership(&_SimpleFaucet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SimpleFaucet *SimpleFaucetTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SimpleFaucet.Contract.RenounceOwnership(&_SimpleFaucet.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SimpleFaucet *SimpleFaucetTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SimpleFaucet *SimpleFaucetSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.RenounceRole(&_SimpleFaucet.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SimpleFaucet *SimpleFaucetTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.RenounceRole(&_SimpleFaucet.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SimpleFaucet *SimpleFaucetTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SimpleFaucet *SimpleFaucetSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.RevokeRole(&_SimpleFaucet.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SimpleFaucet *SimpleFaucetTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.RevokeRole(&_SimpleFaucet.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SimpleFaucet *SimpleFaucetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SimpleFaucet *SimpleFaucetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.TransferOwnership(&_SimpleFaucet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SimpleFaucet *SimpleFaucetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.TransferOwnership(&_SimpleFaucet.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SimpleFaucet *SimpleFaucetTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SimpleFaucet *SimpleFaucetSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.UpgradeTo(&_SimpleFaucet.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_SimpleFaucet *SimpleFaucetTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.UpgradeTo(&_SimpleFaucet.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SimpleFaucet *SimpleFaucetTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SimpleFaucet.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SimpleFaucet *SimpleFaucetSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.UpgradeToAndCall(&_SimpleFaucet.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SimpleFaucet *SimpleFaucetTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.UpgradeToAndCall(&_SimpleFaucet.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_SimpleFaucet *SimpleFaucetTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleFaucet.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_SimpleFaucet *SimpleFaucetSession) Withdraw() (*types.Transaction, error) {
	return _SimpleFaucet.Contract.Withdraw(&_SimpleFaucet.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_SimpleFaucet *SimpleFaucetTransactorSession) Withdraw() (*types.Transaction, error) {
	return _SimpleFaucet.Contract.Withdraw(&_SimpleFaucet.TransactOpts)
}

// SimpleFaucetAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the SimpleFaucet contract.
type SimpleFaucetAdminChangedIterator struct {
	Event *SimpleFaucetAdminChanged // Event containing the contract specifics and raw log

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
func (it *SimpleFaucetAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleFaucetAdminChanged)
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
		it.Event = new(SimpleFaucetAdminChanged)
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
func (it *SimpleFaucetAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleFaucetAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleFaucetAdminChanged represents a AdminChanged event raised by the SimpleFaucet contract.
type SimpleFaucetAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SimpleFaucet *SimpleFaucetFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*SimpleFaucetAdminChangedIterator, error) {

	logs, sub, err := _SimpleFaucet.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &SimpleFaucetAdminChangedIterator{contract: _SimpleFaucet.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_SimpleFaucet *SimpleFaucetFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *SimpleFaucetAdminChanged) (event.Subscription, error) {

	logs, sub, err := _SimpleFaucet.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleFaucetAdminChanged)
				if err := _SimpleFaucet.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_SimpleFaucet *SimpleFaucetFilterer) ParseAdminChanged(log types.Log) (*SimpleFaucetAdminChanged, error) {
	event := new(SimpleFaucetAdminChanged)
	if err := _SimpleFaucet.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleFaucetBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the SimpleFaucet contract.
type SimpleFaucetBeaconUpgradedIterator struct {
	Event *SimpleFaucetBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *SimpleFaucetBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleFaucetBeaconUpgraded)
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
		it.Event = new(SimpleFaucetBeaconUpgraded)
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
func (it *SimpleFaucetBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleFaucetBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleFaucetBeaconUpgraded represents a BeaconUpgraded event raised by the SimpleFaucet contract.
type SimpleFaucetBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SimpleFaucet *SimpleFaucetFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*SimpleFaucetBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _SimpleFaucet.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &SimpleFaucetBeaconUpgradedIterator{contract: _SimpleFaucet.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_SimpleFaucet *SimpleFaucetFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *SimpleFaucetBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _SimpleFaucet.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleFaucetBeaconUpgraded)
				if err := _SimpleFaucet.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_SimpleFaucet *SimpleFaucetFilterer) ParseBeaconUpgraded(log types.Log) (*SimpleFaucetBeaconUpgraded, error) {
	event := new(SimpleFaucetBeaconUpgraded)
	if err := _SimpleFaucet.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleFaucetInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SimpleFaucet contract.
type SimpleFaucetInitializedIterator struct {
	Event *SimpleFaucetInitialized // Event containing the contract specifics and raw log

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
func (it *SimpleFaucetInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleFaucetInitialized)
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
		it.Event = new(SimpleFaucetInitialized)
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
func (it *SimpleFaucetInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleFaucetInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleFaucetInitialized represents a Initialized event raised by the SimpleFaucet contract.
type SimpleFaucetInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SimpleFaucet *SimpleFaucetFilterer) FilterInitialized(opts *bind.FilterOpts) (*SimpleFaucetInitializedIterator, error) {

	logs, sub, err := _SimpleFaucet.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SimpleFaucetInitializedIterator{contract: _SimpleFaucet.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SimpleFaucet *SimpleFaucetFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SimpleFaucetInitialized) (event.Subscription, error) {

	logs, sub, err := _SimpleFaucet.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleFaucetInitialized)
				if err := _SimpleFaucet.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SimpleFaucet *SimpleFaucetFilterer) ParseInitialized(log types.Log) (*SimpleFaucetInitialized, error) {
	event := new(SimpleFaucetInitialized)
	if err := _SimpleFaucet.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleFaucetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SimpleFaucet contract.
type SimpleFaucetOwnershipTransferredIterator struct {
	Event *SimpleFaucetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SimpleFaucetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleFaucetOwnershipTransferred)
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
		it.Event = new(SimpleFaucetOwnershipTransferred)
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
func (it *SimpleFaucetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleFaucetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleFaucetOwnershipTransferred represents a OwnershipTransferred event raised by the SimpleFaucet contract.
type SimpleFaucetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SimpleFaucet *SimpleFaucetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SimpleFaucetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SimpleFaucet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SimpleFaucetOwnershipTransferredIterator{contract: _SimpleFaucet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SimpleFaucet *SimpleFaucetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SimpleFaucetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SimpleFaucet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleFaucetOwnershipTransferred)
				if err := _SimpleFaucet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SimpleFaucet *SimpleFaucetFilterer) ParseOwnershipTransferred(log types.Log) (*SimpleFaucetOwnershipTransferred, error) {
	event := new(SimpleFaucetOwnershipTransferred)
	if err := _SimpleFaucet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleFaucetRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SimpleFaucet contract.
type SimpleFaucetRoleAdminChangedIterator struct {
	Event *SimpleFaucetRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SimpleFaucetRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleFaucetRoleAdminChanged)
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
		it.Event = new(SimpleFaucetRoleAdminChanged)
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
func (it *SimpleFaucetRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleFaucetRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleFaucetRoleAdminChanged represents a RoleAdminChanged event raised by the SimpleFaucet contract.
type SimpleFaucetRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SimpleFaucet *SimpleFaucetFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SimpleFaucetRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SimpleFaucet.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SimpleFaucetRoleAdminChangedIterator{contract: _SimpleFaucet.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SimpleFaucet *SimpleFaucetFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SimpleFaucetRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SimpleFaucet.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleFaucetRoleAdminChanged)
				if err := _SimpleFaucet.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SimpleFaucet *SimpleFaucetFilterer) ParseRoleAdminChanged(log types.Log) (*SimpleFaucetRoleAdminChanged, error) {
	event := new(SimpleFaucetRoleAdminChanged)
	if err := _SimpleFaucet.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleFaucetRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SimpleFaucet contract.
type SimpleFaucetRoleGrantedIterator struct {
	Event *SimpleFaucetRoleGranted // Event containing the contract specifics and raw log

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
func (it *SimpleFaucetRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleFaucetRoleGranted)
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
		it.Event = new(SimpleFaucetRoleGranted)
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
func (it *SimpleFaucetRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleFaucetRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleFaucetRoleGranted represents a RoleGranted event raised by the SimpleFaucet contract.
type SimpleFaucetRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SimpleFaucet *SimpleFaucetFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SimpleFaucetRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SimpleFaucet.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SimpleFaucetRoleGrantedIterator{contract: _SimpleFaucet.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SimpleFaucet *SimpleFaucetFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SimpleFaucetRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SimpleFaucet.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleFaucetRoleGranted)
				if err := _SimpleFaucet.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SimpleFaucet *SimpleFaucetFilterer) ParseRoleGranted(log types.Log) (*SimpleFaucetRoleGranted, error) {
	event := new(SimpleFaucetRoleGranted)
	if err := _SimpleFaucet.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleFaucetRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SimpleFaucet contract.
type SimpleFaucetRoleRevokedIterator struct {
	Event *SimpleFaucetRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SimpleFaucetRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleFaucetRoleRevoked)
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
		it.Event = new(SimpleFaucetRoleRevoked)
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
func (it *SimpleFaucetRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleFaucetRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleFaucetRoleRevoked represents a RoleRevoked event raised by the SimpleFaucet contract.
type SimpleFaucetRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SimpleFaucet *SimpleFaucetFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SimpleFaucetRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SimpleFaucet.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SimpleFaucetRoleRevokedIterator{contract: _SimpleFaucet.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SimpleFaucet *SimpleFaucetFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SimpleFaucetRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SimpleFaucet.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleFaucetRoleRevoked)
				if err := _SimpleFaucet.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SimpleFaucet *SimpleFaucetFilterer) ParseRoleRevoked(log types.Log) (*SimpleFaucetRoleRevoked, error) {
	event := new(SimpleFaucetRoleRevoked)
	if err := _SimpleFaucet.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleFaucetUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SimpleFaucet contract.
type SimpleFaucetUpgradedIterator struct {
	Event *SimpleFaucetUpgraded // Event containing the contract specifics and raw log

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
func (it *SimpleFaucetUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleFaucetUpgraded)
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
		it.Event = new(SimpleFaucetUpgraded)
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
func (it *SimpleFaucetUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleFaucetUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleFaucetUpgraded represents a Upgraded event raised by the SimpleFaucet contract.
type SimpleFaucetUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SimpleFaucet *SimpleFaucetFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SimpleFaucetUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SimpleFaucet.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SimpleFaucetUpgradedIterator{contract: _SimpleFaucet.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SimpleFaucet *SimpleFaucetFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SimpleFaucetUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SimpleFaucet.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleFaucetUpgraded)
				if err := _SimpleFaucet.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_SimpleFaucet *SimpleFaucetFilterer) ParseUpgraded(log types.Log) (*SimpleFaucetUpgraded, error) {
	event := new(SimpleFaucetUpgraded)
	if err := _SimpleFaucet.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
