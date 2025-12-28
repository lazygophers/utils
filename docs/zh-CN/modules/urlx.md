---
title: urlx - URL 操作
---

# urlx - URL 操作

## 概述

urlx 模块提供 URL 操作工具，包括查询参数排序和 URL 构建。

## 函数

### SortQuery()

排序 URL 查询参数。

```go
func SortQuery(query url.Values) url.Values
```

**参数：**
- `query` - URL 查询参数

**返回值：**
- 排序后的查询参数

**示例：**
```go
query := url.Values{}
query.Set("c", "3")
query.Set("a", "1")
query.Set("b", "2")

sorted := urlx.SortQuery(query)
// sorted 是 {"a": "1", "b": "2", "c": "3"}
```

---

## 使用模式

### 查询参数排序

```go
func buildURL(baseURL string, params map[string]string) string {
    query := url.Values{}
    for key, value := range params {
        query.Set(key, value)
    }
    
    sortedQuery := urlx.SortQuery(query)
    
    u, _ := url.Parse(baseURL)
    u.RawQuery = sortedQuery.Encode()
    
    return u.String()
}

func main() {
    params := map[string]string{
        "c": "3",
        "a": "1",
        "b": "2",
    }
    
    url := buildURL("https://example.com/api", params)
    fmt.Printf("URL: %s\n", url)
    // URL: https://example.com/api?a=1&b=2&c=3
}
```

### API 请求构建

```go
func buildAPIRequest(endpoint string, params map[string]string) (*http.Request, error) {
    query := url.Values{}
    for key, value := range params {
        query.Set(key, value)
    }
    
    sortedQuery := urlx.SortQuery(query)
    
    u, err := url.Parse(endpoint)
    if err != nil {
        return nil, err
    }
    
    u.RawQuery = sortedQuery.Encode()
    
    return http.NewRequest("GET", u.String(), nil)
}

func main() {
    params := map[string]string{
        "page":  "1",
        "limit": "10",
        "sort":  "name",
    }
    
    req, err := buildAPIRequest("https://api.example.com/users", params)
    if err != nil {
        log.Fatal(err)
    }
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    
    // 处理响应
}
```

### URL 缓存

```go
func getCacheKey(baseURL string, params map[string]string) string {
    query := url.Values{}
    for key, value := range params {
        query.Set(key, value)
    }
    
    sortedQuery := urlx.SortQuery(query)
    return baseURL + "?" + sortedQuery.Encode()
}

func main() {
    params := map[string]string{
        "id":   "123",
        "type": "user",
    }
    
    cacheKey := getCacheKey("https://api.example.com/data", params)
    fmt.Printf("缓存键: %s\n", cacheKey)
}
```

---

## 最佳实践

### 一致的查询排序

```go
// 好：使用排序查询以获得一致的 URL
func buildConsistentURL(baseURL string, params map[string]string) string {
    query := url.Values{}
    for key, value := range params {
        query.Set(key, value)
    }
    
    sortedQuery := urlx.SortQuery(query)
    
    u, _ := url.Parse(baseURL)
    u.RawQuery = sortedQuery.Encode()
    
    return u.String()
}

// 避免：未排序的查询参数
func buildInconsistentURL(baseURL string, params map[string]string) string {
    query := url.Values{}
    for key, value := range params {
        query.Set(key, value)
    }
    
    u, _ := url.Parse(baseURL)
    u.RawQuery = query.Encode()  // 未排序！
    
    return u.String()
}
```

---

## 相关文档

- [stringx](/zh-CN/modules/stringx) - 字符串工具
- [network](/zh-CN/modules/network) - 网络工具
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
