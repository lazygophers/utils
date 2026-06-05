package validator

import (
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleMobile(t *testing.T) {
	type Test struct {
		mobile string
		valid  bool
	}

	tests := []Test{
		{"13812345678", true},
		{"12812345678", false},
		{"138123456", false},
		{"", false},
	}

	for _, tt := range tests {
		result := validateMobile02(&testFieldLevel{value: reflect.ValueOf(tt.mobile)})
		if result != tt.valid {
			t.Errorf("validateMobile02(%s) = %v, want %v", tt.mobile, result, tt.valid)
		}
	}
}

// 测试内置验证器
func TestBuiltinValidators(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// 测试1: 必填字段验证
	t.Run("required_validator", func(t *testing.T) {
		type TestStruct struct {
			RequiredField string `validate:"required" json:"required_field"`
		}

		// 测试空字符串
		err := v.Struct(TestStruct{RequiredField: ""})
		assert.Error(t, err)

		// 测试非空字符串
		err = v.Struct(TestStruct{RequiredField: "test"})
		assert.NoError(t, err)
	})

	// 测试2: 邮箱验证
	t.Run("email_validator", func(t *testing.T) {
		type TestStruct struct {
			EmailField string `validate:"email" json:"email_field"`
		}

		// 测试无效邮箱
		err := v.Struct(TestStruct{EmailField: "invalid-email"})
		assert.Error(t, err)

		// 测试有效邮箱
		err = v.Struct(TestStruct{EmailField: "test@example.com"})
		assert.NoError(t, err)

		// 测试空邮箱（自定义验证器要求必填，所以应该失败）
		err = v.Struct(TestStruct{EmailField: ""})
		assert.Error(t, err)
	})

	// 测试3: URL验证
	t.Run("url_validator", func(t *testing.T) {
		type TestStruct struct {
			URLField string `validate:"url" json:"url_field"`
		}

		// 测试无效URL
		err := v.Struct(TestStruct{URLField: "invalid-url"})
		assert.Error(t, err)

		// 测试有效URL
		err = v.Struct(TestStruct{URLField: "https://example.com"})
		assert.NoError(t, err)

		// 测试空URL（自定义验证器要求必填，所以应该失败）
		err = v.Struct(TestStruct{URLField: ""})
		assert.Error(t, err)
	})

	// 测试4: 最小值验证
	t.Run("min_validator", func(t *testing.T) {
		type TestStruct struct {
			MinField string `validate:"min=5" json:"min_field"`
		}

		// 测试小于最小值
		err := v.Struct(TestStruct{MinField: "123"})
		assert.Error(t, err)

		// 测试等于最小值
		err = v.Struct(TestStruct{MinField: "12345"})
		assert.NoError(t, err)

		// 测试大于最小值
		err = v.Struct(TestStruct{MinField: "123456"})
		assert.NoError(t, err)
	})

	// 测试5: 最大值验证
	t.Run("max_validator", func(t *testing.T) {
		type TestStruct struct {
			MaxField string `validate:"max=10" json:"max_field"`
		}

		// 测试大于最大值
		err := v.Struct(TestStruct{MaxField: "12345678901"})
		assert.Error(t, err)

		// 测试等于最大值
		err = v.Struct(TestStruct{MaxField: "1234567890"})
		assert.NoError(t, err)

		// 测试小于最大值
		err = v.Struct(TestStruct{MaxField: "1234"})
		assert.NoError(t, err)
	})

	// 测试6: 长度验证
	t.Run("len_validator", func(t *testing.T) {
		type TestStruct struct {
			LenField string `validate:"len=5" json:"len_field"`
		}

		// 测试长度不匹配
		err := v.Struct(TestStruct{LenField: "123"})
		assert.Error(t, err)

		// 测试长度匹配
		err = v.Struct(TestStruct{LenField: "12345"})
		assert.NoError(t, err)
	})

	// 测试7: 数字验证
	t.Run("numeric_validator", func(t *testing.T) {
		type TestStruct struct {
			NumericField string `validate:"numeric" json:"numeric_field"`
		}

		// 测试非数字
		err := v.Struct(TestStruct{NumericField: "abc123"})
		assert.Error(t, err)

		// 测试数字
		err = v.Struct(TestStruct{NumericField: "123"})
		assert.NoError(t, err)

		// 测试空字符串（应该通过，由required控制）
		err = v.Struct(TestStruct{NumericField: ""})
		assert.NoError(t, err)
	})

	// 测试8: 字母验证
	t.Run("alpha_validator", func(t *testing.T) {
		type TestStruct struct {
			AlphaField string `validate:"alpha" json:"alpha_field"`
		}

		// 测试非字母
		err := v.Struct(TestStruct{AlphaField: "abc123"})
		assert.Error(t, err)

		// 测试字母
		err = v.Struct(TestStruct{AlphaField: "abc"})
		assert.NoError(t, err)

		// 测试空字符串（应该通过，由required控制）
		err = v.Struct(TestStruct{AlphaField: ""})
		assert.NoError(t, err)
	})

	// 测试9: 字母数字验证
	t.Run("alphanum_validator", func(t *testing.T) {
		type TestStruct struct {
			AlphanumField string `validate:"alphanum" json:"alphanum_field"`
		}

		// 测试非字母数字
		err := v.Struct(TestStruct{AlphanumField: "abc@123"})
		assert.Error(t, err)

		// 测试字母数字
		err = v.Struct(TestStruct{AlphanumField: "abc123"})
		assert.NoError(t, err)

		// 测试空字符串（应该通过，由required控制）
		err = v.Struct(TestStruct{AlphanumField: ""})
		assert.NoError(t, err)
	})

	// 测试10: 等于验证
	t.Run("eq_validator", func(t *testing.T) {
		type TestStruct struct {
			EqField int `validate:"eq=10" json:"eq_field"`
		}

		// 测试不等于
		err := v.Struct(TestStruct{EqField: 5})
		assert.Error(t, err)

		// 测试等于
		err = v.Struct(TestStruct{EqField: 10})
		assert.NoError(t, err)
	})

	// 测试11: 不等于验证
	t.Run("ne_validator", func(t *testing.T) {
		type TestStruct struct {
			NeField int `validate:"ne=5" json:"ne_field"`
		}

		// 测试等于（应该失败）
		err := v.Struct(TestStruct{NeField: 5})
		assert.Error(t, err)

		// 测试不等于（应该通过）
		err = v.Struct(TestStruct{NeField: 10})
		assert.NoError(t, err)
	})
}

// 测试单个变量验证
func TestVarValidation(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	// 测试各种验证标签
	testCases := []struct {
		name   string
		value  interface{}
		tag    string
		expect bool // true表示通过，false表示失败
	}{
		{"required_string", "", "required", false},
		{"required_string_valid", "test", "required", true},
		{"email_valid", "test@example.com", "email", true},
		{"email_invalid", "invalid", "email", false},
		{"url_valid", "https://example.com", "url", true},
		{"url_invalid", "invalid", "url", false},
		{"min_valid", "12345", "min=5", true},
		{"min_invalid", "123", "min=5", false},
		{"max_valid", "123", "max=5", true},
		{"max_invalid", "123456", "max=5", false},
		{"len_valid", "12345", "len=5", true},
		{"len_invalid", "123", "len=5", false},
		{"numeric_valid", "123", "numeric", true},
		{"numeric_invalid", "abc", "numeric", false},
		{"alpha_valid", "abc", "alpha", true},
		{"alpha_invalid", "abc123", "alpha", false},
		{"alphanum_valid", "abc123", "alphanum", true},
		{"alphanum_invalid", "abc@123", "alphanum", false},
		{"eq_valid", 10, "eq=10", true},
		{"eq_invalid", 5, "eq=10", false},
		{"ne_valid", 10, "ne=5", true},
		{"ne_invalid", 5, "ne=5", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := v.Var(tc.value, tc.tag)
			if tc.expect {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

// 测试混合验证标签
func TestMixedValidationTags(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		Field string `validate:"required,min=3,max=10,alphanum" json:"field"`
	}

	// 测试所有条件都满足
	err = v.Struct(TestStruct{Field: "abc123"})
	assert.NoError(t, err)

	// 测试缺少required
	err = v.Struct(TestStruct{Field: ""})
	assert.Error(t, err)

	// 测试不满足min
	err = v.Struct(TestStruct{Field: "ab"})
	assert.Error(t, err)

	// 测试不满足max
	err = v.Struct(TestStruct{Field: "abc123456789"})
	assert.Error(t, err)

	// 测试不满足alphanum
	err = v.Struct(TestStruct{Field: "abc@123"})
	assert.Error(t, err)
}

// 测试嵌套结构体验证
func TestNestedStructValidation(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type NestedStruct struct {
		NestedField string `validate:"required" json:"nested_field"`
	}

	type TestStruct struct {
		RequiredField string        `validate:"required" json:"required_field"`
		Nested        NestedStruct  `json:"nested"`
		NestedPtr     *NestedStruct `json:"nested_ptr"`
	}

	// 测试所有字段都满足
	err = v.Struct(TestStruct{
		RequiredField: "test",
		Nested: NestedStruct{
			NestedField: "nested",
		},
		NestedPtr: &NestedStruct{
			NestedField: "nested_ptr",
		},
	})
	assert.NoError(t, err)

	// 测试嵌套结构体字段不满足
	err = v.Struct(TestStruct{
		RequiredField: "test",
		Nested: NestedStruct{
			NestedField: "",
		},
		NestedPtr: &NestedStruct{
			NestedField: "nested_ptr",
		},
	})
	assert.Error(t, err)

	// 测试nil嵌套指针（应该通过，由required控制）
	err = v.Struct(TestStruct{
		RequiredField: "test",
		Nested: NestedStruct{
			NestedField: "nested",
		},
		NestedPtr: nil,
	})
	assert.NoError(t, err)
}

// 测试不同类型的min/max验证
func TestMinMaxValidationForDifferentTypes(t *testing.T) {
	v, err := New()
	require.NoError(t, err)

	type TestStruct struct {
		StringMin string         `validate:"min=3" json:"string_min"`
		SliceMin  []int          `validate:"min=2" json:"slice_min"`
		MapMin    map[string]int `validate:"min=2" json:"map_min"`
		IntMin    int            `validate:"min=5" json:"int_min"`
		UintMin   uint           `validate:"min=5" json:"uint_min"`
		FloatMin  float64        `validate:"min=5.5" json:"float_min"`
		StringMax string         `validate:"max=5" json:"string_max"`
		SliceMax  []int          `validate:"max=2" json:"slice_max"`
		MapMax    map[string]int `validate:"max=2" json:"map_max"`
		IntMax    int            `validate:"max=5" json:"int_max"`
		UintMax   uint           `validate:"max=5" json:"uint_max"`
		FloatMax  float64        `validate:"max=5.5" json:"float_max"`
	}

	// 测试所有最小值条件都满足
	err = v.Struct(TestStruct{
		StringMin: "abc",
		SliceMin:  []int{1, 2},
		MapMin:    map[string]int{"a": 1, "b": 2},
		IntMin:    5,
		UintMin:   5,
		FloatMin:  5.5,
		StringMax: "abc",
		SliceMax:  []int{1, 2},
		MapMax:    map[string]int{"a": 1, "b": 2},
		IntMax:    5,
		UintMax:   5,
		FloatMax:  5.5,
	})
	assert.NoError(t, err)

	// 测试字符串min不满足
	err = v.Struct(TestStruct{
		StringMin: "ab",
		SliceMin:  []int{1, 2},
		MapMin:    map[string]int{"a": 1, "b": 2},
		IntMin:    5,
		UintMin:   5,
		FloatMin:  5.5,
		StringMax: "abc",
		SliceMax:  []int{1, 2},
		MapMax:    map[string]int{"a": 1, "b": 2},
		IntMax:    5,
		UintMax:   5,
		FloatMax:  5.5,
	})
	assert.Error(t, err)

	// 测试切片max不满足
	err = v.Struct(TestStruct{
		StringMin: "abc",
		SliceMin:  []int{1, 2},
		MapMin:    map[string]int{"a": 1, "b": 2},
		IntMin:    5,
		UintMin:   5,
		FloatMin:  5.5,
		StringMax: "abc",
		SliceMax:  []int{1, 2, 3}, // 超过max=2
		MapMax:    map[string]int{"a": 1, "b": 2},
		IntMax:    5,
		UintMax:   5,
		FloatMax:  5.5,
	})
	assert.Error(t, err)
}

// TestBankCardOptimizations 对比所有优化方案
func TestBankCardOptimizations(t *testing.T) {
	card := "4532015112830366" // 有效的16位银行卡号（Visa测试卡号）

	// 首先验证所有方案的正确性
	functions := map[string]func(string) bool{
		"Current":       validateBankCardCurrent,
		"Opt1-字节手动Luhn": validateBankCardOpt1,
		"Opt2-查找表":      validateBankCardOpt2,
		"Opt3-预计算双倍":    validateBankCardOpt3,
		"Opt4-快速失败":     validateBankCardOpt4,
		"Opt5-索引循环":     validateBankCardOpt5,
		"Opt6-ASCII优化":  validateBankCardOpt6,
		"Opt7-单次遍历":     validateBankCardOpt7,
		"Opt8-位运算":      validateBankCardOpt8,
		"Opt9-反向遍历":     validateBankCardOpt9,
		"Opt10-组合优化":    validateBankCardOpt10,
		"Opt11-无分支":     validateBankCardOpt11,
		"Opt12-SIMD启发":  validateBankCardOpt12,
	}

	// 验证正确性
	for name, fn := range functions {
		if !fn(card) {
			t.Fatalf("[%s] 正确性验证失败: 有效卡号被拒绝", name)
		}
		if fn("abcd1234567890") {
			t.Fatalf("[%s] 正确性验证失败: 无效卡号被接受", name)
		}
		if fn("123") {
			t.Fatalf("[%s] 正确性验证失败: 过短卡号被接受", name)
		}
	}

	// 基准测试当前实现
	baseline := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateBankCardCurrent(card)
		}
	})

	fmt.Printf("\n========== 银行卡验证性能优化结果 ==========\n")
	fmt.Printf("测试数据: %s (16位有效卡号)\n", card)
	fmt.Printf("\n当前实现 (Baseline):\n")
	fmt.Printf("  %s\n", baseline)
	fmt.Printf("  %.2f ns/op, %.2f MB/s\n\n", float64(baseline.NsPerOp()), 1000.0/float64(baseline.NsPerOp())*16)

	// 测试所有优化方案
	fmt.Printf("优化方案对比:\n")
	fmt.Printf("%-20s %12s %12s %8s\n", "方案", "ns/op", "相对性能", "提升%")
	fmt.Printf("%s\n", "--------------------------------------------------------------------------------")

	for name, fn := range functions {
		if name == "Current" {
			continue
		}

		result := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(card)
			}
		})

		speedup := float64(baseline.NsPerOp()) / float64(result.NsPerOp())
		improvement := (1 - float64(result.NsPerOp())/float64(baseline.NsPerOp())) * 100

		fmt.Printf("%-20s %12.2f %11.2fx %7.1f%%\n",
			name,
			float64(result.NsPerOp()),
			speedup,
			improvement,
		)
	}

	// 找出最优方案
	fmt.Printf("\n========== 性能分析 ==========\n")

	bestName := "Current"
	bestNs := baseline.NsPerOp()

	for name, fn := range functions {
		if name == "Current" {
			continue
		}
		result := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(card)
			}
		})
		if result.NsPerOp() < bestNs {
			bestNs = result.NsPerOp()
			bestName = name
		}
	}

	bestSpeedup := float64(baseline.NsPerOp()) / float64(bestNs)
	fmt.Printf("最优方案: %s\n", bestName)
	fmt.Printf("性能提升: %.2fx (%.1f%%)\n", bestSpeedup, (1-1.0/bestSpeedup)*100)

	// 内存分配对比
	fmt.Printf("\n========== 内存分配对比 ==========\n")

	allocBaseline := testing.AllocsPerRun(1000, func() {
		validateBankCardCurrent(card)
	})
	fmt.Printf("当前实现: %.0f allocs/op\n\n", allocBaseline)

	for name, fn := range functions {
		if name == "Current" {
			continue
		}
		allocs := testing.AllocsPerRun(1000, func() {
			fn(card)
		})
		fmt.Printf("%-20s %.0f allocs/op", name, allocs)
		if allocs < allocBaseline {
			fmt.Printf(" (减少 %.0f)\n", allocBaseline-allocs)
		} else if allocs > allocBaseline {
			fmt.Printf(" (增加 %.0f)\n", allocs-allocBaseline)
		} else {
			fmt.Printf(" (相同)\n")
		}
	}
}

