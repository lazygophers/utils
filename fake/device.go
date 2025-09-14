package fake

import (
	"fmt"

	"github.com/lazygophers/utils/randx"
)

// Device 设备信息结构体
type Device struct {
	Type         string `json:"type"`
	Platform     string `json:"platform"`
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer"`
	OS           string `json:"os"`
	OSVersion    string `json:"os_version"`
	Browser      string `json:"browser"`
	Version      string `json:"version"`
	UserAgent    string `json:"user_agent"`
	ScreenWidth  int    `json:"screen_width"`
	ScreenHeight int    `json:"screen_height"`
}

// DeviceType 设备类型
type DeviceType string

const (
	DeviceTypeDesktop DeviceType = "desktop"
	DeviceTypeLaptop  DeviceType = "laptop"
	DeviceTypeTablet  DeviceType = "tablet"
	DeviceTypeMobile  DeviceType = "mobile"
	DeviceTypeWatch   DeviceType = "watch"
	DeviceTypeTV      DeviceType = "tv"
	DeviceTypeConsole DeviceType = "console"
)

// Platform 平台类型
type Platform string

const (
	PlatformWindows Platform = "Windows"
	PlatformMacOS   Platform = "macOS"
	PlatformLinux   Platform = "Linux"
	PlatformAndroid Platform = "Android"
	PlatformIOS     Platform = "iOS"
	PlatformWeb     Platform = "Web"
)

// DeviceInfo 生成设备信息
func (f *Faker) DeviceInfo() *Device {
	f.incrementCallCount()
	
	deviceTypes := []DeviceType{
		DeviceTypeDesktop, DeviceTypeLaptop, DeviceTypeTablet, 
		DeviceTypeMobile, DeviceTypeWatch, DeviceTypeTV,
	}
	
	deviceType := randx.Choose(deviceTypes)
	
	switch deviceType {
	case DeviceTypeDesktop, DeviceTypeLaptop:
		return f.generateDesktopDevice(deviceType)
	case DeviceTypeTablet:
		return f.generateTabletDevice()
	case DeviceTypeMobile:
		return f.generateMobileDevice()
	case DeviceTypeWatch:
		return f.generateWatchDevice()
	case DeviceTypeTV:
		return f.generateTVDevice()
	default:
		return f.generateDesktopDevice(DeviceTypeDesktop)
	}
}

func (f *Faker) generateDesktopDevice(deviceType DeviceType) *Device {
	platforms := []Platform{PlatformWindows, PlatformMacOS, PlatformLinux}
	platform := randx.Choose(platforms)
	
	var manufacturer, model, os, osVersion, browser, version string
	var screenWidth, screenHeight int
	
	switch platform {
	case PlatformWindows:
		manufacturer = randx.Choose([]string{"Dell", "HP", "Lenovo", "ASUS", "Acer", "MSI", "Custom Build"})
		models := []string{"OptiPlex", "Pavilion", "ThinkPad", "VivoBook", "Aspire", "Gaming", "Workstation"}
		model = randx.Choose(models)
		os = "Windows"
		versions := []string{"10", "11"}
		osVersion = randx.Choose(versions)
		
	case PlatformMacOS:
		manufacturer = "Apple"
		if deviceType == DeviceTypeLaptop {
			models := []string{"MacBook Air", "MacBook Pro"}
			model = randx.Choose(models)
		} else {
			models := []string{"iMac", "Mac Studio", "Mac Pro", "Mac mini"}
			model = randx.Choose(models)
		}
		os = "macOS"
		osVersion = fmt.Sprintf("%d.%d", randx.Intn(5)+10, randx.Intn(10))
		
	case PlatformLinux:
		manufacturer = randx.Choose([]string{"Dell", "HP", "Lenovo", "System76", "Custom Build"})
		model = randx.Choose([]string{"XPS", "ThinkPad", "Galago", "Oryx", "Custom"})
		os = randx.Choose([]string{"Ubuntu", "Fedora", "Arch", "Debian", "openSUSE", "Manjaro"})
		osVersion = fmt.Sprintf("%d.%d", randx.Intn(10)+20, randx.Intn(12)+1)
	}
	
	// 浏览器信息
	browsers := []string{"Chrome", "Firefox", "Safari", "Edge", "Opera"}
	browser = randx.Choose(browsers)
	version = fmt.Sprintf("%d.%d.%d", randx.Intn(100)+80, randx.Intn(10), randx.Intn(1000))
	
	// 屏幕分辨率
	resolutions := [][]int{
		{1920, 1080}, {2560, 1440}, {3840, 2160}, {1366, 768}, {1440, 900},
		{2880, 1800}, {3440, 1440}, {5120, 2880},
	}
	resolution := randx.Choose(resolutions)
	screenWidth, screenHeight = resolution[0], resolution[1]
	
	// 生成用户代理字符串
	userAgent := f.generateUserAgentForDevice(platform, os, osVersion, browser, version, deviceType)
	
	return &Device{
		Type:         string(deviceType),
		Platform:     string(platform),
		Model:        model,
		Manufacturer: manufacturer,
		OS:           os,
		OSVersion:    osVersion,
		Browser:      browser,
		Version:      version,
		UserAgent:    userAgent,
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}
}

