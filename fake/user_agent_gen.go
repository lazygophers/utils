package fake

import (
	"fmt"
	"strings"

	"github.com/lazygophers/utils/randx"
)

// UserAgentOptions 用户代理生成选项
type UserAgentOptions struct {
	DeviceType   DeviceType // 设备类型
	Platform     Platform   // 平台
	Browser      string     // 浏览器
	OS           string     // 操作系统
	OSVersion    string     // 操作系统版本
	Mobile       bool       // 是否移动设备
	Architecture string     // 架构 (x64, ARM, etc.)
}

// BrowserEngine 浏览器引擎信息
type BrowserEngine struct {
	Name    string
	Version string
}

// UserAgentGenerator 用户代理生成器
type UserAgentGenerator struct {
	// 浏览器版本映射
	browserVersions map[string][]string
	// 操作系统版本映射
	osVersions map[Platform][]string
	// 浏览器引擎映射
	browserEngines map[string]BrowserEngine
}

// NewUserAgentGenerator 创建用户代理生成器
func NewUserAgentGenerator() *UserAgentGenerator {
	return &UserAgentGenerator{
		browserVersions: map[string][]string{
			"Chrome": {
				"120.0.6099.109", "119.0.6045.199", "118.0.5993.117", "117.0.5938.149",
				"116.0.5845.187", "115.0.5790.171", "114.0.5735.198", "113.0.5672.126",
				"112.0.5615.137", "111.0.5563.146", "110.0.5481.177", "109.0.5414.119",
			},
			"Firefox": {
				"121.0", "120.0", "119.0", "118.0", "117.0", "116.0", "115.0",
				"114.0", "113.0", "112.0", "111.0", "110.0", "109.0",
			},
			"Safari": {
				"17.2", "17.1", "17.0", "16.6", "16.5", "16.4", "16.3",
				"16.2", "16.1", "16.0", "15.6", "15.5", "15.4",
			},
			"Edge": {
				"120.0.2210.91", "119.0.2151.97", "118.0.2088.76", "117.0.2045.47",
				"116.0.1938.81", "115.0.1901.200", "114.0.1823.79", "113.0.1774.57",
			},
			"Opera": {
				"105.0.4970.48", "104.0.4944.54", "103.0.4928.34", "102.0.4880.46",
				"101.0.4843.25", "100.0.4815.21", "99.0.4788.9", "98.0.4759.15",
			},
		},
		osVersions: map[Platform][]string{
			PlatformWindows: {
				"10.0", "11", "10", "8.1", "8", "7",
			},
			PlatformMacOS: {
				"14.2", "14.1", "14.0", "13.6", "13.5", "13.4", "13.3",
				"12.7", "12.6", "12.5", "11.7", "11.6", "10.15",
			},
			PlatformLinux: {
				"x86_64", "i686", "aarch64", "armv7l",
			},
			PlatformAndroid: {
				"14", "13", "12", "11", "10", "9.0", "8.1", "8.0", "7.1", "7.0",
			},
			PlatformIOS: {
				"17.2", "17.1", "17.0", "16.7", "16.6", "16.5", "16.4",
				"15.8", "15.7", "15.6", "14.8", "14.7",
			},
		},
		browserEngines: map[string]BrowserEngine{
			"Chrome":  {Name: "AppleWebKit", Version: "537.36"},
			"Edge":    {Name: "AppleWebKit", Version: "537.36"},
			"Opera":   {Name: "AppleWebKit", Version: "537.36"},
			"Safari":  {Name: "AppleWebKit", Version: "605.1.15"},
			"Firefox": {Name: "Gecko", Version: "20100101"},
		},
	}
}

