# Generator

Un outil CLI pour générer rapidement des projets à partir de templates prédéfinis.

## 🚀 Fonctionnalités

- 📦 Génération de projets à partir de templates prédéfinis
- 🔍 Liste des templates disponibles
- ℹ️ Informations détaillées sur chaque template
- 🎯 Support de plusieurs types de projets :
  - React TypeScript
  - Go API
  - C++ CMake
  - Rust CLI

## 📋 Prérequis

- Go 1.20 ou supérieur
- Git

## 🛠️ Installation

### Installation locale

```bash
# Cloner le repository
git clone https://github.com/votre-username/generator.git
cd generator

# Installer les dépendances
go mod download

# Compiler le projet
go build -o generator

# Installer globalement (optionnel)
sudo mv generator /usr/local/bin/
```

### Installation via Go

```bash
go install github.com/votre-username/generator@latest
```

## 📁 Structure du Projet

```
.
├── cmd/                    # Commandes CLI
│   ├── generate.go        # Commande de génération
│   ├── info.go           # Commande d'information
│   ├── list.go           # Commande de liste
│   └── root.go           # Commande racine
├── internal/              # Code interne
│   ├── config/           # Configuration
│   └── generator/        # Logique de génération
├── templates/            # Templates de projets
│   ├── react-ts/        # Template React TypeScript
│   ├── go-api/          # Template Go API
│   ├── cpp-cmake/       # Template C++ CMake
│   └── rust-cli/        # Template Rust CLI
├── go.mod               # Dépendances Go
├── go.sum               # Sommes de contrôle
└── main.go             # Point d'entrée
```

## 🎯 Utilisation

### Commandes Disponibles

```bash
# Afficher l'aide
generator --help

# Générer un nouveau projet
generator generate [template] [project_name] [flags]

# Lister les templates disponibles
generator list

# Afficher les informations sur un template
generator info [template]
```

### Exemples

```bash
# Générer un projet React TypeScript
generator generate react-ts mon-app

# Générer un projet Go API dans un dossier spécifique
generator generate go-api mon-api --output ./projects

# Voir les templates disponibles
generator list

# Voir les informations sur le template React TypeScript
generator info react-ts
```

## 🧪 Tests

```bash
# Exécuter tous les tests
go test ./...

# Exécuter les tests avec couverture
go test ./... -cover
```

## 🔧 Configuration

### Variables d'Environnement

- `GENERATOR_TEMPLATES_DIR` : Chemin vers le dossier des templates (par défaut : "./templates")
- `GENERATOR_OUTPUT_DIR` : Dossier de sortie par défaut (optionnel)

### Fichier de Configuration

Créer un fichier `config.yaml` dans le dossier de configuration :

```yaml
templates_dir: "./templates"
output_dir: "./projects"
```

## 📦 Templates Disponibles

### React TypeScript
- React 18
- TypeScript
- Vite
- TailwindCSS
- React Query
- React Router

### Go API
- Gin
- GORM
- Swagger
- JWT
- PostgreSQL
- Tests

### C++ CMake
- C++17
- CMake
- Google Test
- Doxygen
- Clang Format

### Rust CLI
- Clap
- Error handling
- Tests
- Documentation
- Logging

## 🤝 Contribution

1. Fork le projet
2. Créer une branche pour votre fonctionnalité (`git checkout -b feature/AmazingFeature`)
3. Commit vos changements (`git commit -m 'Add some AmazingFeature'`)
4. Push vers la branche (`git push origin feature/AmazingFeature`)
5. Ouvrir une Pull Request

## 📝 Licence

Ce projet est sous licence MIT. Voir le fichier [LICENSE](LICENSE) pour plus de détails.

## 📞 Support

Pour toute question ou problème, veuillez ouvrir une issue sur GitHub.
