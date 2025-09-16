package fake

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"testing"
)

// TestFaker 测试基本的 Faker 功能
func TestFaker(t *testing.T) {
	faker := New()
	if faker == nil {
		t.Fatal("New() returned nil")
	}

	// 测试默认值
	if faker.language != LanguageEnglish {
		t.Errorf("Expected default language %s, got %s", LanguageEnglish, faker.language)
	}

	if faker.country != CountryUS {
		t.Errorf("Expected default country %s, got %s", CountryUS, faker.country)
	}
}

// TestFakerWithOptions 测试带选项的 Faker
func TestFakerWithOptions(t *testing.T) {
	faker := New(
		WithLanguage(LanguageChineseSimplified),
		WithCountry(CountryChina),
		WithGender(GenderMale),
		WithSeed(12345),
	)

	if faker.language != LanguageChineseSimplified {
		t.Errorf("Expected language %s, got %s", LanguageChineseSimplified, faker.language)
	}

	if faker.country != CountryChina {
		t.Errorf("Expected country %s, got %s", CountryChina, faker.country)
	}

	if faker.gender != GenderMale {
		t.Errorf("Expected gender %s, got %s", GenderMale, faker.gender)
	}
}

// TestWithContext 测试上下文支持
func TestWithContext(t *testing.T) {
	ctx := context.Background()
	ctx = ContextWithLanguage(ctx, LanguageFrench)
	ctx = ContextWithCountry(ctx, CountryFrance)
	ctx = ContextWithGender(ctx, GenderFemale)

	faker := WithContext(ctx)

	if faker.language != LanguageFrench {
		t.Errorf("Expected language from context %s, got %s", LanguageFrench, faker.language)
	}

	if faker.country != CountryFrance {
		t.Errorf("Expected country from context %s, got %s", CountryFrance, faker.country)
	}

	if faker.gender != GenderFemale {
		t.Errorf("Expected gender from context %s, got %s", GenderFemale, faker.gender)
	}
}

// TestStats 测试统计功能
func TestStats(t *testing.T) {
	faker := New()

	// 初始统计应该为0
	stats := faker.Stats()
	if stats["call_count"] != 0 {
		t.Errorf("Expected initial call_count 0, got %d", stats["call_count"])
	}

	// 调用一些方法
	_ = faker.Name()
	_ = faker.Email()

	stats = faker.Stats()
	// 由于 Name() 和 Email() 内部可能调用其他方法，所以计数可能大于2
	if stats["call_count"] < 2 {
		t.Errorf("Expected call_count >= 2, got %d", stats["call_count"])
	}
}

// TestClone 测试克隆功能
func TestClone(t *testing.T) {
	original := New(WithLanguage(LanguageChineseSimplified))
	clone := original.Clone()

	if clone.language != original.language {
		t.Error("Clone should preserve language")
	}

	// 克隆应该有独立的统计
	_ = original.Name()
	originalStats := original.Stats()
	cloneStats := clone.Stats()

	if originalStats["call_count"] == cloneStats["call_count"] {
		t.Error("Clone should have independent stats")
	}
}

// TestClearCache 测试缓存清理
func TestClearCache(t *testing.T) {
	faker := New()

	// 生成一些数据以填充缓存
	_ = faker.Name()
	_ = faker.Email()

	// 清理缓存
	faker.ClearCache()

	// 缓存应该被清空
	// 这里我们主要测试不会崩溃
	_ = faker.Name()
}

// TestBasicFunctionality 测试基本功能
func TestBasicFunctionality(t *testing.T) {
	faker := New()

	// 测试名字生成
	name := faker.Name()
	if name == "" {
		t.Error("Name should not be empty")
	}

	// 测试邮箱生成
	email := faker.Email()
	if email == "" || !strings.Contains(email, "@") {
		t.Errorf("Invalid email: %s", email)
	}

	// 测试电话生成
	phone := faker.PhoneNumber()
	if phone == "" {
		t.Error("Phone should not be empty")
	}

	// 测试地址生成
	address := faker.FullAddress()
	if address == nil || address.FullAddress == "" {
		t.Error("Address should not be empty")
	}
}

