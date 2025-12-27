---
title: API 文档
---

# API 文档

本文档提供了 LazyGophers Utils 的详细 API 参考。

## 核心工具

### utils.Must()

断言操作成功，失败时 panic。

```go
func Must[T any](value T, err error) T
```

**参数：**
- `value` - 返回值
- `err` - 错误

**返回：**
- 如果 err 为 nil，返回 value
- 如果 err 不为 nil，panic

**示例：**
```go
data := utils.Must(loadData())
```

### utils.MustSuccess()

断言错误为 nil。

```go
func MustSuccess(err error)
```

**参数：**
- `err` - 错误

**示例：**
```go
utils.MustSuccess(config.Load(&cfg, "config.json"))
```

### utils.MustOk()

断言第二个返回值为 true。

```go
func MustOk[T any](value T, ok bool) T
```

**参数：**
- `value` - 返回值
- `ok` - 是否成功

**返回：**
- 如果 ok 为 true，返回 value
- 如果 ok 为 false，panic

**示例：**
```go
value := utils.MustOk(getValue())
```

### utils.Validate()

验证结构体数据。

```go
func Validate(v any) error
```

**参数：**
- `v` - 要验证的结构体

**返回：**
- 验证错误，如果验证失败

**示例：**
```go
type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
    Age   int    `validate:"min=0,max=150"`
}

user := User{
    Name:  "张三",
    Email: "zhangsan@example.com",
    Age:   25,
}

if err := utils.Validate(&user); err != nil {
    fmt.Printf("验证失败: %v\n", err)
}
```

## 数据处理

### candy.ToInt()

字符串转整数。

```go
func ToInt(s string) int
```

**参数：**
- `s` - 字符串

**返回：**
- 整数值

**示例：**
```go
age := candy.ToInt("25")
```

### candy.ToFloat()

字符串转浮点数。

```go
func ToFloat(s string) float64
```

**参数：**
- `s` - 字符串

**返回：**
- 浮点数值

**示例：**
```go
price := candy.ToFloat("99.99")
```

### candy.ToBool()

字符串转布尔值。

```go
func ToBool(s string) bool
```

**参数：**
- `s` - 字符串

**返回：**
- 布尔值

**示例：**
```go
active := candy.ToBool("true")
```

### candy.ToString()

任意类型转字符串。

```go
func ToString(v any) string
```

**参数：**
- `v` - 任意值

**返回：**
- 字符串

**示例：**
```go
str := candy.ToString(123)
```

## 时间处理

### xtime.NowCalendar()

获取当前日历。

```go
func NowCalendar() *Calendar
```

**返回：**
- 当前日历对象

**示例：**
```go
cal := xtime.NowCalendar()
fmt.Printf("今天: %s\n", cal.String())
```

### Calendar.LunarDate()

获取农历日期。

```go
func (c *Calendar) LunarDate() string
```

**返回：**
- 农历日期字符串

**示例：**
```go
fmt.Printf("农历: %s\n", cal.LunarDate())
```

### Calendar.Animal()

获取生肖。

```go
func (c *Calendar) Animal() string
```

**返回：**
- 生肖字符串

**示例：**
```go
fmt.Printf("生肖: %s\n", cal.Animal())
```

### Calendar.CurrentSolarTerm()

获取当前节气。

```go
func (c *Calendar) CurrentSolarTerm() string
```

**返回：**
- 节气字符串

**示例：**
```go
fmt.Printf("节气: %s\n", cal.CurrentSolarTerm())
```

## 配置管理

### config.Load()

加载配置文件。

```go
func Load(v any, filename string) error
```

**参数：**
- `v` - 配置结构体指针
- `filename` - 配置文件名

**返回：**
- 错误，如果加载失败

**支持的格式：**
- JSON
- YAML
- TOML
- INI
- HCL

**示例：**
```go
type Config struct {
    Database string `json:"database"`
    Port     int    `json:"port"`
    Debug    bool   `json:"debug"`
}

var cfg Config
utils.MustSuccess(config.Load(&cfg, "config.json"))
```

## 并发控制

### routine.NewPool()

创建工作池。

```go
func NewPool(size int) *Pool
```

**参数：**
- `size` - 工作池大小

**返回：**
- 工作池对象

**示例：**
```go
pool := routine.NewPool(10)
defer pool.Close()
```

### Pool.Submit()

提交任务到工作池。

```go
func (p *Pool) Submit(fn func())
```

**参数：**
- `fn` - 要执行的函数

**示例：**
```go
pool.Submit(func() {
    fmt.Println("Task executed")
})
```

### wait.For()

等待条件满足。

```go
func For(timeout time.Duration, condition func() bool) bool
```

**参数：**
- `timeout` - 超时时间
- `condition` - 条件函数

**返回：**
- 是否在超时前满足条件

**示例：**
```go
success := wait.For(5*time.Second, func() bool {
    return pool.Running() == 0
})
```

## 更多 API

完整的 API 文档请访问：

- [pkg.go.dev](https://pkg.go.dev/github.com/lazygophers/utils)
- [GitHub 仓库](https://github.com/lazygophers/utils)

## 相关文档

- [快速开始](/zh-CN/guide/getting-started)
- [模块概览](/zh-CN/modules/overview)
