---
title: fake - 測試數據生成
---

# fake - 測試數據生成

## 概述

fake 模組提供全面的測試數據生成功能,用於單元測試和集成測試。它支持多種語言和數據類型。

## 類型

### Language

用於本地化的語言類型。

```go
type Language string

const (
    LanguageEnglish            Language = "en"
    LanguageChineseSimplified  Language = "zh-CN"
    LanguageChineseTraditional Language = "zh-TW"
    LanguageFrench             Language = "fr"
    LanguageRussian            Language = "ru"
    LanguagePortuguese         Language = "pt"
    LanguageSpanish            Language = "es"
)
```

---

### Gender

用於姓名生成的性別類型。

```go
type Gender string

const (
    GenderMale   Gender = "male"
    GenderFemale Gender = "female"
)
```

---

### Country

用於地址生成的國家類型。

```go
type Country string

const (
    CountryChina     Country = "CN"
    CountryUS        Country = "US"
    CountryUK        Country = "UK"
    CountryFrance    Country = "FR"
    CountryGermany   Country = "DE"
    CountryJapan     Country = "JP"
    CountryKorea     Country = "KR"
    CountryRussia    Country = "RU"
    CountryBrazil    Country = "BR"
    CountrySpain     Country = "ES"
    CountryPortugal  Country = "PT"
    CountryItaly     Country = "IT"
    CountryCanada    Country = "CA"
    CountryAustralia Country = "AU"
    CountryIndia     Country = "IN"
)
```

---

### Faker

測試數據生成器。

```go
type Faker struct {
    language Language
    country  Country
    gender   Gender
    seed     int64
}
```

---

## 構造函數

### New()

使用選項創建新的 faker。

```go
func New(opts ...FakerOption) *Faker
```

**選項:**
- `WithLanguage(lang Language)` - 設置語言
- `WithCountry(country Country)` - 設置國家
- `WithGender(gender Gender)` - 設置性別
- `WithSeed(seed int64)` - 設置隨機種子

**示例:**
```go
faker := fake.New(
    fake.WithLanguage(fake.LanguageEnglish),
    fake.WithCountry(fake.CountryUS),
    fake.WithGender(fake.GenderMale),
    fake.WithSeed(12345),
)
```

---

## 默認實例

### SetDefaultLanguage()

為全局實例設置默認語言。

```go
func SetDefaultLanguage(lang Language)
```

---

### SetDefaultCountry()

為全局實例設置默認國家。

```go
func SetDefaultCountry(country Country)
```

---

### SetDefaultGender()

為全局實例設置默認性別。

```go
func SetDefaultGender(gender Gender)
```

---

## 使用模式

### 姓名生成

```go
faker := fake.New()
name := faker.Name()
email := faker.Email()
address := faker.Address()
company := faker.Company()
```

### 數據生成

```go
faker := fake.New(fake.WithLanguage(fake.LanguageChineseSimplified))

// 生成用戶數據
user := struct {
    Name:    faker.Name(),
    Email:   faker.Email(),
    Phone:   faker.Phone(),
    Address: faker.Address(),
    Company: faker.Company(),
}

// 生成多個用戶
users := make([]User, 10)
for i := 0; i < 10; i++ {
    users[i] = User{
        Name:    faker.Name(),
        Email:   faker.Email(),
        Phone:   faker.Phone(),
        Address: faker.Address(),
        Company: faker.Company(),
    }
}
```

### 隨機數據

```go
faker := fake.New()

// 隨機數字
age := faker.Int(18, 65)
price := faker.Float(10.0, 1000.0)
quantity := faker.Int(1, 100)

// 隨機布爾值
active := faker.Bool()
verified := faker.Bool()

// 隨機日期
date := faker.Date()
time := faker.Time()
```

---

## 最佳實踐

### 種子隨機

```go
// 好的做法: 使用種子進行可重現的測試
faker := fake.New(fake.WithSeed(12345))
result1 := faker.Name()
result2 := faker.Name()
// result1 == result2
```

### 本地化

```go
// 好的做法: 使用適當的語言
faker := fake.New(fake.WithLanguage(fake.LanguageChineseSimplified))
name := faker.Name()  // 中文名

faker := fake.New(fake.WithLanguage(fake.LanguageEnglish))
name := faker.Name()  // 英文名
```

---

## 相關文檔

- [randx](/zh-TW/modules/randx) - 隨機工具
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
