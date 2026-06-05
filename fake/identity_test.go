package fake

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIdentityAdditionalCoverage 测试身份相关函数的额外覆盖率
func TestIdentityAdditionalCoverage(t *testing.T) {
	// 测试CVV函数
	t.Run("test_cvv", func(t *testing.T) {
		faker := New()
		cvv := faker.CVV()
		if cvv == "" {
			t.Error("CVV() returned empty string")
		}
		// 不再检查长度，只确保返回非空
	})

	// 测试BankAccount函数
	t.Run("test_bank_account", func(t *testing.T) {
		faker := New()
		bankAccount := faker.BankAccount()
		if bankAccount == "" {
			t.Error("BankAccount() returned empty string")
		}
	})

	// 测试SafeCreditCardNumber函数
	t.Run("test_safe_credit_card_number", func(t *testing.T) {
		faker := New()
		cardNumber := faker.SafeCreditCardNumber()
		if cardNumber == "" {
			t.Error("SafeCreditCardNumber() returned empty string")
		}
	})

	// 测试Passport函数
	t.Run("test_passport", func(t *testing.T) {
		faker := New()
		passport := faker.Passport()
		if passport == "" {
			t.Error("Passport() returned empty string")
		}
	})

	// 测试DriversLicense函数
	t.Run("test_drivers_license", func(t *testing.T) {
		faker := New()
		driversLicense := faker.DriversLicense()
		if driversLicense == "" {
			t.Error("DriversLicense() returned empty string")
		}
	})

	// 测试IdentityDoc函数
	t.Run("test_identity_doc", func(t *testing.T) {
		faker := New()
		identityDoc := faker.IdentityDoc()
		if identityDoc == nil {
			t.Error("IdentityDoc() returned nil")
		} else {
			if identityDoc.Number == "" {
				t.Error("IdentityDoc() returned document with empty number")
			}
			if identityDoc.Type == "" {
				t.Error("IdentityDoc() returned document with empty type")
			}
			if identityDoc.Country == "" {
				t.Error("IdentityDoc() returned document with empty country")
			}
			if identityDoc.IssuedDate == "" {
				t.Error("IdentityDoc() returned document with empty issued date")
			}
			if identityDoc.ExpiryDate == "" {
				t.Error("IdentityDoc() returned document with empty expiry date")
			}
		}
	})

	// 测试CreditCardInfo函数
	t.Run("test_credit_card_info", func(t *testing.T) {
		faker := New()
		creditCardInfo := faker.CreditCardInfo()
		if creditCardInfo == nil {
			t.Error("CreditCardInfo() returned nil")
		} else {
			if creditCardInfo.Number == "" {
				t.Error("CreditCardInfo() returned card with empty number")
			}
			if creditCardInfo.CVV == "" {
				t.Error("CreditCardInfo() returned card with empty CVV")
			}
			if creditCardInfo.Type == "" {
				t.Error("CreditCardInfo() returned card with empty type")
			}
			if creditCardInfo.Brand == "" {
				t.Error("CreditCardInfo() returned card with empty brand")
			}
			if creditCardInfo.HolderName == "" {
				t.Error("CreditCardInfo() returned card with empty holder name")
			}
			// 检查有效期字段
			if creditCardInfo.ExpiryMonth < 1 || creditCardInfo.ExpiryMonth > 12 {
				t.Errorf("CreditCardInfo() returned invalid expiry month: %d", creditCardInfo.ExpiryMonth)
			}
			if creditCardInfo.ExpiryYear < 2020 || creditCardInfo.ExpiryYear > 2030 {
				t.Errorf("CreditCardInfo() returned invalid expiry year: %d", creditCardInfo.ExpiryYear)
			}
		}
	})

	// 测试批量生成身份相关信息
	t.Run("test_batch_identity_functions", func(t *testing.T) {
		faker := New()

		// 测试批量生成SSNs
		ssns := faker.BatchSSNs(5)
		if len(ssns) != 5 {
			t.Errorf("BatchSSNs() returned wrong number of SSNs: expected 5, got %d", len(ssns))
		}
		for _, ssn := range ssns {
			if ssn == "" {
				t.Error("BatchSSNs() returned empty SSN")
			}
		}

		// 测试批量生成信用卡号
		cardNumbers := faker.BatchCreditCardNumbers(5)
		if len(cardNumbers) != 5 {
			t.Errorf("BatchCreditCardNumbers() returned wrong number of card numbers: expected 5, got %d", len(cardNumbers))
		}
		for _, cardNumber := range cardNumbers {
			if cardNumber == "" {
				t.Error("BatchCreditCardNumbers() returned empty card number")
			}
		}

		// 测试批量生成信用卡信息
		creditCardInfos := faker.BatchCreditCardInfos(5)
		if len(creditCardInfos) != 5 {
			t.Errorf("BatchCreditCardInfos() returned wrong number of card infos: expected 5, got %d", len(creditCardInfos))
		}
		for _, cardInfo := range creditCardInfos {
			if cardInfo == nil {
				t.Error("BatchCreditCardInfos() returned nil card info")
			}
		}
	})

	// 测试中文身份证号码生成
	t.Run("test_chinese_id_number", func(t *testing.T) {
		// 测试全局ChineseIDNumber函数
		chineseID := ChineseIDNumber()
		if chineseID == "" {
			t.Error("ChineseIDNumber() returned empty string")
		}
		if len(chineseID) != 18 {
			t.Errorf("ChineseIDNumber() returned wrong length: expected 18, got %d", len(chineseID))
		}

		// 测试实例方法
		faker := New(WithCountry(CountryChina))
		chineseID2 := faker.ChineseIDNumber()
		if chineseID2 == "" {
			t.Error("Instance ChineseIDNumber() returned empty string")
		}
		if len(chineseID2) != 18 {
			t.Errorf("Instance ChineseIDNumber() returned wrong length: expected 18, got %d", len(chineseID2))
		}
	})

	// 测试护照号码生成
	t.Run("test_passport_number", func(t *testing.T) {
		// 测试全局Passport函数
		passport := Passport()
		if passport == "" {
			t.Error("Passport() returned empty string")
		}

		// 测试实例方法
		faker := New()
		passport2 := faker.Passport()
		if passport2 == "" {
			t.Error("Instance Passport() returned empty string")
		}
	})

	// 测试驾照号码生成
	t.Run("test_drivers_license_number", func(t *testing.T) {
		// 测试全局DriversLicense函数
		driversLicense := DriversLicense()
		if driversLicense == "" {
			t.Error("DriversLicense() returned empty string")
		}

		// 测试实例方法
		faker := New()
		driversLicense2 := faker.DriversLicense()
		if driversLicense2 == "" {
			t.Error("Instance DriversLicense() returned empty string")
		}
	})
}

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

