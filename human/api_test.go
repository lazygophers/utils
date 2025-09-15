package human

import (
	"testing"
	"time"
)

// TestDirectFunctionAPI 测试新的直接函数API设计
func TestDirectFunctionAPI(t *testing.T) {
	// 测试ByteSize - 直接函数调用
	t.Run("ByteSize direct function", func(t *testing.T) {
		if result := ByteSize(1024); result != "1 KB" {
			t.Errorf("ByteSize(1024) = %s, want 1 KB", result)
		}
	})

	// 测试ByteSize with functional options
	t.Run("ByteSize with functional options", func(t *testing.T) {
		result := ByteSize(1536, WithPrecision(2), WithCompact())
		expected := "1.5KB"
		if result != expected {
			t.Errorf("ByteSize(1536, WithPrecision(2), WithCompact()) = %s, want %s", result, expected)
		}
	})

	// 测试BitSpeed - 展示1000进制
	t.Run("BitSpeed direct function (1000-based)", func(t *testing.T) {
		if result := BitSpeed(8000); result != "8 Kbps" {
			t.Errorf("BitSpeed(8000) = %s, want 8 Kbps", result)
		}
	})

	// 测试Speed vs BitSpeed 的区别
	t.Run("Speed vs BitSpeed distinction", func(t *testing.T) {
		value := int64(8000)
		byteSpeed := Speed(value)   // 1024-based: 8000/1024 = 7.8 KB/s
		bitSpeed := BitSpeed(value) // 1000-based: 8000/1000 = 8 Kbps

		if byteSpeed == bitSpeed {
			t.Error("Speed and BitSpeed should produce different results")
		}

		if byteSpeed != "7.8 KB/s" {
			t.Errorf("Speed(%d) = %s, want 7.8 KB/s", value, byteSpeed)
		}

		if bitSpeed != "8 Kbps" {
			t.Errorf("BitSpeed(%d) = %s, want 8 Kbps", value, bitSpeed)
		}
	})

	// 测试Duration - 直接函数
	t.Run("Duration direct function", func(t *testing.T) {
		result := Duration(90 * time.Second)
		expected := "1 minute 30 seconds"
		if result != expected {
			t.Errorf("Duration(90s) = %s, want %s", result, expected)
		}
	})

	// 测试Duration with clock format - 函数式选项
	t.Run("Duration with clock format option", func(t *testing.T) {
		result := Duration(90*time.Second, WithClockFormat())
		expected := "1:30"
		if result != expected {
			t.Errorf("Duration(90s, WithClockFormat()) = %s, want %s", result, expected)
		}
	})

	// 测试RelativeTime - 直接函数
	t.Run("RelativeTime direct function", func(t *testing.T) {
		past := time.Now().Add(-5 * time.Minute)
		result := RelativeTime(past)
		// 相对时间结果可能会有微小变化，只检查不为空
		if result == "" {
			t.Error("RelativeTime should not return empty string")
		}
	})

	// 测试多个选项组合 - 函数式选项链
	t.Run("Multiple functional options chaining", func(t *testing.T) {
		result := ByteSize(1234567, WithPrecision(3), WithCompact(), WithLocale("en"))
		expected := "1.177MB"
		if result != expected {
			t.Errorf("ByteSize with chained options = %s, want %s", result, expected)
		}
	})

	// 测试零值处理
	t.Run("Zero value handling", func(t *testing.T) {
		if result := ByteSize(0); result != "0 B" {
			t.Errorf("ByteSize(0) = %s, want 0 B", result)
		}

		if result := Duration(0, WithClockFormat()); result != "0:00" {
			t.Errorf("Duration(0, WithClockFormat()) = %s, want 0:00", result)
		}
	})
}

// TestInvalidInputsNewAPI 测试新API的无效输入处理
func TestInvalidInputsNewAPI(t *testing.T) {
	// Test invalid unit type in formatWithUnit
	config := DefaultConfig()
	result := formatWithUnit(1.0, 0, config, "invalid")
	if result != "-" {
		t.Errorf("Invalid input should return '-', got %s", result)
	}
}

// BenchmarkNewAPI 新API的基准测试
func BenchmarkNewAPI(b *testing.B) {
	b.Run("ByteSize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ByteSize(1234567)
		}
	})

	b.Run("ByteSize with options", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ByteSize(1234567, WithPrecision(2), WithCompact())
		}
	})

	b.Run("BitSpeed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BitSpeed(8000)
		}
	})

	b.Run("Duration", func(b *testing.B) {
		d := 90 * time.Minute
		for i := 0; i < b.N; i++ {
			Duration(d)
		}
	})

	b.Run("Duration with clock", func(b *testing.B) {
		d := 90 * time.Minute
		for i := 0; i < b.N; i++ {
			Duration(d, WithClockFormat())
		}
	})
}
