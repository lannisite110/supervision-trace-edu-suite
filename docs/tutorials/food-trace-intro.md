# 食品溯源教学 Demo

> **plugin_id**: `edu.cn.trace.food`  
> **TaskType**: `CN_FOOD_TRACE_SIM`  
> **链环境**: Fabric 沙箱 `fabric-local`

## 教学目标

演示食品批次从生产到零售的**虚构溯源链路**，以及 Merkle 证明的**数据结构**。

## 核心概念

| 概念 | 说明 |
|------|------|
| 批次 ID | 贯穿全链路的唯一标识（虚构：`DEMO-BATCH-001`） |
| 内容哈希 | 每个环节对批次快照计算的 SHA-256 摘要 |
| Merkle 证明 | 多环节哈希聚合为根哈希，用于完整性校验教学 |

## Fabric 约定

- 通道：`edu-cn-trace-sandbox`
- 组织：`OrgEduDemo`（单组织教学）
- Chaincode：`plugins/food-trace/chaincode/food_trace.go`

## 合规说明

- 全部使用虚构农场/冷链/门店数据
- 不对接真实监管平台或支付系统
- 禁止主网部署

## 快速验收

```bash
cd ../web3-edu-platform-core
make validate-plugin MANIFEST=../supervision-trace-edu-suite/plugins/food-trace/plugin.manifest.yaml
```

联合调试：

```bash
curl -X POST http://127.0.0.1:8080/api/v1/labs/edu.cn.trace.food/simulate \
  -H 'Content-Type: application/json' \
  -d '{"params":{"batch_id":"DEMO-BATCH-001"},"allowed_chain_ids":["fabric-local"]}'
```
