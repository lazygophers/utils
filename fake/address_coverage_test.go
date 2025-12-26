package fake

import (
	"testing"
)

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
