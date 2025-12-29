---
title: xtime996 - 996 工作时间表
---

# xtime996 - 996 工作时间表

## 概述

xtime996 模块为 996 工作时间表提供时间常量（上午 9 点到晚上 9 点，每周 6 天）。它为此时间表定义工作时间计算。

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
)
```

### 天单位

```go
const (
    Day     = time.Hour * 24
    WorkDay = time.Hour * 12  // 12 小时工作日
    RestDay = Day - WorkDay    // 12 小时休息
)
```

### 周单位

```go
const (
    Week     = Day * 7
    WorkWeek = WorkDay * 6  // 6 个工作日
    RestWeek = Week - WorkWeek  // 1 天休息
)
```

### 月单位

```go
const (
    Month     = Day * 30
    RestMonth = RestDay * 4  // 4 天休息
    WorkMonth = Day*30 - RestMonth  // 26 个工作日
)
```

### 季度单位

```go
const (
    Quarter     = Day * 91
    RestQuarter = RestDay * 14  // 14 天休息
    WorkQuarter = Day*91 - RestQuarter  // 77 个工作日
)
```

### 年单位

```go
const (
    Year     = Day * 365
    RestYear = RestDay * 58  // 58 天休息
    WorkYear = Year - RestYear  // 307 个工作日
)
```

---

## 使用模式

### 计算工作时间

```go
func calculateWeeklyWorkHours() time.Duration {
    return xtime996.WorkWeek
}

func calculateMonthlyWorkHours() time.Duration {
    return xtime996.WorkMonth
}

func calculateYearlyWorkHours() time.Duration {
    return xtime996.WorkYear
}
```

### 检查工作时间

```go
func isWorkTime(t time.Time) bool {
    hour := t.Hour()
    weekday := t.Weekday()
    
    // 996: 上午 9 点到晚上 9 点，周一到周六
    if weekday == time.Sunday {
        return false
    }
    
    return hour >= 9 && hour < 21
}
```

### 计算工作日

```go
func getWorkDaysInMonth(year int, month time.Month) int {
    firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
    lastDay := firstDay.AddDate(0, 1, -1)
    
    workDays := 0
    for d := firstDay; d.Before(lastDay); d = d.AddDate(0, 0, 1) {
        if d.Weekday() != time.Sunday {
            workDays++
        }
    }
    
    return workDays
}
```

### 工作时长计算

```go
func calculateWorkDuration(start, end time.Time) time.Duration {
    duration := end.Sub(start)
    
    // 过滤掉休息日
    workDuration := time.Duration(0)
    current := start
    
    for current.Before(end) {
        if isWorkTime(current) {
            workDuration += time.Hour
        }
        current = current.Add(time.Hour)
    }
    
    return workDuration
}
```

---

## 时间表摘要

### 996 工作时间表

- **工作日**: 周一到周六（6 天）
- **休息日**: 周日（1 天）
- **工作时间**: 上午 9:00 到晚上 9:00（12 小时）
- **休息时间**: 晚上 9:00 到上午 9:00（12 小时）

### 月度分解

- **总天数**: 30 天
- **工作日**: 26 天
- **休息日**: 4 天
- **工作时间**: 312 小时（26 × 12）
- **休息时间**: 408 小时（4 × 24 + 26 × 12）

### 年度分解

- **总天数**: 365 天
- **工作日**: 307 天
- **休息日**: 58 天
- **工作时间**: 3,684 小时（307 × 12）
- **休息时间**: 5,196 小时（58 × 24 + 307 × 12）

---

## 最佳实践

### 工作时间验证

```go
func validateWorkTime(start, end time.Time) error {
    if !isWorkTime(start) {
        return fmt.Errorf("开始时间不是工作时间")
    }
    
    if !isWorkTime(end) {
        return fmt.Errorf("结束时间不是工作时间")
    }
    
    if end.Before(start) {
        return fmt.Errorf("结束时间早于开始时间")
    }
    
    return nil
}
```

### 时间表规划

```go
func planWorkTask(duration time.Duration) (time.Time, error) {
    now := time.Now()
    
    if !isWorkTime(now) {
        // 找到下一个工作时间
        for !isWorkTime(now) {
            now = now.Add(time.Hour)
        }
    }
    
    end := now.Add(duration)
    
    // 检查任务是否适合工作日
    if end.Hour() >= 21 || end.Weekday() == time.Sunday {
        return time.Time{}, fmt.Errorf("任务超过工作时间")
    }
    
    return now, nil
}
```

---

## 相关文档

- [xtime](/zh-CN/modules/xtime) - 高级时间处理
- [xtime955](/zh-CN/modules/xtime955) - 955 工作时间表
- [xtime007](/zh-CN/modules/xtime007) - 24/7 运营
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
