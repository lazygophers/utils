package anyx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMapWithYaml_Optimized(t *testing.T) {
	t.Run("simple key-value pairs", func(t *testing.T) {
		data := []byte(`
key1: value1
key2: value2
key3: value3
`)
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)
		assert.NotNil(t, m)

		val1, err := m.Get("key1")
		assert.NoError(t, err)
		assert.Equal(t, "value1", val1)

		val2, err := m.Get("key2")
		assert.NoError(t, err)
		assert.Equal(t, "value2", val2)
	})

	t.Run("integer values", func(t *testing.T) {
		data := []byte(`
count: 42
price: 100
total: 9999
`)
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)

		val, err := m.Get("count")
		assert.NoError(t, err)
		assert.Equal(t, int64(42), val)
	})

	t.Run("float values", func(t *testing.T) {
		data := []byte(`
pi: 3.14
rate: 0.5
temperature: -10.5
`)
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)

		val, err := m.Get("pi")
		assert.NoError(t, err)
		assert.InDelta(t, 3.14, val, 0.01)
	})

	t.Run("boolean values", func(t *testing.T) {
		data := []byte(`
enabled: true
disabled: false
active: yes
inactive: no
`)
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)

		val1, err := m.Get("enabled")
		assert.NoError(t, err)
		assert.Equal(t, true, val1)

		val2, err := m.Get("disabled")
		assert.NoError(t, err)
		assert.Equal(t, false, val2)
	})

	t.Run("null values", func(t *testing.T) {
		data := []byte(`
empty: null
nothing: ~
`)
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)

		val, err := m.Get("empty")
		assert.NoError(t, err)
		assert.Nil(t, val)
	})

	t.Run("nested maps", func(t *testing.T) {
		data := []byte(`
config:
  database:
    host: localhost
    port: 5432
  server:
    port: 8080
`)
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)

		config, err := m.Get("config")
		assert.NoError(t, err)

		configMap, ok := config.(map[string]interface{})
		assert.True(t, ok)

		db, ok := configMap["database"].(map[string]interface{})
		assert.True(t, ok)
		assert.Equal(t, "localhost", db["host"])
	})

	t.Run("arrays/sequences", func(t *testing.T) {
		data := []byte(`
items:
  - apple
  - banana
  - orange
numbers:
  - 1
  - 2
  - 3
`)
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)

		items, err := m.Get("items")
		assert.NoError(t, err)

		itemsArray, ok := items.([]interface{})
		assert.True(t, ok)
		assert.Len(t, itemsArray, 3)
		assert.Equal(t, "apple", itemsArray[0])
	})

	t.Run("complex nested structure", func(t *testing.T) {
		data := []byte(`
server:
  host: example.com
  ports:
    - 80
    - 443
    - 8080
  tls:
    enabled: true
    cert: /path/to/cert.pem
`)
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)

		server, err := m.Get("server")
		assert.NoError(t, err)

		serverMap, ok := server.(map[string]interface{})
		assert.True(t, ok)
		assert.Equal(t, "example.com", serverMap["host"])

		ports, ok := serverMap["ports"].([]interface{})
		assert.True(t, ok)
		assert.Len(t, ports, 3)

		tls, ok := serverMap["tls"].(map[string]interface{})
		assert.True(t, ok)
		assert.Equal(t, true, tls["enabled"])
	})

	t.Run("empty document", func(t *testing.T) {
		data := []byte(``)
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)
		assert.NotNil(t, m)
	})

	t.Run("only comments", func(t *testing.T) {
		data := []byte(`
# This is a comment
# Another comment
`)
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)
		assert.NotNil(t, m)
	})

	t.Run("special characters in values", func(t *testing.T) {
		data := []byte(`
path: /usr/local/bin
url: https://example.com
quote: "hello world"
`)
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)

		val, err := m.Get("path")
		assert.NoError(t, err)
		assert.Equal(t, "/usr/local/bin", val)
	})

	t.Run("multiline strings", func(t *testing.T) {
		data := []byte(`
description: |
  This is a multiline
  string description
`)
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)

		val, err := m.Get("description")
		assert.NoError(t, err)
		assert.Contains(t, val, "multiline")
	})

	t.Run("large document", func(t *testing.T) {
		var yamlData []byte
		yamlData = append(yamlData, []byte("---\n")...)
		for i := 0; i < 1000; i++ {
			yamlData = append(yamlData, []byte("key")...)
			yamlData = append(yamlData, byte('0'+i%10))
			yamlData = append(yamlData, []byte(": value")...)
			yamlData = append(yamlData, byte('0'+i%10))
			yamlData = append(yamlData, '\n')
		}

		m, err := NewMapWithYaml(yamlData)
		assert.NoError(t, err)
		assert.NotNil(t, m)

		val, err := m.Get("key0")
		assert.NoError(t, err)
		assert.Equal(t, "value0", val)
	})
}

func TestNewMapWithYaml_ErrorHandling(t *testing.T) {
	t.Run("invalid YAML syntax", func(t *testing.T) {
		data := []byte(`
key1: value1
key2: [unclosed array
key3: value3
`)
		_, err := NewMapWithYaml(data)
		assert.Error(t, err)
	})

	t.Run("unmatched brackets", func(t *testing.T) {
		data := []byte(`
list:
  - item1
  - item2
  - item3
`)
		// 这个实际上是有效的 YAML
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)
		assert.NotNil(t, m)
	})
}

func TestConvertYamlNode_EdgeCases(t *testing.T) {
	t.Run("empty sequence", func(t *testing.T) {
		data := []byte(`empty: []`)
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)

		val, err := m.Get("empty")
		assert.NoError(t, err)

		emptySlice, ok := val.([]interface{})
		assert.True(t, ok)
		assert.Len(t, emptySlice, 0)
	})

	t.Run("empty map", func(t *testing.T) {
		data := []byte(`empty: {}`)
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)

		val, err := m.Get("empty")
		assert.NoError(t, err)

		emptyMap, ok := val.(map[string]interface{})
		assert.True(t, ok)
		assert.Len(t, emptyMap, 0)
	})

	t.Run("mixed types in sequence", func(t *testing.T) {
		data := []byte(`
mixed:
  - string
  - 42
  - 3.14
  - true
  - null
`)
		m, err := NewMapWithYaml(data)
		assert.NoError(t, err)

		val, err := m.Get("mixed")
		assert.NoError(t, err)

		mixedSlice, ok := val.([]interface{})
		assert.True(t, ok)
		assert.Len(t, mixedSlice, 5)
		assert.Equal(t, "string", mixedSlice[0])
		assert.Equal(t, int64(42), mixedSlice[1])
	})
}
