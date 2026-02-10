package anyx

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/lazygophers/log"
	"github.com/lazygophers/utils/candy"
	"github.com/lazygophers/utils/json"
	"go.uber.org/atomic"
	"gopkg.in/yaml.v3"
)

type MapAny struct {
	data *sync.Map
	cut  *atomic.Bool
	seq  *atomic.String
}

var (
	ErrNotFound = errors.New("not found")
)

func NewMap(m map[string]interface{}) *MapAny {
	m2 := &MapAny{
		data: &sync.Map{},
		cut:  atomic.NewBool(false),
		seq:  atomic.NewString(""),
	}
	for k, v := range m {
		m2.data.Store(k, v)
	}
	return m2
}

func NewMapWithJson(s []byte) (*MapAny, error) {
	var m map[string]interface{}
	err := json.Unmarshal(s, &m)
	if err != nil {
		return nil, err
	}
	return NewMap(m), nil
}

func NewMapWithYaml(s []byte) (*MapAny, error) {
	var m map[string]interface{}
	err := yaml.Unmarshal(s, &m)
	if err != nil {
		return nil, err
	}
	return NewMap(m), nil
}

func NewMapWithAny(s interface{}) (*MapAny, error) {
	buf, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		return nil, err
	}
	return NewMap(m), nil
}

func (p *MapAny) EnableCut(seq string) *MapAny {
	p.cut.Store(true)
	p.seq.Store(seq)
	return p
}

func (p *MapAny) DisableCut() *MapAny {
	p.cut.Store(false)
	return p
}

func (p *MapAny) Set(key string, value interface{}) {
	p.data.Store(key, value)
}

func (p *MapAny) Get(key string) (interface{}, error) {
	val, ok := p.get(key)
	if !ok {
		return nil, ErrNotFound
	}

	return val, nil
}

func (p *MapAny) get(key string) (interface{}, bool) {
	var val interface{}
	var ok bool

	if val, ok = p.data.Load(key); ok {
		return val, true
	}
	if !p.cut.Load() {
		return nil, false
	}

	seq := p.seq.Load()
	keys := strings.Split(key, seq)

	data := p.data
	var m *MapAny
	for len(keys) > 1 {
		k := keys[0]
		keys = keys[1:]

		val, ok = data.Load(k)
		if !ok {
			return nil, false
		}

		m = p.toMap(val)
		if m == nil {
			return nil, false
		}

		data = m.data
	}

	if len(keys) > 0 {
		if val, ok = data.Load(keys[0]); ok {
			return val, true
		}
	}

	return nil, false
}

func (p *MapAny) Exists(key string) bool {
	_, ok := p.get(key)
	if !ok {
		return false
	}

	return true
}

func (p *MapAny) GetBool(key string) bool {
	val, ok := p.get(key)
	if !ok {
		return false
	}

	return candy.ToBool(val)
}

func (p *MapAny) GetInt(key string) int {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return candy.ToInt(val)
}

func (p *MapAny) GetInt32(key string) int32 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return candy.ToInt32(val)
}

func (p *MapAny) GetInt64(key string) int64 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return candy.ToInt64(val)
}

func (p *MapAny) GetUint16(key string) uint16 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return candy.ToUint16(val)
}

func (p *MapAny) GetUint32(key string) uint32 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return candy.ToUint32(val)
}

func (p *MapAny) GetUint64(key string) uint64 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return candy.ToUint64(val)
}

func (p *MapAny) GetFloat64(key string) float64 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return candy.ToFloat64(val)
}

func (p *MapAny) GetString(key string) string {
	val, ok := p.get(key)
	if !ok {
		return ""
	}

	return candy.ToString(val)
}

func (p *MapAny) GetBytes(key string) []byte {
	val, ok := p.get(key)
	if !ok {
		return []byte("")
	}

	switch x := val.(type) {
	case bool:
		if x {
			return []byte("1")
		}
		return []byte("0")
	case int:
		return []byte(fmt.Sprintf("%d", x))
	case int8:
		return []byte(fmt.Sprintf("%d", x))
	case int16:
		return []byte(fmt.Sprintf("%d", x))
	case int32:
		return []byte(fmt.Sprintf("%d", x))
	case int64:
		return []byte(fmt.Sprintf("%d", x))
	case uint:
		return []byte(fmt.Sprintf("%d", x))
	case uint8:
		return []byte(fmt.Sprintf("%d", x))
	case uint16:
		return []byte(fmt.Sprintf("%d", x))
	case uint32:
		return []byte(fmt.Sprintf("%d", x))
	case uint64:
		return []byte(fmt.Sprintf("%d", x))
	case float32:
		return []byte(fmt.Sprintf("%v", x))
	case float64:
		return []byte(fmt.Sprintf("%v", x))
	case string:
		return []byte(x)
	case []byte:
		return x
	default:
		return []byte("")
	}
}

