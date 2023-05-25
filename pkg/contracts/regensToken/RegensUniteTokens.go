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
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"admins\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"mintCustomVoucher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintGratitude\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200269d3803806200269d8339810160408190526200003491620001c6565b6040805160208101909152600081526200004e81620000d4565b506200005c600033620000e6565b60005b8151811015620000cc57620000b77f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6838381518110620000a357620000a362000298565b6020026020010151620000e660201b60201c565b80620000c381620002ae565b9150506200005f565b505062000431565b6002620000e2828262000365565b5050565b60008281526003602090815260408083206001600160a01b0385168452909152902054620000e2908390839060ff16620000e25760008281526003602090815260408083206001600160a01b03851684529091529020805460ff191660011790556200014f3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b634e487b7160e01b600052604160045260246000fd5b80516001600160a01b0381168114620001c157600080fd5b919050565b60006020808385031215620001da57600080fd5b82516001600160401b0380821115620001f257600080fd5b818501915085601f8301126200020757600080fd5b8151818111156200021c576200021c62000193565b8060051b604051601f19603f8301168101818110858211171562000244576200024462000193565b6040529182528482019250838101850191888311156200026357600080fd5b938501935b828510156200028c576200027c85620001a9565b8452938501939285019262000268565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b600060018201620002cf57634e487b7160e01b600052601160045260246000fd5b5060010190565b600181811c90821680620002eb57607f821691505b6020821081036200030c57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200036057600081815260208120601f850160051c810160208610156200033b5750805b601f850160051c820191505b818110156200035c5782815560010162000347565b5050505b505050565b81516001600160401b0381111562000381576200038162000193565b6200039981620003928454620002d6565b8462000312565b602080601f831160018114620003d15760008415620003b85750858301515b600019600386901b1c1916600185901b1785556200035c565b600085815260208120601f198616915b828110156200040257888601518255948401946001909101908401620003e1565b5085821015620004215787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61225c80620004416000396000f3fe608060405234801561001057600080fd5b50600436106101205760003560e01c80634e1273f4116100ad578063b390c0ab11610071578063b390c0ab14610273578063d539139314610286578063d547741f146102ad578063e985e9c5146102c0578063f242432a146102fc57600080fd5b80634e1273f4146102125780636bd286ca1461023257806391d1485414610245578063a217fddf14610258578063a22cb4651461026057600080fd5b8063248a9ca3116100f4578063248a9ca3146101a357806324f1ae7e146101c65780632eb2c2d6146101d95780632f2ff15d146101ec57806336568abe146101ff57600080fd5b8062fdd58e1461012557806301ffc9a71461014b5780630cc656bd1461016e5780630e89341c14610183575b600080fd5b610138610133366004611726565b61030f565b6040519081526020015b60405180910390f35b61015e610159366004611766565b6103a8565b6040519015158152602001610142565b61018161017c366004611783565b6103b3565b005b61019661019136600461179e565b6103d3565b6040516101429190611807565b6101386101b136600461179e565b60009081526003602052604090206001015490565b6101816101d43660046118db565b6104d2565b6101816101e73660046119c7565b61055a565b6101816101fa366004611a71565b6105a6565b61018161020d366004611a71565b6105cb565b610225610220366004611a9d565b610649565b6040516101429190611ba3565b610181610240366004611bb6565b610773565b61015e610253366004611a71565b610877565b610138600081565b61018161026e366004611c45565b6108a2565b610181610281366004611c81565b6108ad565b6101387f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b6101816102bb366004611a71565b6108b8565b61015e6102ce366004611ca3565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b61018161030a366004611ccd565b6108dd565b60006001600160a01b03831661037f5760405162461bcd60e51b815260206004820152602a60248201527f455243313135353a2061646472657373207a65726f206973206e6f742061207660448201526930b634b21037bbb732b960b11b60648201526084015b60405180910390fd5b506000818152602081815260408083206001600160a01b03861684529091529020545b92915050565b60006103a282610922565b6103d0816002600160405180602001604052806000815250610947565b50565b60008181526004602052604090205460609060ff166104345760405162461bcd60e51b815260206004820152601f60248201527f55524920717565727920666f72206e6f6e6578697374656e7420746f6b656e006044820152606401610376565b6000828152600560205260409020805461044d90611d32565b80601f016020809104026020016040519081016040528092919081815260200182805461047990611d32565b80156104c65780601f1061049b576101008083540402835291602001916104c6565b820191906000526020600020905b8154815290600101906020018083116104a957829003601f168201915b50505050509050919050565b6104fc7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a633610877565b6105485760405162461bcd60e51b815260206004820152601d60248201527f4d7573742068617665206d696e74657220726f6c6520746f206d696e740000006044820152606401610376565b6105558360018484610947565b505050565b6001600160a01b038516331480610576575061057685336102ce565b6105925760405162461bcd60e51b815260040161037690611d6c565b61059f8585858585610971565b5050505050565b6000828152600360205260409020600101546105c181610b54565b6105558383610b5e565b6001600160a01b038116331461063b5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152608401610376565b6106458282610be4565b5050565b606081518351146106ae5760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b6064820152608401610376565b6000835167ffffffffffffffff8111156106ca576106ca61181a565b6040519080825280602002602001820160405280156106f3578160200160208202803683370190505b50905060005b845181101561076b5761073e85828151811061071757610717611dba565b602002602001015185838151811061073157610731611dba565b602002602001015161030f565b82828151811061075057610750611dba565b602090810291909101015261076481611de6565b90506106f9565b509392505050565b61271084116107d75760405162461bcd60e51b815260206004820152602a60248201527f437573746f6d20746f6b656e204944206d75737420626520677265617465722060448201526907468616e2031303030360b41b6064820152608401610376565b60008481526004602052604090205460ff16156108365760405162461bcd60e51b815260206004820181905260248201527f546f6b656e2049442068617320616c7265616479206265656e206d696e7465646044820152606401610376565b61084285858585610947565b6000848152600460209081526040808320805460ff191660011790556005909152902061086f8282611e45565b505050505050565b60009182526003602090815260408084206001600160a01b0393909316845291905290205460ff1690565b610645338383610c4b565b610645338383610d2b565b6000828152600360205260409020600101546108d381610b54565b6105558383610be4565b6001600160a01b0385163314806108f957506108f985336102ce565b6109155760405162461bcd60e51b815260040161037690611d6c565b61059f8585858585610db1565b60006001600160e01b03198216637965db0b60e01b14806103a257506103a282610ee9565b61095384848484610f39565b50506000908152600460205260409020805460ff1916600117905550565b81518351146109d35760405162461bcd60e51b815260206004820152602860248201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206044820152670dad2e6dac2e8c6d60c31b6064820152608401610376565b6001600160a01b0384166109f95760405162461bcd60e51b815260040161037690611f05565b33610a0881878787878761105c565b60005b8451811015610aee576000858281518110610a2857610a28611dba565b602002602001015190506000858381518110610a4657610a46611dba565b602090810291909101810151600084815280835260408082206001600160a01b038e168352909352919091205490915081811015610a965760405162461bcd60e51b815260040161037690611f4a565b6000838152602081815260408083206001600160a01b038e8116855292528083208585039055908b16825281208054849290610ad3908490611f94565b9250508190555050505080610ae790611de6565b9050610a0b565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051610b3e929190611fa7565b60405180910390a461086f81878787878761110b565b6103d08133611266565b610b688282610877565b6106455760008281526003602090815260408083206001600160a01b03851684529091529020805460ff19166001179055610ba03390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b610bee8282610877565b156106455760008281526003602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b816001600160a01b0316836001600160a01b031603610cbe5760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b6064820152608401610376565b6001600160a01b03838116600081815260016020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b81600203610d7b5760405162461bcd60e51b815260206004820181905260248201527f47524154495455444520746f6b656e732063616e6e6f74206265206275726e746044820152606401610376565b610d868383836112bf565b610d90838361030f565b60000361055557506000908152600460205260409020805460ff1916905550565b6001600160a01b038416610dd75760405162461bcd60e51b815260040161037690611f05565b336000610de38561144f565b90506000610df08561144f565b9050610e0083898985858961105c565b6000868152602081815260408083206001600160a01b038c16845290915290205485811015610e415760405162461bcd60e51b815260040161037690611f4a565b6000878152602081815260408083206001600160a01b038d8116855292528083208985039055908a16825281208054889290610e7e908490611f94565b909155505060408051888152602081018890526001600160a01b03808b16928c821692918816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4610ede848a8a8a8a8a61149a565b505050505050505050565b60006001600160e01b03198216636cdb3d1360e11b1480610f1a57506001600160e01b031982166303a24d0760e21b145b806103a257506301ffc9a760e01b6001600160e01b03198316146103a2565b6001600160a01b038416610f995760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736044820152607360f81b6064820152608401610376565b336000610fa58561144f565b90506000610fb28561144f565b9050610fc38360008985858961105c565b6000868152602081815260408083206001600160a01b038b16845290915281208054879290610ff3908490611f94565b909155505060408051878152602081018790526001600160a01b03808a1692600092918716917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a46110538360008989898961149a565b50505050505050565b60005b835181101561105357600084828151811061107c5761107c611dba565b6020026020010151905080600214801561109e57506001600160a01b03871615155b156110fa5760405162461bcd60e51b815260206004820152602660248201527f47524154495455444520746f6b656e732063616e6e6f74206265207472616e7360448201526519995c9c995960d21b6064820152608401610376565b5061110481611de6565b905061105f565b6001600160a01b0384163b1561086f5760405163bc197c8160e01b81526001600160a01b0385169063bc197c819061114f9089908990889088908890600401611fd5565b6020604051808303816000875af192505050801561118a575060408051601f3d908101601f1916820190925261118791810190612033565b60015b61123657611196612050565b806308c379a0036111cf57506111aa61206c565b806111b557506111d1565b8060405162461bcd60e51b81526004016103769190611807565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e2d455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b6064820152608401610376565b6001600160e01b0319811663bc197c8160e01b146110535760405162461bcd60e51b8152600401610376906120f6565b6112708282610877565b6106455761127d81611555565b611288836020611567565b60405160200161129992919061213e565b60408051601f198184030181529082905262461bcd60e51b825261037691600401611807565b6001600160a01b0383166113215760405162461bcd60e51b815260206004820152602360248201527f455243313135353a206275726e2066726f6d20746865207a65726f206164647260448201526265737360e81b6064820152608401610376565b33600061132d8461144f565b9050600061133a8461144f565b905061135a8387600085856040518060200160405280600081525061105c565b6000858152602081815260408083206001600160a01b038a168452909152902054848110156113d75760405162461bcd60e51b8152602060048201526024808201527f455243313135353a206275726e20616d6f756e7420657863656564732062616c604482015263616e636560e01b6064820152608401610376565b6000868152602081815260408083206001600160a01b038b81168086529184528285208a8703905582518b81529384018a90529092908816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4604080516020810190915260009052611053565b6040805160018082528183019092526060916000919060208083019080368337019050509050828160008151811061148957611489611dba565b602090810291909101015292915050565b6001600160a01b0384163b1561086f5760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e61906114de90899089908890889088906004016121b3565b6020604051808303816000875af1925050508015611519575060408051601f3d908101601f1916820190925261151691810190612033565b60015b61152557611196612050565b6001600160e01b0319811663f23a6e6160e01b146110535760405162461bcd60e51b8152600401610376906120f6565b60606103a26001600160a01b03831660145b606060006115768360026121f8565b611581906002611f94565b67ffffffffffffffff8111156115995761159961181a565b6040519080825280601f01601f1916602001820160405280156115c3576020820181803683370190505b509050600360fc1b816000815181106115de576115de611dba565b60200101906001600160f81b031916908160001a905350600f60fb1b8160018151811061160d5761160d611dba565b60200101906001600160f81b031916908160001a90535060006116318460026121f8565b61163c906001611f94565b90505b60018111156116b4576f181899199a1a9b1b9c1cb0b131b232b360811b85600f166010811061167057611670611dba565b1a60f81b82828151811061168657611686611dba565b60200101906001600160f81b031916908160001a90535060049490941c936116ad8161220f565b905061163f565b5083156117035760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610376565b9392505050565b80356001600160a01b038116811461172157600080fd5b919050565b6000806040838503121561173957600080fd5b6117428361170a565b946020939093013593505050565b6001600160e01b0319811681146103d057600080fd5b60006020828403121561177857600080fd5b813561170381611750565b60006020828403121561179557600080fd5b6117038261170a565b6000602082840312156117b057600080fd5b5035919050565b60005b838110156117d25781810151838201526020016117ba565b50506000910152565b600081518084526117f38160208601602086016117b7565b601f01601f19169290920160200192915050565b60208152600061170360208301846117db565b634e487b7160e01b600052604160045260246000fd5b601f8201601f1916810167ffffffffffffffff811182821017156118565761185661181a565b6040525050565b600067ffffffffffffffff8311156118775761187761181a565b60405161188e601f8501601f191660200182611830565b8091508381528484840111156118a357600080fd5b83836020830137600060208583010152509392505050565b600082601f8301126118cc57600080fd5b6117038383356020850161185d565b6000806000606084860312156118f057600080fd5b6118f98461170a565b925060208401359150604084013567ffffffffffffffff81111561191c57600080fd5b611928868287016118bb565b9150509250925092565b600067ffffffffffffffff82111561194c5761194c61181a565b5060051b60200190565b600082601f83011261196757600080fd5b8135602061197482611932565b6040516119818282611830565b83815260059390931b85018201928281019150868411156119a157600080fd5b8286015b848110156119bc57803583529183019183016119a5565b509695505050505050565b600080600080600060a086880312156119df57600080fd5b6119e88661170a565b94506119f66020870161170a565b9350604086013567ffffffffffffffff80821115611a1357600080fd5b611a1f89838a01611956565b94506060880135915080821115611a3557600080fd5b611a4189838a01611956565b93506080880135915080821115611a5757600080fd5b50611a64888289016118bb565b9150509295509295909350565b60008060408385031215611a8457600080fd5b82359150611a946020840161170a565b90509250929050565b60008060408385031215611ab057600080fd5b823567ffffffffffffffff80821115611ac857600080fd5b818501915085601f830112611adc57600080fd5b81356020611ae982611932565b604051611af68282611830565b83815260059390931b8501820192828101915089841115611b1657600080fd5b948201945b83861015611b3b57611b2c8661170a565b82529482019490820190611b1b565b96505086013592505080821115611b5157600080fd5b50611b5e85828601611956565b9150509250929050565b600081518084526020808501945080840160005b83811015611b9857815187529582019590820190600101611b7c565b509495945050505050565b6020815260006117036020830184611b68565b600080600080600060a08688031215611bce57600080fd5b611bd78661170a565b94506020860135935060408601359250606086013567ffffffffffffffff80821115611c0257600080fd5b611c0e89838a016118bb565b93506080880135915080821115611c2457600080fd5b508601601f81018813611c3657600080fd5b611a648882356020840161185d565b60008060408385031215611c5857600080fd5b611c618361170a565b915060208301358015158114611c7657600080fd5b809150509250929050565b60008060408385031215611c9457600080fd5b50508035926020909101359150565b60008060408385031215611cb657600080fd5b611cbf8361170a565b9150611a946020840161170a565b600080600080600060a08688031215611ce557600080fd5b611cee8661170a565b9450611cfc6020870161170a565b93506040860135925060608601359150608086013567ffffffffffffffff811115611d2657600080fd5b611a64888289016118bb565b600181811c90821680611d4657607f821691505b602082108103611d6657634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602e908201527f455243313135353a2063616c6c6572206973206e6f7420746f6b656e206f776e60408201526d195c881bdc88185c1c1c9bdd995960921b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611df857611df8611dd0565b5060010190565b601f82111561055557600081815260208120601f850160051c81016020861015611e265750805b601f850160051c820191505b8181101561086f57828155600101611e32565b815167ffffffffffffffff811115611e5f57611e5f61181a565b611e7381611e6d8454611d32565b84611dff565b602080601f831160018114611ea85760008415611e905750858301515b600019600386901b1c1916600185901b17855561086f565b600085815260208120601f198616915b82811015611ed757888601518255948401946001909101908401611eb8565b5085821015611ef55787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b808201808211156103a2576103a2611dd0565b604081526000611fba6040830185611b68565b8281036020840152611fcc8185611b68565b95945050505050565b6001600160a01b0386811682528516602082015260a06040820181905260009061200190830186611b68565b82810360608401526120138186611b68565b9050828103608084015261202781856117db565b98975050505050505050565b60006020828403121561204557600080fd5b815161170381611750565b600060033d11156120695760046000803e5060005160e01c5b90565b600060443d101561207a5790565b6040516003193d81016004833e81513d67ffffffffffffffff81602484011181841117156120aa57505050505090565b82850191508151818111156120c25750505050505090565b843d87010160208285010111156120dc5750505050505090565b6120eb60208286010187611830565b509095945050505050565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b7f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008152600083516121768160178501602088016117b7565b7001034b99036b4b9b9b4b733903937b6329607d1b60179184019182015283516121a78160288401602088016117b7565b01602801949350505050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a0608082018190526000906121ed908301846117db565b979650505050505050565b80820281158282048414176103a2576103a2611dd0565b60008161221e5761221e611dd0565b50600019019056fea26469706673582212201c138ed6b954a1d81381ec7ea991d11395bdebf31d50f4a13833003e6d5d34b764736f6c63430008140033",
}

// RegensTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use RegensTokenMetaData.ABI instead.
var RegensTokenABI = RegensTokenMetaData.ABI

// RegensTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegensTokenMetaData.Bin instead.
var RegensTokenBin = RegensTokenMetaData.Bin

// DeployRegensToken deploys a new Ethereum contract, binding an instance of RegensToken to it.
func DeployRegensToken(auth *bind.TransactOpts, backend bind.ContractBackend, admins []common.Address) (common.Address, *types.Transaction, *RegensToken, error) {
	parsed, err := RegensTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegensTokenBin), backend, admins)
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
