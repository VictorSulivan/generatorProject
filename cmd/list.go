package cmd

import (
	"fmt"
	"os"
	"strings"

	"generator/internal/config"
	"generator/internal/generator"

	"github.com/spf13/cobra"
)

// listCmd représente la commande "list"
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Liste tous les templates disponibles",
	Long: `Affiche la liste de tous les templates disponibles avec leurs descriptions.

Exemple:
  generator list`,
	Run: func(cmd *cobra.Command, args []string) {
		// Charger la configuration
		cfg := config.Load()

		// Créer le générateur
		gen := generator.New(cfg)

		// Lister les templates
		templates, err := gen.ListTemplates()
		if err != nil {
			fmt.Printf("❌ Erreur : %v\n", err)
			os.Exit(1)
		}

		fmt.Println("📦 Templates disponibles :")
		fmt.Println()

		for _, template := range templates {
			// Lire la description du template
			info, err := gen.GetTemplateInfo(template)
			if err != nil {
				fmt.Printf("  • %s\n    Aucune description disponible\n\n", template)
				continue
			}

			// Extraire la première ligne non vide du README
			lines := strings.Split(info, "\n")
			description := "Aucune description disponible"
			for _, line := range lines {
				if strings.TrimSpace(line) != "" {
					description = strings.TrimSpace(line)
					break
				}
			}

			fmt.Printf("  • %s\n    %s\n\n", template, description)
		}
	},
} 