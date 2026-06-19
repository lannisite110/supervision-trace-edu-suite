// CharityLedger — 慈善流水存证教学 Chaincode（模拟募捐，不对接官方支付）
// 演示：捐赠流水哈希上链 + 审计轨迹

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type DonationEntry struct {
	EntryID     string  `json:"entryId"`
	CampaignID  string  `json:"campaignId"`
	DonorRef    string  `json:"donorRef"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	ContentHash string  `json:"contentHash"`
	Timestamp   string  `json:"timestamp"`
}

type LedgerSummary struct {
	CampaignID   string  `json:"campaignId"`
	TotalEntries int     `json:"totalEntries"`
	TotalAmount  float64 `json:"totalAmount"`
	LedgerHash   string  `json:"ledgerHash"`
}

type CharityLedger struct {
	contractapi.Contract
}

func (c *CharityLedger) InitLedger(ctx contractapi.TransactionContextInterface) error {
	entries := []DonationEntry{
		{
			EntryID: "DEMO-ENTRY-001", CampaignID: "DEMO-CAMPAIGN-001",
			DonorRef: "虚构捐赠者A", Amount: 100.0, Currency: "CNY",
			ContentHash: "sha256:donation-demo-001", Timestamp: "2026-06-01T12:00:00Z",
		},
		{
			EntryID: "DEMO-ENTRY-002", CampaignID: "DEMO-CAMPAIGN-001",
			DonorRef: "虚构捐赠者B", Amount: 50.0, Currency: "CNY",
			ContentHash: "sha256:donation-demo-002", Timestamp: "2026-06-02T15:30:00Z",
		},
	}
	for _, e := range entries {
		if err := ctx.GetStub().PutState("entry:"+e.EntryID, mustJSON(e)); err != nil {
			return err
		}
	}
	return nil
}

func (c *CharityLedger) RecordDonation(ctx contractapi.TransactionContextInterface, entryID, campaignID, donorRef string, amount float64, currency, contentHash, timestamp string) error {
	entry := DonationEntry{
		EntryID: entryID, CampaignID: campaignID, DonorRef: donorRef,
		Amount: amount, Currency: currency, ContentHash: contentHash, Timestamp: timestamp,
	}
	return ctx.GetStub().PutState("entry:"+entryID, mustJSON(entry))
}

func (c *CharityLedger) GetEntry(ctx contractapi.TransactionContextInterface, entryID string) (*DonationEntry, error) {
	data, err := ctx.GetStub().GetState("entry:" + entryID)
	if err != nil || data == nil {
		return nil, fmt.Errorf("entry not found: %s", entryID)
	}
	var entry DonationEntry
	if err := json.Unmarshal(data, &entry); err != nil {
		return nil, err
	}
	return &entry, nil
}

func (c *CharityLedger) GetCampaignSummary(ctx contractapi.TransactionContextInterface, campaignID string) (*LedgerSummary, error) {
	iter, err := ctx.GetStub().GetStateByRange("entry:", "entry;")
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	var total float64
	var count int
	for iter.HasNext() {
		kv, err := iter.Next()
		if err != nil {
			return nil, err
		}
		var entry DonationEntry
		if err := json.Unmarshal(kv.Value, &entry); err != nil {
			return nil, err
		}
		if entry.CampaignID == campaignID {
			total += entry.Amount
			count++
		}
	}
	return &LedgerSummary{
		CampaignID: campaignID, TotalEntries: count,
		TotalAmount: total, LedgerHash: "ledger-merkle-root-demo-9b2c",
	}, nil
}

func mustJSON(v any) []byte {
	b, _ := json.Marshal(v)
	return b
}

func main() {
	chaincode, err := contractapi.NewChaincode(&CharityLedger{})
	if err != nil {
		panic(err)
	}
	if err := chaincode.Start(); err != nil {
		panic(err)
	}
}
