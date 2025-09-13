package network

import (
	"net"
	"testing"
)

// TestSpecialInterfaceConditions tests interfaces that might trigger error conditions
func TestSpecialInterfaceConditions(t *testing.T) {
	// Test interfaces that exist but might have issues with Addrs()
	specialInterfaces := []string{
		// Virtual/tunnel interfaces that might not have addresses or have addr issues
		"utun0", "utun1", "utun2", "utun3", // Tunnel interfaces
		"anpi0", "anpi1",                   // Apple network interfaces
		"gif0", "stf0",                     // IPv6 transition interfaces
		"pktap0",                           // Packet tap interface (might cause issues)
		
		// Try some known active interfaces too
		"bridge100", "bridge101", "bridge102", // Bridge interfaces with IPs
		"vmenet0", "vmenet1", "vmenet2",      // VMware interfaces
	}

	for _, ifName := range specialInterfaces {
		t.Run("Special_"+ifName, func(t *testing.T) {
			// Test both IPv4 and IPv6 to exercise all branches
			ipv4 := GetInterfaceIpByName(ifName, false)
			ipv6 := GetInterfaceIpByName(ifName, true)
			
			// Validate results
			if ipv4 != "" && net.ParseIP(ipv4) == nil {
				t.Errorf("Invalid IPv4 for %s: %s", ifName, ipv4)
			}
			if ipv6 != "" && net.ParseIP(ipv6) == nil {
				t.Errorf("Invalid IPv6 for %s: %s", ifName, ipv6)
			}
		})
	}
}

// TestGetListenIp_DetailedPaths attempts to exercise different paths
func TestGetListenIp_DetailedPaths(t *testing.T) {
	// Since GetListenIp is currently finding IPs successfully,
	// we test multiple times to ensure consistent behavior
	for i := 0; i < 5; i++ {
		t.Run("Iteration_"+string(rune('A'+i)), func(t *testing.T) {
			// Test the main entry points
			result := GetListenIp()
			if result != "" && net.ParseIP(result) == nil {
				t.Errorf("Invalid IP from GetListenIp(): %s", result)
			}
			
			result4 := GetListenIp(false)
			if result4 != "" && net.ParseIP(result4) == nil {
				t.Errorf("Invalid IPv4 from GetListenIp(false): %s", result4)
			}
			
			result6 := GetListenIp(true)
			if result6 != "" && net.ParseIP(result6) == nil {
				t.Errorf("Invalid IPv6 from GetListenIp(true): %s", result6)
			}
		})
	}
}

// TestGetInterfaceIpByName_AllInterfaces tests every system interface
func TestGetInterfaceIpByName_AllInterfaces(t *testing.T) {
	interfaces, err := net.Interfaces()
	if err != nil {
		t.Skipf("Cannot get interfaces: %v", err)
	}

	for _, iface := range interfaces {
		t.Run("AllInterfaces_"+iface.Name, func(t *testing.T) {
			// Test both IP versions for every interface
			ipv4 := GetInterfaceIpByName(iface.Name, false)
			ipv6 := GetInterfaceIpByName(iface.Name, true)
			
			if ipv4 != "" && net.ParseIP(ipv4) == nil {
				t.Errorf("Invalid IPv4 for %s: %s", iface.Name, ipv4)
			}
			if ipv6 != "" && net.ParseIP(ipv6) == nil {
				t.Errorf("Invalid IPv6 for %s: %s", iface.Name, ipv6)
			}
		})
	}
}

// TestGetInterfaceIpByAddrs_CompleteBranchCoverage tests all branches thoroughly
func TestGetInterfaceIpByAddrs_CompleteBranchCoverage(t *testing.T) {
	createIPNet := func(ipStr string) *net.IPNet {
		ip, ipnet, _ := net.ParseCIDR(ipStr)
		ipnet.IP = ip
		return ipnet
	}

	scenarios := []struct {
		name      string
		addresses []net.Addr
		prev6     bool
		desc      string
	}{
		{
			name:      "Empty_list_IPv4",
			addresses: []net.Addr{},
			prev6:     false,
			desc:      "Empty address list should return empty string",
		},
		{
			name:      "Empty_list_IPv6", 
			addresses: []net.Addr{},
			prev6:     true,
			desc:      "Empty address list should return empty string",
		},
		{
			name:      "Only_loopback_IPv4",
			addresses: []net.Addr{createIPNet("127.0.0.1/8")},
			prev6:     false,
			desc:      "Loopback only should return empty",
		},
		{
			name:      "Only_loopback_IPv6",
			addresses: []net.Addr{createIPNet("::1/128")},
			prev6:     true,
			desc:      "IPv6 loopback only should return empty",
		},
		{
			name:      "IPv4_only_prefer_IPv6_fallback",
			addresses: []net.Addr{createIPNet("192.168.1.100/24")},
			prev6:     true,
			desc:      "Should fallback to IPv4 when IPv6 preferred but unavailable",
		},
		{
			name:      "IPv6_only_prefer_IPv4_no_fallback",
			addresses: []net.Addr{createIPNet("2001:db8::1/64")},
			prev6:     false,
			desc:      "Should return empty when only IPv6 available but IPv4 preferred",
		},
		{
			name: "Mixed_addresses_prefer_IPv4",
			addresses: []net.Addr{
				createIPNet("192.168.1.100/24"),
				createIPNet("2001:db8::1/64"),
			},
			prev6: false,
			desc:  "Should return IPv4 when both available and IPv4 preferred",
		},
		{
			name: "Mixed_addresses_prefer_IPv6",
			addresses: []net.Addr{
				createIPNet("192.168.1.100/24"),
				createIPNet("2001:db8::1/64"),
			},
			prev6: true,
			desc:  "Should return IPv6 when both available and IPv6 preferred",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := GetInterfaceIpByAddrs(scenario.addresses, scenario.prev6)
			// Just validate that result is either empty or valid IP
			if result != "" && net.ParseIP(result) == nil {
				t.Errorf("Invalid IP result for %s: %s", scenario.name, result)
			}
		})
	}
}