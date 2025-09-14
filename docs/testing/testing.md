# Testing Documentation

<!-- Language selector -->
[🇺🇸 English](#english) | [🇨🇳 简体中文](#简体中文) | [🇭🇰 繁體中文](#繁體中文) | [🇷🇺 Русский](#русский) | [🇫🇷 Français](#français) | [🇸🇦 العربية](#العربية) | [🇪🇸 Español](#español)

---

## English

### Overview
This document provides comprehensive testing guidelines, strategies, and best practices for the LazyGophers Utils library. It covers unit testing, integration testing, benchmarking, and test automation.

### Testing Philosophy
- **Comprehensive Coverage**: Maintain >85% test coverage across all modules
- **Quality Assurance**: Every public function must have corresponding tests
- **Performance Validation**: Benchmark critical operations for performance regressions
- **Cross-Platform Testing**: Ensure compatibility across different operating systems and architectures

### Test Structure

#### Project Test Organization
```
utils/
├── candy/
│   ├── string_test.go
│   ├── math_test.go
│   └── collection_test.go
├── json/
│   ├── marshal_test.go
│   ├── unmarshal_test.go
│   └── file_test.go
└── routine/
    ├── goroutine_test.go
    └── lifecycle_test.go
```

### Testing Commands

#### Basic Test Execution
```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for specific package
go test ./candy

# Run specific test function
go test -run TestToString ./candy
```

#### Coverage Analysis
```bash
# Generate coverage report
go test -cover ./...

# Detailed coverage with HTML report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Coverage threshold check
go test -cover ./... | grep -E "(PASS|FAIL)" | grep -v "coverage: 100.0%"
```

#### Benchmark Testing
```bash
# Run benchmarks
go test -bench=. ./...

# Run specific benchmark
go test -bench=BenchmarkToString ./candy

# Memory allocation profiling
go test -bench=. -benchmem ./...

# CPU profiling
go test -bench=. -cpuprofile=cpu.prof ./...
```

### Testing Patterns

#### Unit Testing Best Practices
```go
func TestToString(t *testing.T) {
    tests := []struct {
        name     string
        input    interface{}
        expected string
    }{
        {"integer", 42, "42"},
        {"float", 3.14, "3.14"},
        {"string", "hello", "hello"},
        {"boolean", true, "true"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := candy.ToString(tt.input)
            if result != tt.expected {
                t.Errorf("ToString(%v) = %v, want %v", 
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

#### Table-Driven Tests
```go
func TestMapAny_Get(t *testing.T) {
    testCases := []struct {
        name          string
        data          map[string]interface{}
        key           string
        expected      interface{}
        expectError   bool
    }{
        {
            name:        "existing key",
            data:        map[string]interface{}{"key": "value"},
            key:         "key",
            expected:    "value",
            expectError: false,
        },
        {
            name:        "non-existing key",
            data:        map[string]interface{}{},
            key:         "missing",
            expected:    nil,
            expectError: true,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            m := anyx.NewMap(tc.data)
            result, err := m.Get(tc.key)
            
            if tc.expectError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tc.expected, result)
            }
        })
    }
}
```

#### Mock and Stub Testing
```go
type MockLogger struct {
    logs []string
}

func (m *MockLogger) Log(message string) {
    m.logs = append(m.logs, message)
}

func TestRoutineWithMockLogger(t *testing.T) {
    mockLogger := &MockLogger{}
    
    routine.Go(func() error {
        mockLogger.Log("test message")
        return nil
    })
    
    // Wait for goroutine completion
    time.Sleep(100 * time.Millisecond)
    
    assert.Contains(t, mockLogger.logs, "test message")
}
```

### Performance Testing

#### Benchmark Examples
```go
func BenchmarkJSONMarshal(b *testing.B) {
    data := map[string]interface{}{
        "name":   "John Doe",
        "age":    30,
        "active": true,
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := json.Marshal(data)
        if err != nil {
            b.Fatal(err)
        }
    }
}

