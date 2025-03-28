# Template Rust CLI

Ce template fournit une base solide pour créer des applications CLI en Rust, incluant une structure de projet professionnelle et des bonnes pratiques.

## 🚀 Fonctionnalités

- 🛠️ clap pour le parsing des arguments
- 📝 anyhow pour la gestion des erreurs
- 🔄 tokio pour l'asynchrone
- 📦 serde pour la sérialisation
- 🧪 Tests unitaires et d'intégration
- 📚 Documentation avec rustdoc

## 📁 Structure du Projet

```
.
├── src/
│   ├── commands/     # Commandes CLI
│   ├── config/       # Configuration
│   ├── error/        # Gestion des erreurs
│   ├── utils/        # Utilitaires
│   └── main.rs       # Point d'entrée
├── tests/            # Tests d'intégration
├── examples/         # Exemples d'utilisation
└── docs/            # Documentation
```

## 🛠️ Installation

```bash
# Installation des dépendances
cargo build

# Exécution
cargo run -- --command hello

# Tests
cargo test
```

## 📝 Bonnes Pratiques

### Architecture

1. **Command Pattern**
   - Chaque commande dans un module séparé
   - Interface commune pour les commandes
   - Facilite l'ajout de nouvelles commandes

2. **Error Handling**
   - Erreurs personnalisées avec thiserror
   - Propagation d'erreurs avec anyhow
   - Messages d'erreur clairs

### Organisation du Code

1. **Commandes** (`src/commands/`)
   - Implémentation des commandes
   - Validation des arguments
   - Logique métier

2. **Configuration** (`src/config/`)
   - Chargement de la configuration
   - Variables d'environnement
   - Valeurs par défaut

3. **Erreurs** (`src/error/`)
   - Types d'erreurs personnalisés
   - Conversions d'erreurs
   - Messages d'erreur

### Gestion des États

- Configuration globale
- Cache si nécessaire
- État de la commande

## 🔧 Configuration

### Cargo.toml

```toml
[package]
name = "{{.ProjectName}}"
version = "0.1.0"
edition = "2021"

[dependencies]
clap = { version = "4.4", features = ["derive"] }
anyhow = "1.0"
tokio = { version = "1.0", features = ["full"] }
```

### Arguments CLI

```rust
#[derive(Parser, Debug)]
struct Args {
    #[arg(short, long)]
    command: String,
}
```

## 🧪 Tests

### Tests Unitaires

```rust
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_command() {
        assert_eq!(2 + 2, 4);
    }
}
```

### Tests d'Intégration

```rust
#[test]
fn test_cli() {
    let mut cmd = Command::cargo_bin("{{.ProjectName}}").unwrap();
    cmd.arg("--command").arg("hello");
    cmd.assert().success();
}
```

## 📦 Déploiement

### Build Release

```bash
# Build optimisé
cargo build --release

# Installation
cargo install --path .
```

### Distribution

1. Créer une release :
   ```bash
   cargo publish
   ```

2. Installation depuis crates.io :
   ```bash
   cargo install {{.ProjectName}}
   ```

## 🔍 Linting

### rustfmt

```bash
# Vérification du format
cargo fmt -- --check

# Correction automatique
cargo fmt
```

### clippy

```bash
# Vérification
cargo clippy

# Correction automatique
cargo clippy --fix
```

## 📚 Documentation

### Documentation du Code

```rust
/// Description de la fonction
/// 
/// # Arguments
/// 
/// * `arg` - Description de l'argument
/// 
/// # Returns
/// 
/// Description du retour
/// 
/// # Examples
/// 
/// ```
/// let result = function(arg);
/// assert_eq!(result, expected);
/// ```
```

### Documentation du Projet

```bash
# Génération de la documentation
cargo doc --open
```

## 📚 Ressources

- [Documentation Rust](https://www.rust-lang.org/learn)
- [Documentation clap](https://docs.rs/clap)
- [Documentation tokio](https://docs.rs/tokio)
- [Documentation serde](https://docs.rs/serde) 