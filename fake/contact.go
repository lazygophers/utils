package fake

import (
	"fmt"
	"strings"

	"github.com/lazygophers/utils/randx"
)

// PhoneNumber 生成电话号码
func (f *Faker) PhoneNumber() string {

	switch f.country {
	case CountryUS, CountryCanada:
		return f.generateNorthAmericanPhone()
	case CountryChina:
		return f.generateChinesePhone()
	case CountryUK:
		return f.generateUKPhone()
	case CountryFrance:
		return f.generateFrenchPhone()
	case CountryGermany:
		return f.generateGermanPhone()
	case CountryJapan:
		return f.generateJapanesePhone()
	case CountryKorea:
		return f.generateKoreanPhone()
	default:
		return f.generateNorthAmericanPhone()
	}
}

func (f *Faker) generateNorthAmericanPhone() string {
	// 北美电话号码格式: +1 (XXX) XXX-XXXX
	areaCode := randx.Intn(900) + 100 // 100-999
	exchange := randx.Intn(900) + 100 // 100-999
	number := randx.Intn(10000)       // 0000-9999

	formats := []string{
		"+1 (%d) %d-%04d",
		"(%d) %d-%04d",
		"%d-%d-%04d",
		"%d.%d.%04d",
		"%d %d %04d",
	}

	format := randx.Choose(formats)
	return fmt.Sprintf(format, areaCode, exchange, number)
}

func (f *Faker) generateChinesePhone() string {
	// 中国手机号码格式
	prefixes := []string{"130", "131", "132", "133", "134", "135", "136", "137", "138", "139",
		"150", "151", "152", "153", "155", "156", "157", "158", "159",
		"180", "181", "182", "183", "184", "185", "186", "187", "188", "189",
		"170", "171", "172", "173", "174", "175", "176", "177", "178"}

	prefix := randx.Choose(prefixes)
	suffix := randx.Intn(100000000) // 8位数字

	return fmt.Sprintf("+86 %s %08d", prefix, suffix)
}

func (f *Faker) generateUKPhone() string {
	// 英国电话号码格式: +44 XXXX XXXXXX
	areaCode := randx.Intn(9000) + 1000 // 1000-9999
	number := randx.Intn(1000000)       // 000000-999999

	return fmt.Sprintf("+44 %d %06d", areaCode, number)
}

func (f *Faker) generateFrenchPhone() string {
	// 法国电话号码格式: +33 X XX XX XX XX
	first := randx.Intn(9) + 1 // 1-9
	parts := make([]int, 4)
	for i := range parts {
		parts[i] = randx.Intn(100) // 00-99
	}

	return fmt.Sprintf("+33 %d %02d %02d %02d %02d", first, parts[0], parts[1], parts[2], parts[3])
}

func (f *Faker) generateGermanPhone() string {
	// 德国电话号码格式: +49 XXX XXXXXXX
	areaCode := randx.Intn(900) + 100 // 100-999
	number := randx.Intn(10000000)    // 0000000-9999999

	return fmt.Sprintf("+49 %d %07d", areaCode, number)
}

func (f *Faker) generateJapanesePhone() string {
	// 日本电话号码格式: +81 XX-XXXX-XXXX
	area := randx.Intn(90) + 10   // 10-99
	exchange := randx.Intn(10000) // 0000-9999
	number := randx.Intn(10000)   // 0000-9999

	return fmt.Sprintf("+81 %d-%04d-%04d", area, exchange, number)
}

func (f *Faker) generateKoreanPhone() string {
	// 韩国电话号码格式: +82 XX-XXX-XXXX
	area := randx.Intn(90) + 10  // 10-99
	exchange := randx.Intn(1000) // 000-999
	number := randx.Intn(10000)  // 0000-9999

	return fmt.Sprintf("+82 %d-%03d-%04d", area, exchange, number)
}

