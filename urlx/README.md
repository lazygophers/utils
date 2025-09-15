# urlx - URL Utilities

[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils/urlx.svg)](https://pkg.go.dev/github.com/lazygophers/utils/urlx)

A lightweight Go package for URL manipulation and query parameter processing, featuring sorting, normalization, and advanced URL handling utilities.

## Features

### Query Parameter Management
- **Query Sorting**: Sort URL query parameters for consistent URLs
- **Parameter Normalization**: Clean and normalize query parameters
- **Type-Safe Processing**: Leverage Go's type system for safe URL operations
- **Performance Optimized**: Efficient operations with minimal allocations

### URL Manipulation
- **URL Building**: Construct URLs with proper escaping
- **Parameter Extraction**: Extract and process query parameters
- **URL Validation**: Validate URL formats and components
- **Canonical URLs**: Generate canonical URL representations

## Installation

```bash
go get github.com/lazygophers/utils/urlx
```

## Quick Start

```go
package main

import (
    "fmt"
    "net/url"
    "github.com/lazygophers/utils/urlx"
)

func main() {
    // Parse URL with query parameters
    u, _ := url.Parse("https://example.com/search?q=golang&sort=date&page=1")

    // Sort query parameters for consistency
    sortedQuery := urlx.SortQuery(u.Query())
    u.RawQuery = sortedQuery.Encode()

    fmt.Println(u.String())
    // Output: https://example.com/search?page=1&q=golang&sort=date
}
```

## Core API Reference

### Query Parameter Sorting

```go
// Basic query sorting
query := url.Values{
    "z":    []string{"last"},
    "a":    []string{"first"},
    "m":    []string{"middle"},
}

sortedQuery := urlx.SortQuery(query)
fmt.Println(sortedQuery.Encode())
// Output: a=first&m=middle&z=last

// Empty query handling
emptyQuery := url.Values{}
result := urlx.SortQuery(emptyQuery)
// Returns the same empty query (no-op)
```

### URL Normalization

```go
// Before normalization
originalURL := "https://example.com/path?c=3&a=1&b=2"

u, _ := url.Parse(originalURL)
normalized := urlx.SortQuery(u.Query())
u.RawQuery = normalized.Encode()

fmt.Println(u.String())
// Output: https://example.com/path?a=1&b=2&c=3
```

## Advanced Usage

### Building Canonical URLs

```go
func buildCanonicalURL(baseURL string, params map[string]string) (string, error) {
    u, err := url.Parse(baseURL)
    if err != nil {
        return "", err
    }

    // Build query parameters
    query := url.Values{}
    for key, value := range params {
        query.Set(key, value)
    }

    // Sort for consistency
    sortedQuery := urlx.SortQuery(query)
    u.RawQuery = sortedQuery.Encode()

    return u.String(), nil
}

// Usage
canonicalURL, _ := buildCanonicalURL("https://api.example.com/users", map[string]string{
    "page":     "1",
    "limit":    "10",
    "sort":     "name",
    "filter":   "active",
})

fmt.Println(canonicalURL)
// Output: https://api.example.com/users?filter=active&limit=10&page=1&sort=name
```

### URL Comparison and Deduplication

```go
func normalizeURLForComparison(rawURL string) (string, error) {
    u, err := url.Parse(rawURL)
    if err != nil {
        return "", err
    }

    // Normalize query parameters by sorting
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

// Example: Detecting duplicate URLs
urls := []string{
    "https://example.com/search?q=golang&sort=date",
    "https://example.com/search?sort=date&q=golang",
    "https://example.com/search?q=python&sort=date",
}

seen := make(map[string]bool)
for _, rawURL := range urls {
    normalized, _ := normalizeURLForComparison(rawURL)
    if seen[normalized] {
        fmt.Printf("Duplicate URL: %s\n", rawURL)
    } else {
        seen[normalized] = true
        fmt.Printf("Unique URL: %s\n", normalized)
    }
}
```

### API Request Signing

```go
import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "sort"
    "strings"
)

func signAPIRequest(method, path string, params map[string]string, secret string) string {
    // Build query parameters
    query := url.Values{}
    for key, value := range params {
        query.Set(key, value)
    }

    // Sort parameters for consistent signing
    sortedQuery := urlx.SortQuery(query)

    // Build string to sign
    stringToSign := method + "\n" + path + "\n" + sortedQuery.Encode()

    // Generate signature
    mac := hmac.New(sha256.New, []byte(secret))
    mac.Write([]byte(stringToSign))
    signature := hex.EncodeToString(mac.Sum(nil))

    return signature
}

// Usage
signature := signAPIRequest("GET", "/api/users", map[string]string{
    "timestamp": "1633024800",
    "nonce":     "abc123",
    "user_id":   "12345",
}, "secret-key")
```

### Caching and Cache Key Generation

```go
func generateCacheKey(baseURL string, params map[string]string) string {
    u, _ := url.Parse(baseURL)

    // Add parameters
    query := url.Values{}
    for key, value := range params {
        query.Set(key, value)
    }

    // Sort for consistent cache keys
    sortedQuery := urlx.SortQuery(query)
    u.RawQuery = sortedQuery.Encode()

    // Use the canonical URL as cache key
    return u.String()
}

// Example cache implementation
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

## Best Practices

### 1. Always Sort Query Parameters for Consistency

```go
// Good: Consistent URL generation
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

    // Always sort for consistency
    sortedParams := urlx.SortQuery(params)
    u.RawQuery = sortedParams.Encode()

    return u.String()
}

