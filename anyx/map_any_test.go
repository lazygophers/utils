package anyx

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapAny_BasicOperations(t *testing.T) {
	t.Run("Set and Get", func(t *testing.T) {
		m := NewMap(nil)
		m.Set("name", "Alice")
		val, err := m.Get("name")
		assert.NoError(t, err)
		assert.Equal(t, "Alice", val)
	})

	t.Run("Exists", func(t *testing.T) {
		m := NewMap(nil)
		m.Set("age", 30)
		assert.True(t, m.Exists("age"))
		assert.False(t, m.Exists("invalid"))
	})

	t.Run("Keys and Values", func(t *testing.T) {
		m := NewMap(map[string]interface{}{
			"k1": "v1",
			"k2": 2,
		})
		keys := []string{}
		values := []interface{}{}
		m.data.Range(func(key, value interface{}) bool {
			keys = append(keys, key.(string))
			values = append(values, value)
			return true
		})
		assert.ElementsMatch(t, []string{"k1", "k2"}, keys)
		assert.ElementsMatch(t, []interface{}{"v1", 2}, values)
	})
}

func TestMapAny_Concurrency(t *testing.T) {
	m := NewMap(nil)
	var wg sync.WaitGroup

	// 并发写入
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			m.Set("key", idx)
		}(i)
	}

	// 并发读取
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Get("key")
		}()
	}

	wg.Wait()
	val, _ := m.Get("key")
	assert.NotNil(t, val)
}

func TestMapAny_BoundaryConditions(t *testing.T) {
	t.Run("Nil Value", func(t *testing.T) {
		m := NewMap(nil)
		m.Set("nil", nil)
		val, err := m.Get("nil")
		assert.NoError(t, err)
		assert.Nil(t, val)
	})

	t.Run("Empty Map", func(t *testing.T) {
		m := NewMap(nil)
		val, err := m.Get("missing")
		assert.ErrorIs(t, err, ErrNotFound)
		assert.Nil(t, val)
	})

	t.Run("Type Conversion", func(t *testing.T) {
		m := NewMap(map[string]interface{}{"num": "not_a_number"})
		assert.Equal(t, 0, m.GetInt("num"))
		assert.Equal(t, "", m.GetString("invalid"))
	})
}

func TestMapAny_TypeAccessors(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"bool":   true,
		"int":    42,
		"string": "hello",
		"float":  3.14,
	})

	tests := []struct {
		key      string
		expected interface{}
	}{
		{"bool", true},
		{"int", 42},
		{"string", "hello"},
		{"float", 3.14},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			switch tt.expected.(type) {
			case bool:
				assert.Equal(t, tt.expected, m.GetBool(tt.key))
			case int:
				assert.Equal(t, tt.expected, m.GetInt(tt.key))
			case string:
				assert.Equal(t, tt.expected, m.GetString(tt.key))
			case float64:
				assert.Equal(t, tt.expected, m.GetFloat64(tt.key))
			}
		})
	}
}

func TestMapAny_Range(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"a": 1,
		"b": "hello",
		"c": true,
	})

	keys := make(map[string]interface{})
	count := 0

	m.Range(func(key, value interface{}) bool {
		keys[key.(string)] = value
		count++
		return true
	})

	assert.Equal(t, 3, count)
	assert.Equal(t, 1, keys["a"])
	assert.Equal(t, "hello", keys["b"])
	assert.Equal(t, true, keys["c"])

	t.Run("Stop Range", func(t *testing.T) {
		innerCount := 0
		m.Range(func(key, value interface{}) bool {
			innerCount++
			return false // Stop after first item
		})
		assert.Equal(t, 1, innerCount)
	})
}


func TestNewMapWithJson(t *testing.T) {
	// 测试有效的 JSON
	jsonData := []byte(`{"name": "张三", "age": 25, "active": true}`)
	m, err := NewMapWithJson(jsonData)
	require.NoError(t, err)
	require.NotNil(t, m)
	
	// 验证数据
	require.Equal(t, "张三", m.GetString("name"))
	require.Equal(t, 25, m.GetInt("age"))
	require.True(t, m.GetBool("active"))
	
	// 测试无效的 JSON
	invalidJson := []byte(`{"name": "张三",`)
	_, err = NewMapWithJson(invalidJson)
	require.Error(t, err)
	
	// 测试空 JSON
	m, err = NewMapWithJson([]byte("{}"))
	require.NoError(t, err)
	require.NotNil(t, m)
	require.Equal(t, 0, len(m.ToMap()))
}

