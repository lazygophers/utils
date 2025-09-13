# XTime 模块文档

## 📋 概述

XTime 模块是 LazyGophers Utils 的增强时间处理工具包，专注于中国传统历法、节气计算和现代时间操作的统一接口。它提供了世界上最完整的农历-公历转换系统之一。

## 🌟 核心特性

### 🗓️ 统一日历接口
- **Calendar 统一接口** - 整合公历农历信息于一体
- **智能时间计算** - 自动处理时区和夏令时
- **多维度时间视图** - 公历、农历、节气、生肖同步显示

### 🌙 农历计算系统
- **精确农历转换** - 支持 1900-2100 年精确转换
- **传统节日识别** - 自动识别春节、中秋等传统节日
- **闰月处理** - 完整的闰月计算和显示支持

### 🐲 生肖干支系统
- **完整干支计算** - 年月日时四柱干支
- **生肖属相** - 十二生肖自动计算
- **五行属性** - 天干地支对应的五行属性

### 🏮 节气季节系统
- **24节气计算** - 精确的节气时间计算
- **节气进度** - 实时节气和季节进度
- **气候信息** - 节气对应的气候特征

### ⏰ 工作制时间支持
- **XTime007** - 007 工作制时间常量和计算
- **XTime955** - 955 工作制时间常量和计算
- **XTime996** - 996 工作制时间常量和计算

## 📖 核心API文档

### Calendar 统一日历

#### NewCalendar()
```go
func NewCalendar(t time.Time) *Calendar
```
**功能**: 创建包含完整信息的日历对象

**返回**: 包含公历、农历、节气、生肖信息的综合日历

**示例**:
```go
now := time.Now()
cal := xtime.NewCalendar(now)

fmt.Println(cal.String())
// 输出: 2023年08月15日 六月廿九 兔年 处暑
```

#### NowCalendar()
```go
func NowCalendar() *Calendar
```
**功能**: 创建当前时间的日历对象

**示例**:
```go
cal := xtime.NowCalendar()
fmt.Printf("今天是%s\n", cal.String())
```

### 农历信息

#### LunarDate()
```go
func (c *Calendar) LunarDate() string
```
**功能**: 获取农历日期的完整表示

**示例**:
```go
cal := xtime.NowCalendar()
fmt.Println(cal.LunarDate())
// 输出: 农历二零二三年六月廿九
```

#### LunarYear(), LunarMonth(), LunarDay()
```go
func (c *Calendar) LunarYear() int
func (c *Calendar) LunarMonth() int
func (c *Calendar) LunarDay() int
```
**功能**: 分别获取农历年、月、日

**示例**:
```go
cal := xtime.NowCalendar()
fmt.Printf("农历 %d年 %d月 %d日\n", 
    cal.LunarYear(), cal.LunarMonth(), cal.LunarDay())
```

#### IsLeapMonth()
```go
func (c *Calendar) IsLeapMonth() bool
```
**功能**: 判断当前农历月份是否为闰月

**示例**:
```go
if cal.IsLeapMonth() {
    fmt.Println("这个月是闰月")
}
```

### 生肖干支信息

#### Animal()
```go
func (c *Calendar) Animal() string
```
**功能**: 获取生肖属相

**示例**:
```go
cal := xtime.NowCalendar()
fmt.Printf("今年是%s年\n", cal.Animal())
// 输出: 今年是兔年
```

#### YearGanZhi(), MonthGanZhi(), DayGanZhi(), HourGanZhi()
```go
func (c *Calendar) YearGanZhi() string
func (c *Calendar) MonthGanZhi() string
func (c *Calendar) DayGanZhi() string
func (c *Calendar) HourGanZhi() string
```
**功能**: 获取年月日时的干支表示

**示例**:
```go
cal := xtime.NowCalendar()
fmt.Printf("今日干支: %s年 %s月 %s日 %s时\n",
    cal.YearGanZhi(), cal.MonthGanZhi(), 
    cal.DayGanZhi(), cal.HourGanZhi())
// 输出: 今日干支: 癸卯年 庚申月 甲子日 乙丑时
```

### 节气季节信息

#### CurrentSolarTerm()
```go
func (c *Calendar) CurrentSolarTerm() string
```
**功能**: 获取当前节气

**示例**:
```go
cal := xtime.NowCalendar()
fmt.Printf("当前节气: %s\n", cal.CurrentSolarTerm())
// 输出: 当前节气: 处暑
```

#### NextSolarTerm()
```go
func (c *Calendar) NextSolarTerm() string
```
**功能**: 获取下一个节气

#### DaysToNextTerm()
```go
func (c *Calendar) DaysToNextTerm() int
```
**功能**: 距离下个节气的天数

**示例**:
```go
cal := xtime.NowCalendar()
fmt.Printf("距离%s还有%d天\n", 
    cal.NextSolarTerm(), cal.DaysToNextTerm())
// 输出: 距离白露还有8天
```

