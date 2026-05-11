package validator

import (
	"reflect"
	"testing"
)

func TestAlphaOptimized(t *testing.T) {
	validator := Alpha()

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"纯字母", "HelloWorld", true},
		{"混合大小写", "HeLLoWoRLd", true},
		{"长字符串", "TheQuickBrownFoxJumpsOverTheLazyDog", true},
		{"单字符", "H", true},
		{"空字符串", "", false},
		{"包含数字", "Hello123", false},
		{"包含特殊字符", "Test@123", false},
		{"纯数字", "123", false},
		{"包含空格", "Hello World", false},
		{"包含下划线", "Hello_World", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := &fieldLevel{field: reflect.ValueOf(tt.input)}
			result := validator(fl)
			if result != tt.expected {
				t.Errorf("Alpha(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestAlphanumOptimized(t *testing.T) {
	validator := Alphanum()

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"字母数字", "User123", true},
		{"混合", "AbC123xYz", true},
		{"长字符串", "UserID1234567890ABCDEFghijklmnopQRSTUVWXYZ9876543210", true},
		{"单字符字母", "A", true},
		{"单字符数字", "1", true},
		{"空字符串", "", false},
		{"包含特殊字符", "Test@123", false},
		{"包含空格", "Hello World", false},
		{"包含短横线", "Test-123", false},
		{"包含下划线", "Test_123", false},
		{"纯字母", "HelloWorld", true},
		{"纯数字", "123456", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := &fieldLevel{field: reflect.ValueOf(tt.input)}
			result := validator(fl)
			if result != tt.expected {
				t.Errorf("Alphanum(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestAlphaPerformance(t *testing.T) {
	validator := Alpha()
	fl := &fieldLevel{field: reflect.ValueOf("HelloWorldTest")}

	// 运行多次确保没有性能退化
	iterations := 100000
	for i := 0; i < iterations; i++ {
		if !validator(fl) {
			t.Error("Alpha validation failed")
		}
	}
}

func TestAlphanumPerformance(t *testing.T) {
	validator := Alphanum()
	fl := &fieldLevel{field: reflect.ValueOf("User123456789")}

	// 运行多次确保没有性能退化
	iterations := 100000
	for i := 0; i < iterations; i++ {
		if !validator(fl) {
			t.Error("Alphanum validation failed")
		}
	}
}
