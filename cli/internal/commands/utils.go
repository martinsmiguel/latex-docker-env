package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	cleanAll bool
)

var CleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove arquivos temporários",
	Long: `Remove arquivos temporários gerados durante a compilação.

Por padrão, remove apenas arquivos auxiliares (.aux, .log, etc.).
Use --all para remover também o PDF gerado.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cleanProject()
	},
}

var ShellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Acessa o shell do container",
	Long: `Abre um shell interativo dentro do container LaTeX.

Útil para:
- Instalar pacotes LaTeX adicionais
- Executar comandos personalizados
- Debug de problemas de compilação`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return openShell()
	},
}

var LogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Mostra logs do container",
	Long: `Exibe os logs do container LaTeX Docker.

Útil para debug de problemas de compilação ou inicialização.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return showLogs()
	},
}

func init() {
	CleanCmd.Flags().BoolVar(&cleanAll, "all", false, "Remove também o PDF gerado")
}

func cleanProject() error {
	fmt.Println(">> Limpando arquivos temporários...")

	distDir := "dist"
	if _, err := os.Stat(distDir); os.IsNotExist(err) {
		fmt.Println("[INFO] Diretório dist não existe, nada para limpar")
		return nil
	}

	// Padrões de arquivos temporários
	tempPatterns := []string{
		"*.aux", "*.log", "*.bbl", "*.blg", "*.fls",
		"*.fdb_latexmk", "*.synctex.gz", "*.out",
		"*.toc", "*.lot", "*.lof", "*.nav", "*.snm",
		"*.vrb", "*.idx", "*.ind", "*.ilg",
	}

	cleanedCount := 0

	for _, pattern := range tempPatterns {
		matches, err := filepath.Glob(filepath.Join(distDir, pattern))
		if err != nil {
			continue
		}

		for _, match := range matches {
			if err := os.Remove(match); err != nil {
				fmt.Printf("[WARN] Não foi possível remover %s: %v\n", match, err)
			} else {
				fmt.Printf("[REMOVED] %s\n", match)
				cleanedCount++
			}
		}
	}

	// Remover PDF se solicitado
	if cleanAll {
		pdfPath := filepath.Join(distDir, "main.pdf")
		if _, err := os.Stat(pdfPath); err == nil {
			if err := os.Remove(pdfPath); err != nil {
				fmt.Printf("[WARN] Não foi possível remover %s: %v\n", pdfPath, err)
			} else {
				fmt.Printf("[REMOVED] %s\n", pdfPath)
				cleanedCount++
			}
		}
	}

	if cleanedCount == 0 {
		fmt.Println("[INFO] Nenhum arquivo temporário encontrado")
	} else {
		fmt.Printf("[SUCCESS] %d arquivo(s) removido(s)\n", cleanedCount)
	}

	return nil
}

func openShell() error {
	fmt.Println(">> Abrindo shell do container...")
	fmt.Println("[INFO] Digite 'exit' para sair do container")

	// Verificar se container está rodando
	cmd := exec.Command("docker", "compose", "-f", "config/docker/docker-compose.yml", "ps", "-q", "latex-env")
	output, err := cmd.Output()
	if err != nil || len(output) == 0 {
		return fmt.Errorf("container não está rodando. Execute 'ltx build' primeiro")
	}

	// Abrir shell interativo
	cmd = exec.Command("docker", "compose", "-f", "config/docker/docker-compose.yml", "exec", "latex-env", "/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func showLogs() error {
	fmt.Println(">> Mostrando logs do container...")

	// Mostrar logs do container
	cmd := exec.Command("docker", "compose", "-f", "config/docker/docker-compose.yml", "logs", "--tail", "50", "latex-env")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
