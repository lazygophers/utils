# xtime996

> 早上 9 点到晚上 9 点，一周 6 天

## 时间常量定义

| 常量名称        | 值（纳秒）               | 描述                                    |
|-------------|---------------------|---------------------------------------|
| `WorkDay`   | `43200000000000`    | 单日工作时长（12小时）                          |
| `WorkWeek`  | `259200000000000`   | 工作周时长（6个工作日）                          |
| `WorkMonth` | `842400000000000`   | 工作月时长（26个工作日）                         |
| `WorkYear`  | `10610400000000000` | 工作年时长（307个工作日，即 6 工作日/周 × 51 周 + 1 天） |

## 贡献指南

> 📌 请开发者遵循以下规范提交代码贡献

### 开发规范

1. 保持时间常量定义的原子性（建议按最小单位定义）
2. 所有业务时间模型需使用`WorkDay`作为基础单位
3. 提交前必须验证`WorkYear`与`WorkWeek`的乘法关系（307天 = 6天/周 × 51周 + 1天）
4. 任何修改需新增对应的测试用例（参考xtime955的测试模式）
5. 文档更新需同步到`README.md`和`xtime.go`注释

## 核心特性

> 🌟 遵循 **996 工作制**（早9点至晚9点，每周6天）

### 特性说明

- **工作时长定义**：单日工作时长为 12 小时（`WorkDay`），工作周为 6 天（`WorkWeek`），工作月为 26 天（`WorkMonth`），工作年为 307
  天（`WorkYear`）
- **适用场景**：适用于需要精确计算 996 工作制业务时间的项目（如考勤系统、工时统计）
- **扩展性**：支持基于 `WorkDay` 的自定义组合计算（例如：`WorkWeek = WorkDay * 6`）
- **无休息日模型**：默认不包含非工作时间（需开发者自行扩展）

### 与其他模型差异

| 特性     | xtime996（996 工作制） | xtime007（7-11 模型） |
|--------|-------------------|-------------------|
| 单日工作时长 | 12 小时（9:00-21:00） | 24 小时（全天候）        |
| 单周工作日  | 6 天               | 7 天               |
| 月度计算   | 26 个工作日（307 天模型）  | 30 天              |
| 适用行业   | 互联网、制造业等 996 常见行业 | 通用业务场景            |
