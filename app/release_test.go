//go:build release

package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleaseInit(t *testing.T) {
	// Test that PackageType is set to Release in release build
	assert.Equal(t, Release, PackageType, "PackageType should be Release in release build")
}

func TestReleasePackageType(t *testing.T) {
	// Verify the package type string representation
	assert.Equal(t, "release", PackageType.String())
	assert.Equal(t, "release", PackageType.Debug())
}

// BenchmarkReleasePackageType benchmarks PackageType access in release build
func BenchmarkReleasePackageType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PackageType
	}
}

// BenchmarkReleasePackageTypeString benchmarks PackageType.String() in release build
func BenchmarkReleasePackageTypeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PackageType.String()
	}
}
