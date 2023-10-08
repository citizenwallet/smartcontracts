// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package account

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

// AccountMetaData contains all meta data concerning the Account contract.
var AccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SimpleAccountInitialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"func\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dest\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"func\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200198938038062001989833981016040819052620000349162000113565b6001600160a01b0381166080526200004b62000052565b5062000145565b600054610100900460ff1615620000bf5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161462000111576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6000602082840312156200012657600080fd5b81516001600160a01b03811681146200013e57600080fd5b9392505050565b6080516117f7620001926000396000818161028601528181610608015281816106900152818161077e0152818161092901528181610bbf01528181610d350152610eb701526117f76000f3fe6080604052600436106100f65760003560e01c80638da5cb5b1161008a578063c399ec8811610059578063c399ec88146102f9578063c4d66de81461030e578063d087d2881461032e578063f23a6e611461034357600080fd5b80638da5cb5b14610239578063b0d691fe14610277578063b61d27f6146102aa578063bc197c81146102ca57600080fd5b80633a871cdd116100c65780633a871cdd146101c357806347e1da2a146101f15780634a58db19146102115780634d44560d1461021957600080fd5b806223de291461010257806301ffc9a714610129578063150b7a021461015e5780631626ba7e146101a357600080fd5b366100fd57005b600080fd5b34801561010e57600080fd5b5061012761011d36600461122a565b5050505050505050565b005b34801561013557600080fd5b506101496101443660046112db565b610370565b60405190151581526020015b60405180910390f35b34801561016a57600080fd5b5061018a610179366004611305565b630a85bd0160e11b95945050505050565b6040516001600160e01b03199091168152602001610155565b3480156101af57600080fd5b5061018a6101be366004611378565b6103c2565b3480156101cf57600080fd5b506101e36101de3660046113c4565b610450565b604051908152602001610155565b3480156101fd57600080fd5b5061012761020c36600461145d565b61046f565b610127610606565b34801561022557600080fd5b506101276102343660046114f7565b610686565b34801561024557600080fd5b5060005461025f906201000090046001600160a01b031681565b6040516001600160a01b039091168152602001610155565b34801561028357600080fd5b507f000000000000000000000000000000000000000000000000000000000000000061025f565b3480156102b657600080fd5b506101276102c5366004611523565b61070f565b3480156102d657600080fd5b5061018a6102e536600461157f565b63bc197c8160e01b98975050505050505050565b34801561030557600080fd5b506101e361075e565b34801561031a57600080fd5b5061012761032936600461161d565b6107ef565b34801561033a57600080fd5b506101e3610902565b34801561034f57600080fd5b5061018a61035e36600461163a565b63f23a6e6160e01b9695505050505050565b60006001600160e01b03198216630a85bd0160e11b14806103a157506001600160e01b03198216630271189760e51b145b806103bc57506001600160e01b031982166301ffc9a760e01b145b92915050565b60008060029054906101000a90046001600160a01b03166001600160a01b03166104228585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061095892505050565b6001600160a01b03160361043e5750630b135d3f60e11b610449565b506001600160e01b03195b9392505050565b600061045a610bb4565b6104648484610c2e565b905061044982610cdd565b610477610d2a565b848114801561048d575082158061048d57508281145b6104d45760405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b60448201526064015b60405180910390fd5b600083900361058a5760005b85811015610584576105728787838181106104fd576104fd6116a4565b9050602002016020810190610512919061161d565b6000858585818110610526576105266116a4565b905060200281019061053891906116ba565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610dbd92505050565b8061057c81611717565b9150506104e0565b506105fe565b60005b858110156105fc576105ea8787838181106105aa576105aa6116a4565b90506020020160208101906105bf919061161d565b8686848181106105d1576105d16116a4565b90506020020135858585818110610526576105266116a4565b806105f481611717565b91505061058d565b505b505050505050565b7f000000000000000000000000000000000000000000000000000000000000000060405163b760faf960e01b81523060048201526001600160a01b03919091169063b760faf99034906024016000604051808303818588803b15801561066b57600080fd5b505af115801561067f573d6000803e3d6000fd5b5050505050565b61068e610e2d565b7f000000000000000000000000000000000000000000000000000000000000000060405163040b850f60e31b81526001600160a01b03848116600483015260248201849052919091169063205c287890604401600060405180830381600087803b1580156106fb57600080fd5b505af11580156105fe573d6000803e3d6000fd5b610717610d2a565b610758848484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610dbd92505050565b50505050565b6040516370a0823160e01b81523060048201526000906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a08231906024015b602060405180830381865afa1580156107c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ea9190611730565b905090565b600054610100900460ff161580801561080f5750600054600160ff909116105b806108295750303b158015610829575060005460ff166001145b61088c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016104cb565b6000805460ff1916600117905580156108af576000805461ff0019166101001790555b6108b882610e84565b80156108fe576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b604051631aab3f0d60e11b8152306004820152600060248201819052906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906335567e1a906044016107a9565b600081516041146109bf5760405162461bcd60e51b815260206004820152603a60248201526000805160206117a283398151915260448201527f3a20696e76616c6964207369676e6174757265206c656e67746800000000000060648201526084016104cb565b6000826040815181106109d4576109d46116a4565b016020015160f81c905060006109ea8482610f00565b905060006109f9856020610f00565b90506fa2a8918ca85bafe22016d0b997e4df60600160ff1b03811115610a755760405162461bcd60e51b815260206004820152603d60248201526000805160206117a283398151915260448201527f3a20696e76616c6964207369676e6174757265202773272076616c756500000060648201526084016104cb565b8260ff16601b14158015610a8d57508260ff16601c14155b15610aee5760405162461bcd60e51b815260206004820152603d60248201526000805160206117a283398151915260448201527f3a20696e76616c6964207369676e6174757265202776272076616c756500000060648201526084016104cb565b60408051600081526020810180835288905260ff851691810191909152606081018390526080810182905260019060a0016020604051602081039080840390855afa158015610b41573d6000803e3d6000fd5b5050604051601f1901519450506001600160a01b038416610bab5760405162461bcd60e51b815260206004820152603060248201526000805160206117a283398151915260448201526f1d1024a72b20a624a22fa9a4a3a722a960811b60648201526084016104cb565b50505092915050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610c2c5760405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e740000000060448201526064016104cb565b565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c829052603c8120610cab610c6e6101408601866116ba565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508593925050610f669050565b6000546201000090046001600160a01b03908116911614610cd05760019150506103bc565b5060009392505050565b50565b8015610cda57604051600090339060001990849084818181858888f193505050503d806000811461067f576040519150601f19603f3d011682016040523d82523d6000602084013e61067f565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610d7157506000546201000090046001600160a01b031633145b610c2c5760405162461bcd60e51b815260206004820181905260248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e7460448201526064016104cb565b600080846001600160a01b03168484604051610dd99190611749565b60006040518083038185875af1925050503d8060008114610e16576040519150601f19603f3d011682016040523d82523d6000602084013e610e1b565b606091505b50915091508161067f57805160208201fd5b6000546201000090046001600160a01b0316331480610e4b57503330145b610c2c5760405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b60448201526064016104cb565b6000805462010000600160b01b031916620100006001600160a01b038481168202929092178084556040519190048216927f0000000000000000000000000000000000000000000000000000000000000000909216917f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de91a350565b6000610f0d826020611778565b83511015610f5d5760405162461bcd60e51b815260206004820181905260248201527f72656164427974657333323a20696e76616c69642064617461206c656e67746860448201526064016104cb565b50016020015190565b6000806000610f758585610f8a565b91509150610f8281610fcf565b509392505050565b6000808251604103610fc05760208301516040840151606085015160001a610fb487828585611119565b94509450505050610fc8565b506000905060025b9250929050565b6000816004811115610fe357610fe361178b565b03610feb5750565b6001816004811115610fff57610fff61178b565b0361104c5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016104cb565b60028160048111156110605761106061178b565b036110ad5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016104cb565b60038160048111156110c1576110c161178b565b03610cda5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016104cb565b6000806fa2a8918ca85bafe22016d0b997e4df60600160ff1b0383111561114657506000905060036111ca565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561119a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166111c3576000600192509250506111ca565b9150600090505b94509492505050565b6001600160a01b0381168114610cda57600080fd5b60008083601f8401126111fa57600080fd5b50813567ffffffffffffffff81111561121257600080fd5b602083019150836020828501011115610fc857600080fd5b60008060008060008060008060c0898b03121561124657600080fd5b8835611251816111d3565b97506020890135611261816111d3565b96506040890135611271816111d3565b955060608901359450608089013567ffffffffffffffff8082111561129557600080fd5b6112a18c838d016111e8565b909650945060a08b01359150808211156112ba57600080fd5b506112c78b828c016111e8565b999c989b5096995094979396929594505050565b6000602082840312156112ed57600080fd5b81356001600160e01b03198116811461044957600080fd5b60008060008060006080868803121561131d57600080fd5b8535611328816111d3565b94506020860135611338816111d3565b935060408601359250606086013567ffffffffffffffff81111561135b57600080fd5b611367888289016111e8565b969995985093965092949392505050565b60008060006040848603121561138d57600080fd5b83359250602084013567ffffffffffffffff8111156113ab57600080fd5b6113b7868287016111e8565b9497909650939450505050565b6000806000606084860312156113d957600080fd5b833567ffffffffffffffff8111156113f057600080fd5b8401610160818703121561140357600080fd5b95602085013595506040909401359392505050565b60008083601f84011261142a57600080fd5b50813567ffffffffffffffff81111561144257600080fd5b6020830191508360208260051b8501011115610fc857600080fd5b6000806000806000806060878903121561147657600080fd5b863567ffffffffffffffff8082111561148e57600080fd5b61149a8a838b01611418565b909850965060208901359150808211156114b357600080fd5b6114bf8a838b01611418565b909650945060408901359150808211156114d857600080fd5b506114e589828a01611418565b979a9699509497509295939492505050565b6000806040838503121561150a57600080fd5b8235611515816111d3565b946020939093013593505050565b6000806000806060858703121561153957600080fd5b8435611544816111d3565b935060208501359250604085013567ffffffffffffffff81111561156757600080fd5b611573878288016111e8565b95989497509550505050565b60008060008060008060008060a0898b03121561159b57600080fd5b88356115a6816111d3565b975060208901356115b6816111d3565b9650604089013567ffffffffffffffff808211156115d357600080fd5b6115df8c838d01611418565b909850965060608b01359150808211156115f857600080fd5b6116048c838d01611418565b909650945060808b01359150808211156112ba57600080fd5b60006020828403121561162f57600080fd5b8135610449816111d3565b60008060008060008060a0878903121561165357600080fd5b863561165e816111d3565b9550602087013561166e816111d3565b94506040870135935060608701359250608087013567ffffffffffffffff81111561169857600080fd5b6114e589828a016111e8565b634e487b7160e01b600052603260045260246000fd5b6000808335601e198436030181126116d157600080fd5b83018035915067ffffffffffffffff8211156116ec57600080fd5b602001915036819003821315610fc857600080fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161172957611729611701565b5060010190565b60006020828403121561174257600080fd5b5051919050565b6000825160005b8181101561176a5760208186018101518583015201611750565b506000920191825250919050565b808201808211156103bc576103bc611701565b634e487b7160e01b600052602160045260246000fdfe5369676e617475726556616c696461746f72237265636f7665725369676e6572a26469706673582212205cbf0bd1b0e6f3b9d89c32c22db831d7ac66fc0f7967b42f16bf7c55bdf4a6dc64736f6c63430008140033",
}

// AccountABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountMetaData.ABI instead.
var AccountABI = AccountMetaData.ABI

// AccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccountMetaData.Bin instead.
var AccountBin = AccountMetaData.Bin

// DeployAccount deploys a new Ethereum contract, binding an instance of Account to it.
func DeployAccount(auth *bind.TransactOpts, backend bind.ContractBackend, anEntryPoint common.Address) (common.Address, *types.Transaction, *Account, error) {
	parsed, err := AccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccountBin), backend, anEntryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Account{AccountCaller: AccountCaller{contract: contract}, AccountTransactor: AccountTransactor{contract: contract}, AccountFilterer: AccountFilterer{contract: contract}}, nil
}

// Account is an auto generated Go binding around an Ethereum contract.
type Account struct {
	AccountCaller     // Read-only binding to the contract
	AccountTransactor // Write-only binding to the contract
	AccountFilterer   // Log filterer for contract events
}

// AccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountSession struct {
	Contract     *Account          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountCallerSession struct {
	Contract *AccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountTransactorSession struct {
	Contract     *AccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountRaw struct {
	Contract *Account // Generic contract binding to access the raw methods on
}

// AccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountCallerRaw struct {
	Contract *AccountCaller // Generic read-only contract binding to access the raw methods on
}

// AccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountTransactorRaw struct {
	Contract *AccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccount creates a new instance of Account, bound to a specific deployed contract.
func NewAccount(address common.Address, backend bind.ContractBackend) (*Account, error) {
	contract, err := bindAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Account{AccountCaller: AccountCaller{contract: contract}, AccountTransactor: AccountTransactor{contract: contract}, AccountFilterer: AccountFilterer{contract: contract}}, nil
}

// NewAccountCaller creates a new read-only instance of Account, bound to a specific deployed contract.
func NewAccountCaller(address common.Address, caller bind.ContractCaller) (*AccountCaller, error) {
	contract, err := bindAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountCaller{contract: contract}, nil
}

// NewAccountTransactor creates a new write-only instance of Account, bound to a specific deployed contract.
func NewAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountTransactor, error) {
	contract, err := bindAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountTransactor{contract: contract}, nil
}

