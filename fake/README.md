# Fake - 多语言假数据生成器

`fake` 包是一个高性能、多语言支持的假数据生成器，为 Go 应用提供丰富的测试数据生成功能。

## 特性

- 🌍 **多语言支持**: 支持英文、简体中文、繁体中文、法语、俄语、葡萄牙语、西班牙语
- 🏗️ **模块化设计**: 每个功能模块独立，支持按需加载
- ⚡ **高性能**: 使用对象池和缓存优化，支持并发安全
- 🎯 **精确控制**: 支持性别、国家、语言等多维度配置
- 📁 **嵌入式数据**: 使用 embed 技术，减少外部依赖
- 🏷️ **构建标签**: 支持通过 build tag 选择特定语言支持
- 🧪 **100% 测试覆盖**: 完整的测试套件确保代码质量

## 安装

```bash
go get github.com/lazygophers/utils/fake
```

## 快速开始

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/fake"
)

func main() {
    // 使用默认配置
    fmt.Println("Name:", fake.Name())
    fmt.Println("Email:", fake.Email())
    fmt.Println("Phone:", fake.PhoneNumber())
    fmt.Println("Address:", fake.AddressLine())
    fmt.Println("Company:", fake.CompanyName())
    fmt.Println("User Agent:", fake.RandomUserAgent())
}
```

### 创建自定义 Faker

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/fake"
)

func main() {
    // 创建中文环境的 Faker
    faker := fake.New(
        fake.WithLanguage(fake.LanguageChineseSimplified),
        fake.WithCountry(fake.CountryChina),
        fake.WithGender(fake.GenderMale),
    )

    fmt.Println("中文姓名:", faker.Name())
    fmt.Println("手机号码:", faker.PhoneNumber())
    fmt.Println("详细地址:", faker.FullAddress().FullAddress)
    fmt.Println("公司名称:", faker.CompanyName())
}
```

## 功能模块

### 姓名生成

```go
faker := fake.New()

// 基本姓名
fmt.Println(faker.Name())           // "John Smith"
fmt.Println(faker.FirstName())      // "John"
fmt.Println(faker.LastName())       // "Smith"

// 带前缀后缀的完整姓名
fmt.Println(faker.FormattedName())  // "Dr. John Smith Jr."
fmt.Println(faker.NamePrefix())     // "Dr."
fmt.Println(faker.NameSuffix())     // "Jr."

// 批量生成
names := faker.BatchNames(10)
```

### 地址生成

```go
faker := fake.New()

// 基本地址组件
fmt.Println(faker.Street())         // "123 Main Street"
fmt.Println(faker.City())           // "New York"
fmt.Println(faker.State())          // "New York"
fmt.Println(faker.ZipCode())        // "10001"
fmt.Println(faker.CountryName())    // "United States"

// 完整地址
address := faker.FullAddress()
fmt.Printf("%+v\\n", address)

// 坐标
lat, lng := faker.Coordinate()
fmt.Printf("坐标: %.6f, %.6f\\n", lat, lng)
```

### 联系信息

```go
faker := fake.New()

// 电话号码
fmt.Println(faker.PhoneNumber())    // "+1 (555) 123-4567"
fmt.Println(faker.MobileNumber())   // "+1 (555) 987-6543"

// 邮箱地址
fmt.Println(faker.Email())          // "john.doe@gmail.com"
fmt.Println(faker.CompanyEmail())   // "john.doe@company.com"
fmt.Println(faker.SafeEmail())      // "john.doe@example.com"

// 网络相关
fmt.Println(faker.URL())            // "https://example.com/path"
fmt.Println(faker.IPv4())           // "192.168.1.1"
fmt.Println(faker.IPv6())           // "2001:db8::1"
fmt.Println(faker.MAC())            // "00:1a:2b:3c:4d:5e"
```

### 公司信息

```go
faker := fake.New()

// 基本公司信息
fmt.Println(faker.CompanyName())    // "Tech Solutions Inc."
fmt.Println(faker.CompanySuffix())  // "Inc."
fmt.Println(faker.Industry())       // "Technology"
fmt.Println(faker.JobTitle())       // "Senior Software Engineer"
fmt.Println(faker.Department())     // "Engineering"

// 完整公司信息
company := faker.CompanyInfo()
fmt.Printf("%+v\\n", company)

// 商业术语
fmt.Println(faker.BS())             // "implement strategic solutions"
fmt.Println(faker.Catchphrase())    // "Innovation in Technology, Excellence in Service"
```

