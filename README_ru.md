# LazyGophers Utils

> 🚀 Многофункциональная высокопроизводительная библиотека утилит Go, делающая разработку на Go более эффективной

**🌍 Языки**: [English](README.md) • [中文](README_zh.md) • [繁體中文](README_zh-hant.md) • [Español](README_es.md) • [Français](README_fr.md) • [Русский](README_ru.md) • [العربية](README_ar.md)

[![Go Version](https://img.shields.io/badge/Go-1.24.0-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-AGPL%20v3-green.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils.svg)](https://pkg.go.dev/github.com/lazygophers/utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/utils)](https://goreportcard.com/report/github.com/lazygophers/utils)
[![GitHub releases](https://img.shields.io/github/release/lazygophers/utils.svg)](https://github.com/lazygophers/utils/releases)
[![GoProxy.cn Downloads](https://goproxy.cn/stats/github.com/lazygophers/utils/badges/download-count.svg)](https://goproxy.cn/stats/github.com/lazygophers/utils)
[![Test Coverage](https://img.shields.io/badge/coverage-85%25-brightgreen.svg)](https://github.com/lazygophers/utils/actions)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/lazygophers/utils)

## 📋 Содержание

- [Обзор Проекта](#-обзор-проекта)
- [Основные Возможности](#-основные-возможности)
- [Быстрый Старт](#-быстрый-старт)
- [Документация](#-документация)
- [Основные Модули](#-основные-модули)
- [Функциональные Модули](#-функциональные-модули)
- [Примеры Использования](#-примеры-использования)
- [Данные Производительности](#-данные-производительности)
- [Вклад в Проект](#-вклад-в-проект)
- [Лицензия](#-лицензия)
- [Поддержка Сообщества](#-поддержка-сообщества)

## 💡 Обзор Проекта

LazyGophers Utils — это комплексная высокопроизводительная библиотека утилит Go, которая предоставляет более 20 профессиональных модулей, покрывающих различные потребности ежедневной разработки. Она использует модульный дизайн для импорта по требованию с нулевыми конфликтами зависимостей.

**Философия Дизайна**: Простой, Эффективный, Надежный

## ✨ Основные Возможности

| Возможность | Описание | Преимущество |
|-------------|----------|--------------|
| 🧩 **Модульный Дизайн** | Более 20 независимых модулей | Импорт по требованию, уменьшение размера |
| ⚡ **Высокая Производительность** | Протестировано с бенчмарками | Ответ в микросекундах, дружелюбно к памяти |
| 🛡️ **Типобезопасность** | Полное использование дженериков | Проверка ошибок во время компиляции |
| 🔒 **Безопасность Конкурентности** | Дизайн, дружелюбный к горутинам | Готов для продакшена |
| 📚 **Хорошо Документирован** | Покрытие документацией 95%+ | Легко изучить и использовать |
| 🧪 **Хорошо Протестирован** | Покрытие тестами 85%+ | Гарантия качества |

## 🚀 Быстрый Старт

### Установка

```bash
go get github.com/lazygophers/utils
```

### Базовое Использование

```go
package main

import (
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // Обработка ошибок
    value := utils.Must(getValue())
    
    // Преобразование типов
    age := candy.ToInt("25")
    
    // Обработка времени
    cal := xtime.NowCalendar()
    fmt.Println(cal.String()) // 2023年08月15日 六月廿九 兔年 处暑
}
```

## 📖 Документация

### 📁 Документация Модулей
- **Основные Модули**: [Обработка Ошибок](must.go) | [База Данных](orm.go) | [Валидация](validate.go)
- **Обработка Данных**: [candy](candy/) | [json](json/) | [stringx](stringx/)
- **Инструменты Времени**: [xtime](xtime/) | [xtime996](xtime/xtime996/) | [xtime955](xtime/xtime955/)
- **Системные Инструменты**: [config](config/) | [runtime](runtime/) | [osx](osx/)
- **Сеть и Безопасность**: [network](network/) | [cryptox](cryptox/) | [pgp](pgp/)
- **Конкурентность и Управление**: [routine](routine/) | [wait](wait/) | [hystrix](hystrix/)

Для полной документации обратитесь к [центру документации](docs/).

## 🎯 Примеры Использования

### Пример Полного Приложения

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
    // 1. Загрузить конфигурацию
    var cfg AppConfig
    utils.MustSuccess(config.Load(&cfg, "config.json"))
    
    // 2. Валидировать конфигурацию
    utils.MustSuccess(utils.Validate(&cfg))
    
    // 3. Преобразование типов
    portStr := candy.ToString(cfg.Port)
    
    // 4. Обработка времени
    cal := xtime.NowCalendar()
    log.Printf("Приложение запущено: %s", cal.String())
    
    // 5. Запустить сервер
    startServer(cfg)
}
```

## 📊 Данные Производительности

| Операция | Время | Выделение Памяти | vs Стандартная Библиотека |
|----------|-------|-------------------|---------------------------|
| `candy.ToInt()` | 12.3 ns/op | 0 B/op | **В 3.2 раза быстрее** |
| `json.Marshal()` | 156 ns/op | 64 B/op | **В 1.8 раза быстрее** |
| `xtime.Now()` | 45.2 ns/op | 0 B/op | **В 2.1 раза быстрее** |
| `utils.Must()` | 2.1 ns/op | 0 B/op | **Нулевые накладные расходы** |

## 🤝 Вклад в Проект

Мы приветствуем вклад любого рода!

1. 🍴 Сделайте форк проекта
2. 🌿 Создайте ветку функций
3. 📝 Напишите код и тесты
4. 🧪 Убедитесь, что тесты проходят
5. 📤 Отправьте PR

## 📄 Лицензия

Этот проект лицензирован под GNU Affero General Public License v3.0.

См. файл [LICENSE](LICENSE) для подробностей.

## 🌟 Поддержка Сообщества

### Получить Помощь

- 📖 **Документация**: [Полная Документация](docs/)
- 🐛 **Отчеты о Багах**: [GitHub Issues](https://github.com/lazygophers/utils/issues)
- 💬 **Обсуждения**: [GitHub Discussions](https://github.com/lazygophers/utils/discussions)
- ❓ **Вопросы и Ответы**: [Stack Overflow](https://stackoverflow.com/questions/tagged/lazygophers-utils)

---

<div align="center">

**Если этот проект помогает вам, пожалуйста, поставьте нам ⭐ Звезду!**

[🚀 Начать](#-быстрый-старт) • [📖 Просмотреть Документацию](docs/) • [🤝 Присоединиться к Сообществу](https://github.com/lazygophers/utils/discussions)

</div>