package cryptox

import (
	"testing"
)

func TestSHA3_224(t *testing.T) {
	input := "test"
	expected := "3797bf0afbbfca4a7bbba7602a2b552746876517a7f9b7ce2db0ae7b"
	result := SHA3_224(input)
	if result != expected {
		t.Errorf("SHA3_224(%s) = %s; want %s", input, result, expected)
	}
}

func TestSHA3_256(t *testing.T) {
	input := "test"
	expected := "36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80"
	result := SHA3_256(input)
	if result != expected {
		t.Errorf("SHA3_256(%s) = %s; want %s", input, result, expected)
	}
}

func TestSHA3_384(t *testing.T) {
	input := "test"
	expected := "e516dabb23b6e30026863543282780a3ae0dccf05551cf0295178d7ff0f1b41eecb9db3ff219007c4e097260d58621bd"
	result := SHA3_384(input)
	if result != expected {
		t.Errorf("SHA3_384(%s) = %s; want %s", input, result, expected)
	}
}

func TestSHA3_512(t *testing.T) {
	input := "test"
	expected := "9ece086e9bac491fac5c1d1046ca11d737b92a2b2ebd93f005d7b710110c0a678288166e7fbe796883a4f2e9b3ca9f484f521d0ce464345cc1aec96779149c14"
	result := SHA3_512(input)
	if result != expected {
		t.Errorf("SHA3_512(%s) = %s; want %s", input, result, expected)
	}
}

func TestSHAKE128(t *testing.T) {
	input := "test"
	size := 32
	expected := "d3b0aa9cd8b7255622cebc631e867d4093d6f6010191a53973c45fec9b07c774"
	result, err := SHAKE128(input, size)
	if err != nil {
		t.Errorf("SHAKE128(%s, %d) returned an error: %v", input, size, err)
	} else if result != expected {
		t.Errorf("SHAKE128(%s, %d) = %s; want %s", input, size, result, expected)
	}
}

func TestSHAKE256(t *testing.T) {
	input := "test"
	size := 32
	expected := "b54ff7255705a71ee2925e4a3e30e41aed489a579d5595e0df13e32e1e4dd202"
	result, err := SHAKE256(input, size)
	if err != nil {
		t.Errorf("SHAKE256(%s, %d) returned an error: %v", input, size, err)
	} else if result != expected {
		t.Errorf("SHAKE256(%s, %d) = %s; want %s", input, size, result, expected)
	}
}

func TestKeccak256(t *testing.T) {
	input := "test"
	expected := "9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658"
	result := Keccak256(input)
	if result != expected {
		t.Errorf("Keccak256(%s) = %s; want %s", input, result, expected)
	}
}

// Test error conditions for SHA3/SHAKE functions
func TestSHA3ErrorConditions(t *testing.T) {
	data := "test"

	// Test SHAKE128 with invalid size
	_, err := SHAKE128(data, 0)
	if err == nil {
		t.Error("Expected error for SHAKE128 with size 0")
	}

	_, err = SHAKE128(data, -1)
	if err == nil {
		t.Error("Expected error for SHAKE128 with negative size")
	}

	// Test SHAKE256 with invalid size
	_, err = SHAKE256(data, 0)
	if err == nil {
		t.Error("Expected error for SHAKE256 with size 0")
	}

	_, err = SHAKE256(data, -1)
	if err == nil {
		t.Error("Expected error for SHAKE256 with negative size")
	}
}
