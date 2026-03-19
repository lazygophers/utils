package main

import (
	"fmt"
	"time"

	"github.com/lazygophers/utils/xtime"
)

func main() {
	// 示例1: 获取当前日历信息
	cal := xtime.NowCalendar()

	fmt.Printf("公历日期: %s\n", cal.Format("2006-01-02"))
	fmt.Printf("农历日期: %s\n", cal.LunarDate())
	fmt.Printf("简短农历: %s\n", cal.LunarDateShort())
	fmt.Printf("生肖: %s\n", cal.Animal())
	fmt.Printf("生肖年: %s\n", cal.AnimalWithYear())
	fmt.Printf("干支年: %s\n", cal.YearGanZhi())
	fmt.Printf("干支月: %s\n", cal.MonthGanZhi())
	fmt.Printf("干支日: %s\n", cal.DayGanZhi())

	// 示例2: 节气信息
	fmt.Printf("当前节气: %s\n", cal.CurrentSolarTerm())
	fmt.Printf("下个节气: %s\n", cal.NextSolarTerm())
	fmt.Printf("季节: %s\n", cal.Season())
	fmt.Printf("季节进度: %.2f%%\n", cal.SeasonProgress()*100)

	// 示例3: 检查闰月
	if cal.IsLunarLeapYear() {
		fmt.Printf("今年是闰年，闰%d月\n", cal.LunarLeapMonth())
	}

	// 示例4: 指定日期的农历信息
	date := time.Date(2023, 1, 22, 0, 0, 0, 0, time.Local)
	lunar := xtime.WithLunar(date)
	fmt.Printf("\n2023-01-22 是农历: %s%s%s\n",
		lunar.YearAlias(),
		lunar.MonthAlias(),
		lunar.DayAlias())
	fmt.Printf("生肖: %s\n", lunar.Animal())

	// 示例5: 时间范围操作
	now := xtime.With(time.Now())
	fmt.Printf("\n今天开始: %s\n", now.BeginningOfDay().Format("2006-01-02 15:04:05"))
	fmt.Printf("今天结束: %s\n", now.EndOfDay().Format("2006-01-02 15:04:05"))
	fmt.Printf("本周开始: %s\n", now.BeginningOfWeek().Format("2006-01-02"))
	fmt.Printf("本周结束: %s\n", now.EndOfWeek().Format("2006-01-02"))
	fmt.Printf("本月开始: %s\n", now.BeginningOfMonth().Format("2006-01-02"))
	fmt.Printf("本月结束: %s\n", now.EndOfMonth().Format("2006-01-02"))
}
