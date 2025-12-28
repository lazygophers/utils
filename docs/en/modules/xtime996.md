---
title: xtime996 - 996 Work Schedule
---

# xtime996 - 996 Work Schedule

## Overview

The xtime996 module provides time constants for the 996 work schedule (9 AM to 9 PM, 6 days per week). It defines work time calculations for this schedule.

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
)
```

### Day Units

```go
const (
    Day     = time.Hour * 24
    WorkDay = time.Hour * 12  // 12 hours work day
    RestDay = Day - WorkDay    // 12 hours rest
)
```

### Week Units

```go
const (
    Week     = Day * 7
    WorkWeek = WorkDay * 6  // 6 work days
    RestWeek = Week - WorkWeek  // 1 day rest
)
```

### Month Units

```go
const (
    Month     = Day * 30
    RestMonth = RestDay * 4  // 4 rest days
    WorkMonth = Day*30 - RestMonth  // 26 work days
)
```

### Quarter Units

```go
const (
    Quarter     = Day * 91
    RestQuarter = RestDay * 14  // 14 rest days
    WorkQuarter = Day*91 - RestQuarter  // 77 work days
)
```

### Year Units

```go
const (
    Year     = Day * 365
    RestYear = RestDay * 58  // 58 rest days
    WorkYear = Year - RestYear  // 307 work days
)
```

---

## Usage Patterns

### Calculate Work Hours

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

### Check Work Time

```go
func isWorkTime(t time.Time) bool {
    hour := t.Hour()
    weekday := t.Weekday()
    
    // 996: 9 AM to 9 PM, Monday to Saturday
    if weekday == time.Sunday {
        return false
    }
    
    return hour >= 9 && hour < 21
}
```

### Calculate Work Days

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

### Work Duration Calculation

```go
func calculateWorkDuration(start, end time.Time) time.Duration {
    duration := end.Sub(start)
    
    // Filter out rest days
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

## Schedule Summary

### 996 Work Schedule

- **Work Days**: Monday to Saturday (6 days)
- **Rest Days**: Sunday (1 day)
- **Work Hours**: 9:00 AM to 9:00 PM (12 hours)
- **Rest Hours**: 9:00 PM to 9:00 AM (12 hours)

### Monthly Breakdown

- **Total Days**: 30 days
- **Work Days**: 26 days
- **Rest Days**: 4 days
- **Work Hours**: 312 hours (26 × 12)
- **Rest Hours**: 408 hours (4 × 24 + 26 × 12)

### Yearly Breakdown

- **Total Days**: 365 days
- **Work Days**: 307 days
- **Rest Days**: 58 days
- **Work Hours**: 3,684 hours (307 × 12)
- **Rest Hours**: 5,196 hours (58 × 24 + 307 × 12)

---

## Best Practices

### Work Time Validation

```go
func validateWorkTime(start, end time.Time) error {
    if !isWorkTime(start) {
        return fmt.Errorf("start time is not work time")
    }
    
    if !isWorkTime(end) {
        return fmt.Errorf("end time is not work time")
    }
    
    if end.Before(start) {
        return fmt.Errorf("end time is before start time")
    }
    
    return nil
}
```

### Schedule Planning

```go
func planWorkTask(duration time.Duration) (time.Time, error) {
    now := time.Now()
    
    if !isWorkTime(now) {
        // Find next work time
        for !isWorkTime(now) {
            now = now.Add(time.Hour)
        }
    }
    
    end := now.Add(duration)
    
    // Check if task fits in work day
    if end.Hour() >= 21 || end.Weekday() == time.Sunday {
        return time.Time{}, fmt.Errorf("task exceeds work hours")
    }
    
    return now, nil
}
```

---

## Related Documentation

- [xtime](/en/modules/xtime) - Advanced time processing
- [xtime955](/en/modules/xtime955) - 955 work schedule
- [xtime007](/en/modules/xtime007) - 24/7 operations
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
