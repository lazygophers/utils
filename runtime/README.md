# runtime - Runtime System Information

The `runtime` package provides utilities for gathering system and runtime information about the Go application and the environment it's running in. It offers comprehensive system metrics, memory statistics, and runtime diagnostics.

## Features

- **System Information**: CPU, memory, disk, and network information
- **Runtime Metrics**: Go runtime statistics and memory usage
- **Process Information**: Current process details and resource usage
- **Environment Detection**: Operating system and architecture detection
- **Performance Monitoring**: Runtime performance metrics and profiling data
- **Cross-Platform Support**: Works on Linux, macOS, and Windows

## Installation

```bash
go get github.com/lazygophers/utils/runtime
```

## Usage Examples

### Basic System Information

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/runtime"
)

func main() {
    // Get system information
    sysInfo := runtime.GetSystemInfo()
    fmt.Printf("OS: %s\n", sysInfo.OS)
    fmt.Printf("Architecture: %s\n", sysInfo.Arch)
    fmt.Printf("CPU Cores: %d\n", sysInfo.CPUCores)
    fmt.Printf("Total Memory: %d MB\n", sysInfo.TotalMemory/1024/1024)

    // Get runtime information
    runtimeInfo := runtime.GetRuntimeInfo()
    fmt.Printf("Go Version: %s\n", runtimeInfo.GoVersion)
    fmt.Printf("Goroutines: %d\n", runtimeInfo.NumGoroutines)
    fmt.Printf("CGO Calls: %d\n", runtimeInfo.NumCGOCalls)
}
```

### Memory Statistics

```go
// Get detailed memory statistics
memStats := runtime.GetMemoryStats()
fmt.Printf("Allocated Memory: %d KB\n", memStats.Alloc/1024)
fmt.Printf("Total Allocations: %d\n", memStats.TotalAlloc)
fmt.Printf("System Memory: %d KB\n", memStats.Sys/1024)
fmt.Printf("GC Cycles: %d\n", memStats.NumGC)
fmt.Printf("Last GC: %v\n", memStats.LastGC)

// Get heap information
heapInfo := runtime.GetHeapInfo()
fmt.Printf("Heap Size: %d KB\n", heapInfo.HeapAlloc/1024)
fmt.Printf("Heap Objects: %d\n", heapInfo.HeapObjects)
fmt.Printf("Next GC: %d KB\n", heapInfo.NextGC/1024)
```

### Process Information

```go
// Get process information
processInfo := runtime.GetProcessInfo()
fmt.Printf("PID: %d\n", processInfo.PID)
fmt.Printf("PPID: %d\n", processInfo.PPID)
fmt.Printf("Executable: %s\n", processInfo.Executable)
fmt.Printf("Working Directory: %s\n", processInfo.WorkingDir)
fmt.Printf("Start Time: %v\n", processInfo.StartTime)
fmt.Printf("CPU Usage: %.2f%%\n", processInfo.CPUPercent)
fmt.Printf("Memory Usage: %d KB\n", processInfo.MemoryUsage/1024)
```

### CPU Information

```go
// Get CPU information
cpuInfo := runtime.GetCPUInfo()
fmt.Printf("CPU Model: %s\n", cpuInfo.ModelName)
fmt.Printf("CPU Cores: %d\n", cpuInfo.Cores)
fmt.Printf("CPU Threads: %d\n", cpuInfo.Threads)
fmt.Printf("CPU Frequency: %.2f GHz\n", cpuInfo.MHz/1000)
fmt.Printf("Cache Size: %d KB\n", cpuInfo.CacheSize)

