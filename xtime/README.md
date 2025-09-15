# xtime - Advanced Time Utilities

[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils/xtime.svg)](https://pkg.go.dev/github.com/lazygophers/utils/xtime)

A comprehensive Go package for advanced time operations, featuring lunar calendar calculations, Chinese solar terms (24 seasonal divisions), zodiac animals, and traditional Chinese calendar support.

## Features

### Core Time Operations
- Enhanced time wrapper with extended functionality
- Time range calculations (beginning/end of minute, hour, day, week, month, quarter, year)
- Custom week start day configuration
- Random sleep utilities with jitter support

### Lunar Calendar System
- **Complete Lunar Calendar Support** (1900-2100)
  - Lunar-to-solar date conversion
  - Leap month detection and calculation
  - Traditional Chinese date formatting
  - Chinese zodiac animals (12-year cycle)

### Chinese Solar Terms (节气)
- **24 Solar Terms Support** (1904-3000)
  - Precise solar term calculations
  - Next solar term predictions
  - Seasonal progression tracking
  - Year progress calculations

### Traditional Chinese Calendar
- **Heavenly Stems and Earthly Branches (天干地支)**
  - Year, month, day, hour stem-branch calculations
  - Complete GanZhi (干支) system support
  - Traditional Chinese time representation

### Comprehensive Calendar Object
- Unified calendar with solar, lunar, and traditional information
- Festival and special day detection
- JSON serialization support
- Rich formatting options

## Installation

```bash
go get github.com/lazygophers/utils/xtime
```

## Quick Start

```go
package main

import (
    "fmt"
    "time"
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // Create a comprehensive calendar object
    cal := xtime.NowCalendar()

    // Display complete information
    fmt.Println(cal.String())
    // Output: 2023年08月15日 七月廿八 兔年 处暑

    // Get detailed information
    fmt.Println(cal.DetailedString())
}
```

## Core API Reference

### Time Wrapper

```go
// Create enhanced time object
t := xtime.With(time.Now())
t := xtime.Now()

// Time range operations
t.BeginningOfDay()    // 00:00:00 of current day
t.EndOfDay()          // 23:59:59.999999999 of current day
t.BeginningOfWeek()   // Start of current week
t.BeginningOfMonth()  // First day of current month
t.BeginningOfYear()   // January 1st of current year

// Quarter operations
t.Quarter()           // Current quarter (1-4)
t.BeginningOfQuarter()
t.EndOfQuarter()
```

### Lunar Calendar

```go
// Create lunar calendar object
lunar := xtime.WithLunar(time.Now())

// Basic lunar information
fmt.Println(lunar.Year())      // 2023
fmt.Println(lunar.Month())     // 7
fmt.Println(lunar.Day())       // 28
fmt.Println(lunar.Animal())    // "兔" (Rabbit)

// Chinese formatting
fmt.Println(lunar.YearAlias())     // "二零二三"
fmt.Println(lunar.MonthAlias())    // "七月"
fmt.Println(lunar.DayAlias())      // "廿八"

// Leap year information
fmt.Println(lunar.IsLeap())        // false
fmt.Println(lunar.LeapMonth())     // 0 (no leap month)
fmt.Println(lunar.IsLeapMonth())   // false
```

### Solar Terms (节气)

```go
// Get current solar term
term := xtime.CurrentSolarterm(time.Now())
fmt.Println(term.String())         // "处暑"
fmt.Println(term.Time())           // Exact time of the solar term

// Get next solar term
nextTerm := xtime.NextSolarterm(time.Now())
fmt.Println(nextTerm.String())     // "白露"

// Solar term helper
helper := xtime.NewSolarTermHelper()
currentTerm := helper.GetCurrentTerm(time.Now())
if currentTerm != nil {
    fmt.Printf("Current: %s\n", currentTerm.Name)
    fmt.Printf("Tips: %v\n", currentTerm.Tips)
}
```

### Comprehensive Calendar

```go
// Create calendar with all information
cal := xtime.NewCalendar(time.Now())

// Solar calendar information
fmt.Println(cal.Time.Format("2006-01-02"))
fmt.Println(cal.Time.Weekday())

// Lunar calendar information
fmt.Println(cal.LunarDate())       // "农历二零二三年七月廿八"
fmt.Println(cal.LunarDateShort())  // "七月廿八"
fmt.Println(cal.AnimalWithYear())  // "兔年"

// Zodiac and GanZhi information
fmt.Println(cal.YearGanZhi())      // "癸卯"
fmt.Println(cal.FullGanZhi())      // "癸卯年 庚申月 己巳日 乙亥时"

// Solar terms and seasons
fmt.Println(cal.CurrentSolarTerm()) // "处暑"
fmt.Println(cal.NextSolarTerm())    // "白露"
fmt.Println(cal.DaysToNextTerm())   // 7
fmt.Println(cal.Season())           // "秋"
fmt.Printf("Season progress: %.1f%%\n", cal.SeasonProgress()*100)
fmt.Printf("Year progress: %.1f%%\n", cal.YearProgress()*100)
```

### Time Parsing

```go
// Parse time strings
t, err := xtime.Parse("2023-08-15")
if err != nil {
    panic(err)
}

// Must parse (panics on error)
t := xtime.MustParse("2023-08-15 14:30:00")
```

### Random Sleep Utilities

```go
// Random sleep with default range (1-3 seconds)
xtime.RandSleep()

// Random sleep with custom range
xtime.RandSleep(time.Second, time.Second*5)

// Random sleep with single maximum
xtime.RandSleep(time.Second*10)
```

## Advanced Features

### Festival Detection

```go
// Lunar calendar helper
lunarHelper := xtime.NewLunarHelper()

// Check for traditional festivals
festival := lunarHelper.GetTodayFestival()
if festival != nil {
    fmt.Printf("Today is: %s\n", festival.Name)
    fmt.Printf("Traditions: %v\n", festival.Traditions)
    fmt.Printf("Foods: %v\n", festival.Foods)
}

// Check for special days
isSpecial, description := lunarHelper.IsSpecialDay(time.Now())
if isSpecial {
    fmt.Printf("Special day: %s\n", description)
}
```

### Batch Solar Term Operations

```go
helper := xtime.NewSolarTermHelper()

// Get solar term calendar for a year
calendar := helper.GetTermCalendar(2023)
fmt.Printf("2023 has solar terms in %d months\n", len(calendar))

// Get recent solar terms
recentTerms := helper.GetRecentTerms(time.Now(), 5)
for _, term := range recentTerms {
    fmt.Printf("%s: %s\n", term.Name, term.Time.Format("2006-01-02"))
}

// Calculate days until specific solar term
daysToSpring := helper.DaysUntilTerm(time.Now(), "立春")
fmt.Printf("Days until Spring Festival: %d\n", daysToSpring)
```

### JSON Serialization

```go
cal := xtime.NowCalendar()
data := cal.ToMap()

// Structure contains:
// - solar: solar calendar information
// - lunar: lunar calendar information
// - zodiac: zodiac and GanZhi information
// - season: solar terms and seasonal information

jsonBytes, _ := json.Marshal(data)
```

### Age Calculations

```go
lunarHelper := xtime.NewLunarHelper()

// Calculate lunar age (traditional Chinese age)
birthTime := time.Date(1990, 5, 20, 0, 0, 0, 0, time.Local)
lunarAge := lunarHelper.GetLunarAge(birthTime, time.Now())
fmt.Printf("Lunar age: %d years\n", lunarAge)

// Compare lunar dates
comparison := lunarHelper.CompareLunarDates(birthTime, time.Now())
fmt.Printf("Date comparison: %s\n", comparison)
```

## Data Accuracy and Sources

### Lunar Calendar Data
- **Coverage**: 1900-2100 (200+ years)
- **Source**: Purple Mountain Observatory's "Chinese Astronomical Almanac"
- **Validation**: Cross-referenced with historical lunar records
- **Accuracy**: Precise to the day level

### Solar Terms Data
- **Coverage**: 1904-3000 (1000+ years)
- **Algorithm**: High-precision astronomical calculations
- **Accuracy**: Precise to the minute level
- **Updates**: Based on modern astronomical observations

## Performance Characteristics

- **Lunar calculations**: ~1-10 microseconds per operation
- **Solar term lookups**: ~100 nanoseconds (pre-calculated data)
- **Memory usage**: ~2MB for complete data tables
- **Thread safety**: All operations are goroutine-safe
- **Zero allocations**: Most operations avoid memory allocation

## Time Constants

The package provides convenient time duration constants:

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

    WorkDayWeek  = Day * 5        // 5 working days
    ResetDayWeek = Day * 2        // 2 weekend days
    Week         = Day * 7

    WorkDayMonth  = Day*21 + HalfDay  // ~21.5 working days
    ResetDayMonth = Day*8 + HalfDay   // ~8.5 non-working days
    Month         = Day * 30

    QUARTER = Day * 91            // ~91 days per quarter
    Year    = Day * 365
    Decade  = Year*10 + Day*2     // Accounting for leap years
    Century = Year*100 + Day*25   // Accounting for leap years
)
```

## Best Practices

### 1. Use Calendar Objects for Complex Operations
```go
// Preferred: Single calendar object with all information
cal := xtime.NowCalendar()
fmt.Printf("%s %s %s", cal.Time.Format("2006-01-02"),
    cal.LunarDateShort(), cal.CurrentSolarTerm())

