package network

import "testing"

func TestIsLocalIp(t *testing.T) {
	tests := []struct {
		name     string
		ip       string
		expected bool
	}{
		{"IPv4 Private 192.168.x.x", "192.168.1.1", true},
		{"IPv4 Private 10.x.x.x", "10.0.0.1", true},
		{"IPv4 Private 172.16-31.x.x", "172.16.0.1", true},
		{"IPv4 Localhost", "127.0.0.1", true},
		{"IPv4 Public Google DNS", "8.8.8.8", false},
		{"IPv4 Public Cloudflare DNS", "1.1.1.1", false},
		{"IPv6 Loopback", "::1", true},
		{"IPv6 Private fc00::/7", "fc00::1", true},
		{"IPv6 Private fd00::/8", "fd00::1", true},
		{"IPv6 Link-local fe80::/10", "fe80::1", true},
		{"IPv6 Public Google DNS", "2001:4860:4860::8888", false},
		{"IPv6 Public Cloudflare DNS", "2606:4700:4700::1111", false},
		{"Invalid IP empty", "", false},
		{"Invalid IP format", "invalid-ip", false},
		{"Invalid IP incomplete", "192.168.1", false},
		{"Invalid IP overflow", "256.256.256.256", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsLocalIp(tt.ip)
			if result != tt.expected {
				t.Errorf("IsLocalIp(%s) = %v, expected %v", tt.ip, result, tt.expected)
			}
		})
	}
}