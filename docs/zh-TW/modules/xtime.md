---
title: xtime - 高級時間處理
---

# xtime - 高級時間處理

## 概述

xtime 模組提供高級時間處理，支援農曆、生肖和節氣。它包含全面的日曆資訊和時間常量。

## 時間常量

### 基本時間單位

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

### 工作時間常量

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

## 核心類型

### Calendar

包含陽曆和農曆資料的全面日曆資訊。

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

生肖和干支資訊。

```go
type ZodiacInfo struct {
    Animal      string // 生肖：鼠、牛、虎...
    SkyTrunk    string // 天干：甲、乙、丙...
    EarthBranch string // 地支：子、丑、寅...
    YearGanZhi  string // 年干支：甲子、乙丑...
    MonthGanZhi string // 月干支
    DayGanZhi   string // 日干支
    HourGanZhi  string // 時干支
}
```

---

### SeasonInfo

節氣和季節資訊。

```go
type SeasonInfo struct {
    CurrentTerm    string    // 當前節氣
    NextTerm       string    // 下個節氣
    NextTermTime   time.Time // 下個節氣時間
    Season         string    // 季節：春、夏、秋、冬
    SeasonProgress float64   // 季節進度(0-1)
    YearProgress   float64   // 年度進度(0-1)
}
```

---

## 建構函數

### NewCalendar()

建立包含完整農曆和節氣資訊的日曆物件。

```go
func NewCalendar(t time.Time) *Calendar
```

**參數：**
- `t` - 建立日曆的時間

**返回值：**
- 包含完整資訊的日曆物件

**示例：**
```go
t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
cal := xtime.NewCalendar(t)
fmt.Println(cal.String())
```

---

### NowCalendar()

獲取當前日曆資訊。

```go
func NowCalendar() *Calendar
```

**返回值：**
- 當前時間的日曆物件

**示例：**
```go
cal := xtime.NowCalendar()
fmt.Printf("今天: %s\n", cal.String())
```

---

## 農曆方法

### Lunar()

獲取農曆日期資訊。

```go
func (c *Calendar) Lunar() *Lunar
```

**返回值：**
- 農曆日期物件

---

### LunarDate()

獲取農曆日期字串，格式：農曆二零二三年八月十五。

```go
func (c *Calendar) LunarDate() string
```

**示例：**
```go
cal := xtime.NowCalendar()
fmt.Println(cal.LunarDate())
// 輸出: 農曆二零二三年八月十五
```

---

### LunarDateShort()

獲取簡短農曆日期，格式：八月十五。

```go
func (c *Calendar) LunarDateShort() string
```

---

### IsLunarLeapYear()

檢查是否為農曆閏年。

```go
func (c *Calendar) IsLunarLeapYear() bool
```

---

### LunarLeapMonth()

獲取農曆閏月（0 表示無閏月）。

```go
func (c *Calendar) LunarLeapMonth() int64
```

---

## 生肖方法

### Animal()

獲取生肖。

```go
func (c *Calendar) Animal() string
```

**示例：**
```go
cal := xtime.NowCalendar()
fmt.Printf("生肖: %s\n", cal.Animal())
// 輸出: 龍
```

---

### AnimalWithYear()

獲取生肖年份，格式：兔年。

```go
func (c *Calendar) AnimalWithYear() string
```

---

### YearGanZhi()

獲取年干支，格式：癸卯。

```go
func (c *Calendar) YearGanZhi() string
```

---

### MonthGanZhi()

獲取月干支。

```go
func (c *Calendar) MonthGanZhi() string
```

---

### DayGanZhi()

獲取日干支。

```go
func (c *Calendar) DayGanZhi() string
```

---

### HourGanZhi()

獲取時干支。

```go
func (c *Calendar) HourGanZhi() string
```

---

### FullGanZhi()

獲取完整干支資訊，格式：癸卯年 甲申月 己巳日 乙亥時。

```go
func (c *Calendar) FullGanZhi() string
```

---

## 節氣方法

### CurrentSolarTerm()

獲取當前節氣。

```go
func (c *Calendar) CurrentSolarTerm() string
```

**示例：**
```go
cal := xtime.NowCalendar()
fmt.Printf("節氣: %s\n", cal.CurrentSolarTerm())
// 輸出: 立春
```

---

### NextSolarTerm()

獲取下個節氣。

```go
func (c *Calendar) NextSolarTerm() string
```

---

### NextSolarTermTime()

獲取下個節氣時間。

```go
func (c *Calendar) NextSolarTermTime() time.Time
```

---

### DaysToNextTerm()

獲取距離下個節氣的天數。

```go
func (c *Calendar) DaysToNextTerm() int
```

---

### Season()

獲取當前季節。

```go
func (c *Calendar) Season() string
```

**返回值：**
- "春"、"夏"、"秋" 或 "冬"

---

### SeasonProgress()

獲取季節進度（0-1）。

```go
func (c *Calendar) SeasonProgress() float64
```

---

### YearProgress()

獲取年度進度（0-1）。

```go
func (c *Calendar) YearProgress() float64
```

---

## 格式化方法

### String()

獲取完整日曆資訊字串。

```go
func (c *Calendar) String() string
```

**示例：**
```go
cal := xtime.NowCalendar()
fmt.Println(cal.String())
// 輸出: 2024年01月01日 八月十五 龙年 立春
```

---

### DetailedString()

獲取詳細日曆資訊。

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

轉換為映射格式用於 JSON 序列化。

```go
func (c *Calendar) ToMap() map[string]interface{}
```

**返回值：**
- 包含陽曆、農曆、生肖和季節資訊的映射

---

## 使用模式

### 顯示當前日期資訊

```go
func showCurrentDate() {
    cal := xtime.NowCalendar()
    
    fmt.Println("陽曆：", cal.Time.Format("2006年01月02日"))
    fmt.Println("農曆：", cal.LunarDate())
    fmt.Println("生肖：", cal.AnimalWithYear())
    fmt.Println("干支：", cal.FullGanZhi())
    fmt.Println("節氣：", cal.CurrentSolarTerm())
    fmt.Println("季節：", cal.Season())
}
```

### 計算農曆年齡

```go
func getLunarAge(birthDate time.Time) int {
    birthCal := xtime.NewCalendar(birthDate)
    currentCal := xtime.NowCalendar()
    
    return int(currentCal.Lunar().Year() - birthCal.Lunar().Year())
}
```

### 檢查節氣

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

### 基於季節的操作

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

## 最佳實踐

### 日曆建立

```go
// 好：建立一次日曆並重用
cal := xtime.NowCalendar()
fmt.Println(cal.LunarDate())
fmt.Println(cal.Animal())

// 避免：建立多個日曆
fmt.Println(xtime.NowCalendar().LunarDate())
fmt.Println(xtime.NowCalendar().Animal())
```

### 時區處理

```go
// 好：指定時區
loc, _ := time.LoadLocation("Asia/Shanghai")
t := time.Date(2024, 1, 1, 0, 0, 0, 0, loc)
cal := xtime.NewCalendar(t)

// 避免：對本地日期使用 UTC
t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
cal := xtime.NewCalendar(t)
```

---

## 相關文檔

- [xtime996](/zh-TW/modules/xtime996) - 996 工作時間表
- [xtime955](/zh-TW/modules/xtime955) - 955 工作時間表
- [xtime007](/zh-TW/modules/xtime007) - 24/7 運營
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
