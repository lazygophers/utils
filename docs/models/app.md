# App Package Documentation

<!-- Language selector -->
[🇺🇸 English](#english) | [🇨🇳 简体中文](#简体中文) | [🇭🇰 繁體中文](#繁體中文) | [🇷🇺 Русский](#русский) | [🇫🇷 Français](#français) | [🇸🇦 العربية](#العربية) | [🇪🇸 Español](#español)

---

## English

### Overview
The `app` package provides application lifecycle management and environment information utilities for Go applications. It handles build-time information, release types, and application metadata management.

### Key Features
- **Release Type Management**: Support for Debug, Test, Alpha, Beta, and Release environments
- **Build Information**: Access to Git commit, branch, and build metadata
- **Application Metadata**: Organization, name, version, and description management
- **Environment Variables**: Comprehensive build environment information
- **Build Tags**: Conditional compilation based on release type

### Core Components

#### Release Types
```go
type ReleaseType uint8

const (
    Debug ReleaseType = iota
    Test
    Alpha
    Beta
    Release
)
```

#### Application Information
```go
var (
    Organization = "lazygophers"  // Default organization
    Name         string          // Application name
    Version      string          // Application version
    Description  string          // Application description
    PackageType  ReleaseType     // Current release type
)
```

#### Build Environment Information
```go
var (
    Commit      string  // Full Git commit hash
    ShortCommit string  // Short Git commit hash
    Branch      string  // Git branch name
    Tag         string  // Git tag

    BuildDate   string  // Build timestamp

    GoVersion   string  // Go compiler version
    GoOS        string  // Target operating system
    Goarch      string  // Target architecture
    Goarm       string  // ARM version (if applicable)
    Goamd64     string  // AMD64 version (if applicable)
    Gomips      string  // MIPS version (if applicable)
)
```

### Usage Examples

#### Basic Application Setup
```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/app"
)

func init() {
    app.Name = "MyApplication"
    app.Version = "1.0.0"
    app.Description = "A sample application"
}

func main() {
    fmt.Printf("Application: %s v%s\n", app.Name, app.Version)
    fmt.Printf("Organization: %s\n", app.Organization)
    fmt.Printf("Release Type: %s\n", app.PackageType.String())
}
```

#### Release Type Detection
```go
func configureLogging() {
    switch app.PackageType {
    case app.Debug:
        log.SetLevel(log.DebugLevel)
    case app.Test:
        log.SetLevel(log.WarnLevel)
    case app.Alpha, app.Beta:
        log.SetLevel(log.InfoLevel)
    case app.Release:
        log.SetLevel(log.ErrorLevel)
    }
}
```

#### Build Information Display
```go
func showBuildInfo() {
    fmt.Printf("Build Information:\n")
    fmt.Printf("  Commit: %s\n", app.Commit)
    fmt.Printf("  Branch: %s\n", app.Branch)
    fmt.Printf("  Build Date: %s\n", app.BuildDate)
    fmt.Printf("  Go Version: %s\n", app.GoVersion)
    fmt.Printf("  Target: %s/%s\n", app.GoOS, app.Goarch)
}
```

### Build Tag Usage

#### Conditional Compilation
```go
//go:build debug
package main

func init() {
    enableDebugFeatures()
}
```

```go
//go:build release
package main

func init() {
    enableProductionOptimizations()
}
```

#### Build Commands
```bash
# Build for debug environment
go build -tags debug

# Build for test environment  
go build -tags test

# Build for alpha environment
go build -tags alpha

# Build for beta environment
go build -tags beta

# Build for release environment
go build -tags release
```

### Advanced Features

#### Version Management with Build Information
```go
func GetVersionInfo() map[string]string {
    return map[string]string{
        "name":         app.Name,
        "version":      app.Version,
        "organization": app.Organization,
        "release_type": app.PackageType.String(),
        "commit":       app.ShortCommit,
        "branch":       app.Branch,
        "build_date":   app.BuildDate,
        "go_version":   app.GoVersion,
    }
}
```

#### Environment-Specific Configuration
```go
func getConfigFile() string {
    switch app.PackageType {
    case app.Debug:
        return "config.debug.json"
    case app.Test:
        return "config.test.json"
    case app.Alpha:
        return "config.alpha.json"
    case app.Beta:
        return "config.beta.json"
    case app.Release:
        return "config.production.json"
    default:
        return "config.json"
    }
}
```

### Integration with Build Systems

#### Makefile Integration
```makefile
VERSION := $(shell git describe --tags --always --dirty)
COMMIT := $(shell git rev-parse HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)

build:
	go build -ldflags "-X github.com/lazygophers/utils/app.Version=$(VERSION) \
	                   -X github.com/lazygophers/utils/app.Commit=$(COMMIT) \
	                   -X github.com/lazygophers/utils/app.Branch=$(BRANCH) \
	                   -X github.com/lazygophers/utils/app.BuildDate=$(BUILD_DATE)" \
	         -tags release
```

### Best Practices
1. **Set Application Metadata**: Always initialize Name, Version, and Description
2. **Use Appropriate Release Types**: Choose the correct build tag for your environment
3. **Leverage Build Information**: Use commit and build date for debugging and support
4. **Environment-Specific Logic**: Use release types for conditional behavior
5. **Version Display**: Include version information in help/about commands

### Common Patterns
```go
// Application startup banner
func printBanner() {
    fmt.Printf(`
%s v%s (%s)
Organization: %s
Built: %s from %s
Go: %s on %s/%s
`, 
        app.Name, 
        app.Version, 
        app.PackageType.String(),
        app.Organization,
        app.BuildDate,
        app.ShortCommit,
        app.GoVersion,
        app.GoOS,
        app.Goarch,
    )
}

// Health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
    health := map[string]interface{}{
        "status":       "ok",
        "version":      app.Version,
        "release_type": app.PackageType.String(),
        "commit":       app.ShortCommit,
        "build_date":   app.BuildDate,
    }
    
    json.NewEncoder(w).Encode(health)
}
```

---

## 简体中文

### 概述
`app` 包为 Go 应用程序提供应用生命周期管理和环境信息工具。它处理构建时信息、发布类型和应用元数据管理。

### 主要特性
- **发布类型管理**: 支持 Debug、Test、Alpha、Beta 和 Release 环境
- **构建信息**: 访问 Git 提交、分支和构建元数据
- **应用元数据**: 组织、名称、版本和描述管理
- **环境变量**: 综合构建环境信息
- **构建标签**: 基于发布类型的条件编译

### 核心组件

#### 发布类型
```go
type ReleaseType uint8

const (
    Debug ReleaseType = iota
    Test
    Alpha
    Beta
    Release
)
```

#### 应用信息
```go
var (
    Organization = "lazygophers"  // 默认组织
    Name         string          // 应用名称
    Version      string          // 应用版本
    Description  string          // 应用描述
    PackageType  ReleaseType     // 当前发布类型
)
```

### 使用示例

#### 基本应用设置
```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/app"
)

func init() {
    app.Name = "我的应用"
    app.Version = "1.0.0"
    app.Description = "示例应用程序"
}

func main() {
    fmt.Printf("应用程序: %s v%s\n", app.Name, app.Version)
    fmt.Printf("组织: %s\n", app.Organization)
    fmt.Printf("发布类型: %s\n", app.PackageType.String())
}
```

#### 发布类型检测
```go
func configureLogging() {
    switch app.PackageType {
    case app.Debug:
        log.SetLevel(log.DebugLevel)
    case app.Test:
        log.SetLevel(log.WarnLevel)
    case app.Alpha, app.Beta:
        log.SetLevel(log.InfoLevel)
    case app.Release:
        log.SetLevel(log.ErrorLevel)
    }
}
```

### 最佳实践
1. **设置应用元数据**: 始终初始化 Name、Version 和 Description
2. **使用合适的发布类型**: 为您的环境选择正确的构建标签
3. **利用构建信息**: 使用提交信息和构建日期进行调试和支持
4. **环境特定逻辑**: 使用发布类型进行条件行为控制

---

## 繁體中文

### 概述
`app` 套件為 Go 應用程式提供應用生命週期管理和環境資訊工具。它處理建置時資訊、發布型別和應用程式元資料管理。

### 主要特性
- **發布型別管理**: 支援 Debug、Test、Alpha、Beta 和 Release 環境
- **建置資訊**: 存取 Git 提交、分支和建置元資料
- **應用程式元資料**: 組織、名稱、版本和描述管理
- **環境變數**: 綜合建置環境資訊

### 核心組件
```go
var (
    Organization = "lazygophers"  // 預設組織
    Name         string          // 應用程式名稱
    Version      string          // 應用程式版本
    Description  string          // 應用程式描述
    PackageType  ReleaseType     // 目前發布型別
)
```

### 使用範例
```go
func init() {
    app.Name = "我的應用程式"
    app.Version = "1.0.0"
    app.Description = "範例應用程式"
}
```

### 最佳實務
1. **設定應用程式元資料**: 始終初始化 Name、Version 和 Description
2. **使用適當的發布型別**: 為您的環境選擇正確的建置標籤

---

## Русский

### Обзор
Пакет `app` предоставляет утилиты для управления жизненным циклом приложения и информации об окружении для Go-приложений. Он обрабатывает информацию времени сборки, типы релизов и управление метаданными приложения.

### Основные возможности
- **Управление типами релизов**: Поддержка окружений Debug, Test, Alpha, Beta и Release
- **Информация о сборке**: Доступ к Git commit, ветке и метаданным сборки
- **Метаданные приложения**: Управление организацией, именем, версией и описанием
- **Переменные окружения**: Комплексная информация об окружении сборки

### Основные компоненты
```go
var (
    Organization = "lazygophers"  // Организация по умолчанию
    Name         string          // Имя приложения
    Version      string          // Версия приложения
    Description  string          // Описание приложения
    PackageType  ReleaseType     // Текущий тип релиза
)
```

### Примеры использования
```go
func init() {
    app.Name = "Мое приложение"
    app.Version = "1.0.0"
    app.Description = "Пример приложения"
}
```

### Лучшие практики
1. **Установка метаданных приложения**: Всегда инициализируйте Name, Version и Description
2. **Использование подходящих типов релизов**: Выберите правильный тег сборки для вашего окружения

---

## Français

### Aperçu
Le package `app` fournit des utilitaires de gestion du cycle de vie des applications et d'informations sur l'environnement pour les applications Go. Il gère les informations de temps de construction, les types de version et la gestion des métadonnées d'application.

### Caractéristiques principales
- **Gestion des types de version**: Support pour les environnements Debug, Test, Alpha, Beta et Release
- **Informations de construction**: Accès au commit Git, à la branche et aux métadonnées de construction
- **Métadonnées d'application**: Gestion de l'organisation, du nom, de la version et de la description
- **Variables d'environnement**: Informations complètes sur l'environnement de construction

### Composants principaux
```go
var (
    Organization = "lazygophers"  // Organisation par défaut
    Name         string          // Nom de l'application
    Version      string          // Version de l'application
    Description  string          // Description de l'application
    PackageType  ReleaseType     // Type de version actuel
)
```

### Exemples d'utilisation
```go
func init() {
    app.Name = "Mon application"
    app.Version = "1.0.0"
    app.Description = "Application d'exemple"
}
```

### Meilleures pratiques
1. **Définir les métadonnées de l'application**: Toujours initialiser Name, Version et Description
2. **Utiliser les types de version appropriés**: Choisir le bon tag de construction pour votre environnement

---

## العربية

### نظرة عامة
توفر حزمة `app` أدوات إدارة دورة حياة التطبيق ومعلومات البيئة لتطبيقات Go. تتعامل مع معلومات وقت البناء، وأنواع الإصدارات، وإدارة البيانات الوصفية للتطبيق.

### الميزات الرئيسية
- **إدارة أنواع الإصدار**: دعم لبيئات Debug و Test و Alpha و Beta و Release
- **معلومات البناء**: الوصول إلى Git commit والفرع والبيانات الوصفية للبناء
- **البيانات الوصفية للتطبيق**: إدارة المؤسسة والاسم والإصدار والوصف
- **متغيرات البيئة**: معلومات شاملة عن بيئة البناء

### المكونات الأساسية
```go
var (
    Organization = "lazygophers"  // المؤسسة الافتراضية
    Name         string          // اسم التطبيق
    Version      string          // إصدار التطبيق
    Description  string          // وصف التطبيق
    PackageType  ReleaseType     // نوع الإصدار الحالي
)
```

### أمثلة الاستخدام
```go
func init() {
    app.Name = "تطبيقي"
    app.Version = "1.0.0"
    app.Description = "تطبيق مثال"
}
```

### أفضل الممارسات
1. **تعيين البيانات الوصفية للتطبيق**: قم دائماً بتهيئة Name و Version و Description
2. **استخدام أنواع الإصدار المناسبة**: اختر علامة البناء الصحيحة لبيئتك

---

## Español

### Descripción general
El paquete `app` proporciona utilidades de gestión del ciclo de vida de aplicaciones e información del entorno para aplicaciones Go. Maneja información de tiempo de construcción, tipos de lanzamiento y gestión de metadatos de aplicación.

### Características principales
- **Gestión de tipos de lanzamiento**: Soporte para entornos Debug, Test, Alpha, Beta y Release
- **Información de construcción**: Acceso a Git commit, rama y metadatos de construcción
- **Metadatos de aplicación**: Gestión de organización, nombre, versión y descripción
- **Variables de entorno**: Información completa del entorno de construcción

### Componentes principales
```go
var (
    Organization = "lazygophers"  // Organización por defecto
    Name         string          // Nombre de la aplicación
    Version      string          // Versión de la aplicación
    Description  string          // Descripción de la aplicación
    PackageType  ReleaseType     // Tipo de lanzamiento actual
)
```

### Ejemplos de uso
```go
func init() {
    app.Name = "Mi aplicación"
    app.Version = "1.0.0"
    app.Description = "Aplicación de ejemplo"
}
```

### Mejores prácticas
1. **Establecer metadatos de aplicación**: Siempre inicializar Name, Version y Description
2. **Usar tipos de lanzamiento apropiados**: Elegir la etiqueta de construcción correcta para su entorno