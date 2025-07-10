package template

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/martinsmiguel/latex-docker-env/cli/internal/colors"
	"github.com/martinsmiguel/latex-docker-env/cli/pkg/types"
)

type Loader struct {
	registry *Registry
}

func NewLoader(registry *Registry) *Loader {
	return &Loader{registry: registry}
}

func (l *Loader) CreateProject(templateName string, projectInfo *types.ProjectInfo, targetDir string) error {
	tmpl, err := l.registry.GetTemplate(templateName)
	if err != nil {
		return err
	}

	colors.Printf(">> Usando template: %s (%s)\n", tmpl.Metadata.Name, tmpl.Metadata.Description)

	// Criar estrutura de diretórios base
	if err := l.createBaseDirectories(targetDir); err != nil {
		return err
	}

	// Se o template tem definição de arquivos no metadata, usar sistema dinâmico
	if len(tmpl.Metadata.Files) > 0 {
		return l.createFromMetadata(tmpl, projectInfo, targetDir)
	}

	// Caso contrário, usar detecção automática de arquivos
	return l.createFromAutoDetection(tmpl, projectInfo, targetDir)
}

func (l *Loader) createBaseDirectories(targetDir string) error {
	dirs := []string{
		targetDir,
		filepath.Join(targetDir, "chapters"),
		filepath.Join(targetDir, "images"),
		filepath.Join(targetDir, "styles"),
		"dist",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("erro ao criar diretório %s: %w", dir, err)
		}
	}

	return nil
}

func (l *Loader) processTemplateFile(tmpl *types.Template, file types.TemplateFile, projectInfo *types.ProjectInfo, targetDir string) error {
	sourcePath := filepath.Join(tmpl.Path, file.Source)
	destPath := filepath.Join(targetDir, file.Destination)

	// Verificar se arquivo fonte existe
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		if file.Required {
			return fmt.Errorf("arquivo obrigatório não encontrado: %s", sourcePath)
		}
		return nil
	}

	// Criar diretório de destino se necessário
	destDir := filepath.Dir(destPath)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	if file.Template {
		// Processar como template Go
		return l.processGoTemplate(sourcePath, destPath, projectInfo, tmpl.Metadata.Variables)
	} else {
		// Copiar arquivo diretamente
		return l.copyFile(sourcePath, destPath)
	}
}

func (l *Loader) processGoTemplate(sourcePath, destPath string, projectInfo *types.ProjectInfo, variables map[string]string) error {
	content, err := os.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	// Substituir padrões simples primeiro (para compatibilidade com templates antigos)
	contentStr := string(content)
	contentStr = strings.ReplaceAll(contentStr, "{TITLE}", projectInfo.Title)
	contentStr = strings.ReplaceAll(contentStr, "{AUTHOR}", projectInfo.Author)
	contentStr = strings.ReplaceAll(contentStr, "{DATE}", "\\today")

	// Normalizar caminhos LaTeX para estrutura padrão do projeto
	contentStr = l.normalizeLaTeXPaths(contentStr)

	// Processar template Go se contém {{}}
	if strings.Contains(contentStr, "{{") {
		tmpl, err := template.New("template").Delims("{{", "}}").Parse(contentStr)
		if err != nil {
			return err
		}

		file, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer func() {
			if err := file.Close(); err != nil {
				fmt.Printf("Erro ao fechar arquivo: %v\n", err)
			}
		}()

		// Combinar dados do projeto com variáveis do template
		data := map[string]interface{}{
			"Title":     projectInfo.Title,
			"Author":    projectInfo.Author,
			"Type":      projectInfo.Type,
			"Language":  projectInfo.Language,
			"Variables": variables,
		}

		if err := tmpl.Execute(file, data); err != nil {
			return fmt.Errorf("erro ao executar template: %w", err)
		}
	} else {
		// Salvar conteúdo processado
		err = os.WriteFile(destPath, []byte(contentStr), 0644)
	}

	if err == nil {
		colors.Printf("[SUCCESS] Criado: %s\n", destPath)
	}

	return err
}