func (p *MapAny) GetMap(key string) *MapAny {
	val, ok := p.get(key)
	if !ok {
		return NewMap(nil)
	}

	return p.toMap(val)
}

func (p *MapAny) toMap(val interface{}) *MapAny {
	switch x := val.(type) {
	case bool, int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return NewMap(nil)
	case string:
		var m map[string]interface{}
		err := json.Unmarshal([]byte(x), &m)
		if err != nil {
			return NewMap(nil)
		}
		return NewMap(m)
	case []byte:
		var m map[string]interface{}
		err := json.Unmarshal(x, &m)
		if err != nil {
			return NewMap(nil)
		}
		return NewMap(m)
	case map[string]interface{}:
		return NewMap(x)
	case map[interface{}]interface{}:
		m := NewMap(nil)
		for k, v := range x {
			m.Set(candy.ToString(k), v)
		}
		return m
	default:
		buf, err := json.Marshal(x)
		if err != nil {
			return NewMap(nil)
		}
		var m map[string]interface{}
		err = json.Unmarshal(buf, &m)
		if err != nil {
			return NewMap(nil)
		}
		return NewMap(m)
	}
}

func (p *MapAny) GetSlice(key string) []interface{} {
	val, ok := p.get(key)
	if !ok {
		return nil
	}

	switch x := val.(type) {
	case []bool:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []int:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []int8:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []int16:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []int32:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []int64:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []uint:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []uint8:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []uint16:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []uint32:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []uint64:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []float32:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []float64:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []string:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case [][]byte:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []interface{}:
		return x
	default:
		return []interface{}{}
	}
}

func (p *MapAny) GetStringSlice(key string) []string {
	val, ok := p.get(key)
	if !ok {
		return nil
	}

	switch x := val.(type) {
	case []bool:
		var v []string
		for _, val := range x {
			v = append(v, candy.ToString(val))
		}
		return v
	case []int:
		var v []string
		for _, val := range x {
			v = append(v, candy.ToString(val))
		}
		return v
	case []int8:
		var v []string
		for _, val := range x {
			v = append(v, candy.ToString(val))
		}
		return v
	case []int16:
		var v []string
		for _, val := range x {
			v = append(v, candy.ToString(val))
		}
		return v
	case []int32:
		var v []string
		for _, val := range x {
			v = append(v, candy.ToString(val))
		}
		return v
	case []int64:
		var v []string
		for _, val := range x {
			v = append(v, candy.ToString(val))
		}
		return v
	case []uint:
		var v []string
		for _, val := range x {
			v = append(v, candy.ToString(val))
		}
		return v
	case []uint8:
		var v []string
		for _, val := range x {
			v = append(v, candy.ToString(val))
		}
		return v
	case []uint16:
		var v []string
		for _, val := range x {
			v = append(v, candy.ToString(val))
		}
		return v
	case []uint32:
		var v []string
		for _, val := range x {
			v = append(v, candy.ToString(val))
		}
		return v
	case []uint64:
		var v []string
		for _, val := range x {
			v = append(v, candy.ToString(val))
		}
		return v
	case []float32:
		var v []string
		for _, val := range x {
			v = append(v, candy.ToString(val))
		}
		return v
	case []float64:
		var v []string
		for _, val := range x {
			v = append(v, candy.ToString(val))
		}
		return v
	case []string:
		return x
	case [][]byte:
		var v []string
		for _, val := range x {
			v = append(v, string(val))
		}
		return v
	case []interface{}:
		var v []string
		for _, val := range x {
			v = append(v, candy.ToString(val))
		}
		return v
	default:
		return []string{}
	}
}

