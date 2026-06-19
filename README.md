<p align="center">
  <img src="assets/icon.png" alt="Web3 Education Platform" width="128"/>
</p>

# Supervision Trace Education Suite

国内民生溯源**教学 Demo**（子库2）：食品 / 医疗存证 / 慈善流水。

| 插件 | plugin_id |
|------|-----------|
| 食品溯源 | `edu.cn.trace.food` |
| 医疗防篡改 | `edu.cn.trace.medical` |
| 慈善流水 | `edu.cn.trace.charity` |

- Fabric 沙箱 only（`fabric-local`）
- 虚构数据 only
- 任务书：[TASK.md](TASK.md)
- 联合调试：[docs/INTEGRATION.md](docs/INTEGRATION.md)

## 快速开始

```bash
make validate    # manifest 校验 × 3
make smoke       # 联合调试冒烟（自动启后端栈）
make dev-frontend  # 专用 Vue 面板预览 :5174
```

License: PolyForm Noncommercial 1.0.0
