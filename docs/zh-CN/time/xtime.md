---
    title: xtime
    ---

    # xtime

    面向时间领域模型而不是简单格式化：包含日历、农历、节气、生肖和一组时间常量。

    ## 适用场景

    - 你需要同时处理公历与农历信息。
- 你需要节气、生肖、干支等中文语境日历信息。
- 你要在时间逻辑上继续进入 007 / 955 / 996 子包。

    ## 你会接触到什么

    - `xtime.NowCalendar()`：拿到当前日历对象。
- `Calendar.LunarDate()` / `Animal()` / `CurrentSolarTerm()`：读取农历、生肖、节气。
- `Calendar.String()` / `DetailedString()` / `ToMap()`：输出展示或序列化友好的结构。

## 快速示例

```go
cal := xtime.NowCalendar()
fmt.Println(cal.String())
fmt.Println(cal.LunarDate())
fmt.Println(cal.CurrentSolarTerm())
```

    ## 使用建议

    - 把它当作领域模型来读，而不是“又一个时间工具包”。
- 如果你要改动或依赖农历换算结果，先确认语义而不是猜测格式。

    ## 相关文档

    - [模块总览](/modules/overview)
    - [API 概览](/api/overview)
