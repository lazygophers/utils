---
title: fake - 测试数据生成
---

# fake - 测试数据生成

## 概述

fake 模块提供全面的测试数据生成功能,用于单元测试和集成测试。它支持多种语言和数据类型。

## 类型

### Language

用于本地化的语言类型。

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

用于姓名生成的性别类型。

```go
type Gender string

const (
    GenderMale   Gender = "male"
    GenderFemale Gender = "female"
)
```

---

### Country

用于地址生成的国家类型。

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

测试数据生成器。

```go
type Faker struct {
    language Language
    country  Country
    gender   Gender
    seed     int64
}
```

---

## 构造函数

### New()

使用选项创建新的 faker。

```go
func New(opts ...FakerOption) *Faker
```

**选项:**
- `WithLanguage(lang Language)` - 设置语言
- `WithCountry(country Country)` - 设置国家
- `WithGender(gender Gender)` - 设置性别
- `WithSeed(seed int64)` - 设置随机种子

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

## 默认实例

### SetDefaultLanguage()

为全局实例设置默认语言。

```go
func SetDefaultLanguage(lang Language)
```

---

### SetDefaultCountry()

为全局实例设置默认国家。

```go
func SetDefaultCountry(country Country)
```

---

### SetDefaultGender()

为全局实例设置默认性别。

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

### 数据生成

```go
faker := fake.New(fake.WithLanguage(fake.LanguageChineseSimplified))

// 生成用户数据
user := struct {
    Name:    faker.Name(),
    Email:   faker.Email(),
    Phone:   faker.Phone(),
    Address: faker.Address(),
    Company: faker.Company(),
}

// 生成多个用户
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

### 随机数据

```go
faker := fake.New()

// 随机数字
age := faker.Int(18, 65)
price := faker.Float(10.0, 1000.0)
quantity := faker.Int(1, 100)

// 随机布尔值
active := faker.Bool()
verified := faker.Bool()

// 随机日期
date := faker.Date()
time := faker.Time()
```

---

## 最佳实践

### 种子随机

```go
// 好的做法: 使用种子进行可重现的测试
faker := fake.New(fake.WithSeed(12345))
result1 := faker.Name()
result2 := faker.Name()
// result1 == result2
```

### 本地化

```go
// 好的做法: 使用适当的语言
faker := fake.New(fake.WithLanguage(fake.LanguageChineseSimplified))
name := faker.Name()  // 中文名

faker := fake.New(fake.WithLanguage(fake.LanguageEnglish))
name := faker.Name()  // 英文名
```

---

## 相关文档

- [randx](/zh-CN/modules/randx) - 随机工具
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
