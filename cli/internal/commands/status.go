package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/martinsmiguel/latex-docker-env/cli/internal/colors"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Mostra o status do ambiente",
	Long: `Exibe informações detalhadas sobre o estado atual do projeto:

- Status da CLI e configurações
- Status do Docker e containers
- Informações do projeto LaTeX
- Status da última compilação`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return showStatus()
	},
}

func showStatus() error {
	fmt.Println("=== Status do LaTeX Docker Environment ===")
	fmt.Println()

	// Status da CLI
	showCLIStatus()
	fmt.Println()

	// Status do Docker
	if err := showDockerStatus(); err != nil {
		colors.Printf("[ERROR] Erro ao verificar Docker: %v\n", err)
	}
	fmt.Println()

	// Status do Projeto
	showProjectStatus()

	return nil
}

func showCLIStatus() {
	fmt.Println("=== LaTeX CLI ===")
	fmt.Println("Versão: 2.0.0")

	workDir, _ := os.Getwd()
	fmt.Printf("Diretório do projeto: %s\n", workDir)

	configPath := "config/latex-cli.conf"
	if _, err := os.Stat(configPath); err == nil {
		fmt.Printf("Arquivo de configuração: %s\n", configPath)
	} else {
		fmt.Printf("Arquivo de configuração: não encontrado\n")
	}

	fmt.Println("Configurações:")
	fmt.Println("  Engine LaTeX: pdflatex")
	fmt.Println("  Diretório fonte: src")
	fmt.Println("  Diretório de saída: dist")
	fmt.Println("  Container: latex-env")
}

func showDockerStatus() error {
	fmt.Println("=== Status do Docker ===")

	// Verificar se Docker está disponível
	if err := exec.Command("docker", "version").Run(); err != nil {
		fmt.Println("✗ Docker não está disponível")
		return err
	}
	fmt.Println("✓ Docker está disponível")

	// Obter versão do Docker
	cmd := exec.Command("docker", "version", "--format", "{{.Server.Version}}")
	if output, err := cmd.Output(); err == nil {
		fmt.Printf("Versão: %s", string(output))
	}

	// Verificar Docker Compose
	if err := exec.Command("docker", "compose", "version").Run(); err != nil {
		fmt.Println("✗ Docker Compose não está disponível")
	} else {
		fmt.Println("✓ Docker Compose está disponível")
	}

	// Verificar container
	return checkContainerStatus()
}

func checkContainerStatus() error {
	// Verificar se container está rodando
	cmd := exec.Command("docker", "compose", "-f", "config/docker/docker-compose.yml", "ps", "-q", "latex-env")
	output, err := cmd.Output()

	if err != nil || len(strings.TrimSpace(string(output))) == 0 {
		fmt.Println("✗ Container latex-env não está executando")
		return nil
	}

	fmt.Println("✓ Container latex-env está executando")

	// Verificar saúde do container
	containerID := strings.TrimSpace(string(output))
	cmd = exec.Command("docker", "inspect", "--format", "{{.State.Health.Status}}", containerID)
	if healthOutput, err := cmd.Output(); err == nil {
		health := strings.TrimSpace(string(healthOutput))
		if health == "healthy" {
			fmt.Println("✓ Container está saudável")
		} else {
			fmt.Printf("⚠ Container health: %s\n", health)
		}
	}

	return nil
}

func showProjectStatus() {
	fmt.Println("=== Status do Projeto ===")

	sourceDir := "src"
	mainTexPath := filepath.Join(sourceDir, "main.tex")

	// Verificar se projeto está inicializado
	if _, err := os.Stat(mainTexPath); os.IsNotExist(err) {
		fmt.Println("✗ Projeto não inicializado")
		fmt.Println("  Execute 'ltx init' para começar")
		return
	}

	fmt.Println("✓ Projeto inicializado")

	// Tentar extrair título e autor do main.tex
	if content, err := os.ReadFile(mainTexPath); err == nil {
		contentStr := string(content)

		if title := extractFromLatex(contentStr, "\\title{"); title != "" {
			fmt.Printf("  Título: %s\n", title)
		} else {
			fmt.Println("  Título: Não encontrado")
		}

		if author := extractFromLatex(contentStr, "\\author{"); author != "" {
			fmt.Printf("  Autor: %s\n", author)
		} else {
			fmt.Println("  Autor: Não encontrado")
		}
	}

	// Contar arquivos LaTeX
	latexFiles := countLatexFiles(sourceDir)
	fmt.Printf("  Arquivos LaTeX: %d\n", latexFiles)

	// Verificar PDF
	pdfPath := "dist/main.pdf"
	if stat, err := os.Stat(pdfPath); err == nil {
		fmt.Printf("✓ PDF disponível (compilado em: %s)\n", stat.ModTime().Format("2006-01-02 15:04:05"))

		// Contar capítulos (arquivos .tex em chapters/)
		chaptersDir := filepath.Join(sourceDir, "chapters")
		if chapters := countFiles(chaptersDir, ".tex"); chapters > 0 {
			fmt.Printf("  Capítulos: %d\n", chapters)
		}

		// Contar referências
		referencesPath := filepath.Join(sourceDir, "references.bib")
		if refs := countBibEntries(referencesPath); refs > 0 {
			fmt.Printf("  Referências bibliográficas: %d\n", refs)
		}
	} else {
		fmt.Println("✗ PDF não encontrado")
		fmt.Println("  Execute 'ltx build' para compilar")
	}
}

func extractFromLatex(content, command string) string {
	start := strings.Index(content, command)
	if start == -1 {
		return ""
	}

	start += len(command)
	depth := 1

	for i := start; i < len(content) && depth > 0; i++ {
		switch content[i] {
		case '{':
			depth++
		case '}':
			depth--
		}

		if depth == 0 {
			result := content[start:i]
			// Limpar formatação básica
			result = strings.ReplaceAll(result, "\\textbf{", "")
			result = strings.ReplaceAll(result, "}", "")
			return strings.TrimSpace(result)
		}
	}

	return ""
}

func countLatexFiles(dir string) int {
	count := 0
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() && filepath.Ext(path) == ".tex" {
			count++
		}
		return nil
	})
	return count
}

func countFiles(dir, ext string) int {
	count := 0
	if entries, err := os.ReadDir(dir); err == nil {
		for _, entry := range entries {
			if !entry.IsDir() && filepath.Ext(entry.Name()) == ext {
				count++
			}
		}
	}
	return count
}

func countBibEntries(bibPath string) int {
	content, err := os.ReadFile(bibPath)
	if err != nil {
		return 0
	}

	return strings.Count(string(content), "@")
}
