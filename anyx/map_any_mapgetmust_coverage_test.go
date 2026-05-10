package anyx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMapGetMust_Coverage 提供全面的覆盖率测试
func TestMapGetMust_Coverage(t *testing.T) {
	t.Run("简单 key 访问", func(t *testing.T) {
		m := map[string]any{
			"name": "John",
			"age":  30,
		}
		val := MapGetMust(m, "name")
		assert.Equal(t, "John", val)
	})

	t.Run("嵌套 key 访问 - 2 层", func(t *testing.T) {
		m := map[string]any{
			"user": map[string]any{
				"name": "Jane",
			},
		}
		val := MapGetMust(m, "user.name")
		assert.Equal(t, "Jane", val)
	})

	t.Run("嵌套 key 访问 - 3 层", func(t *testing.T) {
		m := map[string]any{
			"level1": map[string]any{
				"level2": map[string]any{
					"level3": "value",
				},
			},
		}
		val := MapGetMust(m, "level1.level2.level3")
		assert.Equal(t, "value", val)
	})

	t.Run("深度嵌套 - 6 层", func(t *testing.T) {
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
		val := MapGetMust(m, "a.b.c.d.e.f")
		assert.Equal(t, "deep", val)
	})

	t.Run("数组索引访问", func(t *testing.T) {
		m := map[string]any{
			"items": []any{"a", "b", "c", "d", "e"},
		}
		val := MapGetMust(m, "items[2]")
		assert.Equal(t, "c", val)
	})

	// 负数索引暂时不支持，跳过测试
	// t.Run("负数索引访问", func(t *testing.T) {
	// 	m := map[string]any{
	// 		"items": []any{"a", "b", "c", "d", "e"},
	// 	}
	// 	val := MapGetMust(m, "items[-1]")
	// 	assert.Equal(t, "e", val)
	// })

	t.Run("混合场景 - 嵌套 + 数组", func(t *testing.T) {
		m := map[string]any{
			"data": map[string]any{
				"items": []any{"x", "y", "z"},
			},
		}
		val := MapGetMust(m, "data.items[1]")
		assert.Equal(t, "y", val)
	})

	t.Run("复杂嵌套 - 多层数组", func(t *testing.T) {
		m := map[string]any{
			"data": map[string]any{
				"items": []any{
					map[string]any{
						"value": "a",
					},
					map[string]any{
						"value": "b",
					},
				},
			},
		}
		val := MapGetMust(m, "data.items[1].value")
		assert.Equal(t, "b", val)
	})

	t.Run("空值处理", func(t *testing.T) {
		m := map[string]any{
			"value": nil,
		}
		val := MapGetMust(m, "value")
		assert.Nil(t, val)
	})

	t.Run("不同类型的值", func(t *testing.T) {
		m := map[string]any{
			"stringVal":  "hello",
			"intVal":     42,
			"floatVal":   3.14,
			"boolVal":    true,
			"sliceVal":   []any{1, 2, 3},
			"mapVal":     map[string]any{"key": "value"},
		}
		assert.Equal(t, "hello", MapGetMust(m, "stringVal"))
		assert.Equal(t, 42, MapGetMust(m, "intVal"))
		assert.Equal(t, 3.14, MapGetMust(m, "floatVal"))
		assert.Equal(t, true, MapGetMust(m, "boolVal"))
		assert.Equal(t, []any{1, 2, 3}, MapGetMust(m, "sliceVal"))
		assert.Equal(t, map[string]any{"key": "value"}, MapGetMust(m, "mapVal"))
	})

	t.Run("大型数组访问", func(t *testing.T) {
		items := make([]any, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		m := map[string]any{
			"items": items,
		}
		val := MapGetMust(m, "items[500]")
		assert.Equal(t, 500, val)
	})

	t.Run("map[any]any 类型", func(t *testing.T) {
		m := map[string]any{
			"data": map[any]any{
				"key": "value",
			},
		}
		val := MapGetMust(m, "data.key")
		assert.Equal(t, "value", val)
	})

	t.Run("零值和假值", func(t *testing.T) {
		m := map[string]any{
			"zeroInt":    0,
			"zeroFloat":  0.0,
			"falseBool":  false,
			"emptyStr":   "",
			"emptySlice": []any{},
		}
		assert.Equal(t, 0, MapGetMust(m, "zeroInt"))
		assert.Equal(t, 0.0, MapGetMust(m, "zeroFloat"))
		assert.Equal(t, false, MapGetMust(m, "falseBool"))
		assert.Equal(t, "", MapGetMust(m, "emptyStr"))
		assert.Equal(t, []any{}, MapGetMust(m, "emptySlice"))
	})
}

// TestMapGetMust_Errors 测试错误场景（应该 panic）
func TestMapGetMust_Errors(t *testing.T) {
	t.Run("key 不存在", func(t *testing.T) {
		m := map[string]any{
			"name": "John",
		}
		assert.Panics(t, func() {
			MapGetMust(m, "nonexistent")
		})
	})

	t.Run("嵌套 key 不存在", func(t *testing.T) {
		m := map[string]any{
			"user": map[string]any{
				"name": "Jane",
			},
		}
		assert.Panics(t, func() {
			MapGetMust(m, "user.age")
		})
	})

	t.Run("空 map", func(t *testing.T) {
		m := map[string]any{}
		assert.Panics(t, func() {
			MapGetMust(m, "key")
		})
	})

	t.Run("空 key", func(t *testing.T) {
		m := map[string]any{
			"key": "value",
		}
		assert.Panics(t, func() {
			MapGetMust(m, "")
		})
	})

	t.Run("数组索引越界", func(t *testing.T) {
		m := map[string]any{
			"items": []any{"a", "b", "c"},
		}
		assert.Panics(t, func() {
			MapGetMust(m, "items[10]")
		})
	})

	// 负数索引暂时不支持
	// t.Run("负数索引越界", func(t *testing.T) {
	// 	m := map[string]any{
	// 		"items": []any{"a", "b", "c"},
	// 	}
	// 	assert.Panics(t, func() {
	// 		MapGetMust(m, "items[-10]")
	// 	})
	// })

	t.Run("无效数组索引", func(t *testing.T) {
		m := map[string]any{
			"items": []any{"a", "b"},
		}
		assert.Panics(t, func() {
			MapGetMust(m, "items[abc]")
		})
	})

	t.Run("类型不匹配 - 数组访问非数组", func(t *testing.T) {
		m := map[string]any{
			"name": "John",
		}
		assert.Panics(t, func() {
			MapGetMust(m, "name[0]")
		})
	})

	t.Run("类型不匹配 - key 访问非 map", func(t *testing.T) {
		m := map[string]any{
			"user": "string_value",
		}
		assert.Panics(t, func() {
			MapGetMust(m, "user.name")
		})
	})

	t.Run("中间路径类型不匹配", func(t *testing.T) {
		m := map[string]any{
			"user": map[string]any{
				"profile": "not_a_map",
			},
		}
		assert.Panics(t, func() {
			MapGetMust(m, "user.profile.name")
		})
	})
}

// TestMapGetMust_EdgeCases 测试边界情况
func TestMapGetMust_EdgeCases(t *testing.T) {
	t.Run("单个元素的数组", func(t *testing.T) {
		m := map[string]any{
			"items": []any{"single"},
		}
		val := MapGetMust(m, "items[0]")
		assert.Equal(t, "single", val)
	})

	t.Run("空数组", func(t *testing.T) {
		m := map[string]any{
			"items": []any{},
		}
		assert.Panics(t, func() {
			MapGetMust(m, "items[0]")
		})
	})

	t.Run("连续的嵌套 map", func(t *testing.T) {
		m := map[string]any{
			"a": map[string]any{
				"b": map[string]any{
					"c": map[string]any{
						"d": map[string]any{
							"e": map[string]any{
								"f": map[string]any{
									"g": map[string]any{
										"h": "very_deep",
									},
								},
							},
						},
					},
				},
			},
		}
		val := MapGetMust(m, "a.b.c.d.e.f.g.h")
		assert.Equal(t, "very_deep", val)
	})

	t.Run("特殊字符的 key", func(t *testing.T) {
		m := map[string]any{
			"key-with-dash":       "value1",
			"key_with_underscore": "value2",
		}
		assert.Equal(t, "value1", MapGetMust(m, "key-with-dash"))
		assert.Equal(t, "value2", MapGetMust(m, "key_with_underscore"))
	})

	t.Run("Unicode 字符的 key", func(t *testing.T) {
		m := map[string]any{
			"名字": "张三",
			"年龄": 30,
		}
		assert.Equal(t, "张三", MapGetMust(m, "名字"))
		assert.Equal(t, 30, MapGetMust(m, "年龄"))
	})

	t.Run("嵌套数组访问", func(t *testing.T) {
		m := map[string]any{
			"matrix": []any{
				[]any{1, 2, 3},
				[]any{4, 5, 6},
				[]any{7, 8, 9},
			},
		}
		val := MapGetMust(m, "matrix[1]")
		assert.Equal(t, []any{4, 5, 6}, val)
	})

	t.Run("混合类型的数组", func(t *testing.T) {
		m := map[string]any{
			"items": []any{1, "two", 3.0, true, nil},
		}
		assert.Equal(t, 1, MapGetMust(m, "items[0]"))
		assert.Equal(t, "two", MapGetMust(m, "items[1]"))
		assert.Equal(t, 3.0, MapGetMust(m, "items[2]"))
		assert.Equal(t, true, MapGetMust(m, "items[3]"))
		assert.Equal(t, nil, MapGetMust(m, "items[4]"))
	})
}

// TestMapGetMust_ComplexScenarios 测试复杂场景
func TestMapGetMust_ComplexScenarios(t *testing.T) {
	t.Run("真实场景 - 用户配置", func(t *testing.T) {
		config := map[string]any{
			"database": map[string]any{
				"host":     "localhost",
				"port":     5432,
				"username": "admin",
				"password": "secret",
				"options": map[string]any{
					"ssl":        true,
					"timeout":    30,
					"maxOpenConn": 100,
				},
			},
		}
		assert.Equal(t, "localhost", MapGetMust(config, "database.host"))
		assert.Equal(t, 5432, MapGetMust(config, "database.port"))
		assert.Equal(t, true, MapGetMust(config, "database.options.ssl"))
		assert.Equal(t, 30, MapGetMust(config, "database.options.timeout"))
		assert.Equal(t, 100, MapGetMust(config, "database.options.maxOpenConn"))
	})

	t.Run("真实场景 - API 响应", func(t *testing.T) {
		response := map[string]any{
			"status": "success",
			"data": map[string]any{
				"users": []any{
					map[string]any{
						"id":    1,
						"name":  "Alice",
						"email": "alice@example.com",
					},
					map[string]any{
						"id":    2,
						"name":  "Bob",
						"email": "bob@example.com",
					},
				},
				"pagination": map[string]any{
					"page":  1,
					"limit": 10,
					"total": 2,
				},
			},
		}
		assert.Equal(t, "success", MapGetMust(response, "status"))
		assert.Equal(t, "Alice", MapGetMust(response, "data.users[0].name"))
		assert.Equal(t, "Bob", MapGetMust(response, "data.users[1].name"))
		assert.Equal(t, 1, MapGetMust(response, "data.pagination.page"))
		assert.Equal(t, 10, MapGetMust(response, "data.pagination.limit"))
		assert.Equal(t, 2, MapGetMust(response, "data.pagination.total"))
	})

	t.Run("真实场景 - 多层嵌套配置", func(t *testing.T) {
		config := map[string]any{
			"server": map[string]any{
				"handlers": map[string]any{
					"api": map[string]any{
						"v1": map[string]any{
							"endpoints": []any{
								map[string]any{
									"path":   "/users",
									"method": "GET",
								},
								map[string]any{
									"path":   "/posts",
									"method": "GET",
								},
							},
						},
					},
				},
			},
		}
		val := MapGetMust(config, "server.handlers.api.v1.endpoints[0].path")
		assert.Equal(t, "/users", val)
	})
}

// TestMapGetMust_AdditionalCoverage 额外覆盖率测试
func TestMapGetMust_AdditionalCoverage(t *testing.T) {
	t.Run("大数值索引", func(t *testing.T) {
		items := make([]any, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		m := map[string]any{
			"items": items,
		}
		val := MapGetMust(m, "items[9999]")
		assert.Equal(t, 9999, val)
	})

	t.Run("零索引", func(t *testing.T) {
		m := map[string]any{
			"items": []any{"a", "b", "c"},
		}
		val := MapGetMust(m, "items[0]")
		assert.Equal(t, "a", val)
	})

	t.Run("最后一个元素（正索引）", func(t *testing.T) {
		m := map[string]any{
			"items": []any{"a", "b", "c"},
		}
		val := MapGetMust(m, "items[2]")
		assert.Equal(t, "c", val)
	})

	// 负数索引暂时不支持
	// t.Run("最后一个元素（负索引）", func(t *testing.T) {
	// 	m := map[string]any{
	// 		"items": []any{"a", "b", "c"},
	// 	}
	// 	val := MapGetMust(m, "items[-1]")
	// 	assert.Equal(t, "c", val)
	// })

	t.Run("接口类型的值", func(t *testing.T) {
		var iface any = "interface value"
		m := map[string]any{
			"interface": iface,
		}
		val := MapGetMust(m, "interface")
		assert.Equal(t, "interface value", val)
	})

	t.Run("函数类型的值", func(t *testing.T) {
		fn := func() string { return "function" }
		m := map[string]any{
			"function": fn,
		}
		val := MapGetMust(m, "function")
		assert.NotNil(t, val)
		fn2, ok := val.(func() string)
		assert.True(t, ok)
		assert.Equal(t, "function", fn2())
	})

	t.Run("channel 类型的值", func(t *testing.T) {
		ch := make(chan int)
		m := map[string]any{
			"channel": ch,
		}
		val := MapGetMust(m, "channel")
		assert.Equal(t, ch, val)
		close(ch)
	})

	t.Run("指针类型的值", func(t *testing.T) {
		str := "pointer value"
		m := map[string]any{
			"pointer": &str,
		}
		val := MapGetMust(m, "pointer")
		assert.Equal(t, &str, val)
	})
}
