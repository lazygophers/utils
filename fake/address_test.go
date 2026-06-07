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

// 测试Street函数的各种情况
func TestStreetEdgeCases(t *testing.T) {
	// 创建不同语言的Faker实例
	fakers := []*Faker{
		New(WithLanguage(LanguageEnglish), WithCountry(CountryUS)),
		New(WithLanguage(LanguageChineseSimplified), WithCountry(CountryChina)),
		New(WithLanguage(LanguageChineseTraditional), WithCountry(CountryChina)),
	}

	for _, f := range fakers {
		// 测试Street函数
		street := f.Street()
		if street == "" {
			t.Error("Street() returned empty string")
		}
	}
}

// 测试CountryName函数的各种情况
func TestCountryNameEdgeCases(t *testing.T) {
	// 测试所有支持的语言
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageSpanish,
		LanguageFrench,
		LanguageRussian,
		LanguagePortuguese,
	}

	for _, lang := range languages {
		f := New(WithLanguage(lang), WithCountry(CountryUS))
		country := f.CountryName()
		if country == "" {
			t.Errorf("CountryName() returned empty string for language %s", lang)
		}
	}
}

// 测试FullAddress函数的各种情况
func TestFullAddressEdgeCases(t *testing.T) {
	// langCountryCase 描述语言与国家组合。
	type langCountryCase struct {
		language Language
		country  Country
	}
	// 测试不同国家和语言组合
	tests := []langCountryCase{}

	// 为所有语言和国家组合生成测试用例
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
	}

	countries := []Country{
		CountryUS,
		CountryUK, // UK没有州的概念
		CountryChina,
		CountryCanada,
	}

	for _, lang := range languages {
		for _, country := range countries {
			tests = append(tests, langCountryCase{lang, country})
		}
	}

	for _, tt := range tests {
		f := New(WithLanguage(tt.language), WithCountry(tt.country))
		address := f.FullAddress()
		if address == nil {
			t.Errorf("FullAddress() returned nil for language %s, country %s", tt.language, tt.country)
			continue
		}

		// 验证地址字段
		if address.Street == "" {
			t.Errorf("Address.Street is empty for language %s, country %s", tt.language, tt.country)
		}
		if address.City == "" {
			t.Errorf("Address.City is empty for language %s, country %s", tt.language, tt.country)
		}
		if address.ZipCode == "" {
			t.Errorf("Address.ZipCode is empty for language %s, country %s", tt.language, tt.country)
		}
		if address.Country == "" {
			t.Errorf("Address.Country is empty for language %s, country %s", tt.language, tt.country)
		}
		if address.FullAddress == "" {
			t.Errorf("Address.FullAddress is empty for language %s, country %s", tt.language, tt.country)
		}
	}
}

// 测试AddressLine函数的各种情况
func TestAddressLineEdgeCases(t *testing.T) {
	// 测试不同国家和语言组合
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
	}

	countries := []Country{
		CountryUS,
		CountryUK, // UK没有州的概念
		CountryChina,
		CountryCanada,
	}

	for _, lang := range languages {
		for _, country := range countries {
			f := New(WithLanguage(lang), WithCountry(country))
			addressLine := f.AddressLine()
			if addressLine == "" {
				t.Errorf("AddressLine() returned empty string for language %s, country %s", lang, country)
			}
		}
	}
}

// 测试ZipCode函数的各种情况
func TestZipCodeEdgeCases(t *testing.T) {
	// 测试不同国家的邮政编码格式
	countries := []Country{
		CountryUS,
		CountryChina,
		CountryUK,
		CountryCanada,
	}

	for _, country := range countries {
		f := New(WithLanguage(LanguageEnglish), WithCountry(country))
		zipCode := f.ZipCode()
		if zipCode == "" {
			t.Errorf("ZipCode() returned empty string for country %s", country)
		}
	}
}

// 测试默认街道和城市生成函数
func testDefaultStreetAndCity(t *testing.T) {
	// 测试getDefaultStreet函数
	street := getDefaultStreet()
	if street == "" {
		t.Error("getDefaultStreet() returned empty string")
	}

	// 测试getDefaultCity函数
	city := getDefaultCity()
	if city == "" {
		t.Error("getDefaultCity() returned empty string")
	}
}

