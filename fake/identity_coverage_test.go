package fake

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试美国社会安全号码生成
func TestSSN(t *testing.T) {
	f := New()
	ssn := f.SSN()
	assert.NotEmpty(t, ssn)
	assert.Len(t, ssn, 11) // XXX-XX-XXXX格式
	assert.Contains(t, ssn, "-")
}

// 测试中国身份证号码生成
func TestChineseIDNumber(t *testing.T) {
	f := New()
	id := f.ChineseIDNumber()
	assert.NotEmpty(t, id)
	assert.Len(t, id, 18) // 中国身份证18位
}

// 测试护照号码生成
func TestPassport(t *testing.T) {
	f := New()
	passport := f.Passport()
	assert.NotEmpty(t, passport)

	// 测试不同国家的护照格式
	f = New(WithCountry(CountryUS))
	usPassport := f.Passport()
	assert.NotEmpty(t, usPassport)

	f = New(WithCountry(CountryChina))
	chinaPassport := f.Passport()
	assert.NotEmpty(t, chinaPassport)
	assert.Contains(t, chinaPassport, "E") // 中国护照以E开头

	f = New(WithCountry(CountryCanada))
	canadaPassport := f.Passport()
	assert.NotEmpty(t, canadaPassport)
}

// 测试驾照号码生成
func TestDriversLicense(t *testing.T) {
	f := New()
	license := f.DriversLicense()
	assert.NotEmpty(t, license)

	// 测试不同国家的驾照格式
	f = New(WithCountry(CountryUS))
	usLicense := f.DriversLicense()
	assert.NotEmpty(t, usLicense)

	f = New(WithCountry(CountryChina))
	chinaLicense := f.DriversLicense()
	assert.NotEmpty(t, chinaLicense)
}

// 测试身份证件信息生成
func TestIdentityDoc(t *testing.T) {
	f := New()
	doc := f.IdentityDoc()
	assert.NotNil(t, doc)
	assert.NotEmpty(t, doc.Type)
	assert.NotEmpty(t, doc.Number)
	assert.NotEmpty(t, doc.Country)
	assert.NotEmpty(t, doc.IssuedDate)
	assert.NotEmpty(t, doc.ExpiryDate)

	// 测试不同国家的身份证件
	f = New(WithCountry(CountryChina))
	chinaDoc := f.IdentityDoc()
	assert.NotNil(t, chinaDoc)
	assert.Equal(t, string(CountryChina), chinaDoc.Country)
}

// 测试信用卡号码生成
func TestCreditCardNumber(t *testing.T) {
	f := New()
	cardNumber := f.CreditCardNumber()
	assert.NotEmpty(t, cardNumber)
	// 信用卡号码长度应为15或16位
	assert.True(t, len(cardNumber) == 15 || len(cardNumber) == 16)
}

// 测试信用卡CVV码生成
func TestCVV(t *testing.T) {
	f := New()
	cvv := f.CVV()
	assert.NotEmpty(t, cvv)
	// CVV应为3或4位
	assert.True(t, (len(cvv) == 3 || len(cvv) == 4), "CVV长度应为3或4位，实际为%d位: %s", len(cvv), cvv)
}

// 测试银行账号生成
func TestIdentityBankAccount(t *testing.T) {
	f := New()
	account := f.BankAccount()
	assert.NotEmpty(t, account)

	// 测试美国银行账号
	f = New(WithCountry(CountryUS))
	usAccount := f.BankAccount()
	assert.NotEmpty(t, usAccount)

	// 测试中国银行账号
	f = New(WithCountry(CountryChina))
	chinaAccount := f.BankAccount()
	assert.NotEmpty(t, chinaAccount)

	// 测试其他国家银行账号
	f = New(WithCountry(CountryUK))
	ukAccount := f.BankAccount()
	assert.NotEmpty(t, ukAccount)
}

// 测试国际银行账号生成
func TestIdentityIBAN(t *testing.T) {
	f := New()
	iban := f.IBAN()
	assert.NotEmpty(t, iban)
	// 标准IBAN长度在15-34位之间
	assert.True(t, len(iban) >= 15 && len(iban) <= 34, "IBAN长度应在15-34位之间，实际为%d位: %s", len(iban), iban)

	// 测试不同国家的IBAN格式
	f = New(WithCountry(CountryGermany))
	deIBAN := f.IBAN()
	assert.NotEmpty(t, deIBAN)
	assert.True(t, strings.HasPrefix(deIBAN, "DE"))

	f = New(WithCountry(CountryFrance))
	frIBAN := f.IBAN()
	assert.NotEmpty(t, frIBAN)
	assert.True(t, strings.HasPrefix(frIBAN, "FR"))

	f = New(WithCountry(CountryUK))
	gbIBAN := f.IBAN()
	assert.NotEmpty(t, gbIBAN)
	assert.True(t, strings.HasPrefix(gbIBAN, "GB"))
}

