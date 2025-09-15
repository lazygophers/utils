# xtime - 高级时间工具包

[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils/xtime.svg)](https://pkg.go.dev/github.com/lazygophers/utils/xtime)

一个功能全面的 Go 时间处理包，提供农历计算、二十四节气、生肖干支、传统中国历法等高级时间功能。

## 核心特性

### 基础时间操作
- 增强的时间封装器，扩展原生时间功能
- 时间范围计算（分钟/小时/日/周/月/季度/年的开始和结束）
- 自定义周起始日配置
- 随机延时工具，支持抖动

### 农历系统
- **完整农历支持** (1900-2100年)
  - 农历与公历日期转换
  - 闰月检测与计算
  - 传统中文日期格式化
  - 十二生肖推算

### 二十四节气
- **二十四节气计算** (1904-3000年)
  - 精确节气时间计算
  - 下个节气预测
  - 季节进度跟踪
  - 年度进度计算

### 传统中国历法
- **天干地支系统**
  - 年月日时干支计算
  - 完整干支体系支持
  - 传统中文时间表示

### 综合日历对象
- 公历、农历、传统历法统一对象
- 节日和特殊日期检测
- JSON 序列化支持
- 丰富的格式化选项

## 安装

```bash
go get github.com/lazygophers/utils/xtime
```

## 快速开始

```go
package main

import (
    "fmt"
    "time"
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // 创建综合日历对象
    cal := xtime.NowCalendar()

    // 显示完整信息
    fmt.Println(cal.String())
    // 输出: 2023年08月15日 七月廿八 兔年 处暑

    // 获取详细信息
    fmt.Println(cal.DetailedString())
}
```

## 核心 API 参考

### 时间封装器

```go
// 创建增强时间对象
t := xtime.With(time.Now())
t := xtime.Now()

// 时间范围操作
t.BeginningOfDay()    // 当日 00:00:00
t.EndOfDay()          // 当日 23:59:59.999999999
t.BeginningOfWeek()   // 本周开始
t.BeginningOfMonth()  // 本月第一天
t.BeginningOfYear()   // 本年1月1日

// 季度操作
t.Quarter()           // 当前季度 (1-4)
t.BeginningOfQuarter()
t.EndOfQuarter()
```

### 农历

```go
// 创建农历对象
lunar := xtime.WithLunar(time.Now())

// 基础农历信息
fmt.Println(lunar.Year())      // 2023
fmt.Println(lunar.Month())     // 7
fmt.Println(lunar.Day())       // 28
fmt.Println(lunar.Animal())    // "兔"

// 中文格式化
fmt.Println(lunar.YearAlias())     // "二零二三"
fmt.Println(lunar.MonthAlias())    // "七月"
fmt.Println(lunar.DayAlias())      // "廿八"

// 闰年信息
fmt.Println(lunar.IsLeap())        // false
fmt.Println(lunar.LeapMonth())     // 0 (无闰月)
fmt.Println(lunar.IsLeapMonth())   // false
```

### 二十四节气

```go
// 获取当前节气
term := xtime.CurrentSolarterm(time.Now())
fmt.Println(term.String())         // "处暑"
fmt.Println(term.Time())           // 节气的精确时间

// 获取下个节气
nextTerm := xtime.NextSolarterm(time.Now())
fmt.Println(nextTerm.String())     // "白露"

// 节气助手
helper := xtime.NewSolarTermHelper()
currentTerm := helper.GetCurrentTerm(time.Now())
if currentTerm != nil {
    fmt.Printf("当前节气: %s\n", currentTerm.Name)
    fmt.Printf("养生贴士: %v\n", currentTerm.Tips)
}
```

### 综合日历

```go
// 创建包含所有信息的日历
cal := xtime.NewCalendar(time.Now())

// 公历信息
fmt.Println(cal.Time.Format("2006-01-02"))
fmt.Println(cal.Time.Weekday())

// 农历信息
fmt.Println(cal.LunarDate())       // "农历二零二三年七月廿八"
fmt.Println(cal.LunarDateShort())  // "七月廿八"
fmt.Println(cal.AnimalWithYear())  // "兔年"

// 生肖干支信息
fmt.Println(cal.YearGanZhi())      // "癸卯"
fmt.Println(cal.FullGanZhi())      // "癸卯年 庚申月 己巳日 乙亥时"

// 节气和季节
fmt.Println(cal.CurrentSolarTerm()) // "处暑"
fmt.Println(cal.NextSolarTerm())    // "白露"
fmt.Println(cal.DaysToNextTerm())   // 7
fmt.Println(cal.Season())           // "秋"
fmt.Printf("季节进度: %.1f%%\n", cal.SeasonProgress()*100)
fmt.Printf("年度进度: %.1f%%\n", cal.YearProgress()*100)
```

### 时间解析

```go
// 解析时间字符串
t, err := xtime.Parse("2023-08-15")
if err != nil {
    panic(err)
}

// 强制解析（错误时panic）
t := xtime.MustParse("2023-08-15 14:30:00")
```

### 随机睡眠工具

```go
// 默认范围随机睡眠 (1-3秒)
xtime.RandSleep()

// 自定义范围随机睡眠
xtime.RandSleep(time.Second, time.Second*5)

// 单个最大值随机睡眠
xtime.RandSleep(time.Second*10)
```

## 高级功能

### 节日检测

```go
// 农历助手
lunarHelper := xtime.NewLunarHelper()

// 检查传统节日
festival := lunarHelper.GetTodayFestival()
if festival != nil {
    fmt.Printf("今天是: %s\n", festival.Name)
    fmt.Printf("传统习俗: %v\n", festival.Traditions)
    fmt.Printf("传统美食: %v\n", festival.Foods)
}

// 检查特殊日子
isSpecial, description := lunarHelper.IsSpecialDay(time.Now())
if isSpecial {
    fmt.Printf("特殊日子: %s\n", description)
}
```

### 批量节气操作

```go
helper := xtime.NewSolarTermHelper()

// 获取某年的节气日历
calendar := helper.GetTermCalendar(2023)
fmt.Printf("2023年有 %d 个月份包含节气\n", len(calendar))

// 获取最近的节气
recentTerms := helper.GetRecentTerms(time.Now(), 5)
for _, term := range recentTerms {
    fmt.Printf("%s: %s\n", term.Name, term.Time.Format("2006-01-02"))
}

// 计算到特定节气的天数
daysToSpring := helper.DaysUntilTerm(time.Now(), "立春")
fmt.Printf("距离立春还有: %d 天\n", daysToSpring)
```

### JSON 序列化

```go
cal := xtime.NowCalendar()
data := cal.ToMap()

// 数据结构包含:
// - solar: 公历信息
// - lunar: 农历信息
// - zodiac: 生肖干支信息
// - season: 节气季节信息

jsonBytes, _ := json.Marshal(data)
```

### 年龄计算

```go
lunarHelper := xtime.NewLunarHelper()

// 计算农历年龄（传统中国年龄算法）
birthTime := time.Date(1990, 5, 20, 0, 0, 0, 0, time.Local)
lunarAge := lunarHelper.GetLunarAge(birthTime, time.Now())
fmt.Printf("农历虚岁: %d 岁\n", lunarAge)

// 比较农历日期
comparison := lunarHelper.CompareLunarDates(birthTime, time.Now())
fmt.Printf("日期对比: %s\n", comparison)
```

## 数据准确性和来源

### 农历数据
- **覆盖范围**: 1900-2100年（200+年）
- **数据来源**: 紫金山天文台《中国天文年历》
- **验证机制**: 与历史农历记录交叉验证
- **精确度**: 日级别精确

### 节气数据
- **覆盖范围**: 1904-3000年（1000+年）
- **算法基础**: 高精度天文计算
- **精确度**: 分钟级别精确
- **数据更新**: 基于现代天文观测

## 性能特性

- **农历计算**: 每次操作约1-10微秒
- **节气查询**: 约100纳秒（预计算数据）
- **内存使用**: 完整数据表约2MB
- **线程安全**: 所有操作都是协程安全的
- **零分配**: 大多数操作避免内存分配

## 时间常量

包提供了便捷的时间段常量：

```go
const (
    Nanosecond  = time.Nanosecond
    Microsecond = time.Microsecond
    Millisecond = time.Millisecond
    Second      = time.Second
    Minute      = time.Minute
    HalfHour    = time.Minute * 30
    Hour        = time.Hour
    HalfDay     = time.Hour * 12
    Day         = time.Hour * 24

    WorkDayWeek  = Day * 5        // 5个工作日
    ResetDayWeek = Day * 2        // 2个休息日
    Week         = Day * 7

    WorkDayMonth  = Day*21 + HalfDay  // 约21.5个工作日
    ResetDayMonth = Day*8 + HalfDay   // 约8.5个非工作日
    Month         = Day * 30

    QUARTER = Day * 91            // 每季度约91天
    Year    = Day * 365
    Decade  = Year*10 + Day*2     // 考虑闰年
    Century = Year*100 + Day*25   // 考虑闰年
)
```

## 最佳实践

### 1. 复杂操作使用日历对象
```go
// 推荐：使用单个包含所有信息的日历对象
cal := xtime.NowCalendar()
fmt.Printf("%s %s %s", cal.Time.Format("2006-01-02"),
    cal.LunarDateShort(), cal.CurrentSolarTerm())

// 避免：多次分别调用
t := xtime.Now()
lunar := xtime.WithLunar(t.Time)
term := xtime.CurrentSolarterm(t.Time)
```

### 2. 多日期批量操作
```go
// 对多个日期高效
helper := xtime.NewSolarTermHelper()
calendar := helper.GetTermCalendar(2023)

// 对单个日期效率较低
for i := 1; i <= 12; i++ {
    date := time.Date(2023, time.Month(i), 1, 0, 0, 0, 0, time.Local)
    term := xtime.CurrentSolarterm(date)
}
```

### 3. 重用助手对象
```go
// 创建一次，多次使用
lunarHelper := xtime.NewLunarHelper()
termHelper := xtime.NewSolarTermHelper()

for _, date := range dates {
    festival := lunarHelper.GetTodayFestival()
    term := termHelper.GetCurrentTerm(date)
}
```

### 4. 正确处理时区
```go
// 指定时区以确保计算准确
beijing := time.FixedZone("CST", 8*3600)
t := time.Date(2023, 8, 15, 12, 0, 0, 0, beijing)
cal := xtime.NewCalendar(t)
```

## 错误处理

包遵循 Go 的惯用错误处理方式：

```go
// 可能失败的函数返回错误
t, err := xtime.Parse("invalid-date")
if err != nil {
    log.Fatal(err)
}

// Must 函数在错误时 panic（谨慎使用）
t := xtime.MustParse("2023-08-15")  // 解析失败时 panic
```

## 相关包

- `github.com/lazygophers/utils/candy` - 类型转换工具
- `github.com/lazygophers/utils/randx` - 随机数生成
- `github.com/jinzhu/now` - 增强时间解析（依赖）

## 示例

查看 [examples.go](examples.go) 获取全面的使用示例，包括：
- 基础日历操作
- 农历功能
- 节气计算
- 节日检测
- 批量操作
- JSON 序列化
- 实际使用场景

## 贡献

此包是 LazyGophers Utils 集合的一部分。贡献指南：

1. 遵循 Go 编码标准
2. 为新功能添加完整测试
3. 更新 API 变更的文档
4. 确保向后兼容性

## 许可证

此包是 LazyGophers Utils 项目的一部分。许可证信息请查看主仓库。

---

*注意：此包专注于传统中国历法系统。对于伊斯兰历、希伯来历或其他历法系统，请考虑使用专门的包。*