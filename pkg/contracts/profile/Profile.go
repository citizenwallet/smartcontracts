// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package profile

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

// LinkableLinkData is an auto generated low-level Go binding around an user-defined struct.
type LinkableLinkData struct {
	Name        string
	Link        common.Address
	Description string
	LinkType    string
	Meta        string
}

// ProfileMetaData contains all meta data concerning the Profile contract.
var ProfileMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"linkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"name\":\"addLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearLinks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinks\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"linkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"internalType\":\"structLinkable.LinkData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProfile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"links\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"linkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"icon\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"linkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"name\":\"updateLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"icon\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"name\":\"updateProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001bd138038062001bd18339810160408190526200003491620001b6565b6001600160a01b03811660809081526040805160a0810182526000928101838152815281516020818101845284825280830191909152825180820184528481528284015282519081019092529181526060820152805160089081906200009b90826200028d565b5060208201516001820190620000b290826200028d565b5060408201516002820190620000c990826200028d565b5060608201516003820190620000e090826200028d565b50620000ee915050620000f5565b5062000359565b600054610100900460ff1615620001625760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811614620001b4576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b600060208284031215620001c957600080fd5b81516001600160a01b0381168114620001e157600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200021357607f821691505b6020821081036200023457634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200028857600081815260208120601f850160051c81016020861015620002635750805b601f850160051c820191505b8181101562000284578281556001016200026f565b5050505b505050565b81516001600160401b03811115620002a957620002a9620001e8565b620002c181620002ba8454620001fe565b846200023a565b602080601f831160018114620002f95760008415620002e05750858301515b600019600386901b1c1916600185901b17855562000284565b600085815260208120601f198616915b828110156200032a5788860151825594840194600190910190840162000309565b5085821015620003495787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6080516118556200037c600039600081816101a4015261106c01526118556000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80639483f25c1161008c578063b0d691fe11610066578063b0d691fe146101a2578063c4d66de8146101c8578063d6afc9b1146101db578063f66ddcbb146101e357600080fd5b80639483f25c14610164578063ab60636c14610177578063b080ae4f1461018f57600080fd5b80634bd13de6146100d45780636510c584146100e957806375ca64fe146100fc578063881d8a40146101045780638ccec7b6146101315780638da5cb5b14610139575b600080fd5b6100e76100e236600461127d565b6101f8565b005b6100e76100f7366004611343565b6102d8565b6100e76103c6565b610117610112366004611401565b6103dc565b604051610128959493929190611460565b60405180910390f35b6100e7610651565b60075461014c906001600160a01b031681565b6040516001600160a01b039091168152602001610128565b6100e76101723660046114c8565b6106f6565b61017f610774565b6040516101289493929190611575565b6100e761019d366004611401565b6109b0565b7f000000000000000000000000000000000000000000000000000000000000000061014c565b6100e76101d63660046115cd565b610a21565b61017f610b4b565b6101eb610da3565b60405161012891906115ef565b610200611061565b60006040518060a00160405280878152602001866001600160a01b031681526020018581526020018481526020018381525090508060068881548110610248576102486116c1565b600091825260209091208251600590920201908190610267908261175f565b5060208201516001820180546001600160a01b0319166001600160a01b03909216919091179055604082015160028201906102a2908261175f565b50606082015160038201906102b7908261175f565b50608082015160048201906102cc908261175f565b50505050505050505050565b6102e0611061565b6040805160a0810182528681526001600160a01b0386166020820152908101849052606081018390526080810182905260068054600181018255600091909152815182916005027ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f01908190610356908261175f565b5060208201516001820180546001600160a01b0319166001600160a01b0390921691909117905560408201516002820190610391908261175f565b50606082015160038201906103a6908261175f565b50608082015160048201906103bb908261175f565b505050505050505050565b6103ce611061565b6103da600660006110ee565b565b600681815481106103ec57600080fd5b906000526020600020906005020160009150905080600001805461040f906116d7565b80601f016020809104026020016040519081016040528092919081815260200182805461043b906116d7565b80156104885780601f1061045d57610100808354040283529160200191610488565b820191906000526020600020905b81548152906001019060200180831161046b57829003601f168201915b505050600184015460028501805494956001600160a01b039092169491935091506104b2906116d7565b80601f01602080910402602001604051908101604052809291908181526020018280546104de906116d7565b801561052b5780601f106105005761010080835404028352916020019161052b565b820191906000526020600020905b81548152906001019060200180831161050e57829003601f168201915b505050505090806003018054610540906116d7565b80601f016020809104026020016040519081016040528092919081815260200182805461056c906116d7565b80156105b95780601f1061058e576101008083540402835291602001916105b9565b820191906000526020600020905b81548152906001019060200180831161059c57829003601f168201915b5050505050908060040180546105ce906116d7565b80601f01602080910402602001604051908101604052809291908181526020018280546105fa906116d7565b80156106475780601f1061061c57610100808354040283529160200191610647565b820191906000526020600020905b81548152906001019060200180831161062a57829003601f168201915b5050505050905085565b610659611061565b6040805160a0810182526000608082018181528252825160208181018552828252808401919091528351808201855282815283850152835190810190935282526060810191909152805160089081906106b2908261175f565b50602082015160018201906106c7908261175f565b50604082015160028201906106dc908261175f565b50606082015160038201906106f1908261175f565b505050565b6106fe611061565b60408051608081018252858152602081018590529081018390526060810182905260088061072c878261175f565b5060208201516001820190610741908261175f565b5060408201516002820190610756908261175f565b506060820151600382019061076b908261175f565b50505050505050565b600880548190610783906116d7565b80601f01602080910402602001604051908101604052809291908181526020018280546107af906116d7565b80156107fc5780601f106107d1576101008083540402835291602001916107fc565b820191906000526020600020905b8154815290600101906020018083116107df57829003601f168201915b505050505090806001018054610811906116d7565b80601f016020809104026020016040519081016040528092919081815260200182805461083d906116d7565b801561088a5780601f1061085f5761010080835404028352916020019161088a565b820191906000526020600020905b81548152906001019060200180831161086d57829003601f168201915b50505050509080600201805461089f906116d7565b80601f01602080910402602001604051908101604052809291908181526020018280546108cb906116d7565b80156109185780601f106108ed57610100808354040283529160200191610918565b820191906000526020600020905b8154815290600101906020018083116108fb57829003601f168201915b50505050509080600301805461092d906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610959906116d7565b80156109a65780601f1061097b576101008083540402835291602001916109a6565b820191906000526020600020905b81548152906001019060200180831161098957829003601f168201915b5050505050905084565b6109b8611061565b600681815481106109cb576109cb6116c1565b600091825260208220600590910201906109e58282611112565b6001820180546001600160a01b0319169055610a05600283016000611112565b610a13600383016000611112565b6106f1600483016000611112565b600054610100900460ff1615808015610a415750600054600160ff909116105b80610a5b5750303b158015610a5b575060005460ff166001145b610ac35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015610ae6576000805461ff0019166101001790555b600780546001600160a01b0319166001600160a01b0384161790558015610b47576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6060806060806008600001600860010160086002016008600301838054610b71906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610b9d906116d7565b8015610bea5780601f10610bbf57610100808354040283529160200191610bea565b820191906000526020600020905b815481529060010190602001808311610bcd57829003601f168201915b50505050509350828054610bfd906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610c29906116d7565b8015610c765780601f10610c4b57610100808354040283529160200191610c76565b820191906000526020600020905b815481529060010190602001808311610c5957829003601f168201915b50505050509250818054610c89906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610cb5906116d7565b8015610d025780601f10610cd757610100808354040283529160200191610d02565b820191906000526020600020905b815481529060010190602001808311610ce557829003601f168201915b50505050509150808054610d15906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610d41906116d7565b8015610d8e5780601f10610d6357610100808354040283529160200191610d8e565b820191906000526020600020905b815481529060010190602001808311610d7157829003601f168201915b50505050509050935093509350935090919293565b60606006805480602002602001604051908101604052809291908181526020016000905b8282101561105857838290600052602060002090600502016040518060a0016040529081600082018054610dfa906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610e26906116d7565b8015610e735780601f10610e4857610100808354040283529160200191610e73565b820191906000526020600020905b815481529060010190602001808311610e5657829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282018054604090920191610ea3906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610ecf906116d7565b8015610f1c5780601f10610ef157610100808354040283529160200191610f1c565b820191906000526020600020905b815481529060010190602001808311610eff57829003601f168201915b50505050508152602001600382018054610f35906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610f61906116d7565b8015610fae5780601f10610f8357610100808354040283529160200191610fae565b820191906000526020600020905b815481529060010190602001808311610f9157829003601f168201915b50505050508152602001600482018054610fc7906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610ff3906116d7565b80156110405780601f1061101557610100808354040283529160200191611040565b820191906000526020600020905b81548152906001019060200180831161102357829003601f168201915b50505050508152505081526020019060010190610dc7565b50505050905090565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806110a257506007546001600160a01b031633145b6103da5760405162461bcd60e51b815260206004820181905260248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152606401610aba565b508054600082556005029060005260206000209081019061110f919061114c565b50565b50805461111e906116d7565b6000825580601f1061112e575050565b601f01602090049060005260206000209081019061110f91906111a9565b808211156111a55760006111608282611112565b6001820180546001600160a01b0319169055611180600283016000611112565b61118e600383016000611112565b61119c600483016000611112565b5060050161114c565b5090565b5b808211156111a557600081556001016111aa565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126111e557600080fd5b813567ffffffffffffffff80821115611200576112006111be565b604051601f8301601f19908116603f01168101908282118183101715611228576112286111be565b8160405283815286602085880101111561124157600080fd5b836020870160208301376000602085830101528094505050505092915050565b80356001600160a01b038116811461127857600080fd5b919050565b60008060008060008060c0878903121561129657600080fd5b86359550602087013567ffffffffffffffff808211156112b557600080fd5b6112c18a838b016111d4565b96506112cf60408a01611261565b955060608901359150808211156112e557600080fd5b6112f18a838b016111d4565b9450608089013591508082111561130757600080fd5b6113138a838b016111d4565b935060a089013591508082111561132957600080fd5b5061133689828a016111d4565b9150509295509295509295565b600080600080600060a0868803121561135b57600080fd5b853567ffffffffffffffff8082111561137357600080fd5b61137f89838a016111d4565b965061138d60208901611261565b955060408801359150808211156113a357600080fd5b6113af89838a016111d4565b945060608801359150808211156113c557600080fd5b6113d189838a016111d4565b935060808801359150808211156113e757600080fd5b506113f4888289016111d4565b9150509295509295909350565b60006020828403121561141357600080fd5b5035919050565b6000815180845260005b8181101561144057602081850181015186830182015201611424565b506000602082860101526020601f19601f83011685010191505092915050565b60a08152600061147360a083018861141a565b6001600160a01b03871660208401528281036040840152611494818761141a565b905082810360608401526114a8818661141a565b905082810360808401526114bc818561141a565b98975050505050505050565b600080600080608085870312156114de57600080fd5b843567ffffffffffffffff808211156114f657600080fd5b611502888389016111d4565b9550602087013591508082111561151857600080fd5b611524888389016111d4565b9450604087013591508082111561153a57600080fd5b611546888389016111d4565b9350606087013591508082111561155c57600080fd5b50611569878288016111d4565b91505092959194509250565b608081526000611588608083018761141a565b828103602084015261159a818761141a565b905082810360408401526115ae818661141a565b905082810360608401526115c2818561141a565b979650505050505050565b6000602082840312156115df57600080fd5b6115e882611261565b9392505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156116b357603f19898403018552815160a0815181865261163c8287018261141a565b838b01516001600160a01b0316878c0152898401518782038b8901529092509050611667828261141a565b91505060608083015186830382880152611681838261141a565b925050506080808301519250858203818701525061169f818361141a565b968901969450505090860190600101611616565b509098975050505050505050565b634e487b7160e01b600052603260045260246000fd5b600181811c908216806116eb57607f821691505b60208210810361170b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156106f157600081815260208120601f850160051c810160208610156117385750805b601f850160051c820191505b8181101561175757828155600101611744565b505050505050565b815167ffffffffffffffff811115611779576117796111be565b61178d8161178784546116d7565b84611711565b602080601f8311600181146117c257600084156117aa5750858301515b600019600386901b1c1916600185901b178555611757565b600085815260208120601f198616915b828110156117f1578886015182559484019460019091019084016117d2565b508582101561180f5787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea264697066735822122007c4820eb9bcd5928b9bafea57dc0195c2fbcfd372caddc56337958f3eefc14564736f6c63430008140033",
}

