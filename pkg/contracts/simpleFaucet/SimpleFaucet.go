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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REDEEM_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"_redeemInterval\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"_redeemAdmin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemInterval\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"redeemTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"redeemed\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"time\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a080604052346100315730608052611cfa9081610037823960805181818161098401528181610ae7015261109d0152f35b600080fdfe608060408181526004918236101561001657600080fd5b600092833560e01c91826301ffc9a7146113d157508163248a9ca3146113a65781632f2ff15d146112f957816336568abe146112675781633659cfe6146110755781633ccfd60b14610d6a5781634f1ef28614610a6c5781635018c20914610a3157816352d1902d1461096e578163715018a61461090e57816377202ba1146106405781638da5cb5b1461061757816391d14854146105d05781639f4568ef1461058e578163a217fddf14610573578163a2d2c71c14610545578163a5ff3f801461051d578163aa8c217c146104fd578163be040fb0146103cd578163d547741f14610391578163f2fde38b146102ff578163f984c55c1461014e575063fc0c546a1461012257600080fd5b3461014a578160031936011261014a5761012d5490516001600160a01b039091168152602090f35b5080fd5b919050346102fb576020806003193601126102f7578361016c61143f565b61012f5465ffffffffffff42811696929181169081156102da576101a06101a992338752610130885282898820541661152d565b1687101561155d565b61012d5485516370a0823160e01b81523084820152926001600160a01b0390911691908584602481865afa9384156102d0579086949392918694610297575b5094610232956101fe61012e54809610156115ba565b885163a9059cbb60e01b81526001600160a01b03909316908301908152602081019490945290948593849291839160400190565b03925af1801561028d57906101309291610260575b503385525282209065ffffffffffff1982541617905580f35b61027f90823d8411610286575b61027781836114a3565b810190611611565b5038610247565b503d61026d565b83513d87823e3d90fd5b8581969295503d83116102c9575b6102af81836114a3565b810103126102c5579251859390926102326101e8565b8480fd5b503d6102a5565b87513d87823e3d90fd5b6102f2915033855261013086528685205416156114e1565b6101a9565b8380fd5b8280fd5b9050346102fb5760203660031901126102fb5761031a61143f565b916103236118fb565b6001600160a01b0383161561033f578361033c84611953565b80f35b906020608492519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152fd5b919050346102fb57806003193601126102fb5761033c91356103c860016103b6611424565b93838752609760205286200154611629565b611825565b919050346102fb57826003193601126102fb5761012f5465ffffffffffff428116939181169081156104df5761041461041d9233885261013060205282868920541661152d565b1684101561155d565b61012d5482516370a0823160e01b815230838201526020926001600160a01b03909216918382602481865afa9182156104d55790849291889261049e575b50926102329361047161012e54809410156115ba565b865163a9059cbb60e01b81523391810191825260208201939093529193849283918a918391604090910190565b8381949293503d83116104ce575b6104b681836114a3565b810103126104ca579051839161023261045b565b8680fd5b503d6104ac565b85513d89823e3d90fd5b6104f891503386526101306020528386205416156114e1565b61041d565b50503461014a578160031936011261014a5760209061012e549051908152f35b50503461014a578160031936011261014a5760209065ffffffffffff61012f54169051908152f35b50503461014a578160031936011261014a5761012f54905160309190911c6001600160a01b03168152602090f35b50503461014a578160031936011261014a5751908152602090f35b50503461014a57602036600319011261014a5760209165ffffffffffff9082906001600160a01b036105be61143f565b16815261013085522054169051908152f35b9050346102fb57816003193601126102fb578160209360ff926105f1611424565b90358252609786528282206001600160a01b039091168252855220549151911615158152f35b50503461014a578160031936011261014a5760335490516001600160a01b039091168152602090f35b9050346102fb5760a03660031901126102fb5761065b61143f565b6001600160a01b039160243583811692919083900361090a576064359365ffffffffffff85168095036104ca57608435948186168087036109065788549160ff97888460081c1615978880996108fa575b80156108e4575b1561088a5760ff198581166001178d559489610879575b506106e38a8d5460081c166106de8161189b565b61189b565b6106ec33611953565b6106fb8a8d5460081c1661189b565b61012d80546001600160a01b031916909117905560443561012e5561012f80546001600160d01b03191690921760309190911b6601000000000000600160d01b03161790557f281081d9b36b37208f0d8dfce5adc7e00d31ece09269aaa8d0bfa43e6840a338808a5260976020908152898b20838c528152898b205490989193911615610830575b50505061078e6118fb565b8216156107de575061079f90611953565b6107a7578280f35b7f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989161ff001984541684555160018152a138808280f35b845162461bcd60e51b8152908101849052602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b828a5260978852888a20828b5288526001898b209182541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8a80a4388080610783565b61ffff1916610101178c55386106ca565b8a5162461bcd60e51b8152602081890152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b50303b1580156106b3575060018a8616146106b3565b5060018a8616106106ac565b8880fd5b8580fd5b833461096b578060031936011261096b576109276118fb565b603380546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b80fd5b9050823461096b578060031936011261096b57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036109c95760208251600080516020611ca58339815191528152f35b6020608492519162461bcd60e51b8352820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152fd5b50503461014a578160031936011261014a57602090517f281081d9b36b37208f0d8dfce5adc7e00d31ece09269aaa8d0bfa43e6840a3388152f35b918091506003193601126102fb57610a8261143f565b906024359067ffffffffffffffff82116102c557366023830112156102c55781840135610aae816114c5565b610aba835191826114a3565b818152866020948583019336602482840101116102fb578060248893018637830101526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811690610b153083141561199c565b610b32600080516020611ca58339815191529282845416146119fd565b610b3a6118fb565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610b745750505050505061033c9150611b83565b869293949596169085516352d1902d60e01b815287818a81865afa8a9181610d37575b50610bf657865162461bcd60e51b8152808a01899052602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608490fd5b979192939695949703610ce25750610c0d82611b83565b7fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8780a285845115801590610cda575b610c4b575b50505050505080f35b80610cc496845196610c5c88611455565b602788527f416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c87890152660819985a5b195960ca1b868901525190845af4913d15610cd0573d610cb6610cad826114c5565b925192836114a3565b81528681943d92013e611c13565b50388080808085610c42565b5060609250611c13565b506001610c3d565b835162461bcd60e51b8152908101859052602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608490fd5b9091508881813d8311610d63575b610d4f81836114a3565b81010312610d5f57519038610b97565b8a80fd5b503d610d45565b919050346102fb57826003193601126102fb577f281081d9b36b37208f0d8dfce5adc7e00d31ece09269aaa8d0bfa43e6840a3389182845260209260978452828520338652845260ff838620541615610eb4575061012d5461012f5483516370a0823160e01b8152308185015286949390926001600160a01b039081169260301c16908684602481865afa938415610eaa579087949392918794610e75575b50855163a9059cbb60e01b81526001600160a01b03909216908201908152602081019390935294859283919082906040015b03925af1908115610e6c5750610e4f578280f35b81610e6592903d106102865761027781836114a3565b5038808280f35b513d85823e3d90fd5b8581969295503d8311610ea3575b610e8d81836114a3565b8101031261090a57925186939092610e3b610e09565b503d610e83565b85513d88823e3d90fd5b848491610ec033611a85565b855191610ecc83611487565b60428352848301936060368637835115611062576030855383519060019182101561104f5790607860218601536041915b818311610fe457505050610fa257610f9e938693610f8a93610f7b604894610f529a519a8576020b1b1b2b9b9a1b7b73a3937b61d1030b1b1b7bab73a1604d1b8d9788015282519283916037890191016117d6565b8401917001034b99036b4b9b9b4b733903937b6329607d1b6037840152518093868401906117d6565b010360288101875201856114a3565b5162461bcd60e51b815292839283016117f9565b0390fd5b50505080606493519262461bcd60e51b845283015260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152fd5b909192600f8116601081101561103c576f181899199a1a9b1b9c1cb0b131b232b360811b901a6110148588611a5e565b53881c92801561102957600019019190610efd565b634e487b7160e01b825260118952602482fd5b634e487b7160e01b835260328a52602483fd5b634e487b7160e01b815260328852602490fd5b634e487b7160e01b815260328752602490fd5b919050346102fb576020806003193601126102f75761109261143f565b916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166110ca3082141561199c565b6110e7600080516020611ca58339815191529183835416146119fd565b6110ef6118fb565b8251908482019282841067ffffffffffffffff851117611254578385528883527f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff16156111495750505050505061033c9150611b83565b869293949596169085516352d1902d60e01b815287818a81865afa8a9181611225575b506111cb57865162461bcd60e51b8152808a01899052602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608490fd5b979192939695949703610ce257506111e282611b83565b7fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8780a28584511580159061121e57610c4b5750505050505080f35b5080610c3d565b9091508881813d831161124d575b61123d81836114a3565b81010312610d5f5751903861116c565b503d611233565b634e487b7160e01b895260418852602489fd5b8391503461014a578260031936011261014a57611282611424565b90336001600160a01b0383160361129e579061033c9135611825565b608490602085519162461bcd60e51b8352820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152fd5b9050346102fb57816003193601126102fb573590611315611424565b90828452609760205261132d60018286200154611629565b828452609760209081528185206001600160a01b039093168086529290528084205460ff161561135b578380f35b82845260976020528084208285526020528320600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8480a43880808380f35b9050346102fb5760203660031901126102fb5781602093600192358152609785522001549051908152f35b8491346102fb5760203660031901126102fb573563ffffffff60e01b81168091036102fb5760209250637965db0b60e01b8114908115611413575b5015158152f35b6301ffc9a760e01b1490508361140c565b602435906001600160a01b038216820361143a57565b600080fd5b600435906001600160a01b038216820361143a57565b6060810190811067ffffffffffffffff82111761147157604052565b634e487b7160e01b600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761147157604052565b90601f8019910116810190811067ffffffffffffffff82111761147157604052565b67ffffffffffffffff811161147157601f01601f191660200190565b156114e857565b60405162461bcd60e51b815260206004820152601e60248201527f53696d706c654661756365743a20616c72656164792072656465656d656400006044820152606490fd5b91909165ffffffffffff8080941691160191821161154757565b634e487b7160e01b600052601160045260246000fd5b1561156457565b60405162461bcd60e51b815260206004820152602860248201527f53696d706c654661756365743a2072656465656d20696e74657276616c206e6f6044820152671d081c185cdcd95960c21b6064820152608490fd5b156115c157565b60405162461bcd60e51b815260206004820152602260248201527f53696d706c654661756365743a20696e73756666696369656e742062616c616e604482015261636560f01b6064820152608490fd5b9081602091031261143a5751801515810361143a5790565b600081815260209060978252604092838220338352835260ff8483205416156116525750505050565b61165b33611a85565b84519161166783611487565b604283528483019360603686378351156117c257603085538351906001918210156117c25790607860218601536041915b8183116117545750505061171257610f529385936116fc936116ed604894610f9e9951988576020b1b1b2b9b9a1b7b73a3937b61d1030b1b1b7bab73a1604d1b8b9788015282519283916037890191016117d6565b010360288101855201836114a3565b5162461bcd60e51b8152918291600483016117f9565b60648486519062461bcd60e51b825280600483015260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152fd5b909192600f811660108110156117ae576f181899199a1a9b1b9c1cb0b131b232b360811b901a6117848588611a5e565b5360041c92801561179a57600019019190611698565b634e487b7160e01b82526011600452602482fd5b634e487b7160e01b83526032600452602483fd5b634e487b7160e01b81526032600452602490fd5b60005b8381106117e95750506000910152565b81810151838201526020016117d9565b6040916020825261181981518092816020860152602086860191016117d6565b601f01601f1916010190565b906000918083526097602052604083209160018060a01b03169182845260205260ff60408420541661185657505050565b8083526097602052604083208284526020526040832060ff1981541690557ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b339380a4565b156118a257565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b6033546001600160a01b0316330361190f57565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b603380546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b156119a357565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608490fd5b15611a0457565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b6064820152608490fd5b908151811015611a6f570160200190565b634e487b7160e01b600052603260045260246000fd5b60405190611a9282611455565b602a8252602082016040368237825115611a6f57603090538151600190811015611a6f57607860218401536029905b808211611b15575050611ad15790565b606460405162461bcd60e51b815260206004820152602060248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152fd5b9091600f81166010811015611b6e576f181899199a1a9b1b9c1cb0b131b232b360811b901a611b448486611a5e565b5360041c918015611b59576000190190611ac1565b60246000634e487b7160e01b81526011600452fd5b60246000634e487b7160e01b81526032600452fd5b803b15611bb857600080516020611ca583398151915280546001600160a01b0319166001600160a01b03909216919091179055565b60405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608490fd5b91929015611c755750815115611c27575090565b3b15611c305790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015611c885750805190602001fd5b60405162461bcd60e51b8152908190610f9e90600483016117f956fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca2646970667358221220e6df4789e039dd5feadbe1815d6b155b9784c3ce8afd48f33859d1225990534764736f6c63430008140033",
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

// RedeemTo is a paid mutator transaction binding the contract method 0xf984c55c.
//
// Solidity: function redeemTo(address to) returns()
func (_SimpleFaucet *SimpleFaucetTransactor) RedeemTo(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.contract.Transact(opts, "redeemTo", to)
}

// RedeemTo is a paid mutator transaction binding the contract method 0xf984c55c.
//
// Solidity: function redeemTo(address to) returns()
func (_SimpleFaucet *SimpleFaucetSession) RedeemTo(to common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.RedeemTo(&_SimpleFaucet.TransactOpts, to)
}

// RedeemTo is a paid mutator transaction binding the contract method 0xf984c55c.
//
// Solidity: function redeemTo(address to) returns()
func (_SimpleFaucet *SimpleFaucetTransactorSession) RedeemTo(to common.Address) (*types.Transaction, error) {
	return _SimpleFaucet.Contract.RedeemTo(&_SimpleFaucet.TransactOpts, to)
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
