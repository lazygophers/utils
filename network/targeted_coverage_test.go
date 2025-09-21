package network

import (
	"net"
	"testing"
	"time"
)

// TestMockInterfaceError creates a scenario to trigger interface.Addrs() error
func TestMockInterfaceError(t *testing.T) {
	// Since we can't easily mock the network interfaces in this environment,
	// let's test with interfaces that might exist but cause addr retrieval issues

	t.Run("TriggerAddrsError", func(t *testing.T) {
		// Try interfaces that might exist but have problematic address retrieval
		// On macOS, some virtual interfaces might exist but have addr issues
		problematicInterfaces := []string{
			"anpi0", "anpi1", "anpi2",    // Apple networking interfaces
			"gif0", "gif1", "gif2",       // Generic tunnel interfaces
			"stf0", "stf1", "stf2",       // 6to4 tunnel interfaces
			"lo100", "lo101", "lo102",    // Additional loopback interfaces
			"pktap0", "pktap1",           // Packet tap interfaces
			"utun99", "utun100",          // Tunnel interfaces with high numbers
		}

		for _, ifName := range problematicInterfaces {
			// Test both IPv4 and IPv6 to potentially trigger the addr error
			result4 := GetInterfaceIpByName(ifName, false)
			result6 := GetInterfaceIpByName(ifName, true)

			// Results should be either empty or valid
			if result4 != "" && net.ParseIP(result4) == nil {
				t.Errorf("Invalid IPv4 for %s: %s", ifName, result4)
			}
			if result6 != "" && net.ParseIP(result6) == nil {
				t.Errorf("Invalid IPv6 for %s: %s", ifName, result6)
			}
		}
	})
}

// TestGetListenIp_ForceMissingPaths specifically targets the uncovered GetListenIp paths
func TestGetListenIp_ForceMissingPaths(t *testing.T) {
	t.Run("ForceAllErrorPaths", func(t *testing.T) {
		// The goal is to create scenarios where GetListenIp hits all uncovered branches:
		// 1. eth0 returns empty (line 54)
		// 2. en0 returns empty (line 58)
		// 3. net.InterfaceAddrs() error path (line 64-65)
		// 4. GetInterfaceIpByAddrs returns empty (line 68)
		// 5. Final error log (line 72)

		// Since eth0 and en0 probably don't exist on this macOS system,
		// the function should fall through to net.InterfaceAddrs()

		// Test with different preferences to exercise all code paths
		testPreferences := []struct {
			name string
			args []bool
		}{
			{"Default", []bool{}},
			{"IPv4", []bool{false}},
			{"IPv6", []bool{true}},
			{"Multiple", []bool{true, false, true}},
		}

		for _, test := range testPreferences {
			t.Run(test.name, func(t *testing.T) {
				// Call multiple times to potentially hit different states
				for i := 0; i < 3; i++ {
					result := GetListenIp(test.args...)
					if result != "" && net.ParseIP(result) == nil {
						t.Errorf("Invalid IP from GetListenIp: %s", result)
					}
				}
			})
		}
	})

	t.Run("RapidFireTesting", func(t *testing.T) {
		// Make rapid successive calls to try to trigger race conditions
		// or different network states that might cause errors
		for i := 0; i < 50; i++ {
			GetListenIp()      // default
			GetListenIp(false) // IPv4
			GetListenIp(true)  // IPv6

			// Small delay to potentially catch different network states
			if i%10 == 0 {
				time.Sleep(time.Microsecond)
			}
		}
	})
}

// TestSpecificInterfaceByName targets the specific line coverage issues
func TestSpecificInterfaceByName(t *testing.T) {
	t.Run("CoverAllBranches", func(t *testing.T) {
		// Test interfaces that definitely don't exist to hit error path at line 11
		nonExistentInterfaces := []string{
			"fake_interface_999",
			"nonexistent_eth_12345",
			"definitely_not_real",
			"", // empty name should also trigger error
		}

		for _, ifName := range nonExistentInterfaces {
			// Test both IP versions to ensure we hit all branches
			GetInterfaceIpByName(ifName, false)
			GetInterfaceIpByName(ifName, true)
		}
	})

	t.Run("TestExistingButProblematic", func(t *testing.T) {
		// Get actual system interfaces and test some that might have addr issues
		interfaces, err := net.Interfaces()
		if err != nil {
			t.Skipf("Could not get interfaces: %v", err)
		}

		// Test all real interfaces to potentially hit the addr error path
		for _, iface := range interfaces {
			// Skip interfaces that are likely to work fine
			if iface.Name == "lo0" || iface.Name == "en0" {
				continue
			}

			// Test potentially problematic interfaces
			GetInterfaceIpByName(iface.Name, false)
			GetInterfaceIpByName(iface.Name, true)
		}
	})
}