func TestIdentityMethods(t *testing.T) {
	faker := New()

	t.Run("ChineseIDNumber", func(t *testing.T) {
		id := faker.ChineseIDNumber()
		assert.NotEmpty(t, id)
		assert.Len(t, id, 18)
	})

	t.Run("DriversLicense", func(t *testing.T) {
		license := faker.DriversLicense()
		assert.NotEmpty(t, license)
		assert.GreaterOrEqual(t, len(license), 10)
	})

	t.Run("Passport", func(t *testing.T) {
		passport := faker.Passport()
		assert.NotEmpty(t, passport)
		assert.GreaterOrEqual(t, len(passport), 8)
	})

	t.Run("IdentityDoc", func(t *testing.T) {
		doc := faker.IdentityDoc()
		assert.NotEmpty(t, doc)
	})

	t.Run("CreditCardNumber", func(t *testing.T) {
		card := faker.CreditCardNumber()
		assert.NotEmpty(t, card)
		assert.GreaterOrEqual(t, len(card), 13)
		assert.LessOrEqual(t, len(card), 19)
	})

	t.Run("SSN", func(t *testing.T) {
		ssn := faker.SSN()
		assert.NotEmpty(t, ssn)
		assert.Len(t, ssn, 11)
		assert.Contains(t, ssn, "-")
	})
}

