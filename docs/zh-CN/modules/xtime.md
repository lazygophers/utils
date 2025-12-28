---
title: xtime - 高级时间处理
---

# xtime - 高级时间处理

## 概述

xtime 模块提供高级时间处理，支持农历、生肖和节气。它包含全面的日历信息和时间常量。

## 时间常量

### 基本时间单位

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
)
```

### 工作时间常量

```go
const (
    WorkDayWeek  = Day * 5
    ResetDayWeek = Day * 2
    Week         = Day * 7
    
    WorkDayMonth  = Day*21 + HalfDay
    ResetDayMonth = Day*8 + HalfDay
    Month         = Day * 30
    
    QUARTER = Day * 91
    Year    = Day * 365
    Decade  = Year*10 + Day*2
    Century = Year*100 + Day*25
)
```

---

## 核心类型

### Calendar

包含阳历和农历数据的全面日历信息。

```go
type Calendar struct {
    *Time
    lunar  *Lunar
    zodiac ZodiacInfo
    season SeasonInfo
}
```

---

### ZodiacInfo

生肖和干支信息。

```go
type ZodiacInfo struct {
    Animal      string // 生肖：鼠、牛、虎...
    SkyTrunk    string // 天干：甲、乙、丙...
    EarthBranch string // 地支：子、丑、寅...
    YearGanZhi  string // 年干支：甲子、乙丑...
    MonthGanZhi string // 月干支
    DayGanZhi   string // 日干支
    HourGanZhi  string // 时干支
}
```

---

### SeasonInfo

节气和季节信息。

```go
type SeasonInfo struct {
    CurrentTerm    string    // 当前节气
    NextTerm       string    // 下个节气
    NextTermTime   time.Time // 下个节气时间
    Season         string    // 季节：春、夏、秋、冬
    SeasonProgress float64   // 季节进度(0-1)
    YearProgress   float64   // 年度进度(0-1)
}
```

---

## 构造函数

### NewCalendar()

创建包含完整农历和节气信息的日历对象。

```go
func NewCalendar(t time.Time) *Calendar
```

**参数：**
- `t` - 创建日历的时间

**返回值：**
- 包含完整信息的日历对象

**示例：**
```go
t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
cal := xtime.NewCalendar(t)
fmt.Println(cal.String())
```

---

### NowCalendar()

获取当前日历信息。

```go
func NowCalendar() *Calendar
```

**返回值：**
- 当前时间的日历对象

**示例：**
```go
cal := xtime.NowCalendar()
fmt.Printf("今天: %s\n", cal.String())
```

---

## 农历方法

### Lunar()

获取农历日期信息。

```go
func (c *Calendar) Lunar() *Lunar
```

**返回值：**
- 农历日期对象

---

### LunarDate()

获取农历日期字符串，格式：农历二零二三年八月十五。

```go
func (c *Calendar) LunarDate() string
```

**示例：**
```go
cal := xtime.NowCalendar()
fmt.Println(cal.LunarDate())
// 输出: 农历二零二三年八月十五
```

---

### LunarDateShort()

获取简短农历日期，格式：八月十五。

```go
func (c *Calendar) LunarDateShort() string
```

---

### IsLunarLeapYear()

检查是否为农历闰年。

```go
func (c *Calendar) IsLunarLeapYear() bool
```

---

### LunarLeapMonth()

获取农历闰月（0 表示无闰月）。

```go
func (c *Calendar) LunarLeapMonth() int64
```

---

## 生肖方法

### Animal()

获取生肖。

```go
func (c *Calendar) Animal() string
```

**示例：**
```go
cal := xtime.NowCalendar()
fmt.Printf("生肖: %s\n", cal.Animal())
// 输出: 龙
```

---

### AnimalWithYear()

获取生肖年份，格式：兔年。

```go
func (c *Calendar) AnimalWithYear() string
```

---

### YearGanZhi()

获取年干支，格式：癸卯。

```go
func (c *Calendar) YearGanZhi() string
```

---

### MonthGanZhi()

获取月干支。

```go
func (c *Calendar) MonthGanZhi() string
```

---

### DayGanZhi()

获取日干支。

```go
func (c *Calendar) DayGanZhi() string
```

---

### HourGanZhi()

获取时干支。

```go
func (c *Calendar) HourGanZhi() string
```

---

### FullGanZhi()

获取完整干支信息，格式：癸卯年 甲申月 己巳日 乙亥时。

```go
func (c *Calendar) FullGanZhi() string
```

---

## 节气方法

### CurrentSolarTerm()

获取当前节气。

```go
func (c *Calendar) CurrentSolarTerm() string
```

**示例：**
```go
cal := xtime.NowCalendar()
fmt.Printf("节气: %s\n", cal.CurrentSolarTerm())
// 输出: 立春
```

---

### NextSolarTerm()

获取下个节气。

```go
func (c *Calendar) NextSolarTerm() string
```

---

### NextSolarTermTime()

获取下个节气时间。

```go
func (c *Calendar) NextSolarTermTime() time.Time
```

---

### DaysToNextTerm()

获取距离下个节气的天数。

```go
func (c *Calendar) DaysToNextTerm() int
```

---

### Season()

获取当前季节。

```go
func (c *Calendar) Season() string
```

**返回值：**
- "春"、"夏"、"秋" 或 "冬"

---

### SeasonProgress()

获取季节进度（0-1）。

```go
func (c *Calendar) SeasonProgress() float64
```

---

### YearProgress()

获取年度进度（0-1）。

```go
func (c *Calendar) YearProgress() float64
```

---

## 格式化方法

### String()

获取完整日历信息字符串。

```go
func (c *Calendar) String() string
```

**示例：**
```go
cal := xtime.NowCalendar()
fmt.Println(cal.String())
// 输出: 2024年01月01日 八月十五 龙年 立春
```

---

### DetailedString()

获取详细日历信息。

```go
func (c *Calendar) DetailedString() string
```

**示例：**
```go
cal := xtime.NowCalendar()
fmt.Println(cal.DetailedString())
```

---

### ToMap()

转换为映射格式用于 JSON 序列化。

```go
func (c *Calendar) ToMap() map[string]interface{}
```

**返回值：**
- 包含阳历、农历、生肖和季节信息的映射

---

## 使用模式

### 显示当前日期信息

```go
func showCurrentDate() {
    cal := xtime.NowCalendar()
    
    fmt.Println("阳历：", cal.Time.Format("2006年01月02日"))
    fmt.Println("农历：", cal.LunarDate())
    fmt.Println("生肖：", cal.AnimalWithYear())
    fmt.Println("干支：", cal.FullGanZhi())
    fmt.Println("节气：", cal.CurrentSolarTerm())
    fmt.Println("季节：", cal.Season())
}
```

### 计算农历年龄

```go
func getLunarAge(birthDate time.Time) int {
    birthCal := xtime.NewCalendar(birthDate)
    currentCal := xtime.NowCalendar()
    
    return int(currentCal.Lunar().Year() - birthCal.Lunar().Year())
}
```

### 检查节气

```go
func isSolarTerm(date time.Time, term string) bool {
    cal := xtime.NewCalendar(date)
    return cal.CurrentSolarTerm() == term
}

