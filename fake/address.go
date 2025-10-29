package fake

import (
	"fmt"

	"github.com/lazygophers/utils/randx"
)

// Address 地址信息结构体
type Address struct {
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state,omitempty"`
	Province    string `json:"province,omitempty"`
	ZipCode     string `json:"zip_code"`
	Country     string `json:"country"`
	FullAddress string `json:"full_address"`
}

// Street 生成街道地址
func (f *Faker) Street() string {

	values, weights, err := getWeightedItems(f.language, "addresses", "streets")
	if err != nil {
		// 回退到英语
		if f.language != LanguageEnglish {
			values, weights, err = getWeightedItems(LanguageEnglish, "addresses", "streets")
		}
		if err != nil {
			return getDefaultStreet()
		}
	}

	street := randx.WeightedChoose(values, weights)

	// 生成街道号码
	streetNumber := randx.Intn(9999) + 1

	return fmt.Sprintf("%d %s", streetNumber, street)
}

// City 生成城市名称
func (f *Faker) City() string {

	values, weights, err := getWeightedItems(f.language, "addresses", "cities")
	if err != nil {
		// 回退到英语
		if f.language != LanguageEnglish {
			values, weights, err = getWeightedItems(LanguageEnglish, "addresses", "cities")
		}
		if err != nil {
			return getDefaultCity()
		}
	}

	return randx.WeightedChoose(values, weights)
}

// State 生成州/省份名称
func (f *Faker) State() string {

	switch f.country {
	case CountryUS:
		return f.generateUSState()
	case CountryChina:
		return f.generateChineseProvince()
	case CountryCanada:
		return f.generateCanadianProvince()
	default:
		return f.generateGenericState()
	}
}

func (f *Faker) generateUSState() string {
	states := []string{
		"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado",
		"Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho",
		"Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana",
		"Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota",
		"Mississippi", "Missouri", "Montana", "Nebraska", "Nevada",
		"New Hampshire", "New Jersey", "New Mexico", "New York",
		"North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon",
		"Pennsylvania", "Rhode Island", "South Carolina", "South Dakota",
		"Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington",
		"West Virginia", "Wisconsin", "Wyoming",
	}
	return randx.Choose(states)
}

func (f *Faker) generateChineseProvince() string {
	provinces := []string{
		"北京市", "天津市", "上海市", "重庆市",
		"河北省", "山西省", "辽宁省", "吉林省", "黑龙江省",
		"江苏省", "浙江省", "安徽省", "福建省", "江西省", "山东省",
		"河南省", "湖北省", "湖南省", "广东省", "海南省",
		"四川省", "贵州省", "云南省", "陕西省", "甘肃省", "青海省",
		"内蒙古自治区", "广西壮族自治区", "西藏自治区", "宁夏回族自治区", "新疆维吾尔自治区",
		"香港特别行政区", "澳门特别行政区", "台湾省",
	}
	return randx.Choose(provinces)
}

func (f *Faker) generateCanadianProvince() string {
	provinces := []string{
		"Alberta", "British Columbia", "Manitoba", "New Brunswick",
		"Newfoundland and Labrador", "Northwest Territories", "Nova Scotia",
		"Nunavut", "Ontario", "Prince Edward Island", "Quebec",
		"Saskatchewan", "Yukon",
	}
	return randx.Choose(provinces)
}

func (f *Faker) generateGenericState() string {
	return ""
}

// ZipCode 生成邮政编码
func (f *Faker) ZipCode() string {

	switch f.country {
	case CountryUS:
		return fmt.Sprintf("%05d", randx.Intn(99999))
	case CountryChina:
		return fmt.Sprintf("%06d", randx.Intn(999999))
	case CountryUK:
		letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		return fmt.Sprintf("%c%c%d %d%c%c",
			letters[randx.Intn(len(letters))],
			letters[randx.Intn(len(letters))],
			randx.Intn(10),
			randx.Intn(10),
			letters[randx.Intn(len(letters))],
			letters[randx.Intn(len(letters))],
		)
	case CountryCanada:
		letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		return fmt.Sprintf("%c%d%c %d%c%d",
			letters[randx.Intn(len(letters))],
			randx.Intn(10),
			letters[randx.Intn(len(letters))],
			randx.Intn(10),
			letters[randx.Intn(len(letters))],
			randx.Intn(10),
		)
	default:
		return fmt.Sprintf("%05d", randx.Intn(99999))
	}
}

