package types

import "time"

// Config representa a configuração da CLI
type Config struct {
	LatexEngine   string `mapstructure:"latex_engine"`
	OutputDir     string `mapstructure:"output_dir"`
	SourceDir     string `mapstructure:"source_dir"`
	ContainerName string `mapstructure:"container_name"`
	ImageName     string `mapstructure:"image_name"`
	WatchDebounce string `mapstructure:"watch_debounce"`
}

// ProjectInfo contém informações do projeto LaTeX
type ProjectInfo struct {
	Title       string
	Author      string
	Type        string // article, book, thesis, etc.
	Language    string
	Bibliography bool
}

// BuildOptions representa opções de compilação
type BuildOptions struct {
	Engine      string
	OutputDir   string
	CleanFirst  bool
	Verbose     bool
	Watch       bool
}

// TemplateMetadata representa os metadados de um template
type TemplateMetadata struct {
	Name         string            `yaml:"name"`
	Description  string            `yaml:"description"`
	Type         string            `yaml:"type"`         // article, thesis, book, presentation
	Author       string            `yaml:"author"`
	Version      string            `yaml:"version"`
	Language     string            `yaml:"language"`
	Dependencies []string          `yaml:"dependencies"` // pacotes LaTeX necessários
	Files        []TemplateFile    `yaml:"files"`
	Variables    map[string]string `yaml:"variables"`    // variáveis personalizáveis
	CreatedAt    time.Time         `yaml:"created_at"`
}

// TemplateFile representa um arquivo dentro de um template
type TemplateFile struct {
	Source      string `yaml:"source"`      // arquivo no template
	Destination string `yaml:"destination"` // onde será copiado
	Required    bool   `yaml:"required"`
	Template    bool   `yaml:"template"`    // se deve processar como template Go
}

// Template representa um template completo
type Template struct {
	Metadata TemplateMetadata
	Path     string
}
