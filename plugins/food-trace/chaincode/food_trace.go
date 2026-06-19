// FoodTrace — 食品溯源教学 Chaincode（虚构数据，Fabric 沙箱）
// 通道: edu-cn-trace-sandbox | 组织: OrgEduDemo | chainId: fabric-local

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type BatchRecord struct {
	BatchID     string `json:"batchId"`
	ProductName string `json:"productName"`
	Origin      string `json:"origin"`
	Stage       string `json:"stage"`
	ContentHash string `json:"contentHash"`
	Timestamp   string `json:"timestamp"`
}

type MerkleProof struct {
	BatchID    string   `json:"batchId"`
	RootHash   string   `json:"rootHash"`
	LeafHashes []string `json:"leafHashes"`
	ProofPath  []string `json:"proofPath"`
}

type FoodTrace struct {
	contractapi.Contract
}

func (f *FoodTrace) InitLedger(ctx contractapi.TransactionContextInterface) error {
	demo := []BatchRecord{
		{
			BatchID: "DEMO-BATCH-001", ProductName: "有机蔬菜礼盒",
			Origin: "虚构农场A", Stage: "采收", ContentHash: "a3f2c1…demo",
			Timestamp: "2026-06-01T08:00:00Z",
		},
		{
			BatchID: "DEMO-BATCH-001", ProductName: "有机蔬菜礼盒",
			Origin: "虚构冷链中心", Stage: "冷链运输", ContentHash: "b4e3d2…demo",
			Timestamp: "2026-06-02T14:30:00Z",
		},
		{
			BatchID: "DEMO-BATCH-001", ProductName: "有机蔬菜礼盒",
			Origin: "虚构零售门店", Stage: "上架销售", ContentHash: "c5f4e3…demo",
			Timestamp: "2026-06-03T09:15:00Z",
		},
	}
	for i, rec := range demo {
		key := fmt.Sprintf("batch:%s:%d", rec.BatchID, i)
		if err := ctx.GetStub().PutState(key, mustJSON(rec)); err != nil {
			return err
		}
	}
	return nil
}

func (f *FoodTrace) CreateBatch(ctx contractapi.TransactionContextInterface, batchID, productName, origin, stage, contentHash, timestamp string) error {
	rec := BatchRecord{
		BatchID: batchID, ProductName: productName, Origin: origin,
		Stage: stage, ContentHash: contentHash, Timestamp: timestamp,
	}
	key := fmt.Sprintf("batch:%s:%s", batchID, stage)
	return ctx.GetStub().PutState(key, mustJSON(rec))
}

func (f *FoodTrace) GetBatchHistory(ctx contractapi.TransactionContextInterface, batchID string) ([]BatchRecord, error) {
	iter, err := ctx.GetStub().GetStateByRange(fmt.Sprintf("batch:%s:", batchID), fmt.Sprintf("batch:%s;", batchID))
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	var records []BatchRecord
	for iter.HasNext() {
		kv, err := iter.Next()
		if err != nil {
			return nil, err
		}
		var rec BatchRecord
		if err := json.Unmarshal(kv.Value, &rec); err != nil {
			return nil, err
		}
		records = append(records, rec)
	}
	return records, nil
}

func (f *FoodTrace) GetMerkleProof(ctx contractapi.TransactionContextInterface, batchID string) (*MerkleProof, error) {
	records, err := f.GetBatchHistory(ctx, batchID)
	if err != nil {
		return nil, err
	}
	leaves := make([]string, len(records))
	for i, r := range records {
		leaves[i] = r.ContentHash
	}
	return &MerkleProof{
		BatchID: batchID, RootHash: "merkle-root-demo-7f3a",
		LeafHashes: leaves, ProofPath: []string{"hash→parent→root (教学简化)"},
	}, nil
}

func mustJSON(v any) []byte {
	b, _ := json.Marshal(v)
	return b
}

func main() {
	chaincode, err := contractapi.NewChaincode(&FoodTrace{})
	if err != nil {
		panic(err)
	}
	if err := chaincode.Start(); err != nil {
		panic(err)
	}
}
