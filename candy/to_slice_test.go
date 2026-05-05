package candy

import (
	"reflect"
	"testing"
)

func TestToFloat64Slice(t *testing.T) {
	if got := ToFloat64Slice(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	tests := []struct {
		name string
		in   any
		want []float64
	}{
		{name: "bool", in: []bool{true, false}, want: []float64{1, 0}},
		{name: "int", in: []int{1, 2}, want: []float64{1, 2}},
		{name: "float32", in: []float32{1.5, 2.5}, want: []float64{1.5, 2.5}},
		{name: "float64", in: []float64{1.5, 2.5}, want: []float64{1.5, 2.5}},
		{name: "string", in: []string{"1", "2"}, want: []float64{1, 2}},
		{name: "any", in: []any{"1", 2}, want: []float64{1, 2}},
		{name: "int8", in: []int8{1, 2}, want: []float64{1, 2}},
		{name: "uint", in: []uint{1, 2}, want: []float64{1, 2}},
		{name: "int16", in: []int16{1, 2}, want: []float64{1, 2}},
		{name: "int32", in: []int32{1, 2}, want: []float64{1, 2}},
		{name: "int64", in: []int64{1, 2}, want: []float64{1, 2}},
		{name: "uint8", in: []uint8{1, 2}, want: []float64{1, 2}},
		{name: "uint16", in: []uint16{1, 2}, want: []float64{1, 2}},
		{name: "uint32", in: []uint32{1, 2}, want: []float64{1, 2}},
		{name: "uint64", in: []uint64{1, 2}, want: []float64{1, 2}},
		{name: "[][]byte", in: [][]byte{[]byte("1"), []byte("2")}, want: []float64{1, 2}},
		{name: "unsupported", in: 123, want: []float64{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat64Slice(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got=%v want=%v", got, tt.want)
			}
		})
	}
}

func TestToInt64Slice(t *testing.T) {
	tests := []struct {
		name string
		in   any
		want []int64
	}{
		{name: "bool", in: []bool{true, false}, want: []int64{1, 0}},
		{name: "int", in: []int{1, 2}, want: []int64{1, 2}},
		{name: "int64", in: []int64{1, 2}, want: []int64{1, 2}},
		{name: "float64", in: []float64{1.5, 2.5}, want: []int64{1, 2}},
		{name: "string", in: []string{"1", "2"}, want: []int64{1, 2}},
		{name: "any", in: []any{"1", 2}, want: []int64{1, 2}},
		{name: "int8", in: []int8{1, 2}, want: []int64{1, 2}},
		{name: "uint", in: []uint{1, 2}, want: []int64{1, 2}},
		{name: "int16", in: []int16{1, 2}, want: []int64{1, 2}},
		{name: "int32", in: []int32{1, 2}, want: []int64{1, 2}},
		{name: "uint8", in: []uint8{1, 2}, want: []int64{1, 2}},
		{name: "uint16", in: []uint16{1, 2}, want: []int64{1, 2}},
		{name: "uint32", in: []uint32{1, 2}, want: []int64{1, 2}},
		{name: "uint64", in: []uint64{1, 2}, want: []int64{1, 2}},
		{name: "float32", in: []float32{1.5, 2.5}, want: []int64{1, 2}},
		{name: "[][]byte", in: [][]byte{[]byte("1"), []byte("2")}, want: []int64{1, 2}},
		{name: "unsupported", in: 123, want: []int64{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt64Slice(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got=%v want=%v", got, tt.want)
			}
		})
	}
}

func TestToStringSlice(t *testing.T) {
	if got := ToStringSlice(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	if got := ToStringSlice("a,b"); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToStringSlice("x"); !reflect.DeepEqual(got, []string{"x"}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToStringSlice([]any{"a", 1}); !reflect.DeepEqual(got, []string{"a", "1"}) {
		t.Fatalf("unexpected: %v", got)
	}

	// 测试 nil 切片
	var nilSlice []string
	if got := ToStringSlice(nilSlice); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	// 测试非切片、非 string、非 nil 的类型
	if got := ToStringSlice(123); !reflect.DeepEqual(got, []string{"123"}) {
		t.Fatalf("unexpected: %v", got)
	}

	// 测试更多切片类型
	if got := ToStringSlice([]int{1, 2}); !reflect.DeepEqual(got, []string{"1", "2"}) {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestToArrayString(t *testing.T) {
	if got := ToArrayString("a,b"); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestToUint64Slice(t *testing.T) {
	if got := ToUint64Slice(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	tests := []struct {
		name string
		in   any
		want []uint64
	}{
		{name: "bool", in: []bool{true, false}, want: []uint64{1, 0}},
		{name: "int", in: []int{1, 2}, want: []uint64{1, 2}},
		{name: "uint64", in: []uint64{1, 2}, want: []uint64{1, 2}},
		{name: "float64", in: []float64{1.5, 2.5}, want: []uint64{1, 2}},
		{name: "string", in: []string{"1", "2"}, want: []uint64{1, 2}},
		{name: "any", in: []any{"1", 2}, want: []uint64{1, 2}},
		{name: "int8", in: []int8{1, 2}, want: []uint64{1, 2}},
		{name: "uint32", in: []uint32{1, 2}, want: []uint64{1, 2}},
		{name: "int16", in: []int16{1, 2}, want: []uint64{1, 2}},
		{name: "int32", in: []int32{1, 2}, want: []uint64{1, 2}},
		{name: "int64", in: []int64{1, 2}, want: []uint64{1, 2}},
		{name: "uint8", in: []uint8{1, 2}, want: []uint64{1, 2}},
		{name: "uint16", in: []uint16{1, 2}, want: []uint64{1, 2}},
		{name: "uint", in: []uint{1, 2}, want: []uint64{1, 2}},
		{name: "float32", in: []float32{1.5, 2.5}, want: []uint64{1, 2}},
		{name: "[][]byte", in: [][]byte{[]byte("1"), []byte("2")}, want: []uint64{1, 2}},
		{name: "unsupported", in: 123, want: []uint64{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint64Slice(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got=%v want=%v", got, tt.want)
			}
		})
	}
}

func TestToUint32Slice(t *testing.T) {
	if got := ToUint32Slice(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	tests := []struct {
		name string
		in   any
		want []uint32
	}{
		{name: "bool", in: []bool{true, false}, want: []uint32{1, 0}},
		{name: "int", in: []int{1, 2}, want: []uint32{1, 2}},
		{name: "uint32", in: []uint32{1, 2}, want: []uint32{1, 2}},
		{name: "float64", in: []float64{1.5, 2.5}, want: []uint32{1, 2}},
		{name: "string", in: []string{"1", "2"}, want: []uint32{1, 2}},
		{name: "any", in: []any{"1", 2}, want: []uint32{1, 2}},
		{name: "int8", in: []int8{1, 2}, want: []uint32{1, 2}},
		{name: "uint64", in: []uint64{1, 2}, want: []uint32{1, 2}},
		{name: "int16", in: []int16{1, 2}, want: []uint32{1, 2}},
		{name: "int32", in: []int32{1, 2}, want: []uint32{1, 2}},
		{name: "int64", in: []int64{1, 2}, want: []uint32{1, 2}},
		{name: "uint8", in: []uint8{1, 2}, want: []uint32{1, 2}},
		{name: "uint16", in: []uint16{1, 2}, want: []uint32{1, 2}},
		{name: "uint", in: []uint{1, 2}, want: []uint32{1, 2}},
		{name: "float32", in: []float32{1.5, 2.5}, want: []uint32{1, 2}},
		{name: "[][]byte", in: [][]byte{[]byte("1"), []byte("2")}, want: []uint32{1, 2}},
		{name: "unsupported", in: 123, want: []uint32{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint32Slice(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got=%v want=%v", got, tt.want)
			}
		})
	}
}

func TestToInterfaceSlice(t *testing.T) {
	if got := ToInterfaceSlice(nil); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	if got := ToInterfaceSlice([]int{1, 2, 3}); !reflect.DeepEqual(got, []any{1, 2, 3}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToInterfaceSlice([]string{"a", "b"}); !reflect.DeepEqual(got, []any{"a", "b"}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToInterfaceSlice(123); len(got) != 0 {
		t.Fatalf("unexpected: %v", got)
	}

	// 测试更多类型分支
	if got := ToInterfaceSlice([]bool{true, false}); !reflect.DeepEqual(got, []any{true, false}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToInterfaceSlice([]int8{1, 2}); !reflect.DeepEqual(got, []any{int8(1), int8(2)}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToInterfaceSlice([]int16{1, 2}); !reflect.DeepEqual(got, []any{int16(1), int16(2)}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToInterfaceSlice([]int32{1, 2}); !reflect.DeepEqual(got, []any{int32(1), int32(2)}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToInterfaceSlice([]int64{1, 2}); !reflect.DeepEqual(got, []any{int64(1), int64(2)}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToInterfaceSlice([]uint{1, 2}); !reflect.DeepEqual(got, []any{uint(1), uint(2)}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToInterfaceSlice([]uint8{1, 2}); !reflect.DeepEqual(got, []any{uint8(1), uint8(2)}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToInterfaceSlice([]uint16{1, 2}); !reflect.DeepEqual(got, []any{uint16(1), uint16(2)}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToInterfaceSlice([]uint32{1, 2}); !reflect.DeepEqual(got, []any{uint32(1), uint32(2)}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToInterfaceSlice([]uint64{1, 2}); !reflect.DeepEqual(got, []any{uint64(1), uint64(2)}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToInterfaceSlice([]float32{1.5, 2.5}); !reflect.DeepEqual(got, []any{float32(1.5), float32(2.5)}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToInterfaceSlice([]float64{1.5, 2.5}); !reflect.DeepEqual(got, []any{float64(1.5), float64(2.5)}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToInterfaceSlice([][]byte{[]byte("a"), []byte("b")}); !reflect.DeepEqual(got, []any{[]byte("a"), []byte("b")}) {
		t.Fatalf("unexpected: %v", got)
	}

	if got := ToInterfaceSlice([]any{1, "a"}); !reflect.DeepEqual(got, []any{1, "a"}) {
		t.Fatalf("unexpected: %v", got)
	}
}