// TestMultiLanguage 测试多语言支持
func TestMultiLanguage(t *testing.T) {
	// 英语
	enFaker := New(WithLanguage(LanguageEnglish))
	enName := enFaker.Name()
	if enName == "" {
		t.Error("English name should not be empty")
	}

	// 中文
	cnFaker := New(WithLanguage(LanguageChineseSimplified))
	cnName := cnFaker.Name()
	if cnName == "" {
		t.Error("Chinese name should not be empty")
	}

	t.Logf("EN Name: %s", enName)
	t.Logf("CN Name: %s", cnName)
}

// TestGlobalFunctions 测试全局函数
func TestGlobalFunctions(t *testing.T) {
	name := Name()
	if name == "" {
		t.Error("Global Name() should not be empty")
	}

	email := Email()
	if email == "" || !strings.Contains(email, "@") {
		t.Errorf("Global Email() invalid: %s", email)
	}

	userAgent := RandomUserAgent()
	if userAgent == "" {
		t.Error("Global RandomUserAgent() should not be empty")
	}
}

// TestAllFeaturesWork 验证所有主要功能是否正常工作
func TestAllFeaturesWork(t *testing.T) {
	faker := New()

	// 基础数据生成
	tests := []struct {
		name string
		fn   func() string
	}{
		{"Name", faker.Name},
		{"FirstName", faker.FirstName},
		{"LastName", faker.LastName},
		{"Email", faker.Email},
		{"PhoneNumber", faker.PhoneNumber},
		{"Street", faker.Street},
		{"City", faker.City},
		{"ZipCode", faker.ZipCode},
		{"CountryName", faker.CountryName},
		{"CompanyName", faker.CompanyName},
		{"JobTitle", faker.JobTitle},
		{"Industry", faker.Industry},
		{"SSN", faker.SSN},
		{"CreditCardNumber", faker.CreditCardNumber},
		{"IPv4", faker.IPv4},
		{"MAC", faker.MAC},
		{"URL", faker.URL},
		{"Word", faker.Word},
		{"Sentence", faker.Sentence},
		{"Title", faker.Title},
		{"UserAgent", faker.UserAgent},
		{"Browser", faker.Browser},
		{"OS", faker.OS},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.fn()
			if result == "" {
				t.Errorf("%s returned empty string", test.name)
			}
		})
	}
}

// TestComplexObjectsWork 验证复杂对象生成是否正常
func TestComplexObjectsWork(t *testing.T) {
	faker := New()

	// 测试复杂对象
	address := faker.FullAddress()
	if address == nil || address.FullAddress == "" {
		t.Error("FullAddress should not be nil or empty")
	}

	company := faker.CompanyInfo()
	if company == nil || company.Name == "" {
		t.Error("CompanyInfo should not be nil or have empty name")
	}

	creditCard := faker.CreditCardInfo()
	if creditCard == nil || creditCard.Number == "" {
		t.Error("CreditCardInfo should not be nil or have empty number")
	}

	device := faker.DeviceInfo()
	if device == nil || device.Type == "" {
		t.Error("DeviceInfo should not be nil or have empty type")
	}

	identity := faker.IdentityDoc()
	if identity == nil || identity.Number == "" {
		t.Error("IdentityDoc should not be nil or have empty number")
	}
}

// TestMultiLanguageFunctional 验证多语言功能是否正常
func TestMultiLanguageFunctional(t *testing.T) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
	}

	for _, lang := range languages {
		t.Run(string(lang), func(t *testing.T) {
			faker := New(WithLanguage(lang))

			name := faker.Name()
			if name == "" {
				t.Errorf("Name should not be empty for language %s", lang)
			}

			company := faker.CompanyName()
			if company == "" {
				t.Errorf("CompanyName should not be empty for language %s", lang)
			}
		})
	}
}

