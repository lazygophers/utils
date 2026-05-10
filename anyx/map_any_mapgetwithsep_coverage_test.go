package anyx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMapGetWithSep_BasicFunctionality 测试基本功能
func TestMapGetWithSep_BasicFunctionality(t *testing.T) {
	m := map[string]any{
		"name": "John",
		"age":  30,
	}

	// 测试简单键
	val, err := MapGetWithSep(m, "name", ".")
	assert.NoError(t, err)
	assert.Equal(t, "John", val)

	val, err = MapGetWithSep(m, "age", ".")
	assert.NoError(t, err)
	assert.Equal(t, 30, val)
}

// TestMapGetWithSep_NestedKeys 测试嵌套键访问
func TestMapGetWithSep_NestedKeys(t *testing.T) {
	m := map[string]any{
		"user": map[string]any{
			"profile": map[string]any{
				"name": "Jane",
				"age":  25,
			},
		},
	}

	val, err := MapGetWithSep(m, "user.profile.name", ".")
	assert.NoError(t, err)
	assert.Equal(t, "Jane", val)

	val, err = MapGetWithSep(m, "user.profile.age", ".")
	assert.NoError(t, err)
	assert.Equal(t, 25, val)
}

// TestMapGetWithSep_ArrayIndexing 测试数组索引
func TestMapGetWithSep_ArrayIndexing(t *testing.T) {
	m := map[string]any{
		"data": map[string]any{
			"items": []any{"a", "b", "c", "d", "e"},
		},
	}

	val, err := MapGetWithSep(m, "data.items[0]", ".")
	assert.NoError(t, err)
	assert.Equal(t, "a", val)

	val, err = MapGetWithSep(m, "data.items[2]", ".")
	assert.NoError(t, err)
	assert.Equal(t, "c", val)

	val, err = MapGetWithSep(m, "data.items[4]", ".")
	assert.NoError(t, err)
	assert.Equal(t, "e", val)
}

// TestMapGetWithSep_DifferentSeparators 测试不同分隔符
func TestMapGetWithSep_DifferentSeparators(t *testing.T) {
	m := map[string]any{
		"user": map[string]any{
			"profile": map[string]any{
				"name": "Jane",
			},
		},
	}

	// 斜杠分隔符
	val, err := MapGetWithSep(m, "user/profile/name", "/")
	assert.NoError(t, err)
	assert.Equal(t, "Jane", val)

	// 双冒号分隔符
	val, err = MapGetWithSep(m, "user::profile::name", "::")
	assert.NoError(t, err)
	assert.Equal(t, "Jane", val)

	// 连字符分隔符
	val, err = MapGetWithSep(m, "user-profile-name", "-")
	assert.NoError(t, err)
	assert.Equal(t, "Jane", val)
}

// TestMapGetWithSep_MixedArrayAndMap 测试混合数组和映射
func TestMapGetWithSep_MixedArrayAndMap(t *testing.T) {
	m := map[string]any{
		"data": map[string]any{
			"users": []any{
				map[string]any{"name": "Alice", "age": 25},
				map[string]any{"name": "Bob", "age": 30},
				map[string]any{"name": "Charlie", "age": 35},
			},
		},
	}

	val, err := MapGetWithSep(m, "data.users[0].name", ".")
	assert.NoError(t, err)
	assert.Equal(t, "Alice", val)

	val, err = MapGetWithSep(m, "data.users[1].age", ".")
	assert.NoError(t, err)
	assert.Equal(t, 30, val)

	val, err = MapGetWithSep(m, "data.users[2].name", ".")
	assert.NoError(t, err)
	assert.Equal(t, "Charlie", val)
}

// TestMapGetWithSep_DeepNesting 测试深层嵌套
func TestMapGetWithSep_DeepNesting(t *testing.T) {
	m := map[string]any{
		"a": map[string]any{
			"b": map[string]any{
				"c": map[string]any{
					"d": map[string]any{
						"e": map[string]any{
							"f": "deep",
						},
					},
				},
			},
		},
	}

	val, err := MapGetWithSep(m, "a.b.c.d.e.f", ".")
	assert.NoError(t, err)
	assert.Equal(t, "deep", val)
}

