package config

import (
	"github.com/spf13/viper"
	"github.com/martinsmiguel/latex-docker-env/cli/pkg/types"
)

const (
	DefaultLatexImage = "blang/latex:ubuntu"
	DefaultOutputDir  = "dist"
	DefaultSourceDir  = "src"
)

func GetLatexImage() string {
	image := viper.GetString("latex_image")
	if image == "" {
		return DefaultLatexImage
	}
	return image
}

func GetConfig() *types.Config {
	return &types.Config{
		LatexEngine:   viper.GetString("latex_engine"),
		OutputDir:     viper.GetString("output_dir"),
		SourceDir:     viper.GetString("source_dir"),
		ContainerName: viper.GetString("container_name"),
		ImageName:     viper.GetString("image_name"),
		WatchDebounce: viper.GetString("watch_debounce"),
	}
}

func SetDefaults() {
	viper.SetDefault("latex_engine", "xelatex")
	viper.SetDefault("output_dir", DefaultOutputDir)
	viper.SetDefault("source_dir", DefaultSourceDir)
	viper.SetDefault("container_name", "latex-env")
	viper.SetDefault("latex_image", DefaultLatexImage)
	viper.SetDefault("watch_debounce", "500ms")
}
