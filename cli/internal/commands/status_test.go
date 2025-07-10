package commands

import (
	"os"
	"path/filepath"
	"testing"
)

func TestShowStatus(t *testing.T) {
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

	// Criar estrutura básica de projeto
	dirs := []string{"src", "dist", "tmp", "src/chapters", "src/images"}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Erro ao criar diretório %s: %v", dir, err)
		}
	}

	// Criar arquivos de teste
	files := map[string]string{
		"src/main.tex":              "\\documentclass{article}\\begin{document}Test\\end{document}",
		"src/chapters/intro.tex":    "\\chapter{Introduction}",
		"src/chapters/chapter1.tex": "\\chapter{Chapter 1}",
		"src/references.bib":        "@article{test, title={Test}}",
		"src/images/logo.png":       "fake png content",
	}

	for file, content := range files {
		if err := os.WriteFile(file, []byte(content), 0644); err != nil {
			t.Fatalf("Erro ao criar arquivo %s: %v", file, err)
		}
	}

	// Testar função showStatus - não deve falhar mesmo sem Docker
	err = showStatus()
	if err != nil {
		t.Errorf("showStatus() error = %v", err)
	}
}

func TestExtractFromLatex(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		command  string
		expected string
	}{
		{
			name:     "extrair título",
			content:  "\\title{My Great Title}\\begin{document}",
			command:  "\\title",
			expected: "My Great Title",
		},
		{
			name:     "extrair autor",
			content:  "\\author{John Doe}\\begin{document}",
			command:  "\\author",
			expected: "John Doe",
		},
		{
			name:     "sem título",
			content:  "\\begin{document}Content\\end{document}",
			command:  "\\title",
			expected: "",
		},
		{
			name:     "título com chaves extras",
			content:  "\\title{My {Complex} Title}\\begin{document}",
			command:  "\\title",
			expected: "My {Complex} Title",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractFromLatex(tt.content, tt.command)
			if result != tt.expected {
				t.Errorf("extractFromLatex() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCountLatexFiles(t *testing.T) {
	// Criar diretório temporário
	tempDir := t.TempDir()

	// Criar estrutura de arquivos
	files := []string{
		"main.tex",
		"chapter1.tex",
		"chapter2.tex",
		"styles.sty",
		"references.bib",
		"image.png",
		"document.pdf",
	}

	for _, file := range files {
		filePath := filepath.Join(tempDir, file)
		if err := os.WriteFile(filePath, []byte("content"), 0644); err != nil {
			t.Fatalf("Erro ao criar arquivo %s: %v", file, err)
		}
	}

	count := countLatexFiles(tempDir)
	expected := 3 // main.tex, chapter1.tex, chapter2.tex
	if count != expected {
		t.Errorf("countLatexFiles() = %d, want %d", count, expected)
	}
}

func TestCountFiles(t *testing.T) {
	// Criar diretório temporário
	tempDir := t.TempDir()

	// Criar arquivos de teste
	files := []string{"file1.txt", "file2.png", "file3.pdf"}
	for _, file := range files {
		filePath := filepath.Join(tempDir, file)
		if err := os.WriteFile(filePath, []byte("content"), 0644); err != nil {
			t.Fatalf("Erro ao criar arquivo %s: %v", file, err)
		}
	}

	count := countFiles(tempDir, ".txt")
	expected := 1 // file1.txt
	if count != expected {
		t.Errorf("countFiles() = %d, want %d", count, expected)
	}

	// Testar com outra extensão
	count = countFiles(tempDir, ".png")
	expected = 1 // file2.png
	if count != expected {
		t.Errorf("countFiles() = %d, want %d", count, expected)
	}

	// Testar contagem total (simulando uma extensão que não existe)
	count = countFiles(tempDir, ".xyz")
	expected = 0
	if count != expected {
		t.Errorf("countFiles() = %d, want %d", count, expected)
	}
}

func TestCountBibEntries(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected int
	}{
		{
			name:     "arquivo bib vazio",
			content:  "",
			expected: 0,
		},
		{
			name: "uma entrada",
			content: `@article{key1,
				title={Test Title},
				author={Test Author}
			}`,
			expected: 1,
		},
		{
			name: "múltiplas entradas",
			content: `@article{key1,
				title={Test Title 1}
			}
			@book{key2,
				title={Test Book}
			}
			@inproceedings{key3,
				title={Test Proceedings}
			}`,
			expected: 3,
		},
		{
			name: "entradas com comentários",
			content: `% This is a comment
			@article{key1,
				title={Test Title}
			}
			% Another comment
			@book{key2,
				title={Test Book}
			}`,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Criar arquivo temporário
			tempFile := filepath.Join(t.TempDir(), "test.bib")
			if err := os.WriteFile(tempFile, []byte(tt.content), 0644); err != nil {
				t.Fatalf("Erro ao criar arquivo de teste: %v", err)
			}

			count := countBibEntries(tempFile)
			if count != tt.expected {
				t.Errorf("countBibEntries() = %d, want %d", count, tt.expected)
			}
		})
	}
}
