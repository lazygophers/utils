# Politique de Sécurité

## Versions Supportées

Nous maintenons activement et fournissons des mises à jour de sécurité pour les versions suivantes de LazyGophers Utils :

| Version | Supportée          | Version Go Requise  | Statut           |
| ------- | ------------------ | ------------------- | ---------------- |
| 1.x.x   | :white_check_mark: | Go 1.24+           | Développement Actif |
| 0.x.x   | :white_check_mark: | Go 1.24+           | Corrections de Sécurité Uniquement |

## Considérations de Sécurité

### Composants Cryptographiques

Cette bibliothèque inclut des utilitaires cryptographiques dans le package `cryptox` :

- **Chiffrement/déchiffrement AES** : Utilise la bibliothèque crypto standard de Go
- **Génération et opérations de clés RSA** : Implémente les pratiques standard de l'industrie
- **Fonctions de hachage** : SHA-256, SHA-512 et autres algorithmes sécurisés
- **Blowfish et ChaCha20** : Algorithmes de chiffrement supplémentaires
- **PGP/GPG** : Implémentation OpenPGP pour la messagerie sécurisée

⚠️ **Important** : Bien que nous suivions les meilleures pratiques de sécurité, effectuez toujours votre propre révision de sécurité avant d'utiliser les fonctions cryptographiques en environnement de production.

### Validation des Entrées

Les fonctions utilitaires de cette bibliothèque (en particulier dans les packages `candy`, `stringx` et `config`) effectuent des conversions de types et l'analyse de données. Bien que nous implémentions des pratiques de programmation défensive :

- La désinfection des entrées est effectuée le cas échéant
- Les échecs de conversion de type sont gérés gracieusement
- Le chargement de configuration inclut des étapes de validation

### Dépendances

Nous auditons régulièrement nos dépendances pour les vulnérabilités connues :

- Toutes les dépendances sont épinglées à des versions spécifiques
- Nous utilisons `go mod tidy` et `go mod verify` dans notre pipeline CI/CD
- Le scan de sécurité est effectué via golangci-lint et gosec

## Signaler une Vulnérabilité

### Comment Signaler

Si vous découvrez une vulnérabilité de sécurité dans LazyGophers Utils, veuillez la signaler de manière responsable :

1. **Email** : Envoyez les détails à `security@lazygophers.com` (si disponible) ou créez un problème privé
2. **Avis de Sécurité GitHub** : Utilisez la fonction de rapport de vulnérabilité privée de GitHub
3. **Contact Direct** : Contactez directement les mainteneurs via GitHub

### Informations à Inclure

Veuillez inclure les informations suivantes dans votre rapport :

- **Description** : Description claire de la vulnérabilité
- **Localisation** : Package/fichier/fonction spécifique affecté
- **Impact** : Impact de sécurité potentiel et vecteurs d'attaque
- **Reproduction** : Étapes pour reproduire la vulnérabilité
- **Correction Suggérée** : Si vous avez des idées pour la remédiation

### Chronologie de Réponse

Nous nous engageons à traiter les problèmes de sécurité rapidement :

- **Accusé de réception** : Dans les 48 heures du rapport
- **Évaluation Initiale** : Dans les 5 jours ouvrables
- **Chronologie de Résolution** :
  - Vulnérabilités critiques : 7-14 jours
  - Haute gravité : 14-30 jours
  - Gravité moyenne/faible : 30-90 jours

### Politique de Divulgation

Nous suivons les pratiques de divulgation responsable :

1. Nous travaillerons avec vous pour comprendre et reproduire le problème
2. Nous développerons et testerons une correction
3. Nous coordonnerons le timing de divulgation avec vous
4. Crédit sera donné aux chercheurs qui signalent les vulnérabilités de manière responsable

## Meilleures Pratiques de Sécurité pour les Utilisateurs

### Directives Générales

- **Restez à Jour** : Utilisez toujours la dernière version de la bibliothèque
- **Révisez le Code** : Effectuez des révisions de sécurité pour les cas d'usage en production
- **Validez les Entrées** : Validez toujours les entrées externes dans vos applications
- **Suivez les Principes** : Appliquez les principes de sécurité de défense en profondeur

### Usage Cryptographique

Lors de l'utilisation du package `cryptox` :

- **Gestion des Clés** : Utilisez des pratiques sécurisées de génération et stockage de clés
- **Nombres Aléatoires** : Assurez-vous d'une entropie appropriée pour les opérations cryptographiques
- **Sélection d'Algorithme** : Choisissez des algorithmes appropriés pour votre modèle de menace
- **Implémentation** : Suivez les meilleures pratiques cryptographiques dans votre application

### Sécurité de Configuration

Lors de l'utilisation du package `config` :

- **Permissions de Fichier** : Restreignez l'accès aux fichiers de configuration contenant des secrets
- **Variables d'Environnement** : Utilisez des méthodes sécurisées pour stocker la configuration sensible
- **Validation** : Validez toujours les données de configuration avant utilisation

## Tests de Sécurité

### Scan de Sécurité Automatisé

Notre pipeline CI/CD inclut :

- **Analyse Statique** : golangci-lint avec des linters axés sur la sécurité
- **Scan de Vulnérabilités** : gosec pour les problèmes de sécurité spécifiques à Go
- **Scan de Dépendances** : Vérifications régulières des dépendances vulnérables
- **Qualité du Code** : Linting et tests complets

### Révision de Sécurité Manuelle

Nous effectuons des révisions de sécurité manuelles pour :

- Toutes les implémentations cryptographiques
- Logique de validation et d'analyse des entrées
- Gestion des erreurs et divulgation d'informations
- Modèles d'authentification et d'autorisation

## Informations de Contact

Pour les questions ou préoccupations liées à la sécurité :

- **Dépôt du Projet** : [https://github.com/lazygophers/utils](https://github.com/lazygophers/utils)
- **Problèmes** : Utilisez GitHub Issues pour les bugs non liés à la sécurité
- **Rapports de Sécurité** : Suivez le processus de rapport de vulnérabilité ci-dessus

## Journal des Modifications

### Mises à Jour de Sécurité

Nous maintenons un journal des mises à jour de sécurité pour la transparence :

- **Version 1.0.x** : Révision de sécurité initiale et durcissement
- **Versions futures** : Les mises à jour de sécurité seront documentées ici

---

**Disponible dans d'autres langues :** [English](SECURITY.md) | [简体中文](SECURITY_zh.md) | [繁體中文](SECURITY_zh-Hant.md) | [Русский](SECURITY_ru.md) | [Español](SECURITY_es.md) | [العربية](SECURITY_ar.md)

**Note** : Cette politique de sécurité est sujette à des mises à jour. Veuillez vérifier la dernière version dans le dépôt.

Dernière mise à jour : 2025-09-13