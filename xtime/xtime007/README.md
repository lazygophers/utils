# 996时间计算工具

## 模块功能
提供符合特殊工作制(007)的定制化时间计算能力，支持非标准工作时间的统计与转换。

## 核心函数
- `CalculateWorkHours(startTime, endTime time.Time) float64`
  计算指定时间段内的有效工作时长（按996规则）
- `IsOvertime(hour float64) bool`
  判断是否超出标准工作时长
- `GetWorkSchedule() []time.Time`
  获取当日工作时间排期

## 使用示例
```go
package main

import (
    "xtime/xtime007"
    "fmt"
    "time"
)

func main() {
    start := time.Date(2023, 11, 1, 9, 0, 0, 0, time.Local)
    end := time.Date(2023, 11, 1, 21, 30, 0, 0, time.Local)

    hours := xtime007.CalculateWorkHours(start, end)
    fmt.Printf("今日工作时长: %.1f 小时\n", hours) // 输出：今日工作时长: 12.5 小时

    if xtime007.IsOvertime(hours) {
        fmt.Println("已触发加班机制")
    }
}
```