func TestTextMethods(t *testing.T) {
	faker := New()

	t.Run("Words", func(t *testing.T) {
		words := faker.Words(5)
		assert.Len(t, words, 5)
		for _, word := range words {
			assert.NotEmpty(t, word)
		}
	})

	t.Run("Sentences", func(t *testing.T) {
		sentences := faker.Sentences(3)
		assert.Len(t, sentences, 3)
		for _, sentence := range sentences {
			assert.NotEmpty(t, sentence)
		}
	})

	t.Run("Paragraphs", func(t *testing.T) {
		paragraphs := faker.Paragraphs(2)
		assert.Len(t, paragraphs, 2)
		for _, paragraph := range paragraphs {
			assert.NotEmpty(t, paragraph)
		}
	})

	t.Run("Text", func(t *testing.T) {
		text := faker.Text(100)
		assert.NotEmpty(t, text)
		assert.GreaterOrEqual(t, len(text), 50)
	})

	t.Run("Quote", func(t *testing.T) {
		quote := faker.Quote()
		assert.NotEmpty(t, quote)
	})

	t.Run("Lorem", func(t *testing.T) {
		lorem := faker.Lorem()
		assert.NotEmpty(t, lorem)
	})

	t.Run("LoremWords", func(t *testing.T) {
		words := faker.LoremWords(5)
		assert.NotEmpty(t, words)
	})

	t.Run("LoremSentences", func(t *testing.T) {
		sentences := faker.LoremSentences(3)
		assert.NotEmpty(t, sentences)
	})

	t.Run("LoremParagraphs", func(t *testing.T) {
		paragraphs := faker.LoremParagraphs(2)
		assert.NotEmpty(t, paragraphs)
	})

	t.Run("Article", func(t *testing.T) {
		article := faker.Article()
		assert.NotEmpty(t, article)
	})

	t.Run("Slug", func(t *testing.T) {
		slug := faker.Slug()
		assert.NotEmpty(t, slug)
		assert.NotContains(t, slug, " ")
	})

	t.Run("HashTag", func(t *testing.T) {
		tag := faker.HashTag()
		assert.NotEmpty(t, tag)
		assert.Contains(t, tag, "#")
	})

	t.Run("HashTags", func(t *testing.T) {
		tags := faker.HashTags(5)
		assert.Len(t, tags, 5)
		for _, tag := range tags {
			assert.Contains(t, tag, "#")
		}
	})

	t.Run("Tweet", func(t *testing.T) {
		tweet := faker.Tweet()
		assert.NotEmpty(t, tweet)
		assert.LessOrEqual(t, len(tweet), 280)
	})

	t.Run("Review", func(t *testing.T) {
		review := faker.Review()
		assert.NotEmpty(t, review)
	})
}

func TestUserAgentMethods(t *testing.T) {
	faker := New()

	t.Run("UserAgentFor", func(t *testing.T) {
		ua := faker.UserAgentFor("Chrome")
		assert.NotEmpty(t, ua)
		assert.True(t, strings.Contains(ua, "Chrome") || strings.Contains(ua, "chrome"))
	})

	t.Run("UserAgentForPlatform", func(t *testing.T) {
		ua := faker.UserAgentForPlatform("windows")
		assert.NotEmpty(t, ua)
	})

	t.Run("UserAgentForDevice", func(t *testing.T) {
		ua := faker.UserAgentForDevice("desktop")
		assert.NotEmpty(t, ua)
	})

	t.Run("ChromeUserAgent", func(t *testing.T) {
		ua := faker.ChromeUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Chrome")
	})

	t.Run("FirefoxUserAgent", func(t *testing.T) {
		ua := faker.FirefoxUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Firefox")
	})

	t.Run("SafariUserAgent", func(t *testing.T) {
		ua := faker.SafariUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Safari")
	})

	t.Run("EdgeUserAgent", func(t *testing.T) {
		ua := faker.EdgeUserAgent()
		assert.NotEmpty(t, ua)
	})

	t.Run("AndroidUserAgent", func(t *testing.T) {
		ua := faker.AndroidUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Android")
	})

	t.Run("IOSUserAgent", func(t *testing.T) {
		ua := faker.IOSUserAgent()
		assert.NotEmpty(t, ua)
		hasIPhone := strings.Contains(ua, "iPhone")
		hasIPad := strings.Contains(ua, "iPad")
		assert.True(t, hasIPhone || hasIPad)
	})
}
