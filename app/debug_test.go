//go:build debug

package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebugInit(t *testing.T) {
	// Test that PackageType is set to Debug in debug build
	assert.Equal(t, Debug, PackageType, "PackageType should be Debug in debug build")
}

func TestDebugPackageType(t *testing.T) {
	// Verify the package type string representation
	assert.Equal(t, "debug", PackageType.String())
	assert.Equal(t, "debug", PackageType.Debug())
}

// BenchmarkDebugPackageType benchmarks PackageType access in debug build
func BenchmarkDebugPackageType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PackageType
	}
}

// BenchmarkDebugPackageTypeString benchmarks PackageType.String() in debug build
func BenchmarkDebugPackageTypeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PackageType.String()
	}
}
