# LazyGophers Utils

> 🚀 مكتبة أدوات Go غنية بالميزات وعالية الأداء تجعل التطوير بـ Go أكثر كفاءة

**🌍 اللغات**: [English](README.md) • [中文](README_zh.md) • [繁體中文](README_zh-hant.md) • [Español](README_es.md) • [Français](README_fr.md) • [Русский](README_ru.md) • [العربية](README_ar.md)

[![Go Version](https://img.shields.io/badge/Go-1.18+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-AGPL%20v3-green.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils.svg)](https://pkg.go.dev/github.com/lazygophers/utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/utils)](https://goreportcard.com/report/github.com/lazygophers/utils)
[![GitHub releases](https://img.shields.io/github/release/lazygophers/utils.svg)](https://github.com/lazygophers/utils/releases)
[![GoProxy.cn Downloads](https://goproxy.cn/stats/github.com/lazygophers/utils/badges/download-count.svg)](https://goproxy.cn/stats/github.com/lazygophers/utils)
[![Test Coverage](https://img.shields.io/badge/coverage-85%25-brightgreen.svg)](https://github.com/lazygophers/utils/actions)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/lazygophers/utils)

## 📋 جدول المحتويات

- [نظرة عامة على المشروع](#-نظرة-عامة-على-المشروع)
- [الميزات الأساسية](#-الميزات-الأساسية)
- [البدء السريع](#-البدء-السريع)
- [التوثيق](#-التوثيق)
- [الوحدات الأساسية](#-الوحدات-الأساسية)
- [وحدات الميزات](#-وحدات-الميزات)
- [أمثلة الاستخدام](#-أمثلة-الاستخدام)
- [بيانات الأداء](#-بيانات-الأداء)
- [المساهمة](#-المساهمة)
- [الرخصة](#-الرخصة)
- [دعم المجتمع](#-دعم-المجتمع)

## 💡 نظرة عامة على المشروع

LazyGophers Utils هي مكتبة أدوات Go شاملة وعالية الأداء توفر أكثر من 20 وحدة احترافية تغطي احتياجات متنوعة في التطوير اليومي. تتبنى تصميمًا معياريًا للاستيراد عند الطلب مع عدم وجود تعارضات في التبعيات.

**فلسفة التصميم**: بسيط، كفؤ، موثوق

## ✨ الميزات الأساسية

| الميزة | الوصف | الميزة |
|-------|--------|--------|
| 🧩 **التصميم المعياري** | أكثر من 20 وحدة مستقلة | استيراد عند الطلب، تقليل الحجم |
| ⚡ **أداء عالي** | مختبر بالمقاييس المرجعية | استجابة بالميكروثانية، صديق للذاكرة |
| 🛡️ **آمن من ناحية النوع** | استخدام كامل للأنواع العامة | فحص الأخطاء في وقت التجميع |
| 🔒 **آمن للتزامن** | تصميم صديق للـ goroutines | جاهز للإنتاج |
| 📚 **موثق جيداً** | تغطية التوثيق 95%+ | سهل التعلم والاستخدام |
| 🧪 **مختبر جيداً** | تغطية الاختبار 85%+ | ضمان الجودة |

## 🚀 البدء السريع

### التثبيت

```bash
go get github.com/lazygophers/utils
```

### الاستخدام الأساسي

```go
package main

import (
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // معالجة الأخطاء
    value := utils.Must(getValue())
    
    // تحويل الأنواع
    age := candy.ToInt("25")
    
    // معالجة الوقت
    cal := xtime.NowCalendar()
    fmt.Println(cal.String()) // 2023年08月15日 六月廿九 兔年 处暑
}
```

## 📖 التوثيق

### 📁 توثيق الوحدات
- **الوحدات الأساسية**: [معالجة الأخطاء](must.go) | [قاعدة البيانات](orm.go) | [التحقق](validate.go)
- **معالجة البيانات**: [candy](candy/) | [json](json/) | [stringx](stringx/)
- **أدوات الوقت**: [xtime](xtime/) | [xtime996](xtime/xtime996/) | [xtime955](xtime/xtime955/)
- **أدوات النظام**: [config](config/) | [runtime](runtime/) | [osx](osx/)
- **الشبكة والأمان**: [network](network/) | [cryptox](cryptox/) | [pgp](pgp/)
- **التزامن والتحكم**: [routine](routine/) | [wait](wait/) | [hystrix](hystrix/)

للحصول على التوثيق الكامل، راجع [مركز التوثيق](docs/).

## 🎯 أمثلة الاستخدام

### مثال تطبيق كامل

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
    // 1. تحميل التكوين
    var cfg AppConfig
    utils.MustSuccess(config.Load(&cfg, "config.json"))
    
    // 2. التحقق من التكوين
    utils.MustSuccess(utils.Validate(&cfg))
    
    // 3. تحويل الأنواع
    portStr := candy.ToString(cfg.Port)
    
    // 4. معالجة الوقت
    cal := xtime.NowCalendar()
    log.Printf("تم بدء التطبيق: %s", cal.String())
    
    // 5. بدء الخادم
    startServer(cfg)
}
```

## 📊 بيانات الأداء

| العملية | الوقت | تخصيص الذاكرة | مقابل المكتبة القياسية |
|---------|-------|---------------|-----------------------|
| `candy.ToInt()` | 12.3 ns/op | 0 B/op | **أسرع بـ 3.2 مرة** |
| `json.Marshal()` | 156 ns/op | 64 B/op | **أسرع بـ 1.8 مرة** |
| `xtime.Now()` | 45.2 ns/op | 0 B/op | **أسرع بـ 2.1 مرة** |
| `utils.Must()` | 2.1 ns/op | 0 B/op | **عدم وجود نفقات إضافية** |

## 🤝 المساهمة

نرحب بالمساهمات من جميع الأنواع!

1. 🍴 قم بعمل fork للمشروع
2. 🌿 أنشئ فرع ميزة
3. 📝 اكتب الكود والاختبارات
4. 🧪 تأكد من نجاح الاختبارات
5. 📤 أرسل PR

## 📄 الرخصة

هذا المشروع مرخص تحت رخصة GNU Affero General Public License v3.0.

راجع ملف [LICENSE](LICENSE) للحصول على التفاصيل.

## 🌟 دعم المجتمع

### الحصول على المساعدة

- 📖 **التوثيق**: [التوثيق الكامل](docs/)
- 🐛 **تقارير الأخطاء**: [GitHub Issues](https://github.com/lazygophers/utils/issues)
- 💬 **المناقشات**: [GitHub Discussions](https://github.com/lazygophers/utils/discussions)
- ❓ **الأسئلة والأجوبة**: [Stack Overflow](https://stackoverflow.com/questions/tagged/lazygophers-utils)

---

<div align="center">

**إذا كان هذا المشروع يساعدك، يرجى إعطاؤنا ⭐ نجمة!**

[🚀 البدء](#-البدء-السريع) • [📖 عرض التوثيق](docs/) • [🤝 انضم للمجتمع](https://github.com/lazygophers/utils/discussions)

</div>