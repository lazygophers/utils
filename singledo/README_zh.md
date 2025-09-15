# SingleDo - 单例执行模式实现

`singledo` 模块提供线程安全的单例执行模式实现，具有智能缓存和去重功能。它确保昂贵的操作在指定时间窗口内只执行一次，非常适合缓存昂贵的计算、API 调用和资源密集型操作。

## 功能特性

- **线程安全执行**: 保证每个键同时只有一次执行
- **基于时间的缓存**: 结果按配置的持续时间进行缓存
- **泛型支持**: 使用 Go 泛型完全类型安全
- **去重**: 对同一键的多个并发调用共享相同结果
- **分组管理**: 通过键组织操作以实现独立缓存
- **零内存分配**: 高效实现，开销最小
- **Panic 恢复**: 优雅处理执行函数中的 panic

## 安装

```bash
go get github.com/lazygophers/utils
```

## 使用方法

### 基本单次执行

```go
package main

import (
    "fmt"
    "time"
    "github.com/lazygophers/utils/singledo"
)

func main() {
    // 创建具有 5 分钟缓存持续时间的 Single 实例
    single := singledo.NewSingle[string](5 * time.Minute)

    // 将被缓存的昂贵操作
    expensiveOperation := func() (string, error) {
        fmt.Println("执行昂贵操作...")
        time.Sleep(2 * time.Second) // 模拟昂贵工作
        return "计算结果", nil
    }

    // 第一次调用 - 执行函数
    result1, err := single.Do(expensiveOperation)
    fmt.Printf("结果 1: %s, 错误: %v\n", result1, err)

    // 在缓存窗口内的第二次调用 - 返回缓存结果
    result2, err := single.Do(expensiveOperation)
    fmt.Printf("结果 2: %s, 错误: %v\n", result2, err)

    // 如需要可手动重置缓存
    single.Reset()
}
```

### 并发去重

```go
package main

import (
    "fmt"
    "sync"
    "time"
    "github.com/lazygophers/utils/singledo"
)

func main() {
    single := singledo.NewSingle[int](1 * time.Minute)

    expensiveComputation := func() (int, error) {
        fmt.Println("计算中...")
        time.Sleep(3 * time.Second)
        return 42, nil
    }

    var wg sync.WaitGroup

    // 启动多个 goroutine
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            result, err := single.Do(expensiveComputation)
            fmt.Printf("Goroutine %d 得到结果: %d, 错误: %v\n", id, result, err)
        }(i)
    }

    wg.Wait()
    // 只会打印一次 "计算中..."，所有 goroutine 得到相同结果
}
```

### 基于分组的键管理

```go
package main

import (
    "fmt"
    "time"
    "github.com/lazygophers/utils/singledo"
)

func main() {
    // 创建用于管理多个缓存操作的分组
    group := singledo.NewSingleGroup[string](2 * time.Minute)

    // 不同键的不同操作
    fetchUserData := func() (string, error) {
        fmt.Println("获取用户数据...")
        time.Sleep(1 * time.Second)
        return "用户数据", nil
    }

    fetchConfigData := func() (string, error) {
        fmt.Println("获取配置数据...")
        time.Sleep(1 * time.Second)
        return "配置数据", nil
    }

    // 使用不同键执行操作
    userData, _ := group.Do("user:123", fetchUserData)
    configData, _ := group.Do("config:app", fetchConfigData)

    fmt.Printf("用户: %s, 配置: %s\n", userData, configData)

    // 在缓存窗口内的后续调用返回缓存结果
    userData2, _ := group.Do("user:123", fetchUserData)
    fmt.Printf("缓存的用户数据: %s\n", userData2)
}
```

### API 响应缓存

```go
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
    "github.com/lazygophers/utils/singledo"
)

type APIResponse struct {
    Data    map[string]interface{} `json:"data"`
    Status  string                 `json:"status"`
}

func main() {
    // 缓存 API 响应 10 分钟
    apiCache := singledo.NewSingleGroup[*APIResponse](10 * time.Minute)

    fetchFromAPI := func(endpoint string) func() (*APIResponse, error) {
        return func() (*APIResponse, error) {
            fmt.Printf("对 %s 进行 API 调用...\n", endpoint)

            resp, err := http.Get("https://api.example.com/" + endpoint)
            if err != nil {
                return nil, err
            }
            defer resp.Body.Close()

            body, err := io.ReadAll(resp.Body)
            if err != nil {
                return nil, err
            }

            var apiResp APIResponse
            if err := json.Unmarshal(body, &apiResp); err != nil {
                return nil, err
            }

            return &apiResp, nil
        }
    }

    // 对同一端点的多次调用将被去重
    result1, err := apiCache.Do("users", fetchFromAPI("users"))
    if err != nil {
        fmt.Printf("错误: %v\n", err)
        return
    }

    result2, err := apiCache.Do("users", fetchFromAPI("users"))
    if err != nil {
        fmt.Printf("错误: %v\n", err)
        return
    }

    fmt.Printf("相同实例: %v\n", result1 == result2) // true - 相同的缓存实例
}
```

