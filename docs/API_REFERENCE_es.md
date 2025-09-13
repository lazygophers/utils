# lazygophers/utils Referencia API

Una biblioteca integral de utilidades de Go con diseño modular, proporcionando funcionalidad esencial para tareas comunes de desarrollo. Esta biblioteca enfatiza la seguridad de tipos, optimización de rendimiento y sigue los estándares Go 1.24+.

## Instalación

```go
go get github.com/lazygophers/utils
```

## Tabla de Contenidos

- [Paquete Principal (utils)](#paquete-principal-utils)
- [Conversión de Tipos (candy)](#conversión-de-tipos-candy)
- [Manipulación de Cadenas (stringx)](#manipulación-de-cadenas-stringx)
- [Operaciones de Mapas (anyx)](#operaciones-de-mapas-anyx)
- [Control de Concurrencia (wait)](#control-de-concurrencia-wait)
- [Circuit Breaker (hystrix)](#circuit-breaker-hystrix)
- [Gestión de Configuración (config)](#gestión-de-configuración-config)
- [Operaciones Criptográficas (cryptox)](#operaciones-criptográficas-cryptox)

---

## Paquete Principal (utils)

El paquete raíz proporciona utilidades fundamentales incluyendo manejo de errores, operaciones de base de datos y validación.

### Funciones Must

Utilidades de manejo de errores que entran en pánico ante error - útiles para inicialización y operaciones críticas.

```go
// Must combina verificación de errores y retorno de valor
func Must[T any](value T, err error) T

// MustOk entra en pánico si ok es false
func MustOk[T any](value T, ok bool) T

// MustSuccess entra en pánico si el error no es nil
func MustSuccess(err error)
```

**Ejemplo de uso:**
```go
import "github.com/lazygophers/utils"

// Manejo de errores durante inicialización
config := utils.Must(loadConfig())

// Operaciones de conversión
value := utils.MustOk(m["key"])
```

### Integración de Base de Datos

```go
// Scan escanea campos de base de datos a estructura, soporta deserialización JSON
func Scan(src interface{}, dst interface{}) error

// Value convierte estructura a valor de base de datos, soporta serialización JSON
func Value(m interface{}) (driver.Value, error)
```

---

## Conversión de Tipos (candy)

Herramientas integrales de conversión de tipos y manipulación de slices, con 99.3% de cobertura de pruebas.

### Conversiones de Tipos Básicas

```go
// Conversión booleana
func ToBool(v interface{}) bool

// Conversión de cadena
func ToString(v interface{}) string

// Conversiones enteras
func ToInt(v interface{}) int
func ToInt32(v interface{}) int32
func ToInt64(v interface{}) int64

// Conversiones de punto flotante
func ToFloat32(v interface{}) float32
func ToFloat64(v interface{}) float64
```

**Ejemplo de uso:**
```go
import "github.com/lazygophers/utils/candy"

// Conversión inteligente de tipos
str := candy.ToString(123)        // "123"
num := candy.ToInt("456")         // 456
flag := candy.ToBool("true")      // true
```

### Operaciones de Slices

```go
// Filter filtra elementos de slice
func Filter[T any](slice []T, predicate func(T) bool) []T

// Map transforma elementos de slice
func Map[T, R any](slice []T, mapper func(T) R) []R

// Unique elimina elementos duplicados
func Unique[T comparable](slice []T) []T
```

---

## Manipulación de Cadenas (stringx)

Manipulación de cadenas de alto rendimiento, con 96.4% de cobertura de pruebas.

### Conversiones Zero-Copy

```go
// ToString conversión eficiente de cadena (zero-copy)
func ToString(b []byte) string

// ToBytes conversión eficiente de byte (zero-copy)
func ToBytes(s string) []byte
```

### Conversión de Nombres

```go
// Camel2Snake conversión de camelCase a snake_case
func Camel2Snake(s string) string

// Snake2Camel conversión de snake_case a camelCase
func Snake2Camel(s string) string
```

**Ejemplo de uso:**
```go
import "github.com/lazygophers/utils/stringx"

// Conversión de nombres
snake := stringx.Camel2Snake("UserName")     // "user_name"
camel := stringx.Snake2Camel("user_name")    // "UserName"

// Conversiones zero-copy (alto rendimiento)
bytes := stringx.ToBytes("hello")
str := stringx.ToString([]byte("world"))
```

---

## Operaciones de Mapas (anyx)

Operaciones de mapas type-agnostic y extracción de valores, con 99.0% de cobertura de pruebas.

### Estructura MapAny

```go
// NewMap crea un nuevo MapAny
func NewMap(m map[string]interface{}) *MapAny

// Get obtiene un valor
func (m *MapAny) Get(key string) interface{}

// Set establece un valor
func (m *MapAny) Set(key string, value interface{})
```

### Extracción Type-Safe

```go
// Obtener valores con tipos específicos
func (m *MapAny) GetString(key string) string
func (m *MapAny) GetInt(key string) int
func (m *MapAny) GetBool(key string) bool
```

---

## Control de Concurrencia (wait)

Utilidades avanzadas de concurrencia y sincronización.

### Procesamiento Asíncrono

```go
// Async procesa elementos de slice de forma asíncrona
func Async[T any](items []T, workerCount int, processor func(T)) error
```

**Ejemplo de uso:**
```go
import "github.com/lazygophers/utils/wait"

urls := []string{
    "https://api1.com",
    "https://api2.com",
}

// Procesamiento concurrente de URLs
err := wait.Async(urls, 2, func(url string) {
    resp, err := http.Get(url)
    // Procesar respuesta...
})
```

---

## Circuit Breaker (hystrix)

Implementación de alto rendimiento del patrón circuit breaker.

```go
type CircuitBreaker struct {
    // Implementación circuit breaker de alto rendimiento
}

func NewCircuitBreaker(config Config) *CircuitBreaker
func (cb *CircuitBreaker) Call(fn func() error) error
```

---

## Gestión de Configuración (config)

Carga de configuración multi-formato.

```go
// LoadConfig carga archivo de configuración
func LoadConfig(filename string, config interface{}) error
```

Formatos soportados:
- JSON (.json)
- YAML (.yaml, .yml)
- TOML (.toml)
- INI (.ini)

---

## Operaciones Criptográficas (cryptox)

Operaciones criptográficas integrales con 100% de cobertura de pruebas.

### Cifrado AES

```go
// Cifrado/descifrado AES
func Encrypt(plaintext, key []byte) ([]byte, error)
func Decrypt(ciphertext, key []byte) ([]byte, error)
```

### Operaciones RSA

```go
// Generación de claves RSA
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error)
```

### Funciones Hash

```go
// Hashes comunes
func MD5(data []byte) string
func SHA256(data []byte) string
func SHA512(data []byte) string
```

---

## Características de Rendimiento

- **Operaciones Zero Allocation**: Muchas funciones logran cero asignaciones de memoria
- **Operaciones Atómicas**: Implementaciones lock-free para escenarios de alta concurrencia
- **Estructuras Memory Aligned**: Optimizadas para eficiencia de caché de CPU
- **Optimizaciones de Genéricos**: Operaciones type-safe sin sobrecarga de reflexión runtime

## Patrón de Manejo de Errores

Todos los paquetes siguen un patrón consistente de manejo de errores:
1. Uso de `github.com/lazygophers/log` para registrar errores
2. Retorno de mensajes de error significativos
3. Provisión de funciones `Must*` para operaciones críticas

Esta referencia proporciona una guía integral para usar la biblioteca LazyGophers Utils, ayudando a los desarrolladores a construir eficientemente aplicaciones Go de alta calidad.