// GenerateUserAgent 根据选项生成用户代理字符串
func (g *UserAgentGenerator) GenerateUserAgent(opts UserAgentOptions) string {
	// 如果没有指定浏览器，随机选择一个
	if opts.Browser == "" {
		browsers := []string{"Chrome", "Firefox", "Safari", "Edge", "Opera"}
		opts.Browser = randx.Choose(browsers)
	}

	// 如果没有指定平台，根据设备类型推断
	if opts.Platform == "" {
		opts.Platform = g.inferPlatformFromDevice(opts.DeviceType)
	}

	// 如果没有指定操作系统，根据平台推断
	if opts.OS == "" {
		opts.OS = g.inferOSFromPlatform(opts.Platform)
	}

	// 如果没有指定操作系统版本，随机生成
	if opts.OSVersion == "" {
		if versions, exists := g.osVersions[opts.Platform]; exists {
			opts.OSVersion = randx.Choose(versions)
		}
	}

	// 如果没有指定架构，使用默认值
	if opts.Architecture == "" {
		opts.Architecture = g.getDefaultArchitecture(opts.Platform)
	}

	// 生成浏览器版本
	browserVersion := g.getBrowserVersion(opts.Browser)

	// 根据平台和浏览器生成用户代理
	return g.buildUserAgent(opts, browserVersion)
}

// inferPlatformFromDevice 根据设备类型推断平台
func (g *UserAgentGenerator) inferPlatformFromDevice(deviceType DeviceType) Platform {
	switch deviceType {
	case DeviceTypeMobile:
		platforms := []Platform{PlatformAndroid, PlatformIOS}
		return randx.Choose(platforms)
	case DeviceTypeTablet:
		platforms := []Platform{PlatformAndroid, PlatformIOS}
		return randx.Choose(platforms)
	case DeviceTypeDesktop, DeviceTypeLaptop:
		platforms := []Platform{PlatformWindows, PlatformMacOS, PlatformLinux}
		return randx.Choose(platforms)
	default:
		return PlatformWindows
	}
}

// inferOSFromPlatform 根据平台推断操作系统
func (g *UserAgentGenerator) inferOSFromPlatform(platform Platform) string {
	switch platform {
	case PlatformWindows:
		return "Windows NT"
	case PlatformMacOS:
		return "Mac OS X"
	case PlatformLinux:
		return "Linux"
	case PlatformAndroid:
		return "Android"
	case PlatformIOS:
		return "iPhone OS"
	default:
		return "Windows NT"
	}
}

// getDefaultArchitecture 获取平台的默认架构
func (g *UserAgentGenerator) getDefaultArchitecture(platform Platform) string {
	switch platform {
	case PlatformWindows:
		return "WOW64" // 或者 "Win64; x64"
	case PlatformMacOS:
		return "Intel"
	case PlatformLinux:
		return "x86_64"
	case PlatformAndroid, PlatformIOS:
		return "Mobile"
	default:
		return "x64"
	}
}

// getBrowserVersion 获取浏览器版本
func (g *UserAgentGenerator) getBrowserVersion(browser string) string {
	if versions, exists := g.browserVersions[browser]; exists {
		return randx.Choose(versions)
	}
	// 如果没有预定义版本，生成一个随机版本
	return fmt.Sprintf("%d.%d.%d.%d",
		randx.Intn(50)+80, // 主版本号 80-130
		randx.Intn(10),    // 次版本号 0-9
		randx.Intn(9999),  // 修订版本号 0-9999
		randx.Intn(999),   // 构建号 0-999
	)
}

// buildUserAgent 构建用户代理字符串
func (g *UserAgentGenerator) buildUserAgent(opts UserAgentOptions, browserVersion string) string {
	switch opts.Platform {
	case PlatformWindows:
		return g.buildWindowsUserAgent(opts, browserVersion)
	case PlatformMacOS:
		return g.buildMacOSUserAgent(opts, browserVersion)
	case PlatformLinux:
		return g.buildLinuxUserAgent(opts, browserVersion)
	case PlatformAndroid:
		return g.buildAndroidUserAgent(opts, browserVersion)
	case PlatformIOS:
		return g.buildIOSUserAgent(opts, browserVersion)
	default:
		return g.buildGenericUserAgent(opts, browserVersion)
	}
}

