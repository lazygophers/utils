# runtime - 运行时系统信息

`runtime` 包提供了收集关于 Go 应用程序及其运行环境的系统和运行时信息的工具。它提供全面的系统指标、内存统计和运行时诊断。

## 功能特性

- **系统信息**: CPU、内存、磁盘和网络信息
- **运行时指标**: Go 运行时统计和内存使用
- **进程信息**: 当前进程详细信息和资源使用
- **环境检测**: 操作系统和架构检测
- **性能监控**: 运行时性能指标和性能分析数据
- **跨平台支持**: 在 Linux、macOS 和 Windows 上工作

## 安装

```bash
go get github.com/lazygophers/utils/runtime
```

## 使用示例

### 基本系统信息

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/runtime"
)

func main() {
    // 获取系统信息
    sysInfo := runtime.GetSystemInfo()
    fmt.Printf("操作系统: %s\n", sysInfo.OS)
    fmt.Printf("架构: %s\n", sysInfo.Arch)
    fmt.Printf("CPU 核心: %d\n", sysInfo.CPUCores)
    fmt.Printf("总内存: %d MB\n", sysInfo.TotalMemory/1024/1024)

    // 获取运行时信息
    runtimeInfo := runtime.GetRuntimeInfo()
    fmt.Printf("Go 版本: %s\n", runtimeInfo.GoVersion)
    fmt.Printf("Goroutines: %d\n", runtimeInfo.NumGoroutines)
    fmt.Printf("CGO 调用: %d\n", runtimeInfo.NumCGOCalls)
}
```

### 内存统计

```go
// 获取详细内存统计
memStats := runtime.GetMemoryStats()
fmt.Printf("已分配内存: %d KB\n", memStats.Alloc/1024)
fmt.Printf("总分配: %d\n", memStats.TotalAlloc)
fmt.Printf("系统内存: %d KB\n", memStats.Sys/1024)
fmt.Printf("GC 周期: %d\n", memStats.NumGC)
fmt.Printf("上次 GC: %v\n", memStats.LastGC)

// 获取堆信息
heapInfo := runtime.GetHeapInfo()
fmt.Printf("堆大小: %d KB\n", heapInfo.HeapAlloc/1024)
fmt.Printf("堆对象: %d\n", heapInfo.HeapObjects)
fmt.Printf("下次 GC: %d KB\n", heapInfo.NextGC/1024)
```

### 进程信息

```go
// 获取进程信息
processInfo := runtime.GetProcessInfo()
fmt.Printf("PID: %d\n", processInfo.PID)
fmt.Printf("PPID: %d\n", processInfo.PPID)
fmt.Printf("可执行文件: %s\n", processInfo.Executable)
fmt.Printf("工作目录: %s\n", processInfo.WorkingDir)
fmt.Printf("启动时间: %v\n", processInfo.StartTime)
fmt.Printf("CPU 使用率: %.2f%%\n", processInfo.CPUPercent)
fmt.Printf("内存使用: %d KB\n", processInfo.MemoryUsage/1024)
```

### CPU 信息

```go
// 获取 CPU 信息
cpuInfo := runtime.GetCPUInfo()
fmt.Printf("CPU 型号: %s\n", cpuInfo.ModelName)
fmt.Printf("CPU 核心: %d\n", cpuInfo.Cores)
fmt.Printf("CPU 线程: %d\n", cpuInfo.Threads)
fmt.Printf("CPU 频率: %.2f GHz\n", cpuInfo.MHz/1000)
fmt.Printf("缓存大小: %d KB\n", cpuInfo.CacheSize)

