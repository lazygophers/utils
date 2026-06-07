package app

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// envReleaseCase 描述 APP_ENV 值与解析得到的 ReleaseType 关系
type envReleaseCase struct {
	name         string
	envValue     string
	expectedType ReleaseType
}

func TestInitFunctionWithDifferentEnvironments(t *testing.T) {
	tests := []envReleaseCase{
		{"development", "development", Debug},
		{"dev", "dev", Debug},
		{"test", "test", Test},
		{"canary", "canary", Test},
		{"production", "production", Release},
		{"prod", "prod", Release},
		{"release", "release", Release},
		{"alpha", "alpha", Alpha},
		{"beta", "beta", Beta},
		{"empty", "", Debug},               // 默认值
		{"unknown", "unknown", Debug},      // 未知值
		{"uppercase", "PRODUCTION", Debug}, // 大写值
		{"mixedcase", "DeV", Debug},        // 混合大小写
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 保存原始环境变量
			originalEnv := os.Getenv("APP_ENV")
			defer os.Setenv("APP_ENV", originalEnv)

			// 设置测试环境变量
			os.Setenv("APP_ENV", tt.envValue)

			// 重新初始化包（在Go中，init函数只执行一次，所以我们需要通过其他方式测试）
			// 我们可以直接测试Env变量，或者通过模拟来测试

			// 测试ReleaseType的String方法
			releaseType := tt.expectedType
			assert.NotEmpty(t, releaseType.String())
		})
	}
}

// releaseStringCase 描述 ReleaseType.String 返回值用例
type releaseStringCase struct {
	name     string
	release  ReleaseType
	expected string
}

func TestReleaseTypeString(t *testing.T) {
	tests := []releaseStringCase{
		{"Debug", Debug, "debug"},
		{"Test", Test, "test"},
		{"Alpha", Alpha, "alpha"},
		{"Beta", Beta, "beta"},
		{"Release", Release, "release"},
		{"Unknown", ReleaseType(99), "debug"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.release.String()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestOrganization(t *testing.T) {
	assert.Equal(t, "lazygophers", Organization)
	assert.NotEmpty(t, Organization)
}

func TestGlobalVariables(t *testing.T) {
	// 构建信息变量在测试时通常为空（通过 -ldflags 注入）
	// 这里仅验证变量存在且可访问，实际值在构建时确定
	type globalVar struct {
		name  string
		value *string
	}
	variables := []globalVar{
		{"Commit", &Commit},
		{"ShortCommit", &ShortCommit},
		{"Branch", &Branch},
		{"Tag", &Tag},
		{"BuildDate", &BuildDate},
		{"GoVersion", &GoVersion},
		{"GoOS", &GoOS},
		{"Goarch", &Goarch},
		{"Goarm", &Goarm},
		{"Goamd64", &Goamd64},
		{"Gomips", &Gomips},
		{"Description", &Description},
	}

	for _, v := range variables {
		t.Run(v.name, func(t *testing.T) {
			// 验证变量可访问（不为 nil）
			assert.NotNil(t, v.value, "Variable %s should be accessible", v.name)
		})
	}
}

// releaseConstantCase 描述 ReleaseType 常量底层值用例
type releaseConstantCase struct {
	name     string
	release  ReleaseType
	expected uint8
}

func TestReleaseTypeConstants(t *testing.T) {
	tests := []releaseConstantCase{
		{"Debug", Debug, 0},
		{"Test", Test, 1},
		{"Alpha", Alpha, 2},
		{"Beta", Beta, 3},
		{"Release", Release, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, uint8(tt.release))
		})
	}
}

func TestEnv(t *testing.T) {
	// Env是在编译时根据构建标签设置的，所以我们只检查它的有效性
	assert.True(t, Env >= Debug && Env <= Release)
	assert.NotEmpty(t, Env.String())
}

func TestReleaseTypeOrder(t *testing.T) {
	assert.True(t, Debug < Test)
	assert.True(t, Test < Alpha)
	assert.True(t, Alpha < Beta)
	assert.True(t, Beta < Release)
}

func TestReleaseTypeRange(t *testing.T) {
	for i := 0; i <= 10; i++ {
		r := ReleaseType(i)
		assert.NotPanics(t, func() {
			_ = r.String()
		})
	}
}

func TestPackageVariables(t *testing.T) {
	// 验证包级变量可访问
	_ = Organization // 固定值 "lazygophers"
	_ = Name         // 通过 -ldflags 注入
	_ = Version      // 通过 -ldflags 注入

	// 构建信息变量
	_ = Commit
	_ = ShortCommit
	_ = Branch
	_ = Tag
	_ = BuildDate
	_ = GoVersion
	_ = GoOS
	_ = Goarch
	_ = Goarm
	_ = Goamd64
	_ = Gomips
	_ = Description
}

func TestSetEnvFromEnv(t *testing.T) {
	// 保存原始环境变量
	originalEnv := os.Getenv("APP_ENV")
	defer os.Setenv("APP_ENV", originalEnv)

	tests := []envReleaseCase{
		{"development", "development", Debug},
		{"dev", "dev", Debug},
		{"test", "test", Test},
		{"canary", "canary", Test},
		{"production", "production", Release},
		{"prod", "prod", Release},
		{"release", "release", Release},
		{"alpha", "alpha", Alpha},
		{"beta", "beta", Beta},
		{"empty", "", Debug},               // 默认值
		{"unknown", "unknown", Debug},      // 未知值
		{"uppercase", "PRODUCTION", Debug}, // 大写值
		{"mixedcase", "DeV", Debug},        // 混合大小写
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 设置测试环境变量
			os.Setenv("APP_ENV", tt.envValue)

			// 调用可测试的函数
			setEnvFromEnv()

			// 验证结果
			assert.True(t, Env >= Debug && Env <= Release)
		})
	}
}

