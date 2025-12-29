---
title: urlx - URL Manipulation
---

# urlx - URL Manipulation

## Overview

The urlx module provides URL manipulation utilities including query parameter sorting and URL building.

## Functions

### SortQuery()

Sort URL query parameters.

```go
func SortQuery(query url.Values) url.Values
```

**Parameters:**
- `query` - URL query parameters

**Returns:**
- Sorted query parameters

**Example:**
```go
query := url.Values{}
query.Set("c", "3")
query.Set("a", "1")
query.Set("b", "2")

sorted := urlx.SortQuery(query)
// sorted is {"a": "1", "b": "2", "c": "3"}
```

---

## Usage Patterns

### Query Parameter Sorting

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

### API Request Building

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
    
    // Process response
}
```

### URL Caching

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
    fmt.Printf("Cache Key: %s\n", cacheKey)
}
```

---

## Best Practices

### Consistent Query Ordering

```go
// Good: Use sorted query for consistent URLs
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

// Avoid: Unsorted query parameters
func buildInconsistentURL(baseURL string, params map[string]string) string {
    query := url.Values{}
    for key, value := range params {
        query.Set(key, value)
    }
    
    u, _ := url.Parse(baseURL)
    u.RawQuery = query.Encode()  // Unsorted!
    
    return u.String()
}
```

---

## Related Documentation

- [stringx](/en/modules/stringx) - String utilities
- [network](/en/modules/network) - Network utilities
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
