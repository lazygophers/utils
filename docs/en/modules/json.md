---
title: json - JSON Processing
---

# json - JSON Processing

## Overview

The json module provides enhanced JSON handling with better error messages and platform-specific optimizations. It wraps the standard library JSON encoder/decoder with improved functionality.

## Functions

### Marshal()

Encode value to JSON.

```go
func Marshal(v any) ([]byte, error)
```

**Parameters:**
- `v` - Value to encode

**Returns:**
- JSON-encoded bytes
- Error if encoding fails

**Example:**
```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

user := User{Name: "John", Email: "john@example.com"}
data, err := json.Marshal(user)
if err != nil {
    log.Errorf("Failed to marshal: %v", err)
}
// data is []byte(`{"name":"John","email":"john@example.com"}`)
```

---

### Unmarshal()

Decode JSON data into value.

```go
func Unmarshal(data []byte, v any) error
```

**Parameters:**
- `data` - JSON-encoded bytes
- `v` - Destination value (must be a pointer)

**Returns:**
- Error if decoding fails

**Example:**
```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

data := []byte(`{"name":"John","email":"john@example.com"}`)
var user User
if err := json.Unmarshal(data, &user); err != nil {
    log.Errorf("Failed to unmarshal: %v", err)
}
```

---

### MarshalString()

Encode value to JSON string.

```go
func MarshalString(v any) (string, error)
```

**Parameters:**
- `v` - Value to encode

**Returns:**
- JSON-encoded string
- Error if encoding fails

**Example:**
```go
user := User{Name: "John", Email: "john@example.com"}
str, err := json.MarshalString(user)
if err != nil {
    log.Errorf("Failed to marshal: %v", err)
}
// str is `{"name":"John","email":"john@example.com"}`
```

---

### UnmarshalString()

Decode JSON string into value.

```go
func UnmarshalString(data string, v any) error
```

**Parameters:**
- `data` - JSON-encoded string
- `v` - Destination value (must be a pointer)

**Returns:**
- Error if decoding fails

**Example:**
```go
data := `{"name":"John","email":"john@example.com"}`
var user User
if err := json.UnmarshalString(data, &user); err != nil {
    log.Errorf("Failed to unmarshal: %v", err)
}
```

---

### NewEncoder()

Create a new JSON encoder.

```go
func NewEncoder(w io.Writer) *json.Encoder
```

**Parameters:**
- `w` - Writer to encode to

**Returns:**
- JSON encoder

**Example:**
```go
file, err := os.Create("users.json")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

encoder := json.NewEncoder(file)
encoder.Encode(User{Name: "John", Email: "john@example.com"})
encoder.Encode(User{Name: "Jane", Email: "jane@example.com"})
```

---

### NewDecoder()

Create a new JSON decoder.

```go
func NewDecoder(r io.Reader) *json.Decoder
```

**Parameters:**
- `r` - Reader to decode from

**Returns:**
- JSON decoder

**Example:**
```go
file, err := os.Open("users.json")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

decoder := json.NewDecoder(file)
for decoder.More() {
    var user User
    if err := decoder.Decode(&user); err != nil {
        log.Errorf("Failed to decode: %v", err)
        break
    }
    // Process user
}
```

---

## Usage Patterns

### HTTP API

```go
func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    if err := db.Create(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
```

### File I/O

```go
func saveUsers(users []User, path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()
    
    encoder := json.NewEncoder(file)
    for _, user := range users {
        if err := encoder.Encode(user); err != nil {
            return err
        }
    }
    return nil
}

func loadUsers(path string) ([]User, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    var users []User
    decoder := json.NewDecoder(file)
    for decoder.More() {
        var user User
        if err := decoder.Decode(&user); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}
```

### Configuration

```go
type Config struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Debug    bool   `json:"debug"`
    Database struct {
        Name     string `json:"name"`
        User     string `json:"user"`
        Password string `json:"password"`
    } `json:"database"`
}

func loadConfig(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    
    var cfg Config
    if err := json.Unmarshal(data, &cfg); err != nil {
        return nil, err
    }
    
    return &cfg, nil
}

func saveConfig(cfg *Config, path string) error {
    data, err := json.Marshal(cfg)
    if err != nil {
        return err
    }
    
    return os.WriteFile(path, data, 0644)
}
```

