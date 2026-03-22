package app

import (
	"fmt"
	"strconv"
	"strings"
)

// Version 表示应用程序版本信息
// 支持格式：
//   - 1.2.3 (标准语义化版本)
//   - 1.2.3.4 (四段版本)
//   - v1.2.3 (带 v 前缀)
//   - 1.2.3-alpha (预发布版本)
//   - 1.2.3-beta.1 (带预发布版本号)
//   - 1.2.3+meta (带构建元数据)
//   - 1.2.3-alpha+meta (预发布+元数据)
type SemVer struct {
	Major      int    // 主版本号
	Minor      int    // 次版本号
	Patch      int    // 修订号
	Build      int    // 构建号（可选第4段）
	Prerelease string // 预发布标识（如 alpha, beta, rc）
	Metadata   string // 构建元数据（+ 后面的内容）
}

// Parse 解析版本字符串
func Parse(version string) (*SemVer, error) {
	v := &SemVer{}

	// 移除 v 前缀
	version = strings.TrimPrefix(version, "v")
	version = strings.TrimPrefix(version, "V")

	// 分离元数据
	if idx := strings.Index(version, "+"); idx != -1 {
		v.Metadata = version[idx+1:]
		version = version[:idx]
	}

	// 分离预发布标识
	if idx := strings.Index(version, "-"); idx != -1 {
		v.Prerelease = version[idx+1:]
		version = version[:idx]
	}

	// 解析版本号
	parts := strings.Split(version, ".")
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid version format: %s", version)
	}

	// 解析主版本号
	major, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid major version: %s", parts[0])
	}
	v.Major = int(major)

	// 解析次版本号
	minor, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid minor version: %s", parts[1])
	}
	v.Minor = int(minor)

	// 解析修订号
	patch, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid patch version: %s", parts[2])
	}
	v.Patch = int(patch)

	// 解析构建号（可选）
	if len(parts) >= 4 {
		build, err := strconv.ParseUint(parts[3], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid build version: %s", parts[3])
		}
		v.Build = int(build)
	}

	return v, nil
}

// MustParse 解析版本字符串，出错时 panic
func MustParse(version string) *SemVer {
	v, err := Parse(version)
	if err != nil {
		panic(err)
	}
	return v
}

// String 返回版本字符串
func (v *SemVer) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch))

	if v.Build > 0 {
		sb.WriteString(fmt.Sprintf(".%d", v.Build))
	}

	if v.Prerelease != "" {
		sb.WriteString("-")
		sb.WriteString(v.Prerelease)
	}

	if v.Metadata != "" {
		sb.WriteString("+")
		sb.WriteString(v.Metadata)
	}

	return sb.String()
}

// Short 返回短格式版本（不包含预发布和元数据）
func (v *SemVer) Short() string {
	if v.Build > 0 {
		return fmt.Sprintf("%d.%d.%d.%d", v.Major, v.Minor, v.Patch, v.Build)
	}
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// Compare 比较版本
// 返回值：-1 表示 v < other, 0 表示相等, 1 表示 v > other
func (v *SemVer) Compare(other *SemVer) int {
	// 比较主版本号
	if v.Major != other.Major {
		if v.Major < other.Major {
			return -1
		}
		return 1
	}

	// 比较次版本号
	if v.Minor != other.Minor {
		if v.Minor < other.Minor {
			return -1
		}
		return 1
	}

	// 比较修订号
	if v.Patch != other.Patch {
		if v.Patch < other.Patch {
			return -1
		}
		return 1
	}

	// 比较构建号
	if v.Build != other.Build {
		if v.Build < other.Build {
			return -1
		}
		return 1
	}

	// 预发布版本 < 正式版本
	vIsPre := v.Prerelease != ""
	otherIsPre := other.Prerelease != ""

	if vIsPre && !otherIsPre {
		return -1
	}
	if !vIsPre && otherIsPre {
		return 1
	}

	// 两者都是预发布版本，比较预发布标识
	if vIsPre && otherIsPre {
		// 简单的字符串比较（可以根据需要增强）
		if v.Prerelease < other.Prerelease {
			return -1
		}
		if v.Prerelease > other.Prerelease {
			return 1
		}
	}

	return 0
}

// Equals 判断版本是否相等
func (v *SemVer) Equals(other *SemVer) bool {
	return v.Compare(other) == 0
}

// LessThan 判断是否小于指定版本
func (v *SemVer) LessThan(other *SemVer) bool {
	return v.Compare(other) < 0
}

// GreaterThan 判断是否大于指定版本
func (v *SemVer) GreaterThan(other *SemVer) bool {
	return v.Compare(other) > 0
}

// LessThanOrEqual 判断是否小于等于指定版本
func (v *SemVer) LessThanOrEqual(other *SemVer) bool {
	return v.Compare(other) <= 0
}

// GreaterThanOrEqual 判断是否大于等于指定版本
func (v *SemVer) GreaterThanOrEqual(other *SemVer) bool {
	return v.Compare(other) >= 0
}

// IsPrerelease 判断是否为预发布版本
func (v *SemVer) IsPrerelease() bool {
	return v.Prerelease != ""
}

// IsStable 判断是否为稳定版本（非预发布）
func (v *SemVer) IsStable() bool {
	return !v.IsPrerelease()
}

// HasMetadata 判断是否有构建元数据
func (v *SemVer) HasMetadata() bool {
	return v.Metadata != ""
}

// cachedVersion 存储解析后的版本对象
var cachedVersion *SemVer

func init() {
	// 在包初始化时解析版本号（如果已设置）
	if Version != "" {
		if v, err := Parse(Version); err == nil {
			cachedVersion = v
		}
		// 如果解析失败，GetVersion() 将返回默认版本
	}
}

// GetVersion 获取当前版本对象
func GetVersion() *SemVer {
	if cachedVersion != nil {
		return cachedVersion
	}

	// 解析全局 Version 变量
	if Version != "" {
		v, err := Parse(Version)
		if err == nil {
			cachedVersion = v
			return cachedVersion
		}
	}

	// 返回默认版本
	return &SemVer{Major: 0, Minor: 0, Patch: 1}
}

// SetVersion 设置版本字符串
func SetVersion(version string) error {
	v, err := Parse(version)
	if err != nil {
		return err
	}
	Version = version
	cachedVersion = v
	return nil
}

// MustSetVersion 设置版本字符串，出错时 panic
func MustSetVersion(version string) {
	if err := SetVersion(version); err != nil {
		panic(err)
	}
}
