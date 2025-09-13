package xtime

import (
	"fmt"
	"time"
)

// Calendar 日历信息，包含公历和农历的综合信息
type Calendar struct {
	*Time             // 嵌入增强的时间类型
	lunar  *Lunar     // 农历信息
	zodiac ZodiacInfo // 生肖天干地支信息
	season SeasonInfo // 节气季节信息
}

// ZodiacInfo 生肖天干地支信息
type ZodiacInfo struct {
	Animal      string // 生肖：鼠、牛、虎...
	SkyTrunk    string // 天干：甲、乙、丙...
	EarthBranch string // 地支：子、丑、寅...
	YearGanZhi  string // 年干支：甲子、乙丑...
	MonthGanZhi string // 月干支
	DayGanZhi   string // 日干支
	HourGanZhi  string // 时干支
}

// SeasonInfo 节气季节信息
type SeasonInfo struct {
	CurrentTerm    string    // 当前节气
	NextTerm       string    // 下个节气
	NextTermTime   time.Time // 下个节气时间
	Season         string    // 季节：春、夏、秋、冬
	SeasonProgress float64   // 季节进度(0-1)
	YearProgress   float64   // 年度进度(0-1)
}

// 天干地支常量
var (
	skyTrunks     = [10]string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	earthBranches = [12]string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	seasons       = [4]string{"春", "夏", "秋", "冬"}
)

// NewCalendar 创建日历对象，包含完整的农历和节气信息
func NewCalendar(t time.Time) *Calendar {
	cal := &Calendar{
		Time:  With(t),
		lunar: WithLunar(t),
	}

	cal.zodiac = cal.calculateZodiac()
	cal.season = cal.calculateSeason()

	return cal
}

// NowCalendar 获取当前日历信息
func NowCalendar() *Calendar {
	return NewCalendar(time.Now())
}

// 农历相关方法

// Lunar 获取农历日期信息
func (c *Calendar) Lunar() *Lunar {
	return c.lunar
}

// LunarDate 获取农历日期字符串，格式：农历二零二三年八月十五
func (c *Calendar) LunarDate() string {
	year := c.lunar.YearAlias()
	month := c.lunar.MonthAlias()
	day := c.lunar.DayAlias()
	return fmt.Sprintf("农历%s年%s%s", year, month, day)
}

// LunarDateShort 获取简短农历日期，格式：八月十五
func (c *Calendar) LunarDateShort() string {
	return c.lunar.MonthAlias() + c.lunar.DayAlias()
}

// IsLunarLeapYear 是否农历闰年
func (c *Calendar) IsLunarLeapYear() bool {
	return c.lunar.IsLeap()
}

// LunarLeapMonth 获取农历闰月（0表示无闰月）
func (c *Calendar) LunarLeapMonth() int64 {
	return c.lunar.LeapMonth()
}

// 生肖天干地支相关方法

// Animal 生肖
func (c *Calendar) Animal() string {
	return c.zodiac.Animal
}

// AnimalWithYear 生肖年，格式：兔年
func (c *Calendar) AnimalWithYear() string {
	return c.zodiac.Animal + "年"
}

// YearGanZhi 年干支，格式：癸卯
func (c *Calendar) YearGanZhi() string {
	return c.zodiac.YearGanZhi
}

// MonthGanZhi 月干支
func (c *Calendar) MonthGanZhi() string {
	return c.zodiac.MonthGanZhi
}

// DayGanZhi 日干支
func (c *Calendar) DayGanZhi() string {
	return c.zodiac.DayGanZhi
}

// HourGanZhi 时干支
func (c *Calendar) HourGanZhi() string {
	return c.zodiac.HourGanZhi
}

// FullGanZhi 完整干支信息，格式：癸卯年 甲申月 己巳日 乙亥时
func (c *Calendar) FullGanZhi() string {
	return fmt.Sprintf("%s年 %s月 %s日 %s时",
		c.zodiac.YearGanZhi, c.zodiac.MonthGanZhi,
		c.zodiac.DayGanZhi, c.zodiac.HourGanZhi)
}

// 节气季节相关方法

// CurrentSolarTerm 当前节气
func (c *Calendar) CurrentSolarTerm() string {
	return c.season.CurrentTerm
}

// NextSolarTerm 下个节气
func (c *Calendar) NextSolarTerm() string {
	return c.season.NextTerm
}

// NextSolarTermTime 下个节气时间
func (c *Calendar) NextSolarTermTime() time.Time {
	return c.season.NextTermTime
}

// DaysToNextTerm 距离下个节气的天数
func (c *Calendar) DaysToNextTerm() int {
	return int(c.season.NextTermTime.Sub(c.Time.Time).Hours() / 24)
}

// Season 当前季节
func (c *Calendar) Season() string {
	return c.season.Season
}

// SeasonProgress 季节进度(0-1)
func (c *Calendar) SeasonProgress() float64 {
	return c.season.SeasonProgress
}

// YearProgress 年度进度(0-1)
func (c *Calendar) YearProgress() float64 {
	return c.season.YearProgress
}

// 格式化输出方法

// String 完整的日历信息字符串
func (c *Calendar) String() string {
	return fmt.Sprintf("%s %s %s %s",
		c.Time.Format("2006年01月02日"),
		c.LunarDateShort(),
		c.AnimalWithYear(),
		c.CurrentSolarTerm())
}

