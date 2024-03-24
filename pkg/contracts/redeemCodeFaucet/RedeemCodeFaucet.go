// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package redeemCodeFaucet

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

// RedeemCodeFaucetMetaData contains all meta data concerning the RedeemCodeFaucet contract.
var RedeemCodeFaucetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CodeRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REDEEM_CODE_CREATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"redeemableAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"}],\"name\":\"addRedeemCode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callInterval\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"codeCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_codeCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"_callInterval\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"_codeCreator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"}],\"name\":\"redeemed\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"time\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a080604052346100315730608052611ddf90816100378239608051818181610580015281816106a80152610aaa0152f35b600080fdfe608060408181526004918236101561001657600080fd5b600092833560e01c91826301ffc9a7146110ed57508163163df9e814610e4d578163248a9ca314610e225781632f2ff15d14610d0a57816336568abe14610c785781633659cfe614610a82578382633ccfd60b14610960575081633d8b9e071461092f5781634f1ef2861461062d57816352d1902d1461056a578163715018a61461050a5781638da5cb5b146104e157816391d148541461049a578163931409081461047157816393c0a4aa14610447578163a217fddf1461042c578163d547741f146103f0578163db006a75146103d3578163e939509614610240578163ed05582b14610210578163f2fde38b1461017a57508063f8ec16de146101505763fc0c546a1461012457600080fd5b3461014c578160031936011261014c5761012d5490516001600160a01b039091168152602090f35b5080fd5b503461014c578160031936011261014c5760209065ffffffffffff61012d5460a01c169051908152f35b90503461020c57602036600319011261020c57610195611140565b9161019e6119c0565b6001600160a01b038316156101ba57836101b784611a18565b80f35b906020608492519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152fd5b8280fd5b50503461014c578060031936011261014c57602090610239610230611140565b60243590611212565b9051908152f35b9190503461020c57606036600319011261020c576024358235610261611171565b9361026a61158d565b821561036a5765ffffffffffff8095169480421686111561030c5782875261013160205284872054166102be5750845261012f6020528184205561013060205282209065ffffffffffff1982541617905580f35b608490602085519162461bcd60e51b8352820152602260248201527f52656465656d436f64654661756365743a20616c72656164792072656465656d604482015261195960f21b6064820152fd5b845162461bcd60e51b8152602081840152603260248201527f52656465656d436f64654661756365743a2076616c6964556e74696c206d75736044820152717420626520696e207468652066757475726560701b6064820152608490fd5b608490602085519162461bcd60e51b8352820152603c60248201527f52656465656d436f64654661756365743a2072656465656d61626c65416d6f7560448201527f6e74206d7573742062652067726561746572207468616e207a65726f000000006064820152fd5b83903461014c57602036600319011261014c576101b790356112b1565b9190503461020c578060031936011261020c576101b79135610427600161041561115b565b93838752609760205286200154611791565b6118ea565b50503461014c578160031936011261014c5751908152602090f35b50503461014c578160031936011261014c5761012e5490516001600160a01b039091168152602090f35b50503461014c578160031936011261014c5760209051600080516020611d6a8339815191528152f35b90503461020c578160031936011261020c578160209360ff926104bb61115b565b90358252609786528282206001600160a01b039091168252855220549151911615158152f35b50503461014c578160031936011261014c5760335490516001600160a01b039091168152602090f35b83346105675780600319360112610567576105236119c0565b603380546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b80fd5b90508234610567578060031936011261056757507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031630036105c55760208251600080516020611d8a8339815191528152f35b6020608492519162461bcd60e51b8352820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152fd5b9180915060031936011261020c57610643611140565b906024359067ffffffffffffffff821161092b573660238301121561092b578184013561066f816111f6565b61067b835191826111d4565b8181528660209485830193366024828401011161020c578060248893018637830101526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116906106d630831415611a61565b6106f3600080516020611d8a833981519152928284541614611ac2565b6106fb6119c0565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610735575050505050506101b79150611c48565b869293949596169085516352d1902d60e01b815287818a81865afa8a91816108f8575b506107b757865162461bcd60e51b8152808a01899052602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608490fd5b9791929396959497036108a357506107ce82611c48565b7fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8780a28584511580159061089b575b61080c575b50505050505080f35b806108859684519661081d88611186565b602788527f416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c87890152660819985a5b195960ca1b868901525190845af4913d15610891573d61087761086e826111f6565b925192836111d4565b81528681943d92013e611cd8565b50388080808085610803565b5060609250611cd8565b5060016107fe565b835162461bcd60e51b8152908101859052602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608490fd5b9091508881813d8311610924575b61091081836111d4565b8101031261092057519038610758565b8a80fd5b503d610906565b8480fd5b90503461020c57602036600319011261020c578160209365ffffffffffff9235815261013185522054169051908152f35b9150913461014c578160031936011261014c5761097b61158d565b61012d5461012e5482516370a0823160e01b8152308682015260209590926001600160a01b039081169216908684602481865afa938415610a78579087949392918794610a3f575b50855163a9059cbb60e01b81526001600160a01b03909216908201908152602081019390935294859283919082906040015b03925af1908115610a365750610a09578280f35b81610a2892903d10610a2f575b610a2081836111d4565b810190611299565b5038808280f35b503d610a16565b513d85823e3d90fd5b8581969295503d8311610a71575b610a5781836111d4565b81010312610a6d579251869390926109f56109c3565b8580fd5b503d610a4d565b85513d88823e3d90fd5b9190503461020c57602080600319360112610c7457610a9f611140565b916001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116610ad730821415611a61565b610af4600080516020611d8a833981519152918383541614611ac2565b610afc6119c0565b8251908482019282841067ffffffffffffffff851117610c61578385528883527f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610b56575050505050506101b79150611c48565b869293949596169085516352d1902d60e01b815287818a81865afa8a9181610c32575b50610bd857865162461bcd60e51b8152808a01899052602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608490fd5b9791929396959497036108a35750610bef82611c48565b7fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8780a285845115801590610c2b5761080c5750505050505080f35b50806107fe565b9091508881813d8311610c5a575b610c4a81836111d4565b8101031261092057519038610b79565b503d610c40565b634e487b7160e01b895260418852602489fd5b8380fd5b8391503461014c578260031936011261014c57610c9361115b565b90336001600160a01b03831603610caf57906101b791356118ea565b608490602085519162461bcd60e51b8352820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152fd5b90503461020c578160031936011261020c57803591610d2761115b565b9183855260209060978252610d4160018488200154611791565b600080516020611d6a8339815191528514610dcc5750838552609781528185206001600160a01b039093168086529281528185205460ff1615610d82578480f35b8385526097815281852083865290528320805460ff1916600117905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8480a4388080808480f35b915162461bcd60e51b815291820152602c60248201527f416363657373436f6e74726f6c3a2063616e6e6f74206772616e7420636f646560448201526b43726561746f7220726f6c6560a01b6064820152608490fd5b90503461020c57602036600319011261020c5781602093600192358152609785522001549051908152f35b90503461020c57608036600319011261020c57610e68611140565b6024356001600160a01b0381811693909291849003610a6d57610e89611171565b926064358181168091036110e95787549060ff96878360081c1615968780986110dd575b80156110c7575b1561106d5760ff198481166001178c55938861105c575b50610ee4898c5460081c16610edf81611960565b611960565b610eed33611a18565b610efc898c5460081c16611960565b61012d80546001600160d01b03191690921760a09190911b65ffffffffffff60a01b1617905561012e80546001600160a01b03191682179055600080516020611d6a833981519152808a5260976020908152898b20838c528152898b205490989193911615611013575b505050610f716119c0565b821615610fc15750610f8290611a18565b610f8a578280f35b7f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989161ff001984541684555160018152a138808280f35b845162461bcd60e51b8152908101849052602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b828a5260978852888a20828b5288526001898b209182541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8a80a4388080610f66565b61ffff1916610101178b5538610ecb565b895162461bcd60e51b8152602081880152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b50303b158015610eb45750600189851614610eb4565b50600189851610610ead565b8780fd5b84913461020c57602036600319011261020c573563ffffffff60e01b811680910361020c5760209250637965db0b60e01b811490811561112f575b5015158152f35b6301ffc9a760e01b14905083611128565b600435906001600160a01b038216820361115657565b600080fd5b602435906001600160a01b038216820361115657565b6044359065ffffffffffff8216820361115657565b6060810190811067ffffffffffffffff8211176111a257604052565b634e487b7160e01b600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176111a257604052565b90601f8019910116810190811067ffffffffffffffff8211176111a257604052565b67ffffffffffffffff81116111a257601f01601f191660200190565b906040519060208201926bffffffffffffffffffffffff199060601b16835260348201524660548201523060601b60748201526068815260a0810181811067ffffffffffffffff8211176111a25760405251902090565b91909165ffffffffffff8080941691160191821161128357565b634e487b7160e01b600052601160045260246000fd5b90816020910312611156575180151581036111565790565b9065ffffffffffff80421661012d91808354928160009788953387528361013396879281806020809c87825260409c8d9283832065ffffffffffff199e8f9860a01c168882541617905561013281528383208981549889161790555220541691169061131c91611269565b161015611562575061012e546001600160a01b039161133c918316611212565b96878a526101319283885280878c20541661150e57888b526101308852868b2054168410156114ca5754169584516370a0823160e01b815230600482015286816024818b5afa90811561143c578a91611499575b50818a5261012f90818852868b20541161144657818a52865284892054855163a9059cbb60e01b81523360048201526024810182905290978790829060449082908e905af1801561143c57918693917f739f8ee835b17b45c3440318f23a9a52f3b7b3194bd17cd54344eeb2aa59d908999a9b9361141f575b50825286522091825416179055519283523392a2565b61143590893d8b11610a2f57610a2081836111d4565b5038611409565b86513d8c823e3d90fd5b855162461bcd60e51b815260048101889052602660248201527f52656465656d436f64654661756365743a20696e73756666696369656e742062604482015265616c616e636560d01b6064820152608490fd5b90508681813d83116114c3575b6114b081836111d4565b810103126114bf575138611390565b8980fd5b503d6114a6565b855162461bcd60e51b815260048101889052601e60248201527f52656465656d436f64654661756365743a20636f6465206578706972656400006044820152606490fd5b865162461bcd60e51b815260048101899052602760248201527f52656465656d436f64654661756365743a20636f646520616c72656164792072604482015266195919595b595960ca1b6064820152608490fd5b9250508096979294955460a01c1694338352522092611585845493828516611269565b169116179055565b3360009081527f3360641542e133d9ce267924c5018ab505f382f31072a9fc8400a9ffb02e864b6020908152604080832054909290600080516020611d6a8339815191529060ff16156115e05750505050565b6115e933611b4a565b8451916115f5836111b8565b6042835284830193606036863783511561177d576030855383519060019182101561177d5790607860218601536041915b81831161170f575050506116cd5761167b9385936116b3936116a46048946116c99951988576020b1b1b2b9b9a1b7b73a3937b61d1030b1b1b7bab73a1604d1b8b97880152825192839160378901910161189b565b8401917001034b99036b4b9b9b4b733903937b6329607d1b60378401525180938684019061189b565b010360288101855201836111d4565b5162461bcd60e51b8152918291600483016118be565b0390fd5b60648486519062461bcd60e51b825280600483015260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152fd5b909192600f81166010811015611769576f181899199a1a9b1b9c1cb0b131b232b360811b901a61173f8588611b23565b5360041c92801561175557600019019190611626565b634e487b7160e01b82526011600452602482fd5b634e487b7160e01b83526032600452602483fd5b634e487b7160e01b81526032600452602490fd5b600081815260209060978252604092838220338352835260ff8483205416156117ba5750505050565b6117c333611b4a565b8451916117cf836111b8565b6042835284830193606036863783511561177d576030855383519060019182101561177d5790607860218601536041915b818311611855575050506116cd5761167b9385936116b3936116a46048946116c99951988576020b1b1b2b9b9a1b7b73a3937b61d1030b1b1b7bab73a1604d1b8b97880152825192839160378901910161189b565b909192600f81166010811015611769576f181899199a1a9b1b9c1cb0b131b232b360811b901a6118858588611b23565b5360041c92801561175557600019019190611800565b60005b8381106118ae5750506000910152565b818101518382015260200161189e565b604091602082526118de815180928160208601526020868601910161189b565b601f01601f1916010190565b906000918083526097602052604083209160018060a01b03169182845260205260ff60408420541661191b57505050565b8083526097602052604083208284526020526040832060ff1981541690557ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b339380a4565b1561196757565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b6033546001600160a01b031633036119d457565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b603380546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b15611a6857565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608490fd5b15611ac957565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b6064820152608490fd5b908151811015611b34570160200190565b634e487b7160e01b600052603260045260246000fd5b60405190611b5782611186565b602a8252602082016040368237825115611b3457603090538151600190811015611b3457607860218401536029905b808211611bda575050611b965790565b606460405162461bcd60e51b815260206004820152602060248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152fd5b9091600f81166010811015611c33576f181899199a1a9b1b9c1cb0b131b232b360811b901a611c098486611b23565b5360041c918015611c1e576000190190611b86565b60246000634e487b7160e01b81526011600452fd5b60246000634e487b7160e01b81526032600452fd5b803b15611c7d57600080516020611d8a83398151915280546001600160a01b0319166001600160a01b03909216919091179055565b60405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608490fd5b91929015611d3a5750815115611cec575090565b3b15611cf55790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015611d4d5750805190602001fd5b60405162461bcd60e51b81529081906116c990600483016118be56fe707a390ecff002f977980b7598d697e51a8a18e7708b8f7ada4c1e67e5ea4808360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca2646970667358221220633e36e1486aacfa5fb7959efae2432138345b9575eabd812baabe86791a80ed64736f6c63430008140033",
}

