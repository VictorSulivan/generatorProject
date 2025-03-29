package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	TemplatesDir string
	OutputDir    string
}

// getProjectRoot retourne le chemin racine du projet
func getProjectRoot() string {
	// Obtenir le chemin de l'exécutable
	exe, err := os.Executable()
	if err != nil {
		return "."
	}

	// Si l'exécutable est dans /usr/local/bin, le projet est dans ~/project/DevProject/generator-project
	if filepath.Dir(exe) == "/usr/local/bin" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "."
		}
		return filepath.Join(homeDir, "project", "DevProject", "generator-project")
	}

	// Sinon, chercher le fichier go.mod en remontant dans l'arborescence
	dir := filepath.Dir(exe)
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "."
}

// getTemplatesPath retourne le chemin vers le dossier templates
func getTemplatesPath() string {
	projectRoot := getProjectRoot()
	return filepath.Join(projectRoot, "templates")
}

// createDefaultConfig crée un fichier de configuration par défaut
func createDefaultConfig(configDir string) error {
	defaultConfig := `# Configuration par défaut de generator
output_dir: ~/projects  # Répertoire par défaut pour les nouveaux projets
`
	
	// Créer le répertoire de configuration s'il n'existe pas
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}
	
	// Créer le fichier de configuration
	configFile := filepath.Join(configDir, "config.yaml")
	return os.WriteFile(configFile, []byte(defaultConfig), 0644)
}

func Load() *Config {
	config := &Config{
		TemplatesDir: getEnvOrDefault("GENERATOR_TEMPLATES_DIR", getTemplatesPath()),
		OutputDir:    getEnvOrDefault("GENERATOR_OUTPUT_DIR", ""),
	}

	// Définir le répertoire de configuration
	configDir := filepath.Join(os.Getenv("HOME"), ".config", "generator")
	configFile := filepath.Join(configDir, "config.yaml")

	// Créer la configuration par défaut si elle n'existe pas
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		if err := createDefaultConfig(configDir); err != nil {
			// En cas d'erreur, on continue avec les valeurs par défaut
			return config
		}
	}

	// TODO: Implémenter le chargement de la configuration YAML
	return config
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
} 