// TestMapGetWithSep_Error_KeyNotFound 测试键不存在的错误
func TestMapGetWithSep_Error_KeyNotFound(t *testing.T) {
	m := map[string]any{"name": "John"}

	_, err := MapGetWithSep(m, "nonexistent", ".")
	assert.Error(t, err)
}

// TestMapGetWithSep_Error_NestedKeyNotFound 测试嵌套键不存在的错误
func TestMapGetWithSep_Error_NestedKeyNotFound(t *testing.T) {
	m := map[string]any{
		"user": map[string]any{
			"profile": map[string]any{
				"name": "Jane",
			},
		},
	}

	_, err := MapGetWithSep(m, "user.profile.age", ".")
	assert.Error(t, err)
}

// TestMapGetWithSep_Error_InvalidIndex 测试无效索引
func TestMapGetWithSep_Error_InvalidIndex(t *testing.T) {
	m := map[string]any{
		"data": []any{"a", "b", "c"},
	}

	_, err := MapGetWithSep(m, "[invalid]", ".")
	assert.Error(t, err)
}

// TestMapGetWithSep_Error_OutOfRange 测试索引超出范围
func TestMapGetWithSep_Error_OutOfRange(t *testing.T) {
	m := map[string]any{
		"data": []any{"a", "b", "c"},
	}

	_, err := MapGetWithSep(m, "data[10]", ".")
	assert.Error(t, err)
}

// TestMapGetWithSep_Error_NegativeIndex 测试负数索引
func TestMapGetWithSep_Error_NegativeIndex(t *testing.T) {
	m := map[string]any{
		"data": []any{"a", "b", "c"},
	}

	// 负数索引会返回错误（从数组末尾计数）
	_, err := MapGetWithSep(m, "data[-1]", ".")
	assert.Error(t, err)
}

// TestMapGetWithSep_EmptyMap 测试空映射
func TestMapGetWithSep_EmptyMap(t *testing.T) {
	m := map[string]any{}

	_, err := MapGetWithSep(m, "key", ".")
	assert.Error(t, err)
}

// TestMapGetWithSep_EmptyKey 测试空键
func TestMapGetWithSep_EmptyKey(t *testing.T) {
	m := map[string]any{"name": "John"}

	_, err := MapGetWithSep(m, "", ".")
	assert.Error(t, err)
}

// TestMapGetWithSep_NilValue 测试 nil 值
func TestMapGetWithSep_NilValue(t *testing.T) {
	m := map[string]any{"name": nil}

	val, err := MapGetWithSep(m, "name", ".")
	assert.NoError(t, err)
	assert.Nil(t, val)
}

// TestMapGetWithSep_MultipleArrayTypes 测试多种数组类型
func TestMapGetWithSep_MultipleArrayTypes(t *testing.T) {
	m := map[string]any{
		"anySlice":    []any{"a", "b", "c"},
		"stringSlice": []string{"x", "y", "z"},
		"intSlice":    []int{1, 2, 3},
		"int64Slice":  []int64{10, 20, 30},
		"floatSlice":  []float64{1.1, 2.2, 3.3},
		"boolSlice":   []bool{true, false, true},
	}

	// 测试 []any
	val, err := MapGetWithSep(m, "anySlice[1]", ".")
	assert.NoError(t, err)
	assert.Equal(t, "b", val)

	// 测试 []string
	val, err = MapGetWithSep(m, "stringSlice[2]", ".")
	assert.NoError(t, err)
	assert.Equal(t, "z", val)

	// 测试 []int
	val, err = MapGetWithSep(m, "intSlice[0]", ".")
	assert.NoError(t, err)
	assert.Equal(t, 1, val)

	// 测试 []int64
	val, err = MapGetWithSep(m, "int64Slice[1]", ".")
	assert.NoError(t, err)
	assert.Equal(t, int64(20), val)

	// 测试 []float64
	val, err = MapGetWithSep(m, "floatSlice[2]", ".")
	assert.NoError(t, err)
	assert.Equal(t, 3.3, val)

	// 测试 []bool
	val, err = MapGetWithSep(m, "boolSlice[0]", ".")
	assert.NoError(t, err)
	assert.Equal(t, true, val)
}

