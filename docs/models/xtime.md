# Xtime Package Documentation

<!-- Language selector -->
[🇺🇸 English](#english) | [🇨🇳 简体中文](#简体中文) | [🇭🇰 繁體中文](#繁體中文) | [🇷🇺 Русский](#русский) | [🇫🇷 Français](#français) | [🇸🇦 العربية](#العربية) | [🇪🇸 Español](#español)

---

## English

### Overview
The `xtime` package provides enhanced time operations with lunar calendar support, Chinese traditional calendar features, and advanced time manipulation utilities. It extends Go's standard `time` package with additional functionality for Asian calendar systems and improved time range operations.

### Key Features
- **Enhanced Time Type**: Wrapper around Go's time.Time with additional methods
- **Lunar Calendar Support**: Chinese lunar calendar calculations and conversions
- **Solar Terms**: Traditional Chinese 24 solar terms calculations
- **Zodiac Animals**: Chinese zodiac year animals and elements
- **Time Range Operations**: Beginning/End of various time periods
- **Flexible Parsing**: Enhanced time parsing with multiple formats
- **Random Sleep**: Utility for random duration sleeps

### Core Components

#### Time Configuration
```go
type Config struct {
    WeekStartDay time.Weekday
    TimeLocation *time.Location
    TimeFormats  []string
}

type Time struct {
    time.Time
    *Config
}
```

#### Basic Usage
```go
// Current time with enhanced features
now := xtime.Now()
fmt.Println(now.Format("2006-01-02 15:04:05"))

// Parse time with flexible formats
t, err := xtime.Parse("2023-08-15", "15:04:05")
t := xtime.MustParse("2023-08-15 14:30:00")

// Create enhanced time from standard time
stdTime := time.Now()
enhanced := xtime.With(stdTime)
```

### Time Range Operations

#### Beginning/End of Time Periods
```go
now := xtime.Now()

// Minute precision
minuteStart := now.BeginningOfMinute()   // 14:30:00.000
minuteEnd := now.EndOfMinute()           // 14:30:59.999

// Hour precision
hourStart := now.BeginningOfHour()       // 14:00:00.000
hourEnd := now.EndOfHour()               // 14:59:59.999

// Day precision
dayStart := now.BeginningOfDay()         // 00:00:00.000
dayEnd := now.EndOfDay()                 // 23:59:59.999

// Week precision (configurable start day)
weekStart := now.BeginningOfWeek()       // Monday 00:00:00
weekEnd := now.EndOfWeek()               // Sunday 23:59:59

// Month precision
monthStart := now.BeginningOfMonth()     // 1st day 00:00:00
monthEnd := now.EndOfMonth()             // Last day 23:59:59

// Quarter precision
quarterStart := now.BeginningOfQuarter() // Q1: Jan 1, Q2: Apr 1, etc.
quarterEnd := now.EndOfQuarter()         // Q1: Mar 31, Q2: Jun 30, etc.

// Year precision
yearStart := now.BeginningOfYear()       // Jan 1 00:00:00
yearEnd := now.EndOfYear()               // Dec 31 23:59:59
```

#### Global Time Range Functions
```go
// Current time ranges (without creating Time instance)
today := xtime.BeginningOfDay()          // Today at 00:00:00
thisWeek := xtime.BeginningOfWeek()      // This week's start
thisMonth := xtime.BeginningOfMonth()    // This month's start
thisYear := xtime.BeginningOfYear()      // This year's start

// Current time endings
todayEnd := xtime.EndOfDay()             // Today at 23:59:59
weekEnd := xtime.EndOfWeek()             // This week's end
monthEnd := xtime.EndOfMonth()           // This month's end
yearEnd := xtime.EndOfYear()             // This year's end

// Quarter information
currentQuarter := xtime.Quarter()        // Returns 1, 2, 3, or 4
```

### Lunar Calendar Features

#### Lunar Calendar Integration
```go
// Create calendar with lunar support
cal := xtime.NowCalendar()

// Lunar date information
lunarDate := cal.LunarDate()             // "农历二零二三年六月廿九"
lunarMonth := cal.LunarMonth()           // Lunar month number
lunarDay := cal.LunarDay()               // Lunar day number

// Zodiac information
animal := cal.Animal()                   // "兔" (Rabbit for 2023)
element := cal.Element()                 // "水" (Water element)

// Solar terms
currentTerm := cal.CurrentSolarTerm()    // "处暑" (End of Heat)
nextTerm := cal.NextSolarTerm()          // "白露" (White Dew)
```

#### Traditional Chinese Features
```go
// Heavenly Stems and Earthly Branches
stems := cal.HeavenlyStems()             // 天干
branches := cal.EarthlyBranches()        // 地支

// Traditional festivals detection
isFestival := cal.IsTraditionalFestival()
festivalName := cal.GetFestivalName()   // "中秋节", "春节", etc.

// Lucky/Unlucky day calculations
isLuckyDay := cal.IsLuckyDay()
dayScore := cal.GetDayScore()            // 1-10 scale
```

### Advanced Usage Patterns

#### Time Range Queries
```go
func getDataInRange(start, end time.Time) []Data {
    startTime := xtime.With(start).BeginningOfDay()
    endTime := xtime.With(end).EndOfDay()
    
    return database.Query(
        "SELECT * FROM data WHERE created_at BETWEEN ? AND ?",
        startTime.Time, endTime.Time,
    )
}

// Week-based reporting
func generateWeeklyReport(date time.Time) Report {
    week := xtime.With(date)
    start := week.BeginningOfWeek()
    end := week.EndOfWeek()
    
    return Report{
        Period: fmt.Sprintf("%s - %s", start.Format("Jan 2"), end.Format("Jan 2")),
        Data:   getDataInRange(start.Time, end.Time),
    }
}
```

#### Business Time Calculations
```go
// Quarter-based business logic
func getQuarterlyGoals(year int) map[int]Goals {
    goals := make(map[int]Goals)
    
    for q := 1; q <= 4; q++ {
        // Create time for start of quarter
        qStart := time.Date(year, time.Month((q-1)*3+1), 1, 0, 0, 0, 0, time.Local)
        quarter := xtime.With(qStart)
        
        goals[q] = Goals{
            Quarter: q,
            Start:   quarter.BeginningOfQuarter(),
            End:     quarter.EndOfQuarter(),
            Target:  calculateQuarterlyTarget(q),
        }
    }
    
    return goals
}
```

#### Time-based Cache Keys
```go
// Generate time-based cache keys
func getCacheKey(operation string, t time.Time) string {
    timeKey := xtime.With(t)
    
    switch operation {
    case "hourly":
        return fmt.Sprintf("%s:hour:%s", operation, timeKey.BeginningOfHour().Format("2006010215"))
    case "daily":
        return fmt.Sprintf("%s:day:%s", operation, timeKey.BeginningOfDay().Format("20060102"))
    case "weekly":
        return fmt.Sprintf("%s:week:%s", operation, timeKey.BeginningOfWeek().Format("200601"))
    case "monthly":
        return fmt.Sprintf("%s:month:%s", operation, timeKey.BeginningOfMonth().Format("200601"))
    default:
        return fmt.Sprintf("%s:%d", operation, t.Unix())
    }
}
```

### Utility Functions

#### Random Sleep
```go
// Random sleep for testing or rate limiting
xtime.RandSleep()                        // Random duration up to 1 second
xtime.RandSleep(5*time.Second)           // Random duration up to 5 seconds
xtime.RandSleep(1*time.Second, 5*time.Second) // Random between 1-5 seconds

// Usage in retry logic
func retryWithBackoff(operation func() error, maxRetries int) error {
    for i := 0; i < maxRetries; i++ {
        if err := operation(); err == nil {
            return nil
        }
        
        if i < maxRetries-1 {
            // Exponential backoff with jitter
            baseDelay := time.Duration(i+1) * time.Second
            xtime.RandSleep(baseDelay, baseDelay*2)
        }
    }
    return fmt.Errorf("operation failed after %d retries", maxRetries)
}
```

#### Unix Timestamp Utilities
```go
// Get current timestamps
unixSeconds := xtime.NowUnix()           // Current Unix timestamp
unixMillis := xtime.NowUnixMilli()       // Current Unix timestamp in milliseconds

// Time comparisons and calculations
now := xtime.Now()
tomorrow := now.AddDate(0, 0, 1)
daysDiff := int(tomorrow.Sub(now.Time).Hours() / 24)
```

### Configuration and Customization

#### Week Start Day Configuration
```go
// Configure week to start on Sunday (default is Monday)
t := xtime.Now()
t.WeekStartDay = time.Sunday

sundayWeekStart := t.BeginningOfWeek()   // Week starts on Sunday
sundayWeekEnd := t.EndOfWeek()           // Week ends on Saturday
```

#### Time Zone Handling
```go
// Work with different time zones
utc := time.UTC
est := time.FixedZone("EST", -5*60*60)

t := xtime.Now()
t.TimeLocation = utc

utcTime := t.In(utc)
estTime := t.In(est)
```

### Integration Examples

#### Web API Time Ranges
```go
// HTTP handler for time range queries
func handleTimeRangeQuery(w http.ResponseWriter, r *http.Request) {
    rangeType := r.URL.Query().Get("range") // "day", "week", "month", "quarter", "year"
    
    now := xtime.Now()
    var start, end *xtime.Time
    
    switch rangeType {
    case "day":
        start, end = now.BeginningOfDay(), now.EndOfDay()
    case "week":
        start, end = now.BeginningOfWeek(), now.EndOfWeek()
    case "month":
        start, end = now.BeginningOfMonth(), now.EndOfMonth()
    case "quarter":
        start, end = now.BeginningOfQuarter(), now.EndOfQuarter()
    case "year":
        start, end = now.BeginningOfYear(), now.EndOfYear()
    default:
        http.Error(w, "Invalid range type", http.StatusBadRequest)
        return
    }
    
    data := fetchData(start.Time, end.Time)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "range": rangeType,
        "start": start.Format(time.RFC3339),
        "end":   end.Format(time.RFC3339),
        "data":  data,
    })
}
```

### Best Practices
1. **Use Enhanced Time Consistently**: Prefer `xtime.Time` over `time.Time` for enhanced functionality
2. **Configure Week Start Day**: Set appropriate week start day for your locale/business needs
3. **Handle Time Zones Properly**: Always be explicit about time zones in business logic
4. **Use Appropriate Precision**: Choose the right time range function for your use case
5. **Lunar Calendar Features**: Leverage lunar calendar features for Asian market applications

### Common Patterns
```go
// Business hours calculation
func isBusinessHour(t time.Time) bool {
    enhanced := xtime.With(t)
    hour := enhanced.Hour()
    weekday := enhanced.Weekday()
    
    return weekday >= time.Monday && weekday <= time.Friday && 
           hour >= 9 && hour < 17
}

// Reporting period helper
type ReportPeriod struct {
    Name  string
    Start *xtime.Time
    End   *xtime.Time
}

func getReportingPeriods(baseTime time.Time) []ReportPeriod {
    t := xtime.With(baseTime)
    
    return []ReportPeriod{
        {"This Hour", t.BeginningOfHour(), t.EndOfHour()},
        {"Today", t.BeginningOfDay(), t.EndOfDay()},
        {"This Week", t.BeginningOfWeek(), t.EndOfWeek()},
        {"This Month", t.BeginningOfMonth(), t.EndOfMonth()},
        {"This Quarter", t.BeginningOfQuarter(), t.EndOfQuarter()},
        {"This Year", t.BeginningOfYear(), t.EndOfYear()},
    }
}
```

---

## 简体中文

### 概述
`xtime` 包提供增强的时间操作，支持农历、中国传统历法功能和高级时间操作工具。它扩展了 Go 标准 `time` 包，为亚洲历法系统和改进的时间范围操作提供额外功能。

### 主要特性
- **增强时间类型**: Go time.Time 的包装器，具有额外方法
- **农历支持**: 中国农历计算和转换
- **二十四节气**: 中国传统二十四节气计算
- **生肖动物**: 中国生肖年动物和五行元素
- **时间范围操作**: 各种时间段的开始/结束
- **灵活解析**: 支持多种格式的增强时间解析
- **随机睡眠**: 随机持续时间睡眠工具

### 核心组件

#### 时间配置
```go
type Config struct {
    WeekStartDay time.Weekday  // 周起始日
    TimeLocation *time.Location  // 时区位置
    TimeFormats  []string  // 时间格式
}

type Time struct {
    time.Time
    *Config
}
```

#### 基本使用
```go
// 当前时间与增强功能
now := xtime.Now()
fmt.Println(now.Format("2006-01-02 15:04:05"))

// 灵活格式解析时间
t, err := xtime.Parse("2023-08-15", "15:04:05")
t := xtime.MustParse("2023-08-15 14:30:00")

// 从标准时间创建增强时间
stdTime := time.Now()
enhanced := xtime.With(stdTime)
```

### 时间范围操作

#### 时间段的开始/结束
```go
now := xtime.Now()

// 分钟精度
minuteStart := now.BeginningOfMinute()   // 14:30:00.000
minuteEnd := now.EndOfMinute()           // 14:30:59.999

// 小时精度
hourStart := now.BeginningOfHour()       // 14:00:00.000
hourEnd := now.EndOfHour()               // 14:59:59.999

// 天精度
dayStart := now.BeginningOfDay()         // 00:00:00.000
dayEnd := now.EndOfDay()                 // 23:59:59.999

// 周精度（可配置起始日）
weekStart := now.BeginningOfWeek()       // 周一 00:00:00
weekEnd := now.EndOfWeek()               // 周日 23:59:59

// 月精度
monthStart := now.BeginningOfMonth()     // 1日 00:00:00
monthEnd := now.EndOfMonth()             // 最后一日 23:59:59

// 季度精度
quarterStart := now.BeginningOfQuarter() // Q1: 1月1日, Q2: 4月1日等
quarterEnd := now.EndOfQuarter()         // Q1: 3月31日, Q2: 6月30日等

// 年精度
yearStart := now.BeginningOfYear()       // 1月1日 00:00:00
yearEnd := now.EndOfYear()               // 12月31日 23:59:59
```

### 农历功能

#### 农历集成
```go
// 创建支持农历的日历
cal := xtime.NowCalendar()

// 农历日期信息
lunarDate := cal.LunarDate()             // "农历二零二三年六月廿九"
lunarMonth := cal.LunarMonth()           // 农历月份数字
lunarDay := cal.LunarDay()               // 农历日期数字

// 生肖信息
animal := cal.Animal()                   // "兔"（2023年的生肖）
element := cal.Element()                 // "水"（五行元素）

// 节气
currentTerm := cal.CurrentSolarTerm()    // "处暑"
nextTerm := cal.NextSolarTerm()          // "白露"
```

### 最佳实践
1. **一致使用增强时间**: 优先使用 `xtime.Time` 而非 `time.Time` 以获得增强功能
2. **配置周起始日**: 为您的地区/业务需求设置适当的周起始日
3. **正确处理时区**: 在业务逻辑中始终明确时区
4. **使用适当精度**: 为您的用例选择正确的时间范围函数

---

## 繁體中文

### 概述
`xtime` 套件提供增強的時間操作，支援農曆、中國傳統曆法功能和進階時間操作工具。它擴展了 Go 標準 `time` 套件，為亞洲曆法系統和改進的時間範圍操作提供額外功能。

### 主要特性
- **增強時間型別**: Go time.Time 的包裝器，具有額外方法
- **農曆支援**: 中國農曆計算和轉換
- **二十四節氣**: 中國傳統二十四節氣計算
- **生肖動物**: 中國生肖年動物和五行元素

### 核心組件
```go
type Time struct {
    time.Time
    *Config
}

// 當前時間與增強功能
now := xtime.Now()
fmt.Println(now.Format("2006-01-02 15:04:05"))
```

### 最佳實務
1. **一致使用增強時間**: 優先使用 `xtime.Time` 而非 `time.Time` 以獲得增強功能
2. **配置週起始日**: 為您的地區/業務需求設定適當的週起始日

---

## Русский

### Обзор
Пакет `xtime` предоставляет улучшенные операции со временем с поддержкой лунного календаря, функций китайского традиционного календаря и расширенных утилит манипуляции временем.

### Основные возможности
- **Улучшенный тип времени**: Обертка вокруг Go time.Time с дополнительными методами
- **Поддержка лунного календаря**: Вычисления и преобразования китайского лунного календаря
- **Солнечные термины**: Вычисления традиционных китайских 24 солнечных терминов
- **Животные зодиака**: Китайские животные зодиака и элементы

### Основные компоненты
```go
type Time struct {
    time.Time
    *Config
}

// Текущее время с улучшенными возможностями
now := xtime.Now()
fmt.Println(now.Format("2006-01-02 15:04:05"))
```

### Лучшие практики
1. **Последовательно используйте улучшенное время**: Предпочитайте `xtime.Time` над `time.Time` для расширенной функциональности
2. **Настройте день начала недели**: Установите подходящий день начала недели для вашего региона/бизнеса

---

## Français

### Aperçu
Le package `xtime` fournit des opérations de temps améliorées avec support du calendrier lunaire, des fonctionnalités de calendrier traditionnel chinois et des utilitaires avancés de manipulation du temps.

### Caractéristiques principales
- **Type de temps amélioré**: Wrapper autour de Go time.Time avec des méthodes supplémentaires
- **Support du calendrier lunaire**: Calculs et conversions du calendrier lunaire chinois
- **Termes solaires**: Calculs des 24 termes solaires traditionnels chinois
- **Animaux du zodiaque**: Animaux du zodiaque chinois et éléments

### Composants principaux
```go
type Time struct {
    time.Time
    *Config
}

// Temps actuel avec fonctionnalités améliorées
now := xtime.Now()
fmt.Println(now.Format("2006-01-02 15:04:05"))
```

### Meilleures pratiques
1. **Utilisez le temps amélioré de manière cohérente**: Préférez `xtime.Time` à `time.Time` pour une fonctionnalité étendue
2. **Configurez le jour de début de semaine**: Définissez le jour de début de semaine approprié pour votre région/besoins commerciaux

---

## العربية

### نظرة عامة
توفر حزمة `xtime` عمليات وقت محسنة مع دعم التقويم القمري، وميزات التقويم الصيني التقليدي، وأدوات متقدمة للتلاعب بالوقت.

### الميزات الرئيسية
- **نوع وقت محسن**: غلاف حول Go time.Time مع طرق إضافية
- **دعم التقويم القمري**: حسابات وتحويلات التقويم القمري الصيني
- **المصطلحات الشمسية**: حسابات المصطلحات الشمسية الصينية التقليدية البالغة 24
- **حيوانات البروج**: حيوانات البروج الصينية والعناصر

### المكونات الأساسية
```go
type Time struct {
    time.Time
    *Config
}

// الوقت الحالي مع الميزات المحسنة
now := xtime.Now()
fmt.Println(now.Format("2006-01-02 15:04:05"))
```

### أفضل الممارسات
1. **استخدم الوقت المحسن بثبات**: فضل `xtime.Time` على `time.Time` للوظائف الموسعة
2. **كوّن يوم بداية الأسبوع**: حدد يوم بداية الأسبوع المناسب لمنطقتك/احتياجات عملك

---

## Español

### Descripción general
El paquete `xtime` proporciona operaciones de tiempo mejoradas con soporte de calendario lunar, características de calendario tradicional chino y utilidades avanzadas de manipulación de tiempo.

### Características principales
- **Tipo de tiempo mejorado**: Wrapper alrededor de Go time.Time con métodos adicionales
- **Soporte de calendario lunar**: Cálculos y conversiones del calendario lunar chino
- **Términos solares**: Cálculos de los 24 términos solares tradicionales chinos
- **Animales del zodíaco**: Animales del zodíaco chino y elementos

### Componentes principales
```go
type Time struct {
    time.Time
    *Config
}

// Tiempo actual con características mejoradas
now := xtime.Now()
fmt.Println(now.Format("2006-01-02 15:04:05"))
```

### Mejores prácticas
1. **Use tiempo mejorado consistentemente**: Prefiera `xtime.Time` sobre `time.Time` para funcionalidad extendida
2. **Configure día de inicio de semana**: Establezca el día de inicio de semana apropiado para su región/necesidades de negocio