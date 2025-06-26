// Lunar 包提供农历日期处理功能，支持农历与公历转换、闰月判断、生肖计算及汉字格式化。
//
// 核心功能包括：
// - 农历/公历日期转换
// - 闰月计算与判断
// - 生肖（十二生肖）推算
// - 传统汉字日期格式化（如 "二零一八年闰六月"）
// - 农历日期结构操作
//
// 主要类型：
// - Lunar：包含农历日期信息和公历时间功能
//
// 示例用法：
//
//	ts := time.Now().Unix()
//	lunar := xtime.WithLunar(time.Now())
//	fmt.Println(lunar.YearAlias(), lunar.MonthAlias())
package xtime

import (
	"fmt"
	"strings"
	"time"
)

// Lunar 结构体表示农历日期信息，包含公历时间字段和农历计算参数
type Lunar struct {
	time.Time

	year, month, day int64
	monthIsLeap      bool
}

// LeapMonth 返回农历年份的闰月信息。返回值为0表示无闰月，1-12表示对应的闰月。
func (p *Lunar) LeapMonth() int64 {
	return leapMonth(p.year)
}

// IsLeap 判断农历年份是否为闰年（存在闰月）。
func (p *Lunar) IsLeap() bool {
	return p.LeapMonth() != 0
}

// IsLeapMonth 检查当前月份是否为闰月。
func (p *Lunar) IsLeapMonth() bool { return p.monthIsLeap }

// Animal 获取农历年份对应的生肖名称（十二生肖）。
// 计算方式：根据"鼠牛虎兔龙蛇马羊猴鸡狗猪"顺序循环计算，公式为 (年份-3) % 12。
func (p *Lunar) Animal() string {
	order := OrderMod(p.year-3, 12)

	if 1 <= order && order <= 12 {
	return animalAlias[(order-1)%12]
}

	return ""
}

// YearAlias 汉字表示年(二零一八)
func (p *Lunar) YearAlias() string {
	s := fmt.Sprintf("%d", p.year)
	for i, replace := range numberAlias {
		s = strings.Replace(s, fmt.Sprintf("%d", i), replace, -1)
	}
	return s
}

// MonthAlias 返回农历月份的汉字表示（如"闰六月"）。
func (p *Lunar) MonthAlias() string {
	pre := ""
	if p.monthIsLeap {
		pre = "闰"
	}
	return pre + lunarMonthAlias[p.month-1] + "月"
}

// DayAlias 汉字表示日(初一, 初十...)
func (p *Lunar) DayAlias() (alias string) {
	switch p.day {
	case 10:
		alias = "初十"
	case 20:
		alias = "二十"
	case 30:
		alias = "三十"
	default:
		alias = dateAlias[(int)(p.day/10)] + numberAlias[p.day%10]
	}
	return
}

// Year 年
func (p *Lunar) Year() int64 {
	return p.year
}

// Month 月
func (p *Lunar) Month() int64 {
	return p.month
}

// Day 日
func (p *Lunar) Day() int64 {
	return p.day
	}

// Day 日
func (p *Lunar) Date() string {
	return fmt.Sprintf("%02d-%02d-%02d", p.year, p.month, p.day)
}

// MonthDayAlise 返回月份和日期的组合格式（如"闰6-15"）。
func (p *Lunar) MonthDayAlise() string {
	if p.monthIsLeap {
		return fmt.Sprintf("闰%d-%d", p.Month(), p.Day())
	}
	return fmt.Sprintf("%d-%d", p.Month(), p.Day())
}

// Equals 返回两个对象是否相同
func (p *Lunar) Equals(b *Lunar) bool {
	return p.Time.Equal(b.Time)
}

func WithLunarTime(t *Time) *Lunar {
	return WithLunar(t.Time)
}

func WithLunar(t time.Time) *Lunar {
	year, month, day, isLeap := FromSolarTimestamp(t.Unix())
	return &Lunar{
		Time:        t,
		year:        year,
		month:       month,
		day:         day,
		monthIsLeap: isLeap,
	}
}