// buildWindowsUserAgent 构建Windows用户代理
func (g *UserAgentGenerator) buildWindowsUserAgent(opts UserAgentOptions, browserVersion string) string {
	engine := g.browserEngines[opts.Browser]
	osVersion := opts.OSVersion

	// Windows版本映射
	if osVersion == "11" {
		osVersion = "10.0"
	} else if osVersion == "10" {
		osVersion = "10.0"
	}

	arch := "Win64; x64"
	if opts.Architecture == "WOW64" {
		arch = "WOW64"
	}

	switch opts.Browser {
	case "Chrome", "Edge", "Opera":
		return fmt.Sprintf("Mozilla/5.0 (Windows NT %s; %s) AppleWebKit/%s (KHTML, like Gecko) %s/%s Safari/537.36",
			osVersion, arch, engine.Version, opts.Browser, browserVersion)
	case "Firefox":
		return fmt.Sprintf("Mozilla/5.0 (Windows NT %s; %s; rv:%s) Gecko/%s Firefox/%s",
			osVersion, arch, browserVersion, engine.Version, browserVersion)
	case "Safari":
		return fmt.Sprintf("Mozilla/5.0 (Windows NT %s; %s) AppleWebKit/%s (KHTML, like Gecko) Version/%s Safari/%s",
			osVersion, arch, engine.Version, browserVersion, engine.Version)
	default:
		return fmt.Sprintf("Mozilla/5.0 (Windows NT %s; %s) %s/%s",
			osVersion, arch, opts.Browser, browserVersion)
	}
}

// buildMacOSUserAgent 构建macOS用户代理
func (g *UserAgentGenerator) buildMacOSUserAgent(opts UserAgentOptions, browserVersion string) string {
	engine := g.browserEngines[opts.Browser]
	osVersion := strings.ReplaceAll(opts.OSVersion, ".", "_")

	switch opts.Browser {
	case "Chrome", "Edge", "Opera":
		return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X %s) AppleWebKit/%s (KHTML, like Gecko) %s/%s Safari/537.36",
			osVersion, engine.Version, opts.Browser, browserVersion)
	case "Firefox":
		return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X %s) Gecko/%s Firefox/%s",
			osVersion, engine.Version, browserVersion)
	case "Safari":
		return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X %s) AppleWebKit/%s (KHTML, like Gecko) Version/%s Safari/%s",
			osVersion, engine.Version, browserVersion, engine.Version)
	default:
		return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X %s) %s/%s",
			osVersion, opts.Browser, browserVersion)
	}
}

// buildLinuxUserAgent 构建Linux用户代理
func (g *UserAgentGenerator) buildLinuxUserAgent(opts UserAgentOptions, browserVersion string) string {
	engine := g.browserEngines[opts.Browser]
	arch := opts.Architecture

	switch opts.Browser {
	case "Chrome", "Edge", "Opera":
		return fmt.Sprintf("Mozilla/5.0 (X11; Linux %s) AppleWebKit/%s (KHTML, like Gecko) %s/%s Safari/537.36",
			arch, engine.Version, opts.Browser, browserVersion)
	case "Firefox":
		return fmt.Sprintf("Mozilla/5.0 (X11; Linux %s; rv:%s) Gecko/%s Firefox/%s",
			arch, browserVersion, engine.Version, browserVersion)
	default:
		return fmt.Sprintf("Mozilla/5.0 (X11; Linux %s) %s/%s",
			arch, opts.Browser, browserVersion)
	}
}

