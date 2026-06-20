# 食品溯源教学 · 分步实验

> **插件**: `edu.cn.trace.food` · **TaskType**: `CN_FOOD_TRACE_SIM` · **Namespace**: `ns-domain-cn`  
> **Chaincode**: `plugins/food-trace/chaincode/food_trace.go` · **Job**: `k8s/overlays/ns-domain-cn/food-trace-job.yaml`  
> **路线**: [TRACE_LEARNING_PATH.md](../TRACE_LEARNING_PATH.md) 第 1 步

---

## 学习目标

- 理解食品批次溯源的**虚构链路**（采收 → 冷链 → 零售）  
- 使用 UI 提交 simulate，阅读 `batch_id` / `merkle_proof_enabled` 等 audit_hints  
- 了解 Fabric 沙箱通道 `edu-cn-trace-sandbox` 与 `fabric-local` chainId  

---

## 前置条件

- 主库四服务已启动（见 [QUICK_DEPLOY.md](../../web3-edu-platform-core/docs/QUICK_DEPLOY.md)）  
- `make register-plugins PLUGINS_DIR=..` 已完成  

---

## 背景原理

`food_trace.go` 教学 Chaincode 维护 `BatchRecord` 与 `MerkleProof` 结构。本平台在 **教学层** 展示时间线与 Merkle 根，不连接真实监管或支付系统。

---

## 分步实验

### 步骤 1：打开 Lab

浏览器访问：`http://127.0.0.1:5173/labs/edu.cn.trace.food`

### 步骤 2：UI 操作

1. 确认批次 ID 为 `DEMO-BATCH-001`（可修改）  
2. 浏览 **溯源链路** 与 **Merkle 证明** 三列面板  
3. 查看 **Fabric 沙箱** 卡片（通道 / 组织 / chaincode 路径）  
4. 点击 **提交仿真实验**  

### 步骤 3：curl 验证

```bash
curl -s -X POST http://127.0.0.1:8080/api/v1/labs/edu.cn.trace.food/simulate \
  -H 'Content-Type: application/json' \
  -d '{
    "user_prompt": "食品批次溯源教学演示",
    "params": {
      "batch_id": "DEMO-BATCH-001",
      "product_name": "有机蔬菜礼盒"
    },
    "allowed_chain_ids": ["fabric-local"],
    "task_type": "CN_FOOD_TRACE_SIM"
  }' | jq '.evaluation'
```

**期望**：`compliance_passed: true`，audit_hints 含 `batch_id=DEMO-BATCH-001`、`channel=edu-cn-trace-sandbox`。

### 步骤 4：自检清单

- [ ] 规则评估卡片显示合规通过  
- [ ] 任务状态变为 `completed`  
- [ ] 修改 batch_id 后 hints 同步更新  

---

## K8s Job（可选）

Job 运行 `peer version` 并校验 chaincode 路径格式：

```bash
cd ../../web3-edu-platform-core
make fabric-bootstrap
make k8s-job-smoke
```

---

## 合规说明

- 全部使用虚构农场/冷链/门店数据  
- 不对接真实监管平台或支付系统  
- 禁止 `target_network=mainnet`  

---

> **合规声明**: Fabric 沙箱 `fabric-local` only；全部虚构数据；不对接真实监管、支付或 EMR 系统；禁止主网部署。
