package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestToOptimizations 验证优化后的函数仍然保持正确性
func TestToOptimizations(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		wantInt  int
		wantUint uint
		wantF64  float64
		wantBool bool
		wantStr  string
	}{
		{
			name:     "integer values",
			input:    42,
			wantInt:  42,
			wantUint: 42,
			wantF64:  42.0,
			wantBool: true,
			wantStr:  "42",
		},
		{
			name:     "string integer",
			input:    "123",
			wantInt:  123,
			wantUint: 123,
			wantF64:  123.0,
			wantBool: true,
			wantStr:  "123",
		},
		{
			name:     "float value",
			input:    123.45,
			wantInt:  123,
			wantUint: 123,
			wantF64:  123.45,
			wantBool: true,
			wantStr:  "123.450000",
		},
		{
			name:     "boolean true",
			input:    true,
			wantInt:  1,
			wantUint: 1,
			wantF64:  1.0,
			wantBool: true,
			wantStr:  "1",
		},
		{
			name:     "boolean false",
			input:    false,
			wantInt:  0,
			wantUint: 0,
			wantF64:  0.0,
			wantBool: false,
			wantStr:  "0",
		},
		{
			name:     "nil value",
			input:    nil,
			wantInt:  0,
			wantUint: 0,
			wantF64:  0.0,
			wantBool: false,
			wantStr:  "",
		},
		{
			name:     "byte slice",
			input:    []byte("hello"),
			wantInt:  0, // can't parse
			wantUint: 0, // can't parse
			wantF64:  0, // can't parse
			wantBool: true,
			wantStr:  "hello",
		},
		{
			name:     "negative int",
			input:    -42,
			wantInt:  -42,
			wantUint: 0, // negative returns 0
			wantF64:  -42.0,
			wantBool: true,
			wantStr:  "-42",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test ToInt
			if got := ToInt(tt.input); got != tt.wantInt {
				t.Errorf("ToInt(%v) = %v, want %v", tt.input, got, tt.wantInt)
			}

			// Test ToUint
			if got := ToUint(tt.input); got != tt.wantUint {
				t.Errorf("ToUint(%v) = %v, want %v", tt.input, got, tt.wantUint)
			}

			// Test ToFloat64
			if got := ToFloat64(tt.input); got != tt.wantF64 {
				t.Errorf("ToFloat64(%v) = %v, want %v", tt.input, got, tt.wantF64)
			}

			// Test ToBool
			if got := ToBool(tt.input); got != tt.wantBool {
				t.Errorf("ToBool(%v) = %v, want %v", tt.input, got, tt.wantBool)
			}

			// Test ToString
			if got := ToString(tt.input); got != tt.wantStr {
				t.Errorf("ToString(%v) = %v, want %v", tt.input, got, tt.wantStr)
			}

			// Test ToBytes
			if tt.input != nil {
				gotBytes := ToBytes(tt.input)
				wantBytes := []byte(tt.wantStr)
				assert.Equal(t, wantBytes, gotBytes, "ToBytes(%v) should match", tt.input)
			} else {
				if got := ToBytes(tt.input); got != nil {
					t.Errorf("ToBytes(nil) = %v, want nil", got)
				}
			}
		})
	}
}

// TestEdgeCases 测试边界情况
func TestEdgeCases(t *testing.T) {
	// Test large numbers
	assert.Equal(t, int(2147483647), ToInt(int32(2147483647)))
	assert.Equal(t, uint(4294967295), ToUint(uint32(4294967295)))

	// Test zero values
	assert.Equal(t, 0, ToInt(0))
	assert.Equal(t, uint(0), ToUint(0))
	assert.Equal(t, 0.0, ToFloat64(0))
	assert.Equal(t, false, ToBool(0))
	assert.Equal(t, "0", ToString(0))

	// Test empty string
	assert.Equal(t, 0, ToInt(""))
	assert.Equal(t, uint(0), ToUint(""))
	assert.Equal(t, 0.0, ToFloat64(""))
	assert.Equal(t, false, ToBool(""))
	assert.Equal(t, "", ToString(""))
	assert.Equal(t, []byte{}, ToBytes("")) // empty string returns empty byte slice, not nil

	// Test special float values
	assert.Equal(t, false, ToBool(float64(0.0)))
	assert.Equal(t, true, ToBool(float64(1.0)))
	assert.Equal(t, true, ToBool(float64(0.5))) // non-zero float is true
}

// TestTypeConversions 测试各种类型转换
func TestTypeConversions(t *testing.T) {
	// Integer to integer conversions
	assert.Equal(t, int(8), ToInt(int8(8)))
	assert.Equal(t, int(16), ToInt(int16(16)))
	assert.Equal(t, int(32), ToInt(int32(32)))
	assert.Equal(t, int(64), ToInt(int64(64)))

	assert.Equal(t, uint(8), ToUint(uint8(8)))
	assert.Equal(t, uint(16), ToUint(uint16(16)))
	assert.Equal(t, uint(32), ToUint(uint32(32)))
	assert.Equal(t, uint(64), ToUint(uint64(64)))

	// Float conversions
	assert.Equal(t, 123, ToInt(float32(123.45)))
	assert.Equal(t, 123, ToInt(float64(123.45)))
	assert.InDelta(t, float64(123.45), ToFloat64(float32(123.45)), 0.001) // float32 precision loss

	// Bool to numeric
	assert.Equal(t, 1, ToInt(true))
	assert.Equal(t, 0, ToInt(false))
	assert.Equal(t, uint(1), ToUint(true))
	assert.Equal(t, uint(0), ToUint(false))
	assert.Equal(t, 1.0, ToFloat64(true))
	assert.Equal(t, 0.0, ToFloat64(false))

	// String parsing
	assert.Equal(t, 123, ToInt("123"))
	assert.Equal(t, uint(123), ToUint("123"))
	assert.Equal(t, 123.45, ToFloat64("123.45"))
	assert.Equal(t, true, ToBool("true"))
	assert.Equal(t, false, ToBool("false"))
}