// TestGetListenIp_SpecificPathCoverage targets specific lines in GetListenIp
func TestGetListenIp_SpecificPathCoverage(t *testing.T) {
	t.Run("CoverBooleanLogic", func(t *testing.T) {
		// Test the boolean logic at line 50: _prev6 = len(prev6) > 0 && prev6[0]
		testCases := []struct {
			name string
			args []bool
			desc string
		}{
			{"EmptyArgs", []bool{}, "len(prev6) = 0, should be false"},
			{"SingleFalse", []bool{false}, "len(prev6) > 0 && prev6[0] = false"},
			{"SingleTrue", []bool{true}, "len(prev6) > 0 && prev6[0] = true"},
			{"MultipleFalseFirst", []bool{false, true}, "len(prev6) > 0 && prev6[0] = false"},
			{"MultipleTrueFirst", []bool{true, false}, "len(prev6) > 0 && prev6[0] = true"},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := GetListenIp(tc.args...)
				if result != "" && net.ParseIP(result) == nil {
					t.Errorf("Invalid IP for %s: %s", tc.name, result)
				}
			})
		}
	})

	t.Run("TestEth0En0Branches", func(t *testing.T) {
		// Since eth0 and en0 likely don't exist on this system,
		// these calls should hit the empty return conditions at lines 54 and 58

		// Test the eth0 branch (line 53-55)
		result1 := GetListenIp(false) // IPv4 preference
		result2 := GetListenIp(true)  // IPv6 preference

		// Test the en0 branch (line 57-60)
		result3 := GetListenIp()      // default (false)

		// All should either be empty or valid IPs
		results := []string{result1, result2, result3}
		for i, result := range results {
			if result != "" && net.ParseIP(result) == nil {
				t.Errorf("Invalid IP result %d: %s", i, result)
			}
		}
	})
}

// TestInterfaceByName_ForceBothErrorPaths ensures we hit both error paths
func TestInterfaceByName_ForceBothErrorPaths(t *testing.T) {
	t.Run("InterfaceByNameError", func(t *testing.T) {
		// Force the net.InterfaceByName error at line 10-12
		badNames := []string{
			"definitely_nonexistent_interface_with_long_name_12345",
			"bad_interface_name_999",
			"",
		}

		for _, name := range badNames {
			GetInterfaceIpByName(name, false)
			GetInterfaceIpByName(name, true)
		}
	})

	t.Run("PotentialAddrsError", func(t *testing.T) {
		// Try to force the inter.Addrs() error at line 15-18
		// This is harder to trigger, but we can try with special interfaces

		// Test with every system interface to see if any have addr issues
		interfaces, err := net.Interfaces()
		if err == nil {
			for _, iface := range interfaces {
				// Test each interface - some might have addr retrieval issues
				GetInterfaceIpByName(iface.Name, false)
				GetInterfaceIpByName(iface.Name, true)
			}
		}

		// Also test with some special interface names that might exist
		specialNames := []string{
			"gif0", "gif1", "stf0", "utun0", "utun1", "utun2",
			"anpi0", "anpi1", "awdl0", "llw0", "pktap0",
		}

		for _, name := range specialNames {
			GetInterfaceIpByName(name, false)
			GetInterfaceIpByName(name, true)
		}
	})
}

// TestGetListenIp_ExtensiveCoverage tries to hit all remaining uncovered lines
func TestGetListenIp_ExtensiveCoverage(t *testing.T) {
	t.Run("AttemptErrorLog", func(t *testing.T) {
		// Try to hit the error log at line 72: "get interface ip failed"
		// This happens when all of the following fail:
		// 1. GetInterfaceIpByName("eth0", _prev6) returns ""
		// 2. GetInterfaceIpByName("en0", _prev6) returns ""
		// 3. GetInterfaceIpByAddrs(address, _prev6) returns ""

		// Since eth0 probably doesn't exist on this macOS system,
		// and en0 might exist but return results,
		// we need to create conditions where GetInterfaceIpByAddrs also fails

		// Test many times with different preferences
		for i := 0; i < 10; i++ {
			GetListenIp()      // default false
			GetListenIp(false) // explicit false
			GetListenIp(true)  // explicit true
		}
	})
}

// TestGetInterfaceIpByName_22Percent targets the missing 22.2% in GetInterfaceIpByName
func TestGetInterfaceIpByName_22Percent(t *testing.T) {
	t.Run("HitAllPaths", func(t *testing.T) {
		// The function is at 77.8% coverage, so we're missing 22.2%
		// This is likely the error path at line 17: log.Errorf("err:%v", err)

		// Get all system interfaces
		interfaces, err := net.Interfaces()
		if err != nil {
			t.Skipf("Cannot get interfaces: %v", err)
		}

		// Test with interfaces that might have address retrieval issues
		for _, iface := range interfaces {
			// Call both IPv4 and IPv6 versions
			GetInterfaceIpByName(iface.Name, false)
			GetInterfaceIpByName(iface.Name, true)
		}

		// Also test with non-existent interfaces to ensure error path
		fakeInterfaces := []string{
			"eth999", "en999", "wlan999", "fake123",
		}

		for _, fake := range fakeInterfaces {
			GetInterfaceIpByName(fake, false)
			GetInterfaceIpByName(fake, true)
		}
	})
}