### 设备信息

```go
faker := fake.New()

// 用户代理
fmt.Println(faker.UserAgent())      // 从内置列表随机选择
fmt.Println(faker.MobileUserAgent()) // 移动端用户代理
fmt.Println(faker.DesktopUserAgent()) // 桌面端用户代理

// 设备信息
device := faker.DeviceInfo()
fmt.Printf("%+v\\n", device)

// 操作系统和浏览器
fmt.Println(faker.OS())             // "Windows"
fmt.Println(faker.Browser())        // "Chrome"
```

### 身份证件和金融

```go
faker := fake.New()

// 身份证件
fmt.Println(faker.SSN())            // "123-45-6789"
fmt.Println(faker.ChineseIDNumber()) // "110101199001010123"
fmt.Println(faker.Passport())       // "123456789"
fmt.Println(faker.DriversLicense())  // "D1234567890"

// 完整身份证件信息
identity := faker.IdentityDoc()
fmt.Printf("%+v\\n", identity)

// 银行卡信息
fmt.Println(faker.CreditCardNumber()) // "4111111111111111"
fmt.Println(faker.CVV())              // "123"
fmt.Println(faker.BankAccount())      // "1234567890123456"
fmt.Println(faker.IBAN())             // "DE89370400440532013000"

// 完整信用卡信息
card := faker.CreditCardInfo()
fmt.Printf("%+v\\n", card)

// 安全的测试卡号
fmt.Println(faker.SafeCreditCardNumber()) // 测试专用卡号
```

### 文本内容

```go
faker := fake.New()

// 基本文本
fmt.Println(faker.Word())           // "lorem"
fmt.Println(faker.Sentence())       // "Lorem ipsum dolor sit amet."
fmt.Println(faker.Paragraph())      // 一段文本

// 批量文本
words := faker.Words(5)             // 5个单词
sentences := faker.Sentences(3)     // 3个句子
paragraphs := faker.Paragraphs(2)   // 2个段落

// 指定长度文本
text := faker.Text(200)             // 200字符文本

// 特殊文本格式
fmt.Println(faker.Title())          // "Lorem Ipsum Dolor"
fmt.Println(faker.Quote())          // "\"Lorem ipsum dolor sit amet\""
fmt.Println(faker.Slug())           // "lorem-ipsum-dolor"
fmt.Println(faker.HashTag())        // "#lorem"

// 社交媒体内容
fmt.Println(faker.Tweet())          // 推特格式消息
fmt.Println(faker.Review())         // 评论内容
fmt.Println(faker.Article())        // 完整文章

// Lorem Ipsum
fmt.Println(faker.Lorem())          // 经典 Lorem Ipsum
fmt.Println(faker.LoremWords(10))   // 10个 Lorem 单词
fmt.Println(faker.LoremSentences(3)) // 3个 Lorem 句子
fmt.Println(faker.LoremParagraphs(2)) // 2个 Lorem 段落
```

## 多语言支持

### 支持的语言

- `fake.LanguageEnglish` - 英语
- `fake.LanguageChineseSimplified` - 简体中文
- `fake.LanguageChineseTraditional` - 繁体中文
- `fake.LanguageFrench` - 法语
- `fake.LanguageRussian` - 俄语
- `fake.LanguagePortuguese` - 葡萄牙语
- `fake.LanguageSpanish` - 西班牙语

### 语言切换示例

```go
// 英语环境
enFaker := fake.New(fake.WithLanguage(fake.LanguageEnglish))
fmt.Println("EN:", enFaker.Name()) // "John Smith"

// 简体中文环境
cnFaker := fake.New(fake.WithLanguage(fake.LanguageChineseSimplified))
fmt.Println("CN:", cnFaker.Name()) // "王伟"

// 法语环境
frFaker := fake.New(fake.WithLanguage(fake.LanguageFrench))
fmt.Println("FR:", frFaker.Name()) // "Pierre Dupont"
```

## 高性能特性

### 对象池化

```go
// 使用对象池减少内存分配
fake.WithPooledFaker(func(faker *fake.Faker) {
    name := faker.Name()
    email := faker.Email()
    // ... 其他操作
})

// 预热对象池
fake.WarmupPools()

// 获取池统计信息
stats := fake.GetPoolStats()
fmt.Printf("Pool stats: %+v\\n", stats)
```

### 并行生成

```go
// 并行生成大量数据
names := fake.ParallelGenerate(1000, func(faker *fake.Faker) string {
    return faker.Name()
})

emails := fake.ParallelGenerate(1000, func(faker *fake.Faker) string {
    return faker.Email()
})
```

