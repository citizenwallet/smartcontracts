// Generated code, do not modify. Run `build_runner build` to re-generate!
// @dart=2.12
// ignore_for_file: no_leading_underscores_for_library_prefixes
import 'package:web3dart/web3dart.dart' as _i1;

final _contractAbi = _i1.ContractAbi.fromJson(
  '[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"faucet","type":"address"},{"indexed":false,"internalType":"address","name":"token","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"},{"indexed":false,"internalType":"uint48","name":"redeemInterval","type":"uint48"}],"name":"RedeemCodeFaucetCreated","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"faucet","type":"address"},{"indexed":false,"internalType":"address","name":"token","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"},{"indexed":false,"internalType":"uint48","name":"redeemInterval","type":"uint48"}],"name":"SimpleFaucetCreated","type":"event"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"uint256","name":"salt","type":"uint256"},{"internalType":"contract IERC20Upgradeable","name":"_token","type":"address"},{"internalType":"uint48","name":"_redeemInterval","type":"uint48"},{"internalType":"address","name":"_codeCreator","type":"address"}],"name":"createRedeemCodeFaucet","outputs":[{"internalType":"contract RedeemCodeFaucet","name":"ret","type":"address"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"uint256","name":"salt","type":"uint256"},{"internalType":"contract IERC20Upgradeable","name":"_token","type":"address"},{"internalType":"uint256","name":"_amount","type":"uint256"},{"internalType":"uint48","name":"_redeemInterval","type":"uint48"},{"internalType":"address","name":"_redeemAdmin","type":"address"}],"name":"createSimpleFaucet","outputs":[{"internalType":"contract SimpleFaucet","name":"ret","type":"address"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"uint256","name":"salt","type":"uint256"},{"internalType":"contract IERC20Upgradeable","name":"_token","type":"address"},{"internalType":"uint48","name":"_redeemInterval","type":"uint48"},{"internalType":"address","name":"_codeCreator","type":"address"}],"name":"getRedeemCodeFaucetAddress","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"uint256","name":"salt","type":"uint256"},{"internalType":"contract IERC20Upgradeable","name":"_token","type":"address"},{"internalType":"uint256","name":"_amount","type":"uint256"},{"internalType":"uint48","name":"_redeemInterval","type":"uint48"},{"internalType":"address","name":"_redeemAdmin","type":"address"}],"name":"getSimpleFaucetAddress","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"redeemCodeFaucetImplementation","outputs":[{"internalType":"contract RedeemCodeFaucet","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"simpleFaucetImplementation","outputs":[{"internalType":"contract SimpleFaucet","name":"","type":"address"}],"stateMutability":"view","type":"function"}]',
  'FaucetFactory',
);

class FaucetFactory extends _i1.GeneratedContract {
  FaucetFactory({
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
  Future<String> createRedeemCodeFaucet(
    _i1.EthereumAddress owner,
    BigInt salt,
    _i1.EthereumAddress _token,
    BigInt _redeemInterval,
    _i1.EthereumAddress _codeCreator, {
    required _i1.Credentials credentials,
    _i1.Transaction? transaction,
  }) async {
    final function = self.abi.functions[1];
    assert(checkSignature(function, '4a40c132'));
    final params = [
      owner,
      salt,
      _token,
      _redeemInterval,
      _codeCreator,
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
  Future<String> createSimpleFaucet(
    _i1.EthereumAddress owner,
    BigInt salt,
    _i1.EthereumAddress _token,
    BigInt _amount,
    BigInt _redeemInterval,
    _i1.EthereumAddress _redeemAdmin, {
    required _i1.Credentials credentials,
    _i1.Transaction? transaction,
  }) async {
    final function = self.abi.functions[2];
    assert(checkSignature(function, 'c8a66cb0'));
    final params = [
      owner,
      salt,
      _token,
      _amount,
      _redeemInterval,
      _redeemAdmin,
    ];
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
  Future<_i1.EthereumAddress> getRedeemCodeFaucetAddress(
    _i1.EthereumAddress owner,
    BigInt salt,
    _i1.EthereumAddress _token,
    BigInt _redeemInterval,
    _i1.EthereumAddress _codeCreator, {
    _i1.BlockNum? atBlock,
  }) async {
    final function = self.abi.functions[3];
    assert(checkSignature(function, '65895758'));
    final params = [
      owner,
      salt,
      _token,
      _redeemInterval,
      _codeCreator,
    ];
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
  Future<_i1.EthereumAddress> getSimpleFaucetAddress(
    _i1.EthereumAddress owner,
    BigInt salt,
    _i1.EthereumAddress _token,
    BigInt _amount,
    BigInt _redeemInterval,
    _i1.EthereumAddress _redeemAdmin, {
    _i1.BlockNum? atBlock,
  }) async {
    final function = self.abi.functions[4];
    assert(checkSignature(function, '09e5223e'));
    final params = [
      owner,
      salt,
      _token,
      _amount,
      _redeemInterval,
      _redeemAdmin,
    ];
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
  Future<_i1.EthereumAddress> redeemCodeFaucetImplementation(
      {_i1.BlockNum? atBlock}) async {
    final function = self.abi.functions[5];
    assert(checkSignature(function, '76c40417'));
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
  Future<_i1.EthereumAddress> simpleFaucetImplementation(
      {_i1.BlockNum? atBlock}) async {
    final function = self.abi.functions[6];
    assert(checkSignature(function, '6057b63d'));
    final params = [];
    final response = await read(
      function,
      params,
      atBlock,
    );
    return (response[0] as _i1.EthereumAddress);
  }

  /// Returns a live stream of all RedeemCodeFaucetCreated events emitted by this contract.
  Stream<RedeemCodeFaucetCreated> redeemCodeFaucetCreatedEvents({
    _i1.BlockNum? fromBlock,
    _i1.BlockNum? toBlock,
  }) {
    final event = self.event('RedeemCodeFaucetCreated');
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
      return RedeemCodeFaucetCreated(
        decoded,
        result,
      );
    });
  }

  /// Returns a live stream of all SimpleFaucetCreated events emitted by this contract.
  Stream<SimpleFaucetCreated> simpleFaucetCreatedEvents({
    _i1.BlockNum? fromBlock,
    _i1.BlockNum? toBlock,
  }) {
    final event = self.event('SimpleFaucetCreated');
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
      return SimpleFaucetCreated(
        decoded,
        result,
      );
    });
  }
}

class RedeemCodeFaucetCreated {
  RedeemCodeFaucetCreated(
    List<dynamic> response,
    this.event,
  )   : faucet = (response[0] as _i1.EthereumAddress),
        token = (response[1] as _i1.EthereumAddress),
        amount = (response[2] as BigInt),
        redeemInterval = (response[3] as BigInt);

  final _i1.EthereumAddress faucet;

  final _i1.EthereumAddress token;

  final BigInt amount;

  final BigInt redeemInterval;

  final _i1.FilterEvent event;
}

class SimpleFaucetCreated {
  SimpleFaucetCreated(
    List<dynamic> response,
    this.event,
  )   : faucet = (response[0] as _i1.EthereumAddress),
        token = (response[1] as _i1.EthereumAddress),
        amount = (response[2] as BigInt),
        redeemInterval = (response[3] as BigInt);

  final _i1.EthereumAddress faucet;

  final _i1.EthereumAddress token;

  final BigInt amount;

  final BigInt redeemInterval;

  final _i1.FilterEvent event;
}
