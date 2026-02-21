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
}

func TestToArrayString(t *testing.T) {
	if got := ToArrayString("a,b"); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Fatalf("unexpected: %v", got)
	}
}
