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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"linkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"name\":\"addLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearLinks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinks\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"linkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"internalType\":\"structLinkable.LinkData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProfile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"links\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"linkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"linkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"name\":\"updateLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"meta\",\"type\":\"string\"}],\"name\":\"updateProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801562000010575f80fd5b50604051620019333803806200193383398101604081905262000033916200018b565b6001600160a01b03811660809081526040805191820181525f606083018181528352815160208181018452828252808501919091528251908101835290815290820152805160089081906200008990826200025a565b5060208201516001820190620000a090826200025a565b5060408201516002820190620000b790826200025a565b50620000c5915050620000cc565b5062000322565b5f54610100900460ff1615620001385760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116101562000189575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f602082840312156200019c575f80fd5b81516001600160a01b0381168114620001b3575f80fd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c90821680620001e357607f821691505b6020821081036200020257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111562000255575f81815260208120601f850160051c81016020861015620002305750805b601f850160051c820191505b8181101562000251578281556001016200023c565b5050505b505050565b81516001600160401b03811115620002765762000276620001ba565b6200028e81620002878454620001ce565b8462000208565b602080601f831160018114620002c4575f8415620002ac5750858301515b5f19600386901b1c1916600185901b17855562000251565b5f85815260208120601f198616915b82811015620002f457888601518255948401946001909101908401620002d3565b50858210156200031257878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b6080516115f86200033b5f395f610e8d01526115f85ff3fe608060405234801561000f575f80fd5b50600436106100b1575f3560e01c80638da5cb5b1161006e5780638da5cb5b1461012d578063ab60636c14610158578063b080ae4f1461016f578063c4d66de814610182578063d6afc9b114610195578063f66ddcbb1461019d575f80fd5b80631105a5eb146100b55780634bd13de6146100ca5780636510c584146100dd57806375ca64fe146100f0578063881d8a40146100f85780638ccec7b614610125575b5f80fd5b6100c86100c3366004611071565b6101b2565b005b6100c86100d836600461110e565b610213565b6100c86100eb3660046111cc565b6102f1565b6100c86103de565b61010b610106366004611282565b6103f3565b60405161011c9594939291906112dc565b60405180910390f35b6100c861065b565b600754610140906001600160a01b031681565b6040516001600160a01b03909116815260200161011c565b6101606106db565b60405161011c93929190611343565b6100c861017d366004611282565b610883565b6100c8610190366004611385565b6108f0565b610160610a14565b6101a5610bd0565b60405161011c91906113a5565b6101ba610e82565b60408051606081018252848152602081018490529081018290526008806101e186826114f2565b50602082015160018201906101f690826114f2565b506040820151600282019061020b90826114f2565b505050505050565b61021b610e82565b5f6040518060a00160405280878152602001866001600160a01b031681526020018581526020018481526020018381525090508060068881548110610262576102626115ae565b5f9182526020909120825160059092020190819061028090826114f2565b5060208201516001820180546001600160a01b0319166001600160a01b03909216919091179055604082015160028201906102bb90826114f2565b50606082015160038201906102d090826114f2565b50608082015160048201906102e590826114f2565b50505050505050505050565b6102f9610e82565b6040805160a0810182528681526001600160a01b03861660208201529081018490526060810183905260808101829052600680546001810182555f91909152815182916005027ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0190819061036e90826114f2565b5060208201516001820180546001600160a01b0319166001600160a01b03909216919091179055604082015160028201906103a990826114f2565b50606082015160038201906103be90826114f2565b50608082015160048201906103d390826114f2565b505050505050505050565b6103e6610e82565b6103f160065f610f0f565b565b60068181548110610402575f80fd5b905f5260205f2090600502015f91509050805f01805461042190611475565b80601f016020809104026020016040519081016040528092919081815260200182805461044d90611475565b80156104985780601f1061046f57610100808354040283529160200191610498565b820191905f5260205f20905b81548152906001019060200180831161047b57829003601f168201915b505050600184015460028501805494956001600160a01b039092169491935091506104c290611475565b80601f01602080910402602001604051908101604052809291908181526020018280546104ee90611475565b80156105395780601f1061051057610100808354040283529160200191610539565b820191905f5260205f20905b81548152906001019060200180831161051c57829003601f168201915b50505050509080600301805461054e90611475565b80601f016020809104026020016040519081016040528092919081815260200182805461057a90611475565b80156105c55780601f1061059c576101008083540402835291602001916105c5565b820191905f5260205f20905b8154815290600101906020018083116105a857829003601f168201915b5050505050908060040180546105da90611475565b80601f016020809104026020016040519081016040528092919081815260200182805461060690611475565b80156106515780601f1061062857610100808354040283529160200191610651565b820191905f5260205f20905b81548152906001019060200180831161063457829003601f168201915b5050505050905085565b610663610e82565b604080516080810182525f606082018181528252825160208181018552828252808401919091528351908101845290815291810191909152805160089081906106ac90826114f2565b50602082015160018201906106c190826114f2565b50604082015160028201906106d690826114f2565b505050565b6008805481906106ea90611475565b80601f016020809104026020016040519081016040528092919081815260200182805461071690611475565b80156107615780601f1061073857610100808354040283529160200191610761565b820191905f5260205f20905b81548152906001019060200180831161074457829003601f168201915b50505050509080600101805461077690611475565b80601f01602080910402602001604051908101604052809291908181526020018280546107a290611475565b80156107ed5780601f106107c4576101008083540402835291602001916107ed565b820191905f5260205f20905b8154815290600101906020018083116107d057829003601f168201915b50505050509080600201805461080290611475565b80601f016020809104026020016040519081016040528092919081815260200182805461082e90611475565b80156108795780601f1061085057610100808354040283529160200191610879565b820191905f5260205f20905b81548152906001019060200180831161085c57829003601f168201915b5050505050905083565b61088b610e82565b6006818154811061089e5761089e6115ae565b5f91825260208220600590910201906108b78282610f30565b6001820180546001600160a01b03191690556108d6600283015f610f30565b6108e3600383015f610f30565b6106d6600483015f610f30565b5f54610100900460ff161580801561090e57505f54600160ff909116105b806109275750303b15801561092757505f5460ff166001145b61098f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff1916600117905580156109b0575f805461ff0019166101001790555b600780546001600160a01b0319166001600160a01b0384161790558015610a10575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b606080606060085f0160086001016008600201828054610a3390611475565b80601f0160208091040260200160405190810160405280929190818152602001828054610a5f90611475565b8015610aaa5780601f10610a8157610100808354040283529160200191610aaa565b820191905f5260205f20905b815481529060010190602001808311610a8d57829003601f168201915b50505050509250818054610abd90611475565b80601f0160208091040260200160405190810160405280929190818152602001828054610ae990611475565b8015610b345780601f10610b0b57610100808354040283529160200191610b34565b820191905f5260205f20905b815481529060010190602001808311610b1757829003601f168201915b50505050509150808054610b4790611475565b80601f0160208091040260200160405190810160405280929190818152602001828054610b7390611475565b8015610bbe5780601f10610b9557610100808354040283529160200191610bbe565b820191905f5260205f20905b815481529060010190602001808311610ba157829003601f168201915b50505050509050925092509250909192565b60606006805480602002602001604051908101604052809291908181526020015f905b82821015610e79578382905f5260205f2090600502016040518060a00160405290815f82018054610c2390611475565b80601f0160208091040260200160405190810160405280929190818152602001828054610c4f90611475565b8015610c9a5780601f10610c7157610100808354040283529160200191610c9a565b820191905f5260205f20905b815481529060010190602001808311610c7d57829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282018054604090920191610cca90611475565b80601f0160208091040260200160405190810160405280929190818152602001828054610cf690611475565b8015610d415780601f10610d1857610100808354040283529160200191610d41565b820191905f5260205f20905b815481529060010190602001808311610d2457829003601f168201915b50505050508152602001600382018054610d5a90611475565b80601f0160208091040260200160405190810160405280929190818152602001828054610d8690611475565b8015610dd15780601f10610da857610100808354040283529160200191610dd1565b820191905f5260205f20905b815481529060010190602001808311610db457829003601f168201915b50505050508152602001600482018054610dea90611475565b80601f0160208091040260200160405190810160405280929190818152602001828054610e1690611475565b8015610e615780601f10610e3857610100808354040283529160200191610e61565b820191905f5260205f20905b815481529060010190602001808311610e4457829003601f168201915b50505050508152505081526020019060010190610bf3565b50505050905090565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610ec357506007546001600160a01b031633145b6103f15760405162461bcd60e51b815260206004820181905260248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152606401610986565b5080545f8255600502905f5260205f2090810190610f2d9190610f67565b50565b508054610f3c90611475565b5f825580601f10610f4b575050565b601f0160209004905f5260205f2090810190610f2d9190610fc0565b80821115610fbc575f610f7a8282610f30565b6001820180546001600160a01b0319169055610f99600283015f610f30565b610fa6600383015f610f30565b610fb3600483015f610f30565b50600501610f67565b5090565b5b80821115610fbc575f8155600101610fc1565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610ff7575f80fd5b813567ffffffffffffffff8082111561101257611012610fd4565b604051601f8301601f19908116603f0116810190828211818310171561103a5761103a610fd4565b81604052838152866020858801011115611052575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f805f60608486031215611083575f80fd5b833567ffffffffffffffff8082111561109a575f80fd5b6110a687838801610fe8565b945060208601359150808211156110bb575f80fd5b6110c787838801610fe8565b935060408601359150808211156110dc575f80fd5b506110e986828701610fe8565b9150509250925092565b80356001600160a01b0381168114611109575f80fd5b919050565b5f805f805f8060c08789031215611123575f80fd5b86359550602087013567ffffffffffffffff80821115611141575f80fd5b61114d8a838b01610fe8565b965061115b60408a016110f3565b95506060890135915080821115611170575f80fd5b61117c8a838b01610fe8565b94506080890135915080821115611191575f80fd5b61119d8a838b01610fe8565b935060a08901359150808211156111b2575f80fd5b506111bf89828a01610fe8565b9150509295509295509295565b5f805f805f60a086880312156111e0575f80fd5b853567ffffffffffffffff808211156111f7575f80fd5b61120389838a01610fe8565b9650611211602089016110f3565b95506040880135915080821115611226575f80fd5b61123289838a01610fe8565b94506060880135915080821115611247575f80fd5b61125389838a01610fe8565b93506080880135915080821115611268575f80fd5b5061127588828901610fe8565b9150509295509295909350565b5f60208284031215611292575f80fd5b5035919050565b5f81518084525f5b818110156112bd576020818501810151868301820152016112a1565b505f602082860101526020601f19601f83011685010191505092915050565b60a081525f6112ee60a0830188611299565b6001600160a01b0387166020840152828103604084015261130f8187611299565b905082810360608401526113238186611299565b905082810360808401526113378185611299565b98975050505050505050565b606081525f6113556060830186611299565b82810360208401526113678186611299565b9050828103604084015261137b8185611299565b9695505050505050565b5f60208284031215611395575f80fd5b61139e826110f3565b9392505050565b5f6020808301818452808551808352604092508286019150828160051b8701018488015f5b8381101561146757603f19898403018552815160a081518186526113f082870182611299565b838b01516001600160a01b0316878c0152898401518782038b890152909250905061141b8282611299565b915050606080830151868303828801526114358382611299565b92505050608080830151925085820381870152506114538183611299565b9689019694505050908601906001016113ca565b509098975050505050505050565b600181811c9082168061148957607f821691505b6020821081036114a757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156106d6575f81815260208120601f850160051c810160208610156114d35750805b601f850160051c820191505b8181101561020b578281556001016114df565b815167ffffffffffffffff81111561150c5761150c610fd4565b6115208161151a8454611475565b846114ad565b602080601f831160018114611553575f841561153c5750858301515b5f19600386901b1c1916600185901b17855561020b565b5f85815260208120601f198616915b8281101561158157888601518255948401946001909101908401611562565b508582101561159e57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52603260045260245ffdfea26469706673582212206ab4d8a3d29e41620ebf14e90e29985c10f03bf6ab4ea1a9cbc931a4bc16841864736f6c63430008140033",
}

// ProfileABI is the input ABI used to generate the binding from.
// Deprecated: Use ProfileMetaData.ABI instead.
var ProfileABI = ProfileMetaData.ABI

// ProfileBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProfileMetaData.Bin instead.
var ProfileBin = ProfileMetaData.Bin

// DeployProfile deploys a new Ethereum contract, binding an instance of Profile to it.
func DeployProfile(auth *bind.TransactOpts, backend bind.ContractBackend, entryPoint common.Address) (common.Address, *types.Transaction, *Profile, error) {
	parsed, err := ProfileMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProfileBin), backend, entryPoint)
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
