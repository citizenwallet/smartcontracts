// Generated code, do not modify. Run `build_runner build` to re-generate!
// @dart=2.12
// ignore_for_file: no_leading_underscores_for_library_prefixes
import 'package:web3dart/web3dart.dart' as _i1;

final _contractAbi = _i1.ContractAbi.fromJson(
  '[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"_owner","type":"address"},{"indexed":true,"internalType":"address","name":"_operator","type":"address"},{"indexed":false,"internalType":"bool","name":"_approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"_operator","type":"address"},{"indexed":true,"internalType":"address","name":"_from","type":"address"},{"indexed":true,"internalType":"address","name":"_to","type":"address"},{"indexed":false,"internalType":"uint256[]","name":"_ids","type":"uint256[]"},{"indexed":false,"internalType":"uint256[]","name":"_values","type":"uint256[]"}],"name":"TransferBatch","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"_operator","type":"address"},{"indexed":true,"internalType":"address","name":"_from","type":"address"},{"indexed":true,"internalType":"address","name":"_to","type":"address"},{"indexed":false,"internalType":"uint256","name":"_id","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"_value","type":"uint256"}],"name":"TransferSingle","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"string","name":"_value","type":"string"},{"indexed":true,"internalType":"uint256","name":"_id","type":"uint256"}],"name":"URI","type":"event"}]',
  'ERC1155',
);

class ERC1155 extends _i1.GeneratedContract {
  ERC1155({
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

  /// Returns a live stream of all ApprovalForAll events emitted by this contract.
  Stream<ApprovalForAll> approvalForAllEvents({
    _i1.BlockNum? fromBlock,
    _i1.BlockNum? toBlock,
  }) {
    final event = self.event('ApprovalForAll');
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
      return ApprovalForAll(
        decoded,
        result,
      );
    });
  }

  /// Returns a live stream of all TransferBatch events emitted by this contract.
  Stream<TransferBatch> transferBatchEvents({
    _i1.BlockNum? fromBlock,
    _i1.BlockNum? toBlock,
  }) {
    final event = self.event('TransferBatch');
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
      return TransferBatch(
        decoded,
        result,
      );
    });
  }

  /// Returns a live stream of all TransferSingle events emitted by this contract.
  Stream<TransferSingle> transferSingleEvents({
    _i1.BlockNum? fromBlock,
    _i1.BlockNum? toBlock,
  }) {
    final event = self.event('TransferSingle');
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
      return TransferSingle(
        decoded,
        result,
      );
    });
  }

  /// Returns a live stream of all URI events emitted by this contract.
  Stream<URI> uRIEvents({
    _i1.BlockNum? fromBlock,
    _i1.BlockNum? toBlock,
  }) {
    final event = self.event('URI');
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
      return URI(
        decoded,
        result,
      );
    });
  }
}

class ApprovalForAll {
  ApprovalForAll(
    List<dynamic> response,
    this.event,
  )   : owner = (response[0] as _i1.EthereumAddress),
        operator = (response[1] as _i1.EthereumAddress),
        approved = (response[2] as bool);

  final _i1.EthereumAddress owner;

  final _i1.EthereumAddress operator;

  final bool approved;

  final _i1.FilterEvent event;
}

class TransferBatch {
  TransferBatch(
    List<dynamic> response,
    this.event,
  )   : operator = (response[0] as _i1.EthereumAddress),
        from = (response[1] as _i1.EthereumAddress),
        to = (response[2] as _i1.EthereumAddress),
        ids = (response[3] as List<dynamic>).cast<BigInt>(),
        values = (response[4] as List<dynamic>).cast<BigInt>();

  final _i1.EthereumAddress operator;

  final _i1.EthereumAddress from;

  final _i1.EthereumAddress to;

  final List<BigInt> ids;

  final List<BigInt> values;

  final _i1.FilterEvent event;
}

class TransferSingle {
  TransferSingle(
    List<dynamic> response,
    this.event,
  )   : operator = (response[0] as _i1.EthereumAddress),
        from = (response[1] as _i1.EthereumAddress),
        to = (response[2] as _i1.EthereumAddress),
        id = (response[3] as BigInt),
        value = (response[4] as BigInt);

  final _i1.EthereumAddress operator;

  final _i1.EthereumAddress from;

  final _i1.EthereumAddress to;

  final BigInt id;

  final BigInt value;

  final _i1.FilterEvent event;
}

class URI {
  URI(
    List<dynamic> response,
    this.event,
  )   : value = (response[0] as String),
        id = (response[1] as BigInt);

  final String value;

  final BigInt id;

  final _i1.FilterEvent event;
}