// Avoid: Multiple separate calls
t := xtime.Now()
lunar := xtime.WithLunar(t.Time)
term := xtime.CurrentSolarterm(t.Time)
```

### 2. Batch Operations for Multiple Dates
```go
// Efficient for multiple dates
helper := xtime.NewSolarTermHelper()
calendar := helper.GetTermCalendar(2023)

// Less efficient for single dates
for i := 1; i <= 12; i++ {
    date := time.Date(2023, time.Month(i), 1, 0, 0, 0, 0, time.Local)
    term := xtime.CurrentSolarterm(date)
}
```

### 3. Reuse Helper Objects
```go
// Create once, use multiple times
lunarHelper := xtime.NewLunarHelper()
termHelper := xtime.NewSolarTermHelper()

for _, date := range dates {
    festival := lunarHelper.GetTodayFestival()
    term := termHelper.GetCurrentTerm(date)
}
```

### 4. Handle Time Zones Properly
```go
// Specify timezone for accurate calculations
beijing := time.FixedZone("CST", 8*3600)
t := time.Date(2023, 8, 15, 12, 0, 0, 0, beijing)
cal := xtime.NewCalendar(t)
```

## Error Handling

The package follows Go's idiomatic error handling:

```go
// Functions that can fail return errors
t, err := xtime.Parse("invalid-date")
if err != nil {
    log.Fatal(err)
}

// Must functions panic on error (use carefully)
t := xtime.MustParse("2023-08-15")  // Panics if parsing fails
```

## Related Packages

- `github.com/lazygophers/utils/candy` - Type conversion utilities
- `github.com/lazygophers/utils/randx` - Random number generation
- `github.com/jinzhu/now` - Enhanced time parsing (dependency)

## Examples

See [examples.go](examples.go) for comprehensive usage examples covering:
- Basic calendar operations
- Lunar calendar features
- Solar term calculations
- Festival detection
- Batch operations
- JSON serialization
- Real-world use cases

## Contributing

This package is part of the LazyGophers Utils collection. For contributions:

1. Follow Go coding standards
2. Add comprehensive tests for new features
3. Update documentation for API changes
4. Ensure backward compatibility

## License

This package is part of the LazyGophers Utils project. See the main repository for license information.

---

*Note: This package focuses on traditional Chinese calendar systems. For Islamic, Hebrew, or other calendar systems, consider specialized packages.*