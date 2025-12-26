package network

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInterfaceIpByAddrs(t *testing.T) {
	t.Run("ipv4_addresses", func(t *testing.T) {
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
		result := GetInterfaceIpByAddrs(addresses, false)
		assert.NotEmpty(t, result)
		assert.Contains(t, result, ".")
	})

	t.Run("ipv6_addresses", func(t *testing.T) {
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
		result := GetInterfaceIpByAddrs(addresses, true)
		assert.NotEmpty(t, result)
		assert.Contains(t, result, ":")
	})

	t.Run("mixed_addresses_ipv4", func(t *testing.T) {
		addresses := []net.Addr{
			&net.IPNet{
				IP:   net.ParseIP("fe80::1"),
				Mask: net.CIDRMask(64, 128),
			},
			&net.IPNet{
				IP:   net.ParseIP("192.168.1.1"),
				Mask: net.CIDRMask(24, 32),
			},
		}
		result := GetInterfaceIpByAddrs(addresses, false)
		assert.NotEmpty(t, result)
		assert.Contains(t, result, ".")
	})

	t.Run("mixed_addresses_ipv6", func(t *testing.T) {
		addresses := []net.Addr{
			&net.IPNet{
				IP:   net.ParseIP("fe80::1"),
				Mask: net.CIDRMask(64, 128),
			},
			&net.IPNet{
				IP:   net.ParseIP("192.168.1.1"),
				Mask: net.CIDRMask(24, 32),
			},
		}
		result := GetInterfaceIpByAddrs(addresses, true)
		assert.NotEmpty(t, result)
		assert.Contains(t, result, ":")
	})

	t.Run("loopback_addresses", func(t *testing.T) {
		addresses := []net.Addr{
			&net.IPNet{
				IP:   net.ParseIP("127.0.0.1"),
				Mask: net.CIDRMask(8, 32),
			},
			&net.IPNet{
				IP:   net.ParseIP("::1"),
				Mask: net.CIDRMask(128, 128),
			},
		}
		result := GetInterfaceIpByAddrs(addresses, false)
		assert.Empty(t, result)
	})

	t.Run("empty_addresses", func(t *testing.T) {
		addresses := []net.Addr{}
		result := GetInterfaceIpByAddrs(addresses, false)
		assert.Empty(t, result)
	})

	t.Run("ipv6_fallback_to_ipv4", func(t *testing.T) {
		addresses := []net.Addr{
			&net.IPNet{
				IP:   net.ParseIP("fe80::1"),
				Mask: net.CIDRMask(64, 128),
			},
		}
		result := GetInterfaceIpByAddrs(addresses, false)
		assert.Empty(t, result)
	})
}

func TestGetInterfaceIpByName(t *testing.T) {
	t.Run("existing_interface_lo0", func(t *testing.T) {
		result := GetInterfaceIpByName("lo0", false)
		assert.Empty(t, result)
	})

	t.Run("existing_interface_lo0_ipv6", func(t *testing.T) {
		result := GetInterfaceIpByName("lo0", true)
		assert.NotEmpty(t, result)
		assert.Contains(t, result, ":")
	})

	t.Run("non_existing_interface", func(t *testing.T) {
		result := GetInterfaceIpByName("nonexistent123", false)
		assert.Empty(t, result)
	})

	t.Run("empty_interface_name", func(t *testing.T) {
		result := GetInterfaceIpByName("", false)
		assert.Empty(t, result)
	})
}

func TestGetListenIp(t *testing.T) {
	t.Run("default_ipv4", func(t *testing.T) {
		result := GetListenIp()
		if result == "" {
			t.Skip("No non-loopback IPv4 address found")
		}
		assert.NotEmpty(t, result)
	})

	t.Run("explicit_ipv4", func(t *testing.T) {
		result := GetListenIp(false)
		if result == "" {
			t.Skip("No non-loopback IPv4 address found")
		}
		assert.NotEmpty(t, result)
	})

	t.Run("ipv6", func(t *testing.T) {
		result := GetListenIp(true)
		if result == "" {
			t.Skip("No non-loopback IPv6 address found")
		}
		assert.NotEmpty(t, result)
	})

	t.Run("multiple_params", func(t *testing.T) {
		result := GetListenIp(false, true)
		if result == "" {
			t.Skip("No non-loopback address found")
		}
		assert.NotEmpty(t, result)
	})
}
