// Generated code, do not modify. Run `build_runner build` to re-generate!
// @dart=2.12
// ignore_for_file: no_leading_underscores_for_library_prefixes
import 'package:web3dart/web3dart.dart' as _i1;

final _contractAbi = _i1.ContractAbi.fromJson(
  '[{"inputs":[{"internalType":"contract IEntryPoint","name":"anEntryPoint","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint8","name":"version","type":"uint8"}],"name":"Initialized","type":"event"},{"inputs":[{"internalType":"string","name":"name","type":"string"},{"internalType":"address","name":"link","type":"address"},{"internalType":"string","name":"description","type":"string"},{"internalType":"string","name":"linkType","type":"string"},{"internalType":"string","name":"meta","type":"string"}],"name":"addLink","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"clearLinks","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"clearProfile","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"entryPoint","outputs":[{"internalType":"contract IEntryPoint","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getLinks","outputs":[{"components":[{"internalType":"string","name":"name","type":"string"},{"internalType":"address","name":"link","type":"address"},{"internalType":"string","name":"description","type":"string"},{"internalType":"string","name":"linkType","type":"string"},{"internalType":"string","name":"meta","type":"string"}],"internalType":"struct Linkable.LinkData[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getProfile","outputs":[{"internalType":"string","name":"","type":"string"},{"internalType":"string","name":"","type":"string"},{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"anOwner","type":"address"}],"name":"initialize","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"links","outputs":[{"internalType":"string","name":"name","type":"string"},{"internalType":"address","name":"link","type":"address"},{"internalType":"string","name":"description","type":"string"},{"internalType":"string","name":"linkType","type":"string"},{"internalType":"string","name":"meta","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"profile","outputs":[{"internalType":"string","name":"name","type":"string"},{"internalType":"string","name":"description","type":"string"},{"internalType":"string","name":"meta","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"index","type":"uint256"}],"name":"removeLink","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"index","type":"uint256"},{"internalType":"string","name":"name","type":"string"},{"internalType":"address","name":"link","type":"address"},{"internalType":"string","name":"description","type":"string"},{"internalType":"string","name":"linkType","type":"string"},{"internalType":"string","name":"meta","type":"string"}],"name":"updateLink","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"name","type":"string"},{"internalType":"string","name":"description","type":"string"},{"internalType":"string","name":"meta","type":"string"}],"name":"updateProfile","outputs":[],"stateMutability":"nonpayable","type":"function"}]',
  'Profile',
);

class Profile extends _i1.GeneratedContract {
  Profile({
    required _i1.EthereumAddress address,
    required _i1.Web3Client client,
    int? chainId,
  }) : super(
          _i1.DeployedContract(
            _contractAbi,
            address,
          ),
          client,
          chainId,
        );

  /// The optional [transaction] parameter can be used to override parameters
  /// like the gas price, nonce and max gas. The `data` and `to` fields will be
  /// set by the contract.
  Future<String> addLink(
    String name,
    _i1.EthereumAddress link,
    String description,
    String linkType,
    String meta, {
    required _i1.Credentials credentials,
    _i1.Transaction? transaction,
  }) async {
    final function = self.abi.functions[1];
    assert(checkSignature(function, '6510c584'));
    final params = [
      name,
      link,
      description,
      linkType,
      meta,
    ];
    return write(
      credentials,
      transaction,
      function,
      params,
    );
  }

  /// The optional [transaction] parameter can be used to override parameters
  /// like the gas price, nonce and max gas. The `data` and `to` fields will be
  /// set by the contract.
  Future<String> clearLinks({
    required _i1.Credentials credentials,
    _i1.Transaction? transaction,
  }) async {
    final function = self.abi.functions[2];
    assert(checkSignature(function, '75ca64fe'));
    final params = [];
    return write(
      credentials,
      transaction,
      function,
      params,
    );
  }

  /// The optional [transaction] parameter can be used to override parameters
  /// like the gas price, nonce and max gas. The `data` and `to` fields will be
  /// set by the contract.
  Future<String> clearProfile({
    required _i1.Credentials credentials,
    _i1.Transaction? transaction,
  }) async {
    final function = self.abi.functions[3];
    assert(checkSignature(function, '8ccec7b6'));
    final params = [];
    return write(
      credentials,
      transaction,
      function,
      params,
    );
  }

  /// The optional [atBlock] parameter can be used to view historical data. When
  /// set, the function will be evaluated in the specified block. By default, the
  /// latest on-chain block will be used.
  Future<_i1.EthereumAddress> entryPoint({_i1.BlockNum? atBlock}) async {
    final function = self.abi.functions[4];
    assert(checkSignature(function, 'b0d691fe'));
    final params = [];
    final response = await read(
      function,
      params,
      atBlock,
    );
    return (response[0] as _i1.EthereumAddress);
  }

  /// The optional [atBlock] parameter can be used to view historical data. When
  /// set, the function will be evaluated in the specified block. By default, the
  /// latest on-chain block will be used.
  Future<List<dynamic>> getLinks({_i1.BlockNum? atBlock}) async {
    final function = self.abi.functions[5];
    assert(checkSignature(function, 'f66ddcbb'));
    final params = [];
    final response = await read(
      function,
      params,
      atBlock,
    );
    return (response[0] as List<dynamic>).cast<dynamic>();
  }

