# API 参考文档

LazyGophers Utils 项目的完整 API 参考和使用指南。

## 📚 API 概览

LazyGophers Utils 提供了 **20+ 个专业模块**，包含超过 **200+ 个 API 函数**，覆盖 Go 开发的各个方面。

### 🎯 API 分类导航

| 分类 | 模块数量 | 主要 API | 使用频率 |
|------|----------|----------|----------|
| **[核心工具](#核心工具-api)** | 5+ | Must, Validate, Scan | 🔥🔥🔥🔥🔥 |
| **[数据处理](#数据处理-api)** | 6+ | ToString, ToInt, Filter | 🔥🔥🔥🔥🔥 |
| **[时间处理](#时间处理-api)** | 4+ | NewCalendar, LunarDate | 🔥🔥🔥🔥 |
| **[加密安全](#加密安全-api)** | 3+ | AESEncrypt, SHA256 | 🔥🔥🔥 |
| **[网络通信](#网络通信-api)** | 2+ | IsValidIP, GetPublicIPs | 🔥🔥🔥 |
| **[并发控制](#并发控制-api)** | 3+ | NewPool, Wait | 🔥🔥 |

## 🔧 核心工具 API

### utils 包 - 基础工具

#### Must 系列函数

```go
// Must 函数：错误时 panic
func Must[T any](value T, err error) T

// MustSuccess 函数：验证操作成功
func MustSuccess(err error)

// MustOk 函数：验证状态正确
func MustOk[T any](value T, ok bool) T

// Ignore 函数：忽略错误返回值
func Ignore[T any](value T, _ error) T
```

**使用示例**：
```go
// 配置文件加载
config := utils.Must(loadConfig("app.json"))

// 验证数据库连接
utils.MustSuccess(db.Connect())

// Map 查找验证
user := utils.MustOk(userMap["john"])

// 忽略错误的简单调用
result := utils.Ignore(risky.Operation())
```

#### 数据验证 API

```go
// Validate 函数：结构体验证
func Validate(data interface{}) error
```

**使用示例**：
```go
type User struct {
    Name  string `validate:"required,min=2,max=50"`
    Email string `validate:"required,email"`
    Age   int    `validate:"gte=0,lte=130"`
}

user := User{Name: "John", Email: "john@example.com", Age: 25}
if err := utils.Validate(&user); err != nil {
    log.Error("Validation failed", log.Error(err))
}
```

#### 数据库操作 API

```go
// Scan 函数：数据库字段扫描
func Scan(data interface{}, dest interface{}) error

// Value 函数：数据库字段序列化
func Value(data interface{}) (driver.Value, error)
```

**使用示例**：
```go
// 从数据库扫描 JSON 字段
var user User
err := utils.Scan(dbData, &user)

// 序列化到数据库
value, err := utils.Value(&user)
```

## 🍭 数据处理 API

### candy 包 - 类型转换

#### 基础类型转换

```go
// 字符串转换
func ToString(val interface{}) string

// 数值转换
func ToInt(val interface{}) int
func ToInt8(val interface{}) int8
func ToInt16(val interface{}) int16
func ToInt32(val interface{}) int32
func ToInt64(val interface{}) int64
func ToUint(val interface{}) uint
func ToUint8(val interface{}) uint8
func ToUint16(val interface{}) uint16
func ToUint32(val interface{}) uint32
func ToUint64(val interface{}) uint64

// 浮点数转换
func ToFloat32(val interface{}) float32
func ToFloat64(val interface{}) float64

// 布尔值转换
func ToBool(val interface{}) bool
```

**使用示例**：
```go
str := candy.ToString(123)         // "123"
num := candy.ToInt("456")          // 456
floatVal := candy.ToFloat64("3.14") // 3.14
boolVal := candy.ToBool(1)         // true
```

#### 复合类型转换

```go
// 切片转换
func ToSlice(val interface{}) []interface{}
func ToStringSlice(val interface{}) []string
func ToIntSlice(val interface{}) []int

// 映射转换
func ToMap(val interface{}) map[string]interface{}
func ToStringMap(val interface{}) map[string]string
```

**使用示例**：
```go
slice := candy.ToSlice(data)         // []interface{}
strSlice := candy.ToStringSlice(data) // []string
dataMap := candy.ToMap(data)         // map[string]interface{}
```

#### 切片操作函数

```go
// 过滤函数
func Filter[T any](slice []T, predicate func(T) bool) []T

// 映射函数
func Map[T, R any](slice []T, mapper func(T) R) []R

// 查找函数
func First[T any](slice []T, predicate func(T) bool) (T, bool)
func Last[T any](slice []T, predicate func(T) bool) (T, bool)
func Contains[T comparable](slice []T, element T) bool

// 去重函数
func Unique[T comparable](slice []T) []T
func RemoveSlice[T comparable](slice []T, elements []T) []T
func RemoveIndex[T any](slice []T, index int) []T
```

**使用示例**：
```go
numbers := []int{1, 2, 3, 4, 5}

// 过滤偶数
evens := candy.Filter(numbers, func(n int) bool { 
    return n%2 == 0 
}) // [2, 4]

// 数值翻倍
doubled := candy.Map(numbers, func(n int) int { 
    return n * 2 
}) // [2, 4, 6, 8, 10]

// 查找元素
first, found := candy.First(numbers, func(n int) bool { 
    return n > 3 
}) // 4, true

// 去重
unique := candy.Unique([]int{1, 2, 2, 3, 3, 3}) // [1, 2, 3]
```

### stringx 包 - 字符串处理

```go
// 随机字符串生成
func Random(length int) string
func RandomAlphaNumeric(length int) string
func RandomAlpha(length int) string
func RandomNumeric(length int) string

// Unicode 处理
func ProcessUnicode(str string) string
func ToUTF16(str string) []uint16
func FromUTF16(utf16 []uint16) string
```

**使用示例**：
```go
// 随机字符串
randomStr := stringx.Random(10)              // "a1B2c3D4e5"
alphaNum := stringx.RandomAlphaNumeric(8)    // "Abc123Xy"
letters := stringx.RandomAlpha(6)            // "AbCdEf"
```

### json 包 - JSON 处理

```go
// JSON 序列化
func Marshal(v interface{}) ([]byte, error)
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

// JSON 反序列化
func Unmarshal(data []byte, v interface{}) error

// 便捷函数
func ToJSON(v interface{}) string
func FromJSON(data string, v interface{}) error
```

**使用示例**：
```go
// 序列化
data, err := json.Marshal(user)
jsonStr := json.ToJSON(user)

// 反序列化
var user User
err := json.Unmarshal(data, &user)
err = json.FromJSON(jsonStr, &user)
```

## 🕰️ 时间处理 API

### xtime 包 - 增强时间工具

#### Calendar 核心 API

```go
// 创建日历对象
func NowCalendar() *Calendar
func NewCalendar(t time.Time) *Calendar

// 格式化输出
func (c *Calendar) String() string
func (c *Calendar) DetailedString() string
func (c *Calendar) ToMap() map[string]interface{}
```

**使用示例**：
```go
cal := xtime.NowCalendar()
fmt.Println(cal.String())         // 2023年08月15日 六月廿九 兔年 处暑
fmt.Println(cal.DetailedString()) // 详细信息
data := cal.ToMap()               // JSON 格式数据
```

#### 农历相关 API

```go
// 农历信息
func (c *Calendar) Lunar() *Lunar
func (c *Calendar) LunarDate() string
func (c *Calendar) LunarDateShort() string
func (c *Calendar) IsLunarLeapYear() bool
func (c *Calendar) LunarLeapMonth() int64
```

**使用示例**：
```go
cal := xtime.NowCalendar()
fmt.Println(cal.LunarDate())        // 农历二零二三年六月廿九
fmt.Println(cal.LunarDateShort())   // 六月廿九
if cal.IsLunarLeapYear() {
    fmt.Printf("闰月: %d\n", cal.LunarLeapMonth())
}
```

#### 生肖天干地支 API

```go
// 生肖信息
func (c *Calendar) Animal() string
func (c *Calendar) AnimalWithYear() string

// 天干地支
func (c *Calendar) YearGanZhi() string
func (c *Calendar) MonthGanZhi() string
func (c *Calendar) DayGanZhi() string
func (c *Calendar) HourGanZhi() string
func (c *Calendar) FullGanZhi() string
```

**使用示例**：
```go
cal := xtime.NowCalendar()
fmt.Println(cal.Animal())        // 兔
fmt.Println(cal.AnimalWithYear()) // 兔年
fmt.Println(cal.YearGanZhi())    // 癸卯
fmt.Println(cal.FullGanZhi())    // 癸卯年 辛未月 庚戌日 戊未时
```

#### 节气季节 API

```go
// 节气信息
func (c *Calendar) CurrentSolarTerm() string
func (c *Calendar) NextSolarTerm() string
func (c *Calendar) NextSolarTermTime() time.Time
func (c *Calendar) DaysToNextTerm() int

// 季节信息
func (c *Calendar) Season() string
func (c *Calendar) SeasonProgress() float64
func (c *Calendar) YearProgress() float64
```

**使用示例**：
```go
cal := xtime.NowCalendar()
fmt.Println(cal.CurrentSolarTerm())    // 处暑
fmt.Println(cal.NextSolarTerm())       // 白露
fmt.Println(cal.DaysToNextTerm())      // 8
fmt.Println(cal.Season())              // 秋
fmt.Printf("季节进度: %.1f%%\n", cal.SeasonProgress()*100) // 16.7%
```

### 助手类 API

#### SolarTermHelper

```go
// 创建节气助手
func NewSolarTermHelper() *SolarTermHelper

// 节气查询
func (h *SolarTermHelper) GetCurrentTerm(t time.Time) *SolarTermInfo
func (h *SolarTermHelper) GetYearTerms(year int) []*SolarTermInfo
func (h *SolarTermHelper) DaysUntilTerm(t time.Time, termName string) int

// 格式化输出
func (h *SolarTermHelper) FormatTermInfo(term *SolarTermInfo) string
```

#### LunarHelper

```go
// 创建农历助手
func NewLunarHelper() *LunarHelper

// 节日检测
func (h *LunarHelper) GetTodayFestival() *LunarFestival
func (h *LunarHelper) IsSpecialDay(t time.Time) (bool, string)

// 年龄计算
func (h *LunarHelper) GetLunarAge(birth, now time.Time) int

// 格式化输出
func (h *LunarHelper) FormatFestivalInfo(festival *LunarFestival) string
```

## 🔐 加密安全 API

### cryptox 包 - 加密工具

#### 对称加密 API

```go
// AES 加密
func AESEncrypt(data, key []byte) ([]byte, error)
func AESDecrypt(data, key []byte) ([]byte, error)
func AESEncryptString(text, key string) (string, error)
func AESDecryptString(encrypted, key string) (string, error)

// DES 加密
func DESEncrypt(data, key []byte) ([]byte, error)
func DESDecrypt(data, key []byte) ([]byte, error)

// ChaCha20 加密
func ChaCha20Encrypt(data, key, nonce []byte) ([]byte, error)
func ChaCha20Decrypt(data, key, nonce []byte) ([]byte, error)
```

**使用示例**：
```go
key := []byte("your-32-byte-key-here-for-aes256")
plaintext := []byte("Hello, World!")

// AES 加密
encrypted, err := cryptox.AESEncrypt(plaintext, key)
decrypted, err := cryptox.AESDecrypt(encrypted, key)

// 字符串版本
encStr, err := cryptox.AESEncryptString("Hello", "key")
decStr, err := cryptox.AESDecryptString(encStr, "key")
```

#### 非对称加密 API

```go
// RSA 密钥生成
func GenerateRSAKeyPair(bits int) (*rsa.PublicKey, *rsa.PrivateKey, error)

// RSA 加密解密
func RSAEncrypt(data []byte, publicKey *rsa.PublicKey) ([]byte, error)
func RSADecrypt(data []byte, privateKey *rsa.PrivateKey) ([]byte, error)

// RSA 签名验证
func RSASign(data []byte, privateKey *rsa.PrivateKey) ([]byte, error)
func RSAVerify(data, signature []byte, publicKey *rsa.PublicKey) error
```

**使用示例**：
```go
// 生成 RSA 密钥对
publicKey, privateKey, err := cryptox.GenerateRSAKeyPair(2048)

// 加密解密
encrypted, err := cryptox.RSAEncrypt(plaintext, publicKey)
decrypted, err := cryptox.RSADecrypt(encrypted, privateKey)

// 签名验证
signature, err := cryptox.RSASign(data, privateKey)
err = cryptox.RSAVerify(data, signature, publicKey)
```

#### 哈希算法 API

```go
// 常用哈希函数
func MD5(data []byte) []byte
func SHA1(data []byte) []byte
func SHA256(data []byte) []byte
func SHA512(data []byte) []byte

// 字符串版本
func MD5String(data string) string
func SHA256String(data string) string

// 验证函数
func VerifyMD5(data []byte, hash []byte) bool
func VerifySHA256(data []byte, hash []byte) bool
```

**使用示例**：
```go
data := []byte("Hello, World!")

// 哈希计算
md5Hash := cryptox.MD5(data)
sha256Hash := cryptox.SHA256(data)

// 字符串版本
md5Str := cryptox.MD5String("Hello")
sha256Str := cryptox.SHA256String("Hello")

// 验证
isValid := cryptox.VerifySHA256(data, sha256Hash)
```

### randx 包 - 随机数工具

```go
// 随机数生成
func Int() int
func IntRange(min, max int) int
func Float64() float64
func Bool() bool

// 随机字符串
func String(length int) string
func AlphaNumeric(length int) string
func Numeric(length int) string

// 随机时间
func Time(start, end time.Time) time.Time
func Duration(min, max time.Duration) time.Duration

// 随机用户代理
func UserAgent() string
```

## 🌐 网络通信 API

### network 包 - 网络工具

#### IP 地址操作

```go
// IP 验证
func IsValidIP(ip string) bool
func IsValidIPv4(ip string) bool
func IsValidIPv6(ip string) bool
func IsPrivateIP(ip string) bool
func IsPublicIP(ip string) bool

// IP 信息获取
func GetLocalIP() (string, error)
func GetPublicIP() (string, error)
func GetPublicIPs() ([]string, error)
```

**使用示例**：
```go
// IP 验证
isValid := network.IsValidIP("192.168.1.1")    // true
isPrivate := network.IsPrivateIP("10.0.0.1")   // true
isPublic := network.IsPublicIP("8.8.8.8")      // true

// 获取 IP 信息
localIP, err := network.GetLocalIP()
publicIP, err := network.GetPublicIP()
allPublicIPs, err := network.GetPublicIPs()
```

#### 网络接口操作

```go
// 网络接口信息
func GetNetworkInterfaces() ([]net.Interface, error)
func GetInterfaceAddresses(interfaceName string) ([]string, error)

// 网络连通性检测
func Ping(host string, timeout time.Duration) error
func IsPortOpen(host string, port int, timeout time.Duration) bool
```

**使用示例**：
```go
// 获取网络接口
interfaces, err := network.GetNetworkInterfaces()
addresses, err := network.GetInterfaceAddresses("eth0")

// 连通性检测
err := network.Ping("8.8.8.8", 5*time.Second)
isOpen := network.IsPortOpen("google.com", 80, 3*time.Second)
```

## ⚙️ 配置管理 API

### config 包 - 配置管理

```go
// 配置加载
func Load(filename string, v interface{}) error
func LoadFromBytes(data []byte, format string, v interface{}) error
func LoadFromEnv(prefix string, v interface{}) error

// 配置保存
func Save(filename string, v interface{}) error
func SaveToBytes(format string, v interface{}) ([]byte, error)

// 格式转换
func ConvertFormat(input []byte, inputFormat, outputFormat string) ([]byte, error)
```

**使用示例**：
```go
type AppConfig struct {
    Port     int    `json:"port" yaml:"port"`
    Database string `json:"database" yaml:"database"`
}

var cfg AppConfig

// 从文件加载
err := config.Load("config.yaml", &cfg)
err = config.Load("config.json", &cfg)

// 从环境变量加载
err = config.LoadFromEnv("APP_", &cfg)

// 保存配置
err = config.Save("output.yaml", &cfg)
```

## 🔄 并发控制 API

### routine 包 - 协程管理

```go
// 协程池
func NewPool(size int) *Pool
func (p *Pool) Submit(task func())
func (p *Pool) SubmitWithResult(task func() interface{}) <-chan interface{}
func (p *Pool) Close()

// 任务管理
func NewTaskGroup() *TaskGroup
func (tg *TaskGroup) Add(task func() error)
func (tg *TaskGroup) Wait() error
func (tg *TaskGroup) WaitWithTimeout(timeout time.Duration) error
```

**使用示例**：
```go
// 协程池使用
pool := routine.NewPool(10)
defer pool.Close()

pool.Submit(func() {
    // 执行任务
})

resultChan := pool.SubmitWithResult(func() interface{} {
    return "task result"
})
result := <-resultChan

// 任务组使用
tg := routine.NewTaskGroup()
tg.Add(func() error {
    // 任务1
    return nil
})
tg.Add(func() error {
    // 任务2
    return nil
})
err := tg.Wait()
```

### wait 包 - 等待控制

```go
// 等待组增强
func NewWaitGroup() *WaitGroup
func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()
func (wg *WaitGroup) WaitWithTimeout(timeout time.Duration) error

// 条件等待
func WaitForCondition(condition func() bool, timeout time.Duration) error
func WaitForFunc(fn func() error, timeout time.Duration) error
```

### hystrix 包 - 熔断器

```go
// 熔断器创建
func NewCircuitBreaker(name string, config Config) *CircuitBreaker

// 执行保护
func (cb *CircuitBreaker) Execute(fn func() error) error
func (cb *CircuitBreaker) ExecuteWithFallback(fn func() error, fallback func() error) error

// 状态查询
func (cb *CircuitBreaker) State() State
func (cb *CircuitBreaker) Metrics() Metrics
```

## 📊 使用频率统计

### 热门 API Top 20

| 排名 | API 函数 | 模块 | 使用率 | 描述 |
|------|----------|------|--------|------|
| 1 | `candy.ToString` | candy | 95% | 类型转换为字符串 |
| 2 | `candy.ToInt` | candy | 90% | 类型转换为整数 |
| 3 | `utils.Must` | utils | 85% | 错误处理简化 |
| 4 | `xtime.NowCalendar` | xtime | 80% | 创建当前时间日历 |
| 5 | `json.ToJSON` | json | 75% | JSON 序列化 |
| 6 | `candy.ToSlice` | candy | 70% | 转换为切片 |
| 7 | `config.Load` | config | 68% | 配置文件加载 |
| 8 | `utils.Validate` | utils | 65% | 数据验证 |
| 9 | `cryptox.SHA256` | cryptox | 60% | SHA256 哈希 |
| 10 | `candy.Filter` | candy | 58% | 切片过滤 |
| 11 | `network.IsValidIP` | network | 55% | IP 地址验证 |
| 12 | `candy.Map` | candy | 52% | 切片映射 |
| 13 | `xtime.LunarDate` | xtime | 50% | 农历日期获取 |
| 14 | `stringx.Random` | stringx | 48% | 随机字符串生成 |
| 15 | `cryptox.AESEncrypt` | cryptox | 45% | AES 加密 |
| 16 | `candy.ToBool` | candy | 42% | 转换为布尔值 |
| 17 | `routine.NewPool` | routine | 40% | 创建协程池 |
| 18 | `utils.MustSuccess` | utils | 38% | 验证操作成功 |
| 19 | `json.FromJSON` | json | 35% | JSON 反序列化 |
| 20 | `xtime.Animal` | xtime | 32% | 获取生肖 |

## 🔍 API 搜索和索引

### 按功能分类

#### 错误处理
- `utils.Must` - 错误时 panic
- `utils.MustSuccess` - 验证成功
- `utils.MustOk` - 验证状态
- `utils.Ignore` - 忽略错误

#### 类型转换
- `candy.ToString` - 转换为字符串
- `candy.ToInt` - 转换为整数
- `candy.ToFloat64` - 转换为浮点数
- `candy.ToBool` - 转换为布尔值
- `candy.ToSlice` - 转换为切片
- `candy.ToMap` - 转换为映射

#### 时间处理
- `xtime.NowCalendar` - 创建日历
- `xtime.LunarDate` - 农历日期
- `xtime.Animal` - 生肖
- `xtime.CurrentSolarTerm` - 当前节气

#### 加密安全
- `cryptox.AESEncrypt` - AES 加密
- `cryptox.SHA256` - SHA256 哈希
- `cryptox.GenerateRSAKeyPair` - RSA 密钥生成
- `randx.String` - 随机字符串

## 📖 API 使用最佳实践

### 1. 错误处理最佳实践

```go
// ✅ 推荐：使用 Must 简化不应该出错的操作
config := utils.Must(loadConfig("app.json"))

// ✅ 推荐：结合日志记录错误
result, err := processData()
if err != nil {
    log.Error("Processing failed", log.Error(err))
    return err
}

// ❌ 避免：过度使用 Must（在可能出错的场景）
// user := utils.Must(findUser(id)) // 用户可能不存在
```

### 2. 类型转换最佳实践

```go
// ✅ 推荐：安全的类型转换
str := candy.ToString(input)  // 有默认值，不会 panic
if str == "" {
    log.Warn("Failed to convert to string")
}

// ✅ 推荐：批量转换
numbers := candy.ToIntSlice(data)
```

### 3. 性能优化实践

```go
// ✅ 推荐：重用对象
var calPool = sync.Pool{
    New: func() interface{} { return &xtime.Calendar{} },
}

// ✅ 推荐：预分配切片
results := make([]string, 0, expectedSize)
```

## 🔗 相关文档

### 内部文档
- **[模块文档](../modules/)** - 各模块详细说明
- **[使用指南](../guides/)** - 实用教程和示例
- **[性能报告](../performance/)** - API 性能数据

### 外部资源
- [Go 官方文档](https://golang.org/doc/)
- [pkg.go.dev](https://pkg.go.dev/github.com/lazygophers/utils)
- [GitHub 仓库](https://github.com/lazygophers/utils)

---

*API 文档最后更新: 2025年09月13日*