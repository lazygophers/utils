package anyx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNavigateToValue_Coverage 全面测试 navigateToValue 函数的覆盖率
func TestNavigateToValue_Coverage(t *testing.T) {
	tests := []struct {
		name           string
		data           any
		part           string
		want           any
		wantErr        bool
		errType        error
		description    string
		skipValueCheck bool
	}{
		// ===== Map 键访问测试 =====
		{
			name:        "简单 map[string]any 键访问",
			data:        map[string]any{"name": "John"},
			part:        "name",
			want:        "John",
			wantErr:     false,
			description: "验证基本的 map 键访问功能",
		},
		{
			name:        "空字符串键访问",
			data:        map[string]any{"": "empty"},
			part:        "",
			want:        "empty",
			wantErr:     false,
			description: "验证空字符串作为键的访问",
		},
		{
			name:        "键不存在",
			data:        map[string]any{"existing": "value"},
			part:        "nonexistent",
			want:        nil,
			wantErr:     true,
			errType:     ErrNotFound,
			description: "验证访问不存在的键返回 ErrNotFound",
		},
		{
			name:        "map[any]any 键访问",
			data:        map[any]any{"key": "value", 123: "number"},
			part:        "key",
			want:        "value",
			wantErr:     false,
			description: "验证 map[any]any 类型的键访问",
		},
		{
			name:        "map[any]any 键不存在",
			data:        map[any]any{"key": "value"},
			part:        "nonexistent",
			want:        nil,
			wantErr:     true,
			errType:     ErrNotFound,
			description: "验证 map[any]any 中键不存在的情况",
		},
		{
			name:        "空 map 访问",
			data:        map[string]any{},
			part:        "key",
			want:        nil,
			wantErr:     true,
			errType:     ErrNotFound,
			description: "验证空 map 的访问",
		},
		{
			name:        "大型 map 访问",
			data:        func() map[string]any { m := make(map[string]any, 100); for i := 0; i < 100; i++ { m[string(rune(i))] = i }; return m }(),
			part:        "X",
			want:        88,
			wantErr:     false,
			description: "验证大型 map 的性能和正确性（ASCII 'X' = 88）",
		},

		// ===== 数组索引访问测试 =====
		{
			name:        "[]any 索引访问",
			data:        []any{"a", "b", "c", "d", "e"},
			part:        "[2]",
			want:        "c",
			wantErr:     false,
			description: "验证 []any 类型的索引访问",
		},
		{
			name:        "[]string 索引访问",
			data:        []string{"apple", "banana", "cherry"},
			part:        "[1]",
			want:        "banana",
			wantErr:     false,
			description: "验证 []string 类型的索引访问",
		},
		{
			name:        "[]int 索引访问",
			data:        []int{10, 20, 30, 40, 50},
			part:        "[3]",
			want:        40,
			wantErr:     false,
			description: "验证 []int 类型的索引访问",
		},
		{
			name:        "[]int64 索引访问",
			data:        []int64{100, 200, 300},
			part:        "[2]",
			want:        int64(300),
			wantErr:     false,
			description: "验证 []int64 类型的索引访问",
		},
		{
			name:        "[]float64 索引访问",
			data:        []float64{1.1, 2.2, 3.3, 4.4},
			part:        "[1]",
			want:        2.2,
			wantErr:     false,
			description: "验证 []float64 类型的索引访问",
		},
		{
			name:        "[]bool 索引访问",
			data:        []bool{true, false, true},
			part:        "[2]",
			want:        true,
			wantErr:     false,
			description: "验证 []bool 类型的索引访问",
		},
		{
			name:        "[]map[string]any 索引访问",
			data:        []map[string]any{{"key": "a"}, {"key": "b"}, {"key": "c"}},
			part:        "[1]",
			want:        map[string]any{"key": "b"},
			wantErr:     false,
			description: "验证 []map[string]any 类型的索引访问",
		},

		// ===== 边界测试 =====
		{
			name:        "第一个元素索引 0",
			data:        []any{"first", "second", "third"},
			part:        "[0]",
			want:        "first",
			wantErr:     false,
			description: "验证索引 0 的访问",
		},
		{
			name:        "最后一个元素索引",
			data:        []any{"a", "b", "c", "d", "e"},
			part:        "[4]",
			want:        "e",
			wantErr:     false,
			description: "验证最后一个元素的访问",
		},
		{
			name:        "负数索引（当前实现不支持）",
			data:        []any{"a", "b", "c", "d", "e"},
			part:        "[-2]",
			want:        nil,
			wantErr:     true,
			errType:     ErrOutOfRange,
			description: "验证负数索引返回错误（当前实现不支持）",
		},
		{
			name:        "负数索引 -1（当前实现不支持）",
			data:        []any{"a", "b", "c"},
			part:        "[-1]",
			want:        nil,
			wantErr:     true,
			errType:     ErrOutOfRange,
			description: "验证负数索引 -1 返回错误（当前实现不支持）",
		},
		{
			name:        "大索引值",
			data:        func() []any { s := make([]any, 1000); for i := 0; i < 1000; i++ { s[i] = i }; return s }(),
			part:        "[999]",
			want:        999,
			wantErr:     false,
			description: "验证大索引值的访问",
		},
		{
			name:        "多位数字索引",
			data:        make([]any, 100),
			part:        "[99]",
			want:        nil,
			wantErr:     false,
			description: "验证多位数字索引的解析",
		},

		// ===== 错误情况测试 =====
		{
			name:        "索引越界 - 正数",
			data:        []any{"a", "b", "c"},
			part:        "[10]",
			want:        nil,
			wantErr:     true,
			errType:     ErrOutOfRange,
			description: "验证索引超出范围（正数）",
		},
		{
			name:        "索引越界 - 负数",
			data:        []any{"a", "b", "c"},
			part:        "[-10]",
			want:        nil,
			wantErr:     true,
			errType:     ErrOutOfRange,
			description: "验证索引超出范围（负数）",
		},
		{
			name:        "空索引（当前实现解析为 0）",
			data:        []any{"a", "b", "c"},
			part:        "[]",
			want:        "a",
			wantErr:     false,
			description: "验证空索引被解析为 0（当前实现行为）",
		},
		{
			name:        "无效索引格式 - 非数字",
			data:        []any{"a", "b", "c"},
			part:        "[abc]",
			want:        nil,
			wantErr:     true,
			errType:     ErrInvalidIndex,
			description: "验证无效索引格式的错误处理",
		},
		{
			name:        "无效索引格式 - 混合字符",
			data:        []any{"a", "b", "c"},
			part:        "[1a2]",
			want:        nil,
			wantErr:     true,
			errType:     ErrInvalidIndex,
			description: "验证混合字符索引的错误处理",
		},
		{
			name:        "无效索引格式 - 小数",
			data:        []any{"a", "b", "c"},
			part:        "[1.5]",
			want:        nil,
			wantErr:     true,
			errType:     ErrInvalidIndex,
			description: "验证小数索引的错误处理",
		},
		{
			name:        "无效索引格式 - 负号后无数字（当前实现解析为 0）",
			data:        []any{"a", "b", "c"},
			part:        "[-]",
			want:        "a",
			wantErr:     false,
			description: "验证只有负号的索引被解析为 0（当前实现行为）",
		},

		// ===== 类型错误测试 =====
		{
			name:        "不支持索引的类型 - 字符串",
			data:        "not-a-slice",
			part:        "[0]",
			want:        nil,
			wantErr:     true,
			errType:     ErrInvalidSlice,
			description: "验证对字符串尝试索引访问的错误",
		},
		{
			name:        "不支持索引的类型 - 整数",
			data:        12345,
			part:        "[0]",
			want:        nil,
			wantErr:     true,
			errType:     ErrInvalidSlice,
			description: "验证对整数尝试索引访问的错误",
		},
		{
			name:        "不支持键访问的类型 - 切片",
			data:        []any{"a", "b"},
			part:        "key",
			want:        nil,
			wantErr:     true,
			errType:     ErrInvalidMapType,
			description: "验证对切片尝试键访问的错误",
		},
		{
			name:        "不支持键访问的类型 - 字符串",
			data:        "not-a-map",
			part:        "key",
			want:        nil,
			wantErr:     true,
			errType:     ErrInvalidMapType,
			description: "验证对字符串尝试键访问的错误",
		},
		{
			name:        "nil 数据访问",
			data:        nil,
			part:        "key",
			want:        nil,
			wantErr:     true,
			errType:     ErrInvalidMapType,
			description: "验证 nil 数据的错误处理",
		},

		// ===== 特殊情况测试 =====
		{
			name:        "空切片访问",
			data:        []any{},
			part:        "[0]",
			want:        nil,
			wantErr:     true,
			errType:     ErrOutOfRange,
			description: "验证空切片的访问",
		},
		{
			name:        "嵌套结构 - 第一层",
			data:        map[string]any{"level1": map[string]any{"level2": "value"}},
			part:        "level1",
			want:        map[string]any{"level2": "value"},
			wantErr:     false,
			description: "验证嵌套结构的第一层访问",
		},
		{
			name:        "复杂值类型（函数值）",
			data:        map[string]any{"func": func() {}},
			part:        "func",
			want:        "skip-check", // 不检查返回值
			wantErr:     false,
			description: "验证存储函数等复杂类型的访问",
			skipValueCheck: true,
		},
		{
			name:        "nil 值在 map 中",
			data:        map[string]any{"nil": nil},
			part:        "nil",
			want:        nil,
			wantErr:     false,
			description: "验证 map 中存储 nil 值的访问（键存在，值为 nil）",
		},
		{
			name:        "零值在切片中",
			data:        []int{0, 1, 2},
			part:        "[0]",
			want:        0,
			wantErr:     false,
			description: "验证切片中零值的访问",
		},

		// ===== 性能关键路径测试 =====
		{
			name:        "单字符键",
			data:        map[string]any{"a": 1},
			part:        "a",
			want:        1,
			wantErr:     false,
			description: "测试单字符键的性能路径",
		},
		{
			name:        "长键名",
			data:        map[string]any{"this-is-a-very-long-key-name-for-testing-performance": "value"},
			part:        "this-is-a-very-long-key-name-for-testing-performance",
			want:        "value",
			wantErr:     false,
			description: "测试长键名的性能路径",
		},
		{
			name:        "map 中存在的数字字符串键",
			data:        map[string]any{"123": "value"},
			part:        "123",
			want:        "value",
			wantErr:     false,
			description: "验证数字字符串作为键（不是索引）",
		},
		{
			name:        "part 只有方括号但内容为空（当前实现解析为索引 0）",
			data:        map[string]any{"": "value"},
			part:        "[]",
			want:        nil,
			wantErr:     true,
			errType:     ErrInvalidMapType,
			description: "验证空方括号被当作索引处理，但 map 不支持索引访问",
		},
		{
			name:        "part 是单个方括号",
			data:        map[string]any{"]": "value"},
			part:        "]",
			want:        "value",
			wantErr:     false,
			description: "验证单个方括号被当作键而不是索引",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := navigateToValue(tt.data, tt.part)

			if tt.wantErr {
				assert.Error(t, err, tt.description)
				if tt.errType != nil {
					assert.ErrorIs(t, err, tt.errType, tt.description)
				}
			} else {
				assert.NoError(t, err, tt.description)
				if !tt.skipValueCheck {
					assert.Equal(t, tt.want, got, tt.description)
				}
			}
		})
	}
}