// ProfileABI is the input ABI used to generate the binding from.
// Deprecated: Use ProfileMetaData.ABI instead.
var ProfileABI = ProfileMetaData.ABI

// ProfileBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProfileMetaData.Bin instead.
var ProfileBin = ProfileMetaData.Bin

// DeployProfile deploys a new Ethereum contract, binding an instance of Profile to it.
func DeployProfile(auth *bind.TransactOpts, backend bind.ContractBackend, anEntryPoint common.Address) (common.Address, *types.Transaction, *Profile, error) {
	parsed, err := ProfileMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProfileBin), backend, anEntryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Profile{ProfileCaller: ProfileCaller{contract: contract}, ProfileTransactor: ProfileTransactor{contract: contract}, ProfileFilterer: ProfileFilterer{contract: contract}}, nil
}

// Profile is an auto generated Go binding around an Ethereum contract.
type Profile struct {
	ProfileCaller     // Read-only binding to the contract
	ProfileTransactor // Write-only binding to the contract
	ProfileFilterer   // Log filterer for contract events
}

// ProfileCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProfileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProfileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProfileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProfileSession struct {
	Contract     *Profile          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProfileCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProfileCallerSession struct {
	Contract *ProfileCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ProfileTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProfileTransactorSession struct {
	Contract     *ProfileTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ProfileRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProfileRaw struct {
	Contract *Profile // Generic contract binding to access the raw methods on
}

// ProfileCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProfileCallerRaw struct {
	Contract *ProfileCaller // Generic read-only contract binding to access the raw methods on
}

// ProfileTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProfileTransactorRaw struct {
	Contract *ProfileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProfile creates a new instance of Profile, bound to a specific deployed contract.
func NewProfile(address common.Address, backend bind.ContractBackend) (*Profile, error) {
	contract, err := bindProfile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Profile{ProfileCaller: ProfileCaller{contract: contract}, ProfileTransactor: ProfileTransactor{contract: contract}, ProfileFilterer: ProfileFilterer{contract: contract}}, nil
}

// NewProfileCaller creates a new read-only instance of Profile, bound to a specific deployed contract.
func NewProfileCaller(address common.Address, caller bind.ContractCaller) (*ProfileCaller, error) {
	contract, err := bindProfile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileCaller{contract: contract}, nil
}

// NewProfileTransactor creates a new write-only instance of Profile, bound to a specific deployed contract.
func NewProfileTransactor(address common.Address, transactor bind.ContractTransactor) (*ProfileTransactor, error) {
	contract, err := bindProfile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileTransactor{contract: contract}, nil
}

// NewProfileFilterer creates a new log filterer instance of Profile, bound to a specific deployed contract.
func NewProfileFilterer(address common.Address, filterer bind.ContractFilterer) (*ProfileFilterer, error) {
	contract, err := bindProfile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProfileFilterer{contract: contract}, nil
}

// bindProfile binds a generic wrapper to an already deployed contract.
func bindProfile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProfileMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profile *ProfileRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Profile.Contract.ProfileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profile *ProfileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profile.Contract.ProfileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profile *ProfileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profile.Contract.ProfileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profile *ProfileCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Profile.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profile *ProfileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profile.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profile *ProfileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profile.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Profile *ProfileCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Profile *ProfileSession) EntryPoint() (common.Address, error) {
	return _Profile.Contract.EntryPoint(&_Profile.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Profile *ProfileCallerSession) EntryPoint() (common.Address, error) {
	return _Profile.Contract.EntryPoint(&_Profile.CallOpts)
}

// GetLinks is a free data retrieval call binding the contract method 0xf66ddcbb.
//
// Solidity: function getLinks() view returns((string,address,string,string,string)[])
func (_Profile *ProfileCaller) GetLinks(opts *bind.CallOpts) ([]LinkableLinkData, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getLinks")

	if err != nil {
		return *new([]LinkableLinkData), err
	}

	out0 := *abi.ConvertType(out[0], new([]LinkableLinkData)).(*[]LinkableLinkData)

	return out0, err

}

// GetLinks is a free data retrieval call binding the contract method 0xf66ddcbb.
//
// Solidity: function getLinks() view returns((string,address,string,string,string)[])
func (_Profile *ProfileSession) GetLinks() ([]LinkableLinkData, error) {
	return _Profile.Contract.GetLinks(&_Profile.CallOpts)
}

// GetLinks is a free data retrieval call binding the contract method 0xf66ddcbb.
//
// Solidity: function getLinks() view returns((string,address,string,string,string)[])
func (_Profile *ProfileCallerSession) GetLinks() ([]LinkableLinkData, error) {
	return _Profile.Contract.GetLinks(&_Profile.CallOpts)
}

// GetProfile is a free data retrieval call binding the contract method 0xd6afc9b1.
//
// Solidity: function getProfile() view returns(string, string, string, string)
func (_Profile *ProfileCaller) GetProfile(opts *bind.CallOpts) (string, string, string, string, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getProfile")

	if err != nil {
		return *new(string), *new(string), *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)

	return out0, out1, out2, out3, err

}

// GetProfile is a free data retrieval call binding the contract method 0xd6afc9b1.
//
// Solidity: function getProfile() view returns(string, string, string, string)
func (_Profile *ProfileSession) GetProfile() (string, string, string, string, error) {
	return _Profile.Contract.GetProfile(&_Profile.CallOpts)
}

// GetProfile is a free data retrieval call binding the contract method 0xd6afc9b1.
//
// Solidity: function getProfile() view returns(string, string, string, string)
func (_Profile *ProfileCallerSession) GetProfile() (string, string, string, string, error) {
	return _Profile.Contract.GetProfile(&_Profile.CallOpts)
}

// Links is a free data retrieval call binding the contract method 0x881d8a40.
//
// Solidity: function links(uint256 ) view returns(string name, address link, string description, string linkType, string meta)
func (_Profile *ProfileCaller) Links(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name        string
	Link        common.Address
	Description string
	LinkType    string
	Meta        string
}, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "links", arg0)

	outstruct := new(struct {
		Name        string
		Link        common.Address
		Description string
		LinkType    string
		Meta        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Link = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Description = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.LinkType = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Meta = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// Links is a free data retrieval call binding the contract method 0x881d8a40.
//
// Solidity: function links(uint256 ) view returns(string name, address link, string description, string linkType, string meta)
func (_Profile *ProfileSession) Links(arg0 *big.Int) (struct {
	Name        string
	Link        common.Address
	Description string
	LinkType    string
	Meta        string
}, error) {
	return _Profile.Contract.Links(&_Profile.CallOpts, arg0)
}

// Links is a free data retrieval call binding the contract method 0x881d8a40.
//
// Solidity: function links(uint256 ) view returns(string name, address link, string description, string linkType, string meta)
func (_Profile *ProfileCallerSession) Links(arg0 *big.Int) (struct {
	Name        string
	Link        common.Address
	Description string
	LinkType    string
	Meta        string
}, error) {
	return _Profile.Contract.Links(&_Profile.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Profile *ProfileCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Profile *ProfileSession) Owner() (common.Address, error) {
	return _Profile.Contract.Owner(&_Profile.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Profile *ProfileCallerSession) Owner() (common.Address, error) {
	return _Profile.Contract.Owner(&_Profile.CallOpts)
}

// Profile is a free data retrieval call binding the contract method 0xab60636c.
//
// Solidity: function profile() view returns(string name, string description, string icon, string meta)
func (_Profile *ProfileCaller) Profile(opts *bind.CallOpts) (struct {
	Name        string
	Description string
	Icon        string
	Meta        string
}, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "profile")

	outstruct := new(struct {
		Name        string
		Description string
		Icon        string
		Meta        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Icon = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Meta = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Profile is a free data retrieval call binding the contract method 0xab60636c.
//
// Solidity: function profile() view returns(string name, string description, string icon, string meta)
func (_Profile *ProfileSession) Profile() (struct {
	Name        string
	Description string
	Icon        string
	Meta        string
}, error) {
	return _Profile.Contract.Profile(&_Profile.CallOpts)
}

// Profile is a free data retrieval call binding the contract method 0xab60636c.
//
// Solidity: function profile() view returns(string name, string description, string icon, string meta)
func (_Profile *ProfileCallerSession) Profile() (struct {
	Name        string
	Description string
	Icon        string
	Meta        string
}, error) {
	return _Profile.Contract.Profile(&_Profile.CallOpts)
}

// AddLink is a paid mutator transaction binding the contract method 0x6510c584.
//
// Solidity: function addLink(string name, address link, string description, string linkType, string meta) returns()
func (_Profile *ProfileTransactor) AddLink(opts *bind.TransactOpts, name string, link common.Address, description string, linkType string, meta string) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "addLink", name, link, description, linkType, meta)
}

// AddLink is a paid mutator transaction binding the contract method 0x6510c584.
//
// Solidity: function addLink(string name, address link, string description, string linkType, string meta) returns()
func (_Profile *ProfileSession) AddLink(name string, link common.Address, description string, linkType string, meta string) (*types.Transaction, error) {
	return _Profile.Contract.AddLink(&_Profile.TransactOpts, name, link, description, linkType, meta)
}

// AddLink is a paid mutator transaction binding the contract method 0x6510c584.
//
// Solidity: function addLink(string name, address link, string description, string linkType, string meta) returns()
func (_Profile *ProfileTransactorSession) AddLink(name string, link common.Address, description string, linkType string, meta string) (*types.Transaction, error) {
	return _Profile.Contract.AddLink(&_Profile.TransactOpts, name, link, description, linkType, meta)
}

// ClearLinks is a paid mutator transaction binding the contract method 0x75ca64fe.
//
// Solidity: function clearLinks() returns()
func (_Profile *ProfileTransactor) ClearLinks(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "clearLinks")
}

// ClearLinks is a paid mutator transaction binding the contract method 0x75ca64fe.
//
// Solidity: function clearLinks() returns()
func (_Profile *ProfileSession) ClearLinks() (*types.Transaction, error) {
	return _Profile.Contract.ClearLinks(&_Profile.TransactOpts)
}

// ClearLinks is a paid mutator transaction binding the contract method 0x75ca64fe.
//
// Solidity: function clearLinks() returns()
func (_Profile *ProfileTransactorSession) ClearLinks() (*types.Transaction, error) {
	return _Profile.Contract.ClearLinks(&_Profile.TransactOpts)
}

// ClearProfile is a paid mutator transaction binding the contract method 0x8ccec7b6.
//
// Solidity: function clearProfile() returns()
func (_Profile *ProfileTransactor) ClearProfile(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "clearProfile")
}

// ClearProfile is a paid mutator transaction binding the contract method 0x8ccec7b6.
//
// Solidity: function clearProfile() returns()
func (_Profile *ProfileSession) ClearProfile() (*types.Transaction, error) {
	return _Profile.Contract.ClearProfile(&_Profile.TransactOpts)
}

// ClearProfile is a paid mutator transaction binding the contract method 0x8ccec7b6.
//
// Solidity: function clearProfile() returns()
func (_Profile *ProfileTransactorSession) ClearProfile() (*types.Transaction, error) {
	return _Profile.Contract.ClearProfile(&_Profile.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_Profile *ProfileTransactor) Initialize(opts *bind.TransactOpts, anOwner common.Address) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "initialize", anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_Profile *ProfileSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _Profile.Contract.Initialize(&_Profile.TransactOpts, anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_Profile *ProfileTransactorSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _Profile.Contract.Initialize(&_Profile.TransactOpts, anOwner)
}

// RemoveLink is a paid mutator transaction binding the contract method 0xb080ae4f.
//
// Solidity: function removeLink(uint256 index) returns()
func (_Profile *ProfileTransactor) RemoveLink(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "removeLink", index)
}

// RemoveLink is a paid mutator transaction binding the contract method 0xb080ae4f.
//
// Solidity: function removeLink(uint256 index) returns()
func (_Profile *ProfileSession) RemoveLink(index *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.RemoveLink(&_Profile.TransactOpts, index)
}

// RemoveLink is a paid mutator transaction binding the contract method 0xb080ae4f.
//
// Solidity: function removeLink(uint256 index) returns()
func (_Profile *ProfileTransactorSession) RemoveLink(index *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.RemoveLink(&_Profile.TransactOpts, index)
}

// UpdateLink is a paid mutator transaction binding the contract method 0x4bd13de6.
//
// Solidity: function updateLink(uint256 index, string name, address link, string description, string linkType, string meta) returns()
func (_Profile *ProfileTransactor) UpdateLink(opts *bind.TransactOpts, index *big.Int, name string, link common.Address, description string, linkType string, meta string) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "updateLink", index, name, link, description, linkType, meta)
}

