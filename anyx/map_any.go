package anyx

import (
	"errors"
	"fmt"
	"github.com/lazygophers/utils/json"
	"go.uber.org/atomic"
	"gopkg.in/yaml.v3"
	"math"
	"strings"
	"sync"
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
		return nil, false
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

	return ToBool(val)
}

func (p *MapAny) GetInt(key string) int {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return ToInt(val)
}

func (p *MapAny) GetInt32(key string) int32 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return ToInt32(val)
}

func (p *MapAny) GetInt64(key string) int64 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return ToInt64(val)
}

func (p *MapAny) GetUint16(key string) uint16 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return ToUint16(val)
}

func (p *MapAny) GetUint32(key string) uint32 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return ToUint32(val)
}

func (p *MapAny) GetUint64(key string) uint64 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return ToUint64(val)
}

func (p *MapAny) GetFloat64(key string) float64 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return ToFloat64(val)
}

func (p *MapAny) GetString(key string) string {
	val, ok := p.get(key)
	if !ok {
		return ""
	}

	return ToString(val)
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

func (p *MapAny) GetUint64Slice(key string) []uint64 {
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

func (p *MapAny) GetInt64Slice(key string) []int64 {
	val, ok := p.get(key)
	if !ok {
		return nil
	}

	return ToInt64Slice(val)
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

func (p *MapAny) Range(f func(key, value interface{})) bool {
	return p.Range(f)
}
