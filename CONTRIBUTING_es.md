# Guía de Contribución

¡Bienvenido a contribuir al proyecto LazyGophers Utils! Agradecemos enormemente cada contribución de la comunidad.

[![Contributors](https://img.shields.io/badge/Contributors-Welcome-brightgreen.svg)](#cómo-contribuir)
[![Code Style](https://img.shields.io/badge/Code%20Style-Go%20Standard-blue.svg)](#estándares-de-código)

## 🤝 Cómo Contribuir

### Tipos de Contribuciones

Damos la bienvenida a los siguientes tipos de contribuciones:

- 🐛 **Correcciones de Errores** - Arreglar problemas conocidos
- ✨ **Nuevas Características** - Añadir nuevas funciones utilitarias o módulos
- 📚 **Mejoras de Documentación** - Mejorar documentación, añadir ejemplos
- 🎨 **Optimización de Código** - Optimización de rendimiento, refactorización
- 🧪 **Mejoras de Pruebas** - Aumentar cobertura de pruebas, arreglar problemas de pruebas
- 🌐 **Internacionalización** - Añadir soporte multi-idioma

### Proceso de Contribución

#### 1. Preparación

**Fork del Proyecto**
```bash
# 1. Haz fork de este proyecto a tu cuenta de GitHub
# 2. Clona tu fork localmente
git clone https://github.com/TU_USUARIO/utils.git
cd utils

# 3. Añade el proyecto original como repositorio upstream
git remote add upstream https://github.com/lazygophers/utils.git

# 4. Crea una nueva rama de característica
git checkout -b feature/tu-caracteristica-genial
```

**Configurar Entorno de Desarrollo**
```bash
# Instalar dependencias
go mod tidy

# Verificar entorno
go version  # Requiere Go 1.24.0+
go test ./... # Asegurar que todas las pruebas pasen
```

#### 2. Fase de Desarrollo

**Escribir Código**
- Seguir [Estándares de Código](#estándares-de-código)
- Escribir casos de prueba para nuevas características
- Asegurar que la cobertura de pruebas no baje del nivel actual
- Añadir comentarios de documentación necesarios

**Estándares de Commit**
```bash
# Usar formato de mensaje de commit estandarizado
git commit -m "feat(module): añadir nueva función utilitaria

- Añadir función FormatDuration
- Soportar múltiples formatos de salida de tiempo
- Añadir casos de prueba completos
- Actualizar documentación relacionada

Closes #123"
```

**Formato de Mensaje de Commit**:
```
<tipo>(<ámbito>): <asunto>

<cuerpo>

<pie de página>
```

**Categorías de Tipo**:
- `feat`: Nuevas características
- `fix`: Correcciones de errores  
- `docs`: Actualizaciones de documentación
- `style`: Ajustes de formateo de código
- `refactor`: Refactorización de código
- `perf`: Optimización de rendimiento
- `test`: Relacionado con pruebas
- `chore`: Actualizaciones de herramientas de build o dependencias

**Rango de Ámbito** (opcional):
- `candy`: módulo candy
- `xtime`: módulo xtime
- `config`: módulo config
- `cryptox`: módulo cryptox
- etc...

#### 3. Pruebas y Validación

**Ejecutar Pruebas**
```bash
# Ejecutar todas las pruebas
go test -v ./...

# Verificar cobertura de pruebas
go test -cover -v ./...

# Ejecutar pruebas de benchmark
go test -bench=. ./...

# Verificar formateo de código
go fmt ./...

# Análisis estático
go vet ./...
```

**Pruebas de Rendimiento**
```bash
# Ejecutar pruebas de rendimiento
go test -bench=BenchmarkTuFuncion -benchmem ./...

# Asegurar que no hay regresión significativa de rendimiento
```

#### 4. Crear Pull Request

**Push a Tu Fork**
```bash
git push origin feature/tu-caracteristica-genial
```

**Crear PR**
1. Visita la página del proyecto en GitHub
2. Haz clic en "New Pull Request"
3. Selecciona tu rama
4. Completa la descripción del PR (consulta [Plantilla de PR](#plantilla-de-pr))
5. Asegura que todas las verificaciones pasen

#### 5. Revisión de Código

- Los mantenedores revisarán tu código
- Haz modificaciones basadas en retroalimentación
- Mantén comunicación y actitud cooperativa
- Será fusionado después de que las pruebas pasen

## 📝 Estándares de Código

### Estilo de Código Go

**Estándares Básicos**
```go
// ✅ Buen ejemplo
package candy

import (
    "context"
    "fmt"
    "time"
    
    "github.com/lazygophers/log"
)

// FormatDuration formatea duración de tiempo en cadena legible por humanos
// Soporta múltiples niveles de precisión, elige automáticamente unidades apropiadas
//
// Parámetros:
//   - duration: duración de tiempo a formatear
//   - precision: nivel de precisión (1-3)
//
// Retorna:
//   - string: cadena formateada, como "2 horas 30 minutos"
//
// Ejemplo:
//   FormatDuration(90*time.Minute, 2) // retorna "1 hora 30 minutos"
//   FormatDuration(45*time.Second, 1) // retorna "45 segundos"
func FormatDuration(duration time.Duration, precision int) string {
    if duration == 0 {
        return "0 segundos"
    }
    
    // Lógica de implementación...
    return result
}
```

**Convenciones de Nomenclatura**
- Usar CamelCase
- Nombres de función empiezan con verbos: `Get`, `Set`, `Format`, `Parse`
- Constantes usan ALL_CAPS: `const MaxRetries = 3`
- Miembros privados usan minúsculas: `internalHelper`
- Nombres de paquete usan palabras únicas en minúsculas: `candy`, `xtime`

**Estándares de Comentarios**
- Todas las funciones públicas deben tener comentarios
- Comentarios empiezan con nombre de función
- Incluir descripciones de parámetros y valores de retorno  
- Proporcionar ejemplos de uso
- Comentarios en inglés, concisos y claros

**Manejo de Errores**
```go
// ✅ Enfoque recomendado de manejo de errores
func ProcessData(data []byte) (*Result, error) {
    if len(data) == 0 {
        log.Warn("Datos vacíos proporcionados")
        return nil, fmt.Errorf("los datos no pueden estar vacíos")
    }
    
    result, err := parseData(data)
    if err != nil {
        log.Error("Falló al analizar datos", log.Error(err))
        return nil, fmt.Errorf("análisis de datos falló: %w", err)
    }
    
    return result, nil
}
```

### Estándares de Estructura de Proyecto

**Organización de Módulos**
```
utils/
├── README.md           # Resumen del proyecto
├── CONTRIBUTING.md     # Guía de contribución  
├── SECURITY.md        # Política de seguridad
├── go.mod             # Definición del módulo Go
├── must.go            # Funciones utilitarias principales
├── candy/             # Herramientas de procesamiento de datos
│   ├── README.md      # Documentación del módulo
│   ├── to_string.go   # Conversión de tipos
│   └── to_string_test.go
├── xtime/             # Herramientas de procesamiento de tiempo  
│   ├── README.md      # Documentación de uso detallada
│   ├── TESTING.md     # Reportes de pruebas
│   ├── PERFORMANCE.md # Reportes de rendimiento
│   ├── calendar.go    # Funcionalidad de calendario
│   └── calendar_test.go
└── ...
```

**Nomenclatura de Archivos**
- Usar letras minúsculas y guiones bajos: `to_string.go`
- Sufijo de archivo de prueba: `_test.go`
- Pruebas de benchmark: `_benchmark_test.go`
- Archivos de documentación: `README.md`, `TESTING.md`

### Estándares de Pruebas

**Requisitos de Cobertura de Pruebas**
- Cobertura de pruebas de nuevas características debe ser ≥ 90%
- No puede reducir la cobertura general de pruebas
- Incluir casos normales y casos límite
- Rutas de manejo de errores deben ser probadas

**Ejemplo de Prueba**
```go
func TestFormatDuration(t *testing.T) {
    testCases := []struct {
        name      string
        duration  time.Duration
        precision int
        want      string
    }{
        {
            name:      "tiempo cero",
            duration:  0,
            precision: 1,
            want:      "0 segundos",
        },
        {
            name:      "90 minutos alta precisión",
            duration:  90 * time.Minute,
            precision: 2,
            want:      "1 hora 30 minutos",
        },
        // Más casos de prueba...
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            got := FormatDuration(tc.duration, tc.precision)
            assert.Equal(t, tc.want, got)
        })
    }
}

// Prueba de benchmark
func BenchmarkFormatDuration(b *testing.B) {
    duration := 90 * time.Minute
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = FormatDuration(duration, 2)
    }
}
```

## 🎯 Áreas Clave de Desarrollo

### Alta Prioridad

1. **Mejora del Módulo xtime**
   - Mejora de funcionalidad de calendario lunar y términos solares
   - Optimización de rendimiento
   - Más características específicas culturales

2. **Extensión del Módulo candy**  
   - Funciones de conversión de tipos
   - Herramientas de procesamiento de datos
   - Optimización de rendimiento

3. **Mejora de Cobertura de Pruebas**
   - Meta: Todos los módulos > 90%
   - Suplemento de casos límite
   - Mejora de pruebas de rendimiento

### Prioridad Media

4. **Nuevos Módulos Utilitarios**
   - Funciones utilitarias AI/ML
   - Herramientas de integración de servicios en la nube
   - Herramientas de microservicios

5. **Mejora de Documentación**
   - Documentación de referencia API
   - Guía de mejores prácticas
   - Guía de optimización de rendimiento

### Contribuciones Bienvenidas

- 🌏 **Soporte Multi-idioma** - Documentación en inglés, internacionalización de mensajes de error
- 📊 **Más Soporte de Formatos de Datos** - Procesamiento XML, YAML, TOML
- 🔧 **Herramientas de Desarrollo** - Generación de código, gestión de configuración
- 🎨 **Herramientas UI/UX** - Procesamiento de colores, salida formateada
- 🔐 **Herramientas de Seguridad** - Cifrado/descifrado, verificación de firma

## 📋 Plantilla de PR

Por favor usa la siguiente plantilla al crear un PR:

```markdown
## Descripción del Cambio

Breve descripción del contenido y propósito de este cambio.

## Tipo de Cambio

- [ ] Corrección de error
- [ ] Nueva característica
- [ ] Actualización de documentación
- [ ] Optimización de rendimiento  
- [ ] Refactorización de código
- [ ] Mejora de pruebas

## Cambios Detallados

### Nuevas Características
- Añadida función `FormatDuration`
- Soporte para múltiples niveles de precisión
- Añadido display de unidades de tiempo en chino

### Problemas Corregidos  
- Corregido error de conversión de zona horaria (#123)
- Resuelto problema de fuga de memoria

### Optimización de Rendimiento
- Optimizado rendimiento de concatenación de cadenas
- Reducida asignación de memoria en 30%

## Descripción de Pruebas

- [ ] Todas las pruebas pasan
- [ ] Añadidos nuevos casos de prueba
- [ ] Cobertura de pruebas ≥ 90%
- [ ] Pruebas de benchmark pasan

**Cobertura de Pruebas**: 92.5%

## Actualizaciones de Documentación

- [ ] Actualizado README.md
- [ ] Añadidos comentarios de función
- [ ] Actualizado código de ejemplo

## Compatibilidad

- [ ] Compatible hacia atrás
- [ ] Requiere actualización de versión (explicar razón)
- [ ] Cambios que rompen compatibilidad (explicación detallada)

## Lista de Verificación

- [ ] Código sigue estándares del proyecto
- [ ] Pasó verificación de formato `go fmt`
- [ ] Pasó verificación estática `go vet`
- [ ] Todas las pruebas pasan
- [ ] Documentación actualizada
- [ ] Mensajes de commit siguen estándares

## Problemas Relacionados

Closes #123
Refs #456

## Capturas de Pantalla/Demo

Proporciona capturas de pantalla o demos si es necesario.
```

## 🐛 Reportes de Errores

¿Encontraste un error? Por favor usa la siguiente plantilla para crear un Issue:

```markdown
## Descripción del Error

Breve descripción del problema encontrado.

## Pasos de Reproducción

1. Ejecutar paso 1
2. Ejecutar paso 2  
3. Observar resultado

## Comportamiento Esperado

Describe el comportamiento correcto que esperas ver.

## Comportamiento Actual

Describe el comportamiento erróneo actual observado.

## Información del Entorno

- **Sistema Operativo**: macOS 12.0
- **Versión de Go**: 1.24.0
- **Versión de Utils**: v1.2.0
- **Otra información relevante**:

## Logs de Error

```
pegar logs de error aquí
```

## Ejemplo Mínimo Reproducible

```go
package main

import (
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // Código mínimo de reproducción de error
}
```
```

## ✨ Solicitudes de Características

¿Quieres una nueva característica? Por favor usa la siguiente plantilla:

```markdown
## Descripción de la Característica

Describe la característica que te gustaría añadir.

## Casos de Uso

Describe cuándo se usaría esta característica.

## Diseño de API Sugerido

```go
// Firma de función sugerida y uso
func NewAwesomeFunction(param string) (Result, error) {
    // ...
}
```

## Soluciones Alternativas

¿Has considerado otras soluciones?

## Información Adicional

Otra información relevante o referencias.
```

## 🏆 Reconocimiento de Contribuyentes

### Reconocimiento por Tipo de Contribución

Daremos diferentes reconocimientos basados en tipos de contribución:

- 🥇 **Contribuyentes Principales** - Activos a largo plazo, contribuciones importantes de características
- 🥈 **Contribuyentes Activos** - Múltiples contribuciones valiosas  
- 🥉 **Contribuyentes de Comunidad** - Correcciones de errores, mejoras de documentación
- 🌟 **Primeros Contribuyentes** - Bienvenida a primeras contribuciones

### Estadísticas de Contribución

Mostraremos contribuyentes en los siguientes lugares:

- Lista de contribuyentes README.md
- Reconocimientos en notas de lanzamiento
- Sitio web del proyecto (si está disponible)
- Reportes anuales de contribuyentes

## 💬 Comunicación

### Obtener Ayuda

- 📖 **Problemas de Documentación**: Revisa README.md para cada módulo
- 🐛 **Reportes de Errores**: [GitHub Issues](https://github.com/lazygophers/utils/issues)
- 💡 **Discusiones de Características**: [GitHub Discussions](https://github.com/lazygophers/utils/discussions)
- ❓ **Preguntas de Uso**: [GitHub Discussions Q&A](https://github.com/lazygophers/utils/discussions/categories/q-a)

### Estándares de Discusión

Por favor sigue estos estándares de comunicación:

- Usa lenguaje amigable y profesional
- Proporciona descripciones detalladas de problemas y sugerencias
- Proporciona información de contexto suficiente
- Respeta diferentes puntos de vista y opiniones
- Participa activamente en discusiones constructivas

## 📜 Licencia

Este proyecto está licenciado bajo [GNU Affero General Public License v3.0](LICENSE).

**Contribuir significa estar de acuerdo**:
- Posees los derechos de autor del código enviado
- Aceptas liberar código bajo licencia AGPL v3.0
- Sigues el código de conducta de contribuyentes del proyecto

## 🙏 Reconocimientos

¡Gracias a todos los desarrolladores que han contribuido al proyecto LazyGophers Utils!

**Agradecimientos Especiales**:
- Todos los contribuyentes que enviaron Issues y PRs
- Miembros de la comunidad que proporcionaron sugerencias y retroalimentación
- Voluntarios que ayudaron a mejorar la documentación

---

**Disponible en otros idiomas:** [English](CONTRIBUTING.md) | [简体中文](CONTRIBUTING_zh.md) | [繁體中文](CONTRIBUTING_zh-Hant.md) | [Français](CONTRIBUTING_fr.md) | [Русский](CONTRIBUTING_ru.md) | [العربية](CONTRIBUTING_ar.md)

**¡Happy Coding! 🎉**

¡Siéntete libre de contactar al equipo de mantenedores en cualquier momento si tienes preguntas. Estamos felices de ayudarte a comenzar tu viaje de contribución!