### Pretty Printing

```go
func prettyPrint(v interface{}) error {
    data, err := json.MarshalIndent(v, "", "  ")
    if err != nil {
        return err
    }
    fmt.Println(string(data))
    return nil
}

func prettyPrintToFile(v interface{}, path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()
    
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    return encoder.Encode(v)
}
```

---

## Advanced Features

### Custom Marshal/Unmarshal

```go
type Time struct {
    time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf(`"%s"`, t.Time.Format("2006-01-02"))), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
    str := string(data)
    if str == "null" {
        return nil
    }
    
    str = strings.Trim(str, `"`)
    parsed, err := time.Parse("2006-01-02", str)
    if err != nil {
        return err
    }
    
    t.Time = parsed
    return nil
}

type Event struct {
    Name string `json:"name"`
    Date Time  `json:"date"`
}
```

### Streaming Processing

```go
func processLargeJSON(reader io.Reader, processor func(User) error) error {
    decoder := json.NewDecoder(reader)
    
    for decoder.More() {
        var user User
        if err := decoder.Decode(&user); err != nil {
            return err
        }
        
        if err := processor(user); err != nil {
            return err
        }
    }
    
    return nil
}

func processUsersFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()
    
    return processLargeJSON(file, func(user User) error {
        fmt.Printf("Processing: %s\n", user.Name)
        return nil
    })
}
```

### Error Handling

```go
func safeUnmarshal(data []byte, v interface{}) error {
    if err := json.Unmarshal(data, v); err != nil {
        // Log detailed error information
        log.Errorf("JSON unmarshal failed: %v", err)
        log.Errorf("Data: %s", string(data))
        
        // Return wrapped error
        return fmt.Errorf("failed to unmarshal JSON: %w", err)
    }
    return nil
}

func validateJSON(data []byte) error {
    var v interface{}
    if err := json.Unmarshal(data, &v); err != nil {
        return fmt.Errorf("invalid JSON: %w", err)
    }
    return nil
}
```

---

## Best Practices

### Error Handling

```go
// Good: Handle errors properly
data, err := json.Marshal(user)
if err != nil {
    log.Errorf("Failed to marshal user: %v", err)
    return err
}

// Good: Validate JSON before unmarshaling
if !json.Valid(data) {
    return fmt.Errorf("invalid JSON data")
}

var user User
if err := json.Unmarshal(data, &user); err != nil {
    log.Errorf("Failed to unmarshal user: %v", err)
    return err
}
```

### Memory Efficiency

```go
// Good: Use streaming for large files
func processLargeFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()
    
    decoder := json.NewDecoder(file)
    for decoder.More() {
        var item interface{}
        if err := decoder.Decode(&item); err != nil {
            return err
        }
        // Process item
    }
    return nil
}

// Avoid: Loading entire file into memory
func processLargeFileBad(path string) error {
    data, err := os.ReadFile(path)  // Loads entire file
    if err != nil {
        return err
    }
    
    var items []interface{}
    if err := json.Unmarshal(data, &items); err != nil {
        return err
    }
    // Process items
    return nil
}
```

### Performance

```go
// Good: Reuse encoder/decoder
var encoder = json.NewEncoder(os.Stdout)
var decoder = json.NewDecoder(os.Stdin)

func process(data []byte) error {
    var v interface{}
    return decoder.Decode(&v)
}

// Avoid: Creating new encoder/decoder each time
func processBad(data []byte) error {
    decoder := json.NewDecoder(bytes.NewReader(data))  // Expensive
    var v interface{}
    return decoder.Decode(&v)
}
```

---

## Platform-Specific Optimizations

The json module automatically selects the best implementation based on the platform:

### Linux AMD64

Uses `sonic` JSON library for maximum performance:
- **3.5x faster** than standard library
- Zero-allocation optimizations
- SIMD-accelerated parsing

### Other Platforms

Uses standard library JSON with enhanced error messages:
- Compatible with all platforms
- Better error messages for debugging
- Standard library performance

---

## Related Documentation

- [candy](/en/modules/candy) - Type conversion
- [stringx](/en/modules/stringx) - String utilities
- [anyx](/en/modules/anyx) - Interface{} helpers
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
