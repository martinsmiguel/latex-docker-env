package template

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/martinsmiguel/latex-docker-env/cli/pkg/types"
	"gopkg.in/yaml.v3"
)

type Registry struct {
	templates map[string]*types.Template
	paths     []string
}

func NewRegistry() *Registry {
	return &Registry{
		templates: make(map[string]*types.Template),
		paths:     []string{},
	}
}

func (r *Registry) AddTemplatePath(path string) {
	r.paths = append(r.paths, path)
}

func (r *Registry) LoadTemplates() error {
	r.templates = make(map[string]*types.Template)

	for _, basePath := range r.paths {
		if err := r.loadTemplatesFromPath(basePath); err != nil {
			return fmt.Errorf("erro ao carregar templates de %s: %w", basePath, err)
		}
	}

	return nil
}

func (r *Registry) loadTemplatesFromPath(basePath string) error {
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		return nil // Não é erro se o diretório não existir
	}

	entries, err := os.ReadDir(basePath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		templatePath := filepath.Join(basePath, entry.Name())
		metadataPath := filepath.Join(templatePath, "template.yaml")

		var template *types.Template

		if _, err := os.Stat(metadataPath); os.IsNotExist(err) {
			// Template sem metadata - criar automaticamente
			template = r.createAutoTemplate(templatePath, entry.Name())
		} else {
			// Template com metadata
			template, err = r.loadTemplate(templatePath)
			if err != nil {
				fmt.Printf("Aviso: erro ao carregar template %s: %v\n", entry.Name(), err)
				continue
			}
		}

		if template != nil {
			r.templates[template.Metadata.Name] = template
		}
	}

	return nil
}

func (r *Registry) loadTemplate(templatePath string) (*types.Template, error) {
	metadataPath := filepath.Join(templatePath, "template.yaml")

	data, err := os.ReadFile(metadataPath)
	if err != nil {
		return nil, err
	}

	var metadata types.TemplateMetadata
	if err := yaml.Unmarshal(data, &metadata); err != nil {
		return nil, err
	}

	return &types.Template{
		Metadata: metadata,
		Path:     templatePath,
	}, nil
}

func (r *Registry) GetTemplate(name string) (*types.Template, error) {
	template, exists := r.templates[name]
	if !exists {
		return nil, fmt.Errorf("template '%s' não encontrado", name)
	}
	return template, nil
}

func (r *Registry) ListTemplates() []*types.Template {
	templates := make([]*types.Template, 0, len(r.templates))
	for _, template := range r.templates {
		templates = append(templates, template)
	}

	// Ordenar por nome
	sort.Slice(templates, func(i, j int) bool {
		return templates[i].Metadata.Name < templates[j].Metadata.Name
	})

	return templates
}

func (r *Registry) GetTemplatesByType(templateType string) []*types.Template {
	var filtered []*types.Template

	for _, template := range r.templates {
		if strings.EqualFold(template.Metadata.Type, templateType) {
			filtered = append(filtered, template)
		}
	}

	return filtered
}

func (r *Registry) TemplateExists(name string) bool {
	_, exists := r.templates[name]
	return exists
}

// Cria template automaticamente para diretórios sem metadata
func (r *Registry) createAutoTemplate(templatePath, dirName string) *types.Template {
	// Detectar tipo baseado no conteúdo
	templateType := r.detectTemplateType(templatePath)

	return &types.Template{
		Metadata: types.TemplateMetadata{
			Name:        dirName,
			Description: fmt.Sprintf("Template detectado automaticamente: %s", dirName),
			Type:        templateType,
			Author:      "Auto-detectado",
			Version:     "1.0.0",
			Language:    "multilingual",
			Dependencies: []string{},
			Files:       []types.TemplateFile{}, // Vazio para usar detecção automática
			Variables:   make(map[string]string),
		},
		Path: templatePath,
	}
}

// Detecta o tipo de template baseado nos arquivos presentes
func (r *Registry) detectTemplateType(templatePath string) string {
	var hasBeamer bool
	var hasBook bool
	var hasThesis bool

	err := filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		if strings.HasSuffix(info.Name(), ".tex") {
			content, err := os.ReadFile(path)
			if err != nil {
				return nil
			}

			contentStr := strings.ToLower(string(content))

			if strings.Contains(contentStr, "\\documentclass{beamer}") ||
			   strings.Contains(contentStr, "beamer") {
				hasBeamer = true
			}

			if strings.Contains(contentStr, "\\documentclass{book}") ||
			   strings.Contains(contentStr, "\\chapter") {
				hasBook = true
			}

			if strings.Contains(contentStr, "thesis") ||
			   strings.Contains(contentStr, "dissertation") {
				hasThesis = true
			}
		}

		return nil
	})

	if err != nil {
		// Em caso de erro, retornar tipo padrão
		return "article"
	}

	switch {
	case hasBeamer:
		return "presentation"
	case hasThesis:
		return "thesis"
	case hasBook:
		return "book"
	default:
		return "article"
	}
}