func isBeforeSolarTerm(date time.Time, term string) bool {
    cal := xtime.NewCalendar(date)
    return cal.DaysToNextTerm() > 0 && cal.NextSolarTerm() == term
}
```

### 基于季节的操作

```go
func getSeasonInfo(date time.Time) (season string, progress float64) {
    cal := xtime.NewCalendar(date)
    return cal.Season(), cal.SeasonProgress()
}

func isSeason(date time.Time, season string) bool {
    cal := xtime.NewCalendar(date)
    return cal.Season() == season
}
```

---

## 最佳实践

### 日历创建

```go
// 好：创建一次日历并重用
cal := xtime.NowCalendar()
fmt.Println(cal.LunarDate())
fmt.Println(cal.Animal())

// 避免：创建多个日历
fmt.Println(xtime.NowCalendar().LunarDate())
fmt.Println(xtime.NowCalendar().Animal())
```

### 时区处理

```go
// 好：指定时区
loc, _ := time.LoadLocation("Asia/Shanghai")
t := time.Date(2024, 1, 1, 0, 0, 0, 0, loc)
cal := xtime.NewCalendar(t)

// 避免：对本地日期使用 UTC
t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
cal := xtime.NewCalendar(t)
```

---

## 相关文档

- [xtime996](/zh-CN/modules/xtime996) - 996 工作时间表
- [xtime955](/zh-CN/modules/xtime955) - 955 工作时间表
- [xtime007](/zh-CN/modules/xtime007) - 24/7 运营
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
