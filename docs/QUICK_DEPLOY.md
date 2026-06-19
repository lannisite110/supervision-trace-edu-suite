# 快速部署 · Supervision Trace Edu Suite

```bash
# 五仓 clone（见主库 QUICK_DEPLOY）
cd ~/web3home/web3-edu-platform-core
make register-plugins PLUGINS_DIR=..
make fabric-bootstrap    # 可选 Fabric 沙箱
make ci-gate
```

本仓插件 ID：`edu.cn.trace.food` · `edu.cn.trace.charity` · `edu.cn.trace.medical`

主库文档：[QUICK_DEPLOY](https://github.com/lannisite110/web3-edu-platform-core/blob/main/docs/QUICK_DEPLOY.md)