// FromSolarTimestamp 将公历Unix时间戳转换为农历日期信息。
// 参数：
//
//	ts: 公历Unix时间戳（秒）
//
// 返回：
//
//	lunarYear: 农历年份
//	lunarMonth: 农历月份（1-12）
//	lunarDay: 农历日期（1-30）
//	lunarMonthIsLeap: 是否为闰月
//
// 注意：输入时间戳需在1900-01-31到2099-12-31范围内。
func FromSolarTimestamp(ts int64) (lunarYear, lunarMonth, lunarDay int64, lunarMonthIsLeap bool) {
	var (
		i, offset, leap         int64
		daysOfYear, daysOfMonth int64
		isLeap                  bool
	)
	// 与 1900-01-31 相差多少天
	t := time.Unix(ts, 0)
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	t2 := time.Date(1900, 1, 31, 0, 0, 0, 0, time.Local)
	offset = (t1.Unix() - t2.Unix()) / 86400

	for i = 1900; i < 2101 && offset > 0; i++ {
		daysOfYear = daysOfLunarYear(i)
		offset -= daysOfYear
	}
	if offset < 0 {
		offset += daysOfYear
		i--
}

	// 农历年
	lunarYear = i
	// 闰哪个月
	leap = leapMonth(i)

	isLeap = false

	// 使用当年的天数 offset，逐个减去农历每月天数，计算当前日期
	// 包含闰月处理逻辑
	for i = 1; i < 13 && offset > 0; i++ {
		// 闰月
		if leap > 0 && i == (leap+1) && !isLeap {
			i--
			isLeap = true
			// 计算农历月天数
			daysOfMonth = leapDays(lunarYear)
		} else {
			// 计算农历普通月天数
			daysOfMonth = lunarDays(lunarYear, i)
		}
		// 解除闰月
		if true == isLeap && i == (leap+1) {
			isLeap = false
		}
		offset -= daysOfMonth
	}
	// 当 offset 为 0 且月份是闰月时进行校正
	if 0 == offset && leap > 0 && i == leap+1 {
		if isLeap {
			isLeap = false
		} else {
			isLeap = true
			i--
		}
	}
	if offset < 0 {
		offset += daysOfMonth
		i--
	}
	// 农历月
	lunarMonth = i
	// 农历日
	lunarDay = offset + 1
	// 农历是否为闰月
	lunarMonthIsLeap = isLeap

	return
}

func OrderMod(a, b int64) (result int64) {
	result = a % b
	if result == 0 {
		result = b
	}
	return
}

// daysOfLunarYear 计算指定农历年的总天数。
func daysOfLunarYear(year int64) int64 {
	var (
		i, sum int64
	)
	sum = 29 * 12
	for i = 0x8000; i > 0x8; i >>= 1 {
		if (lunars[year-1900] & i) != 0 {
			sum++
		}
	}
	return sum + leapDays(year)
}

// leapMonth 获取指定年份的闰月月份（0表示无闰月）。
func leapMonth(year int64) int64 {
	return lunars[year-1900] & 0xf
}

func leapDays(year int64) (days int64) {
	if leapMonth(year) == 0 {
		days = 0
	} else if (lunars[year-1900] & 0x10000) != 0 {
		days = 30
	} else {
		days = 29
	}
	return
}

func lunarDays(year, month int64) (days int64) {
	if month > 12 || month < 1 {
		days = 0
	} else if (lunars[year-1900] & (0x10000 >> uint64(month))) != 0 {
		days = 30
	} else {
		days = 29
	}
	return
}

// numberAlias 存储数字到汉字的映射（0-9）。
var numberAlias = [...]string{
	"零", "一", "二", "三", "四",
	"五", "六", "七", "八", "九",
}

// dateAlias 存储日期格式化前缀（如"初"、"十"）。
var dateAlias = [...]string{
	"初", "十", "廿", "卅",
}

// lunarMonthAlias 存储农历月份汉字别名（正月至腊月）。
var lunarMonthAlias = [...]string{
	"正", "二", "三", "四", "五", "六",
	"七", "八", "九", "十", "冬", "腊",
}

// lunars 存储农历历法数据（1900-2100年）。
var lunars = [...]int64{
	0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2, // 1900-1909
	0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977, // 1910-1919
	0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970, // 1920-1929
	0x06566, 0x0d4a0, 0x0ea50, 0x06e95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950, // 1930-1939
	0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557, // 1940-1949
	0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5b0, 0x14573, 0x052b0, 0x0a9a8, 0x0e950, 0x06aa0, // 1950-1959
	0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0, // 1960-1969
	0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b6a0, 0x195a6, // 1970-1979
	0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570, // 1980-1989
	0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x055c0, 0x0ab60, 0x096d5, 0x092e0, // 1990-1999
	0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5, // 2000-2009
	0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930, // 2010-2019
	0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530, // 2020-2029
	0x05aa0, 0x076a3, 0x096d0, 0x04afb, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45, // 2030-2039
	0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0, // 2040-2049
	0x14b63, 0x09370, 0x049f8, 0x04970, 0x064b0, 0x168a6, 0x0ea50, 0x06b20, 0x1a6c4, 0x0aae0, // 2050-2059
	0x0a2e0, 0x0d2e3, 0x0c960, 0x0d557, 0x0d4a0, 0x0da50, 0x05d55, 0x056a0, 0x0a6d0, 0x055d4, // 2060-2069
	0x052d0, 0x0a9b8, 0x0a950, 0x0b4a0, 0x0b6a6, 0x0ad50, 0x055a0, 0x0aba4, 0x0a5b0, 0x052b0, // 2070-2079
	0x0b273, 0x06930, 0x07337, 0x06aa0, 0x0ad50, 0x14b55, 0x04b60, 0x0a570, 0x054e4, 0x0d160, // 2080-2089
	0x0e968, 0x0d520, 0x0daa0, 0x16aa6, 0x056d0, 0x04ae0, 0x0a9d4, 0x0a2d0, 0x0d150, 0x0f252, // 2090-2099
	0x0d520, // 2100
}

// 动物生肖常量数组
var animalAlias = [...]string{
	"鼠", "牛", "虎", "兔", "龙", "蛇",
	"马", "羊", "猴", "鸡", "狗", "猪",
}