// Get CPU usage
cpuUsage := runtime.GetCPUUsage()
fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage.Total)
fmt.Printf("User Time: %.2f%%\n", cpuUsage.User)
fmt.Printf("System Time: %.2f%%\n", cpuUsage.System)
```

### Disk Information

```go
// Get disk usage information
diskInfo := runtime.GetDiskInfo()
for _, disk := range diskInfo {
    fmt.Printf("Device: %s\n", disk.Device)
    fmt.Printf("Mount Point: %s\n", disk.Mountpoint)
    fmt.Printf("File System: %s\n", disk.Fstype)
    fmt.Printf("Total: %d GB\n", disk.Total/1024/1024/1024)
    fmt.Printf("Used: %d GB\n", disk.Used/1024/1024/1024)
    fmt.Printf("Free: %d GB\n", disk.Free/1024/1024/1024)
    fmt.Printf("Usage: %.2f%%\n", disk.UsedPercent)
    fmt.Println("---")
}
```

### Network Information

```go
// Get network interfaces
interfaces := runtime.GetNetworkInterfaces()
for _, iface := range interfaces {
    fmt.Printf("Interface: %s\n", iface.Name)
    fmt.Printf("MAC Address: %s\n", iface.HardwareAddr)
    fmt.Printf("MTU: %d\n", iface.MTU)
    fmt.Printf("Flags: %v\n", iface.Flags)

    for _, addr := range iface.Addrs {
        fmt.Printf("Address: %s\n", addr)
    }
    fmt.Println("---")
}

// Get network statistics
netStats := runtime.GetNetworkStats()
fmt.Printf("Bytes Sent: %d\n", netStats.BytesSent)
fmt.Printf("Bytes Received: %d\n", netStats.BytesRecv)
fmt.Printf("Packets Sent: %d\n", netStats.PacketsSent)
fmt.Printf("Packets Received: %d\n", netStats.PacketsRecv)
```

## API Reference

### System Information Types

```go
type SystemInfo struct {
    OS           string    // Operating system
    Arch         string    // Architecture
    Hostname     string    // System hostname
    Platform     string    // Platform details
    CPUCores     int       // Number of CPU cores
    TotalMemory  uint64    // Total system memory in bytes
    Uptime       time.Duration // System uptime
    BootTime     time.Time // System boot time
}

type RuntimeInfo struct {
    GoVersion     string    // Go version
    NumGoroutines int       // Number of goroutines
    NumCGOCalls   int64     // Number of CGO calls
    GOMAXPROCS    int       // GOMAXPROCS setting
    Compiler      string    // Go compiler
}
```

### Memory Information Types

```go
type MemoryStats struct {
    Alloc        uint64    // Allocated memory
    TotalAlloc   uint64    // Total allocations
    Sys          uint64    // System memory
    Lookups      uint64    // Pointer lookups
    Mallocs      uint64    // Malloc calls
    Frees        uint64    // Free calls
    NumGC        uint32    // GC cycles
    LastGC       time.Time // Last GC time
    PauseTotal   time.Duration // Total GC pause time
}

type HeapInfo struct {
    HeapAlloc    uint64    // Heap allocated memory
    HeapSys      uint64    // Heap system memory
    HeapIdle     uint64    // Heap idle memory
    HeapInuse    uint64    // Heap in-use memory
    HeapReleased uint64    // Released to OS
    HeapObjects  uint64    // Number of objects
    NextGC       uint64    // Next GC target
}
```

### Process Information Types

```go
type ProcessInfo struct {
    PID          int       // Process ID
    PPID         int       // Parent process ID
    Executable   string    // Executable path
    WorkingDir   string    // Working directory
    StartTime    time.Time // Process start time
    CPUPercent   float64   // CPU usage percentage
    MemoryUsage  uint64    // Memory usage in bytes
    OpenFiles    int       // Number of open files
    Threads      int       // Number of threads
}
```

### CPU Information Types

```go
type CPUInfo struct {
    ModelName   string    // CPU model name
    Cores       int       // Number of cores
    Threads     int       // Number of threads
    MHz         float64   // CPU frequency
    CacheSize   int32     // Cache size in KB
    VendorID    string    // Vendor ID
    Family      string    // CPU family
}

