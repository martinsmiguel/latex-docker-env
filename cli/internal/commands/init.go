package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/martinsmiguel/latex-docker-env/cli/pkg/types"
)

var (
	initTitle    string
	initAuthor   string
	initTemplate string
	initForce    bool
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Inicializa um novo documento LaTeX",
	Long: `Inicializa um novo documento LaTeX com base em templates.

Templates disponíveis:
  default - Template básico para documentos gerais
  article - Template para artigos científicos
  thesis  - Template para teses e dissertações
  report  - Template para relatórios técnicos
  book    - Template para livros`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return initProject()
	},
}

func init() {
	InitCmd.Flags().StringVarP(&initTitle, "title", "t", "", "Título do documento")
	InitCmd.Flags().StringVarP(&initAuthor, "author", "a", "", "Nome do autor")
	InitCmd.Flags().StringVar(&initTemplate, "template", "default", "Template a usar")
	InitCmd.Flags().BoolVarP(&initForce, "force", "f", false, "Sobrescreve arquivos existentes")
}

func initProject() error {
	fmt.Println(">> Inicializando novo documento LaTeX...")

	// Verificar se já existe projeto
	sourceDir := "src"
	mainTexPath := filepath.Join(sourceDir, "main.tex")

	if _, err := os.Stat(mainTexPath); err == nil && !initForce {
		return fmt.Errorf("já existe um documento em %s. Use --force para sobrescrever", mainTexPath)
	}

	// Criar estrutura de diretórios
	dirs := []string{
		sourceDir,
		filepath.Join(sourceDir, "chapters"),
		filepath.Join(sourceDir, "images"),
		"dist",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("erro ao criar diretório %s: %w", dir, err)
		}
	}

	// Obter dados do projeto
	projectInfo := &types.ProjectInfo{
		Title:    getTitle(),
		Author:   getAuthor(),
		Type:     initTemplate,
		Language: "portuguese",
		Bibliography: true,
	}

	// Criar arquivos do template
	if err := createFromTemplate(projectInfo, sourceDir); err != nil {
		return fmt.Errorf("erro ao criar arquivos do template: %w", err)
	}

	fmt.Printf("[SUCCESS] Documento LaTeX inicializado com sucesso!\n")
	fmt.Printf("[INFO] Arquivos criados em: %s\n", sourceDir)
	fmt.Printf("[INFO] Para compilar: ltx build\n")

	return nil
}

func getTitle() string {
	if initTitle != "" {
		return initTitle
	}
	return "Meu Documento"
}

func getAuthor() string {
	if initAuthor != "" {
		return initAuthor
	}
	return "Autor"
}

func createFromTemplate(info *types.ProjectInfo, sourceDir string) error {
	// Template principal
	mainTemplate := `\documentclass{article}
\input{src/preamble}

\title{\textbf{[[.Title]]}}
\author{[[.Author]]}

\begin{document}
\maketitle

\section{Introdução}
Este é o texto inicial do documento. Para mais informações, veja \cite{exemplo}.

[[if .HasChapters]]
\input{src/chapters/introduction}
\input{src/chapters/methodology}
\input{src/chapters/results}
\input{src/chapters/conclusion}
[[end]]

[[if .Bibliography]]
\bibliography{src/references}
\bibliographystyle{plain}
[[end]]
\end{document}
`

	// Preamble template
	preambleTemplate := `% Codificação e idioma
\usepackage[utf8]{inputenc}
\usepackage[T1]{fontenc}
\usepackage[portuguese]{babel}

% Packages essenciais
\usepackage{graphicx}
\usepackage{amsmath}
\usepackage{amssymb}
\usepackage{hyperref}
\usepackage{geometry}
\usepackage{natbib}

% Configurações de página
\geometry{a4paper, margin=2.5cm}

% Configurações do hyperref
\hypersetup{
    colorlinks=true,
    linkcolor=blue,
    filecolor=magenta,
    urlcolor=cyan,
    citecolor=red
}
`

	// References template
	referencesTemplate := `@article{exemplo,
  title={Título do Artigo de Exemplo},
  author={Autor, Nome},
  journal={Journal Name},
  year={2023},
  volume={1},
  pages={1--10}
}

@book{livro-exemplo,
  title={Título do Livro},
  author={Sobrenome, Nome},
  publisher={Editora},
  year={2023}
}
`

	// Criar arquivo main.tex
	if err := createFileFromTemplate(filepath.Join(sourceDir, "main.tex"), mainTemplate, map[string]interface{}{
		"Title":       info.Title,
		"Author":      info.Author,
		"HasChapters": initTemplate != "article",
		"Bibliography": info.Bibliography,
	}); err != nil {
		return err
	}
	fmt.Printf("[SUCCESS] Criado: %s\n", filepath.Join(sourceDir, "main.tex"))

	// Criar preamble.tex
	if err := os.WriteFile(filepath.Join(sourceDir, "preamble.tex"), []byte(preambleTemplate), 0644); err != nil {
		return err
	}
	fmt.Printf("[SUCCESS] Criado: %s\n", filepath.Join(sourceDir, "preamble.tex"))

	// Criar references.bib
	if err := os.WriteFile(filepath.Join(sourceDir, "references.bib"), []byte(referencesTemplate), 0644); err != nil {
		return err
	}
	fmt.Printf("[SUCCESS] Criado: %s\n", filepath.Join(sourceDir, "references.bib"))

	// Criar capítulos se não for artigo
	if initTemplate != "article" {
		chapters := []struct {
			name    string
			content string
		}{
			{"introduction", "\\section{Introdução}\n\nEste é o capítulo de introdução.\n\n"},
			{"methodology", "\\section{Metodologia}\n\nDescreva a metodologia utilizada.\n\n"},
			{"results", "\\section{Resultados}\n\nApresente os resultados obtidos.\n\n"},
			{"conclusion", "\\section{Conclusão}\n\nApresente as conclusões do trabalho.\n\n"},
		}

		for _, chapter := range chapters {
			chapterPath := filepath.Join(sourceDir, "chapters", chapter.name+".tex")
			if err := os.WriteFile(chapterPath, []byte(chapter.content), 0644); err != nil {
				return err
			}
			fmt.Printf("[SUCCESS] Criado: %s\n", chapterPath)
		}
	}

	return nil
}

func createFileFromTemplate(filePath, templateStr string, data interface{}) error {
	tmpl, err := template.New("template").Delims("[[", "]]").Parse(templateStr)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}
