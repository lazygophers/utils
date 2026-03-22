package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseVersion(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		major       int
		minor       int
		patch       int
		build       int
		prerelease  string
		metadata    string
		expectError bool
	}{
		// 标准语义化版本
		{"标准版本", "1.2.3", 1, 2, 3, 0, "", "", false},
		// 四段版本
		{"四段版本", "1.2.3.4", 1, 2, 3, 4, "", "", false},
		// 带 v 前缀
		{"v前缀小写", "v1.2.3", 1, 2, 3, 0, "", "", false},
		{"v前缀大写", "V1.2.3", 1, 2, 3, 0, "", "", false},
		// 预发布版本
		{"alpha版本", "1.2.3-alpha", 1, 2, 3, 0, "alpha", "", false},
		{"beta版本", "1.2.3-beta", 1, 2, 3, 0, "beta", "", false},
		{"rc版本", "1.2.3-rc", 1, 2, 3, 0, "rc", "", false},
		{"beta带点", "1.2.3-beta.1", 1, 2, 3, 0, "beta.1", "", false},
		// 元数据
		{"元数据", "1.2.3+meta", 1, 2, 3, 0, "", "meta", false},
		{"长元数据", "1.2.3+20210101", 1, 2, 3, 0, "", "20210101", false},
		// 预发布+元数据
		{"预发布+元数据", "1.2.3-alpha+meta", 1, 2, 3, 0, "alpha", "meta", false},
		{"完整版", "1.2.3-beta.2+20210101", 1, 2, 3, 0, "beta.2", "20210101", false},
		// 零版本
		{"零版本", "0.0.1", 0, 0, 1, 0, "", "", false},
		{"全零", "0.0.0", 0, 0, 0, 0, "", "", false},
		// 大版本号
		{"大版本号", "100.200.300", 100, 200, 300, 0, "", "", false},
		// 无效格式
		{"无效-太少段", "1.2", 0, 0, 0, 0, "", "", true},
		{"无效-非数字", "a.b.c", 0, 0, 0, 0, "", "", true},
		{"无效-空", "", 0, 0, 0, 0, "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := Parse(tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, v)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, v)
				assert.Equal(t, tt.major, v.Major)
				assert.Equal(t, tt.minor, v.Minor)
				assert.Equal(t, tt.patch, v.Patch)
				assert.Equal(t, tt.build, v.Build)
				assert.Equal(t, tt.prerelease, v.Prerelease)
				assert.Equal(t, tt.metadata, v.Metadata)
			}
		})
	}
}

func TestVersionString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"标准版本", "1.2.3", "1.2.3"},
		{"四段版本", "1.2.3.4", "1.2.3.4"},
		{"预发布版本", "1.2.3-alpha", "1.2.3-alpha"},
		{"元数据版本", "1.2.3+meta", "1.2.3+meta"},
		{"完整版", "1.2.3-beta.2+meta", "1.2.3-beta.2+meta"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := MustParse(tt.input)
			assert.Equal(t, tt.expected, v.String())
		})
	}
}

func TestVersionShort(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"标准版本", "1.2.3", "1.2.3"},
		{"四段版本", "1.2.3.4", "1.2.3.4"},
		{"预发布版本（忽略）", "1.2.3-alpha", "1.2.3"},
		{"元数据版本（忽略）", "1.2.3+meta", "1.2.3"},
		{"完整版（忽略额外）", "1.2.3-beta.2+meta", "1.2.3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := MustParse(tt.input)
			assert.Equal(t, tt.expected, v.Short())
		})
	}
}

