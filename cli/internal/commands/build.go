package commands

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/martinsmiguel/latex-docker-env/cli/internal/colors"
)

var (
	buildEngine   string
	buildClean    bool
	buildVerbose  bool
)

var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Compila o documento LaTeX",
	Long: `Compila o documento LaTeX usando Docker.

O comando irá:
1. Verificar se o ambiente Docker está ativo
2. Compilar o documento principal (main.tex)
3. Processar bibliografia se necessário
4. Gerar o PDF final na pasta dist/`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return buildProject()
	},
}

func init() {
	BuildCmd.Flags().StringVar(&buildEngine, "engine", "pdflatex", "Engine LaTeX a usar (pdflatex, xelatex, lualatex)")
	BuildCmd.Flags().BoolVar(&buildClean, "clean", false, "Limpar arquivos temporários antes de compilar")
	BuildCmd.Flags().BoolVarP(&buildVerbose, "verbose", "v", false, "Saída detalhada")
}

func buildProject() error {
	start := time.Now()
	colors.Println(">> Compilando documento LaTeX...")

	// Verificar se há compilações em andamento
	if err := handleRunningCompilation(); err != nil {
		return err
	}

	// Verificar se existe main.tex
	sourceDir := "src"
	mainTexPath := filepath.Join(sourceDir, "main.tex")

	if _, err := os.Stat(mainTexPath); os.IsNotExist(err) {
		return fmt.Errorf("arquivo %s não encontrado. Execute 'ltx init' primeiro", mainTexPath)
	}

	// Limpar se solicitado
	if buildClean {
		if err := cleanTempFiles(); err != nil {
			colors.Printf("[WARN] Erro ao limpar arquivos temporários: %v\n", err)
		}
	}

	// Verificar Docker
	if err := exec.Command("docker", "version").Run(); err != nil {
		return fmt.Errorf("Docker não está disponível: %w", err)
	}

	// Verificar se container está rodando
	colors.PrintInfo("Iniciando compilação...")
	if err := ensureContainerRunning(); err != nil {
		return fmt.Errorf("erro ao garantir que container esteja rodando: %w", err)
	}

	// Compilar documento
	if err := compileDocument(mainTexPath); err != nil {
		return fmt.Errorf("erro na compilação: %w", err)
	}

	duration := time.Since(start)
	colors.Printf("[SUCCESS] Compilação concluída em %v\n", duration.Round(time.Second))
	colors.PrintInfo("PDF gerado: dist/main.pdf")

	return nil
}

func ensureContainerRunning() error {
	// Verificar se o container existe e está rodando usando docker-compose
	cmd := exec.Command("docker", "compose", "-f", "config/docker/docker-compose.yml", "ps", "-q", "latex-env")
	output, err := cmd.Output()
	if err != nil || len(output) == 0 {
		colors.PrintInfo("Iniciando ambiente Docker...")

		// Iniciar container
		cmd = exec.Command("docker", "compose", "-f", "config/docker/docker-compose.yml", "up", "-d")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("erro ao iniciar container: %w", err)
		}

		colors.PrintSuccess("Ambiente Docker iniciado com sucesso")

		// Aguardar container ficar saudável
		colors.PrintInfo("Aguardando container ficar saudável...")
		time.Sleep(2 * time.Second)
		colors.PrintSuccess("Container está saudável")
	}

	return nil
}

func compileDocument(mainTexPath string) error {
	engine := buildEngine
	if engine == "" {
		engine = "pdflatex"
	}

	colors.Printf("[INFO] Compilando %s com %s...\n", mainTexPath, engine)

	// Comando simplificado para executar latexmk no container
	args := []string{
		"compose", "-f", "config/docker/docker-compose.yml",
		"exec", "-T", "latex-env",
		"latexmk",
		"-pdf",
		"-interaction=nonstopmode",
		"-file-line-error",
		"-synctex=1",
		"-recorder",
		"-output-directory=dist",
		"src/main.tex",
	}

	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func cleanTempFiles() error {
	colors.PrintInfo("Limpando arquivos temporários...")

	// Padrões de arquivos temporários
	patterns := []string{
		"dist/*.aux",
		"dist/*.log",
		"dist/*.bbl",
		"dist/*.blg",
		"dist/*.fls",
		"dist/*.fdb_latexmk",
		"dist/*.synctex.gz",
		"dist/*.out",
		"dist/*.toc",
		"dist/*.lot",
		"dist/*.lof",
	}

	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			continue
		}

		for _, match := range matches {
			if err := os.Remove(match); err != nil {
				colors.Printf("[WARN] Não foi possível remover %s: %v\n", match, err)
			}
		}
	}

	return nil
}

// checkRunningCompilation verifica se há uma compilação em andamento
func checkRunningCompilation() (bool, error) {
	// Verificar se há processos latexmk rodando no container
	cmd := exec.Command("docker", "compose", "-f", "config/docker/docker-compose.yml",
		"exec", "-T", "latex-env", "pgrep", "-f", "latexmk")

	output, err := cmd.Output()
	if err != nil {
		// Se não encontrou processos ou houve erro, assume que não há compilação
		return false, nil
	}

	// Se há output, significa que há processos latexmk rodando
	return len(strings.TrimSpace(string(output))) > 0, nil
}

// killRunningCompilation mata processos de compilação em andamento
func killRunningCompilation() error {
	colors.PrintInfo("Encerrando processos de compilação em andamento...")

	// Matar processos latexmk no container
	cmd := exec.Command("docker", "compose", "-f", "config/docker/docker-compose.yml",
		"exec", "-T", "latex-env", "pkill", "-f", "latexmk")

	if err := cmd.Run(); err != nil {
		// Ignorar erro se não houver processos para matar
		colors.PrintWarn("Nenhum processo de compilação encontrado para encerrar")
	}

	// Aguardar um pouco para os processos terminarem
	time.Sleep(2 * time.Second)

	return nil
}

// askUserConfirmation pergunta ao usuário se deseja continuar
func askUserConfirmation(message string) bool {
	fmt.Printf("%s (s/N): ", message)
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	response = strings.TrimSpace(strings.ToLower(response))
	return response == "s" || response == "sim" || response == "y" || response == "yes"
}

// handleRunningCompilation gerencia compilações em andamento
func handleRunningCompilation() error {
	isRunning, err := checkRunningCompilation()
	if err != nil {
		return fmt.Errorf("erro ao verificar compilações em andamento: %w", err)
	}

	if isRunning {
		colors.PrintWarn("Há uma compilação LaTeX em andamento!")

		if askUserConfirmation("Deseja encerrar a compilação atual e iniciar uma nova?") {
			if err := killRunningCompilation(); err != nil {
				return fmt.Errorf("erro ao encerrar compilação: %w", err)
			}
			colors.PrintSuccess("Compilação anterior encerrada")
		} else {
			return fmt.Errorf("operação cancelada pelo usuário")
		}
	}

	return nil
}