func (f *Faker) generateMobileDevice() *Device {
	platforms := []Platform{PlatformAndroid, PlatformIOS}
	platform := randx.Choose(platforms)
	
	var manufacturer, model, os, osVersion, browser, version string
	var screenWidth, screenHeight int
	
	switch platform {
	case PlatformAndroid:
		manufacturers := []string{"Samsung", "Google", "OnePlus", "Xiaomi", "Huawei", "Sony", "LG", "Motorola"}
		manufacturer = randx.Choose(manufacturers)
		
		switch manufacturer {
		case "Samsung":
			models := []string{"Galaxy S23", "Galaxy S22", "Galaxy Note 20", "Galaxy A54", "Galaxy A34"}
			model = randx.Choose(models)
		case "Google":
			models := []string{"Pixel 7", "Pixel 6", "Pixel 7a", "Pixel 6a"}
			model = randx.Choose(models)
		case "OnePlus":
			models := []string{"OnePlus 11", "OnePlus 10T", "OnePlus Nord"}
			model = randx.Choose(models)
		default:
			model = fmt.Sprintf("%s-%d", manufacturer, randx.Intn(900)+100)
		}
		
		os = "Android"
		osVersion = fmt.Sprintf("%d.%d", randx.Intn(5)+9, randx.Intn(10))
		browser = randx.Choose([]string{"Chrome", "Samsung Internet", "Firefox", "Opera"})
		
	case PlatformIOS:
		manufacturer = "Apple"
		models := []string{"iPhone 14", "iPhone 13", "iPhone 12", "iPhone SE", "iPhone 14 Pro", "iPhone 13 Pro"}
		model = randx.Choose(models)
		os = "iOS"
		osVersion = fmt.Sprintf("%d.%d", randx.Intn(5)+13, randx.Intn(8))
		browser = randx.Choose([]string{"Safari", "Chrome", "Firefox", "Opera"})
	}
	
	version = fmt.Sprintf("%d.%d.%d", randx.Intn(100)+80, randx.Intn(10), randx.Intn(1000))
	
	// 手机屏幕分辨率
	resolutions := [][]int{
		{375, 812}, {390, 844}, {414, 896}, {360, 800}, {412, 915},
		{375, 667}, {414, 736}, {320, 568}, {360, 640}, {411, 731},
	}
	resolution := randx.Choose(resolutions)
	screenWidth, screenHeight = resolution[0], resolution[1]
	
	userAgent := f.generateUserAgentForDevice(platform, os, osVersion, browser, version, DeviceTypeMobile)
	
	return &Device{
		Type:         string(DeviceTypeMobile),
		Platform:     string(platform),
		Model:        model,
		Manufacturer: manufacturer,
		OS:           os,
		OSVersion:    osVersion,
		Browser:      browser,
		Version:      version,
		UserAgent:    userAgent,
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}
}

