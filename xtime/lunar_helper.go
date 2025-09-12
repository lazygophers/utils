package xtime

import (
	"fmt"
	"time"
)

// LunarHelper 农历查询助手
type LunarHelper struct{}

// NewLunarHelper 创建农历助手
func NewLunarHelper() *LunarHelper {
	return &LunarHelper{}
}

// LunarFestival 农历节日信息
type LunarFestival struct {
	Name        string `json:"name"`        // 节日名称
	Month       int    `json:"month"`       // 农历月份
	Day         int    `json:"day"`         // 农历日期
	Description string `json:"description"` // 节日描述
	Traditions  []string `json:"traditions"` // 传统习俗
	Foods       []string `json:"foods"`      // 传统食物
}

// 传统农历节日定义
var lunarFestivals = []LunarFestival{
	{
		Name: "春节", Month: 1, Day: 1,
		Description: "农历新年，中华民族最重要的传统节日",
		Traditions:  []string{"放鞭炮", "贴春联", "拜年", "给压岁钱", "守岁"},
		Foods:       []string{"饺子", "年糕", "鱼", "汤圆"},
	},
	{
		Name: "元宵节", Month: 1, Day: 15,
		Description: "正月十五，春节的最后一天",
		Traditions:  []string{"赏花灯", "猜灯谜", "舞龙舞狮", "踩高跷"},
		Foods:       []string{"汤圆", "元宵"},
	},
	{
		Name: "龙抬头", Month: 2, Day: 2,
		Description: "二月二龙抬头，祈求风调雨顺",
		Traditions:  []string{"理发", "吃龙须面", "撒灰引龙"},
		Foods:       []string{"龙须面", "春饼", "猪头肉"},
	},
	{
		Name: "上巳节", Month: 3, Day: 3,
		Description: "三月三上巳节，古代踏青节日",
		Traditions:  []string{"踏青", "祓禊", "曲水流觞"},
		Foods:       []string{"青团", "艾草粑粑"},
	},
	{
		Name: "端午节", Month: 5, Day: 5,
		Description: "纪念屈原，驱邪避疫的节日",
		Traditions:  []string{"赛龙舟", "挂艾草", "佩香囊", "系五彩绳"},
		Foods:       []string{"粽子", "雄黄酒", "五毒饼"},
	},
	{
		Name: "七夕节", Month: 7, Day: 7,
		Description: "牛郎织女相会，中国的情人节",
		Traditions:  []string{"乞巧", "穿针引线", "拜织女", "吃巧果"},
		Foods:       []string{"巧果", "巧芽面", "花瓜"},
	},
	{
		Name: "中元节", Month: 7, Day: 15,
		Description: "祭祀祖先，普渡众生的节日",
		Traditions:  []string{"祭祖", "放河灯", "烧纸钱", "做法事"},
		Foods:       []string{"鸭子", "包子", "花馍"},
	},
	{
		Name: "中秋节", Month: 8, Day: 15,
		Description: "团圆节，赏月祈福的节日",
		Traditions:  []string{"赏月", "吃月饼", "家庭团聚", "猜灯谜"},
		Foods:       []string{"月饼", "桂花酒", "柚子", "田螺"},
	},
	{
		Name: "重阳节", Month: 9, Day: 9,
		Description: "九九重阳，登高祈福的节日",
		Traditions:  []string{"登高", "赏菊", "佩茱萸", "敬老"},
		Foods:       []string{"重阳糕", "菊花酒", "螃蟹"},
	},
	{
		Name: "下元节", Month: 10, Day: 15,
		Description: "水官解厄，祈求平安的节日",
		Traditions:  []string{"祭祖", "祈福", "吃麻腐包子"},
		Foods:       []string{"麻腐包子", "红豆糯米饭"},
	},
	{
		Name: "腊八节", Month: 12, Day: 8,
		Description: "佛祖成道日，喝腊八粥的节日",
		Traditions:  []string{"喝腊八粥", "泡腊八蒜", "做腊八豆腐"},
		Foods:       []string{"腊八粥", "腊八蒜", "腊八豆腐"},
	},
	{
		Name: "小年", Month: 12, Day: 23,
		Description: "祭灶王爷，开始准备过年",
		Traditions:  []string{"祭灶王", "扫房子", "剪窗花", "写春联"},
		Foods:       []string{"灶糖", "火烧", "年糕"},
	},
	{
		Name: "除夕", Month: 12, Day: 30,
		Description: "农历年的最后一天，合家团圆",
		Traditions:  []string{"年夜饭", "守岁", "贴春联", "放鞭炮", "给压岁钱"},
		Foods:       []string{"饺子", "鱼", "年糕", "汤圆"},
	},
}

