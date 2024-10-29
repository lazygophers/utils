package anyx

import (
	"errors"
	"fmt"
	"github.com/lazygophers/utils/json"
	"golang.org/x/exp/constraints"
	"gopkg.in/yaml.v3"
	"maps"
	"math"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"go.uber.org/atomic"
)

type Map struct {
	data *sync.Map
	cut  *atomic.Bool
	seq  *atomic.String
}

var (
	ErrNotFound = errors.New("not found")
)

type ValueType int

const (
	ValueUnknown ValueType = iota
	ValueNumber
	ValueString
	ValueBool
)

func NewMap(m map[string]interface{}) *Map {
	m2 := &Map{
		data: &sync.Map{},
		cut:  atomic.NewBool(false),
		seq:  atomic.NewString(""),
	}
	for k, v := range m {
		m2.data.Store(k, v)
	}
	return m2
}

func NewMapWithJson(s []byte) (*Map, error) {
	var m map[string]interface{}
	err := json.Unmarshal(s, &m)
	if err != nil {
		return nil, err
	}
	return NewMap(m), nil
}

func NewMapWithYaml(s []byte) (*Map, error) {
	var m map[string]interface{}
	err := yaml.Unmarshal(s, &m)
	if err != nil {
		return nil, err
	}
	return NewMap(m), nil
}

func NewMapWithAny(s interface{}) (*Map, error) {
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

func (p *Map) EnableCut(seq string) *Map {
	p.cut.Store(true)
	p.seq.Store(seq)
	return p
}

func (p *Map) DisableCut() *Map {
	p.cut.Store(false)
	return p
}

func (p *Map) Set(key string, value interface{}) {
	p.data.Store(key, value)
}

func (p *Map) Get(key string) (interface{}, error) {
	val, ok := p.get(key)
	if !ok {
		return nil, ErrNotFound
	}

	return val, nil
}

func (p *Map) get(key string) (interface{}, bool) {
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
	var m *Map
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
		return nil, false
	}

	return nil, false
}

func CheckValueType(val interface{}) ValueType {
	switch val.(type) {
	case bool:
		return ValueBool
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return ValueNumber
	case string, []byte:
		return ValueString
	default:
		return ValueUnknown
	}
}

func (p *Map) Exists(key string) bool {
	_, ok := p.get(key)
	if !ok {
		return false
	}

	return true
}

func (p *Map) GetBool(key string) bool {
	val, ok := p.get(key)
	if !ok {
		return false
	}

	return ToBool(val)
}

func (p *Map) GetInt(key string) int {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return ToInt(val)
}

func (p *Map) GetInt32(key string) int32 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return ToInt32(val)
}

func (p *Map) GetInt64(key string) int64 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return ToInt64(val)
}

func (p *Map) GetUint16(key string) uint16 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return uint16(x)
	case int8:
		return uint16(x)
	case int16:
		return uint16(x)
	case int32:
		return uint16(x)
	case int64:
		return uint16(x)
	case uint:
		return uint16(x)
	case uint8:
		return uint16(x)
	case uint16:
		return x
	case uint32:
		return uint16(x)
	case uint64:
		return uint16(x)
	case float32:
		return uint16(x)
	case float64:
		return uint16(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 16)
		if err != nil {
			return 0
		}
		return uint16(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 16)
		if err != nil {
			return 0
		}
		return uint16(val)
	default:
		return 0
	}
}

func (p *Map) GetUint32(key string) uint32 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return ToUint32(val)
}

func (p *Map) GetUint64(key string) uint64 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return ToUint64(val)
}

func (p *Map) GetFloat64(key string) float64 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return ToFloat64(val)
}

func (p *Map) GetString(key string) string {
	val, ok := p.get(key)
	if !ok {
		return ""
	}

	return ToString(val)
}

