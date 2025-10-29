package fake

import (
	"fmt"
	"strconv"
	"time"

	"github.com/lazygophers/utils/randx"
)

// IdentityDocument 身份证件信息结构体
type IdentityDocument struct {
	Type       string `json:"type"`
	Number     string `json:"number"`
	Country    string `json:"country"`
	IssuedDate string `json:"issued_date"`
	ExpiryDate string `json:"expiry_date"`
	IsValid    bool   `json:"is_valid"`
}

// CreditCard 银行卡信息结构体
type CreditCard struct {
	Number      string `json:"number"`
	Type        string `json:"type"`
	Brand       string `json:"brand"`
	CVV         string `json:"cvv"`
	ExpiryMonth int    `json:"expiry_month"`
	ExpiryYear  int    `json:"expiry_year"`
	HolderName  string `json:"holder_name"`
	IsValid     bool   `json:"is_valid"`
}

// SSN 生成美国社会安全号码
func (f *Faker) SSN() string {

	// SSN格式: XXX-XX-XXXX
	// 第一部分: 001-899 (避免 000, 666, 900-999)
	area := randx.Intn(899) + 1
	if area == 666 {
		area = 665 // 避免666
	}

	// 第二部分: 01-99
	group := randx.Intn(99) + 1

	// 第三部分: 0001-9999
	serial := randx.Intn(9999) + 1

	return fmt.Sprintf("%03d-%02d-%04d", area, group, serial)
}

// ChineseIDNumber 生成中国身份证号码
func (f *Faker) ChineseIDNumber() string {

	// 地区代码 (前6位)
	areaCodes := []string{
		"110101", "110102", "110105", "110106", // 北京
		"310101", "310104", "310105", "310106", // 上海
		"440101", "440103", "440104", "440105", // 广州
		"440301", "440303", "440304", "440305", // 深圳
		"120101", "120102", "120103", "120104", // 天津
		"500101", "500102", "500103", "500104", // 重庆
		"320101", "320102", "320104", "320105", // 南京
		"330101", "330102", "330103", "330104", // 杭州
		"510101", "510104", "510105", "510106", // 成都
		"420101", "420102", "420103", "420104", // 武汉
	}
	areaCode := randx.Choose(areaCodes)

	// 出生日期 (第7-14位)
	year := randx.Intn(50) + 1960 // 1960-2009
	month := randx.Intn(12) + 1
	day := randx.Intn(28) + 1 // 简化处理，使用1-28天
	birthDate := fmt.Sprintf("%04d%02d%02d", year, month, day)

	// 顺序码 (第15-17位)
	sequence := randx.Intn(999) + 1

	// 前17位
	id17 := fmt.Sprintf("%s%s%03d", areaCode, birthDate, sequence)

	// 计算校验码 (第18位)
	checksum := f.calculateChineseIDChecksum(id17)

	return id17 + checksum
}

func (f *Faker) calculateChineseIDChecksum(id17 string) string {
	// 校验码计算权重
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 校验码对应表
	checksums := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

	sum := 0
	for i, char := range id17 {
		digit, _ := strconv.Atoi(string(char))
		sum += digit * weights[i]
	}

	return checksums[sum%11]
}

// Passport 生成护照号码
func (f *Faker) Passport() string {

	switch f.country {
	case CountryUS:
		// 美国护照: 9位数字
		return fmt.Sprintf("%09d", randx.Intn(999999999))

	case CountryUK:
		// 英国护照: 9位数字
		return fmt.Sprintf("%09d", randx.Intn(999999999))

	case CountryChina:
		// 中国护照: E + 8位数字
		return fmt.Sprintf("E%08d", randx.Intn(99999999))

	case CountryCanada:
		// 加拿大护照: 2个字母 + 6位数字
		letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		return fmt.Sprintf("%c%c%06d",
			letters[randx.Intn(len(letters))],
			letters[randx.Intn(len(letters))],
			randx.Intn(999999))

	default:
		// 通用格式: 字母 + 数字
		letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		return fmt.Sprintf("%c%08d",
			letters[randx.Intn(len(letters))],
			randx.Intn(99999999))
	}
}

