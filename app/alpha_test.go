//go:build alpha

package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlphaInit(t *testing.T) {
	// Test that PackageType is set to Alpha in alpha build
	assert.Equal(t, Alpha, PackageType, "PackageType should be Alpha in alpha build")
}

func TestAlphaPackageType(t *testing.T) {
	// Verify the package type string representation
	assert.Equal(t, "alpha", PackageType.String())
	assert.Equal(t, "alpha", PackageType.Debug())
}

// BenchmarkAlphaPackageType benchmarks PackageType access in alpha build
func BenchmarkAlphaPackageType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PackageType
	}
}

// BenchmarkAlphaPackageTypeString benchmarks PackageType.String() in alpha build
func BenchmarkAlphaPackageTypeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PackageType.String()
	}
}
