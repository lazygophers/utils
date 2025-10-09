package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvVariables(t *testing.T) {
	tests := []struct {
		name     string
		varPtr   *string
		varName  string
		testVal  string
	}{
		{"Commit", &Commit, "Commit", "abc123def456"},
		{"ShortCommit", &ShortCommit, "ShortCommit", "abc123d"},
		{"Branch", &Branch, "Branch", "main"},
		{"Tag", &Tag, "Tag", "v1.2.3"},
		{"BuildDate", &BuildDate, "BuildDate", "2024-10-09T12:00:00Z"},
		{"GoVersion", &GoVersion, "GoVersion", "go1.24"},
		{"GoOS", &GoOS, "GoOS", "linux"},
		{"Goarch", &Goarch, "Goarch", "amd64"},
		{"Goarm", &Goarm, "Goarm", "7"},
		{"Goamd64", &Goamd64, "Goamd64", "v3"},
		{"Gomips", &Gomips, "Gomips", "hardfloat"},
		{"Description", &Description, "Description", "Test application"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original value
			original := *tt.varPtr

			// Test assignment
			*tt.varPtr = tt.testVal
			assert.Equal(t, tt.testVal, *tt.varPtr, "%s should be settable", tt.varName)

			// Test that the variable exists and can be read
			assert.NotNil(t, tt.varPtr, "%s should not be nil", tt.varName)

			// Restore original value
			*tt.varPtr = original
		})
	}
}

func TestEnvVariablesReadOnly(t *testing.T) {
	// Test that variables can be read without modifying
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
}

func TestEnvVariablesConcurrency(t *testing.T) {
	// Test concurrent read access to env variables
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func() {
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
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestEnvVariablesEmptyDefault(t *testing.T) {
	// Test that variables start as empty strings by default
	// (unless set by build flags)
	// This is more of a documentation test
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
	}

	for varName, varPtr := range variables {
		t.Run("Default_"+varName, func(t *testing.T) {
			// Just verify the variable is accessible
			// The actual value depends on build flags
			_ = *varPtr
		})
	}
}

// BenchmarkEnvVariableAccess benchmarks direct access to env variables
func BenchmarkEnvVariableAccess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Commit
		_ = ShortCommit
		_ = Branch
		_ = Tag
	}
}

// BenchmarkEnvVariableWrite benchmarks writing to env variables
func BenchmarkEnvVariableWrite(b *testing.B) {
	original := Commit
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Commit = "test-commit"
	}

	Commit = original
}

// BenchmarkEnvVariableConcurrentRead benchmarks concurrent reads
func BenchmarkEnvVariableConcurrentRead(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Commit
			_ = Branch
			_ = Tag
		}
	})
}
