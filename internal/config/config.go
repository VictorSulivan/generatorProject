package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	TemplatesDir string
	OutputDir    string
}

func Load() *Config {
	config := &Config{
		TemplatesDir: getEnvOrDefault("GENERATOR_TEMPLATES_DIR", "templates"),
		OutputDir:    getEnvOrDefault("GENERATOR_OUTPUT_DIR", ""),
	}

	// Charger la configuration depuis le fichier YAML si présent
	configFile := filepath.Join(os.Getenv("HOME"), ".config", "generator", "config.yaml")
	if _, err := os.Stat(configFile); err == nil {
		// TODO: Implémenter le chargement de la configuration YAML
	}

	return config
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
} 