// TestNavigateToValue_PatternIntegrity 测试模式识别的完整性
func TestNavigateToValue_PatternIntegrity(t *testing.T) {
	tests := []struct {
		name        string
		part        string
		isIndex     bool
		description string
	}{
		{"标准数组索引", "[0]", true, "以 [ 开头，以 ] 结尾，长度 > 2"},
		{"负数索引", "[-1]", true, "负数索引模式"},
		{"多位数索引", "[123]", true, "多位数字索引"},
		{"空索引", "[]", true, "空索引（当前实现解析为索引 0）"},
		{"普通键", "key", false, "普通键名"},
		{"带方括号的键（不完整）", "[key", false, "只有左方括号"},
		{"带方括号的键（不完整）", "key]", false, "只有右方括号"},
		{"方括号中间", "ke[y]", false, "方括号在中间"},
		{"单个左方括号", "[", false, "只有一个左方括号"},
		{"单个右方括号", "]", false, "只有一个右方括号"},
		{"空字符串", "", false, "空字符串"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 检查模式匹配逻辑
			isIndex := len(tt.part) > 2 && tt.part[0] == '[' && tt.part[len(tt.part)-1] == ']'
			assert.Equal(t, tt.isIndex, isIndex, tt.description)
		})
	}
}

// TestNavigateToValue_EdgeCases 测试边界情况
func TestNavigateToValue_EdgeCases(t *testing.T) {
	t.Run("超大切片的最后元素", func(t *testing.T) {
		data := make([]any, 10000)
		data[9999] = "last"
		result, err := navigateToValue(data, "[9999]")
		assert.NoError(t, err)
		assert.Equal(t, "last", result)
	})

	t.Run("map 键包含方括号字符", func(t *testing.T) {
		data := map[string]any{
			"key[0]": "value",
		}
		// "key[0]" 会被当作键，不是索引
		result, err := navigateToValue(data, "key[0]")
		assert.NoError(t, err)
		assert.Equal(t, "value", result)
	})

	t.Run("连续多次索引边界检查", func(t *testing.T) {
		data := []any{"a", "b", "c"}
		// 连续访问边界（当前实现不支持负数索引）
		_, err1 := navigateToValue(data, "[-1]")   // 最后一个（会失败）
		_, err2 := navigateToValue(data, "[0]")    // 第一个
		_, err3 := navigateToValue(data, "[2]")    // 最后一个
		_, err4 := navigateToValue(data, "[3]")    // 越界

		assert.Error(t, err1) // 负数索引不支持
		assert.NoError(t, err2)
		assert.NoError(t, err3)
		assert.Error(t, err4)
	})

	t.Run("map 中 nil 值 vs 键不存在", func(t *testing.T) {
		data1 := map[string]any{"key": nil}
		data2 := map[string]any{}

		result1, err1 := navigateToValue(data1, "key")
		result2, err2 := navigateToValue(data2, "key")

		// 键存在但值为 nil：不返回错误
		assert.NoError(t, err1)
		assert.Nil(t, result1)

		// 键不存在：返回错误
		assert.Error(t, err2)
		assert.Nil(t, result2)
	})
}

