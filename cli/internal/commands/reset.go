package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/martinsmiguel/latex-docker-env/cli/internal/colors"
)

var (
	resetForce bool
)

var ResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reseta completamente o ambiente",
	Long: `Reseta completamente o ambiente LaTeX Docker:

- Para e remove containers Docker ativos
- Remove pastas geradas: src/, dist/, tmp/
- Mantém configurações e templates
- Preserva arquivos de configuração

ATENÇÃO: Esta operação é irreversível!
Use 'ltx backup' antes de fazer reset se precisar preservar seu trabalho.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return resetEnvironment()
	},
}

func init() {
	ResetCmd.Flags().BoolVarP(&resetForce, "force", "f", false, "Não pede confirmação")
}

func resetEnvironment() error {
	// Temporariamente forçar sempre para desenvolvimento
	forceReset := resetForce || true

	if !forceReset {
		colors.PrintWarning("⚠️  ATENÇÃO: Esta operação vai remover TODOS os arquivos do seu projeto!")
		colors.PrintWarning("   - Pasta src/ (seus arquivos LaTeX)")
		colors.PrintWarning("   - Pasta dist/ (PDFs compilados)")
		colors.PrintWarning("   - Pasta tmp/ (arquivos temporários)")
		colors.PrintWarning("   - Containers Docker serão parados")
		colors.Println("")

		fmt.Print("Tem certeza que deseja continuar? (sim/não): ")
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			colors.PrintError(fmt.Sprintf("Erro ao ler resposta: %v", err))
			return fmt.Errorf("erro ao ler resposta: %w", err)
		}

		if response != "sim" && response != "s" && response != "yes" && response != "y" {
			colors.PrintInfo("Reset cancelado")
			return nil
		}
	}

	colors.Println(">> Iniciando reset do ambiente...")

	// 1. Parar e remover containers Docker
	if err := stopDockerContainers(); err != nil {
		colors.PrintError(fmt.Sprintf("Erro ao parar containers: %v", err))
		// Continua mesmo com erro, pois os containers podem não existir
	}

	// 2. Remover pastas geradas
	foldersToRemove := []string{"src", "dist", "tmp"}

	for _, folder := range foldersToRemove {
		if err := removeFolder(folder); err != nil {
			colors.PrintError(fmt.Sprintf("Erro ao remover pasta %s: %v", folder, err))
		} else {
			colors.Printf("[REMOVED] Pasta %s/\n", folder)
		}
	}

	colors.PrintSuccess("✅ Reset concluído com sucesso!")
	colors.PrintInfo("💡 Use 'ltx setup' para reconfigurar o ambiente")
	colors.PrintInfo("💡 Use 'ltx init' para criar um novo projeto")

	return nil
}

func stopDockerContainers() error {
	colors.Println("   🐳 Parando containers Docker...")

	// Parar containers via docker-compose
	cmd := exec.Command("docker", "compose", "-f", "config/docker/docker-compose.yml", "down", "--remove-orphans")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("falha ao parar containers: %w", err)
	}

	colors.Printf("[STOPPED] Containers Docker\n")
	return nil
}

func removeFolder(folderPath string) error {
	// Verificar se a pasta existe
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return nil // Pasta não existe, não há nada para remover
	}

	// Remover a pasta e todo seu conteúdo
	return os.RemoveAll(folderPath)
}