func BenchmarkCandyToString(b *testing.B) {
    values := []interface{}{42, 3.14, "hello", true}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for _, v := range values {
            _ = candy.ToString(v)
        }
    }
}
```

#### Memory Profiling
```go
func BenchmarkMemoryUsage(b *testing.B) {
    b.ReportAllocs()
    
    for i := 0; i < b.N; i++ {
        data := make([]string, 1000)
        for j := range data {
            data[j] = candy.ToString(j)
        }
    }
}
```

### Integration Testing

#### Cross-Module Integration
```go
func TestConfigurationFlow(t *testing.T) {
    // Test complete configuration loading flow
    configData := `{
        "database": {
            "host": "localhost",
            "port": 5432
        }
    }`
    
    // Write test config file
    tmpFile, err := os.CreateTemp("", "test-config-*.json")
    require.NoError(t, err)
    defer os.Remove(tmpFile.Name())
    
    _, err = tmpFile.WriteString(configData)
    require.NoError(t, err)
    tmpFile.Close()
    
    // Test loading and parsing
    var config Config
    err = json.UnmarshalFromFile(tmpFile.Name(), &config)
    require.NoError(t, err)
    
    assert.Equal(t, "localhost", config.Database.Host)
    assert.Equal(t, 5432, config.Database.Port)
}
```

### Test Utilities

#### Helper Functions
```go
// Test helper for creating temporary files
func createTempJSON(t *testing.T, data interface{}) string {
    tmpFile, err := os.CreateTemp("", "test-*.json")
    require.NoError(t, err)
    
    err = json.NewEncoder(tmpFile).Encode(data)
    require.NoError(t, err)
    tmpFile.Close()
    
    t.Cleanup(func() {
        os.Remove(tmpFile.Name())
    })
    
    return tmpFile.Name()
}

// Test helper for goroutine synchronization
func waitForGoroutine(t *testing.T, timeout time.Duration, fn func()) {
    done := make(chan bool, 1)
    
    go func() {
        fn()
        done <- true
    }()
    
    select {
    case <-done:
        // Success
    case <-time.After(timeout):
        t.Fatal("Goroutine did not complete within timeout")
    }
}
```

### Test Configuration

#### GitHub Actions CI
```yaml
name: Tests
on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        go-version: [1.18, 1.19, 1.20, 1.21]
    
    steps:
    - uses: actions/checkout@v3
    - uses: actions/setup-go@v3
      with:
        go-version: ${{ matrix.go-version }}
    
    - name: Run tests
      run: go test -v -race -coverprofile=coverage.out ./...
    
    - name: Upload coverage
      uses: codecov/codecov-action@v3
      with:
        file: ./coverage.out
```

### Testing Makefile Targets
```makefile
# Test commands
.PHONY: test test-coverage test-race test-bench

test:
	go test ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

test-race:
	go test -race ./...

test-bench:
	go test -bench=. -benchmem ./...

test-integration:
	go test -tags=integration ./...

# Excluding PGP package from tests (as per project requirements)
TEST_PACKAGES := $(shell go list ./... | grep -v pgp)

test-excluding-pgp:
	go test $(TEST_PACKAGES)
```

### Quality Gates

#### Coverage Requirements
- **Minimum Coverage**: 85% across all packages
- **Critical Paths**: 95% coverage for core utilities
- **New Features**: 100% coverage requirement

#### Performance Benchmarks
```go
// Performance regression detection
func BenchmarkPerformanceRegression(b *testing.B) {
    baseline := 100 * time.Nanosecond // Expected performance
    
    start := time.Now()
    for i := 0; i < b.N; i++ {
        // Operation under test
        _ = candy.ToString(42)
    }
    elapsed := time.Since(start) / time.Duration(b.N)
    
    if elapsed > baseline {
        b.Errorf("Performance regression detected: %v > %v", elapsed, baseline)
    }
}
```

---

## 简体中文

### 概述
本文档为 LazyGophers Utils 库提供全面的测试指南、策略和最佳实践。涵盖单元测试、集成测试、基准测试和测试自动化。

### 测试理念
- **全面覆盖**: 在所有模块中维持 >85% 的测试覆盖率
- **质量保证**: 每个公共函数都必须有相应的测试
- **性能验证**: 对关键操作进行基准测试以检测性能回归
- **跨平台测试**: 确保在不同操作系统和架构上的兼容性

### 测试命令

#### 基本测试执行
```bash
# 运行所有测试
go test ./...

# 运行详细输出的测试
go test -v ./...

# 运行特定包的测试
go test ./candy

# 运行特定测试函数
go test -run TestToString ./candy
```

#### 覆盖率分析
```bash
# 生成覆盖率报告
go test -cover ./...

# 带 HTML 报告的详细覆盖率
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### 测试模式

#### 单元测试最佳实践
```go
func TestToString(t *testing.T) {
    tests := []struct {
        name     string
        input    interface{}
        expected string
    }{
        {"整数", 42, "42"},
        {"浮点数", 3.14, "3.14"},
        {"字符串", "hello", "hello"},
        {"布尔值", true, "true"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := candy.ToString(tt.input)
            if result != tt.expected {
                t.Errorf("ToString(%v) = %v, 期望 %v", 
                    tt.input, result, tt.expected)
            }
        })
    }
}
```