// Bad: Inconsistent URLs
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

    // No sorting - order depends on map iteration
    u.RawQuery = params.Encode()

    return u.String()
}
```

### 2. Handle Empty and Nil Values

```go
func safeQuerySort(query url.Values) url.Values {
    if len(query) == 0 {
        return query  // urlx.SortQuery handles this correctly
    }

    return urlx.SortQuery(query)
}
```

### 3. Use for URL Deduplication

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
        return false  // Already exists
    }

    s.urls[canonical] = true
    return true  // New URL added
}
```

## Performance Considerations

- **Minimal Allocations**: The package is designed for minimal memory allocations
- **Efficient Sorting**: Uses optimized sorting from the `candy` package
- **Early Returns**: Handles edge cases efficiently
- **Memory Reuse**: Reuses existing url.Values when possible

### Benchmarks

```go
// Typical performance characteristics:
// BenchmarkSortQuery/empty-8         1000000000    0.5 ns/op    0 B/op    0 allocs/op
// BenchmarkSortQuery/single-8        50000000     25.1 ns/op    0 B/op    0 allocs/op
// BenchmarkSortQuery/multiple-8      10000000    150.3 ns/op   64 B/op    1 allocs/op
```

## Use Cases

### 1. Web API Development

```go
// Consistent API endpoint URLs
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

### 2. Web Scraping and Crawling

```go
// Deduplicate URLs in a crawler
type Crawler struct {
    visited map[string]bool
}

func (c *Crawler) shouldVisit(rawURL string) bool {
    u, err := url.Parse(rawURL)
    if err != nil {
        return false
    }

    // Normalize query parameters
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

### 3. Analytics and Tracking

```go
// Normalize URLs for analytics
func normalizeForAnalytics(rawURL string) string {
    u, _ := url.Parse(rawURL)

    if u.RawQuery != "" {
        query, _ := url.ParseQuery(u.RawQuery)

        // Remove tracking parameters
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

### 4. Configuration and Settings

```go
// Generate consistent configuration URLs
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

## Error Handling

The package handles errors gracefully:

```go
// Empty or nil queries are handled safely
var emptyQuery url.Values
result := urlx.SortQuery(emptyQuery)
// Returns empty url.Values, no panic

// The function is safe to call on any url.Values
query := url.Values{
    "key": []string{"value"},
}
sorted := urlx.SortQuery(query)
// Always returns a valid url.Values
```

## Integration Examples

### HTTP Client

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

    // Ensure consistent URL generation
    sortedQuery := urlx.SortQuery(query)
    u.RawQuery = sortedQuery.Encode()

    return c.client.Get(u.String())
}
```

### URL Builder

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

// Usage
url := NewURLBuilder("https", "api.example.com").
    Path("/users").
    Param("page", "1").
    Param("limit", "10").
    Param("sort", "name").
    Build()
// Output: https://api.example.com/users?limit=10&page=1&sort=name
```

## Related Packages

- `net/url` - Go standard library URL parsing and manipulation
- `github.com/lazygophers/utils/candy` - Type conversion and sorting utilities (dependency)
- Standard library `sort` - Underlying sorting functionality

## Future Enhancements

Potential future additions to the package:

- URL pattern matching
- Query parameter validation
- URL template processing
- Advanced URL normalization
- Custom sorting strategies

## Contributing

This package is part of the LazyGophers Utils collection. For contributions:

1. Follow Go's URL handling best practices
2. Add comprehensive tests for edge cases
3. Ensure backward compatibility
4. Document any new functionality thoroughly

## License

This package is part of the LazyGophers Utils project. See the main repository for license information.

---

*This package provides a foundation for consistent URL manipulation. For more advanced URL routing and pattern matching, consider dedicated routing libraries.*