---
title: API 文檔
---

# API 文檔

本文檔提供 LazyGophers Utils 的詳細 API 參考。

## 核心工具

### utils.Must()

斷言操作成功，失敗時 panic。

```go
func Must[T any](value T, err error) T
```

**參數:**
- `value` - 返回值
- `err` - 錯誤

**返回:**
- 如果 `err` 為 nil，返回 `value`
- 如果 `err` 不為 nil，則 panic

**示例:**
```go
data := utils.Must(loadData())
```

### utils.MustSuccess()

斷言錯誤為 nil。

```go
func MustSuccess(err error)
```

**參數:**
- `err` - 錯誤

**示例:**
```go
utils.MustSuccess(config.Load(&cfg, "config.json"))
```

### utils.MustOk()

斷言第二個返回值為 true。

```go
func MustOk[T any](value T, ok bool) T
```

**參數:**
- `value` - 返回值
- `ok` - 成功標志

**返回:**
- 如果 `ok` 為 true，返回 `value`
- 如果 `ok` 為 false，則 panic

**示例:**
```go
value := utils.MustOk(getValue())
```

### utils.Validate()

驗證結構體數據。

```go
func Validate(v any) error
```

**參數:**
- `v` - 要驗證的結構體

**返回:**
- 如果驗證失敗，返回驗證錯誤

**示例:**
```go
type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
    Age   int    `validate:"min=0,max=150"`
}

user := User{
    Name:  "John Doe",
    Email: "john@example.com",
    Age:   25,
}

if err := utils.Validate(&user); err != nil {
    fmt.Printf("Validation failed: %v\n", err)
}
```

## 數據處理

### candy.ToInt()

將字符串轉換為整數。

```go
func ToInt(s string) int
```

**參數:**
- `s` - 字符串

**返回:**
- 整數值

**示例:**
```go
age := candy.ToInt("25")
```

### candy.ToFloat()

將字符串轉換為浮點數。

```go
func ToFloat(s string) float64
```

**參數:**
- `s` - 字符串

**返回:**
- 浮點數值

**示例:**
```go
price := candy.ToFloat("99.99")
```

### candy.ToBool()

將字符串轉換為布爾值。

```go
func ToBool(s string) bool
```

**參數:**
- `s` - 字符串

**返回:**
- 布爾值

**示例:**
```go
active := candy.ToBool("true")
```

### candy.ToString()

將任意類型轉換為字符串。

```go
func ToString(v any) string
```

**參數:**
- `v` - 任意值

**返回:**
- 字符串

**示例:**
```go
str := candy.ToString(123)
```

## 時間處理

### xtime.NowCalendar()

獲取當前日曆。

```go
func NowCalendar() *Calendar
```

**返回:**
- 當前日曆對象

**示例:**
```go
cal := xtime.NowCalendar()
fmt.Printf("Today: %s\n", cal.String())
```

### Calendar.LunarDate()

獲取農曆日期。

```go
func (c *Calendar) LunarDate() string
```

**返回:**
- 農曆日期字符串

**示例:**
```go
fmt.Printf("Lunar: %s\n", cal.LunarDate())
```

### Calendar.Animal()

獲取生肖動物。

```go
func (c *Calendar) Animal() string
```

**返回:**
- 生肖動物字符串

**示例:**
```go
fmt.Printf("Animal: %s\n", cal.Animal())
```

### Calendar.CurrentSolarTerm()

獲取當前節氣。

```go
func (c *Calendar) CurrentSolarTerm() string
```

**返回:**
- 節氣字符串

**示例:**
```go
fmt.Printf("Solar Term: %s\n", cal.CurrentSolarTerm())
```

## 配置管理

### config.Load()

加載配置文件。

```go
func Load(v any, filename string) error
```

**參數:**
- `v` - 配置結構體指針
- `filename` - 配置文件名

**返回:**
- 如果加載失敗，返回錯誤

**支持的格式:**
- JSON
- YAML
- TOML
- INI
- HCL

**示例:**
```go
type Config struct {
    Database string `json:"database"`
    Port     int    `json:"port"`
    Debug    bool   `json:"debug"`
}

var cfg Config
utils.MustSuccess(config.Load(&cfg, "config.json"))
```

## 並發控制

### routine.NewPool()

創建工作池。

```go
func NewPool(size int) *Pool
```

**參數:**
- `size` - 池大小

**返回:**
- 工作池對象

**示例:**
```go
pool := routine.NewPool(10)
defer pool.Close()
```

### Pool.Submit()

向工作池提交任務。

```go
func (p *Pool) Submit(fn func())
```

**參數:**
- `fn` - 要執行的函數

**示例:**
```go
pool.Submit(func() {
    fmt.Println("Task executed")
})
```

### wait.For()

等待條件。

```go
func For(timeout time.Duration, condition func() bool) bool
```

**參數:**
- `timeout` - 超時時間
- `condition` - 條件函數

**返回:**
- 是否在超時前滿足條件

**示例:**
```go
success := wait.For(5*time.Second, func() bool {
    return pool.Running() == 0
})
```

## 更多 API

完整的 API 文檔，請訪問：

- [pkg.go.dev](https://pkg.go.dev/github.com/lazygophers/utils)
- [GitHub 倉庫](https://github.com/lazygophers/utils)

## 相關文檔

- [入門指南](/zh-TW/guide/getting-started)
- [模組概覽](/zh-TW/modules/overview)
