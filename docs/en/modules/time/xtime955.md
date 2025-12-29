---
title: xtime955 - 955 Work Schedule
---

# xtime955 - 955 Work Schedule

## Overview

The xtime955 module provides time constants for the 955 work schedule (9 AM to 5 PM, 5 days per week). It defines work time calculations for this schedule.

## Time Constants

### Basic Time Units

```go
const (
    Nanosecond  = time.Nanosecond
    Microsecond = time.Microsecond
    Millisecond = time.Millisecond
    Second      = time.Second
    Minute      = time.Minute
)
```

### Work Time Cycle

```go
const (
    HalfHour = time.Minute * 30
    Hour     = time.Hour
    Day      = time.Hour * 24
    WorkDay  = time.Hour * 8  // 8 hours work day
    RestDay  = Day - WorkDay  // 16 hours rest
)
```

### Week Cycle

```go
const (
    Week     = Day * 7
    WorkWeek = WorkDay * 5  // 5 work days
    RestWeek = Week - WorkWeek  // 2 days rest
)
```

### Quarter Cycle

```go
const (
    Month     = Day * 30
    WorkMonth = Day * 22  // 22 work days
    RestMonth = Month - WorkMonth  // 8 days rest
)
```

### Year Cycle

```go
const (
    Quarter     = Day * 91
    WorkQuarter = WorkMonth * 3  // 66 work days
    RestQuarter = Quarter - WorkQuarter  // 25 days rest
)
```

---

## Usage Patterns

### Calculate Work Hours

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

### Check Work Time

```go
func isWorkTime(t time.Time) bool {
    hour := t.Hour()
    weekday := t.Weekday()
    
    // 955: 9 AM to 5 PM, Monday to Friday
    if weekday == time.Saturday || weekday == time.Sunday {
        return false
    }
    
    return hour >= 9 && hour < 17
}
```

### Calculate Work Days

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

## Schedule Summary

### 955 Work Schedule

- **Work Days**: Monday to Friday (5 days)
- **Rest Days**: Saturday and Sunday (2 days)
- **Work Hours**: 9:00 AM to 5:00 PM (8 hours)
- **Rest Hours**: 5:00 PM to 9:00 AM (16 hours)

### Monthly Breakdown

- **Total Days**: 30 days
- **Work Days**: 22 days
- **Rest Days**: 8 days
- **Work Hours**: 176 hours (22 × 8)
- **Rest Hours**: 544 hours (8 × 24 + 22 × 16)

### Yearly Breakdown

- **Total Days**: 365 days
- **Work Days**: 250 days
- **Rest Days**: 115 days
- **Work Hours**: 2,000 hours (250 × 8)
- **Rest Hours**: 6,760 hours (115 × 24 + 250 × 16)

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
        // Find next work day
        for !isWorkTime(now) {
            now = now.Add(time.Hour)
        }
    }
    
    end := now.Add(duration)
    
    // Check if task fits in work day
    if end.Hour() >= 17 || end.Weekday() == time.Saturday || end.Weekday() == time.Sunday {
        return time.Time{}, fmt.Errorf("task exceeds work hours")
    }
    
    return now, nil
}
```

---

## Related Documentation

- [xtime](/en/modules/xtime) - Advanced time processing
- [xtime996](/en/modules/xtime996) - 996 work schedule
- [xtime007](/en/modules/xtime007) - 24/7 operations
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
