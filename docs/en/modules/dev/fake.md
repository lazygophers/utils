---
title: fake - Test Data Generation
---

# fake - Test Data Generation

## Overview

The fake module provides comprehensive test data generation for unit testing and integration testing. It supports multiple languages and data types.

## Types

### Language

Language type for localization.

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
)
```

---

### Gender

Gender type for name generation.

```go
type Gender string

const (
    GenderMale   Gender = "male"
    GenderFemale Gender = "female"
)
```

---

### Country

Country type for address generation.

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
)
```

---

### Faker

Test data generator.

```go
type Faker struct {
    language Language
    country  Country
    gender   Gender
    seed     int64
}
```

---

## Constructor

### New()

Create new faker with options.

```go
func New(opts ...FakerOption) *Faker
```

**Options:**
- `WithLanguage(lang Language)` - Set language
- `WithCountry(country Country)` - Set country
- `WithGender(gender Gender)` - Set gender
- `WithSeed(seed int64)` - Set random seed

**Example:**
```go
faker := fake.New(
    fake.WithLanguage(fake.LanguageEnglish),
    fake.WithCountry(fake.CountryUS),
    fake.WithGender(fake.GenderMale),
    fake.WithSeed(12345),
)
)
```

---

## Default Instance

### SetDefaultLanguage()

Set default language for global instance.

```go
func SetDefaultLanguage(lang Language)
```

---

### SetDefaultCountry()

Set default country for global instance.

```go
func SetDefaultCountry(country Country)
```

---

### SetDefaultGender()

Set default gender for global instance.

```go
func SetDefaultGender(gender Gender)
```

---

## Usage Patterns

### Name Generation

```go
faker := fake.New()
name := faker.Name()
email := faker.Email()
address := faker.Address()
company := faker.Company()
```

### Data Generation

```go
faker := fake.New(fake.WithLanguage(fake.LanguageChineseSimplified))

// Generate user data
user := struct {
    Name:    faker.Name(),
    Email:   faker.Email(),
    Phone:   faker.Phone(),
    Address:  faker.Address(),
    Company: faker.Company(),
}

// Generate multiple users
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

### Random Data

```go
faker := fake.New()

// Random numbers
age := faker.Int(18, 65)
price := faker.Float(10.0, 1000.0)
quantity := faker.Int(1, 100)

// Random booleans
active := faker.Bool()
verified := faker.Bool()

// Random dates
date := faker.Date()
time := faker.Time()
```

---

## Best Practices

### Seeded Random

```go
// Good: Use seed for reproducible tests
faker := fake.New(fake.WithSeed(12345))
result1 := faker.Name()
result2 := faker.Name()
// result1 == result2
```

### Localization

```go
// Good: Use appropriate language
faker := fake.New(fake.WithLanguage(fake.LanguageChineseSimplified))
name := faker.Name()  // Chinese name

faker := fake.New(fake.WithLanguage(fake.LanguageEnglish))
name := faker.Name()  // English name
```

---

## Related Documentation

- [randx](/en/modules/randx) - Random utilities
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
