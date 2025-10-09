//go:build beta

package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBetaInit(t *testing.T) {
	// Test that PackageType is set to Beta in beta build
	assert.Equal(t, Beta, PackageType, "PackageType should be Beta in beta build")
}

func TestBetaPackageType(t *testing.T) {
	// Verify the package type string representation
	assert.Equal(t, "beta", PackageType.String())
	assert.Equal(t, "beta", PackageType.Debug())
}

// BenchmarkBetaPackageType benchmarks PackageType access in beta build
func BenchmarkBetaPackageType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PackageType
	}
}

// BenchmarkBetaPackageTypeString benchmarks PackageType.String() in beta build
func BenchmarkBetaPackageTypeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PackageType.String()
	}
}