// TestBatchGeneration 验证批量生成功能
func TestBatchGeneration(t *testing.T) {
	faker := New()

	// 测试批量生成
	names := faker.BatchNames(10)
	if len(names) != 10 {
		t.Errorf("Expected 10 names, got %d", len(names))
	}

	for i, name := range names {
		if name == "" {
			t.Errorf("Name at index %d is empty", i)
		}
	}

	// 测试并行生成
	emails := ParallelGenerate(100, func(f *Faker) string {
		return f.Email()
	})

	if len(emails) != 100 {
		t.Errorf("Expected 100 emails, got %d", len(emails))
	}

	for i, email := range emails {
		if email == "" || !strings.Contains(email, "@") {
			t.Errorf("Invalid email at index %d: %s", i, email)
		}
	}
}

// TestGlobalFunctionsSanity 验证全局函数的正常性
func TestGlobalFunctionsSanity(t *testing.T) {
	globalTests := []struct{
		name string
		value string
	}{
		{"Name", Name()},
		{"Email", Email()},
		{"PhoneNumber", PhoneNumber()},
		{"RandomUserAgent", RandomUserAgent()},
		{"CompanyName", CompanyName()},
		{"Word", Word()},
		{"Title", Title()},
	}

	for _, test := range globalTests {
		if test.value == "" {
			t.Errorf("Global %s returned empty string", test.name)
		}
	}
}

// TestNoMemoryLeaks 简单的内存泄漏测试
func TestNoMemoryLeaks(t *testing.T) {
	faker := New()

	// 生成大量数据，检查是否有明显的内存问题
	for i := 0; i < 10000; i++ {
		_ = faker.Name()
		_ = faker.Email()

		// 定期清理缓存
		if i%1000 == 0 {
			faker.ClearCache()
		}
	}

	// 如果到这里没有崩溃，说明基本没有严重的内存问题
	stats := faker.Stats()
	t.Logf("Stats after 10k generations: %+v", stats)
}

// TestRandomUserAgent 测试用户代理生成
func TestRandomUserAgent(t *testing.T) {
	// Test basic functionality
	ua := RandomUserAgent()

	if ua == "" {
		t.Error("RandomUserAgent() returned empty string")
	}

	// Test that returned user agent has basic structure
	if !strings.HasPrefix(ua, "Mozilla/") {
		t.Errorf("RandomUserAgent() returned invalid user agent (should start with Mozilla/): %q", ua)
	}

	// Should contain a browser name
	hasBrowser := strings.Contains(ua, "Chrome") || strings.Contains(ua, "Firefox") ||
		strings.Contains(ua, "Safari") || strings.Contains(ua, "Edge") || strings.Contains(ua, "Opera")
	if !hasBrowser {
		t.Errorf("RandomUserAgent() returned invalid user agent (missing browser name): %q", ua)
	}
}

func TestRandomUserAgentReturnsValidUserAgent(t *testing.T) {
	ua := RandomUserAgent()

	// All user agents should contain "Mozilla"
	if !strings.Contains(ua, "Mozilla") {
		t.Errorf("RandomUserAgent() returned invalid user agent (missing Mozilla): %q", ua)
	}

	// Should contain either a browser engine (AppleWebKit, Gecko) or browser name
	hasEngine := strings.Contains(ua, "AppleWebKit") || strings.Contains(ua, "Gecko")
	hasBrowser := strings.Contains(ua, "Chrome") || strings.Contains(ua, "Firefox") ||
		strings.Contains(ua, "Safari") || strings.Contains(ua, "Edge") || strings.Contains(ua, "Opera")

	if !hasEngine && !hasBrowser {
		t.Errorf("RandomUserAgent() returned invalid user agent (missing engine or browser): %q", ua)
	}
}

func TestRandomUserAgentConsistency(t *testing.T) {
	// Test that function doesn't panic and returns consistent format
	for i := 0; i < 100; i++ {
		ua := RandomUserAgent()

		if ua == "" {
			t.Fatalf("RandomUserAgent() returned empty string on iteration %d", i)
		}

		// Each user agent should be reasonably long
		if len(ua) < 30 {
			t.Errorf("RandomUserAgent() returned suspiciously short user agent: %q", ua)
		}

		// Should not contain line breaks or tabs
		if strings.Contains(ua, "\n") || strings.Contains(ua, "\t") || strings.Contains(ua, "\r") {
			t.Errorf("RandomUserAgent() returned user agent with invalid characters: %q", ua)
		}
	}
}