### 质量门槛

#### 覆盖率要求
- **最低覆盖率**: 所有包 85%
- **关键路径**: 核心工具 95% 覆盖率
- **新特性**: 100% 覆盖率要求

---

## 繁體中文

### 概述
本文件為 LazyGophers Utils 函式庫提供全面的測試指南、策略和最佳實務。涵蓋單元測試、整合測試、效能測試和測試自動化。

### 測試理念
- **全面覆蓋**: 在所有模組中維持 >85% 的測試覆蓋率
- **品質保證**: 每個公開函數都必須有相應的測試
- **效能驗證**: 對關鍵操作進行效能測試以檢測效能衰退
- **跨平台測試**: 確保在不同作業系統和架構上的相容性

### 測試命令
```bash
# 執行所有測試
go test ./...

# 執行詳細輸出的測試
go test -v ./...
```

### 品質門檻
- **最低覆蓋率**: 所有套件 85%
- **關鍵路徑**: 核心工具 95% 覆蓋率

---

## Русский

### Обзор
Этот документ предоставляет комплексные руководящие принципы тестирования, стратегии и лучшие практики для библиотеки LazyGophers Utils. Он охватывает модульное тестирование, интеграционное тестирование, бенчмаркинг и автоматизацию тестов.

### Философия тестирования
- **Комплексное покрытие**: Поддержание >85% покрытия тестами во всех модулях
- **Обеспечение качества**: Каждая публичная функция должна иметь соответствующие тесты
- **Проверка производительности**: Бенчмаркинг критических операций для выявления регрессий производительности

### Команды тестирования
```bash
# Запустить все тесты
go test ./...

# Запустить тесты с подробным выводом
go test -v ./...
```

### Пороги качества
- **Минимальное покрытие**: 85% по всем пакетам
- **Критические пути**: 95% покрытие для основных утилит

---

## Français

### Aperçu
Ce document fournit des directives de test complètes, des stratégies et des meilleures pratiques pour la bibliothèque LazyGophers Utils. Il couvre les tests unitaires, les tests d'intégration, le benchmarking et l'automatisation des tests.

### Philosophie de test
- **Couverture complète**: Maintenir >85% de couverture de test à travers tous les modules
- **Assurance qualité**: Chaque fonction publique doit avoir des tests correspondants
- **Validation de performance**: Benchmark des opérations critiques pour les régressions de performance

### Commandes de test
```bash
# Exécuter tous les tests
go test ./...

# Exécuter les tests avec sortie détaillée
go test -v ./...
```

### Seuils de qualité
- **Couverture minimale**: 85% pour tous les packages
- **Chemins critiques**: 95% de couverture pour les utilitaires principaux

---

## العربية

### نظرة عامة
توفر هذه الوثيقة إرشادات اختبار شاملة واستراتيجيات وأفضل الممارسات لمكتبة LazyGophers Utils. تغطي اختبار الوحدة، واختبار التكامل، والقياس المعياري، وأتمتة الاختبار.

### فلسفة الاختبار
- **التغطية الشاملة**: الحفاظ على >85% تغطية اختبار عبر جميع الوحدات
- **ضمان الجودة**: كل وظيفة عامة يجب أن تحتوي على اختبارات مقابلة
- **التحقق من الأداء**: قياس العمليات الحرجة لتراجع الأداء

### أوامر الاختبار
```bash
# تشغيل جميع الاختبارات
go test ./...

# تشغيل الاختبارات مع مخرجات مفصلة
go test -v ./...
```

### عتبات الجودة
- **الحد الأدنى للتغطية**: 85% عبر جميع الحزم
- **المسارات الحرجة**: 95% تغطية للأدوات الأساسية

---

## Español

### Descripción general
Este documento proporciona directrices de prueba integrales, estrategias y mejores prácticas para la biblioteca LazyGophers Utils. Cubre pruebas unitarias, pruebas de integración, benchmarking y automatización de pruebas.

### Filosofía de pruebas
- **Cobertura integral**: Mantener >85% de cobertura de pruebas en todos los módulos
- **Aseguramiento de calidad**: Cada función pública debe tener pruebas correspondientes
- **Validación de rendimiento**: Benchmark de operaciones críticas para regresiones de rendimiento

### Comandos de prueba
```bash
# Ejecutar todas las pruebas
go test ./...

# Ejecutar pruebas con salida detallada
go test -v ./...
```

### Umbrales de calidad
- **Cobertura mínima**: 85% en todos los paquetes
- **Rutas críticas**: 95% de cobertura para utilidades principales