// 获取 CPU 使用率
cpuUsage := runtime.GetCPUUsage()
fmt.Printf("CPU 使用率: %.2f%%\n", cpuUsage.Total)
fmt.Printf("用户时间: %.2f%%\n", cpuUsage.User)
fmt.Printf("系统时间: %.2f%%\n", cpuUsage.System)
```

### 磁盘信息

```go
// 获取磁盘使用信息
diskInfo := runtime.GetDiskInfo()
for _, disk := range diskInfo {
    fmt.Printf("设备: %s\n", disk.Device)
    fmt.Printf("挂载点: %s\n", disk.Mountpoint)
    fmt.Printf("文件系统: %s\n", disk.Fstype)
    fmt.Printf("总计: %d GB\n", disk.Total/1024/1024/1024)
    fmt.Printf("已用: %d GB\n", disk.Used/1024/1024/1024)
    fmt.Printf("可用: %d GB\n", disk.Free/1024/1024/1024)
    fmt.Printf("使用率: %.2f%%\n", disk.UsedPercent)
    fmt.Println("---")
}
```

### 网络信息

```go
// 获取网络接口
interfaces := runtime.GetNetworkInterfaces()
for _, iface := range interfaces {
    fmt.Printf("接口: %s\n", iface.Name)
    fmt.Printf("MAC 地址: %s\n", iface.HardwareAddr)
    fmt.Printf("MTU: %d\n", iface.MTU)
    fmt.Printf("标志: %v\n", iface.Flags)

    for _, addr := range iface.Addrs {
        fmt.Printf("地址: %s\n", addr)
    }
    fmt.Println("---")
}

// 获取网络统计
netStats := runtime.GetNetworkStats()
fmt.Printf("发送字节: %d\n", netStats.BytesSent)
fmt.Printf("接收字节: %d\n", netStats.BytesRecv)
fmt.Printf("发送包: %d\n", netStats.PacketsSent)
fmt.Printf("接收包: %d\n", netStats.PacketsRecv)
```

## API 参考

### 系统信息类型

```go
type SystemInfo struct {
    OS           string    // 操作系统
    Arch         string    // 架构
    Hostname     string    // 系统主机名
    Platform     string    // 平台详细信息
    CPUCores     int       // CPU 核心数
    TotalMemory  uint64    // 总系统内存（字节）
    Uptime       time.Duration // 系统运行时间
    BootTime     time.Time // 系统启动时间
}

type RuntimeInfo struct {
    GoVersion     string    // Go 版本
    NumGoroutines int       // Goroutines 数量
    NumCGOCalls   int64     // CGO 调用数
    GOMAXPROCS    int       // GOMAXPROCS 设置
    Compiler      string    // Go 编译器
}
```

### 内存信息类型

```go
type MemoryStats struct {
    Alloc        uint64    // 已分配内存
    TotalAlloc   uint64    // 总分配
    Sys          uint64    // 系统内存
    Lookups      uint64    // 指针查找
    Mallocs      uint64    // Malloc 调用
    Frees        uint64    // Free 调用
    NumGC        uint32    // GC 周期
    LastGC       time.Time // 上次 GC 时间
    PauseTotal   time.Duration // 总 GC 暂停时间
}

type HeapInfo struct {
    HeapAlloc    uint64    // 堆已分配内存
    HeapSys      uint64    // 堆系统内存
    HeapIdle     uint64    // 堆空闲内存
    HeapInuse    uint64    // 堆使用中内存
    HeapReleased uint64    // 释放给操作系统
    HeapObjects  uint64    // 对象数量
    NextGC       uint64    // 下次 GC 目标
}
```

### 进程信息类型

```go
type ProcessInfo struct {
    PID          int       // 进程 ID
    PPID         int       // 父进程 ID
    Executable   string    // 可执行文件路径
    WorkingDir   string    // 工作目录
    StartTime    time.Time // 进程启动时间
    CPUPercent   float64   // CPU 使用百分比
    MemoryUsage  uint64    // 内存使用（字节）
    OpenFiles    int       // 打开文件数
    Threads      int       // 线程数
}
```

### CPU 信息类型

```go
type CPUInfo struct {
    ModelName   string    // CPU 型号名称
    Cores       int       // 核心数
    Threads     int       // 线程数
    MHz         float64   // CPU 频率
    CacheSize   int32     // 缓存大小（KB）
    VendorID    string    // 供应商 ID
    Family      string    // CPU 系列
}

