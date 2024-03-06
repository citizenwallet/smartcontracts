// DERC20.js

module.exports = '60806040523480156200001157600080fd5b506040516200362e3803806200362e833981810160405260608110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660018202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019080838360005b83811015620000c3578082015181840152602081019050620000a6565b50505050905090810190601f168015620000f15780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011557600080fd5b838201915060208201858111156200012c57600080fd5b82518660018202830111640100000000821117156200014a57600080fd5b8083526020830192505050908051906020019080838360005b838110156200018057808201518184015260208101905062000163565b50505050905090810190601f168015620001ae5780820380516001836020036101000a031916815260200191505b5060405260200180519060200190929190505050826040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525060058282828888816003908051906020019062000218929190620005f5565b50806004908051906020019062000231929190620005f5565b506012600560006101000a81548160ff021916908360ff1602179055505050604051806080016040528060528152602001620035dc6052913980519060200120838051906020012083805190602001208330604051602001808681526020018581526020018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200195505050505050604051602081830303815290604052805190602001206007819055505050505050506200030e816200039160201b60201c565b620003326000801b62000326620003af60201b60201c565b6200046660201b60201c565b6200038860405180807f4445504f5349544f525f524f4c45000000000000000000000000000000000000815250600e01905060405180910390206200037c620003af60201b60201c565b6200046660201b60201c565b505050620006a4565b80600560006101000a81548160ff021916908360ff16021790555050565b60003073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156200045c5760606000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050600080369050905073ffffffffffffffffffffffffffffffffffffffff81830151169250505062000460565b3390505b80905090565b6200047882826200047c60201b60201c565b5050565b620004ab81600660008581526020019081526020016000206000016200052060201b620029c61790919060201c565b156200051c57620004c1620003af60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600062000550836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200055860201b60201c565b905092915050565b60006200056c8383620005d260201b60201c565b620005c7578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050620005cc565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200063857805160ff191683800117855562000669565b8280016001018555821562000669579182015b82811115620006685782518255916020019190600101906200064b565b5b5090506200067891906200067c565b5090565b620006a191905b808211156200069d57600081600090555060010162000683565b5090565b90565b612f2880620006b46000396000f3fe6080604052600436106101c65760003560e01c806339509351116100f7578063a217fddf11610095578063ca15c87311610064578063ca15c87314610c80578063cf2c52cb14610ccf578063d547741f14610d75578063dd62ed3e14610dd0576101cd565b8063a217fddf14610b44578063a3b0b5a314610b6f578063a457c2d714610b9a578063a9059cbb14610c0d576101cd565b80638acfcaf7116100d15780638acfcaf7146109915780639010d07c146109bc57806391d1485414610a4157806395d89b4114610ab4576101cd565b8063395093511461088e578063626381a01461090157806370a082311461092c576101cd565b806323b872dd116101645780632e1a7d4d1161013e5780632e1a7d4d1461076c5780632f2ff15d146107a7578063313ce5671461080257806336568abe14610833576101cd565b806323b872dd14610625578063248a9ca3146106b85780632d0335ab14610707576101cd565b80630c53c51c116101a05780630c53c51c146103655780630dd7531a146104da5780630f7e59701461056a57806318160ddd146105fa576101cd565b806306fdde03146101d2578063095ea7b3146102625780630b54817c146102d5576101cd565b366101cd57005b600080fd5b3480156101de57600080fd5b506101e7610e55565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561022757808201518184015260208101905061020c565b50505050905090810190601f1680156102545780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561026e57600080fd5b506102bb6004803603604081101561028557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ef7565b604051808215151515815260200191505060405180910390f35b3480156102e157600080fd5b506102ea610f15565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561032a57808201518184015260208101905061030f565b50505050905090810190601f1680156103575780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61045f600480360360a081101561037b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156103b857600080fd5b8201836020820111156103ca57600080fd5b803590602001918460018302840111640100000000831117156103ec57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019092919080359060200190929190803560ff169060200190929190505050610f4e565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561049f578082015181840152602081019050610484565b50505050905090810190601f1680156104cc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156104e657600080fd5b506104ef6113a4565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561052f578082015181840152602081019050610514565b50505050905090810190601f16801561055c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561057657600080fd5b5061057f6113dd565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105bf5780820151818401526020810190506105a4565b50505050905090810190601f1680156105ec5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561060657600080fd5b5061060f611416565b6040518082815260200191505060405180910390f35b34801561063157600080fd5b5061069e6004803603606081101561064857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611420565b604051808215151515815260200191505060405180910390f35b3480156106c457600080fd5b506106f1600480360360208110156106db57600080fd5b81019080803590602001909291905050506114f9565b6040518082815260200191505060405180910390f35b34801561071357600080fd5b506107566004803603602081101561072a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611519565b6040518082815260200191505060405180910390f35b34801561077857600080fd5b506107a56004803603602081101561078f57600080fd5b8101908080359060200190929190505050611562565b005b3480156107b357600080fd5b50610800600480360360408110156107ca57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611576565b005b34801561080e57600080fd5b50610817611600565b604051808260ff1660ff16815260200191505060405180910390f35b34801561083f57600080fd5b5061088c6004803603604081101561085657600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611617565b005b34801561089a57600080fd5b506108e7600480360360408110156108b157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506116b0565b604051808215151515815260200191505060405180910390f35b34801561090d57600080fd5b50610916611763565b6040518082815260200191505060405180910390f35b34801561093857600080fd5b5061097b6004803603602081101561094f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611769565b6040518082815260200191505060405180910390f35b34801561099d57600080fd5b506109a66117b1565b6040518082815260200191505060405180910390f35b3480156109c857600080fd5b506109ff600480360360408110156109df57600080fd5b8101908080359060200190929190803590602001909291905050506117b6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610a4d57600080fd5b50610a9a60048036036040811015610a6457600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506117e8565b604051808215151515815260200191505060405180910390f35b348015610ac057600080fd5b50610ac961181a565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610b09578082015181840152602081019050610aee565b50505050905090810190601f168015610b365780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610b5057600080fd5b50610b596118bc565b6040518082815260200191505060405180910390f35b348015610b7b57600080fd5b50610b846118c3565b6040518082815260200191505060405180910390f35b348015610ba657600080fd5b50610bf360048036036040811015610bbd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506118fc565b604051808215151515815260200191505060405180910390f35b348015610c1957600080fd5b50610c6660048036036040811015610c3057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506119c9565b604051808215151515815260200191505060405180910390f35b348015610c8c57600080fd5b50610cb960048036036020811015610ca357600080fd5b81019080803590602001909291905050506119e7565b6040518082815260200191505060405180910390f35b348015610cdb57600080fd5b50610d7360048036036040811015610cf257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190640100000000811115610d2f57600080fd5b820183602082011115610d4157600080fd5b80359060200191846001830284011164010000000083111715610d6357600080fd5b9091929391929390505050611a0e565b005b348015610d8157600080fd5b50610dce60048036036040811015610d9857600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611ae0565b005b348015610ddc57600080fd5b50610e3f60048036036040811015610df357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611b6a565b6040518082815260200191505060405180910390f35b606060038054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610eed5780601f10610ec257610100808354040283529160200191610eed565b820191906000526020600020905b815481529060010190602001808311610ed057829003601f168201915b5050505050905090565b6000610f0b610f04611bf1565b8484611ca6565b6001905092915050565b6040518060400160405280600281526020017f3a9900000000000000000000000000000000000000000000000000000000000081525081565b6060610f58612c3f565b6040518060600160405280600860008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481526020018873ffffffffffffffffffffffffffffffffffffffff168152602001878152509050610fd78782878787611e9d565b61102c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180612e146021913960400191505060405180910390fd5b61107f6001600860008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611f4f90919063ffffffff16565b600860008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b873388604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561118a57808201518184015260208101905061116f565b50505050905090810190601f1680156111b75780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a1600060603073ffffffffffffffffffffffffffffffffffffffff16888a6040516020018083805190602001908083835b6020831061121957805182526020820191506020810190506020830392506111f6565b6001836020036101000a0380198251168184511680821785525050505050509050018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401925050506040516020818303038152906040526040518082805190602001908083835b602083106112b65780518252602082019150602081019050602083039250611293565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611318576040519150601f19603f3d011682016040523d82523d6000602084013e61131d565b606091505b509150915081611395576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f46756e6374696f6e2063616c6c206e6f74207375636365737366756c0000000081525060200191505060405180910390fd5b80935050505095945050505050565b6040518060400160405280600181526020017f050000000000000000000000000000000000000000000000000000000000000081525081565b6040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525081565b6000600254905090565b600061142d848484611fd7565b6114ee84611439611bf1565b6114e985604051806060016040528060288152602001612dec60289139600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600061149f611bf1565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546122989092919063ffffffff16565b611ca6565b600190509392505050565b600060066000838152602001908152602001600020600201549050919050565b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b61157361156d611bf1565b82612358565b50565b61159d6006600084815260200190815260200160002060020154611598611bf1565b6117e8565b6115f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f815260200180612cbc602f913960400191505060405180910390fd5b6115fc828261251c565b5050565b6000600560009054906101000a900460ff16905090565b61161f611bf1565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146116a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f815260200180612ec4602f913960400191505060405180910390fd5b6116ac82826125b0565b5050565b60006117596116bd611bf1565b8461175485600160006116ce611bf1565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611f4f90919063ffffffff16565b611ca6565b6001905092915050565b613a9981565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600581565b60006117e0826006600086815260200190815260200160002060000161264490919063ffffffff16565b905092915050565b6000611812826006600086815260200190815260200160002060000161265e90919063ffffffff16565b905092915050565b606060048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156118b25780601f10611887576101008083540402835291602001916118b2565b820191906000526020600020905b81548152906001019060200180831161189557829003601f168201915b5050505050905090565b6000801b81565b60405180807f4445504f5349544f525f524f4c45000000000000000000000000000000000000815250600e019050604051809103902081565b60006119bf611909611bf1565b846119ba85604051806060016040528060258152602001612e9f6025913960016000611933611bf1565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546122989092919063ffffffff16565b611ca6565b6001905092915050565b60006119dd6119d6611bf1565b8484611fd7565b6001905092915050565b6000611a076006600084815260200190815260200160002060000161268e565b9050919050565b60405180807f4445504f5349544f525f524f4c45000000000000000000000000000000000000815250600e0190506040518091039020611a5581611a50611bf1565b6117e8565b611aaa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180612d986024913960400191505060405180910390fd5b600083836020811015611abc57600080fd5b81019080803590602001909291905050509050611ad985826126a3565b5050505050565b611b076006600084815260200190815260200160002060020154611b02611bf1565b6117e8565b611b5c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526030815260200180612dbc6030913960400191505060405180910390fd5b611b6682826125b0565b5050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60003073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415611c9c5760606000368080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050600080369050905073ffffffffffffffffffffffffffffffffffffffff818301511692505050611ca0565b3390505b80905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611d2c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180612e7b6024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611db2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180612d506022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b60006001611eb2611ead8761286a565b612910565b83868660405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015611f0c573d6000803e3d6000fd5b5050506020604051035173ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614905095945050505050565b600080828401905083811015611fcd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561205d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180612e566025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156120e3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180612c996023913960400191505060405180910390fd5b6120ee838383612977565b61215981604051806060016040528060268152602001612d72602691396000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546122989092919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506121ec816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611f4f90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000838311158290612345576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561230a5780820151818401526020810190506122ef565b50505050905090810190601f1680156123375780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156123de576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180612e356021913960400191505060405180910390fd5b6123ea82600083612977565b61245581604051806060016040528060228152602001612ceb602291396000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546122989092919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506124ac8160025461297c90919063ffffffff16565b600281905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b61254481600660008581526020019081526020016000206000016129c690919063ffffffff16565b156125ac57612551611bf1565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6125d881600660008581526020019081526020016000206000016129f690919063ffffffff16565b15612640576125e5611bf1565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b60006126538360000183612a26565b60001c905092915050565b6000612686836000018373ffffffffffffffffffffffffffffffffffffffff1660001b612aa9565b905092915050565b600061269c82600001612acc565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415612746576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b61275260008383612977565b61276781600254611f4f90919063ffffffff16565b6002819055506127be816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611f4f90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b6000604051806080016040528060438152602001612d0d604391398051906020012082600001518360200151846040015180519060200120604051602001808581526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b600061291a612add565b8260405160200180807f190100000000000000000000000000000000000000000000000000000000000081525060020183815260200182815260200192505050604051602081830303815290604052805190602001209050919050565b505050565b60006129be83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250612298565b905092915050565b60006129ee836000018373ffffffffffffffffffffffffffffffffffffffff1660001b612ae7565b905092915050565b6000612a1e836000018373ffffffffffffffffffffffffffffffffffffffff1660001b612b57565b905092915050565b600081836000018054905011612a87576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180612c776022913960400191505060405180910390fd5b826000018281548110612a9657fe5b9060005260206000200154905092915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081600001805490509050919050565b6000600754905090565b6000612af38383612aa9565b612b4c578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050612b51565b600090505b92915050565b60008083600101600084815260200190815260200160002054905060008114612c335760006001820390506000600186600001805490500390506000866000018281548110612ba257fe5b9060005260206000200154905080876000018481548110612bbf57fe5b9060005260206000200181905550600183018760010160008381526020019081526020016000208190555086600001805480612bf757fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050612c39565b60009150505b92915050565b604051806060016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152509056fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e647345524332303a207472616e7366657220746f20746865207a65726f2061646472657373416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f206772616e7445524332303a206275726e20616d6f756e7420657863656564732062616c616e63654d6574615472616e73616374696f6e2875696e74323536206e6f6e63652c616464726573732066726f6d2c62797465732066756e6374696f6e5369676e61747572652945524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e63654368696c6445524332303a20494e53554646494349454e545f5045524d495353494f4e53416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f207265766f6b6545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e63655369676e657220616e64207369676e617475726520646f206e6f74206d6174636845524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636520726f6c657320666f722073656c66a2646970667358221220cf5ae10ce4835c90d8edfbf7420ad78b008e2a0dd9dbbfd8bba9f476382c243c64736f6c63430006060033454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e7432353620636861696e49642c6164647265737320766572696679696e67436f6e747261637429000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000b44756d6d7920455243323000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000064445524332300000000000000000000000000000000000000000000000000000';