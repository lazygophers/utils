package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDeviceInfoExtended 测试设备信息生成的扩展功能
func TestDeviceInfoExtended(t *testing.T) {
	faker := New()

	// 测试DeviceInfo函数
	t.Run("test_device_info", func(t *testing.T) {
		device := faker.DeviceInfo()
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
	})

	// 测试各种设备类型
	t.Run("test_device_types", func(t *testing.T) {
		// 测试多次调用，确保覆盖不同设备类型
		deviceTypes := make(map[string]bool)
		for i := 0; i < 20; i++ {
			device := faker.DeviceInfo()
			deviceTypes[device.Type] = true
		}
		
		// 确保至少生成了几种不同的设备类型
		assert.Greater(t, len(deviceTypes), 3)
	})

	// 测试批量生成设备信息
	t.Run("test_batch_device_infos", func(t *testing.T) {
		devices := faker.BatchDeviceInfos(5)
		assert.Len(t, devices, 5)
		for _, device := range devices {
			assert.NotNil(t, device)
			assert.NotEmpty(t, device.Type)
		}
	})

	// 测试批量生成用户代理
	t.Run("test_batch_user_agents", func(t *testing.T) {
		userAgents := faker.BatchUserAgents(5)
		assert.Len(t, userAgents, 5)
		for _, ua := range userAgents {
			assert.NotEmpty(t, ua)
		}
	})
}

// TestDeviceGenerationFunctions 测试各种设备生成函数
func TestDeviceGenerationFunctions(t *testing.T) {
	faker := New()

	// 测试生成桌面设备
	t.Run("test_generate_desktop_device", func(t *testing.T) {
		// 测试桌面设备
		desktop := faker.generateDesktopDevice(DeviceTypeDesktop)
		assert.NotNil(t, desktop)
		assert.Equal(t, string(DeviceTypeDesktop), desktop.Type)
		
		// 测试笔记本设备
		laptop := faker.generateDesktopDevice(DeviceTypeLaptop)
		assert.NotNil(t, laptop)
		assert.Equal(t, string(DeviceTypeLaptop), laptop.Type)
	})

	// 测试生成移动设备
	t.Run("test_generate_mobile_device", func(t *testing.T) {
		mobile := faker.generateMobileDevice()
		assert.NotNil(t, mobile)
		assert.Equal(t, string(DeviceTypeMobile), mobile.Type)
	})

	// 测试生成平板设备
	t.Run("test_generate_tablet_device", func(t *testing.T) {
		tablet := faker.generateTabletDevice()
		assert.NotNil(t, tablet)
		assert.Equal(t, string(DeviceTypeTablet), tablet.Type)
	})

	// 测试生成手表设备
	t.Run("test_generate_watch_device", func(t *testing.T) {
		watch := faker.generateWatchDevice()
		assert.NotNil(t, watch)
		assert.Equal(t, string(DeviceTypeWatch), watch.Type)
	})

	// 测试生成电视设备
	t.Run("test_generate_tv_device", func(t *testing.T) {
		tv := faker.generateTVDevice()
		assert.NotNil(t, tv)
		assert.Equal(t, string(DeviceTypeTV), tv.Type)
	})

	// 测试生成用户代理
	t.Run("test_generate_user_agent_for_device", func(t *testing.T) {
		opts := UserAgentOptions{
			Platform:   PlatformWindows,
			OS:         "Windows NT",
			OSVersion:  "10.0",
			Browser:    "Chrome",
			DeviceType: DeviceTypeDesktop,
		}
		
		ua := faker.generateUserAgentForDevice(
			opts.Platform,
			opts.OS,
			opts.OSVersion,
			opts.Browser,
			"120.0.0.0",
			opts.DeviceType,
		)
		assert.NotEmpty(t, ua)
	})
}

// TestUserAgentFunctions 测试用户代理生成函数
func TestUserAgentFunctions(t *testing.T) {
	faker := New()

	// 测试移动端用户代理生成
	t.Run("test_mobile_user_agent", func(t *testing.T) {
		ua := faker.MobileUserAgent()
		assert.NotEmpty(t, ua)
	})

	// 测试桌面端用户代理生成
	t.Run("test_desktop_user_agent", func(t *testing.T) {
		ua := faker.DesktopUserAgent()
		assert.NotEmpty(t, ua)
	})
}

// TestBrowserAndOSFunctions 测试浏览器和操作系统生成函数
func TestBrowserAndOSFunctions(t *testing.T) {
	faker := New()

	// 测试浏览器生成
	t.Run("test_browser", func(t *testing.T) {
		browser := faker.Browser()
		assert.NotEmpty(t, browser)
	})

	// 测试操作系统生成
	t.Run("test_os", func(t *testing.T) {
		os := faker.OS()
		assert.NotEmpty(t, os)
	})
}

// TestGlobalDeviceFunctionsExtended 测试全局设备生成函数
func TestGlobalDeviceFunctionsExtended(t *testing.T) {
	// 测试全局DeviceInfo函数
	t.Run("test_global_device_info", func(t *testing.T) {
		device := DeviceInfo()
		assert.NotNil(t, device)
		assert.NotEmpty(t, device.Type)
	})

	// 测试全局MobileUserAgent函数
	t.Run("test_global_mobile_user_agent", func(t *testing.T) {
		ua := MobileUserAgent()
		assert.NotEmpty(t, ua)
	})

	// 测试全局DesktopUserAgent函数
	t.Run("test_global_desktop_user_agent", func(t *testing.T) {
		ua := DesktopUserAgent()
		assert.NotEmpty(t, ua)
	})

	// 测试全局Browser函数
	t.Run("test_global_browser", func(t *testing.T) {
		browser := Browser()
		assert.NotEmpty(t, browser)
	})

	// 测试全局OS函数
	t.Run("test_global_os", func(t *testing.T) {
		os := OS()
		assert.NotEmpty(t, os)
	})
}
