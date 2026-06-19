# 医疗防篡改存证教学 Demo

> **plugin_id**: `edu.cn.trace.medical`  
> **TaskType**: `CN_MEDICAL_TAMPER_DEMO`  
> **链环境**: Fabric 沙箱 `fabric-local`

## 教学目标

演示**病历内容哈希上链**与**异常修改检测算法**的数据结构，帮助理解防篡改存证原理。

## 核心概念

| 概念 | 说明 |
|------|------|
| 内容哈希 | 病历快照的 SHA-256 摘要上链存证 |
| 篡改检测 | 比对链上哈希与提交哈希，触发异常告警 |
| 版本号 | 记录每次合法更新的版本递增 |

## 合规红线

| 禁止 | 替代表述 |
|------|----------|
| 真实医院 EMR/HIS 对接 | 病历哈希存证**数据结构**演示 |
| 等保三级认证宣称 | 等保**概念**教学参考 |

## Fabric 约定

- 通道：`edu-cn-trace-sandbox`
- 组织：`OrgEduDemo`
- Chaincode：`plugins/medical-tamper/chaincode/medical_tamper.go`

## 快速验收

```bash
cd ../web3-edu-platform-core
make validate-plugin MANIFEST=../supervision-trace-edu-suite/plugins/medical-tamper/plugin.manifest.yaml
```

联合调试：

```bash
curl -X POST http://127.0.0.1:8080/api/v1/labs/edu.cn.trace.medical/simulate \
  -H 'Content-Type: application/json' \
  -d '{"params":{"record_id":"DEMO-RECORD-001"},"allowed_chain_ids":["fabric-local"]}'
```

---

> **合规声明**: Fabric 沙箱 `fabric-local` only；全部虚构数据；不对接真实监管、支付或 EMR 系统；禁止主网部署。
