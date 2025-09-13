# Contribuir a LazyGophers Utils

¡Gracias por tu interés en contribuir a LazyGophers Utils! Este documento proporciona directrices e información para los contribuyentes.

## 🚀 Comenzando

### Prerrequisitos

- Go 1.24.0 o más reciente
- Git
- Make (opcional, para automatización)

### Configuración de Desarrollo

1. **Fork y Clone**
   ```bash
   git clone https://github.com/your-username/utils.git
   cd utils
   ```

2. **Instalar Dependencias**
   ```bash
   go mod tidy
   ```

3. **Verificar Configuración**
   ```bash
   go test ./...
   ```

## 📋 Directrices de Desarrollo

### Estilo de Código

1. **Seguir Estándares de Go**
   - Usar `gofmt` para formateo
   - Seguir prácticas efectivas de Go
   - Usar nombres de variables y funciones significativos

2. **Directrices Específicas de Paquetes**
   - Cada paquete debe ser independiente y reutilizable
   - Minimizar dependencias externas
   - Usar genéricos para seguridad de tipos cuando sea apropiado

3. **Documentación**
   - Todas las funciones públicas deben tener comentarios de documentación en chino
   - Incluir ejemplos de uso para funciones complejas
   - Documentar características de rendimiento para funciones críticas

### Directrices de Rendimiento

1. **Optimización de Memoria**
   - Usar pools de objetos para operaciones de alta frecuencia
   - Preferir operaciones zero-copy cuando sea posible
   - Minimizar asignaciones de memoria en rutas calientes

2. **Concurrencia**
   - Usar operaciones atómicas en lugar de mutex cuando sea posible
   - Asegurar thread-safety para operaciones concurrentes
   - Diseñar algoritmos lock-free cuando sea apropiado

## 🧪 Requisitos de Pruebas

### Pruebas Unitarias

1. **Cobertura de Pruebas**
   - Apuntar a 90%+ de cobertura de pruebas para código nuevo
   - Probar rutas de éxito y error
   - Incluir casos extremos y condiciones límite

2. **Organización de Pruebas**
   ```bash
   # Ejecutar pruebas para paquete específico
   go test ./candy
   
   # Ejecutar con cobertura
   go test -cover ./...
   ```

## 📝 Directrices de Commit

### Formato de Mensaje de Commit

```
<type>(<scope>): <description>

<body>

<footer>
```

**Tipos:**
- `feat`: Nueva funcionalidad
- `fix`: Corrección de error
- `perf`: Mejora de rendimiento
- `refactor`: Refactorización de código
- `test`: Agregar o actualizar pruebas
- `docs`: Cambios de documentación

## 🔍 Proceso de Revisión de Código

### Directrices de Pull Request

1. **Antes de Enviar**
   - Asegurar que todas las pruebas pasen
   - Ejecutar `go fmt ./...`
   - Ejecutar `go vet ./...`
   - Actualizar documentación si es necesario

2. **Criterios de Revisión**
   - Calidad y legibilidad del código
   - Cobertura y calidad de pruebas
   - Impacto en el rendimiento
   - Cambios que rompen compatibilidad
   - Completitud de documentación

## 🏗️ Directrices de Arquitectura

### Diseño de Paquetes

1. **Responsabilidad Única**
   - Cada paquete debe tener un propósito claro y enfocado
   - Evitar mezclar funcionalidad no relacionada
   - Mantener APIs públicas mínimas y limpias

2. **Dependencias**
   - Minimizar dependencias externas
   - Preferir biblioteca estándar cuando sea posible
   - Documentar justificación de dependencias

## 📚 Documentación

### Documentación de API

```go
// ToString convierte cualquier tipo a string
// Soporta conversión de tipos básicos, slices, maps y structs
// Usa serialización JSON para tipos complejos
//
// Características de rendimiento:
// - Conversión tipos básicos: O(1)
// - Conversión tipos complejos: O(n) donde n es complejidad de serialización
func ToString(v interface{}) string
```

## 🐛 Directrices de Issues

### Reportes de Errores

```markdown
**Descripción del Error**
Descripción clara del error

**Pasos para Reproducir**
1. Primer paso
2. Segundo paso
3. Tercer paso

**Comportamiento Esperado**
Qué debería pasar

**Comportamiento Real**
Qué pasó realmente

**Entorno**
- Versión de Go:
- SO:
- Versión del paquete:
```

## 🤝 Directrices de Comunidad

### Código de Conducta

1. **Sé Respetuoso**
   - Trata a todos los contribuyentes con respeto
   - Sé constructivo en la retroalimentación
   - Da la bienvenida a los recién llegados

2. **Sé Colaborativo**
   - Comparte conocimiento y ayuda a otros
   - Proporciona revisiones claras y útiles
   - Comunícate abiertamente sobre desafíos

## 📖 Recursos de Aprendizaje

### Mejores Prácticas de Go
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Comentarios de Revisión de Código Go](https://github.com/golang/go/wiki/CodeReviewComments)

### Optimización de Rendimiento
- [Consejos de Rendimiento Go](https://github.com/golang/go/wiki/Performance)
- [Profiling de Programas Go](https://blog.golang.org/profiling-go-programs)

### Pruebas
- [Pruebas en Go](https://golang.org/doc/code.html#Testing)

## 📞 Obtener Ayuda

Si necesitas ayuda o tienes preguntas:

1. Revisa la documentación existente
2. Busca issues existentes
3. Crea un nuevo issue con descripción clara
4. Únete a nuestras discusiones comunitarias

¡Gracias por contribuir a LazyGophers Utils! Tus contribuciones ayudan a hacer esta biblioteca mejor para todos.