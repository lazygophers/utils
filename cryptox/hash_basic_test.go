package cryptox

import (
	"testing"
)

func TestMd5(t *testing.T) {
	input := "test"
	expected := "098f6bcd4621d373cade4e832627b4f6"
	result := Md5(input)
	if result != expected {
		t.Errorf("Md5(%s) = %s; want %s", input, result, expected)
	}
}

func TestSHA1(t *testing.T) {
	input := "test"
	expected := "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"
	result := SHA1(input)
	if result != expected {
		t.Errorf("SHA1(%s) = %s; want %s", input, result, expected)
	}
}

func TestSha224(t *testing.T) {
	input := "test"
	expected := "90a3ed9e32b2aaf4c61c410eb925426119e1a9dc53d4286ade99a809"
	result := Sha224(input)
	if result != expected {
		t.Errorf("Sha224(%s) = %s; want %s", input, result, expected)
	}
}

func TestSha256(t *testing.T) {
	input := "test"
	expected := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	result := Sha256(input)
	if result != expected {
		t.Errorf("Sha256(%s) = %s; want %s", input, result, expected)
	}
}

func TestSha384(t *testing.T) {
	input := "test"
	expected := "768412320f7b0aa5812fce428dc4706b3cae50e02a64caa16a782249bfe8efc4b7ef1ccb126255d196047dfedf17a0a9"
	result := Sha384(input)
	if result != expected {
		t.Errorf("Sha384(%s) = %s; want %s", input, result, expected)
	}
}

func TestSha512(t *testing.T) {
	input := "test"
	expected := "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff"
	result := Sha512(input)
	if result != expected {
		t.Errorf("Sha512(%s) = %s; want %s", input, result, expected)
	}
}

func TestSha512_224(t *testing.T) {
	input := "test"
	expected := "06001bf08dfb17d2b54925116823be230e98b5c6c278303bc4909a8c"
	result := Sha512_224(input)
	if result != expected {
		t.Errorf("Sha512_224(%s) = %s; want %s", input, result, expected)
	}
}

func TestSha512_256(t *testing.T) {
	input := "test"
	expected := "3d37fe58435e0d87323dee4a2c1b339ef954de63716ee79f5747f94d974f913f"
	result := Sha512_256(input)
	if result != expected {
		t.Errorf("Sha512_256(%s) = %s; want %s", input, result, expected)
	}
}
