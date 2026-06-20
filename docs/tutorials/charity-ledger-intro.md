# 慈善流水存证 · 分步实验

> **插件**: `edu.cn.trace.charity` · **TaskType**: `CN_CHARITY_LEDGER_DEMO` · **Namespace**: `ns-domain-cn`  
> **Chaincode**: `plugins/charity-ledger/chaincode/charity_ledger.go` · **Job**: `k8s/overlays/ns-domain-cn/charity-ledger-job.yaml`  
> **路线**: [TRACE_LEARNING_PATH.md](../TRACE_LEARNING_PATH.md) 第 3 步

---

## 学习目标

- 理解模拟募捐**流水哈希上链**与活动账本 Merkle 根  
- 区分**公开字段**与**隐私字段**（hash 上链）  
- 阅读 `campaign_id` / `entry_count` / `ledger_hash` audit_hints  

---

## 前置条件

- 主库 backend 已启动  
- 建议顺序：food → medical → charity（见 [TRACE_LEARNING_PATH.md](../TRACE_LEARNING_PATH.md)）  

---

## 背景原理

每笔虚构捐赠生成独立内容哈希；同一 `campaign_id` 下流水聚合为账本根哈希。**不对接**微信/支付宝等官方支付渠道。

---

## 分步实验

### 步骤 1：打开 Lab

`http://127.0.0.1:5173/labs/edu.cn.trace.charity`

### 步骤 2：UI 操作

1. 活动 ID：`DEMO-CAMPAIGN-001`  
2. 浏览捐赠流水表与 **公开 vs 隐私字段** 对照表  
3. 提交存证实验  

### 步骤 3：curl 验证

```bash
curl -s -X POST http://127.0.0.1:8080/api/v1/labs/edu.cn.trace.charity/simulate \
  -H 'Content-Type: application/json' \
  -d '{
    "user_prompt": "模拟募捐流水哈希上链",
    "params": {
      "campaign_id": "DEMO-CAMPAIGN-001",
      "entry_count": 2,
      "ledger_hash": "ledger-merkle-root-demo-9b2c"
    },
    "allowed_chain_ids": ["fabric-local"],
    "task_type": "CN_CHARITY_LEDGER_DEMO"
  }' | jq '.evaluation'
```

**期望**：`compliance_passed: true`，hints 含 `entry_count=2`。

### 步骤 4：合规负例

```bash
curl -s -X POST http://127.0.0.1:8080/api/v1/labs/edu.cn.trace.charity/simulate \
  -H 'Content-Type: application/json' \
  -d '{
    "user_prompt": "对接微信支付真实募捐",
    "params": {},
    "allowed_chain_ids": ["fabric-local"]
  }' | jq '.evaluation.compliance_passed'
```

**期望**：`false`。

### 步骤 5：自检清单

- [ ] 三列 grid：活动 / 流水 / 字段策略  
- [ ] eval-card 显示 entry_count  
- [ ] 任务 completed + report JSON  

---

## 合规说明

- 虚构捐赠者代号，不含真实 PII  
- 禁止官方募捐支付对接表述  

---

> **合规声明**: Fabric 沙箱 `fabric-local` only；全部虚构数据；不对接真实监管、支付或 EMR 系统；禁止主网部署。
