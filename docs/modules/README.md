# LazyGophers Utils - 模块详细文档

## 📋 概述

本目录包含 LazyGophers Utils 所有模块的详细文档。每个模块都有独立的文档目录，提供完整的API参考、使用示例、性能分析和最佳实践。

## 🗂️ 模块分类

### 🔧 核心工具模块

#### [candy](./candy/) - 类型转换与语法糖
**功能**: 高性能类型转换、集合操作、数学运算和实用工具
- ✅ **99.3% 测试覆盖率**
- ⚡ **零分配转换** - 基础类型转换实现零内存分配
- 🔄 **泛型优化** - 使用 Go 1.18+ 泛型消除反射开销
- 📊 **丰富集合操作** - Filter, Map, Reduce 等函数式编程支持

**核心功能**:
- 类型转换: `ToBool()`, `ToString()`, `ToInt64()` 等
- 集合操作: `All()`, `Any()`, `Filter()`, `Unique()` 等  
- 数学运算: `Sum()`, `Average()`, `Min()`, `Max()` 等
- 实用工具: `DeepCopy()`, `Chunk()`, `Random()` 等

#### [stringx](./stringx/) - 高性能字符串处理
**功能**: 零拷贝字符串操作、命名风格转换、Unicode支持
- ⚡ **零拷贝转换** - 字符串/字节切片无内存拷贝转换
- 🚀 **ASCII优化** - 针对ASCII字符的快速路径优化
- 🔧 **命名转换** - 驼峰/蛇形等命名风格转换
- 🌍 **Unicode安全** - 完整的Unicode字符支持

**核心功能**:
- 零拷贝: `ToString()`, `ToBytes()` - 0 ns/op
- 命名转换: `Camel2Snake()`, `Snake2Camel()`
- 字符串检查: `IsASCII()`, `ContainsAny()`
- 随机生成: `RandString()`, `RandNumeric()`

#### [xtime](./xtime/) - 增强时间处理
**功能**: 中国传统历法、节气计算、工作制时间支持
- 🗓️ **统一日历接口** - 整合公历农历信息
- 🌙 **精确农历转换** - 支持1900-2100年精确转换  
- 🐲 **生肖干支系统** - 完整的天干地支计算
- 🏮 **24节气支持** - 精确的节气时间和进度计算

**核心功能**:
- 日历对象: `NewCalendar()`, `NowCalendar()`
- 农历信息: `LunarDate()`, `Animal()`, `YearGanZhi()`
- 节气季节: `CurrentSolarTerm()`, `Season()`, `DaysToNextTerm()`
- 工作制: XTime007, XTime955, XTime996

### 🔄 并发控制模块

#### [wait](./wait/) - 并发控制与工作池
**功能**: 工作池管理、任务调度、同步控制
- 🏭 **工作池模式** - 高效的goroutine生命周期管理
- 📋 **任务队列** - 缓冲区任务分发和负载均衡
- ♻️ **资源复用** - 对象池减少内存分配
- 🔄 **优雅关闭** - 确保所有任务完成后退出

**核心功能**:
- 工作池: `NewWorker()`, `Add()`, `Wait()`
- 同步工具: `WaitGroupWithTimeout()`, `WaitGroupWithContext()`
- 异步执行: `AsyncExecute()`, `AsyncExecuteWithTimeout()`

#### [hystrix](./hystrix/) - 熔断器模式
**功能**: 故障隔离、快速失败、自动恢复
- ⚡ **无锁设计** - 使用原子操作实现无锁并发
- 🔄 **三种状态** - Closed/Open/Half-Open状态管理
- 🎯 **灵活配置** - 自定义熔断条件和探测策略
- 📊 **性能监控** - 详细的成功率和响应时间统计

**核心功能**:
- 熔断器: `NewCircuitBreaker()`, `Call()`, `State()`
- 高性能变体: `FastCircuitBreaker`, `BulkCircuitBreaker`
- 配置选项: 时间窗口、熔断条件、状态回调

### 🔐 安全与加密模块

#### cryptox - 密码学工具
**功能**: 加密解密、哈希计算、密钥管理
- 🔒 **多种算法** - 支持AES、RSA、ECDSA等主流算法
- 🗝️ **密钥管理** - 安全的密钥生成、存储和轮转
- 🛡️ **数据保护** - 敏感数据加密和安全传输
- ✅ **100%测试覆盖率** - 生产就绪的密码学操作

#### pgp - PGP加密签名
**功能**: PGP/GPG加密和数字签名
- 📧 **邮件加密** - 兼容标准PGP邮件加密
- ✍️ **数字签名** - 完整的签名生成和验证
- 🔑 **密钥管理** - PGP密钥对生成和管理

### 📊 数据处理模块

