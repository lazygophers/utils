package network

import (
	"net"
	"testing"
)

// TestGetInterfaceIpByName_EdgeCasesCoverage focuses on specific missing coverage paths
func TestGetInterfaceIpByName_EdgeCasesCoverage(t *testing.T) {
	// Test to try to hit the inter.Addrs() error path (line 17-18)
	// We test with interfaces that might exist but have address retrieval issues
	problematicInterfaces := []string{
		"anpi0", "anpi1", "gif0", "stf0", // macOS virtual interfaces that might have address issues
		"ap1", "awdl0", "llw0", // Apple wireless interfaces
		"pktap0",         // Packet capture interface
		"utun0", "utun1", // Tunnel interfaces
	}

	for _, ifName := range problematicInterfaces {
		t.Run("EdgeCase_"+ifName, func(t *testing.T) {
			// Test both IP versions to exercise all paths
			ipv4 := GetInterfaceIpByName(ifName, false)
			ipv6 := GetInterfaceIpByName(ifName, true)

			// Either empty or valid IP
			if ipv4 != "" && net.ParseIP(ipv4) == nil {
				t.Errorf("Invalid IPv4 for %s: %s", ifName, ipv4)
			}
			if ipv6 != "" && net.ParseIP(ipv6) == nil {
				t.Errorf("Invalid IPv6 for %s: %s", ifName, ipv6)
			}
		})
	}
}

// TestGetListenIp_PathCoverage focuses on covering missing GetListenIp paths
func TestGetListenIp_PathCoverage(t *testing.T) {
	t.Run("EdgePathTesting", func(t *testing.T) {
		// Multiple calls to try to hit different network states
		// This tries to exercise the specific uncovered paths in GetListenIp

		tests := []struct {
			name string
			args []bool
		}{
			{"Default", []bool{}},
			{"IPv4", []bool{false}},
			{"IPv6", []bool{true}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := GetListenIp(test.args...)
				if result != "" && net.ParseIP(result) == nil {
					t.Errorf("Invalid IP from GetListenIp: %s", result)
				}
			})
		}
	})
}

// TestSystemInterfaceReality tests with interfaces that definitely exist on this system
func TestSystemInterfaceReality(t *testing.T) {
	// These interfaces definitely exist based on our earlier ifconfig output
	knownInterfaces := []string{"lo0", "en0", "bridge100", "bridge101", "en5", "en6"}

	for _, iface := range knownInterfaces {
		t.Run("Known_"+iface, func(t *testing.T) {
			// Test with both preferences
			ipv4 := GetInterfaceIpByName(iface, false)
			ipv6 := GetInterfaceIpByName(iface, true)

			if ipv4 != "" && net.ParseIP(ipv4) == nil {
				t.Errorf("Invalid IPv4 for known interface %s: %s", iface, ipv4)
			}
			if ipv6 != "" && net.ParseIP(ipv6) == nil {
				t.Errorf("Invalid IPv6 for known interface %s: %s", iface, ipv6)
			}
		})
	}
}

// TestGetInterfaceIpByAddrs_AllBranches ensures we hit all branches in GetInterfaceIpByAddrs
func TestGetInterfaceIpByAddrs_AllBranches(t *testing.T) {
	createIPNet := func(ipStr string) *net.IPNet {
		ip, ipnet, _ := net.ParseCIDR(ipStr)
		ipnet.IP = ip
		return ipnet
	}

	tests := []struct {
		name      string
		addresses []net.Addr
		prev6     bool
		expected  string
	}{
		{
			name:      "Empty addresses IPv4",
			addresses: []net.Addr{},
			prev6:     false,
			expected:  "",
		},
		{
			name:      "Empty addresses IPv6",
			addresses: []net.Addr{},
			prev6:     true,
			expected:  "",
		},
		{
			name:      "Only IPv4 prefer IPv6 fallback",
			addresses: []net.Addr{createIPNet("192.168.1.100/24")},
			prev6:     true,
			expected:  "192.168.1.100",
		},
		{
			name:      "Only IPv6 prefer IPv4 no fallback",
			addresses: []net.Addr{createIPNet("2001:db8::1/64")},
			prev6:     false,
			expected:  "",
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