// 测试完整信用卡信息生成
func TestIdentityCreditCardInfo(t *testing.T) {
	f := New()
	cardInfo := f.CreditCardInfo()
	assert.NotNil(t, cardInfo)
	assert.NotEmpty(t, cardInfo.Number)
	assert.NotEmpty(t, cardInfo.Type)
	assert.NotEmpty(t, cardInfo.Brand)
	assert.NotEmpty(t, cardInfo.CVV)
	assert.Greater(t, cardInfo.ExpiryMonth, 0)
	assert.LessOrEqual(t, cardInfo.ExpiryMonth, 12)
	assert.Greater(t, cardInfo.ExpiryYear, 0)
	assert.NotEmpty(t, cardInfo.HolderName)

	// 测试不同品牌信用卡的识别
	for i := 0; i < 10; i++ {
		cardInfo := f.CreditCardInfo()
		assert.NotNil(t, cardInfo)
		assert.NotEmpty(t, cardInfo.Brand)
	}
}

// 测试安全信用卡号码生成
func TestIdentitySafeCreditCardNumber(t *testing.T) {
	f := New()
	safeCardNumber := f.SafeCreditCardNumber()
	assert.NotEmpty(t, safeCardNumber)
	// 安全信用卡号码应为15或16位
	assert.True(t, len(safeCardNumber) == 15 || len(safeCardNumber) == 16)

	// 测试多次调用，确保返回不同的测试卡号
	cardNumbers := make(map[string]bool)
	for i := 0; i < 5; i++ {
		cardNumber := f.SafeCreditCardNumber()
		cardNumbers[cardNumber] = true
	}
	assert.Greater(t, len(cardNumbers), 1) // 至少返回两种不同的卡号
}

// 测试批量生成函数
func TestIdentityBatchSSNs(t *testing.T) {
	f := New()
	ssns := f.BatchSSNs(5)
	assert.Len(t, ssns, 5)
	for _, ssn := range ssns {
		assert.NotEmpty(t, ssn)
		assert.Len(t, ssn, 11)
	}
}

func TestIdentityBatchCreditCardNumbers(t *testing.T) {
	f := New()
	cardNumbers := f.BatchCreditCardNumbers(5)
	assert.Len(t, cardNumbers, 5)
	for _, cardNumber := range cardNumbers {
		assert.NotEmpty(t, cardNumber)
	}
}

func TestIdentityBatchCreditCardInfos(t *testing.T) {
	f := New()
	cardInfos := f.BatchCreditCardInfos(5)
	assert.Len(t, cardInfos, 5)
	for _, cardInfo := range cardInfos {
		assert.NotNil(t, cardInfo)
	}
}

// 测试全局函数
func TestGlobalIdentityFunctions(t *testing.T) {
	// 测试全局SSN函数
	ssn := SSN()
	assert.NotEmpty(t, ssn)

	// 测试全局ChineseIDNumber函数
	chineseID := ChineseIDNumber()
	assert.NotEmpty(t, chineseID)

	// 测试全局Passport函数
	passport := Passport()
	assert.NotEmpty(t, passport)

	// 测试全局DriversLicense函数
	license := DriversLicense()
	assert.NotEmpty(t, license)

	// 测试全局IdentityDoc函数
	identityDoc := IdentityDoc()
	assert.NotNil(t, identityDoc)

	// 测试全局CreditCardNumber函数
	creditCardNumber := CreditCardNumber()
	assert.NotEmpty(t, creditCardNumber)

	// 测试全局CVV函数
	cvv := CVV()
	assert.NotEmpty(t, cvv)

	// 测试全局BankAccount函数
	bankAccount := BankAccount()
	assert.NotEmpty(t, bankAccount)

	// 测试全局IBAN函数
	iban := IBAN()
	assert.NotEmpty(t, iban)

	// 测试全局CreditCardInfo函数
	creditCardInfo := CreditCardInfo()
	assert.NotNil(t, creditCardInfo)

	// 测试全局SafeCreditCardNumber函数
	safeCreditCardNumber := SafeCreditCardNumber()
	assert.NotEmpty(t, safeCreditCardNumber)
}

// 测试Luhn校验位计算
func TestCalculateLuhnCheckDigit(t *testing.T) {
	f := New()
	// 使用已知的有效卡号前缀测试
	checkDigit := f.calculateLuhnCheckDigit("411111111111111")
	assert.Equal(t, 1, checkDigit) // 4111111111111111是有效的Visa卡号

	checkDigit = f.calculateLuhnCheckDigit("555555555555444")
	assert.Equal(t, 4, checkDigit) // 5555555555554444是有效的MasterCard卡号
}

// 测试中国身份证校验码计算
func TestCalculateChineseIDChecksum(t *testing.T) {
	f := New()
	// 使用已知的前17位身份证号码测试
	id17 := "11010119900307123"
	checksum := f.calculateChineseIDChecksum(id17)
	// 计算结果应该是一个有效的校验码
	assert.NotEmpty(t, checksum)
	assert.Len(t, checksum, 1)
}
