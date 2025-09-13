# Guide de Contribution

Bienvenue pour contribuer au projet LazyGophers Utils ! Nous apprécions grandement chaque contribution de la communauté.

[![Contributors](https://img.shields.io/badge/Contributors-Welcome-brightgreen.svg)](#comment-contribuer)
[![Code Style](https://img.shields.io/badge/Code%20Style-Go%20Standard-blue.svg)](#normes-de-code)

## 🤝 Comment Contribuer

### Types de Contributions

Nous accueillons les types de contributions suivants :

- 🐛 **Corrections de Bugs** - Corriger les problèmes connus
- ✨ **Nouvelles Fonctionnalités** - Ajouter de nouvelles fonctions utilitaires ou modules
- 📚 **Améliorations de Documentation** - Améliorer la documentation, ajouter des exemples
- 🎨 **Optimisation de Code** - Optimisation des performances, refactorisation
- 🧪 **Améliorations de Tests** - Augmenter la couverture de tests, corriger les problèmes de tests
- 🌐 **Internationalisation** - Ajouter le support multi-langues

### Processus de Contribution

#### 1. Préparation

**Fork du Projet**
```bash
# 1. Fork ce projet vers votre compte GitHub
# 2. Cloner votre fork localement
git clone https://github.com/VOTRE_NOM_UTILISATEUR/utils.git
cd utils

# 3. Ajouter le projet original comme dépôt en amont
git remote add upstream https://github.com/lazygophers/utils.git

# 4. Créer une nouvelle branche de fonctionnalité
git checkout -b feature/votre-fonctionnalite-geniale
```

**Configuration de l'Environnement de Développement**
```bash
# Installer les dépendances
go mod tidy

# Vérifier l'environnement
go version  # Nécessite Go 1.24.0+
go test ./... # S'assurer que tous les tests passent
```

#### 2. Phase de Développement

**Écrire du Code**
- Suivre les [Normes de Code](#normes-de-code)
- Écrire des cas de test pour les nouvelles fonctionnalités
- S'assurer que la couverture de tests ne diminue pas par rapport au niveau actuel
- Ajouter les commentaires de documentation nécessaires

**Normes de Commit**
```bash
# Utiliser le format de message de commit standardisé
git commit -m "feat(module): ajouter nouvelle fonction utilitaire

- Ajouter la fonction FormatDuration
- Supporter plusieurs formats de sortie de temps
- Ajouter des cas de test complets
- Mettre à jour la documentation associée

Closes #123"
```

**Format de Message de Commit** :
```
<type>(<portée>): <sujet>

<corps>

<pied de page>
```

**Catégories de Type** :
- `feat`: Nouvelles fonctionnalités
- `fix`: Corrections de bugs  
- `docs`: Mises à jour de documentation
- `style`: Ajustements de formatage de code
- `refactor`: Refactorisation de code
- `perf`: Optimisation des performances
- `test`: Relatif aux tests
- `chore`: Mises à jour d'outils de build ou de dépendances

**Portée** (optionnel) :
- `candy`: module candy
- `xtime`: module xtime
- `config`: module config
- `cryptox`: module cryptox
- etc...

#### 3. Tests et Validation

**Exécuter les Tests**
```bash
# Exécuter tous les tests
go test -v ./...

# Vérifier la couverture de tests
go test -cover -v ./...

# Exécuter les tests de benchmark
go test -bench=. ./...

# Vérifier le formatage du code
go fmt ./...

# Analyse statique
go vet ./...
```

**Tests de Performance**
```bash
# Exécuter les tests de performance
go test -bench=BenchmarkVotreFonction -benchmem ./...

# S'assurer qu'il n'y a pas de régression de performance significative
```

#### 4. Créer une Pull Request

**Pousser vers Votre Fork**
```bash
git push origin feature/votre-fonctionnalite-geniale
```

**Créer une PR**
1. Visiter la page du projet sur GitHub
2. Cliquer sur "New Pull Request"
3. Sélectionner votre branche
4. Remplir la description de la PR (se référer au [Modèle de PR](#modele-de-pr))
5. S'assurer que toutes les vérifications passent

#### 5. Révision de Code

- Les mainteneurs réviseront votre code
- Effectuer des modifications basées sur les retours
- Maintenir une communication et une attitude coopérative
- Sera fusionné après que les tests passent

## 📝 Normes de Code

### Style de Code Go

**Normes de Base**
```go
// ✅ Bon exemple
package candy

import (
    "context"
    "fmt"
    "time"
    
    "github.com/lazygophers/log"
)

// FormatDuration formate la durée en chaîne lisible par l'homme
// Supporte plusieurs niveaux de précision, choisit automatiquement les unités appropriées
//
// Paramètres :
//   - duration: durée à formater
//   - precision: niveau de précision (1-3)
//
// Retourne :
//   - string: chaîne formatée, comme "2 heures 30 minutes"
//
// Exemple :
//   FormatDuration(90*time.Minute, 2) // retourne "1 heure 30 minutes"
//   FormatDuration(45*time.Second, 1) // retourne "45 secondes"
func FormatDuration(duration time.Duration, precision int) string {
    if duration == 0 {
        return "0 secondes"
    }
    
    // Logique d'implémentation...
    return result
}
```

**Conventions de Nommage**
- Utiliser CamelCase
- Les noms de fonction commencent par des verbes : `Get`, `Set`, `Format`, `Parse`
- Les constantes utilisent ALL_CAPS : `const MaxRetries = 3`
- Les membres privés utilisent des minuscules : `internalHelper`
- Les noms de package utilisent des mots uniques en minuscules : `candy`, `xtime`

**Normes de Commentaires**
- Toutes les fonctions publiques doivent avoir des commentaires
- Les commentaires commencent par le nom de la fonction
- Inclure les descriptions des paramètres et valeurs de retour  
- Fournir des exemples d'utilisation
- Commentaires en anglais, concis et clairs

**Gestion d'Erreurs**
```go
// ✅ Approche recommandée de gestion d'erreurs
func ProcessData(data []byte) (*Result, error) {
    if len(data) == 0 {
        log.Warn("Données vides fournies")
        return nil, fmt.Errorf("les données ne peuvent pas être vides")
    }
    
    result, err := parseData(data)
    if err != nil {
        log.Error("Échec de l'analyse des données", log.Error(err))
        return nil, fmt.Errorf("échec de l'analyse des données : %w", err)
    }
    
    return result, nil
}
```

### Normes de Structure de Projet

**Organisation des Modules**
```
utils/
├── README.md           # Aperçu du projet
├── CONTRIBUTING.md     # Guide de contribution  
├── SECURITY.md        # Politique de sécurité
├── go.mod             # Définition du module Go
├── must.go            # Fonctions utilitaires principales
├── candy/             # Outils de traitement de données
│   ├── README.md      # Documentation du module
│   ├── to_string.go   # Conversion de type
│   └── to_string_test.go
├── xtime/             # Outils de traitement du temps  
│   ├── README.md      # Documentation d'utilisation détaillée
│   ├── TESTING.md     # Rapports de test
│   ├── PERFORMANCE.md # Rapports de performance
│   ├── calendar.go    # Fonctionnalité de calendrier
│   └── calendar_test.go
└── ...
```

**Nommage des Fichiers**
- Utiliser des lettres minuscules et des underscores : `to_string.go`
- Suffixe de fichier de test : `_test.go`
- Tests de benchmark : `_benchmark_test.go`
- Fichiers de documentation : `README.md`, `TESTING.md`

### Normes de Test

**Exigences de Couverture de Test**
- La couverture de test des nouvelles fonctionnalités doit être ≥ 90%
- Ne peut pas réduire la couverture de test globale
- Inclure les cas normaux et les cas limites
- Les chemins de gestion d'erreurs doivent être testés

**Exemple de Test**
```go
func TestFormatDuration(t *testing.T) {
    testCases := []struct {
        name      string
        duration  time.Duration
        precision int
        want      string
    }{
        {
            name:      "temps zéro",
            duration:  0,
            precision: 1,
            want:      "0 secondes",
        },
        {
            name:      "90 minutes haute précision",
            duration:  90 * time.Minute,
            precision: 2,
            want:      "1 heure 30 minutes",
        },
        // Plus de cas de test...
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            got := FormatDuration(tc.duration, tc.precision)
            assert.Equal(t, tc.want, got)
        })
    }
}

// Test de benchmark
func BenchmarkFormatDuration(b *testing.B) {
    duration := 90 * time.Minute
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = FormatDuration(duration, 2)
    }
}
```

## 🎯 Domaines de Développement Clés

### Haute Priorité

1. **Amélioration du Module xtime**
   - Amélioration des fonctionnalités de calendrier lunaire et de termes solaires
   - Optimisation des performances
   - Plus de fonctionnalités spécifiques à la culture

2. **Extension du Module candy**  
   - Fonctions de conversion de type
   - Outils de traitement de données
   - Optimisation des performances

3. **Amélioration de la Couverture de Test**
   - Objectif : Tous les modules > 90%
   - Supplément de cas limites
   - Amélioration des tests de performance

### Priorité Moyenne

4. **Nouveaux Modules Utilitaires**
   - Fonctions utilitaires AI/ML
   - Outils d'intégration de services cloud
   - Outils de microservices

5. **Amélioration de la Documentation**
   - Documentation de référence API
   - Guide des meilleures pratiques
   - Guide d'optimisation des performances

### Contributions Bienvenues

- 🌏 **Support Multi-langues** - Documentation anglaise, internationalisation des messages d'erreur
- 📊 **Plus de Support de Formats de Données** - Traitement XML, YAML, TOML
- 🔧 **Outils de Développement** - Génération de code, gestion de configuration
- 🎨 **Outils UI/UX** - Traitement des couleurs, sortie formatée
- 🔐 **Outils de Sécurité** - Chiffrement/déchiffrement, vérification de signature

## 📋 Modèle de PR

Veuillez utiliser le modèle suivant lors de la création d'une PR :

```markdown
## Description du Changement

Brève description du contenu et du but de ce changement.

## Type de Changement

- [ ] Correction de bug
- [ ] Nouvelle fonctionnalité
- [ ] Mise à jour de documentation
- [ ] Optimisation des performances  
- [ ] Refactorisation de code
- [ ] Amélioration des tests

## Changements Détaillés

### Nouvelles Fonctionnalités
- Ajouté la fonction `FormatDuration`
- Support de plusieurs niveaux de précision
- Ajouté l'affichage des unités de temps en chinois

### Problèmes Corrigés  
- Corrigé le bug de conversion de fuseau horaire (#123)
- Résolu le problème de fuite mémoire

### Optimisation des Performances
- Optimisé les performances de concaténation de chaînes
- Réduit l'allocation mémoire de 30%

## Description des Tests

- [ ] Tous les tests passent
- [ ] Ajouté de nouveaux cas de test
- [ ] Couverture de test ≥ 90%
- [ ] Tests de benchmark passent

**Couverture de Test** : 92.5%

## Mises à Jour de Documentation

- [ ] Mis à jour README.md
- [ ] Ajouté des commentaires de fonction
- [ ] Mis à jour le code d'exemple

## Compatibilité

- [ ] Compatible vers l'arrière
- [ ] Nécessite une mise à jour de version (expliquer la raison)
- [ ] Changements cassants (explication détaillée)

## Liste de Vérification

- [ ] Le code suit les normes du projet
- [ ] Passé la vérification de format `go fmt`
- [ ] Passé la vérification statique `go vet`
- [ ] Tous les tests passent
- [ ] Documentation mise à jour
- [ ] Messages de commit suivent les normes

## Problèmes Liés

Closes #123
Refs #456

## Captures d'écran/Démo

Fournir des captures d'écran ou des démos si nécessaire.
```

## 🐛 Rapports de Bugs

Trouvé un bug ? Veuillez utiliser le modèle suivant pour créer un Issue :

```markdown
## Description du Bug

Brève description du problème rencontré.

## Étapes de Reproduction

1. Exécuter l'étape 1
2. Exécuter l'étape 2  
3. Observer le résultat

## Comportement Attendu

Décrire le comportement correct que vous vous attendez à voir.

## Comportement Réel

Décrire le comportement erroné réellement observé.

## Informations d'Environnement

- **Système d'Exploitation** : macOS 12.0
- **Version Go** : 1.24.0
- **Version Utils** : v1.2.0
- **Autres informations pertinentes** :

## Journaux d'Erreur

```
coller les journaux d'erreur ici
```

## Exemple Minimal Reproductible

```go
package main

import (
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // Code minimal de reproduction d'erreur
}
```
```

## ✨ Demandes de Fonctionnalités

Vous voulez une nouvelle fonctionnalité ? Veuillez utiliser le modèle suivant :

```markdown
## Description de la Fonctionnalité

Décrire la fonctionnalité que vous aimeriez ajouter.

## Cas d'Utilisation

Décrire quand cette fonctionnalité serait utilisée.

## Conception d'API Suggérée

```go
// Signature de fonction suggérée et utilisation
func NewAwesomeFunction(param string) (Result, error) {
    // ...
}
```

## Solutions Alternatives

Avez-vous considéré d'autres solutions ?

## Informations Supplémentaires

Autres informations pertinentes ou références.
```

## 🏆 Reconnaissance des Contributeurs

### Reconnaissance par Type de Contribution

Nous donnerons différentes reconnaissances basées sur les types de contribution :

- 🥇 **Contributeurs Principaux** - Actifs à long terme, contributions de fonctionnalités importantes
- 🥈 **Contributeurs Actifs** - Multiples contributions précieuses  
- 🥉 **Contributeurs Communautaires** - Corrections de bugs, améliorations de documentation
- 🌟 **Premiers Contributeurs** - Accueil des premières contributions

### Statistiques de Contribution

Nous présenterons les contributeurs dans les endroits suivants :

- Liste des contributeurs README.md
- Remerciements dans les notes de version
- Site web du projet (si disponible)
- Rapports annuels des contributeurs

## 💬 Communication

### Obtenir de l'Aide

- 📖 **Problèmes de Documentation** : Vérifier README.md pour chaque module
- 🐛 **Rapports de Bugs** : [GitHub Issues](https://github.com/lazygophers/utils/issues)
- 💡 **Discussions de Fonctionnalités** : [GitHub Discussions](https://github.com/lazygophers/utils/discussions)
- ❓ **Questions d'Utilisation** : [GitHub Discussions Q&A](https://github.com/lazygophers/utils/discussions/categories/q-a)

### Normes de Discussion

Veuillez suivre ces normes de communication :

- Utiliser un langage amical et professionnel
- Fournir des descriptions de problèmes détaillées et des suggestions
- Fournir suffisamment d'informations de contexte
- Respecter différents points de vue et opinions
- Participer activement aux discussions constructives

## 📜 Licence

Ce projet est sous licence [GNU Affero General Public License v3.0](LICENSE).

**Contribuer signifie accepter** :
- Vous possédez le copyright du code soumis
- Acceptez de publier le code sous licence AGPL v3.0
- Suivez le code de conduite des contributeurs du projet

## 🙏 Remerciements

Merci à tous les développeurs qui ont contribué au projet LazyGophers Utils !

**Remerciements Spéciaux** :
- Tous les contributeurs qui ont soumis des Issues et PRs
- Membres de la communauté qui ont fourni des suggestions et des retours
- Bénévoles qui ont aidé à améliorer la documentation

---

**Disponible dans d'autres langues :** [English](CONTRIBUTING.md) | [简体中文](CONTRIBUTING_zh.md) | [繁體中文](CONTRIBUTING_zh-Hant.md) | [Русский](CONTRIBUTING_ru.md) | [Español](CONTRIBUTING_es.md) | [العربية](CONTRIBUTING_ar.md)

**Happy Coding ! 🎉**

N'hésitez pas à contacter l'équipe de mainteneurs à tout moment si vous avez des questions. Nous sommes heureux de vous aider à commencer votre parcours de contribution !