// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package profactory

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

// ProfactoryMetaData contains all meta data concerning the Profactory contract.
var ProfactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ProfileCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"createGratitudeToken\",\"outputs\":[{\"internalType\":\"contractProfile\",\"name\":\"ret\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getGratitudeTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gratitudeImplementation\",\"outputs\":[{\"internalType\":\"contractProfile\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b5060405161228638038061228683398101604081905261002e91610084565b8060405161003b90610077565b6001600160a01b039091168152602001604051809103905ff080158015610064573d5f803e3d5ffd5b506001600160a01b0316608052506100b1565b6119338061095383390190565b5f60208284031215610094575f80fd5b81516001600160a01b03811681146100aa575f80fd5b9392505050565b60805161087d6100d65f395f818160770152818161011c01526101e4015261087d5ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c806301dfecf014610043578063c56ce58814610072578063d87c2ee514610099575b5f80fd5b6100566100513660046102bd565b6100ac565b6040516001600160a01b03909116815260200160405180910390f35b6100567f000000000000000000000000000000000000000000000000000000000000000081565b6100566100a73660046102bd565b6101a7565b5f806100b884846101a7565b90506001600160a01b0381163b80156100d3575090506101a1565b6040516001600160a01b038616907ff10ec97e920cde75887622583772c637595a40b1f7777c18e51ea36cd127475c905f90a26040516001600160a01b038616602482015284907f00000000000000000000000000000000000000000000000000000000000000009060440160408051601f198184030181529181526020820180516001600160e01b031663189acdbd60e31b17905251610173906102b0565b61017e929190610314565b8190604051809103905ff590508015801561019b573d5f803e3d5ffd5b50925050505b92915050565b5f610279825f1b604051806020016101be906102b0565b601f1982820381018352601f9091011660408190526001600160a01b03871660248201527f00000000000000000000000000000000000000000000000000000000000000009060440160408051601f19818403018152918152602080830180516001600160e01b031663189acdbd60e31b179052905161024093929101610314565b60408051601f198184030181529082905261025e9291602001610355565b60405160208183030381529060405280519060200120610280565b9392505050565b5f6102798383305f604051836040820152846020820152828152600b8101905060ff815360559020949350505050565b6104c48061038483390190565b5f80604083850312156102ce575f80fd5b82356001600160a01b03811681146102e4575f80fd5b946020939093013593505050565b5f5b8381101561030c5781810151838201526020016102f4565b50505f910152565b60018060a01b0383168152604060208201525f82518060408401526103408160608501602087016102f2565b601f01601f1916919091016060019392505050565b5f83516103668184602088016102f2565b83519083019061037a8183602088016102f2565b0194935050505056fe60806040526040516104c43803806104c4833981016040819052610022916102d2565b61002d82825f610034565b50506103e7565b61003d8361005f565b5f825111806100495750805b1561005a57610058838361009e565b505b505050565b610068816100ca565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606100c3838360405180606001604052806027815260200161049d6027913961017d565b9392505050565b6001600160a01b0381163b61013c5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f80856001600160a01b031685604051610199919061039a565b5f60405180830381855af49150503d805f81146101d1576040519150601f19603f3d011682016040523d82523d5f602084013e6101d6565b606091505b5090925090506101e8868383876101f2565b9695505050505050565b606083156102605782515f03610259576001600160a01b0385163b6102595760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610133565b508161026a565b61026a8383610272565b949350505050565b8151156102825781518083602001fd5b8060405162461bcd60e51b815260040161013391906103b5565b634e487b7160e01b5f52604160045260245ffd5b5f5b838110156102ca5781810151838201526020016102b2565b50505f910152565b5f80604083850312156102e3575f80fd5b82516001600160a01b03811681146102f9575f80fd5b60208401519092506001600160401b0380821115610315575f80fd5b818501915085601f830112610328575f80fd5b81518181111561033a5761033a61029c565b604051601f8201601f19908116603f011681019083821181831017156103625761036261029c565b8160405282815288602084870101111561037a575f80fd5b61038b8360208301602088016102b0565b80955050505050509250929050565b5f82516103ab8184602087016102b0565b9190910192915050565b602081525f82518060208401526103d38160408501602087016102b0565b601f01601f19169190910160400192915050565b60aa806103f35f395ff3fe608060405236601057600e6013565b005b600e5b601f601b6021565b6057565b565b5f60527f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b365f80375f80365f845af43d5f803e8080156070573d5ff35b3d5ffdfea2646970667358221220970731ac93e3e3f75616b8fbf2f918d6e475eaf2beff067bf9785eb862a4015f64736f6c63430008140033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122047aa2e9f02e2cb52502e78453f15ca2e71faec92e424823657c8ea18cf4922a064736f6c6343000814003360a060405234801562000010575f80fd5b50604051620019333803806200193383398101604081905262000033916200018b565b6001600160a01b03811660809081526040805191820181525f606083018181528352815160208181018452828252808501919091528251908101835290815290820152805160089081906200008990826200025a565b5060208201516001820190620000a090826200025a565b5060408201516002820190620000b790826200025a565b50620000c5915050620000cc565b5062000322565b5f54610100900460ff1615620001385760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116101562000189575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f602082840312156200019c575f80fd5b81516001600160a01b0381168114620001b3575f80fd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c90821680620001e357607f821691505b6020821081036200020257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111562000255575f81815260208120601f850160051c81016020861015620002305750805b601f850160051c820191505b8181101562000251578281556001016200023c565b5050505b505050565b81516001600160401b03811115620002765762000276620001ba565b6200028e81620002878454620001ce565b8462000208565b602080601f831160018114620002c4575f8415620002ac5750858301515b5f19600386901b1c1916600185901b17855562000251565b5f85815260208120601f198616915b82811015620002f457888601518255948401946001909101908401620002d3565b50858210156200031257878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b6080516115f86200033b5f395f610e8d01526115f85ff3fe608060405234801561000f575f80fd5b50600436106100b1575f3560e01c80638da5cb5b1161006e5780638da5cb5b1461012d578063ab60636c14610158578063b080ae4f1461016f578063c4d66de814610182578063d6afc9b114610195578063f66ddcbb1461019d575f80fd5b80631105a5eb146100b55780634bd13de6146100ca5780636510c584146100dd57806375ca64fe146100f0578063881d8a40146100f85780638ccec7b614610125575b5f80fd5b6100c86100c3366004611071565b6101b2565b005b6100c86100d836600461110e565b610213565b6100c86100eb3660046111cc565b6102f1565b6100c86103de565b61010b610106366004611282565b6103f3565b60405161011c9594939291906112dc565b60405180910390f35b6100c861065b565b600754610140906001600160a01b031681565b6040516001600160a01b03909116815260200161011c565b6101606106db565b60405161011c93929190611343565b6100c861017d366004611282565b610883565b6100c8610190366004611385565b6108f0565b610160610a14565b6101a5610bd0565b60405161011c91906113a5565b6101ba610e82565b60408051606081018252848152602081018490529081018290526008806101e186826114f2565b50602082015160018201906101f690826114f2565b506040820151600282019061020b90826114f2565b505050505050565b61021b610e82565b5f6040518060a00160405280878152602001866001600160a01b031681526020018581526020018481526020018381525090508060068881548110610262576102626115ae565b5f9182526020909120825160059092020190819061028090826114f2565b5060208201516001820180546001600160a01b0319166001600160a01b03909216919091179055604082015160028201906102bb90826114f2565b50606082015160038201906102d090826114f2565b50608082015160048201906102e590826114f2565b50505050505050505050565b6102f9610e82565b6040805160a0810182528681526001600160a01b03861660208201529081018490526060810183905260808101829052600680546001810182555f91909152815182916005027ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0190819061036e90826114f2565b5060208201516001820180546001600160a01b0319166001600160a01b03909216919091179055604082015160028201906103a990826114f2565b50606082015160038201906103be90826114f2565b50608082015160048201906103d390826114f2565b505050505050505050565b6103e6610e82565b6103f160065f610f0f565b565b60068181548110610402575f80fd5b905f5260205f2090600502015f91509050805f01805461042190611475565b80601f016020809104026020016040519081016040528092919081815260200182805461044d90611475565b80156104985780601f1061046f57610100808354040283529160200191610498565b820191905f5260205f20905b81548152906001019060200180831161047b57829003601f168201915b505050600184015460028501805494956001600160a01b039092169491935091506104c290611475565b80601f01602080910402602001604051908101604052809291908181526020018280546104ee90611475565b80156105395780601f1061051057610100808354040283529160200191610539565b820191905f5260205f20905b81548152906001019060200180831161051c57829003601f168201915b50505050509080600301805461054e90611475565b80601f016020809104026020016040519081016040528092919081815260200182805461057a90611475565b80156105c55780601f1061059c576101008083540402835291602001916105c5565b820191905f5260205f20905b8154815290600101906020018083116105a857829003601f168201915b5050505050908060040180546105da90611475565b80601f016020809104026020016040519081016040528092919081815260200182805461060690611475565b80156106515780601f1061062857610100808354040283529160200191610651565b820191905f5260205f20905b81548152906001019060200180831161063457829003601f168201915b5050505050905085565b610663610e82565b604080516080810182525f606082018181528252825160208181018552828252808401919091528351908101845290815291810191909152805160089081906106ac90826114f2565b50602082015160018201906106c190826114f2565b50604082015160028201906106d690826114f2565b505050565b6008805481906106ea90611475565b80601f016020809104026020016040519081016040528092919081815260200182805461071690611475565b80156107615780601f1061073857610100808354040283529160200191610761565b820191905f5260205f20905b81548152906001019060200180831161074457829003601f168201915b50505050509080600101805461077690611475565b80601f01602080910402602001604051908101604052809291908181526020018280546107a290611475565b80156107ed5780601f106107c4576101008083540402835291602001916107ed565b820191905f5260205f20905b8154815290600101906020018083116107d057829003601f168201915b50505050509080600201805461080290611475565b80601f016020809104026020016040519081016040528092919081815260200182805461082e90611475565b80156108795780601f1061085057610100808354040283529160200191610879565b820191905f5260205f20905b81548152906001019060200180831161085c57829003601f168201915b5050505050905083565b61088b610e82565b6006818154811061089e5761089e6115ae565b5f91825260208220600590910201906108b78282610f30565b6001820180546001600160a01b03191690556108d6600283015f610f30565b6108e3600383015f610f30565b6106d6600483015f610f30565b5f54610100900460ff161580801561090e57505f54600160ff909116105b806109275750303b15801561092757505f5460ff166001145b61098f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff1916600117905580156109b0575f805461ff0019166101001790555b600780546001600160a01b0319166001600160a01b0384161790558015610a10575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b606080606060085f0160086001016008600201828054610a3390611475565b80601f0160208091040260200160405190810160405280929190818152602001828054610a5f90611475565b8015610aaa5780601f10610a8157610100808354040283529160200191610aaa565b820191905f5260205f20905b815481529060010190602001808311610a8d57829003601f168201915b50505050509250818054610abd90611475565b80601f0160208091040260200160405190810160405280929190818152602001828054610ae990611475565b8015610b345780601f10610b0b57610100808354040283529160200191610b34565b820191905f5260205f20905b815481529060010190602001808311610b1757829003601f168201915b50505050509150808054610b4790611475565b80601f0160208091040260200160405190810160405280929190818152602001828054610b7390611475565b8015610bbe5780601f10610b9557610100808354040283529160200191610bbe565b820191905f5260205f20905b815481529060010190602001808311610ba157829003601f168201915b50505050509050925092509250909192565b60606006805480602002602001604051908101604052809291908181526020015f905b82821015610e79578382905f5260205f2090600502016040518060a00160405290815f82018054610c2390611475565b80601f0160208091040260200160405190810160405280929190818152602001828054610c4f90611475565b8015610c9a5780601f10610c7157610100808354040283529160200191610c9a565b820191905f5260205f20905b815481529060010190602001808311610c7d57829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282018054604090920191610cca90611475565b80601f0160208091040260200160405190810160405280929190818152602001828054610cf690611475565b8015610d415780601f10610d1857610100808354040283529160200191610d41565b820191905f5260205f20905b815481529060010190602001808311610d2457829003601f168201915b50505050508152602001600382018054610d5a90611475565b80601f0160208091040260200160405190810160405280929190818152602001828054610d8690611475565b8015610dd15780601f10610da857610100808354040283529160200191610dd1565b820191905f5260205f20905b815481529060010190602001808311610db457829003601f168201915b50505050508152602001600482018054610dea90611475565b80601f0160208091040260200160405190810160405280929190818152602001828054610e1690611475565b8015610e615780601f10610e3857610100808354040283529160200191610e61565b820191905f5260205f20905b815481529060010190602001808311610e4457829003601f168201915b50505050508152505081526020019060010190610bf3565b50505050905090565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610ec357506007546001600160a01b031633145b6103f15760405162461bcd60e51b815260206004820181905260248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152606401610986565b5080545f8255600502905f5260205f2090810190610f2d9190610f67565b50565b508054610f3c90611475565b5f825580601f10610f4b575050565b601f0160209004905f5260205f2090810190610f2d9190610fc0565b80821115610fbc575f610f7a8282610f30565b6001820180546001600160a01b0319169055610f99600283015f610f30565b610fa6600383015f610f30565b610fb3600483015f610f30565b50600501610f67565b5090565b5b80821115610fbc575f8155600101610fc1565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610ff7575f80fd5b813567ffffffffffffffff8082111561101257611012610fd4565b604051601f8301601f19908116603f0116810190828211818310171561103a5761103a610fd4565b81604052838152866020858801011115611052575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f805f60608486031215611083575f80fd5b833567ffffffffffffffff8082111561109a575f80fd5b6110a687838801610fe8565b945060208601359150808211156110bb575f80fd5b6110c787838801610fe8565b935060408601359150808211156110dc575f80fd5b506110e986828701610fe8565b9150509250925092565b80356001600160a01b0381168114611109575f80fd5b919050565b5f805f805f8060c08789031215611123575f80fd5b86359550602087013567ffffffffffffffff80821115611141575f80fd5b61114d8a838b01610fe8565b965061115b60408a016110f3565b95506060890135915080821115611170575f80fd5b61117c8a838b01610fe8565b94506080890135915080821115611191575f80fd5b61119d8a838b01610fe8565b935060a08901359150808211156111b2575f80fd5b506111bf89828a01610fe8565b9150509295509295509295565b5f805f805f60a086880312156111e0575f80fd5b853567ffffffffffffffff808211156111f7575f80fd5b61120389838a01610fe8565b9650611211602089016110f3565b95506040880135915080821115611226575f80fd5b61123289838a01610fe8565b94506060880135915080821115611247575f80fd5b61125389838a01610fe8565b93506080880135915080821115611268575f80fd5b5061127588828901610fe8565b9150509295509295909350565b5f60208284031215611292575f80fd5b5035919050565b5f81518084525f5b818110156112bd576020818501810151868301820152016112a1565b505f602082860101526020601f19601f83011685010191505092915050565b60a081525f6112ee60a0830188611299565b6001600160a01b0387166020840152828103604084015261130f8187611299565b905082810360608401526113238186611299565b905082810360808401526113378185611299565b98975050505050505050565b606081525f6113556060830186611299565b82810360208401526113678186611299565b9050828103604084015261137b8185611299565b9695505050505050565b5f60208284031215611395575f80fd5b61139e826110f3565b9392505050565b5f6020808301818452808551808352604092508286019150828160051b8701018488015f5b8381101561146757603f19898403018552815160a081518186526113f082870182611299565b838b01516001600160a01b0316878c0152898401518782038b890152909250905061141b8282611299565b915050606080830151868303828801526114358382611299565b92505050608080830151925085820381870152506114538183611299565b9689019694505050908601906001016113ca565b509098975050505050505050565b600181811c9082168061148957607f821691505b6020821081036114a757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156106d6575f81815260208120601f850160051c810160208610156114d35750805b601f850160051c820191505b8181101561020b578281556001016114df565b815167ffffffffffffffff81111561150c5761150c610fd4565b6115208161151a8454611475565b846114ad565b602080601f831160018114611553575f841561153c5750858301515b5f19600386901b1c1916600185901b17855561020b565b5f85815260208120601f198616915b8281101561158157888601518255948401946001909101908401611562565b508582101561159e57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52603260045260245ffdfea26469706673582212206ab4d8a3d29e41620ebf14e90e29985c10f03bf6ab4ea1a9cbc931a4bc16841864736f6c63430008140033",
}

// ProfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProfactoryMetaData.ABI instead.
var ProfactoryABI = ProfactoryMetaData.ABI

// ProfactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProfactoryMetaData.Bin instead.
var ProfactoryBin = ProfactoryMetaData.Bin

// DeployProfactory deploys a new Ethereum contract, binding an instance of Profactory to it.
func DeployProfactory(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *Profactory, error) {
	parsed, err := ProfactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProfactoryBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Profactory{ProfactoryCaller: ProfactoryCaller{contract: contract}, ProfactoryTransactor: ProfactoryTransactor{contract: contract}, ProfactoryFilterer: ProfactoryFilterer{contract: contract}}, nil
}

// Profactory is an auto generated Go binding around an Ethereum contract.
type Profactory struct {
	ProfactoryCaller     // Read-only binding to the contract
	ProfactoryTransactor // Write-only binding to the contract
	ProfactoryFilterer   // Log filterer for contract events
}

// ProfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProfactorySession struct {
	Contract     *Profactory       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProfactoryCallerSession struct {
	Contract *ProfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ProfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProfactoryTransactorSession struct {
	Contract     *ProfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProfactoryRaw struct {
	Contract *Profactory // Generic contract binding to access the raw methods on
}

// ProfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProfactoryCallerRaw struct {
	Contract *ProfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ProfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProfactoryTransactorRaw struct {
	Contract *ProfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProfactory creates a new instance of Profactory, bound to a specific deployed contract.
func NewProfactory(address common.Address, backend bind.ContractBackend) (*Profactory, error) {
	contract, err := bindProfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Profactory{ProfactoryCaller: ProfactoryCaller{contract: contract}, ProfactoryTransactor: ProfactoryTransactor{contract: contract}, ProfactoryFilterer: ProfactoryFilterer{contract: contract}}, nil
}

// NewProfactoryCaller creates a new read-only instance of Profactory, bound to a specific deployed contract.
func NewProfactoryCaller(address common.Address, caller bind.ContractCaller) (*ProfactoryCaller, error) {
	contract, err := bindProfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProfactoryCaller{contract: contract}, nil
}

// NewProfactoryTransactor creates a new write-only instance of Profactory, bound to a specific deployed contract.
func NewProfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProfactoryTransactor, error) {
	contract, err := bindProfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProfactoryTransactor{contract: contract}, nil
}

// NewProfactoryFilterer creates a new log filterer instance of Profactory, bound to a specific deployed contract.
func NewProfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProfactoryFilterer, error) {
	contract, err := bindProfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProfactoryFilterer{contract: contract}, nil
}

// bindProfactory binds a generic wrapper to an already deployed contract.
func bindProfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProfactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profactory *ProfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Profactory.Contract.ProfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profactory *ProfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profactory.Contract.ProfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profactory *ProfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profactory.Contract.ProfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profactory *ProfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Profactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profactory *ProfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profactory *ProfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profactory.Contract.contract.Transact(opts, method, params...)
}

