# 联合调试指南 · supervision-trace-edu-suite

## 前置条件

- 主库 `web3-edu-platform-core` ≥ v0.1.0
- Go 1.22+、Python 3.12+、Node 20+

## 一键冒烟（推荐）

```bash
cd supervision-trace-edu-suite
chmod +x scripts/joint-debug-smoke.sh
make smoke
```

脚本将自动：注册插件 → 校验 manifest → 合规扫描 → 启动后端栈 → 对 3 个插件执行 evaluate / simulate / status / report → 验证 mainnet 拦截。

## 手动联合调试

### 1. 注册插件

```bash
cd ../web3-edu-platform-core
make register-plugins PLUGINS_DIR=..
```

### 2. 启动后端（三个终端）

```bash
make run-rule-engine   # :8081
make run-scheduler     # :8082
make run-gateway       # :8080
```

### 3a. 主库通用前端

```bash
cd frontend-web && npm run dev   # :5173
```

左侧导航「国内领域」下可见 3 个 Lab（v0.1 使用通用 `LabView`）。

### 3b. 子库专用面板预览（推荐）

```bash
cd supervision-trace-edu-suite
make dev-frontend   # :5174
```

访问：
- http://127.0.0.1:5174/labs/edu.cn.trace.food
- http://127.0.0.1:5174/labs/edu.cn.trace.medical
- http://127.0.0.1:5174/labs/edu.cn.trace.charity

### 4. API 验证

```bash
# 食品溯源
curl -X POST http://127.0.0.1:8080/api/v1/labs/edu.cn.trace.food/simulate \
  -H 'Content-Type: application/json' \
  -d '{"params":{"batch_id":"DEMO-BATCH-001"},"allowed_chain_ids":["fabric-local"]}'
```

## K8s 资源

| 文件 | 说明 |
|------|------|
| `k8s/overlays/ns-domain-cn/fabric-sandbox-configmap.yaml` | Fabric 沙箱共享配置 |
| `k8s/overlays/ns-domain-cn/*-job.yaml` | 各插件 Job 模板 |

## 合规提醒

- `allowed_chain_ids` 使用 `fabric-local`
- 禁止 `target_network: mainnet`（网关返回 403）
- 文档与 UI 不含真实监管/支付/EMR 对接表述