// UpdateLink is a paid mutator transaction binding the contract method 0x4bd13de6.
//
// Solidity: function updateLink(uint256 index, string name, address link, string description, string linkType, string meta) returns()
func (_Profile *ProfileSession) UpdateLink(index *big.Int, name string, link common.Address, description string, linkType string, meta string) (*types.Transaction, error) {
	return _Profile.Contract.UpdateLink(&_Profile.TransactOpts, index, name, link, description, linkType, meta)
}

// UpdateLink is a paid mutator transaction binding the contract method 0x4bd13de6.
//
// Solidity: function updateLink(uint256 index, string name, address link, string description, string linkType, string meta) returns()
func (_Profile *ProfileTransactorSession) UpdateLink(index *big.Int, name string, link common.Address, description string, linkType string, meta string) (*types.Transaction, error) {
	return _Profile.Contract.UpdateLink(&_Profile.TransactOpts, index, name, link, description, linkType, meta)
}

// UpdateProfile is a paid mutator transaction binding the contract method 0x9483f25c.
//
// Solidity: function updateProfile(string name, string description, string icon, string meta) returns()
func (_Profile *ProfileTransactor) UpdateProfile(opts *bind.TransactOpts, name string, description string, icon string, meta string) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "updateProfile", name, description, icon, meta)
}

