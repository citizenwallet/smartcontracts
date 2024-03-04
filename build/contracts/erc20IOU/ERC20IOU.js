// ERC20IOU.js

module.exports = '60a06040523060805234801561001457600080fd5b5060805161176b61004c6000396000818161023b01528181610284015281816103230152818161036301526103f6015261176b6000f3fe60806040526004361061009c5760003560e01c8063715018a611610064578063715018a61461016a5780638da5cb5b1461017f578063c4d66de8146101b1578063f2fde38b146101d1578063fae0d7dc146101f1578063fc0c546a1461021157600080fd5b80633659cfe6146100a15780633d8b9e07146100c35780634f1ef2861461011457806352d1902d14610127578063536f3fff1461014a575b600080fd5b3480156100ad57600080fd5b506100c16100bc36600461126a565b610231565b005b3480156100cf57600080fd5b506100f86100de366004611287565b60ca6020526000908152604090205465ffffffffffff1681565b60405165ffffffffffff90911681526020015b60405180910390f35b6100c16101223660046112b6565b610319565b34801561013357600080fd5b5061013c6103e9565b60405190815260200161010b565b34801561015657600080fd5b506100c1610165366004611395565b61049c565b34801561017657600080fd5b506100c161084c565b34801561018b57600080fd5b506033546001600160a01b03165b6040516001600160a01b03909116815260200161010b565b3480156101bd57600080fd5b506100c16101cc36600461126a565b610860565b3480156101dd57600080fd5b506100c16101ec36600461126a565b610994565b3480156101fd57600080fd5b5061013c61020c36600461144f565b610a0a565b34801561021d57600080fd5b5060c954610199906001600160a01b031681565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036102825760405162461bcd60e51b8152600401610279906114a6565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166102cb6000805160206116cf833981519152546001600160a01b031690565b6001600160a01b0316146102f15760405162461bcd60e51b8152600401610279906114f2565b6102fa81610a88565b6040805160008082526020820190925261031691839190610a90565b50565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036103615760405162461bcd60e51b8152600401610279906114a6565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166103aa6000805160206116cf833981519152546001600160a01b031690565b6001600160a01b0316146103d05760405162461bcd60e51b8152600401610279906114f2565b6103d982610a88565b6103e582826001610a90565b5050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104895760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610279565b506000805160206116cf83398151915290565b4265ffffffffffff808616908216108015906104c757508565ffffffffffff168165ffffffffffff16105b61051e5760405162461bcd60e51b815260206004820152602260248201527f4552433230494f553a2065787069726564206f72206e6f742076616c69642079604482015261195d60f21b6064820152608401610279565b600061052d8989898989610a0a565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c810191909152605c0160408051601f198184030181529181528151602092830120600081815260ca90935291205490915065ffffffffffff16156105de5760405162461bcd60e51b815260206004820152601a60248201527f4552433230494f553a20616c72656164792072656465656d65640000000000006044820152606401610279565b883b156106bd57604051630b135d3f60e11b815289906001600160a01b03821690631626ba7e906106179085908990899060040161153e565b602060405180830381865afa158015610634573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106589190611574565b6001600160e01b031916631626ba7e60e01b146106b75760405162461bcd60e51b815260206004820152601b60248201527f4552433230494f553a20696e76616c6964207369676e617475726500000000006044820152606401610279565b5061075d565b886001600160a01b03166107078286868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610c0092505050565b6001600160a01b03161461075d5760405162461bcd60e51b815260206004820152601b60248201527f4552433230494f553a20696e76616c6964207369676e617475726500000000006044820152606401610279565b60c9546040516323b872dd60e01b81526001600160a01b038b81166004830152336024830152604482018b9052909116906323b872dd906064016020604051808303816000875af11580156107b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107da919061159e565b50600081815260ca6020908152604091829020805465ffffffffffff191665ffffffffffff8616179055905189815233916001600160a01b038c16917fd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9910160405180910390a3505050505050505050565b610854610e67565b61085e6000610ec1565b565b600054610100900460ff16158080156108805750600054600160ff909116105b8061089a5750303b15801561089a575060005460ff166001145b6108fd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610279565b6000805460ff191660011790558015610920576000805461ff0019166101001790555b610928610f13565b610930610f42565b60c980546001600160a01b0319166001600160a01b03841617905580156103e5576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b61099c610e67565b6001600160a01b038116610a015760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610279565b61031681610ec1565b6040516bffffffffffffffffffffffff19606087811b82166020840152603483018790526001600160d01b031960d087811b8216605486015286901b16605a84015280830184905246608084015230901b1660a082015260009060b40160405160208183030381529060405280519060200120905095945050505050565b610316610e67565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610ac857610ac383610f69565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610b22575060408051601f3d908101601f19168201909252610b1f918101906115c0565b60015b610b855760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610279565b6000805160206116cf8339815191528114610bf45760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610279565b50610ac3838383611005565b60008151604114610c675760405162461bcd60e51b815260206004820152603a602482015260008051602061171683398151915260448201527f3a20696e76616c6964207369676e6174757265206c656e6774680000000000006064820152608401610279565b600082604081518110610c7c57610c7c6115d9565b016020015160f81c90506000610c928482611030565b90506000610ca1856020611030565b90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0811115610d275760405162461bcd60e51b815260206004820152603d602482015260008051602061171683398151915260448201527f3a20696e76616c6964207369676e6174757265202773272076616c75650000006064820152608401610279565b8260ff16601b14158015610d3f57508260ff16601c14155b15610da05760405162461bcd60e51b815260206004820152603d602482015260008051602061171683398151915260448201527f3a20696e76616c6964207369676e6174757265202776272076616c75650000006064820152608401610279565b60408051600081526020810180835288905260ff851691810191909152606081018390526080810182905260019060a0016020604051602081039080840390855afa158015610df3573d6000803e3d6000fd5b5050604051601f1901519450506001600160a01b038416610e5d5760405162461bcd60e51b8152602060048201526030602482015260008051602061171683398151915260448201526f1d1024a72b20a624a22fa9a4a3a722a960811b6064820152608401610279565b5050505b92915050565b6033546001600160a01b0316331461085e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610279565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16610f3a5760405162461bcd60e51b8152600401610279906115ef565b61085e611096565b600054610100900460ff1661085e5760405162461bcd60e51b8152600401610279906115ef565b6001600160a01b0381163b610fd65760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610279565b6000805160206116cf83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b61100e836110c6565b60008251118061101b5750805b15610ac35761102a8383611106565b50505050565b600061103d82602061163a565b8351101561108d5760405162461bcd60e51b815260206004820181905260248201527f72656164427974657333323a20696e76616c69642064617461206c656e6774686044820152606401610279565b50016020015190565b600054610100900460ff166110bd5760405162461bcd60e51b8152600401610279906115ef565b61085e33610ec1565b6110cf81610f69565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b606061112b83836040518060600160405280602781526020016116ef60279139611132565b9392505050565b6060600080856001600160a01b03168560405161114f919061167f565b600060405180830381855af49150503d806000811461118a576040519150601f19603f3d011682016040523d82523d6000602084013e61118f565b606091505b50915091506111a0868383876111aa565b9695505050505050565b60608315611219578251600003611212576001600160a01b0385163b6112125760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610279565b5081611223565b611223838361122b565b949350505050565b81511561123b5781518083602001fd5b8060405162461bcd60e51b8152600401610279919061169b565b6001600160a01b038116811461031657600080fd5b60006020828403121561127c57600080fd5b813561112b81611255565b60006020828403121561129957600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156112c957600080fd5b82356112d481611255565b9150602083013567ffffffffffffffff808211156112f157600080fd5b818501915085601f83011261130557600080fd5b813581811115611317576113176112a0565b604051601f8201601f19908116603f0116810190838211818310171561133f5761133f6112a0565b8160405282815288602084870101111561135857600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b803565ffffffffffff8116811461139057600080fd5b919050565b600080600080600080600060c0888a0312156113b057600080fd5b87356113bb81611255565b9650602088013595506113d06040890161137a565b94506113de6060890161137a565b93506080880135925060a088013567ffffffffffffffff8082111561140257600080fd5b818a0191508a601f83011261141657600080fd5b81358181111561142557600080fd5b8b602082850101111561143757600080fd5b60208301945080935050505092959891949750929550565b600080600080600060a0868803121561146757600080fd5b853561147281611255565b9450602086013593506114876040870161137a565b92506114956060870161137a565b949793965091946080013592915050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b83815260406020820152816040820152818360608301376000818301606090810191909152601f909201601f1916010192915050565b60006020828403121561158657600080fd5b81516001600160e01b03198116811461112b57600080fd5b6000602082840312156115b057600080fd5b8151801515811461112b57600080fd5b6000602082840312156115d257600080fd5b5051919050565b634e487b7160e01b600052603260045260246000fd5b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b80820180821115610e6157634e487b7160e01b600052601160045260246000fd5b60005b8381101561167657818101518382015260200161165e565b50506000910152565b6000825161169181846020870161165b565b9190910192915050565b60208152600082518060208401526116ba81604085016020870161165b565b601f01601f1916919091016040019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645369676e617475726556616c696461746f72237265636f7665725369676e6572a26469706673582212201a241b3cc8d8bd39f776d7853968b359ecae0407e91e553a2ea4022501b8d00d64736f6c63430008140033';