func (p *MapAny) GetUint64Slice(key string) []uint64 {
	val, ok := p.get(key)
	if !ok {
		return nil
	}

	switch x := val.(type) {
	case []bool:
		var v []uint64
		for _, val := range x {
			v = append(v, candy.ToUint64(val))
		}
		return v
	case []int:
		var v []uint64
		for _, val := range x {
			v = append(v, candy.ToUint64(val))
		}
		return v
	case []int8:
		var v []uint64
		for _, val := range x {
			v = append(v, candy.ToUint64(val))
		}
		return v
	case []int16:
		var v []uint64
		for _, val := range x {
			v = append(v, candy.ToUint64(val))
		}
		return v
	case []int32:
		var v []uint64
		for _, val := range x {
			v = append(v, candy.ToUint64(val))
		}
		return v
	case []int64:
		var v []uint64
		for _, val := range x {
			v = append(v, candy.ToUint64(val))
		}
		return v
	case []uint:
		var v []uint64
		for _, val := range x {
			v = append(v, candy.ToUint64(val))
		}
		return v
	case []uint8:
		var v []uint64
		for _, val := range x {
			v = append(v, candy.ToUint64(val))
		}
		return v
	case []uint16:
		var v []uint64
		for _, val := range x {
			v = append(v, candy.ToUint64(val))
		}
		return v
	case []uint32:
		var v []uint64
		for _, val := range x {
			v = append(v, candy.ToUint64(val))
		}
		return v
	case []uint64:
		var v []uint64
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []float32:
		var v []uint64
		for _, val := range x {
			v = append(v, candy.ToUint64(val))
		}
		return v
	case []float64:
		var v []uint64
		for _, val := range x {
			v = append(v, candy.ToUint64(val))
		}
		return v
	case []string:
		var v []uint64
		for _, val := range x {
			v = append(v, candy.ToUint64(val))
		}
		return v
	case [][]byte:
		var v []uint64
		for _, val := range x {
			v = append(v, candy.ToUint64(val))
		}
		return v
	case []interface{}:
		var v []uint64
		for _, val := range x {
			v = append(v, candy.ToUint64(val))
		}
		return v
	default:
		return []uint64{}
	}
}

func (p *MapAny) GetInt64Slice(key string) []int64 {
	val, ok := p.get(key)
	if !ok {
		return nil
	}

	return candy.ToInt64Slice(val)
}

func (p *MapAny) GetUint32Slice(key string) []uint32 {
	val, ok := p.get(key)
	if !ok {
		return nil
	}

	switch x := val.(type) {
	case []bool:
		var v []uint32
		for _, val := range x {
			v = append(v, candy.ToUint32(val))
		}
		return v
	case []int:
		var v []uint32
		for _, val := range x {
			v = append(v, candy.ToUint32(val))
		}
		return v
	case []int8:
		var v []uint32
		for _, val := range x {
			v = append(v, candy.ToUint32(val))
		}
		return v
	case []int16:
		var v []uint32
		for _, val := range x {
			v = append(v, candy.ToUint32(val))
		}
		return v
	case []int32:
		var v []uint32
		for _, val := range x {
			v = append(v, candy.ToUint32(val))
		}
		return v
	case []int64:
		var v []uint32
		for _, val := range x {
			v = append(v, candy.ToUint32(val))
		}
		return v
	case []uint:
		var v []uint32
		for _, val := range x {
			v = append(v, candy.ToUint32(val))
		}
		return v
	case []uint8:
		var v []uint32
		for _, val := range x {
			v = append(v, candy.ToUint32(val))
		}
		return v
	case []uint16:
		var v []uint32
		for _, val := range x {
			v = append(v, candy.ToUint32(val))
		}
		return v
	case []uint32:
		return x
	case []uint64:
		var v []uint32
		for _, val := range x {
			v = append(v, candy.ToUint32(val))
		}
		return v
	case []float32:
		var v []uint32
		for _, val := range x {
			v = append(v, candy.ToUint32(val))
		}
		return v
	case []float64:
		var v []uint32
		for _, val := range x {
			v = append(v, candy.ToUint32(val))
		}
		return v
	case []string:
		var v []uint32
		for _, val := range x {
			v = append(v, candy.ToUint32(val))
		}
		return v
	case [][]byte:
		var v []uint32
		for _, val := range x {
			v = append(v, candy.ToUint32(val))
		}
		return v
	case []interface{}:
		var v []uint32
		for _, val := range x {
			v = append(v, candy.ToUint32(val))
		}
		return v
	default:
		return []uint32{}
	}
}

