package fake

import (
	"testing"
)

// TestStreet 测试街道地址生成
func TestStreet(t *testing.T) {
	faker := New()
	street := faker.Street()

	if street == "" {
		t.Error("Street() returned empty string")
	}

	// 街道地址应该包含数字
	hasNumber := false
	for _, char := range street {
		if char >= '0' && char <= '9' {
			hasNumber = true
			break
		}
	}
	if !hasNumber {
		t.Error("Street address should contain numbers")
	}
}

// TestCity 测试城市生成
func TestCity(t *testing.T) {
	faker := New()
	city := faker.City()

	if city == "" {
		t.Error("City() returned empty string")
	}
}

// TestState 测试州/省份生成
func TestState(t *testing.T) {
	// 测试美国州
	usFaker := New(WithCountry(CountryUS))
	usState := usFaker.State()
	if usState == "" {
		t.Error("US State() returned empty string")
	}

	// 测试中国省份
	cnFaker := New(WithCountry(CountryChina))
	cnProvince := cnFaker.State()
	if cnProvince == "" {
		t.Error("China State() returned empty string")
	}
}

// TestZipCode 测试邮政编码生成
func TestZipCode(t *testing.T) {
	// 测试美国邮编
	usFaker := New(WithCountry(CountryUS))
	usZip := usFaker.ZipCode()
	if len(usZip) != 5 {
		t.Errorf("US zip code should be 5 digits, got %s", usZip)
	}

	// 测试中国邮编
	cnFaker := New(WithCountry(CountryChina))
	cnZip := cnFaker.ZipCode()
	if len(cnZip) != 6 {
		t.Errorf("China zip code should be 6 digits, got %s", cnZip)
	}
}

// TestFullAddress 测试完整地址生成
func TestFullAddress(t *testing.T) {
	faker := New()
	address := faker.FullAddress()

	if address == nil {
		t.Fatal("FullAddress() returned nil")
	}

	if address.Street == "" {
		t.Error("Address street should not be empty")
	}

	if address.City == "" {
		t.Error("Address city should not be empty")
	}

	if address.ZipCode == "" {
		t.Error("Address zip code should not be empty")
	}

	if address.Country == "" {
		t.Error("Address country should not be empty")
	}

	if address.FullAddress == "" {
		t.Error("Address full address should not be empty")
	}
}

// TestGlobalAddressFunctions 测试全局地址函数
func TestGlobalAddressFunctions(t *testing.T) {
	street := Street()
	if street == "" {
		t.Error("Global Street() returned empty string")
	}

	city := City()
	if city == "" {
		t.Error("Global City() returned empty string")
	}

	zipCode := ZipCode()
	if zipCode == "" {
		t.Error("Global ZipCode() returned empty string")
	}

	countryName := CountryName()
	if countryName == "" {
		t.Error("Global CountryName() returned empty string")
	}

	fullAddress := FullAddress()
	if fullAddress == nil {
		t.Error("Global FullAddress() returned nil")
	}

	addressLine := AddressLine()
	if addressLine == "" {
		t.Error("Global AddressLine() returned empty string")
	}
}
