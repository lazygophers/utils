package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrongPasswordOldFunc(t *testing.T) {
	assert.True(t, validateStrongPasswordOld("MyPass123!"))
	assert.False(t, validateStrongPasswordOld("weak"))
}

func TestStrongPasswordFastFunc(t *testing.T) {
	assert.True(t, validateStrongPasswordFast("MyPass123!"))
	assert.False(t, validateStrongPasswordFast("weak"))
}

func TestStrongPasswordPerformanceRuns(t *testing.T) {
	// Run the performance comparison test
	TestStrongPassword_Performance(t)
}

func BenchmarkOldWrapper(b *testing.B) {
	BenchmarkStrongPassword_Old(b)
}

func BenchmarkNewWrapper(b *testing.B) {
	BenchmarkStrongPassword_New(b)
}

func TestBenchmarkOldWrapper(t *testing.T) {
	passwords := []string{
		"Abc123!@",
		"Password123!",
		"SecurePass#2024",
		"MyP@ssw0rd",
		"Test@1234",
	}
	for i := 0; i < 100; i++ {
		for _, p := range passwords {
			validateStrongPasswordOld(p)
		}
	}
}

func TestBenchmarkNewWrapper(t *testing.T) {
	passwords := []string{
		"Abc123!@",
		"Password123!",
		"SecurePass#2024",
		"MyP@ssw0rd",
		"Test@1234",
	}
	for i := 0; i < 100; i++ {
		for _, p := range passwords {
			validateStrongPasswordFast(p)
		}
	}
}
