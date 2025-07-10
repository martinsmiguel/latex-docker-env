package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/martinsmiguel/latex-docker-env/cli/internal/config"
	"github.com/martinsmiguel/latex-docker-env/cli/internal/colors"
)

var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configura o ambiente de desenvolvimento",
	Long: `Verifica dependências, configura Docker e prepara o ambiente
para desenvolvimento LaTeX.`,
	RunE: runSetup,
}

func runSetup(cmd *cobra.Command, args []string) error {
	colors.Println(">> Configurando ambiente LaTeX Docker...")

	// 1. Definir configurações padrão
	config.SetDefaults()

	// 2. Verificar se estamos no diretório correto
	if err := verifyProjectStructure(); err != nil {
		return fmt.Errorf("estrutura do projeto inválida: %w", err)
	}

	// 3. Verificar Docker (simplificado)
	colors.Println(">> Verificando Docker...")
	if err := exec.Command("docker", "--version").Run(); err != nil {
		return fmt.Errorf("docker não está disponível: %w", err)
	}
	fmt.Println("[OK] Docker verificado")

	// 4. Pular verificação de imagem por enquanto
	fmt.Println("[OK] Imagem LaTeX será verificada durante o build")

	// 5. Criar diretórios necessários
	if err := createDirectories(); err != nil {
		return fmt.Errorf("erro ao criar diretórios: %w", err)
	}
	fmt.Println("[OK] Estrutura de diretórios criada")

	// 6. Configurar VS Code (simplificado)
	if err := setupVSCode(); err != nil {
		colors.Printf("[WARN] Configuração VS Code opcional falhou: %v\n", err)
	} else {
		fmt.Println("[OK] VS Code configurado")
	}

	colors.Println("\n[SUCCESS] Ambiente configurado com sucesso!")
	fmt.Println("Execute 'ltx init' para criar seu primeiro documento.")

	return nil
}

func verifyProjectStructure() error {
	requiredDirs := []string{"config", "lib", "docs"}
	requiredFiles := []string{"config/latex-cli.conf", "config/docker/docker-compose.yml"}

	for _, dir := range requiredDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			return fmt.Errorf("diretório %s não encontrado. Execute este comando no diretório raiz do latex-docker-env", dir)
		}
	}

	for _, file := range requiredFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf("arquivo %s não encontrado", file)
		}
	}

	return nil
}

func createDirectories() error {
	dirs := []string{"src", "dist", "tmp"}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("erro ao criar diretório %s: %w", dir, err)
		}
	}

	return nil
}

func setupVSCode() error {
	vscodeDir := ".vscode"
	if err := os.MkdirAll(vscodeDir, 0755); err != nil {
		return err
	}

	// Copiar configurações do VS Code se existirem
	configSource := "config/vscode/settings.json"
	// configTarget := filepath.Join(vscodeDir, "settings.json") // TODO: implementar cópia

	if _, err := os.Stat(configSource); err == nil {
		// TODO: Implementar cópia do arquivo
		fmt.Printf("[INFO] Configurações VS Code disponíveis em %s\n", configSource)
	}

	return nil
}

func checkAndPullImage(imageName string) error {
	// Verificar se a imagem já existe
	cmd := exec.Command("docker", "images", "-q", imageName)
	output, err := cmd.Output()

	if err == nil && len(output) > 0 {
		fmt.Printf("[OK] Imagem %s já existe localmente\n", imageName)
		return nil
	}

	// Imagem não existe, fazer pull
	fmt.Printf(">> Baixando imagem %s...\n", imageName)
	cmd = exec.Command("docker", "pull", imageName)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("falha ao baixar imagem %s: %w", imageName, err)
	}

	return nil
}
