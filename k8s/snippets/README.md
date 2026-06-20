# Fabric 沙箱 Job 门禁

Phase 0 起，溯源三插件 Job 使用 `hyperledger/fabric-tools:2.5` 执行 **真实工具链检查**，不再仅 `echo` 占位。

| 检查项 | 命令 |
|--------|------|
| Peer | `peer version` |
| 通道/组织 | 环境变量 `FABRIC_CHANNEL` / `FABRIC_ORG` |
| Chaincode 引用 | `CHAINCODE_PATH` manifest 路径 |

本地 K8s 冒烟：`cd ../web3-edu-platform-core && make k8s-job-smoke`（busybox 回退见 DEV.md）。

Fabric 引导：`make fabric-bootstrap`（主库）。
