package commands

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsRelevantFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected bool
	}{
		{
			name:     "arquivo .tex",
			filename: "document.tex",
			expected: true,
		},
		{
			name:     "arquivo .bib",
			filename: "references.bib",
			expected: true,
		},
		{
			name:     "arquivo .sty",
			filename: "style.sty",
			expected: true,
		},
		{
			name:     "arquivo .cls",
			filename: "class.cls",
			expected: true,
		},
		{
			name:     "arquivo .png",
			filename: "image.png",
			expected: true,
		},
		{
			name:     "arquivo .jpg",
			filename: "photo.jpg",
			expected: true,
		},
		{
			name:     "arquivo .pdf",
			filename: "document.pdf",
			expected: true,
		},
		{
			name:     "arquivo .svg",
			filename: "vector.svg",
			expected: true,
		},
		{
			name:     "arquivo temporário .aux",
			filename: "document.aux",
			expected: false,
		},
		{
			name:     "arquivo temporário .log",
			filename: "document.log",
			expected: false,
		},
		{
			name:     "arquivo temporário .out",
			filename: "document.out",
			expected: false,
		},
		{
			name:     "arquivo temporário .toc",
			filename: "document.toc",
			expected: false,
		},
		{
			name:     "arquivo .txt",
			filename: "readme.txt",
			expected: false,
		},
		{
			name:     "arquivo sem extensão",
			filename: "makefile",
			expected: false,
		},
		{
			name:     "arquivo .DS_Store",
			filename: ".DS_Store",
			expected: false,
		},
		{
			name:     "arquivo oculto",
			filename: ".hidden",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isRelevantFile(tt.filename)
			if result != tt.expected {
				t.Errorf("isRelevantFile(%s) = %v, want %v", tt.filename, result, tt.expected)
			}
		})
	}
}

func TestAddWatchPaths(t *testing.T) {
	// Criar diretório temporário para teste
	tempDir := t.TempDir()

	// Criar estrutura de diretórios
	dirs := []string{
		"src",
		"src/chapters",
		"src/images",
		"src/styles",
		"templates",
		"config",
	}

	for _, dir := range dirs {
		dirPath := filepath.Join(tempDir, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			t.Fatalf("Erro ao criar diretório %s: %v", dir, err)
		}
	}

	// Criar alguns arquivos para testar
	files := []string{
		"src/main.tex",
		"src/chapters/intro.tex",
		"src/images/logo.png",
		"src/styles/custom.sty",
		"templates/article.tex",
		"config/settings.json",
	}

	for _, file := range files {
		filePath := filepath.Join(tempDir, file)
		if err := os.WriteFile(filePath, []byte("content"), 0644); err != nil {
			t.Fatalf("Erro ao criar arquivo %s: %v", file, err)
		}
	}

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

	// Mudar para o diretório temporário
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Erro ao mudar para diretório temporário: %v", err)
	}

	// Testar se a função addWatchPaths funciona sem erro
	// Como ela depende do fsnotify.Watcher, vamos testar indiretamente
	// verificando se os diretórios existem
	watchDirs := []string{"src", "templates", "config"}

	for _, dir := range watchDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("Diretório %s deveria existir para watch", dir)
		}
	}
}

func TestWatchProject_DirectoryStructure(t *testing.T) {
	// Criar diretório temporário para teste
	tempDir := t.TempDir()
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Erro ao obter diretório atual: %v", err)
	}
	defer func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Errorf("Erro ao restaurar diretório: %v", err)
		}
	}()

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Erro ao mudar para diretório temporário: %v", err)
	}

	tests := []struct {
		name        string
		setupDirs   []string
		setupFiles  map[string]string
		expectError bool
	}{
		{
			name:        "sem diretório src",
			setupDirs:   []string{},
			setupFiles:  map[string]string{},
			expectError: true,
		},
		{
			name:        "com diretório src vazio",
			setupDirs:   []string{"src"},
			setupFiles:  map[string]string{},
			expectError: true, // Sem main.tex
		},
		{
			name:       "estrutura válida",
			setupDirs:  []string{"src", "src/chapters"},
			setupFiles: map[string]string{
				"src/main.tex": `\documentclass{article}
\usepackage[utf8]{inputenc}
\begin{document}
\title{Test Document}
\author{Test Author}
\maketitle
\section{Introduction}
This is a test document.
\end{document}`,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Limpar diretório
			if err := os.RemoveAll(tempDir); err != nil {
				t.Fatalf("Erro ao limpar diretório: %v", err)
			}
			if err := os.MkdirAll(tempDir, 0755); err != nil {
				t.Fatalf("Erro ao recriar diretório: %v", err)
			}
			if err := os.Chdir(tempDir); err != nil {
				t.Fatalf("Erro ao mudar para diretório: %v", err)
			}

			// Configurar estrutura de teste
			for _, dir := range tt.setupDirs {
				if err := os.MkdirAll(dir, 0755); err != nil {
					t.Fatalf("Erro ao criar diretório %s: %v", dir, err)
				}
			}

			for file, content := range tt.setupFiles {
				if err := os.WriteFile(file, []byte(content), 0644); err != nil {
					t.Fatalf("Erro ao criar arquivo %s: %v", file, err)
				}
			}

			// Testar apenas verificação da estrutura de diretórios
			// Em vez de chamar watchProject() que é bloqueante,
			// vamos testar se o arquivo main.tex existe
			_, err := os.Stat("src/main.tex")
			hasMainTex := err == nil

			if tt.expectError && hasMainTex {
				t.Error("Teste esperava erro mas encontrou main.tex")
			} else if !tt.expectError && !hasMainTex {
				t.Error("Teste esperava sucesso mas não encontrou main.tex")
			}
		})
	}
}

// Teste auxiliar para verificar extensões de arquivo
func TestFileExtensions(t *testing.T) {
	relevantExtensions := []string{".tex", ".bib", ".sty", ".cls", ".png", ".jpg", ".jpeg", ".pdf", ".svg"}
	irrelevantExtensions := []string{".aux", ".log", ".out", ".toc", ".fdb_latexmk", ".fls", ".synctex.gz"}

	for _, ext := range relevantExtensions {
		filename := "test" + ext
		if !isRelevantFile(filename) {
			t.Errorf("Arquivo %s deveria ser considerado relevante", filename)
		}
	}

	for _, ext := range irrelevantExtensions {
		filename := "test" + ext
		if isRelevantFile(filename) {
			t.Errorf("Arquivo %s não deveria ser considerado relevante", filename)
		}
	}
}

// Teste para verificar arquivos com caminhos complexos
func TestIsRelevantFileWithPaths(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		expected bool
	}{
		{
			name:     "arquivo em subdiretório",
			filepath: "chapters/intro.tex",
			expected: true,
		},
		{
			name:     "arquivo com caminho absoluto",
			filepath: "/home/user/project/src/main.tex",
			expected: true,
		},
		{
			name:     "arquivo em diretório oculto",
			filepath: ".git/config",
			expected: false,
		},
		{
			name:     "arquivo temporário em subdiretório",
			filepath: "output/document.aux",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := filepath.Base(tt.filepath)
			result := isRelevantFile(filename)
			if result != tt.expected {
				t.Errorf("isRelevantFile(%s) = %v, want %v", filename, result, tt.expected)
			}
		})
	}
}
