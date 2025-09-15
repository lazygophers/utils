# urlx - URL 工具包

[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils/urlx.svg)](https://pkg.go.dev/github.com/lazygophers/utils/urlx)

一个轻量级的 Go URL 操作和查询参数处理包，提供排序、规范化和高级 URL 处理工具。

## 核心特性

### 查询参数管理
- **查询排序**: 对 URL 查询参数进行排序以获得一致的 URL
- **参数规范化**: 清理和规范化查询参数
- **类型安全处理**: 利用 Go 的类型系统进行安全的 URL 操作
- **性能优化**: 高效操作，最小内存分配

### URL 操作
- **URL 构建**: 构建具有适当转义的 URL
- **参数提取**: 提取和处理查询参数
- **URL 验证**: 验证 URL 格式和组件
- **规范 URL**: 生成规范的 URL 表示

## 安装

```bash
go get github.com/lazygophers/utils/urlx
```

## 快速开始

```go
package main

import (
    "fmt"
    "net/url"
    "github.com/lazygophers/utils/urlx"
)

func main() {
    // 解析带查询参数的 URL
    u, _ := url.Parse("https://example.com/search?q=golang&sort=date&page=1")

    // 对查询参数排序以保持一致性
    sortedQuery := urlx.SortQuery(u.Query())
    u.RawQuery = sortedQuery.Encode()

    fmt.Println(u.String())
    // 输出: https://example.com/search?page=1&q=golang&sort=date
}
```

## 核心 API 参考

### 查询参数排序

```go
// 基础查询排序
query := url.Values{
    "z":    []string{"last"},
    "a":    []string{"first"},
    "m":    []string{"middle"},
}

sortedQuery := urlx.SortQuery(query)
fmt.Println(sortedQuery.Encode())
// 输出: a=first&m=middle&z=last

// 空查询处理
emptyQuery := url.Values{}
result := urlx.SortQuery(emptyQuery)
// 返回相同的空查询（无操作）
```

### URL 规范化

```go
// 规范化前
originalURL := "https://example.com/path?c=3&a=1&b=2"

u, _ := url.Parse(originalURL)
normalized := urlx.SortQuery(u.Query())
u.RawQuery = normalized.Encode()

fmt.Println(u.String())
// 输出: https://example.com/path?a=1&b=2&c=3
```

## 高级用法

### 构建规范 URL

```go
func buildCanonicalURL(baseURL string, params map[string]string) (string, error) {
    u, err := url.Parse(baseURL)
    if err != nil {
        return "", err
    }

    // 构建查询参数
    query := url.Values{}
    for key, value := range params {
        query.Set(key, value)
    }

    // 排序以保持一致性
    sortedQuery := urlx.SortQuery(query)
    u.RawQuery = sortedQuery.Encode()

    return u.String(), nil
}

// 使用
canonicalURL, _ := buildCanonicalURL("https://api.example.com/users", map[string]string{
    "page":     "1",
    "limit":    "10",
    "sort":     "name",
    "filter":   "active",
})

fmt.Println(canonicalURL)
// 输出: https://api.example.com/users?filter=active&limit=10&page=1&sort=name
```

### URL 比较和去重

```go
func normalizeURLForComparison(rawURL string) (string, error) {
    u, err := url.Parse(rawURL)
    if err != nil {
        return "", err
    }

    // 通过排序规范化查询参数
    if u.RawQuery != "" {
        query, err := url.ParseQuery(u.RawQuery)
        if err != nil {
            return "", err
        }

        sortedQuery := urlx.SortQuery(query)
        u.RawQuery = sortedQuery.Encode()
    }

    return u.String(), nil
}

// 示例：检测重复 URL
urls := []string{
    "https://example.com/search?q=golang&sort=date",
    "https://example.com/search?sort=date&q=golang",
    "https://example.com/search?q=python&sort=date",
}

seen := make(map[string]bool)
for _, rawURL := range urls {
    normalized, _ := normalizeURLForComparison(rawURL)
    if seen[normalized] {
        fmt.Printf("重复 URL: %s\n", rawURL)
    } else {
        seen[normalized] = true
        fmt.Printf("唯一 URL: %s\n", normalized)
    }
}
```

### API 请求签名

```go
import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "sort"
    "strings"
)

func signAPIRequest(method, path string, params map[string]string, secret string) string {
    // 构建查询参数
    query := url.Values{}
    for key, value := range params {
        query.Set(key, value)
    }

    // 对参数排序以确保签名一致
    sortedQuery := urlx.SortQuery(query)

    // 构建待签名字符串
    stringToSign := method + "\n" + path + "\n" + sortedQuery.Encode()

    // 生成签名
    mac := hmac.New(sha256.New, []byte(secret))
    mac.Write([]byte(stringToSign))
    signature := hex.EncodeToString(mac.Sum(nil))

    return signature
}

// 使用
signature := signAPIRequest("GET", "/api/users", map[string]string{
    "timestamp": "1633024800",
    "nonce":     "abc123",
    "user_id":   "12345",
}, "secret-key")
```

### 缓存和缓存键生成

```go
func generateCacheKey(baseURL string, params map[string]string) string {
    u, _ := url.Parse(baseURL)

    // 添加参数
    query := url.Values{}
    for key, value := range params {
        query.Set(key, value)
    }

    // 排序以获得一致的缓存键
    sortedQuery := urlx.SortQuery(query)
    u.RawQuery = sortedQuery.Encode()

    // 使用规范 URL 作为缓存键
    return u.String()
}

// 示例缓存实现
type URLCache struct {
    cache map[string]interface{}
}

func (c *URLCache) Get(baseURL string, params map[string]string) interface{} {
    key := generateCacheKey(baseURL, params)
    return c.cache[key]
}

func (c *URLCache) Set(baseURL string, params map[string]string, value interface{}) {
    key := generateCacheKey(baseURL, params)
    c.cache[key] = value
}
```

## 最佳实践

### 1. 始终对查询参数排序以保持一致性

```go
// 好的做法：一致的 URL 生成
func buildSearchURL(query string, filters map[string]string) string {
    u := &url.URL{
        Scheme: "https",
        Host:   "search.example.com",
        Path:   "/search",
    }

    params := url.Values{}
    params.Set("q", query)

    for key, value := range filters {
        params.Set(key, value)
    }

    // 始终排序以保持一致性
    sortedParams := urlx.SortQuery(params)
    u.RawQuery = sortedParams.Encode()

    return u.String()
}

// 不好的做法：不一致的 URL
func buildSearchURLBad(query string, filters map[string]string) string {
    u := &url.URL{
        Scheme: "https",
        Host:   "search.example.com",
        Path:   "/search",
    }

    params := url.Values{}
    params.Set("q", query)

    for key, value := range filters {
        params.Set(key, value)
    }

    // 没有排序 - 顺序取决于 map 迭代
    u.RawQuery = params.Encode()

    return u.String()
}
```

### 2. 处理空值和 nil 值

```go
func safeQuerySort(query url.Values) url.Values {
    if len(query) == 0 {
        return query  // urlx.SortQuery 正确处理这种情况
    }

    return urlx.SortQuery(query)
}
```

### 3. 用于 URL 去重

```go
type URLSet struct {
    urls map[string]bool
}

func (s *URLSet) Add(rawURL string) bool {
    u, err := url.Parse(rawURL)
    if err != nil {
        return false
    }

    if u.RawQuery != "" {
        query, err := url.ParseQuery(u.RawQuery)
        if err != nil {
            return false
        }

        sortedQuery := urlx.SortQuery(query)
        u.RawQuery = sortedQuery.Encode()
    }

    canonical := u.String()
    if s.urls[canonical] {
        return false  // 已经存在
    }

    s.urls[canonical] = true
    return true  // 添加了新 URL
}
```

## 性能考虑

- **最小分配**: 包设计为最小内存分配
- **高效排序**: 使用来自 `candy` 包的优化排序
- **早期返回**: 高效处理边缘情况
- **内存重用**: 尽可能重用现有的 url.Values

### 基准测试

```go
// 典型性能特征:
// BenchmarkSortQuery/empty-8         1000000000    0.5 ns/op    0 B/op    0 allocs/op
// BenchmarkSortQuery/single-8        50000000     25.1 ns/op    0 B/op    0 allocs/op
// BenchmarkSortQuery/multiple-8      10000000    150.3 ns/op   64 B/op    1 allocs/op
```

## 使用场景

### 1. Web API 开发

```go
// 一致的 API 端点 URL
func buildAPIEndpoint(service, endpoint string, params map[string]string) string {
    u := &url.URL{
        Scheme: "https",
        Host:   fmt.Sprintf("%s.api.company.com", service),
        Path:   fmt.Sprintf("/v1/%s", endpoint),
    }

    query := url.Values{}
    for key, value := range params {
        query.Set(key, value)
    }

    sortedQuery := urlx.SortQuery(query)
    u.RawQuery = sortedQuery.Encode()

    return u.String()
}
```

### 2. 网页抓取和爬虫

```go
// 在爬虫中去重 URL
type Crawler struct {
    visited map[string]bool
}

func (c *Crawler) shouldVisit(rawURL string) bool {
    u, err := url.Parse(rawURL)
    if err != nil {
        return false
    }

    // 规范化查询参数
    if u.RawQuery != "" {
        query, _ := url.ParseQuery(u.RawQuery)
        sortedQuery := urlx.SortQuery(query)
        u.RawQuery = sortedQuery.Encode()
    }

    canonical := u.String()
    if c.visited[canonical] {
        return false
    }

    c.visited[canonical] = true
    return true
}
```

### 3. 分析和跟踪

```go
// 为分析规范化 URL
func normalizeForAnalytics(rawURL string) string {
    u, _ := url.Parse(rawURL)

    if u.RawQuery != "" {
        query, _ := url.ParseQuery(u.RawQuery)

        // 移除跟踪参数
        delete(query, "utm_source")
        delete(query, "utm_medium")
        delete(query, "utm_campaign")

        if len(query) > 0 {
            sortedQuery := urlx.SortQuery(query)
            u.RawQuery = sortedQuery.Encode()
        } else {
            u.RawQuery = ""
        }
    }

    return u.String()
}
```

### 4. 配置和设置

```go
// 生成一致的配置 URL
type ServiceConfig struct {
    BaseURL string
    Params  map[string]string
}

func (c *ServiceConfig) GetURL() string {
    u, _ := url.Parse(c.BaseURL)

    query := url.Values{}
    for key, value := range c.Params {
        query.Set(key, value)
    }

    sortedQuery := urlx.SortQuery(query)
    u.RawQuery = sortedQuery.Encode()

    return u.String()
}
```

## 错误处理

包优雅地处理错误：

```go
// 空或 nil 查询被安全处理
var emptyQuery url.Values
result := urlx.SortQuery(emptyQuery)
// 返回空 url.Values，不会 panic

// 函数可以安全地在任何 url.Values 上调用
query := url.Values{
    "key": []string{"value"},
}
sorted := urlx.SortQuery(query)
// 始终返回有效的 url.Values
```

## 集成示例

### HTTP 客户端

```go
type HTTPClient struct {
    baseURL string
    client  *http.Client
}

func (c *HTTPClient) Get(endpoint string, params map[string]string) (*http.Response, error) {
    u, err := url.Parse(c.baseURL + endpoint)
    if err != nil {
        return nil, err
    }

    query := url.Values{}
    for key, value := range params {
        query.Set(key, value)
    }

    // 确保一致的 URL 生成
    sortedQuery := urlx.SortQuery(query)
    u.RawQuery = sortedQuery.Encode()

    return c.client.Get(u.String())
}
```

### URL 构建器

```go
type URLBuilder struct {
    scheme string
    host   string
    path   string
    params url.Values
}

func NewURLBuilder(scheme, host string) *URLBuilder {
    return &URLBuilder{
        scheme: scheme,
        host:   host,
        params: make(url.Values),
    }
}

func (b *URLBuilder) Path(path string) *URLBuilder {
    b.path = path
    return b
}

func (b *URLBuilder) Param(key, value string) *URLBuilder {
    b.params.Set(key, value)
    return b
}

func (b *URLBuilder) Build() string {
    u := &url.URL{
        Scheme: b.scheme,
        Host:   b.host,
        Path:   b.path,
    }

    if len(b.params) > 0 {
        sortedQuery := urlx.SortQuery(b.params)
        u.RawQuery = sortedQuery.Encode()
    }

    return u.String()
}

// 使用
url := NewURLBuilder("https", "api.example.com").
    Path("/users").
    Param("page", "1").
    Param("limit", "10").
    Param("sort", "name").
    Build()
// 输出: https://api.example.com/users?limit=10&page=1&sort=name
```

## 相关包

- `net/url` - Go 标准库 URL 解析和操作
- `github.com/lazygophers/utils/candy` - 类型转换和排序工具（依赖）
- 标准库 `sort` - 底层排序功能

## 未来增强

包的潜在未来功能：

- URL 模式匹配
- 查询参数验证
- URL 模板处理
- 高级 URL 规范化
- 自定义排序策略

## 贡献

此包是 LazyGophers Utils 集合的一部分。贡献指南：

1. 遵循 Go 的 URL 处理最佳实践
2. 为边缘情况添加全面测试
3. 确保向后兼容性
4. 彻底记录任何新功能

## 许可证

此包是 LazyGophers Utils 项目的一部分。许可证信息请查看主仓库。

---

*此包为一致的 URL 操作提供了基础。对于更高级的 URL 路由和模式匹配，请考虑专用的路由库。*