func TestNewMapWithYaml(t *testing.T) {
	// 测试有效的 YAML
	yamlData := []byte("name: 张三\nage: 25\nactive: true\n")
	m, err := NewMapWithYaml(yamlData)
	require.NoError(t, err)
	require.NotNil(t, m)
	
	// 验证数据
	require.Equal(t, "张三", m.GetString("name"))
	require.Equal(t, 25, m.GetInt("age"))
	require.True(t, m.GetBool("active"))
	
	// 测试无效的 YAML
	invalidYaml := []byte("name: 张三\nage: 25\ninvalid: [")
	_, err = NewMapWithYaml(invalidYaml)
	require.Error(t, err)
	
	// 测试空 YAML
	m, err = NewMapWithYaml([]byte("{}"))
	require.NoError(t, err)
	require.NotNil(t, m)
	require.Equal(t, 0, len(m.ToMap()))
}

func TestNewMapWithAny(t *testing.T) {
	// 测试结构体
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	person := Person{Name: "张三", Age: 25}
	
	m, err := NewMapWithAny(person)
	require.NoError(t, err)
	require.NotNil(t, m)
	
	// 验证数据
	require.Equal(t, "张三", m.GetString("name"))
	require.Equal(t, 25, m.GetInt("age"))
	
	// 测试 map
	data := map[string]interface{}{
		"name": "李四",
		"age":  30,
	}
	m, err = NewMapWithAny(data)
	require.NoError(t, err)
	require.NotNil(t, m)
	require.Equal(t, "李四", m.GetString("name"))
	require.Equal(t, 30, m.GetInt("age"))
	
	// 测试无法序列化的类型
	_, err = NewMapWithAny(func() {})
	require.Error(t, err)
}

func TestMapAny_EnableCut(t *testing.T) {
	m := NewMap(nil)
	
	// 测试启用 cut
	m2 := m.EnableCut(".")
	assert.Same(t, m, m2) // 应该返回同一个实例
	assert.True(t, m.cut.Load())
	assert.Equal(t, ".", m.seq.Load())
	
	// 测试链式调用
	m3 := m.EnableCut(":").EnableCut("|")
	assert.True(t, m3.cut.Load())
	assert.Equal(t, "|", m3.seq.Load())
}

func TestMapAny_DisableCut(t *testing.T) {
	m := NewMap(nil).EnableCut(".")
	
	// 测试禁用 cut
	m2 := m.DisableCut()
	assert.Same(t, m, m2) // 应该返回同一个实例
	assert.False(t, m2.cut.Load())
}

func TestMapAny_GetInt32(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"int32":  int32(42),
		"int64":  int64(100),
		"string": "123",
		"float":  3.14,
	})
	
	assert.Equal(t, int32(42), m.GetInt32("int32"))
	assert.Equal(t, int32(100), m.GetInt32("int64"))
	assert.Equal(t, int32(123), m.GetInt32("string"))
	assert.Equal(t, int32(0), m.GetInt32("missing"))
}

func TestMapAny_GetInt64(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"int32":  int32(42),
		"int64":  int64(100),
		"string": "123",
		"float":  3.14,
	})
	
	assert.Equal(t, int64(42), m.GetInt64("int32"))
	assert.Equal(t, int64(100), m.GetInt64("int64"))
	assert.Equal(t, int64(123), m.GetInt64("string"))
	assert.Equal(t, int64(0), m.GetInt64("missing"))
}

func TestMapAny_GetUint16(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"uint16": uint16(42),
		"int":    -100,
		"string": "123",
		"float":  3.14,
	})
	
	assert.Equal(t, uint16(42), m.GetUint16("uint16"))
	assert.Equal(t, uint16(65436), m.GetUint16("int")) // 负数转换为 uint16 会溢出
	assert.Equal(t, uint16(123), m.GetUint16("string"))
	assert.Equal(t, uint16(0), m.GetUint16("missing"))
}

func TestMapAny_GetUint32(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"uint32": uint32(42),
		"int":    -100,
		"string": "123",
		"float":  3.14,
	})
	
	assert.Equal(t, uint32(42), m.GetUint32("uint32"))
	assert.Equal(t, uint32(0xffffff9c), m.GetUint32("int")) // 负数转换为 uint32 会溢出
	assert.Equal(t, uint32(123), m.GetUint32("string"))
	assert.Equal(t, uint32(0), m.GetUint32("missing"))
}

