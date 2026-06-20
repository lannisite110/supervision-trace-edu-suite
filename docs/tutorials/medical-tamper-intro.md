# 医疗防篡改存证 · 分步实验

> **插件**: `edu.cn.trace.medical` · **TaskType**: `CN_MEDICAL_TAMPER_DEMO` · **Namespace**: `ns-domain-cn`  
> **Chaincode**: `plugins/medical-tamper/chaincode/medical_tamper.go` · **Job**: `k8s/overlays/ns-domain-cn/medical-tamper-job.yaml`  
> **路线**: [TRACE_LEARNING_PATH.md](../TRACE_LEARNING_PATH.md) 第 2 步

---

## 学习目标

- 理解病历**内容哈希上链**与链下校验的关系  
- 通过修改「待校验哈希」触发 `tamper_detected` hints  
- 浏览**版本哈希链**理解合法更新 vs 异常篡改（教学）  

---

## 前置条件

- 已完成 [food-trace-intro.md](food-trace-intro.md) 或熟悉 Fabric 沙箱基础  
- 主库 backend 已启动  

---

## 背景原理

链上存证的是病历快照的 SHA-256 摘要。校验时将**提交哈希**与**链上哈希**比对；不一致则触发教学用「异常修改检测」告警。不对接真实 EMR/HIS。

---

## 分步实验

### 步骤 1：打开 Lab

`http://127.0.0.1:5173/labs/edu.cn.trace.medical`

### 步骤 2：未篡改场景

1. 记录 ID：`DEMO-RECORD-001`  
2. 待校验哈希保持 `sha256:demo-medical-hash-abc123`  
3. 提交 → 应显示 **哈希一致**  

### 步骤 3：篡改场景

1. 将待校验哈希改为任意值（如 `sha256:tampered-demo`）  
2. 提交 → UI 与规则评估均应显示 **tamper_detected=true**  

### 步骤 4：curl 验证

```bash
curl -s -X POST http://127.0.0.1:8080/api/v1/labs/edu.cn.trace.medical/simulate \
  -H 'Content-Type: application/json' \
  -d '{
    "user_prompt": "病历哈希防篡改存证演示",
    "params": {
      "record_id": "DEMO-RECORD-001",
      "submitted_hash": "sha256:tampered-demo"
    },
    "allowed_chain_ids": ["fabric-local"],
    "task_type": "CN_MEDICAL_TAMPER_DEMO"
  }' | jq '.evaluation.audit_hints'
```

**期望**：hints 含 `tamper_detected=true`、`stored_hash=sha256:demo-medical-hash-abc123`。

### 步骤 5：自检清单

- [ ] 三列 grid：存证 / 检测结果 / 版本哈希链  
- [ ] 链上 vs 提交哈希对比区可见  
- [ ] 合规拒绝：prompt 含「等保三级认证」应被 rule-engine 拦截  

---

## 合规红线

| 禁止 | 替代表述 |
|------|----------|
| 真实 EMR/HIS 对接 | 病历哈希**数据结构**演示 |
| 等保三级认证宣称 | 等保**概念**教学参考 |

---

> **合规声明**: Fabric 沙箱 `fabric-local` only；全部虚构数据；不对接真实监管、支付或 EMR 系统；禁止主网部署。