// TestAddressCoverage 测试地址相关函数的覆盖率
func TestAddressCoverage(t *testing.T) {
	// 测试加拿大省份生成
	t.Run("test_generate_canadian_province", func(t *testing.T) {
		caFaker := New(WithCountry(CountryCanada))
		province := caFaker.State()
		if province == "" {
			t.Error("Canada State() returned empty string")
		}
	})

	// 测试通用州/省份生成
	t.Run("test_generate_generic_state", func(t *testing.T) {
		// 使用一个不存在的国家，触发generic state生成
		genericFaker := New(WithCountry("UnknownCountry"))
		state := genericFaker.State()
		// generic state可以返回空字符串，所以不检查是否为空
		_ = state
	})

	// 测试传统中文国家名称生成
	t.Run("test_generate_traditional_chinese_country_name", func(t *testing.T) {
		zhTwFaker := New(WithLanguage(LanguageChineseTraditional))
		country := zhTwFaker.CountryName()
		if country == "" {
			t.Error("Traditional Chinese CountryName() returned empty string")
		}
	})

	// 测试纬度生成
	t.Run("test_latitude", func(t *testing.T) {
		faker := New()
		latitude := faker.Latitude()
		// 纬度应该在-90到90之间
		if latitude < -90 || latitude > 90 {
			t.Errorf("Latitude() returned invalid value: %f", latitude)
		}
	})

	// 测试经度生成
	t.Run("test_longitude", func(t *testing.T) {
		faker := New()
		longitude := faker.Longitude()
		// 经度应该在-180到180之间
		if longitude < -180 || longitude > 180 {
			t.Errorf("Longitude() returned invalid value: %f", longitude)
		}
	})

	// 测试坐标生成
	t.Run("test_coordinate", func(t *testing.T) {
		faker := New()
		latitude, longitude := faker.Coordinate()
		// 纬度应该在-90到90之间，经度应该在-180到180之间
		if latitude < -90 || latitude > 90 {
			t.Errorf("Coordinate() returned invalid latitude: %f", latitude)
		}
		if longitude < -180 || longitude > 180 {
			t.Errorf("Coordinate() returned invalid longitude: %f", longitude)
		}
	})

	// 测试批量生成街道
	t.Run("test_batch_streets", func(t *testing.T) {
		faker := New()
		streets := faker.BatchStreets(5)
		if len(streets) != 5 {
			t.Errorf("BatchStreets() returned wrong number of streets: expected 5, got %d", len(streets))
		}
		for _, street := range streets {
			if street == "" {
				t.Error("BatchStreets() returned empty string")
			}
		}
	})

	// 测试批量生成城市
	t.Run("test_batch_cities", func(t *testing.T) {
		faker := New()
		cities := faker.BatchCities(5)
		if len(cities) != 5 {
			t.Errorf("BatchCities() returned wrong number of cities: expected 5, got %d", len(cities))
		}
		for _, city := range cities {
			if city == "" {
				t.Error("BatchCities() returned empty string")
			}
		}
	})

	// 测试批量生成州/省份
	t.Run("test_batch_states", func(t *testing.T) {
		faker := New(WithCountry(CountryUS))
		states := faker.BatchStates(5)
		if len(states) != 5 {
			t.Errorf("BatchStates() returned wrong number of states: expected 5, got %d", len(states))
		}
		for _, state := range states {
			if state == "" {
				t.Error("BatchStates() returned empty string")
			}
		}
	})

	// 测试批量生成邮政编码
	t.Run("test_batch_zip_codes", func(t *testing.T) {
		faker := New(WithCountry(CountryUS))
		zipCodes := faker.BatchZipCodes(5)
		if len(zipCodes) != 5 {
			t.Errorf("BatchZipCodes() returned wrong number of zip codes: expected 5, got %d", len(zipCodes))
		}
		for _, zipCode := range zipCodes {
			if zipCode == "" {
				t.Error("BatchZipCodes() returned empty string")
			}
		}
	})

	// 测试批量生成完整地址
	t.Run("test_batch_full_addresses", func(t *testing.T) {
		faker := New()
		addresses := faker.BatchFullAddresses(5)
		if len(addresses) != 5 {
			t.Errorf("BatchFullAddresses() returned wrong number of addresses: expected 5, got %d", len(addresses))
		}
		for _, address := range addresses {
			if address == nil {
				t.Error("BatchFullAddresses() returned nil")
			}
			if address.Street == "" {
				t.Error("BatchFullAddresses() returned address with empty street")
			}
		}
	})

	// 测试全局纬度函数
	t.Run("test_global_latitude", func(t *testing.T) {
		latitude := Latitude()
		if latitude < -90 || latitude > 90 {
			t.Errorf("Global Latitude() returned invalid value: %f", latitude)
		}
	})

	// 测试全局经度函数
	t.Run("test_global_longitude", func(t *testing.T) {
		longitude := Longitude()
		if longitude < -180 || longitude > 180 {
			t.Errorf("Global Longitude() returned invalid value: %f", longitude)
		}
	})

	// 测试全局坐标函数
	t.Run("test_global_coordinate", func(t *testing.T) {
		latitude, longitude := Coordinate()
		if latitude < -90 || latitude > 90 {
			t.Errorf("Global Coordinate() returned invalid latitude: %f", latitude)
		}
		if longitude < -180 || longitude > 180 {
			t.Errorf("Global Coordinate() returned invalid longitude: %f", longitude)
		}
	})

	// 测试英国邮编生成
	t.Run("test_uk_zip_code", func(t *testing.T) {
		ukFaker := New(WithCountry(CountryUK))
		zipCode := ukFaker.ZipCode()
		if zipCode == "" {
			t.Error("UK ZipCode() returned empty string")
		}
	})

	// 测试加拿大邮编生成
	t.Run("test_canada_zip_code", func(t *testing.T) {
		caFaker := New(WithCountry(CountryCanada))
		zipCode := caFaker.ZipCode()
		if zipCode == "" {
			t.Error("Canada ZipCode() returned empty string")
		}
	})

	// 测试不同语言的国家名称生成
	t.Run("test_different_language_country_names", func(t *testing.T) {
		// 英文国家名称
		enFaker := New(WithLanguage(LanguageEnglish))
		enCountry := enFaker.CountryName()
		if enCountry == "" {
			t.Error("English CountryName() returned empty string")
		}

		// 简体中文国家名称
		zhCnFaker := New(WithLanguage(LanguageChineseSimplified))
		zhCnCountry := zhCnFaker.CountryName()
		if zhCnCountry == "" {
			t.Error("Simplified Chinese CountryName() returned empty string")
		}

		// 传统中文国家名称
		zhTwFaker := New(WithLanguage(LanguageChineseTraditional))
		zhTwCountry := zhTwFaker.CountryName()
		if zhTwCountry == "" {
			t.Error("Traditional Chinese CountryName() returned empty string")
		}
	})

	// 测试地址行生成
	t.Run("test_address_line", func(t *testing.T) {
		faker := New()
		addressLine := faker.AddressLine()
		if addressLine == "" {
			t.Error("AddressLine() returned empty string")
		}

		// 测试中文地址行
		zhFaker := New(WithLanguage(LanguageChineseSimplified))
		zhAddressLine := zhFaker.AddressLine()
		if zhAddressLine == "" {
			t.Error("Chinese AddressLine() returned empty string")
		}
	})

	// 测试默认街道和城市生成
	t.Run("test_default_street_and_city", func(t *testing.T) {
		// 测试getDefaultStreet函数
		street := getDefaultStreet()
		if street == "" {
			t.Error("getDefaultStreet() returned empty string")
		}

		// 测试getDefaultCity函数
		city := getDefaultCity()
		if city == "" {
			t.Error("getDefaultCity() returned empty string")
		}
	})

	// 测试全局State函数
	t.Run("test_global_state", func(t *testing.T) {
		state := State()
		// State()可以返回空字符串，所以不检查是否为空
		_ = state
	})
}