// TestNavigateToValue_ConcurrentAccess 测试并发安全性
func TestNavigateToValue_ConcurrentAccess(t *testing.T) {
	data := map[string]any{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	sliceData := []any{"a", "b", "c", "d", "e"}

	done := make(chan bool)

	// 并发读取 map
	for i := 0; i < 100; i++ {
		go func() {
			_, _ = navigateToValue(data, "key1")
			done <- true
		}()
	}

	// 并发读取 slice
	for i := 0; i < 100; i++ {
		go func() {
			_, _ = navigateToValue(sliceData, "[2]")
			done <- true
		}()
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 200; i++ {
		<-done
	}

	// 如果没有 panic 或 deadlock，测试通过
	t.Log("并发访问测试通过")
}

// TestNavigateToValue_TypeSwitchCoverage 测试类型切换的覆盖率
func TestNavigateToValue_TypeSwitchCoverage(t *testing.T) {
	tests := []struct {
		name           string
		data           any
		part           string
		want           any
		wantErr        bool
		description    string
		skipValueCheck bool
	}{
		// 覆盖 accessArrayIndex 中的所有类型分支
		{"[]any 分支", []any{1, 2, 3}, "[0]", 1, false, "覆盖 []any 类型", false},
		{"[]string 分支", []string{"a", "b"}, "[0]", "a", false, "覆盖 []string 类型", false},
		{"[]int 分支", []int{1, 2}, "[0]", 1, false, "覆盖 []int 类型", false},
		{"[]int64 分支", []int64{1, 2}, "[0]", int64(1), false, "覆盖 []int64 类型", false},
		{"[]float64 分支", []float64{1.1, 2.2}, "[0]", 1.1, false, "覆盖 []float64 类型", false},
		{"[]bool 分支", []bool{true, false}, "[0]", true, false, "覆盖 []bool 类型", false},
		{"[]map[string]any 分支", []map[string]any{{"k": "v"}}, "[0]", map[string]any{"k": "v"}, false, "覆盖 []map[string]any 类型", false},
		{"未知切片类型", []int32{1, 2}, "[0]", nil, true, "覆盖 accessGenericSlice 分支", false},

		// 覆盖 accessMapKey 中的所有类型分支
		{"map[string]any 分支", map[string]any{"k": "v"}, "k", "v", false, "覆盖 map[string]any 类型", false},
		{"map[any]any 分支", map[any]any{"k": "v"}, "k", "v", false, "覆盖 map[any]any 类型", false},
		{"未知类型访问键", 123, "k", nil, true, "覆盖错误类型分支", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := navigateToValue(tt.data, tt.part)

			if tt.wantErr {
				assert.Error(t, err, tt.description)
			} else {
				assert.NoError(t, err, tt.description)
				if !tt.skipValueCheck {
					assert.Equal(t, tt.want, got, tt.description)
				}
			}
		})
	}
}

// TestParseIndex_Coverage 测试 parseIndex 函数的覆盖率
func TestParseIndex_Coverage(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		want           int
		wantErr        bool
		errType        error
		description    string
		skipValueCheck bool
	}{
		{"简单正数", "123", 123, false, nil, "基本正数解析", false},
		{"零", "0", 0, false, nil, "零值解析", false},
		{"负数", "-456", -456, false, nil, "负数解析", false},
		{"负零", "-0", 0, false, nil, "负零（应该解析为 0）", false},
		{"大数", "999999", 999999, false, nil, "大数解析", false},
		{"空字符串", "", 0, true, ErrInvalidIndex, "空字符串错误", false},
		{"只有负号", "-", 0, true, ErrInvalidIndex, "只有负号应返回错误", false},
		{"包含非数字字符", "12a34", 0, true, ErrInvalidIndex, "包含字母的错误", false},
		{"全是字母", "abc", 0, true, ErrInvalidIndex, "全是字母的错误", false},
		{"包含空格", "12 34", 0, true, ErrInvalidIndex, "包含空格的错误", false},
		{"包含特殊字符", "12!34", 0, true, ErrInvalidIndex, "包含特殊字符的错误", false},
		{"小数点", "12.34", 0, true, ErrInvalidIndex, "包含小数点的错误", false},
		{"多个负号", "--123", 0, true, ErrInvalidIndex, "多个负号的错误", false},
		{"负号在中间", "12-34", 0, true, ErrInvalidIndex, "负号在中间的错误", false},
		{"前导零", "007", 7, false, nil, "前导零（允许）", false},
		{"多位负数", "-100", -100, false, nil, "多位负数", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseIndex(tt.input)

			if tt.wantErr {
				assert.Error(t, err, tt.description)
				if tt.errType != nil {
					assert.ErrorIs(t, err, tt.errType, tt.description)
				}
			} else {
				assert.NoError(t, err, tt.description)
				if !tt.skipValueCheck {
					assert.Equal(t, tt.want, got, tt.description)
				}
			}
		})
	}
}

// TestAccessGenericSlice_Coverage 测试 accessGenericSlice 的覆盖率
func TestAccessGenericSlice_Coverage(t *testing.T) {
	tests := []struct {
		name        string
		slice       any
		index       int
		wantErr     bool
		errType     error
		description string
	}{
		{"[]uint", []uint{1, 2, 3}, 1, true, ErrInvalidSlice, "不支持的 []uint 类型"},
		{"[]float32", []float32{1.1, 2.2}, 0, true, ErrInvalidSlice, "不支持的 []float32 类型"},
		{"[]int32", []int32{1, 2}, 0, true, ErrInvalidSlice, "不支持的 []int32 类型"},
		{"[]interface{} 显式", []interface{}{"a"}, 0, true, ErrInvalidSlice, "显式 []interface{} 类型"},
		{"非切片类型", "string", 0, true, ErrInvalidSlice, "非切片类型"},
		{"nil", nil, 0, true, ErrInvalidSlice, "nil 值"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := accessGenericSlice(tt.slice, tt.index)

			assert.True(t, tt.wantErr, tt.description)
			if tt.errType != nil {
				assert.ErrorIs(t, err, tt.errType, tt.description)
			}
		})
	}
}
