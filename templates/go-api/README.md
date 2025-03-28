# Template Go API

Ce template fournit une base solide pour créer des APIs RESTful avec Go, incluant une architecture propre et des bonnes pratiques.

## 🚀 Fonctionnalités

- 🛠️ Gin pour le routing HTTP
- 📝 Swagger pour la documentation API
- 🔐 GORM pour l'ORM
- 🔑 JWT pour l'authentification
- 📦 Structure modulaire
- 🧪 Tests unitaires et d'intégration

## 📁 Structure du Projet

```
.
├── cmd/
│   └── api/           # Point d'entrée de l'application
├── internal/
│   ├── handlers/      # Gestionnaires HTTP
│   ├── middleware/    # Middlewares (auth, logging, etc.)
│   ├── models/        # Modèles de données
│   ├── repository/    # Couche d'accès aux données
│   └── service/       # Logique métier
├── pkg/
│   ├── config/        # Configuration
│   └── utils/         # Utilitaires
├── api/
│   ├── routes/        # Définition des routes
│   └── swagger/       # Documentation API
└── docs/             # Documentation générale
```

## 🛠️ Installation

```bash
# Installation des dépendances
go mod download

# Démarrage du serveur
go run cmd/api/main.go

# Génération de la documentation Swagger
swag init
```

## 📝 Bonnes Pratiques

### Architecture

1. **Clean Architecture**
   - Séparation claire des responsabilités
   - Dépendances vers l'intérieur
   - Couches indépendantes

2. **Repository Pattern**
   - Abstraction de la couche de données
   - Facilite les tests
   - Permet de changer de base de données

### Organisation du Code

1. **Handlers** (`internal/handlers/`)
   - Gestion des requêtes HTTP
   - Validation des entrées
   - Transformation des réponses

2. **Services** (`internal/service/`)
   - Logique métier
   - Orchestration des opérations
   - Validation des règles métier

3. **Repositories** (`internal/repository/`)
   - Accès aux données
   - Requêtes complexes
   - Cache si nécessaire

### Gestion des Erreurs

- Utiliser des erreurs personnalisées
- Middleware de gestion d'erreurs global
- Logging structuré

### Tests

- Tests unitaires pour chaque couche
- Tests d'intégration pour les flux complets
- Mocks pour les dépendances externes

## 🔧 Configuration

### Variables d'Environnement

Copiez `.env.example` vers `.env` et configurez :

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=app
```

### Base de Données

1. Créer la base de données :
   ```sql
   CREATE DATABASE app;
   ```

2. Exécuter les migrations :
   ```bash
   go run cmd/migrate/main.go
   ```

## 🧪 Tests

```bash
# Tests unitaires
go test ./...

# Tests avec couverture
go test ./... -cover

# Tests d'intégration
go test ./... -tags=integration
```

## 📦 Déploiement

### Docker

```bash
# Build de l'image
docker build -t app .

# Exécution du conteneur
docker run -p 8080:8080 app
```

### Kubernetes

1. Créer les manifests :
   ```bash
   kubectl apply -f k8s/
   ```

2. Vérifier le déploiement :
   ```bash
   kubectl get pods
   ```

## 🔍 Linting

```bash
# Installation de golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Vérification du code
golangci-lint run
```

## 📚 Ressources

- [Documentation Go](https://golang.org/doc/)
- [Documentation Gin](https://gin-gonic.com/)
- [Documentation GORM](https://gorm.io/)
- [Documentation Swagger](https://swaggo.github.io/swaggo.io/) 