func TestRandomUserAgentDistribution(t *testing.T) {
	// Test that function returns different user agents over multiple calls
	// This is probabilistic, but with 255+ user agents, we should get variety

	results := make(map[string]int)
	iterations := 1000

	for i := 0; i < iterations; i++ {
		ua := RandomUserAgent()
		results[ua]++
	}

	// We should get at least 50 different user agents in 1000 calls
	// (this is very conservative given we have 255+ options)
	if len(results) < 50 {
		t.Errorf("RandomUserAgent() showed poor distribution: only %d unique user agents in %d calls", len(results), iterations)
	}

	// No single user agent should appear more than 5% of the time
	// (again, very conservative)
	maxAllowed := iterations / 20
	for ua, count := range results {
		if count > maxAllowed {
			t.Errorf("RandomUserAgent() returned %q too frequently: %d times out of %d calls", ua, count, iterations)
		}
	}
}

func TestUserAgentGeneration(t *testing.T) {
	// Test that generated user agents have proper structure
	for i := 0; i < 10; i++ {
		ua := RandomUserAgent()

		if ua == "" {
			t.Errorf("Generated user agent %d is empty", i)
		}

		if len(ua) < 30 {
			t.Errorf("Generated user agent %d is too short: %q", i, ua)
		}

		// Should start with Mozilla
		if !strings.HasPrefix(ua, "Mozilla/") {
			t.Errorf("Generated user agent %d doesn't start with Mozilla/: %q", i, ua)
		}

		// Should contain key browser components (at least one browser name)
		hasBrowser := strings.Contains(ua, "Chrome") || strings.Contains(ua, "Firefox") ||
			strings.Contains(ua, "Safari") || strings.Contains(ua, "Edge") || strings.Contains(ua, "Opera")
		if !hasBrowser {
			t.Errorf("Generated user agent %d missing browser name: %q", i, ua)
		}
	}
}

func TestUserAgentVariety(t *testing.T) {
	// Test that we generate different user agents
	agents := make(map[string]bool)

	for i := 0; i < 20; i++ {
		ua := RandomUserAgent()
		agents[ua] = true
	}

	// We should have generated multiple different user agents
	if len(agents) < 2 {
		t.Errorf("Generated user agents lack variety: only %d unique agents", len(agents))
	}
}

func TestRandomUserAgentBrowserTypes(t *testing.T) {
	// Test that we have different types of browsers in our list
	chromeCount := 0
	windowsCount := 0
	linuxCount := 0
	macCount := 0
	androidCount := 0

	// Sample a reasonable number to check distribution
	for i := 0; i < 100; i++ {
		ua := RandomUserAgent()

		if strings.Contains(ua, "Chrome") {
			chromeCount++
		}
		if strings.Contains(ua, "Windows") {
			windowsCount++
		}
		if strings.Contains(ua, "Linux") {
			linuxCount++
		}
		if strings.Contains(ua, "Macintosh") || strings.Contains(ua, "Mac OS X") {
			macCount++
		}
		if strings.Contains(ua, "Android") {
			androidCount++
		}
	}

	// We should have Chrome user agents (most of our list is Chrome)
	if chromeCount == 0 {
		t.Error("No Chrome user agents found in sample")
	}

	// We should have Windows user agents
	if windowsCount == 0 {
		t.Error("No Windows user agents found in sample")
	}

	// We should have some mobile (Android) user agents
	if androidCount == 0 {
		t.Error("No Android user agents found in sample")
	}
}

func TestRandomUserAgentNoEmptyOrNil(t *testing.T) {
	// Test edge cases to ensure function is robust
	for i := 0; i < 50; i++ {
		ua := RandomUserAgent()

		if ua == "" {
			t.Errorf("RandomUserAgent() returned empty string on call %d", i)
		}

		// Test for common invalid values
		invalidValues := []string{
			"<nil>",
			"null",
			"undefined",
			" ",
			"\t",
			"\n",
		}

		for _, invalid := range invalidValues {
			if ua == invalid {
				t.Errorf("RandomUserAgent() returned invalid value: %q", ua)
			}
		}
	}
}

