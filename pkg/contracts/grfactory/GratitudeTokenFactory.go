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
	Bin: "0x60a060405234801561000f575f80fd5b5060405161276a38038061276a83398101604081905261002e91610084565b8060405161003b90610077565b6001600160a01b039091168152602001604051809103905ff080158015610064573d5f803e3d5ffd5b506001600160a01b0316608052506100b1565b611e178061095383390190565b5f60208284031215610094575f80fd5b81516001600160a01b03811681146100aa575f80fd5b9392505050565b60805161087d6100d65f395f818160770152818161011c01526101e4015261087d5ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c806301dfecf014610043578063c56ce58814610072578063d87c2ee514610099575b5f80fd5b6100566100513660046102bd565b6100ac565b6040516001600160a01b03909116815260200160405180910390f35b6100567f000000000000000000000000000000000000000000000000000000000000000081565b6100566100a73660046102bd565b6101a7565b5f806100b884846101a7565b90506001600160a01b0381163b80156100d3575090506101a1565b6040516001600160a01b038616907f632ffdb1f5dbeb99afea447aa1f5f6cc4b23d0f61f074784c8c03a4134d7c15a905f90a26040516001600160a01b038616602482015284907f00000000000000000000000000000000000000000000000000000000000000009060440160408051601f198184030181529181526020820180516001600160e01b031663189acdbd60e31b17905251610173906102b0565b61017e929190610314565b8190604051809103905ff590508015801561019b573d5f803e3d5ffd5b50925050505b92915050565b5f610279825f1b604051806020016101be906102b0565b601f1982820381018352601f9091011660408190526001600160a01b03871660248201527f00000000000000000000000000000000000000000000000000000000000000009060440160408051601f19818403018152918152602080830180516001600160e01b031663189acdbd60e31b179052905161024093929101610314565b60408051601f198184030181529082905261025e9291602001610355565b60405160208183030381529060405280519060200120610280565b9392505050565b5f6102798383305f604051836040820152846020820152828152600b8101905060ff815360559020949350505050565b6104c48061038483390190565b5f80604083850312156102ce575f80fd5b82356001600160a01b03811681146102e4575f80fd5b946020939093013593505050565b5f5b8381101561030c5781810151838201526020016102f4565b50505f910152565b60018060a01b0383168152604060208201525f82518060408401526103408160608501602087016102f2565b601f01601f1916919091016060019392505050565b5f83516103668184602088016102f2565b83519083019061037a8183602088016102f2565b0194935050505056fe60806040526040516104c43803806104c4833981016040819052610022916102d2565b61002d82825f610034565b50506103e7565b61003d8361005f565b5f825111806100495750805b1561005a57610058838361009e565b505b505050565b610068816100ca565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606100c3838360405180606001604052806027815260200161049d6027913961017d565b9392505050565b6001600160a01b0381163b61013c5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f80856001600160a01b031685604051610199919061039a565b5f60405180830381855af49150503d805f81146101d1576040519150601f19603f3d011682016040523d82523d5f602084013e6101d6565b606091505b5090925090506101e8868383876101f2565b9695505050505050565b606083156102605782515f03610259576001600160a01b0385163b6102595760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610133565b508161026a565b61026a8383610272565b949350505050565b8151156102825781518083602001fd5b8060405162461bcd60e51b815260040161013391906103b5565b634e487b7160e01b5f52604160045260245ffd5b5f5b838110156102ca5781810151838201526020016102b2565b50505f910152565b5f80604083850312156102e3575f80fd5b82516001600160a01b03811681146102f9575f80fd5b60208401519092506001600160401b0380821115610315575f80fd5b818501915085601f830112610328575f80fd5b81518181111561033a5761033a61029c565b604051601f8201601f19908116603f011681019083821181831017156103625761036261029c565b8160405282815288602084870101111561037a575f80fd5b61038b8360208301602088016102b0565b80955050505050509250929050565b5f82516103ab8184602087016102b0565b9190910192915050565b602081525f82518060208401526103d38160408501602087016102b0565b601f01601f19169190910160400192915050565b60aa806103f35f395ff3fe608060405236601057600e6013565b005b600e5b601f601b6021565b6057565b565b5f60527f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b365f80375f80365f845af43d5f803e8080156070573d5ff35b3d5ffdfea2646970667358221220970731ac93e3e3f75616b8fbf2f918d6e475eaf2beff067bf9785eb862a4015f64736f6c63430008140033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220960c2e9cfd58239a5af385c77925499d9be826e6074292d3233ce5eae51f090c64736f6c6343000814003360c06040523060805234801562000014575f80fd5b5060405162001e1738038062001e17833981016040819052620000379162000114565b6001600160a01b03811660a0526200004e62000055565b5062000143565b5f54610100900460ff1615620000c15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116101562000112575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f6020828403121562000125575f80fd5b81516001600160a01b03811681146200013c575f80fd5b9392505050565b60805160a051611c8f620001885f395f8181610bdb015261108501525f81816104f001528181610539015281816106140152818161065401526106e10152611c8f5ff3fe608060405260043610610131575f3560e01c806352d1902d116100a8578063925247251161006d578063925247251461035357806395d89b4114610372578063a457c2d714610386578063a9059cbb146103a5578063c4d66de8146103c4578063dd62ed3e146103e3575f80fd5b806352d1902d1461029657806370a08231146102aa57806379cc6790146102de5780638a0b0360146102fd5780638da5cb5b1461031c575f80fd5b8063313ce567116100f9578063313ce567146101ec5780633659cfe614610207578063395093511461022657806340c10f191461024557806342966c68146102645780634f1ef28614610283575f80fd5b806306fdde0314610135578063095ea7b31461015f57806318160ddd1461018e578063203abe34146101ac57806323b872dd146101cd575b5f80fd5b348015610140575f80fd5b50610149610402565b60405161015691906115b2565b60405180910390f35b34801561016a575f80fd5b5061017e6101793660046115ff565b610492565b6040519015158152602001610156565b348015610199575f80fd5b506035545b604051908152602001610156565b3480156101b7575f80fd5b506101cb6101c63660046116fe565b6104ab565b005b3480156101d8575f80fd5b5061017e6101e7366004611738565b6104c8565b3480156101f7575f80fd5b5060405160128152602001610156565b348015610212575f80fd5b506101cb610221366004611771565b6104e6565b348015610231575f80fd5b5061017e6102403660046115ff565b6105c9565b348015610250575f80fd5b506101cb61025f3660046115ff565b6105ea565b34801561026f575f80fd5b506101cb61027e36600461178a565b610600565b6101cb6102913660046117a1565b61060a565b3480156102a1575f80fd5b5061019e6106d5565b3480156102b5575f80fd5b5061019e6102c4366004611771565b6001600160a01b03165f9081526033602052604090205490565b3480156102e9575f80fd5b506101cb6102f83660046115ff565b610786565b348015610308575f80fd5b506101cb610317366004611841565b61079b565b348015610327575f80fd5b5060fb5461033b906001600160a01b031681565b6040516001600160a01b039091168152602001610156565b34801561035e575f80fd5b506101cb61036d366004611883565b6107ed565b34801561037d575f80fd5b506101496108d0565b348015610391575f80fd5b5061017e6103a03660046115ff565b6108df565b3480156103b0575f80fd5b5061017e6103bf3660046115ff565b610964565b3480156103cf575f80fd5b506101cb6103de366004611771565b610977565b3480156103ee575f80fd5b5061019e6103fd366004611938565b610a83565b60606036805461041190611969565b80601f016020809104026020016040519081016040528092919081815260200182805461043d90611969565b80156104885780601f1061045f57610100808354040283529160200191610488565b820191905f5260205f20905b81548152906001019060200180831161046b57829003601f168201915b5050505050905090565b5f3361049f818585610aad565b60019150505b92915050565b6104b3610bd0565b6104c581670de0b6b3a764000061079b565b50565b5f6104d1610bd0565b6104dc848484610c5f565b90505b9392505050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036105375760405162461bcd60e51b815260040161052e906119a1565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661057f5f80516020611c13833981519152546001600160a01b031690565b6001600160a01b0316146105a55760405162461bcd60e51b815260040161052e906119ed565b6105ae81610c77565b604080515f808252602082019092526104c591839190610c7f565b5f3361049f8185856105db8383610a83565b6105e59190611a4d565b610aad565b6105f2610bd0565b6105fc8282610de9565b5050565b6104c53382610ea8565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036106525760405162461bcd60e51b815260040161052e906119a1565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661069a5f80516020611c13833981519152546001600160a01b031690565b6001600160a01b0316146106c05760405162461bcd60e51b815260040161052e906119ed565b6106c982610c77565b6105fc82826001610c7f565b5f306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107745760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000606482015260840161052e565b505f80516020611c1383398151915290565b610791823383610fda565b6105fc8282610ea8565b6107a3610bd0565b5f5b82518110156107e8575f8382815181106107c1576107c1611a60565b602002602001015190506107d58184610de9565b50806107e081611a74565b9150506107a5565b505050565b6107f5610bd0565b805182511461086c5760405162461bcd60e51b815260206004820152603760248201527f726563697069656e747320616e6420616d6f756e747320617272617973206d7560448201527f73742068617665207468652073616d65206c656e677468000000000000000000606482015260840161052e565b5f5b82518110156107e8575f83828151811061088a5761088a611a60565b602002602001015190505f8383815181106108a7576108a7611a60565b602002602001015190506108bb8282610de9565b505080806108c890611a74565b91505061086e565b60606037805461041190611969565b5f33816108ec8286610a83565b90508381101561094c5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b606482015260840161052e565b6109598286868403610aad565b506001949350505050565b5f61096d610bd0565b6104df8383611052565b5f54610100900460ff161580801561099557505f54600160ff909116105b806109ae5750303b1580156109ae57505f5460ff166001145b610a115760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161052e565b5f805460ff191660011790558015610a32575f805461ff0019166101001790555b610a3b8261105f565b80156105fc575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6001600160a01b039182165f90815260346020908152604080832093909416825291909152205490565b6001600160a01b038316610b0f5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b606482015260840161052e565b6001600160a01b038216610b705760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b606482015260840161052e565b6001600160a01b038381165f8181526034602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610c11575060fb546001600160a01b031633145b610c5d5760405162461bcd60e51b815260206004820181905260248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e74604482015260640161052e565b565b5f33610c6c858285610fda565b610959858585611126565b6104c5610bd0565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610cb2576107e8836112cf565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610d0c575060408051601f3d908101601f19168201909252610d0991810190611a8c565b60015b610d6f5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b606482015260840161052e565b5f80516020611c138339815191528114610ddd5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b606482015260840161052e565b506107e883838361136a565b6001600160a01b038216610e3f5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161052e565b8060355f828254610e509190611a4d565b90915550506001600160a01b0382165f818152603360209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6001600160a01b038216610f085760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b606482015260840161052e565b6001600160a01b0382165f9081526033602052604090205481811015610f7b5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b606482015260840161052e565b6001600160a01b0383165f8181526033602090815260408083208686039055603580548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b5f610fe58484610a83565b90505f19811461104c578181101561103f5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161052e565b61104c8484848403610aad565b50505050565b5f3361049f818585611126565b60fb80546001600160a01b0319166001600160a01b0383811691821790925560405190917f000000000000000000000000000000000000000000000000000000000000000016907fe90b83f24696d649f6cdd7a0b35764256ae51f4998f6b266afb6fc6f2d50c602905f90a36111166040518060400160405280600e81526020016d23b930ba34ba3ab232aa37b5b2b760911b8152506040518060400160405280600381526020016211d51560ea1b81525061138e565b61111e6113be565b6104c56113be565b6001600160a01b03831661118a5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161052e565b6001600160a01b0382166111ec5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b606482015260840161052e565b6001600160a01b0383165f90815260336020526040902054818110156112635760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b606482015260840161052e565b6001600160a01b038085165f8181526033602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906112c29086815260200190565b60405180910390a361104c565b6001600160a01b0381163b61133c5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b606482015260840161052e565b5f80516020611c1383398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b611373836113e4565b5f8251118061137f5750805b156107e85761104c8383611423565b5f54610100900460ff166113b45760405162461bcd60e51b815260040161052e90611aa3565b6105fc8282611513565b5f54610100900460ff16610c5d5760405162461bcd60e51b815260040161052e90611aa3565b6113ed816112cf565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606001600160a01b0383163b61148b5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b606482015260840161052e565b5f80846001600160a01b0316846040516114a59190611aee565b5f60405180830381855af49150503d805f81146114dd576040519150601f19603f3d011682016040523d82523d5f602084013e6114e2565b606091505b509150915061150a8282604051806060016040528060278152602001611c3360279139611552565b95945050505050565b5f54610100900460ff166115395760405162461bcd60e51b815260040161052e90611aa3565b60366115458382611b56565b5060376107e88282611b56565b606083156115615750816104df565b6104df83838151156115765781518083602001fd5b8060405162461bcd60e51b815260040161052e91906115b2565b5f5b838110156115aa578181015183820152602001611592565b50505f910152565b602081525f82518060208401526115d0816040850160208701611590565b601f01601f19169190910160400192915050565b80356001600160a01b03811681146115fa575f80fd5b919050565b5f8060408385031215611610575f80fd5b611619836115e4565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561166457611664611627565b604052919050565b5f67ffffffffffffffff82111561168557611685611627565b5060051b60200190565b5f82601f83011261169e575f80fd5b813560206116b36116ae8361166c565b61163b565b82815260059290921b840181019181810190868411156116d1575f80fd5b8286015b848110156116f3576116e6816115e4565b83529183019183016116d5565b509695505050505050565b5f6020828403121561170e575f80fd5b813567ffffffffffffffff811115611724575f80fd5b6117308482850161168f565b949350505050565b5f805f6060848603121561174a575f80fd5b611753846115e4565b9250611761602085016115e4565b9150604084013590509250925092565b5f60208284031215611781575f80fd5b6104df826115e4565b5f6020828403121561179a575f80fd5b5035919050565b5f80604083850312156117b2575f80fd5b6117bb836115e4565b915060208084013567ffffffffffffffff808211156117d8575f80fd5b818601915086601f8301126117eb575f80fd5b8135818111156117fd576117fd611627565b61180f601f8201601f1916850161163b565b91508082528784828501011115611824575f80fd5b80848401858401375f848284010152508093505050509250929050565b5f8060408385031215611852575f80fd5b823567ffffffffffffffff811115611868575f80fd5b6118748582860161168f565b95602094909401359450505050565b5f8060408385031215611894575f80fd5b823567ffffffffffffffff808211156118ab575f80fd5b6118b78683870161168f565b93506020915081850135818111156118cd575f80fd5b85019050601f810186136118df575f80fd5b80356118ed6116ae8261166c565b81815260059190911b8201830190838101908883111561190b575f80fd5b928401925b8284101561192957833582529284019290840190611910565b80955050505050509250929050565b5f8060408385031215611949575f80fd5b611952836115e4565b9150611960602084016115e4565b90509250929050565b600181811c9082168061197d57607f821691505b60208210810361199b57634e487b7160e01b5f52602260045260245ffd5b50919050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b808201808211156104a5576104a5611a39565b634e487b7160e01b5f52603260045260245ffd5b5f60018201611a8557611a85611a39565b5060010190565b5f60208284031215611a9c575f80fd5b5051919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b5f8251611aff818460208701611590565b9190910192915050565b601f8211156107e8575f81815260208120601f850160051c81016020861015611b2f5750805b601f850160051c820191505b81811015611b4e57828155600101611b3b565b505050505050565b815167ffffffffffffffff811115611b7057611b70611627565b611b8481611b7e8454611969565b84611b09565b602080601f831160018114611bb7575f8415611ba05750858301515b5f19600386901b1c1916600185901b178555611b4e565b5f85815260208120601f198616915b82811015611be557888601518255948401946001909101908401611bc6565b5085821015611c0257878501515f19600388901b60f8161c191681555b5050505050600190811b0190555056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220b1962d272abdb504a9b85dc88e41caf02923e08faf16419a0eaefa123597d9eb64736f6c63430008140033",
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