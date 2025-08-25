# network

网络工具包，提供丰富的网络相关功能，包括 HTTP 客户端、WebSocket、TCP/UDP 操作等。

## 功能特性

- **HTTP 客户端**：支持 GET、POST、PUT、DELETE 等方法
- **WebSocket 客户端**：实时双向通信
- **TCP/UDP 操作**：底层网络协议支持
- **连接池管理**：复用连接，提升性能
- **请求拦截器**：统一处理请求/响应
- **超时控制**：防止请求阻塞
- **重试机制**：自动重试失败请求

## 安装

```bash
go get github.com/lazygophers/utils/network
```

## 快速开始

### HTTP 请求

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/lazygophers/utils/network"
)

func main() {
    // 创建 HTTP 客户端
    client := network.NewHTTPClient()
    
    // 发送 GET 请求
    resp, err := client.Get(context.Background(), "https://api.example.com/users")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    
    fmt.Printf("Status: %s\n", resp.Status)
}
```

### POST 请求

```go
func createUser() {
    client := network.NewHTTPClient()
    
    data := map[string]interface{}{
        "name":  "John Doe",
        "email": "john@example.com",
    }
    
    resp, err := client.PostJSON(
        context.Background(),
        "https://api.example.com/users",
        data,
    )
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    
    if resp.StatusCode == 201 {
        fmt.Println("User created successfully")
    }
}
```

### WebSocket 连接

```go
func handleWebSocket() {
    ws, err := network.NewWebSocket("ws://localhost:8080/ws")
    if err != nil {
        log.Fatal(err)
    }
    defer ws.Close()
    
    // 发送消息
    err = ws.WriteJSON(map[string]string{
        "type": "message",
        "data": "Hello WebSocket",
    })
    if err != nil {
        log.Fatal(err)
    }
    
    // 读取消息
    var msg map[string]interface{}
    err = ws.ReadJSON(&msg)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Received: %v\n", msg)
}
```

### TCP 服务器

```go
func startTCPServer() {
    server, err := network.NewTCPServer(":8080")
    if err != nil {
        log.Fatal(err)
    }
    
    server.HandleFunc(func(conn network.TCPConn) {
        defer conn.Close()
        
        // 读取数据
        buf := make([]byte, 1024)
        n, err := conn.Read(buf)
        if err != nil {
            log.Printf("Read error: %v", err)
            return
        }
        
        fmt.Printf("Received: %s\n", buf[:n])
        
        // 发送响应
        _, err = conn.Write([]byte("Hello from server"))
        if err != nil {
            log.Printf("Write error: %v", err)
        }
    })
    
    log.Fatal(server.Start())
}
```

## 高级用法

### 拦截器

```go
func setupInterceptors() {
    client := network.NewHTTPClient(
        network.WithRequestInterceptor(func(req *http.Request) error {
            req.Header.Set("Authorization", "Bearer token")
            return nil
        }),
        network.WithResponseInterceptor(func(resp *http.Response) error {
            log.Printf("Response status: %d", resp.StatusCode)
            return nil
        }),
    )
}
```

### 连接池配置

```go
func configurePool() {
    client := network.NewHTTPClient(
        network.WithMaxIdleConns(100),
        network.WithMaxIdleConnsPerHost(10),
        network.WithIdleConnTimeout(30 * time.Second),
    )
}
```

### 重试机制

```go
func withRetry() {
    client := network.NewHTTPClient(
        network.WithRetry(network.RetryConfig{
            MaxAttempts: 3,
            WaitTime:    time.Second,
            MaxWaitTime: 10 * time.Second,
            Retryable: func(resp *http.Response, err error) bool {
                return err != nil || resp.StatusCode >= 500
            },
        }),
    )
}
```

## API 文档

详细的 API 文档请访问：[GoDoc Reference](https://pkg.go.dev/github.com/lazygophers/utils/network)

## 许可证

本项目采用 AGPL-3.0 许可证。详见 [LICENSE](../LICENSE) 文件。