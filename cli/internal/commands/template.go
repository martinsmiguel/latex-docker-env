package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/martinsmiguel/latex-docker-env/cli/internal/colors"
	"github.com/martinsmiguel/latex-docker-env/cli/pkg/types"
)

var TemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Gerencia templates do LaTeX",
	Long:  `Comandos para listar, validar e gerenciar templates LaTeX.`,
}

var listTemplatesCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista templates disponíveis",
	Long:  `Lista todos os templates LaTeX disponíveis no sistema.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return listTemplates()
	},
}

var validateTemplateCmd = &cobra.Command{
	Use:   "validate [template-path]",
	Short: "Valida um template",
	Long:  `Valida a estrutura e metadados de um template.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return validateTemplate(args[0])
	},
}

func init() {
	TemplateCmd.AddCommand(listTemplatesCmd)
	TemplateCmd.AddCommand(validateTemplateCmd)
}

func listTemplates() error {
	registry := getTemplateRegistry()

	if err := registry.LoadTemplates(); err != nil {
		return fmt.Errorf("erro ao carregar templates: %w", err)
	}

	templates := registry.ListTemplates()

	if len(templates) == 0 {
		colors.PrintInfo("Nenhum template encontrado.")
		colors.Println("\n💡 Para adicionar templates:")
		colors.Println("   1. Extraia o template na pasta 'templates/' ou 'user-templates/'")
		colors.Println("   2. Certifique-se que contém um arquivo 'template.yaml'")
		colors.Println("   3. Execute 'ltx template list' novamente")
		return nil
	}

	colors.Println(">> Templates disponíveis:")
	colors.Println("")

	// Agrupar por tipo
	typeGroups := make(map[string][]*types.Template)
	for _, tmpl := range templates {
		typeGroups[tmpl.Metadata.Type] = append(typeGroups[tmpl.Metadata.Type], tmpl)
	}

	for templateType, templates := range typeGroups {
		colors.Printf("📂 %s\n", strings.ToUpper(templateType))
		for _, tmpl := range templates {
			colors.Printf("  📄 %s\n", tmpl.Metadata.Name)
			colors.Printf("     %s\n", tmpl.Metadata.Description)
			colors.Printf("     Por: %s (v%s)\n", tmpl.Metadata.Author, tmpl.Metadata.Version)
			if len(tmpl.Metadata.Dependencies) > 0 {
				colors.Printf("     Deps: %s\n", strings.Join(tmpl.Metadata.Dependencies, ", "))
			}
			colors.Printf("     📍 %s\n", tmpl.Path)
		}
		colors.Println("")
	}

	colors.PrintInfo(fmt.Sprintf("Total: %d templates encontrados", len(templates)))

	return nil
}

func validateTemplate(templatePath string) error {
	registry := getTemplateRegistry()

	colors.Printf(">> Validando template em: %s\n", templatePath)

	// Tentar carregar o template específico usando método público
	if err := registry.LoadTemplates(); err != nil {
		colors.PrintError(fmt.Sprintf("❌ Erro ao carregar templates: %v", err))
		return err
	}

	// Verificar se foi carregado com sucesso
	dirName := filepath.Base(templatePath)
	tmpl, err := registry.GetTemplate(dirName)
	if err != nil {
		colors.PrintError(fmt.Sprintf("❌ Erro ao encontrar template: %v", err))
		return err
	}

	// Verificar se template.yaml existe
	metadataPath := filepath.Join(templatePath, "template.yaml")
	hasMetadata := true
	if _, err := os.Stat(metadataPath); os.IsNotExist(err) {
		hasMetadata = false
		colors.Printf("   ℹ️  Template auto-detectado (sem template.yaml)\n")
	}

	colors.PrintSuccess("✅ Template válido!")
	colors.Printf("   Nome: %s\n", tmpl.Metadata.Name)
	colors.Printf("   Tipo: %s\n", tmpl.Metadata.Type)
	colors.Printf("   Autor: %s\n", tmpl.Metadata.Author)
	colors.Printf("   Versão: %s\n", tmpl.Metadata.Version)
	colors.Printf("   Arquivos: %d\n", len(tmpl.Metadata.Files))

	if hasMetadata {
		// Validar arquivos apenas se há metadados explícitos
		var missingFiles []string
		for _, file := range tmpl.Metadata.Files {
			sourcePath := filepath.Join(templatePath, file.Source)
			if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
				if file.Required {
					missingFiles = append(missingFiles, file.Source)
				} else {
					colors.Printf("   ⚠️  Arquivo opcional não encontrado: %s\n", file.Source)
				}
			}
		}

		if len(missingFiles) > 0 {
			colors.PrintError("❌ Arquivos obrigatórios não encontrados:")
			for _, file := range missingFiles {
				colors.Printf("   - %s\n", file)
			}
			return fmt.Errorf("template incompleto")
		}

		colors.PrintSuccess("✅ Todos os arquivos obrigatórios estão presentes")
	} else {
		colors.Printf("   ℹ️  Validação baseada na auto-detecção\n")
		// Para templates auto-detectados, apenas verificar se existe um arquivo .tex principal
		hasMainTex := false
		entries, err := os.ReadDir(templatePath)
		if err == nil {
			for _, entry := range entries {
				if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".tex") {
					hasMainTex = true
					break
				}
			}
		}

		if hasMainTex {
			colors.PrintSuccess("✅ Template contém arquivos LaTeX (.tex)")
		} else {
			colors.PrintError("❌ Nenhum arquivo .tex encontrado")
			return fmt.Errorf("template sem arquivos LaTeX")
		}
	}

	return nil
}