// CountryName 生成国家名称
func (f *Faker) CountryName() string {

	switch f.language {
	case LanguageEnglish:
		return f.generateEnglishCountryName()
	case LanguageChineseSimplified:
		return f.generateChineseCountryName()
	case LanguageChineseTraditional:
		return f.generateTraditionalChineseCountryName()
	default:
		return f.generateEnglishCountryName()
	}
}

func (f *Faker) generateEnglishCountryName() string {
	countries := []string{
		"United States", "China", "Japan", "Germany", "United Kingdom",
		"France", "India", "Italy", "Brazil", "Canada", "Russia", "South Korea",
		"Spain", "Australia", "Mexico", "Indonesia", "Netherlands", "Saudi Arabia",
		"Turkey", "Taiwan", "Belgium", "Argentina", "Thailand", "Bangladesh",
		"Egypt", "Nigeria", "Vietnam", "Malaysia", "South Africa", "Philippines",
		"Chile", "Finland", "Romania", "Czech Republic", "New Zealand", "Portugal",
		"Peru", "Greece", "Iraq", "Algeria", "Qatar", "Kazakhstan", "Hungary",
		"Kuwait", "Morocco", "Ecuador", "Ethiopia", "Slovakia", "Kenya", "Angola",
	}
	return randx.Choose(countries)
}

func (f *Faker) generateChineseCountryName() string {
	countries := []string{
		"中国", "美国", "日本", "德国", "英国", "法国", "印度", "意大利",
		"巴西", "加拿大", "俄罗斯", "韩国", "西班牙", "澳大利亚", "墨西哥",
		"印度尼西亚", "荷兰", "沙特阿拉伯", "土耳其", "台湾", "比利时",
		"阿根廷", "泰国", "孟加拉国", "埃及", "尼日利亚", "越南", "马来西亚",
		"南非", "菲律宾", "智利", "芬兰", "罗马尼亚", "捷克", "新西兰",
		"葡萄牙", "秘鲁", "希腊", "伊拉克", "阿尔及利亚", "卡塔尔",
		"哈萨克斯坦", "匈牙利", "科威特", "摩洛哥", "厄瓜多尔", "埃塞俄比亚",
		"斯洛伐克", "肯尼亚", "安哥拉",
	}
	return randx.Choose(countries)
}

func (f *Faker) generateTraditionalChineseCountryName() string {
	countries := []string{
		"中國", "美國", "日本", "德國", "英國", "法國", "印度", "意大利",
		"巴西", "加拿大", "俄羅斯", "韓國", "西班牙", "澳大利亞", "墨西哥",
		"印度尼西亞", "荷蘭", "沙特阿拉伯", "土耳其", "台灣", "比利時",
		"阿根廷", "泰國", "孟加拉國", "埃及", "尼日利亞", "越南", "馬來西亞",
		"南非", "菲律賓", "智利", "芬蘭", "羅馬尼亞", "捷克", "新西蘭",
		"葡萄牙", "秘魯", "希臘", "伊拉克", "阿爾及利亞", "卡塔爾",
		"哈薩克斯坦", "匈牙利", "科威特", "摩洛哥", "厄瓜多爾", "埃塞俄比亞",
		"斯洛伐克", "肯尼亞", "安哥拉",
	}
	return randx.Choose(countries)
}

