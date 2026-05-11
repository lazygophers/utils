package defaults

import (
	"reflect"
	"testing"
	"time"
)

// ========== 验证优化的有效性 ==========

// 验证接口类型优化（预期提升 12.94%）
func BenchmarkSetInterfaceDefault_Optimized(b *testing.B) {
	type TestStruct struct {
		Field interface{} `default:"test"`
	}
	var ts TestStruct
	vv := reflect.ValueOf(&ts.Field).Elem()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vv.Set(reflect.Zero(vv.Type()))
		if err := setInterfaceDefault(vv, "test", &Options{}); err != nil {
			b.Fatal(err)
		}
	}
}

// 验证指针类型优化（预期提升 6.18%）
func BenchmarkSetPtrDefault_Optimized(b *testing.B) {
	type TestStruct struct {
		Field *string `default:"hello"`
	}
	var ts TestStruct
	vv := reflect.ValueOf(&ts.Field).Elem()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vv.Set(reflect.Zero(vv.Type()))
		if err := setPtrDefault(vv, "hello", &Options{}); err != nil {
			b.Fatal(err)
		}
	}
}

// 验证 Channel 类型优化（预期提升 2.85%）
func BenchmarkSetChanDefault_Optimized(b *testing.B) {
	type TestStruct struct {
		Field chan int `default:"10"`
	}
	var ts TestStruct
	vv := reflect.ValueOf(&ts.Field).Elem()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vv.Set(reflect.Zero(vv.Type()))
		if err := setChanDefault(vv, "10", &Options{}); err != nil {
			b.Fatal(err)
		}
	}
}

// 验证时间类型优化（预期提升 0.33%）
func BenchmarkSetTimeDefault_Optimized(b *testing.B) {
	type TestStruct struct {
		Field time.Time `default:"2024-01-15T10:30:00Z"`
	}
	var ts TestStruct
	vv := reflect.ValueOf(&ts.Field).Elem()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := setTimeDefault(vv, "2024-01-15T10:30:00Z", &Options{}); err != nil {
			b.Fatal(err)
		}
	}
}

// ========== 综合性能测试 ==========

func BenchmarkSetDefaults_ComplexStruct(b *testing.B) {
	type ComplexStruct struct {
		TimeField  time.Time     `default:"2024-01-15T10:30:00Z"`
		PtrField   *string       `default:"hello"`
		IfField    interface{}   `default:"test"`
		ChanField  chan int      `default:"10"`
		StringField string        `default:"world"`
		IntField    int           `default:"42"`
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var cs ComplexStruct
		SetDefaults(&cs)
	}
}

// ========== 边界条件测试 ==========

func BenchmarkSetInterface_Default_Empty(b *testing.B) {
	var field interface{}
	vv := reflect.ValueOf(&field).Elem()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setInterfaceDefault(vv, "", &Options{})
	}
}

func BenchmarkSetChan_Default_Zero(b *testing.B) {
	var field chan int
	vv := reflect.ValueOf(&field).Elem()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setChanDefault(vv, "0", &Options{})
	}
}

func BenchmarkSetPtr_NilAlreadyInitialized(b *testing.B) {
	str := "initialized"
	var field *string = &str
	vv := reflect.ValueOf(&field).Elem()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setPtrDefault(vv, "newvalue", &Options{})
	}
}