// NewAccountFilterer creates a new log filterer instance of Account, bound to a specific deployed contract.
func NewAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountFilterer, error) {
	contract, err := bindAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountFilterer{contract: contract}, nil
}

// bindAccount binds a generic wrapper to an already deployed contract.
func bindAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Account *AccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Account.Contract.AccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Account *AccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.Contract.AccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Account *AccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Account.Contract.AccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Account *AccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Account.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Account *AccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Account *AccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Account.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Account *AccountCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Account *AccountSession) EntryPoint() (common.Address, error) {
	return _Account.Contract.EntryPoint(&_Account.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Account *AccountCallerSession) EntryPoint() (common.Address, error) {
	return _Account.Contract.EntryPoint(&_Account.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Account *AccountCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Account *AccountSession) GetDeposit() (*big.Int, error) {
	return _Account.Contract.GetDeposit(&_Account.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Account *AccountCallerSession) GetDeposit() (*big.Int, error) {
	return _Account.Contract.GetDeposit(&_Account.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Account *AccountCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Account *AccountSession) GetNonce() (*big.Int, error) {
	return _Account.Contract.GetNonce(&_Account.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Account *AccountCallerSession) GetNonce() (*big.Int, error) {
	return _Account.Contract.GetNonce(&_Account.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_Account *AccountCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_Account *AccountSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _Account.Contract.IsValidSignature(&_Account.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_Account *AccountCallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _Account.Contract.IsValidSignature(&_Account.CallOpts, _hash, _signature)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Account *AccountCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Account *AccountSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Account.Contract.OnERC1155BatchReceived(&_Account.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Account *AccountCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Account.Contract.OnERC1155BatchReceived(&_Account.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Account *AccountCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Account *AccountSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Account.Contract.OnERC1155Received(&_Account.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Account *AccountCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Account.Contract.OnERC1155Received(&_Account.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Account *AccountCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Account *AccountSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Account.Contract.OnERC721Received(&_Account.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Account *AccountCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Account.Contract.OnERC721Received(&_Account.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Account *AccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Account *AccountSession) Owner() (common.Address, error) {
	return _Account.Contract.Owner(&_Account.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Account *AccountCallerSession) Owner() (common.Address, error) {
	return _Account.Contract.Owner(&_Account.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Account *AccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Account *AccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Account.Contract.SupportsInterface(&_Account.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Account *AccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Account.Contract.SupportsInterface(&_Account.CallOpts, interfaceId)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Account *AccountCaller) TokensReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "tokensReceived", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Account *AccountSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _Account.Contract.TokensReceived(&_Account.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Account *AccountCallerSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _Account.Contract.TokensReceived(&_Account.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Account *AccountTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Account *AccountSession) AddDeposit() (*types.Transaction, error) {
	return _Account.Contract.AddDeposit(&_Account.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Account *AccountTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _Account.Contract.AddDeposit(&_Account.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Account *AccountTransactor) Execute(opts *bind.TransactOpts, dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "execute", dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Account *AccountSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Account.Contract.Execute(&_Account.TransactOpts, dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Account *AccountTransactorSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Account.Contract.Execute(&_Account.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_Account *AccountTransactor) ExecuteBatch(opts *bind.TransactOpts, dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "executeBatch", dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_Account *AccountSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _Account.Contract.ExecuteBatch(&_Account.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] value, bytes[] func) returns()
func (_Account *AccountTransactorSession) ExecuteBatch(dest []common.Address, value []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _Account.Contract.ExecuteBatch(&_Account.TransactOpts, dest, value, arg2)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_Account *AccountTransactor) Initialize(opts *bind.TransactOpts, anOwner common.Address) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "initialize", anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_Account *AccountSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _Account.Contract.Initialize(&_Account.TransactOpts, anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_Account *AccountTransactorSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _Account.Contract.Initialize(&_Account.TransactOpts, anOwner)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Account *AccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Account *AccountSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Account.Contract.ValidateUserOp(&_Account.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Account *AccountTransactorSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Account.Contract.ValidateUserOp(&_Account.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Account *AccountTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Account *AccountSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Account.Contract.WithdrawDepositTo(&_Account.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Account *AccountTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Account.Contract.WithdrawDepositTo(&_Account.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Account *AccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Account *AccountSession) Receive() (*types.Transaction, error) {
	return _Account.Contract.Receive(&_Account.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Account *AccountTransactorSession) Receive() (*types.Transaction, error) {
	return _Account.Contract.Receive(&_Account.TransactOpts)
}

// AccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Account contract.
type AccountInitializedIterator struct {
	Event *AccountInitialized // Event containing the contract specifics and raw log

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
func (it *AccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountInitialized)
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
		it.Event = new(AccountInitialized)
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
func (it *AccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountInitialized represents a Initialized event raised by the Account contract.
type AccountInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Account *AccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*AccountInitializedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AccountInitializedIterator{contract: _Account.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Account *AccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AccountInitialized) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountInitialized)
				if err := _Account.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Account *AccountFilterer) ParseInitialized(log types.Log) (*AccountInitialized, error) {
	event := new(AccountInitialized)
	if err := _Account.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountSimpleAccountInitializedIterator is returned from FilterSimpleAccountInitialized and is used to iterate over the raw logs and unpacked data for SimpleAccountInitialized events raised by the Account contract.
type AccountSimpleAccountInitializedIterator struct {
	Event *AccountSimpleAccountInitialized // Event containing the contract specifics and raw log

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
func (it *AccountSimpleAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountSimpleAccountInitialized)
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
		it.Event = new(AccountSimpleAccountInitialized)
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
func (it *AccountSimpleAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountSimpleAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountSimpleAccountInitialized represents a SimpleAccountInitialized event raised by the Account contract.
type AccountSimpleAccountInitialized struct {
	EntryPoint common.Address
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSimpleAccountInitialized is a free log retrieval operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_Account *AccountFilterer) FilterSimpleAccountInitialized(opts *bind.FilterOpts, entryPoint []common.Address, owner []common.Address) (*AccountSimpleAccountInitializedIterator, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Account.contract.FilterLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &AccountSimpleAccountInitializedIterator{contract: _Account.contract, event: "SimpleAccountInitialized", logs: logs, sub: sub}, nil
}

// WatchSimpleAccountInitialized is a free log subscription operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_Account *AccountFilterer) WatchSimpleAccountInitialized(opts *bind.WatchOpts, sink chan<- *AccountSimpleAccountInitialized, entryPoint []common.Address, owner []common.Address) (event.Subscription, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Account.contract.WatchLogs(opts, "SimpleAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountSimpleAccountInitialized)
				if err := _Account.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
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

// ParseSimpleAccountInitialized is a log parse operation binding the contract event 0x47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de.
//
// Solidity: event SimpleAccountInitialized(address indexed entryPoint, address indexed owner)
func (_Account *AccountFilterer) ParseSimpleAccountInitialized(log types.Log) (*AccountSimpleAccountInitialized, error) {
	event := new(AccountSimpleAccountInitialized)
	if err := _Account.contract.UnpackLog(event, "SimpleAccountInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