// DriversLicense 生成驾照号码
func (f *Faker) DriversLicense() string {

	switch f.country {
	case CountryUS:
		// 美国驾照格式因州而异，使用通用格式
		letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		digits := randx.Intn(99999999)
		return fmt.Sprintf("%c%c%08d",
			letters[randx.Intn(len(letters))],
			letters[randx.Intn(len(letters))],
			digits)

	case CountryChina:
		// 中国驾照: 地区码 + 12位数字
		areaCodes := []string{"1101", "3101", "4401", "4403", "1201", "5001"}
		areaCode := randx.Choose(areaCodes)
		return fmt.Sprintf("%s%012d", areaCode, randx.Intn(999999999999))

	default:
		return fmt.Sprintf("%012d", randx.Intn(999999999999))
	}
}

// IdentityDoc 生成身份证件信息
func (f *Faker) IdentityDoc() *IdentityDocument {

	docTypes := []string{"ID Card", "Passport", "Driver License", "Social Security"}
	docType := randx.Choose(docTypes)

	var number string
	switch docType {
	case "ID Card":
		if f.country == CountryChina {
			number = f.ChineseIDNumber()
		} else {
			number = f.SSN()
		}
	case "Passport":
		number = f.Passport()
	case "Driver License":
		number = f.DriversLicense()
	case "Social Security":
		number = f.SSN()
	}

	// 生成发证日期和过期日期
	issuedDate := time.Now().AddDate(-randx.Intn(10), -randx.Intn(12), -randx.Intn(30))
	expiryDate := issuedDate.AddDate(randx.Intn(10)+5, 0, 0) // 5-15年有效期

	return &IdentityDocument{
		Type:       docType,
		Number:     number,
		Country:    string(f.country),
		IssuedDate: issuedDate.Format("2006-01-02"),
		ExpiryDate: expiryDate.Format("2006-01-02"),
		IsValid:    time.Now().Before(expiryDate),
	}
}

// CreditCardNumber 生成信用卡号码
func (f *Faker) CreditCardNumber() string {

	// 信用卡品牌前缀
	prefixes := map[string][]string{
		"Visa":             {"4"},
		"MasterCard":       {"5"},
		"American Express": {"34", "37"},
		"Discover":         {"6"},
		"JCB":              {"35"},
		"UnionPay":         {"62"},
	}

	brands := []string{"Visa", "MasterCard", "American Express", "Discover", "JCB", "UnionPay"}
	brand := randx.Choose(brands)
	brandPrefixes := prefixes[brand]
	prefix := randx.Choose(brandPrefixes)

	var totalLength int
	switch brand {
	case "American Express":
		totalLength = 15
	case "Visa", "MasterCard", "Discover", "JCB", "UnionPay":
		totalLength = 16
	default:
		totalLength = 16
	}

	// 生成剩余数字
	remaining := totalLength - len(prefix) - 1 // -1 for check digit
	numberPart := prefix

	for i := 0; i < remaining; i++ {
		numberPart += fmt.Sprintf("%d", randx.Intn(10))
	}

	// 计算Luhn校验位
	checkDigit := f.calculateLuhnCheckDigit(numberPart)

	return numberPart + fmt.Sprintf("%d", checkDigit)
}

func (f *Faker) calculateLuhnCheckDigit(number string) int {
	sum := 0
	alternate := true

	// 从右到左遍历
	for i := len(number) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(number[i]))

		if alternate {
			digit *= 2
			if digit > 9 {
				digit = (digit % 10) + 1
			}
		}

		sum += digit
		alternate = !alternate
	}

	return (10 - (sum % 10)) % 10
}

// CVV 生成信用卡CVV码
func (f *Faker) CVV() string {

	// 大部分卡是3位，American Express是4位
	if randx.Float32() < 0.1 { // 10% 概率是4位
		return fmt.Sprintf("%04d", randx.Intn(10000))
	}
	return fmt.Sprintf("%03d", randx.Intn(1000))
}

// BankAccount 生成银行账号
func (f *Faker) BankAccount() string {

	switch f.country {
	case CountryUS:
		// 美国银行账号: 8-17位数字
		length := randx.Intn(10) + 8
		return fmt.Sprintf("%0*d", length, randx.Intn(int(1e17)))

	case CountryChina:
		// 中国银行账号: 16-19位数字
		length := randx.Intn(4) + 16
		return fmt.Sprintf("%0*d", length, randx.Intn(int(1e18)))

	default:
		// 通用格式: 10-16位数字
		length := randx.Intn(7) + 10
		return fmt.Sprintf("%0*d", length, randx.Intn(int(1e15)))
	}
}

