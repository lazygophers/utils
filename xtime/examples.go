package xtime

import (
	"fmt"
	"time"
)

// ExampleUsage 展示xtime包的各种使用方法
func ExampleUsage() {
	fmt.Println("=== xtime 农历节气 SDK 使用示例 ===")

	// 1. 基础日历功能
	fmt.Println("1. 创建日历对象")
	cal := NowCalendar()
	fmt.Printf("完整信息: %s\n", cal.String())
	fmt.Printf("详细信息:\n%s\n\n", cal.DetailedString())

	// 2. 农历功能
	fmt.Println("2. 农历功能")
	fmt.Printf("农历日期: %s\n", cal.LunarDate())
	fmt.Printf("农历简写: %s\n", cal.LunarDateShort())
	fmt.Printf("生肖: %s\n", cal.AnimalWithYear())
	fmt.Printf("年干支: %s\n", cal.YearGanZhi())
	fmt.Printf("完整干支: %s\n", cal.FullGanZhi())
	fmt.Printf("是否闰年: %t\n", cal.IsLunarLeapYear())
	if cal.IsLunarLeapYear() {
		fmt.Printf("闰月: %d月\n", cal.LunarLeapMonth())
	}
	fmt.Println()

	// 3. 节气功能
	fmt.Println("3. 节气功能")
	fmt.Printf("当前节气: %s\n", cal.CurrentSolarTerm())
	fmt.Printf("下个节气: %s\n", cal.NextSolarTerm())
	fmt.Printf("距下个节气: %d天\n", cal.DaysToNextTerm())
	fmt.Printf("当前季节: %s\n", cal.Season())
	fmt.Printf("季节进度: %.1f%%\n", cal.SeasonProgress()*100)
	fmt.Printf("年度进度: %.1f%%\n", cal.YearProgress()*100)
	fmt.Println()

	// 4. 农历助手功能
	fmt.Println("4. 农历助手功能")
	lunarHelper := NewLunarHelper()

	// 检查今天是否是节日
	if festival := lunarHelper.GetTodayFestival(); festival != nil {
		fmt.Printf("今天是节日: %s\n", lunarHelper.FormatFestivalInfo(festival))
	} else {
		fmt.Println("今天不是传统节日")
	}

	// 检查是否是特殊日子
	if isSpecial, desc := lunarHelper.IsSpecialDay(time.Now()); isSpecial {
		fmt.Printf("特殊日子: %s\n", desc)
	}

	// 获取完整农历信息
	lunarInfo := lunarHelper.GetLunarInfo(time.Now())
	fmt.Printf("农历信息: %+v\n\n", lunarInfo)

	// 5. 节气助手功能
	fmt.Println("5. 节气助手功能")
	termHelper := NewSolarTermHelper()

	// 当前节气详情
	if currentTerm := termHelper.GetCurrentTerm(time.Now()); currentTerm != nil {
		fmt.Printf("当前节气详情:\n%s\n", termHelper.FormatTermInfo(currentTerm))
	}

	// 下个节气详情
	if nextTerm := termHelper.GetNextTerm(time.Now()); nextTerm != nil {
		fmt.Printf("下个节气详情:\n%s\n", termHelper.FormatTermInfo(nextTerm))
	}

	// 获取今年的节气日历
	calendar := termHelper.GetTermCalendar(time.Now().Year())
	fmt.Printf("今年节气分布: %d个月有节气\n", len(calendar))

	// 6. JSON序列化
	fmt.Println("6. JSON序列化")
	jsonData := cal.ToMap()
	fmt.Printf("JSON数据结构包含 %d 个主要部分\n", len(jsonData))
	fmt.Printf("- solar: 公历信息\n- lunar: 农历信息\n- zodiac: 生肖干支\n- season: 节气季节\n\n")

	// 7. 批量查询示例
	fmt.Println("7. 批量查询示例")

	// 获取最近的节气
	recentTerms := termHelper.GetRecentTerms(time.Now(), 5)
	fmt.Printf("最近的 %d 个节气:\n", len(recentTerms))
	for i, term := range recentTerms {
		status := "未来"
		if term.Time.Before(time.Now()) {
			status = "过去"
		}
		fmt.Printf("  %d. %s (%s) - %s\n", i+1, term.Name,
			term.Time.Format("01-02"), status)
	}

	// 8. 实用工具方法
	fmt.Println("\n8. 实用工具方法")

	// 距离特定节气的天数
	daysToSpring := termHelper.DaysUntilTerm(time.Now(), "立春")
	if daysToSpring >= 0 {
		fmt.Printf("距离立春还有 %d 天\n", daysToSpring)
	}

	// 生日相关（示例：1990年5月20日出生）
	birthTime := time.Date(1990, 5, 20, 0, 0, 0, 0, time.Local)
	lunarAge := lunarHelper.GetLunarAge(birthTime, time.Now())
	fmt.Printf("农历虚岁: %d 岁\n", lunarAge)

	// 比较农历日期
	comparison := lunarHelper.CompareLunarDates(birthTime, time.Now())
	fmt.Printf("生日对比: %s\n", comparison)
}

