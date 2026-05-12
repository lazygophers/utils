package xtime

import (
	"time"
)

// newCalendarLazy 延迟计算版本（仅创建基础结构，不计算 Zodiac/Season）
// 适用场景：仅需要公历和农历基础信息
func newCalendarLazy(t time.Time) *Calendar {
	return &Calendar{
		Time:  With(t),
		lunar: WithLunar(t),
		// zodiac 和 season 懒加载，首次访问时计算
	}
}

// newCalendarCachedZodiac 缓存 Zodiac 计算版本
// 适用场景：频繁调用，且时间跨度不大
func newCalendarCachedZodiac(t time.Time) *Calendar {
	cal := &Calendar{
		Time:  With(t),
		lunar: WithLunar(t),
	}

	// Zodiac 计算可以缓存（相同年份）
	cal.zodiac = cal.calculateZodiac()
	cal.season = cal.calculateSeason() // 保留原 Season 计算（含节气查询）

	return cal
}

// newCalendarSimpleSeason 简化 Season 计算版本
// 移除节气查询，仅基于月份计算季节
// 适用场景：不需要精确节气信息
func newCalendarSimpleSeason(t time.Time) *Calendar {
	cal := &Calendar{
		Time:  With(t),
		lunar: WithLunar(t),
	}

	cal.zodiac = cal.calculateZodiac()
	cal.season = cal.calculateSeasonSimple() // 简化版 Season

	return cal
}

// newCalendarMinimal 完全简化版本
// 仅包含公历、农历、生肖（不含干支、节气、季节）
// 适用场景：只需要基础日历信息
func newCalendarMinimal(t time.Time) *Calendar {
	cal := &Calendar{
		Time:  With(t),
		lunar: WithLunar(t),
	}

	// 仅计算生肖（最快）
	cal.zodiac = cal.calculateZodiacMinimal()
	// season 使用零值

	return cal
}

// newCalendarPrealloc 预分配版本
// 优化内存分配和数组查找
func newCalendarPrealloc(t time.Time) *Calendar {
	cal := &Calendar{
		Time:  With(t),
		lunar: WithLunar(t),
	}

	cal.zodiac = cal.calculateZodiac()
	cal.season = cal.calculateSeasonPrealloc()

	return cal
}

// calculateSeasonSimple 简化版 Season 计算（移除节气查询）
func (c *Calendar) calculateSeasonSimple() SeasonInfo {
	now := c.Time.Time
	month := int(now.Month())
	day := now.Day()

	var season string
	var seasonProgress float64
	var yearProgress float64

	// 季节计算
	switch {
	case month >= 3 && month <= 5:
		season = "春"
		seasonProgress = float64(month-3)/3.0 + float64(day)/90.0
	case month >= 6 && month <= 8:
		season = "夏"
		seasonProgress = float64(month-6)/3.0 + float64(day)/90.0
	case month >= 9 && month <= 11:
		season = "秋"
		seasonProgress = float64(month-9)/3.0 + float64(day)/90.0
	default: // 12, 1, 2
		season = "冬"
		if month == 12 {
			seasonProgress = float64(day)/90.0
		} else {
			seasonProgress = float64(month+3)/3.0 + float64(day)/90.0
		}
	}

	// 年度进度
	yearStart := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	yearEnd := time.Date(now.Year()+1, 1, 1, 0, 0, 0, 0, time.Local)
	yearProgress = float64(now.Sub(yearStart)) / float64(yearEnd.Sub(yearStart))

	// 简化节气名称（基于月份）
	solarTerms := [][]string{
		{"小寒", "大寒"},                    // 1月
		{"立春", "雨水"},                    // 2月
		{"惊蛰", "春分"},                    // 3月
		{"清明", "谷雨"},                    // 4月
		{"立夏", "小满"},                    // 5月
		{"芒种", "夏至"},                    // 6月
		{"小暑", "大暑"},                    // 7月
		{"立秋", "处暑"},                    // 8月
		{"白露", "秋分"},                    // 9月
		{"寒露", "霜降"},                    // 10月
		{"立冬", "小雪"},                    // 11月
		{"大雪", "冬至"},                    // 12月
	}
	termIdx := (month - 1)
	if termIdx < 0 {
		termIdx = 11
	}
	dayIdx := day / 16
	if dayIdx > 1 {
		dayIdx = 1
	}

	return SeasonInfo{
		CurrentTerm:    solarTerms[termIdx][dayIdx],
		NextTerm:       "", // 简化版不计算
		NextTermTime:   time.Time{}, // 零值
		Season:         season,
		SeasonProgress: seasonProgress,
		YearProgress:   yearProgress,
	}
}