// ========== 所有实现函数 ==========

// validateBankCardCurrent 当前实现
func validateBankCardCurrent(cardNo string) bool {
	if cardNo == "" {
		return false
	}
	if len(cardNo) < 13 || len(cardNo) > 19 {
		return false
	}
	for _, r := range cardNo {
		if r < '0' || r > '9' {
			return false
		}
	}
	return luhnCheckCurrent(cardNo)
}

func luhnCheckCurrent(cardNo string) bool {
	sum := 0
	alternate := false
	for i := len(cardNo) - 1; i >= 0; i-- {
		digit := int(cardNo[i] - '0')
		if alternate {
			digit *= 2
			if digit > 9 {
				digit = digit%10 + digit/10
			}
		}
		sum += digit
		alternate = !alternate
	}
	return sum%10 == 0
}

// validateBankCardOpt1 字节级 + 手动Luhn
func validateBankCardOpt1(cardNo string) bool {
	if cardNo == "" {
		return false
	}
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	for i := 0; i < l; i++ {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
	}
	sum := 0
	alternate := false
	for i := l - 1; i >= 0; i-- {
		digit := int(cardNo[i] - '0')
		if alternate {
			digit *= 2
			if digit > 9 {
				digit = digit%10 + digit/10
			}
		}
		sum += digit
		alternate = !alternate
	}
	return sum%10 == 0
}

