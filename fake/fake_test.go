package fake

import (
	"strings"
	"testing"
)

func TestRandomUserAgent(t *testing.T) {
	// Test basic functionality
	ua := RandomUserAgent()
	
	if ua == "" {
		t.Error("RandomUserAgent() returned empty string")
	}
	
	// Test that returned user agent is from our list
	found := false
	for _, expected := range userAgents {
		if ua == expected {
			found = true
			break
		}
	}
	
	if !found {
		t.Errorf("RandomUserAgent() returned unexpected user agent: %q", ua)
	}
}

func TestRandomUserAgentReturnsValidUserAgent(t *testing.T) {
	ua := RandomUserAgent()
	
	// All user agents in our list should contain "Mozilla"
	if !strings.Contains(ua, "Mozilla") {
		t.Errorf("RandomUserAgent() returned invalid user agent (missing Mozilla): %q", ua)
	}
	
	// All user agents should contain "AppleWebKit"
	if !strings.Contains(ua, "AppleWebKit") {
		t.Errorf("RandomUserAgent() returned invalid user agent (missing AppleWebKit): %q", ua)
	}
	
	// All user agents should contain "Safari"
	if !strings.Contains(ua, "Safari") {
		t.Errorf("RandomUserAgent() returned invalid user agent (missing Safari): %q", ua)
	}
}

func TestRandomUserAgentConsistency(t *testing.T) {
	// Test that function doesn't panic and returns consistent format
	for i := 0; i < 100; i++ {
		ua := RandomUserAgent()
		
		if ua == "" {
			t.Fatalf("RandomUserAgent() returned empty string on iteration %d", i)
		}
		
		// Each user agent should be reasonably long
		if len(ua) < 50 {
			t.Errorf("RandomUserAgent() returned suspiciously short user agent: %q", ua)
		}
		
		// Should not contain line breaks or tabs
		if strings.Contains(ua, "\n") || strings.Contains(ua, "\t") || strings.Contains(ua, "\r") {
			t.Errorf("RandomUserAgent() returned user agent with invalid characters: %q", ua)
		}
	}
}

func TestRandomUserAgentDistribution(t *testing.T) {
	// Test that function returns different user agents over multiple calls
	// This is probabilistic, but with 255+ user agents, we should get variety
	
	results := make(map[string]int)
	iterations := 1000
	
	for i := 0; i < iterations; i++ {
		ua := RandomUserAgent()
		results[ua]++
	}
	
	// We should get at least 50 different user agents in 1000 calls
	// (this is very conservative given we have 255+ options)
	if len(results) < 50 {
		t.Errorf("RandomUserAgent() showed poor distribution: only %d unique user agents in %d calls", len(results), iterations)
	}
	
	// No single user agent should appear more than 5% of the time 
	// (again, very conservative)
	maxAllowed := iterations / 20
	for ua, count := range results {
		if count > maxAllowed {
			t.Errorf("RandomUserAgent() returned %q too frequently: %d times out of %d calls", ua, count, iterations)
		}
	}
}

func TestUserAgentsListContent(t *testing.T) {
	// Test the userAgents slice has expected properties
	if len(userAgents) == 0 {
		t.Fatal("userAgents slice is empty")
	}
	
	// Check that we have a reasonable number of user agents
	if len(userAgents) < 100 {
		t.Errorf("userAgents slice too small: %d entries", len(userAgents))
	}
	
	// Check each user agent for basic validity
	for i, ua := range userAgents {
		if ua == "" {
			t.Errorf("userAgents[%d] is empty", i)
		}
		
		if len(ua) < 30 {
			t.Errorf("userAgents[%d] is too short: %q", i, ua)
		}
		
		// Should start with Mozilla
		if !strings.HasPrefix(ua, "Mozilla/") {
			t.Errorf("userAgents[%d] doesn't start with Mozilla/: %q", i, ua)
		}
		
		// Should contain key browser components
		requiredComponents := []string{"Mozilla", "AppleWebKit", "Safari"}
		for _, component := range requiredComponents {
			if !strings.Contains(ua, component) {
				t.Errorf("userAgents[%d] missing %s: %q", i, component, ua)
			}
		}
	}
}

func TestUserAgentsListUniqueness(t *testing.T) {
	// Test that there are no duplicate user agents
	seen := make(map[string]int)
	
	for i, ua := range userAgents {
		if prevIndex, exists := seen[ua]; exists {
			t.Errorf("Duplicate user agent found at indices %d and %d: %q", prevIndex, i, ua)
		}
		seen[ua] = i
	}
}

func TestRandomUserAgentBrowserTypes(t *testing.T) {
	// Test that we have different types of browsers in our list
	chromeCount := 0
	windowsCount := 0
	linuxCount := 0
	macCount := 0
	androidCount := 0
	
	// Sample a reasonable number to check distribution
	for i := 0; i < 100; i++ {
		ua := RandomUserAgent()
		
		if strings.Contains(ua, "Chrome") {
			chromeCount++
		}
		if strings.Contains(ua, "Windows") {
			windowsCount++
		}
		if strings.Contains(ua, "Linux") {
			linuxCount++
		}
		if strings.Contains(ua, "Macintosh") || strings.Contains(ua, "Mac OS X") {
			macCount++
		}
		if strings.Contains(ua, "Android") {
			androidCount++
		}
	}
	
	// We should have Chrome user agents (most of our list is Chrome)
	if chromeCount == 0 {
		t.Error("No Chrome user agents found in sample")
	}
	
	// We should have Windows user agents
	if windowsCount == 0 {
		t.Error("No Windows user agents found in sample")
	}
	
	// We should have some mobile (Android) user agents
	if androidCount == 0 {
		t.Error("No Android user agents found in sample")
	}
}

func TestRandomUserAgentNoEmptyOrNil(t *testing.T) {
	// Test edge cases to ensure function is robust
	for i := 0; i < 50; i++ {
		ua := RandomUserAgent()
		
		if ua == "" {
			t.Errorf("RandomUserAgent() returned empty string on call %d", i)
		}
		
		// Test for common invalid values
		invalidValues := []string{
			"<nil>",
			"null",
			"undefined",
			" ",
			"\t",
			"\n",
		}
		
		for _, invalid := range invalidValues {
			if ua == invalid {
				t.Errorf("RandomUserAgent() returned invalid value: %q", ua)
			}
		}
	}
}

// Benchmark tests
func BenchmarkRandomUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandomUserAgent()
	}
}

func BenchmarkRandomUserAgentAllocation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = RandomUserAgent()
	}
}

// Test that the function works correctly with concurrent access
func TestRandomUserAgentConcurrency(t *testing.T) {
	results := make(chan string, 100)
	
	// Launch 10 goroutines
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				results <- RandomUserAgent()
			}
		}()
	}
	
	// Collect results
	for i := 0; i < 100; i++ {
		ua := <-results
		if ua == "" {
			t.Error("Concurrent access resulted in empty user agent")
		}
		
		// Verify it's a valid user agent from our list
		found := false
		for _, expected := range userAgents {
			if ua == expected {
				found = true
				break
			}
		}
		
		if !found {
			t.Errorf("Concurrent access returned invalid user agent: %q", ua)
		}
	}
}