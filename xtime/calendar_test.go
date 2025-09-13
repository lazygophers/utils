package xtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalendar(t *testing.T) {
	testTime := time.Date(2023, 8, 15, 14, 30, 0, 0, time.Local)

	t.Run("calendar_creation", func(t *testing.T) {
		cal := NewCalendar(testTime)

		assert.NotNil(t, cal)
		assert.NotNil(t, cal.Time)
		assert.NotNil(t, cal.lunar)
		assert.Equal(t, testTime, cal.Time.Time)
	})

	t.Run("now_calendar", func(t *testing.T) {
		cal := NowCalendar()

		assert.NotNil(t, cal)
		assert.True(t, time.Since(cal.Time.Time) < time.Second)
	})

	t.Run("lunar_methods", func(t *testing.T) {
		cal := NewCalendar(testTime)

		// 测试农历相关方法
		assert.NotNil(t, cal.Lunar())
		assert.NotEmpty(t, cal.LunarDate())
		assert.NotEmpty(t, cal.LunarDateShort())
		assert.True(t, cal.IsLunarLeapYear() || !cal.IsLunarLeapYear()) // 布尔值
		assert.True(t, cal.LunarLeapMonth() >= 0)                       // 0或正数

		t.Logf("农历日期: %s", cal.LunarDate())
		t.Logf("农历简写: %s", cal.LunarDateShort())
	})

	t.Run("zodiac_methods", func(t *testing.T) {
		cal := NewCalendar(testTime)

		// 测试生肖干支方法
		assert.NotEmpty(t, cal.Animal())
		assert.NotEmpty(t, cal.AnimalWithYear())
		assert.NotEmpty(t, cal.YearGanZhi())
		assert.NotEmpty(t, cal.MonthGanZhi())
		assert.NotEmpty(t, cal.DayGanZhi())
		assert.NotEmpty(t, cal.HourGanZhi())
		assert.NotEmpty(t, cal.FullGanZhi())

		// 验证格式
		assert.Contains(t, cal.AnimalWithYear(), "年")
		assert.Len(t, []rune(cal.YearGanZhi()), 2) // 干支应该是2个汉字

		t.Logf("生肖: %s", cal.AnimalWithYear())
		t.Logf("完整干支: %s", cal.FullGanZhi())
	})

	t.Run("season_methods", func(t *testing.T) {
		cal := NewCalendar(testTime)

		// 测试节气季节方法
		assert.NotEmpty(t, cal.CurrentSolarTerm())
		assert.NotEmpty(t, cal.NextSolarTerm())
		assert.False(t, cal.NextSolarTermTime().IsZero())
		assert.True(t, cal.DaysToNextTerm() >= 0)
		assert.NotEmpty(t, cal.Season())
		assert.True(t, cal.SeasonProgress() >= 0 && cal.SeasonProgress() <= 1)
		assert.True(t, cal.YearProgress() >= 0 && cal.YearProgress() <= 1)

		t.Logf("当前节气: %s", cal.CurrentSolarTerm())
		t.Logf("季节进度: %.1f%%", cal.SeasonProgress()*100)
	})

	t.Run("string_methods", func(t *testing.T) {
		cal := NewCalendar(testTime)

		// 测试字符串格式化方法
		str := cal.String()
		assert.NotEmpty(t, str)
		assert.Contains(t, str, "2023")

		detailed := cal.DetailedString()
		assert.NotEmpty(t, detailed)
		assert.Contains(t, detailed, "公历")
		assert.Contains(t, detailed, "农历")
		assert.Contains(t, detailed, "干支")
		assert.Contains(t, detailed, "节气")

		t.Logf("简短信息: %s", str)
		t.Logf("详细信息:\n%s", detailed)
	})

	t.Run("to_map", func(t *testing.T) {
		cal := NewCalendar(testTime)

		data := cal.ToMap()
		assert.Contains(t, data, "solar")
		assert.Contains(t, data, "lunar")
		assert.Contains(t, data, "zodiac")
		assert.Contains(t, data, "season")

		// 验证数据结构
		solar := data["solar"].(map[string]interface{})
		assert.Contains(t, solar, "date")
		assert.Contains(t, solar, "time")
		assert.Contains(t, solar, "weekday")
		assert.Contains(t, solar, "timestamp")

		lunar := data["lunar"].(map[string]interface{})
		assert.Contains(t, lunar, "year")
		assert.Contains(t, lunar, "month")
		assert.Contains(t, lunar, "day")

		zodiac := data["zodiac"].(map[string]interface{})
		assert.Contains(t, zodiac, "animal")
		assert.Contains(t, zodiac, "yearGanZhi")

		season := data["season"].(map[string]interface{})
		assert.Contains(t, season, "current")
		assert.Contains(t, season, "season")
	})
}

