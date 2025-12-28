---
title: xtime996 - 996 工作時間表
---

# xtime996 - 996 工作時間表

## 概述

xtime996 模組為 996 工作時間表提供時間常量（上午 9 點到晚上 9 點，每周 6 天）。它為此時間表定義工作時間計算。

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
)
```

### 天單位

```go
const (
    Day     = time.Hour * 24
    WorkDay = time.Hour * 12  // 12 小時工作日
    RestDay = Day - WorkDay    // 12 小時休息
)
```

### 周單位

```go
const (
    Week     = Day * 7
    WorkWeek = WorkDay * 6  // 6 個工作日
    RestWeek = Week - WorkWeek  // 1 天休息
)
```

### 月單位

```go
const (
    Month     = Day * 30
    RestMonth = RestDay * 4  // 4 天休息
    WorkMonth = Day*30 - RestMonth  // 26 個工作日
)
```

### 季度單位

```go
const (
    Quarter     = Day * 91
    RestQuarter = RestDay * 14  // 14 天休息
    WorkQuarter = Day*91 - RestQuarter  // 77 個工作日
)
```

### 年單位

```go
const (
    Year     = Day * 365
    RestYear = RestDay * 58  // 58 天休息
    WorkYear = Year - RestYear  // 307 個工作日
)
```

---

## 使用模式

### 計算工作時間

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

### 檢查工作時間

```go
func isWorkTime(t time.Time) bool {
    hour := t.Hour()
    weekday := t.Weekday()
    
    // 996: 上午 9 點到晚上 9 點，周一到周六
    if weekday == time.Sunday {
        return false
    }
    
    return hour >= 9 && hour < 21
}
```

### 計算工作日

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

### 工作時長計算

```go
func calculateWorkDuration(start, end time.Time) time.Duration {
    duration := end.Sub(start)
    
    // 過濾掉休息日
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

## 時間表摘要

### 996 工作時間表

- **工作日**: 周一到周六（6 天）
- **休息日**: 周日（1 天）
- **工作時間**: 上午 9:00 到晚上 9:00（12 小時）
- **休息時間**: 晚上 9:00 到上午 9:00（12 小時）

### 月度分解

- **總天數**: 30 天
- **工作日**: 26 天
- **休息日**: 4 天
- **工作時間**: 312 小時（26 × 12）
- **休息時間**: 408 小時（4 × 24 + 26 × 12）

### 年度分解

- **總天數**: 365 天
- **工作日**: 307 天
- **休息日**: 58 天
- **工作時間**: 3,684 小時（307 × 12）
- **休息時間**: 5,196 小時（58 × 24 + 307 × 12）

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
        // 找到下一個工作時間
        for !isWorkTime(now) {
            now = now.Add(time.Hour)
        }
    }
    
    end := now.Add(duration)
    
    // 檢查任務是否適合工作日
    if end.Hour() >= 21 || end.Weekday() == time.Sunday {
        return time.Time{}, fmt.Errorf("任務超過工作時間")
    }
    
    return now, nil
}
```

---

## 相關文檔

- [xtime](/zh-TW/modules/xtime) - 高級時間處理
- [xtime955](/zh-TW/modules/xtime955) - 955 工作時間表
- [xtime007](/zh-TW/modules/xtime007) - 24/7 運營
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
