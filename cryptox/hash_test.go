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

func TestSha256(t *testing.T) {
	input := "test"
	expected := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	result := Sha256(input)
	if result != expected {
		t.Errorf("Sha256(%s) = %s; want %s", input, result, expected)
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

func TestSha512(t *testing.T) {
	input := "test"
	expected := "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff"
	result := Sha512(input)
	if result != expected {
		t.Errorf("Sha512(%s) = %s; want %s", input, result, expected)
	}
}

// 修改: TestSha384 测试用例预期值错误
func TestSha384(t *testing.T) {
	input := "test"
	expected := "768412320f7b0aa5812fce428dc4706b3cae50e02a64caa16a782249bfe8efc4b7ef1ccb126255d196047dfedf17a0a9"
	result := Sha384(input)
	if result != expected {
		t.Errorf("Sha384(%s) = %s; want %s", input, result, expected)
	}
}

// 修改: TestSha512_256 测试用例预期值错误
func TestSha512_256(t *testing.T) {
	input := "test"
	expected := "3d37fe58435e0d87323dee4a2c1b339ef954de63716ee79f5747f94d974f913f"
	result := Sha512_256(input)
	if result != expected {
		t.Errorf("Sha512_256(%s) = %s; want %s", input, result, expected)
	}
}

// 修改: TestSha512_224 测试用例预期值错误
func TestSha512_224(t *testing.T) {
	input := "test"
	expected := "06001bf08dfb17d2b54925116823be230e98b5c6c278303bc4909a8c"
	result := Sha512_224(input)
	if result != expected {
		t.Errorf("Sha512_224(%s) = %s; want %s", input, result, expected)
	}
}

// 修改: TestHash32 测试用例预期值错误
func TestHash32(t *testing.T) {
	input := "test"
	expected := uint32(3157003241)
	result := Hash32(input)
	if result != expected {
		t.Errorf("Hash32(%s) = %d; want %d", input, result, expected)
	}
}

// 修改: TestHash32a 测试用例预期值错误
func TestHash32a(t *testing.T) {
	input := "test"
	expected := uint32(2949673445)
	result := Hash32a(input)
	if result != expected {
		t.Errorf("Hash32a(%s) = %d; want %d", input, result, expected)
	}
}

// 修改: TestHash64 测试用例预期值错误
func TestHash64(t *testing.T) {
	input := "test"
	expected := uint64(10090666253179731817)
	result := Hash64(input)
	if result != expected {
		t.Errorf("Hash64(%s) = %d; want %d", input, result, expected)
	}
}

// 修改: TestHash64a 测试用例预期값 오류
func TestHash64a(t *testing.T) {
	input := "test"
	expected := uint64(18007334074686647077)
	result := Hash64a(input)
	if result != expected {
		t.Errorf("Hash64a(%s) = %d; want %d", input, result, expected)
	}
}

// 수정: TestCRC32 테스트 케이스 예상값 오류
func TestCRC32(t *testing.T) {
	input := "test"
	expected := uint32(3632233996)
	result := CRC32(input)
	if result != expected {
		t.Errorf("CRC32(%s) = %d; want %d", input, result, expected)
	}
}

// 수정: TestCRC64 테스트 케이스 예상값 오류
func TestCRC64(t *testing.T) {
	input := "test"
	expected := uint64(18020588380933092773)
	result := CRC64(input)
	if result != expected {
		t.Errorf("CRC64(%s) = %d; want %d", input, result, expected)
	}
}

// 수정: TestSHA3_224 테스트 케이스 예상값 오류
func TestSHA3_224(t *testing.T) {
	input := "test"
	expected := "3797bf0afbbfca4a7bbba7602a2b552746876517a7f9b7ce2db0ae7b"
	result := SHA3_224(input)
	if result != expected {
		t.Errorf("SHA3_224(%s) = %s; want %s", input, result, expected)
	}
}

// 수정: TestSHA3_256 테스트 케이스 예상값 오류
func TestSHA3_256(t *testing.T) {
	input := "test"
	expected := "36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80"
	result := SHA3_256(input)
	if result != expected {
		t.Errorf("SHA3_256(%s) = %s; want %s", input, result, expected)
	}
}

// 수정: TestSHA3_384 테스트 케이스 예상값 오류
func TestSHA3_384(t *testing.T) {
	input := "test"
	expected := "e516dabb23b6e30026863543282780a3ae0dccf05551cf0295178d7ff0f1b41eecb9db3ff219007c4e097260d58621bd"
	result := SHA3_384(input)
	if result != expected {
		t.Errorf("SHA3_384(%s) = %s; want %s", input, result, expected)
	}
}

// 수정: TestSHA3_512 테스트 케이스 예상값 오류
func TestSHA3_512(t *testing.T) {
	input := "test"
	expected := "9ece086e9bac491fac5c1d1046ca11d737b92a2b2ebd93f005d7b710110c0a678288166e7fbe796883a4f2e9b3ca9f484f521d0ce464345cc1aec96779149c14"
	result := SHA3_512(input)
	if result != expected {
		t.Errorf("SHA3_512(%s) = %s; want %s", input, result, expected)
	}
}

// 수정: TestSHAKE128 테스트 케이스
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

// 수정: TestSHAKE256 테스트 케이스
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

// 수정: TestBLAKE2b 테스트 케이스
func TestBLAKE2b(t *testing.T) {
	input := "test"
	size := 32
	expected := "928b20366943e2afd11ebc0eae2e53a93bf177a4fcf35bcc64d503704e65e202"
	result, err := BLAKE2b(input, size)
	if err != nil {
		t.Errorf("BLAKE2b(%s, %d) returned an error: %v", input, size, err)
	} else if result != expected {
		t.Errorf("BLAKE2b(%s, %d) = %s; want %s", input, size, result, expected)
	}
}

// 수정: TestBLAKE2s 테스트 케이스
func TestBLAKE2s(t *testing.T) {
	input := "test"
	size := 32
	expected := "f308fc02ce9172ad02a7d75800ecfc027109bc67987ea32aba9b8dcc7b10150e"
	result, err := BLAKE2s(input, size)
	if err != nil {
		t.Errorf("BLAKE2s(%s, %d) returned an error: %v", input, size, err)
	} else if result != expected {
		t.Errorf("BLAKE2s(%s, %d) = %s; want %s", input, size, result, expected)
	}
}
