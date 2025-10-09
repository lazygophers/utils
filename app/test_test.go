//go:build test

package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestInit(t *testing.T) {
	// Test that PackageType is set to Test in test build
	assert.Equal(t, Test, PackageType, "PackageType should be Test in test build")
}

func TestTestPackageType(t *testing.T) {
	// Verify the package type string representation
	assert.Equal(t, "test", PackageType.String())
	assert.Equal(t, "test", PackageType.Debug())
}

// BenchmarkTestPackageType benchmarks PackageType access in test build
func BenchmarkTestPackageType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PackageType
	}
}

// BenchmarkTestPackageTypeString benchmarks PackageType.String() in test build
func BenchmarkTestPackageTypeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PackageType.String()
	}
}
