// MedicalTamper — 医疗防篡改存证教学 Chaincode（虚构数据，不对接真实 EMR）
// 演示：病历内容哈希上链 + 篡改检测算法

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type MedicalRecord struct {
	RecordID    string `json:"recordId"`
	PatientRef  string `json:"patientRef"`
	ContentHash string `json:"contentHash"`
	Version     int    `json:"version"`
	Timestamp   string `json:"timestamp"`
	Operator    string `json:"operator"`
}

type TamperCheckResult struct {
	RecordID       string `json:"recordId"`
	StoredHash     string `json:"storedHash"`
	SubmittedHash  string `json:"submittedHash"`
	IsTampered     bool   `json:"isTampered"`
	DetectionNote  string `json:"detectionNote"`
}

type MedicalTamper struct {
	contractapi.Contract
}

func (m *MedicalTamper) InitLedger(ctx contractapi.TransactionContextInterface) error {
	demo := MedicalRecord{
		RecordID: "DEMO-RECORD-001", PatientRef: "虚构患者-XYZ",
		ContentHash: "sha256:demo-medical-hash-abc123",
		Version: 1, Timestamp: "2026-06-01T10:00:00Z", Operator: "教学演示员",
	}
	return ctx.GetStub().PutState("record:"+demo.RecordID, mustJSON(demo))
}

func (m *MedicalTamper) StoreRecordHash(ctx contractapi.TransactionContextInterface, recordID, patientRef, contentHash, timestamp, operator string) error {
	rec := MedicalRecord{
		RecordID: recordID, PatientRef: patientRef, ContentHash: contentHash,
		Version: 1, Timestamp: timestamp, Operator: operator,
	}
	return ctx.GetStub().PutState("record:"+recordID, mustJSON(rec))
}

func (m *MedicalTamper) GetRecord(ctx contractapi.TransactionContextInterface, recordID string) (*MedicalRecord, error) {
	data, err := ctx.GetStub().GetState("record:" + recordID)
	if err != nil || data == nil {
		return nil, fmt.Errorf("record not found: %s", recordID)
	}
	var rec MedicalRecord
	if err := json.Unmarshal(data, &rec); err != nil {
		return nil, err
	}
	return &rec, nil
}

func (m *MedicalTamper) CheckTamper(ctx contractapi.TransactionContextInterface, recordID, submittedHash string) (*TamperCheckResult, error) {
	rec, err := m.GetRecord(ctx, recordID)
	if err != nil {
		return nil, err
	}
	tampered := rec.ContentHash != submittedHash
	note := "哈希一致，未检测到篡改"
	if tampered {
		note = "哈希不一致，触发异常修改检测算法（教学演示）"
	}
	return &TamperCheckResult{
		RecordID: recordID, StoredHash: rec.ContentHash,
		SubmittedHash: submittedHash, IsTampered: tampered, DetectionNote: note,
	}, nil
}

func mustJSON(v any) []byte {
	b, _ := json.Marshal(v)
	return b
}

func main() {
	chaincode, err := contractapi.NewChaincode(&MedicalTamper{})
	if err != nil {
		panic(err)
	}
	if err := chaincode.Start(); err != nil {
		panic(err)
	}
}