// Test that the function works correctly with concurrent access
func TestRandomUserAgentConcurrency(t *testing.T) {
	results := make(chan string, 100)

	// Launch 10 goroutines
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				results <- RandomUserAgent()
			}
		}()
	}

	// Collect results
	for i := 0; i < 100; i++ {
		ua := <-results
		if ua == "" {
			t.Error("Concurrent access resulted in empty user agent")
		}

		// Verify it's a valid user agent structure
		if !strings.HasPrefix(ua, "Mozilla/") {
			t.Errorf("Concurrent access returned invalid user agent (should start with Mozilla/): %q", ua)
		}

		// Should contain a browser name
		hasBrowser := strings.Contains(ua, "Chrome") || strings.Contains(ua, "Firefox") ||
			strings.Contains(ua, "Safari") || strings.Contains(ua, "Edge") || strings.Contains(ua, "Opera")
		if !hasBrowser {
			t.Errorf("Concurrent access returned invalid user agent (missing browser name): %q", ua)
		}
	}
}

// Example_basicUsage 基本使用示例
func Example_basicUsage() {
	// 基本使用
	name := Name()
	email := Email()
	phone := PhoneNumber()

	// 验证输出不为空
	if name == "" || email == "" || phone == "" {
		fmt.Println("Error: Empty output")
	} else {
		fmt.Println("Basic usage example works correctly")
	}

	// Output:
	// Basic usage example works correctly
}

// Example_customFaker 自定义 Faker 示例
func Example_customFaker() {
	// 创建中文 Faker
	faker := New(
		WithLanguage(LanguageChineseSimplified),
		WithCountry(CountryChina),
		WithGender(GenderMale),
	)

	name := faker.Name()
	phone := faker.PhoneNumber()
	address := faker.FullAddress()

	// 验证输出不为空且符合预期格式
	if name == "" || phone == "" || address == nil || address.FullAddress == "" {
		fmt.Println("Error: Empty output")
	} else {
		fmt.Println("Custom faker example works correctly")
	}

	// Output:
	// Custom faker example works correctly
}

// =============================================================================
// Benchmark Tests
// =============================================================================

// BenchmarkRandomUserAgent 基准测试：RandomUserAgent函数
func BenchmarkRandomUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandomUserAgent()
	}
}

func BenchmarkRandomUserAgentAllocation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = RandomUserAgent()
	}
}

// BenchmarkNamePerformance 基准测试：名字生成性能
func BenchmarkNamePerformance(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = faker.Name()
	}
}

func BenchmarkEmail(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = faker.Email()
	}
}

func BenchmarkPhoneNumber(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = faker.PhoneNumber()
	}
}

func BenchmarkUserAgent(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = faker.UserAgent()
	}
}

func BenchmarkAddress(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = faker.FullAddress()
	}
}

func BenchmarkDeviceInfo(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = faker.DeviceInfo()
	}
}

// 基准测试 - 全局函数 vs 实例方法
func BenchmarkGlobalName(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = Name()
	}
}

func BenchmarkInstanceName(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = faker.Name()
	}
}

// 基准测试 - 批量生成对比
func BenchmarkBatchNamesComparison(b *testing.B) {
	faker := New()

	b.Run("Batch10", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = faker.BatchNames(10)
		}
	})

	b.Run("Batch100", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = faker.BatchNames(100)
		}
	})

	b.Run("Batch1000", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = faker.BatchNames(1000)
		}
	})
}

// 基准测试 - 并发性能
func BenchmarkConcurrentName(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = faker.Name()
		}
	})
}

func BenchmarkConcurrentEmail(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = faker.Email()
		}
	})
}

func BenchmarkConcurrentUserAgent(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = faker.UserAgent()
		}
	})
}

// 基准测试 - 内存分配
func BenchmarkMemoryAllocation(b *testing.B) {
	faker := New()

	b.Run("SingleInstance", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = faker.Name()
			_ = faker.Email()
			_ = faker.PhoneNumber()
		}
	})

	b.Run("MultipleInstances", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			f := New()
			_ = f.Name()
			_ = f.Email()
			_ = f.PhoneNumber()
		}
	})
}

