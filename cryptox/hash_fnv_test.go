package cryptox

import (
	"testing"
)

func TestHash32(t *testing.T) {
	input := "test"
	expected := uint32(3157003241)
	result := Hash32(input)
	if result != expected {
		t.Errorf("Hash32(%s) = %d; want %d", input, result, expected)
	}
}

func TestHash32a(t *testing.T) {
	input := "test"
	expected := uint32(2949673445)
	result := Hash32a(input)
	if result != expected {
		t.Errorf("Hash32a(%s) = %d; want %d", input, result, expected)
	}
}

func TestHash64(t *testing.T) {
	input := "test"
	expected := uint64(10090666253179731817)
	result := Hash64(input)
	if result != expected {
		t.Errorf("Hash64(%s) = %d; want %d", input, result, expected)
	}
}

func TestHash64a(t *testing.T) {
	input := "test"
	expected := uint64(18007334074686647077)
	result := Hash64a(input)
	if result != expected {
		t.Errorf("Hash64a(%s) = %d; want %d", input, result, expected)
	}
}