// QuickExample 快速示例
func QuickExample() {
	// 最简单的使用方法
	cal := NowCalendar()
	fmt.Printf("今天是：%s\n", cal.String())

	// 检查节日
	lunarHelper := NewLunarHelper()
	if festival := lunarHelper.GetTodayFestival(); festival != nil {
		fmt.Printf("今天是%s！\n", festival.Name)
	}

	// 查看节气
	termHelper := NewSolarTermHelper()
	if term := termHelper.GetCurrentTerm(time.Now()); term != nil {
		fmt.Printf("当前节气：%s\n", term.Name)
		fmt.Printf("养生提示：%v\n", term.Tips)
	}
}

// GetTodayLucky 获取今日运势（示例功能）
func GetTodayLucky() map[string]string {
	cal := NowCalendar()
	lunar := cal.Lunar()

	// 基于农历日期计算简单运势（示例算法）
	dayNum := lunar.Day()
	monthNum := lunar.Month()

	lucky := map[string]string{
		"overall": "平稳",
		"love":    "一般",
		"career":  "顺利",
		"wealth":  "小有收获",
		"health":  "注意休息",
	}

	// 简单的运势算法示例
	switch (dayNum + monthNum) % 5 {
	case 0:
		lucky["overall"] = "大吉"
		lucky["love"] = "桃花运旺"
		lucky["career"] = "贵人相助"
	case 1:
		lucky["overall"] = "吉"
		lucky["wealth"] = "财运亨通"
	case 2:
		lucky["overall"] = "平"
	case 3:
		lucky["health"] = "身体健康"
		lucky["career"] = "工作顺心"
	case 4:
		lucky["overall"] = "需谨慎"
		lucky["love"] = "感情稳定"
	}

	return lucky
}

// FormatTodayInfo 格式化今日信息
func FormatTodayInfo() string {
	cal := NowCalendar()
	lunarHelper := NewLunarHelper()
	termHelper := NewSolarTermHelper()

	info := fmt.Sprintf(`📅 今日信息 📅

🌞 公历：%s %s
🌙 农历：%s
🐲 生肖：%s
🌿 节气：%s（下个：%s，%d天后）
🍂 季节：%s（进度：%.1f%%）
📊 年度进度：%.1f%%`,
		cal.Time.Format("2006年01月02日"),
		cal.Time.Weekday(),
		cal.LunarDate(),
		cal.AnimalWithYear(),
		cal.CurrentSolarTerm(),
		cal.NextSolarTerm(),
		cal.DaysToNextTerm(),
		cal.Season(),
		cal.SeasonProgress()*100,
		cal.YearProgress()*100)

	// 添加节日信息
	if festival := lunarHelper.GetTodayFestival(); festival != nil {
		info += fmt.Sprintf("\n🎉 今日节日：%s", festival.Name)
		info += fmt.Sprintf("\n🏮 传统习俗：%v", festival.Traditions)
		info += fmt.Sprintf("\n🍜 传统美食：%v", festival.Foods)
	}

	// 添加特殊日子信息
	if isSpecial, desc := lunarHelper.IsSpecialDay(time.Now()); isSpecial {
		info += fmt.Sprintf("\n✨ 特殊意义：%s", desc)
	}

	// 添加节气养生
	if term := termHelper.GetCurrentTerm(time.Now()); term != nil && len(term.Tips) > 0 {
		info += fmt.Sprintf("\n💡 养生贴士：%v", term.Tips)
	}

	return info
}
