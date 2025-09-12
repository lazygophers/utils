package cryptox

import (
	"testing"
)

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