func TestVersionCompare(t *testing.T) {
	tests := []struct {
		name     string
		v1       string
		v2       string
		expected int
	}{
		{"相等", "1.2.3", "1.2.3", 0},
		{"主版本大", "2.0.0", "1.9.9", 1},
		{"主版本小", "1.9.9", "2.0.0", -1},
		{"次版本大", "1.5.0", "1.4.9", 1},
		{"次版本小", "1.4.9", "1.5.0", -1},
		{"修订号大", "1.2.5", "1.2.4", 1},
		{"修订号小", "1.2.4", "1.2.5", -1},
		{"构建号大", "1.2.3.5", "1.2.3.4", 1},
		{"构建号小", "1.2.3.4", "1.2.3.5", -1},
		{"预发布 < 正式", "1.2.3-alpha", "1.2.3", -1},
		{"正式 > 预发布", "1.2.3", "1.2.3-beta", 1},
		{"预发布比较", "1.2.3-alpha", "1.2.3-beta", -1},
		{"预发布相等", "1.2.3-alpha", "1.2.3-alpha", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := MustParse(tt.v1)
			v2 := MustParse(tt.v2)
			result := v1.Compare(v2)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestVersionComparisonMethods(t *testing.T) {
	v100 := MustParse("1.0.0")
	v200 := MustParse("2.0.0")
	v100Copy := MustParse("1.0.0")

	assert.True(t, v100.Equals(v100Copy))
	assert.False(t, v100.Equals(v200))

	assert.True(t, v100.LessThan(v200))
	assert.False(t, v200.LessThan(v100))

	assert.True(t, v200.GreaterThan(v100))
	assert.False(t, v100.GreaterThan(v200))

	assert.True(t, v100.LessThanOrEqual(v200))
	assert.True(t, v100.LessThanOrEqual(v100Copy))

	assert.True(t, v200.GreaterThanOrEqual(v100))
	assert.True(t, v100.GreaterThanOrEqual(v100Copy))
}

func TestVersionIsPrerelease(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		isPrerelease bool
		isStable     bool
	}{
		{"稳定版", "1.2.3", false, true},
		{"alpha", "1.2.3-alpha", true, false},
		{"beta", "1.2.3-beta", true, false},
		{"rc", "1.2.3-rc.1", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := MustParse(tt.input)
			assert.Equal(t, tt.isPrerelease, v.IsPrerelease())
			assert.Equal(t, tt.isStable, v.IsStable())
		})
	}
}

func TestVersionHasMetadata(t *testing.T) {
	assert.False(t, MustParse("1.2.3").HasMetadata())
	assert.True(t, MustParse("1.2.3+meta").HasMetadata())
	assert.True(t, MustParse("1.2.3-alpha+meta").HasMetadata())
}

func TestVersionGetters(t *testing.T) {
	v := MustParse("1.2.3-beta.2+meta")

	assert.Equal(t, 1, v.Major)
	assert.Equal(t, 2, v.Minor)
	assert.Equal(t, 3, v.Patch)
	assert.Equal(t, 0, v.Build)
	assert.Equal(t, "beta.2", v.Prerelease)
	assert.Equal(t, "meta", v.Metadata)
}

func TestVersionGettersWithBuild(t *testing.T) {
	v := MustParse("1.2.3.4")

	assert.Equal(t, 1, v.Major)
	assert.Equal(t, 2, v.Minor)
	assert.Equal(t, 3, v.Patch)
	assert.Equal(t, 4, v.Build)
}

func TestMustParse(t *testing.T) {
	assert.NotPanics(t, func() {
		MustParse("1.2.3")
	})

	assert.Panics(t, func() {
		MustParse("invalid")
	})
}

func TestSetVersion(t *testing.T) {
	// 保存原始版本
	originalVersion := Version
	defer func() { Version = originalVersion }()

	err := SetVersion("2.0.0")
	assert.NoError(t, err)
	assert.Equal(t, "2.0.0", Version)
	assert.Equal(t, "2.0.0", GetVersion().String())

	err = SetVersion("invalid")
	assert.Error(t, err)
}

func TestMustSetVersion(t *testing.T) {
	// 保存原始版本
	originalVersion := Version
	defer func() { Version = originalVersion }()

	assert.NotPanics(t, func() {
		MustSetVersion("3.0.0")
	})

	assert.Panics(t, func() {
		MustSetVersion("invalid")
	})
}

func TestGetVersion(t *testing.T) {
	// 保存原始版本
	originalVersion := Version
	originalCached := cachedVersion
	defer func() {
		Version = originalVersion
		cachedVersion = originalCached
	}()

	// 测试默认版本
	Version = ""
	cachedVersion = nil
	v := GetVersion()
	assert.Equal(t, 0, v.Major)
	assert.Equal(t, 0, v.Minor)
	assert.Equal(t, 1, v.Patch)

	// 测试设置的版本
	Version = "1.2.3"
	cachedVersion = nil
	v = GetVersion()
	assert.Equal(t, "1.2.3", v.String())
}

func TestVersionInitParsesVersion(t *testing.T) {
	// 注意：这个测试验证 init 函数的行为
	// 由于 init 只执行一次，我们无法直接测试重复 init
	// 但 GetVersion() 会返回缓存的版本

	v := GetVersion()
	// 如果全局 Version 变量被设置（通过 ldflags），应该被正确解析
	// 否则返回默认版本
	assert.NotNil(t, v)
}
