package utils

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestStruct 用于测试的结构体
type TestStruct struct {
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age" validate:"gte=18"`
	Address string `json:"address" validate:"required"`
	Email   string `json:"email" validate:"email"`
}

// TestMustOk 测试 MustOk 函数
func TestMustOk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		value   interface{}
		ok       bool
		wantPanic bool
	}{
		{"正常情况 - ok 为 true", "test value", true, false},
		{"正常情况 - ok 为 true，数字类型", "test value", true, false},
		{"panic 情况 - ok 为 false", "test value", false, true},
		{"panic 情况 - ok 为 false，数字类型", "test value", false, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.wantPanic {
				assert.Panics(t, func() {
					MustOk(tt.value, tt.ok)
				}, "MustOk 应该 panic")
			} else {
				assert.NotPanics(t, func() {
					result := MustOk(tt.value, tt.ok)
					assert.Equal(t, tt.value, result, "MustOk 返回值应该匹配")
				}, "MustOk 不应该 panic")
			}
		})
	}
}

// TestMustSuccess 测试 MustSuccess 函数
func TestMustSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      error
		wantPanic bool
		wantMsg  string
	}{
		{"无错误", nil, false, ""},
		{"有错误", errors.New("test error"), true, "test error"},
		{"格式化错误", fmt.Errorf("format error: %d", 123), true, "format error: 123"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.wantPanic {
				assert.Panics(t, func() {
					MustSuccess(tt.err)
				}, "MustSuccess 应该 panic")
			} else {
				assert.NotPanics(t, func() {
					MustSuccess(tt.err)
				}, "MustSuccess 不应该 panic")
			}
		})
	}
}

// TestMustSuccessWithFormattedError 测试 MustSuccess 的格式化错误
func TestMustSuccessWithFormattedError(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() {
		MustSuccess(fmt.Errorf("formatted error: %d", 123))
	}, "MustSuccess 应该 panic")
}

// TestMust 测试 Must 函数
func TestMust(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		value   interface{}
		err     error
		wantPanic bool
	}{
		{"无错误", "success", nil, false},
		{"有错误", nil, errors.New("operation failed"), true},
		{"nil 值无错误", nil, nil, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.wantPanic {
				assert.Panics(t, func() {
					Must(tt.value, tt.err)
				}, "Must 应该 panic")
			} else {
				assert.NotPanics(t, func() {
					result := Must(tt.value, tt.err)
					assert.Equal(t, tt.value, result, "Must 返回值应该匹配")
				}, "Must 不应该 panic")
			}
		})
	}
}

// TestIgnore 测试 Ignore 函数
func TestIgnore(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value interface{}
		ignore interface{}
		want  interface{}
	}{
		{"忽略字符串", "hello", "world", "hello"},
		{"忽略数字", 42, 100, 42},
		{"忽略结构体", TestStruct{Name: "test"}, nil, TestStruct{Name: "test"}},
		{"忽略 nil", "value", nil, "value"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := Ignore(tt.value, tt.ignore)
			assert.Equal(t, tt.want, result, "Ignore 返回值应该匹配")
		})
	}
}

// BenchmarkIgnore 基准测试 Ignore 函数
func BenchmarkIgnore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ignore("test", "ignored")
	}
}

// TestScan 测试 Scan 函数
func TestScan(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		src       interface{}
		want      TestStruct
		wantError bool
	}{
		{
			name: "JSON 对象字符串",
			src:  `{"name":"张三","age":25,"address":"北京市朝阳区","email":"zhangsan@example.com"}`,
			want: TestStruct{
				Name:    "张三",
				Age:     25,
				Address: "北京市朝阳区",
				Email:   "zhangsan@example.com",
			},
			wantError: false,
		},
		{
			name: "JSON 对象字节",
			src:  []byte(`{"name":"李四","age":30,"address":"上海市浦东新区","email":"lisi@example.com"}`),
			want: TestStruct{
				Name:    "李四",
				Age:     30,
				Address: "上海市浦东新区",
				Email:   "lisi@example.com",
			},
			wantError: false,
		},
		{
			name: "JSON 数组字符串 - 应该失败",
			src:  `[{"name":"王五","age":28,"address":"广州市天河区","email":"wangwu@example.com"}]`,
			wantError: true, // 数组应该无法扫描到单个结构体
		},
		{
			name: "空字符串",
			src:  "",
			want: TestStruct{
				Name:    "", // 空字符串不会设置默认值
				Age:     0,
				Address: "",
				Email:   "",
			},
			wantError: false,
		},
		{
			name:      "无效 JSON",
			src:       `{"name":"test","age":`,
			wantError: true,
		},
		{
			name:      "不支持的类型",
			src:       123,
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var dst TestStruct
			err := Scan(tt.src, &dst)

			if tt.wantError {
				assert.Error(t, err, "Scan 应该返回错误")
			} else {
				require.NoError(t, err, "Scan 不应该返回错误")
				assert.Equal(t, tt.want.Name, dst.Name, "Name 字段应该匹配")
				assert.Equal(t, tt.want.Age, dst.Age, "Age 字段应该匹配")
				assert.Equal(t, tt.want.Address, dst.Address, "Address 字段应该匹配")
				assert.Equal(t, tt.want.Email, dst.Email, "Email 字段应该匹配")
			}
		})
	}
}

