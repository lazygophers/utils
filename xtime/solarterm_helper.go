package xtime

import (
	"fmt"
	"sort"
	"time"
)

// SolarTermHelper 节气查询助手
type SolarTermHelper struct{}

// NewSolarTermHelper 创建节气助手
func NewSolarTermHelper() *SolarTermHelper {
	return &SolarTermHelper{}
}

// SolarTermInfo 节气详细信息
type SolarTermInfo struct {
	Name        string    `json:"name"`        // 节气名称
	Time        time.Time `json:"time"`        // 节气时间
	Index       int       `json:"index"`       // 节气序号(0-23)
	Season      string    `json:"season"`      // 所属季节
	Description string    `json:"description"` // 节气描述
	Tips        []string  `json:"tips"`        // 养生贴士
}

// 节气描述和养生贴士
var solarTermDescriptions = map[string]struct {
	season      string
	description string
	tips        []string
}{
	"小寒": {"冬", "一年中最寒冷时节的开始", []string{"注意保暖", "适当进补", "早睡晚起"}},
	"大寒": {"冬", "一年中最寒冷的时候", []string{"防寒保暖", "温补肾阳", "避免熬夜"}},
	"立春": {"春", "春季的开始，万物复苏", []string{"早睡早起", "春捂秋冻", "疏肝理气"}},
	"雨水": {"春", "降雨量增多，气温回升", []string{"防春困", "调理脾胃", "适度运动"}},
	"惊蛰": {"春", "春雷始鸣，蛰虫惊醒", []string{"清淡饮食", "调节情绪", "预防感冒"}},
	"春分": {"春", "昼夜等长，寒暑平衡", []string{"平衡饮食", "调节作息", "踏青运动"}},
	"清明": {"春", "天气清朗，万物显明", []string{"踏青赏花", "清淡饮食", "情绪调节"}},
	"谷雨": {"春", "雨量充足，有利谷物生长", []string{"祛湿健脾", "适度运动", "调理肝脏"}},
	"立夏": {"夏", "夏季的开始，气温升高", []string{"清心降火", "适当午休", "多吃清淡"}},
	"小满": {"夏", "夏作物籽粒开始饱满", []string{"防暑降温", "祛湿健脾", "心情舒畅"}},
	"芒种": {"夏", "有芒作物成熟，可以收割", []string{"清热解毒", "调理心情", "适量运动"}},
	"夏至": {"夏", "白昼最长，黑夜最短", []string{"午休补眠", "清热降火", "多吃苦味"}},
	"小暑": {"夏", "气候炎热，但还不是最热", []string{"防暑降温", "清心安神", "饮食清淡"}},
	"大暑": {"夏", "一年中最热的时候", []string{"避免中暑", "多饮水", "心静自然凉"}},
	"立秋": {"秋", "秋季的开始，暑去凉来", []string{"润肺养阴", "早睡早起", "贴秋膘"}},
	"处暑": {"秋", "暑气至此结束", []string{"滋阴润燥", "调理脾胃", "适度运动"}},
	"白露": {"秋", "天气转凉，露水出现", []string{"预防感冒", "养肺润燥", "适当秋冻"}},
	"秋分": {"秋", "昼夜等长，凉爽宜人", []string{"平补肺气", "调节情绪", "登高远望"}},
	"寒露": {"秋", "气温更低，露水更凉", []string{"保暖防寒", "滋阴润肺", "适量进补"}},
	"霜降": {"秋", "开始出现霜冻", []string{"温补脾肾", "预防感冒", "调理肠胃"}},
	"立冬": {"冬", "冬季的开始，准备过冬", []string{"适度进补", "早睡晚起", "保暖防寒"}},
	"小雪": {"冬", "开始下雪，但雪量不大", []string{"温补肾阳", "防寒保暖", "调节情绪"}},
	"大雪": {"冬", "降雪量增多，天气更冷", []string{"大补特补", "早睡晚起", "避免外感"}},
	"冬至": {"冬", "白昼最短，黑夜最长", []string{"温补阳气", "适度进补", "早睡晚起"}},
}

// GetCurrentTerm 获取当前节气信息
func (h *SolarTermHelper) GetCurrentTerm(t time.Time) *SolarTermInfo {
	terms := h.GetYearTerms(t.Year())

	var current *SolarTermInfo
	for i, term := range terms {
		if t.After(term.Time) {
			current = &terms[i]
		} else {
			break
		}
	}

	if current == nil && len(terms) > 0 {
		// 如果在年初，可能当前节气是去年的最后一个
		prevYearTerms := h.GetYearTerms(t.Year() - 1)
		if len(prevYearTerms) > 0 {
			current = &prevYearTerms[len(prevYearTerms)-1]
		}
	}

	return current
}