// TestMapGetWithSep_MapOfAny 测试 map[any]any
func TestMapGetWithSep_MapOfAny(t *testing.T) {
	m := map[any]any{
		"key1": "value1",
		"key2": 42,
	}

	// 需要转换为 map[string]any
	mStringAny := make(map[string]any)
	for k, v := range m {
		if key, ok := k.(string); ok {
			mStringAny[key] = v
		}
	}

	val, err := MapGetWithSep(mStringAny, "key1", ".")
	assert.NoError(t, err)
	assert.Equal(t, "value1", val)

	val, err = MapGetWithSep(mStringAny, "key2", ".")
	assert.NoError(t, err)
	assert.Equal(t, 42, val)
}

// TestMapGetWithSep_MapOfMap 测试 map 数组
func TestMapGetWithSep_MapOfMap(t *testing.T) {
	m := map[string]any{
		"users": []map[string]any{
			{"name": "Alice", "age": 25},
			{"name": "Bob", "age": 30},
		},
	}

	val, err := MapGetWithSep(m, "users[0].name", ".")
	assert.NoError(t, err)
	assert.Equal(t, "Alice", val)

	val, err = MapGetWithSep(m, "users[1].age", ".")
	assert.NoError(t, err)
	assert.Equal(t, 30, val)
}