// buildAndroidUserAgent 构建Android用户代理
func (g *UserAgentGenerator) buildAndroidUserAgent(opts UserAgentOptions, browserVersion string) string {
	engine := g.browserEngines[opts.Browser]

	// 随机生成设备型号
	deviceModels := []string{
		"SM-G973F", "SM-G981B", "SM-G998B", "SM-A525F", "SM-A715F",
		"Pixel 6", "Pixel 7", "OnePlus 9", "Mi 11", "HUAWEI P40",
	}
	deviceModel := randx.Choose(deviceModels)

	switch opts.Browser {
	case "Chrome":
		mobileStr := ""
		if opts.Mobile {
			mobileStr = " Mobile"
		}
		return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; %s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s%s Safari/537.36",
			opts.OSVersion, deviceModel, engine.Version, browserVersion, mobileStr)
	case "Firefox":
		mobileStr := ""
		if opts.Mobile {
			mobileStr = " Mobile"
		}
		return fmt.Sprintf("Mozilla/5.0 (Mobile; rv:%s) Gecko/%s Firefox/%s%s",
			browserVersion, engine.Version, browserVersion, mobileStr)
	default:
		return fmt.Sprintf("Mozilla/5.0 (Linux; Android %s; %s) %s/%s",
			opts.OSVersion, deviceModel, opts.Browser, browserVersion)
	}
}

// buildIOSUserAgent 构建iOS用户代理
func (g *UserAgentGenerator) buildIOSUserAgent(opts UserAgentOptions, browserVersion string) string {
	engine := g.browserEngines[opts.Browser]
	osVersion := strings.ReplaceAll(opts.OSVersion, ".", "_")

	deviceType := "iPhone"
	if opts.DeviceType == DeviceTypeTablet {
		deviceType = "iPad"
	}

	cpuType := fmt.Sprintf("CPU %s OS", deviceType)
	if deviceType == "iPad" {
		cpuType = "CPU OS"
	}

	switch opts.Browser {
	case "Safari", "Chrome", "Edge", "Opera":
		return fmt.Sprintf("Mozilla/5.0 (%s; %s %s like Mac OS X) AppleWebKit/%s (KHTML, like Gecko) Version/%s Mobile/15E148 Safari/604.1",
			deviceType, cpuType, osVersion, engine.Version, browserVersion)
	case "Firefox":
		return fmt.Sprintf("Mozilla/5.0 (%s; %s %s like Mac OS X) Gecko/%s FxiOS/%s.0 Mobile/15E148 Safari/605.1.15",
			deviceType, cpuType, osVersion, engine.Version, browserVersion)
	default:
		return fmt.Sprintf("Mozilla/5.0 (%s; %s %s like Mac OS X) %s/%s",
			deviceType, cpuType, osVersion, opts.Browser, browserVersion)
	}
}

// buildGenericUserAgent 构建通用用户代理
func (g *UserAgentGenerator) buildGenericUserAgent(opts UserAgentOptions, browserVersion string) string {
	return fmt.Sprintf("Mozilla/5.0 (%s %s) %s/%s",
		opts.Platform, opts.OSVersion, opts.Browser, browserVersion)
}

// 全局用户代理生成器实例
var defaultUserAgentGen = NewUserAgentGenerator()

// GenerateUserAgent 使用默认生成器生成用户代理
func (f *Faker) GenerateUserAgent(opts UserAgentOptions) string {
	f.incrementCallCount()
	return defaultUserAgentGen.GenerateUserAgent(opts)
}

// GenerateRandomUserAgent 生成随机用户代理
func (f *Faker) GenerateRandomUserAgent() string {
	f.incrementCallCount()

	// 随机选择设备类型
	deviceTypes := []DeviceType{DeviceTypeDesktop, DeviceTypeLaptop, DeviceTypeMobile, DeviceTypeTablet}
	deviceType := randx.Choose(deviceTypes)

	// 根据设备类型设置mobile选项
	mobile := deviceType == DeviceTypeMobile || deviceType == DeviceTypeTablet

	opts := UserAgentOptions{
		DeviceType: deviceType,
		Mobile:     mobile,
	}

	return defaultUserAgentGen.GenerateUserAgent(opts)
}