func (f *Faker) generateTabletDevice() *Device {
	platforms := []Platform{PlatformAndroid, PlatformIOS}
	platform := randx.Choose(platforms)
	
	var manufacturer, model, os, osVersion string
	
	switch platform {
	case PlatformAndroid:
		manufacturers := []string{"Samsung", "Google", "Lenovo", "Huawei", "Amazon"}
		manufacturer = randx.Choose(manufacturers)
		
		switch manufacturer {
		case "Samsung":
			model = randx.Choose([]string{"Galaxy Tab S8", "Galaxy Tab A8", "Galaxy Tab S7"})
		case "Google":
			model = "Pixel Tablet"
		case "Amazon":
			model = randx.Choose([]string{"Fire HD 10", "Fire HD 8", "Fire 7"})
		default:
			model = fmt.Sprintf("%s Tab", manufacturer)
		}
		
		os = "Android"
		osVersion = fmt.Sprintf("%d.%d", randx.Intn(5)+9, randx.Intn(10))
		
	case PlatformIOS:
		manufacturer = "Apple"
		model = randx.Choose([]string{"iPad Pro", "iPad Air", "iPad", "iPad mini"})
		os = "iPadOS"
		osVersion = fmt.Sprintf("%d.%d", randx.Intn(5)+13, randx.Intn(8))
	}
	
	browser := randx.Choose([]string{"Chrome", "Safari", "Firefox", "Opera"})
	version := fmt.Sprintf("%d.%d.%d", randx.Intn(100)+80, randx.Intn(10), randx.Intn(1000))
	
	// 平板屏幕分辨率
	resolutions := [][]int{
		{768, 1024}, {810, 1080}, {1200, 1920}, {1536, 2048}, {834, 1194},
		{1024, 1366}, {800, 1280}, {1600, 2560},
	}
	resolution := randx.Choose(resolutions)
	screenWidth, screenHeight := resolution[0], resolution[1]
	
	userAgent := f.generateUserAgentForDevice(platform, os, osVersion, browser, version, DeviceTypeTablet)
	
	return &Device{
		Type:         string(DeviceTypeTablet),
		Platform:     string(platform),
		Model:        model,
		Manufacturer: manufacturer,
		OS:           os,
		OSVersion:    osVersion,
		Browser:      browser,
		Version:      version,
		UserAgent:    userAgent,
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}
}

func (f *Faker) generateWatchDevice() *Device {
	platforms := []Platform{PlatformAndroid, PlatformIOS}
	platform := randx.Choose(platforms)
	
	var manufacturer, model, os, osVersion string
	
	switch platform {
	case PlatformAndroid:
		manufacturer = randx.Choose([]string{"Samsung", "Fossil", "Garmin", "Fitbit"})
		model = randx.Choose([]string{"Galaxy Watch", "Gear S3", "Vivoactive", "Versa"})
		os = "Wear OS"
		osVersion = fmt.Sprintf("3.%d", randx.Intn(5))
		
	case PlatformIOS:
		manufacturer = "Apple"
		model = randx.Choose([]string{"Apple Watch Series 8", "Apple Watch SE", "Apple Watch Ultra"})
		os = "watchOS"
		osVersion = fmt.Sprintf("%d.%d", randx.Intn(3)+8, randx.Intn(5))
	}
	
	// 智能手表通常使用内置浏览器或简化版浏览器
	browser := "Watch Browser"
	version := "1.0"
	
	// 手表屏幕分辨率
	screenWidth, screenHeight := 396, 484 // Apple Watch typical resolution
	
	userAgent := fmt.Sprintf("WatchOS/%s (%s; %s %s)", osVersion, manufacturer, model, os)
	
	return &Device{
		Type:         string(DeviceTypeWatch),
		Platform:     string(platform),
		Model:        model,
		Manufacturer: manufacturer,
		OS:           os,
		OSVersion:    osVersion,
		Browser:      browser,
		Version:      version,
		UserAgent:    userAgent,
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}
}