// 基准测试 - 缓存效果
func BenchmarkWithCache(b *testing.B) {
	faker := New()

	// 预热缓存
	for i := 0; i < 100; i++ {
		_ = faker.Name()
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = faker.Name()
	}
}

func BenchmarkWithoutCache(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		faker := New()
		_ = faker.Name()
	}
}

// 基准测试 - 不同语言
func BenchmarkMultiLanguage(b *testing.B) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageFrench,
	}

	for _, lang := range languages {
		b.Run(string(lang), func(b *testing.B) {
			faker := New(WithLanguage(lang))
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				_ = faker.Name()
			}
		})
	}
}

// 压力测试 - 内存使用监控
func BenchmarkMemoryStress(b *testing.B) {
	faker := New()

	b.Run("LargeDataGeneration", func(b *testing.B) {
		var m1, m2 runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&m1)

		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			// 生成大量数据
			_ = faker.BatchNames(100)
			_ = faker.BatchEmails(100)
			_ = faker.BatchUserAgents(100)
		}

		b.StopTimer()
		runtime.GC()
		runtime.ReadMemStats(&m2)

		b.Logf("Memory allocated: %d bytes", m2.TotalAlloc-m1.TotalAlloc)
		b.Logf("Memory used: %d bytes", m2.Alloc-m1.Alloc)
		b.Logf("GC cycles: %d", m2.NumGC-m1.NumGC)
	})
}

// 基准测试 - CPU使用优化
func BenchmarkCPUEfficiency(b *testing.B) {
	faker := New()

	b.Run("StandardGeneration", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = faker.Name()
			_ = faker.Email()
			_ = faker.PhoneNumber()
			_ = faker.UserAgent()
		}
	})

	b.Run("BatchGeneration", func(b *testing.B) {
		b.ReportAllocs()
		batchSize := b.N / 4
		if batchSize <= 0 {
			batchSize = 1
		}

		for i := 0; i < 4; i++ {
			_ = faker.BatchNames(batchSize)
			_ = faker.BatchEmails(batchSize)
			_ = faker.BatchUserAgents(batchSize)
		}
	})
}

// 随机度测试
func BenchmarkRandomness(b *testing.B) {
	faker := New()

	b.Run("NameRandomness", func(b *testing.B) {
		results := make(map[string]int)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			name := faker.Name()
			results[name]++
		}

		b.StopTimer()
		uniqueCount := len(results)
		b.Logf("Unique names: %d out of %d (%.2f%%)",
			uniqueCount, b.N, float64(uniqueCount)/float64(b.N)*100)
	})

	b.Run("EmailRandomness", func(b *testing.B) {
		results := make(map[string]int)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			email := faker.Email()
			results[email]++
		}

		b.StopTimer()
		uniqueCount := len(results)
		b.Logf("Unique emails: %d out of %d (%.2f%%)",
			uniqueCount, b.N, float64(uniqueCount)/float64(b.N)*100)
	})
}

// 锁竞争测试
func BenchmarkLockContention(b *testing.B) {
	faker := New()

	b.Run("NoContention", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}
	})

	b.Run("HighContention", func(b *testing.B) {
		const numGoroutines = 100
		var wg sync.WaitGroup

		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < b.N/numGoroutines; j++ {
					_ = faker.Name()
				}
			}()
		}
		wg.Wait()
	})
}

// 基准测试 - 不同配置下的性能
func BenchmarkConfigurationImpact(b *testing.B) {
	b.Run("DefaultConfig", func(b *testing.B) {
		faker := New()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}
	})

	b.Run("WithSeed", func(b *testing.B) {
		faker := New(WithSeed(12345))
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}
	})

	b.Run("WithMultipleOptions", func(b *testing.B) {
		faker := New(
			WithLanguage(LanguageChineseSimplified),
			WithCountry(CountryChina),
			WithGender(GenderFemale),
			WithSeed(12345),
		)
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}
	})
}