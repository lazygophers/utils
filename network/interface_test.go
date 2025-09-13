package network

import (
	"net"
	"testing"
)

func TestGetInterfaceIpByName(t *testing.T) {
	tests := []struct {
		name          string
		interfaceName string
		prev6         bool
		shouldContain string // 部分匹配，因为实际IP可能变化
	}{
		{"eth0 IPv4", "eth0", false, ""},                // 可能不存在，返回空字符串
		{"eth0 IPv6", "eth0", true, ""},                 // 可能不存在，返回空字符串
		{"invalid interface", "nonexistent", false, ""}, // 不存在的接口
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetInterfaceIpByName(tt.interfaceName, tt.prev6)
			// 由于网络接口的存在性不确定，我们主要测试函数不会panic
			// 结果应该是空字符串或有效IP地址
			if result != "" {
				// 如果返回非空，应该是有效IP
				if net.ParseIP(result) == nil {
					t.Errorf("GetInterfaceIpByName(%s, %v) returned invalid IP: %s",
						tt.interfaceName, tt.prev6, result)
				}
			}
		})
	}
}

func TestGetInterfaceIpByAddrs(t *testing.T) {
	// 创建模拟地址列表
	createIPNet := func(ipStr string) *net.IPNet {
		ip, ipnet, _ := net.ParseCIDR(ipStr)
		ipnet.IP = ip // 使用实际IP而非网络地址
		return ipnet
	}

	tests := []struct {
		name      string
		addresses []net.Addr
		prev6     bool
		expected  string
	}{
		{
			name:      "IPv4 only, prefer IPv4",
			addresses: []net.Addr{createIPNet("192.168.1.100/24")},
			prev6:     false,
			expected:  "192.168.1.100",
		},
		{
			name:      "IPv6 only, prefer IPv6",
			addresses: []net.Addr{createIPNet("2001:db8::1/64")},
			prev6:     true,
			expected:  "2001:db8::1",
		},
		{
			name: "Both IPv4 and IPv6, prefer IPv4",
			addresses: []net.Addr{
				createIPNet("192.168.1.100/24"),
				createIPNet("2001:db8::1/64"),
			},
			prev6:    false,
			expected: "192.168.1.100",
		},
		{
			name: "Both IPv4 and IPv6, prefer IPv6",
			addresses: []net.Addr{
				createIPNet("192.168.1.100/24"),
				createIPNet("2001:db8::1/64"),
			},
			prev6:    true,
			expected: "2001:db8::1",
		},
		{
			name: "IPv4 only, prefer IPv6 (fallback to IPv4)",
			addresses: []net.Addr{
				createIPNet("192.168.1.100/24"),
			},
			prev6:    true,
			expected: "192.168.1.100",
		},
		{
			name:      "Loopback address (should be ignored)",
			addresses: []net.Addr{createIPNet("127.0.0.1/8")},
			prev6:     false,
			expected:  "",
		},
		{
			name:      "IPv6 loopback (should be ignored)",
			addresses: []net.Addr{createIPNet("::1/128")},
			prev6:     true,
			expected:  "",
		},
		{
			name:      "Empty address list",
			addresses: []net.Addr{},
			prev6:     false,
			expected:  "",
		},
		{
			name: "Mixed with loopback",
			addresses: []net.Addr{
				createIPNet("127.0.0.1/8"),      // loopback, should be ignored
				createIPNet("192.168.1.100/24"), // valid IPv4
			},
			prev6:    false,
			expected: "192.168.1.100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetInterfaceIpByAddrs(tt.addresses, tt.prev6)
			if result != tt.expected {
				t.Errorf("GetInterfaceIpByAddrs() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

// mockAddr implements net.Addr for testing
type mockAddr struct{}

func (m mockAddr) Network() string { return "mock" }
func (m mockAddr) String() string  { return "mock-address" }

func TestGetInterfaceIpByAddrs_InvalidAddr(t *testing.T) {
	// 测试非IPNet类型的地址
	addresses := []net.Addr{mockAddr{}}
	result := GetInterfaceIpByAddrs(addresses, false)
	if result != "" {
		t.Errorf("Expected empty result for invalid address type, got %s", result)
	}
}

func TestGetListenIp(t *testing.T) {
	tests := []struct {
		name  string
		prev6 []bool
	}{
		{"Default IPv4", []bool{}},
		{"Explicit IPv4", []bool{false}},
		{"Explicit IPv6", []bool{true}},
		{"Multiple args (first used)", []bool{true, false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetListenIp(tt.prev6...)
			// 结果应该是空字符串或有效IP地址
			if result != "" {
				if net.ParseIP(result) == nil {
					t.Errorf("GetListenIp() returned invalid IP: %s", result)
				}
			}
			// 函数不应该panic，即使没有找到合适的接口
		})
	}
}

func TestGetInterfaceIpByName_ErrorHandling(t *testing.T) {
	tests := []struct {
		name          string
		interfaceName string
		prev6         bool
	}{
		{"Valid interface eth0 IPv4", "eth0", false},
		{"Valid interface eth0 IPv6", "eth0", true},
		{"Valid interface en0 IPv4", "en0", false},
		{"Valid interface en0 IPv6", "en0", true},
		{"Invalid interface", "definitely_nonexistent_interface_12345", false},
		{"Empty interface name", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetInterfaceIpByName(tt.interfaceName, tt.prev6)
			// 对于无效接口，应该返回空字符串
			// 对于有效接口，可能返回IP或空字符串（如果没有配置该类型的IP）
			if result != "" {
				if net.ParseIP(result) == nil {
					t.Errorf("GetInterfaceIpByName(%s, %v) returned invalid IP: %s",
						tt.interfaceName, tt.prev6, result)
				}
			}
		})
	}
}
