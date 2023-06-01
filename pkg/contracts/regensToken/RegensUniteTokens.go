// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package regensToken

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

// RegensTokenMetaData contains all meta data concerning the RegensToken contract.
var RegensTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"mintCustomVoucher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintGratitude\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b50604051620027a6380380620027a683398101604081905262000033916200027c565b60408051602081019091525f81526200004c81620000f0565b5060066200005b838262000418565b5060076200006a828262000418565b50620000775f3362000102565b5f5b8351811015620000e657620000d17f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6858381518110620000bd57620000bd620004e0565b60200260200101516200010260201b60201c565b80620000dd81620004f4565b91505062000079565b5050505062000519565b6002620000fe828262000418565b5050565b5f8281526003602090815260408083206001600160a01b0385168452909152902054620000fe908390839060ff16620000fe575f8281526003602090815260408083206001600160a01b03851684529091529020805460ff19166001179055620001693390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715620001ec57620001ec620001ad565b604052919050565b5f82601f83011262000204575f80fd5b81516001600160401b03811115620002205762000220620001ad565b602062000236601f8301601f19168201620001c1565b82815285828487010111156200024a575f80fd5b5f5b83811015620002695785810183015182820184015282016200024c565b505f928101909101919091529392505050565b5f805f606084860312156200028f575f80fd5b83516001600160401b0380821115620002a6575f80fd5b818601915086601f830112620002ba575f80fd5b8151602082821115620002d157620002d1620001ad565b8160051b620002e2828201620001c1565b928352848101820192828101908b851115620002fc575f80fd5b958301955b848710156200033557865192506001600160a01b038316831462000324575f8081fd5b828252958301959083019062000301565b928a0151929850919450505050808211156200034f575f80fd5b6200035d87838801620001f4565b9350604086015191508082111562000373575f80fd5b506200038286828701620001f4565b9150509250925092565b600181811c90821680620003a157607f821691505b602082108103620003c057634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111562000413575f81815260208120601f850160051c81016020861015620003ee5750805b601f850160051c820191505b818110156200040f57828155600101620003fa565b5050505b505050565b81516001600160401b03811115620004345762000434620001ad565b6200044c816200044584546200038c565b84620003c6565b602080601f83116001811462000482575f84156200046a5750858301515b5f19600386901b1c1916600185901b1785556200040f565b5f85815260208120601f198616915b82811015620004b25788860151825594840194600190910190840162000491565b5085821015620004d057878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52603260045260245ffd5b5f600182016200051257634e487b7160e01b5f52601160045260245ffd5b5060010190565b61227f80620005275f395ff3fe608060405234801561000f575f80fd5b5060043610610131575f3560e01c80634e1273f4116100b4578063a22cb46511610079578063a22cb4651461027e578063b390c0ab14610291578063d5391393146102a4578063d547741f146102cb578063e985e9c5146102de578063f242432a14610319575f80fd5b80634e1273f4146102295780636bd286ca1461024957806391d148541461025c57806395d89b411461026f578063a217fddf14610277575f80fd5b8063248a9ca3116100fa578063248a9ca3146101bb57806324f1ae7e146101dd5780632eb2c2d6146101f05780632f2ff15d1461020357806336568abe14610216575f80fd5b8062fdd58e1461013557806301ffc9a71461015b57806306fdde031461017e5780630cc656bd146101935780630e89341c146101a8575b5f80fd5b6101486101433660046117a2565b61032c565b6040519081526020015b60405180910390f35b61016e6101693660046117df565b6103c3565b6040519015158152602001610152565b6101866103cd565b6040516101529190611847565b6101a66101a1366004611859565b610459565b005b6101866101b6366004611872565b610478565b6101486101c9366004611872565b5f9081526003602052604090206001015490565b6101a66101eb366004611943565b610573565b6101a66101fe366004611a27565b6105fb565b6101a6610211366004611aca565b610647565b6101a6610224366004611aca565b61066b565b61023c610237366004611af4565b6106e9565b6040516101529190611bf2565b6101a6610257366004611c04565b610811565b61016e61026a366004611aca565b610913565b61018661093d565b6101485f81565b6101a661028c366004611c8c565b61094a565b6101a661029f366004611cc5565b610955565b6101487f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b6101a66102d9366004611aca565b610960565b61016e6102ec366004611ce5565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205460ff1690565b6101a6610327366004611d0d565b610984565b5f6001600160a01b03831661039b5760405162461bcd60e51b815260206004820152602a60248201527f455243313135353a2061646472657373207a65726f206973206e6f742061207660448201526930b634b21037bbb732b960b11b60648201526084015b60405180910390fd5b505f818152602081815260408083206001600160a01b03861684529091529020545b92915050565b5f6103bd826109c9565b600680546103da90611d6d565b80601f016020809104026020016040519081016040528092919081815260200182805461040690611d6d565b80156104515780601f1061042857610100808354040283529160200191610451565b820191905f5260205f20905b81548152906001019060200180831161043457829003601f168201915b505050505081565b610475816002600160405180602001604052805f8152506109ed565b50565b5f8181526004602052604090205460609060ff166104d85760405162461bcd60e51b815260206004820152601f60248201527f55524920717565727920666f72206e6f6e6578697374656e7420746f6b656e006044820152606401610392565b5f82815260056020526040902080546104f090611d6d565b80601f016020809104026020016040519081016040528092919081815260200182805461051c90611d6d565b80156105675780601f1061053e57610100808354040283529160200191610567565b820191905f5260205f20905b81548152906001019060200180831161054a57829003601f168201915b50505050509050919050565b61059d7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a633610913565b6105e95760405162461bcd60e51b815260206004820152601d60248201527f4d7573742068617665206d696e74657220726f6c6520746f206d696e740000006044820152606401610392565b6105f683600184846109ed565b505050565b6001600160a01b038516331480610617575061061785336102ec565b6106335760405162461bcd60e51b815260040161039290611da5565b6106408585858585610a16565b5050505050565b5f8281526003602052604090206001015461066181610bf4565b6105f68383610bfe565b6001600160a01b03811633146106db5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152608401610392565b6106e58282610c83565b5050565b6060815183511461074e5760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b6064820152608401610392565b5f835167ffffffffffffffff81111561076957610769611889565b604051908082528060200260200182016040528015610792578160200160208202803683370190505b5090505f5b8451811015610809576107dc8582815181106107b5576107b5611df3565b60200260200101518583815181106107cf576107cf611df3565b602002602001015161032c565b8282815181106107ee576107ee611df3565b602090810291909101015261080281611e1b565b9050610797565b509392505050565b61271084116108755760405162461bcd60e51b815260206004820152602a60248201527f437573746f6d20746f6b656e204944206d75737420626520677265617465722060448201526907468616e2031303030360b41b6064820152608401610392565b5f8481526004602052604090205460ff16156108d35760405162461bcd60e51b815260206004820181905260248201527f546f6b656e2049442068617320616c7265616479206265656e206d696e7465646044820152606401610392565b6108df858585856109ed565b5f848152600460209081526040808320805460ff191660011790556005909152902061090b8282611e78565b505050505050565b5f9182526003602090815260408084206001600160a01b0393909316845291905290205460ff1690565b600780546103da90611d6d565b6106e5338383610ce9565b6106e5338383610dc8565b5f8281526003602052604090206001015461097a81610bf4565b6105f68383610c83565b6001600160a01b0385163314806109a057506109a085336102ec565b6109bc5760405162461bcd60e51b815260040161039290611da5565b6106408585858585610e4c565b5f6001600160e01b03198216637965db0b60e01b14806103bd57506103bd82610f80565b6109f984848484610fcf565b50505f908152600460205260409020805460ff1916600117905550565b8151835114610a785760405162461bcd60e51b815260206004820152602860248201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206044820152670dad2e6dac2e8c6d60c31b6064820152608401610392565b6001600160a01b038416610a9e5760405162461bcd60e51b815260040161039290611f34565b33610aad8187878787876110ec565b5f5b8451811015610b8e575f858281518110610acb57610acb611df3565b602002602001015190505f858381518110610ae857610ae8611df3565b6020908102919091018101515f84815280835260408082206001600160a01b038e168352909352919091205490915081811015610b375760405162461bcd60e51b815260040161039290611f79565b5f838152602081815260408083206001600160a01b038e8116855292528083208585039055908b16825281208054849290610b73908490611fc3565b9250508190555050505080610b8790611e1b565b9050610aaf565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051610bde929190611fd6565b60405180910390a461090b818787878787611199565b61047581336112f3565b610c088282610913565b6106e5575f8281526003602090815260408083206001600160a01b03851684529091529020805460ff19166001179055610c3f3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b610c8d8282610913565b156106e5575f8281526003602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b816001600160a01b0316836001600160a01b031603610d5c5760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b6064820152608401610392565b6001600160a01b038381165f81815260016020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b81600203610e185760405162461bcd60e51b815260206004820181905260248201527f47524154495455444520746f6b656e732063616e6e6f74206265206275726e746044820152606401610392565b610e2383838361134c565b610e2d838361032c565b5f036105f657505f908152600460205260409020805460ff1916905550565b6001600160a01b038416610e725760405162461bcd60e51b815260040161039290611f34565b335f610e7d856114d5565b90505f610e89856114d5565b9050610e998389898585896110ec565b5f868152602081815260408083206001600160a01b038c16845290915290205485811015610ed95760405162461bcd60e51b815260040161039290611f79565b5f878152602081815260408083206001600160a01b038d8116855292528083208985039055908a16825281208054889290610f15908490611fc3565b909155505060408051888152602081018890526001600160a01b03808b16928c821692918816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4610f75848a8a8a8a8a61151e565b505050505050505050565b5f6001600160e01b03198216636cdb3d1360e11b1480610fb057506001600160e01b031982166303a24d0760e21b145b806103bd57506301ffc9a760e01b6001600160e01b03198316146103bd565b6001600160a01b03841661102f5760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736044820152607360f81b6064820152608401610392565b335f61103a856114d5565b90505f611046856114d5565b9050611056835f898585896110ec565b5f868152602081815260408083206001600160a01b038b16845290915281208054879290611085908490611fc3565b909155505060408051878152602081018790526001600160a01b03808a16925f92918716917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a46110e3835f8989898961151e565b50505050505050565b5f5b83518110156110e3575f84828151811061110a5761110a611df3565b6020026020010151905080600214801561112c57506001600160a01b03871615155b156111885760405162461bcd60e51b815260206004820152602660248201527f47524154495455444520746f6b656e732063616e6e6f74206265207472616e7360448201526519995c9c995960d21b6064820152608401610392565b5061119281611e1b565b90506110ee565b6001600160a01b0384163b1561090b5760405163bc197c8160e01b81526001600160a01b0385169063bc197c81906111dd9089908990889088908890600401612003565b6020604051808303815f875af1925050508015611217575060408051601f3d908101601f1916820190925261121491810190612060565b60015b6112c35761122361207b565b806308c379a00361125c5750611237612094565b80611242575061125e565b8060405162461bcd60e51b81526004016103929190611847565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e2d455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b6064820152608401610392565b6001600160e01b0319811663bc197c8160e01b146110e35760405162461bcd60e51b81526004016103929061211d565b6112fd8282610913565b6106e55761130a816115d8565b6113158360206115ea565b604051602001611326929190612165565b60408051601f198184030181529082905262461bcd60e51b825261039291600401611847565b6001600160a01b0383166113ae5760405162461bcd60e51b815260206004820152602360248201527f455243313135353a206275726e2066726f6d20746865207a65726f206164647260448201526265737360e81b6064820152608401610392565b335f6113b9846114d5565b90505f6113c5846114d5565b90506113e383875f858560405180602001604052805f8152506110ec565b5f858152602081815260408083206001600160a01b038a1684529091529020548481101561145f5760405162461bcd60e51b8152602060048201526024808201527f455243313135353a206275726e20616d6f756e7420657863656564732062616c604482015263616e636560e01b6064820152608401610392565b5f868152602081815260408083206001600160a01b038b81168086529184528285208a8703905582518b81529384018a90529092908816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a460408051602081019091525f90526110e3565b6040805160018082528183019092526060915f91906020808301908036833701905050905082815f8151811061150d5761150d611df3565b602090810291909101015292915050565b6001600160a01b0384163b1561090b5760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e619061156290899089908890889088906004016121d9565b6020604051808303815f875af192505050801561159c575060408051601f3d908101601f1916820190925261159991810190612060565b60015b6115a85761122361207b565b6001600160e01b0319811663f23a6e6160e01b146110e35760405162461bcd60e51b81526004016103929061211d565b60606103bd6001600160a01b03831660145b60605f6115f883600261221d565b611603906002611fc3565b67ffffffffffffffff81111561161b5761161b611889565b6040519080825280601f01601f191660200182016040528015611645576020820181803683370190505b509050600360fc1b815f8151811061165f5761165f611df3565b60200101906001600160f81b03191690815f1a905350600f60fb1b8160018151811061168d5761168d611df3565b60200101906001600160f81b03191690815f1a9053505f6116af84600261221d565b6116ba906001611fc3565b90505b6001811115611731576f181899199a1a9b1b9c1cb0b131b232b360811b85600f16601081106116ee576116ee611df3565b1a60f81b82828151811061170457611704611df3565b60200101906001600160f81b03191690815f1a90535060049490941c9361172a81612234565b90506116bd565b5083156117805760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610392565b9392505050565b80356001600160a01b038116811461179d575f80fd5b919050565b5f80604083850312156117b3575f80fd5b6117bc83611787565b946020939093013593505050565b6001600160e01b031981168114610475575f80fd5b5f602082840312156117ef575f80fd5b8135611780816117ca565b5f5b838110156118145781810151838201526020016117fc565b50505f910152565b5f81518084526118338160208601602086016117fa565b601f01601f19169290920160200192915050565b602081525f611780602083018461181c565b5f60208284031215611869575f80fd5b61178082611787565b5f60208284031215611882575f80fd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b601f8201601f1916810167ffffffffffffffff811182821017156118c3576118c3611889565b6040525050565b5f67ffffffffffffffff8311156118e3576118e3611889565b6040516118fa601f8501601f19166020018261189d565b80915083815284848401111561190e575f80fd5b838360208301375f60208583010152509392505050565b5f82601f830112611934575f80fd5b611780838335602085016118ca565b5f805f60608486031215611955575f80fd5b61195e84611787565b925060208401359150604084013567ffffffffffffffff811115611980575f80fd5b61198c86828701611925565b9150509250925092565b5f67ffffffffffffffff8211156119af576119af611889565b5060051b60200190565b5f82601f8301126119c8575f80fd5b813560206119d582611996565b6040516119e2828261189d565b83815260059390931b8501820192828101915086841115611a01575f80fd5b8286015b84811015611a1c5780358352918301918301611a05565b509695505050505050565b5f805f805f60a08688031215611a3b575f80fd5b611a4486611787565b9450611a5260208701611787565b9350604086013567ffffffffffffffff80821115611a6e575f80fd5b611a7a89838a016119b9565b94506060880135915080821115611a8f575f80fd5b611a9b89838a016119b9565b93506080880135915080821115611ab0575f80fd5b50611abd88828901611925565b9150509295509295909350565b5f8060408385031215611adb575f80fd5b82359150611aeb60208401611787565b90509250929050565b5f8060408385031215611b05575f80fd5b823567ffffffffffffffff80821115611b1c575f80fd5b818501915085601f830112611b2f575f80fd5b81356020611b3c82611996565b604051611b49828261189d565b83815260059390931b8501820192828101915089841115611b68575f80fd5b948201945b83861015611b8d57611b7e86611787565b82529482019490820190611b6d565b96505086013592505080821115611ba2575f80fd5b50611baf858286016119b9565b9150509250929050565b5f8151808452602080850194508084015f5b83811015611be757815187529582019590820190600101611bcb565b509495945050505050565b602081525f6117806020830184611bb9565b5f805f805f60a08688031215611c18575f80fd5b611c2186611787565b94506020860135935060408601359250606086013567ffffffffffffffff80821115611c4b575f80fd5b611c5789838a01611925565b93506080880135915080821115611c6c575f80fd5b508601601f81018813611c7d575f80fd5b611abd888235602084016118ca565b5f8060408385031215611c9d575f80fd5b611ca683611787565b915060208301358015158114611cba575f80fd5b809150509250929050565b5f8060408385031215611cd6575f80fd5b50508035926020909101359150565b5f8060408385031215611cf6575f80fd5b611cff83611787565b9150611aeb60208401611787565b5f805f805f60a08688031215611d21575f80fd5b611d2a86611787565b9450611d3860208701611787565b93506040860135925060608601359150608086013567ffffffffffffffff811115611d61575f80fd5b611abd88828901611925565b600181811c90821680611d8157607f821691505b602082108103611d9f57634e487b7160e01b5f52602260045260245ffd5b50919050565b6020808252602e908201527f455243313135353a2063616c6c6572206973206e6f7420746f6b656e206f776e60408201526d195c881bdc88185c1c1c9bdd995960921b606082015260800190565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b5f60018201611e2c57611e2c611e07565b5060010190565b601f8211156105f6575f81815260208120601f850160051c81016020861015611e595750805b601f850160051c820191505b8181101561090b57828155600101611e65565b815167ffffffffffffffff811115611e9257611e92611889565b611ea681611ea08454611d6d565b84611e33565b602080601f831160018114611ed9575f8415611ec25750858301515b5f19600386901b1c1916600185901b17855561090b565b5f85815260208120601f198616915b82811015611f0757888601518255948401946001909101908401611ee8565b5085821015611f2457878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b808201808211156103bd576103bd611e07565b604081525f611fe86040830185611bb9565b8281036020840152611ffa8185611bb9565b95945050505050565b6001600160a01b0386811682528516602082015260a0604082018190525f9061202e90830186611bb9565b82810360608401526120408186611bb9565b90508281036080840152612054818561181c565b98975050505050505050565b5f60208284031215612070575f80fd5b8151611780816117ca565b5f60033d11156120915760045f803e505f5160e01c5b90565b5f60443d10156120a15790565b6040516003193d81016004833e81513d67ffffffffffffffff81602484011181841117156120d157505050505090565b82850191508151818111156120e95750505050505090565b843d87010160208285010111156121035750505050505090565b6121126020828601018761189d565b509095945050505050565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b7f416363657373436f6e74726f6c3a206163636f756e742000000000000000000081525f835161219c8160178501602088016117fa565b7001034b99036b4b9b9b4b733903937b6329607d1b60179184019182015283516121cd8160288401602088016117fa565b01602801949350505050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a0608082018190525f906122129083018461181c565b979650505050505050565b80820281158282048414176103bd576103bd611e07565b5f8161224257612242611e07565b505f19019056fea2646970667358221220028540a0d2603bb7b5751b4f6302a98b8415c5bb792d072abb4b29988817db9d64736f6c63430008140033",
}

// RegensTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use RegensTokenMetaData.ABI instead.
var RegensTokenABI = RegensTokenMetaData.ABI

// RegensTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegensTokenMetaData.Bin instead.
var RegensTokenBin = RegensTokenMetaData.Bin

// DeployRegensToken deploys a new Ethereum contract, binding an instance of RegensToken to it.
func DeployRegensToken(auth *bind.TransactOpts, backend bind.ContractBackend, admins []common.Address, _name string, _symbol string) (common.Address, *types.Transaction, *RegensToken, error) {
	parsed, err := RegensTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegensTokenBin), backend, admins, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RegensToken{RegensTokenCaller: RegensTokenCaller{contract: contract}, RegensTokenTransactor: RegensTokenTransactor{contract: contract}, RegensTokenFilterer: RegensTokenFilterer{contract: contract}}, nil
}

// RegensToken is an auto generated Go binding around an Ethereum contract.
type RegensToken struct {
	RegensTokenCaller     // Read-only binding to the contract
	RegensTokenTransactor // Write-only binding to the contract
	RegensTokenFilterer   // Log filterer for contract events
}

// RegensTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegensTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegensTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegensTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegensTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegensTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegensTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegensTokenSession struct {
	Contract     *RegensToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegensTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegensTokenCallerSession struct {
	Contract *RegensTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RegensTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegensTokenTransactorSession struct {
	Contract     *RegensTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RegensTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegensTokenRaw struct {
	Contract *RegensToken // Generic contract binding to access the raw methods on
}

// RegensTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegensTokenCallerRaw struct {
	Contract *RegensTokenCaller // Generic read-only contract binding to access the raw methods on
}

// RegensTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegensTokenTransactorRaw struct {
	Contract *RegensTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegensToken creates a new instance of RegensToken, bound to a specific deployed contract.
func NewRegensToken(address common.Address, backend bind.ContractBackend) (*RegensToken, error) {
	contract, err := bindRegensToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegensToken{RegensTokenCaller: RegensTokenCaller{contract: contract}, RegensTokenTransactor: RegensTokenTransactor{contract: contract}, RegensTokenFilterer: RegensTokenFilterer{contract: contract}}, nil
}

// NewRegensTokenCaller creates a new read-only instance of RegensToken, bound to a specific deployed contract.
func NewRegensTokenCaller(address common.Address, caller bind.ContractCaller) (*RegensTokenCaller, error) {
	contract, err := bindRegensToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegensTokenCaller{contract: contract}, nil
}

// NewRegensTokenTransactor creates a new write-only instance of RegensToken, bound to a specific deployed contract.
func NewRegensTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*RegensTokenTransactor, error) {
	contract, err := bindRegensToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegensTokenTransactor{contract: contract}, nil
}

