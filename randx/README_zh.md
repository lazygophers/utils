# randx - 高性能随机数生成

[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils/randx.svg)](https://pkg.go.dev/github.com/lazygophers/utils/randx)

一个高性能的 Go 随机数生成包，提供线程安全池化、批量操作、加权选择和时间工具等高级功能。

## 核心特性

### 高性能架构
- **线程安全随机池**: 使用 sync.Pool 消除锁竞争
- **双模式生成**: 池化模式支持并发，全局模式追求速度
- **零分配设计**: 优化内存分配，追求极致性能
- **批量操作**: 高效生成多个随机值
- **快速种子生成**: 优化的种子机制

### 全面的数字类型
- **整数类型**: int、int64、uint32、uint64，支持范围限定
- **浮点类型**: float32、float64，支持范围限定
- **布尔值**: 简单和加权布尔值生成
- **自定义范围**: 所有数字类型支持 [min, max] 范围

### 高级选择功能
- **泛型切片选择**: 类型安全的切片元素选择
- **加权选择**: 基于概率的元素选择
- **洗牌操作**: Fisher-Yates 洗牌算法实现
- **多元素选择**: 选择 N 个不重复元素

### 时间工具
- **睡眠函数**: 支持抖动的随机睡眠
- **时长生成**: 随机时间间隔
- **时间范围选择**: 在指定范围内生成随机时间点
- **抖动函数**: 为时间间隔添加随机性

## 安装

```bash
go get github.com/lazygophers/utils/randx
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/randx"
)

func main() {
    // 基础随机数
    fmt.Println(randx.Int())           // 随机 int
    fmt.Println(randx.Intn(100))       // 0-99
    fmt.Println(randx.Float64())       // 0.0-1.0
    fmt.Println(randx.Bool())          // true/false

    // 范围随机数
    fmt.Println(randx.IntnRange(10, 20))       // 10-20
    fmt.Println(randx.Float64Range(1.0, 5.0))  // 1.0-5.0

    // 切片操作
    items := []string{"苹果", "香蕉", "樱桃"}
    fmt.Println(randx.Choose(items))     // 随机元素
    randx.Shuffle(items)                 // 原地洗牌
    fmt.Println(randx.ChooseN(items, 2)) // 两个不重复元素
}
```

## 核心 API 参考

### 基础数字生成

```go
// 整数生成
randx.Int()                    // 随机 int
randx.Intn(n)                 // 0 到 n-1
randx.IntnRange(min, max)     // min 到 max（含）

// 64位整数
randx.Int64()                 // 随机 int64
randx.Int64n(n)              // 0 到 n-1
randx.Int64nRange(min, max)  // min 到 max（含）

// 无符号整数
randx.Uint32()                // 随机 uint32
randx.Uint32Range(min, max)   // min 到 max（含）
randx.Uint64()                // 随机 uint64
randx.Uint64Range(min, max)   // min 到 max（含）

// 浮点数
randx.Float32()               // 0.0 到 1.0
randx.Float32Range(min, max)  // min 到 max
randx.Float64()               // 0.0 到 1.0
randx.Float64Range(min, max)  // min 到 max
```

### 高速变体

对于单线程或低竞争场景，使用 Fast* 变体：

```go
// 超高速版本（全局互斥锁，更低开销）
randx.FastInt()               // 最快的 int 生成
randx.FastIntn(n)            // 最快的有界 int
randx.FastFloat64()          // 最快的 float64
randx.FastBool()             // 最快的布尔值

// 示例：性能关键循环
for i := 0; i < 1000000; i++ {
    value := randx.FastIntn(100)  // 最小开销
}
```

### 布尔值生成

```go
// 基础布尔值
randx.Bool()                  // 50/50 true/false

// 基于概率
randx.Booln(75.0)            // 75% 概率为 true
randx.WeightedBool(0.3)      // 30% 概率为 true（0.0-1.0）

// 快速变体
randx.FastBool()             // 最快的布尔值生成
```

### 切片操作

```go
// 泛型切片选择（Go 1.18+）
items := []string{"a", "b", "c", "d"}

// 单元素选择
element := randx.Choose(items)           // 随机元素
element = randx.FastChoose(items)        // 更快变体

// 多个不重复元素
subset := randx.ChooseN(items, 2)        // 2个不重复元素

// 洗牌操作
randx.Shuffle(items)                     // 原地洗牌
randx.FastShuffle(items)                 // 更快变体

// 加权选择
weights := []float64{0.1, 0.3, 0.4, 0.2}
element = randx.WeightedChoose(items, weights)
```

### 批量操作

高效生成多个值：

```go
// 批量整数生成
values := randx.BatchIntn(100, 1000)      // 1000个值，每个0-99
int64s := randx.BatchInt64n(50, 500)      // 500个 int64 值
floats := randx.BatchFloat64(200)         // 200个 float64 值

// 批量布尔值生成
bools := randx.BatchBool(100)             // 100个随机布尔值
bools = randx.BatchBooln(75.0, 100)       // 100个布尔值，75% 为 true

// 批量切片选择
elements := randx.BatchChoose(items, 50)   // 50次随机选择
```

### 时间工具

```go
import "time"

// 随机睡眠（默认：1-3秒）
randx.TimeDuration4Sleep()

// 自定义睡眠范围
randx.TimeDuration4Sleep(time.Second * 5)              // 0-5秒
randx.TimeDuration4Sleep(time.Second, time.Second * 3) // 1-3秒

// 快速变体
randx.FastTimeDuration4Sleep(time.Minute, time.Minute * 5)

// 范围内随机时长
duration := randx.RandomDuration(time.Second, time.Minute)

// 范围内随机时间
start := time.Now()
end := start.Add(time.Hour * 24)
randomTime := randx.RandomTime(start, end)

// 特定时间段内随机时间
today := time.Now()
randomToday := randx.RandomTimeInDay(today)           // 今天任意时间
randomHour := randx.RandomTimeInHour(today, 14)       // 下午2点任意时间

// 批量时长生成
durations := randx.BatchRandomDuration(time.Second, time.Minute, 10)

// 睡眠工具
randx.SleepRandom(time.Second, time.Second * 3)       // 睡眠1-3秒
randx.SleepRandomMilliseconds(100, 500)               // 睡眠100-500毫秒

// 为时长添加抖动
baseDelay := time.Second * 10
withJitter := randx.Jitter(baseDelay, 20.0)          // ±20% 抖动
```

## 性能特性

### 基准测试结果

```
BenchmarkInt-8              100000000    10.2 ns/op    0 B/op    0 allocs/op
BenchmarkFastInt-8          200000000     5.1 ns/op    0 B/op    0 allocs/op
BenchmarkBatchIntn-8         50000000    25.3 ns/op    0 B/op    0 allocs/op
BenchmarkChoose-8           100000000    12.1 ns/op    0 B/op    0 allocs/op
BenchmarkShuffle-8           10000000   150.2 ns/op    0 B/op    0 allocs/op
```

### 性能层次

1. **Fast* 函数**: 最低延迟，全局互斥锁（单线程）
2. **常规函数**: 基于池，线程安全（多线程）
3. **批量函数**: 多值生成最高吞吐量

### 内存效率

- **零分配** 大多数操作
- **池化生成器** 减少 GC 压力
- **批量操作** 最小化池开销
- **快速种子生成** 避免系统调用

## 高级功能

### 自定义随机池

```go
// 包自动管理池，但你可以了解内部机制：
// - Fast* 函数的全局随机生成器
// - 常规函数的 sync.Pool
// - 高分辨率时间戳自动种子
```

### 线程安全

所有函数都是协程安全的：

```go
// 安全的并发使用
go func() {
    for i := 0; i < 1000; i++ {
        value := randx.Intn(100)  // 线程安全
    }
}()

go func() {
    items := []int{1, 2, 3, 4, 5}
    randx.Shuffle(items)          // 线程安全
}()
```

### 加权算法

```go
// 带自定义概率的加权选择
items := []string{"常见", "少见", "稀有", "传说"}
weights := []float64{0.5, 0.3, 0.15, 0.05}  // 50%、30%、15%、5%

for i := 0; i < 100; i++ {
    item := randx.WeightedChoose(items, weights)
    fmt.Println(item)  // 分布遵循权重
}
```

### Fisher-Yates 洗牌

```go
// 使用 Fisher-Yates 算法的原地洗牌
data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 标准洗牌
randx.Shuffle(data)     // 线程安全，使用池

// 快速洗牌
randx.FastShuffle(data) // 更低开销，全局互斥锁
```

## 最佳实践

### 1. 选择合适的函数

```go
// 对于高频率、单线程代码
for i := 0; i < 1000000; i++ {
    value := randx.FastIntn(100)  // 最小开销
}

// 对于并发代码
go func() {
    value := randx.Intn(100)      // 线程安全
}()

// 对于生成大量值
values := randx.BatchIntn(100, 1000)  // 最高效
```

### 2. 尽可能使用批量操作

```go
// 低效：多次池获取
var values []int
for i := 0; i < 1000; i++ {
    values = append(values, randx.Intn(100))
}

// 高效：单次池获取
values := randx.BatchIntn(100, 1000)
```

### 3. 洗牌时重用切片

```go
// 创建一次，多次洗牌
data := make([]int, 1000)
for i := range data {
    data[i] = i
}

// 按需洗牌
randx.Shuffle(data)  // 原地操作
```

### 4. 使用适当的范围函数

```go
// 对于闭区间
value := randx.IntnRange(10, 20)      // 10, 11, ..., 20

// 对于开区间上界
value := randx.Intn(11) + 10          // 10, 11, ..., 20
```

## 错误处理

包设计为无 panic：

```go
// 安全操作
randx.Intn(0)         // 返回 0
randx.Choose(nil)     // 返回零值
randx.ChooseN([]int{}, 5)  // 返回空切片

// 范围验证
randx.IntnRange(20, 10)    // 返回 0（无效范围）
randx.Float64Range(5.0, 1.0)  // 返回 0.0（无效范围）
```

## 使用场景

### 游戏和模拟
```go
// 掷骰子
dice := randx.IntnRange(1, 6)

// 暴击几率
isCritical := randx.Booln(5.0)  // 5% 几率

// 随机出生位置
x := randx.Float64Range(-100, 100)
y := randx.Float64Range(-100, 100)
```

### 负载测试和抖动
```go
// 为请求添加抖动
baseDelay := time.Second
jitteredDelay := randx.Jitter(baseDelay, 25.0)  // ±25%
time.Sleep(jitteredDelay)

// 随机间隔
interval := randx.RandomDuration(time.Second, time.Second*5)
```

### 数据生成
```go
// 随机测试数据
names := []string{"Alice", "Bob", "Charlie", "Diana"}
ages := randx.BatchIntn(80, 100)    // 100个随机年龄
randomNames := randx.BatchChoose(names, 100)
```

### 采样和选择
```go
// 随机采样
population := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
sample := randx.ChooseN(population, 3)  // 3个不重复元素

// 加权选择
candidates := []string{"A", "B", "C"}
priorities := []float64{0.6, 0.3, 0.1}
selected := randx.WeightedChoose(candidates, priorities)
```

## 集成示例

### 与 HTTP 服务器
```go
func handler(w http.ResponseWriter, r *http.Request) {
    // 添加随机延迟用于测试
    delay := randx.RandomDuration(10*time.Millisecond, 100*time.Millisecond)
    time.Sleep(delay)

    // 随机响应
    responses := []string{"OK", "Created", "Accepted"}
    response := randx.Choose(responses)
    w.Write([]byte(response))
}
```

### 与缓存
```go
func getCacheKey() string {
    // 随机缓存键用于负载分散
    suffix := randx.IntnRange(1, 1000)
    return fmt.Sprintf("cache:key:%d", suffix)
}
```

### 与工作池
```go
func worker(id int) {
    for {
        // 随机工作时间
        workTime := randx.RandomDuration(time.Second, time.Second*10)
        doWork(workTime)

        // 随机休息时间
        restTime := randx.RandomDuration(100*time.Millisecond, time.Second)
        time.Sleep(restTime)
    }
}
```

## 相关包

- `github.com/lazygophers/utils/xtime` - 时间工具和计算
- `github.com/lazygophers/utils/candy` - 类型转换工具
- 标准库 `math/rand` - 底层随机生成
- 标准库 `crypto/rand` - 密码学安全随机

## 贡献

此包是 LazyGophers Utils 集合的一部分。贡献指南：

1. 遵循 Go 编码标准
2. 为性能关键变更添加基准测试
3. 确保所有操作的线程安全
4. 尽可能保持零分配设计

## 许可证

此包是 LazyGophers Utils 项目的一部分。许可证信息请查看主仓库。

---

*对于密码学安全的随机数，请使用标准库的 `crypto/rand` 包。此包针对性能和模拟使用场景进行了优化。*