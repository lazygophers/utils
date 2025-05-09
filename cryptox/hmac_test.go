package cryptox

import (
	"testing"
)

func TestHmacMd5(t *testing.T) {
	key := "testKey"
	data := "testData"
	result := HmacMd5(key, data)
	if result == "" {
		t.Errorf("HmacMd5 returned empty string")
	}
}

func TestHmacSha1(t *testing.T) {
	key := "testKey"
	data := "testData"
	result := HmacSha1(key, data)
	if result == "" {
		t.Errorf("HmacSha1 returned empty string")
	}
}

func TestHmacSha256(t *testing.T) {
	key := "testKey"
	data := "testData"
	result := HmacSha256(key, data)
	if result == "" {
		t.Errorf("HmacSha256 returned empty string")
	}
}

func TestHmacSha224(t *testing.T) {
	key := "testKey"
	data := "testData"
	result := HmacSha224(key, data)
	if result == "" {
		t.Errorf("HmacSha224 returned empty string")
	}
}

func TestHmacSha512(t *testing.T) {
	key := "testKey"
	data := "testData"
	result := HmacSha512(key, data)
	if result == "" {
		t.Errorf("HmacSha512 returned empty string")
	}
}

func TestHmacSha384(t *testing.T) {
	key := "testKey"
	data := "testData"
	result := HmacSha384(key, data)
	if result == "" {
		t.Errorf("HmacSha384 returned empty string")
	}
}

func TestHmacSha3_256(t *testing.T) {
	key := "testKey"
	data := "testData"
	result := HmacSha3_256(key, data)
	if result == "" {
		t.Errorf("HmacSha3_256 returned empty string")
	}
}

func TestHmacSha3_384(t *testing.T) {
	key := "testKey"
	data := "testData"
	result := HmacSha3_384(key, data)
	if result == "" {
		t.Errorf("HmacSha3_384 returned empty string")
	}
}

func TestHmacSha3_512(t *testing.T) {
	key := "testKey"
	data := "testData"
	result := HmacSha3_512(key, data)
	if result == "" {
		t.Errorf("HmacSha3_512 returned empty string")
	}
}