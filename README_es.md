# LazyGophers Utils

> 🚀 Una biblioteca de utilidades Go rica en funciones y de alto rendimiento que hace que el desarrollo en Go sea más eficiente

**🌍 Idiomas**: [English](README.md) • [中文](README_zh.md) • [繁體中文](README_zh-hant.md) • [Español](README_es.md) • [Français](README_fr.md) • [Русский](README_ru.md) • [العربية](README_ar.md)

[![Go Version](https://img.shields.io/badge/Go-1.18+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-AGPL%20v3-green.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils.svg)](https://pkg.go.dev/github.com/lazygophers/utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/utils)](https://goreportcard.com/report/github.com/lazygophers/utils)
[![GitHub releases](https://img.shields.io/github/release/lazygophers/utils.svg)](https://github.com/lazygophers/utils/releases)
[![GoProxy.cn Downloads](https://goproxy.cn/stats/github.com/lazygophers/utils/badges/download-count.svg)](https://goproxy.cn/stats/github.com/lazygophers/utils)
[![Test Coverage](https://img.shields.io/badge/coverage-85%25-brightgreen.svg)](https://github.com/lazygophers/utils/actions)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/lazygophers/utils)

## 📋 Tabla de Contenidos

- [Descripción del Proyecto](#-descripción-del-proyecto)
- [Características Principales](#-características-principales)
- [Inicio Rápido](#-inicio-rápido)
- [Documentación](#-documentación)
- [Módulos Principales](#-módulos-principales)
- [Módulos de Funcionalidad](#-módulos-de-funcionalidad)
- [Ejemplos de Uso](#-ejemplos-de-uso)
- [Datos de Rendimiento](#-datos-de-rendimiento)
- [Contribuir](#-contribuir)
- [Licencia](#-licencia)
- [Soporte de la Comunidad](#-soporte-de-la-comunidad)

## 💡 Descripción del Proyecto

LazyGophers Utils es una biblioteca de utilidades Go integral y de alto rendimiento que proporciona más de 20 módulos profesionales que cubren diversas necesidades en el desarrollo diario. Adopta un diseño modular para importaciones bajo demanda con cero conflictos de dependencias.

**Filosofía de Diseño**: Simple, Eficiente, Confiable

## ✨ Características Principales

| Característica | Descripción | Ventaja |
|----------------|-------------|---------|
| 🧩 **Diseño Modular** | Más de 20 módulos independientes | Importar bajo demanda, reducir tamaño |
| ⚡ **Alto Rendimiento** | Probado con benchmarks | Respuesta en microsegundos, amigable con la memoria |
| 🛡️ **Tipo Seguro** | Uso completo de genéricos | Verificación de errores en tiempo de compilación |
| 🔒 **Seguro para Concurrencia** | Diseño amigable con goroutines | Listo para producción |
| 📚 **Bien Documentado** | Cobertura de documentación 95%+ | Fácil de aprender y usar |
| 🧪 **Bien Probado** | Cobertura de pruebas 85%+ | Garantía de calidad |

## 🚀 Inicio Rápido

### Instalación

```bash
go get github.com/lazygophers/utils
```

### Uso Básico

```go
package main

import (
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // Manejo de errores
    value := utils.Must(getValue())
    
    // Conversión de tipos
    age := candy.ToInt("25")
    
    // Procesamiento de tiempo
    cal := xtime.NowCalendar()
    fmt.Println(cal.String()) // 2023年08月15日 六月廿九 兔年 处暑
}
```

## 📖 Documentación

### 📁 Documentación de Módulos
- **Módulos Principales**: [Manejo de Errores](must.go) | [Base de Datos](orm.go) | [Validación](validate.go)
- **Procesamiento de Datos**: [candy](candy/) | [json](json/) | [stringx](stringx/)
- **Herramientas de Tiempo**: [xtime](xtime/) | [xtime996](xtime/xtime996/) | [xtime955](xtime/xtime955/)
- **Herramientas del Sistema**: [config](config/) | [runtime](runtime/) | [osx](osx/)
- **Red y Seguridad**: [network](network/) | [cryptox](cryptox/) | [pgp](pgp/)
- **Concurrencia y Control**: [routine](routine/) | [wait](wait/) | [hystrix](hystrix/)

### 📋 Referencia Rápida
- [🔧 Guía de Instalación](#-inicio-rápido)
- [📝 Ejemplos de Uso](#-ejemplos-de-uso)
- [📚 Índice de Documentación Completa](docs/) - Centro de documentación integral
- [🎯 Buscar Módulos por Escenario](docs/#-búsqueda-rápida) - Posicionamiento rápido por casos de uso
- [🏗️ Documentación de Arquitectura](docs/architecture_en.md) - Inmersión profunda en el diseño del sistema

### 🌍 Documentación Multiidioma
- [English](README.md) - Versión en inglés
- [中文](README_zh.md) - Versión en chino
- [繁體中文](README_zh-hant.md) - Chino tradicional
- [Français](README_fr.md) - Versión en francés
- [Русский](README_ru.md) - Versión en ruso
- [العربية](README_ar.md) - Versión en árabe

## 🔧 Módulos Principales

### Manejo de Errores (`must.go`)
```go
// Asegurar el éxito de la operación, pánico en caso de falla
value := utils.Must(getValue())

// Verificar que no hay error
utils.MustSuccess(doSomething())

// Verificar estado booleano
result := utils.MustOk(checkCondition())
```

### Operaciones de Base de Datos (`orm.go`)
```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age" default:"18"`
}

// Escanear datos de base de datos a estructura
err := utils.Scan(dbData, &user)

// Convertir estructura a valor de base de datos
value, err := utils.Value(user)
```

### Validación de Datos (`validate.go`)
```go
type Config struct {
    Email string `validate:"required,email"`
    Port  int    `validate:"min=1,max=65535"`
}

// Validación rápida
err := utils.Validate(&config)
```

## 📦 Módulos de Funcionalidad

Para una lista detallada de módulos de funcionalidad y ejemplos de uso, consulte la [documentación completa](docs/).

## 🎯 Ejemplos de Uso

### Ejemplo de Aplicación Completa

```go
package main

import (
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/config"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/xtime"
)

type AppConfig struct {
    Port     int    `json:"port" default:"8080" validate:"min=1,max=65535"`
    Database string `json:"database" validate:"required"`
    Debug    bool   `json:"debug" default:"false"`
}

func main() {
    // 1. Cargar configuración
    var cfg AppConfig
    utils.MustSuccess(config.Load(&cfg, "config.json"))
    
    // 2. Validar configuración
    utils.MustSuccess(utils.Validate(&cfg))
    
    // 3. Conversión de tipos
    portStr := candy.ToString(cfg.Port)
    
    // 4. Procesamiento de tiempo
    cal := xtime.NowCalendar()
    log.Printf("Aplicación iniciada: %s", cal.String())
    
    // 5. Iniciar servidor
    startServer(cfg)
}
```

## 📊 Datos de Rendimiento

| Operación | Tiempo | Asignación de Memoria | vs Biblioteca Estándar |
|-----------|--------|-----------------------|-------------------------|
| `candy.ToInt()` | 12.3 ns/op | 0 B/op | **3.2x más rápido** |
| `json.Marshal()` | 156 ns/op | 64 B/op | **1.8x más rápido** |
| `xtime.Now()` | 45.2 ns/op | 0 B/op | **2.1x más rápido** |
| `utils.Must()` | 2.1 ns/op | 0 B/op | **Cero sobrecarga** |

## 🤝 Contribuir

¡Damos la bienvenida a contribuciones de todo tipo!

1. 🍴 Haz fork del proyecto
2. 🌿 Crea una rama de función
3. 📝 Escribe código y pruebas
4. 🧪 Asegúrate de que las pruebas pasen
5. 📤 Envía un PR

## 📄 Licencia

Este proyecto está licenciado bajo la Licencia Pública General Affero de GNU v3.0.

Consulta el archivo [LICENSE](LICENSE) para más detalles.

## 🌟 Soporte de la Comunidad

### Obtener Ayuda

- 📖 **Documentación**: [Documentación Completa](docs/)
- 🐛 **Reportes de Bugs**: [GitHub Issues](https://github.com/lazygophers/utils/issues)
- 💬 **Discusiones**: [GitHub Discussions](https://github.com/lazygophers/utils/discussions)
- ❓ **Preguntas y Respuestas**: [Stack Overflow](https://stackoverflow.com/questions/tagged/lazygophers-utils)

---

<div align="center">

**¡Si este proyecto te ayuda, por favor danos una ⭐ Estrella!**

[🚀 Comenzar](#-inicio-rápido) • [📖 Ver Documentación](docs/) • [🤝 Unirse a la Comunidad](https://github.com/lazygophers/utils/discussions)

</div>