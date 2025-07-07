package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestGetLatexImage(t *testing.T) {
	// Resetar viper para cada teste
	viper.Reset()

	tests := []struct {
		name     string
		setValue string
		expected string
	}{
		{
			name:     "valor padrão",
			setValue: "",
			expected: DefaultLatexImage,
		},
		{
			name:     "valor customizado",
			setValue: "custom/latex:latest",
			expected: "custom/latex:latest",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Reset()
			if tt.setValue != "" {
				viper.Set("latex_image", tt.setValue)
			}

			result := GetLatexImage()
			if result != tt.expected {
				t.Errorf("GetLatexImage() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSetDefaults(t *testing.T) {
	viper.Reset()
	SetDefaults()

	tests := []struct {
		key      string
		expected interface{}
	}{
		{"latex_engine", "xelatex"},
		{"output_dir", DefaultOutputDir},
		{"source_dir", DefaultSourceDir},
		{"container_name", "latex-env"},
		{"latex_image", DefaultLatexImage},
		{"watch_debounce", "500ms"},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			result := viper.Get(tt.key)
			if result != tt.expected {
				t.Errorf("SetDefaults() - %s = %v, expected %v", tt.key, result, tt.expected)
			}
		})
	}
}
