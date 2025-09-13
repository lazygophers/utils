# Política de Seguridad

## Versiones Soportadas

Mantenemos activamente y proporcionamos actualizaciones de seguridad para las siguientes versiones de LazyGophers Utils:

| Versión | Soportada          | Versión Go Requerida | Estado           |
| ------- | ------------------ | -------------------- | ---------------- |
| 1.x.x   | :white_check_mark: | Go 1.24+            | Desarrollo Activo |
| 0.x.x   | :white_check_mark: | Go 1.24+            | Solo Correcciones de Seguridad |

## Consideraciones de Seguridad

### Componentes Criptográficos

Esta biblioteca incluye utilidades criptográficas en el paquete `cryptox`:

- **Cifrado/descifrado AES**: Utiliza la biblioteca criptográfica estándar de Go
- **Generación y operaciones de claves RSA**: Implementa prácticas estándar de la industria
- **Funciones hash**: SHA-256, SHA-512 y otros algoritmos seguros
- **Blowfish y ChaCha20**: Algoritmos de cifrado adicionales
- **PGP/GPG**: Implementación OpenPGP para mensajería segura

⚠️ **Importante**: Aunque seguimos las mejores prácticas de seguridad, siempre realice su propia revisión de seguridad antes de usar funciones criptográficas en entornos de producción.

### Validación de Entrada

Las funciones de utilidad en esta biblioteca (especialmente en los paquetes `candy`, `stringx` y `config`) realizan conversiones de tipos y análisis de datos. Aunque implementamos prácticas de programación defensiva:

- La sanitización de entrada se realiza cuando es aplicable
- Los fallos de conversión de tipos se manejan elegantemente
- La carga de configuración incluye pasos de validación

### Dependencias

Auditamos regularmente nuestras dependencias para vulnerabilidades conocidas:

- Todas las dependencias están fijadas a versiones específicas
- Usamos `go mod tidy` y `go mod verify` en nuestro pipeline CI/CD
- El escaneo de seguridad se realiza a través de golangci-lint y gosec

## Reportar una Vulnerabilidad

### Cómo Reportar

Si descubre una vulnerabilidad de seguridad en LazyGophers Utils, por favor repórtela responsablemente:

1. **Email**: Envíe detalles a `security@lazygophers.com` (si está disponible) o cree un problema privado
2. **Asesoría de Seguridad de GitHub**: Use la función de reporte privado de vulnerabilidades de GitHub
3. **Contacto Directo**: Contacte a los mantenedores directamente a través de GitHub

### Información a Incluir

Por favor incluya la siguiente información en su reporte:

- **Descripción**: Descripción clara de la vulnerabilidad
- **Ubicación**: Paquete/archivo/función específico afectado
- **Impacto**: Impacto potencial de seguridad y vectores de ataque
- **Reproducción**: Pasos para reproducir la vulnerabilidad
- **Corrección Sugerida**: Si tiene ideas para la remediación

### Cronograma de Respuesta

Estamos comprometidos a abordar los problemas de seguridad rápidamente:

- **Reconocimiento**: Dentro de 48 horas del reporte
- **Evaluación Inicial**: Dentro de 5 días hábiles
- **Cronograma de Resolución**:
  - Vulnerabilidades críticas: 7-14 días
  - Alta severidad: 14-30 días
  - Media/Baja severidad: 30-90 días

### Política de Divulgación

Seguimos prácticas de divulgación responsable:

1. Trabajaremos con usted para entender y reproducir el problema
2. Desarrollaremos y probaremos una corrección
3. Coordinaremos el tiempo de divulgación con usted
4. Se dará crédito a los investigadores que reporten vulnerabilidades responsablemente

## Mejores Prácticas de Seguridad para Usuarios

### Pautas Generales

- **Manténgase Actualizado**: Siempre use la versión más reciente de la biblioteca
- **Revise el Código**: Realice revisiones de seguridad para casos de uso en producción
- **Valide las Entradas**: Siempre valide las entradas externas en sus aplicaciones
- **Siga los Principios**: Aplique principios de seguridad de defensa en profundidad

### Uso Criptográfico

Al usar el paquete `cryptox`:

- **Gestión de Claves**: Use prácticas seguras de generación y almacenamiento de claves
- **Números Aleatorios**: Asegure la entropía adecuada para operaciones criptográficas
- **Selección de Algoritmo**: Elija algoritmos apropiados para su modelo de amenazas
- **Implementación**: Siga las mejores prácticas criptográficas en su aplicación

### Seguridad de Configuración

Al usar el paquete `config`:

- **Permisos de Archivo**: Restrinja el acceso a archivos de configuración que contengan secretos
- **Variables de Entorno**: Use métodos seguros para almacenar configuración sensible
- **Validación**: Siempre valide los datos de configuración antes del uso

## Pruebas de Seguridad

### Escaneo de Seguridad Automatizado

Nuestro pipeline CI/CD incluye:

- **Análisis Estático**: golangci-lint con linters enfocados en seguridad
- **Escaneo de Vulnerabilidades**: gosec para problemas de seguridad específicos de Go
- **Escaneo de Dependencias**: Verificaciones regulares de dependencias vulnerables
- **Calidad del Código**: Linting y pruebas exhaustivas

### Revisión Manual de Seguridad

Realizamos revisiones manuales de seguridad para:

- Todas las implementaciones criptográficas
- Lógica de validación y análisis de entrada
- Manejo de errores y divulgación de información
- Patrones de autenticación y autorización

## Información de Contacto

Para preguntas o preocupaciones relacionadas con seguridad:

- **Repositorio del Proyecto**: [https://github.com/lazygophers/utils](https://github.com/lazygophers/utils)
- **Problemas**: Use GitHub Issues para errores no relacionados con seguridad
- **Reportes de Seguridad**: Siga el proceso de reporte de vulnerabilidades anterior

## Registro de Cambios

### Actualizaciones de Seguridad

Mantenemos un registro de actualizaciones de seguridad para transparencia:

- **Versión 1.0.x**: Revisión inicial de seguridad y endurecimiento
- **Versiones futuras**: Las actualizaciones de seguridad se documentarán aquí

---

**Disponible en otros idiomas:** [English](SECURITY.md) | [简体中文](SECURITY_zh.md) | [繁體中文](SECURITY_zh-Hant.md) | [Français](SECURITY_fr.md) | [Русский](SECURITY_ru.md) | [العربية](SECURITY_ar.md)

**Nota**: Esta política de seguridad está sujeta a actualizaciones. Por favor verifique la versión más reciente en el repositorio.

Última actualización: 2025-09-13