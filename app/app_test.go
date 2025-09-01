package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestReleaseTypeString 测试 ReleaseType 的 String 方法
func TestReleaseTypeString(t *testing.T) {
	tests := []struct {
		give    ReleaseType
		want    string
		desc    string
	}{
		{Debug, "debug", "Debug 类型应该返回 debug"},
		{Test, "test", "Test 类型应该返回 test"},
		{Alpha, "alpha", "Alpha 类型应该返回 alpha"},
		{Beta, "beta", "Beta 类型应该返回 beta"},
		{Release, "release", "Release 类型应该返回 release"},
		{ReleaseType(99), "debug", "未知类型应该返回 debug"},
	}

	for _, tt := range tests {
		tt := tt // 避免竞态
		t.Run(tt.desc, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.give.String(), "String() 方法返回值错误")
			assert.Equal(t, tt.want, tt.give.ʔ(), "ʔ() 方法返回值错误")
		})
	}
}

// TestPackageType 测试 PackageType 全局变量
func TestPackageType(t *testing.T) {
	// 测试 PackageType 的零值
	assert.Equal(t, ReleaseType(0), PackageType, "PackageType 初始值应该为零值")
	
	// 测试 PackageType 可以被修改
	originalType := PackageType
	PackageType = Alpha
	assert.Equal(t, Alpha, PackageType, "PackageType 应该可以被修改")
	
	// 恢复原值
	PackageType = originalType
	assert.Equal(t, originalType, PackageType, "PackageType 应该可以恢复原值")
}

// TestReleaseTypeValues 测试 ReleaseType 常量值
func TestReleaseTypeValues(t *testing.T) {
	// 测试常量值的递增性
	assert.True(t, Debug < Test, "Debug 应该小于 Test")
	assert.True(t, Test < Alpha, "Test 应该小于 Alpha")
	assert.True(t, Alpha < Beta, "Alpha 应该小于 Beta")
	assert.True(t, Beta < Release, "Beta 应该小于 Release")
}

// BenchmarkReleaseTypeString 性能测试
func BenchmarkReleaseTypeString(b *testing.B) {
	types := []ReleaseType{Debug, Test, Alpha, Beta, Release}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = types[i%len(types)].String()
	}
}