// GetGratitudeTokenAddress is a free data retrieval call binding the contract method 0xd87c2ee5.
//
// Solidity: function getGratitudeTokenAddress(address owner, uint256 salt) view returns(address)
func (_Profactory *ProfactoryCaller) GetGratitudeTokenAddress(opts *bind.CallOpts, owner common.Address, salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Profactory.contract.Call(opts, &out, "getGratitudeTokenAddress", owner, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGratitudeTokenAddress is a free data retrieval call binding the contract method 0xd87c2ee5.
//
// Solidity: function getGratitudeTokenAddress(address owner, uint256 salt) view returns(address)
func (_Profactory *ProfactorySession) GetGratitudeTokenAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _Profactory.Contract.GetGratitudeTokenAddress(&_Profactory.CallOpts, owner, salt)
}

// GetGratitudeTokenAddress is a free data retrieval call binding the contract method 0xd87c2ee5.
//
// Solidity: function getGratitudeTokenAddress(address owner, uint256 salt) view returns(address)
func (_Profactory *ProfactoryCallerSession) GetGratitudeTokenAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _Profactory.Contract.GetGratitudeTokenAddress(&_Profactory.CallOpts, owner, salt)
}

// GratitudeImplementation is a free data retrieval call binding the contract method 0xc56ce588.
//
// Solidity: function gratitudeImplementation() view returns(address)
func (_Profactory *ProfactoryCaller) GratitudeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Profactory.contract.Call(opts, &out, "gratitudeImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GratitudeImplementation is a free data retrieval call binding the contract method 0xc56ce588.
//
// Solidity: function gratitudeImplementation() view returns(address)
func (_Profactory *ProfactorySession) GratitudeImplementation() (common.Address, error) {
	return _Profactory.Contract.GratitudeImplementation(&_Profactory.CallOpts)
}

// GratitudeImplementation is a free data retrieval call binding the contract method 0xc56ce588.
//
// Solidity: function gratitudeImplementation() view returns(address)
func (_Profactory *ProfactoryCallerSession) GratitudeImplementation() (common.Address, error) {
	return _Profactory.Contract.GratitudeImplementation(&_Profactory.CallOpts)
}

// CreateGratitudeToken is a paid mutator transaction binding the contract method 0x01dfecf0.
//
// Solidity: function createGratitudeToken(address owner, uint256 salt) returns(address ret)
func (_Profactory *ProfactoryTransactor) CreateGratitudeToken(opts *bind.TransactOpts, owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _Profactory.contract.Transact(opts, "createGratitudeToken", owner, salt)
}

// CreateGratitudeToken is a paid mutator transaction binding the contract method 0x01dfecf0.
//
// Solidity: function createGratitudeToken(address owner, uint256 salt) returns(address ret)
func (_Profactory *ProfactorySession) CreateGratitudeToken(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _Profactory.Contract.CreateGratitudeToken(&_Profactory.TransactOpts, owner, salt)
}

// CreateGratitudeToken is a paid mutator transaction binding the contract method 0x01dfecf0.
//
// Solidity: function createGratitudeToken(address owner, uint256 salt) returns(address ret)
func (_Profactory *ProfactoryTransactorSession) CreateGratitudeToken(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _Profactory.Contract.CreateGratitudeToken(&_Profactory.TransactOpts, owner, salt)
}

// ProfactoryProfileCreatedIterator is returned from FilterProfileCreated and is used to iterate over the raw logs and unpacked data for ProfileCreated events raised by the Profactory contract.
type ProfactoryProfileCreatedIterator struct {
	Event *ProfactoryProfileCreated // Event containing the contract specifics and raw log

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
func (it *ProfactoryProfileCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfactoryProfileCreated)
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
		it.Event = new(ProfactoryProfileCreated)
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
func (it *ProfactoryProfileCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfactoryProfileCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfactoryProfileCreated represents a ProfileCreated event raised by the Profactory contract.
type ProfactoryProfileCreated struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProfileCreated is a free log retrieval operation binding the contract event 0xf10ec97e920cde75887622583772c637595a40b1f7777c18e51ea36cd127475c.
//
// Solidity: event ProfileCreated(address indexed owner)
func (_Profactory *ProfactoryFilterer) FilterProfileCreated(opts *bind.FilterOpts, owner []common.Address) (*ProfactoryProfileCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Profactory.contract.FilterLogs(opts, "ProfileCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ProfactoryProfileCreatedIterator{contract: _Profactory.contract, event: "ProfileCreated", logs: logs, sub: sub}, nil
}

// WatchProfileCreated is a free log subscription operation binding the contract event 0xf10ec97e920cde75887622583772c637595a40b1f7777c18e51ea36cd127475c.
//
// Solidity: event ProfileCreated(address indexed owner)
func (_Profactory *ProfactoryFilterer) WatchProfileCreated(opts *bind.WatchOpts, sink chan<- *ProfactoryProfileCreated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Profactory.contract.WatchLogs(opts, "ProfileCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfactoryProfileCreated)
				if err := _Profactory.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
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

// ParseProfileCreated is a log parse operation binding the contract event 0xf10ec97e920cde75887622583772c637595a40b1f7777c18e51ea36cd127475c.
//
// Solidity: event ProfileCreated(address indexed owner)
func (_Profactory *ProfactoryFilterer) ParseProfileCreated(log types.Log) (*ProfactoryProfileCreated, error) {
	event := new(ProfactoryProfileCreated)
	if err := _Profactory.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
