package candy

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// ToBytes 优化方案
func toBytes_Current(val interface{}) []byte {
	return ToBytes(val)
}

func toBytes_Strconv(val interface{}) []byte {
	switch x := val.(type) {
	case bool:
		if x {
			return []byte("1")
		}
		return []byte("0")
	case int:
		return []byte(strconv.FormatInt(int64(x), 10))
	case int64:
		return []byte(strconv.FormatInt(x, 10))
	case string:
		return []byte(x)
	case []byte:
		return x
	case nil:
		return nil
	default:
		return []byte(fmt.Sprintf("%v", x))
	}
}

func toBytes_FastPath(val interface{}) []byte {
	switch x := val.(type) {
	case string:
		return []byte(x)
	case []byte:
		return x
	case nil:
		return nil
	case bool:
		if x {
			return []byte("1")
		}
		return []byte("0")
	case int:
		return []byte(strconv.FormatInt(int64(x), 10))
	case int64:
		return []byte(strconv.FormatInt(x, 10))
	default:
		return []byte(fmt.Sprintf("%v", x))
	}
}

// 基准测试
func BenchmarkToBytes_Bool_Current(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toBytes_Current(true)
	}
}

func BenchmarkToBytes_Bool_Strconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toBytes_Strconv(true)
	}
}

func BenchmarkToBytes_Bool_FastPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toBytes_FastPath(true)
	}
}

func BenchmarkToBytes_Int_Current(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toBytes_Current(42)
	}
}

func BenchmarkToBytes_Int_Strconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toBytes_Strconv(42)
	}
}

func BenchmarkToBytes_Int_FastPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toBytes_FastPath(42)
	}
}

func BenchmarkToBytes_Int64_Current(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toBytes_Current(int64(9223372036854775807))
	}
}

func BenchmarkToBytes_Int64_Strconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toBytes_Strconv(int64(9223372036854775807))
	}
}

func BenchmarkToBytes_Int64_FastPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toBytes_FastPath(int64(9223372036854775807))
	}
}

func BenchmarkToBytes_Float64_Current(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toBytes_Current(3.141592653589793)
	}
}

func BenchmarkToBytes_String_Current(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toBytes_Current("hello world, test string")
	}
}

func BenchmarkToBytes_String_FastPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toBytes_FastPath("hello world, test string")
	}
}

func BenchmarkToBytes_Bytes_Current(b *testing.B) {
	data := []byte("test data")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = toBytes_Current(data)
	}
}

func BenchmarkToBytes_Bytes_FastPath(b *testing.B) {
	data := []byte("test data")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = toBytes_FastPath(data)
	}
}

func BenchmarkToBytes_Duration_Current(b *testing.B) {
	d := time.Second * 5
	for i := 0; i < b.N; i++ {
		_ = toBytes_Current(d)
	}
}

func BenchmarkToBytes_Nil_Current(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toBytes_Current(nil)
	}
}