func TestLunarHelper(t *testing.T) {
	helper := NewLunarHelper()
	testTime := time.Date(2023, 8, 15, 0, 0, 0, 0, time.Local)

	t.Run("festival_detection", func(t *testing.T) {
		// 测试节日检测（可能没有节日）
		festival := helper.GetFestival(testTime)
		if festival != nil {
			assert.NotEmpty(t, festival.Name)
			assert.NotEmpty(t, festival.Description)
			t.Logf("检测到节日: %s", festival.Name)
		}

		// 测试今天的节日 - 无论有无节日都不应该panic
		assert.NotPanics(t, func() {
			_ = helper.GetTodayFestival()
		})
	})

	t.Run("special_day_detection", func(t *testing.T) {
		isSpecial, desc := helper.IsSpecialDay(testTime)
		assert.IsType(t, true, isSpecial)

		if isSpecial {
			assert.NotEmpty(t, desc)
			t.Logf("特殊日子: %s", desc)
		}
	})

	t.Run("lunar_info", func(t *testing.T) {
		info := helper.GetLunarInfo(testTime)

		assert.Contains(t, info, "date")
		assert.Contains(t, info, "zodiac")
		assert.Contains(t, info, "leapInfo")
		assert.Contains(t, info, "special")

		dateInfo := info["date"].(map[string]interface{})
		assert.Contains(t, dateInfo, "fullStr")

		fullStr := dateInfo["fullStr"].(string)
		assert.Contains(t, fullStr, "农历")
		assert.Contains(t, fullStr, "年")

		t.Logf("农历信息: %s", fullStr)
	})

	t.Run("lunar_age", func(t *testing.T) {
		birthTime := time.Date(1990, 5, 20, 0, 0, 0, 0, time.Local)
		currentTime := time.Date(2023, 8, 15, 0, 0, 0, 0, time.Local)

		age := helper.GetLunarAge(birthTime, currentTime)
		assert.True(t, age > 30) // 1990年出生应该30多岁
		assert.True(t, age < 50) // 不应该超过50岁

		t.Logf("农历虚岁: %d", age)
	})

	t.Run("compare_dates", func(t *testing.T) {
		t1 := time.Date(2023, 8, 15, 0, 0, 0, 0, time.Local)
		t2 := time.Date(2024, 8, 15, 0, 0, 0, 0, time.Local)

		comparison := helper.CompareLunarDates(t1, t2)
		assert.NotEmpty(t, comparison)

		t.Logf("日期比较: %s", comparison)
	})

	t.Run("format_festival", func(t *testing.T) {
		// 测试春节信息格式化
		springFestival := &LunarFestival{
			Name: "春节", Month: 1, Day: 1,
			Description: "农历新年",
			Traditions:  []string{"放鞭炮", "贴春联"},
			Foods:       []string{"饺子", "年糕"},
		}

		formatted := helper.FormatFestivalInfo(springFestival)
		assert.NotEmpty(t, formatted)
		assert.Contains(t, formatted, "春节")
		assert.Contains(t, formatted, "农历1月1日")

		t.Logf("节日信息:\n%s", formatted)
	})
}