// RedeemCodeFaucetABI is the input ABI used to generate the binding from.
// Deprecated: Use RedeemCodeFaucetMetaData.ABI instead.
var RedeemCodeFaucetABI = RedeemCodeFaucetMetaData.ABI

// RedeemCodeFaucetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RedeemCodeFaucetMetaData.Bin instead.
var RedeemCodeFaucetBin = RedeemCodeFaucetMetaData.Bin

// DeployRedeemCodeFaucet deploys a new Ethereum contract, binding an instance of RedeemCodeFaucet to it.
func DeployRedeemCodeFaucet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RedeemCodeFaucet, error) {
	parsed, err := RedeemCodeFaucetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RedeemCodeFaucetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RedeemCodeFaucet{RedeemCodeFaucetCaller: RedeemCodeFaucetCaller{contract: contract}, RedeemCodeFaucetTransactor: RedeemCodeFaucetTransactor{contract: contract}, RedeemCodeFaucetFilterer: RedeemCodeFaucetFilterer{contract: contract}}, nil
}

// RedeemCodeFaucet is an auto generated Go binding around an Ethereum contract.
type RedeemCodeFaucet struct {
	RedeemCodeFaucetCaller     // Read-only binding to the contract
	RedeemCodeFaucetTransactor // Write-only binding to the contract
	RedeemCodeFaucetFilterer   // Log filterer for contract events
}