### 数据库查询缓存

```go
package main

import (
    "database/sql"
    "fmt"
    "time"
    "github.com/lazygophers/utils/singledo"
)

type User struct {
    ID   int    `db:"id"`
    Name string `db:"name"`
    Email string `db:"email"`
}

func main() {
    // 缓存数据库查询 5 分钟
    queryCache := singledo.NewSingleGroup[*User](5 * time.Minute)

    // 模拟数据库连接
    var db *sql.DB // 初始化您的数据库连接

    fetchUser := func(userID int) func() (*User, error) {
        return func() (*User, error) {
            fmt.Printf("查询数据库用户 ID: %d\n", userID)

            query := "SELECT id, name, email FROM users WHERE id = ?"
            row := db.QueryRow(query, userID)

            user := &User{}
            err := row.Scan(&user.ID, &user.Name, &user.Email)
            if err != nil {
                return nil, err
            }

            return user, nil
        }
    }

    // 缓存键包括用户 ID
    userKey := fmt.Sprintf("user:%d", 123)

    // 第一次调用访问数据库
    user1, err := queryCache.Do(userKey, fetchUser(123))
    if err != nil {
        fmt.Printf("错误: %v\n", err)
        return
    }

    // 第二次调用返回缓存结果
    user2, err := queryCache.Do(userKey, fetchUser(123))
    if err != nil {
        fmt.Printf("错误: %v\n", err)
        return
    }

    fmt.Printf("用户 1: %+v\n", user1)
    fmt.Printf("相同的缓存用户: %v\n", user1 == user2)
}
```

### 复杂数据处理

```go
package main

import (
    "crypto/md5"
    "fmt"
    "time"
    "github.com/lazygophers/utils/singledo"
)

type ProcessingResult struct {
    Hash      string
    Size      int
    Processed time.Time
}

func main() {
    // 缓存处理结果 30 分钟
    processor := singledo.NewSingleGroup[*ProcessingResult](30 * time.Minute)

    processData := func(data []byte) func() (*ProcessingResult, error) {
        return func() (*ProcessingResult, error) {
            fmt.Printf("处理 %d 字节数据...\n", len(data))

            // 模拟昂贵的处理
            time.Sleep(2 * time.Second)

            hash := fmt.Sprintf("%x", md5.Sum(data))

            return &ProcessingResult{
                Hash:      hash,
                Size:      len(data),
                Processed: time.Now(),
            }, nil
        }
    }

    data1 := []byte("Hello, World!")
    data2 := []byte("Hello, World!") // 相同内容
    data3 := []byte("Different data")

    // 使用内容哈希作为缓存键
    key1 := fmt.Sprintf("data:%x", md5.Sum(data1))
    key2 := fmt.Sprintf("data:%x", md5.Sum(data2))
    key3 := fmt.Sprintf("data:%x", md5.Sum(data3))

    result1, _ := processor.Do(key1, processData(data1))
    result2, _ := processor.Do(key2, processData(data2)) // 相同键，缓存结果
    result3, _ := processor.Do(key3, processData(data3)) // 不同键，新处理

    fmt.Printf("结果 1: %+v\n", result1)
    fmt.Printf("结果 2: %+v\n", result2)
    fmt.Printf("结果 3: %+v\n", result3)
    fmt.Printf("结果 1 和 2 是相同实例: %v\n", result1 == result2)
}
```

## API 参考

### Single 类型

#### `NewSingle[T any](wait time.Duration) *Single[T]`
为类型 T 创建具有指定缓存持续时间的新 Single 实例。

**参数:**
- `wait`: 缓存成功结果的持续时间

**返回值:**
- `*Single[T]`: 新的 Single 实例

#### `(s *Single[T]) Do(fn func() (T, error)) (T, error)`
如果尚未缓存或正在进行，执行函数，如果可用则返回缓存结果。

**参数:**
- `fn`: 要执行的函数（每个缓存窗口最多调用一次）

**返回值:**
- `T`: 函数的结果或缓存值
- `error`: 函数的错误或 nil

