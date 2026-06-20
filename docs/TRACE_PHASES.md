# 民生溯源子库 · 分阶段路线图（规划）

> **仓库**: `supervision-trace-edu-suite` · 当前 **v0.3.0** → 目标 **v0.4.0**  
> **插件**: 3（食品 / 医疗 / 慈善）· **Namespace**: `ns-domain-cn` · Fabric 沙箱  
> **参照**: [web3-hot-topic-labs/docs/HOT_TOPIC_PHASES.md](../web3-hot-topic-labs/docs/HOT_TOPIC_PHASES.md)（子库1 已完成 Phase 0–4）

---

## 现状评估（v0.3.0）

| 项 | 状态 |
|----|------|
| 前端 | `FoodTraceLab` 已有三列 grid（溯源链 + Merkle）；医疗/慈善面板较简 |
| 规则 | 3 个 evaluate 有基本逻辑，弱于热点 Phase 3 结构化 hints |
| 合约/Chaincode | Go chaincode 教学版存在，需与 UI 字段对齐 |
| K8s Job | 可能仍为 echo 占位（待 audit） |
| 教程 | 3 篇 intro 偏短，无总路线文档 |

**优势**：食品溯源 Lab 已是平台 UI 标杆，适合作为本子库 Phase 2 模板。

---

## Phase 0 — 共享 composable + Job 审计（预计 1–2 天）

| 任务 | 说明 |
|------|------|
| 抽取/复用 | `useLabSimulate`（可 symlink 或复制到 `plugins/frontend/shared/`） |
| 统一 parseHints | 与热点子库相同 audit_hints 解析 |
| K8s 审计 | 检查 `k8s/overlays/ns-domain-cn/*` 是否 echo；改为 Fabric/toolchain gate |
| 验收 | `make register-plugins` + 3 插件联调 PASS |

---

## Phase 1 — 总入口文档（预计 1 天）

| 任务 | 说明 |
|------|------|
| `TRACE_LEARNING_PATH.md` | 3 插件顺序：food → medical → charity |
| `docs/tutorials/README.md` | 教程索引 |
| 主库 LEARNING_PATH | 新增 **阶段 3B：国内溯源** 交叉链接 |

---

## Phase 2 — 三插件领域 UI 对齐（预计 3–4 天）

| 插件 | 目标 UI |
|------|---------|
| `edu.cn.trace.food` | 维持并增强：批次时间线、Merkle、Fabric 通道示意 |
| `edu.cn.trace.medical` | 病历哈希链、防篡改对比（改前/改后 hash） |
| `edu.cn.trace.charity` | 捐赠流水账本、公开可查字段 vs 隐私字段 |

参考：`FoodTraceLab.vue` 网格布局复制到另两个 Lab。

---

## Phase 3 — Chaincode + 规则 + Job（预计 3–5 天）

| 任务 | 说明 |
|------|------|
| 规则 | params（batch_id、record_hash）驱动 audit_hints |
| Chaincode | 与 manifest 路径一致；Fabric bootstrap 文档链接 |
| Job | `fabric-local` 通道 smoke 或 chaincode 编译 gate |
| 合规 | 教程强调虚构数据、不对接 HIS/监管 |

---

## Phase 4 — 教程扩写 + tag v0.4.0（预计 2 天）

- 3 篇 tutorial 扩写（学习目标 / curl / 自检 / 合规）
- `make tutorial-audit` PASS
- 打 tag **v0.4.0**，主库 CHANGELOG 引用 `TRACE_LEARNING_PATH.md`

---

## 建议时间线

| 周 | 子库 | 里程碑 |
|----|------|--------|
| 已完成 | web3-hot-topic-labs | **v0.4.0** Phase 0–4 ✅ |
| 第 1–2 周 | **supervision-trace-edu-suite** | Phase 0–2 |
| 第 3 周 | supervision-trace-edu-suite | Phase 3–4 → v0.4.0 |
| 第 4–5 周 | enterprise-gov-edu-demo | 政企 3 插件（招投标/多签/供应链） |
| 第 6–7 周 | global-social-edu-sandbox | 海外沙箱 5 插件 |

---

## 启动命令（Trace Phase 0）

```bash
cd ~/web3home/supervision-trace-edu-suite
grep -r "Simulation completed" k8s/ || echo "no echo jobs"
cd ../web3-edu-platform-core
make register-plugins PLUGINS_DIR=..
make integration-all-plugins  # 确认 3 个 trace 插件 PASS
```

---

## 相关文档

- [TASK.md](../TASK.md) · [SUB_REPO_READING_ORDER.md](../../SUB_REPO_READING_ORDER.md) 子库2 章节
- 主库 Fabric：`make fabric-bootstrap`