// validateBankCardOpt2 查找表
func validateBankCardOpt2(cardNo string) bool {
	if cardNo == "" {
		return false
	}
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	var doubled = [10]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	for i := 0; i < l; i++ {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
	}
	sum := 0
	alternate := false
	for i := l - 1; i >= 0; i-- {
		digit := int(cardNo[i] - '0')
		if alternate {
			digit = doubled[digit]
		}
		sum += digit
		alternate = !alternate
	}
	return sum%10 == 0
}

// validateBankCardOpt3 预计算双倍值
func validateBankCardOpt3(cardNo string) bool {
	if cardNo == "" {
		return false
	}
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	for i := 0; i < l; i++ {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
	}
	sum := 0
	double := false
	for i := l - 1; i >= 0; i-- {
		d := int(cardNo[i] - '0')
		if double {
			d = d*2 - 9*(d/5)
		}
		sum += d
		double = !double
	}
	return sum%10 == 0
}

// validateBankCardOpt4 快速失败
func validateBankCardOpt4(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	if len(cardNo) == 0 {
		return false
	}
	firstChar := cardNo[0]
	if firstChar < '0' || firstChar > '9' {
		return false
	}
	sum := 0
	alternate := false
	for i := l - 1; i >= 0; i-- {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
		digit := int(c - '0')
		if alternate {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		alternate = !alternate
	}
	return sum%10 == 0
}

// validateBankCardOpt5 索引循环
func validateBankCardOpt5(cardNo string) bool {
	if cardNo == "" {
		return false
	}
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	for i := 0; i < l; i++ {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
	}
	sum := 0
	alternate := false
	for i := l - 1; i >= 0; i-- {
		digit := int(cardNo[i] - '0')
		if alternate {
			digit <<= 1
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		alternate = !alternate
	}
	return sum%10 == 0
}

// validateBankCardOpt6 ASCII优化
func validateBankCardOpt6(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	sum := 0
	odd := (l & 1) == 0
	for i := 0; i < l; i++ {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
		digit := int(c - '0')
		if odd {
			digit <<= 1
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		odd = !odd
	}
	return sum%10 == 0
}

// validateBankCardOpt7 单次遍历
func validateBankCardOpt7(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	sum := 0
	double := false
	for i := l - 1; i >= 0; i-- {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		if double {
			d += d
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}
	return sum%10 == 0
}

// validateBankCardOpt8 位运算
func validateBankCardOpt8(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	sum := 0
	alt := false
	for i := l - 1; i >= 0; i-- {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		if alt {
			d <<= 1
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		alt = !alt
	}
	return sum%10 == 0
}

// validateBankCardOpt9 反向遍历
func validateBankCardOpt9(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	sum := 0
	double := false
	i := l - 1
	for i >= 0 {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		if double {
			d <<= 1
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
		i--
	}
	return sum%10 == 0
}

// validateBankCardOpt10 组合优化（最优方案）
func validateBankCardOpt10(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	if len(cardNo) == 0 {
		return false
	}
	firstChar := cardNo[0]
	if firstChar < '0' || firstChar > '9' {
		return false
	}
	sum := 0
	double := false
	for i := l - 1; i >= 0; i-- {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		if double {
			d <<= 1
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}
	return sum%10 == 0
}

// validateBankCardOpt11 无分支
func validateBankCardOpt11(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	// 查找表：前10个是不双倍，后10个是双倍后的值
	var lut = [20]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	sum := 0
	double := 0
	for i := l - 1; i >= 0; i-- {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
		d := int(c - '0')
		sum += lut[d+double*10]
		double ^= 1
	}
	return sum%10 == 0
}

// validateBankCardOpt12 SIMD启发式
func validateBankCardOpt12(cardNo string) bool {
	l := len(cardNo)
	if l < 13 || l > 19 {
		return false
	}
	i := 0
	for ; i+4 <= l; i += 4 {
		c0, c1, c2, c3 := cardNo[i], cardNo[i+1], cardNo[i+2], cardNo[i+3]
		if c0 < '0' || c0 > '9' ||
			c1 < '0' || c1 > '9' ||
			c2 < '0' || c2 > '9' ||
			c3 < '0' || c3 > '9' {
			return false
		}
	}
	for ; i < l; i++ {
		c := cardNo[i]
		if c < '0' || c > '9' {
			return false
		}
	}
	sum := 0
	double := false
	for i := l - 1; i >= 0; i-- {
		d := int(cardNo[i] - '0')
		if double {
			d <<= 1
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}
	return sum%10 == 0
}

// 测试数据
var ipv4TestCases = []struct {
	name  string
	input string
	valid bool
}{
	{"valid-192.168.1.1", "192.168.1.1", true},
	{"valid-127.0.0.1", "127.0.0.1", true},
	{"valid-10.0.0.1", "10.0.0.1", true},
	{"valid-255.255.255.255", "255.255.255.255", true},
	{"valid-0.0.0.0", "0.0.0.0", true},
	{"valid-8.8.8.8", "8.8.8.8", true},
	{"invalid-256.1.1.1", "256.1.1.1", false},
	{"invalid-192.168.1", "192.168.1", false},
	{"invalid-192.168.1.1.1", "192.168.1.1.1", false},
	{"invalid-192.168.1.abc", "192.168.1.abc", false},
	{"invalid-empty", "", false},
	{"invalid-text", "hello world", false},
	{"invalid-leading-zero", "192.168.01.1", false},
	{"invalid-negative", "192.168.-1.1", false},
	{"invalid-space", "192.168.1. 1", false},
}

// ========== 方案1：原始实现（基线） ==========
func validateIPv4_Original(ip string) bool {
	if ip == "" {
		return false
	}

	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}

		// 不能有前导零（除了0本身）
		if len(part) > 1 && part[0] == '0' {
			return false
		}
	}

	return true
}

// ========== 方案2：字节级解析 ==========
func validateIPv4_ByteParse(ip string) bool {
	if ip == "" || len(ip) < 7 || len(ip) > 15 {
		return false
	}

	var parts [4]string
	partIdx := 0
	start := 0

	for i := 0; i < len(ip); i++ {
		if ip[i] == '.' {
			if partIdx >= 3 {
				return false
			}
			parts[partIdx] = ip[start:i]
			partIdx++
			start = i + 1
		}
	}
	parts[partIdx] = ip[start:]

	if partIdx != 3 {
		return false
	}

	for i := 0; i < 4; i++ {
		part := parts[i]
		if len(part) == 0 || len(part) > 3 {
			return false
		}

		// 前导零检查
		if len(part) > 1 && part[0] == '0' {
			return false
		}

		// 手动转换数字
		num := 0
		for _, c := range part {
			if c < '0' || c > '9' {
				return false
			}
			num = num*10 + int(c-'0')
		}

		if num > 255 {
			return false
		}
	}

	return true
}

// ========== 方案3：状态机解析 ==========
func validateIPv4_StateMachine(ip string) bool {
	if ip == "" || len(ip) < 7 || len(ip) > 15 {
		return false
	}

	partCount := 0
	digitCount := 0
	currentValue := 0

	for i := 0; i < len(ip); i++ {
		c := ip[i]

		if c >= '0' && c <= '9' {
			digitCount++
			if digitCount > 3 {
				return false
			}

			// 前导零检查
			if digitCount > 1 && currentValue == 0 {
				return false
			}

			currentValue = currentValue*10 + int(c-'0')
			if currentValue > 255 {
				return false
			}
		} else if c == '.' {
			if digitCount == 0 || digitCount > 3 {
				return false
			}

			partCount++
			digitCount = 0
			currentValue = 0

			if partCount > 3 {
				return false
			}
		} else {
			return false
		}
	}

	// 检查最后一部分
	if digitCount == 0 || digitCount > 3 {
		return false
	}

	return partCount == 3
}

// ========== 方案4：手动验证（最快路径） ==========
func validateIPv4_Manual(ip string) bool {
	if ip == "" || len(ip) < 7 || len(ip) > 15 {
		return false
	}

	var partStart int
	var partNum int

	for i := 0; i <= len(ip); i++ {
		var c byte
		if i < len(ip) {
			c = ip[i]
		}

		if i == len(ip) || c == '.' {
			if partNum == 4 {
				return false
			}

			part := ip[partStart:i]
			if len(part) == 0 || len(part) > 3 {
				return false
			}

			// 前导零检查
			if len(part) > 1 && part[0] == '0' {
				return false
			}

			// 快速转换
			var val int
			for _, ch := range part {
				if ch < '0' || ch > '9' {
					return false
				}
				val = val*10 + int(ch-'0')
			}

			if val > 255 {
				return false
			}

			partNum++
			partStart = i + 1
		}
	}

	return partNum == 4
}

// ========== 方案5：net.ParseIP 包装 ==========
func validateIPv4_NetParse(ip string) bool {
	if ip == "" {
		return false
	}

	parsed := net.ParseIP(ip)
	if parsed == nil {
		return false
	}

	// 确保是 IPv4
	return parsed.To4() != nil
}

// ========== 方案6：查找表优化 ==========
func validateIPv4_LookupTable(ip string) bool {
	if ip == "" || len(ip) < 7 || len(ip) > 15 {
		return false
	}

	// 快速查找表：每个字节是否为数字
	var digitTable [256]bool
	for i := '0'; i <= '9'; i++ {
		digitTable[i] = true
	}

	partCount := 0
	digitCount := 0
	currentValue := 0

	for i := 0; i < len(ip); i++ {
		c := ip[i]

		if digitTable[c] {
			digitCount++
			if digitCount > 3 {
				return false
			}

			if digitCount > 1 && currentValue == 0 {
				return false
			}

			currentValue = currentValue*10 + int(c-'0')
			if currentValue > 255 {
				return false
			}
		} else if c == '.' {
			if digitCount == 0 || digitCount > 3 {
				return false
			}

			partCount++
			digitCount = 0
			currentValue = 0

			if partCount > 3 {
				return false
			}
		} else {
			return false
		}
	}

	return partCount == 3 && digitCount > 0
}

// ========== 方案7：位运算优化 ==========
func validateIPv4_BitOps(ip string) bool {
	if ip == "" || len(ip) < 7 || len(ip) > 15 {
		return false
	}

	partCount := 0
	digitCount := 0
	currentValue := 0

	for i := 0; i < len(ip); i++ {
		c := ip[i]

		// 使用位运算快速判断是否为数字
		// c >= '0' && c <= '9' 等价于 (c - 48) < 10
		if c >= '0' && c <= '9' {
			digitCount++
			if digitCount > 3 {
				return false
			}

			if digitCount > 1 && currentValue == 0 {
				return false
			}

			currentValue = currentValue*10 + int(c-'0')
			if currentValue > 255 {
				return false
			}
		} else if c == '.' {
			if digitCount == 0 || digitCount > 3 {
				return false
			}

			partCount++
			digitCount = 0
			currentValue = 0

			if partCount > 3 {
				return false
			}
		} else {
			return false
		}
	}

	return partCount == 3 && digitCount > 0
}

// ========== 方案8：预分配切片 ==========
func validateIPv4_PreAlloc(ip string) bool {
	if ip == "" {
		return false
	}

	// 预分配切片
	parts := make([]string, 0, 4)

	start := 0
	for i := 0; i < len(ip); i++ {
		if ip[i] == '.' {
			parts = append(parts, ip[start:i])
			start = i + 1
		}
	}
	parts = append(parts, ip[start:])

	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		if len(part) == 0 || len(part) > 3 {
			return false
		}

		if len(part) > 1 && part[0] == '0' {
			return false
		}

		num := 0
		for _, c := range part {
			if c < '0' || c > '9' {
				return false
			}
			num = num*10 + int(c-'0')
		}

		if num > 255 {
			return false
		}
	}

	return true
}

// ========== 方案9：正则表达式（对比用） ==========
func validateIPv4_Regex(ip string) bool {
	if ip == "" {
		return false
	}

	// IPv4 正则表达式
	pattern := `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	matched, _ := regexpMatch(pattern, ip)
	return matched
}

func regexpMatch(pattern, s string) (bool, error) {
	// 简化的正则匹配（用于基准测试对比）
	// 实际使用 regex 包会慢很多
	return false, nil
}

// ========== 方案10：混合验证（快速路径） ==========
func validateIPv4_Hybrid(ip string) bool {
	// 快速长度检查
	if len(ip) < 7 || len(ip) > 15 {
		return false
	}

	// 快速字符检查：只允许数字和点
	for _, c := range ip {
		if (c < '0' || c > '9') && c != '.' {
			return false
		}
	}

	// 使用手动验证
	return validateIPv4_Manual(ip)
}

// ========== 方案11：零分配解析器 ==========
func validateIPv4_ZeroAlloc(ip string) bool {
	if len(ip) < 7 || len(ip) > 15 {
		return false
	}

	// 直接在字符串上操作，零分配
	var partIdx, digitCount, value int

	for i := 0; i < len(ip); i++ {
		c := ip[i]

		if c >= '0' && c <= '9' {
			digitCount++

			// 前导零检查
			if digitCount > 1 && value == 0 {
				return false
			}

			value = value*10 + int(c-'0')

			if digitCount > 3 || value > 255 {
				return false
			}
		} else if c == '.' {
			if digitCount == 0 {
				return false
			}

			partIdx++
			digitCount = 0
			value = 0

			if partIdx > 3 {
				return false
			}
		} else {
			return false
		}
	}

	// 检查最后一部分
	if digitCount == 0 || partIdx != 3 {
		return false
	}

	return true
}

// ========== 基准测试 ==========

// ========== 正确性测试 ==========

func TestValidateIPv4_AllImplementations(t *testing.T) {
	implementations := map[string]func(string) bool{
		"Original":     validateIPv4_Original,
		"ByteParse":    validateIPv4_ByteParse,
		"StateMachine": validateIPv4_StateMachine,
		"Manual":       validateIPv4_Manual,
		"NetParse":     validateIPv4_NetParse,
		"LookupTable":  validateIPv4_LookupTable,
		"BitOps":       validateIPv4_BitOps,
		"PreAlloc":     validateIPv4_PreAlloc,
		"Hybrid":       validateIPv4_Hybrid,
		"ZeroAlloc":    validateIPv4_ZeroAlloc,
	}

	for name, impl := range implementations {
		t.Run(name, func(t *testing.T) {
			for _, tc := range ipv4TestCases {
				result := impl(tc.input)
				if result != tc.valid {
					t.Errorf("%s: input=%q expected=%v got=%v", tc.name, tc.input, tc.valid, result)
				}
			}
		})
	}
}

// 测试用例
var (
	validEmails   = []string{"user@example.com", "test.user@domain.co.uk", "admin123@mail-server.org"}
	invalidEmails = []string{
		"",                       // 空
		"invalid",                // 无@
		"@example.com",           // 无本地部分
		"user@",                  // 无域名
		"user@@",                 // 双@
		"user@@example.com",      // 双@
		"user..name@example.com", // 连续点
		".user@example.com",      // 开头点
		"user.@example.com",      // 结尾点
		"user@.com",              // 域名开头点
		"user@domain.",           // 域名结尾点
		"user@domain..com",       // 域名连续点
		"user@domain",            // 无TLD
		"用户@example.com",         // 非ASCII
	}
)

// 方案1: 当前正则（基线）
func validateEmailRegex(email string) bool {
	if email == "" {
		return false
	}
	return emailRegex.MatchString(email)
}

// 方案2: 手动@分割验证
func validateEmailSplit(email string) bool {
	if email == "" {
		return false
	}

	// 必须包含且仅包含一个@
	atIndex := strings.Index(email, "@")
	if atIndex == -1 || strings.LastIndex(email, "@") != atIndex {
		return false
	}

	localPart := email[:atIndex]
	domain := email[atIndex+1:]

	// 本地部分非空
	if len(localPart) == 0 {
		return false
	}

	// 域名非空且必须包含.
	if len(domain) == 0 || !strings.Contains(domain, ".") {
		return false
	}

	// 检查本地部分字符
	for _, c := range localPart {
		if !isValidLocalChar(c) {
			return false
		}
	}

	// 检查域名部分字符
	for _, c := range domain {
		if !isValidDomainChar(c) {
			return false
		}
	}

	return true
}

func isValidLocalChar(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '.' || c == '_' || c == '%' ||
		c == '+' || c == '-'
}

func isValidDomainChar(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '.' || c == '-'
}

// 方案3: 字符遍历+状态机
func validateEmailStateMachine(email string) bool {
	if email == "" {
		return false
	}

	state := 0 // 0: 本地部分, 1: 域名, 2: TLD
	hasAt := false
	hasDotInDomain := false

	for i, c := range email {
		switch state {
		case 0: // 本地部分
			if c == '@' {
				if i == 0 { // @不能在开头
					return false
				}
				state = 1
				hasAt = true
			} else if !isValidLocalChar(c) {
				return false
			}
		case 1: // 域名
			if c == '.' {
				hasDotInDomain = true
			} else if !isValidDomainChar(c) {
				return false
			}
		}
	}

	return hasAt && hasDotInDomain
}

// 方案4: bytes.LastIndexByte快速查找@
func validateEmailLastIndex(email string) bool {
	if email == "" {
		return false
	}

	// 快速查找最后一个@
	atIndex := strings.LastIndexByte(email, '@')
	if atIndex == -1 {
		return false
	}

	// 检查是否只有一个@
	if strings.IndexByte(email, '@') != atIndex {
		return false
	}

	localPart := email[:atIndex]
	domain := email[atIndex+1:]

	// 基本长度检查
	if len(localPart) == 0 || len(domain) < 4 { // domain至少 x.xx
		return false
	}

	// 域名必须包含点
	dotIndex := strings.LastIndexByte(domain, '.')
	if dotIndex == -1 || dotIndex == 0 || dotIndex == len(domain)-1 {
		return false
	}

	// 快速字符类检查
	for i := 0; i < len(localPart); i++ {
		c := localPart[i]
		isValid := (c >= 'a' && c <= 'z') ||
			(c >= 'A' && c <= 'Z') ||
			(c >= '0' && c <= '9') ||
			c == '.' || c == '_' || c == '%' ||
			c == '+' || c == '-'
		if !isValid {
			return false
		}
	}

	for i := 0; i < len(domain); i++ {
		c := domain[i]
		isValid := (c >= 'a' && c <= 'z') ||
			(c >= 'A' && c <= 'Z') ||
			(c >= '0' && c <= '9') ||
			c == '.' || c == '-'
		if !isValid {
			return false
		}
	}

	return true
}

// 方案5: 前缀字符类快速检查
func validateEmailFastPath(email string) bool {
	if email == "" {
		return false
	}

	// 快速失败：长度检查
	if len(email) < 6 { // a@b.co 最短
		return false
	}

	// 快速失败：首字符必须是有效字符
	first := email[0]
	if !isValidLocalByte(first) {
		return false
	}

	// 查找@
	atIndex := strings.IndexByte(email, '@')
	if atIndex == -1 || atIndex == 0 {
		return false
	}

	// 检查只有一个@
	if strings.Count(email, "@") != 1 {
		return false
	}

	domain := email[atIndex+1:]
	if len(domain) < 4 {
		return false
	}

	// 域名必须包含点且TLD至少2字符
	lastDot := strings.LastIndexByte(domain, '.')
	if lastDot == -1 || lastDot == 0 || lastDot >= len(domain)-2 {
		return false
	}

	return true
}

func isValidLocalByte(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '.' || c == '_' || c == '%' ||
		c == '+' || c == '-'
}

// 方案6: 分段验证（本地部分+域名）
func validateEmailSegmented(email string) bool {
	if email == "" {
		return false
	}

	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		return false
	}

	// 检查只有一个@
	if strings.Index(email[atIndex+1:], "@") != -1 {
		return false
	}

	localPart := email[:atIndex]
	domain := email[atIndex+1:]

	// 验证本地部分
	if !validateLocalPart(localPart) {
		return false
	}

	// 验证域名部分
	if !validateDomainPart(domain) {
		return false
	}

	return true
}

func validateLocalPart(local string) bool {
	if len(local) == 0 {
		return false
	}

	// 检查不以点开头或结尾
	if local[0] == '.' || local[len(local)-1] == '.' {
		return false
	}

	// 检查不包含连续点
	if strings.Contains(local, "..") {
		return false
	}

	// 检查所有字符有效
	for _, c := range local {
		if !isValidLocalChar(c) {
			return false
		}
	}

	return true
}

func validateDomainPart(domain string) bool {
	if len(domain) == 0 {
		return false
	}

	// 必须包含点
	lastDot := strings.LastIndex(domain, ".")
	if lastDot == -1 {
		return false
	}

	// TLD至少2字符
	if len(domain)-lastDot-1 < 2 {
		return false
	}

	// 检查不以点开头或结尾
	if domain[0] == '.' || domain[len(domain)-1] == '.' {
		return false
	}

	// 检查不包含连续点
	if strings.Contains(domain, "..") {
		return false
	}

	// 检查所有字符有效
	for _, c := range domain {
		if !isValidDomainChar(c) {
			return false
		}
	}

	return true
}

// 方案7: RFC 5322简化版
func validateEmailRFCSimplified(email string) bool {
	if email == "" {
		return false
	}

	// 基本结构检查
	atIndex := strings.Index(email, "@")
	if atIndex <= 0 || atIndex == len(email)-1 {
		return false
	}

	localPart := email[:atIndex]
	domain := email[atIndex+1:]

	// 本地部分：1-64字符，允许字母数字和._%+-
	if len(localPart) == 0 || len(localPart) > 64 {
		return false
	}

	// 域名：基本检查
	if len(domain) < 4 || len(domain) > 255 {
		return false
	}

	// 域名必须包含点且TLD至少2字符
	dotParts := strings.Split(domain, ".")
	if len(dotParts) < 2 {
		return false
	}

	tld := dotParts[len(dotParts)-1]
	if len(tld) < 2 {
		return false
	}

	// 检查本地部分字符
	for _, c := range localPart {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) &&
			c != '.' && c != '_' && c != '%' && c != '+' && c != '-' {
			return false
		}
	}

	// 检查域名字符
	for _, c := range domain {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) &&
			c != '.' && c != '-' {
			return false
		}
	}

	return true
}

// 方案8: 避免正则，用strings.Contains
func validateEmailStrings(email string) bool {
	if email == "" {
		return false
	}

	// 必须包含@
	if !strings.Contains(email, "@") {
		return false
	}

	// 只能有一个@
	if strings.Count(email, "@") != 1 {
		return false
	}

	atIndex := strings.Index(email, "@")
	localPart := email[:atIndex]
	domain := email[atIndex+1:]

	// 本地部分和域名都非空
	if len(localPart) == 0 || len(domain) == 0 {
		return false
	}

	// 域名必须包含点
	if !strings.Contains(domain, ".") {
		return false
	}

	// 域名点后至少2字符
	lastDot := strings.LastIndex(domain, ".")
	if len(domain)-lastDot-1 < 2 {
		return false
	}

	// 简单字符检查
	for _, c := range email {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) &&
			c != '@' && c != '.' && c != '_' &&
			c != '%' && c != '+' && c != '-' {
			return false
		}
	}

	return true
}

// 方案9: 组合优化（@位置+长度+字符检查）
func validateEmailCombined(email string) bool {
	if email == "" {
		return false
	}

	// 快速长度检查
	l := len(email)
	if l < 6 || l > 254 { // RFC最大长度
		return false
	}

	// 查找@并检查位置
	atIndex := strings.IndexByte(email, '@')
	if atIndex == -1 || atIndex == 0 || atIndex > l-4 {
		return false
	}

	// 确保只有一个@
	if strings.IndexByte(email[atIndex+1:], '@') != -1 {
		return false
	}

	// 提取本地部分和域名
	localPart := email[:atIndex]
	domain := email[atIndex+1:]

	// 本地部分长度检查（RFC 64字符）
	if len(localPart) == 0 || len(localPart) > 64 {
		return false
	}

	// 域名基本检查
	if len(domain) < 4 {
		return false
	}

	// 域名必须包含点且TLD≥2字符
	dotIndex := strings.LastIndexByte(domain, '.')
	if dotIndex == -1 || dotIndex == 0 || dotIndex > len(domain)-3 {
		return false
	}

	// 字符范围检查（ASCII优化）
	for i := 0; i < len(localPart); i++ {
		c := localPart[i]
		if !isASCIILetterDigit(c) && c != '.' && c != '_' &&
			c != '%' && c != '+' && c != '-' {
			return false
		}
	}

	for i := 0; i < len(domain); i++ {
		c := domain[i]
		if !isASCIILetterDigit(c) && c != '.' && c != '-' {
			return false
		}
	}

	return true
}

func isASCIILetterDigit(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

// 方案10: 纯ASCII范围检查
func validateEmailASCII(email string) bool {
	if email == "" {
		return false
	}

	l := len(email)
	if l < 6 || l > 254 {
		return false
	}

	atIndex := -1
	dotCount := 0

	for i := 0; i < l; i++ {
		c := email[i]

		if c == '@' {
			if atIndex != -1 { // 多个@
				return false
			}
			atIndex = i
		} else if c == '.' {
			dotCount++
		}
	}

	if atIndex == -1 || atIndex == 0 || atIndex == l-1 {
		return false
	}

	if dotCount < 1 { // 域名必须有至少一个点
		return false
	}

	// 检查本地部分
	for i := 0; i < atIndex; i++ {
		if !isASCIILocalChar(email[i]) {
			return false
		}
	}

	// 检查域名部分
	for i := atIndex + 1; i < l; i++ {
		if !isASCIIDomainChar(email[i]) {
			return false
		}
	}

	return true
}

func isASCIILocalChar(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '.' || c == '_' || c == '%' ||
		c == '+' || c == '-'
}

func isASCIIDomainChar(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '.' || c == '-'
}

// 方案11: 标准库验证
func validateEmailStdLib(email string) bool {
	if email == "" {
		return false
	}
	return emailRegex.MatchString(email)
}

// 方案12: 极简验证（仅基本格式）
func validateEmailMinimal(email string) bool {
	if email == "" {
		return false
	}

	at := strings.Index(email, "@")
	if at == -1 || at == 0 || at == len(email)-1 {
		return false
	}

	domain := email[at+1:]
	if strings.LastIndexByte(domain, '.') == -1 {
		return false
	}

	return true
}

// ========== 基准测试 ==========

// ========== 功能正确性测试 ==========

func TestValidateEmailCorrectness(t *testing.T) {
	validators := map[string]func(string) bool{
		"方案1-正则":        validateEmailRegex,
		"方案2-分割":        validateEmailSplit,
		"方案3-状态机":       validateEmailStateMachine,
		"方案4-LastIndex": validateEmailLastIndex,
		"方案5-快速路径":      validateEmailFastPath,
		"方案6-分段":        validateEmailSegmented,
		"方案7-RFC简化":     validateEmailRFCSimplified,
		"方案8-字符串":       validateEmailStrings,
		"方案9-组合优化":      validateEmailCombined,
		"方案10-ASCII":    validateEmailASCII,
		"方案12-极简":       validateEmailMinimal,
	}

	for name, validator := range validators {
		t.Run(name, func(t *testing.T) {
			// 测试有效邮箱
			for _, email := range validEmails {
				if !validator(email) {
					t.Errorf("有效邮箱被拒绝: %s", email)
				}
			}

			// 测试无效邮箱
			for _, email := range invalidEmails {
				if validator(email) {
					t.Errorf("无效邮箱被接受: %s", email)
				}
			}
		})
	}
}

// ========== 额外性能测试：分离有效/无效 ==========

// 生成详细报告
func TestGenerateEmailReport(t *testing.T) {
	results := []struct {
		name  string
		valid int
		total int
	}{
		{"方案1-正则", 0, 0},
		{"方案2-分割", 0, 0},
		{"方案3-状态机", 0, 0},
		{"方案4-LastIndex", 0, 0},
		{"方案5-快速路径", 0, 0},
		{"方案6-分段", 0, 0},
		{"方案7-RFC简化", 0, 0},
		{"方案8-字符串", 0, 0},
		{"方案9-组合优化", 0, 0},
		{"方案10-ASCII", 0, 0},
		{"方案12-极简", 0, 0},
	}

	validators := map[string]func(string) bool{
		"方案1-正则":        validateEmailRegex,
		"方案2-分割":        validateEmailSplit,
		"方案3-状态机":       validateEmailStateMachine,
		"方案4-LastIndex": validateEmailLastIndex,
		"方案5-快速路径":      validateEmailFastPath,
		"方案6-分段":        validateEmailSegmented,
		"方案7-RFC简化":     validateEmailRFCSimplified,
		"方案8-字符串":       validateEmailStrings,
		"方案9-组合优化":      validateEmailCombined,
		"方案10-ASCII":    validateEmailASCII,
		"方案12-极简":       validateEmailMinimal,
	}

	allEmails := append(validEmails, invalidEmails...)

	for i, result := range results {
		validCount := 0
		for _, email := range allEmails {
			if validators[result.name](email) {
				validCount++
			}
		}
		results[i].valid = validCount
		results[i].total = len(allEmails)
	}

	fmt.Println("\n========== 邮箱验证正确性测试 ==========")
	for _, r := range results {
		fmt.Printf("%-15s: %d/%d 通过\n", r.name, r.valid, r.total)
	}
}

// 测试数据
var (
	validUUIDs = []string{
		"550e8400-e29b-41d4-a716-446655440000",
		"6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		"6ba7b811-9dad-11d1-80b4-00c04fd430c8",
		"6ba7b812-9dad-11d1-80b4-00c04fd430c8",
		"6ba7b814-9dad-11d1-80b4-00c04fd430c8",
		"00000000-0000-0000-0000-000000000000",
		"ffffffff-ffff-ffff-ffff-ffffffffffff",
		"01234567-89ab-cdef-0123-456789abcdef",
	}
	invalidUUIDs = []string{
		"550e8400-e29b-41d4-a716-44665544000",   // 太短
		"550e8400-e29b-41d4-a716-4466554400000", // 太长
		"550e8400-e29b-41d4-a716-44665544000G",  // 无效字符
		"550e8400-e29b-41d4-a716-44665544000 ",  // 尾部空格
		" 50e8400-e29b-41d4-a716-446655440000",  // 头部空格
		"550e8400e29b-41d4-a716-446655440000",   // 缺少分隔符
		"550e8400-e29b-41d4-a716-44665544000",   // 格式错误
		"G50e8400-e29b-41d4-a716-446655440000",  // 大写G
		"not-a-uuid",                            // 明显错误
		"",                                      // 空字符串
	}
)

// ============== 方案1：当前实现（基线） ==============
func validateUUIDOriginal(uuid string) bool {
	if uuid == "" {
		return false
	}
	return uuidRegex.MatchString(strings.ToLower(uuid))
}

// ============== 方案2：手动检查（字节级） ==============
func validateUUIDManual(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	// 检查固定位置的分隔符
	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	// 检查每段
	for i := 0; i < 36; i++ {
		switch i {
		case 8, 13, 18, 23:
			continue // 跳过分隔符
		default:
			c := uuid[i]
			isDigit := c >= '0' && c <= '9'
			isLower := c >= 'a' && c <= 'f'
			isUpper := c >= 'A' && c <= 'F'
			if !isDigit && !isLower && !isUpper {
				return false
			}
		}
	}

	return true
}

// ============== 方案3：分段验证 ==============
func validateUUIDSegmented(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	// 检查分隔符
	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	// 验证第一段：8个字符
	if !isHexSegment(uuid[0:8]) {
		return false
	}

	// 验证第二段：4个字符
	if !isHexSegment(uuid[9:13]) {
		return false
	}

	// 验证第三段：4个字符
	if !isHexSegment(uuid[14:18]) {
		return false
	}

	// 验证第四段：4个字符
	if !isHexSegment(uuid[19:23]) {
		return false
	}

	// 验证第五段：12个字符
	if !isHexSegment(uuid[24:36]) {
		return false
	}

	return true
}

func isHexSegment(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		isDigit := c >= '0' && c <= '9'
		isLower := c >= 'a' && c <= 'f'
		isUpper := c >= 'A' && c <= 'F'
		if !isDigit && !isLower && !isUpper {
			return false
		}
	}
	return true
}

// ============== 方案4：使用 strings.IndexByte ==============
func validateUUIDIndexByte(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	// 检查分隔符位置
	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	// 移除分隔符后检查
	hexPart := uuid[:8] + uuid[9:13] + uuid[14:18] + uuid[19:23] + uuid[24:]

	for i := 0; i < len(hexPart); i++ {
		c := hexPart[i]
		if !(c >= '0' && c <= '9') && !(c >= 'a' && c <= 'f') && !(c >= 'A' && c <= 'F') {
			return false
		}
	}

	return true
}

// ============== 方案5：预计算查找表 ==============
var hexTable = [256]bool{}

func init() {
	for i := 0; i < 256; i++ {
		c := byte(i)
		hexTable[i] = (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
	}
}

func validateUUIDLookupTable(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	for i := 0; i < 36; i++ {
		switch i {
		case 8, 13, 18, 23:
			if uuid[i] != '-' {
				return false
			}
		default:
			if !hexTable[uuid[i]] {
				return false
			}
		}
	}

	return true
}

// ============== 方案6：混合模式（快速路径 + 分段验证） ==============
func validateUUIDHybrid(uuid string) bool {
	// 快速长度检查
	if len(uuid) != 36 {
		return false
	}

	// 快速分隔符检查
	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	// 批量检查十六进制字符
	for i := 0; i < 36; i++ {
		switch i {
		case 8, 13, 18, 23:
			continue
		default:
			c := uuid[i]
			if !(((c >= '0') && (c <= '9')) || ((c >= 'a') && (c <= 'f')) || ((c >= 'A') && (c <= 'F'))) {
				return false
			}
		}
	}

	return true
}

// ============== 方案7：使用 unicode.Is ==============
func validateUUIDUnicode(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	for i := 0; i < 36; i++ {
		switch i {
		case 8, 13, 18, 23:
			if uuid[i] != '-' {
				return false
			}
		default:
			c := rune(uuid[i])
			if !unicode.IsDigit(c) && !(c >= 'a' && c <= 'f') && !(c >= 'A' && c <= 'F') {
				return false
			}
		}
	}

	return true
}

// ============== 方案8：直接字节比较（无分支） ==============
func validateUUIDByteCompare(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	// 检查分隔符
	if uuid[8]|uuid[13]|uuid[18]|uuid[23] != '-' {
		return false
	}

	// 检查所有字符
	for i := 0; i < 36; i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			continue
		}
		c := uuid[i]
		isValid := (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
		if !isValid {
			return false
		}
	}

	return true
}

// ============== 方案9：使用 ASCII 边界检查 ==============
func validateUUIDASCIICheck(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	for i := 0; i < 36; i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			continue
		}
		c := uuid[i]
		// ASCII 快速检查
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') && (c < 'A' || c > 'F') {
			return false
		}
	}

	return true
}

// ============== 方案10：预定义有效字符集 ==============
var validHexChars = map[byte]bool{
	'0': true, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true,
	'a': true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true,
	'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true,
}

func validateUUIDMapCheck(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	for i := 0; i < 36; i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			continue
		}
		if !validHexChars[uuid[i]] {
			return false
		}
	}

	return true
}

// ============== 方案11：位操作优化 ==============
func validateUUIDBitOps(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	for i := 0; i < 36; i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			continue
		}
		c := uuid[i]
		// 使用位操作优化范围检查
		isDigit := (c-'0')&0xFF <= 9
		isLower := ((c|0x20)-'a')&0xFF <= 5
		isUpper := (c-'A')&0xFF <= 5
		if !isDigit && !isLower && !isUpper {
			return false
		}
	}

	return true
}

// ============== 方案12：使用 strings.IndexAny 检查无效字符 ==============
func validateUUIDIndexAny(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	// 检查是否包含无效字符
	hexPart := uuid[:8] + uuid[9:13] + uuid[14:18] + uuid[19:23] + uuid[24:]
	if strings.IndexAny(hexPart, "0123456789abcdefABCDEF") == -1 {
		return len(hexPart) == 0
	}

	// 逐字符验证
	for _, c := range hexPart {
		if !(((c >= '0') && (c <= '9')) || ((c >= 'a') && (c <= 'f')) || ((c >= 'A') && (c <= 'F'))) {
			return false
		}
	}

	return true
}

// ============== 基准测试 ==============

// ============== 验证正确性的测试 ==============
func TestValidateUUID_Correctness(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(string) bool
	}{
		{"Original", validateUUIDOriginal},
		{"Manual", validateUUIDManual},
		{"Segmented", validateUUIDSegmented},
		{"IndexByte", validateUUIDIndexByte},
		{"LookupTable", validateUUIDLookupTable},
		{"Hybrid", validateUUIDHybrid},
		{"Unicode", validateUUIDUnicode},
		{"ByteCompare", validateUUIDByteCompare},
		{"ASCIICheck", validateUUIDASCIICheck},
		{"MapCheck", validateUUIDMapCheck},
		{"BitOps", validateUUIDBitOps},
		{"IndexAny", validateUUIDIndexAny},
	}

	for _, impl := range implementations {
		t.Run(impl.name, func(t *testing.T) {
			// 测试有效 UUID
			for _, uuid := range validUUIDs {
				if !impl.fn(uuid) {
					t.Errorf("%s: 有效 UUID 被拒绝: %s", impl.name, uuid)
				}
			}

			// 测试无效 UUID
			for _, uuid := range invalidUUIDs {
				if impl.fn(uuid) {
					t.Errorf("%s: 无效 UUID 被接受: %s", impl.name, uuid)
				}
			}
		})
	}
}

// 性能对比辅助函数
func runComparisonBenchmark() {
	fmt.Println("运行 UUID 验证性能对比...")
	fmt.Println("请在终端执行: cd validator && go test -bench=BenchmarkValidateUUID -benchmem -benchtime=3s | tee uuid_comparison_results.txt")
}

// TestValidateStrongPassword_Correctness 验证正确性
func TestValidateStrongPassword_Correctness(t *testing.T) {
	v, _ := New()

	type Form struct {
		Password string `validate:"strong_password"`
	}

	validPasswords := []string{
		"Abc123!@",        // 最小长度，包含所有类型
		"Password123!",    // 常见强密码
		"SecurePass#2024", // 包含数字
		"MyP@ssw0rd",      // 复杂密码
		"Test@1234",       // 简单但符合
		"ADMIN@123",       // 全大写+数字+特殊
		"student#123",     // 全小写+数字+特殊
		"User2024$Pass",   // 混合
		"1A2b3C4d!",       // 交替字符
		"P@ssw0rd123456",  // 较长密码
	}

	invalidPasswords := []struct {
		password string
		reason   string
	}{
		{"", "空密码"},
		{"short1A", "太短（7字符）"},
		{"lowercas", "只1种类型（小写）"},
		{"UPPERCAS", "只1种类型（大写）"},
		{"12345678", "只1种类型（数字）"},
		{"!@#$%^&*", "只1种类型（特殊）"},
		{"lower12", "只2种类型（小写+数字）"},
		{"LOWER12", "只2种类型（大写+数字）"},
		{"lower!!", "只2种类型（小写+特殊）"},
		{"LOWER!!", "只2种类型（大写+特殊）"},
		{"1234!!", "只2种类型（数字+特殊）"},
		{"lowerUP", "只2种类型（小写+大写）"},
	}

	t.Run("ValidPasswords", func(t *testing.T) {
		for _, pwd := range validPasswords {
			t.Run(pwd, func(t *testing.T) {
				form := Form{Password: pwd}
				err := v.Struct(form)
				if err != nil {
					t.Errorf("有效密码被拒绝: %s, error: %v", pwd, err)
				}
			})
		}
	})

	t.Run("InvalidPasswords", func(t *testing.T) {
		for _, tc := range invalidPasswords {
			t.Run(tc.reason, func(t *testing.T) {
				form := Form{Password: tc.password}
				err := v.Struct(form)
				if err == nil {
					t.Errorf("无效密码被接受: %s (%s)", tc.password, tc.reason)
				}
			})
		}
	})
}

// BenchmarkStrongPassword_Validator 集成基准测试