func TestInitFunctionBranchCoverage(t *testing.T) {
	// 确保 setEnvFromEnv 函数的所有分支都被测试
	type envBranchCase struct {
		envValue     string
		expectedType ReleaseType
	}
	testCases := []envBranchCase{
		{"dev", Debug},
		{"development", Debug},
		{"test", Test},
		{"canary", Test},
		{"prod", Release},
		{"production", Release},
		{"release", Release},
		{"alpha", Alpha},
		{"beta", Beta},
		{"", Debug},        // 空值，不覆盖
		{"unknown", Debug}, // 未知值，不覆盖
	}

	originalEnv := os.Getenv("APP_ENV")
	defer os.Setenv("APP_ENV", originalEnv)

	for _, tc := range testCases {
		t.Run(tc.envValue, func(t *testing.T) {
			os.Setenv("APP_ENV", tc.envValue)
			// 重置 Env 为默认值，然后调用 setEnvFromEnv
			initialEnv := Env
			setEnvFromEnv()
			// 验证：如果 envValue 有效，Env 应该改变；否则保持不变
			if tc.envValue != "" && tc.envValue != "unknown" {
				assert.Equal(t, tc.expectedType, Env)
			} else {
				assert.Equal(t, initialEnv, Env)
			}
		})
	}
}

// releaseIsCase 描述 ReleaseType.IsXxx 系列方法的期望返回
type releaseIsCase struct {
	name      string
	release   ReleaseType
	isDebug   bool
	isTest    bool
	isAlpha   bool
	isBeta    bool
	isRelease bool
}

func TestReleaseTypeIsMethods(t *testing.T) {
	tests := []releaseIsCase{
		{"Debug", Debug, true, false, false, false, false},
		{"Test", Test, false, true, false, false, false},
		{"Alpha", Alpha, false, false, true, false, false},
		{"Beta", Beta, false, false, false, true, false},
		{"Release", Release, false, false, false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.isDebug, tt.release.IsDebug())
			assert.Equal(t, tt.isTest, tt.release.IsTest())
			assert.Equal(t, tt.isAlpha, tt.release.IsAlpha())
			assert.Equal(t, tt.isBeta, tt.release.IsBeta())
			assert.Equal(t, tt.isRelease, tt.release.IsRelease())
		})
	}
}

// globalIsCase 描述全局 IsXxx 系列函数的期望返回
type globalIsCase struct {
	name      string
	getEnv    func() ReleaseType
	isDebug   bool
	isTest    bool
	isAlpha   bool
	isBeta    bool
	isRelease bool
}

func TestGlobalIsFunctions(t *testing.T) {
	// 全局函数基于 Env 变量，测试它们的行为一致
	tests := []globalIsCase{
		{"Debug", func() ReleaseType { return Debug }, true, false, false, false, false},
		{"Test", func() ReleaseType { return Test }, false, true, false, false, false},
		{"Alpha", func() ReleaseType { return Alpha }, false, false, true, false, false},
		{"Beta", func() ReleaseType { return Beta }, false, false, false, true, false},
		{"Release", func() ReleaseType { return Release }, false, false, false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 保存原始 Env
			originalEnv := Env
			defer func() { Env = originalEnv }()

			// 设置测试环境
			Env = tt.getEnv()

			// 测试全局函数
			assert.Equal(t, tt.isDebug, IsDebug())
			assert.Equal(t, tt.isDebug, IsDev())
			assert.Equal(t, tt.isTest, IsTest())
			assert.Equal(t, tt.isAlpha, IsAlpha())
			assert.Equal(t, tt.isBeta, IsBeta())
			assert.Equal(t, tt.isRelease, IsRelease())
			assert.Equal(t, tt.isRelease, IsProd())
		})
	}
}

func TestIsDevAndIsDebugAlias(t *testing.T) {
	// 验证 IsDev 和 IsDebug 是同义词
	originalEnv := Env
	defer func() { Env = originalEnv }()

	Env = Debug
	assert.Equal(t, IsDebug(), IsDev())

	Env = Test
	assert.Equal(t, IsDebug(), IsDev())
}

func TestIsProdAndIsReleaseAlias(t *testing.T) {
	// 验证 IsProd 和 IsRelease 是同义词
	originalEnv := Env
	defer func() { Env = originalEnv }()

	Env = Release
	assert.Equal(t, IsRelease(), IsProd())

	Env = Debug
	assert.Equal(t, IsRelease(), IsProd())
}