func (p *MapAny) ToSyncMap() *sync.Map {
	var m sync.Map
	p.data.Range(func(key, value interface{}) bool {
		m.Store(key, value)
		return true
	})
	return &m
}

func (p *MapAny) ToMap() map[string]interface{} {
	m := map[string]interface{}{}
	p.data.Range(func(key, value interface{}) bool {
		k := candy.ToString(key)

		switch x := value.(type) {
		case float32:
			if math.Floor(float64(x)) == float64(x) {
				m[k] = int32(x)
			} else {
				m[k] = x
			}
		case float64:
			if math.Floor(x) == x {
				m[k] = int64(x)
			} else {
				m[k] = x
			}
		case *MapAny:
			m[k] = x.ToMap()
		case bool,
			int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64,
			string, []byte:
			m[k] = x
		default:
			m[k] = x
		}

		return true
	})
	return m
}

func (p *MapAny) Clone() *MapAny {
	return &MapAny{
		data: p.ToSyncMap(),
		cut:  atomic.NewBool(p.cut.Load()),
		seq:  atomic.NewString(p.seq.Load()),
	}
}

func (p *MapAny) Range(f func(key, value interface{}) bool) {
	p.data.Range(func(key, value interface{}) bool {
		return f(key, value)
	})
}

func MapGet(m map[string]any, key string) (any, error) {
	return mapGetWithSeparator(m, key, ".")
}

func MapGetIgnore(m map[string]any, key string) (value any) {
	value, _ = mapGetWithSeparator(m, key, ".")
	return value
}

func MapGetMust(m map[string]any, key string) any {
	value, err := mapGetWithSeparator(m, key, ".")
	if err != nil {
		log.Panicf("err:%v", err)
	}
	return value
}

// MapGetWithSep gets a value from a nested map using a custom separator.
// Supports array/slice indexing with [index] syntax.
func MapGetWithSep(m map[string]any, key string, sep string) (any, error) {
	return mapGetWithSeparator(m, key, sep)
}

// MapExists checks if a key exists in the map.
// Returns true if the key is found, false otherwise.
// Supports nested key access with "." separator.
func MapExists(m map[string]any, key string) bool {
	_, err := mapGetWithSeparator(m, key, ".")
	return err == nil
}

// MapExistsWithSep checks if a key exists in the map using a custom separator.
// Returns true if the key is found, false otherwise.
func MapExistsWithSep(m map[string]any, key string, sep string) bool {
	_, err := mapGetWithSeparator(m, key, sep)
	return err == nil
}

// New error types for detailed error reporting
var (
	ErrInvalidIndex   = errors.New("invalid index")
	ErrTypeMismatch   = errors.New("type mismatch")
	ErrOutOfRange     = errors.New("index out of range")
	ErrInvalidMapType = errors.New("invalid map type")
	ErrInvalidSlice   = errors.New("invalid slice type")
	ErrEmptyKey       = errors.New("empty key")
)

// mapGetWithSeparator is the internal implementation for nested map access
func mapGetWithSeparator(m map[string]any, key string, sep string) (any, error) {
	// Handle empty map
	if len(m) == 0 {
		return nil, ErrNotFound
	}

	// Handle empty key
	if key == "" {
		return nil, ErrEmptyKey
	}

	// Split the key into parts
	parts := splitKey(key, sep)
	if len(parts) == 0 {
		return nil, ErrEmptyKey
	}

	var current any = m
	for i, part := range parts {
		val, err := navigateToValue(current, part)
		if err != nil {
			// Check if this is the last part and we can provide more context
			if i == len(parts)-1 {
				return nil, fmt.Errorf("%w: key '%s' not found at path '%s'", err, part, joinPath(parts[:i+1], sep))
			}
			return nil, fmt.Errorf("%w: at path '%s', key '%s' not found", err, joinPath(parts[:i], sep), part)
		}
		current = val
	}

	return current, nil
}