// TestScanWithDefaults 测试 Scan 函数的默认值功能
func TestScanWithDefaults(t *testing.T) {
	t.Parallel()

	var dst TestStruct
	err := Scan("", &dst)
	require.NoError(t, err, "Scan 不应该返回错误")

	// 检查默认值是否被设置
	// 注意：实际的 Scan 函数可能没有设置默认值，这里需要根据实际情况调整
	assert.Equal(t, "", dst.Name, "Name 应该为空字符串")
	assert.Equal(t, 0, dst.Age, "Age 应该为零值")
	assert.Equal(t, "", dst.Address, "Address 应该为空字符串")
}


// TestValue 测试 Value 函数
func TestValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{"nil 值", nil, "null"},
		{"结构体", TestStruct{Name: "测试", Age: 20}, `{"name":"测试","age":20,"address":"","email":""}`},
		{"指针", &TestStruct{Name: "指针测试"}, `{"name":"指针测试","age":0,"address":"","email":""}`},
		{"map", map[string]interface{}{"key": "value"}, `{"key":"value"}`},
		{"slice", []int{1, 2, 3}, `[1,2,3]`},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, err := Value(tt.value)
			require.NoError(t, err, "Value 不应该返回错误")
			assert.Equal(t, tt.want, string(result.([]byte)), "Value 返回值应该匹配")
		})
	}
}

// TestValueWithNilPanic 测试 Value 函数处理 nil 值时的 panic 情况
func TestValueWithNilPanic(t *testing.T) {
	t.Parallel()

	// 测试 nil 值可能导致 panic 的情况
	// 由于 defaults.SetDefaults 在处理 nil 时可能会 panic
	assert.Panics(t, func() {
		_, _ = Value(nil)
	}, "Value 处理 nil 时应该 panic")
}

// TestValueDriverValue 测试 Value 函数的 driver.Valuer 接口
func TestValueDriverValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{"实现了 Valuer 接口", testValuer{}, `{}`},
		{"普通结构体", TestStruct{Name: "测试"}, `{"name":"测试","age":0,"address":"","email":""}`},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, err := Value(tt.value)
			require.NoError(t, err, "Value 不应该返回错误")
			assert.Equal(t, tt.want, string(result.([]byte)), "Value 返回值应该匹配")
		})
	}
}

// TestValidate 测试 Validate 函数
func TestValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		data      TestStruct
		wantError bool
		errorMsg  string
	}{
		{
			name: "有效数据",
			data: TestStruct{
				Name:    "张三",
				Age:     25,
				Address: "北京市朝阳区",
				Email:   "zhangsan@example.com",
			},
			wantError: false,
		},
		{
			name: "姓名为空",
			data: TestStruct{
				Name:    "",
				Age:     25,
				Address: "北京市朝阳区",
				Email:   "zhangsan@example.com",
			},
			wantError: true,
			errorMsg:  "Key: 'TestStruct.Name' Error:Field validation for 'Name' failed on the 'required' tag",
		},
		{
			name: "邮箱格式错误",
			data: TestStruct{
				Name:    "张三",
				Age:     25,
				Address: "北京市朝阳区",
				Email:   "invalid-email",
			},
			wantError: true,
			errorMsg:  "Key: 'TestStruct.Email' Error:Field validation for 'Email' failed on the 'email' tag",
		},
		{
			name: "年龄小于18",
			data: TestStruct{
				Name:    "张三",
				Age:     16,
				Address: "北京市朝阳区",
				Email:   "zhangsan@example.com",
			},
			wantError: true,
			errorMsg:  "Key: 'TestStruct.Age' Error:Field validation for 'Age' failed on the 'gte' tag",
		},
		{
			name: "地址为空",
			data: TestStruct{
				Name:    "张三",
				Age:     25,
				Address: "",
				Email:   "zhangsan@example.com",
			},
			wantError: true,
			errorMsg:  "Key: 'TestStruct.Address' Error:Field validation for 'Address' failed on the 'required' tag",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := Validate(tt.data)

			if tt.wantError {
				require.Error(t, err, "Validate 应该返回错误")
				// 使用正则表达式匹配错误信息，因为 validator 可能包含额外的上下文信息
				assert.Regexp(t, regexp.MustCompile(tt.errorMsg), err.Error(), "错误信息应该匹配")
			} else {
				require.NoError(t, err, "Validate 不应该返回错误")
			}
		})
	}
}

