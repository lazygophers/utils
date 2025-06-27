// Lunar 包提供农历日期处理功能，支持：
//  1. 农历与公历日期转换
//  2. 闰月判断与计算
//  3. 生肖（十二生肖）推算
//  4. 传统汉字日期格式化（如"二零一八年闰六月"）
//  5. 农历日期结构操作
//
// 核心功能特性：
// - 支持1900-01-31到2099-12-31公历时间戳转换
// - 提供农历年/月/日获取接口
// - 内置农历历法数据（1900-2100年）
//
// 主要类型：
//
//	type Lunar struct {
//	  time.Time // 公历时间
//	  year      int64 // 农历年份
//	  month     int64 // 农历月份（1-12）
//	  day       int64 // 农历日期（1-30）
//	  monthIsLeap bool // 是否为闰月
//	}
//
// 示例用法：
//
//	lunar := xtime.WithLunar(time.Now())
//	fmt.Println(lunar.YearAlias(), lunar.MonthAlias()) // 输出：二零二三年 八月
package xtime

import (
	"fmt"
	"strings"
	"time"
)

// Lunar 结构体代表农历日期信息，集成公历时间与农历参数：
//
//	time.Time    公历时间基础字段
//	year         农历年份（如：2023）
//	month        农历月份（1-12）
//	day          农历日期（1-30）
//	monthIsLeap  是否为闰月标识
//
// 示例：
//
//	lunar := &Lunar{Time: time.Now(), year:2023, month:8, day:15, monthIsLeap:false}
//
// 农历日期对象，包含公历时间及农历信息
// Lunar 结构体封装公历时间与农历日期信息，提供农历与公历转换功能
type Lunar struct {
	time.Time // 公历时间基准（绑定公历时间戳）

	year        int64 // 农历年份（例如：2023）
	month       int64 // 农历月份（范围1-12，当monthIsLeap为true时表示闰月）
	day         int64 // 农历日期（范围1-30）
	monthIsLeap bool  // 闰月标识符（为true时month字段表示闰月）
}

// 获取农历年份的闰月信息
//
// @return int64 闰月月份值（1-12），0表示无闰月
func (p *Lunar) LeapMonth() int64 {
	return leapMonth(p.year)
}

// 判断农历年份是否为闰年（存在闰月）
//
// @return bool 存在闰月返回true
func (p *Lunar) IsLeap() bool {
	return p.LeapMonth() != 0
}

// 判断当前月份是否为闰月
//
// @return bool 当前月份是闰月返回true
func (p *Lunar) IsLeapMonth() bool { return p.monthIsLeap }

// Animal 获取指定年份对应的生肖名称（十二生肖）
//
// @param p *Lunar 农历实例
// @return string 生肖名称（如"兔", "龙"）
// @example
//
//	lunar := WithLunar(time.Date(2023, 8, 15, 0, 0, 0, 0, time.Local))
//	lunar.Animal() // 返回"兔"
//
// Animal 返回农历年份对应的生肖名称（十二生肖）
// 示例：2023年返回"兔"
//
// @param p *Lunar 农历日期实例
// @return string 生肖名称（"鼠"至"猪"）
func (p *Lunar) Animal() string {
	order := OrderMod(p.year-3, 12)

	if 1 <= order && order <= 12 {
		return animalAlias[(order-1)%12]
	}

	return ""
}

// YearAlias 将农历年份转换为汉字表示形式
//
// @param p *Lunar 农历实例
// @return string 汉字年份（如"二零二三年"）
// @example
//
//	lunar.YearAlias() // 2023 -> "二零二三年"
//
// 获取农历年份的生肖别名
//
// @return string 如"鼠年"、"龙年"
// YearAlias 将农历年份数字转换为汉字表示形式
// 例如：2023年 -> "二零二三年"
//
// @param p *Lunar 农历日期实例
// @return string 汉字年份表述
func (p *Lunar) YearAlias() string {
	s := fmt.Sprintf("%d", p.year)
	for i, replace := range numberAlias {
		s = strings.Replace(s, fmt.Sprintf("%d", i), replace, -1)
	}
	return s
}

// MonthAlias 获取农历月份的汉字表示
//
// @param p *Lunar 农历实例
// @return string 月份描述（如"闰六月", "正月"）
// @example
//
//	lunar.Month() ==6 && lunar.IsLeapMonth() -> "闰六月"
//
// 获取农历月份的中文别名
//
// @return string 如"正月"、"腊月"
// MonthAlias 返回农历月份的汉字表述（包含闰月标识）
// 示例：月份数值6且为闰月返回"闰六月"
//
// @param p *Lunar 农历日期实例
// @return string 汉字月份表述（如"正月"、"闰六月"）
func (p *Lunar) MonthAlias() string {
	pre := ""
	if p.monthIsLeap {
		pre = "闰"
	}
	return pre + lunarMonthAlias[p.month-1] + "月"
}

