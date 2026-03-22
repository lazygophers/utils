package app

var Organization = "lazygophers"

var Name string

var Version string

type ReleaseType uint8

const (
	Debug ReleaseType = iota
	Test
	Alpha
	Beta
	Release
)

func (p ReleaseType) String() string {
	switch p {
	case Release:
		return "release"
	case Beta:
		return "beta"
	case Alpha:
		return "alpha"
	case Test:
		return "test"
	case Debug:
		fallthrough
	default:
		return "debug"
	}
}

// IsDebug 判断是否为调试环境
func (p ReleaseType) IsDebug() bool {
	return p == Debug
}

// IsTest 判断是否为测试环境
func (p ReleaseType) IsTest() bool {
	return p == Test
}

// IsAlpha 判断是否为 Alpha 版本
func (p ReleaseType) IsAlpha() bool {
	return p == Alpha
}

// IsBeta 判断是否为 Beta 版本
func (p ReleaseType) IsBeta() bool {
	return p == Beta
}

// IsRelease 判断是否为正式环境
func (p ReleaseType) IsRelease() bool {
	return p == Release
}

// IsDev 判断当前环境是否为调试环境（全局便捷函数）
func IsDev() bool {
	return Env == Debug
}

// IsDebug 判断当前环境是否为调试环境（全局便捷函数）
func IsDebug() bool {
	return Env == Debug
}

// IsTest 判断当前环境是否为测试环境（全局便捷函数）
func IsTest() bool {
	return Env == Test
}

// IsAlpha 判断当前环境是否为 Alpha 版本（全局便捷函数）
func IsAlpha() bool {
	return Env == Alpha
}

// IsBeta 判断当前环境是否为 Beta 版本（全局便捷函数）
func IsBeta() bool {
	return Env == Beta
}

// IsProd 判断当前环境是否为正式环境（全局便捷函数）
func IsProd() bool {
	return Env == Release
}

// IsRelease 判断当前环境是否为正式环境（全局便捷函数）
func IsRelease() bool {
	return Env == Release
}

