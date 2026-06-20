<p align="center">
  <img src="assets/icon.png" alt="Web3 Education Platform" width="128"/>
</p>

# Supervision Trace Education Suite

**v0.4.0** · 国内民生溯源**教学 Demo**（子库2）· 主库 LEARNING_PATH **阶段 3B**

食品 / 医疗防篡改 / 慈善流水 · Fabric 沙箱 only。

## 文档

- [完整学习路径](docs/TRACE_LEARNING_PATH.md) · [教程索引](docs/tutorials/README.md)
- [学习路径（子库入口）](docs/LEARNING_PATH.md) · [快速部署](docs/QUICK_DEPLOY.md)
- [分阶段工程路线](docs/TRACE_PHASES.md)

依赖 [web3-edu-platform-core](https://github.com/lannisite110/web3-edu-platform-core) **v1.1.0**。

| 插件 | plugin_id | TaskType |
|------|-----------|----------|
| 食品溯源 | `edu.cn.trace.food` | `CN_FOOD_TRACE_SIM` |
| 医疗防篡改 | `edu.cn.trace.medical` | `CN_MEDICAL_TAMPER_DEMO` |
| 慈善流水 | `edu.cn.trace.charity` | `CN_CHARITY_LEDGER_DEMO` |

- Fabric 沙箱 only（`fabric-local`）· 虚构数据 only
- 变更记录：[CHANGELOG.md](CHANGELOG.md) · 任务书：[TASK.md](TASK.md)

## 快速开始

```bash
make validate    # manifest 校验 × 3
make smoke       # 联合调试冒烟（自动启后端栈）
make dev-frontend  # 专用 Vue 面板预览 :5174
```

主库注册：

```bash
cd ../web3-edu-platform-core
make register-plugins PLUGINS_DIR=..
make tutorial-audit PLUGINS_DIR=..
```

License: PolyForm Noncommercial 1.0.0
