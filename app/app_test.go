package app

import (
	"testing"
)

func TestReleaseTypeString(t *testing.T) {
	tests := []struct {
		name     string
		release  ReleaseType
		expected string
	}{
		{"Debug", Debug, "debug"},
		{"Test", Test, "test"},
		{"Alpha", Alpha, "alpha"},
		{"Beta", Beta, "beta"},
		{"Release", Release, "release"},
		{"Unknown", ReleaseType(99), "debug"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.release.String()
			if result != tt.expected {
				t.Errorf("ReleaseType.String() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestReleaseTypeDebug(t *testing.T) {
	tests := []struct {
		name     string
		release  ReleaseType
		expected string
	}{
		{"Debug", Debug, "debug"},
		{"Test", Test, "test"},
		{"Alpha", Alpha, "alpha"},
		{"Beta", Beta, "beta"},
		{"Release", Release, "release"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.release.Debug()
			if result != tt.expected {
				t.Errorf("ReleaseType.Debug() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestOrganization(t *testing.T) {
	if Organization == "" {
		t.Error("Organization should not be empty")
	}

	expectedOrganization := "lazygophers"
	if Organization != expectedOrganization {
		t.Errorf("Organization = %s, expected %s", Organization, expectedOrganization)
	}
}

func TestGlobalVariables(t *testing.T) {
	tests := []struct {
		name  string
		value string
		check bool
	}{
		{"Commit", Commit, false},
		{"ShortCommit", ShortCommit, false},
		{"Branch", Branch, false},
		{"Tag", Tag, false},
		{"BuildDate", BuildDate, false},
		{"GoVersion", GoVersion, false},
		{"GoOS", GoOS, false},
		{"Goarch", Goarch, false},
		{"Goarm", Goarm, false},
		{"Goamd64", Goamd64, false},
		{"Gomips", Gomips, false},
		{"Description", Description, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.check && tt.value == "" {
				t.Errorf("%s should not be empty", tt.name)
			}
			t.Logf("%s = %s", tt.name, tt.value)
		})
	}
}

func TestReleaseTypeConstants(t *testing.T) {
	tests := []struct {
		name     string
		release  ReleaseType
		expected uint8
	}{
		{"Debug", Debug, 0},
		{"Test", Test, 1},
		{"Alpha", Alpha, 2},
		{"Beta", Beta, 3},
		{"Release", Release, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if uint8(tt.release) != tt.expected {
				t.Errorf("ReleaseType value = %d, expected %d", tt.release, tt.expected)
			}
		})
	}
}

func TestPackageTypeIsValid(t *testing.T) {
	validTypes := []ReleaseType{Debug, Test, Alpha, Beta, Release}

	for _, validType := range validTypes {
		t.Run(validType.String(), func(t *testing.T) {
			if validType < Debug || validType > Release {
				t.Errorf("ReleaseType %v is not within valid range", validType)
			}
		})
	}
}

func TestReleaseTypeStringConsistency(t *testing.T) {
	tests := []struct {
		release  ReleaseType
		expected string
	}{
		{Debug, "debug"},
		{Test, "test"},
		{Alpha, "alpha"},
		{Beta, "beta"},
		{Release, "release"},
	}

	for _, tt := range tests {
		t.Run(tt.release.String(), func(t *testing.T) {
			str1 := tt.release.String()
			str2 := tt.release.Debug()

			if str1 != str2 {
				t.Errorf("String() and Debug() returned different values: %s vs %s", str1, str2)
			}

			if str1 != tt.expected {
				t.Errorf("String() = %s, expected %s", str1, tt.expected)
			}
		})
	}
}

func TestPackageTypeValue(t *testing.T) {
	packageType := PackageType

	if packageType < Debug || packageType > Release {
		t.Errorf("PackageType %v is not within valid range", packageType)
	}

	t.Logf("PackageType = %v (%s)", packageType, packageType.String())
}

func TestReleaseTypeValues(t *testing.T) {
	tests := []struct {
		name    string
		value   ReleaseType
		numeric uint8
	}{
		{"Debug", Debug, 0},
		{"Test", Test, 1},
		{"Alpha", Alpha, 2},
		{"Beta", Beta, 3},
		{"Release", Release, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if uint8(tt.value) != tt.numeric {
				t.Errorf("ReleaseType %v has numeric value %d, expected %d", tt.value, uint8(tt.value), tt.numeric)
			}
		})
	}
}
