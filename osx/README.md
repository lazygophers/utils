# osx

macOS 系统专用工具包，提供访问 macOS 特有功能的便捷接口。

## 功能特性

- **系统信息获取**：获取 macOS 版本、硬件信息等
- **Finder 操作**：文件选择、桌面路径获取等
- **通知中心**：发送系统通知
- **剪贴板操作**：读取和写入剪贴板内容
- **系统偏好设置**：访问和修改系统设置
- **安全相关**：钥匙串访问、安全范围等
- **图形界面**：窗口管理、屏幕操作等

## 安装

```bash
go get github.com/lazygophers/utils/osx
```

## 快速开始

### 获取系统信息

```go
package main

import (
    "fmt"
    
    "github.com/lazygophers/utils/osx"
)

func main() {
    // 获取 macOS 版本
    version, err := osx.GetMacOSVersion()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("macOS Version: %s\n", version)
    
    // 获取硬件信息
    info, err := osx.GetHardwareInfo()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Model: %s\n", info.Model)
    fmt.Printf("Memory: %d GB\n", info.MemoryGB)
}
```

### 发送系统通知

```go
func sendNotification() {
    err := osx.SendNotification(osx.Notification{
        Title:    "应用通知",
        Subtitle: "来自 Go 应用",
        Body:     "这是一条测试通知消息",
        Sound:    true,
    })
    if err != nil {
        fmt.Printf("Failed to send notification: %v\n", err)
    }
}
```

### 剪贴板操作

```go
func clipboardExample() {
    // 写入文本到剪贴板
    err := osx.SetClipboardString("Hello from Go!")
    if err != nil {
        fmt.Printf("Error setting clipboard: %v\n", err)
        return
    }
    
    // 从剪贴板读取文本
    text, err := osx.GetClipboardString()
    if err != nil {
        fmt.Printf("Error getting clipboard: %v\n", err)
        return
    }
    fmt.Printf("Clipboard content: %s\n", text)
}
```

### Finder 操作

```go
func finderExample() {
    // 获取选中的文件
    selectedFiles, err := osx.GetFinderSelection()
    if err != nil {
        fmt.Printf("Error getting selection: %v\n", err)
        return
    }
    
    fmt.Println("Selected files:")
    for _, file := range selectedFiles {
        fmt.Printf("- %s\n", file)
    }
    
    // 获取桌面路径
    desktopPath, err := osx.GetDesktopPath()
    if err != nil {
        fmt.Printf("Error getting desktop path: %v\n", err)
        return
    }
    fmt.Printf("Desktop path: %s\n", desktopPath)
}
```

### 系统偏好设置

```go
func preferencesExample() {
    // 检查深色模式
    isDarkMode, err := osx.IsDarkMode()
    if err != nil {
        fmt.Printf("Error checking dark mode: %v\n", err)
        return
    }
    fmt.Printf("Dark mode enabled: %t\n", isDarkMode)
    
    // 获取屏幕亮度
    brightness, err := osx.GetDisplayBrightness()
    if err != nil {
        fmt.Printf("Error getting brightness: %v\n", err)
        return
    }
    fmt.Printf("Display brightness: %.2f\n", brightness)
}
```

## 高级用法

### 钥匙串访问

```go
func keychainExample() {
    // 添加密码到钥匙串
    err := osx.AddKeychainPassword("myapp.example.com", "username", "password123")
    if err != nil {
        fmt.Printf("Error adding password: %v\n", err)
        return
    }
    
    // 从钥匙串获取密码
    password, err := osx.GetKeychainPassword("myapp.example.com", "username")
    if err != nil {
        fmt.Printf("Error getting password: %v\n", err)
        return
    }
    fmt.Printf("Retrieved password: %s\n", password)
}
```

### 窗口管理

```go
func windowExample() {
    // 获取前端窗口信息
    window, err := osx.GetFrontmostWindow()
    if err != nil {
        fmt.Printf("Error getting frontmost window: %v\n", err)
        return
    }
    
    fmt.Printf("Frontmost app: %s\n", window.OwnerName)
    fmt.Printf("Window title: %s\n", window.Title)
    
    // 获取窗口尺寸
    bounds, err := osx.GetWindowBounds(window)
    if err != nil {
        fmt.Printf("Error getting window bounds: %v\n", err)
        return
    }
    fmt.Printf("Window size: %dx%d\n", bounds.Width, bounds.Height)
}
```

### 系统事件监听

```go
func eventListener() {
    // 监听系统睡眠事件
    sleepCh, err := osx.WatchSystemSleep()
    if err != nil {
        fmt.Printf("Error setting up sleep watcher: %v\n", err)
        return
    }
    
    go func() {
        for event := range sleepCh {
            fmt.Printf("System is about to sleep: %v\n", event)
        }
    }()
    
    // 监听应用切换事件
    appCh, err := osx.WatchAppSwitch()
    if err != nil {
        fmt.Printf("Error setting up app switch watcher: %v\n", err)
        return
    }
    
    go func() {
        for app := range appCh {
            fmt.Printf("Switched to app: %s\n", app)
        }
    }()
}
```

## 权限要求

某些功能需要特定的权限：

- **通知权限**：需要在系统偏好设置中允许应用发送通知
- **辅助功能权限**：窗口控制和模拟操作需要辅助功能权限
- **完全磁盘访问权限**：访问某些系统文件需要此权限

## API 文档

详细的 API 文档请访问：[GoDoc Reference](https://pkg.go.dev/github.com/lazygophers/utils/osx)

## 许可证

本项目采用 AGPL-3.0 许可证。详见 [LICENSE](../LICENSE) 文件。