// RedeemCodeFaucetCaller is an auto generated read-only Go binding around an Ethereum contract.
type RedeemCodeFaucetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemCodeFaucetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RedeemCodeFaucetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemCodeFaucetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RedeemCodeFaucetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemCodeFaucetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RedeemCodeFaucetSession struct {
	Contract     *RedeemCodeFaucet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RedeemCodeFaucetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RedeemCodeFaucetCallerSession struct {
	Contract *RedeemCodeFaucetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// RedeemCodeFaucetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RedeemCodeFaucetTransactorSession struct {
	Contract     *RedeemCodeFaucetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RedeemCodeFaucetRaw is an auto generated low-level Go binding around an Ethereum contract.
type RedeemCodeFaucetRaw struct {
	Contract *RedeemCodeFaucet // Generic contract binding to access the raw methods on
}

// RedeemCodeFaucetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RedeemCodeFaucetCallerRaw struct {
	Contract *RedeemCodeFaucetCaller // Generic read-only contract binding to access the raw methods on
}

// RedeemCodeFaucetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RedeemCodeFaucetTransactorRaw struct {
	Contract *RedeemCodeFaucetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRedeemCodeFaucet creates a new instance of RedeemCodeFaucet, bound to a specific deployed contract.
func NewRedeemCodeFaucet(address common.Address, backend bind.ContractBackend) (*RedeemCodeFaucet, error) {
	contract, err := bindRedeemCodeFaucet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucet{RedeemCodeFaucetCaller: RedeemCodeFaucetCaller{contract: contract}, RedeemCodeFaucetTransactor: RedeemCodeFaucetTransactor{contract: contract}, RedeemCodeFaucetFilterer: RedeemCodeFaucetFilterer{contract: contract}}, nil
}

// NewRedeemCodeFaucetCaller creates a new read-only instance of RedeemCodeFaucet, bound to a specific deployed contract.
func NewRedeemCodeFaucetCaller(address common.Address, caller bind.ContractCaller) (*RedeemCodeFaucetCaller, error) {
	contract, err := bindRedeemCodeFaucet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetCaller{contract: contract}, nil
}

// NewRedeemCodeFaucetTransactor creates a new write-only instance of RedeemCodeFaucet, bound to a specific deployed contract.
func NewRedeemCodeFaucetTransactor(address common.Address, transactor bind.ContractTransactor) (*RedeemCodeFaucetTransactor, error) {
	contract, err := bindRedeemCodeFaucet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetTransactor{contract: contract}, nil
}

// NewRedeemCodeFaucetFilterer creates a new log filterer instance of RedeemCodeFaucet, bound to a specific deployed contract.
func NewRedeemCodeFaucetFilterer(address common.Address, filterer bind.ContractFilterer) (*RedeemCodeFaucetFilterer, error) {
	contract, err := bindRedeemCodeFaucet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetFilterer{contract: contract}, nil
}

// bindRedeemCodeFaucet binds a generic wrapper to an already deployed contract.
func bindRedeemCodeFaucet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RedeemCodeFaucetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RedeemCodeFaucet *RedeemCodeFaucetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RedeemCodeFaucet.Contract.RedeemCodeFaucetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RedeemCodeFaucet *RedeemCodeFaucetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RedeemCodeFaucetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RedeemCodeFaucet *RedeemCodeFaucetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RedeemCodeFaucetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RedeemCodeFaucet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.DEFAULTADMINROLE(&_RedeemCodeFaucet.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.DEFAULTADMINROLE(&_RedeemCodeFaucet.CallOpts)
}

// REDEEMCODECREATORROLE is a free data retrieval call binding the contract method 0x93140908.
//
// Solidity: function REDEEM_CODE_CREATOR_ROLE() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) REDEEMCODECREATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "REDEEM_CODE_CREATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REDEEMCODECREATORROLE is a free data retrieval call binding the contract method 0x93140908.
//
// Solidity: function REDEEM_CODE_CREATOR_ROLE() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) REDEEMCODECREATORROLE() ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.REDEEMCODECREATORROLE(&_RedeemCodeFaucet.CallOpts)
}

