// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voucher

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

// VoucherMetaData contains all meta data concerning the Voucher contract.
var VoucherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"},{\"internalType\":\"contractITokenEntryPoint\",\"name\":\"aTokenEntryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"VoucherInitialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"func\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dest\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"func\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"recoverOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenEntryPoint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e0346200019957601f6200215c38819003918201601f19168301916001600160401b038311848410176200019e578084926040948552833981010312620001995780516001600160a01b03918282168203620001995760200151918216820362000199573060805260a05260c05260005460ff8160081c16620001445760ff8082160362000108575b604051611fa79081620001b58239608051818181610566015281816109ca0152610adf015260a0518181816106d0015281816107a30152818161083d01528181610cb701528181610e5701528181610eeb01528181611029015281816112ee015261133f015260c0518181816102e801528181610d6d01526113dd0152f35b60ff90811916176000557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160ff8152a13862000089565b60405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001b575b361561001957600080fd5b005b60003560e01c806223de291461018a57806301ffc9a7146101855780630b699abe14610180578063150b7a021461017b5780631626ba7e1461017657806318dfb3c7146101715780633659cfe61461016c5780633a871cdd146101675780634a58db19146101625780634d44560d1461015d5780634f1ef2861461015857806352d1902d146101535780638129fc1c1461014e5780638da5cb5b14610149578063b0d691fe14610144578063b61d27f61461013f578063bb5032061461013a578063bc197c8114610135578063c399ec8814610130578063d087d2881461012b578063f23a6e6114610126578063f2fde38b146101215763fff35b720361000e57610fd7565b610faf565b610f55565b610eb8565b610e2b565b610d9c565b610d57565b610ce6565b610ca1565b610c74565b610b91565b610acc565b610978565b61080c565b610794565b61069b565b61053f565b610457565b6103d6565b61037c565b6102ac565b61023e565b6101d2565b6001600160a01b038116036101a057565b600080fd5b9181601f840112156101a0578235916001600160401b0383116101a057602083818601950101116101a057565b346101a05760c03660031901126101a0576101ee60043561018f565b6101f960243561018f565b61020460443561018f565b6001600160401b036084358181116101a0576102249036906004016101a5565b505060a4359081116101a0576100199036906004016101a5565b346101a05760203660031901126101a05760043563ffffffff60e01b81168091036101a057602090630a85bd0160e11b811490811561029b575b811561028a575b506040519015158152f35b6301ffc9a760e01b1490503861027f565b630271189760e51b81149150610278565b346101a05760203660031901126101a0576004356102c98161018f565b604051638da5cb5b60e01b81526001600160a01b0391906020816004817f000000000000000000000000000000000000000000000000000000000000000087165afa90811561037757600091610330575b506100199261032b91163314611865565b6118bc565b906020823d821161036f575b8161034960209383610905565b8101031261036c57506100199261032b91516103648161018f565b91509261031a565b80fd5b3d915061033c565b611417565b346101a05760803660031901126101a05761039860043561018f565b6103a360243561018f565b6064356001600160401b0381116101a0576103c29036906004016101a5565b5050604051630a85bd0160e11b8152602090f35b346101a05760403660031901126101a0576024356001600160401b0381116101a05761041461040b60209236906004016101a5565b9060043561157e565b6040516001600160e01b03199091168152f35b9181601f840112156101a0578235916001600160401b0383116101a0576020808501948460051b0101116101a057565b346101a05760403660031901126101a0576001600160401b036004358181116101a057610488903690600401610427565b90916024359081116101a0576104a2903690600401610427565b6104aa611335565b8083036105045760005b8181106104bd57005b6104ee6104d36104ce8387896110dd565b6110f2565b6104e86104e1848688611131565b3691610941565b906114c6565b60001981146104ff576001016104b4565b6110b1565b60405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b6044820152606490fd5b346101a05760203660031901126101a05760043561055c8161018f565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116919061059530841415611948565b6105b2600080516020611f328339815191529382855416146119a9565b6105ba6114f5565b604051906105c7826108cf565b600082527f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff16156106015750506100199150611cb4565b6020600491604094939451928380926352d1902d60e01b825286165afa6000918161065c575b506106495760405162461bcd60e51b81528061064560048201611c65565b0390fd5b610019936106579114611c07565b611d44565b61067e91925060203d8111610685575b6106768183610905565b8101906114e6565b9038610627565b503d61066c565b90816101609103126101a05790565b346101a05760603660031901126101a0576004356001600160401b0381116101a0576106cb90369060040161068c565b6044357f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633036107445761070e6107269260243590611423565b908061072a575b506040519081529081906020820190565b0390f35b600080808093338219f15061073d611496565b5038610715565b60405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e74000000006044820152606490fd5b60009103126101a057565b60008060031936011261036c577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681813b1561036c5760405163b760faf960e01b8152306004820152918290602490829034905af1801561037757610800575080f35b610809906108b7565b80f35b346101a0576000604036600319011261036c5760043561082b8161018f565b6108336114f5565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691839190833b1561089d5760449083604051958694859363040b850f60e31b855216600484015260243560248401525af1801561037757610800575080f35b8280fd5b634e487b7160e01b600052604160045260246000fd5b6001600160401b0381116108ca57604052565b6108a1565b602081019081106001600160401b038211176108ca57604052565b606081019081106001600160401b038211176108ca57604052565b90601f801991011681019081106001600160401b038211176108ca57604052565b6001600160401b0381116108ca57601f01601f191660200190565b92919261094d82610926565b9161095b6040519384610905565b8294818452818301116101a0578281602093846000960137010152565b60403660031901126101a0576004356109908161018f565b6024356001600160401b0381116101a057366023820112156101a0576109c0903690602481600401359101610941565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169291906109fa30851415611948565b610a17600080516020611f328339815191529482865416146119a9565b610a1f6114f5565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610a555750506100199150611cb4565b6020600491604094939451928380926352d1902d60e01b825286165afa60009181610aac575b50610a995760405162461bcd60e51b81528061064560048201611c65565b61001993610aa79114611c07565b611e00565b610ac591925060203d8111610685576106768183610905565b9038610a7b565b346101a05760003660031901126101a0577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163003610b2657604051600080516020611f328339815191528152602090f35b60405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608490fd5b346101a05760008060031936011261036c578054610bc660ff8260081c161580928193610c66575b8115610c46575b5061114c565b80610bd9600160ff196000541617600055565b610c2d575b610be66111af565b610bed5780f35b610bfd61ff001960005416600055565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a180f35b610c4161010061ff00196000541617600055565b610bde565b303b15915081610c58575b5038610bc0565b6001915060ff161438610c51565b600160ff8216109150610bb9565b346101a05760003660031901126101a05760005460405160109190911c6001600160a01b03168152602090f35b346101a05760003660031901126101a0576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346101a05760603660031901126101a057600435610d038161018f565b6044356001600160401b0381116101a057600091610d31610d29849336906004016101a5565b6104e1611335565b9060208251920190602435905af1610d47611496565b9015610d4f57005b602081519101fd5b346101a05760003660031901126101a0576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b346101a05760a03660031901126101a057610db860043561018f565b610dc360243561018f565b6001600160401b036044358181116101a057610de3903690600401610427565b50506064358181116101a057610dfd903690600401610427565b50506084359081116101a057610e179036906004016101a5565b505060405163bc197c8160e01b8152602090f35b346101a05760003660031901126101a0576040516370a0823160e01b81523060048201526020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa801561037757602091600091610e9b575b50604051908152f35b610eb29150823d8111610685576106768183610905565b38610e92565b346101a05760003660031901126101a057604051631aab3f0d60e11b8152306004820152600060248201526020816044817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa80156103775761072691600091610f3757506040519081529081906020820190565b610f4f915060203d8111610685576106768183610905565b38610715565b346101a05760a03660031901126101a057610f7160043561018f565b610f7c60243561018f565b6084356001600160401b0381116101a057610f9b9036906004016101a5565b505060405163f23a6e6160e01b8152602090f35b346101a05760203660031901126101a057610019600435610fcf8161018f565b61032b6114f5565b346101a05760403660031901126101a0576004356001600160401b0381116101a05761100790369060040161068c565b61101360243582611423565b1580611027575b6040519015158152602090f35b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b156101a05760408051630bd28e3b60e01b815260209490940135901c60048401526000908390602490829084905af19182156103775761072692611098575b5061101a565b806110a56110ab926108b7565b80610789565b38611092565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b91908110156110ed5760051b0190565b6110c7565b356110fc8161018f565b90565b903590601e19813603018212156101a057018035906001600160401b0382116101a0576020019181360383136101a057565b908210156110ed576111489160051b8101906110ff565b9091565b1561115357565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b60ff60005460081c161561126457600080516020611f32833981519152546001600160a01b031660008181526001602052604090205460ff1661121f576001600160a01b0316600090815260016020526040902061121490805460ff19166001179055565b61121d336112bd565b565b60405162461bcd60e51b815260206004820152601960248201527f566f75636865723a20616c7265616479206d69677261746564000000000000006044820152606490fd5b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b6000805462010000600160b01b031916601083901b62010000600160b01b03161790556001600160a01b03908116907f0000000000000000000000000000000000000000000000000000000000000000167f706f6c731664f703c9b00eb42d1b09183bc2ac9772b13f96c2c0975df783d8eb600080a3565b60018060a01b03807f0000000000000000000000000000000000000000000000000000000000000000163314908115611405575b81156113d9575b501561137857565b60405162461bcd60e51b815260206004820152603360248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e74604482015272081bdc88151bdad95b915b9d1c9e541bda5b9d606a1b6064820152608490fd5b90507f000000000000000000000000000000000000000000000000000000000000000016331438611370565b80915060005460101c16331490611369565b6040513d6000823e3d90fd5b907f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52603c60002061148661147e60018060a01b03926114786104e18560005460101c16966101408101906110ff565b90611b4a565b919091611a2a565b160361149157600090565b600190565b3d156114c1573d906114a782610926565b916114b56040519384610905565b82523d6000602084013e565b606090565b600091829182602083519301915af16114dd611496565b9015610d4f5750565b908160209103126101a0575190565b6000543360109190911c6001600160a01b0316148015611575575b1561151757565b60405162461bcd60e51b815260206004820152603060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526f081bdc881d1a194818dbdb9d1c9858dd60821b6064820152608490fd5b50333014611510565b919061158b913691610941565b90611599604183511461173e565b6115bc6115b66115a88461179e565b516001600160f81b03191690565b60f81c90565b6115ce6115c884611801565b93611855565b916fa2a8918ca85bafe22016d0b997e4df60600160ff1b0383116116e55760ff8216601b81141590816116d9575b5061168057611630600093602095604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa15610377576000516001600160a01b03166116558115156117ae565b60005460101c6001600160a01b03160361167457630b135d3f60e11b90565b6001600160e01b031990565b60405162461bcd60e51b815260206004820152603d6024820152600080516020611f5283398151915260448201527f3a20696e76616c6964207369676e6174757265202776272076616c75650000006064820152608490fd5b601c91501415386115fc565b60405162461bcd60e51b815260206004820152603d6024820152600080516020611f5283398151915260448201527f3a20696e76616c6964207369676e6174757265202773272076616c75650000006064820152608490fd5b1561174557565b60405162461bcd60e51b815260206004820152603a6024820152600080516020611f5283398151915260448201527f3a20696e76616c6964207369676e6174757265206c656e6774680000000000006064820152608490fd5b8051604010156110ed5760600190565b156117b557565b60405162461bcd60e51b81526020600482015260306024820152600080516020611f5283398151915260448201526f1d1024a72b20a624a22fa9a4a3a722a960811b6064820152608490fd5b6020815110611811576020015190565b606460405162461bcd60e51b815260206004820152602060248201527f72656164427974657333323a20696e76616c69642064617461206c656e6774686044820152fd5b6040815110611811576040015190565b1561186c57565b60405162461bcd60e51b815260206004820152602260248201527f4f776e61626c653a206e6f7420546f6b656e456e747279506f696e74206f776e60448201526132b960f11b6064820152608490fd5b6001600160a01b038116156118f4576000805462010000600160b01b03191660109290921b62010000600160b01b0316919091179055565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b1561194f57565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608490fd5b156119b057565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b6064820152608490fd5b60051115611a1457565b634e487b7160e01b600052602160045260246000fd5b611a3381611a0a565b80611a3b5750565b611a4481611a0a565b60018103611a915760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606490fd5b611a9a81611a0a565b60028103611ae75760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606490fd5b80611af3600392611a0a565b14611afa57565b60405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608490fd5b906041815114600014611b7457611148916020820151906060604084015193015160001a90611b7e565b5050600090600290565b939291906fa2a8918ca85bafe22016d0b997e4df60600160ff1b038311611bfa5791611bcb9160209360405196879094939260ff6060936080840197845216602083015260408201520152565b836000948592838052039060015afa156103775781516001600160a01b03811615611bf4579190565b50600190565b5050509050600090600390565b15611c0e57565b60405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608490fd5b60809060208152602e60208201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960408201526d6f6e206973206e6f74205555505360901b60608201520190565b803b15611ce957600080516020611f3283398151915280546001600160a01b0319166001600160a01b03909216919091179055565b60405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608490fd5b90611d4e82611cb4565b6001600160a01b0382167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a2805115801590611df8575b611d90575050565b611df59160008060405193611da4856108ea565b602785527f416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c6020860152660819985a5b195960ca1b6040860152602081519101845af4611def611496565b91611e9f565b50565b506000611d88565b90611e0a82611cb4565b6001600160a01b0382167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a2805115801590611e4b57611d90575050565b506001611d88565b15611e5a57565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b91929015611ebf5750815115611eb3575090565b6110fc903b1515611e53565b825190915015611ed25750805190602001fd5b6040519062461bcd60e51b82528160208060048301528251908160248401526000935b828510611f18575050604492506000838284010152601f80199101168101030190fd5b8481018201518686016044015293810193859350611ef556fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5369676e617475726556616c696461746f72237265636f7665725369676e6572a2646970667358221220fe9cf65327428811471273cb6d3e1a5d77100d90a7bb07520307e97d098a654564736f6c63430008140033",
}