func (l *Loader) copyFile(sourcePath, destPath string) error {
	// Se for arquivo .tex, aplicar normalização de caminhos
	if strings.HasSuffix(strings.ToLower(sourcePath), ".tex") {
		content, err := os.ReadFile(sourcePath)
		if err != nil {
			return err
		}

		// Normalizar caminhos LaTeX
		normalizedContent := l.normalizeLaTeXPaths(string(content))

		err = os.WriteFile(destPath, []byte(normalizedContent), 0644)
		if err == nil {
			colors.Printf("[SUCCESS] Copiado: %s (normalizado)\n", destPath)
		}
		return err
	}

	// Para outros arquivos, copiar diretamente
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := sourceFile.Close(); err != nil {
			fmt.Printf("Erro ao fechar arquivo origem: %v\n", err)
		}
	}()

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := destFile.Close(); err != nil {
			fmt.Printf("Erro ao fechar arquivo destino: %v\n", err)
		}
	}()

	_, err = io.Copy(destFile, sourceFile)
	if err == nil {
		colors.Printf("[SUCCESS] Copiado: %s\n", destPath)
	}

	return err
}

// Cria projeto baseado nos metadados definidos
func (l *Loader) createFromMetadata(tmpl *types.Template, projectInfo *types.ProjectInfo, targetDir string) error {
	for _, file := range tmpl.Metadata.Files {
		if err := l.processTemplateFile(tmpl, file, projectInfo, targetDir); err != nil {
			if file.Required {
				return fmt.Errorf("erro ao processar arquivo obrigatório %s: %w", file.Source, err)
			}
			colors.Printf("[WARN] Erro ao processar arquivo opcional %s: %v\n", file.Source, err)
		}
	}
	return nil
}

// Detecta automaticamente arquivos no template e os processa
func (l *Loader) createFromAutoDetection(tmpl *types.Template, projectInfo *types.ProjectInfo, targetDir string) error {
	colors.Printf(">> Detectando arquivos automaticamente em: %s\n", tmpl.Path)

	return filepath.Walk(tmpl.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Pular o arquivo de metadata
		if info.Name() == "template.yaml" {
			return nil
		}

		// Pular diretórios
		if info.IsDir() {
			return nil
		}

		// Calcular caminho relativo
		relPath, err := filepath.Rel(tmpl.Path, path)
		if err != nil {
			return err
		}

		// Criar arquivo template fictício para processamento
		templateFile := types.TemplateFile{
			Source:      relPath,
			Destination: l.mapDestination(relPath, targetDir),
			Required:    false,
			Template:    l.isTemplateFile(path),
		}

		return l.processTemplateFile(tmpl, templateFile, projectInfo, targetDir)
	})
}

// Mapeia arquivos do template para destinos apropriados
func (l *Loader) mapDestination(sourcePath, targetDir string) string {
	// Regras de mapeamento inteligente
	switch {
	case strings.HasSuffix(sourcePath, ".tex"):
		// Arquivos .tex principais vão para a raiz
		if strings.Contains(sourcePath, "main") || strings.Contains(sourcePath, "template") {
			return "main.tex"
		}
		// Outros .tex vão para chapters se estiverem em subdiretórios
		if strings.Contains(sourcePath, "/") {
			return filepath.Join("chapters", filepath.Base(sourcePath))
		}
		return sourcePath

	case strings.HasSuffix(sourcePath, ".sty") || strings.HasSuffix(sourcePath, ".cls"):
		// Arquivos de estilo vão para styles
		return filepath.Join("styles", filepath.Base(sourcePath))

	case strings.HasSuffix(sourcePath, ".bib"):
		// Bibliografia vai para references.bib
		return "references.bib"

	case isImageFile(sourcePath):
		// Imagens vão para images
		return filepath.Join("images", filepath.Base(sourcePath))

	default:
		// Outros arquivos mantêm estrutura original
		return sourcePath
	}
}