type CPUUsage struct {
    Total       float64   // 总 CPU 使用率
    User        float64   // 用户时间百分比
    System      float64   // 系统时间百分比
    Idle        float64   // 空闲时间百分比
}
```

### 函数

#### 系统信息
- `GetSystemInfo() SystemInfo` - 获取一般系统信息
- `GetRuntimeInfo() RuntimeInfo` - 获取 Go 运行时信息
- `GetHostname() string` - 获取系统主机名
- `GetUptime() time.Duration` - 获取系统运行时间

#### 内存信息
- `GetMemoryStats() MemoryStats` - 获取内存统计
- `GetHeapInfo() HeapInfo` - 获取堆信息
- `ForceGC()` - 强制垃圾回收
- `ReadMemStats() *runtime.MemStats` - 获取原始内存统计

#### 进程信息
- `GetProcessInfo() ProcessInfo` - 获取当前进程信息
- `GetPID() int` - 获取进程 ID
- `GetPPID() int` - 获取父进程 ID
- `GetExecutablePath() string` - 获取可执行文件路径

#### CPU 信息
- `GetCPUInfo() []CPUInfo` - 获取 CPU 信息
- `GetCPUUsage() CPUUsage` - 获取 CPU 使用统计
- `GetCPUCount() int` - 获取 CPU 数量
- `GetLoadAverage() (float64, float64, float64)` - 获取系统负载平均值

#### 磁盘信息
- `GetDiskInfo() []DiskInfo` - 获取磁盘使用信息
- `GetDiskUsage(path string) DiskUsage` - 获取特定路径的使用情况

#### 网络信息
- `GetNetworkInterfaces() []NetworkInterface` - 获取网络接口
- `GetNetworkStats() NetworkStats` - 获取网络统计

## 高级使用示例

### 系统监控

```go
// 创建系统监控器
monitor := runtime.NewSystemMonitor()

// 开始监控
monitor.Start(5 * time.Second) // 每 5 秒更新一次

// 获取持续更新
for update := range monitor.Updates() {
    fmt.Printf("CPU: %.2f%%, 内存: %d MB, Goroutines: %d\n",
        update.CPU.Total,
        update.Memory.Alloc/1024/1024,
        update.Runtime.NumGoroutines)
}
```

### 资源使用跟踪

```go
// 跟踪一段时间内的资源使用
tracker := runtime.NewResourceTracker()

// 开始跟踪
tracker.Start()

// 执行一些工作
doSomeWork()

// 获取使用报告
report := tracker.GetReport()
fmt.Printf("峰值内存: %d MB\n", report.PeakMemory/1024/1024)
fmt.Printf("平均 CPU: %.2f%%\n", report.AverageCPU)
fmt.Printf("持续时间: %v\n", report.Duration)
```

### 性能分析

```go
// 开始性能分析
profiler := runtime.NewProfiler()
profiler.Start()

// 执行操作
performOperations()

// 获取性能分析数据
profile := profiler.Stop()
fmt.Printf("CPU 性能分析: %s\n", profile.CPUProfile)
fmt.Printf("内存性能分析: %s\n", profile.MemoryProfile)
fmt.Printf("Goroutine 性能分析: %s\n", profile.GoroutineProfile)
```

## 最佳实践

1. **定期监控**: 使用定期监控而不是连续轮询
2. **资源清理**: 完成后始终停止监控器和跟踪器
3. **错误处理**: 当系统信息不可用时处理错误
4. **跨平台**: 在所有目标平台上测试兼容性
5. **性能影响**: 注意收集系统信息有开销

## 性能考虑

- **缓存**: 系统信息被缓存以减少开销
- **延迟加载**: 信息仅在请求时加载
- **高效 API**: 在可能的情况下使用高效的系统调用
- **最小分配**: 对频繁调用的函数最小化内存分配

## 平台特定说明

### Linux
- 使用 `/proc` 文件系统获取详细系统信息
- 支持容器环境的 cgroups
- 来自 `/proc/net/dev` 的网络统计

### macOS
- 使用 `sysctl` 获取系统信息
- Activity Monitor 集成
- 详细指标的 Core Foundation API

### Windows
- 使用 WMI（Windows 管理工具）
- 指标的性能计数器
- 进程信息的 Windows API

## 错误处理

```go
// 安全错误处理
sysInfo, err := runtime.GetSystemInfoSafe()
if err != nil {
    log.Printf("获取系统信息失败: %v", err)
    // 使用默认值或回退
}

// 检查是否在容器中运行
if runtime.IsContainer() {
    // 为容器化环境调整行为
}
```

## 相关包

- `app` - 应用程序生命周期管理
- `network` - 网络工具和信息
- `osx` - 操作系统特定的文件和进程操作