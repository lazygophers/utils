---
title: network - 网络工具
---

# network - 网络工具

## 概述

network 模块提供网络接口发现和 IP 地址检索工具，支持 IPv4 和 IPv6 双栈。

## 函数

### GetInterfaceIpByName()

按接口名称获取 IP 地址。

```go
func GetInterfaceIpByName(ifaceName string) (string, error)
```

**参数：**
- `ifaceName` - 网络接口名称（例如 "eth0"、"en0"）

**返回值：**
- 接口的 IP 地址
- 如果接口不存在或发生错误，返回错误

**示例：**
```go
ip, err := network.GetInterfaceIpByName("eth0")
if err != nil {
    log.Errorf("获取 IP 失败: %v", err)
} else {
    fmt.Printf("IP 地址: %s\n", ip)
}
```

---

### GetInterfaceIpByAddrs()

按地址列表获取接口 IP。

```go
func GetInterfaceIpByAddrs(addrs []net.Addr) (string, error)
```

**参数：**
- `addrs` - 网络地址列表

**返回值：**
- 第一个有效 IP 地址
- 如果未找到有效 IP，返回错误

**示例：**
```go
interfaces, err := net.Interfaces()
if err != nil {
    log.Fatal(err)
}

for _, iface := range interfaces {
    addrs, err := iface.Addrs()
    if err != nil {
        continue
    }
    
    ip, err := network.GetInterfaceIpByAddrs(addrs)
    if err == nil {
        fmt.Printf("接口 %s: %s\n", iface.Name, ip)
    }
}
```

---

### GetListenIp()

获取监听 IP 地址。

```go
func GetListenIp() (string, error)
```

**返回值：**
- 适合监听的 IP 地址
- 如果未找到有效 IP，返回错误

**行为：**
- 优先返回非回环地址
- 如果没有非回环地址，返回回环地址
- 支持 IPv4 和 IPv6

**示例：**
```go
ip, err := network.GetListenIp()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("监听 IP: %s\n", ip)
server := &http.Server{
    Addr: ip + ":8080",
}
```

---

## 使用模式

### 接口发现

```go
func listInterfaces() error {
    interfaces, err := net.Interfaces()
    if err != nil {
        return err
    }
    
    for _, iface := range interfaces {
        addrs, err := iface.Addrs()
        if err != nil {
            continue
        }
        
        ip, err := network.GetInterfaceIpByAddrs(addrs)
        if err == nil {
            fmt.Printf("%s: %s\n", iface.Name, ip)
        }
    }
    
    return nil
}
```

### 服务器绑定

```go
func startServer() error {
    ip, err := network.GetListenIp()
    if err != nil {
        return err
    }
    
    addr := ip + ":8080"
    listener, err := net.Listen("tcp", addr)
    if err != nil {
        return err
    }
    
    log.Infof("监听在 %s", addr)
    
    // 接受连接
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Errorf("接受连接失败: %v", err)
            continue
        }
        
        go handleConnection(conn)
    }
}
```

### 特定接口

```go
func getInterfaceIP(ifaceName string) (string, error) {
    ip, err := network.GetInterfaceIpByName(ifaceName)
    if err != nil {
        return "", fmt.Errorf("接口 %s 未找到: %w", ifaceName, err)
    }
    
    return ip, nil
}

func main() {
    ip, err := getInterfaceIP("eth0")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("eth0 IP: %s\n", ip)
}
```

### 双栈支持

```go
func getIPv4Address() (string, error) {
    ip, err := network.GetListenIp()
    if err != nil {
        return "", err
    }
    
    // 检查是否为 IPv4
    parsedIP := net.ParseIP(ip)
    if parsedIP == nil {
        return "", fmt.Errorf("无效的 IP 地址")
    }
    
    if parsedIP.To4() == nil {
        return "", fmt.Errorf("不是 IPv4 地址")
    }
    
    return ip, nil
}
```

---

## 最佳实践

### 错误处理

```go
// 好：优雅地处理接口错误
func getSafeInterfaceIP(ifaceName string) string {
    ip, err := network.GetInterfaceIpByName(ifaceName)
    if err != nil {
        log.Warnf("接口 %s 未找到，使用默认值: %v", ifaceName, err)
        return "0.0.0.0"
    }
    
    return ip
}

// 好：验证 IP 地址
func validateIP(ip string) error {
    parsedIP := net.ParseIP(ip)
    if parsedIP == nil {
        return fmt.Errorf("无效的 IP 地址: %s", ip)
    }
    return nil
}
```

### 接口选择

```go
// 好：选择适当的接口
func selectInterface() (string, error) {
    // 尝试特定接口
    if ip, err := network.GetInterfaceIpByName("eth0"); err == nil {
        return ip, nil
    }
    
    // 回退到任何可用接口
    return network.GetListenIp()
}

// 好：按优先级选择接口
func selectInterfaceByPriority() (string, error) {
    interfaces := []string{"eth0", "en0", "wlan0"}
    
    for _, iface := range interfaces {
        if ip, err := network.GetInterfaceIpByName(iface); err == nil {
            return ip, nil
        }
    }
    
    return network.GetListenIp()
}
```

---

## 相关文档

- [urlx](/zh-CN/modules/urlx) - URL 操作
- [cryptox](/zh-CN/modules/cryptox) - 加密函数
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
