# xtime007

> 早上 0 点到晚上 24 点，一周七天

提供基于 24 小时 x 7 天 工作制的时间计算常量

# 安装配置

```bash
go get -u github.com/lazygophers/utils/xtime/xtime007
```

添加到项目依赖：
```go
import "github.com/lazygophers/utils/xtime/xtime007"
```
# 包功能描述

## 使用示例

```go
// 计算年度工作时间占比
fmt.Println(xtime007.WorkYear.Seconds()) // 31536000 秒

// 验证季度模型
if xtime007.Quarter == xtime007.WorkQuarter {
    fmt.Println("季度模型已对齐")
}

// 检查月度休息时间
fmt.Println(xtime007.WorkMonth - xtime007.RestMonth) // 2592000000000000 微秒
```

## 核心特性
* **业务时间模型** - 提供标准工作日/周/月/年模型
* **全天候支持** - 所有时间单位基于24小时制
* **复合单位** - 通过预定义组合简化时间计算

## 注意事项
> 1. 月份默认按30天计算
> 2. 季度定义为91天（3个月）
> 3. 所有时间单位继承自 time.Duration 类型

# 注释汇总

| 常量名称     | 描述                     | 计算方式           | 适用场景             |
|--------------|--------------------------|--------------------|----------------------|
| WorkDay      | 标准工作日定义为24小时   | time.Hour * 24     | 持续运行服务         |
| RestDay      | 非工作日时间模型         | 0                  | 连续运行无休息场景   |
| WorkWeek     | 7天完整工作周期          | WorkDay * 7        | 周维度业务统计       |
| RestWeek     | 非工作周时间模型         | 0                  | 无休息日业务场景     |
| WorkMonth    | 30天标准工作月           | Day * 30           | 月度结算系统         |
| RestMonth    | 无休息日模型             | 0                  | 连续运行无休息场景   |
| Quarter      | 91天季度基准             | Day * 91           | 季度维度数据分析     |
| RestQuarter  | 无休息日模型             | 0                  | 连续运行无休息场景   |
| WorkQuarter  | 91天标准季度             | Quarter            | 季度维度数据分析     |
| WorkYear     | 365天自然年工作模型      | Day * 365          | 年度计划制定         |
| RestYear     | 无休息日模型             | 0                  | 连续运行无休息场景   |
