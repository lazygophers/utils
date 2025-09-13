package cryptox

import (
	"testing"
)

func TestCRC32(t *testing.T) {
	input := "test"
	expected := uint32(3632233996)
	result := CRC32(input)
	if result != expected {
		t.Errorf("CRC32(%s) = %d; want %d", input, result, expected)
	}
}

func TestCRC64(t *testing.T) {
	input := "test"
	expected := uint64(18020588380933092773)
	result := CRC64(input)
	if result != expected {
		t.Errorf("CRC64(%s) = %d; want %d", input, result, expected)
	}
}