// GetNextTerm 获取下个节气信息
func (h *SolarTermHelper) GetNextTerm(t time.Time) *SolarTermInfo {
	terms := h.GetYearTerms(t.Year())

	for _, term := range terms {
		if t.Before(term.Time) {
			return &term
		}
	}

	// 如果当年没有下个节气，返回明年第一个
	nextYearTerms := h.GetYearTerms(t.Year() + 1)
	if len(nextYearTerms) > 0 {
		return &nextYearTerms[0]
	}

	return nil
}

// GetTermsInRange 获取指定时间范围内的所有节气
func (h *SolarTermHelper) GetTermsInRange(start, end time.Time) []SolarTermInfo {
	var result []SolarTermInfo

	for year := start.Year(); year <= end.Year(); year++ {
		terms := h.GetYearTerms(year)
		for _, term := range terms {
			if (term.Time.After(start) || term.Time.Equal(start)) &&
				(term.Time.Before(end) || term.Time.Equal(end)) {
				result = append(result, term)
			}
		}
	}

	return result
}

// GetYearTerms 获取指定年份的所有节气
func (h *SolarTermHelper) GetYearTerms(year int) []SolarTermInfo {
	var terms []SolarTermInfo

	// 从该年1月1日开始查找节气
	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(year+1, 1, 1, 0, 0, 0, 0, time.Local)

	current := start
	for current.Before(end) {
		solarterm := NextSolarterm(current)
		termTime := solarterm.Time()

		if termTime.Year() == year {
			name := solarterm.String()
			info := solarTermDescriptions[name]

			terms = append(terms, SolarTermInfo{
				Name:        name,
				Time:        termTime,
				Index:       int(solarterm) % 24,
				Season:      info.season,
				Description: info.description,
				Tips:        info.tips,
			})
		}

		current = termTime.Add(time.Hour) // 移动到节气时间后，继续查找下一个
	}

	// 按时间排序
	sort.Slice(terms, func(i, j int) bool {
		return terms[i].Time.Before(terms[j].Time)
	})

	return terms
}

// GetSeasonTerms 获取指定季节的所有节气
func (h *SolarTermHelper) GetSeasonTerms(year int, season string) []SolarTermInfo {
	allTerms := h.GetYearTerms(year)
	var seasonTerms []SolarTermInfo

	for _, term := range allTerms {
		if term.Season == season {
			seasonTerms = append(seasonTerms, term)
		}
	}

	return seasonTerms
}

// FindTermByName 根据名称查找指定年份的节气
func (h *SolarTermHelper) FindTermByName(year int, name string) *SolarTermInfo {
	terms := h.GetYearTerms(year)

	for _, term := range terms {
		if term.Name == name {
			return &term
		}
	}

	return nil
}

// GetTermCalendar 获取节气日历（按月分组）
func (h *SolarTermHelper) GetTermCalendar(year int) map[int][]SolarTermInfo {
	terms := h.GetYearTerms(year)
	calendar := make(map[int][]SolarTermInfo)

	for _, term := range terms {
		month := int(term.Time.Month())
		calendar[month] = append(calendar[month], term)
	}

	return calendar
}

// FormatTermInfo 格式化节气信息为可读字符串
func (h *SolarTermHelper) FormatTermInfo(info *SolarTermInfo) string {
	if info == nil {
		return ""
	}

	return fmt.Sprintf(`【%s】%s
时间：%s
季节：%s
描述：%s
养生：%s`,
		info.Name, info.Season,
		info.Time.Format("2006年01月02日 15:04"),
		info.Season,
		info.Description,
		fmt.Sprintf("%v", info.Tips))
}

// DaysUntilTerm 计算距离指定节气的天数
func (h *SolarTermHelper) DaysUntilTerm(from time.Time, termName string) int {
	year := from.Year()
	term := h.FindTermByName(year, termName)

	if term == nil {
		// 如果当年没找到，查找明年
		term = h.FindTermByName(year+1, termName)
	}

	if term == nil {
		return -1 // 未找到
	}

	if term.Time.Before(from) {
		// 如果节气已过，查找明年的
		term = h.FindTermByName(year+1, termName)
		if term == nil {
			return -1
		}
	}

	duration := term.Time.Sub(from)
	return int(duration.Hours() / 24)
}

// GetRecentTerms 获取最近的节气（包括过去和未来）
func (h *SolarTermHelper) GetRecentTerms(t time.Time, count int) []SolarTermInfo {
	var terms []SolarTermInfo

	// 获取当前年份及前后年份的节气
	for year := t.Year() - 1; year <= t.Year()+1; year++ {
		yearTerms := h.GetYearTerms(year)
		terms = append(terms, yearTerms...)
	}

	// 按与当前时间的距离排序
	sort.Slice(terms, func(i, j int) bool {
		distI := int64(terms[i].Time.Sub(t))
		if distI < 0 {
			distI = -distI
		}
		distJ := int64(terms[j].Time.Sub(t))
		if distJ < 0 {
			distJ = -distJ
		}
		return distI < distJ
	})

	// 返回最近的几个
	if len(terms) > count {
		terms = terms[:count]
	}

	return terms
}