func TestMapAny_GetUint64(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"uint64": uint64(42),
		"int":    -100,
		"string": "123",
		"float":  3.14,
	})
	
	assert.Equal(t, uint64(42), m.GetUint64("uint64"))
	assert.Equal(t, uint64(18446744073709551516), m.GetUint64("int")) // 负数转换为 uint64 会溢出
	assert.Equal(t, uint64(123), m.GetUint64("string"))
	assert.Equal(t, uint64(0), m.GetUint64("missing"))
}

func TestMapAny_GetBytes(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"bool":     true,
		"int":      42,
		"string":   "hello",
		"bytes":    []byte("world"),
		"struct":   struct{ Field int }{Field: 1},
	})
	
	assert.Equal(t, []byte("1"), m.GetBytes("bool"))
	assert.Equal(t, []byte("42"), m.GetBytes("int"))
	assert.Equal(t, []byte("hello"), m.GetBytes("string"))
	assert.Equal(t, []byte("world"), m.GetBytes("bytes"))
	assert.Equal(t, []byte(""), m.GetBytes("struct"))
	assert.Equal(t, []byte(""), m.GetBytes("missing"))
}

func TestMapAny_GetMap(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"map": map[string]interface{}{
			"key": "value",
		},
		"string":   `{"key": "value"}`,
		"bytes":    []byte(`{"key": "value"}`),
		"number":   42,
		"missing":  nil,
	})
	
	// 测试获取 map
	result := m.GetMap("map")
	assert.Equal(t, "value", result.GetString("key"))
	
	// 测试从字符串解析
	result = m.GetMap("string")
	assert.Equal(t, "value", result.GetString("key"))
	
	// 测试从字节解析
	result = m.GetMap("bytes")
	assert.Equal(t, "value", result.GetString("key"))
	
	// 测试数字类型
	result = m.GetMap("number")
	assert.Equal(t, 0, len(result.ToMap()))
	
	// 测试缺失的键
	result = m.GetMap("missing")
	assert.Equal(t, 0, len(result.ToMap()))
}

func TestMapAny_GetSlice(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"boolSlice":   []bool{true, false},
		"intSlice":    []int{1, 2, 3},
		"interfaceSlice": []interface{}{1, "two", true},
		"string":      "not a slice",
		"missing":     nil,
	})
	
	// 测试 bool 切片
	result := m.GetSlice("boolSlice")
	assert.Equal(t, []interface{}{true, false}, result)
	
	// 测试 int 切片
	result = m.GetSlice("intSlice")
	assert.Equal(t, []interface{}{1, 2, 3}, result)
	
	// 测试 interface 切片
	result = m.GetSlice("interfaceSlice")
	assert.Equal(t, []interface{}{1, "two", true}, result)
	
	// 测试非切片类型
	result = m.GetSlice("string")
	assert.Empty(t, result)
	
	// 测试缺失的键
	result = m.GetSlice("missing")
	assert.Empty(t, result)
}

func TestMapAny_GetStringSlice(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"stringSlice": []string{"a", "b", "c"},
		"intSlice":    []int{1, 2, 3},
		"interfaceSlice": []interface{}{"one", 2, true},
		"string":      "not a slice",
		"missing":     nil,
	})
	
	// 测试字符串切片
	result := m.GetStringSlice("stringSlice")
	assert.Equal(t, []string{"a", "b", "c"}, result)
	
	// 测试 int 切片转换
	result = m.GetStringSlice("intSlice")
	assert.Equal(t, []string{"1", "2", "3"}, result)
	
	// 测试 interface 切片转换
	result = m.GetStringSlice("interfaceSlice")
	assert.Equal(t, []string{"one", "2", "1"}, result)
	
	// 测试非切片类型
	result = m.GetStringSlice("string")
	assert.Equal(t, []string{}, result)
	
	// 测试缺失的键
	result = m.GetStringSlice("missing")
	assert.Equal(t, []string{}, result)
}