// REDEEMCODECREATORROLE is a free data retrieval call binding the contract method 0x93140908.
//
// Solidity: function REDEEM_CODE_CREATOR_ROLE() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) REDEEMCODECREATORROLE() ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.REDEEMCODECREATORROLE(&_RedeemCodeFaucet.CallOpts)
}

// CallInterval is a free data retrieval call binding the contract method 0xf8ec16de.
//
// Solidity: function callInterval() view returns(uint48)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) CallInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "callInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallInterval is a free data retrieval call binding the contract method 0xf8ec16de.
//
// Solidity: function callInterval() view returns(uint48)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) CallInterval() (*big.Int, error) {
	return _RedeemCodeFaucet.Contract.CallInterval(&_RedeemCodeFaucet.CallOpts)
}

// CallInterval is a free data retrieval call binding the contract method 0xf8ec16de.
//
// Solidity: function callInterval() view returns(uint48)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) CallInterval() (*big.Int, error) {
	return _RedeemCodeFaucet.Contract.CallInterval(&_RedeemCodeFaucet.CallOpts)
}

// CodeCreator is a free data retrieval call binding the contract method 0x93c0a4aa.
//
// Solidity: function codeCreator() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) CodeCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "codeCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CodeCreator is a free data retrieval call binding the contract method 0x93c0a4aa.
//
// Solidity: function codeCreator() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) CodeCreator() (common.Address, error) {
	return _RedeemCodeFaucet.Contract.CodeCreator(&_RedeemCodeFaucet.CallOpts)
}