func (p *Map) GetBytes(key string) []byte {
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

func (p *Map) GetMap(key string) *Map {
	val, ok := p.get(key)
	if !ok {
		return NewMap(nil)
	}

	return p.toMap(val)
}

func (p *Map) toMap(val interface{}) *Map {
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
			m.Set(ToString(k), v)
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

func (p *Map) GetSlice(key string) []interface{} {
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

func (p *Map) GetStringSlice(key string) []string {
	val, ok := p.get(key)
	if !ok {
		return nil
	}

	switch x := val.(type) {
	case []bool:
		var v []string
		for _, val := range x {
			v = append(v, ToString(val))
		}
		return v
	case []int:
		var v []string
		for _, val := range x {
			v = append(v, ToString(val))
		}
		return v
	case []int8:
		var v []string
		for _, val := range x {
			v = append(v, ToString(val))
		}
		return v
	case []int16:
		var v []string
		for _, val := range x {
			v = append(v, ToString(val))
		}
		return v
	case []int32:
		var v []string
		for _, val := range x {
			v = append(v, ToString(val))
		}
		return v
	case []int64:
		var v []string
		for _, val := range x {
			v = append(v, ToString(val))
		}
		return v
	case []uint:
		var v []string
		for _, val := range x {
			v = append(v, ToString(val))
		}
		return v
	case []uint8:
		var v []string
		for _, val := range x {
			v = append(v, ToString(val))
		}
		return v
	case []uint16:
		var v []string
		for _, val := range x {
			v = append(v, ToString(val))
		}
		return v
	case []uint32:
		var v []string
		for _, val := range x {
			v = append(v, ToString(val))
		}
		return v
	case []uint64:
		var v []string
		for _, val := range x {
			v = append(v, ToString(val))
		}
		return v
	case []float32:
		var v []string
		for _, val := range x {
			v = append(v, ToString(val))
		}
		return v
	case []float64:
		var v []string
		for _, val := range x {
			v = append(v, ToString(val))
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
			v = append(v, ToString(val))
		}
		return v
	default:
		return []string{}
	}
}

func (p *Map) GetUint64Slice(key string) []uint64 {
	val, ok := p.get(key)
	if !ok {
		return nil
	}

	switch x := val.(type) {
	case []bool:
		var v []uint64
		for _, val := range x {
			v = append(v, ToUint64(val))
		}
		return v
	case []int:
		var v []uint64
		for _, val := range x {
			v = append(v, ToUint64(val))
		}
		return v
	case []int8:
		var v []uint64
		for _, val := range x {
			v = append(v, ToUint64(val))
		}
		return v
	case []int16:
		var v []uint64
		for _, val := range x {
			v = append(v, ToUint64(val))
		}
		return v
	case []int32:
		var v []uint64
		for _, val := range x {
			v = append(v, ToUint64(val))
		}
		return v
	case []int64:
		var v []uint64
		for _, val := range x {
			v = append(v, ToUint64(val))
		}
		return v
	case []uint:
		var v []uint64
		for _, val := range x {
			v = append(v, ToUint64(val))
		}
		return v
	case []uint8:
		var v []uint64
		for _, val := range x {
			v = append(v, ToUint64(val))
		}
		return v
	case []uint16:
		var v []uint64
		for _, val := range x {
			v = append(v, ToUint64(val))
		}
		return v
	case []uint32:
		var v []uint64
		for _, val := range x {
			v = append(v, ToUint64(val))
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
			v = append(v, ToUint64(val))
		}
		return v
	case []float64:
		var v []uint64
		for _, val := range x {
			v = append(v, ToUint64(val))
		}
		return v
	case []string:
		var v []uint64
		for _, val := range x {
			v = append(v, ToUint64(val))
		}
		return v
	case [][]byte:
		var v []uint64
		for _, val := range x {
			v = append(v, ToUint64(val))
		}
		return v
	case []interface{}:
		var v []uint64
		for _, val := range x {
			v = append(v, ToUint64(val))
		}
		return v
	default:
		return []uint64{}
	}
}

func (p *Map) GetInt64Slice(key string) []int64 {
	val, ok := p.get(key)
	if !ok {
		return nil
	}

	return ToInt64Slice(val)
}

func (p *Map) GetUint32Slice(key string) []uint32 {
	val, ok := p.get(key)
	if !ok {
		return nil
	}

	switch x := val.(type) {
	case []bool:
		var v []uint32
		for _, val := range x {
			v = append(v, ToUint32(val))
		}
		return v
	case []int:
		var v []uint32
		for _, val := range x {
			v = append(v, ToUint32(val))
		}
		return v
	case []int8:
		var v []uint32
		for _, val := range x {
			v = append(v, ToUint32(val))
		}
		return v
	case []int16:
		var v []uint32
		for _, val := range x {
			v = append(v, ToUint32(val))
		}
		return v
	case []int32:
		var v []uint32
		for _, val := range x {
			v = append(v, ToUint32(val))
		}
		return v
	case []int64:
		var v []uint32
		for _, val := range x {
			v = append(v, ToUint32(val))
		}
		return v
	case []uint:
		var v []uint32
		for _, val := range x {
			v = append(v, ToUint32(val))
		}
		return v
	case []uint8:
		var v []uint32
		for _, val := range x {
			v = append(v, ToUint32(val))
		}
		return v
	case []uint16:
		var v []uint32
		for _, val := range x {
			v = append(v, ToUint32(val))
		}
		return v
	case []uint32:
		return x
	case []uint64:
		var v []uint32
		for _, val := range x {
			v = append(v, ToUint32(val))
		}
		return v
	case []float32:
		var v []uint32
		for _, val := range x {
			v = append(v, ToUint32(val))
		}
		return v
	case []float64:
		var v []uint32
		for _, val := range x {
			v = append(v, ToUint32(val))
		}
		return v
	case []string:
		var v []uint32
		for _, val := range x {
			v = append(v, ToUint32(val))
		}
		return v
	case [][]byte:
		var v []uint32
		for _, val := range x {
			v = append(v, ToUint32(val))
		}
		return v
	case []interface{}:
		var v []uint32
		for _, val := range x {
			v = append(v, ToUint32(val))
		}
		return v
	default:
		return []uint32{}
	}
}

func (p *Map) ToSyncMap() *sync.Map {
	var m sync.Map
	p.data.Range(func(key, value interface{}) bool {
		m.Store(key, value)
		return true
	})
	return &m
}

func (p *Map) ToMap() map[string]interface{} {
	m := map[string]interface{}{}
	p.data.Range(func(key, value interface{}) bool {
		k := ToString(key)

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
		case *Map:
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

func (p *Map) Clone() *Map {
	return &Map{
		data: p.ToSyncMap(),
		cut:  atomic.NewBool(p.cut.Load()),
		seq:  atomic.NewString(p.seq.Load()),
	}
}

func (p *Map) Range(f func(key, value interface{})) bool {
	return p.Range(f)
}

func MapKeysString(m interface{}) []string {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.String {
		panic("map key type required string")
	}

	result := make([]string, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, v.String())
	}

	return result
}

func MapKeysUint32(m interface{}) []uint32 {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Uint32 {
		panic("map key type required string")
	}

	result := make([]uint32, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, uint32(v.Uint()))
	}

	return result
}

func MapKeysUint64(m interface{}) []uint64 {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Uint64 {
		panic("map key type required string")
	}

	result := make([]uint64, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, v.Uint())
	}

	return result
}

func MapKeysInt32(m interface{}) []int32 {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Int32 {
		panic("map key type required string")
	}

	result := make([]int32, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, int32(v.Int()))
	}

	return result
}

func MapKeysInt64(m interface{}) []int64 {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Int64 {
		panic("map key type required string")
	}

	result := make([]int64, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, v.Int())
	}

	return result
}

func MapValues(m interface{}) interface{} {
	vo := reflect.ValueOf(m)
	if vo.Kind() != reflect.Map {
		panic("required map type")
	}

	elType := vo.Type().Elem()
	list := reflect.MakeSlice(reflect.SliceOf(elType), 0, vo.Len())

	for _, key := range vo.MapKeys() {
		list = reflect.Append(list, vo.MapIndex(key))
	}

	return list.Interface()
}

func MergeMap[K constraints.Ordered, V any](source, target map[K]V) map[K]V {
	res := maps.Clone(source)

	if len(target) > 0 {
		for k, v := range target {
			res[k] = v
		}
	}

	return res
}
