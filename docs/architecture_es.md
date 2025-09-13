# LazyGophers Utils - Documentación de Arquitectura

## 🏗️ Visión General

LazyGophers Utils es una biblioteca integral de utilidades de Go diseñada con enfoque en modularidad, rendimiento y experiencia del desarrollador. La biblioteca sigue las mejores prácticas modernas de Go, incluyendo el uso extensivo de genéricos, operaciones atómicas y optimizaciones zero-copy.

## 📊 Estadísticas del Proyecto

- **Total de Paquetes**: 25 módulos independientes
- **Líneas de Código**: 56,847 líneas
- **Archivos Go**: 323 archivos
- **Cobertura de Pruebas**: 85.8%
- **Versión Go**: 1.24.0+

## 🎯 Principios de Diseño

### 1. Arquitectura Modular
Cada paquete está diseñado como un módulo independiente que puede ser importado y utilizado por separado, minimizando las dependencias y el tamaño del binario.

### 2. Enfoque Rendimiento-Primero
- Uso extensivo de operaciones atómicas para operaciones thread-safe
- Algoritmos lock-free donde sea posible
- Optimizaciones de alineación de memoria
- Conversiones string/byte zero-copy usando operaciones unsafe
- Pool de objetos para operaciones de alta frecuencia

### 3. Seguridad de Tipos
Uso extensivo de genéricos Go 1.18+ para proporcionar APIs type-safe mientras se mantiene el rendimiento.

### 4. Manejo Consistente de Errores
Todos los paquetes siguen un patrón consistente de manejo de errores: registrar errores usando `github.com/lazygophers/log` antes de devolverlos.

## 🏛️ Arquitectura de Paquetes

### Paquetes Principales
Estos paquetes forman la base de la biblioteca:

#### `utils` (Paquete Raíz)
- **Propósito**: Utilidades fundamentales para manejo de errores, operaciones de base de datos y validación
- **Características Clave**:
  - `Must[T any](value T, err error) T` - Wrapper panic-on-error
  - `Scan()` y `Value()` - Utilidades de integración de base de datos
  - `Validate()` - Validación de struct usando go-playground/validator
- **Dependencias**: Dependencias externas mínimas

#### `candy`
- **Propósito**: Conversión de tipos integral y manipulación de slices
- **Tamaño**: 143 archivos, 15,963 líneas de código
- **Funciones Principales**:
  - Conversión de tipos: `ToBool()`, `ToString()`, `ToInt*()`, `ToFloat*()`
  - Programación funcional: `All()`, `Any()`, `Filter()`, `Map()`
  - Utilidades de slices: `Unique()`, `Sort()`, `Shuffle()`, `Chunk()`
- **Rendimiento**: Optimizado con genéricos para seguridad de tipos

#### `json`
- **Propósito**: Operaciones JSON mejoradas con optimización de rendimiento
- **Características**:
  - Optimización específica de plataforma usando sonic en plataformas soportadas
  - Wrapper de API consistente para diferentes implementaciones de JSON
- **Crítico para Rendimiento**: Sí

### Paquetes de Infraestructura

#### `runtime`
- **Propósito**: Utilidades de runtime e información del sistema
- **Características Principales**:
  - Manejo y recuperación de panic integral
  - Utilidades de directorios del sistema (`ExecDir()`, `UserHomeDir()`, etc.)
  - Utilidades de detección de plataforma

#### `routine`
- **Propósito**: Gestión mejorada de goroutine
- **Características**:
  - Ejecución segura de goroutine con recuperación de panic
  - Trazado de Goroutine usando `github.com/petermattis/goid`
  - Limpieza automática y manejo de errores

#### `app`
- **Propósito**: Ciclo de vida de aplicación e información de construcción
- **Características**:
  - Metadatos de construcción (commit, branch, tag, build date)
  - Detección de entorno
  - Información de versión

### Paquetes de Utilidades