// splitKey splits a key into parts, handling both separators and array indices
// splitKey splits a key into parts, handling both separators and array indices
// splitKey splits a key into parts, handling both separators and array indices
// splitKey splits a key into parts, handling both separators and array indices
// splitKey splits a key into parts, handling both separators and array indices
// splitKey splits a key into parts, handling both separators and array indices
func splitKey(key string, sep string) []string {
	var parts []string
	current := new(strings.Builder)
	inBrackets := false
	afterBrackets := false
	sepLen := len(sep)
	i := 0
	endsWithSep := strings.HasSuffix(key, sep)

	for i < len(key) {
		c := key[i]
		switch {
		case c == '[':
			inBrackets = true
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
			current.WriteByte(c)
			afterBrackets = false
		case c == ']':
			inBrackets = false
			current.WriteByte(c)
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
			afterBrackets = true
		case !inBrackets && i+sepLen <= len(key) && key[i:i+sepLen] == sep:
			// Only add empty part if we're not just after brackets
			if current.Len() > 0 || !afterBrackets {
				parts = append(parts, current.String())
			}
			current.Reset()
			i += sepLen - 1 // Skip the separator
			afterBrackets = false
		default:
			current.WriteByte(c)
			afterBrackets = false
		}
		i++
	}

	// Add the last part if non-empty, or if key ends with separator
	if current.Len() > 0 || endsWithSep {
		parts = append(parts, current.String())
	}

	return parts
}

// navigateToValue navigates through a value using a key part (which may be an array index)
func navigateToValue(current any, part string) (any, error) {
	// Check if this is an array index
	if len(part) > 2 && part[0] == '[' && part[len(part)-1] == ']' {
		indexStr := part[1 : len(part)-1]
		return accessArrayIndex(current, indexStr)
	}

	// Regular key access
	return accessMapKey(current, part)
}

// accessArrayIndex accesses an array/slice by index
// accessArrayIndex accesses an array/slice by index
func accessArrayIndex(current any, indexStr string) (any, error) {
	// Parse the index
	index, err := parseIndex(indexStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidIndex, indexStr)
	}

	switch v := current.(type) {
	case []any:
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("%w: index %d out of range [0, %d)", ErrOutOfRange, index, len(v))
		}
		return v[index], nil
	case []string:
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("%w: index %d out of range [0, %d)", ErrOutOfRange, index, len(v))
		}
		return v[index], nil
	case []int:
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("%w: index %d out of range [0, %d)", ErrOutOfRange, index, len(v))
		}
		return v[index], nil
	case []int64:
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("%w: index %d out of range [0, %d)", ErrOutOfRange, index, len(v))
		}
		return v[index], nil
	case []float64:
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("%w: index %d out of range [0, %d)", ErrOutOfRange, index, len(v))
		}
		return v[index], nil
	case []bool:
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("%w: index %d out of range [0, %d)", ErrOutOfRange, index, len(v))
		}
		return v[index], nil
	case []map[string]any:
		if index < 0 || index >= len(v) {
			return nil, fmt.Errorf("%w: index %d out of range [0, %d)", ErrOutOfRange, index, len(v))
		}
		return v[index], nil
	default:
		// Try to handle other slice types via reflection-like approach
		return accessGenericSlice(v, index)
	}
}

// parseIndex parses an index string to an integer
func parseIndex(s string) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("%w: empty index", ErrInvalidIndex)
	}

	// Handle negative indices
	negative := false
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}

	var result int
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("%w: '%s'", ErrInvalidIndex, s)
		}
		result = result*10 + int(c-'0')
	}

	if negative {
		result = -result
	}

	return result, nil
}

// accessGenericSlice handles generic slice types
func accessGenericSlice(slice any, index int) (any, error) {
	// Use reflection-like approach to handle various slice types
	// For now, return a type mismatch error
	return nil, fmt.Errorf("%w: cannot access index %d on type %T", ErrInvalidSlice, index, slice)
}

// accessMapKey accesses a map using a string key
func accessMapKey(current any, key string) (any, error) {
	switch v := current.(type) {
	case map[string]any:
		val, ok := v[key]
		if !ok {
			return nil, ErrNotFound
		}
		return val, nil
	case map[any]any:
		val, ok := v[key]
		if !ok {
			return nil, ErrNotFound
		}
		return val, nil
	default:
		return nil, fmt.Errorf("%w: cannot access key '%s' on type %T", ErrInvalidMapType, key, current)
	}
}

// joinPath joins parts with a separator for error messages
// joinPath joins parts with a separator for error messages
// joinPath joins parts with a separator for error messages
func joinPath(parts []string, sep string) string {
	switch len(parts) {
	case 0:
		return ""
	case 1:
		return parts[0]
	}

	result := parts[0]
	for _, part := range parts[1:] {
		result += sep + part
	}
	return result
}
