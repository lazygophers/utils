---
title: network - Network Utilities
---

# network - Network Utilities

## Overview

The network module provides network utilities for interface IP address detection and network operations.

## Functions

### GetInterfaceIpByName()

Get IP address by interface name.

```go
func GetInterfaceIpByName(name string, prev6 bool) string
```

**Parameters:**
- `name` - Interface name (e.g., "eth0", "en0")
- `prev6` - Prefer IPv6 if true

**Returns:**
- IP address string
- Empty string if not found

**Example:**
```go
ip := network.GetInterfaceIpByName("eth0", false)
fmt.Printf("Interface IP: %s\n", ip)
```

---

### GetInterfaceIpByAddrs()

Get IP address from interface addresses.

```go
func GetInterfaceIpByAddrs(address []net.Addr, prev6 bool) string
```

**Parameters:**
- `address` - Interface addresses
- `prev6` - Prefer IPv6 if true

**Returns:**
- IP address string
- Empty string if not found

**Example:**
```go
inter, err := net.InterfaceByName("eth0")
if err != nil {
    log.Fatal(err)
}

addrs, err := inter.Addrs()
if err != nil {
    log.Fatal(err)
}

ip := network.GetInterfaceIpByAddrs(addrs, false)
fmt.Printf("IP: %s\n", ip)
```

---

### GetListenIp()

Get listen IP address for network interfaces.

```go
func GetListenIp(prev6 ...bool) string
```

**Parameters:**
- `prev6` - Prefer IPv6 if true (optional)

**Returns:**
- IP address string
- Empty string if not found

**Search Order:**
1. eth0 interface
2. en0 interface
3. First available non-loopback interface

**Example:**
```go
ip := network.GetListenIp()
fmt.Printf("Listen IP: %s\n", ip)

ip6 := network.GetListenIp(true)
fmt.Printf("Listen IPv6: %s\n", ip6)
```

---

## Usage Patterns

### Server Configuration

```go
func startServer() {
    ip := network.GetListenIp()
    port := 8080
    
    addr := fmt.Sprintf("%s:%d", ip, port)
    listener, err := net.Listen("tcp", addr)
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    
    log.Infof("Server listening on %s", addr)
    // Server logic
}
```

### Interface Discovery

```go
func discoverInterfaces() {
    interfaces, err := net.Interfaces()
    if err != nil {
        log.Errorf("Failed to get interfaces: %v", err)
        return
    }
    
    for _, inter := range interfaces {
        ip := network.GetInterfaceIpByName(inter.Name, false)
        if ip != "" {
            log.Infof("Interface %s: %s", inter.Name, ip)
        }
    }
}
```

### Dual Stack Support

```go
func getServerAddresses() (string, string) {
    ipv4 := network.GetListenIp(false)
    ipv6 := network.GetListenIp(true)
    
    return ipv4, ipv6
}

func startDualStackServer() {
    ipv4, ipv6 := getServerAddresses()
    
    if ipv4 != "" {
        go startServer(ipv4, 8080)
    }
    
    if ipv6 != "" {
        go startServer(ipv6, 8080)
    }
}
```

### Service Discovery

```go
func findServicePort(service string) int {
    ip := network.GetListenIp()
    
    // Try common ports
    ports := []int{8080, 8081, 8082, 8083, 8084}
    
    for _, port := range ports {
        addr := fmt.Sprintf("%s:%d", ip, port)
        conn, err := net.DialTimeout("tcp", addr, time.Second)
        if err == nil {
            conn.Close()
            return port
        }
    }
    
    return 0
}
```

---

## Best Practices

### Error Handling

```go
// Good: Handle missing interfaces
func getInterfaceIP(name string) (string, error) {
    ip := network.GetInterfaceIpByName(name, false)
    if ip == "" {
        return "", fmt.Errorf("interface %s not found", name)
    }
    return ip, nil
}

// Good: Fallback to localhost
func getListenAddress() string {
    ip := network.GetListenIp()
    if ip == "" {
        log.Warn("No network interface found, using localhost")
        return "127.0.0.1"
    }
    return ip
}
```

### IPv6 Support

```go
// Good: Support both IPv4 and IPv6
func startServer() {
    ipv4 := network.GetListenIp(false)
    ipv6 := network.GetListenIp(true)
    
    if ipv4 != "" {
        go func() {
            addr := fmt.Sprintf("%s:8080", ipv4)
            listener, err := net.Listen("tcp", addr)
            if err != nil {
                log.Errorf("Failed to listen on IPv4: %v", err)
                return
            }
            log.Infof("Listening on IPv4: %s", addr)
            // Server logic
        }()
    }
    
    if ipv6 != "" {
        go func() {
            addr := fmt.Sprintf("[%s]:8080", ipv6)
            listener, err := net.Listen("tcp6", addr)
            if err != nil {
                log.Errorf("Failed to listen on IPv6: %v", err)
                return
            }
            log.Infof("Listening on IPv6: %s", addr)
            // Server logic
        }()
    }
}
```

---

## Related Documentation

- [runtime](/en/modules/runtime) - Runtime information
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
