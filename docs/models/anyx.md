# Anyx Package Documentation

<!-- Language selector -->
[🇺🇸 English](#english) | [🇨🇳 简体中文](#简体中文) | [🇭🇰 繁體中文](#繁體中文) | [🇷🇺 Русский](#русский) | [🇫🇷 Français](#français) | [🇸🇦 العربية](#العربية) | [🇪🇸 Español](#español)

---

## English

### Overview
The `anyx` package provides powerful utilities for working with any type values in Go, offering thread-safe map operations with dynamic type conversion and flexible data access patterns.

### Key Features
- **MapAny Structure**: Thread-safe map with dynamic typing support
- **Type Conversion**: Automatic conversion between different data types
- **Nested Access**: Support for nested key access with configurable separators
- **Multiple Format Support**: JSON and YAML integration
- **Thread Safety**: Built with `sync.Map` for concurrent access
- **Flexible Initialization**: Multiple construction methods

### Core Components

#### MapAny Structure
```go
type MapAny struct {
    data *sync.Map
    cut  *atomic.Bool
    seq  *atomic.String
}
```

#### Construction Methods
```go
// Create from map
m := anyx.NewMap(map[string]interface{}{
    "name": "John",
    "age":  30,
})

// Create from JSON
jsonData := []byte(`{"name": "Alice", "score": 95.5}`)
m, err := anyx.NewMapWithJson(jsonData)

// Create from YAML
yamlData := []byte(`name: Bob\nage: 25`)
m, err := anyx.NewMapWithYaml(yamlData)

// Create from any type
user := User{Name: "Charlie", Age: 35}
m, err := anyx.NewMapWithAny(user)
```

#### Data Access Methods
```go
// Basic get operations
value, err := m.Get("name")           // Get raw value
name := m.GetString("name")           // Get as string
age := m.GetInt("age")               // Get as int
score := m.GetFloat64("score")       // Get as float64
active := m.GetBool("active")        // Get as bool

// Check existence
exists := m.Exists("email")          // Check if key exists

// Collection access
items := m.GetSlice("items")         // Get as []interface{}
tags := m.GetStringSlice("tags")     // Get as []string
ids := m.GetUint64Slice("ids")       // Get as []uint64
```

#### Nested Access
```go
// Enable nested access with dot notation
m.EnableCut(".")

// Access nested values
data := map[string]interface{}{
    "user": map[string]interface{}{
        "profile": map[string]interface{}{
            "name": "Deep Value",
        },
    },
}
m := anyx.NewMap(data)
m.EnableCut(".")

name := m.GetString("user.profile.name") // Returns "Deep Value"
```

### Advanced Features

#### Custom Separator
```go
// Use custom separator for nested access
m.EnableCut("->")
value := m.GetString("level1->level2->key")

// Disable nested access
m.DisableCut()
```

#### Data Manipulation
```go
// Set values
m.Set("newKey", "newValue")

// Get nested map
nestedMap := m.GetMap("user")

// Convert to standard types
syncMap := m.ToSyncMap()
regularMap := m.ToMap()

// Clone the map
cloned := m.Clone()
```

#### Iteration
```go
// Range over all key-value pairs
m.Range(func(key, value interface{}) bool {
    fmt.Printf("Key: %v, Value: %v\n", key, value)
    return true // continue iteration
})
```

### Best Practices
1. **Enable Cut for Nested Access**: Use `EnableCut()` when working with nested data structures
2. **Type Safety**: Always handle the possibility of type conversion failures
3. **Error Handling**: Check errors when creating MapAny from external data
4. **Performance**: Use appropriate getter methods for known types
5. **Thread Safety**: MapAny is thread-safe, but external mutations should be synchronized

### Common Patterns
```go
// Configuration management
configData := `{
    "database": {
        "host": "localhost",
        "port": 5432,
        "credentials": {
            "username": "admin",
            "password": "secret"
        }
    }
}`
config, _ := anyx.NewMapWithJson([]byte(configData))
config.EnableCut(".")

dbHost := config.GetString("database.host")
dbPort := config.GetInt("database.port")
username := config.GetString("database.credentials.username")

// Dynamic data processing
func processUserData(data interface{}) {
    userMap, err := anyx.NewMapWithAny(data)
    if err != nil {
        return
    }
    
    if userMap.Exists("profile") {
        profile := userMap.GetMap("profile")
        name := profile.GetString("name")
        age := profile.GetInt("age")
        // Process profile data...
    }
}
```

---

## 简体中文

### 概述
`anyx` 包为 Go 提供处理任意类型值的强大工具，提供支持动态类型转换和灵活数据访问模式的线程安全映射操作。

### 主要特性
- **MapAny 结构**: 支持动态类型的线程安全映射
- **类型转换**: 不同数据类型间的自动转换
- **嵌套访问**: 支持可配置分隔符的嵌套键访问
- **多格式支持**: JSON 和 YAML 集成
- **线程安全**: 基于 `sync.Map` 构建，支持并发访问
- **灵活初始化**: 多种构造方法

### 核心组件

#### MapAny 结构
```go
type MapAny struct {
    data *sync.Map
    cut  *atomic.Bool
    seq  *atomic.String
}
```

#### 构造方法
```go
// 从 map 创建
m := anyx.NewMap(map[string]interface{}{
    "name": "张三",
    "age":  30,
})

// 从 JSON 创建
jsonData := []byte(`{"name": "李四", "score": 95.5}`)
m, err := anyx.NewMapWithJson(jsonData)

// 从 YAML 创建
yamlData := []byte(`name: 王五\nage: 25`)
m, err := anyx.NewMapWithYaml(yamlData)
```

#### 数据访问方法
```go
// 基本获取操作
value, err := m.Get("name")           // 获取原始值
name := m.GetString("name")           // 获取字符串
age := m.GetInt("age")               // 获取整数
score := m.GetFloat64("score")       // 获取浮点数
active := m.GetBool("active")        // 获取布尔值

// 检查存在性
exists := m.Exists("email")          // 检查键是否存在
```

### 最佳实践
1. **嵌套访问启用 Cut**: 处理嵌套数据结构时使用 `EnableCut()`
2. **类型安全**: 始终处理类型转换失败的可能性
3. **错误处理**: 从外部数据创建 MapAny 时检查错误
4. **性能考虑**: 为已知类型使用适当的获取器方法

---

## 繁體中文

### 概述
`anyx` 套件為 Go 提供處理任意型別值的強大工具，提供支援動態型別轉換和靈活資料存取模式的執行緒安全對應操作。

### 主要特性
- **MapAny 結構**: 支援動態型別的執行緒安全對應
- **型別轉換**: 不同資料型別間的自動轉換
- **巢狀存取**: 支援可設定分隔符的巢狀鍵存取
- **多格式支援**: JSON 和 YAML 整合
- **執行緒安全**: 基於 `sync.Map` 建構，支援並發存取

### 核心組件
```go
// 從 map 建立
m := anyx.NewMap(map[string]interface{}{
    "name": "張三",
    "age":  30,
})

// 從 JSON 建立
jsonData := []byte(`{"name": "李四", "score": 95.5}`)
m, err := anyx.NewMapWithJson(jsonData)
```

### 最佳實務
1. **巢狀存取啟用 Cut**: 處理巢狀資料結構時使用 `EnableCut()`
2. **型別安全**: 始終處理型別轉換失敗的可能性
3. **錯誤處理**: 從外部資料建立 MapAny 時檢查錯誤

---

## Русский

### Обзор
Пакет `anyx` предоставляет мощные утилиты для работы с значениями любого типа в Go, предлагая потокобезопасные операции с картами с поддержкой динамического преобразования типов и гибких шаблонов доступа к данным.

### Основные возможности
- **Структура MapAny**: Потокобезопасная карта с поддержкой динамических типов
- **Преобразование типов**: Автоматическое преобразование между различными типами данных
- **Вложенный доступ**: Поддержка вложенного доступа к ключам с настраиваемыми разделителями
- **Поддержка множественных форматов**: Интеграция JSON и YAML

### Основные компоненты
```go
// Создание из карты
m := anyx.NewMap(map[string]interface{}{
    "name": "Иван",
    "age":  30,
})

// Создание из JSON
jsonData := []byte(`{"name": "Алиса", "score": 95.5}`)
m, err := anyx.NewMapWithJson(jsonData)
```

### Лучшие практики
1. **Включите Cut для вложенного доступа**: Используйте `EnableCut()` при работе с вложенными структурами данных
2. **Безопасность типов**: Всегда обрабатывайте возможность неудачи преобразования типов

---

## Français

### Aperçu
Le package `anyx` fournit des utilitaires puissants pour travailler avec des valeurs de tout type en Go, offrant des opérations de carte thread-safe avec support de conversion de type dynamique et des modèles d'accès aux données flexibles.

### Caractéristiques principales
- **Structure MapAny**: Carte thread-safe avec support des types dynamiques
- **Conversion de type**: Conversion automatique entre différents types de données
- **Accès imbriqué**: Support pour l'accès aux clés imbriquées avec des séparateurs configurables
- **Support multi-format**: Intégration JSON et YAML

### Composants principaux
```go
// Créer à partir d'une carte
m := anyx.NewMap(map[string]interface{}{
    "name": "Jean",
    "age":  30,
})

// Créer à partir de JSON
jsonData := []byte(`{"name": "Alice", "score": 95.5}`)
m, err := anyx.NewMapWithJson(jsonData)
```

### Meilleures pratiques
1. **Activez Cut pour l'accès imbriqué**: Utilisez `EnableCut()` lors du travail avec des structures de données imbriquées
2. **Sécurité des types**: Gérez toujours la possibilité d'échecs de conversion de type

---

## العربية

### نظرة عامة
توفر حزمة `anyx` أدوات قوية للعمل مع قيم أي نوع في Go، وتقدم عمليات خريطة آمنة للخيوط مع دعم تحويل النوع الديناميكي وأنماط الوصول المرنة للبيانات.

### الميزات الرئيسية
- **هيكل MapAny**: خريطة آمنة للخيوط مع دعم الأنواع الديناميكية
- **تحويل النوع**: التحويل التلقائي بين أنواع البيانات المختلفة
- **الوصول المتداخل**: دعم للوصول للمفاتيح المتداخلة مع فواصل قابلة للتكوين
- **دعم متعدد التنسيقات**: تكامل JSON و YAML

### المكونات الأساسية
```go
// إنشاء من الخريطة
m := anyx.NewMap(map[string]interface{}{
    "name": "أحمد",
    "age":  30,
})

// إنشاء من JSON
jsonData := []byte(`{"name": "فاطمة", "score": 95.5}`)
m, err := anyx.NewMapWithJson(jsonData)
```

### أفضل الممارسات
1. **تفعيل Cut للوصول المتداخل**: استخدم `EnableCut()` عند العمل مع هياكل البيانات المتداخلة
2. **أمان النوع**: تعامل دائماً مع إمكانية فشل تحويل النوع

---

## Español

### Descripción general
El paquete `anyx` proporciona utilidades poderosas para trabajar con valores de cualquier tipo en Go, ofreciendo operaciones de mapa thread-safe con soporte de conversión de tipo dinámico y patrones de acceso de datos flexibles.

### Características principales
- **Estructura MapAny**: Mapa thread-safe con soporte de tipos dinámicos
- **Conversión de tipo**: Conversión automática entre diferentes tipos de datos
- **Acceso anidado**: Soporte para acceso a claves anidadas con separadores configurables
- **Soporte multi-formato**: Integración JSON y YAML

### Componentes principales
```go
// Crear desde mapa
m := anyx.NewMap(map[string]interface{}{
    "name": "Juan",
    "age":  30,
})

// Crear desde JSON
jsonData := []byte(`{"name": "Alicia", "score": 95.5}`)
m, err := anyx.NewMapWithJson(jsonData)
```

### Mejores prácticas
1. **Habilite Cut para acceso anidado**: Use `EnableCut()` cuando trabaje con estructuras de datos anidadas
2. **Seguridad de tipos**: Siempre maneje la posibilidad de fallas en la conversión de tipos