// MobileNumber 生成手机号码
func (f *Faker) MobileNumber() string {

	switch f.country {
	case CountryUS, CountryCanada:
		// 北美手机号码使用特定的区号
		areaCodes := []int{201, 202, 203, 206, 212, 213, 214, 215, 216, 301, 302, 303, 305, 312, 313, 314, 315, 404, 405, 407, 408, 410, 412, 413, 414, 415, 417, 501, 502, 503, 504, 505, 507, 508, 510, 512, 513, 515, 516, 517, 518, 601, 602, 603, 605, 606, 607, 608, 609, 612, 614, 615, 616, 617, 618, 619, 701, 702, 703, 704, 707, 708, 712, 713, 714, 715, 716, 717, 718, 719, 801, 802, 803, 804, 805, 806, 808, 810, 812, 813, 814, 815, 816, 817, 818, 901, 903, 904, 906, 907, 908, 909, 910, 912, 913, 914, 915, 916, 917, 918, 919}
		areaCode := randx.Choose(areaCodes)
		exchange := randx.Intn(900) + 100
		number := randx.Intn(10000)
		return fmt.Sprintf("+1 (%d) %d-%04d", areaCode, exchange, number)

	case CountryChina:
		return f.generateChinesePhone()

	default:
		return f.PhoneNumber()
	}
}

// Email 生成邮箱地址
func (f *Faker) Email() string {

	username := f.generateEmailUsername()
	domain := f.generateEmailDomain()

	return fmt.Sprintf("%s@%s", username, domain)
}

func (f *Faker) generateEmailUsername() string {
	// 生成用户名部分
	patterns := []string{
		"%s.%s",   // john.doe
		"%s_%s",   // john_doe
		"%s%s",    // johndoe
		"%s.%s%d", // john.doe123
		"%s_%s%d", // john_doe123
		"%s%d",    // john123
	}

	firstName := strings.ToLower(f.FirstName())
	lastName := strings.ToLower(f.LastName())

	// 移除空格和特殊字符
	firstName = strings.ReplaceAll(firstName, " ", "")
	lastName = strings.ReplaceAll(lastName, " ", "")

	// 移除非ASCII字符（针对中文名字等）
	firstName = f.toASCII(firstName)
	lastName = f.toASCII(lastName)

	pattern := randx.Choose(patterns)

	switch pattern {
	case "%s.%s":
		return fmt.Sprintf(pattern, firstName, lastName)
	case "%s_%s":
		return fmt.Sprintf(pattern, firstName, lastName)
	case "%s%s":
		return fmt.Sprintf(pattern, firstName, lastName)
	case "%s.%s%d":
		return fmt.Sprintf(pattern, firstName, lastName, randx.Intn(999)+1)
	case "%s_%s%d":
		return fmt.Sprintf(pattern, firstName, lastName, randx.Intn(999)+1)
	case "%s%d":
		return fmt.Sprintf(pattern, firstName, randx.Intn(9999)+1)
	default:
		return fmt.Sprintf("%s.%s", firstName, lastName)
	}
}

func (f *Faker) generateEmailDomain() string {
	domains := []string{
		"gmail.com", "yahoo.com", "hotmail.com", "outlook.com", "icloud.com",
		"aol.com", "mail.com", "protonmail.com", "yandex.com", "zoho.com",
		"live.com", "msn.com", "qq.com", "163.com", "126.com",
		"sina.com", "sohu.com", "yeah.net", "foxmail.com", "139.com",
		"company.com", "business.net", "enterprise.org", "startup.io", "tech.co",
	}

	return randx.Choose(domains)
}

// CompanyEmail 生成企业邮箱地址
func (f *Faker) CompanyEmail() string {

	username := f.generateEmailUsername()
	domain := f.generateCompanyEmailDomain()

	return fmt.Sprintf("%s@%s", username, domain)
}

func (f *Faker) generateCompanyEmailDomain() string {
	domains := []string{
		"company.com", "corporation.net", "business.org", "enterprise.co",
		"tech.io", "startup.com", "firm.net", "group.org", "holdings.com",
		"solutions.co", "systems.net", "services.com", "consulting.biz",
		"technologies.info", "innovations.org", "dynamics.com", "global.net",
		"international.org", "worldwide.com", "united.co", "associates.net",
	}

	return randx.Choose(domains)
}

// SafeEmail 生成安全的邮箱地址（使用example.com等域名）
func (f *Faker) SafeEmail() string {

	username := f.generateEmailUsername()
	domains := []string{"example.com", "example.org", "example.net", "test.com", "sample.org"}
	domain := randx.Choose(domains)

	return fmt.Sprintf("%s@%s", username, domain)
}

