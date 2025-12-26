package fake

import (
	"testing"
)

// TestGlobalIdentityFunctionsAdditional 测试identity.go中的全局便捷函数
func TestGlobalIdentityFunctionsAdditional(t *testing.T) {
	// 测试SSN全局函数
	ssn := SSN()
	if ssn == "" {
		t.Error("Global SSN() should not return empty string")
	}

	// 测试ChineseIDNumber全局函数
	chineseID := ChineseIDNumber()
	if chineseID == "" {
		t.Error("Global ChineseIDNumber() should not return empty string")
	}

	// 测试Passport全局函数
	passport := Passport()
	if passport == "" {
		t.Error("Global Passport() should not return empty string")
	}

	// 测试DriversLicense全局函数
	driversLicense := DriversLicense()
	if driversLicense == "" {
		t.Error("Global DriversLicense() should not return empty string")
	}

	// 测试IdentityDoc全局函数
	identityDoc := IdentityDoc()
	if identityDoc == nil {
		t.Error("Global IdentityDoc() should not return nil")
	} else if identityDoc.Number == "" {
		t.Error("Global IdentityDoc() should return a document with a number")
	}

	// 测试CreditCardNumber全局函数
	creditCardNumber := CreditCardNumber()
	if creditCardNumber == "" {
		t.Error("Global CreditCardNumber() should not return empty string")
	}

	// 测试CVV全局函数
	cvv := CVV()
	if cvv == "" {
		t.Error("Global CVV() should not return empty string")
	}

	// 测试BankAccount全局函数
	bankAccount := BankAccount()
	if bankAccount == "" {
		t.Error("Global BankAccount() should not return empty string")
	}

	// 测试IBAN全局函数
	iban := IBAN()
	if iban == "" {
		t.Error("Global IBAN() should not return empty string")
	}

	// 测试CreditCardInfo全局函数
	creditCardInfo := CreditCardInfo()
	if creditCardInfo == nil {
		t.Error("Global CreditCardInfo() should not return nil")
	} else if creditCardInfo.Number == "" {
		t.Error("Global CreditCardInfo() should return a credit card with a number")
	}

	// 测试SafeCreditCardNumber全局函数
	safeCreditCardNumber := SafeCreditCardNumber()
	if safeCreditCardNumber == "" {
		t.Error("Global SafeCreditCardNumber() should not return empty string")
	}
}

// TestIdentityMethodsAdditional 测试identity.go中的实例方法
func TestIdentityMethodsAdditional(t *testing.T) {
	faker := New()

	// 测试SSN方法
	ssn := faker.SSN()
	if ssn == "" {
		t.Error("SSN() should not return empty string")
	}

	// 测试ChineseIDNumber方法
	chineseID := faker.ChineseIDNumber()
	if chineseID == "" {
		t.Error("ChineseIDNumber() should not return empty string")
	}

	// 测试Passport方法
	passport := faker.Passport()
	if passport == "" {
		t.Error("Passport() should not return empty string")
	}

	// 测试DriversLicense方法
	driversLicense := faker.DriversLicense()
	if driversLicense == "" {
		t.Error("DriversLicense() should not return empty string")
	}

	// 测试IdentityDoc方法
	identityDoc := faker.IdentityDoc()
	if identityDoc == nil {
		t.Error("IdentityDoc() should not return nil")
	} else if identityDoc.Number == "" {
		t.Error("IdentityDoc() should return a document with a number")
	}

	// 测试CreditCardNumber方法
	creditCardNumber := faker.CreditCardNumber()
	if creditCardNumber == "" {
		t.Error("CreditCardNumber() should not return empty string")
	}

	// 测试CVV方法
	cvv := faker.CVV()
	if cvv == "" {
		t.Error("CVV() should not return empty string")
	}

	// 测试BankAccount方法
	bankAccount := faker.BankAccount()
	if bankAccount == "" {
		t.Error("BankAccount() should not return empty string")
	}

	// 测试IBAN方法
	iban := faker.IBAN()
	if iban == "" {
		t.Error("IBAN() should not return empty string")
	}

	// 测试CreditCardInfo方法
	creditCardInfo := faker.CreditCardInfo()
	if creditCardInfo == nil {
		t.Error("CreditCardInfo() should not return nil")
	} else if creditCardInfo.Number == "" {
		t.Error("CreditCardInfo() should return a credit card with a number")
	}

	// 测试SafeCreditCardNumber方法
	safeCreditCardNumber := faker.SafeCreditCardNumber()
	if safeCreditCardNumber == "" {
		t.Error("SafeCreditCardNumber() should not return empty string")
	}
}

// TestBatchIdentityFunctionsAdditional 测试批量生成函数
func TestBatchIdentityFunctionsAdditional(t *testing.T) {
	faker := New()

	// 测试BatchSSNs方法
	ssns := faker.BatchSSNs(5)
	if len(ssns) != 5 {
		t.Errorf("BatchSSNs(5) should return 5 items, got %d", len(ssns))
	}
	for i, ssn := range ssns {
		if ssn == "" {
			t.Errorf("SSN at index %d should not be empty", i)
		}
	}

	// 测试BatchCreditCardNumbers方法
	creditCardNumbers := faker.BatchCreditCardNumbers(5)
	if len(creditCardNumbers) != 5 {
		t.Errorf("BatchCreditCardNumbers(5) should return 5 items, got %d", len(creditCardNumbers))
	}
	for i, cc := range creditCardNumbers {
		if cc == "" {
			t.Errorf("Credit card number at index %d should not be empty", i)
		}
	}

	// 测试BatchCreditCardInfos方法
	creditCardInfos := faker.BatchCreditCardInfos(5)
	if len(creditCardInfos) != 5 {
		t.Errorf("BatchCreditCardInfos(5) should return 5 items, got %d", len(creditCardInfos))
	}
	for i, cci := range creditCardInfos {
		if cci == nil || cci.Number == "" {
			t.Errorf("Credit card info at index %d should not be nil or have empty number", i)
		}
	}
}