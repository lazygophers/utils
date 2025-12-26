package fake

import (
	"testing"
)

// TestAdditionalCoverage 测试覆盖率较低的函数
func TestAdditionalCoverage(t *testing.T) {
	// 测试NamePrefix函数
	t.Run("test_name_prefix", func(t *testing.T) {
		prefix := NamePrefix()
		if prefix == "" {
			t.Error("NamePrefix() should not return empty string")
		}
	})

	// 测试Tweet函数
	t.Run("test_tweet", func(t *testing.T) {
		tweet := Tweet()
		if tweet == "" {
			t.Error("Tweet() should not return empty string")
		}
	})

	// 测试UserAgent函数
	t.Run("test_user_agent", func(t *testing.T) {
		userAgent := UserAgent()
		if userAgent == "" {
			t.Error("UserAgent() should not return empty string")
		}
	})

	// 测试FormattedName函数
	t.Run("test_formatted_name", func(t *testing.T) {
		formattedName := FormattedName()
		if formattedName == "" {
			t.Error("FormattedName() should not return empty string")
		}
	})

	// 测试Pool相关函数
	t.Run("test_pool_functions", func(t *testing.T) {
		// 测试GetPoolStats
		GetPoolStats()
		// 只检查函数执行，不验证具体值

		// 测试WarmupPools
		WarmupPools()
	})

	// 测试BatchNamesOptimized
	t.Run("test_batch_names_optimized", func(t *testing.T) {
		WithPooledFaker(func(f *Faker) {
			names := f.BatchNamesOptimized(10)
			if len(names) != 10 {
				t.Errorf("Expected 10 names, got %d", len(names))
			}
		})
	})

	// 测试IncrementOnlyName
	t.Run("test_increment_only_name", func(t *testing.T) {
		inc := NewIncrementOnly()
		name1 := inc.IncrementOnlyName()
		name2 := inc.IncrementOnlyName()
		// IncrementOnlyName() may return the same name on consecutive calls,
		// but it should not return empty string
		if name1 == "" {
			t.Error("IncrementOnlyName() should not return empty string")
		}
		if name2 == "" {
			t.Error("IncrementOnlyName() should not return empty string")
		}
	})

	// 测试StaticName
	t.Run("test_static_name", func(t *testing.T) {
		static := NewStatic()
		name1 := static.StaticName()
		name2 := static.StaticName()
		if name1 != name2 {
			t.Error("StaticName() should return the same name on consecutive calls")
		}
	})
}

// TestBankAccountAdditional 测试BankAccount函数
func TestBankAccountAdditional(t *testing.T) {
	// 测试BankAccount函数
	for i := 0; i < 10; i++ {
		bankAccount := BankAccount()
		if bankAccount == "" {
			t.Error("BankAccount() should not return empty string")
		}
	}
}

// TestSafeCreditCardNumberAdditional 测试SafeCreditCardNumber函数
func TestSafeCreditCardNumberAdditional(t *testing.T) {
	// 测试SafeCreditCardNumber函数
	for i := 0; i < 10; i++ {
		cardNumber := SafeCreditCardNumber()
		if cardNumber == "" {
			t.Error("SafeCreditCardNumber() should not return empty string")
		}
		// 验证返回的是卡号格式
		if len(cardNumber) < 13 {
			t.Error("SafeCreditCardNumber() should return a valid credit card number format")
		}
	}
}

// TestUserAgentAdditional 测试UserAgent函数
func TestUserAgentAdditional(t *testing.T) {
	// 测试UserAgent函数
	userAgent := UserAgent()
	if userAgent == "" {
		t.Error("UserAgent() should not return empty string")
	}

	// 测试GenerateUserAgent函数
	gen := NewUserAgentGenerator()
	ua := gen.GenerateUserAgent(UserAgentOptions{Browser: "Chrome", Platform: "Windows"})
	if ua == "" {
		t.Error("GenerateUserAgent() should not return empty string")
	}
}

// TestTweetAdditional 测试Tweet函数
func TestTweetAdditional(t *testing.T) {
	// 测试Tweet函数
	tweet := Tweet()
	if tweet == "" {
		t.Error("Tweet() should not return empty string")
	}
	// 推文长度应该在合理范围内
	if len(tweet) > 280 {
		t.Error("Tweet() should return a string with length <= 280")
	}
}

// TestNameFunctionsAdditional 测试名称相关函数
func TestNameFunctionsAdditional(t *testing.T) {
	// 测试FormattedName函数
	formattedName := FormattedName()
	if formattedName == "" {
		t.Error("FormattedName() should not return empty string")
	}

	// 测试NamePrefix函数
	prefix := NamePrefix()
	if prefix == "" {
		t.Error("NamePrefix() should not return empty string")
	}
}
