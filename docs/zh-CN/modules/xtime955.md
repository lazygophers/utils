---
title: xtime955 - 955 工作时间表
---

# xtime955 - 955 工作时间表

## 概述

xtime955 模块为 955 工作时间表提供时间常量（上午 9 点到下午 5 点，每周 5 天）。它为此时间表定义工作时间计算。

## 时间常量

### 基本时间单位

```go
const (
    Nanosecond  = time.Nanosecond
    Microsecond = time.Microsecond
    Millisecond = time.Millisecond
    Second      = time.Second
    Minute      = time.Minute
)
```

### 工作时间周期

```go
const (
    HalfHour = time.Minute * 30
    Hour     = time.Hour
    Day      = time.Hour * 24
    WorkDay  = time.Hour * 8  // 8 小时工作日
    RestDay  = Day - WorkDay  // 16 小时休息
)
```

### 周周期

```go
const (
    Week     = Day * 7
    WorkWeek = WorkDay * 5  // 5 个工作日
    RestWeek = Week - WorkWeek  // 2 天休息
)
```

### 季度周期

```go
const (
    Month     = Day * 30
    WorkMonth = Day * 22  // 22 个工作日
    RestMonth = Month - WorkMonth  // 8 天休息
)
```

### 年周期

```go
const (
    Quarter     = Day * 91
    WorkQuarter = WorkMonth * 3  // 66 个工作日
    RestQuarter = Quarter - WorkQuarter  // 25 天休息
)
```

---

## 使用模式

### 计算工作时间

```go
func calculateWeeklyWorkHours() time.Duration {
    return xtime955.WorkWeek
}

func calculateMonthlyWorkHours() time.Duration {
    return xtime955.WorkMonth
}

func calculateYearlyWorkHours() time.Duration {
    return xtime955.WorkYear
}
```

### 检查工作时间

```go
func isWorkTime(t time.Time) bool {
    hour := t.Hour()
    weekday := t.Weekday()
    
    // 955: 上午 9 点到下午 5 点，周一到周五
    if weekday == time.Saturday || weekday == time.Sunday {
        return false
    }
    
    return hour >= 9 && hour < 17
}
```

### 计算工作日

```go
func getWorkDaysInMonth(year int, month time.Month) int {
    firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
    lastDay := firstDay.AddDate(0, 1, -1)
    
    workDays := 0
    for d := firstDay; d.Before(lastDay); d = d.AddDate(0, 0, 1) {
        weekday := d.Weekday()
        if weekday != time.Saturday && weekday != time.Sunday {
            workDays++
        }
    }
    
    return workDays
}
```

---

## 时间表摘要

### 955 工作时间表

- **工作日**: 周一到周五（5 天）
- **休息日**: 周六和周日（2 天）
- **工作时间**: 上午 9:00 到下午 5:00（8 小时）
- **休息时间**: 下午 5:00 到上午 9:00（16 小时）

### 月度分解

- **总天数**: 30 天
- **工作日**: 22 天
- **休息日**: 8 天
- **工作时间**: 176 小时（22 × 8）
- **休息时间**: 544 小时（8 × 24 + 22 × 16）

### 年度分解

- **总天数**: 365 天
- **工作日**: 250 天
- **休息日**: 115 天
- **工作时间**: 2,000 小时（250 × 8）
- **休息时间**: 6,760 小时（115 × 24 + 250 × 16）

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
        // 找到下一个工作日
        for !isWorkTime(now) {
            now = now.Add(time.Hour)
        }
    }
    
    end := now.Add(duration)
    
    // 检查任务是否适合工作日
    if end.Hour() >= 17 || end.Weekday() == time.Saturday || end.Weekday() == time.Sunday {
        return time.Time{}, fmt.Errorf("任务超过工作时间")
    }
    
    return now, nil
}
```

---

## 相关文档

- [xtime](/zh-CN/modules/xtime) - 高级时间处理
- [xtime996](/zh-CN/modules/xtime996) - 996 工作时间表
- [xtime007](/zh-CN/modules/xtime007) - 24/7 运营
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
