# Agent-2 任务书 · supervision-trace-edu-suite

> **仓库**: `supervision-trace-edu-suite`  
> **compliance_tier**: `cn_domain`  
> **Namespace**: `ns-domain-cn`（Fabric 沙箱通道）

---

## 1. 目标

3 个**国内民生存证教学 Demo**（虚构数据，不对接真实监管/支付/医院 HIS）。

| # | 模块 | plugin_id | TaskType |
|---|------|-----------|----------|
| 1 | 食品溯源 | `edu.cn.trace.food` | `CN_FOOD_TRACE_SIM` |
| 2 | 医疗防篡改存证 | `edu.cn.trace.medical` | `CN_MEDICAL_TAMPER_DEMO` |
| 3 | 慈善流水存证 | `edu.cn.trace.charity` | `CN_CHARITY_LEDGER_DEMO` |

---

## 2. 技术栈

- **链**: Hyperledger Fabric 本地沙箱（`fabric-local`），不用公链主网
- **合约**: Chaincode Go（教学简化版）
- **前端**: Vue 面板 3 个，展示批次/哈希/Merkle 证明

---

## 3. 目录结构

```
supervision-trace-edu-suite/
├── plugins/
│   ├── food-trace/
│   ├── medical-tamper/
│   └── charity-ledger/
├── k8s/overlays/ns-domain-cn/
└── docs/tutorials/
```

每个子目录含完整 `plugin.manifest.yaml` + chaincode + rules + frontend。

---

## 4. 合规红线

| 禁止 | 替代表述 |
|------|----------|
| 真实医院 EMR 对接 | 「病历哈希存证**数据结构**演示」 |
| 等保三级认证宣称 | 「等保**概念**教学参考」 |
| 对接官方募捐支付 | 「模拟募捐流水哈希上链」 |
| 纪检/反腐产品化描述 | 「异常修改检测**算法**演示」 |

---

## 5. Fabric 约定

- 通道名: `edu-cn-trace-sandbox`
- 组织: `OrgEduDemo`（单组织教学）
- chainId 字段填: `fabric-local`

---

## 6. 验收

```bash
cd ../web3-edu-platform-core
for m in ../supervision-trace-edu-suite/plugins/*/plugin.manifest.yaml; do
  make validate-plugin MANIFEST="$m"
done
```
