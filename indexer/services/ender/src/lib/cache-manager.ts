import {
  IsolationLevel,
  Transaction,
  assetRefresher,
  marketRefresher,
  perpetualMarketRefresher,
} from '@dydxprotocol-indexer/postgres';

export async function refreshDataCaches(): Promise<void> {
  const txId: number = await Transaction.start();
  await Transaction.setIsolationLevel(txId, IsolationLevel.READ_COMMITTED);

  await Promise.all([
    perpetualMarketRefresher.updatePerpetualMarkets({ txId, readReplica: true }),
    assetRefresher.updateAssets({ txId, readReplica: true }),
    marketRefresher.updateMarkets({ txId, readReplica: true }),
  ]);

  await Transaction.rollback(txId);
}