func (f *Faker) generateTVDevice() *Device {
	manufacturers := []string{"Samsung", "LG", "Sony", "TCL", "Roku", "Amazon", "Google", "Apple"}
	manufacturer := randx.Choose(manufacturers)
	
	var model, os, osVersion string
	
	switch manufacturer {
	case "Samsung":
		model = "Smart TV"
		os = "Tizen"
		osVersion = fmt.Sprintf("%d.%d", randx.Intn(3)+5, randx.Intn(5))
	case "LG":
		model = "webOS TV"
		os = "webOS"
		osVersion = fmt.Sprintf("%d.%d", randx.Intn(3)+5, randx.Intn(5))
	case "Roku":
		model = "Roku TV"
		os = "Roku OS"
		osVersion = fmt.Sprintf("%d.%d", randx.Intn(3)+10, randx.Intn(5))
	case "Apple":
		model = "Apple TV"
		os = "tvOS"
		osVersion = fmt.Sprintf("%d.%d", randx.Intn(3)+14, randx.Intn(5))
	default:
		model = "Android TV"
		os = "Android TV"
		osVersion = fmt.Sprintf("%d.%d", randx.Intn(3)+9, randx.Intn(5))
	}
	
	browser := "TV Browser"
	version := "1.0"
	
	// TV屏幕分辨率
	resolutions := [][]int{
		{1920, 1080}, {3840, 2160}, {2560, 1440}, {1366, 768},
	}
	resolution := randx.Choose(resolutions)
	screenWidth, screenHeight := resolution[0], resolution[1]
	
	userAgent := fmt.Sprintf("SmartTV/%s (%s; %s %s)", osVersion, manufacturer, model, os)
	
	return &Device{
		Type:         string(DeviceTypeTV),
		Platform:     "TV",
		Model:        model,
		Manufacturer: manufacturer,
		OS:           os,
		OSVersion:    osVersion,
		Browser:      browser,
		Version:      version,
		UserAgent:    userAgent,
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}
}

func (f *Faker) generateUserAgentForDevice(platform Platform, os, osVersion, browser, version string, deviceType DeviceType) string {
	// 使用新的用户代理生成器
	opts := UserAgentOptions{
		Platform:    platform,
		Browser:     browser,
		OS:          os,
		OSVersion:   osVersion,
		DeviceType:  deviceType,
		Mobile:      deviceType == DeviceTypeMobile || deviceType == DeviceTypeTablet,
	}
	
	return defaultUserAgentGen.GenerateUserAgent(opts)
}

// MobileUserAgent 生成移动端用户代理
func (f *Faker) MobileUserAgent() string {
	f.incrementCallCount()
	opts := UserAgentOptions{
		DeviceType: DeviceTypeMobile,
		Mobile:     true,
	}
	return defaultUserAgentGen.GenerateUserAgent(opts)
}

// DesktopUserAgent 生成桌面端用户代理
func (f *Faker) DesktopUserAgent() string {
	f.incrementCallCount()
	opts := UserAgentOptions{
		DeviceType: DeviceTypeDesktop,
		Mobile:     false,
	}
	return defaultUserAgentGen.GenerateUserAgent(opts)
}

// Browser 生成浏览器名称
func (f *Faker) Browser() string {
	f.incrementCallCount()
	browsers := []string{"Chrome", "Firefox", "Safari", "Edge", "Opera", "Internet Explorer", "Brave", "Vivaldi"}
	return randx.Choose(browsers)
}

// OS 生成操作系统名称
func (f *Faker) OS() string {
	f.incrementCallCount()
	systems := []string{"Windows", "macOS", "Linux", "Android", "iOS", "Unix", "FreeBSD", "Ubuntu", "Fedora", "CentOS"}
	return randx.Choose(systems)
}

// 批量生成函数
func (f *Faker) BatchDeviceInfos(count int) []*Device {
	results := make([]*Device, count)
	for i := 0; i < count; i++ {
		results[i] = f.DeviceInfo()
	}
	return results
}

func (f *Faker) BatchUserAgents(count int) []string {
	return f.batchGenerate(count, f.UserAgent)
}

// 全局便捷函数
func DeviceInfo() *Device {
	return getDefaultFaker().DeviceInfo()
}

func MobileUserAgent() string {
	return getDefaultFaker().MobileUserAgent()
}

func DesktopUserAgent() string {
	return getDefaultFaker().DesktopUserAgent()
}

func Browser() string {
	return getDefaultFaker().Browser()
}

func OS() string {
	return getDefaultFaker().OS()
}