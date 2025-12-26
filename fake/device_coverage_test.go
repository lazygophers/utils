package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试设备信息生成
func TestDeviceInfo(t *testing.T) {
	f := New()
	device := f.DeviceInfo()
	assert.NotNil(t, device)
	assert.NotEmpty(t, device.Type)
	assert.NotEmpty(t, device.Platform)
	assert.NotEmpty(t, device.Model)
	assert.NotEmpty(t, device.Manufacturer)
	assert.NotEmpty(t, device.OS)
	assert.NotEmpty(t, device.OSVersion)
	assert.NotEmpty(t, device.Browser)
	assert.NotEmpty(t, device.Version)
	assert.NotEmpty(t, device.UserAgent)
	assert.Greater(t, device.ScreenWidth, 0)
	assert.Greater(t, device.ScreenHeight, 0)
}

// 测试生成各种类型的设备信息
func TestGenerateDifferentDeviceTypes(t *testing.T) {
	f := New()

	// 测试桌面设备
	desktopDevice := f.generateDesktopDevice(DeviceTypeDesktop)
	assert.NotNil(t, desktopDevice)
	assert.Equal(t, string(DeviceTypeDesktop), desktopDevice.Type)

	// 测试笔记本设备
	laptopDevice := f.generateDesktopDevice(DeviceTypeLaptop)
	assert.NotNil(t, laptopDevice)
	assert.Equal(t, string(DeviceTypeLaptop), laptopDevice.Type)

	// 测试移动设备
	mobileDevice := f.generateMobileDevice()
	assert.NotNil(t, mobileDevice)
	assert.Equal(t, string(DeviceTypeMobile), mobileDevice.Type)

	// 测试平板设备
	tabletDevice := f.generateTabletDevice()
	assert.NotNil(t, tabletDevice)
	assert.Equal(t, string(DeviceTypeTablet), tabletDevice.Type)

	// 测试手表设备
	watchDevice := f.generateWatchDevice()
	assert.NotNil(t, watchDevice)
	assert.Equal(t, string(DeviceTypeWatch), watchDevice.Type)

	// 测试电视设备
	tvDevice := f.generateTVDevice()
	assert.NotNil(t, tvDevice)
	assert.Equal(t, string(DeviceTypeTV), tvDevice.Type)
}

// 测试移动设备用户代理生成
func TestMobileUserAgent(t *testing.T) {
	f := New()
	userAgent := f.MobileUserAgent()
	assert.NotEmpty(t, userAgent)
}

// 测试桌面设备用户代理生成
func TestDesktopUserAgent(t *testing.T) {
	f := New()
	userAgent := f.DesktopUserAgent()
	assert.NotEmpty(t, userAgent)
}

// 测试浏览器名称生成
func TestBrowser(t *testing.T) {
	f := New()
	browser := f.Browser()
	assert.NotEmpty(t, browser)
}

// 测试操作系统名称生成
func TestOS(t *testing.T) {
	f := New()
	os := f.OS()
	assert.NotEmpty(t, os)
}

// 测试批量生成设备信息
func TestBatchDeviceInfos(t *testing.T) {
	f := New()
	devices := f.BatchDeviceInfos(5)
	assert.Len(t, devices, 5)
	for _, device := range devices {
		assert.NotNil(t, device)
	}
}

// 测试批量生成用户代理
func TestBatchUserAgents(t *testing.T) {
	f := New()
	userAgents := f.BatchUserAgents(5)
	assert.Len(t, userAgents, 5)
	for _, ua := range userAgents {
		assert.NotEmpty(t, ua)
	}
}

// 测试全局函数
func TestGlobalDeviceFunctions(t *testing.T) {
	// 测试全局DeviceInfo函数
	device := DeviceInfo()
	assert.NotNil(t, device)

	// 测试全局MobileUserAgent函数
	mobileUA := MobileUserAgent()
	assert.NotEmpty(t, mobileUA)

	// 测试全局DesktopUserAgent函数
	desktopUA := DesktopUserAgent()
	assert.NotEmpty(t, desktopUA)

	// 测试全局Browser函数
	browser := Browser()
	assert.NotEmpty(t, browser)

	// 测试全局OS函数
	os := OS()
	assert.NotEmpty(t, os)
}

// 测试设备生成函数的各种分支
func TestDeviceGenerationBranches(t *testing.T) {
	f := New()

	// 测试桌面设备的各种平台分支
	windowsDesktop := f.generateDesktopDevice(DeviceTypeDesktop)
	assert.True(t, windowsDesktop.Platform == string(PlatformWindows) ||
		windowsDesktop.Platform == string(PlatformMacOS) ||
		windowsDesktop.Platform == string(PlatformLinux))

	// 测试移动设备的各种平台分支
	mobileDevice := f.generateMobileDevice()
	assert.True(t, mobileDevice.Platform == string(PlatformAndroid) ||
		mobileDevice.Platform == string(PlatformIOS))

	// 测试平板设备的各种平台分支
	tabletDevice := f.generateTabletDevice()
	assert.True(t, tabletDevice.Platform == string(PlatformAndroid) ||
		tabletDevice.Platform == string(PlatformIOS))

	// 测试手表设备的各种平台分支
	watchDevice := f.generateWatchDevice()
	assert.True(t, watchDevice.Platform == string(PlatformAndroid) ||
		watchDevice.Platform == string(PlatformIOS))

	// 测试电视设备的各种平台分支
	tvDevice := f.generateTVDevice()
	assert.NotEmpty(t, tvDevice.OS)
}
