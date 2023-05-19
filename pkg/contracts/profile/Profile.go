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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"linkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"name\":\"addLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearLinks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinks\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"linkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"internalType\":\"structLinkable.LinkData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProfile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"links\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"linkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"linkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"name\":\"updateLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"name\":\"updateProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001a0438038062001a04833981016040819052620000349162000190565b6001600160a01b03811660809081526040805191820181526000606083018181528352815160208181018452828252808501919091528251908101835290815290820152805160089081906200008b908262000267565b5060208201516001820190620000a2908262000267565b5060408201516002820190620000b9908262000267565b50620000c7915050620000ce565b5062000333565b600054610100900460ff16156200013b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156200018e576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b600060208284031215620001a357600080fd5b81516001600160a01b0381168114620001bb57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620001ed57607f821691505b6020821081036200020e57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200026257600081815260208120601f850160051c810160208610156200023d5750805b601f850160051c820191505b818110156200025e5782815560010162000249565b5050505b505050565b81516001600160401b03811115620002835762000283620001c2565b6200029b81620002948454620001d8565b8462000214565b602080601f831160018114620002d35760008415620002ba5750858301515b600019600386901b1c1916600185901b1785556200025e565b600085815260208120601f198616915b828110156200030457888601518255948401946001909101908401620002e3565b5085821015620003235787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6080516116ae62000356600039600081816101a30152610f0701526116ae6000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80638da5cb5b1161008c578063b0d691fe11610066578063b0d691fe146101a1578063c4d66de8146101c7578063d6afc9b1146101da578063f66ddcbb146101e257600080fd5b80638da5cb5b1461014c578063ab60636c14610177578063b080ae4f1461018e57600080fd5b80631105a5eb146100d45780634bd13de6146100e95780636510c584146100fc57806375ca64fe1461010f578063881d8a40146101175780638ccec7b614610144575b600080fd5b6100e76100e23660046110fc565b6101f7565b005b6100e76100f73660046111a0565b610258565b6100e761010a366004611266565b610338565b6100e7610426565b61012a610125366004611324565b61043c565b60405161013b959493929190611383565b60405180910390f35b6100e76106b1565b60075461015f906001600160a01b031681565b6040516001600160a01b03909116815260200161013b565b61017f610732565b60405161013b939291906113eb565b6100e761019c366004611324565b6108e0565b7f000000000000000000000000000000000000000000000000000000000000000061015f565b6100e76101d536600461142e565b610951565b61017f610a7b565b6101ea610c3e565b60405161013b9190611450565b6101ff610efc565b604080516060810182528481526020810184905290810182905260088061022686826115a2565b506020820151600182019061023b90826115a2565b506040820151600282019061025090826115a2565b505050505050565b610260610efc565b60006040518060a00160405280878152602001866001600160a01b0316815260200185815260200184815260200183815250905080600688815481106102a8576102a8611662565b6000918252602090912082516005909202019081906102c790826115a2565b5060208201516001820180546001600160a01b0319166001600160a01b039092169190911790556040820151600282019061030290826115a2565b506060820151600382019061031790826115a2565b506080820151600482019061032c90826115a2565b50505050505050505050565b610340610efc565b6040805160a0810182528681526001600160a01b0386166020820152908101849052606081018390526080810182905260068054600181018255600091909152815182916005027ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f019081906103b690826115a2565b5060208201516001820180546001600160a01b0319166001600160a01b03909216919091179055604082015160028201906103f190826115a2565b506060820151600382019061040690826115a2565b506080820151600482019061041b90826115a2565b505050505050505050565b61042e610efc565b61043a60066000610f89565b565b6006818154811061044c57600080fd5b906000526020600020906005020160009150905080600001805461046f90611522565b80601f016020809104026020016040519081016040528092919081815260200182805461049b90611522565b80156104e85780601f106104bd576101008083540402835291602001916104e8565b820191906000526020600020905b8154815290600101906020018083116104cb57829003601f168201915b505050600184015460028501805494956001600160a01b0390921694919350915061051290611522565b80601f016020809104026020016040519081016040528092919081815260200182805461053e90611522565b801561058b5780601f106105605761010080835404028352916020019161058b565b820191906000526020600020905b81548152906001019060200180831161056e57829003601f168201915b5050505050908060030180546105a090611522565b80601f01602080910402602001604051908101604052809291908181526020018280546105cc90611522565b80156106195780601f106105ee57610100808354040283529160200191610619565b820191906000526020600020905b8154815290600101906020018083116105fc57829003601f168201915b50505050509080600401805461062e90611522565b80601f016020809104026020016040519081016040528092919081815260200182805461065a90611522565b80156106a75780601f1061067c576101008083540402835291602001916106a7565b820191906000526020600020905b81548152906001019060200180831161068a57829003601f168201915b5050505050905085565b6106b9610efc565b6040805160808101825260006060820181815282528251602081810185528282528084019190915283519081018452908152918101919091528051600890819061070390826115a2565b506020820151600182019061071890826115a2565b506040820151600282019061072d90826115a2565b505050565b60088054819061074190611522565b80601f016020809104026020016040519081016040528092919081815260200182805461076d90611522565b80156107ba5780601f1061078f576101008083540402835291602001916107ba565b820191906000526020600020905b81548152906001019060200180831161079d57829003601f168201915b5050505050908060010180546107cf90611522565b80601f01602080910402602001604051908101604052809291908181526020018280546107fb90611522565b80156108485780601f1061081d57610100808354040283529160200191610848565b820191906000526020600020905b81548152906001019060200180831161082b57829003601f168201915b50505050509080600201805461085d90611522565b80601f016020809104026020016040519081016040528092919081815260200182805461088990611522565b80156108d65780601f106108ab576101008083540402835291602001916108d6565b820191906000526020600020905b8154815290600101906020018083116108b957829003601f168201915b5050505050905083565b6108e8610efc565b600681815481106108fb576108fb611662565b600091825260208220600590910201906109158282610fad565b6001820180546001600160a01b0319169055610935600283016000610fad565b610943600383016000610fad565b61072d600483016000610fad565b600054610100900460ff16158080156109715750600054600160ff909116105b8061098b5750303b15801561098b575060005460ff166001145b6109f35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015610a16576000805461ff0019166101001790555b600780546001600160a01b0319166001600160a01b0384161790558015610a77576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6060806060600860000160086001016008600201828054610a9b90611522565b80601f0160208091040260200160405190810160405280929190818152602001828054610ac790611522565b8015610b145780601f10610ae957610100808354040283529160200191610b14565b820191906000526020600020905b815481529060010190602001808311610af757829003601f168201915b50505050509250818054610b2790611522565b80601f0160208091040260200160405190810160405280929190818152602001828054610b5390611522565b8015610ba05780601f10610b7557610100808354040283529160200191610ba0565b820191906000526020600020905b815481529060010190602001808311610b8357829003601f168201915b50505050509150808054610bb390611522565b80601f0160208091040260200160405190810160405280929190818152602001828054610bdf90611522565b8015610c2c5780601f10610c0157610100808354040283529160200191610c2c565b820191906000526020600020905b815481529060010190602001808311610c0f57829003601f168201915b50505050509050925092509250909192565b60606006805480602002602001604051908101604052809291908181526020016000905b82821015610ef357838290600052602060002090600502016040518060a0016040529081600082018054610c9590611522565b80601f0160208091040260200160405190810160405280929190818152602001828054610cc190611522565b8015610d0e5780601f10610ce357610100808354040283529160200191610d0e565b820191906000526020600020905b815481529060010190602001808311610cf157829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282018054604090920191610d3e90611522565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6a90611522565b8015610db75780601f10610d8c57610100808354040283529160200191610db7565b820191906000526020600020905b815481529060010190602001808311610d9a57829003601f168201915b50505050508152602001600382018054610dd090611522565b80601f0160208091040260200160405190810160405280929190818152602001828054610dfc90611522565b8015610e495780601f10610e1e57610100808354040283529160200191610e49565b820191906000526020600020905b815481529060010190602001808311610e2c57829003601f168201915b50505050508152602001600482018054610e6290611522565b80601f0160208091040260200160405190810160405280929190818152602001828054610e8e90611522565b8015610edb5780601f10610eb057610100808354040283529160200191610edb565b820191906000526020600020905b815481529060010190602001808311610ebe57829003601f168201915b50505050508152505081526020019060010190610c62565b50505050905090565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610f3d57506007546001600160a01b031633145b61043a5760405162461bcd60e51b815260206004820181905260248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e7460448201526064016109ea565b5080546000825560050290600052602060002090810190610faa9190610fe7565b50565b508054610fb990611522565b6000825580601f10610fc9575050565b601f016020900490600052602060002090810190610faa9190611044565b80821115611040576000610ffb8282610fad565b6001820180546001600160a01b031916905561101b600283016000610fad565b611029600383016000610fad565b611037600483016000610fad565b50600501610fe7565b5090565b5b808211156110405760008155600101611045565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261108057600080fd5b813567ffffffffffffffff8082111561109b5761109b611059565b604051601f8301601f19908116603f011681019082821181831017156110c3576110c3611059565b816040528381528660208588010111156110dc57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060006060848603121561111157600080fd5b833567ffffffffffffffff8082111561112957600080fd5b6111358783880161106f565b9450602086013591508082111561114b57600080fd5b6111578783880161106f565b9350604086013591508082111561116d57600080fd5b5061117a8682870161106f565b9150509250925092565b80356001600160a01b038116811461119b57600080fd5b919050565b60008060008060008060c087890312156111b957600080fd5b86359550602087013567ffffffffffffffff808211156111d857600080fd5b6111e48a838b0161106f565b96506111f260408a01611184565b9550606089013591508082111561120857600080fd5b6112148a838b0161106f565b9450608089013591508082111561122a57600080fd5b6112368a838b0161106f565b935060a089013591508082111561124c57600080fd5b5061125989828a0161106f565b9150509295509295509295565b600080600080600060a0868803121561127e57600080fd5b853567ffffffffffffffff8082111561129657600080fd5b6112a289838a0161106f565b96506112b060208901611184565b955060408801359150808211156112c657600080fd5b6112d289838a0161106f565b945060608801359150808211156112e857600080fd5b6112f489838a0161106f565b9350608088013591508082111561130a57600080fd5b506113178882890161106f565b9150509295509295909350565b60006020828403121561133657600080fd5b5035919050565b6000815180845260005b8181101561136357602081850181015186830182015201611347565b506000602082860101526020601f19601f83011685010191505092915050565b60a08152600061139660a083018861133d565b6001600160a01b038716602084015282810360408401526113b7818761133d565b905082810360608401526113cb818661133d565b905082810360808401526113df818561133d565b98975050505050505050565b6060815260006113fe606083018661133d565b8281036020840152611410818661133d565b90508281036040840152611424818561133d565b9695505050505050565b60006020828403121561144057600080fd5b61144982611184565b9392505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561151457603f19898403018552815160a0815181865261149d8287018261133d565b838b01516001600160a01b0316878c0152898401518782038b89015290925090506114c8828261133d565b915050606080830151868303828801526114e2838261133d565b9250505060808083015192508582038187015250611500818361133d565b968901969450505090860190600101611477565b509098975050505050505050565b600181811c9082168061153657607f821691505b60208210810361155657634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561072d57600081815260208120601f850160051c810160208610156115835750805b601f850160051c820191505b818110156102505782815560010161158f565b815167ffffffffffffffff8111156115bc576115bc611059565b6115d0816115ca8454611522565b8461155c565b602080601f83116001811461160557600084156115ed5750858301515b600019600386901b1c1916600185901b178555610250565b600085815260208120601f198616915b8281101561163457888601518255948401946001909101908401611615565b50858210156116525787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fdfea26469706673582212203c982fc82e5bc968f964fa9e31c576136a142d10f02ccc6df2e130b83135dc6164736f6c63430008140033",
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
// Solidity: function getProfile() view returns(string, string, string)
func (_Profile *ProfileCaller) GetProfile(opts *bind.CallOpts) (string, string, string, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getProfile")

	if err != nil {
		return *new(string), *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)

	return out0, out1, out2, err

}