type CPUUsage struct {
    Total       float64   // Total CPU usage
    User        float64   // User time percentage
    System      float64   // System time percentage
    Idle        float64   // Idle time percentage
}
```

### Functions

#### System Information
- `GetSystemInfo() SystemInfo` - Get general system information
- `GetRuntimeInfo() RuntimeInfo` - Get Go runtime information
- `GetHostname() string` - Get system hostname
- `GetUptime() time.Duration` - Get system uptime

#### Memory Information
- `GetMemoryStats() MemoryStats` - Get memory statistics
- `GetHeapInfo() HeapInfo` - Get heap information
- `ForceGC()` - Force garbage collection
- `ReadMemStats() *runtime.MemStats` - Get raw memory stats

#### Process Information
- `GetProcessInfo() ProcessInfo` - Get current process information
- `GetPID() int` - Get process ID
- `GetPPID() int` - Get parent process ID
- `GetExecutablePath() string` - Get executable path

#### CPU Information
- `GetCPUInfo() []CPUInfo` - Get CPU information
- `GetCPUUsage() CPUUsage` - Get CPU usage statistics
- `GetCPUCount() int` - Get number of CPUs
- `GetLoadAverage() (float64, float64, float64)` - Get system load average

#### Disk Information
- `GetDiskInfo() []DiskInfo` - Get disk usage information
- `GetDiskUsage(path string) DiskUsage` - Get usage for specific path

#### Network Information
- `GetNetworkInterfaces() []NetworkInterface` - Get network interfaces
- `GetNetworkStats() NetworkStats` - Get network statistics

## Advanced Usage Examples

### System Monitoring

```go
// Create a system monitor
monitor := runtime.NewSystemMonitor()

// Start monitoring
monitor.Start(5 * time.Second) // Update every 5 seconds

// Get continuous updates
for update := range monitor.Updates() {
    fmt.Printf("CPU: %.2f%%, Memory: %d MB, Goroutines: %d\n",
        update.CPU.Total,
        update.Memory.Alloc/1024/1024,
        update.Runtime.NumGoroutines)
}
```

### Resource Usage Tracking

```go
// Track resource usage over time
tracker := runtime.NewResourceTracker()

// Start tracking
tracker.Start()

// Perform some work
doSomeWork()

// Get usage report
report := tracker.GetReport()
fmt.Printf("Peak Memory: %d MB\n", report.PeakMemory/1024/1024)
fmt.Printf("Average CPU: %.2f%%\n", report.AverageCPU)
fmt.Printf("Duration: %v\n", report.Duration)
```

### Performance Profiling

```go
// Start profiling
profiler := runtime.NewProfiler()
profiler.Start()

// Perform operations
performOperations()

// Get profile data
profile := profiler.Stop()
fmt.Printf("CPU Profile: %s\n", profile.CPUProfile)
fmt.Printf("Memory Profile: %s\n", profile.MemoryProfile)
fmt.Printf("Goroutine Profile: %s\n", profile.GoroutineProfile)
```

## Best Practices

1. **Periodic Monitoring**: Use periodic monitoring instead of continuous polling
2. **Resource Cleanup**: Always stop monitors and trackers when done
3. **Error Handling**: Handle errors when system information is unavailable
4. **Cross-Platform**: Test on all target platforms for compatibility
5. **Performance Impact**: Be aware that gathering system info has overhead

## Performance Considerations

- **Caching**: System information is cached to reduce overhead
- **Lazy Loading**: Information is loaded only when requested
- **Efficient APIs**: Uses efficient system calls where possible
- **Minimal Allocations**: Minimizes memory allocations for frequently called functions

## Platform-Specific Notes

### Linux
- Uses `/proc` filesystem for detailed system information
- Supports cgroups for container environments
- Network statistics from `/proc/net/dev`

### macOS
- Uses `sysctl` for system information
- Activity Monitor integration
- Core Foundation APIs for detailed metrics

### Windows
- Uses WMI (Windows Management Instrumentation)
- Performance counters for metrics
- Windows API for process information

## Error Handling

```go
// Safe error handling
sysInfo, err := runtime.GetSystemInfoSafe()
if err != nil {
    log.Printf("Failed to get system info: %v", err)
    // Use defaults or fallback
}

// Check if running in container
if runtime.IsContainer() {
    // Adjust behavior for containerized environment
}
```

## Related Packages

- `app` - Application lifecycle management
- `network` - Network utilities and information
- `osx` - OS-specific file and process operations