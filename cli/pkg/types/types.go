package types

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