### 批量优化

```go
faker := fake.New()

// 优化的批量生成
names := faker.BatchNamesOptimized(1000)
emails := faker.BatchEmailsOptimized(1000)

// 普通批量生成
addresses := faker.BatchFullAddresses(100)
companies := faker.BatchCompanyInfos(50)
```

## 上下文支持

```go
import "context"

ctx := context.Background()
ctx = fake.ContextWithLanguage(ctx, fake.LanguageChineseSimplified)
ctx = fake.ContextWithCountry(ctx, fake.CountryChina)
ctx = fake.ContextWithGender(ctx, fake.GenderMale)

faker := fake.WithContext(ctx)
fmt.Println(faker.Name()) // 中文男性姓名
```

## 构建标签

通过构建标签可以减少最终二进制文件大小，只包含需要的语言支持：

```bash
# 只包含英语支持
go build -tags fake_en

# 只包含简体中文支持  
go build -tags fake_zh_cn

# 包含多种常用语言
go build -tags fake_multi

# 包含所有语言（默认）
go build
```

## 统计和监控

```go
faker := fake.New()

// 生成一些数据
faker.Name()
faker.Email()
faker.PhoneNumber()

// 获取统计信息
stats := faker.Stats()
fmt.Printf("调用次数: %d\\n", stats["call_count"])
fmt.Printf("缓存命中: %d\\n", stats["cache_hits"])
fmt.Printf("生成数据: %d\\n", stats["generated_data"])

// 清理缓存
faker.ClearCache()
```

## 性能测试

```go
func BenchmarkNameGeneration(b *testing.B) {
    faker := fake.New()
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        _ = faker.Name()
    }
}

func BenchmarkBatchGeneration(b *testing.B) {
    faker := fake.New()
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        _ = faker.BatchNamesOptimized(100)
    }
}
```

## 最佳实践

### 1. 选择合适的生成方式

```go
// 单个数据生成
name := fake.Name()

// 少量批量生成
names := fake.BatchGenerate(10, fake.Name)

// 大量数据并行生成
names := fake.ParallelGenerate(10000, func(faker *fake.Faker) string {
    return faker.Name()
})
```

### 2. 合理使用缓存

```go
faker := fake.New()

// 长时间运行的程序应定期清理缓存
ticker := time.NewTicker(time.Hour)
go func() {
    for range ticker.C {
        faker.ClearCache()
    }
}()
```

### 3. 并发安全

```go
// 每个 goroutine 使用独立的 Faker 实例
func generateData() {
    faker := fake.New() // 或使用 fake.GetFaker()
    defer fake.PutFaker(faker) // 如果使用了 GetFaker()
    
    // ... 生成数据
}
```

### 4. 内存管理

```go
// 对于大量数据生成，使用流式处理
func generateLargeDataset(count int, output chan<- string) {
    defer close(output)
    
    fake.WithPooledFaker(func(faker *fake.Faker) {
        for i := 0; i < count; i++ {
            output <- faker.Name()
            
            // 定期清理缓存
            if i%1000 == 0 {
                faker.ClearCache()
            }
        }
    })
}
```

## 错误处理

fake 包设计为尽可能不出错，在数据加载失败时会自动回退到默认值：

```go
// 即使数据文件缺失，也会返回合理的默认值
faker := fake.New(fake.WithLanguage(fake.LanguageRussian))
name := faker.Name() // 如果俄语数据不存在，会回退到英语
```

## 扩展和自定义

### 添加自定义数据

你可以通过修改 `data/` 目录下的 JSON 文件来添加自定义数据：

```json
{
  "language": "en",
  "type": "names",
  "version": "1.0.0",
  "items": [
    {"value": "CustomName", "weight": 1.0, "tags": ["custom"]}
  ]
}
```

### 实现自定义生成器

```go
type CustomFaker struct {
    *fake.Faker
}

func (cf *CustomFaker) CustomData() string {
    // 实现自定义逻辑
    return "custom data"
}

func NewCustomFaker() *CustomFaker {
    return &CustomFaker{
        Faker: fake.New(),
    }
}
```

## 许可证

本项目采用 MIT 许可证。详见 [LICENSE](LICENSE) 文件。

## 贡献

欢迎贡献代码！请查看 [CONTRIBUTING.md](CONTRIBUTING.md) 了解详情。

## 支持

如有问题或建议，请提交 [Issue](https://github.com/lazygophers/utils/issues)。