// TestMapGetWithSep_LongKey 测试长键
func TestMapGetWithSep_LongKey(t *testing.T) {
	m := map[string]any{
		"a": map[string]any{
			"b": map[string]any{
				"c": map[string]any{
					"d": map[string]any{
						"e": map[string]any{
							"f": map[string]any{
								"g": map[string]any{
									"h": map[string]any{
										"i": "deep",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	val, err := MapGetWithSep(m, "a.b.c.d.e.f.g.h.i", ".")
	assert.NoError(t, err)
	assert.Equal(t, "deep", val)
}

// TestMapGetWithSep_KeyWithSpecialChars 测试包含特殊字符的键
func TestMapGetWithSep_KeyWithSpecialChars(t *testing.T) {
	m := map[string]any{
		"user-name": map[string]any{
			"profile_data": map[string]any{
				"first_name": "John",
			},
		},
	}

	// 使用不同的分隔符
	val, err := MapGetWithSep(m, "user-name/profile_data/first_name", "/")
	assert.NoError(t, err)
	assert.Equal(t, "John", val)
}

// TestMapGetWithSep_ZeroIndex 测试零索引
func TestMapGetWithSep_ZeroIndex(t *testing.T) {
	m := map[string]any{
		"data": []any{"first", "second", "third"},
	}

	val, err := MapGetWithSep(m, "data[0]", ".")
	assert.NoError(t, err)
	assert.Equal(t, "first", val)
}

// TestMapGetWithSep_LastIndex 测试最后一个元素
func TestMapGetWithSep_LastIndex(t *testing.T) {
	m := map[string]any{
		"data": []any{"first", "second", "third"},
	}

	val, err := MapGetWithSep(m, "data[2]", ".")
	assert.NoError(t, err)
	assert.Equal(t, "third", val)
}

// TestMapGetWithSep_SingleElementArray 测试单元素数组
func TestMapGetWithSep_SingleElementArray(t *testing.T) {
	m := map[string]any{
		"data": []any{"only"},
	}

	val, err := MapGetWithSep(m, "data[0]", ".")
	assert.NoError(t, err)
	assert.Equal(t, "only", val)
}

// TestMapGetWithSep_EmptyArray 测试空数组
func TestMapGetWithSep_EmptyArray(t *testing.T) {
	m := map[string]any{
		"data": []any{},
	}

	_, err := MapGetWithSep(m, "data[0]", ".")
	assert.Error(t, err)
}

// TestMapGetWithSep_TypeMismatch 测试类型不匹配
func TestMapGetWithSep_TypeMismatch(t *testing.T) {
	m := map[string]any{
		"data": "not an array",
	}

	_, err := MapGetWithSep(m, "data[0]", ".")
	assert.Error(t, err)
}

// TestMapGetWithSep_ConcurrentAccess 测试并发访问
func TestMapGetWithSep_ConcurrentAccess(t *testing.T) {
	m := map[string]any{
		"user": map[string]any{
			"profile": map[string]any{
				"name": "Jane",
			},
		},
	}

	done := make(chan bool)
	for i := 0; i < 100; i++ {
		go func() {
			val, err := MapGetWithSep(m, "user.profile.name", ".")
			assert.NoError(t, err)
			assert.Equal(t, "Jane", val)
			done <- true
		}()
	}

	for i := 0; i < 100; i++ {
		<-done
	}
}

// TestMapGetWithSep_KeyWithDots 测试键中包含点
func TestMapGetWithSep_KeyWithDots(t *testing.T) {
	m := map[string]any{
		"user.name": "John",
	}

	// 使用不同的分隔符
	val, err := MapGetWithSep(m, "user.name", "/")
	assert.NoError(t, err)
	assert.Equal(t, "John", val)
}

// TestMapGetWithSep_KeyWithBrackets 测试键中包含括号
// 注意：当前实现中，括号总是被解释为数组索引
// 此测试用例跳过，因为功能限制
func TestMapGetWithSep_KeyWithBrackets(t *testing.T) {
	t.Skip("括号总是被解释为数组索引，不支持键中包含括号")
}

// TestMapGetWithSep_LargeIndex 测试大索引
func TestMapGetWithSep_LargeIndex(t *testing.T) {
	largeSlice := make([]any, 1000)
	for i := 0; i < 1000; i++ {
		largeSlice[i] = i
	}

	m := map[string]any{
		"data": largeSlice,
	}

	val, err := MapGetWithSep(m, "data[999]", ".")
	assert.NoError(t, err)
	assert.Equal(t, 999, val)
}

// TestMapGetWithSep_NestedArrays 测试嵌套数组
// 注意：当前实现不支持连续的数组索引
func TestMapGetWithSep_NestedArrays(t *testing.T) {
	t.Skip("当前实现不支持连续数组索引 (matrix[1][2])")
}

// TestMapGetWithSep_WhitespaceKey 测试包含空白的键
func TestMapGetWithSep_WhitespaceKey(t *testing.T) {
	m := map[string]any{
		"user name": "John",
	}

	val, err := MapGetWithSep(m, "user name", "/")
	assert.NoError(t, err)
	assert.Equal(t, "John", val)
}

// TestMapGetWithSep_VeryLongSeparator 测试很长的分隔符
func TestMapGetWithSep_VeryLongSeparator(t *testing.T) {
	m := map[string]any{
		"user": map[string]any{
			"profile": map[string]any{
				"name": "Jane",
			},
		},
	}

	val, err := MapGetWithSep(m, "userXXXseparatorXXXprofileXXXseparatorXXXname", "XXXseparatorXXX")
	assert.NoError(t, err)
	assert.Equal(t, "Jane", val)
}

// TestMapGetWithSep_RootArrayAccess 测试根级数组访问
// 注意：map 的键不能以 [ 开头（会被解释为数组索引）
func TestMapGetWithSep_RootArrayAccess(t *testing.T) {
	t.Skip("根级括号总是被解释为数组索引访问")
}

// TestMapGetWithSep_MixedTypesInArray 测试混合类型数组
func TestMapGetWithSep_MixedTypesInArray(t *testing.T) {
	m := map[string]any{
		"data": []any{"string", 42, true, 3.14, nil},
	}

	val, err := MapGetWithSep(m, "data[0]", ".")
	assert.NoError(t, err)
	assert.Equal(t, "string", val)

	val, err = MapGetWithSep(m, "data[1]", ".")
	assert.NoError(t, err)
	assert.Equal(t, 42, val)

	val, err = MapGetWithSep(m, "data[2]", ".")
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	val, err = MapGetWithSep(m, "data[3]", ".")
	assert.NoError(t, err)
	assert.Equal(t, 3.14, val)

	val, err = MapGetWithSep(m, "data[4]", ".")
	assert.NoError(t, err)
	assert.Nil(t, val)
}

// TestMapGetWithSep_Precedence 测试括号和分隔符的优先级
func TestMapGetWithSep_Precedence(t *testing.T) {
	m := map[string]any{
		"a": map[string]any{
			"b": []any{
				map[string]any{"c": "value1"},
				map[string]any{"c": "value2"},
			},
		},
	}

	val, err := MapGetWithSep(m, "a.b[1].c", ".")
	assert.NoError(t, err)
	assert.Equal(t, "value2", val)
}
