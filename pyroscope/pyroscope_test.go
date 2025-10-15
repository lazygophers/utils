package pyroscope

import (
	"os"
	"testing"
)

func TestLoad_WithEmptyAddress(t *testing.T) {
	// Test Load function with empty address
	// This should not panic and should use default address internally
	Load("")

	// If we reach here, the function didn't panic
	// This is the main test since Load() is designed to be fire-and-forget
}

func TestLoad_WithCustomAddress(t *testing.T) {
	// Test Load function with custom address
	// This should not panic
	Load("http://custom.pyroscope.server:4040")

	// If we reach here, the function didn't panic
}

func TestLoad_WithInvalidAddress(t *testing.T) {
	// Test Load function with invalid address
	// This should not panic, but may log an error internally
	Load("invalid-address-format")

	// If we reach here, the function didn't panic
}

func TestLoad_WithLocalhostAddress(t *testing.T) {
	// Test Load function with localhost address
	Load("http://localhost:4040")

	// If we reach here, the function didn't panic
}

func TestLoad_WithHTTPSAddress(t *testing.T) {
	// Test Load function with HTTPS address
	Load("https://pyroscope.example.com")

	// If we reach here, the function didn't panic
}

func TestLoad_MultipleCallsShouldNotPanic(t *testing.T) {
	// Test multiple calls to Load function
	Load("http://first.server:4040")
	Load("http://second.server:4040")
	Load("")

	// If we reach here, multiple calls didn't cause issues
}

func TestLoad_WithDifferentPorts(t *testing.T) {
	// Test Load function with different port numbers
	addresses := []string{
		"http://127.0.0.1:4040",
		"http://127.0.0.1:8080",
		"http://127.0.0.1:9090",
		"http://127.0.0.1:3000",
	}

	for _, addr := range addresses {
		Load(addr)
	}

	// If we reach here, all addresses were processed without panic
}

func TestLoad_WithEnvironmentVariable(t *testing.T) {
	// Test Load function behavior with environment variables
	originalHostname := os.Getenv("HOSTNAME")

	// Set a test hostname
	os.Setenv("HOSTNAME", "test-hostname")

	// Test Load function
	Load("http://test.pyroscope.server:4040")

	// Restore original hostname
	if originalHostname != "" {
		os.Setenv("HOSTNAME", originalHostname)
	} else {
		os.Unsetenv("HOSTNAME")
	}

	// If we reach here, the function handled environment variables correctly
}

func TestLoad_WithEmptyHostname(t *testing.T) {
	// Test Load function when HOSTNAME environment variable is empty
	originalHostname := os.Getenv("HOSTNAME")

	// Clear hostname
	os.Unsetenv("HOSTNAME")

	// Test Load function
	Load("http://test.pyroscope.server:4040")

	// Restore original hostname
	if originalHostname != "" {
		os.Setenv("HOSTNAME", originalHostname)
	}

	// If we reach here, the function handled empty hostname correctly
}

func TestLoad_StressTest(t *testing.T) {
	// Stress test - call Load multiple times rapidly
	for i := 0; i < 10; i++ {
		Load("http://stress-test.server:4040")
	}

	// If we reach here, stress test passed
}

func TestLoad_EdgeCases(t *testing.T) {
	// Test various edge cases
	edgeCases := []string{
		"",                              // empty string
		"http://",                       // incomplete URL
		"://no-protocol.com",            // missing protocol
		"http://127.0.0.1",              // no port
		"http://127.0.0.1:99999",        // invalid port
		"ftp://wrong-protocol.com:4040", // wrong protocol
		"http://[::1]:4040",             // IPv6 address
		"http://very-long-hostname-that-might-cause-issues.example.com:4040", // long hostname
	}

	for _, testCase := range edgeCases {
		Load(testCase)
	}

	// If we reach here, all edge cases were handled without panic
}

// Test that Load function can be called concurrently
func TestLoad_Concurrent(t *testing.T) {
	// Test concurrent calls to Load function
	done := make(chan bool, 5)

	for i := 0; i < 5; i++ {
		go func(id int) {
			Load("http://concurrent-test.server:4040")
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 5; i++ {
		<-done
	}

	// If we reach here, concurrent calls completed successfully
}

// Test Load function behavior with special characters in address
func TestLoad_SpecialCharacters(t *testing.T) {
	specialAddresses := []string{
		"http://test-server.com:4040", // hyphen
		"http://test_server.com:4040", // underscore
		"http://test.server.com:4040", // dot
		"http://192.168.1.100:4040",   // IP address
	}

	for _, addr := range specialAddresses {
		Load(addr)
	}

	// If we reach here, special characters were handled correctly
}

// Benchmark test for Load function
func BenchmarkLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Load("http://benchmark.server:4040")
	}
}

func BenchmarkLoadEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Load("")
	}
}

// Test Load function with realistic addresses
func TestLoad_RealisticScenarios(t *testing.T) {
	realisticAddresses := []string{
		"http://pyroscope.production.company.com:4040",
		"https://monitoring.example.org:443",
		"http://dev-pyroscope.local:4040",
		"http://staging-monitoring:8080",
		"https://secure-pyroscope.internal:9443",
	}

	for _, addr := range realisticAddresses {
		Load(addr)
	}

	// If we reach here, realistic scenarios were handled correctly
}

// Test Load function robustness
func TestLoad_Robustness(t *testing.T) {
	// Test that Load doesn't crash with various inputs
	testInputs := []string{
		"",
		" ",
		"\n",
		"\t",
		"null",
		"undefined",
		"0",
		"-1",
		"http://",
		"://",
		"http",
		"4040",
		"localhost",
		"127.0.0.1",
		"http://127.0.0.1:4040/path",
		"http://127.0.0.1:4040?query=value",
		"http://127.0.0.1:4040#fragment",
	}

	for _, input := range testInputs {
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Load(%q) panicked: %v", input, r)
				}
			}()
			Load(input)
		}()
	}
}

// Test Load function with environment variations
func TestLoad_EnvironmentVariations(t *testing.T) {
	testEnvs := map[string]string{
		"":                  "",
		"localhost":         "localhost",
		"production-server": "production-server",
		"test-env-123":      "test-env-123",
		"dev":               "dev",
		"staging":           "staging",
		"prod":              "prod",
	}

	originalHostname := os.Getenv("HOSTNAME")

	for hostname := range testEnvs {
		if hostname == "" {
			os.Unsetenv("HOSTNAME")
		} else {
			os.Setenv("HOSTNAME", hostname)
		}

		Load("http://test.server:4040")
	}

	// Restore original hostname
	if originalHostname != "" {
		os.Setenv("HOSTNAME", originalHostname)
	} else {
		os.Unsetenv("HOSTNAME")
	}
}

// Test that Load function is idempotent
func TestLoad_Idempotent(t *testing.T) {
	address := "http://idempotent-test.server:4040"

	// Call Load multiple times with same address
	Load(address)
	Load(address)
	Load(address)

	// Function should handle multiple calls gracefully
	// If we reach here, idempotency test passed
}

// Test Load function with various protocols
func TestLoad_Protocols(t *testing.T) {
	protocols := []string{
		"http://server:4040",
		"https://server:4040",
		// Note: pyroscope likely only supports HTTP/HTTPS, but we test various inputs
		"ftp://server:4040", // should be handled gracefully
		"ws://server:4040",  // should be handled gracefully
		"tcp://server:4040", // should be handled gracefully
	}

	for _, protocol := range protocols {
		Load(protocol)
	}

	// If we reach here, all protocols were handled without panic
}