// GetTodayFestival 获取今天的农历节日（如果有）
func (h *LunarHelper) GetTodayFestival() *LunarFestival {
	return h.GetFestival(time.Now())
}

// GetFestival 获取指定日期的农历节日
func (h *LunarHelper) GetFestival(t time.Time) *LunarFestival {
	lunar := WithLunar(t)
	month := int(lunar.Month())
	day := int(lunar.Day())
	
	for _, festival := range lunarFestivals {
		if festival.Month == month && festival.Day == day {
			return &festival
		}
	}
	
	// 检查除夕的特殊情况（腊月最后一天）
	if month == 12 {
		// 获取这个月的天数
		nextMonth := t.AddDate(0, 1, 0)
		nextLunar := WithLunar(nextMonth)
		if nextLunar.Month() == 1 { // 下个月是正月，说明当前是腊月最后一天
			for _, festival := range lunarFestivals {
				if festival.Name == "除夕" {
					return &festival
				}
			}
		}
	}
	
	return nil
}

// GetYearFestivals 获取指定年份的所有农历节日
func (h *LunarHelper) GetYearFestivals(year int) []FestivalDate {
	var festivals []FestivalDate
	
	for _, festival := range lunarFestivals {
		// 转换农历日期为公历日期
		lunarDate := time.Date(year, time.Month(festival.Month), festival.Day, 0, 0, 0, 0, time.Local)
		
		// 这里需要农历转公历的功能，暂时简化处理
		festivals = append(festivals, FestivalDate{
			Festival:  festival,
			SolarDate: lunarDate, // 实际应该转换为公历日期
		})
	}
	
	return festivals
}

// FestivalDate 节日日期信息
type FestivalDate struct {
	Festival  LunarFestival `json:"festival"`
	SolarDate time.Time     `json:"solarDate"`
}

// GetUpcomingFestivals 获取即将到来的农历节日
func (h *LunarHelper) GetUpcomingFestivals(count int) []FestivalDate {
	now := time.Now()
	var upcoming []FestivalDate
	
	// 检查今年剩余的节日
	year := now.Year()
	for i := 0; i < 2; i++ { // 检查今年和明年
		yearFestivals := h.GetYearFestivals(year + i)
		for _, fd := range yearFestivals {
			if fd.SolarDate.After(now) {
				upcoming = append(upcoming, fd)
			}
		}
	}
	
	// 按日期排序并返回指定数量
	if len(upcoming) > count {
		upcoming = upcoming[:count]
	}
	
	return upcoming
}

// FormatFestivalInfo 格式化节日信息
func (h *LunarHelper) FormatFestivalInfo(festival *LunarFestival) string {
	if festival == nil {
		return ""
	}
	
	return fmt.Sprintf(`【%s】农历%d月%d日
描述：%s
习俗：%v
美食：%v`,
		festival.Name, festival.Month, festival.Day,
		festival.Description,
		festival.Traditions,
		festival.Foods)
}

