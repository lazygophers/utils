---
title: xtime - Advanced Time Processing
---

# xtime - Advanced Time Processing

## Overview

The xtime module provides advanced time processing with lunar calendar, zodiac, and solar terms support. It includes comprehensive calendar information and time constants.

## Time Constants

### Basic Time Units

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

### Work Time Constants

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

## Core Types

### Calendar

Comprehensive calendar information with solar and lunar data.

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

Zodiac and Ganzhi information.

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

Solar term and season information.

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

## Constructor Functions

### NewCalendar()

Create calendar object with complete lunar and solar term information.

```go
func NewCalendar(t time.Time) *Calendar
```

**Parameters:**
- `t` - Time to create calendar for

**Returns:**
- Calendar object with complete information

**Example:**
```go
t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
cal := xtime.NewCalendar(t)
fmt.Println(cal.String())
```

---

### NowCalendar()

Get current calendar information.

```go
func NowCalendar() *Calendar
```

**Returns:**
- Calendar object for current time

**Example:**
```go
cal := xtime.NowCalendar()
fmt.Printf("Today: %s\n", cal.String())
```

---

## Lunar Calendar Methods

### Lunar()

Get lunar date information.

```go
func (c *Calendar) Lunar() *Lunar
```

**Returns:**
- Lunar date object

---

### LunarDate()

Get lunar date string, format: 农历二零二三年八月十五.

```go
func (c *Calendar) LunarDate() string
```

**Example:**
```go
cal := xtime.NowCalendar()
fmt.Println(cal.LunarDate())
// Output: 农历二零二三年八月十五
```

---

### LunarDateShort()

Get short lunar date, format: 八月十五.

```go
func (c *Calendar) LunarDateShort() string
```

---

### IsLunarLeapYear()

Check if it's a lunar leap year.

```go
func (c *Calendar) IsLunarLeapYear() bool
```

---

### LunarLeapMonth()

Get lunar leap month (0 means no leap month).

```go
func (c *Calendar) LunarLeapMonth() int64
```

---

## Zodiac Methods

### Animal()

Get zodiac animal.

```go
func (c *Calendar) Animal() string
```

**Example:**
```go
cal := xtime.NowCalendar()
fmt.Printf("Animal: %s\n", cal.Animal())
// Output: 龙
```

---

### AnimalWithYear()

Get zodiac year, format: 兔年.

```go
func (c *Calendar) AnimalWithYear() string
```

---

### YearGanZhi()

Get year Ganzhi, format: 癸卯.

```go
func (c *Calendar) YearGanZhi() string
```

---

### MonthGanZhi()

Get month Ganzhi.

```go
func (c *Calendar) MonthGanZhi() string
```

---

### DayGanZhi()

Get day Ganzhi.

```go
func (c *Calendar) DayGanZhi() string
```

---

### HourGanZhi()

Get hour Ganzhi.

```go
func (c *Calendar) HourGanZhi() string
```

---

### FullGanZhi()

Get complete Ganzhi information, format: 癸卯年 甲申月 己巳日 乙亥时.

```go
func (c *Calendar) FullGanZhi() string
```

---

## Solar Term Methods

### CurrentSolarTerm()

Get current solar term.

```go
func (c *Calendar) CurrentSolarTerm() string
```

**Example:**
```go
cal := xtime.NowCalendar()
fmt.Printf("Solar Term: %s\n", cal.CurrentSolarTerm())
// Output: 立春
```

---

### NextSolarTerm()

Get next solar term.

```go
func (c *Calendar) NextSolarTerm() string
```

---

### NextSolarTermTime()

Get next solar term time.

```go
func (c *Calendar) NextSolarTermTime() time.Time
```

---

### DaysToNextTerm()

Get days to next solar term.

```go
func (c *Calendar) DaysToNextTerm() int
```

---

### Season()

Get current season.

```go
func (c *Calendar) Season() string
```

**Returns:**
- "春", "夏", "秋", or "冬"

---

### SeasonProgress()

Get season progress (0-1).

```go
func (c *Calendar) SeasonProgress() float64
```

---

### YearProgress()

Get year progress (0-1).

```go
func (c *Calendar) YearProgress() float64
```

---

## Formatting Methods

### String()

Get complete calendar information string.

```go
func (c *Calendar) String() string
```

**Example:**
```go
cal := xtime.NowCalendar()
fmt.Println(cal.String())
// Output: 2024年01月01日 八月十五 龙年 立春
```

---

### DetailedString()

Get detailed calendar information.

```go
func (c *Calendar) DetailedString() string
```

**Example:**
```go
cal := xtime.NowCalendar()
fmt.Println(cal.DetailedString())
```

---

### ToMap()

Convert to map format for JSON serialization.

```go
func (c *Calendar) ToMap() map[string]interface{}
```

**Returns:**
- Map with solar, lunar, zodiac, and season information

---

## Usage Patterns

### Display Current Date Information

```go
func showCurrentDate() {
    cal := xtime.NowCalendar()
    
    fmt.Println("公历：", cal.Time.Format("2006年01月02日"))
    fmt.Println("农历：", cal.LunarDate())
    fmt.Println("生肖：", cal.AnimalWithYear())
    fmt.Println("干支：", cal.FullGanZhi())
    fmt.Println("节气：", cal.CurrentSolarTerm())
    fmt.Println("季节：", cal.Season())
}
```

### Calculate Age in Lunar Years

```go
func getLunarAge(birthDate time.Time) int {
    birthCal := xtime.NewCalendar(birthDate)
    currentCal := xtime.NowCalendar()
    
    return int(currentCal.Lunar().Year() - birthCal.Lunar().Year())
}
```

### Check Solar Term

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

### Season-Based Operations

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

## Best Practices

### Calendar Creation

```go
// Good: Create calendar once and reuse
cal := xtime.NowCalendar()
fmt.Println(cal.LunarDate())
fmt.Println(cal.Animal())

// Avoid: Creating multiple calendars
fmt.Println(xtime.NowCalendar().LunarDate())
fmt.Println(xtime.NowCalendar().Animal())
```

### Time Zone Handling

```go
// Good: Specify time zone
loc, _ := time.LoadLocation("Asia/Shanghai")
t := time.Date(2024, 1, 1, 0, 0, 0, 0, loc)
cal := xtime.NewCalendar(t)

// Avoid: Using UTC for local dates
t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
cal := xtime.NewCalendar(t)
```

---

## Related Documentation

- [xtime996](/en/modules/xtime996) - 996 work schedule
- [xtime955](/en/modules/xtime955) - 955 work schedule
- [xtime007](/en/modules/xtime007) - 24/7 operations
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
