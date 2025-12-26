package anyx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试get函数的完整覆盖
func TestGetFunctionCoverage(t *testing.T) {
	// 创建一个带有嵌套结构的map
	m := map[string]interface{}{
		"key1": "value1",
		"nested": map[string]interface{}{
			"key2": 42,
			"deep": map[string]interface{}{
				"key3": true,
			},
		},
	}

	// 测试1: 未启用cut功能，直接查找不存在的key
	t.Run("not_found_without_cut", func(t *testing.T) {
		mapAny := NewMap(m)
		val, ok := mapAny.get("nonexistent")
		assert.False(t, ok)
		assert.Nil(t, val)
	})

	// 测试2: 启用cut功能，查找不存在的嵌套key
	t.Run("not_found_with_cut", func(t *testing.T) {
		mapAny := NewMap(m)
		mapAny.EnableCut(".")
		val, ok := mapAny.get("nonexistent.key")
		assert.False(t, ok)
		assert.Nil(t, val)
	})

	// 测试3: 启用cut功能，查找存在的嵌套key
	t.Run("found_with_cut", func(t *testing.T) {
		mapAny := NewMap(m)
		mapAny.EnableCut(".")
		val, ok := mapAny.get("nested.key2")
		assert.True(t, ok)
		assert.Equal(t, 42, val)
	})

	// 测试4: 启用cut功能，查找深层嵌套的key
	t.Run("deep_nested_with_cut", func(t *testing.T) {
		mapAny := NewMap(m)
		mapAny.EnableCut(".")
		val, ok := mapAny.get("nested.deep.key3")
		assert.True(t, ok)
		assert.Equal(t, true, val)
	})

	// 测试5: 启用cut功能，但嵌套路径中有非map类型
	t.Run("non_map_in_path", func(t *testing.T) {
		m := map[string]interface{}{
			"nested": "not_a_map",
		}
		mapAny := NewMap(m)
		mapAny.EnableCut(".")
		val, ok := mapAny.get("nested.key")
		assert.False(t, ok)
		assert.Nil(t, val)
	})

	// 测试6: 启用cut功能，使用不同的分隔符
	t.Run("different_separator", func(t *testing.T) {
		// 创建一个带有下划线分隔符的map
		m := map[string]interface{}{
			"nested_deep_key3": true,
		}
		mapAny := NewMap(m)
		mapAny.EnableCut("_")
		// 这个查找应该成功，因为我们使用了下划线作为分隔符
		val, ok := mapAny.get("nested_deep_key3")
		assert.True(t, ok)
		assert.Equal(t, true, val)
	})

	// 测试7: 启用cut功能，然后禁用cut功能
	t.Run("enable_then_disable_cut", func(t *testing.T) {
		mapAny := NewMap(m)
		mapAny.EnableCut(".")
		mapAny.DisableCut()
		// 禁用cut后，应该无法查找嵌套key
		val, ok := mapAny.get("nested.key2")
		assert.False(t, ok)
		assert.Nil(t, val)
	})

	// 测试8: 测试NewMapWithAny函数的完整覆盖
	t.Run("new_map_with_any", func(t *testing.T) {
		// 测试使用struct创建MapAny
		type TestStruct struct {
			Key1 string `json:"key1"`
			Key2 int    `json:"key2"`
		}

		testStruct := TestStruct{
			Key1: "value1",
			Key2: 42,
		}

		mapAny, err := NewMapWithAny(testStruct)
		assert.NoError(t, err)
		assert.NotNil(t, mapAny)

		val, err := mapAny.Get("key1")
		assert.NoError(t, err)
		assert.Equal(t, "value1", val)

		val, err = mapAny.Get("key2")
		assert.NoError(t, err)
		assert.Equal(t, 42, val)
	})

	// 测试9: 测试NewMapWithAny函数处理错误情况
	t.Run("new_map_with_any_error", func(t *testing.T) {
		// 测试使用不可序列化的类型
		ch := make(chan int)
		mapAny, err := NewMapWithAny(ch)
		assert.Error(t, err)
		assert.Nil(t, mapAny)
	})

	// 测试10: 测试get函数的空key情况
	t.Run("empty_key", func(t *testing.T) {
		mapAny := NewMap(m)
		val, ok := mapAny.get("")
		assert.False(t, ok)
		assert.Nil(t, val)
	})
}
