# Candy Package Documentation

<!-- Language selector -->
[🇺🇸 English](#english) | [🇨🇳 简体中文](#简体中文) | [🇭🇰 繁體中文](#繁體中文) | [🇷🇺 Русский](#русский) | [🇫🇷 Français](#français) | [🇸🇦 العربية](#العربية) | [🇪🇸 Español](#español)

---

## English

### Overview
The `candy` package provides syntactic sugar utilities for Go, simplifying common programming operations with type-safe generic functions.

### Key Features
- **Mathematical Operations**: Abs, Sqrt, Cbrt, Pow functions for numeric types
- **Collection Utilities**: Chunk, Contains, Each, Reduce functions for slices
- **Statistical Functions**: Average, Max, Min calculations
- **Sorting Utilities**: SortUsing for custom sorting logic
- **Random Operations**: Random element selection
- **Type Safety**: Full generic type support with constraints

### Core Functions

#### Mathematical Functions
```go
// Absolute value calculation
result := candy.Abs(-42) // Returns 42

// Square root
sqrt := candy.Sqrt(16.0) // Returns 4.0

// Cube root  
cbrt := candy.Cbrt(8.0) // Returns 2.0

// Power calculation
power := candy.Pow(2, 3) // Returns 8
```

#### Collection Operations
```go
// Chunk slice into smaller slices
chunks := candy.Chunk([]int{1,2,3,4,5,6}, 2) // Returns [[1,2], [3,4], [5,6]]

// Check if slice contains element
exists := candy.Contains([]string{"a", "b", "c"}, "b") // Returns true

// Apply function to each element
candy.Each([]int{1,2,3}, func(v int) { fmt.Println(v) })

// Reduce slice to single value
sum := candy.Reduce([]int{1,2,3,4}, 0, func(acc, v int) int { return acc + v }) // Returns 10
```

#### Statistical Functions
```go
// Calculate average
avg := candy.Average([]float64{1.0, 2.0, 3.0, 4.0}) // Returns 2.5

// Find maximum value
max := candy.Max([]int{3, 7, 2, 9, 1}) // Returns 9
```

### Best Practices
1. Use appropriate numeric constraints (Integer/Float) for mathematical operations
2. Handle empty slices appropriately in collection functions
3. Leverage generic types for type safety
4. Consider performance implications for large datasets

### Common Patterns
```go
// Pipeline processing
result := candy.Chunk(
    candy.Reduce(numbers, []int{}, filter),
    batchSize,
)

// Statistical analysis
stats := StatInfo{
    Average: candy.Average(data),
    Max:     candy.Max(data),
    Min:     candy.Min(data),
}
```

---

## 简体中文

### 概述
`candy` 包为 Go 提供语法糖工具，通过类型安全的泛型函数简化常见编程操作。

### 主要特性
- **数学运算**: 支持绝对值、平方根、立方根、幂运算等数值类型操作
- **集合工具**: 提供切片的分块、包含、遍历、归约等功能
- **统计函数**: 平均值、最大值、最小值计算
- **排序工具**: 自定义排序逻辑的 SortUsing
- **随机操作**: 随机元素选择
- **类型安全**: 完整的泛型类型支持和约束

### 核心函数

#### 数学函数
```go
// 绝对值计算
result := candy.Abs(-42) // 返回 42

// 平方根
sqrt := candy.Sqrt(16.0) // 返回 4.0

// 立方根
cbrt := candy.Cbrt(8.0) // 返回 2.0

// 幂运算
power := candy.Pow(2, 3) // 返回 8
```

#### 集合操作
```go
// 将切片分块成更小的切片
chunks := candy.Chunk([]int{1,2,3,4,5,6}, 2) // 返回 [[1,2], [3,4], [5,6]]

// 检查切片是否包含元素
exists := candy.Contains([]string{"a", "b", "c"}, "b") // 返回 true

// 对每个元素应用函数
candy.Each([]int{1,2,3}, func(v int) { fmt.Println(v) })

// 将切片归约为单个值
sum := candy.Reduce([]int{1,2,3,4}, 0, func(acc, v int) int { return acc + v }) // 返回 10
```

### 最佳实践
1. 为数学运算使用适当的数值约束（Integer/Float）
2. 在集合函数中适当处理空切片
3. 利用泛型类型确保类型安全
4. 考虑大数据集的性能影响

---

## 繁體中文

### 概述
`candy` 套件為 Go 提供語法糖工具，透過型別安全的泛型函數簡化常見程式設計操作。

### 主要特性
- **數學運算**: 支援絕對值、平方根、立方根、冪運算等數值型別操作
- **集合工具**: 提供切片的分塊、包含、遍歷、歸約等功能
- **統計函數**: 平均值、最大值、最小值計算
- **排序工具**: 自訂排序邏輯的 SortUsing
- **隨機操作**: 隨機元素選擇
- **型別安全**: 完整的泛型型別支援和約束

### 核心函數
```go
// 絕對值計算
result := candy.Abs(-42) // 回傳 42

// 平方根
sqrt := candy.Sqrt(16.0) // 回傳 4.0
```

### 最佳實務
1. 為數學運算使用適當的數值約束（Integer/Float）
2. 在集合函數中適當處理空切片
3. 利用泛型型別確保型別安全

---

## Русский

### Обзор
Пакет `candy` предоставляет утилиты синтаксического сахара для Go, упрощая общие операции программирования с типобезопасными универсальными функциями.

### Основные возможности
- **Математические операции**: функции Abs, Sqrt, Cbrt, Pow для числовых типов
- **Утилиты коллекций**: функции Chunk, Contains, Each, Reduce для срезов
- **Статистические функции**: вычисления Average, Max, Min
- **Утилиты сортировки**: SortUsing для пользовательской логики сортировки
- **Случайные операции**: выбор случайного элемента
- **Безопасность типов**: полная поддержка универсальных типов с ограничениями

### Основные функции
```go
// Вычисление абсолютного значения
result := candy.Abs(-42) // Возвращает 42

// Квадратный корень
sqrt := candy.Sqrt(16.0) // Возвращает 4.0
```

### Лучшие практики
1. Используйте соответствующие числовые ограничения (Integer/Float) для математических операций
2. Правильно обрабатывайте пустые срезы в функциях коллекций

---

## Français

### Aperçu
Le package `candy` fournit des utilitaires de sucre syntaxique pour Go, simplifiant les opérations de programmation courantes avec des fonctions génériques type-safe.

### Caractéristiques principales
- **Opérations mathématiques**: fonctions Abs, Sqrt, Cbrt, Pow pour les types numériques
- **Utilitaires de collection**: fonctions Chunk, Contains, Each, Reduce pour les tranches
- **Fonctions statistiques**: calculs Average, Max, Min
- **Utilitaires de tri**: SortUsing pour la logique de tri personnalisée
- **Opérations aléatoires**: sélection d'éléments aléatoires
- **Sécurité des types**: prise en charge complète des types génériques avec contraintes

### Fonctions principales
```go
// Calcul de la valeur absolue
result := candy.Abs(-42) // Retourne 42

// Racine carrée
sqrt := candy.Sqrt(16.0) // Retourne 4.0
```

### Meilleures pratiques
1. Utilisez des contraintes numériques appropriées (Integer/Float) pour les opérations mathématiques
2. Gérez correctement les tranches vides dans les fonctions de collection

---

## العربية

### نظرة عامة
توفر حزمة `candy` أدوات السكر النحوي لـ Go، مما يبسط عمليات البرمجة الشائعة باستخدام وظائف عامة آمنة النوع.

### الميزات الرئيسية
- **العمليات الرياضية**: وظائف Abs، Sqrt، Cbrt، Pow للأنواع الرقمية
- **أدوات المجموعة**: وظائف Chunk، Contains، Each، Reduce للشرائح
- **الوظائف الإحصائية**: حسابات Average، Max، Min
- **أدوات الفرز**: SortUsing لمنطق الفرز المخصص
- **العمليات العشوائية**: اختيار العناصر العشوائية
- **أمان النوع**: دعم كامل للأنواع العامة مع القيود

### الوظائف الأساسية
```go
// حساب القيمة المطلقة
result := candy.Abs(-42) // يعيد 42

// الجذر التربيعي
sqrt := candy.Sqrt(16.0) // يعيد 4.0
```

### أفضل الممارسات
1. استخدم القيود الرقمية المناسبة (Integer/Float) للعمليات الرياضية
2. تعامل بشكل مناسب مع الشرائح الفارغة في وظائف المجموعة

---

## Español

### Descripción general
El paquete `candy` proporciona utilidades de azúcar sintáctico para Go, simplificando operaciones de programación comunes con funciones genéricas type-safe.

### Características principales
- **Operaciones matemáticas**: funciones Abs, Sqrt, Cbrt, Pow para tipos numéricos
- **Utilidades de colección**: funciones Chunk, Contains, Each, Reduce para slices
- **Funciones estadísticas**: cálculos Average, Max, Min
- **Utilidades de ordenación**: SortUsing para lógica de ordenación personalizada
- **Operaciones aleatorias**: selección de elementos aleatorios
- **Seguridad de tipos**: soporte completo de tipos genéricos con restricciones

### Funciones principales
```go
// Cálculo de valor absoluto
result := candy.Abs(-42) // Devuelve 42

// Raíz cuadrada
sqrt := candy.Sqrt(16.0) // Devuelve 4.0
```

### Mejores prácticas
1. Use restricciones numéricas apropiadas (Integer/Float) para operaciones matemáticas
2. Maneje apropiadamente los slices vacíos en funciones de colección