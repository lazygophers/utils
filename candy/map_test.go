package candy

import (
	"reflect"
	"testing"
)

func TestMapKeys(t *testing.T) {
	tests := []struct {
		name string
		m    map[int]string
		want []int
	}{
		{
			name: "empty map",
			m:    map[int]string{},
			want: []int{},
		},
		{
			name: "single element",
			m:    map[int]string{1: "a"},
			want: []int{1},
		},
		{
			name: "multiple elements",
			m:    map[int]string{1: "a", 2: "b", 3: "c"},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeys(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeys() length = %v, want %v", len(got), len(tt.want))
				return
			}

			// Convert to maps for comparison since order is not guaranteed
			gotMap := make(map[int]bool)
			wantMap := make(map[int]bool)
			for _, v := range got {
				gotMap[v] = true
			}
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysInt(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want []int
	}{
		{
			name: "nil input",
			m:    nil,
			want: nil,
		},
		{
			name: "non-map input",
			m:    "not a map",
			want: nil,
		},
		{
			name: "empty int map",
			m:    map[int]string{},
			want: []int{},
		},
		{
			name: "int map",
			m:    map[int]string{1: "a", 2: "b", 3: "c"},
			want: []int{1, 2, 3},
		},
		{
			name: "map with non-int keys",
			m:    map[string]int{"a": 1, "b": 2},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysInt(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeysInt() length = %v, want %v", len(got), len(tt.want))
				return
			}

			// Convert to maps for comparison since order is not guaranteed
			gotMap := make(map[int]bool)
			wantMap := make(map[int]bool)
			for _, v := range got {
				gotMap[v] = true
			}
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeysInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysInt8(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want []int8
	}{
		{
			name: "nil input",
			m:    nil,
			want: nil,
		},
		{
			name: "non-map input",
			m:    "not a map",
			want: nil,
		},
		{
			name: "empty int8 map",
			m:    map[int8]string{},
			want: []int8{},
		},
		{
			name: "int8 map",
			m:    map[int8]string{1: "a", 2: "b", -3: "c"},
			want: []int8{1, 2, -3},
		},
		{
			name: "map with non-int8 keys",
			m:    map[string]int{"a": 1, "b": 2},
			want: []int8{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysInt8(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeysInt8() length = %v, want %v", len(got), len(tt.want))
				return
			}

			gotMap := make(map[int8]bool)
			wantMap := make(map[int8]bool)
			for _, v := range got {
				gotMap[v] = true
			}
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeysInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysInt16(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want []int16
	}{
		{
			name: "nil input",
			m:    nil,
			want: nil,
		},
		{
			name: "non-map input",
			m:    "not a map",
			want: nil,
		},
		{
			name: "empty int16 map",
			m:    map[int16]string{},
			want: []int16{},
		},
		{
			name: "int16 map",
			m:    map[int16]string{100: "a", 200: "b", -300: "c"},
			want: []int16{100, 200, -300},
		},
		{
			name: "map with non-int16 keys",
			m:    map[string]int{"a": 1, "b": 2},
			want: []int16{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysInt16(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeysInt16() length = %v, want %v", len(got), len(tt.want))
				return
			}

			gotMap := make(map[int16]bool)
			wantMap := make(map[int16]bool)
			for _, v := range got {
				gotMap[v] = true
			}
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeysInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysInt32(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want []int32
	}{
		{
			name: "nil input",
			m:    nil,
			want: nil,
		},
		{
			name: "non-map input",
			m:    "not a map",
			want: nil,
		},
		{
			name: "empty int32 map",
			m:    map[int32]string{},
			want: []int32{},
		},
		{
			name: "int32 map",
			m:    map[int32]string{1000: "a", 2000: "b", -3000: "c"},
			want: []int32{1000, 2000, -3000},
		},
		{
			name: "map with non-int32 keys",
			m:    map[string]int{"a": 1, "b": 2},
			want: []int32{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysInt32(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeysInt32() length = %v, want %v", len(got), len(tt.want))
				return
			}

			gotMap := make(map[int32]bool)
			wantMap := make(map[int32]bool)
			for _, v := range got {
				gotMap[v] = true
			}
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeysInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysInt64(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want []int64
	}{
		{
			name: "nil input",
			m:    nil,
			want: nil,
		},
		{
			name: "non-map input",
			m:    "not a map",
			want: nil,
		},
		{
			name: "empty int64 map",
			m:    map[int64]string{},
			want: []int64{},
		},
		{
			name: "int64 map",
			m:    map[int64]string{1000000: "a", 2000000: "b", -3000000: "c"},
			want: []int64{1000000, 2000000, -3000000},
		},
		{
			name: "map with non-int64 keys",
			m:    map[string]int{"a": 1, "b": 2},
			want: []int64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysInt64(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeysInt64() length = %v, want %v", len(got), len(tt.want))
				return
			}

			gotMap := make(map[int64]bool)
			wantMap := make(map[int64]bool)
			for _, v := range got {
				gotMap[v] = true
			}
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeysInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysUint(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want []uint
	}{
		{
			name: "nil input",
			m:    nil,
			want: nil,
		},
		{
			name: "non-map input",
			m:    "not a map",
			want: nil,
		},
		{
			name: "empty uint map",
			m:    map[uint]string{},
			want: []uint{},
		},
		{
			name: "uint map",
			m:    map[uint]string{1: "a", 2: "b", 3: "c"},
			want: []uint{1, 2, 3},
		},
		{
			name: "map with non-uint keys",
			m:    map[string]int{"a": 1, "b": 2},
			want: []uint{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysUint(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeysUint() length = %v, want %v", len(got), len(tt.want))
				return
			}

			gotMap := make(map[uint]bool)
			wantMap := make(map[uint]bool)
			for _, v := range got {
				gotMap[v] = true
			}
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeysUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysUint8(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want []uint8
	}{
		{
			name: "nil input",
			m:    nil,
			want: nil,
		},
		{
			name: "non-map input",
			m:    "not a map",
			want: nil,
		},
		{
			name: "empty uint8 map",
			m:    map[uint8]string{},
			want: []uint8{},
		},
		{
			name: "uint8 map",
			m:    map[uint8]string{1: "a", 2: "b", 255: "c"},
			want: []uint8{1, 2, 255},
		},
		{
			name: "map with non-uint8 keys",
			m:    map[string]int{"a": 1, "b": 2},
			want: []uint8{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysUint8(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeysUint8() length = %v, want %v", len(got), len(tt.want))
				return
			}

			gotMap := make(map[uint8]bool)
			wantMap := make(map[uint8]bool)
			for _, v := range got {
				gotMap[v] = true
			}
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeysUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysUint16(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want []uint16
	}{
		{
			name: "nil input",
			m:    nil,
			want: nil,
		},
		{
			name: "non-map input",
			m:    "not a map",
			want: nil,
		},
		{
			name: "empty uint16 map",
			m:    map[uint16]string{},
			want: []uint16{},
		},
		{
			name: "uint16 map",
			m:    map[uint16]string{100: "a", 200: "b", 65535: "c"},
			want: []uint16{100, 200, 65535},
		},
		{
			name: "map with non-uint16 keys",
			m:    map[string]int{"a": 1, "b": 2},
			want: []uint16{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysUint16(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeysUint16() length = %v, want %v", len(got), len(tt.want))
				return
			}

			gotMap := make(map[uint16]bool)
			wantMap := make(map[uint16]bool)
			for _, v := range got {
				gotMap[v] = true
			}
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeysUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysUint32(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want []uint32
	}{
		{
			name: "nil input",
			m:    nil,
			want: nil,
		},
		{
			name: "non-map input",
			m:    "not a map",
			want: nil,
		},
		{
			name: "empty uint32 map",
			m:    map[uint32]string{},
			want: []uint32{},
		},
		{
			name: "uint32 map",
			m:    map[uint32]string{1000: "a", 2000: "b", 4294967295: "c"},
			want: []uint32{1000, 2000, 4294967295},
		},
		{
			name: "map with non-uint32 keys",
			m:    map[string]int{"a": 1, "b": 2},
			want: []uint32{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysUint32(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeysUint32() length = %v, want %v", len(got), len(tt.want))
				return
			}

			gotMap := make(map[uint32]bool)
			wantMap := make(map[uint32]bool)
			for _, v := range got {
				gotMap[v] = true
			}
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeysUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysUint64(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want []uint64
	}{
		{
			name: "nil input",
			m:    nil,
			want: nil,
		},
		{
			name: "non-map input",
			m:    "not a map",
			want: nil,
		},
		{
			name: "empty uint64 map",
			m:    map[uint64]string{},
			want: []uint64{},
		},
		{
			name: "uint64 map",
			m:    map[uint64]string{1000000: "a", 2000000: "b", 18446744073709551615: "c"},
			want: []uint64{1000000, 2000000, 18446744073709551615},
		},
		{
			name: "map with non-uint64 keys",
			m:    map[string]int{"a": 1, "b": 2},
			want: []uint64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysUint64(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeysUint64() length = %v, want %v", len(got), len(tt.want))
				return
			}

			gotMap := make(map[uint64]bool)
			wantMap := make(map[uint64]bool)
			for _, v := range got {
				gotMap[v] = true
			}
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeysUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysFloat32(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want []float32
	}{
		{
			name: "nil input",
			m:    nil,
			want: nil,
		},
		{
			name: "non-map input",
			m:    "not a map",
			want: nil,
		},
		{
			name: "empty float32 map",
			m:    map[float32]string{},
			want: []float32{},
		},
		{
			name: "float32 map",
			m:    map[float32]string{1.1: "a", 2.2: "b", -3.3: "c"},
			want: []float32{1.1, 2.2, -3.3},
		},
		{
			name: "map with non-float32 keys",
			m:    map[string]int{"a": 1, "b": 2},
			want: []float32{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysFloat32(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeysFloat32() length = %v, want %v", len(got), len(tt.want))
				return
			}

			gotMap := make(map[float32]bool)
			wantMap := make(map[float32]bool)
			for _, v := range got {
				gotMap[v] = true
			}
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeysFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysFloat64(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want []float64
	}{
		{
			name: "nil input",
			m:    nil,
			want: nil,
		},
		{
			name: "non-map input",
			m:    "not a map",
			want: nil,
		},
		{
			name: "empty float64 map",
			m:    map[float64]string{},
			want: []float64{},
		},
		{
			name: "float64 map",
			m:    map[float64]string{1.111: "a", 2.222: "b", -3.333: "c"},
			want: []float64{1.111, 2.222, -3.333},
		},
		{
			name: "map with non-float64 keys",
			m:    map[string]int{"a": 1, "b": 2},
			want: []float64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysFloat64(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeysFloat64() length = %v, want %v", len(got), len(tt.want))
				return
			}

			gotMap := make(map[float64]bool)
			wantMap := make(map[float64]bool)
			for _, v := range got {
				gotMap[v] = true
			}
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeysFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysString(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want []string
	}{
		{
			name: "nil input",
			m:    nil,
			want: nil,
		},
		{
			name: "non-map input",
			m:    "not a map",
			want: nil,
		},
		{
			name: "empty string map",
			m:    map[string]int{},
			want: []string{},
		},
		{
			name: "string map",
			m:    map[string]int{"a": 1, "b": 2, "c": 3},
			want: []string{"a", "b", "c"},
		},
		{
			name: "map with non-string keys",
			m:    map[int]string{1: "a", 2: "b"},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysString(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeysString() length = %v, want %v", len(got), len(tt.want))
				return
			}

			gotMap := make(map[string]bool)
			wantMap := make(map[string]bool)
			for _, v := range got {
				gotMap[v] = true
			}
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeysString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapKeysBytes(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want [][]byte
	}{
		{
			name: "nil input",
			m:    nil,
			want: nil,
		},
		{
			name: "non-map input",
			m:    "not a map",
			want: nil,
		},
		{
			name: "map with non-bytes keys",
			m:    map[string]int{"a": 1, "b": 2},
			want: [][]byte{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysBytes(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeysBytes() length = %v, want %v", len(got), len(tt.want))
				return
			}

			// For byte slices, we need to compare element by element
			gotMap := make(map[string]bool)
			wantMap := make(map[string]bool)
			for _, v := range got {
				gotMap[string(v)] = true
			}
			for _, v := range tt.want {
				wantMap[string(v)] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeysBytes() result mismatch")
			}
		})
	}
}

func TestMapKeysAny(t *testing.T) {
	tests := []struct {
		name string
		m    interface{}
		want []interface{}
	}{
		{
			name: "nil input",
			m:    nil,
			want: nil,
		},
		{
			name: "non-map input",
			m:    "not a map",
			want: nil,
		},
		{
			name: "empty map",
			m:    map[string]int{},
			want: []interface{}{},
		},
		{
			name: "string keys map",
			m:    map[string]int{"a": 1, "b": 2, "c": 3},
			want: []interface{}{"a", "b", "c"},
		},
		{
			name: "int keys map",
			m:    map[int]string{1: "a", 2: "b", 3: "c"},
			want: []interface{}{1, 2, 3},
		},
		{
			name: "mixed type keys map",
			m:    map[interface{}]string{"a": "value1", 2: "value2", true: "value3"},
			want: []interface{}{"a", 2, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapKeysAny(tt.m)
			if len(got) != len(tt.want) {
				t.Errorf("MapKeysAny() length = %v, want %v", len(got), len(tt.want))
				return
			}

			// Convert to maps for comparison since order is not guaranteed
			gotMap := make(map[interface{}]bool)
			wantMap := make(map[interface{}]bool)
			for _, v := range got {
				gotMap[v] = true
			}
			for _, v := range tt.want {
				wantMap[v] = true
			}

			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("MapKeysAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark tests
func BenchmarkMapKeys(b *testing.B) {
	m := make(map[int]string, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = "value"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapKeys(m)
	}
}

func BenchmarkMapKeysInt(b *testing.B) {
	m := make(map[int]string, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = "value"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapKeysInt(m)
	}
}

func BenchmarkMapKeysString(b *testing.B) {
	m := make(map[string]int, 1000)
	for i := 0; i < 1000; i++ {
		m[makeKey(i)] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapKeysString(m)
	}
}

func BenchmarkMapKeysAny(b *testing.B) {
	m := make(map[interface{}]string, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = "value"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapKeysAny(m)
	}
}

// Helper function to generate unique keys for benchmarks
func makeKey(i int) string {
	return "key_" + string(rune('a'+(i%26))) + string(rune('a'+((i/26)%26)))
}
