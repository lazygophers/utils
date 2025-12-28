---
title: network - 網絡工具
---

# network - 網絡工具

## 概述

network 模組提供網絡接口發現和 IP 地址檢索工具，支援 IPv4 和 IPv6 雙棧。

## 函數

### GetInterfaceIpByName()

按接口名稱獲取 IP 地址。

```go
func GetInterfaceIpByName(ifaceName string) (string, error)
```

**參數：**
- `ifaceName` - 網絡接口名稱（例如 "eth0"、"en0"）

**返回值：**
- 接口的 IP 地址
- 如果接口不存在或發生錯誤，返回錯誤

**示例：**
```go
ip, err := network.GetInterfaceIpByName("eth0")
if err != nil {
    log.Errorf("獲取 IP 失敗: %v", err)
} else {
    fmt.Printf("IP 地址: %s\n", ip)
}
```

---

### GetInterfaceIpByAddrs()

按地址列表獲取接口 IP。

```go
func GetInterfaceIpByAddrs(addrs []net.Addr) (string, error)
```

**參數：**
- `addrs` - 網絡地址列表

**返回值：**
- 第一個有效 IP 地址
- 如果未找到有效 IP，返回錯誤

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

獲取監聽 IP 地址。

```go
func GetListenIp() (string, error)
```

**返回值：**
- 適合監聽的 IP 地址
- 如果未找到有效 IP，返回錯誤

**行為：**
- 優先返回非回環地址
- 如果沒有非回環地址，返回回環地址
- 支援 IPv4 和 IPv6

**示例：**
```go
ip, err := network.GetListenIp()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("監聽 IP: %s\n", ip)
server := &http.Server{
    Addr: ip + ":8080",
}
```

---

## 使用模式

### 接口發現

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

### 服務器綁定

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
    
    log.Infof("監聽在 %s", addr)
    
    // 接受連接
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Errorf("接受連接失敗: %v", err)
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

### 雙棧支援

```go
func getIPv4Address() (string, error) {
    ip, err := network.GetListenIp()
    if err != nil {
        return "", err
    }
    
    // 檢查是否為 IPv4
    parsedIP := net.ParseIP(ip)
    if parsedIP == nil {
        return "", fmt.Errorf("無效的 IP 地址")
    }
    
    if parsedIP.To4() == nil {
        return "", fmt.Errorf("不是 IPv4 地址")
    }
    
    return ip, nil
}
```

---

## 最佳實踐

### 錯誤處理

```go
// 好：優雅地處理接口錯誤
func getSafeInterfaceIP(ifaceName string) string {
    ip, err := network.GetInterfaceIpByName(ifaceName)
    if err != nil {
        log.Warnf("接口 %s 未找到，使用默認值: %v", ifaceName, err)
        return "0.0.0.0"
    }
    
    return ip
}

// 好：驗證 IP 地址
func validateIP(ip string) error {
    parsedIP := net.ParseIP(ip)
    if parsedIP == nil {
        return fmt.Errorf("無效的 IP 地址: %s", ip)
    }
    return nil
}
```

### 接口選擇

```go
// 好：選擇適當的接口
func selectInterface() (string, error) {
    // 嘗試特定接口
    if ip, err := network.GetInterfaceIpByName("eth0"); err == nil {
        return ip, nil
    }
    
    // 回退到任何可用接口
    return network.GetListenIp()
}

// 好：按優先級選擇接口
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

## 相關文檔

- [urlx](/zh-TW/modules/urlx) - URL 操作
- [cryptox](/zh-TW/modules/cryptox) - 加密函數
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
