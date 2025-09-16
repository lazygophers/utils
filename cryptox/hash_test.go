package cryptox

import (
	"testing"
)

// ==== MERGED FROM hash_basic_test.go ====

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

// ==== MERGED FROM hash_crc_test.go ====

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

// ==== MERGED FROM hash_fnv_test.go ====

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

// ==== MERGED FROM hash_hmac_test.go ====

func TestHMACMd5(t *testing.T) {
	key := "key"
	message := "test"
	expected := "1d4a2743c056e467ff3f09c9af31de7e"
	result := HMACMd5(key, message)
	if result != expected {
		t.Errorf("HMACMd5(%s, %s) = %s; want %s", key, message, result, expected)
	}
}

func TestHMACSHA1(t *testing.T) {
	key := "key"
	message := "test"
	expected := "671f54ce0c540f78ffe1e26dcf9c2a047aea4fda"
	result := HMACSHA1(key, message)
	if result != expected {
		t.Errorf("HMACSHA1(%s, %s) = %s; want %s", key, message, result, expected)
	}
}

func TestHMACSHA256(t *testing.T) {
	key := "key"
	message := "test"
	expected := "02afb56304902c656fcb737cdd03de6205bb6d401da2812efd9b2d36a08af159"
	result := HMACSHA256(key, message)
	if result != expected {
		t.Errorf("HMACSHA256(%s, %s) = %s; want %s", key, message, result, expected)
	}
}

func TestHMACSHA384(t *testing.T) {
	key := "key"
	message := "test"
	expected := "160a099ad9d6dadb46311cb4e6dfe98aca9ca519c2e0fedc8dc45da419b1173039cc131f0b5f68b2bbc2b635109b57a8"
	result := HMACSHA384(key, message)
	if result != expected {
		t.Errorf("HMACSHA384(%s, %s) = %s; want %s", key, message, result, expected)
	}
}

func TestHMACSHA512(t *testing.T) {
	key := "key"
	message := "test"
	expected := "287a0fb89a7fbdfa5b5538636918e537a5b83065e4ff331268b7aaa115dde047a9b0f4fb5b828608fc0b6327f10055f7637b058e9e0dbb9e698901a3e6dd461c"
	result := HMACSHA512(key, message)
	if result != expected {
		t.Errorf("HMACSHA512(%s, %s) = %s; want %s", key, message, result, expected)
	}
}