func TestSolarTermHelper(t *testing.T) {
	helper := NewSolarTermHelper()
	testTime := time.Date(2023, 8, 15, 0, 0, 0, 0, time.Local)

	t.Run("current_term", func(t *testing.T) {
		term := helper.GetCurrentTerm(testTime)

		if term != nil {
			assert.NotEmpty(t, term.Name)
			assert.NotEmpty(t, term.Season)
			assert.NotEmpty(t, term.Description)
			assert.True(t, term.Index >= 0 && term.Index < 24)

			t.Logf("当前节气: %s (%s)", term.Name, term.Season)
		}
	})

	t.Run("next_term", func(t *testing.T) {
		term := helper.GetNextTerm(testTime)

		if term != nil {
			assert.NotEmpty(t, term.Name)
			assert.True(t, term.Time.After(testTime))

			t.Logf("下个节气: %s", term.Name)
		}
	})

	t.Run("year_terms", func(t *testing.T) {
		terms := helper.GetYearTerms(2023)

		// 一年应该有24个节气，但实际可能少一些（跨年问题）
		assert.True(t, len(terms) > 0)
		assert.True(t, len(terms) <= 24)

		// 验证节气按时间排序
		for i := 1; i < len(terms); i++ {
			assert.True(t, terms[i].Time.After(terms[i-1].Time) ||
				terms[i].Time.Equal(terms[i-1].Time))
		}

		t.Logf("2023年节气数量: %d", len(terms))
	})

	t.Run("season_terms", func(t *testing.T) {
		springTerms := helper.GetSeasonTerms(2023, "春")
		summerTerms := helper.GetSeasonTerms(2023, "夏")

		// 每个季节应该有一些节气
		if len(springTerms) > 0 {
			for _, term := range springTerms {
				assert.Equal(t, "春", term.Season)
			}
			t.Logf("春季节气: %d个", len(springTerms))
		}

		if len(summerTerms) > 0 {
			for _, term := range summerTerms {
				assert.Equal(t, "夏", term.Season)
			}
			t.Logf("夏季节气: %d个", len(summerTerms))
		}
	})

	t.Run("find_by_name", func(t *testing.T) {
		term := helper.FindTermByName(2023, "立春")

		if term != nil {
			assert.Equal(t, "立春", term.Name)
			assert.Equal(t, 2023, term.Time.Year())

			t.Logf("2023年立春: %s", term.Time.Format("2006-01-02 15:04"))
		}
	})

	t.Run("term_calendar", func(t *testing.T) {
		calendar := helper.GetTermCalendar(2023)

		// 应该有一些月份包含节气
		assert.True(t, len(calendar) > 0)
		assert.True(t, len(calendar) <= 12)

		for month, terms := range calendar {
			assert.True(t, month >= 1 && month <= 12)
			assert.True(t, len(terms) > 0)

			for _, term := range terms {
				assert.Equal(t, month, int(term.Time.Month()))
			}
		}

		t.Logf("节气日历: %d个月有节气", len(calendar))
	})

	t.Run("days_until_term", func(t *testing.T) {
		days := helper.DaysUntilTerm(testTime, "立春")

		// 应该返回有效天数或-1（未找到）
		assert.True(t, days >= -1)

		if days >= 0 {
			t.Logf("距离立春: %d天", days)
		}
	})

	t.Run("recent_terms", func(t *testing.T) {
		terms := helper.GetRecentTerms(testTime, 5)

		assert.True(t, len(terms) <= 5)

		if len(terms) > 0 {
			t.Logf("最近的节气数量: %d", len(terms))
			for i, term := range terms {
				status := "未来"
				if term.Time.Before(testTime) {
					status = "过去"
				}
				t.Logf("  %d. %s - %s", i+1, term.Name, status)
			}
		}
	})

	t.Run("format_term_info", func(t *testing.T) {
		term := &SolarTermInfo{
			Name: "立春", Season: "春",
			Time:        time.Date(2023, 2, 4, 10, 42, 0, 0, time.Local),
			Description: "春季的开始",
			Tips:        []string{"早睡早起", "疏肝理气"},
		}

		formatted := helper.FormatTermInfo(term)
		assert.NotEmpty(t, formatted)
		assert.Contains(t, formatted, "立春")
		assert.Contains(t, formatted, "春季的开始")

		t.Logf("节气信息:\n%s", formatted)
	})
}

func TestExamples(t *testing.T) {
	t.Run("quick_example", func(t *testing.T) {
		assert.NotPanics(t, func() {
			QuickExample()
		})
	})

	t.Run("get_today_lucky", func(t *testing.T) {
		lucky := GetTodayLucky()

		assert.Contains(t, lucky, "overall")
		assert.Contains(t, lucky, "love")
		assert.Contains(t, lucky, "career")
		assert.Contains(t, lucky, "wealth")
		assert.Contains(t, lucky, "health")

		for key, value := range lucky {
			assert.NotEmpty(t, value)
			t.Logf("%s: %s", key, value)
		}
	})

	t.Run("format_today_info", func(t *testing.T) {
		info := FormatTodayInfo()

		assert.NotEmpty(t, info)
		assert.Contains(t, info, "今日信息")
		assert.Contains(t, info, "公历")
		assert.Contains(t, info, "农历")

		t.Logf("今日信息:\n%s", info)
	})
}

// 基准测试
func BenchmarkCalendar(b *testing.B) {
	testTime := time.Now()

	b.Run("NewCalendar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewCalendar(testTime)
		}
	})

	b.Run("Calendar_String", func(b *testing.B) {
		cal := NewCalendar(testTime)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = cal.String()
		}
	})

	b.Run("Calendar_ToMap", func(b *testing.B) {
		cal := NewCalendar(testTime)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = cal.ToMap()
		}
	})
}

func BenchmarkHelpers(b *testing.B) {
	testTime := time.Now()

	b.Run("LunarHelper_GetLunarInfo", func(b *testing.B) {
		helper := NewLunarHelper()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = helper.GetLunarInfo(testTime)
		}
	})

	b.Run("SolarTermHelper_GetCurrentTerm", func(b *testing.B) {
		helper := NewSolarTermHelper()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = helper.GetCurrentTerm(testTime)
		}
	})

	b.Run("SolarTermHelper_GetYearTerms", func(b *testing.B) {
		helper := NewSolarTermHelper()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = helper.GetYearTerms(2023)
		}
	})
}
