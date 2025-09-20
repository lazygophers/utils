package network

import (
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNetworkModuleCompleteCoverage focuses on hitting all uncovered lines
func TestNetworkModuleCompleteCoverage(t *testing.T) {
	// Test GetInterfaceIpByName error handling specifically
	t.Run("GetInterfaceIpByName_ErrorHandling", func(t *testing.T) {
		// Test with definitely non-existent interface to trigger InterfaceByName error
		result := GetInterfaceIpByName("definitely_nonexistent_interface_12345_xyz", false)
		assert.Equal(t, "", result, "Non-existent interface should return empty string")

		result = GetInterfaceIpByName("definitely_nonexistent_interface_12345_xyz", true)
		assert.Equal(t, "", result, "Non-existent interface should return empty string")
	})

	// Test RealIpFromHeader with all possible header combinations
	t.Run("RealIpFromHeader_CompleteCoverage", func(t *testing.T) {
		// Test with empty headers
		header := make(http.Header)
		result := RealIpFromHeader(header)
		assert.Equal(t, "", result, "Empty headers should return empty string")

		// Test Cf-Connecting-Ip with local IP (should be skipped)
		header.Set("Cf-Connecting-Ip", "192.168.1.1")
		result = RealIpFromHeader(header)
		assert.Equal(t, "", result, "Local IP should be skipped")

		// Test Cf-Connecting-Ip with valid public IP
		header = make(http.Header)
		header.Set("Cf-Connecting-Ip", "8.8.8.8")
		result = RealIpFromHeader(header)
		assert.Equal(t, "8.8.8.8", result, "Should return public IP")

		// Test Cf-Pseudo-Ipv4
		header = make(http.Header)
		header.Set("Cf-Pseudo-Ipv4", "1.1.1.1")
		result = RealIpFromHeader(header)
		assert.Equal(t, "1.1.1.1", result, "Should return public IPv4")

		// Test Cf-Connecting-Ipv6
		header = make(http.Header)
		header.Set("Cf-Connecting-Ipv6", "2001:4860:4860::8888")
		result = RealIpFromHeader(header)
		assert.Equal(t, "2001:4860:4860::8888", result, "Should return public IPv6")

		// Test Cf-Pseudo-Ipv6
		header = make(http.Header)
		header.Set("Cf-Pseudo-Ipv6", "2001:4860:4860::8844")
		result = RealIpFromHeader(header)
		assert.Equal(t, "2001:4860:4860::8844", result, "Should return public IPv6")

		// Test Fastly-Client-Ip
		header = make(http.Header)
		header.Set("Fastly-Client-Ip", "9.9.9.9")
		result = RealIpFromHeader(header)
		assert.Equal(t, "9.9.9.9", result, "Should return Fastly client IP")

		// Test True-Client-Ip
		header = make(http.Header)
		header.Set("True-Client-Ip", "76.76.19.19")
		result = RealIpFromHeader(header)
		assert.Equal(t, "76.76.19.19", result, "Should return true client IP")

		// Test X-Real-IP
		header = make(http.Header)
		header.Set("X-Real-IP", "208.67.222.222")
		result = RealIpFromHeader(header)
		assert.Equal(t, "208.67.222.222", result, "Should return X-Real-IP")

		// Test X-Client-IP
		header = make(http.Header)
		header.Set("X-Client-IP", "208.67.220.220")
		result = RealIpFromHeader(header)
		assert.Equal(t, "208.67.220.220", result, "Should return X-Client-IP")

		// Test X-Original-Forwarded-For with multiple IPs
		header = make(http.Header)
		header.Set("X-Original-Forwarded-For", "192.168.1.1, 8.8.8.8, 1.1.1.1")
		result = RealIpFromHeader(header)
		assert.Equal(t, "8.8.8.8", result, "Should return first public IP from list")

		// Test X-Original-Forwarded-For with only local IPs
		header = make(http.Header)
		header.Set("X-Original-Forwarded-For", "192.168.1.1, 10.0.0.1, 172.16.0.1")
		result = RealIpFromHeader(header)
		assert.Equal(t, "", result, "Should return empty for all local IPs")

		// Test X-Original-Forwarded-For with invalid IP
		header = make(http.Header)
		header.Set("X-Original-Forwarded-For", "invalid-ip, 8.8.8.8")
		result = RealIpFromHeader(header)
		assert.Equal(t, "8.8.8.8", result, "Should skip invalid IP and return valid one")

		// Test X-Forwarded-For
		header = make(http.Header)
		header.Set("X-Forwarded-For", "203.0.113.1, 198.51.100.1")
		result = RealIpFromHeader(header)
		assert.Equal(t, "203.0.113.1", result, "Should return first public IP from X-Forwarded-For")

		// Test X-Forwarded-For with whitespace
		header = make(http.Header)
		header.Set("X-Forwarded-For", " 203.0.113.2 , 198.51.100.2 ")
		result = RealIpFromHeader(header)
		assert.Equal(t, "203.0.113.2", result, "Should handle whitespace properly")

		// Test X-Forwarded
		header = make(http.Header)
		header.Set("X-Forwarded", "203.0.113.3, 198.51.100.3")
		result = RealIpFromHeader(header)
		assert.Equal(t, "203.0.113.3", result, "Should return first public IP from X-Forwarded")

		// Test Forwarded-For
		header = make(http.Header)
		header.Set("Forwarded-For", "203.0.113.4, 198.51.100.4")
		result = RealIpFromHeader(header)
		assert.Equal(t, "203.0.113.4", result, "Should return first public IP from Forwarded-For")

		// Test Forwarded
		header = make(http.Header)
		header.Set("Forwarded", "203.0.113.5, 198.51.100.5")
		result = RealIpFromHeader(header)
		assert.Equal(t, "203.0.113.5", result, "Should return first public IP from Forwarded")

		// Test priority order - Cf-Connecting-Ip should have highest priority
		header = make(http.Header)
		header.Set("Cf-Connecting-Ip", "8.8.8.8")
		header.Set("X-Forwarded-For", "1.1.1.1")
		result = RealIpFromHeader(header)
		assert.Equal(t, "8.8.8.8", result, "Cf-Connecting-Ip should have highest priority")

		// Test with all local IPs in various headers
		header = make(http.Header)
		header.Set("Cf-Connecting-Ip", "192.168.1.1")
		header.Set("X-Real-IP", "10.0.0.1")
		header.Set("X-Forwarded-For", "172.16.0.1")
		result = RealIpFromHeader(header)
		assert.Equal(t, "", result, "Should return empty when all IPs are local")
	})

	// Test IsLocalIp with comprehensive IP types
	t.Run("IsLocalIp_CompleteCoverage", func(t *testing.T) {
		// Test invalid IP (should return false according to the actual implementation)
		assert.False(t, IsLocalIp("invalid-ip"), "Invalid IP should return false")

		// Test public IPs
		assert.False(t, IsLocalIp("8.8.8.8"), "8.8.8.8 should not be local")
		assert.False(t, IsLocalIp("1.1.1.1"), "1.1.1.1 should not be local")
		assert.False(t, IsLocalIp("208.67.222.222"), "208.67.222.222 should not be local")

		// Test private IPs
		assert.True(t, IsLocalIp("192.168.1.1"), "192.168.1.1 should be local")
		assert.True(t, IsLocalIp("10.0.0.1"), "10.0.0.1 should be local")
		assert.True(t, IsLocalIp("172.16.0.1"), "172.16.0.1 should be local")

		// Test loopback
		assert.True(t, IsLocalIp("127.0.0.1"), "127.0.0.1 should be local")
		assert.True(t, IsLocalIp("::1"), "::1 should be local")

		// Test link-local
		assert.True(t, IsLocalIp("169.254.1.1"), "169.254.1.1 should be local")
		assert.True(t, IsLocalIp("fe80::1"), "fe80::1 should be local")

		// Test IPv6 private
		assert.True(t, IsLocalIp("fc00::1"), "fc00::1 should be local")
		assert.True(t, IsLocalIp("fd00::1"), "fd00::1 should be local")

		// Test IPv6 public
		assert.False(t, IsLocalIp("2001:4860:4860::8888"), "2001:4860:4860::8888 should not be local")
	})

	// Test GetListenIp comprehensive scenarios
	t.Run("GetListenIp_CompleteCoverage", func(t *testing.T) {
		// Test default (no arguments)
		result := GetListenIp()
		if result != "" {
			assert.True(t, net.ParseIP(result) != nil, "Default GetListenIp should return valid IP or empty")
		}

		// Test explicit IPv4 preference
		result = GetListenIp(false)
		if result != "" {
			assert.True(t, net.ParseIP(result) != nil, "IPv4 GetListenIp should return valid IP or empty")
		}

		// Test explicit IPv6 preference
		result = GetListenIp(true)
		if result != "" {
			assert.True(t, net.ParseIP(result) != nil, "IPv6 GetListenIp should return valid IP or empty")
		}

		// Test with multiple arguments (only first should be used)
		result = GetListenIp(true, false, true)
		if result != "" {
			assert.True(t, net.ParseIP(result) != nil, "Multiple args GetListenIp should return valid IP or empty")
		}
	})

	// Test GetInterfaceIpByAddrs with mock addresses that exercise all paths
	t.Run("GetInterfaceIpByAddrs_AllPaths", func(t *testing.T) {
		// Create helper function
		createIPNet := func(ipStr string) *net.IPNet {
			ip, ipnet, _ := net.ParseCIDR(ipStr)
			ipnet.IP = ip
			return ipnet
		}

		// Test empty address list
		result := GetInterfaceIpByAddrs([]net.Addr{}, false)
		assert.Equal(t, "", result, "Empty address list should return empty")

		result = GetInterfaceIpByAddrs([]net.Addr{}, true)
		assert.Equal(t, "", result, "Empty address list should return empty")

		// Test only loopback addresses (should be filtered out)
		loopbackAddrs := []net.Addr{
			createIPNet("127.0.0.1/8"),
			createIPNet("::1/128"),
		}
		result = GetInterfaceIpByAddrs(loopbackAddrs, false)
		assert.Equal(t, "", result, "Only loopback addresses should return empty")

		result = GetInterfaceIpByAddrs(loopbackAddrs, true)
		assert.Equal(t, "", result, "Only loopback addresses should return empty")

		// Test IPv4 only with IPv6 preference (should fallback to IPv4)
		ipv4OnlyAddrs := []net.Addr{createIPNet("192.168.1.100/24")}
		result = GetInterfaceIpByAddrs(ipv4OnlyAddrs, true)
		assert.Equal(t, "192.168.1.100", result, "IPv4 only with IPv6 preference should fallback")

		// Test IPv6 only with IPv4 preference (should return empty)
		ipv6OnlyAddrs := []net.Addr{createIPNet("2001:db8::1/64")}
		result = GetInterfaceIpByAddrs(ipv6OnlyAddrs, false)
		assert.Equal(t, "", result, "IPv6 only with IPv4 preference should return empty")

		// Test mixed addresses with IPv4 preference
		mixedAddrs := []net.Addr{
			createIPNet("192.168.1.100/24"),
			createIPNet("2001:db8::1/64"),
		}
		result = GetInterfaceIpByAddrs(mixedAddrs, false)
		assert.Equal(t, "192.168.1.100", result, "Mixed addresses with IPv4 preference should return IPv4")

		// Test mixed addresses with IPv6 preference
		result = GetInterfaceIpByAddrs(mixedAddrs, true)
		assert.Equal(t, "2001:db8::1", result, "Mixed addresses with IPv6 preference should return IPv6")

		// Test non-IPNet address type using existing mockAddr from interface_test.go
		nonIPNetAddrs := []net.Addr{mockAddr{}}
		result = GetInterfaceIpByAddrs(nonIPNetAddrs, false)
		assert.Equal(t, "", result, "Non-IPNet addresses should be skipped")
	})

	// Test specific error scenarios for GetInterfaceIpByName
	t.Run("GetInterfaceIpByName_ErrorScenarios", func(t *testing.T) {
		// Test with empty interface name
		result := GetInterfaceIpByName("", false)
		assert.Equal(t, "", result, "Empty interface name should return empty")

		result = GetInterfaceIpByName("", true)
		assert.Equal(t, "", result, "Empty interface name should return empty")

		// Test with various non-existent interface names
		nonExistentNames := []string{
			"fake_interface_xyz",
			"nonexistent12345",
			"test_interface_that_does_not_exist",
		}

		for _, name := range nonExistentNames {
			result = GetInterfaceIpByName(name, false)
			assert.Equal(t, "", result, "Non-existent interface %s should return empty", name)

			result = GetInterfaceIpByName(name, true)
			assert.Equal(t, "", result, "Non-existent interface %s should return empty", name)
		}
	})
}

// TestNetworkRealIpFromHeader_EdgeCases tests edge cases in header processing
func TestNetworkRealIpFromHeader_EdgeCases(t *testing.T) {
	// Test with headers containing empty values
	t.Run("EmptyHeaderValues", func(t *testing.T) {
		header := make(http.Header)
		header.Set("Cf-Connecting-Ip", "")
		header.Set("X-Real-IP", "8.8.8.8")
		result := RealIpFromHeader(header)
		assert.Equal(t, "8.8.8.8", result, "Should skip empty header and use next valid one")
	})

	// Test with headers containing only whitespace
	t.Run("WhitespaceInHeaders", func(t *testing.T) {
		header := make(http.Header)
		header.Set("X-Forwarded-For", "   , , 8.8.8.8,   ")
		result := RealIpFromHeader(header)
		assert.Equal(t, "8.8.8.8", result, "Should handle whitespace and empty entries properly")
	})

	// Test with malformed forwarded headers
	t.Run("MalformedForwardedHeaders", func(t *testing.T) {
		header := make(http.Header)
		header.Set("X-Forwarded-For", "not-an-ip, , 256.256.256.256, 8.8.8.8")
		result := RealIpFromHeader(header)
		assert.Equal(t, "8.8.8.8", result, "Should skip malformed IPs and return valid one")
	})

	// Test with all headers containing local IPs
	t.Run("AllLocalIPs", func(t *testing.T) {
		header := make(http.Header)
		header.Set("Cf-Connecting-Ip", "192.168.1.1")
		header.Set("Cf-Pseudo-Ipv4", "10.0.0.1")
		header.Set("X-Real-IP", "172.16.0.1")
		header.Set("X-Forwarded-For", "127.0.0.1, 192.168.1.2")
		result := RealIpFromHeader(header)
		assert.Equal(t, "", result, "Should return empty when all IPs are local")
	})
}

// TestGetListenIp_SystemInterfaceScenarios tests with actual system interfaces
func TestGetListenIp_SystemInterfaceScenarios(t *testing.T) {
	t.Run("SystemInterfaceTesting", func(t *testing.T) {
		// Get actual system interfaces
		interfaces, err := net.Interfaces()
		if err != nil {
			t.Skip("Cannot get system interfaces")
		}

		// Test each preference type multiple times to exercise all paths
		for i := 0; i < 5; i++ {
			// Test default
			result := GetListenIp()
			if result != "" && net.ParseIP(result) == nil {
				t.Errorf("Invalid IP from GetListenIp(): %s", result)
			}

			// Test IPv4 preference
			result = GetListenIp(false)
			if result != "" && net.ParseIP(result) == nil {
				t.Errorf("Invalid IP from GetListenIp(false): %s", result)
			}

			// Test IPv6 preference
			result = GetListenIp(true)
			if result != "" && net.ParseIP(result) == nil {
				t.Errorf("Invalid IP from GetListenIp(true): %s", result)
			}
		}

		// Also test individual interfaces to potentially hit error paths in GetInterfaceIpByName
		for _, iface := range interfaces {
			if len(iface.Name) > 0 {
				// Test both IP versions for each interface
				ipv4 := GetInterfaceIpByName(iface.Name, false)
				ipv6 := GetInterfaceIpByName(iface.Name, true)

				if ipv4 != "" && net.ParseIP(ipv4) == nil {
					t.Errorf("Invalid IPv4 for interface %s: %s", iface.Name, ipv4)
				}
				if ipv6 != "" && net.ParseIP(ipv6) == nil {
					t.Errorf("Invalid IPv6 for interface %s: %s", iface.Name, ipv6)
				}
			}
		}
	})
}