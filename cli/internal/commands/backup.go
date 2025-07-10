package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/martinsmiguel/latex-docker-env/cli/internal/colors"
)

var (
	backupName   string
	backupCustom string
)

var BackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Cria backup do trabalho atual",
	Long: `Cria um backup do seu trabalho LaTeX atual:

- Copia toda a pasta src/ (arquivos LaTeX)
- Copia PDFs da pasta dist/
- Salva em uma pasta um nível acima do repositório
- Nome automático com timestamp ou nome customizado

O backup é salvo em: ../latex-backups/[nome-do-backup]/

Exemplos:
  ltx backup                           # Backup com timestamp
  ltx backup --name "versao-final"     # Backup com nome específico
  ltx backup --custom "../meus-docs"   # Backup em local customizado`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return createBackup()
	},
}

func init() {
	BackupCmd.Flags().StringVarP(&backupName, "name", "n", "", "Nome do backup (padrão: timestamp)")
	BackupCmd.Flags().StringVar(&backupCustom, "custom", "", "Caminho customizado para o backup")
}

func createBackup() error {
	colors.Println(">> Criando backup do projeto...")

	// Verificar se existe conteúdo para backup
	if !hasContentToBackup() {
		colors.PrintWarning("⚠️  Nenhum conteúdo encontrado para backup")
		colors.PrintInfo("   Execute 'ltx init' para criar um projeto primeiro")
		return nil
	}

	// Determinar nome e caminho do backup
	backupPath, err := getBackupPath()
	if err != nil {
		return fmt.Errorf("erro ao determinar caminho do backup: %w", err)
	}

	// Criar diretório do backup
	if err := os.MkdirAll(backupPath, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório de backup: %w", err)
	}

	colors.Printf("📁 Salvando backup em: %s\n", backupPath)

	// Copiar conteúdo
	backupCount := 0

	// 1. Copiar pasta src/
	if _, err := os.Stat("src"); err == nil {
		srcBackupPath := filepath.Join(backupPath, "src")
		if err := copyDirectory("src", srcBackupPath); err != nil {
			return fmt.Errorf("erro ao copiar pasta src: %w", err)
		}
		colors.Printf("[COPIED] src/ → %s\n", srcBackupPath)
		backupCount++
	}

	// 2. Copiar PDFs da pasta dist/
	if _, err := os.Stat("dist"); err == nil {
		distBackupPath := filepath.Join(backupPath, "dist")
		if err := os.MkdirAll(distBackupPath, 0755); err != nil {
			return fmt.Errorf("erro ao criar pasta dist no backup: %w", err)
		}

		// Copiar apenas arquivos PDF
		pdfFiles, err := filepath.Glob("dist/*.pdf")
		if err == nil && len(pdfFiles) > 0 {
			for _, pdfFile := range pdfFiles {
				fileName := filepath.Base(pdfFile)
				destPath := filepath.Join(distBackupPath, fileName)
				if err := copyFile(pdfFile, destPath); err != nil {
					colors.PrintWarning(fmt.Sprintf("Aviso: não foi possível copiar %s", fileName))
				} else {
					colors.Printf("[COPIED] %s → %s\n", pdfFile, destPath)
					backupCount++
				}
			}
		}
	}

	// 3. Criar arquivo de informações do backup
	if err := createBackupInfo(backupPath); err != nil {
		colors.PrintWarning(fmt.Sprintf("Aviso: não foi possível criar arquivo de informações: %v", err))
	}

	if backupCount == 0 {
		colors.PrintWarning("⚠️  Nenhum arquivo foi copiado para o backup")
	} else {
		colors.PrintSuccess(fmt.Sprintf("✅ Backup criado com sucesso! (%d item(s) copiado(s))", backupCount))
		colors.PrintInfo(fmt.Sprintf("📂 Localização: %s", backupPath))
	}

	return nil
}

func hasContentToBackup() bool {
	// Verificar se existe pasta src com conteúdo
	if srcInfo, err := os.Stat("src"); err == nil && srcInfo.IsDir() {
		return true
	}

	// Verificar se existe pasta dist com PDFs
	if pdfFiles, err := filepath.Glob("dist/*.pdf"); err == nil && len(pdfFiles) > 0 {
		return true
	}

	return false
}

func getBackupPath() (string, error) {
	var basePath string
	var backupDirName string

	// Determinar caminho base
	if backupCustom != "" {
		basePath = backupCustom
	} else {
		// Um nível acima do repositório atual
		currentDir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		basePath = filepath.Join(filepath.Dir(currentDir), "latex-backups")
	}

	// Determinar nome do backup
	if backupName != "" {
		backupDirName = backupName
	} else {
		// Nome com timestamp
		timestamp := time.Now().Format("2006-01-02_15-04-05")
		projectName := filepath.Base(getCurrentDir())
		backupDirName = fmt.Sprintf("%s_%s", projectName, timestamp)
	}

	return filepath.Join(basePath, backupDirName), nil
}

func getCurrentDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		return "latex-project"
	}
	return filepath.Base(currentDir)
}

func copyDirectory(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calcular caminho de destino
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		return copyFile(path, dstPath)
	})
}

func copyFile(src, dst string) error {
	// Criar diretório pai se necessário
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	// Abrir arquivo fonte
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := srcFile.Close(); closeErr != nil {
			fmt.Printf("Erro ao fechar arquivo origem: %v\n", closeErr)
		}
	}()

	// Criar arquivo destino
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := dstFile.Close(); closeErr != nil {
			fmt.Printf("Erro ao fechar arquivo destino: %v\n", closeErr)
		}
	}()

	// Copiar conteúdo
	_, err = dstFile.ReadFrom(srcFile)
	if err != nil {
		return err
	}

	// Copiar permissões
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(dst, srcInfo.Mode())
}

func createBackupInfo(backupPath string) error {
	infoPath := filepath.Join(backupPath, "backup-info.txt")

	content := fmt.Sprintf(`Backup do LaTeX Docker Environment
=====================================

Data/Hora: %s
Projeto: %s
Diretório Original: %s

Conteúdo do Backup:
- src/: Arquivos LaTeX do projeto
- dist/: PDFs compilados

Para restaurar:
1. Copie o conteúdo de src/ de volta para o projeto
2. Execute 'ltx build' para recompilar

Criado por: ltx backup
`,
		time.Now().Format("2006-01-02 15:04:05"),
		getCurrentDir(),
		getCurrentWorkingDir(),
	)

	return os.WriteFile(infoPath, []byte(content), 0644)
}

func getCurrentWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return "desconhecido"
	}
	return dir
}