// DetailedString 详细的日历信息
func (c *Calendar) DetailedString() string {
	return fmt.Sprintf(`公历：%s %s
农历：%s
干支：%s
节气：%s（下个：%s，%d天后）
季节：%s（进度：%.1f%%）`,
		c.Time.Format("2006年01月02日 15:04:05"), c.Time.Weekday().String(),
		c.LunarDate(),
		c.FullGanZhi(),
		c.CurrentSolarTerm(), c.NextSolarTerm(), c.DaysToNextTerm(),
		c.Season(), c.SeasonProgress()*100)
}

// ToMap 转换为map格式，便于JSON序列化
func (c *Calendar) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"solar": map[string]interface{}{
			"date":      c.Time.Format("2006-01-02"),
			"time":      c.Time.Format("15:04:05"),
			"weekday":   c.Time.Weekday().String(),
			"timestamp": c.Time.Unix(),
		},
		"lunar": map[string]interface{}{
			"year":       c.lunar.Year(),
			"month":      c.lunar.Month(),
			"day":        c.lunar.Day(),
			"date":       c.LunarDate(),
			"dateShort":  c.LunarDateShort(),
			"isLeapYear": c.IsLunarLeapYear(),
			"leapMonth":  c.LunarLeapMonth(),
		},
		"zodiac": map[string]interface{}{
			"animal":      c.Animal(),
			"yearGanZhi":  c.YearGanZhi(),
			"monthGanZhi": c.MonthGanZhi(),
			"dayGanZhi":   c.DayGanZhi(),
			"hourGanZhi":  c.HourGanZhi(),
			"fullGanZhi":  c.FullGanZhi(),
		},
		"season": map[string]interface{}{
			"current":        c.CurrentSolarTerm(),
			"next":           c.NextSolarTerm(),
			"nextTime":       c.NextSolarTermTime(),
			"daysToNext":     c.DaysToNextTerm(),
			"season":         c.Season(),
			"seasonProgress": c.SeasonProgress(),
			"yearProgress":   c.YearProgress(),
		},
	}
}

// calculateZodiac 计算生肖天干地支信息
func (c *Calendar) calculateZodiac() ZodiacInfo {
	// 获取农历年月日时信息
	lunarYear := c.lunar.Year()

	// 计算生肖（基于农历年）
	animal := c.lunar.Animal()

	// 计算年干支（基于农历年）
	yearTrunk := skyTrunks[(lunarYear-4)%10]
	yearBranch := earthBranches[(lunarYear-4)%12]
	yearGanZhi := yearTrunk + yearBranch

	// 计算月干支（基于公历月，简化计算）
	month := int(c.Time.Month())
	monthTrunk := skyTrunks[(month-1)%10]
	monthBranch := earthBranches[(month-1)%12]
	monthGanZhi := monthTrunk + monthBranch

	// 计算日干支（基于公历日期，简化计算）
	dayOfYear := c.Time.YearDay()
	dayTrunk := skyTrunks[(dayOfYear-1)%10]
	dayBranch := earthBranches[(dayOfYear-1)%12]
	dayGanZhi := dayTrunk + dayBranch

	// 计算时干支
	hour := c.Time.Hour()
	hourIndex := (hour + 1) / 2 % 12
	hourTrunk := skyTrunks[hour%10]
	hourBranch := earthBranches[hourIndex]
	hourGanZhi := hourTrunk + hourBranch

	return ZodiacInfo{
		Animal:      animal,
		SkyTrunk:    yearTrunk,
		EarthBranch: yearBranch,
		YearGanZhi:  yearGanZhi,
		MonthGanZhi: monthGanZhi,
		DayGanZhi:   dayGanZhi,
		HourGanZhi:  hourGanZhi,
	}
}

// calculateSeason 计算节气季节信息
func (c *Calendar) calculateSeason() SeasonInfo {
	// 获取当前和下个节气
	now := c.Time.Time
	nextSolarterm := NextSolarterm(now)
	nextTermTime := nextSolarterm.Time()

	// 获取当前节气（简化：基于月份）
	month := int(now.Month())
	var currentTerm, season string
	var seasonProgress float64

	switch {
	case month >= 2 && month <= 4: // 春季
		season = "春"
		currentTerm = []string{"立春", "雨水", "惊蛰", "春分", "清明", "谷雨"}[(month-2)*2+now.Day()/15]
		seasonProgress = float64(month-2)/3.0 + float64(now.Day())/90.0
	case month >= 5 && month <= 7: // 夏季
		season = "夏"
		currentTerm = []string{"立夏", "小满", "芒种", "夏至", "小暑", "大暑"}[(month-5)*2+now.Day()/15]
		seasonProgress = float64(month-5)/3.0 + float64(now.Day())/90.0
	case month >= 8 && month <= 10: // 秋季
		season = "秋"
		currentTerm = []string{"立秋", "处暑", "白露", "秋分", "寒露", "霜降"}[(month-8)*2+now.Day()/15]
		seasonProgress = float64(month-8)/3.0 + float64(now.Day())/90.0
	default: // 冬季
		season = "冬"
		currentTerm = []string{"立冬", "小雪", "大雪", "冬至", "小寒", "大寒"}[((month+9)%12)*2+now.Day()/15]
		seasonProgress = float64((month+9)%12)/3.0 + float64(now.Day())/90.0
	}

	// 计算年度进度
	startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	endOfYear := time.Date(now.Year()+1, 1, 1, 0, 0, 0, 0, now.Location())
	yearProgress := float64(now.Sub(startOfYear)) / float64(endOfYear.Sub(startOfYear))

	return SeasonInfo{
		CurrentTerm:    currentTerm,
		NextTerm:       nextSolarterm.String(),
		NextTermTime:   nextTermTime,
		Season:         season,
		SeasonProgress: seasonProgress,
		YearProgress:   yearProgress,
	}
}
