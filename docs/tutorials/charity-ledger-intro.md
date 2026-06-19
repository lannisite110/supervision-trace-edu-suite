# 慈善流水存证教学 Demo

> **plugin_id**: `edu.cn.trace.charity`  
> **TaskType**: `CN_CHARITY_LEDGER_DEMO`  
> **链环境**: Fabric 沙箱 `fabric-local`

## 教学目标

演示**模拟募捐流水哈希上链**与审计轨迹的数据结构，帮助理解公益资金透明存证原理。

## 核心概念

| 概念 | 说明 |
|------|------|
| 捐赠流水 | 每笔虚构捐赠生成独立内容哈希 |
| 活动汇总 | 同一 campaign 下流水聚合为账本 Merkle 根 |
| 审计轨迹 | 链上不可篡改的时间戳与哈希记录 |

## 合规红线

| 禁止 | 替代表述 |
|------|----------|
| 对接官方募捐支付 | 模拟募捐流水哈希上链 |
| 真实捐赠者 PII | 虚构捐赠者代号（如「虚构捐赠者A」） |

## Fabric 约定

- 通道：`edu-cn-trace-sandbox`
- 组织：`OrgEduDemo`
- Chaincode：`plugins/charity-ledger/chaincode/charity_ledger.go`

## 快速验收

```bash
cd ../web3-edu-platform-core
make validate-plugin MANIFEST=../supervision-trace-edu-suite/plugins/charity-ledger/plugin.manifest.yaml
```

联合调试：

```bash
curl -X POST http://127.0.0.1:8080/api/v1/labs/edu.cn.trace.charity/simulate \
  -H 'Content-Type: application/json' \
  -d '{"params":{"campaign_id":"DEMO-CAMPAIGN-001"},"allowed_chain_ids":["fabric-local"]}'
```