// CodeCreator is a free data retrieval call binding the contract method 0x93c0a4aa.
//
// Solidity: function codeCreator() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) CodeCreator() (common.Address, error) {
	return _RedeemCodeFaucet.Contract.CodeCreator(&_RedeemCodeFaucet.CallOpts)
}

// GetHash is a free data retrieval call binding the contract method 0xed05582b.
//
// Solidity: function getHash(address _codeCreator, uint256 code) view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) GetHash(opts *bind.CallOpts, _codeCreator common.Address, code *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "getHash", _codeCreator, code)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0xed05582b.
//
// Solidity: function getHash(address _codeCreator, uint256 code) view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) GetHash(_codeCreator common.Address, code *big.Int) ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.GetHash(&_RedeemCodeFaucet.CallOpts, _codeCreator, code)
}

// GetHash is a free data retrieval call binding the contract method 0xed05582b.
//
// Solidity: function getHash(address _codeCreator, uint256 code) view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) GetHash(_codeCreator common.Address, code *big.Int) ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.GetHash(&_RedeemCodeFaucet.CallOpts, _codeCreator, code)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.GetRoleAdmin(&_RedeemCodeFaucet.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.GetRoleAdmin(&_RedeemCodeFaucet.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RedeemCodeFaucet.Contract.HasRole(&_RedeemCodeFaucet.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RedeemCodeFaucet.Contract.HasRole(&_RedeemCodeFaucet.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) Owner() (common.Address, error) {
	return _RedeemCodeFaucet.Contract.Owner(&_RedeemCodeFaucet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) Owner() (common.Address, error) {
	return _RedeemCodeFaucet.Contract.Owner(&_RedeemCodeFaucet.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) ProxiableUUID() ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.ProxiableUUID(&_RedeemCodeFaucet.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) ProxiableUUID() ([32]byte, error) {
	return _RedeemCodeFaucet.Contract.ProxiableUUID(&_RedeemCodeFaucet.CallOpts)
}

// Redeemed is a free data retrieval call binding the contract method 0x3d8b9e07.
//
// Solidity: function redeemed(bytes32 codeHash) view returns(uint48 time)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) Redeemed(opts *bind.CallOpts, codeHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "redeemed", codeHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Redeemed is a free data retrieval call binding the contract method 0x3d8b9e07.
//
// Solidity: function redeemed(bytes32 codeHash) view returns(uint48 time)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) Redeemed(codeHash [32]byte) (*big.Int, error) {
	return _RedeemCodeFaucet.Contract.Redeemed(&_RedeemCodeFaucet.CallOpts, codeHash)
}