// IBAN 生成国际银行账号 (International Bank Account Number)
func (f *Faker) IBAN() string {

	countryCodes := map[Country]string{
		CountryGermany: "DE",
		CountryFrance:  "FR",
		CountryUK:      "GB",
		CountrySpain:   "ES",
		CountryItaly:   "IT",
	}

	countryCode, exists := countryCodes[f.country]
	if !exists {
		countryCode = "DE" // 默认德国
	}

	// 生成银行代码和账号
	bankCode := fmt.Sprintf("%08d", randx.Intn(99999999))
	accountNumber := fmt.Sprintf("%010d", randx.Intn(9999999999))

	// IBAN校验码计算简化版本
	checkDigits := fmt.Sprintf("%02d", randx.Intn(100))

	return fmt.Sprintf("%s%s%s%s", countryCode, checkDigits, bankCode, accountNumber)
}

// CreditCardInfo 生成完整信用卡信息
func (f *Faker) CreditCardInfo() *CreditCard {

	number := f.CreditCardNumber()
	cvv := f.CVV()

	// 根据卡号确定品牌
	var brand, cardType string
	switch {
	case number[0] == '4':
		brand = "Visa"
		cardType = randx.Choose([]string{"Credit", "Debit", "Prepaid"})
	case number[0] == '5':
		brand = "MasterCard"
		cardType = randx.Choose([]string{"Credit", "Debit", "Prepaid"})
	case number[:2] == "34" || number[:2] == "37":
		brand = "American Express"
		cardType = "Credit"
	case number[0] == '6':
		brand = "Discover"
		cardType = "Credit"
	case number[:2] == "35":
		brand = "JCB"
		cardType = "Credit"
	case number[:2] == "62":
		brand = "UnionPay"
		cardType = randx.Choose([]string{"Credit", "Debit"})
	default:
		brand = "Unknown"
		cardType = "Credit"
	}

	// 生成过期日期
	currentMonth := int(time.Now().Month())
	currentYear := time.Now().Year()

	expiryMonth := randx.Intn(12) + 1
	expiryYear := currentYear + randx.Intn(5) + 1 // 1-5年后过期

	// 检查是否有效
	isValid := true
	if expiryYear == currentYear && expiryMonth <= currentMonth {
		isValid = false
	}

	return &CreditCard{
		Number:      number,
		Type:        cardType,
		Brand:       brand,
		CVV:         cvv,
		ExpiryMonth: expiryMonth,
		ExpiryYear:  expiryYear,
		HolderName:  f.Name(),
		IsValid:     isValid,
	}
}

// SafeCreditCardNumber 生成测试用的安全信用卡号码
func (f *Faker) SafeCreditCardNumber() string {

	// 使用测试卡号前缀，这些不是真实的信用卡号码
	testPrefixes := []string{
		"4000000000000000", // Visa测试卡 (16位)
		"5555555555554444", // MasterCard测试卡 (16位)
		"378282246310005",  // American Express测试卡 (15位)
		"6011111111111117", // Discover测试卡 (16位)
	}

	chosen := randx.Choose(testPrefixes)
	// 对于American Express (15位)，直接返回；对于其他卡，截取前16位
	if len(chosen) == 15 {
		return chosen
	}
	if len(chosen) >= 16 {
		return chosen[:16]
	}
	return chosen
}

// 批量生成函数
func (f *Faker) BatchSSNs(count int) []string {
	return f.batchGenerate(count, f.SSN)
}

func (f *Faker) BatchCreditCardNumbers(count int) []string {
	return f.batchGenerate(count, f.CreditCardNumber)
}

func (f *Faker) BatchCreditCardInfos(count int) []*CreditCard {
	results := make([]*CreditCard, count)
	for i := 0; i < count; i++ {
		results[i] = f.CreditCardInfo()
	}
	return results
}

// 全局便捷函数
func SSN() string {
	return getDefaultFaker().SSN()
}

func ChineseIDNumber() string {
	return getDefaultFaker().ChineseIDNumber()
}

func Passport() string {
	return getDefaultFaker().Passport()
}

func DriversLicense() string {
	return getDefaultFaker().DriversLicense()
}

func IdentityDoc() *IdentityDocument {
	return getDefaultFaker().IdentityDoc()
}

func CreditCardNumber() string {
	return getDefaultFaker().CreditCardNumber()
}

func CVV() string {
	return getDefaultFaker().CVV()
}

func BankAccount() string {
	return getDefaultFaker().BankAccount()
}

func IBAN() string {
	return getDefaultFaker().IBAN()
}

func CreditCardInfo() *CreditCard {
	return getDefaultFaker().CreditCardInfo()
}

func SafeCreditCardNumber() string {
	return getDefaultFaker().SafeCreditCardNumber()
}