// VoucherABI is the input ABI used to generate the binding from.
// Deprecated: Use VoucherMetaData.ABI instead.
var VoucherABI = VoucherMetaData.ABI

// VoucherBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VoucherMetaData.Bin instead.
var VoucherBin = VoucherMetaData.Bin

// DeployVoucher deploys a new Ethereum contract, binding an instance of Voucher to it.
func DeployVoucher(auth *bind.TransactOpts, backend bind.ContractBackend, anEntryPoint common.Address, aTokenEntryPoint common.Address) (common.Address, *types.Transaction, *Voucher, error) {
	parsed, err := VoucherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VoucherBin), backend, anEntryPoint, aTokenEntryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Voucher{VoucherCaller: VoucherCaller{contract: contract}, VoucherTransactor: VoucherTransactor{contract: contract}, VoucherFilterer: VoucherFilterer{contract: contract}}, nil
}

// Voucher is an auto generated Go binding around an Ethereum contract.
type Voucher struct {
	VoucherCaller     // Read-only binding to the contract
	VoucherTransactor // Write-only binding to the contract
	VoucherFilterer   // Log filterer for contract events
}

// VoucherCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoucherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoucherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoucherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoucherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoucherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoucherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoucherSession struct {
	Contract     *Voucher          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoucherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoucherCallerSession struct {
	Contract *VoucherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VoucherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoucherTransactorSession struct {
	Contract     *VoucherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VoucherRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoucherRaw struct {
	Contract *Voucher // Generic contract binding to access the raw methods on
}

// VoucherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoucherCallerRaw struct {
	Contract *VoucherCaller // Generic read-only contract binding to access the raw methods on
}

// VoucherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoucherTransactorRaw struct {
	Contract *VoucherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoucher creates a new instance of Voucher, bound to a specific deployed contract.
func NewVoucher(address common.Address, backend bind.ContractBackend) (*Voucher, error) {
	contract, err := bindVoucher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voucher{VoucherCaller: VoucherCaller{contract: contract}, VoucherTransactor: VoucherTransactor{contract: contract}, VoucherFilterer: VoucherFilterer{contract: contract}}, nil
}

// NewVoucherCaller creates a new read-only instance of Voucher, bound to a specific deployed contract.
func NewVoucherCaller(address common.Address, caller bind.ContractCaller) (*VoucherCaller, error) {
	contract, err := bindVoucher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoucherCaller{contract: contract}, nil
}

// NewVoucherTransactor creates a new write-only instance of Voucher, bound to a specific deployed contract.
func NewVoucherTransactor(address common.Address, transactor bind.ContractTransactor) (*VoucherTransactor, error) {
	contract, err := bindVoucher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoucherTransactor{contract: contract}, nil
}

// NewVoucherFilterer creates a new log filterer instance of Voucher, bound to a specific deployed contract.
func NewVoucherFilterer(address common.Address, filterer bind.ContractFilterer) (*VoucherFilterer, error) {
	contract, err := bindVoucher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoucherFilterer{contract: contract}, nil
}

// bindVoucher binds a generic wrapper to an already deployed contract.
func bindVoucher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoucherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voucher *VoucherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voucher.Contract.VoucherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voucher *VoucherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voucher.Contract.VoucherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voucher *VoucherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voucher.Contract.VoucherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voucher *VoucherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voucher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voucher *VoucherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voucher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voucher *VoucherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voucher.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Voucher *VoucherCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Voucher *VoucherSession) EntryPoint() (common.Address, error) {
	return _Voucher.Contract.EntryPoint(&_Voucher.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Voucher *VoucherCallerSession) EntryPoint() (common.Address, error) {
	return _Voucher.Contract.EntryPoint(&_Voucher.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Voucher *VoucherCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Voucher *VoucherSession) GetDeposit() (*big.Int, error) {
	return _Voucher.Contract.GetDeposit(&_Voucher.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Voucher *VoucherCallerSession) GetDeposit() (*big.Int, error) {
	return _Voucher.Contract.GetDeposit(&_Voucher.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Voucher *VoucherCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Voucher *VoucherSession) GetNonce() (*big.Int, error) {
	return _Voucher.Contract.GetNonce(&_Voucher.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Voucher *VoucherCallerSession) GetNonce() (*big.Int, error) {
	return _Voucher.Contract.GetNonce(&_Voucher.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_Voucher *VoucherCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_Voucher *VoucherSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _Voucher.Contract.IsValidSignature(&_Voucher.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_Voucher *VoucherCallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _Voucher.Contract.IsValidSignature(&_Voucher.CallOpts, _hash, _signature)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Voucher *VoucherCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Voucher *VoucherSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Voucher.Contract.OnERC1155BatchReceived(&_Voucher.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Voucher *VoucherCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Voucher.Contract.OnERC1155BatchReceived(&_Voucher.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Voucher *VoucherCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Voucher *VoucherSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Voucher.Contract.OnERC1155Received(&_Voucher.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Voucher *VoucherCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Voucher.Contract.OnERC1155Received(&_Voucher.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Voucher *VoucherCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Voucher *VoucherSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Voucher.Contract.OnERC721Received(&_Voucher.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Voucher *VoucherCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Voucher.Contract.OnERC721Received(&_Voucher.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voucher *VoucherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voucher *VoucherSession) Owner() (common.Address, error) {
	return _Voucher.Contract.Owner(&_Voucher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voucher *VoucherCallerSession) Owner() (common.Address, error) {
	return _Voucher.Contract.Owner(&_Voucher.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Voucher *VoucherCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Voucher *VoucherSession) ProxiableUUID() ([32]byte, error) {
	return _Voucher.Contract.ProxiableUUID(&_Voucher.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Voucher *VoucherCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Voucher.Contract.ProxiableUUID(&_Voucher.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Voucher *VoucherCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Voucher *VoucherSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Voucher.Contract.SupportsInterface(&_Voucher.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Voucher *VoucherCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Voucher.Contract.SupportsInterface(&_Voucher.CallOpts, interfaceId)
}

// TokenEntryPoint is a free data retrieval call binding the contract method 0xbb503206.
//
// Solidity: function tokenEntryPoint() view returns(address)
func (_Voucher *VoucherCaller) TokenEntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "tokenEntryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenEntryPoint is a free data retrieval call binding the contract method 0xbb503206.
//
// Solidity: function tokenEntryPoint() view returns(address)
func (_Voucher *VoucherSession) TokenEntryPoint() (common.Address, error) {
	return _Voucher.Contract.TokenEntryPoint(&_Voucher.CallOpts)
}

// TokenEntryPoint is a free data retrieval call binding the contract method 0xbb503206.
//
// Solidity: function tokenEntryPoint() view returns(address)
func (_Voucher *VoucherCallerSession) TokenEntryPoint() (common.Address, error) {
	return _Voucher.Contract.TokenEntryPoint(&_Voucher.CallOpts)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Voucher *VoucherCaller) TokensReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	var out []interface{}
	err := _Voucher.contract.Call(opts, &out, "tokensReceived", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Voucher *VoucherSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _Voucher.Contract.TokensReceived(&_Voucher.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Voucher *VoucherCallerSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _Voucher.Contract.TokensReceived(&_Voucher.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Voucher *VoucherTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Voucher *VoucherSession) AddDeposit() (*types.Transaction, error) {
	return _Voucher.Contract.AddDeposit(&_Voucher.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Voucher *VoucherTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _Voucher.Contract.AddDeposit(&_Voucher.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Voucher *VoucherTransactor) Execute(opts *bind.TransactOpts, dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "execute", dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Voucher *VoucherSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Voucher.Contract.Execute(&_Voucher.TransactOpts, dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Voucher *VoucherTransactorSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Voucher.Contract.Execute(&_Voucher.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x18dfb3c7.
//
// Solidity: function executeBatch(address[] dest, bytes[] func) returns()
func (_Voucher *VoucherTransactor) ExecuteBatch(opts *bind.TransactOpts, dest []common.Address, arg1 [][]byte) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "executeBatch", dest, arg1)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x18dfb3c7.
//
// Solidity: function executeBatch(address[] dest, bytes[] func) returns()
func (_Voucher *VoucherSession) ExecuteBatch(dest []common.Address, arg1 [][]byte) (*types.Transaction, error) {
	return _Voucher.Contract.ExecuteBatch(&_Voucher.TransactOpts, dest, arg1)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x18dfb3c7.
//
// Solidity: function executeBatch(address[] dest, bytes[] func) returns()
func (_Voucher *VoucherTransactorSession) ExecuteBatch(dest []common.Address, arg1 [][]byte) (*types.Transaction, error) {
	return _Voucher.Contract.ExecuteBatch(&_Voucher.TransactOpts, dest, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Voucher *VoucherTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Voucher *VoucherSession) Initialize() (*types.Transaction, error) {
	return _Voucher.Contract.Initialize(&_Voucher.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Voucher *VoucherTransactorSession) Initialize() (*types.Transaction, error) {
	return _Voucher.Contract.Initialize(&_Voucher.TransactOpts)
}

// RecoverOwnership is a paid mutator transaction binding the contract method 0x0b699abe.
//
// Solidity: function recoverOwnership(address newOwner) returns()
func (_Voucher *VoucherTransactor) RecoverOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "recoverOwnership", newOwner)
}

// RecoverOwnership is a paid mutator transaction binding the contract method 0x0b699abe.
//
// Solidity: function recoverOwnership(address newOwner) returns()
func (_Voucher *VoucherSession) RecoverOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Voucher.Contract.RecoverOwnership(&_Voucher.TransactOpts, newOwner)
}

// RecoverOwnership is a paid mutator transaction binding the contract method 0x0b699abe.
//
// Solidity: function recoverOwnership(address newOwner) returns()
func (_Voucher *VoucherTransactorSession) RecoverOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Voucher.Contract.RecoverOwnership(&_Voucher.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Voucher *VoucherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Voucher *VoucherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Voucher.Contract.TransferOwnership(&_Voucher.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Voucher *VoucherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Voucher.Contract.TransferOwnership(&_Voucher.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Voucher *VoucherTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Voucher *VoucherSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Voucher.Contract.UpgradeTo(&_Voucher.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Voucher *VoucherTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Voucher.Contract.UpgradeTo(&_Voucher.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Voucher *VoucherTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Voucher *VoucherSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Voucher.Contract.UpgradeToAndCall(&_Voucher.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Voucher *VoucherTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Voucher.Contract.UpgradeToAndCall(&_Voucher.TransactOpts, newImplementation, data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Voucher *VoucherTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Voucher *VoucherSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Voucher.Contract.ValidateUserOp(&_Voucher.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Voucher *VoucherTransactorSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Voucher.Contract.ValidateUserOp(&_Voucher.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp0 is a paid mutator transaction binding the contract method 0xfff35b72.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash) returns(bool)
func (_Voucher *VoucherTransactor) ValidateUserOp0(opts *bind.TransactOpts, userOp UserOperation, userOpHash [32]byte) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "validateUserOp0", userOp, userOpHash)
}

// ValidateUserOp0 is a paid mutator transaction binding the contract method 0xfff35b72.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash) returns(bool)
func (_Voucher *VoucherSession) ValidateUserOp0(userOp UserOperation, userOpHash [32]byte) (*types.Transaction, error) {
	return _Voucher.Contract.ValidateUserOp0(&_Voucher.TransactOpts, userOp, userOpHash)
}

// ValidateUserOp0 is a paid mutator transaction binding the contract method 0xfff35b72.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash) returns(bool)
func (_Voucher *VoucherTransactorSession) ValidateUserOp0(userOp UserOperation, userOpHash [32]byte) (*types.Transaction, error) {
	return _Voucher.Contract.ValidateUserOp0(&_Voucher.TransactOpts, userOp, userOpHash)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Voucher *VoucherTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Voucher.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Voucher *VoucherSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Voucher.Contract.WithdrawDepositTo(&_Voucher.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Voucher *VoucherTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Voucher.Contract.WithdrawDepositTo(&_Voucher.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Voucher *VoucherTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voucher.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Voucher *VoucherSession) Receive() (*types.Transaction, error) {
	return _Voucher.Contract.Receive(&_Voucher.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Voucher *VoucherTransactorSession) Receive() (*types.Transaction, error) {
	return _Voucher.Contract.Receive(&_Voucher.TransactOpts)
}

// VoucherAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Voucher contract.
type VoucherAdminChangedIterator struct {
	Event *VoucherAdminChanged // Event containing the contract specifics and raw log

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
func (it *VoucherAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoucherAdminChanged)
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
		it.Event = new(VoucherAdminChanged)
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
func (it *VoucherAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoucherAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoucherAdminChanged represents a AdminChanged event raised by the Voucher contract.
type VoucherAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Voucher *VoucherFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*VoucherAdminChangedIterator, error) {

	logs, sub, err := _Voucher.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &VoucherAdminChangedIterator{contract: _Voucher.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Voucher *VoucherFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *VoucherAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Voucher.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoucherAdminChanged)
				if err := _Voucher.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Voucher *VoucherFilterer) ParseAdminChanged(log types.Log) (*VoucherAdminChanged, error) {
	event := new(VoucherAdminChanged)
	if err := _Voucher.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoucherBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Voucher contract.
type VoucherBeaconUpgradedIterator struct {
	Event *VoucherBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *VoucherBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoucherBeaconUpgraded)
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
		it.Event = new(VoucherBeaconUpgraded)
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
func (it *VoucherBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoucherBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoucherBeaconUpgraded represents a BeaconUpgraded event raised by the Voucher contract.
type VoucherBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Voucher *VoucherFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*VoucherBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Voucher.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &VoucherBeaconUpgradedIterator{contract: _Voucher.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Voucher *VoucherFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *VoucherBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Voucher.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoucherBeaconUpgraded)
				if err := _Voucher.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Voucher *VoucherFilterer) ParseBeaconUpgraded(log types.Log) (*VoucherBeaconUpgraded, error) {
	event := new(VoucherBeaconUpgraded)
	if err := _Voucher.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoucherInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Voucher contract.
type VoucherInitializedIterator struct {
	Event *VoucherInitialized // Event containing the contract specifics and raw log

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
func (it *VoucherInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoucherInitialized)
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
		it.Event = new(VoucherInitialized)
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
func (it *VoucherInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoucherInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoucherInitialized represents a Initialized event raised by the Voucher contract.
type VoucherInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Voucher *VoucherFilterer) FilterInitialized(opts *bind.FilterOpts) (*VoucherInitializedIterator, error) {

	logs, sub, err := _Voucher.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VoucherInitializedIterator{contract: _Voucher.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Voucher *VoucherFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VoucherInitialized) (event.Subscription, error) {

	logs, sub, err := _Voucher.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoucherInitialized)
				if err := _Voucher.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Voucher *VoucherFilterer) ParseInitialized(log types.Log) (*VoucherInitialized, error) {
	event := new(VoucherInitialized)
	if err := _Voucher.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoucherUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Voucher contract.
type VoucherUpgradedIterator struct {
	Event *VoucherUpgraded // Event containing the contract specifics and raw log

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
func (it *VoucherUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoucherUpgraded)
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
		it.Event = new(VoucherUpgraded)
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
func (it *VoucherUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoucherUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoucherUpgraded represents a Upgraded event raised by the Voucher contract.
type VoucherUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Voucher *VoucherFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VoucherUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Voucher.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VoucherUpgradedIterator{contract: _Voucher.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Voucher *VoucherFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VoucherUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Voucher.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoucherUpgraded)
				if err := _Voucher.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Voucher *VoucherFilterer) ParseUpgraded(log types.Log) (*VoucherUpgraded, error) {
	event := new(VoucherUpgraded)
	if err := _Voucher.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoucherVoucherInitializedIterator is returned from FilterVoucherInitialized and is used to iterate over the raw logs and unpacked data for VoucherInitialized events raised by the Voucher contract.
type VoucherVoucherInitializedIterator struct {
	Event *VoucherVoucherInitialized // Event containing the contract specifics and raw log

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
func (it *VoucherVoucherInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoucherVoucherInitialized)
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
		it.Event = new(VoucherVoucherInitialized)
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
func (it *VoucherVoucherInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoucherVoucherInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoucherVoucherInitialized represents a VoucherInitialized event raised by the Voucher contract.
type VoucherVoucherInitialized struct {
	EntryPoint common.Address
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoucherInitialized is a free log retrieval operation binding the contract event 0x706f6c731664f703c9b00eb42d1b09183bc2ac9772b13f96c2c0975df783d8eb.
//
// Solidity: event VoucherInitialized(address indexed entryPoint, address indexed owner)
func (_Voucher *VoucherFilterer) FilterVoucherInitialized(opts *bind.FilterOpts, entryPoint []common.Address, owner []common.Address) (*VoucherVoucherInitializedIterator, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Voucher.contract.FilterLogs(opts, "VoucherInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &VoucherVoucherInitializedIterator{contract: _Voucher.contract, event: "VoucherInitialized", logs: logs, sub: sub}, nil
}

// WatchVoucherInitialized is a free log subscription operation binding the contract event 0x706f6c731664f703c9b00eb42d1b09183bc2ac9772b13f96c2c0975df783d8eb.
//
// Solidity: event VoucherInitialized(address indexed entryPoint, address indexed owner)
func (_Voucher *VoucherFilterer) WatchVoucherInitialized(opts *bind.WatchOpts, sink chan<- *VoucherVoucherInitialized, entryPoint []common.Address, owner []common.Address) (event.Subscription, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Voucher.contract.WatchLogs(opts, "VoucherInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoucherVoucherInitialized)
				if err := _Voucher.contract.UnpackLog(event, "VoucherInitialized", log); err != nil {
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

// ParseVoucherInitialized is a log parse operation binding the contract event 0x706f6c731664f703c9b00eb42d1b09183bc2ac9772b13f96c2c0975df783d8eb.
//
// Solidity: event VoucherInitialized(address indexed entryPoint, address indexed owner)
func (_Voucher *VoucherFilterer) ParseVoucherInitialized(log types.Log) (*VoucherVoucherInitialized, error) {
	event := new(VoucherVoucherInitialized)
	if err := _Voucher.contract.UnpackLog(event, "VoucherInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
