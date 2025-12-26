package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
		RequiredField string       `validate:"required" json:"required_field"`
		Nested        NestedStruct `json:"nested"`
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
		StringMin  string `validate:"min=3" json:"string_min"`
		SliceMin   []int  `validate:"min=2" json:"slice_min"`
		MapMin     map[string]int `validate:"min=2" json:"map_min"`
		IntMin     int    `validate:"min=5" json:"int_min"`
		UintMin    uint   `validate:"min=5" json:"uint_min"`
		FloatMin   float64 `validate:"min=5.5" json:"float_min"`
		StringMax  string `validate:"max=5" json:"string_max"`
		SliceMax   []int  `validate:"max=2" json:"slice_max"`
		MapMax     map[string]int `validate:"max=2" json:"map_max"`
		IntMax     int    `validate:"max=5" json:"int_max"`
		UintMax    uint   `validate:"max=5" json:"uint_max"`
		FloatMax   float64 `validate:"max=5.5" json:"float_max"`
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
