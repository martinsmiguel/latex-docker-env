package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/martinsmiguel/latex-docker-env/cli/internal/colors"
	"github.com/martinsmiguel/latex-docker-env/cli/internal/template"
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
	colors.Println(">> Limpando arquivos temporários...")

	distDir := "dist"
	if _, err := os.Stat(distDir); os.IsNotExist(err) {
		colors.PrintInfo("Diretório dist não existe, nada para limpar")
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
				colors.Printf("[WARN] Não foi possível remover %s: %v\n", match, err)
			} else {
				colors.Printf("[REMOVED] %s\n", match)
				cleanedCount++
			}
		}
	}

	// Remover PDF se solicitado
	if cleanAll {
		pdfPath := filepath.Join(distDir, "main.pdf")
		if _, err := os.Stat(pdfPath); err == nil {
			if err := os.Remove(pdfPath); err != nil {
				colors.Printf("[WARN] Não foi possível remover %s: %v\n", pdfPath, err)
			} else {
				colors.Printf("[REMOVED] %s\n", pdfPath)
				cleanedCount++
			}
		}
	}

	if cleanedCount == 0 {
		colors.PrintInfo("Nenhum arquivo temporário encontrado")
	} else {
		colors.Printf("[SUCCESS] %d arquivo(s) removido(s)\n", cleanedCount)
	}

	return nil
}

func openShell() error {
	colors.Println(">> Abrindo shell do container...")
	colors.PrintInfo("Digite 'exit' para sair do container")

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
	colors.Println(">> Mostrando logs do container...")

	// Mostrar logs do container
	cmd := exec.Command("docker", "compose", "-f", "config/docker/docker-compose.yml", "logs", "--tail", "50", "latex-env")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// Função utilitária para criar registry de templates
func getTemplateRegistry() *template.Registry {
	registry := template.NewRegistry()

	// Usar sempre caminhos absolutos fixos para debug
	projectRoot := "/Users/miguelmartins/workspace/labs/latex-docker-env"

	registry.AddTemplatePath(filepath.Join(projectRoot, "cli/templates"))
	registry.AddTemplatePath(filepath.Join(projectRoot, "templates"))
	registry.AddTemplatePath(filepath.Join(projectRoot, "user-templates"))

	return registry
}

// Encontra o diretório raiz do projeto latex-docker-env
func findProjectRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	// Procurar pelo arquivo go.mod ou qualquer indicador do projeto
	for dir := wd; dir != "/" && dir != ""; dir = filepath.Dir(dir) {
		// Verificar se existe o arquivo go.mod no subdiretório cli/
		cliPath := filepath.Join(dir, "cli", "go.mod")
		if _, err := os.Stat(cliPath); err == nil {
			return dir
		}

		// Verificar se existe config/latex-cli.conf
		configPath := filepath.Join(dir, "config", "latex-cli.conf")
		if _, err := os.Stat(configPath); err == nil {
			return dir
		}
	}

	return ""
}