  /// The optional [atBlock] parameter can be used to view historical data. When
  /// set, the function will be evaluated in the specified block. By default, the
  /// latest on-chain block will be used.
  Future<GetProfile> getProfile({_i1.BlockNum? atBlock}) async {
    final function = self.abi.functions[6];
    assert(checkSignature(function, 'd6afc9b1'));
    final params = [];
    final response = await read(
      function,
      params,
      atBlock,
    );
    return GetProfile(response);
  }

  /// The optional [transaction] parameter can be used to override parameters
  /// like the gas price, nonce and max gas. The `data` and `to` fields will be
  /// set by the contract.
  Future<String> initialize(
    _i1.EthereumAddress anOwner, {
    required _i1.Credentials credentials,
    _i1.Transaction? transaction,
  }) async {
    final function = self.abi.functions[7];
    assert(checkSignature(function, 'c4d66de8'));
    final params = [anOwner];
    return write(
      credentials,
      transaction,
      function,
      params,
    );
  }

  /// The optional [atBlock] parameter can be used to view historical data. When
  /// set, the function will be evaluated in the specified block. By default, the
  /// latest on-chain block will be used.
  Future<Links> links(
    BigInt $param6, {
    _i1.BlockNum? atBlock,
  }) async {
    final function = self.abi.functions[8];
    assert(checkSignature(function, '881d8a40'));
    final params = [$param6];
    final response = await read(
      function,
      params,
      atBlock,
    );
    return Links(response);
  }

  /// The optional [atBlock] parameter can be used to view historical data. When
  /// set, the function will be evaluated in the specified block. By default, the
  /// latest on-chain block will be used.
  Future<_i1.EthereumAddress> owner({_i1.BlockNum? atBlock}) async {
    final function = self.abi.functions[9];
    assert(checkSignature(function, '8da5cb5b'));
    final params = [];
    final response = await read(
      function,
      params,
      atBlock,
    );
    return (response[0] as _i1.EthereumAddress);
  }

  /// The optional [atBlock] parameter can be used to view historical data. When
  /// set, the function will be evaluated in the specified block. By default, the
  /// latest on-chain block will be used.
  Future<Profile> profile({_i1.BlockNum? atBlock}) async {
    final function = self.abi.functions[10];
    assert(checkSignature(function, 'ab60636c'));
    final params = [];
    final response = await read(
      function,
      params,
      atBlock,
    );
    return Profile(response);
  }

  /// The optional [transaction] parameter can be used to override parameters
  /// like the gas price, nonce and max gas. The `data` and `to` fields will be
  /// set by the contract.
  Future<String> removeLink(
    BigInt index, {
    required _i1.Credentials credentials,
    _i1.Transaction? transaction,
  }) async {
    final function = self.abi.functions[11];
    assert(checkSignature(function, 'b080ae4f'));
    final params = [index];
    return write(
      credentials,
      transaction,
      function,
      params,
    );
  }

  /// The optional [transaction] parameter can be used to override parameters
  /// like the gas price, nonce and max gas. The `data` and `to` fields will be
  /// set by the contract.
  Future<String> updateLink(
    BigInt index,
    String name,
    _i1.EthereumAddress link,
    String description,
    String linkType,
    String meta, {
    required _i1.Credentials credentials,
    _i1.Transaction? transaction,
  }) async {
    final function = self.abi.functions[12];
    assert(checkSignature(function, '4bd13de6'));
    final params = [
      index,
      name,
      link,
      description,
      linkType,
      meta,
    ];
    return write(
      credentials,
      transaction,
      function,
      params,
    );
  }

  /// The optional [transaction] parameter can be used to override parameters
  /// like the gas price, nonce and max gas. The `data` and `to` fields will be
  /// set by the contract.
  Future<String> updateProfile(
    String name,
    String description,
    String meta, {
    required _i1.Credentials credentials,
    _i1.Transaction? transaction,
  }) async {
    final function = self.abi.functions[13];
    assert(checkSignature(function, '1105a5eb'));
    final params = [
      name,
      description,
      meta,
    ];
    return write(
      credentials,
      transaction,
      function,
      params,
    );
  }

  /// Returns a live stream of all Initialized events emitted by this contract.
  Stream<Initialized> initializedEvents({
    _i1.BlockNum? fromBlock,
    _i1.BlockNum? toBlock,
  }) {
    final event = self.event('Initialized');
    final filter = _i1.FilterOptions.events(
      contract: self,
      event: event,
      fromBlock: fromBlock,
      toBlock: toBlock,
    );
    return client.events(filter).map((_i1.FilterEvent result) {
      final decoded = event.decodeResults(
        result.topics!,
        result.data!,
      );
      return Initialized(
        decoded,
        result,
      );
    });
  }
}

class GetProfile {
  GetProfile(List<dynamic> response)
      : var1 = (response[0] as String),
        var2 = (response[1] as String),
        var3 = (response[2] as String);

  final String var1;

  final String var2;

  final String var3;
}

class Links {
  Links(List<dynamic> response)
      : name = (response[0] as String),
        link = (response[1] as _i1.EthereumAddress),
        description = (response[2] as String),
        linkType = (response[3] as String),
        meta = (response[4] as String);

  final String name;

  final _i1.EthereumAddress link;

  final String description;

  final String linkType;

  final String meta;
}

class Profile {
  Profile(List<dynamic> response)
      : name = (response[0] as String),
        description = (response[1] as String),
        meta = (response[2] as String);

  final String name;

  final String description;

  final String meta;
}

class Initialized {
  Initialized(
    List<dynamic> response,
    this.event,
  ) : version = (response[0] as BigInt);

  final BigInt version;

  final _i1.FilterEvent event;
}