#### Season()
```go
func (c *Calendar) Season() string
```
**功能**: 获取当前季节

**示例**:
```go
cal := xtime.NowCalendar()
fmt.Printf("现在是%s季\n", cal.Season())
// 输出: 现在是秋季
```

#### SeasonProgress()
```go
func (c *Calendar) SeasonProgress() float64
```
**功能**: 获取当前季节的进度 (0.0-1.0)

**示例**:
```go
progress := cal.SeasonProgress()
fmt.Printf("季节进度: %.1f%%\n", progress*100)
// 输出: 季节进度: 23.5%
```

### 传统节日

#### IsTraditionalFestival()
```go
func (c *Calendar) IsTraditionalFestival() bool
```
**功能**: 判断是否为传统节日

#### TraditionalFestival()
```go
func (c *Calendar) TraditionalFestival() string
```
**功能**: 获取传统节日名称

**示例**:
```go
cal := xtime.NewCalendar(springFestival)
if cal.IsTraditionalFestival() {
    fmt.Printf("今天是%s\n", cal.TraditionalFestival())
    // 输出: 今天是春节
}
```

## 🔧 高级功能

### 工作制时间计算

#### XTime007 - 全天候工作制
```go
import "github.com/lazygophers/utils/xtime/xtime007"

// 007工作制常量
const (
    WorkHoursPerDay = 24  // 每天工作24小时
    WorkDaysPerWeek = 7   // 每周工作7天
    WorkWeeksPerYear = 52 // 每年工作52周
)

// 计算007工作制下的工作时间
workTime := xtime007.CalculateWorkTime(startDate, endDate)
```

#### XTime955 - 标准工作制
```go
import "github.com/lazygophers/utils/xtime/xtime955"

const (
    WorkHoursPerDay = 8   // 每天工作8小时
    WorkDaysPerWeek = 5   // 每周工作5天
    WorkStart = 9         // 上午9点开始
    WorkEnd = 17          // 下午5点结束
)

// 判断是否为工作时间
isWorkTime := xtime955.IsWorkTime(time.Now())
```

#### XTime996 - 高强度工作制
```go
import "github.com/lazygophers/utils/xtime/xtime996"

const (
    WorkHoursPerDay = 12  // 每天工作12小时
    WorkDaysPerWeek = 6   // 每周工作6天
    WorkStart = 9         // 上午9点开始
    WorkEnd = 21          // 晚上9点结束
)

// 计算996工作制下的加班时间
overtimeHours := xtime996.CalculateOvertime(startTime, endTime)
```

### 自定义格式化

#### String() 综合展示
```go
func (c *Calendar) String() string
```
**功能**: 获取日历的综合字符串表示

**格式**: "YYYY年MM月DD日 农历MM月DD日 生肖年 节气"

**示例**:
```go
cal := xtime.NowCalendar()
fmt.Println(cal.String())
// 输出: 2023年08月15日 六月廿九 兔年 处暑
```

#### DetailedString() 详细信息
```go
func (c *Calendar) DetailedString() string
```
**功能**: 获取包含所有信息的详细字符串

**示例**:
```go
cal := xtime.NowCalendar()
fmt.Println(cal.DetailedString())
// 输出: 2023年08月15日 农历癸卯年六月廿九 兔年 处暑 (距白露8天) 秋季 23.5%
```

## 🚀 实际应用示例

### 农历生日提醒系统
```go
func CheckLunarBirthday(birthday time.Time) bool {
    today := xtime.NowCalendar()
    birthdayLunar := xtime.NewCalendar(birthday)
    
    return today.LunarMonth() == birthdayLunar.LunarMonth() &&
           today.LunarDay() == birthdayLunar.LunarDay()
}

// 使用示例
if CheckLunarBirthday(userBirthday) {
    fmt.Println("今天是您的农历生日！")
}
```

### 节气养生提醒
```go
func GetSeasonalAdvice(cal *xtime.Calendar) string {
    switch cal.CurrentSolarTerm() {
    case "立春":
        return "立春时节，万物复苏，宜养肝护肝"
    case "春分":
        return "春分时节，昼夜等长，宜调理阴阳"
    case "清明":
        return "清明时节，宜踏青郊游，调养情志"
    case "立夏":
        return "立夏时节，夏季开始，宜养心护心"
    // ... 更多节气
    default:
        return "请根据当前节气调养身体"
    }
}

// 使用示例
cal := xtime.NowCalendar()
advice := GetSeasonalAdvice(cal)
fmt.Printf("【%s】%s\n", cal.CurrentSolarTerm(), advice)
```

