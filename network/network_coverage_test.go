package network

import (
	"net"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetListenIpCoverage tests various branches of GetListenIp function
func TestGetListenIpCoverage(t *testing.T) {
	t.Run("get_listen_ip_with_ipv6_preference", func(t *testing.T) {
		// Test with IPv6 preference
		result := GetListenIp(true)
		// Result could be empty, IPv4, or IPv6 address, all are acceptable
		// The function tries IPv6 first but falls back to IPv4 if no IPv6 is available
		if result != "" {
			// Both IPv4 and IPv6 formats are acceptable
			assert.True(t, strings.Contains(result, ".") || strings.Contains(result, ":"))
		}
	})
	
	t.Run("get_listen_ip_without_preference", func(t *testing.T) {
		// Test without preference (should default to IPv4)
		result := GetListenIp()
		// Result could be empty or an IPv4 address, both are acceptable
		if result != "" {
			assert.Contains(t, result, ".")
		}
	})
	
	t.Run("get_listen_ip_multiple_params", func(t *testing.T) {
		// Test with multiple parameters (only first one should be considered)
		result := GetListenIp(false, true, false)
		// Result could be empty or an IPv4 address, both are acceptable
		if result != "" {
			assert.Contains(t, result, ".")
		}
	})
	
	t.Run("get_interface_ip_by_name_invalid", func(t *testing.T) {
		// Test GetInterfaceIpByName with invalid interface name
		result := GetInterfaceIpByName("invalid_interface_name_12345", false)
		assert.Empty(t, result)
	})
	
	t.Run("get_interface_ip_by_name_ipv6_preference", func(t *testing.T) {
		// Test GetInterfaceIpByName with IPv6 preference
		result := GetInterfaceIpByName("lo0", true)
		// On macOS, lo0 should have an IPv6 address
		if result != "" {
			assert.Contains(t, result, ":")
		}
	})
}

// TestGetInterfaceIpByAddrsEdgeCases tests edge cases for GetInterfaceIpByAddrs
func TestGetInterfaceIpByAddrsEdgeCases(t *testing.T) {
	t.Run("get_interface_ip_by_addrs_only_ipv6_prev6_false", func(t *testing.T) {
		// Test with only IPv6 addresses but prev6=false
		addresses := []net.Addr{
			&net.IPNet{
				IP:   net.ParseIP("fe80::1"),
				Mask: net.CIDRMask(64, 128),
			},
			&net.IPNet{
				IP:   net.ParseIP("2001:db8::1"),
				Mask: net.CIDRMask(64, 128),
			},
		}
		result := GetInterfaceIpByAddrs(addresses, false)
		assert.Empty(t, result)
	})
	
	t.Run("get_interface_ip_by_addrs_mixed_prev6_true_returns_v6", func(t *testing.T) {
		// Test with mixed addresses, prev6=true should return IPv6
		addresses := []net.Addr{
			&net.IPNet{
				IP:   net.ParseIP("192.168.1.1"),
				Mask: net.CIDRMask(24, 32),
			},
			&net.IPNet{
				IP:   net.ParseIP("fe80::1"),
				Mask: net.CIDRMask(64, 128),
			},
		}
		result := GetInterfaceIpByAddrs(addresses, true)
		assert.Contains(t, result, ":")
	})
	
	t.Run("get_interface_ip_by_addrs_mixed_prev6_true_returns_v4_if_no_v6", func(t *testing.T) {
		// Test with mixed addresses, prev6=true but no valid IPv6 addresses
		addresses := []net.Addr{
			&net.IPNet{
				IP:   net.ParseIP("192.168.1.1"),
				Mask: net.CIDRMask(24, 32),
			},
			&net.IPNet{
				IP:   net.ParseIP("10.0.0.1"),
				Mask: net.CIDRMask(8, 32),
			},
		}
		result := GetInterfaceIpByAddrs(addresses, true)
		assert.Contains(t, result, ".")
	})
}
