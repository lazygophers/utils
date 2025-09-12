package network

import (
	"net/http"
	"testing"
)

func TestRealIpFromHeader(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string]string
		expected string
	}{
		// Cloudflare headers
		{
			name:     "Cf-Connecting-Ip with public IP",
			headers:  map[string]string{"Cf-Connecting-Ip": "203.0.113.1"},
			expected: "203.0.113.1",
		},
		{
			name:     "Cf-Connecting-Ip with private IP (ignored)",
			headers:  map[string]string{"Cf-Connecting-Ip": "192.168.1.1"},
			expected: "",
		},
		{
			name:     "Cf-Pseudo-Ipv4 with public IP",
			headers:  map[string]string{"Cf-Pseudo-Ipv4": "203.0.113.2"},
			expected: "203.0.113.2",
		},
		{
			name:     "Cf-Pseudo-Ipv4 with private IP (ignored)",
			headers:  map[string]string{"Cf-Pseudo-Ipv4": "10.0.0.1"},
			expected: "",
		},
		{
			name:     "Cf-Connecting-Ipv6 with public IPv6",
			headers:  map[string]string{"Cf-Connecting-Ipv6": "2001:db8::1"},
			expected: "2001:db8::1",
		},
		{
			name:     "Cf-Connecting-Ipv6 with private IPv6 (ignored)",
			headers:  map[string]string{"Cf-Connecting-Ipv6": "fc00::1"},
			expected: "",
		},
		{
			name:     "Cf-Pseudo-Ipv6 with public IPv6",
			headers:  map[string]string{"Cf-Pseudo-Ipv6": "2001:db8::2"},
			expected: "2001:db8::2",
		},

		// Other CDN headers
		{
			name:     "Fastly-Client-Ip with public IP",
			headers:  map[string]string{"Fastly-Client-Ip": "203.0.113.3"},
			expected: "203.0.113.3",
		},
		{
			name:     "True-Client-Ip with public IP",
			headers:  map[string]string{"True-Client-Ip": "203.0.113.4"},
			expected: "203.0.113.4",
		},

		// Common proxy headers
		{
			name:     "X-Real-IP with public IP",
			headers:  map[string]string{"X-Real-IP": "203.0.113.5"},
			expected: "203.0.113.5",
		},
		{
			name:     "X-Client-IP with public IP",
			headers:  map[string]string{"X-Client-IP": "203.0.113.6"},
			expected: "203.0.113.6",
		},

		// Forwarded headers with multiple IPs
		{
			name:     "X-Original-Forwarded-For with multiple IPs",
			headers:  map[string]string{"X-Original-Forwarded-For": "192.168.1.1, 203.0.113.7, 10.0.0.1"},
			expected: "203.0.113.7",
		},
		{
			name:     "X-Original-Forwarded-For with spaces",
			headers:  map[string]string{"X-Original-Forwarded-For": " 192.168.1.1 , 203.0.113.8 , 10.0.0.1 "},
			expected: "203.0.113.8",
		},
		{
			name:     "X-Original-Forwarded-For with all private IPs",
			headers:  map[string]string{"X-Original-Forwarded-For": "192.168.1.1, 10.0.0.1, 172.16.0.1"},
			expected: "",
		},

		{
			name:     "X-Forwarded-For with multiple IPs",
			headers:  map[string]string{"X-Forwarded-For": "192.168.1.1, 203.0.113.9, 10.0.0.1"},
			expected: "203.0.113.9",
		},
		{
			name:     "X-Forwarded-For with spaces",
			headers:  map[string]string{"X-Forwarded-For": " 192.168.1.1 , 203.0.113.10 , 10.0.0.1 "},
			expected: "203.0.113.10",
		},

		{
			name:     "X-Forwarded with multiple IPs",
			headers:  map[string]string{"X-Forwarded": "192.168.1.1, 203.0.113.11, 10.0.0.1"},
			expected: "203.0.113.11",
		},

		{
			name:     "Forwarded-For with multiple IPs",
			headers:  map[string]string{"Forwarded-For": "192.168.1.1, 203.0.113.12, 10.0.0.1"},
			expected: "203.0.113.12",
		},

		{
			name:     "Forwarded with multiple IPs",
			headers:  map[string]string{"Forwarded": "192.168.1.1, 203.0.113.13, 10.0.0.1"},
			expected: "203.0.113.13",
		},

		// Priority tests - earlier headers should take precedence
		{
			name: "Multiple headers - Cloudflare takes priority",
			headers: map[string]string{
				"Cf-Connecting-Ip": "203.0.113.100",
				"X-Real-IP":        "203.0.113.200",
				"X-Forwarded-For":  "203.0.113.300",
			},
			expected: "203.0.113.100",
		},
		{
			name: "Multiple headers - X-Real-IP when no Cloudflare",
			headers: map[string]string{
				"X-Real-IP":       "203.0.113.200",
				"X-Forwarded-For": "203.0.113.300",
			},
			expected: "203.0.113.200",
		},

		// Edge cases
		{
			name:     "Empty header value",
			headers:  map[string]string{"X-Real-IP": ""},
			expected: "",
		},
		{
			name:     "No relevant headers",
			headers:  map[string]string{"User-Agent": "test"},
			expected: "",
		},
		{
			name:     "Empty header",
			headers:  map[string]string{},
			expected: "",
		},

		// Invalid IP handling
		{
			name:     "X-Forwarded-For with invalid IP",
			headers:  map[string]string{"X-Forwarded-For": "invalid-ip, 203.0.113.14"},
			expected: "203.0.113.14",
		},
		{
			name:     "X-Forwarded-For with only empty values",
			headers:  map[string]string{"X-Forwarded-For": " , , "},
			expected: "",
		},

		// IPv6 cases
		{
			name:     "X-Forwarded-For with IPv6",
			headers:  map[string]string{"X-Forwarded-For": "::1, 2001:db8::100"},
			expected: "2001:db8::100",
		},

		// Test all private IP ranges
		{
			name:     "Class A private IP (10.x.x.x)",
			headers:  map[string]string{"X-Real-IP": "10.255.255.255"},
			expected: "",
		},
		{
			name:     "Class B private IP (172.16-31.x.x)",
			headers:  map[string]string{"X-Real-IP": "172.31.255.255"},
			expected: "",
		},
		{
			name:     "Class C private IP (192.168.x.x)",
			headers:  map[string]string{"X-Real-IP": "192.168.255.255"},
			expected: "",
		},
		{
			name:     "Localhost",
			headers:  map[string]string{"X-Real-IP": "127.0.0.1"},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := make(http.Header)
			for key, value := range tt.headers {
				header.Set(key, value)
			}

			result := RealIpFromHeader(header)
			if result != tt.expected {
				t.Errorf("RealIpFromHeader() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

func TestRealIpFromHeader_HeaderCaseSensitivity(t *testing.T) {
	// HTTP headers are case-insensitive, test different cases
	tests := []struct {
		name       string
		headerKey  string
		headerVal  string
		expected   string
	}{
		{"Lowercase cf-connecting-ip", "cf-connecting-ip", "203.0.113.1", "203.0.113.1"},
		{"Uppercase CF-CONNECTING-IP", "CF-CONNECTING-IP", "203.0.113.2", "203.0.113.2"},
		{"Mixed case Cf-Connecting-Ip", "Cf-Connecting-Ip", "203.0.113.3", "203.0.113.3"},
		{"Lowercase x-real-ip", "x-real-ip", "203.0.113.4", "203.0.113.4"},
		{"Uppercase X-REAL-IP", "X-REAL-IP", "203.0.113.5", "203.0.113.5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := make(http.Header)
			header.Set(tt.headerKey, tt.headerVal)

			result := RealIpFromHeader(header)
			if result != tt.expected {
				t.Errorf("RealIpFromHeader() with header %s = %s, expected %s", 
					tt.headerKey, result, tt.expected)
			}
		})
	}
}

func TestRealIpFromHeader_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string]string
		expected string
	}{
		{
			name:     "Single comma with spaces",
			headers:  map[string]string{"X-Forwarded-For": " , "},
			expected: "",
		},
		{
			name:     "Multiple commas only",
			headers:  map[string]string{"X-Forwarded-For": ",,,"},
			expected: "",
		},
		{
			name:     "Mix of valid and invalid IPs",
			headers:  map[string]string{"X-Forwarded-For": "invalid, 192.168.1.1, 203.0.113.99, another-invalid"},
			expected: "203.0.113.99",
		},
		{
			name:     "Very long IP list",
			headers:  map[string]string{"X-Forwarded-For": "192.168.1.1, 10.0.0.1, 172.16.0.1, 203.0.113.200, 192.168.2.1, 10.0.0.2"},
			expected: "203.0.113.200",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := make(http.Header)
			for key, value := range tt.headers {
				header.Set(key, value)
			}

			result := RealIpFromHeader(header)
			if result != tt.expected {
				t.Errorf("RealIpFromHeader() = %s, expected %s", result, tt.expected)
			}
		})
	}
}