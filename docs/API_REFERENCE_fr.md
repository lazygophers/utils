# lazygophers/utils Référence API

Une bibliothèque d'utilitaires Go complète avec conception modulaire, fournissant des fonctionnalités essentielles pour les tâches de développement courantes. Cette bibliothèque met l'accent sur la sécurité des types, l'optimisation des performances et suit les standards Go 1.24+.

## Installation

```go
go get github.com/lazygophers/utils
```

## Table des Matières

- [Package Principal (utils)](#package-principal-utils)
- [Conversion de Types (candy)](#conversion-de-types-candy)
- [Manipulation de Chaînes (stringx)](#manipulation-de-chaînes-stringx)
- [Opérations de Cartes (anyx)](#opérations-de-cartes-anyx)
- [Contrôle de Concurrence (wait)](#contrôle-de-concurrence-wait)
- [Circuit Breaker (hystrix)](#circuit-breaker-hystrix)
- [Gestion de Configuration (config)](#gestion-de-configuration-config)
- [Opérations Cryptographiques (cryptox)](#opérations-cryptographiques-cryptox)

---

## Package Principal (utils)

Le package racine fournit des utilitaires fondamentaux incluant la gestion d'erreurs, les opérations de base de données et la validation.

### Fonctions Must

Utilitaires de gestion d'erreurs qui paniquent sur erreur - utiles pour l'initialisation et les opérations critiques.

```go
// Must combine vérification d'erreur et retour de valeur
func Must[T any](value T, err error) T

// MustOk panique si ok est false
func MustOk[T any](value T, ok bool) T

// MustSuccess panique si l'erreur n'est pas nil
func MustSuccess(err error)
```

**Exemple d'utilisation :**
```go
import "github.com/lazygophers/utils"

// Gestion d'erreurs lors de l'initialisation
config := utils.Must(loadConfig())

// Opérations de conversion
value := utils.MustOk(m["key"])
```

### Intégration Base de Données

```go
// Scan scanne les champs de base de données vers une structure, supporte la désérialisation JSON
func Scan(src interface{}, dst interface{}) error

// Value convertit une structure en valeur de base de données, supporte la sérialisation JSON
func Value(m interface{}) (driver.Value, error)
```

---

## Conversion de Types (candy)

Outils complets de conversion de types et manipulation de slices, avec 99.3% de couverture de test.

### Conversions de Types Basiques

```go
// Conversion booléenne
func ToBool(v interface{}) bool

// Conversion chaîne
func ToString(v interface{}) string

// Conversions entières
func ToInt(v interface{}) int
func ToInt32(v interface{}) int32
func ToInt64(v interface{}) int64

// Conversions flottantes
func ToFloat32(v interface{}) float32
func ToFloat64(v interface{}) float64
```

**Exemple d'utilisation :**
```go
import "github.com/lazygophers/utils/candy"

// Conversion intelligente de types
str := candy.ToString(123)        // "123"
num := candy.ToInt("456")         // 456
flag := candy.ToBool("true")      // true
```

### Opérations de Slices

```go
// Filter filtre les éléments de slice
func Filter[T any](slice []T, predicate func(T) bool) []T

// Map transforme les éléments de slice
func Map[T, R any](slice []T, mapper func(T) R) []R

// Unique supprime les éléments dupliqués
func Unique[T comparable](slice []T) []T
```

---

## Manipulation de Chaînes (stringx)

Manipulation de chaînes haute performance, avec 96.4% de couverture de test.

### Conversions Zero-Copy

```go
// ToString conversion de chaîne efficace (zero-copy)
func ToString(b []byte) string

// ToBytes conversion de byte efficace (zero-copy)
func ToBytes(s string) []byte
```

### Conversion de Noms

```go
// Camel2Snake conversion de camelCase vers snake_case
func Camel2Snake(s string) string

// Snake2Camel conversion de snake_case vers camelCase
func Snake2Camel(s string) string
```

**Exemple d'utilisation :**
```go
import "github.com/lazygophers/utils/stringx"

// Conversion de noms
snake := stringx.Camel2Snake("UserName")     // "user_name"
camel := stringx.Snake2Camel("user_name")    // "UserName"

// Conversions zero-copy (haute performance)
bytes := stringx.ToBytes("hello")
str := stringx.ToString([]byte("world"))
```

---

## Opérations de Cartes (anyx)

Opérations de cartes type-agnostic et extraction de valeurs, avec 99.0% de couverture de test.

### Structure MapAny

```go
// NewMap crée un nouveau MapAny
func NewMap(m map[string]interface{}) *MapAny

// Get obtient une valeur
func (m *MapAny) Get(key string) interface{}

// Set définit une valeur
func (m *MapAny) Set(key string, value interface{})
```

### Extraction Type-Safe

```go
// Obtenir des valeurs avec types spécifiques
func (m *MapAny) GetString(key string) string
func (m *MapAny) GetInt(key string) int
func (m *MapAny) GetBool(key string) bool
```

---

## Contrôle de Concurrence (wait)

Utilitaires avancés de concurrence et synchronisation.

### Traitement Asynchrone

```go
// Async traite les éléments de slice de manière asynchrone
func Async[T any](items []T, workerCount int, processor func(T)) error
```

**Exemple d'utilisation :**
```go
import "github.com/lazygophers/utils/wait"

urls := []string{
    "https://api1.com",
    "https://api2.com",
}

// Traitement concurrent d'URLs
err := wait.Async(urls, 2, func(url string) {
    resp, err := http.Get(url)
    // Traiter la réponse...
})
```

---

## Circuit Breaker (hystrix)

Implémentation haute performance du pattern circuit breaker.

```go
type CircuitBreaker struct {
    // Implémentation circuit breaker haute performance
}

func NewCircuitBreaker(config Config) *CircuitBreaker
func (cb *CircuitBreaker) Call(fn func() error) error
```

---

## Gestion de Configuration (config)

Chargement de configuration multi-format.

```go
// LoadConfig charge le fichier de configuration
func LoadConfig(filename string, config interface{}) error
```

Formats supportés :
- JSON (.json)
- YAML (.yaml, .yml)
- TOML (.toml)
- INI (.ini)

---

## Opérations Cryptographiques (cryptox)

Opérations cryptographiques complètes avec 100% de couverture de test.

### Chiffrement AES

```go
// Chiffrement/déchiffrement AES
func Encrypt(plaintext, key []byte) ([]byte, error)
func Decrypt(ciphertext, key []byte) ([]byte, error)
```

### Opérations RSA

```go
// Génération de clés RSA
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error)
```

### Fonctions de Hachage

```go
// Hachages courants
func MD5(data []byte) string
func SHA256(data []byte) string
func SHA512(data []byte) string
```

---

## Caractéristiques de Performance

- **Opérations Zero Allocation** : Beaucoup de fonctions atteignent zéro allocations mémoire
- **Opérations Atomiques** : Implémentations lock-free pour des scénarios haute concurrence
- **Structures Alignées Mémoire** : Optimisées pour l'efficacité du cache CPU
- **Optimisations Génériques** : Opérations type-safe sans surcharge de réflexion runtime

## Pattern de Gestion d'Erreurs

Tous les packages suivent un pattern de gestion d'erreurs cohérent :
1. Utilisation de `github.com/lazygophers/log` pour enregistrer les erreurs
2. Retour de messages d'erreur significatifs
3. Fourniture de fonctions `Must*` pour les opérations critiques

Cette référence fournit un guide complet pour utiliser la bibliothèque LazyGophers Utils, aidant les développeurs à construire efficacement des applications Go de haute qualité.