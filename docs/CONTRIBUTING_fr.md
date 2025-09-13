# Contribuer à LazyGophers Utils

Merci pour votre intérêt à contribuer à LazyGophers Utils ! Ce document fournit des directives et informations pour les contributeurs.

## 🚀 Commencer

### Prérequis

- Go 1.24.0 ou plus récent
- Git
- Make (optionnel, pour l'automatisation)

### Configuration de Développement

1. **Fork et Clone**
   ```bash
   git clone https://github.com/your-username/utils.git
   cd utils
   ```

2. **Installer les Dépendances**
   ```bash
   go mod tidy
   ```

3. **Vérifier la Configuration**
   ```bash
   go test ./...
   ```

## 📋 Directives de Développement

### Style de Code

1. **Suivre les Standards Go**
   - Utiliser `gofmt` pour le formatage
   - Suivre les pratiques Go efficaces
   - Utiliser des noms de variables et fonctions significatifs

2. **Directives Spécifiques aux Packages**
   - Chaque package doit être indépendant et réutilisable
   - Minimiser les dépendances externes
   - Utiliser des génériques pour la sécurité des types quand approprié

3. **Documentation**
   - Toutes les fonctions publiques doivent avoir des commentaires de documentation en chinois
   - Inclure des exemples d'utilisation pour les fonctions complexes
   - Documenter les caractéristiques de performance pour les fonctions critiques

### Directives de Performance

1. **Optimisation Mémoire**
   - Utiliser des pools d'objets pour les opérations haute fréquence
   - Préférer les opérations zero-copy quand possible
   - Minimiser les allocations mémoire dans les chemins chauds

2. **Concurrence**
   - Utiliser des opérations atomiques au lieu de mutex quand possible
   - Assurer la sécurité thread pour les opérations concurrentes
   - Concevoir des algorithmes lock-free quand approprié

## 🧪 Exigences de Test

### Tests Unitaires

1. **Couverture de Test**
   - Viser 90%+ de couverture de test pour le nouveau code
   - Tester les chemins de succès et d'erreur
   - Inclure les cas limites et conditions aux frontières

2. **Organisation des Tests**
   ```bash
   # Exécuter des tests pour un package spécifique
   go test ./candy
   
   # Exécuter avec couverture
   go test -cover ./...
   ```

## 📝 Directives de Commit

### Format de Message de Commit

```
<type>(<scope>): <description>

<body>

<footer>
```

**Types :**
- `feat` : Nouvelle fonctionnalité
- `fix` : Correction de bug
- `perf` : Amélioration de performance
- `refactor` : Refactorisation de code
- `test` : Ajout ou mise à jour de tests
- `docs` : Modifications de documentation

## 🔍 Processus de Révision de Code

### Directives Pull Request

1. **Avant Soumission**
   - S'assurer que tous les tests passent
   - Exécuter `go fmt ./...`
   - Exécuter `go vet ./...`
   - Mettre à jour la documentation si nécessaire

2. **Critères de Révision**
   - Qualité et lisibilité du code
   - Couverture et qualité des tests
   - Impact sur les performances
   - Changements cassants
   - Complétude de la documentation

## 🏗️ Directives d'Architecture

### Conception de Packages

1. **Responsabilité Unique**
   - Chaque package doit avoir un objectif clair et ciblé
   - Éviter le mélange de fonctionnalités non liées
   - Maintenir des APIs publiques minimales et propres

2. **Dépendances**
   - Minimiser les dépendances externes
   - Préférer la bibliothèque standard quand possible
   - Documenter la justification des dépendances

## 📚 Documentation

### Documentation API

```go
// ToString convertit tout type en string
// Supporte la conversion des types de base, slices, maps et structs
// Utilise la sérialisation JSON pour les types complexes
//
// Caractéristiques de performance :
// - Conversion types de base : O(1)
// - Conversion types complexes : O(n) où n est la complexité de sérialisation
func ToString(v interface{}) string
```

## 🐛 Directives des Issues

### Rapports de Bugs

```markdown
**Description du Bug**
Description claire du bug

**Étapes de Reproduction**
1. Première étape
2. Deuxième étape
3. Troisième étape

**Comportement Attendu**
Ce qui devrait se passer

**Comportement Actuel**
Ce qui s'est réellement passé

**Environnement**
- Version Go :
- OS :
- Version du package :
```

## 🤝 Directives Communautaires

### Code de Conduite

1. **Soyez Respectueux**
   - Traitez tous les contributeurs avec respect
   - Soyez constructif dans les retours
   - Accueillez les nouveaux venus

2. **Soyez Collaboratif**
   - Partagez les connaissances et aidez les autres
   - Fournissez des révisions claires et utiles
   - Communiquez ouvertement sur les défis

## 📖 Ressources d'Apprentissage

### Meilleures Pratiques Go
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Commentaires de Révision de Code Go](https://github.com/golang/go/wiki/CodeReviewComments)

### Optimisation de Performance
- [Conseils de Performance Go](https://github.com/golang/go/wiki/Performance)
- [Profiling des Programmes Go](https://blog.golang.org/profiling-go-programs)

### Tests
- [Tests en Go](https://golang.org/doc/code.html#Testing)

## 📞 Obtenir de l'Aide

Si vous avez besoin d'aide ou avez des questions :

1. Vérifiez la documentation existante
2. Recherchez les issues existantes
3. Créez une nouvelle issue avec une description claire
4. Rejoignez nos discussions communautaires

Merci de contribuer à LazyGophers Utils ! Vos contributions aident à rendre cette bibliothèque meilleure pour tous.