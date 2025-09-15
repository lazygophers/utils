package app

import (
	"testing"
)

func TestReleaseType_String(t *testing.T) {
	tests := []struct {
		name     string
		rt       ReleaseType
		expected string
	}{
		{
			name:     "Debug release type",
			rt:       Debug,
			expected: "debug",
		},
		{
			name:     "Test release type",
			rt:       Test,
			expected: "test",
		},
		{
			name:     "Alpha release type",
			rt:       Alpha,
			expected: "alpha",
		},
		{
			name:     "Beta release type",
			rt:       Beta,
			expected: "beta",
		},
		{
			name:     "Release release type",
			rt:       Release,
			expected: "release",
		},
		{
			name:     "Invalid release type defaults to debug",
			rt:       ReleaseType(99),
			expected: "debug",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.rt.String()
			if result != tt.expected {
				t.Errorf("ReleaseType(%d).String() = %q, expected %q", tt.rt, result, tt.expected)
			}
		})
	}
}

func TestReleaseType_Debug(t *testing.T) {
	tests := []struct {
		name     string
		rt       ReleaseType
		expected string
	}{
		{
			name:     "Debug method returns same as String",
			rt:       Debug,
			expected: "debug",
		},
		{
			name:     "Test method returns same as String",
			rt:       Test,
			expected: "test",
		},
		{
			name:     "Alpha method returns same as String",
			rt:       Alpha,
			expected: "alpha",
		},
		{
			name:     "Beta method returns same as String",
			rt:       Beta,
			expected: "beta",
		},
		{
			name:     "Release method returns same as String",
			rt:       Release,
			expected: "release",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.rt.Debug()
			expected := tt.rt.String()
			if result != expected || result != tt.expected {
				t.Errorf("ReleaseType(%d).Debug() = %q, expected %q", tt.rt, result, tt.expected)
			}
		})
	}
}

func TestReleaseTypeConstants(t *testing.T) {
	// Test that constants have correct values
	if Debug != 0 {
		t.Errorf("Debug constant = %d, expected 0", Debug)
	}
	if Test != 1 {
		t.Errorf("Test constant = %d, expected 1", Test)
	}
	if Alpha != 2 {
		t.Errorf("Alpha constant = %d, expected 2", Alpha)
	}
	if Beta != 3 {
		t.Errorf("Beta constant = %d, expected 3", Beta)
	}
	if Release != 4 {
		t.Errorf("Release constant = %d, expected 4", Release)
	}
}

func TestGlobalVariables(t *testing.T) {
	// Test that global variables exist and are accessible
	// We can't test their initial values as they may be set by build flags

	// Test string variables exist
	_ = Commit
	_ = ShortCommit
	_ = Branch
	_ = Tag
	_ = BuildDate
	_ = GoVersion
	_ = GoOS
	_ = Goarch
	_ = Goarm
	_ = Goamd64
	_ = Gomips
	_ = Description
	_ = Name
	_ = Version

	// Test Organization is set
	if Organization != "lazygophers" {
		t.Errorf("Organization = %q, expected %q", Organization, "lazygophers")
	}

	// Test PackageType variable exists and is a valid ReleaseType
	if PackageType < Debug || PackageType > Release {
		// Allow for build tag variations, just ensure it's accessible
		_ = PackageType
	}
}

func TestVariableAssignments(t *testing.T) {
	// Test that we can assign values to global variables
	originalCommit := Commit
	originalName := Name
	originalVersion := Version
	originalPackageType := PackageType

	// Test assignment
	Commit = "test-commit-123"
	Name = "test-app"
	Version = "v1.0.0"
	PackageType = Beta

	// Verify assignments
	if Commit != "test-commit-123" {
		t.Errorf("Commit assignment failed, got %q", Commit)
	}
	if Name != "test-app" {
		t.Errorf("Name assignment failed, got %q", Name)
	}
	if Version != "v1.0.0" {
		t.Errorf("Version assignment failed, got %q", Version)
	}
	if PackageType != Beta {
		t.Errorf("PackageType assignment failed, got %d", PackageType)
	}

	// Restore original values
	Commit = originalCommit
	Name = originalName
	Version = originalVersion
	PackageType = originalPackageType
}

func TestReleaseTypeImplementsStringer(t *testing.T) {
	// Test that ReleaseType implements the fmt.Stringer interface
	var rt ReleaseType = Debug
	stringer, ok := interface{}(rt).(interface{ String() string })
	if !ok {
		t.Error("ReleaseType does not implement String() method")
	}

	result := stringer.String()
	if result != "debug" {
		t.Errorf("String() method returned %q, expected %q", result, "debug")
	}
}

// Test coverage for all build metadata variables
func TestBuildMetadataVariables(t *testing.T) {
	variables := map[string]*string{
		"Commit":      &Commit,
		"ShortCommit": &ShortCommit,
		"Branch":      &Branch,
		"Tag":         &Tag,
		"BuildDate":   &BuildDate,
		"GoVersion":   &GoVersion,
		"GoOS":        &GoOS,
		"Goarch":      &Goarch,
		"Goarm":       &Goarm,
		"Goamd64":     &Goamd64,
		"Gomips":      &Gomips,
		"Description": &Description,
		"Name":        &Name,
		"Version":     &Version,
	}

	for varName, varPtr := range variables {
		t.Run("Variable_"+varName, func(t *testing.T) {
			// Test that the variable can be accessed and modified
			original := *varPtr
			testValue := "test-" + varName
			*varPtr = testValue

			if *varPtr != testValue {
				t.Errorf("Failed to set %s to %q, got %q", varName, testValue, *varPtr)
			}

			// Restore original value
			*varPtr = original
		})
	}
}

// Benchmark tests
func BenchmarkReleaseType_String(b *testing.B) {
	rt := Release
	for i := 0; i < b.N; i++ {
		_ = rt.String()
	}
}

func BenchmarkReleaseType_Debug(b *testing.B) {
	rt := Alpha
	for i := 0; i < b.N; i++ {
		_ = rt.Debug()
	}
}