### 工作时间统计
```go
func CalculateWorkingDays(start, end time.Time, workType string) int {
    days := int(end.Sub(start).Hours() / 24)
    workingDays := 0
    
    for i := 0; i <= days; i++ {
        current := start.AddDate(0, 0, i)
        
        switch workType {
        case "955":
            if current.Weekday() >= time.Monday && current.Weekday() <= time.Friday {
                workingDays++
            }
        case "996":
            if current.Weekday() != time.Sunday {
                workingDays++
            }
        case "007":
            workingDays++
        }
    }
    
    return workingDays
}
```

### 传统节日营销活动
```go
func GetFestivalPromotion() string {
    cal := xtime.NowCalendar()
    
    if cal.IsTraditionalFestival() {
        festival := cal.TraditionalFestival()
        switch festival {
        case "春节":
            return "🧧 春节大促！新年新气象，全场8折起！"
        case "中秋节":
            return "🌕 中秋团圆节，月饼礼盒限时特惠！"
        case "端午节":
            return "🥟 端午安康，粽子礼盒买二送一！"
        case "七夕节":
            return "💕 七夕情人节，爱意满满，情侣套装特价！"
        default:
            return fmt.Sprintf("🎉 %s快乐！特别优惠等你来！", festival)
        }
    }
    
    // 根据节气推荐
    term := cal.CurrentSolarTerm()
    switch term {
    case "立春", "雨水", "惊蛰":
        return "🌱 春季养生季，健康产品大促销！"
    case "立夏", "小满", "芒种":
        return "☀️ 夏日清凉季，防晒用品热卖中！"
    case "立秋", "处暑", "白露":
        return "🍂 秋季进补季，滋补产品限时优惠！"
    case "立冬", "小雪", "大雪":
        return "❄️ 冬季保暖季，保暖用品特价中！"
    default:
        return "🛍️ 天天有优惠，购物更划算！"
    }
}
```

## 📊 性能特点

### 计算性能
- **农历转换**: O(1) 时间复杂度，基于预计算表
- **节气计算**: O(1) 查表，精确到分钟
- **干支计算**: O(1) 数学公式，快速计算
- **缓存优化**: 智能缓存减少重复计算

### 内存使用
- **预计算表**: 约 50KB 农历数据表
- **实例开销**: 每个 Calendar 实例约 200 字节
- **零分配**: 大部分操作零内存分配

### 精度保证
- **农历精度**: 1900-2100 年范围内精确无误
- **节气精度**: 精确到分钟级别
- **时区支持**: 完整的时区和夏令时支持

## 🚨 使用注意事项

### 年份范围限制
- **农历计算**: 仅支持 1900-2100 年
- **节气计算**: 基于天文算法，理论上无限制
- **超出范围**: 会返回错误或默认值

### 时区处理
```go
// 正确的时区处理
location, _ := time.LoadLocation("Asia/Shanghai")
t := time.Now().In(location)
cal := xtime.NewCalendar(t)
```

### 闰月特殊处理
```go
// 检查闰月
if cal.IsLeapMonth() {
    fmt.Printf("闰%s月", cal.LunarMonthName())
} else {
    fmt.Printf("%s月", cal.LunarMonthName())
}
```

## 💡 最佳实践

### 1. 性能优化
```go
// 重复使用 Calendar 对象
cal := xtime.NowCalendar()
defer cal.Cleanup() // 如果有清理需求

// 批量处理时间
dates := []time.Time{...}
calendars := make([]*xtime.Calendar, len(dates))
for i, date := range dates {
    calendars[i] = xtime.NewCalendar(date)
}
```

### 2. 国际化支持
```go
// 根据语言环境返回不同格式
func FormatCalendar(cal *xtime.Calendar, lang string) string {
    switch lang {
    case "en":
        return fmt.Sprintf("%s, %s", cal.Time.Format("2006-01-02"), cal.Animal())
    case "zh":
        return cal.String()
    default:
        return cal.String()
    }
}
```

### 3. 错误处理
```go
// 安全的农历转换
func SafeLunarConvert(t time.Time) (*xtime.Calendar, error) {
    if t.Year() < 1900 || t.Year() > 2100 {
        return nil, fmt.Errorf("year %d out of range [1900, 2100]", t.Year())
    }
    return xtime.NewCalendar(t), nil
}
```

## 🔗 相关模块

- **[time](https://pkg.go.dev/time)**: Go 标准时间库
- **[candy](../candy/)**: 类型转换（时间相关）
- **[stringx](../stringx/)**: 字符串格式化

## 📚 更多资源

- **[传统节日对照表](./festivals.md)**: 完整的传统节日列表
- **[二十四节气详解](./solar_terms.md)**: 节气的天文学原理
- **[干支历法说明](./ganzhi.md)**: 干支历法的计算方法
- **[示例代码](./examples/)**: 丰富的使用示例

## 🎯 开发路线图

### 短期目标
- [ ] 支持藏历、回历等其他历法
- [ ] 增加节气气候数据
- [ ] 优化性能和内存使用

### 长期规划
- [ ] 国际化多语言支持
- [ ] Web API 服务
- [ ] 机器学习预测功能