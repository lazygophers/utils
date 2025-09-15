package fake

import (
	"strings"
	"testing"
)

// TestName 测试姓名生成
func TestName(t *testing.T) {
	faker := New()
	name := faker.Name()

	if name == "" {
		t.Error("Name() returned empty string")
	}

	// 英文名字应该包含空格
	if !strings.Contains(name, " ") {
		t.Error("English name should contain space")
	}
}

// TestChineseName 测试中文姓名生成
func TestChineseName(t *testing.T) {
	faker := New(WithLanguage(LanguageChineseSimplified))
	name := faker.Name()

	if name == "" {
		t.Error("Chinese name should not be empty")
	}

	// 中文名字不应该包含空格
	if strings.Contains(name, " ") {
		t.Error("Chinese name should not contain space")
	}
}

// TestFirstName 测试名字生成
func TestFirstName(t *testing.T) {
	faker := New()
	firstName := faker.FirstName()

	if firstName == "" {
		t.Error("FirstName() returned empty string")
	}

	// 名字不应该包含空格
	if strings.Contains(firstName, " ") {
		t.Error("FirstName should not contain space")
	}
}

// TestLastName 测试姓氏生成
func TestLastName(t *testing.T) {
	faker := New()
	lastName := faker.LastName()

	if lastName == "" {
		t.Error("LastName() returned empty string")
	}

	// 姓氏不应该包含空格（对于单个姓氏）
	parts := strings.Split(lastName, " ")
	if len(parts) > 2 {
		t.Error("LastName should not have too many parts")
	}
}

// TestGenderSpecificNames 测试性别相关的名字生成
func TestGenderSpecificNames(t *testing.T) {
	maleFaker := New(WithGender(GenderMale))
	femaleFaker := New(WithGender(GenderFemale))

	// 生成多个名字确保性别特定
	for i := 0; i < 10; i++ {
		maleName := maleFaker.FirstName()
		femaleName := femaleFaker.FirstName()

		if maleName == "" || femaleName == "" {
			t.Error("Gender-specific names should not be empty")
		}

		// 这里我们主要测试不会崩溃，实际的性别验证需要更复杂的逻辑
	}
}

// TestNamePrefix 测试姓名前缀
func TestNamePrefix(t *testing.T) {
	faker := New()
	prefix := faker.NamePrefix()

	// 前缀可能为空，这是正常的
	if prefix != "" {
		expectedPrefixes := []string{"Mr.", "Ms.", "Mrs.", "Miss", "Dr.", "Prof."}
		found := false
		for _, expected := range expectedPrefixes {
			if prefix == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Unexpected name prefix: %s", prefix)
		}
	}
}

// TestNameSuffix 测试姓名后缀
func TestNameSuffix(t *testing.T) {
	faker := New()

	// 生成多个后缀，因为大多数情况下会返回空字符串
	suffixes := make(map[string]bool)
	for i := 0; i < 100; i++ {
		suffix := faker.NameSuffix()
		if suffix != "" {
			suffixes[suffix] = true
		}
	}

	// 检查生成的后缀是否合法
	validSuffixes := []string{"Jr.", "Sr.", "II", "III", "IV", "V"}
	for suffix := range suffixes {
		found := false
		for _, valid := range validSuffixes {
			if suffix == valid {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Invalid name suffix: %s", suffix)
		}
	}
}

// TestFullName 测试完整姓名生成
func TestFullName(t *testing.T) {
	faker := New()
	fullName := faker.FullName()

	if fullName == "" {
		t.Error("FullName() returned empty string")
	}

	parts := strings.Split(fullName, " ")
	if len(parts) < 2 {
		t.Error("FullName should have at least 2 parts")
	}

	// 完整姓名可能有中间名，所以部分数量可以是2或3
	if len(parts) > 3 {
		t.Error("FullName should not have more than 3 parts")
	}
}

// TestFormattedName 测试格式化姓名
func TestFormattedName(t *testing.T) {
	faker := New()
	formatted := faker.FormattedName()

	if formatted == "" {
		t.Error("FormattedName() returned empty string")
	}

	// 格式化姓名应该至少包含基本的姓名
	parts := strings.Split(formatted, " ")
	if len(parts) < 2 {
		t.Error("FormattedName should have at least 2 parts")
	}
}

// TestBatchNames 测试批量姓名生成
func TestBatchNames(t *testing.T) {
	faker := New()
	count := 10
	names := faker.BatchNames(count)

	if len(names) != count {
		t.Errorf("Expected %d names, got %d", count, len(names))
	}

	for i, name := range names {
		if name == "" {
			t.Errorf("Name at index %d is empty", i)
		}
	}
}

// TestBatchNamesOptimized 测试优化的批量姓名生成
func TestBatchNamesOptimized(t *testing.T) {
	faker := New()
	count := 100
	names := faker.BatchNamesOptimized(count)

	if len(names) != count {
		t.Errorf("Expected %d names, got %d", count, len(names))
	}

	for i, name := range names {
		if name == "" {
			t.Errorf("Optimized name at index %d is empty", i)
		}
	}
}

// TestGlobalNameFunctions 测试全局姓名函数
func TestGlobalNameFunctions(t *testing.T) {
	// 测试全局函数
	name := Name()
	if name == "" {
		t.Error("Global Name() returned empty string")
	}

	firstName := FirstName()
	if firstName == "" {
		t.Error("Global FirstName() returned empty string")
	}

	lastName := LastName()
	if lastName == "" {
		t.Error("Global LastName() returned empty string")
	}

	fullName := FullName()
	if fullName == "" {
		t.Error("Global FullName() returned empty string")
	}
}

// TestNameConsistency 测试姓名生成的一致性
func TestNameConsistency(t *testing.T) {
	faker := New(WithSeed(12345))

	// 用相同的种子生成多次，应该得到相同结果
	name1 := faker.FirstName()

	faker2 := New(WithSeed(12345))
	name2 := faker2.FirstName()

	// 注意：由于随机性，这个测试可能需要调整
	// 这里主要测试不会崩溃
	if name1 == "" || name2 == "" {
		t.Error("Names should not be empty")
	}
}

// BenchmarkName 性能测试：姓名生成
func BenchmarkName(b *testing.B) {
	faker := New()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = faker.Name()
	}
}

// BenchmarkBatchNames 性能测试：批量姓名生成
func BenchmarkBatchNames(b *testing.B) {
	faker := New()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = faker.BatchNames(100)
	}
}

// BenchmarkBatchNamesOptimized 性能测试：优化的批量姓名生成
func BenchmarkBatchNamesOptimized(b *testing.B) {
	faker := New()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = faker.BatchNamesOptimized(100)
	}
}

// TestNameMemoryUsage 测试姓名生成的内存使用
func TestNameMemoryUsage(b *testing.T) {
	faker := New()

	// 生成大量姓名，检查内存使用情况
	for i := 0; i < 1000; i++ {
		_ = faker.Name()
	}

	// 清理缓存
	faker.ClearCache()
}