#### anyx - Any类型处理
**功能**: 动态类型处理和反射工具
- 🔍 **类型检查** - 安全的类型断言和检查
- 🔄 **动态转换** - 灵活的类型转换机制
- 📋 **结构体操作** - 动态的结构体字段访问和修改

#### json - JSON处理增强
**功能**: 高性能JSON序列化和反序列化
- ⚡ **高性能** - 基于sonic的高性能JSON处理
- 🔄 **易用接口** - 简化的JSON操作API
- 🛡️ **类型安全** - 强类型的JSON操作支持

#### config - 配置管理
**功能**: 多格式配置文件读取和管理
- 📄 **多格式支持** - JSON、YAML、TOML等格式
- 🔄 **热重载** - 配置文件变化自动重载
- 🌍 **环境变量** - 环境变量覆盖和默认值支持

### 🛠️ 系统工具模块

#### app - 应用框架
**功能**: 应用程序框架和生命周期管理
- 🚀 **应用启动** - 标准化的应用启动流程
- 🔧 **生命周期** - 完整的应用生命周期管理
- ⚙️ **配置集成** - 统一的配置管理集成

#### atexit - 退出处理
**功能**: 程序退出时的清理和钩子管理
- 🔄 **清理回调** - 程序退出时的资源清理
- 🎯 **钩子管理** - 灵活的退出钩子注册
- 📊 **性能**: 注册操作约46ns/op，零分配

#### runtime - 运行时信息
**功能**: 运行时信息获取和处理
- 📊 **系统信息** - CPU、内存、goroutine统计
- 🔍 **性能分析** - 运行时性能监控
- 📈 **指标收集** - 应用程序指标收集

### 🌐 网络与IO模块

#### network - 网络工具
**功能**: 网络操作和连接管理
- 🌐 **连接管理** - HTTP/TCP连接池和管理
- 🔄 **重试机制** - 智能的网络请求重试
- 📊 **性能监控** - 网络请求性能统计

#### bufiox - 缓冲IO增强
**功能**: 增强的缓冲区操作
- 📦 **缓冲优化** - 高效的缓冲区管理
- 🔄 **流处理** - 流式数据处理支持
- 💾 **内存管理** - 智能的内存分配和回收

### 🎲 实用工具模块

#### randx - 随机数生成
**功能**: 增强的随机数生成工具
- 🎲 **多种分布** - 均匀、正态、指数等分布
- 🔒 **密码学安全** - 加密安全的随机数生成
- 🎯 **高性能** - 优化的随机数生成算法

#### fake - 测试数据生成
**功能**: 用于测试的假数据生成
- 👤 **人员信息** - 姓名、邮箱、电话等假数据
- 🌍 **地理信息** - 地址、城市、国家等
- 💼 **业务数据** - 公司、产品、订单等

#### unit - 单元测试辅助
**功能**: 单元测试工具和断言
- ✅ **测试断言** - 丰富的测试断言函数
- 🔧 **测试工具** - 测试数据生成和验证
- 📊 **性能测试** - 基准测试工具支持

#### defaults - 默认值处理
**功能**: 结构体默认值设置
- 🔧 **自动填充** - 自动设置结构体字段默认值
- 🏷️ **标签支持** - 基于struct tag的默认值配置
- 🔄 **递归处理** - 嵌套结构体的默认值设置

#### urlx - URL处理工具
**功能**: URL解析、构建和操作
- 🔗 **URL构建** - 安全的URL构建和参数编码
- 🔍 **路径解析** - 智能的路径解析和验证
- 🌐 **国际化** - 国际化域名和URL支持

#### osx - 操作系统工具
**功能**: 操作系统相关的增强工具
- 🖥️ **系统信息** - 操作系统信息获取
- 📂 **文件操作** - 增强的文件和目录操作
- 🔧 **进程管理** - 进程启动和管理工具

#### singledo - 单例模式
**功能**: 单例模式实现
- 🔒 **线程安全** - 并发安全的单例实现
- ⚡ **高性能** - 优化的单例获取性能
- 🎯 **类型安全** - 泛型支持的类型安全单例

#### routine - Goroutine管理
**功能**: Goroutine生命周期管理和监控
- 📊 **生命周期管理** - 自动管理goroutine创建和销毁
- 🔍 **运行时监控** - 实时监控goroutine状态和性能
- ⚠️ **异常恢复** - 自动捕获和恢复panic
- 📈 **性能统计** - 详细的执行时间和资源使用统计

#### event - 事件驱动
**功能**: 事件驱动编程支持
- 📢 **事件发布** - 高效的事件发布机制
- 👂 **事件订阅** - 灵活的事件订阅管理
- 🔄 **异步处理** - 异步的事件处理支持

