# lazygophers/utils API 参考

一个全面的 Go 工具库，具有模块化设计，为常见开发任务提供基本功能。该库强调类型安全、性能优化，并遵循 Go 1.24+ 标准。

## 安装

```go
go get github.com/lazygophers/utils
```

## 目录

- [核心包 (utils)](#核心包-utils)
- [类型转换 (candy)](#类型转换-candy)
- [字符串操作 (stringx)](#字符串操作-stringx)
- [映射操作 (anyx)](#映射操作-anyx)
- [并发控制 (wait)](#并发控制-wait)
- [熔断器 (hystrix)](#熔断器-hystrix)
- [配置管理 (config)](#配置管理-config)
- [加密操作 (cryptox)](#加密操作-cryptox)
- [时间扩展 (xtime)](#时间扩展-xtime)
- [随机数生成 (randx)](#随机数生成-randx)
- [网络工具 (network)](#网络工具-network)

---

## 核心包 (utils)

根包提供基础工具，包括错误处理、数据库操作和验证。

### Must 函数

错误处理工具，在错误时 panic - 适用于初始化和关键操作。

```go
// Must 结合错误检查和值返回
func Must[T any](value T, err error) T

// MustOk 如果 ok 为 false 则 panic
func MustOk[T any](value T, ok bool) T

// MustSuccess 如果错误不为 nil 则 panic
func MustSuccess(err error)

// Ignore 显式忽略返回值
func Ignore[T any](value T, _ any) T
```

**使用示例：**
```go
import "github.com/lazygophers/utils"

// 初始化时的错误处理
config := utils.Must(loadConfig())

// 转换操作
value := utils.MustOk(m["key"])

// 忽略错误返回值
utils.Ignore(writer.Write(data))
```

### 数据库集成

```go
// Scan 将数据库字段扫描到结构体，支持 JSON 反序列化
func Scan(src interface{}, dst interface{}) error

// Value 将结构体转换为数据库值，支持 JSON 序列化
func Value(m interface{}) (driver.Value, error)
```

**使用示例：**
```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

var user User
err := utils.Scan(dbValue, &user)

// 存储到数据库
value, err := utils.Value(user)
```

### 验证

```go
// Validate 使用 go-playground/validator 验证结构体
func Validate(m interface{}) error
```

---

## 类型转换 (candy)

全面的类型转换和切片操作工具，拥有 99.3% 的测试覆盖率。

### 基本类型转换

```go
// 布尔转换
func ToBool(v interface{}) bool

// 字符串转换
func ToString(v interface{}) string

// 整数转换
func ToInt(v interface{}) int
func ToInt32(v interface{}) int32
func ToInt64(v interface{}) int64

// 浮点数转换
func ToFloat32(v interface{}) float32
func ToFloat64(v interface{}) float64
```

**使用示例：**
```go
import "github.com/lazygophers/utils/candy"

// 智能类型转换
str := candy.ToString(123)        // "123"
num := candy.ToInt("456")         // 456
flag := candy.ToBool("true")      // true
```

### 切片操作

```go
// Filter 过滤切片元素
func Filter[T any](slice []T, predicate func(T) bool) []T

// Map 转换切片元素
func Map[T, R any](slice []T, mapper func(T) R) []R

// Unique 去除重复元素
func Unique[T comparable](slice []T) []T

// Sort 排序切片
func Sort[T any](slice []T, less func(i, j int) bool) []T
```

**使用示例：**
```go
numbers := []int{1, 2, 3, 4, 5}

// 过滤偶数
evens := candy.Filter(numbers, func(n int) bool { return n%2 == 0 })

// 转换为字符串
strings := candy.Map(numbers, func(n int) string { return fmt.Sprintf("num_%d", n) })

// 去重
unique := candy.Unique([]int{1, 2, 2, 3, 3, 4})
```

### 函数式编程

```go
// All 检查所有元素是否满足条件
func All[T any](slice []T, predicate func(T) bool) bool

// Any 检查是否有元素满足条件
func Any[T any](slice []T, predicate func(T) bool) bool

// Reduce 归约操作
func Reduce[T, R any](slice []T, initial R, reducer func(R, T) R) R
```

---

## 字符串操作 (stringx)

高性能字符串操作，拥有 96.4% 的测试覆盖率。

### 零拷贝转换

```go
// ToString 高效字符串转换（零拷贝）
func ToString(b []byte) string

// ToBytes 高效字节转换（零拷贝）
func ToBytes(s string) []byte
```

### 命名转换

```go
// Camel2Snake 驼峰转蛇形命名
func Camel2Snake(s string) string

// Snake2Camel 蛇形转驼峰命名
func Snake2Camel(s string) string

// ToKebab 转换为短横线命名
func ToKebab(s string) string
```

**使用示例：**
```go
import "github.com/lazygophers/utils/stringx"

// 命名转换
snake := stringx.Camel2Snake("UserName")     // "user_name"
camel := stringx.Snake2Camel("user_name")    // "UserName"
kebab := stringx.ToKebab("UserName")         // "user-name"

// 零拷贝转换（高性能）
bytes := stringx.ToBytes("hello")
str := stringx.ToString([]byte("world"))
```

---

## 映射操作 (anyx)

类型无关的映射操作和值提取，拥有 99.0% 的测试覆盖率。

### MapAny 结构

```go
type MapAny struct {
    // 内部实现隐藏
}

// NewMap 创建新的 MapAny
func NewMap(m map[string]interface{}) *MapAny

// NewMapWithJson 从 JSON 创建 MapAny
func NewMapWithJson(jsonStr string) *MapAny
```

### 基本操作

```go
// Get 获取值
func (m *MapAny) Get(key string) interface{}

// Set 设置值
func (m *MapAny) Set(key string, value interface{})

// Exists 检查键是否存在
func (m *MapAny) Exists(key string) bool
```

### 类型安全获取

```go
// 获取特定类型的值
func (m *MapAny) GetString(key string) string
func (m *MapAny) GetInt(key string) int
func (m *MapAny) GetInt64(key string) int64
func (m *MapAny) GetBool(key string) bool
func (m *MapAny) GetFloat64(key string) float64
```

**使用示例：**
```go
import "github.com/lazygophers/utils/anyx"

// 创建映射
data := anyx.NewMap(map[string]interface{}{
    "name": "张三",
    "age":  30,
    "active": true,
})

// 类型安全获取
name := data.GetString("name")     // "张三"
age := data.GetInt("age")          // 30
active := data.GetBool("active")   // true

// 嵌套键支持
data.EnableCut()
nested := data.Get("user.profile.name")
```

---

## 并发控制 (wait)

高级并发和同步工具。

### 异步处理

```go
// Async 异步处理切片元素
func Async[T any](items []T, workerCount int, processor func(T)) error

// AsyncUnique 异步处理切片元素（去重）
func AsyncUnique[T comparable](items []T, workerCount int, processor func(T)) error
```

**使用示例：**
```go
import "github.com/lazygophers/utils/wait"

urls := []string{
    "https://api1.com",
    "https://api2.com",
    "https://api3.com",
}

// 并发处理 URL
err := wait.Async(urls, 3, func(url string) {
    resp, err := http.Get(url)
    if err != nil {
        log.Printf("错误访问 %s: %v", url, err)
        return
    }
    defer resp.Body.Close()
    log.Printf("成功访问 %s", url)
})
```

### 增强 WaitGroup

```go
type Group struct {
    // 增强的 WaitGroup 实现
}

func (g *Group) Go(fn func()) error
func (g *Group) Wait() error
```

---

## 熔断器 (hystrix)

高性能熔断器模式实现，拥有 66.7% 的测试覆盖率。

### 基本熔断器

```go
type CircuitBreaker struct {
    // 高性能熔断器实现
}

func NewCircuitBreaker(config Config) *CircuitBreaker
func (cb *CircuitBreaker) Call(fn func() error) error
```

**使用示例：**
```go
import "github.com/lazygophers/utils/hystrix"

// 创建熔断器
cb := hystrix.NewCircuitBreaker(hystrix.Config{
    MaxRequests:        100,
    Interval:          time.Second * 30,
    Timeout:           time.Second * 10,
    MaxFailureRate:    0.5,
})

// 受保护的调用
err := cb.Call(func() error {
    return riskyOperation()
})
```

---

## 配置管理 (config)

多格式配置加载，拥有 95.7% 的测试覆盖率。

```go
// LoadConfig 加载配置文件
func LoadConfig(filename string, config interface{}) error

// 支持的格式
// - JSON (.json)
// - YAML (.yaml, .yml)
// - TOML (.toml)
// - INI (.ini)
```

**使用示例：**
```go
import "github.com/lazygophers/utils/config"

type AppConfig struct {
    Database struct {
        Host     string `json:"host"`
        Port     int    `json:"port"`
        Username string `json:"username"`
        Password string `json:"password"`
    } `json:"database"`
    
    Server struct {
        Port int    `json:"port"`
        Host string `json:"host"`
    } `json:"server"`
}

var cfg AppConfig
err := config.LoadConfig("app.json", &cfg)
if err != nil {
    log.Fatal(err)
}
```

---

## 加密操作 (cryptox)

全面的加密操作，拥有 100% 的测试覆盖率。

### AES 加密

```go
// AES 加密/解密
func Encrypt(plaintext, key []byte) ([]byte, error)
func Decrypt(ciphertext, key []byte) ([]byte, error)

// 不同模式
func EncryptCBC(plaintext, key, iv []byte) ([]byte, error)
func DecryptCBC(ciphertext, key, iv []byte) ([]byte, error)
```

### RSA 操作

```go
// RSA 密钥生成
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error)

// RSA 加密/解密
func RSAEncryptOAEP(plaintext []byte, publicKey *rsa.PublicKey) ([]byte, error)
func RSADecryptOAEP(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error)
```

### 哈希函数

```go
// 常用哈希
func MD5(data []byte) string
func SHA1(data []byte) string
func SHA256(data []byte) string
func SHA512(data []byte) string

// HMAC
func HMACSHA256(data, key []byte) []byte
```

**使用示例：**
```go
import "github.com/lazygophers/utils/cryptox"

// AES 加密
key := []byte("32-byte-long-key-for-aes-256-enc")
plaintext := []byte("敏感数据")

ciphertext, err := cryptox.Encrypt(plaintext, key)
if err != nil {
    log.Fatal(err)
}

// 解密
decrypted, err := cryptox.Decrypt(ciphertext, key)
if err != nil {
    log.Fatal(err)
}

// 哈希
hash := cryptox.SHA256([]byte("数据"))
hmac := cryptox.HMACSHA256([]byte("数据"), []byte("密钥"))
```

---

## 时间扩展 (xtime)

扩展时间操作，包含中国农历支持。

### 农历计算

```go
// 农历日期结构
type LunarDate struct {
    Year  int
    Month int
    Day   int
    Leap  bool
}

// 阳历转农历
func SolarToLunar(year, month, day int) LunarDate

// 农历转阳历
func LunarToSolar(lunar LunarDate) (int, int, int)
```

### 节气计算

```go
// 获取节气
func GetSolarTerm(year int) []SolarTerm

// 节气结构
type SolarTerm struct {
    Name string
    Date time.Time
}
```

**使用示例：**
```go
import "github.com/lazygophers/utils/xtime"

// 农历转换
lunar := xtime.SolarToLunar(2024, 1, 1)
fmt.Printf("农历：%d年%d月%d日", lunar.Year, lunar.Month, lunar.Day)

// 获取节气
terms := xtime.GetSolarTerm(2024)
for _, term := range terms {
    fmt.Printf("%s: %s\n", term.Name, term.Date.Format("2006-01-02"))
}
```

---

## 随机数生成 (randx)

高性能随机数生成，拥有 38% 的测试覆盖率。

### 基本随机数

```go
// 整数随机数
func Int() int
func Intn(n int) int
func Int64() int64
func Int64n(n int64) int64

// 浮点随机数
func Float64() float64
func Float64Range(min, max float64) float64

// 布尔随机数
func Bool() bool
func Booln(n int) bool
```

### 随机选择

```go
// 从切片中随机选择
func Choose[T any](items []T) T

// 随机打乱
func Shuffle[T any](items []T)
```

**使用示例：**
```go
import "github.com/lazygophers/utils/randx"

// 随机数
num := randx.Intn(100)           // 0-99
float := randx.Float64Range(1.0, 10.0)  // 1.0-10.0
flag := randx.Bool()             // true/false

// 随机选择
fruits := []string{"苹果", "香蕉", "橙子"}
fruit := randx.Choose(fruits)    // 随机水果

// 打乱切片
randx.Shuffle(fruits)
```

---

## 网络工具 (network)

网络相关工具，拥有 89.1% 的测试覆盖率。

```go
// 获取真实 IP
func RealIpFromHeader(headers map[string]string) string

// 检查是否为本地 IP
func IsLocalIp(ip string) bool

// 获取网络接口 IP
func GetInterfaceIpByName(name string) (string, error)
```

**使用示例：**
```go
import "github.com/lazygophers/utils/network"

// HTTP 请求中获取真实 IP
headers := map[string]string{
    "X-Forwarded-For": "192.168.1.100",
    "X-Real-IP":       "10.0.0.1",
}
realIP := network.RealIpFromHeader(headers)

// 检查 IP 类型
isLocal := network.IsLocalIp("192.168.1.1")  // true
```

## 性能特性

- **零分配操作**: 许多函数实现零内存分配
- **原子操作**: 高并发场景的无锁实现
- **内存对齐**: 针对 CPU 缓存效率优化
- **泛型优化**: 类型安全操作无运行时反射开销

## 错误处理模式

所有包遵循一致的错误处理模式：
1. 使用 `github.com/lazygophers/log` 记录错误
2. 返回有意义的错误信息
3. 提供 `Must*` 函数用于关键操作

## 最佳实践

1. **类型安全**: 优先使用泛型函数确保类型安全
2. **性能**: 在热路径中使用 stringx 和 candy 包的优化函数
3. **并发**: 使用 wait 包进行安全的并发处理
4. **配置**: 使用 config 包统一管理应用配置
5. **加密**: 使用 cryptox 包进行安全的加密操作

此 API 参考提供了 LazyGophers Utils 库的全面使用指南，帮助开发者高效地使用这些工具构建高质量的 Go 应用程序。