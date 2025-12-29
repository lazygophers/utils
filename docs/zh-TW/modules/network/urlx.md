---
title: urlx - URL 操作
---

# urlx - URL 操作

## 概述

urlx 模組提供 URL 操作工具，包括查詢參數排序和 URL 構建。

## 函數

### SortQuery()

排序 URL 查詢參數。

```go
func SortQuery(query url.Values) url.Values
```

**參數：**
- `query` - URL 查詢參數

**返回值：**
- 排序後的查詢參數

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

### 查詢參數排序

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

### API 請求構建

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
    
    // 處理響應
}
```

### URL 緩存

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
    fmt.Printf("緩存鍵: %s\n", cacheKey)
}
```

---

## 最佳實踐

### 一致的查詢排序

```go
// 好：使用排序查詢以獲得一致的 URL
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

// 避免：未排序的查詢參數
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

## 相關文檔

- [stringx](/zh-TW/modules/stringx) - 字符串工具
- [network](/zh-TW/modules/network) - 網絡工具
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
