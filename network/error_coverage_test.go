package network

import (
	"net"
	"testing"
)

// TestInterfaceErrorConditions focuses on triggering error conditions
func TestInterfaceErrorConditions(t *testing.T) {
	t.Run("GetInterfaceIpByName error paths", func(t *testing.T) {
		// Test with interfaces that definitely don't exist to trigger error path
		errorInterfaces := []string{
			"definitely_nonexistent_interface_999",
			"fake_interface_that_does_not_exist",
			"error_trigger_interface",
		}

		for _, ifName := range errorInterfaces {
			// Test both IPv4 and IPv6 to ensure all error paths are hit
			result4 := GetInterfaceIpByName(ifName, false)
			result6 := GetInterfaceIpByName(ifName, true)

			// Should return empty for non-existent interfaces
			if result4 != "" {
				t.Errorf("Expected empty result for %s IPv4, got %s", ifName, result4)
			}
			if result6 != "" {
				t.Errorf("Expected empty result for %s IPv6, got %s", ifName, result6)
			}
		}
	})

	t.Run("GetListenIp extensive testing", func(t *testing.T) {
		// Test all parameter combinations to hit all branches
		testCases := []struct {
			name string
			args []bool
		}{
			{"No args (default false)", []bool{}},
			{"Explicit false", []bool{false}},
			{"Explicit true", []bool{true}},
			{"Multiple args", []bool{true, false, true}},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := GetListenIp(tc.args...)
				// Result should be empty or valid IP
				if result != "" && net.ParseIP(result) == nil {
					t.Errorf("GetListenIp returned invalid IP: %s", result)
				}
			})
		}
	})

	t.Run("Force specific network scenarios", func(t *testing.T) {
		// Try to exercise specific branches in GetListenIp
		// by calling it many times with different preferences
		for i := 0; i < 20; i++ {
			// Alternate between preferences to try to hit different paths
			if i%2 == 0 {
				GetListenIp(true) // IPv6 preference
			} else {
				GetListenIp(false) // IPv4 preference
			}
		}

		// Also test the default path
		for i := 0; i < 10; i++ {
			GetListenIp() // Default (no args)
		}
	})

	t.Run("Invalid interface names", func(t *testing.T) {
		// Test various invalid interface names to trigger error handling
		invalidNames := []string{
			"",       // Empty name
			"invalid-interface-name-that-is-very-long-and-definitely-does-not-exist",
			"🚀", // Unicode characters
			"interface with spaces",
			"interface\nwith\nnewlines",
			"interface\twith\ttabs",
		}

		for _, name := range invalidNames {
			ipv4 := GetInterfaceIpByName(name, false)
			ipv6 := GetInterfaceIpByName(name, true)

			// Should handle gracefully and return empty
			if ipv4 != "" {
				t.Errorf("Expected empty for invalid interface %q IPv4, got %s", name, ipv4)
			}
			if ipv6 != "" {
				t.Errorf("Expected empty for invalid interface %q IPv6, got %s", name, ipv6)
			}
		}
	})
}

// TestSpecificCoverageImprovements targets uncovered lines specifically
func TestSpecificCoverageImprovements(t *testing.T) {
	t.Run("Target GetInterfaceIpByName error handling", func(t *testing.T) {
		// The function has two error paths:
		// 1. net.InterfaceByName error (line 10-13)
		// 2. inter.Addrs() error (line 16-19)

		// Test case 1: Invalid interface name (triggers net.InterfaceByName error)
		result := GetInterfaceIpByName("absolutely_nonexistent_interface_12345", false)
		if result != "" {
			t.Errorf("Expected empty result for non-existent interface, got %s", result)
		}

		result = GetInterfaceIpByName("absolutely_nonexistent_interface_12345", true)
		if result != "" {
			t.Errorf("Expected empty result for non-existent interface, got %s", result)
		}
	})

	t.Run("Target GetListenIp error scenarios", func(t *testing.T) {
		// GetListenIp has multiple paths:
		// 1. eth0 check (line 53-55)
		// 2. en0 check (line 58-60)
		// 3. net.InterfaceAddrs() fallback (line 62-66)
		// 4. Final GetInterfaceIpByAddrs call (line 68-70)
		// 5. Error logging (line 72)

		// Test multiple times to try to hit the error path where no IP is found
		for i := 0; i < 50; i++ {
			// Vary the preference to exercise different paths
			pref := i%3 == 0 // Sometimes true, sometimes false

			result := GetListenIp(pref)
			if result != "" && net.ParseIP(result) == nil {
				t.Errorf("GetListenIp returned invalid IP: %s", result)
			}
		}
	})

	t.Run("Exercise GetInterfaceIpByAddrs with edge cases", func(t *testing.T) {
		// Test with address lists that might cause specific behaviors
		createIPNet := func(ipStr string) *net.IPNet {
			ip, ipnet, err := net.ParseCIDR(ipStr)
			if err != nil {
				t.Fatalf("Failed to parse CIDR %s: %v", ipStr, err)
			}
			ipnet.IP = ip
			return ipnet
		}

		// Test various scenarios
		testCases := []struct {
			name      string
			addresses []net.Addr
			prev6     bool
		}{
			{
				"Multiple IPv4 addresses",
				[]net.Addr{
					createIPNet("192.168.1.100/24"),
					createIPNet("10.0.0.100/8"),
				},
				false,
			},
			{
				"Multiple IPv6 addresses",
				[]net.Addr{
					createIPNet("2001:db8::1/64"),
					createIPNet("fe80::1/64"),
				},
				true,
			},
			{
				"Mixed with loopback first",
				[]net.Addr{
					createIPNet("127.0.0.1/8"),
					createIPNet("192.168.1.100/24"),
				},
				false,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := GetInterfaceIpByAddrs(tc.addresses, tc.prev6)
				if result != "" && net.ParseIP(result) == nil {
					t.Errorf("GetInterfaceIpByAddrs returned invalid IP: %s", result)
				}
			})
		}
	})
}