package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/martinsmiguel/latex-docker-env/cli/pkg/types"
)

func TestNormalizeLaTeXPaths(t *testing.T) {
	registry := NewRegistry()
	loader := NewLoader(registry)

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "normalizar usepackage options",
			input:    "\\usepackage{misc/options}",
			expected: "\\usepackage{options}",
		},
		{
			name:     "normalizar usepackage styles/options",
			input:    "\\usepackage{styles/options}",
			expected: "\\usepackage{options}",
		},
		{
			name:     "normalizar input frontmatter",
			input:    "\\input{frontmatter/titlepage}",
			expected: "\\input{chapters/titlepage}",
		},
		{
			name:     "normalizar input content",
			input:    "\\input{content/chapter1}",
			expected: "\\input{chapters/chapter1}",
		},
		{
			name:     "normalizar includegraphics frontmatter",
			input:    "\\includegraphics{frontmatter/logo.png}",
			expected: "\\includegraphics{images/logo.png}",
		},
		{
			name:     "normalizar includegraphics com parâmetros",
			input:    "\\includegraphics[width=0.5\\textwidth]{frontmatter/logo.png}",
			expected: "\\includegraphics[width=0.5\\textwidth]{images/logo.png}",
		},
		{
			name:     "normalizar múltiplas referências",
			input: `\usepackage{misc/options}
\input{frontmatter/titlepage}
\includegraphics{frontmatter/logo.png}`,
			expected: `\usepackage{options}
\input{chapters/titlepage}
\includegraphics{images/logo.png}`,
		},
		{
			name:     "remover caminhos relativos",
			input:    "\\input{../chapters/test} \\input{./test}",
			expected: "\\input{chapters/test} \\input{test}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := loader.normalizeLaTeXPaths(tt.input)
			if result != tt.expected {
				t.Errorf("normalizeLaTeXPaths() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestMapDestination(t *testing.T) {
	registry := NewRegistry()
	loader := NewLoader(registry)

	tests := []struct {
		name       string
		sourcePath string
		targetDir  string
		expected   string
	}{
		{
			name:       "arquivo main.tex",
			sourcePath: "main.tex",
			targetDir:  "src",
			expected:   "main.tex",
		},
		{
			name:       "arquivo template.tex",
			sourcePath: "template.tex",
			targetDir:  "src",
			expected:   "main.tex",
		},
		{
			name:       "arquivo .tex em subdiretório",
			sourcePath: "chapters/chapter1.tex",
			targetDir:  "src",
			expected:   "chapters/chapter1.tex",
		},
		{
			name:       "arquivo .sty",
			sourcePath: "options.sty",
			targetDir:  "src",
			expected:   "styles/options.sty",
		},
		{
			name:       "arquivo .cls",
			sourcePath: "book.cls",
			targetDir:  "src",
			expected:   "styles/book.cls",
		},
		{
			name:       "arquivo .bib",
			sourcePath: "bibliography.bib",
			targetDir:  "src",
			expected:   "references.bib",
		},
		{
			name:       "arquivo de imagem .png",
			sourcePath: "logo.png",
			targetDir:  "src",
			expected:   "images/logo.png",
		},
		{
			name:       "arquivo de imagem .jpg",
			sourcePath: "photo.jpg",
			targetDir:  "src",
			expected:   "images/photo.jpg",
		},
		{
			name:       "arquivo de imagem .pdf",
			sourcePath: "diagram.pdf",
			targetDir:  "src",
			expected:   "images/diagram.pdf",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := loader.mapDestination(tt.sourcePath, tt.targetDir)
			if result != tt.expected {
				t.Errorf("mapDestination() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestIsTemplateFile(t *testing.T) {
	registry := NewRegistry()
	loader := NewLoader(registry)

	tests := []struct {
		name     string
		content  string
		expected bool
	}{
		{
			name:     "arquivo com {{.Title}}",
			content:  "\\title{{{.Title}}}",
			expected: true,
		},
		{
			name:     "arquivo com {{.Author}}",
			content:  "\\author{{{.Author}}}",
			expected: true,
		},
		{
			name:     "arquivo com {TITLE}",
			content:  "\\title{{TITLE}}",
			expected: true,
		},
		{
			name:     "arquivo com {AUTHOR}",
			content:  "\\author{{AUTHOR}}",
			expected: true,
		},
		{
			name:     "arquivo sem templates",
			content:  "\\documentclass{article}\n\\begin{document}\nHello World\n\\end{document}",
			expected: false,
		},
		{
			name:     "arquivo binário (simulado)",
			content:  string([]byte{0xFF, 0xFE, 0x00, 0x01}),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Criar arquivo temporário
			tempFile := filepath.Join(t.TempDir(), "test.tex")
			err := os.WriteFile(tempFile, []byte(tt.content), 0644)
			if err != nil {
				t.Fatalf("Erro ao criar arquivo temporário: %v", err)
			}

			result := loader.isTemplateFile(tempFile)
			if result != tt.expected {
				t.Errorf("isTemplateFile() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestCreateBaseDirectories(t *testing.T) {
	registry := NewRegistry()
	loader := NewLoader(registry)

	// Criar diretório temporário
	tempDir := t.TempDir()

	err := loader.createBaseDirectories(tempDir)
	if err != nil {
		t.Fatalf("createBaseDirectories() error = %v", err)
	}

	// Verificar se os diretórios foram criados
	expectedDirs := []string{
		tempDir,
		filepath.Join(tempDir, "chapters"),
		filepath.Join(tempDir, "images"),
		filepath.Join(tempDir, "styles"),
		"dist",
	}

	for _, dir := range expectedDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("Diretório %s não foi criado", dir)
		}
	}
}

func TestCopyFile(t *testing.T) {
	registry := NewRegistry()
	loader := NewLoader(registry)

	tests := []struct {
		name           string
		sourceContent  string
		sourceName     string
		expectNormalize bool
	}{
		{
			name: "arquivo .tex com normalização",
			sourceContent: `\usepackage{misc/options}
\input{frontmatter/title}`,
			sourceName:      "test.tex",
			expectNormalize: true,
		},
		{
			name:            "arquivo não .tex",
			sourceContent:   "binary content",
			sourceName:      "test.png",
			expectNormalize: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Criar diretórios temporários
			tempDir := t.TempDir()
			sourceFile := filepath.Join(tempDir, "source_"+tt.sourceName)
			destFile := filepath.Join(tempDir, "dest_"+tt.sourceName)

			// Criar arquivo fonte
			err := os.WriteFile(sourceFile, []byte(tt.sourceContent), 0644)
			if err != nil {
				t.Fatalf("Erro ao criar arquivo fonte: %v", err)
			}

			// Copiar arquivo
			err = loader.copyFile(sourceFile, destFile)
			if err != nil {
				t.Fatalf("copyFile() error = %v", err)
			}

			// Verificar se arquivo foi copiado
			if _, err := os.Stat(destFile); os.IsNotExist(err) {
				t.Errorf("Arquivo de destino não foi criado")
			}

			// Verificar normalização se esperada
			if tt.expectNormalize && strings.HasSuffix(tt.sourceName, ".tex") {
				content, err := os.ReadFile(destFile)
				if err != nil {
					t.Fatalf("Erro ao ler arquivo de destino: %v", err)
				}

				contentStr := string(content)
				if strings.Contains(contentStr, "misc/options") || strings.Contains(contentStr, "frontmatter/") {
					t.Errorf("Arquivo não foi normalizado corretamente: %s", contentStr)
				}
			}
		})
	}
}

func TestProcessGoTemplate(t *testing.T) {
	registry := NewRegistry()
	loader := NewLoader(registry)

	// Criar arquivo temporário com template
	tempDir := t.TempDir()
	sourceFile := filepath.Join(tempDir, "template.tex")
	destFile := filepath.Join(tempDir, "output.tex")

	templateContent := `\title{{TITLE}}
\author{{AUTHOR}}
\usepackage{misc/options}
\input{frontmatter/intro}`

	err := os.WriteFile(sourceFile, []byte(templateContent), 0644)
	if err != nil {
		t.Fatalf("Erro ao criar arquivo template: %v", err)
	}

	// Dados do projeto
	projectInfo := &types.ProjectInfo{
		Title:    "Test Title",
		Author:   "Test Author",
		Type:     "book",
		Language: "pt-br",
	}

	variables := map[string]string{
		"CustomVar": "CustomValue",
	}

	// Processar template
	err = loader.processGoTemplate(sourceFile, destFile, projectInfo, variables)
	if err != nil {
		t.Fatalf("processGoTemplate() error = %v", err)
	}

	// Verificar resultado
	result, err := os.ReadFile(destFile)
	if err != nil {
		t.Fatalf("Erro ao ler arquivo resultado: %v", err)
	}

	resultStr := string(result)

	// Verificar substituições
	if !strings.Contains(resultStr, "Test Title") {
		t.Errorf("Título não foi substituído corretamente")
	}

	if !strings.Contains(resultStr, "Test Author") {
		t.Errorf("Autor não foi substituído corretamente")
	}

	// Verificar normalização
	if strings.Contains(resultStr, "misc/options") {
		t.Errorf("Paths não foram normalizados")
	}

	if strings.Contains(resultStr, "frontmatter/") {
		t.Errorf("Paths não foram normalizados")
	}
}