// calculateZodiacMinimal 最小化版 Zodiac 计算（仅生肖）
func (c *Calendar) calculateZodiacMinimal() ZodiacInfo {
	// 仅计算生肖
	return ZodiacInfo{
		Animal: c.lunar.Animal(),
		// 其他字段留空
	}
}

// calculateSeasonPrealloc 预分配版 Season 计算
func (c *Calendar) calculateSeasonPrealloc() SeasonInfo {
	now := c.Time.Time
	nextSolarterm := NextSolarterm(now)
	nextTermTime := nextSolarterm.Time()

	month := int(now.Month())
	day := now.Day()

	// 预分配当前节气数组（避免每次创建）
	var currentTerm, season string
	var seasonProgress float64
	var yearProgress float64

	// 季节判断（优化：使用整数比较）
	switch {
	case month >= 2 && month <= 4:
		season = "春"
		idx := (month-2)*2 + day/15
		if idx > 5 {
			idx = 5
		}
		currentTerm = springTerms[idx]
		seasonProgress = float64(month-2)/3.0 + float64(day)/90.0
	case month >= 5 && month <= 7:
		season = "夏"
		idx := (month-5)*2 + day/15
		if idx > 5 {
			idx = 5
		}
		currentTerm = summerTerms[idx]
		seasonProgress = float64(month-5)/3.0 + float64(day)/90.0
	case month >= 8 && month <= 10:
		season = "秋"
		idx := (month-8)*2 + day/15
		if idx > 5 {
			idx = 5
		}
		currentTerm = autumnTerms[idx]
		seasonProgress = float64(month-8)/3.0 + float64(day)/90.0
	default: // 11, 12, 1
		season = "冬"
		if month == 11 {
			idx := day/15
			if idx > 1 {
				idx = 1
			}
			currentTerm = winterTerms1[idx]
		} else if month == 12 {
			idx := day/15
			if idx > 1 {
				idx = 1
			}
			currentTerm = winterTerms2[idx]
		} else { // 1月
			idx := day/15
			if idx > 1 {
				idx = 1
			}
			currentTerm = winterTerms3[idx]
		}
		if month == 12 || month == 1 {
			seasonProgress = float64(day+30) / 90.0
		} else {
			seasonProgress = float64(day) / 90.0
		}
	}

	// 年度进度
	yearStart := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	yearEnd := time.Date(now.Year()+1, 1, 1, 0, 0, 0, 0, time.Local)
	yearProgress = float64(now.Sub(yearStart)) / float64(yearEnd.Sub(yearStart))

	return SeasonInfo{
		CurrentTerm:    currentTerm,
		NextTerm:       nextSolarterm.String(),
		NextTermTime:   nextTermTime,
		Season:         season,
		SeasonProgress: seasonProgress,
		YearProgress:   yearProgress,
	}
}

// 预定义节气数组常量（避免运行时创建）
var (
	springTerms = []string{"立春", "雨水", "惊蛰", "春分", "清明", "谷雨"}
	summerTerms = []string{"立夏", "小满", "芒种", "夏至", "小暑", "大暑"}
	autumnTerms = []string{"立秋", "处暑", "白露", "秋分", "寒露", "霜降"}
	winterTerms1 = []string{"立冬", "小雪"} // 11月
	winterTerms2 = []string{"大雪", "冬至"} // 12月
	winterTerms3 = []string{"小寒", "大寒"} // 1月
)