func (l *Loader) isTemplateFile(path string) bool {
	// Verifica se arquivo contém variáveis de template
	content, err := os.ReadFile(path)
	if err != nil {
		return false
	}

	contentStr := string(content)
	// Procura por padrões de template comuns, mas apenas se são ASCII válidos
	templatePatterns := []string{
		"{{.Title}}",
		"{{.Author}}",
		"{TITLE}",
		"{AUTHOR}",
		"{DATE}",
	}

	for _, pattern := range templatePatterns {
		if strings.Contains(contentStr, pattern) {
			return true
		}
	}

	// Evitar arquivos com caracteres especiais ou binários
	for _, b := range content {
		if b > 127 && b < 160 {
			return false // Caracteres de controle problemáticos
		}
	}

	// Verificar se contém {{ }} genérico apenas se for texto válido
	if strings.Contains(contentStr, "{{") && strings.Contains(contentStr, "}}") {
		// Verificar se não há caracteres problemáticos entre {{ }}
		start := strings.Index(contentStr, "{{")
		end := strings.Index(contentStr, "}}")
		if start != -1 && end != -1 && end > start {
			substr := contentStr[start:end+2]
			// Se contém caracteres não-ASCII problemáticos, não é template
			for _, r := range substr {
				if r > 127 {
					return false
				}
			}
			return true
		}
	}

	return false
}

func isImageFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	imageExts := []string{".png", ".jpg", ".jpeg", ".pdf", ".svg", ".eps"}

	for _, imgExt := range imageExts {
		if ext == imgExt {
			return true
		}
	}
	return false
}

// normalizeLaTeXPaths normaliza caminhos LaTeX para a estrutura padrão do projeto
func (l *Loader) normalizeLaTeXPaths(content string) string {
	// Mapa de substituições conhecidas para normalizar caminhos
	pathReplacements := map[string]string{
		// Pacotes e arquivos de estilo - manter .sty funcionando
		"\\usepackage{misc/options}":        "\\usepackage{options}",
		"\\usepackage{styles/options}":      "\\usepackage{options}",
		"\\usepackage{misc/":                "\\usepackage{", // Remover prefixo misc/
		"\\usepackage{style/":               "\\usepackage{", // Remover prefixo style/
		"\\usepackage{sty/":                 "\\usepackage{", // Remover prefixo sty/
		"\\input{misc/":                     "\\input{styles/",
		"\\input{style/":                    "\\input{styles/",
		"\\input{sty/":                      "\\input{styles/",

		// Imagens - unificar todos os diretórios para images/ (TEXINPATHS resolve)
		"\\includegraphics{frontmatter/":    "\\includegraphics{images/",
		"\\includegraphics{image/":          "\\includegraphics{images/",
		"\\includegraphics{img/":            "\\includegraphics{images/",
		"\\includegraphics{figures/":        "\\includegraphics{images/",
		"\\includegraphics{fig/":            "\\includegraphics{images/",
		"\\includegraphics{graphics/":       "\\includegraphics{images/",
		"\\includegraphics{assets/":         "\\includegraphics{images/",

		// Padrões específicos de imagem com parâmetros
		"]{frontmatter/":                    "]{images/",
		"]{image/":                          "]{images/",
		"]{img/":                            "]{images/",
		"]{figures/":                        "]{images/",
		"]{fig/":                            "]{images/",
		"]{graphics/":                       "]{images/",
		"]{assets/":                         "]{images/",

		// Capítulos e seções - manter paths relativos do src/
		"\\input{frontmatter/":              "\\input{chapters/",
		"\\input{content/":                  "\\input{chapters/",
		"\\input{chapter/":                  "\\input{chapters/",
		"\\input{section/":                  "\\input{chapters/",
		"\\input{sections/":                 "\\input{chapters/",
		"\\include{frontmatter/":            "\\include{chapters/",
		"\\include{content/":                "\\include{chapters/",
		"\\include{chapter/":                "\\include{chapters/",
		"\\include{section/":                "\\include{chapters/",
		"\\include{sections/":               "\\include{chapters/",

		// Backmatter
		"\\input{back/":                     "\\input{chapters/",
		"\\input{backmatter/":               "\\input{chapters/",
	}

	// Aplicar substituições
	for old, new := range pathReplacements {
		content = strings.ReplaceAll(content, old, new)
	}

	// Normalizar caminhos relativos para absolutos (remover ../ e ./)
	content = strings.ReplaceAll(content, "../", "")
	content = strings.ReplaceAll(content, "./", "")

	return content
}
