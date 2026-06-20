# 国内民生溯源 · 完整学习路径

> **子库** `supervision-trace-edu-suite` v0.4.0 · 主库 ≥ v1.1.0 · **3 插件**  
> **Namespace**: `ns-domain-cn` · **链**: Fabric 沙箱 `fabric-local`  
> 交叉引用：主库 [LEARNING_PATH.md](../web3-edu-platform-core/docs/LEARNING_PATH.md) **阶段 3B**

---

## 路线总览（约 1 周业余 / 3 天全职）

```text
food-trace（食品溯源 · 入口）
    → medical-tamper（医疗防篡改）
    → charity-ledger（慈善流水存证）
```

三插件共用同一 Fabric 教学通道 `edu-cn-trace-sandbox` / 组织 `OrgEduDemo`。

---

## 第 0 步：环境与注册（所有 Lab 共用）

```bash
cd ~/web3home/web3-edu-platform-core
make register-plugins PLUGINS_DIR=..
make run-rule-engine & make run-scheduler & make run-gateway &
cd frontend-web && npm run dev
# → http://127.0.0.1:5173
```

Fabric 沙箱（可选 K8s）：

```bash
make fabric-bootstrap    # 应用 ns-domain-cn ConfigMap
make k8s-job-smoke       # 需 kubectl
```

详见主库 [QUICK_DEPLOY.md](../web3-edu-platform-core/docs/QUICK_DEPLOY.md) 与子库 [QUICK_DEPLOY.md](QUICK_DEPLOY.md)。

---

## 第 1 步：食品溯源（必做 · 1–2 天）

| 项 | 内容 |
|----|------|
| 插件 | `edu.cn.trace.food` |
| 教程 | [tutorials/food-trace-intro.md](tutorials/food-trace-intro.md) |
| TaskType | `CN_FOOD_TRACE_SIM` |
| UI | 批次输入 · 溯源时间线 · Merkle 证明 · Fabric 通道示意 |
| Chaincode | `plugins/food-trace/chaincode/food_trace.go` |

**自检**：修改 `batch_id` → 规则评估卡片显示新批次 → 任务 `completed`。

---

## 第 2 步：医疗防篡改（1 天）

| 项 | 内容 |
|----|------|
| 插件 | `edu.cn.trace.medical` |
| 教程 | [tutorials/medical-tamper-intro.md](tutorials/medical-tamper-intro.md) |
| TaskType | `CN_MEDICAL_TAMPER_DEMO` |
| UI | 链上哈希 vs 待校验哈希 · 篡改检测 · 版本哈希链 |
| Chaincode | `plugins/medical-tamper/chaincode/medical_tamper.go` |

**自检**：将待校验哈希改为任意值 → UI 与 `tamper_detected=true` hints 一致。

---

## 第 3 步：慈善流水存证（1 天）

| 项 | 内容 |
|----|------|
| 插件 | `edu.cn.trace.charity` |
| 教程 | [tutorials/charity-ledger-intro.md](tutorials/charity-ledger-intro.md) |
| TaskType | `CN_CHARITY_LEDGER_DEMO` |
| UI | 活动 ID · 捐赠流水表 · 公开/隐私字段对照 |
| Chaincode | `plugins/charity-ledger/chaincode/charity_ledger.go` |

**自检**：修改 `campaign_id` → eval-card 更新 → 合规拒绝「官方支付对接」类 prompt。

---

## 教程索引

完整列表见 [tutorials/README.md](tutorials/README.md)。

---

## 工程化验收

```bash
cd ~/web3home/supervision-trace-edu-suite
make validate && make smoke

cd ../web3-edu-platform-core
make tutorial-audit PLUGINS_DIR=..
make integration-all-plugins   # 含 3 trace 插件
```

分阶段工程路线：[TRACE_PHASES.md](TRACE_PHASES.md)

---

## 合规速查

| 禁止 | 替代表述 |
|------|----------|
| 真实 EMR/HIS | 病历哈希**数据结构**演示 |
| 等保三级认证 | 等保**概念**教学参考 |
| 官方募捐支付 | 模拟流水哈希上链 |
| 主网 / 真实监管对接 | `fabric-local` 沙箱 only |

全平台：[COMPLIANCE_MASTER.md](../../COMPLIANCE_MASTER.md)
