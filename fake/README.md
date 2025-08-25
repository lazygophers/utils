# fake

fake 是一个功能强大的模拟数据生成工具包，专门用于生成各种类型的测试数据、模拟数据样本和虚假数据。它提供了丰富的数据类型支持、灵活的配置选项和高效的生成策略，是开发测试、原型开发和演示场景的理想选择。

## 特性

- **多种数据类型支持**
  - 基本数据类型：字符串、数字、布尔值、日期时间
  - 个人信息：姓名、地址、电话、邮箱、身份证号
  - 网络数据：IP地址、MAC地址、URL、User-Agent
  - 金融数据：信用卡号、银行账号、交易金额
  - 文本数据：段落、句子、单词、字符
  - 商务数据：公司名称、职位、产品名称

- **智能数据生成**
  - 基于规则的智能生成
  - 数据关联性和一致性保证
  - 支持自定义生成函数
  - 支持多语言和地区化数据

- **高性能和可扩展性**
  - 零分配设计，减少GC压力
  - 支持批量生成
  - 线程安全设计
  - 内存池优化

- **灵活的配置选项**
  - 可配置的数据生成器
  - 支持种子值保证可重现性
  - 支持自定义数据源
  - 支持数据过滤和验证

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
    // 生成随机姓名
    name := fake.Name()
    fmt.Println("姓名:", name)
    
    // 生成随机邮箱
    email := fake.Email()
    fmt.Println("邮箱:", email)
    
    // 生成随机手机号
    phone := fake.Phone()
    fmt.Println("手机号:", phone)
    
    // 生成随机地址
    address := fake.Address()
    fmt.Println("地址:", address)
}
```

### 生成特定类型数据

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/fake"
)

func main() {
    // 数字类型
    fmt.Println("随机整数:", fake.Int())
    fmt.Println("范围内的随机数:", fake.IntRange(1, 100))
    fmt.Println("随机浮点数:", fake.Float64())
    
    // 字符串类型
    fmt.Println("随机字符串:", fake.String(10))
    fmt.Println("随机单词:", fake.Word())
    fmt.Println("随机句子:", fake.Sentence())
    fmt.Println("随机段落:", fake.Paragraph())
    
    // 时间类型
    fmt.Println("随机时间:", fake.Time())
    fmt.Println("随机日期:", fake.Date())
    fmt.Println("未来时间:", fake.FutureTime())
    fmt.Println("过去时间:", fake.PastTime())
    
    // 网络类型
    fmt.Println("随机IP:", fake.IP())
    fmt.Println("随机URL:", fake.URL())
    fmt.Println("随机User-Agent:", fake.UserAgent())
}
```

### 使用种子值保证可重现性

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/fake"
)

func main() {
    // 设置种子值
    fake.Seed(12345)
    
    // 生成的数据将是可重现的
    fmt.Println("可重现的姓名:", fake.Name())
    fmt.Println("可重现的邮箱:", fake.Email())
}
```

### 批量生成数据

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/fake"
)

func main() {
    // 批量生成姓名列表
    names := fake.Names(10)
    fmt.Println("10个随机姓名:", names)
    
    // 批量生成用户数据
    type User struct {
        ID    int
        Name  string
        Email string
        Age   int
    }
    
    var users []User
    for i := 0; i < 5; i++ {
        users = append(users, User{
            ID:    i + 1,
            Name:  fake.Name(),
            Email: fake.Email(),
            Age:   fake.IntRange(18, 60),
        })
    }
    
    fmt.Printf("生成的用户数据: %+v\n", users)
}
```

### 自定义数据生成器

```go
package main

import (
    "fmt"
    "math/rand"
    "github.com/lazygophers/utils/fake"
)

func main() {
    // 创建自定义生成器
    customFake := fake.New()
    
    // 添加自定义生成函数
    customFake.AddFunc("custom_id", func() string {
        return fmt.Sprintf("ID-%d", rand.Intn(10000))
    })
    
    // 使用自定义生成器
    fmt.Println("自定义ID:", customFake.Get("custom_id"))
    
    // 或者使用自定义生成器的方法
    customFake.Register("product_code", func() string {
        letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        code := make([]byte, 8)
        for i := range code {
            code[i] = letters[rand.Intn(len(letters))]
        }
        return string(code)
    })
    
    fmt.Println("产品代码:", customFake.Get("product_code"))
}
```

### 结构化数据生成

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/fake"
)

type Person struct {
    Name    string `fake:"name"`
    Email   string `fake:"email"`
    Phone   string `fake:"phone"`
    Address string `fake:"address"`
    Age     int    `fake:"int_range(18,80)"`
}

func main() {
    // 使用结构体标签自动生成数据
    var person Person
    if err := fake.Fill(&person); err != nil {
        panic(err)
    }
    
    fmt.Printf("生成的个人信息: %+v\n", person)
}
```

## 高级功能

### 数据模板

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/fake"
)

func main() {
    // 使用模板生成格式化数据
    template := `{
        "username": "{{.Username}}",
        "email": "{{.Email}}",
        "profile": {
            "age": {{.Age}},
            "city": "{{.City}}"
        }
    }`
    
    data := map[string]interface{}{
        "Username": fake.Username(),
        "Email":    fake.Email(),
        "Age":      fake.IntRange(18, 60),
        "City":     fake.City(),
    }
    
    result, err := fake.ExecuteTemplate(template, data)
    if err != nil {
        panic(err)
    }
    
    fmt.Println("模板生成结果:", result)
}
```

### 数据源配置

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/fake"
)

func main() {
    // 配置自定义数据源
    fake.SetDataSource("names", []string{
        "张三", "李四", "王五", "赵六", "钱七",
    })
    
    // 从自定义数据源获取数据
    name := fake.GetFromSource("names")
    fmt.Println("从自定义数据源获取的姓名:", name)
}
```

### 国际化支持

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/fake"
)

func main() {
    // 设置语言环境
    fake.SetLocale("zh_CN")
    
    // 生成中文数据
    fmt.Println("中文姓名:", fake.Name())
    fmt.Println("中文地址:", fake.Address())
    
    // 切换到英文
    fake.SetLocale("en_US")
    fmt.Println("English Name:", fake.Name())
    fmt.Println("English Address:", fake.Address())
}
```

## 性能优化

### 使用内存池

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/fake"
)

func main() {
    // 启用内存池优化
    fake.EnablePool(true)
    
    // 批量生成数据将获得更好的性能
    names := fake.Names(1000)
    fmt.Printf("生成了 %d 个姓名\n", len(names))
}
```

### 并发生成

```go
package main

import (
    "fmt"
    "sync"
    "github.com/lazygophers/utils/fake"
)

func main() {
    var wg sync.WaitGroup
    results := make(chan string, 10)
    
    // 并发生成数据
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            results <- fake.Name()
        }()
    }
    
    // 收集结果
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // 输出结果
    for name := range results {
        fmt.Println("并发生成的姓名:", name)
    }
}
```

## API 文档

详细的 API 文档请参考 [GoDoc](https://pkg.go.dev/github.com/lazygophers/utils/fake)。

## 许可证

本项目采用 AGPL-3.0 许可证。详见 [LICENSE](../LICENSE) 文件。