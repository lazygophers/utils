# lazygophers/utils Справочник API

Комплексная библиотека утилит Go с модульным дизайном, предоставляющая основной функционал для обычных задач разработки. Эта библиотека делает акцент на безопасности типов, оптимизации производительности и следует стандартам Go 1.24+.

## Установка

```go
go get github.com/lazygophers/utils
```

## Содержание

- [Основной Пакет (utils)](#основной-пакет-utils)
- [Преобразование Типов (candy)](#преобразование-типов-candy)
- [Манипуляция Строками (stringx)](#манипуляция-строками-stringx)
- [Операции с Картами (anyx)](#операции-с-картами-anyx)
- [Управление Concurrency (wait)](#управление-concurrency-wait)
- [Circuit Breaker (hystrix)](#circuit-breaker-hystrix)
- [Управление Конфигурацией (config)](#управление-конфигурацией-config)
- [Криптографические Операции (cryptox)](#криптографические-операции-cryptox)

---

## Основной Пакет (utils)

Корневой пакет предоставляет фундаментальные утилиты, включая обработку ошибок, операции с базой данных и валидацию.

### Функции Must

Утилиты обработки ошибок, которые вызывают panic при ошибке - полезны для инициализации и критических операций.

```go
// Must комбинирует проверку ошибок и возврат значения
func Must[T any](value T, err error) T

// MustOk вызывает panic если ok равно false
func MustOk[T any](value T, ok bool) T

// MustSuccess вызывает panic если ошибка не nil
func MustSuccess(err error)
```

**Пример использования:**
```go
import "github.com/lazygophers/utils"

// Обработка ошибок при инициализации
config := utils.Must(loadConfig())

// Операции преобразования
value := utils.MustOk(m["key"])
```

### Интеграция с Базой Данных

```go
// Scan сканирует поля базы данных в структуру, поддерживает JSON десериализацию
func Scan(src interface{}, dst interface{}) error

// Value преобразует структуру в значение базы данных, поддерживает JSON сериализацию
func Value(m interface{}) (driver.Value, error)
```

---

## Преобразование Типов (candy)

Комплексные инструменты преобразования типов и манипуляции срезами, с 99.3% покрытием тестами.

### Базовые Преобразования Типов

```go
// Логическое преобразование
func ToBool(v interface{}) bool

// Строковое преобразование
func ToString(v interface{}) string

// Целочисленные преобразования
func ToInt(v interface{}) int
func ToInt32(v interface{}) int32
func ToInt64(v interface{}) int64

// Преобразования с плавающей точкой
func ToFloat32(v interface{}) float32
func ToFloat64(v interface{}) float64
```

**Пример использования:**
```go
import "github.com/lazygophers/utils/candy"

// Умное преобразование типов
str := candy.ToString(123)        // "123"
num := candy.ToInt("456")         // 456
flag := candy.ToBool("true")      // true
```

### Операции со Срезами

```go
// Filter фильтрует элементы среза
func Filter[T any](slice []T, predicate func(T) bool) []T

// Map преобразует элементы среза
func Map[T, R any](slice []T, mapper func(T) R) []R

// Unique удаляет дублирующиеся элементы
func Unique[T comparable](slice []T) []T
```

---

## Манипуляция Строками (stringx)

Высокопроизводительная манипуляция строками, с 96.4% покрытием тестами.

### Zero-Copy Преобразования

```go
// ToString эффективное преобразование строк (zero-copy)
func ToString(b []byte) string

// ToBytes эффективное преобразование байт (zero-copy)
func ToBytes(s string) []byte
```

### Преобразование Названий

```go
// Camel2Snake преобразование из camelCase в snake_case
func Camel2Snake(s string) string

// Snake2Camel преобразование из snake_case в camelCase
func Snake2Camel(s string) string
```

**Пример использования:**
```go
import "github.com/lazygophers/utils/stringx"

// Преобразование названий
snake := stringx.Camel2Snake("UserName")     // "user_name"
camel := stringx.Snake2Camel("user_name")    // "UserName"

// Zero-copy преобразования (высокая производительность)
bytes := stringx.ToBytes("hello")
str := stringx.ToString([]byte("world"))
```

---

## Операции с Картами (anyx)

Type-agnostic операции с картами и извлечение значений, с 99.0% покрытием тестами.

### Структура MapAny

```go
// NewMap создает новый MapAny
func NewMap(m map[string]interface{}) *MapAny

// Get получает значение
func (m *MapAny) Get(key string) interface{}

// Set устанавливает значение
func (m *MapAny) Set(key string, value interface{})
```

### Type-Safe Извлечение

```go
// Получение значений с определенными типами
func (m *MapAny) GetString(key string) string
func (m *MapAny) GetInt(key string) int
func (m *MapAny) GetBool(key string) bool
```

---

## Управление Concurrency (wait)

Продвинутые утилиты concurrency и синхронизации.

### Асинхронная Обработка

```go
// Async обрабатывает элементы среза асинхронно
func Async[T any](items []T, workerCount int, processor func(T)) error
```

**Пример использования:**
```go
import "github.com/lazygophers/utils/wait"

urls := []string{
    "https://api1.com",
    "https://api2.com",
}

// Concurrent обработка URL
err := wait.Async(urls, 2, func(url string) {
    resp, err := http.Get(url)
    // Обработать ответ...
})
```

---

## Circuit Breaker (hystrix)

Высокопроизводительная реализация паттерна circuit breaker.

```go
type CircuitBreaker struct {
    // Высокопроизводительная реализация circuit breaker
}

func NewCircuitBreaker(config Config) *CircuitBreaker
func (cb *CircuitBreaker) Call(fn func() error) error
```

---

## Управление Конфигурацией (config)

Мульти-форматная загрузка конфигурации.

```go
// LoadConfig загружает файл конфигурации
func LoadConfig(filename string, config interface{}) error
```

Поддерживаемые форматы:
- JSON (.json)
- YAML (.yaml, .yml)
- TOML (.toml)
- INI (.ini)

---

## Криптографические Операции (cryptox)

Комплексные криптографические операции со 100% покрытием тестами.

### AES Шифрование

```go
// AES шифрование/расшифрование
func Encrypt(plaintext, key []byte) ([]byte, error)
func Decrypt(ciphertext, key []byte) ([]byte, error)
```

### RSA Операции

```go
// Генерация RSA ключей
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error)
```

### Хеш Функции

```go
// Общие хеши
func MD5(data []byte) string
func SHA256(data []byte) string
func SHA512(data []byte) string
```

---

## Характеристики Производительности

- **Zero Allocation Операции**: Многие функции достигают нулевых аллокаций памяти
- **Атомарные Операции**: Lock-free реализации для высоко-concurrent сценариев
- **Memory Aligned Структуры**: Оптимизированы для эффективности CPU кеша
- **Оптимизации Дженериков**: Type-safe операции без runtime рефлексии

## Паттерн Обработки Ошибок

Все пакеты следуют консистентному паттерну обработки ошибок:
1. Использование `github.com/lazygophers/log` для логирования ошибок
2. Возврат значимых сообщений об ошибках
3. Предоставление `Must*` функций для критических операций

Этот справочник предоставляет комплексное руководство по использованию библиотеки LazyGophers Utils, помогая разработчикам эффективно строить высококачественные Go приложения.