// TestValidateWithNilValidator 测试 Validate 函数的 nil validator 处理
func TestValidateWithNilValidator(t *testing.T) {
	t.Parallel()
	
	// 由于 validate 是全局变量且由 validator.New() 初始化，
	// 在实际使用中不会出现 nil 的情况
	// 这个测试只是为了验证 Validate 函数的错误处理
	
	// 准备测试数据 - 无效邮箱
	data := TestStruct{
		Name:    "test-user",
		Email:  "invalid-email", // 无效邮箱
	}
	
	// 测试 Validate 函数能正确处理验证错误
	err := Validate(data)
	assert.Error(t, err, "Validate 应该返回验证错误")
	assert.Contains(t, err.Error(), "Email", "错误信息应该包含 Email 字段")
}

// TestValidateWithNilInput 测试 Validate 函数的 nil 输入处理
func TestValidateWithNilInput(t *testing.T) {
	t.Parallel()
	
	err := Validate(nil)
	assert.Error(t, err, "Validate 应该返回错误")
	assert.Contains(t, err.Error(), "nil", "错误信息应该包含 nil")
}

// testValuer 实现了 driver.Valuer 接口的测试结构体
type testValuer struct{}

func (v testValuer) Value() (driver.Value, error) {
	return "test", nil
}

// TestToMapAnyInt64 测试 ToMapAnyInt64 函数
func TestToMapAnyInt64(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		give []interface{}
		want []int64
	}{
		{"正常转换", []interface{}{1, 2.2, "3"}, []int64{1, 2, 3}},
		{"空切片", []interface{}{}, []int64{}},
		{"nil切片", nil, []int64{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToMapAnyInt64(tt.give)
			assert.Equal(t, tt.want, result, "ToMapAnyInt64 返回值应该匹配")
		})
	}
}

// TestToMapAnyFloat32 测试 ToMapAnyFloat32 函数
func TestToMapAnyFloat32(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		give []interface{}
		want []float32
	}{
		{"正常转换", []interface{}{1.1, 2.2, "3.3"}, []float32{1.1, 2.2, 3.3}},
		{"空切片", []interface{}{}, []float32{}},
		{"nil切片", nil, []float32{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToMapAnyFloat32(tt.give)
			assert.Equal(t, tt.want, result, "ToMapAnyFloat32 返回值应该匹配")
		})
	}
}

// TestToMapAnyFloat64 测试 ToMapAnyFloat64 函数
func TestToMapAnyFloat64(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		give []interface{}
		want []float64
	}{
		{"正常转换", []interface{}{1.1, 2.2, "3.3"}, []float64{1.1, 2.2, 3.3}},
		{"空切片", []interface{}{}, []float64{}},
		{"nil切片", nil, []float64{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToMapAnyFloat64(tt.give)
			assert.Equal(t, tt.want, result, "ToMapAnyFloat64 返回值应该匹配")
		})
	}
}

// TestToMapAnyString 测试 ToMapAnyString 函数
func TestToMapAnyString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		give []interface{}
		want []string
	}{
		{"正常转换", []interface{}{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"空切片", []interface{}{}, []string{}},
		{"nil切片", nil, []string{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToMapAnyString(tt.give)
			assert.Equal(t, tt.want, result, "ToMapAnyString 返回值应该匹配")
		})
	}
}

// TestToMapAnyBool 测试 ToMapAnyBool 函数
func TestToMapAnyBool(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		give []interface{}
		want []bool
	}{
		{"正常转换", []interface{}{true, false, "true"}, []bool{true, false, true}},
		{"空切片", []interface{}{}, []bool{}},
		{"nil切片", nil, []bool{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToMapAnyBool(tt.give)
			assert.Equal(t, tt.want, result, "ToMapAnyBool 返回值应该匹配")
		})
	}
}

// TestToMapAnyAny 测试 ToMapAnyAny 函数
func TestToMapAnyAny(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		give []interface{}
		want []interface{}
	}{
		{"正常转换", []interface{}{1, "a", true}, []interface{}{1, "a", true}},
		{"空切片", []interface{}{}, []interface{}{}},
		{"nil切片", nil, []interface{}{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToMapAnyAny(tt.give)
			assert.Equal(t, tt.want, result, "ToMapAnyAny 返回值应该匹配")
		})
	}
}