**行为:**
- 如果结果已缓存且未过期，立即返回缓存值
- 如果函数当前正在执行，等待完成并返回结果
- 如果没有缓存且没有正在执行，执行函数
- 只有成功的结果（error == nil）会被缓存

#### `(s *Single[T]) Reset()`
清除缓存结果，强制下次调用执行函数。

### Group 类型

#### `NewSingleGroup[T any](wait time.Duration) *Group[T]`
创建用于按键管理多个缓存操作的新 Group 实例。

**参数:**
- `wait`: 每个键缓存成功结果的持续时间

**返回值:**
- `*Group[T]`: 新的 Group 实例

#### `(g *Group[T]) Do(key string, fn func() (T, error)) (T, error)`
如果给定键未缓存或正在进行，执行函数。

**参数:**
- `key`: 操作的唯一标识符
- `fn`: 为此键执行的函数

**返回值:**
- `T`: 函数的结果或键的缓存值
- `error`: 函数的错误或 nil

**行为:**
- 每个键维护独立的缓存和执行状态
- 键永远不会自动清理（简单性的设计选择）
- 适用于有界键集或短期应用程序

## 最佳实践

### 1. 选择适当的缓存持续时间
```go
// 短期数据
userSession := singledo.NewSingle[*Session](5 * time.Minute)

// 配置数据
appConfig := singledo.NewSingle[*Config](1 * time.Hour)

// 静态参考数据
currencies := singledo.NewSingle[[]Currency](24 * time.Hour)
```

### 2. 适当处理错误
```go
result, err := single.Do(func() (string, error) {
    // 只有成功的结果会被缓存
    if someCondition {
        return "", errors.New("临时失败") // 不会缓存
    }
    return "成功", nil // 这将被缓存
})

if err != nil {
    // 处理错误 - 结果将是零值
    log.Printf("操作失败: %v", err)
    return
}
```

### 3. 为分组使用有意义的键
```go
group := singledo.NewSingleGroup[*User](10 * time.Minute)

// 好 - 描述性和唯一
userKey := fmt.Sprintf("user:id:%d", userID)
profileKey := fmt.Sprintf("profile:user:%d:full", userID)

// 避免 - 太通用或易冲突
badKey := fmt.Sprintf("%d", userID)
```

### 4. 考虑分组的内存使用
```go
// 分组永远不会自动清理键
// 对于无界键集，实现清理或使用 Single 实例

// 适合有界集
userCache := singledo.NewSingleGroup[*User](time.Hour)

// 对无界集考虑替代方案
// 或实现您自己的清理机制
```

## 性能注意事项

- **内存**: 分组保留对曾经使用的所有键的引用
- **并发**: 通过高效的 RWMutex 使用实现最小锁竞争
- **CPU**: 缓存命中几乎零开销
- **Goroutines**: 不创建 goroutine，同步执行模型

## 线程安全

singledo 模块完全线程安全：

- 多个 goroutine 可以安全地并发调用 `Do()`
- 每个键同时只会发生一次执行
- 对同一操作的并发调用将等待并接收相同结果
- `Reset()` 可以安全地并发调用（尽管可能导致缓存未命中）

## 错误处理

- 只有成功的结果（error == nil）会被缓存
- 错误立即返回且不被缓存
- 执行函数中的 panic 被恢复并作为错误返回
- 即使发生错误，缓存状态也得到正确维护

## 使用场景

1. **API 响应缓存**: 减少外部 API 调用
2. **数据库查询优化**: 缓存昂贵的查询
3. **计算缓存**: 缓存重计算的结果
4. **资源初始化**: 确保资源只初始化一次
5. **配置加载**: 缓存配置数据
6. **文件处理**: 缓存文件解析/处理的结果
7. **身份验证**: 缓存用户身份验证/授权结果

## 与其他模式的比较

### vs sync.Once
- **优势**: 基于时间的过期、错误处理、随时间多次执行
- **使用场景**: 当您需要定期重新执行而不是一次性初始化时

### vs 手动缓存
- **优势**: 线程安全、并发调用去重、内置过期
- **使用场景**: 当您需要的不仅仅是简单的基于 map 的缓存时

### vs 基于 Channel 的模式
- **优势**: 更简单的 API、无 goroutine 管理、缓存命中的即时结果
- **使用场景**: 当您不需要复杂的流控制或异步执行时

## 相关包

- [`sync`](https://pkg.go.dev/sync): 标准同步原语
- [`context`](https://pkg.go.dev/context): 用于超时和取消（考虑添加 context 支持）
- [`time`](https://pkg.go.dev/time): 基于时间的功能