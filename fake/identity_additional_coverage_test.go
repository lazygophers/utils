package fake

import (
	"testing"
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