#### `stringx`
- **Propósito**: Manipulación de cadenas de alto rendimiento
- **Tamaño**: 11 archivos, 4,385 líneas de código
- **Optimizaciones de Rendimiento**:
  - Conversiones string/byte zero-copy usando operaciones unsafe
  - Optimizaciones fast-path ASCII para conversión de caso
  - Operaciones de cadenas eficientes en memoria
- **Funciones Principales**:
  - `ToString()` / `ToBytes()` - Conversiones zero-copy
  - `Camel2Snake()` - Conversión de caso optimizada
  - Utilidades Unicode con optimizaciones de rendimiento

#### `anyx`
- **Propósito**: Operaciones de map type-agnostic y extracción de valores
- **Tamaño**: 4 archivos, 3,999 líneas de código
- **Capacidades Principales**:
  - Operaciones de map thread-safe con conversión de tipos
  - Soporte de claves anidadas con notación de punto
  - Utilidades de extracción de tipos integrales
- **Dependencias**: Usa paquetes `candy` y `json`

#### `wait`
- **Propósito**: Utilidades avanzadas de concurrencia y sincronización
- **Tamaño**: 6 archivos, 1,323 líneas de código
- **Componentes Principales**:
  - `Async()` - Pool de goroutine con distribución de trabajo
  - `AsyncUnique()` - Deduplicación de tareas en procesamiento concurrente
  - WaitGroup mejorado con características adicionales
  - Pool de objetos para optimización de rendimiento

### Paquetes Especializados

#### `cryptox`
- **Propósito**: Operaciones criptográficas integrales
- **Tamaño**: 40 archivos, 11,254 líneas de código
- **Capacidades**:
  - Cifrado simétrico: AES, DES, Triple DES, ChaCha20
  - Criptografía asimétrica: RSA, ECDSA, ECDH
  - Hashing: Familia SHA, MD5, HMAC, BLAKE2, SHA3
  - Derivación de claves: PBKDF2, Scrypt, Argon2
- **Seguridad**: Implementaciones listas para producción siguiendo mejores prácticas

#### `hystrix`
- **Propósito**: Implementación de patrón circuit breaker de alto rendimiento
- **Tamaño**: 4 archivos, 1,367 líneas de código
- **Optimizaciones de Rendimiento**:
  - Operaciones atómicas para gestión de estado
  - Algoritmos lock-free
  - Alineación de memoria para eficiencia de caché de CPU
  - Tres variantes: estándar, rápido y optimizado por lotes
- **Benchmark**: Operaciones de registro a ~46ns/op con cero asignaciones

#### `xtime`
- **Propósito**: Operaciones de tiempo extendidas con soporte de calendario chino
- **Tamaño**: 21 archivos, 10,744 líneas de código
- **Características Únicas**:
  - Cálculos de calendario lunar chino
  - Cálculo de términos solares
  - Constantes de tiempo de negocio para horarios de trabajo
  - Sub-paquetes para diferentes patrones de trabajo (007, 955, 996)

### Paquetes de Configuración e I/O

#### `config`
- **Propósito**: Carga de configuración multi-formato
- **Formatos Soportados**: JSON, YAML, TOML, INI, HCL
- **Características**: Carga de configuración consciente del entorno con validación

#### `bufiox`
- **Propósito**: Operaciones I/O con buffer y utilidades
- **Características**: Utilidades de escaneo personalizadas para optimización de rendimiento

#### `osx`
- **Propósito**: Interfaz de SO multiplataforma y operaciones de archivos
- **Tamaño**: 9 archivos, 2,554 líneas de código
- **Capacidades**: Utilidades de sistema de archivos con compatibilidad multiplataforma

### Red y Comunicación

#### `network`
- **Propósito**: Utilidades de red y asistentes
- **Características**: Utilidades de dirección IP, detección de interfaz, extracción de IP real

### Utilidades Aleatorias y de Prueba

#### `randx`
- **Propósito**: Generación extendida de números aleatorios y datos
- **Tamaño**: 9 archivos, 2,014 líneas de código
- **Capacidades**:
  - Varias distribuciones de probabilidad
  - Utilidades aleatorias basadas en tiempo
  - Generadores aleatorios optimizados para rendimiento

