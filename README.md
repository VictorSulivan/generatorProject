# Generator Project

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

1. Cloner le projet :
```bash
git clone https://github.com/votre-username/generator-project.git
cd generator-project
```

2. Compiler le projet :
```bash
go build -o generator
```

3. Installer l'exécutable globalement :
```bash
sudo cp generator /usr/local/bin/
sudo chmod +x /usr/local/bin/generator
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

Une fois installé, vous pouvez utiliser la commande `generator` depuis n'importe où :

```bash
# Lister les templates disponibles
generator list

# Voir les informations sur un template
generator info go-api

# Générer un nouveau projet
generator generate go-api mon-api

# Générer dans un dossier spécifique
generator generate go-api mon-api --output /chemin/vers/dossier
```

## 🧪 Tests

```bash
# Exécuter tous les tests
go test ./...

# Exécuter les tests avec couverture
go test ./... -cover
```

## 🔧 Configuration

Le programme utilise deux emplacements principaux :

1. **Templates** : Les templates sont stockés dans le dossier `templates` du projet. Le programme trouve automatiquement ce dossier en fonction de l'emplacement de l'exécutable.

2. **Configuration** : Un fichier de configuration est créé automatiquement dans `~/.config/generator/config.yaml` lors de la première utilisation.

### Variables d'environnement

Vous pouvez personnaliser le comportement du programme en utilisant les variables d'environnement suivantes :

- `GENERATOR_TEMPLATES_DIR` : Chemin vers le dossier des templates (par défaut : dossier `templates` du projet)
- `GENERATOR_OUTPUT_DIR` : Répertoire par défaut pour les nouveaux projets générés

### Fichier de configuration

Le fichier de configuration par défaut (`~/.config/generator/config.yaml`) contient :
```yaml
# Configuration par défaut de generator
output_dir: ~/projects  # Répertoire par défaut pour les nouveaux projets
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

## 🤝 Contribution et Fork

Ce projet est open source et peut être librement forké et utilisé selon vos besoins. Les contributions sont les bienvenues ! N'hésitez pas à :
1. Fork le projet
2. Créer une branche pour votre fonctionnalité
3. Commiter vos changements
4. Pousser vers la branche
5. Ouvrir une Pull Request

## 📝 Licence

Ce projet est sous licence MIT. Voir le fichier [LICENSE](LICENSE) pour plus de détails.

## 📞 Support et Contact

Pour toute question, suggestion ou discussion sur le projet, vous pouvez :
- Ouvrir une issue sur GitHub
- Me contacter sur LinkedIn : [Votre Nom](https://www.linkedin.com/in/votre-profil)
