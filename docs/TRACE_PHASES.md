# 民生溯源子库 · 分阶段路线图（规划）

> **仓库**: `supervision-trace-edu-suite` · **v0.4.0** ✅  
> **插件**: 3（食品 / 医疗 / 慈善）· **Namespace**: `ns-domain-cn` · Fabric 沙箱  
> **参照**: [web3-hot-topic-labs/docs/HOT_TOPIC_PHASES.md](../web3-hot-topic-labs/docs/HOT_TOPIC_PHASES.md)（子库1 已完成 Phase 0–4）

---

## 现状（v0.4.0）

| 项 | 状态 |
|----|------|
| 前端 | 三插件三列 grid + eval-card + Fabric 示意 |
| 规则 | params 驱动 audit_hints（含 tamper / ledger） |
| K8s Job | `peer version` + chaincode 路径 gate |
| 教程 | 3 篇扩写 + [TRACE_LEARNING_PATH.md](TRACE_LEARNING_PATH.md) |
| 主库 | LEARNING_PATH 阶段 3B 已链接 |

---

## Phase 0 — 共享 composable + Job 审计 ✅

| 任务 | 状态 |
|------|------|
| `useLabSimulate` | ✅ |
| `parseHints.ts` | ✅ |
| 三 Lab eval-card | ✅ |
| K8s Job | ✅ |
| 医疗规则 tamper hints | ✅ |
| `scripts/trace-phase0-verify.sh` | ✅ |

---

## Phase 1 — 总入口文档 ✅

| 任务 | 状态 |
|------|------|
| `TRACE_LEARNING_PATH.md` | ✅ food → medical → charity |
| `docs/tutorials/README.md` | ✅ |
| 主库 LEARNING_PATH 3B | ✅ |

---

## Phase 2 — 三插件领域 UI 对齐 ✅

| 插件 | 状态 |
|------|------|
| `edu.cn.trace.food` | ✅ 溯源链 + Merkle + Fabric 沙箱卡片 |
| `edu.cn.trace.medical` | ✅ 篡改对比 + 版本哈希链 |
| `edu.cn.trace.charity` | ✅ 流水表 + 公开/隐私字段 |

---

## Phase 3 — Chaincode + 规则 + Job ✅

| 任务 | 状态 |
|------|------|
| 规则 params hints | ✅ batch / tamper / campaign / entry_count |
| Job chaincode 路径 gate | ✅ `plugins/*/chaincode/*.go` |
| Fabric bootstrap 文档链接 | ✅ TRACE_LEARNING_PATH 第 0 步 |

---

## Phase 4 — 教程扩写 + v0.4.0 ✅

- 3 篇 tutorial 扩写  
- manifest `0.4.0`  
- `make tutorial-audit` / `make smoke` 验收  

**Tag**: `v0.4.0` ✅

---

## P0 收官

| 子库 | 状态 |
|------|------|
| hot / trace / gov / global | ✅ v0.4.0 |
| 主库 | v1.1.0 · LEARNING_PATH 3A–3D |

---

## 相关文档

- [TRACE_LEARNING_PATH.md](TRACE_LEARNING_PATH.md) · [TASK.md](../TASK.md)
- [SUB_REPO_READING_ORDER.md](../../SUB_REPO_READING_ORDER.md) 子库2 章节
