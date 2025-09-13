# lazygophers/utils مرجع API

مكتبة أدوات Go شاملة بتصميم معياري، توفر وظائف أساسية لمهام التطوير الشائعة. تركز هذه المكتبة على أمان الأنواع وتحسين الأداء وتتبع معايير Go 1.24+.

## التثبيت

```go
go get github.com/lazygophers/utils
```

## جدول المحتويات

- [الحزمة الأساسية (utils)](#الحزمة-الأساسية-utils)
- [تحويل الأنواع (candy)](#تحويل-الأنواع-candy)
- [معالجة النصوص (stringx)](#معالجة-النصوص-stringx)
- [عمليات الخرائط (anyx)](#عمليات-الخرائط-anyx)
- [التحكم في التزامن (wait)](#التحكم-في-التزامن-wait)
- [قاطع الدائرة (hystrix)](#قاطع-الدائرة-hystrix)
- [إدارة التكوين (config)](#إدارة-التكوين-config)
- [العمليات التشفيرية (cryptox)](#العمليات-التشفيرية-cryptox)

---

## الحزمة الأساسية (utils)

توفر الحزمة الجذر أدوات أساسية تشمل معالجة الأخطاء وعمليات قاعدة البيانات والتحقق.

### دوال Must

أدوات معالجة الأخطاء التي تسبب panic عند الخطأ - مفيدة للتهيئة والعمليات الحرجة.

```go
// Must يجمع بين فحص الأخطاء وإرجاع القيم
func Must[T any](value T, err error) T

// MustOk يسبب panic إذا كان ok خطأ
func MustOk[T any](value T, ok bool) T

// MustSuccess يسبب panic إذا لم يكن الخطأ nil
func MustSuccess(err error)
```

**مثال الاستخدام:**
```go
import "github.com/lazygophers/utils"

// معالجة الأخطاء أثناء التهيئة
config := utils.Must(loadConfig())

// عمليات التحويل
value := utils.MustOk(m["key"])
```

### تكامل قاعدة البيانات

```go
// Scan مسح حقول قاعدة البيانات إلى هيكل، يدعم إلغاء تسلسل JSON
func Scan(src interface{}, dst interface{}) error

// Value تحويل الهيكل إلى قيمة قاعدة بيانات، يدعم تسلسل JSON
func Value(m interface{}) (driver.Value, error)
```

---

## تحويل الأنواع (candy)

أدوات شاملة لتحويل الأنواع ومعالجة الشرائح، بتغطية اختبار 99.3%.

### تحويلات الأنواع الأساسية

```go
// تحويل منطقي
func ToBool(v interface{}) bool

// تحويل نصي
func ToString(v interface{}) string

// تحويل أعداد صحيحة
func ToInt(v interface{}) int
func ToInt32(v interface{}) int32
func ToInt64(v interface{}) int64

// تحويل أعداد عشرية
func ToFloat32(v interface{}) float32
func ToFloat64(v interface{}) float64
```

**مثال الاستخدام:**
```go
import "github.com/lazygophers/utils/candy"

// تحويل أنواع ذكي
str := candy.ToString(123)        // "123"
num := candy.ToInt("456")         // 456
flag := candy.ToBool("true")      // true
```

### عمليات الشرائح

```go
// Filter تصفية عناصر الشريحة
func Filter[T any](slice []T, predicate func(T) bool) []T

// Map تحويل عناصر الشريحة
func Map[T, R any](slice []T, mapper func(T) R) []R

// Unique إزالة العناصر المكررة
func Unique[T comparable](slice []T) []T
```

---

## معالجة النصوص (stringx)

معالجة نصوص عالية الأداء، بتغطية اختبار 96.4%.

### تحويلات بدون نسخ

```go
// ToString تحويل نصي فعال (بدون نسخ)
func ToString(b []byte) string

// ToBytes تحويل بايت فعال (بدون نسخ)  
func ToBytes(s string) []byte
```

### تحويل أسماء

```go
// Camel2Snake تحويل من الجمل إلى الثعبان
func Camel2Snake(s string) string

// Snake2Camel تحويل من الثعبان إلى الجمل
func Snake2Camel(s string) string
```

**مثال الاستخدام:**
```go
import "github.com/lazygophers/utils/stringx"

// تحويل أسماء
snake := stringx.Camel2Snake("UserName")     // "user_name"
camel := stringx.Snake2Camel("user_name")    // "UserName"

// تحويلات بدون نسخ (أداء عالي)
bytes := stringx.ToBytes("hello")
str := stringx.ToString([]byte("world"))
```

---

## عمليات الخرائط (anyx)

عمليات خرائط مستقلة عن النوع واستخراج القيم، بتغطية اختبار 99.0%.

### هيكل MapAny

```go
// NewMap إنشاء MapAny جديد
func NewMap(m map[string]interface{}) *MapAny

// Get الحصول على قيمة
func (m *MapAny) Get(key string) interface{}

// Set تعيين قيمة
func (m *MapAny) Set(key string, value interface{})
```

### استخراج آمن للأنواع

```go
// الحصول على قيم بأنواع محددة
func (m *MapAny) GetString(key string) string
func (m *MapAny) GetInt(key string) int
func (m *MapAny) GetBool(key string) bool
```

---

## التحكم في التزامن (wait)

أدوات تزامن وتنسيق متقدمة.

### المعالجة غير المتزامنة

```go
// Async معالجة عناصر الشريحة بشكل غير متزامن
func Async[T any](items []T, workerCount int, processor func(T)) error
```

**مثال الاستخدام:**
```go
import "github.com/lazygophers/utils/wait"

urls := []string{
    "https://api1.com",
    "https://api2.com",
}

// معالجة متزامنة للعناوين
err := wait.Async(urls, 2, func(url string) {
    resp, err := http.Get(url)
    // معالجة الاستجابة...
})
```

---

## قاطع الدائرة (hystrix)

تطبيق عالي الأداء لنمط قاطع الدائرة.

```go
type CircuitBreaker struct {
    // تطبيق قاطع دائرة عالي الأداء
}

func NewCircuitBreaker(config Config) *CircuitBreaker
func (cb *CircuitBreaker) Call(fn func() error) error
```

---

## إدارة التكوين (config)

تحميل تكوين متعدد التنسيقات.

```go
// LoadConfig تحميل ملف التكوين
func LoadConfig(filename string, config interface{}) error
```

يدعم التنسيقات:
- JSON (.json)
- YAML (.yaml, .yml)  
- TOML (.toml)
- INI (.ini)

---

## العمليات التشفيرية (cryptox)

عمليات تشفير شاملة بتغطية اختبار 100%.

### تشفير AES

```go
// تشفير/فك تشفير AES
func Encrypt(plaintext, key []byte) ([]byte, error)
func Decrypt(ciphertext, key []byte) ([]byte, error)
```

### عمليات RSA

```go
// توليد مفاتيح RSA
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error)
```

### دوال التجزئة

```go
// تجزئة شائعة
func MD5(data []byte) string
func SHA256(data []byte) string
func SHA512(data []byte) string
```

---

## خصائص الأداء

- **عمليات بدون تخصيص**: العديد من الوظائف تحقق صفر تخصيصات ذاكرة
- **عمليات ذرية**: تطبيقات خالية من الأقفال لسيناريوهات عالية التزامن  
- **هياكل محاذية للذاكرة**: محسنة لكفاءة ذاكرة التخزين المؤقت للمعالج
- **تحسينات المتغيرات العامة**: عمليات آمنة الأنواع بدون انعكاس وقت التشغيل

## نمط معالجة الأخطاء

جميع الحزم تتبع نمط معالجة أخطاء متسق:
1. استخدام `github.com/lazygophers/log` لتسجيل الأخطاء
2. إرجاع رسائل خطأ ذات معنى
3. توفير دوال `Must*` للعمليات الحرجة

هذا المرجع يوفر دليلاً شاملاً لاستخدام مكتبة LazyGophers Utils، مما يساعد المطورين على بناء تطبيقات Go عالية الجودة بكفاءة.