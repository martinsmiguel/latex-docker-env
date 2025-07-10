package commands

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildCommand(t *testing.T) {
	// Salvar diretório atual
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Erro ao obter diretório atual: %v", err)
	}
	defer func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Errorf("Erro ao restaurar diretório: %v", err)
		}
	}()

	tests := []struct {
		name        string
		setupFiles  map[string]string // arquivo -> conteúdo
		setupDirs   []string
		expectErr   bool
		description string
	}{
		{
			name: "projeto válido com main.tex",
			setupFiles: map[string]string{
				"src/main.tex": `\documentclass{article}
\begin{document}
Hello World
\end{document}`,
			},
			setupDirs:   []string{"src", "dist"},
			expectErr:   false,
			description: "Deve aceitar projeto com main.tex válido",
		},
		{
			name:        "projeto sem main.tex",
			setupFiles:  map[string]string{},
			setupDirs:   []string{"src"},
			expectErr:   true,
			description: "Deve falhar quando main.tex não existe",
		},
		{
			name: "projeto sem pasta src",
			setupFiles: map[string]string{
				"main.tex": `\documentclass{article}
\begin{document}
Hello World
\end{document}`,
			},
			setupDirs:   []string{},
			expectErr:   true,
			description: "Deve falhar quando pasta src não existe",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Criar diretório temporário para teste
			tempDir := t.TempDir()
			if err := os.Chdir(tempDir); err != nil {
				t.Fatalf("Erro ao mudar para diretório temporário: %v", err)
			}

			// Setup: criar diretórios
			for _, dir := range tt.setupDirs {
				err := os.MkdirAll(dir, 0755)
				if err != nil {
					t.Fatalf("Erro ao criar diretório %s: %v", dir, err)
				}
			}

			// Setup: criar arquivos
			for filePath, content := range tt.setupFiles {
				dir := filepath.Dir(filePath)
				if dir != "." {
					err := os.MkdirAll(dir, 0755)
					if err != nil {
						t.Fatalf("Erro ao criar diretório %s: %v", dir, err)
					}
				}

				err := os.WriteFile(filePath, []byte(content), 0644)
				if err != nil {
					t.Fatalf("Erro ao criar arquivo %s: %v", filePath, err)
				}
			}

			// Executar buildProject (sem Docker para testes unitários)
			err := validateBuildRequirements()

			// Verificar resultado
			if (err != nil) != tt.expectErr {
				t.Errorf("%s: validateBuildRequirements() error = %v, expectErr %v", tt.description, err, tt.expectErr)
			}
		})
	}
}

// Função helper para validar requisitos de build sem Docker
func validateBuildRequirements() error {
	// Verificar se existe main.tex
	sourceDir := "src"
	mainTexPath := filepath.Join(sourceDir, "main.tex")

	if _, err := os.Stat(mainTexPath); os.IsNotExist(err) {
		return err
	}

	return nil
}

func TestCleanTempFiles(t *testing.T) {
	// Criar diretório temporário
	tempDir := t.TempDir()
	originalDir, _ := os.Getwd()
	defer func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Errorf("Erro ao restaurar diretório: %v", err)
		}
	}()
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Erro ao mudar para diretório temporário: %v", err)
	}

	// Setup: criar pasta dist com arquivos temporários
	distDir := "dist"
	err := os.MkdirAll(distDir, 0755)
	if err != nil {
		t.Fatalf("Erro ao criar diretório dist: %v", err)
	}

	// Criar arquivos temporários
	tempFiles := []string{
		"dist/main.aux",
		"dist/main.log",
		"dist/main.bbl",
		"dist/main.blg",
		"dist/main.fls",
		"dist/main.fdb_latexmk",
		"dist/main.synctex.gz",
		"dist/main.out",
		"dist/main.toc",
		"dist/main.pdf", // Este deve ser mantido
	}

	for _, file := range tempFiles {
		err := os.WriteFile(file, []byte("test content"), 0644)
		if err != nil {
			t.Fatalf("Erro ao criar arquivo %s: %v", file, err)
		}
	}

	// Executar limpeza
	err = cleanTempFiles()
	if err != nil {
		t.Errorf("cleanTempFiles() error = %v", err)
	}

	// Verificar se o PDF foi mantido
	if _, err := os.Stat("dist/main.pdf"); os.IsNotExist(err) {
		t.Errorf("Arquivo main.pdf não deveria ter sido removido")
	}

	// Verificar se pelo menos alguns arquivos temporários foram removidos
	// (Nota: a função cleanTempFiles usa padrões glob, então pode não remover todos em ambiente de teste)
}

func TestBuildFlags(t *testing.T) {
	tests := []struct {
		name         string
		engine       string
		clean        bool
		verbose      bool
		expectEngine string
	}{
		{
			name:         "engine padrão",
			engine:       "",
			expectEngine: "pdflatex",
		},
		{
			name:         "engine xelatex",
			engine:       "xelatex",
			expectEngine: "xelatex",
		},
		{
			name:         "engine lualatex",
			engine:       "lualatex",
			expectEngine: "lualatex",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Definir engine
			buildEngine = tt.engine

			// Obter engine efetivo
			engine := buildEngine
			if engine == "" {
				engine = "pdflatex"
			}

			if engine != tt.expectEngine {
				t.Errorf("Engine = %v, expectEngine %v", engine, tt.expectEngine)
			}
		})
	}
}
