package cmd

import (
	"fmt"
	"os"

	"generator/internal/config"
	"generator/internal/generator"

	"github.com/spf13/cobra"
)

var (
	outputDir string
)

// generateCmd représente la commande "generate"
var generateCmd = &cobra.Command{
	Use:   "generate [template] [project_name]",
	Short: "Génère un nouveau projet à partir d'un template",
	Long: `Génère un nouveau projet à partir d'un template prédéfini.

Templates disponibles:
  - react-ts: Template React avec TypeScript
  - go-api: Template API REST en Go
  - cpp-cmake: Template C++ avec CMake
  - rust-cli: Template CLI en Rust

Exemples:
  generator generate react-ts mon-app
  generator generate go-api mon-api --output ./projects`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		templateName := args[0]
		projectName := args[1]

		// Charger la configuration
		cfg := config.Load()
		if outputDir != "" {
			cfg.OutputDir = outputDir
		}

		// Créer le générateur
		gen := generator.New(cfg)

		// Générer le projet
		if err := gen.Generate(templateName, projectName); err != nil {
			fmt.Printf("❌ Erreur : %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Projet %s créé avec succès !\n", projectName)
	},
}

func init() {
	generateCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Dossier de sortie (optionnel)")
}