// FullAddress 生成完整地址
func (f *Faker) FullAddress() *Address {

	street := f.Street()
	city := f.City()
	state := f.State()
	zipCode := f.ZipCode()
	country := f.CountryName()

	var fullAddress string
	switch f.language {
	case LanguageChineseSimplified, LanguageChineseTraditional:
		// 中文地址格式：国家 省份 城市 街道
		if state != "" {
			fullAddress = fmt.Sprintf("%s %s %s %s", country, state, city, street)
		} else {
			fullAddress = fmt.Sprintf("%s %s %s", country, city, street)
		}
	default:
		// 英文地址格式：街道, 城市, 州 邮编, 国家
		if state != "" {
			fullAddress = fmt.Sprintf("%s, %s, %s %s, %s", street, city, state, zipCode, country)
		} else {
			fullAddress = fmt.Sprintf("%s, %s %s, %s", street, city, zipCode, country)
		}
	}

	return &Address{
		Street:      street,
		City:        city,
		State:       state,
		ZipCode:     zipCode,
		Country:     country,
		FullAddress: fullAddress,
	}
}

// AddressLine 生成地址行（不包含国家）
func (f *Faker) AddressLine() string {

	street := f.Street()
	city := f.City()
	state := f.State()
	zipCode := f.ZipCode()

	switch f.language {
	case LanguageChineseSimplified, LanguageChineseTraditional:
		if state != "" {
			return fmt.Sprintf("%s %s %s %s", state, city, street, zipCode)
		}
		return fmt.Sprintf("%s %s %s", city, street, zipCode)
	default:
		if state != "" {
			return fmt.Sprintf("%s, %s, %s %s", street, city, state, zipCode)
		}
		return fmt.Sprintf("%s, %s %s", street, city, zipCode)
	}
}

// Latitude 生成纬度
func (f *Faker) Latitude() float64 {
	return (randx.Float64() - 0.5) * 180
}

// Longitude 生成经度
func (f *Faker) Longitude() float64 {
	return (randx.Float64() - 0.5) * 360
}

// Coordinate 生成坐标对
func (f *Faker) Coordinate() (float64, float64) {
	return f.Latitude(), f.Longitude()
}

// 批量生成函数
func (f *Faker) BatchStreets(count int) []string {
	return f.batchGenerate(count, f.Street)
}

func (f *Faker) BatchCities(count int) []string {
	return f.batchGenerate(count, f.City)
}

func (f *Faker) BatchStates(count int) []string {
	return f.batchGenerate(count, f.State)
}

func (f *Faker) BatchZipCodes(count int) []string {
	return f.batchGenerate(count, f.ZipCode)
}

func (f *Faker) BatchFullAddresses(count int) []*Address {
	results := make([]*Address, count)
	for i := 0; i < count; i++ {
		results[i] = f.FullAddress()
	}
	return results
}

// 默认值回退函数
func getDefaultStreet() string {
	streets := []string{"Main Street", "Park Avenue", "Oak Street", "Pine Street", "Maple Avenue"}
	streetNumber := randx.Intn(9999) + 1
	return fmt.Sprintf("%d %s", streetNumber, randx.Choose(streets))
}

func getDefaultCity() string {
	cities := []string{"New York", "Los Angeles", "Chicago", "Houston", "Phoenix"}
	return randx.Choose(cities)
}

// 全局便捷函数
func Street() string {
	return getDefaultFaker().Street()
}

func City() string {
	return getDefaultFaker().City()
}

func State() string {
	return getDefaultFaker().State()
}

func ZipCode() string {
	return getDefaultFaker().ZipCode()
}

func CountryName() string {
	return getDefaultFaker().CountryName()
}

func FullAddress() *Address {
	return getDefaultFaker().FullAddress()
}

func AddressLine() string {
	return getDefaultFaker().AddressLine()
}

func Latitude() float64 {
	return getDefaultFaker().Latitude()
}

func Longitude() float64 {
	return getDefaultFaker().Longitude()
}

func Coordinate() (float64, float64) {
	return getDefaultFaker().Coordinate()
}
