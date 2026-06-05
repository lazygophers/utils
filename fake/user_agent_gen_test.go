package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestUserAgentGenerator 测试UserAgentGenerator的功能
func TestUserAgentGenerator(t *testing.T) {
	gen := NewUserAgentGenerator()

	// 测试基本的用户代理生成
	t.Run("test_basic_generate", func(t *testing.T) {
		opts := UserAgentOptions{
			Browser:  "Chrome",
			Platform: PlatformWindows,
			OS:       "Windows NT",
		}

		ua := gen.GenerateUserAgent(opts)
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Chrome")
		assert.Contains(t, ua, "Windows NT")
	})

	// 测试根据设备类型推断平台
	t.Run("test_infer_platform", func(t *testing.T) {
		opts := UserAgentOptions{
			DeviceType: DeviceTypeMobile,
			Browser:    "Chrome",
		}

		ua := gen.GenerateUserAgent(opts)
		assert.NotEmpty(t, ua)
	})

	// 测试不同浏览器的用户代理生成
	t.Run("test_different_browsers", func(t *testing.T) {
		browsers := []string{"Chrome", "Firefox", "Safari", "Edge", "Opera"}
		for _, browser := range browsers {
			opts := UserAgentOptions{
				Browser:  browser,
				Platform: PlatformWindows,
			}

			ua := gen.GenerateUserAgent(opts)
			assert.NotEmpty(t, ua)
			assert.Contains(t, ua, browser)
		}
	})

	// 测试不同平台的用户代理生成
	t.Run("test_different_platforms", func(t *testing.T) {
		platforms := []Platform{PlatformWindows, PlatformMacOS, PlatformLinux, PlatformAndroid, PlatformIOS}
		for _, platform := range platforms {
			opts := UserAgentOptions{
				Browser:  "Chrome",
				Platform: platform,
			}

			ua := gen.GenerateUserAgent(opts)
			assert.NotEmpty(t, ua)
		}
	})
}

// TestUserAgentMethodsExtended 测试Faker的用户代理生成方法
func TestUserAgentMethodsExtended(t *testing.T) {
	faker := New()

	// 测试GenerateUserAgent方法
	t.Run("test_generate_user_agent", func(t *testing.T) {
		opts := UserAgentOptions{
			Browser:  "Chrome",
			Platform: PlatformWindows,
		}

		ua := faker.GenerateUserAgent(opts)
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Chrome")
	})

	// 测试GenerateRandomUserAgent方法
	t.Run("test_generate_random_user_agent", func(t *testing.T) {
		ua := faker.GenerateRandomUserAgent()
		assert.NotEmpty(t, ua)
	})

	// 测试UserAgentFor方法
	t.Run("test_user_agent_for", func(t *testing.T) {
		testCases := []string{"Chrome", "Firefox", "Safari", "Edge", "Opera"}
		for _, browser := range testCases {
			ua := faker.UserAgentFor(browser)
			assert.NotEmpty(t, ua)
		}
	})

	// 测试UserAgentForPlatform方法
	t.Run("test_user_agent_for_platform", func(t *testing.T) {
		testCases := []Platform{PlatformWindows, PlatformMacOS, PlatformLinux, PlatformAndroid, PlatformIOS}
		for _, platform := range testCases {
			ua := faker.UserAgentForPlatform(platform)
			assert.NotEmpty(t, ua)
		}
	})

	// 测试UserAgentForDevice方法
	t.Run("test_user_agent_for_device", func(t *testing.T) {
		testCases := []DeviceType{DeviceTypeDesktop, DeviceTypeLaptop, DeviceTypeMobile, DeviceTypeTablet}
		for _, deviceType := range testCases {
			ua := faker.UserAgentForDevice(deviceType)
			assert.NotEmpty(t, ua)
		}
	})

	// 测试特定浏览器的便捷方法
	t.Run("test_browser_specific_methods", func(t *testing.T) {
		// Chrome
		ua := faker.ChromeUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Chrome")

		// Firefox
		ua = faker.FirefoxUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Firefox")

		// Safari
		ua = faker.SafariUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Safari")

		// Edge
		ua = faker.EdgeUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Edge")
	})

	// 测试特定平台的便捷方法
	t.Run("test_platform_specific_methods", func(t *testing.T) {
		// Android
		ua := faker.AndroidUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Android")

		// iOS
		ua = faker.IOSUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "iPhone")
	})
}

// TestGlobalUserAgentFunctions 测试全局用户代理生成函数
func TestGlobalUserAgentFunctions(t *testing.T) {
	// 测试GenerateUserAgent全局函数
	t.Run("test_global_generate_user_agent", func(t *testing.T) {
		opts := UserAgentOptions{
			Browser:  "Chrome",
			Platform: PlatformWindows,
		}

		ua := GenerateUserAgent(opts)
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Chrome")
	})

	// 测试GenerateRandomUserAgent全局函数
	t.Run("test_global_generate_random_user_agent", func(t *testing.T) {
		ua := GenerateRandomUserAgent()
		assert.NotEmpty(t, ua)
	})

	// 测试UserAgentFor全局函数
	t.Run("test_global_user_agent_for", func(t *testing.T) {
		ua := UserAgentFor("Chrome")
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Chrome")
	})

	// 测试UserAgentForPlatform全局函数
	t.Run("test_global_user_agent_for_platform", func(t *testing.T) {
		ua := UserAgentForPlatform(PlatformWindows)
		assert.NotEmpty(t, ua)
	})

	// 测试UserAgentForDevice全局函数
	t.Run("test_global_user_agent_for_device", func(t *testing.T) {
		ua := UserAgentForDevice(DeviceTypeMobile)
		assert.NotEmpty(t, ua)
	})

	// 测试特定浏览器的全局便捷函数
	t.Run("test_global_browser_specific_functions", func(t *testing.T) {
		// Chrome
		ua := ChromeUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Chrome")

		// Firefox
		ua = FirefoxUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Firefox")

		// Safari
		ua = SafariUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Safari")

		// Edge
		ua = EdgeUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Edge")

		// Android
		ua = AndroidUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Android")

		// iOS
		ua = IOSUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "iPhone")
	})
}

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