// Redeemed is a free data retrieval call binding the contract method 0x3d8b9e07.
//
// Solidity: function redeemed(bytes32 codeHash) view returns(uint48 time)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) Redeemed(codeHash [32]byte) (*big.Int, error) {
	return _RedeemCodeFaucet.Contract.Redeemed(&_RedeemCodeFaucet.CallOpts, codeHash)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RedeemCodeFaucet.Contract.SupportsInterface(&_RedeemCodeFaucet.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RedeemCodeFaucet.Contract.SupportsInterface(&_RedeemCodeFaucet.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RedeemCodeFaucet.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) Token() (common.Address, error) {
	return _RedeemCodeFaucet.Contract.Token(&_RedeemCodeFaucet.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_RedeemCodeFaucet *RedeemCodeFaucetCallerSession) Token() (common.Address, error) {
	return _RedeemCodeFaucet.Contract.Token(&_RedeemCodeFaucet.CallOpts)
}

// AddRedeemCode is a paid mutator transaction binding the contract method 0xe9395096.
//
// Solidity: function addRedeemCode(bytes32 codeHash, uint256 redeemableAmount, uint48 validUntil) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) AddRedeemCode(opts *bind.TransactOpts, codeHash [32]byte, redeemableAmount *big.Int, validUntil *big.Int) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "addRedeemCode", codeHash, redeemableAmount, validUntil)
}

// AddRedeemCode is a paid mutator transaction binding the contract method 0xe9395096.
//
// Solidity: function addRedeemCode(bytes32 codeHash, uint256 redeemableAmount, uint48 validUntil) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) AddRedeemCode(codeHash [32]byte, redeemableAmount *big.Int, validUntil *big.Int) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.AddRedeemCode(&_RedeemCodeFaucet.TransactOpts, codeHash, redeemableAmount, validUntil)
}

// AddRedeemCode is a paid mutator transaction binding the contract method 0xe9395096.
//
// Solidity: function addRedeemCode(bytes32 codeHash, uint256 redeemableAmount, uint48 validUntil) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) AddRedeemCode(codeHash [32]byte, redeemableAmount *big.Int, validUntil *big.Int) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.AddRedeemCode(&_RedeemCodeFaucet.TransactOpts, codeHash, redeemableAmount, validUntil)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.GrantRole(&_RedeemCodeFaucet.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.GrantRole(&_RedeemCodeFaucet.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x163df9e8.
//
// Solidity: function initialize(address owner, address _token, uint48 _callInterval, address _codeCreator) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, _token common.Address, _callInterval *big.Int, _codeCreator common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "initialize", owner, _token, _callInterval, _codeCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0x163df9e8.
//
// Solidity: function initialize(address owner, address _token, uint48 _callInterval, address _codeCreator) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) Initialize(owner common.Address, _token common.Address, _callInterval *big.Int, _codeCreator common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.Initialize(&_RedeemCodeFaucet.TransactOpts, owner, _token, _callInterval, _codeCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0x163df9e8.
//
// Solidity: function initialize(address owner, address _token, uint48 _callInterval, address _codeCreator) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) Initialize(owner common.Address, _token common.Address, _callInterval *big.Int, _codeCreator common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.Initialize(&_RedeemCodeFaucet.TransactOpts, owner, _token, _callInterval, _codeCreator)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 code) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) Redeem(opts *bind.TransactOpts, code *big.Int) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "redeem", code)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 code) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) Redeem(code *big.Int) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.Redeem(&_RedeemCodeFaucet.TransactOpts, code)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 code) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) Redeem(code *big.Int) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.Redeem(&_RedeemCodeFaucet.TransactOpts, code)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) RenounceOwnership() (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RenounceOwnership(&_RedeemCodeFaucet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RenounceOwnership(&_RedeemCodeFaucet.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RenounceRole(&_RedeemCodeFaucet.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RenounceRole(&_RedeemCodeFaucet.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RevokeRole(&_RedeemCodeFaucet.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.RevokeRole(&_RedeemCodeFaucet.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.TransferOwnership(&_RedeemCodeFaucet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.TransferOwnership(&_RedeemCodeFaucet.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.UpgradeTo(&_RedeemCodeFaucet.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.UpgradeTo(&_RedeemCodeFaucet.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.UpgradeToAndCall(&_RedeemCodeFaucet.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.UpgradeToAndCall(&_RedeemCodeFaucet.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RedeemCodeFaucet.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetSession) Withdraw() (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.Withdraw(&_RedeemCodeFaucet.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RedeemCodeFaucet *RedeemCodeFaucetTransactorSession) Withdraw() (*types.Transaction, error) {
	return _RedeemCodeFaucet.Contract.Withdraw(&_RedeemCodeFaucet.TransactOpts)
}

// RedeemCodeFaucetAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetAdminChangedIterator struct {
	Event *RedeemCodeFaucetAdminChanged // Event containing the contract specifics and raw log

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
func (it *RedeemCodeFaucetAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetAdminChanged)
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
		it.Event = new(RedeemCodeFaucetAdminChanged)
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
func (it *RedeemCodeFaucetAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetAdminChanged represents a AdminChanged event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*RedeemCodeFaucetAdminChangedIterator, error) {

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetAdminChangedIterator{contract: _RedeemCodeFaucet.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetAdminChanged) (event.Subscription, error) {

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetAdminChanged)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseAdminChanged(log types.Log) (*RedeemCodeFaucetAdminChanged, error) {
	event := new(RedeemCodeFaucetAdminChanged)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemCodeFaucetBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetBeaconUpgradedIterator struct {
	Event *RedeemCodeFaucetBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *RedeemCodeFaucetBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetBeaconUpgraded)
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
		it.Event = new(RedeemCodeFaucetBeaconUpgraded)
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
func (it *RedeemCodeFaucetBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetBeaconUpgraded represents a BeaconUpgraded event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*RedeemCodeFaucetBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetBeaconUpgradedIterator{contract: _RedeemCodeFaucet.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetBeaconUpgraded)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseBeaconUpgraded(log types.Log) (*RedeemCodeFaucetBeaconUpgraded, error) {
	event := new(RedeemCodeFaucetBeaconUpgraded)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemCodeFaucetCodeRedeemedIterator is returned from FilterCodeRedeemed and is used to iterate over the raw logs and unpacked data for CodeRedeemed events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetCodeRedeemedIterator struct {
	Event *RedeemCodeFaucetCodeRedeemed // Event containing the contract specifics and raw log

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
func (it *RedeemCodeFaucetCodeRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetCodeRedeemed)
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
		it.Event = new(RedeemCodeFaucetCodeRedeemed)
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
func (it *RedeemCodeFaucetCodeRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetCodeRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetCodeRedeemed represents a CodeRedeemed event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetCodeRedeemed struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCodeRedeemed is a free log retrieval operation binding the contract event 0x739f8ee835b17b45c3440318f23a9a52f3b7b3194bd17cd54344eeb2aa59d908.
//
// Solidity: event CodeRedeemed(address indexed receiver, uint256 amount)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterCodeRedeemed(opts *bind.FilterOpts, receiver []common.Address) (*RedeemCodeFaucetCodeRedeemedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "CodeRedeemed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetCodeRedeemedIterator{contract: _RedeemCodeFaucet.contract, event: "CodeRedeemed", logs: logs, sub: sub}, nil
}

// WatchCodeRedeemed is a free log subscription operation binding the contract event 0x739f8ee835b17b45c3440318f23a9a52f3b7b3194bd17cd54344eeb2aa59d908.
//
// Solidity: event CodeRedeemed(address indexed receiver, uint256 amount)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchCodeRedeemed(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetCodeRedeemed, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "CodeRedeemed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetCodeRedeemed)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "CodeRedeemed", log); err != nil {
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

// ParseCodeRedeemed is a log parse operation binding the contract event 0x739f8ee835b17b45c3440318f23a9a52f3b7b3194bd17cd54344eeb2aa59d908.
//
// Solidity: event CodeRedeemed(address indexed receiver, uint256 amount)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseCodeRedeemed(log types.Log) (*RedeemCodeFaucetCodeRedeemed, error) {
	event := new(RedeemCodeFaucetCodeRedeemed)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "CodeRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemCodeFaucetInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetInitializedIterator struct {
	Event *RedeemCodeFaucetInitialized // Event containing the contract specifics and raw log

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
func (it *RedeemCodeFaucetInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetInitialized)
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
		it.Event = new(RedeemCodeFaucetInitialized)
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
func (it *RedeemCodeFaucetInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetInitialized represents a Initialized event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterInitialized(opts *bind.FilterOpts) (*RedeemCodeFaucetInitializedIterator, error) {

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetInitializedIterator{contract: _RedeemCodeFaucet.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetInitialized) (event.Subscription, error) {

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetInitialized)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseInitialized(log types.Log) (*RedeemCodeFaucetInitialized, error) {
	event := new(RedeemCodeFaucetInitialized)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemCodeFaucetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetOwnershipTransferredIterator struct {
	Event *RedeemCodeFaucetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RedeemCodeFaucetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetOwnershipTransferred)
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
		it.Event = new(RedeemCodeFaucetOwnershipTransferred)
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
func (it *RedeemCodeFaucetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetOwnershipTransferred represents a OwnershipTransferred event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RedeemCodeFaucetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetOwnershipTransferredIterator{contract: _RedeemCodeFaucet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetOwnershipTransferred)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseOwnershipTransferred(log types.Log) (*RedeemCodeFaucetOwnershipTransferred, error) {
	event := new(RedeemCodeFaucetOwnershipTransferred)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemCodeFaucetRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetRoleAdminChangedIterator struct {
	Event *RedeemCodeFaucetRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RedeemCodeFaucetRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetRoleAdminChanged)
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
		it.Event = new(RedeemCodeFaucetRoleAdminChanged)
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
func (it *RedeemCodeFaucetRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetRoleAdminChanged represents a RoleAdminChanged event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RedeemCodeFaucetRoleAdminChangedIterator, error) {

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

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetRoleAdminChangedIterator{contract: _RedeemCodeFaucet.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetRoleAdminChanged)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseRoleAdminChanged(log types.Log) (*RedeemCodeFaucetRoleAdminChanged, error) {
	event := new(RedeemCodeFaucetRoleAdminChanged)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemCodeFaucetRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetRoleGrantedIterator struct {
	Event *RedeemCodeFaucetRoleGranted // Event containing the contract specifics and raw log

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
func (it *RedeemCodeFaucetRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetRoleGranted)
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
		it.Event = new(RedeemCodeFaucetRoleGranted)
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
func (it *RedeemCodeFaucetRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetRoleGranted represents a RoleGranted event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RedeemCodeFaucetRoleGrantedIterator, error) {

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

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetRoleGrantedIterator{contract: _RedeemCodeFaucet.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetRoleGranted)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseRoleGranted(log types.Log) (*RedeemCodeFaucetRoleGranted, error) {
	event := new(RedeemCodeFaucetRoleGranted)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemCodeFaucetRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetRoleRevokedIterator struct {
	Event *RedeemCodeFaucetRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RedeemCodeFaucetRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetRoleRevoked)
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
		it.Event = new(RedeemCodeFaucetRoleRevoked)
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
func (it *RedeemCodeFaucetRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetRoleRevoked represents a RoleRevoked event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RedeemCodeFaucetRoleRevokedIterator, error) {

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

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetRoleRevokedIterator{contract: _RedeemCodeFaucet.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetRoleRevoked)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseRoleRevoked(log types.Log) (*RedeemCodeFaucetRoleRevoked, error) {
	event := new(RedeemCodeFaucetRoleRevoked)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RedeemCodeFaucetUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetUpgradedIterator struct {
	Event *RedeemCodeFaucetUpgraded // Event containing the contract specifics and raw log

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
func (it *RedeemCodeFaucetUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemCodeFaucetUpgraded)
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
		it.Event = new(RedeemCodeFaucetUpgraded)
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
func (it *RedeemCodeFaucetUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemCodeFaucetUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemCodeFaucetUpgraded represents a Upgraded event raised by the RedeemCodeFaucet contract.
type RedeemCodeFaucetUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RedeemCodeFaucetUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RedeemCodeFaucetUpgradedIterator{contract: _RedeemCodeFaucet.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RedeemCodeFaucetUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RedeemCodeFaucet.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemCodeFaucetUpgraded)
				if err := _RedeemCodeFaucet.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_RedeemCodeFaucet *RedeemCodeFaucetFilterer) ParseUpgraded(log types.Log) (*RedeemCodeFaucetUpgraded, error) {
	event := new(RedeemCodeFaucetUpgraded)
	if err := _RedeemCodeFaucet.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
