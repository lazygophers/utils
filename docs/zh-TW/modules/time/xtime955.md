---
title: xtime955 - 955 工作時間表
---

# xtime955 - 955 工作時間表

## 概述

xtime955 模組為 955 工作時間表提供時間常量（上午 9 點到下午 5 點，每周 5 天）。它為此時間表定義工作時間計算。

## 時間常量

### 基本時間單位

```go
const (
    Nanosecond  = time.Nanosecond
    Microsecond = time.Microsecond
    Millisecond = time.Millisecond
    Second      = time.Second
    Minute      = time.Minute
)
```

### 工作時間周期

```go
const (
    HalfHour = time.Minute * 30
    Hour     = time.Hour
    Day      = time.Hour * 24
    WorkDay  = time.Hour * 8  // 8 小時工作日
    RestDay  = Day - WorkDay  // 16 小時休息
)
```

### 周周期

```go
const (
    Week     = Day * 7
    WorkWeek = WorkDay * 5  // 5 個工作日
    RestWeek = Week - WorkWeek  // 2 天休息
)
```

### 季度周期

```go
const (
    Month     = Day * 30
    WorkMonth = Day * 22  // 22 個工作日
    RestMonth = Month - WorkMonth  // 8 天休息
)
```

### 年周期

```go
const (
    Quarter     = Day * 91
    WorkQuarter = WorkMonth * 3  // 66 個工作日
    RestQuarter = Quarter - WorkQuarter  // 25 天休息
)
```

---

## 使用模式

### 計算工作時間

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

### 檢查工作時間

```go
func isWorkTime(t time.Time) bool {
    hour := t.Hour()
    weekday := t.Weekday()
    
    // 955: 上午 9 點到下午 5 點，周一到周五
    if weekday == time.Saturday || weekday == time.Sunday {
        return false
    }
    
    return hour >= 9 && hour < 17
}
```

### 計算工作日

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

## 時間表摘要

### 955 工作時間表

- **工作日**: 周一到周五（5 天）
- **休息日**: 周六和周日（2 天）
- **工作時間**: 上午 9:00 到下午 5:00（8 小時）
- **休息時間**: 下午 5:00 到上午 9:00（16 小時）

### 月度分解

- **總天數**: 30 天
- **工作日**: 22 天
- **休息日**: 8 天
- **工作時間**: 176 小時（22 × 8）
- **休息時間**: 544 小時（8 × 24 + 22 × 16）

### 年度分解

- **總天數**: 365 天
- **工作日**: 250 天
- **休息日**: 115 天
- **工作時間**: 2,000 小時（250 × 8）
- **休息時間**: 6,760 小時（115 × 24 + 250 × 16）

---

## 最佳實踐

### 工作時間驗證

```go
func validateWorkTime(start, end time.Time) error {
    if !isWorkTime(start) {
        return fmt.Errorf("開始時間不是工作時間")
    }
    
    if !isWorkTime(end) {
        return fmt.Errorf("結束時間不是工作時間")
    }
    
    if end.Before(start) {
        return fmt.Errorf("結束時間早於開始時間")
    }
    
    return nil
}
```

### 時間表規劃

```go
func planWorkTask(duration time.Duration) (time.Time, error) {
    now := time.Now()
    
    if !isWorkTime(now) {
        // 找到下一個工作日
        for !isWorkTime(now) {
            now = now.Add(time.Hour)
        }
    }
    
    end := now.Add(duration)
    
    // 檢查任務是否適合工作日
    if end.Hour() >= 17 || end.Weekday() == time.Saturday || end.Weekday() == time.Sunday {
        return time.Time{}, fmt.Errorf("任務超過工作時間")
    }
    
    return now, nil
}
```

---

## 相關文檔

- [xtime](/zh-TW/modules/xtime) - 高級時間處理
- [xtime996](/zh-TW/modules/xtime996) - 996 工作時間表
- [xtime007](/zh-TW/modules/xtime007) - 24/7 運營
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