// NewRegensTokenFilterer creates a new log filterer instance of RegensToken, bound to a specific deployed contract.
func NewRegensTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*RegensTokenFilterer, error) {
	contract, err := bindRegensToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegensTokenFilterer{contract: contract}, nil
}

// bindRegensToken binds a generic wrapper to an already deployed contract.
func bindRegensToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegensTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegensToken *RegensTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegensToken.Contract.RegensTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegensToken *RegensTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegensToken.Contract.RegensTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegensToken *RegensTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegensToken.Contract.RegensTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegensToken *RegensTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegensToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegensToken *RegensTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegensToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegensToken *RegensTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegensToken.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RegensToken *RegensTokenCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RegensToken.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RegensToken *RegensTokenSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RegensToken.Contract.DEFAULTADMINROLE(&_RegensToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RegensToken *RegensTokenCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RegensToken.Contract.DEFAULTADMINROLE(&_RegensToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_RegensToken *RegensTokenCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RegensToken.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_RegensToken *RegensTokenSession) MINTERROLE() ([32]byte, error) {
	return _RegensToken.Contract.MINTERROLE(&_RegensToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_RegensToken *RegensTokenCallerSession) MINTERROLE() ([32]byte, error) {
	return _RegensToken.Contract.MINTERROLE(&_RegensToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_RegensToken *RegensTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RegensToken.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_RegensToken *RegensTokenSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _RegensToken.Contract.BalanceOf(&_RegensToken.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_RegensToken *RegensTokenCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _RegensToken.Contract.BalanceOf(&_RegensToken.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_RegensToken *RegensTokenCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _RegensToken.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_RegensToken *RegensTokenSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _RegensToken.Contract.BalanceOfBatch(&_RegensToken.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_RegensToken *RegensTokenCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _RegensToken.Contract.BalanceOfBatch(&_RegensToken.CallOpts, accounts, ids)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RegensToken *RegensTokenCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RegensToken.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RegensToken *RegensTokenSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RegensToken.Contract.GetRoleAdmin(&_RegensToken.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RegensToken *RegensTokenCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RegensToken.Contract.GetRoleAdmin(&_RegensToken.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RegensToken *RegensTokenCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RegensToken.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RegensToken *RegensTokenSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RegensToken.Contract.HasRole(&_RegensToken.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RegensToken *RegensTokenCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RegensToken.Contract.HasRole(&_RegensToken.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_RegensToken *RegensTokenCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _RegensToken.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_RegensToken *RegensTokenSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _RegensToken.Contract.IsApprovedForAll(&_RegensToken.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_RegensToken *RegensTokenCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _RegensToken.Contract.IsApprovedForAll(&_RegensToken.CallOpts, account, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RegensToken *RegensTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RegensToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RegensToken *RegensTokenSession) Name() (string, error) {
	return _RegensToken.Contract.Name(&_RegensToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RegensToken *RegensTokenCallerSession) Name() (string, error) {
	return _RegensToken.Contract.Name(&_RegensToken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RegensToken *RegensTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RegensToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RegensToken *RegensTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RegensToken.Contract.SupportsInterface(&_RegensToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RegensToken *RegensTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RegensToken.Contract.SupportsInterface(&_RegensToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RegensToken *RegensTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RegensToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RegensToken *RegensTokenSession) Symbol() (string, error) {
	return _RegensToken.Contract.Symbol(&_RegensToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RegensToken *RegensTokenCallerSession) Symbol() (string, error) {
	return _RegensToken.Contract.Symbol(&_RegensToken.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_RegensToken *RegensTokenCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _RegensToken.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_RegensToken *RegensTokenSession) Uri(tokenId *big.Int) (string, error) {
	return _RegensToken.Contract.Uri(&_RegensToken.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_RegensToken *RegensTokenCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _RegensToken.Contract.Uri(&_RegensToken.CallOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 id, uint256 amount) returns()
func (_RegensToken *RegensTokenTransactor) Burn(opts *bind.TransactOpts, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _RegensToken.contract.Transact(opts, "burn", id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 id, uint256 amount) returns()
func (_RegensToken *RegensTokenSession) Burn(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _RegensToken.Contract.Burn(&_RegensToken.TransactOpts, id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 id, uint256 amount) returns()
func (_RegensToken *RegensTokenTransactorSession) Burn(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _RegensToken.Contract.Burn(&_RegensToken.TransactOpts, id, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RegensToken *RegensTokenTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegensToken.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RegensToken *RegensTokenSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegensToken.Contract.GrantRole(&_RegensToken.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RegensToken *RegensTokenTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegensToken.Contract.GrantRole(&_RegensToken.TransactOpts, role, account)
}

// MintCoin is a paid mutator transaction binding the contract method 0x24f1ae7e.
//
// Solidity: function mintCoin(address to, uint256 amount, bytes data) returns()
func (_RegensToken *RegensTokenTransactor) MintCoin(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _RegensToken.contract.Transact(opts, "mintCoin", to, amount, data)
}

// MintCoin is a paid mutator transaction binding the contract method 0x24f1ae7e.
//
// Solidity: function mintCoin(address to, uint256 amount, bytes data) returns()
func (_RegensToken *RegensTokenSession) MintCoin(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _RegensToken.Contract.MintCoin(&_RegensToken.TransactOpts, to, amount, data)
}

// MintCoin is a paid mutator transaction binding the contract method 0x24f1ae7e.
//
// Solidity: function mintCoin(address to, uint256 amount, bytes data) returns()
func (_RegensToken *RegensTokenTransactorSession) MintCoin(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _RegensToken.Contract.MintCoin(&_RegensToken.TransactOpts, to, amount, data)
}

// MintCustomVoucher is a paid mutator transaction binding the contract method 0x6bd286ca.
//
// Solidity: function mintCustomVoucher(address to, uint256 id, uint256 amount, bytes data, string _uri) returns()
func (_RegensToken *RegensTokenTransactor) MintCustomVoucher(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int, data []byte, _uri string) (*types.Transaction, error) {
	return _RegensToken.contract.Transact(opts, "mintCustomVoucher", to, id, amount, data, _uri)
}

// MintCustomVoucher is a paid mutator transaction binding the contract method 0x6bd286ca.
//
// Solidity: function mintCustomVoucher(address to, uint256 id, uint256 amount, bytes data, string _uri) returns()
func (_RegensToken *RegensTokenSession) MintCustomVoucher(to common.Address, id *big.Int, amount *big.Int, data []byte, _uri string) (*types.Transaction, error) {
	return _RegensToken.Contract.MintCustomVoucher(&_RegensToken.TransactOpts, to, id, amount, data, _uri)
}

// MintCustomVoucher is a paid mutator transaction binding the contract method 0x6bd286ca.
//
// Solidity: function mintCustomVoucher(address to, uint256 id, uint256 amount, bytes data, string _uri) returns()
func (_RegensToken *RegensTokenTransactorSession) MintCustomVoucher(to common.Address, id *big.Int, amount *big.Int, data []byte, _uri string) (*types.Transaction, error) {
	return _RegensToken.Contract.MintCustomVoucher(&_RegensToken.TransactOpts, to, id, amount, data, _uri)
}

// MintGratitude is a paid mutator transaction binding the contract method 0x0cc656bd.
//
// Solidity: function mintGratitude(address to) returns()
func (_RegensToken *RegensTokenTransactor) MintGratitude(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _RegensToken.contract.Transact(opts, "mintGratitude", to)
}

// MintGratitude is a paid mutator transaction binding the contract method 0x0cc656bd.
//
// Solidity: function mintGratitude(address to) returns()
func (_RegensToken *RegensTokenSession) MintGratitude(to common.Address) (*types.Transaction, error) {
	return _RegensToken.Contract.MintGratitude(&_RegensToken.TransactOpts, to)
}

// MintGratitude is a paid mutator transaction binding the contract method 0x0cc656bd.
//
// Solidity: function mintGratitude(address to) returns()
func (_RegensToken *RegensTokenTransactorSession) MintGratitude(to common.Address) (*types.Transaction, error) {
	return _RegensToken.Contract.MintGratitude(&_RegensToken.TransactOpts, to)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RegensToken *RegensTokenTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegensToken.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RegensToken *RegensTokenSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegensToken.Contract.RenounceRole(&_RegensToken.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RegensToken *RegensTokenTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegensToken.Contract.RenounceRole(&_RegensToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RegensToken *RegensTokenTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegensToken.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RegensToken *RegensTokenSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegensToken.Contract.RevokeRole(&_RegensToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RegensToken *RegensTokenTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegensToken.Contract.RevokeRole(&_RegensToken.TransactOpts, role, account)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_RegensToken *RegensTokenTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _RegensToken.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_RegensToken *RegensTokenSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _RegensToken.Contract.SafeBatchTransferFrom(&_RegensToken.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_RegensToken *RegensTokenTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _RegensToken.Contract.SafeBatchTransferFrom(&_RegensToken.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_RegensToken *RegensTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _RegensToken.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_RegensToken *RegensTokenSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _RegensToken.Contract.SafeTransferFrom(&_RegensToken.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_RegensToken *RegensTokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _RegensToken.Contract.SafeTransferFrom(&_RegensToken.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RegensToken *RegensTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _RegensToken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RegensToken *RegensTokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _RegensToken.Contract.SetApprovalForAll(&_RegensToken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RegensToken *RegensTokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _RegensToken.Contract.SetApprovalForAll(&_RegensToken.TransactOpts, operator, approved)
}

// RegensTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the RegensToken contract.
type RegensTokenApprovalForAllIterator struct {
	Event *RegensTokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *RegensTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegensTokenApprovalForAll)
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
		it.Event = new(RegensTokenApprovalForAll)
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
func (it *RegensTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegensTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegensTokenApprovalForAll represents a ApprovalForAll event raised by the RegensToken contract.
type RegensTokenApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_RegensToken *RegensTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*RegensTokenApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RegensToken.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &RegensTokenApprovalForAllIterator{contract: _RegensToken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_RegensToken *RegensTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *RegensTokenApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RegensToken.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegensTokenApprovalForAll)
				if err := _RegensToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_RegensToken *RegensTokenFilterer) ParseApprovalForAll(log types.Log) (*RegensTokenApprovalForAll, error) {
	event := new(RegensTokenApprovalForAll)
	if err := _RegensToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegensTokenRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RegensToken contract.
type RegensTokenRoleAdminChangedIterator struct {
	Event *RegensTokenRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RegensTokenRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegensTokenRoleAdminChanged)
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
		it.Event = new(RegensTokenRoleAdminChanged)
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
func (it *RegensTokenRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegensTokenRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegensTokenRoleAdminChanged represents a RoleAdminChanged event raised by the RegensToken contract.
type RegensTokenRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RegensToken *RegensTokenFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RegensTokenRoleAdminChangedIterator, error) {

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

	logs, sub, err := _RegensToken.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RegensTokenRoleAdminChangedIterator{contract: _RegensToken.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RegensToken *RegensTokenFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RegensTokenRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _RegensToken.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegensTokenRoleAdminChanged)
				if err := _RegensToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_RegensToken *RegensTokenFilterer) ParseRoleAdminChanged(log types.Log) (*RegensTokenRoleAdminChanged, error) {
	event := new(RegensTokenRoleAdminChanged)
	if err := _RegensToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegensTokenRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RegensToken contract.
type RegensTokenRoleGrantedIterator struct {
	Event *RegensTokenRoleGranted // Event containing the contract specifics and raw log

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
func (it *RegensTokenRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegensTokenRoleGranted)
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
		it.Event = new(RegensTokenRoleGranted)
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
func (it *RegensTokenRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegensTokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegensTokenRoleGranted represents a RoleGranted event raised by the RegensToken contract.
type RegensTokenRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RegensToken *RegensTokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegensTokenRoleGrantedIterator, error) {

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

	logs, sub, err := _RegensToken.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegensTokenRoleGrantedIterator{contract: _RegensToken.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RegensToken *RegensTokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RegensTokenRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RegensToken.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegensTokenRoleGranted)
				if err := _RegensToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_RegensToken *RegensTokenFilterer) ParseRoleGranted(log types.Log) (*RegensTokenRoleGranted, error) {
	event := new(RegensTokenRoleGranted)
	if err := _RegensToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegensTokenRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RegensToken contract.
type RegensTokenRoleRevokedIterator struct {
	Event *RegensTokenRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RegensTokenRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegensTokenRoleRevoked)
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
		it.Event = new(RegensTokenRoleRevoked)
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
func (it *RegensTokenRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegensTokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegensTokenRoleRevoked represents a RoleRevoked event raised by the RegensToken contract.
type RegensTokenRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RegensToken *RegensTokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegensTokenRoleRevokedIterator, error) {

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

	logs, sub, err := _RegensToken.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegensTokenRoleRevokedIterator{contract: _RegensToken.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RegensToken *RegensTokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RegensTokenRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RegensToken.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegensTokenRoleRevoked)
				if err := _RegensToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_RegensToken *RegensTokenFilterer) ParseRoleRevoked(log types.Log) (*RegensTokenRoleRevoked, error) {
	event := new(RegensTokenRoleRevoked)
	if err := _RegensToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegensTokenTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the RegensToken contract.
type RegensTokenTransferBatchIterator struct {
	Event *RegensTokenTransferBatch // Event containing the contract specifics and raw log

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
func (it *RegensTokenTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegensTokenTransferBatch)
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
		it.Event = new(RegensTokenTransferBatch)
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
func (it *RegensTokenTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegensTokenTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegensTokenTransferBatch represents a TransferBatch event raised by the RegensToken contract.
type RegensTokenTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_RegensToken *RegensTokenFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*RegensTokenTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegensToken.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RegensTokenTransferBatchIterator{contract: _RegensToken.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_RegensToken *RegensTokenFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *RegensTokenTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegensToken.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegensTokenTransferBatch)
				if err := _RegensToken.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_RegensToken *RegensTokenFilterer) ParseTransferBatch(log types.Log) (*RegensTokenTransferBatch, error) {
	event := new(RegensTokenTransferBatch)
	if err := _RegensToken.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegensTokenTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the RegensToken contract.
type RegensTokenTransferSingleIterator struct {
	Event *RegensTokenTransferSingle // Event containing the contract specifics and raw log

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
func (it *RegensTokenTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegensTokenTransferSingle)
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
		it.Event = new(RegensTokenTransferSingle)
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
func (it *RegensTokenTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegensTokenTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegensTokenTransferSingle represents a TransferSingle event raised by the RegensToken contract.
type RegensTokenTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_RegensToken *RegensTokenFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*RegensTokenTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegensToken.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RegensTokenTransferSingleIterator{contract: _RegensToken.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_RegensToken *RegensTokenFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *RegensTokenTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RegensToken.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegensTokenTransferSingle)
				if err := _RegensToken.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_RegensToken *RegensTokenFilterer) ParseTransferSingle(log types.Log) (*RegensTokenTransferSingle, error) {
	event := new(RegensTokenTransferSingle)
	if err := _RegensToken.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegensTokenURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the RegensToken contract.
type RegensTokenURIIterator struct {
	Event *RegensTokenURI // Event containing the contract specifics and raw log

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
func (it *RegensTokenURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegensTokenURI)
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
		it.Event = new(RegensTokenURI)
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
func (it *RegensTokenURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegensTokenURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegensTokenURI represents a URI event raised by the RegensToken contract.
type RegensTokenURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_RegensToken *RegensTokenFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*RegensTokenURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegensToken.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &RegensTokenURIIterator{contract: _RegensToken.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_RegensToken *RegensTokenFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *RegensTokenURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RegensToken.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegensTokenURI)
				if err := _RegensToken.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_RegensToken *RegensTokenFilterer) ParseURI(log types.Log) (*RegensTokenURI, error) {
	event := new(RegensTokenURI)
	if err := _RegensToken.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
