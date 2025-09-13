# LazyGophers Utils

> 🚀 Une bibliothèque d'utilitaires Go riche en fonctionnalités et haute performance qui rend le développement Go plus efficace

**🌍 Langues**: [English](README.md) • [中文](README_zh.md) • [繁體中文](README_zh-hant.md) • [Español](README_es.md) • [Français](README_fr.md) • [Русский](README_ru.md) • [العربية](README_ar.md)

[![Go Version](https://img.shields.io/badge/Go-1.18+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-AGPL%20v3-green.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils.svg)](https://pkg.go.dev/github.com/lazygophers/utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/utils)](https://goreportcard.com/report/github.com/lazygophers/utils)
[![GitHub releases](https://img.shields.io/github/release/lazygophers/utils.svg)](https://github.com/lazygophers/utils/releases)
[![GoProxy.cn Downloads](https://goproxy.cn/stats/github.com/lazygophers/utils/badges/download-count.svg)](https://goproxy.cn/stats/github.com/lazygophers/utils)
[![Test Coverage](https://img.shields.io/badge/coverage-85%25-brightgreen.svg)](https://github.com/lazygophers/utils/actions)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/lazygophers/utils)

## 📋 Table des Matières

- [Aperçu du Projet](#-aperçu-du-projet)
- [Caractéristiques Principales](#-caractéristiques-principales)
- [Démarrage Rapide](#-démarrage-rapide)
- [Documentation](#-documentation)
- [Modules Principaux](#-modules-principaux)
- [Modules de Fonctionnalités](#-modules-de-fonctionnalités)
- [Exemples d'Utilisation](#-exemples-dutilisation)
- [Données de Performance](#-données-de-performance)
- [Contribuer](#-contribuer)
- [Licence](#-licence)
- [Support Communautaire](#-support-communautaire)

## 💡 Aperçu du Projet

LazyGophers Utils est une bibliothèque d'utilitaires Go complète et haute performance qui fournit plus de 20 modules professionnels couvrant divers besoins du développement quotidien. Elle adopte une conception modulaire pour des importations à la demande avec zéro conflit de dépendances.

**Philosophie de Conception**: Simple, Efficace, Fiable

## ✨ Caractéristiques Principales

| Caractéristique | Description | Avantage |
|------------------|-------------|----------|
| 🧩 **Conception Modulaire** | Plus de 20 modules indépendants | Importer à la demande, réduire la taille |
| ⚡ **Haute Performance** | Testé avec des benchmarks | Réponse en microsecondes, économe en mémoire |
| 🛡️ **Sécurité de Type** | Utilisation complète des génériques | Vérification d'erreurs à la compilation |
| 🔒 **Sécurité de Concurrence** | Conception amicale aux goroutines | Prêt pour la production |
| 📚 **Bien Documenté** | Couverture de documentation 95%+ | Facile à apprendre et utiliser |
| 🧪 **Bien Testé** | Couverture de tests 85%+ | Assurance qualité |

## 🚀 Démarrage Rapide

### Installation

```bash
go get github.com/lazygophers/utils
```

### Utilisation de Base

```go
package main

import (
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // Gestion d'erreurs
    value := utils.Must(getValue())
    
    // Conversion de types
    age := candy.ToInt("25")
    
    // Traitement du temps
    cal := xtime.NowCalendar()
    fmt.Println(cal.String()) // 2023年08月15日 六月廿九 兔年 处暑
}
```

## 📖 Documentation

### 📁 Documentation des Modules
- **Modules Principaux**: [Gestion d'Erreurs](must.go) | [Base de Données](orm.go) | [Validation](validate.go)
- **Traitement de Données**: [candy](candy/) | [json](json/) | [stringx](stringx/)
- **Outils de Temps**: [xtime](xtime/) | [xtime996](xtime/xtime996/) | [xtime955](xtime/xtime955/)
- **Outils Système**: [config](config/) | [runtime](runtime/) | [osx](osx/)
- **Réseau et Sécurité**: [network](network/) | [cryptox](cryptox/) | [pgp](pgp/)
- **Concurrence et Contrôle**: [routine](routine/) | [wait](wait/) | [hystrix](hystrix/)

Pour une documentation complète, consultez le [centre de documentation](docs/).

## 🎯 Exemples d'Utilisation

### Exemple d'Application Complète

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
    // 1. Charger la configuration
    var cfg AppConfig
    utils.MustSuccess(config.Load(&cfg, "config.json"))
    
    // 2. Valider la configuration
    utils.MustSuccess(utils.Validate(&cfg))
    
    // 3. Conversion de types
    portStr := candy.ToString(cfg.Port)
    
    // 4. Traitement du temps
    cal := xtime.NowCalendar()
    log.Printf("Application démarrée: %s", cal.String())
    
    // 5. Démarrer le serveur
    startServer(cfg)
}
```

## 📊 Données de Performance

| Opération | Temps | Allocation Mémoire | vs Bibliothèque Standard |
|-----------|-------|--------------------|--------------------------|
| `candy.ToInt()` | 12.3 ns/op | 0 B/op | **3.2x plus rapide** |
| `json.Marshal()` | 156 ns/op | 64 B/op | **1.8x plus rapide** |
| `xtime.Now()` | 45.2 ns/op | 0 B/op | **2.1x plus rapide** |
| `utils.Must()` | 2.1 ns/op | 0 B/op | **Zéro surcharge** |

## 🤝 Contribuer

Nous accueillons les contributions de toutes sortes !

1. 🍴 Fork le projet
2. 🌿 Créer une branche de fonctionnalité
3. 📝 Écrire du code et des tests
4. 🧪 S'assurer que les tests passent
5. 📤 Soumettre une PR

## 📄 Licence

Ce projet est sous licence GNU Affero General Public License v3.0.

Voir le fichier [LICENSE](LICENSE) pour plus de détails.

## 🌟 Support Communautaire

### Obtenir de l'Aide

- 📖 **Documentation**: [Documentation Complète](docs/)
- 🐛 **Rapports de Bugs**: [GitHub Issues](https://github.com/lazygophers/utils/issues)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/lazygophers/utils/discussions)
- ❓ **Q&R**: [Stack Overflow](https://stackoverflow.com/questions/tagged/lazygophers-utils)

---

<div align="center">

**Si ce projet vous aide, merci de nous donner une ⭐ Étoile !**

[🚀 Commencer](#-démarrage-rapide) • [📖 Voir la Documentation](docs/) • [🤝 Rejoindre la Communauté](https://github.com/lazygophers/utils/discussions)

</div>