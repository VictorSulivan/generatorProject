package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"generator/internal/config"

	"github.com/spf13/cobra"
)

// infoCmd représente la commande "info"
var infoCmd = &cobra.Command{
	Use:   "info [template]",
	Short: "Affiche les informations détaillées sur un template",
	Long: `Affiche les informations détaillées sur un template spécifique, incluant :
  - Description
  - Fonctionnalités
  - Structure du projet
  - Instructions d'installation
  - Bonnes pratiques

Exemple:
  generator info react-ts`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		templateName := args[0]
		showTemplateInfo(templateName)
	},
}

func showTemplateInfo(templateName string) {
	// Charger la configuration
	cfg := config.Load()
	readmePath := filepath.Join(cfg.TemplatesDir, templateName, "README.md")
	
	// Vérifier si le template existe
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		fmt.Printf("❌ Template introuvable : %s\n", templateName)
		fmt.Println("Utilisez 'generator list' pour voir les templates disponibles")
		return
	}

	// Lire le contenu du README
	data, err := os.ReadFile(readmePath)
	if err != nil {
		fmt.Printf("❌ Erreur lors de la lecture du README : %v\n", err)
		return
	}

	// Afficher le contenu formaté
	content := string(data)
	fmt.Printf("📚 Informations sur le template %s :\n\n", templateName)
	fmt.Println(content)
} 