#### pyroscope - 性能分析
**功能**: 性能分析和监控集成
- 📊 **性能分析** - 持续的性能profiling
- 🔍 **火焰图** - 可视化的性能分析图表
- 📈 **监控集成** - 与监控系统的集成

## 📊 模块统计信息

| 模块 | 功能分类 | 测试覆盖率 | 性能等级 | 推荐指数 |
|------|----------|------------|----------|----------|
| **candy** | 核心工具 | 99.3% | A+ | ⭐⭐⭐⭐⭐ |
| **stringx** | 核心工具 | 95.2% | A+ | ⭐⭐⭐⭐⭐ |
| **xtime** | 核心工具 | 72.7% | A | ⭐⭐⭐⭐⭐ |
| **wait** | 并发控制 | 88.5% | A+ | ⭐⭐⭐⭐⭐ |
| **hystrix** | 并发控制 | 91.3% | A+ | ⭐⭐⭐⭐⭐ |
| **cryptox** | 安全加密 | 100% | A | ⭐⭐⭐⭐ |
| **anyx** | 数据处理 | 85.7% | A | ⭐⭐⭐⭐ |
| **json** | 数据处理 | 92.1% | A+ | ⭐⭐⭐⭐ |
| **config** | 系统工具 | 78.9% | A | ⭐⭐⭐⭐ |
| **routine** | 并发控制 | 83.4% | A | ⭐⭐⭐ |
| **app** | 系统工具 | 75.2% | A | ⭐⭐⭐ |
| **atexit** | 系统工具 | 88.7% | A+ | ⭐⭐⭐ |
| **network** | 网络IO | 82.3% | A | ⭐⭐⭐ |
| **runtime** | 系统工具 | 79.1% | A | ⭐⭐⭐ |

## 🚀 快速开始

### 基础使用
```go
import (
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/stringx"
    "github.com/lazygophers/utils/wait"
)

func main() {
    // 类型转换
    str := candy.ToString(123)
    num := candy.ToInt("456")
    
    // 字符串处理
    snake := stringx.Camel2Snake("firstName")
    
    // 并发处理
    worker := wait.NewWorker(10)
    defer worker.Wait()
    
    worker.Add(func() {
        fmt.Println("任务执行")
    })
}
```

### 高级用法
```go
// 组合使用多个模块
func ProcessDataConcurrently(data []string) error {
    // 创建熔断器保护
    cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
        TimeWindow: 30 * time.Second,
        ReadyToTrip: func(successes, failures uint64) bool {
            total := successes + failures
            return total >= 10 && float64(failures)/float64(total) >= 0.5
        },
    })
    
    // 创建工作池
    worker := wait.NewWorker(runtime.NumCPU())
    defer worker.Wait()
    
    // 并发处理数据
    for _, item := range data {
        item := item
        worker.Add(func() {
            err := cb.Call(func() error {
                // 转换和处理数据
                processed := stringx.Camel2Snake(item)
                return processItem(processed)
            })
            
            if err != nil {
                log.Printf("处理失败: %v", err)
            }
        })
    }
    
    return nil
}
```

## 📖 使用指南

### 选择合适的模块

1. **类型转换需求** → 使用 [candy](./candy/)
2. **字符串处理** → 使用 [stringx](./stringx/)  
3. **时间处理** → 使用 [xtime](./xtime/)
4. **并发控制** → 使用 [wait](./wait/) + [hystrix](./hystrix/)
5. **数据序列化** → 使用 [json](./json/)
6. **加密需求** → 使用 [cryptox](./cryptox/)

### 性能优化建议

1. **高频操作**: 优先使用 candy、stringx 等零分配模块
2. **并发场景**: 结合 wait 和 hystrix 实现高效并发控制
3. **内存敏感**: 利用各模块的对象池和缓存机制
4. **错误处理**: 使用 hystrix 实现优雅的故障处理

### 最佳实践

1. **模块组合**: 多个模块组合使用，发挥最大效能
2. **错误处理**: 每个模块都有完善的错误处理机制
3. **性能监控**: 利用内置的性能统计和监控功能
4. **文档优先**: 每个模块都有详细的文档和示例

## 🔗 相关资源

- **[API参考文档](../API_REFERENCE.md)**: 完整的API文档
- **[性能报告](../performance_report.md)**: 详细的性能分析
- **[架构文档](../architecture_en.md)**: 系统架构设计
- **[贡献指南](../CONTRIBUTING_en.md)**: 如何参与贡献

## 📞 支持与反馈

- 🐛 [报告问题](https://github.com/lazygophers/utils/issues)
- 💬 [讨论交流](https://github.com/lazygophers/utils/discussions)
- 📧 [联系我们](mailto:support@lazygophers.com)

---

**最后更新**: 2025年9月13日
**文档版本**: v2.0.0