func TestMapAny_GetUint64Slice(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"uint64Slice": []uint64{1, 2, 3},
		"intSlice":     []int{-1, 0, 1},
		"stringSlice":  []string{"1", "2", "3"},
		"interfaceSlice": []interface{}{1, "2", "3"},
		"string":       "not a slice",
		"missing":      nil,
	})
	
	// 测试 uint64 切片
	result := m.GetUint64Slice("uint64Slice")
	assert.Equal(t, []uint64{1, 2, 3}, result)
	
	// 测试 int 切片转换（负数会变成大的正数）
	result = m.GetUint64Slice("intSlice")
	assert.Equal(t, []uint64{0xffffffffffffffff, 0, 1}, result)
	
	// 测试字符串切片转换
	result = m.GetUint64Slice("stringSlice")
	assert.Equal(t, []uint64{1, 2, 3}, result)
	
	// 测试非切片类型
	result = m.GetUint64Slice("string")
	assert.Equal(t, []uint64{}, result)
	
	// 测试缺失的键
	result = m.GetUint64Slice("missing")
	assert.Equal(t, []uint64{}, result)
}

func TestMapAny_GetInt64Slice(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"int64Slice":    []int64{1, 2, 3},
		"intSlice":      []int{1, 2, 3},
		"stringSlice":   []string{"1", "2", "3"},
		"interfaceSlice": []interface{}{1, 2, 3},
		"string":        "not a slice",
		"missing":       nil,
	})
	
	// 测试 int64 切片
	result := m.GetInt64Slice("int64Slice")
	assert.Equal(t, []int64{1, 2, 3}, result)
	
	// 测试 int 切片转换
	result = m.GetInt64Slice("intSlice")
	assert.Equal(t, []int64{1, 2, 3}, result)
	
	// 测试字符串切片转换
	result = m.GetInt64Slice("stringSlice")
	assert.Equal(t, []int64{1, 2, 3}, result)
	
	// 测试非切片类型
	result = m.GetInt64Slice("string")
	assert.Empty(t, result)
	
	// 测试缺失的键
	result = m.GetInt64Slice("missing")
	assert.Equal(t, []int64{}, result)
}

func TestMapAny_GetUint32Slice(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"uint32Slice": []uint32{1, 2, 3},
		"intSlice":     []int{-1, 0, 1},
		"stringSlice":  []string{"1", "2", "3"},
		"interfaceSlice": []interface{}{1, 2, 3},
		"string":       "not a slice",
		"missing":      nil,
	})
	
	// 测试 uint32 切片
	result := m.GetUint32Slice("uint32Slice")
	assert.Equal(t, []uint32{1, 2, 3}, result)
	
	// 测试 int 切片转换（负数会变成大的正数）
	result = m.GetUint32Slice("intSlice")
	assert.Equal(t, []uint32{0xffffffff, 0, 1}, result)
	
	// 测试字符串切片转换
	result = m.GetUint32Slice("stringSlice")
	assert.Equal(t, []uint32{1, 2, 3}, result)
	
	// 测试非切片类型
	result = m.GetUint32Slice("string")
	assert.Equal(t, []uint32{}, result)
	
	// 测试缺失的键
	result = m.GetUint32Slice("missing")
	assert.Equal(t, []uint32{}, result)
}

func TestMapAny_ToSyncMap(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	})
	
	syncMap := m.ToSyncMap()
	assert.NotNil(t, syncMap)
	
	// 验证数据
	val, ok := syncMap.Load("key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", val)
	
	val, ok = syncMap.Load("key2")
	assert.True(t, ok)
	assert.Equal(t, 42, val)
}

func TestMapAny_ToMap(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"string": "hello",
		"int":    42,
		"float32": 3.0,
		"float64": 3.14,
		"nested": NewMap(map[string]interface{}{
			"inner": "value",
		}),
	})
	
	result := m.ToMap()
	assert.Equal(t, "hello", result["string"])
	assert.Equal(t, 42, result["int"])
	assert.Equal(t, int64(3), result["float32"]) // 整数 float32 转为 int64
	assert.Equal(t, 3.14, result["float64"])
	assert.Equal(t, map[string]interface{}{"inner": "value"}, result["nested"])
}

func TestMapAny_Clone(t *testing.T) {
	original := NewMap(map[string]interface{}{
		"key1": "value1",
		"key2": 42,
	}).EnableCut(".")
	
	clone := original.Clone()
	
	// 验证数据相同
	assert.Equal(t, "value1", clone.GetString("key1"))
	assert.Equal(t, 42, clone.GetInt("key2"))
	
	// 验证 cut 设置被复制
	assert.True(t, clone.cut.Load())
	assert.Equal(t, ".", clone.seq.Load())
	
	// 验证是不同的实例
	assert.NotSame(t, original, clone)
	
	// 修改 clone 不影响 original
	clone.Set("key1", "modified")
	assert.Equal(t, "value1", original.GetString("key1"))
	assert.Equal(t, "modified", clone.GetString("key1"))
}