#### `fake`
- **Propósito**: Generación de datos falsos para pruebas
- **Características**: Generación de user agent, utilidades de datos de prueba

## 🔗 Gráfico de Dependencias

```
Paquete Raíz (utils)
├── json (operaciones JSON principales)
├── candy (conversión de tipos)
└── ... (dependencias externas mínimas)

Capa de Infraestructura
├── runtime → app
├── routine → runtime, log
└── osx (abstracción de SO)

Capa de Utilidades
├── stringx (manipulación de cadenas)
├── anyx → candy, json
├── wait → routine, runtime
└── xtime (operaciones de tiempo)

Capa Especializada
├── cryptox (operaciones criptográficas)
├── hystrix → randx (circuit breaker)
├── config → json, osx, runtime
└── network (utilidades de red)
```

## 🚀 Características de Rendimiento

### Benchmarks (Apple M3)

| Paquete | Operación | Rendimiento | Memoria |
|---------|-----------|-------------|--------|
| atexit | Register | 46.69 ns/op | 43 B/op, 0 allocs/op |
| atexit | RegisterConcurrent | 43.81 ns/op | 44 B/op, 0 allocs/op |
| atexit | ExecuteCallbacks | 545.9 ns/op | 896 B/op, 1 allocs/op |

### Características de Rendimiento

1. **Operaciones Lock-Free**: Las rutas críticas usan operaciones atómicas en lugar de mutex
2. **Alineación de Memoria**: Estructuras alineadas para rendimiento óptimo de caché de CPU
3. **Operaciones Zero-Copy**: Conversiones string/byte sin asignación de memoria
4. **Pool de Objetos**: Reduce la presión de GC en operaciones de alta frecuencia
5. **Optimizaciones de Genéricos**: Operaciones type-safe sin reflexión en tiempo de ejecución

## 🧪 Pruebas y Calidad

### Cobertura de Pruebas por Paquete
- **candy**: 99.3%
- **anyx**: 99.0%
- **atexit**: 100.0%
- **bufiox**: 100.0%
- **cryptox**: 100.0%
- **defaults**: 100.0%
- **stringx**: 96.4%
- **osx**: 97.7%
- **config**: 95.7%
- **network**: 89.1%

### Garantía de Calidad
- Pruebas unitarias integrales con cobertura de casos extremos
- Pruebas de benchmark para operaciones críticas de rendimiento
- Pruebas de condición de carrera para operaciones concurrentes
- Pruebas de fuga de memoria para operaciones de larga duración

## 🌏 Características Culturales

### Soporte de Calendario Chino (paquete xtime)
- **Calendario Lunar**: Implementación completa del calendario lunar chino
- **Términos Solares**: Cálculo de 24 términos solares chinos tradicionales
- **Horarios de Trabajo**: Soporte para patrones de trabajo chinos (007, 955, 996)
- **Festivales Tradicionales**: Cálculos de festivales tradicionales chinos

## 🔮 Consideraciones Arquitectónicas Futuras

1. **Sistema de Plugins**: Considerar implementar una arquitectura de plugins para extensibilidad
2. **Observabilidad**: Integración mejorada de métricas y trazado
3. **Configuración**: Capacidades de recarga en caliente de configuración
4. **Caché**: Capa de caché distribuido para escenarios de alto rendimiento
5. **Streaming**: Utilidades de streaming mejoradas para procesamiento de datos grandes

## 📈 Escalabilidad

La arquitectura está diseñada para escalar vertical y horizontalmente:

- **Escalado Vertical**: Uso optimizado de memoria y rendimiento de CPU
- **Escalado Horizontal**: Operaciones thread-safe soportan uso concurrente
- **Microservicios**: Cada paquete puede usarse independientemente en diferentes servicios
- **Cloud Native**: Compatible con entornos de contenedores y plataformas en la nube

Esta arquitectura proporciona una base sólida para construir aplicaciones Go de alto rendimiento mientras mantiene la claridad del código y la productividad del desarrollador.