// IsSpecialDay 判断是否是特殊的农历日子
func (h *LunarHelper) IsSpecialDay(t time.Time) (bool, string) {
	lunar := WithLunar(t)
	month := lunar.Month()
	day := lunar.Day()
	
	// 检查是否是节日
	if festival := h.GetFestival(t); festival != nil {
		return true, fmt.Sprintf("今天是%s", festival.Name)
	}
	
	// 检查特殊数字日子
	if month == day {
		return true, fmt.Sprintf("农历%s月%s，日月同数", lunar.MonthAlias(), lunar.DayAlias())
	}
	
	// 检查初一、十五
	if day == 1 {
		return true, fmt.Sprintf("农历%s朔日（初一）", lunar.MonthAlias())
	}
	if day == 15 {
		return true, fmt.Sprintf("农历%s望日（十五）", lunar.MonthAlias())
	}
	
	return false, ""
}

// GetLunarInfo 获取完整的农历信息
func (h *LunarHelper) GetLunarInfo(t time.Time) map[string]interface{} {
	lunar := WithLunar(t)
	festival := h.GetFestival(t)
	isSpecial, specialDesc := h.IsSpecialDay(t)
	
	info := map[string]interface{}{
		"date": map[string]interface{}{
			"year":      lunar.Year(),
			"month":     lunar.Month(),
			"day":       lunar.Day(),
			"yearStr":   lunar.YearAlias(),
			"monthStr":  lunar.MonthAlias(),
			"dayStr":    lunar.DayAlias(),
			"fullStr":   fmt.Sprintf("农历%s年%s%s", lunar.YearAlias(), lunar.MonthAlias(), lunar.DayAlias()),
		},
		"zodiac": map[string]interface{}{
			"animal":     lunar.Animal(),
			"animalYear": lunar.Animal() + "年",
		},
		"leapInfo": map[string]interface{}{
			"isLeapYear":   lunar.IsLeap(),
			"isLeapMonth":  lunar.IsLeapMonth(),
			"leapMonth":    lunar.LeapMonth(),
		},
		"special": map[string]interface{}{
			"isSpecial":   isSpecial,
			"description": specialDesc,
		},
	}
	
	if festival != nil {
		info["festival"] = festival
	}
	
	return info
}

// CompareLunarDates 比较两个农历日期
func (h *LunarHelper) CompareLunarDates(t1, t2 time.Time) string {
	lunar1 := WithLunar(t1)
	lunar2 := WithLunar(t2)
	
	if lunar1.Year() == lunar2.Year() && lunar1.Month() == lunar2.Month() && lunar1.Day() == lunar2.Day() {
		return "同一个农历日期"
	}
	
	if lunar1.Month() == lunar2.Month() && lunar1.Day() == lunar2.Day() {
		return fmt.Sprintf("同一个农历日子：%s%s", lunar1.MonthAlias(), lunar1.DayAlias())
	}
	
	return "不同的农历日期"
}

// GetLunarAge 计算农历虚岁年龄
func (h *LunarHelper) GetLunarAge(birthTime, currentTime time.Time) int {
	birth := WithLunar(birthTime)
	current := WithLunar(currentTime)
	
	age := int(current.Year() - birth.Year())
	
	// 如果还没过农历生日，年龄减1
	if current.Month() < birth.Month() || 
	   (current.Month() == birth.Month() && current.Day() < birth.Day()) {
		age--
	}
	
	return age + 1 // 虚岁比周岁多1
}

// GetNextLunarBirthday 获取下一个农历生日
func (h *LunarHelper) GetNextLunarBirthday(birthTime, currentTime time.Time) time.Time {
	birth := WithLunar(birthTime)
	current := WithLunar(currentTime)
	
	// 先尝试今年的农历生日
	thisYear := current.Year()
	nextYear := thisYear
	
	// 如果今年的生日已过，使用明年的
	if current.Month() > birth.Month() || 
	   (current.Month() == birth.Month() && current.Day() >= birth.Day()) {
		nextYear = thisYear + 1
	}
	
	// 这里需要农历转公历的功能来计算确切的公历日期
	// 暂时返回一个估算的时间
	return time.Date(int(nextYear), time.Month(birth.Month()), int(birth.Day()), 0, 0, 0, 0, time.Local)
}