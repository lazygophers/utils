package anyx_test

import (
	"github.com/lazygophers/utils/anyx"
	"reflect"
	"strconv"
	"testing"
)

func TestPluckUint64(t *testing.T) {
	type item struct {
		Id uint64
	}

	type args struct {
		list      interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "",
			args: args{
				list: []item{
					{
						1,
					},
					{
						2,
					},
				},
				fieldName: "Id",
			},
			want: []uint64{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anyx.PluckUint64(tt.args.list, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PluckUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func pluckUint64(max uint64, b *testing.B) {
	type item struct {
		Id uint64
	}

	var items []*item
	for i := uint64(0); i < max; i++ {
		items = append(items, &item{
			Id: i,
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		anyx.PluckUint64(items, "Id")
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/lazygophers/utils
cpu: 11th Gen Intel(R) Core(TM) i7-11700 @ 2.50GHz
BenchmarkPluckUint64_0-16               606172554                5.894 ns/op           0 B/op          0 allocs/op
BenchmarkPluckUint64_10-16               8633337               429.6 ns/op           136 B/op          4 allocs/op
BenchmarkPluckUint64_100-16              1525666              2379 ns/op             952 B/op          4 allocs/op
BenchmarkPluckUint64_1000-16              155276             22949 ns/op            8248 B/op          4 allocs/op
BenchmarkPluckUint64_2000-16               78535             44629 ns/op           16440 B/op          4 allocs/op
BenchmarkPluckUint64_5000-16               35017            102445 ns/op           41016 B/op          4 allocs/op
BenchmarkPluckUint64_10000-16              17898            210740 ns/op           81976 B/op          4 allocs/op
*/
func BenchmarkPluckUint64_0(b *testing.B)     { pluckUint64(0, b) }
func BenchmarkPluckUint64_10(b *testing.B)    { pluckUint64(10, b) }
func BenchmarkPluckUint64_100(b *testing.B)   { pluckUint64(100, b) }
func BenchmarkPluckUint64_1000(b *testing.B)  { pluckUint64(1000, b) }
func BenchmarkPluckUint64_2000(b *testing.B)  { pluckUint64(2000, b) }
func BenchmarkPluckUint64_5000(b *testing.B)  { pluckUint64(5000, b) }
func BenchmarkPluckUint64_10000(b *testing.B) { pluckUint64(10000, b) }

func TestKeyByV2(t *testing.T) {
	type item struct {
		Id   int
		Name string
	}

	type args struct {
		list      interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "",
			args: args{
				list: []*item{
					{
						Id:   1,
						Name: "a",
					},
					{
						Id:   2,
						Name: "b",
					},
					{
						Id:   3,
						Name: "c",
					},
				},
				fieldName: "Id",
			},
			want: map[int]*item{
				1: {
					Id:   1,
					Name: "a",
				},
				2: {
					Id:   2,
					Name: "b",
				},
				3: {
					Id:   3,
					Name: "c",
				},
			},
		},
		{
			name: "",
			args: args{
				list:      []*item{},
				fieldName: "Id",
			},
			want: map[int]*item{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anyx.KeyBy(tt.args.list, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func keyBy(max int64, b *testing.B) {
	type item struct {
		Id   int64
		Name string
	}

	var items []*item
	for i := int64(0); i < max; i++ {
		items = append(items, &item{
			Id:   i,
			Name: strconv.FormatInt(i, 10),
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = anyx.KeyBy(items, "Id").(map[int64]*item)
	}
}

/*
λ go test -bench='KeyBy*' -benchmem -benchtime=3s
goos: windows
goarch: amd64
pkg: github.com/lazygophers/utils
cpu: 11th Gen Intel(R) Core(TM) i7-11700 @ 2.50GHz
BenchmarkKeyB_0-16           24369027               150.7 ns/op            56 B/op          2 allocs/op
BenchmarkKeyBy_10-16           4481101               792.7 ns/op           371 B/op          4 allocs/op
BenchmarkKeyBy_100-16           525364              6070 ns/op            2920 B/op          5 allocs/op
BenchmarkKeyBy_1000-16           58640             60200 ns/op           41064 B/op          5 allocs/op
BenchmarkKeyBy_2000-16           30552            119240 ns/op           82024 B/op          5 allocs/op
BenchmarkKeyBy_5000-16           12266            291591 ns/op          163946 B/op          5 allocs/op
BenchmarkKeyBy_10000-16           5373            602209 ns/op          319594 B/op          5 allocs/op
*/
func BenchmarkKeyBy_0(b *testing.B)     { keyBy(0, b) }
func BenchmarkKeyBy_10(b *testing.B)    { keyBy(10, b) }
func BenchmarkKeyBy_100(b *testing.B)   { keyBy(100, b) }
func BenchmarkKeyBy_1000(b *testing.B)  { keyBy(1000, b) }
func BenchmarkKeyBy_2000(b *testing.B)  { keyBy(2000, b) }
func BenchmarkKeyBy_5000(b *testing.B)  { keyBy(5000, b) }
func BenchmarkKeyBy_10000(b *testing.B) { keyBy(10000, b) }

func TestPluckStringSlice(t *testing.T) {
	type item struct {
		Id    int
		Names []string
	}
	type args struct {
		list      interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "",
			args: args{
				list: []*item{
					{
						Names: []string{"a", "b", "c"},
					},
					{
						Names: []string{"d", "e", "f"},
					},
					{
						Names: []string{"g", "h", "i"},
					},
				},
				fieldName: "Names",
			},
			want: [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anyx.PluckStringSlice(tt.args.list, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PluckStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPluckString(t *testing.T) {
	tests := []struct {
		name string
		args [][]string
		want []string
	}{
		{
			name: "",
			args: [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
				{"g", "h", "i"},
			},
			want: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anyx.PluckString(tt.args, ""); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PluckString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyByUint64(t *testing.T) {
	type Model struct {
		Id   uint64 `json:"id"`
		Name string `json:"name"`
	}
	t.Log(anyx.KeyByUint64([]*Model{
		{
			Id:   2,
			Name: "2",
		},
		{
			Id:   3,
			Name: "3",
		},
	}, "Id"))
}
