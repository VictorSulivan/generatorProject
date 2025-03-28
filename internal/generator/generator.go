package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"generator/internal/config"
)

type Generator struct {
	config *config.Config
}

func New(cfg *config.Config) *Generator {
	return &Generator{
		config: cfg,
	}
}

func (g *Generator) Generate(templateName, projectName string) error {
	// Vérifier si le template existe
	templatePath := filepath.Join(g.config.TemplatesDir, templateName)
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return fmt.Errorf("template introuvable : %s", templateName)
	}

	// Déterminer le dossier de sortie
	var destPath string
	if g.config.OutputDir != "" {
		destPath = filepath.Join(g.config.OutputDir, projectName)
	} else {
		destPath = projectName
	}

	// Vérifier si le dossier de sortie existe déjà
	if _, err := os.Stat(destPath); !os.IsNotExist(err) {
		return fmt.Errorf("le dossier %s existe déjà", destPath)
	}

	// Copier récursive des fichiers du template vers le projet
	return g.copyDir(templatePath, destPath, projectName)
}

func (g *Generator) copyDir(src, dst, projectName string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(src, path)
		destPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}

		// Lire le contenu du fichier
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Remplacer les variables si c'est un template
		content := string(data)
		content = strings.ReplaceAll(content, "{{.ProjectName}}", projectName)

		// Écrire le fichier de sortie
		return os.WriteFile(destPath, []byte(content), info.Mode())
	})
}

func (g *Generator) ListTemplates() ([]string, error) {
	entries, err := os.ReadDir(g.config.TemplatesDir)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture des templates : %v", err)
	}

	var templates []string
	for _, entry := range entries {
		if entry.IsDir() {
			templates = append(templates, entry.Name())
		}
	}

	return templates, nil
}

func (g *Generator) GetTemplateInfo(templateName string) (string, error) {
	readmePath := filepath.Join(g.config.TemplatesDir, templateName, "README.md")
	
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		return "", fmt.Errorf("template introuvable : %s", templateName)
	}

	data, err := os.ReadFile(readmePath)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la lecture du README : %v", err)
	}

	return string(data), nil
} 