// 便捷方法 - 指定浏览器生成用户代理
func (f *Faker) UserAgentFor(browser string) string {
	f.incrementCallCount()
	opts := UserAgentOptions{
		Browser: browser,
	}
	return defaultUserAgentGen.GenerateUserAgent(opts)
}

// 便捷方法 - 指定平台生成用户代理
func (f *Faker) UserAgentForPlatform(platform Platform) string {
	f.incrementCallCount()
	opts := UserAgentOptions{
		Platform: platform,
	}
	return defaultUserAgentGen.GenerateUserAgent(opts)
}

// 便捷方法 - 指定设备类型生成用户代理
func (f *Faker) UserAgentForDevice(deviceType DeviceType) string {
	f.incrementCallCount()
	opts := UserAgentOptions{
		DeviceType: deviceType,
		Mobile:     deviceType == DeviceTypeMobile || deviceType == DeviceTypeTablet,
	}
	return defaultUserAgentGen.GenerateUserAgent(opts)
}

// 便捷方法 - Chrome用户代理
func (f *Faker) ChromeUserAgent() string {
	f.incrementCallCount()
	opts := UserAgentOptions{
		Browser: "Chrome",
	}
	return defaultUserAgentGen.GenerateUserAgent(opts)
}

// 便捷方法 - Firefox用户代理
func (f *Faker) FirefoxUserAgent() string {
	f.incrementCallCount()
	opts := UserAgentOptions{
		Browser: "Firefox",
	}
	return defaultUserAgentGen.GenerateUserAgent(opts)
}

// 便捷方法 - Safari用户代理
func (f *Faker) SafariUserAgent() string {
	f.incrementCallCount()
	opts := UserAgentOptions{
		Browser: "Safari",
	}
	return defaultUserAgentGen.GenerateUserAgent(opts)
}

// 便捷方法 - Edge用户代理
func (f *Faker) EdgeUserAgent() string {
	f.incrementCallCount()
	opts := UserAgentOptions{
		Browser: "Edge",
	}
	return defaultUserAgentGen.GenerateUserAgent(opts)
}

// 便捷方法 - Android设备用户代理
func (f *Faker) AndroidUserAgent() string {
	f.incrementCallCount()
	opts := UserAgentOptions{
		Platform:   PlatformAndroid,
		DeviceType: DeviceTypeMobile,
		Mobile:     true,
	}
	return defaultUserAgentGen.GenerateUserAgent(opts)
}

// 便捷方法 - iOS设备用户代理
func (f *Faker) IOSUserAgent() string {
	f.incrementCallCount()
	opts := UserAgentOptions{
		Platform:   PlatformIOS,
		DeviceType: DeviceTypeMobile,
		Mobile:     true,
	}
	return defaultUserAgentGen.GenerateUserAgent(opts)
}

// 全局便捷函数
func GenerateUserAgent(opts UserAgentOptions) string {
	return getDefaultFaker().GenerateUserAgent(opts)
}

func GenerateRandomUserAgent() string {
	return getDefaultFaker().GenerateRandomUserAgent()
}

func UserAgentFor(browser string) string {
	return getDefaultFaker().UserAgentFor(browser)
}

func UserAgentForPlatform(platform Platform) string {
	return getDefaultFaker().UserAgentForPlatform(platform)
}

func UserAgentForDevice(deviceType DeviceType) string {
	return getDefaultFaker().UserAgentForDevice(deviceType)
}

func ChromeUserAgent() string {
	return getDefaultFaker().ChromeUserAgent()
}

func FirefoxUserAgent() string {
	return getDefaultFaker().FirefoxUserAgent()
}

func SafariUserAgent() string {
	return getDefaultFaker().SafariUserAgent()
}

func EdgeUserAgent() string {
	return getDefaultFaker().EdgeUserAgent()
}

func AndroidUserAgent() string {
	return getDefaultFaker().AndroidUserAgent()
}

func IOSUserAgent() string {
	return getDefaultFaker().IOSUserAgent()
}

func UserAgent() string {
	return getDefaultFaker().UserAgent()
}