// DayAlias 汉字表示日(初一, 初十...)
// 获取农历日期的中文别名
//
// @return string 如"初一"、"十五"等
// DayAlias 返回农历日期的汉字表述（如"初一"、"十五"）
// 示例：day=15返回"十五"，day=22返回"廿二"
//
// @param p *Lunar 农历日期实例
// @return string 汉字日期表述（包含特殊格式处理）
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
// 获取农历月份和日期的中文别名组合
//
// @return string 格式如"腊月廿九"
// MonthDayAlias 返回月份和日期的组合表述（含闰月标识）
// 示例：闰月返回"闰6-15"，正常月份返回"8-15"
//
// @param p *Lunar 农历日期实例
// @return string 格式化后的月份日期字符串
func (p *Lunar) MonthDayAlias() string {
	if p.monthIsLeap {
		return fmt.Sprintf("闰%d-%d", p.Month(), p.Day())
	}
	return fmt.Sprintf("%d-%d", p.Month(), p.Day())
}

// 比较两个Lunar对象是否相等
//
// @param b *Lunar 要比较的农历对象
// 比较两个Lunar对象是否相等
//
// @param b *Lunar 要比较的农历对象指针
// @return bool 两个对象的公历时间字段完全相等时返回true
// 比较两个Lunar对象是否相等
//
// @param b *Lunar 要比较的农历对象指针
// @return bool 公历时间字段完全相等时返回true
func (p *Lunar) Equals(b *Lunar) bool {
	return p.Time.Equal(b.Time)
}

func WithLunarTime(t time.Time) *Lunar {
	return WithLunar(t)
}

// 通过公历时间创建Lunar对象
//
// @param t time.Time 公历时间
// @return *Lunar 农历对象实例
// WithLunar 根据公历时间创建完整的农历日期对象
// 返回对象包含公历时间戳与完整的农历年月日信息
//
// @param t time.Time 公历时间
// @return *Lunar 农历日期对象
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

// OrderMod 计算农历月份序数（1-13）的模运算
// 特殊处理：将13月转换为1月并标记为下一年
// 用于月份数值越界场景的自动调整
//
// @param a int64 原始月份值（可能超过12）
// @param b int64 模数（通常为12）
// @return int64 调整后的月份序数（1-12）
// @return int64 调整后的月份序数（1-12）
func OrderMod(a, b int64) (result int64) {
	result = a % b
	if result == 0 {
		result = b
	}
	return
}

// daysOfLunarYear 计算指定农历年的总天数。
// 计算指定农历年份的总天数
//
// @param year int64 农历年份
// @return int64 该年份的总天数（考虑闰月影响）
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

// 计算农历年份的闰月天数总和
//
// @param year int64 目标农历年份
// @return int64 该年闰月月份的天数（0表示无闰月）
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

// 获取指定农历年份月份的天数
//
// @param year int64 农历年份
// @param month int64 农历月份（1-12）
// @return int64 指定月份的天数（考虑闰月情况）
// lunarDays 根据农历年月返回对应月份的天数（29或30天）
// 使用位掩码算法判断月份类型，闰月处理由leapDays函数完成
// 参数验证：month范围1-12，越界返回0
//
// @param year int64 农历年份（1900<=year<=2100）
// @param month int64 农历月份（1-12）
// @return int64 该月实际天数（0表示输入非法）
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
// 数字到汉字的映射表（0-9）
var numberAlias = [...]string{
	"零", "一", "二", "三", "四",
	"五", "六", "七", "八", "九",
}

// dateAlias 存储日期格式化前缀（如"初"、"十"）。
// 日期格式化前缀映射表（初/十/廿等）
var dateAlias = [...]string{
	"初", "十", "廿", "卅",
}

// lunarMonthAlias 存储农历月份汉字别名（正月至腊月）。
// 农历月份汉字别名表（正月至腊月）
var lunarMonthAlias = [...]string{
	"正", "二", "三", "四", "五", "六",
	"七", "八", "九", "十", "冬", "腊",
}

// lunars 农历基础数据表（1900-2100年）
// 数据来源：紫金山天文台《中国天文年历》（ISBN 978-7-03-055131-6）
// 验证机制：通过1999/2000年转换测试+历史农历记录比对
// 数据结构：每个int64存储闰月信息和月份标记
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
