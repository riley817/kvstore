package main

import (
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

type KVStoreApplication struct{}

var _ abcitypes.Application = (*KVStoreApplication)(nil)

func NewKVStoreApplication() *KVStoreApplication {
	return &KVStoreApplication{}
}

func (K KVStoreApplication) Info(info abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{}
}

func (K KVStoreApplication) Query(query abcitypes.RequestQuery) abcitypes.ResponseQuery {
	return abcitypes.ResponseQuery{Code: 0}
}

func (K KVStoreApplication) CheckTx(tx abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	return abcitypes.ResponseCheckTx{Code: 0}
}

func (K KVStoreApplication) InitChain(chain abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	return abcitypes.ResponseInitChain{}
}

func (K KVStoreApplication) BeginBlock(block abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	return abcitypes.ResponseBeginBlock{}
}

func (K KVStoreApplication) DeliverTx(tx abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	return abcitypes.ResponseDeliverTx{Code: 0}
}

func (K KVStoreApplication) EndBlock(block abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	return abcitypes.ResponseEndBlock{}
}

func (K KVStoreApplication) Commit() abcitypes.ResponseCommit {
	return abcitypes.ResponseCommit{}
}

func (K KVStoreApplication) ListSnapshots(snapshots abcitypes.RequestListSnapshots) abcitypes.ResponseListSnapshots {
	return abcitypes.ResponseListSnapshots{}
}

func (K KVStoreApplication) OfferSnapshot(snapshot abcitypes.RequestOfferSnapshot) abcitypes.ResponseOfferSnapshot {
	return abcitypes.ResponseOfferSnapshot{}
}

func (K KVStoreApplication) LoadSnapshotChunk(chunk abcitypes.RequestLoadSnapshotChunk) abcitypes.ResponseLoadSnapshotChunk {
	return abcitypes.ResponseLoadSnapshotChunk{}
}

func (K KVStoreApplication) ApplySnapshotChunk(chunk abcitypes.RequestApplySnapshotChunk) abcitypes.ResponseApplySnapshotChunk {
	return abcitypes.ResponseApplySnapshotChunk{}
}