// UpdateProfile is a paid mutator transaction binding the contract method 0x9483f25c.
//
// Solidity: function updateProfile(string name, string description, string icon, string meta) returns()
func (_Profile *ProfileSession) UpdateProfile(name string, description string, icon string, meta string) (*types.Transaction, error) {
	return _Profile.Contract.UpdateProfile(&_Profile.TransactOpts, name, description, icon, meta)
}

// UpdateProfile is a paid mutator transaction binding the contract method 0x9483f25c.
//
// Solidity: function updateProfile(string name, string description, string icon, string meta) returns()
func (_Profile *ProfileTransactorSession) UpdateProfile(name string, description string, icon string, meta string) (*types.Transaction, error) {
	return _Profile.Contract.UpdateProfile(&_Profile.TransactOpts, name, description, icon, meta)
}

// ProfileInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Profile contract.
type ProfileInitializedIterator struct {
	Event *ProfileInitialized // Event containing the contract specifics and raw log

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
func (it *ProfileInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileInitialized)
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
		it.Event = new(ProfileInitialized)
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
func (it *ProfileInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileInitialized represents a Initialized event raised by the Profile contract.
type ProfileInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Profile *ProfileFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProfileInitializedIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProfileInitializedIterator{contract: _Profile.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Profile *ProfileFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProfileInitialized) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileInitialized)
				if err := _Profile.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Profile *ProfileFilterer) ParseInitialized(log types.Log) (*ProfileInitialized, error) {
	event := new(ProfileInitialized)
	if err := _Profile.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