// URL 生成网址
func (f *Faker) URL() string {

	protocols := []string{"http", "https"}
	protocol := randx.Choose(protocols)

	domain := f.generateWebDomain()

	// 30% 概率添加路径
	if randx.Float32() < 0.3 {
		path := f.generateURLPath()
		return fmt.Sprintf("%s://%s/%s", protocol, domain, path)
	}

	return fmt.Sprintf("%s://%s", protocol, domain)
}

func (f *Faker) generateWebDomain() string {
	domains := []string{
		"google.com", "facebook.com", "youtube.com", "twitter.com", "instagram.com",
		"linkedin.com", "reddit.com", "wikipedia.org", "amazon.com", "apple.com",
		"microsoft.com", "netflix.com", "github.com", "stackoverflow.com", "medium.com",
		"example.com", "sample.org", "demo.net", "test.co", "placeholder.io",
	}

	return randx.Choose(domains)
}

func (f *Faker) generateURLPath() string {
	paths := []string{
		"home", "about", "contact", "services", "products", "blog", "news",
		"help", "support", "faq", "login", "register", "profile", "settings",
		"dashboard", "admin", "user", "search", "category", "item", "page",
	}

	// 生成1-3个路径段
	segments := randx.Intn(3) + 1
	pathParts := make([]string, segments)

	for i := 0; i < segments; i++ {
		pathParts[i] = randx.Choose(paths)
	}

	return strings.Join(pathParts, "/")
}

// IPv4 生成IPv4地址
func (f *Faker) IPv4() string {

	return fmt.Sprintf("%d.%d.%d.%d",
		randx.Intn(256),
		randx.Intn(256),
		randx.Intn(256),
		randx.Intn(256))
}

// IPv6 生成IPv6地址
func (f *Faker) IPv6() string {

	parts := make([]string, 8)
	for i := range parts {
		parts[i] = fmt.Sprintf("%04x", randx.Intn(65536))
	}

	return strings.Join(parts, ":")
}

// MAC 生成MAC地址
func (f *Faker) MAC() string {

	parts := make([]string, 6)
	for i := range parts {
		parts[i] = fmt.Sprintf("%02x", randx.Intn(256))
	}

	return strings.Join(parts, ":")
}

// 辅助函数：将非ASCII字符转换为ASCII
func (f *Faker) toASCII(s string) string {
	// 简单的中文转拼音映射（示例）
	chineseToASCII := map[rune]string{
		'李': "li", '王': "wang", '张': "zhang", '刘': "liu", '陈': "chen",
		'杨': "yang", '赵': "zhao", '黄': "huang", '周': "zhou", '吴': "wu",
		'伟': "wei", '强': "qiang", '明': "ming", '军': "jun", '杰': "jie",
		'芳': "fang", '秀': "xiu", '敏': "min", '静': "jing", '丽': "li",
	}

	var result strings.Builder
	for _, r := range s {
		if r < 128 {
			result.WriteRune(r)
		} else if ascii, ok := chineseToASCII[r]; ok {
			result.WriteString(ascii)
		}
		// 忽略其他非ASCII字符
	}

	resultStr := result.String()
	if resultStr == "" {
		// 如果结果为空，生成一个随机的ASCII字符串
		return fmt.Sprintf("user%d", randx.Intn(10000))
	}

	return resultStr
}

// 批量生成函数
func (f *Faker) BatchPhoneNumbers(count int) []string {
	return f.batchGenerate(count, f.PhoneNumber)
}

func (f *Faker) BatchEmails(count int) []string {
	return f.batchGenerate(count, f.Email)
}

func (f *Faker) BatchURLs(count int) []string {
	return f.batchGenerate(count, f.URL)
}

// 全局便捷函数
func PhoneNumber() string {
	return getDefaultFaker().PhoneNumber()
}

func MobileNumber() string {
	return getDefaultFaker().MobileNumber()
}

func Email() string {
	return getDefaultFaker().Email()
}

func CompanyEmail() string {
	return getDefaultFaker().CompanyEmail()
}

func SafeEmail() string {
	return getDefaultFaker().SafeEmail()
}

func URL() string {
	return getDefaultFaker().URL()
}

func IPv4() string {
	return getDefaultFaker().IPv4()
}

func IPv6() string {
	return getDefaultFaker().IPv6()
}

func MAC() string {
	return getDefaultFaker().MAC()
}
