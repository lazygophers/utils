package validator

import (
	"reflect"
	"testing"
)

func BenchmarkMobile01_Regex_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile01(fl)
	}
}

func BenchmarkMobile02_Manual_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile02(fl)
	}
}

func BenchmarkMobile03_Bytes_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile03(fl)
	}
}

func BenchmarkMobile04_FastPath_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile04(fl)
	}
}

func BenchmarkMobile05_Lookup_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile05(fl)
	}
}

func BenchmarkMobile06_LengthFirst_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile06(fl)
	}
}

func BenchmarkMobile09_Combined_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile09(fl)
	}
}

func BenchmarkMobile10_Range_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile10(fl)
	}
}

func BenchmarkMobile11_Unrolled_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile11(fl)
	}
}

func BenchmarkMobile12_BitOps_Valid(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("13812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile12(fl)
	}
}

// 无效前缀测试
func BenchmarkMobile01_Regex_InvalidPrefix(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("12812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile01(fl)
	}
}

func BenchmarkMobile02_Manual_InvalidPrefix(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("12812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile02(fl)
	}
}

func BenchmarkMobile11_Unrolled_InvalidPrefix(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("12812345678")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile11(fl)
	}
}

// 无效长度测试
func BenchmarkMobile01_Regex_InvalidLen(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("138123456")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile01(fl)
	}
}

func BenchmarkMobile02_Manual_InvalidLen(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("138123456")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile02(fl)
	}
}

func BenchmarkMobile11_Unrolled_InvalidLen(b *testing.B) {
	fl := &testFieldLevel{value: reflect.ValueOf("138123456")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateMobile11(fl)
	}
}