// GetProfile is a free data retrieval call binding the contract method 0xd6afc9b1.
//
// Solidity: function getProfile() view returns(string, string, string)
func (_Profile *ProfileSession) GetProfile() (string, string, string, error) {
	return _Profile.Contract.GetProfile(&_Profile.CallOpts)
}

// GetProfile is a free data retrieval call binding the contract method 0xd6afc9b1.
//
// Solidity: function getProfile() view returns(string, string, string)
func (_Profile *ProfileCallerSession) GetProfile() (string, string, string, error) {
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
// Solidity: function profile() view returns(string name, string description, string meta)
func (_Profile *ProfileCaller) Profile(opts *bind.CallOpts) (struct {
	Name        string
	Description string
	Meta        string
}, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "profile")

	outstruct := new(struct {
		Name        string
		Description string
		Meta        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Meta = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// Profile is a free data retrieval call binding the contract method 0xab60636c.
//
// Solidity: function profile() view returns(string name, string description, string meta)
func (_Profile *ProfileSession) Profile() (struct {
	Name        string
	Description string
	Meta        string
}, error) {
	return _Profile.Contract.Profile(&_Profile.CallOpts)
}

// Profile is a free data retrieval call binding the contract method 0xab60636c.
//
// Solidity: function profile() view returns(string name, string description, string meta)
func (_Profile *ProfileCallerSession) Profile() (struct {
	Name        string
	Description string
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

// UpdateProfile is a paid mutator transaction binding the contract method 0x1105a5eb.
//
// Solidity: function updateProfile(string name, string description, string meta) returns()
func (_Profile *ProfileTransactor) UpdateProfile(opts *bind.TransactOpts, name string, description string, meta string) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "updateProfile", name, description, meta)
}

// UpdateProfile is a paid mutator transaction binding the contract method 0x1105a5eb.
//
// Solidity: function updateProfile(string name, string description, string meta) returns()
func (_Profile *ProfileSession) UpdateProfile(name string, description string, meta string) (*types.Transaction, error) {
	return _Profile.Contract.UpdateProfile(&_Profile.TransactOpts, name, description, meta)
}

// UpdateProfile is a paid mutator transaction binding the contract method 0x1105a5eb.
//
// Solidity: function updateProfile(string name, string description, string meta) returns()
func (_Profile *ProfileTransactorSession) UpdateProfile(name string, description string, meta string) (*types.Transaction, error) {
	return _Profile.Contract.UpdateProfile(&_Profile.TransactOpts, name, description, meta)
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
