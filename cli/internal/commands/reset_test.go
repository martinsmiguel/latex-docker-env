package commands

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResetCommand(t *testing.T) {
	// Salvar diretório atual
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Erro ao obter diretório atual: %v", err)
	}
	defer os.Chdir(originalDir)

	// Criar diretório temporário para teste
	tempDir := t.TempDir()
	os.Chdir(tempDir)

	tests := []struct {
		name        string
		setupDirs   []string
		setupFiles  []string
		expectErr   bool
		expectDirs  []string // Diretórios que devem ser removidos
	}{
		{
			name:       "reset com todas as pastas existentes",
			setupDirs:  []string{"src", "dist", "tmp", "config"},
			setupFiles: []string{"src/main.tex", "dist/main.pdf", "tmp/temp.aux"},
			expectErr:  false,
			expectDirs: []string{"src", "dist", "tmp"},
		},
		{
			name:       "reset sem pastas existentes",
			setupDirs:  []string{},
			setupFiles: []string{},
			expectErr:  false,
			expectDirs: []string{},
		},
		{
			name:       "reset com apenas algumas pastas",
			setupDirs:  []string{"src", "config"},
			setupFiles: []string{"src/main.tex", "config/test.conf"},
			expectErr:  false,
			expectDirs: []string{"src"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup: criar diretórios e arquivos de teste
			for _, dir := range tt.setupDirs {
				err := os.MkdirAll(dir, 0755)
				if err != nil {
					t.Fatalf("Erro ao criar diretório %s: %v", dir, err)
				}
			}

			for _, file := range tt.setupFiles {
				dir := filepath.Dir(file)
				if dir != "." {
					err := os.MkdirAll(dir, 0755)
					if err != nil {
						t.Fatalf("Erro ao criar diretório %s: %v", dir, err)
					}
				}

				err := os.WriteFile(file, []byte("test content"), 0644)
				if err != nil {
					t.Fatalf("Erro ao criar arquivo %s: %v", file, err)
				}
			}

			// Forçar reset para evitar interação
			originalForce := resetForce
			resetForce = true
			defer func() { resetForce = originalForce }()

			// Executar reset
			err := resetEnvironment()

			// Verificar resultado
			if (err != nil) != tt.expectErr {
				t.Errorf("resetEnvironment() error = %v, expectErr %v", err, tt.expectErr)
				return
			}

			// Verificar se as pastas esperadas foram removidas
			for _, dir := range tt.expectDirs {
				if _, err := os.Stat(dir); !os.IsNotExist(err) {
					t.Errorf("Diretório %s deveria ter sido removido", dir)
				}
			}

			// Verificar se o diretório config não foi removido (se existia)
			for _, setupDir := range tt.setupDirs {
				if setupDir == "config" {
					if _, err := os.Stat("config"); os.IsNotExist(err) {
						t.Errorf("Diretório config não deveria ter sido removido")
					}
				}
			}
		})
	}
}

func TestRemoveFolder(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(string) error
		path      string
		expectErr bool
	}{
		{
			name: "remover pasta existente",
			setup: func(path string) error {
				return os.MkdirAll(path, 0755)
			},
			path:      "test_folder",
			expectErr: false,
		},
		{
			name: "remover pasta inexistente",
			setup: func(path string) error {
				return nil // Não criar nada
			},
			path:      "nonexistent_folder",
			expectErr: false, // Não deve dar erro
		},
		{
			name: "remover pasta com arquivos",
			setup: func(path string) error {
				if err := os.MkdirAll(path, 0755); err != nil {
					return err
				}
				return os.WriteFile(filepath.Join(path, "test.txt"), []byte("test"), 0644)
			},
			path:      "test_folder_with_files",
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Criar diretório temporário
			tempDir := t.TempDir()
			testPath := filepath.Join(tempDir, tt.path)

			// Setup
			err := tt.setup(testPath)
			if err != nil {
				t.Fatalf("Erro no setup: %v", err)
			}

			// Executar função
			err = removeFolder(testPath)

			// Verificar resultado
			if (err != nil) != tt.expectErr {
				t.Errorf("removeFolder() error = %v, expectErr %v", err, tt.expectErr)
				return
			}

			// Verificar se a pasta foi removida
			if !tt.expectErr {
				if _, err := os.Stat(testPath); !os.IsNotExist(err) {
					t.Errorf("Pasta %s deveria ter sido removida", testPath)
				}
			}
		})
	}
}

func TestStopDockerContainers(t *testing.T) {
	// Este teste verifica se a função não falha mesmo quando o Docker não está disponível
	t.Run("stop containers sem docker", func(t *testing.T) {
		// Criar diretório temporário sem arquivo docker-compose.yml
		tempDir := t.TempDir()
		originalDir, _ := os.Getwd()
		defer os.Chdir(originalDir)
		os.Chdir(tempDir)

		// A função deve lidar graciosamente com a ausência do Docker
		err := stopDockerContainers()

		// Esperamos um erro porque não há arquivo docker-compose.yml
		if err == nil {
			t.Log("stopDockerContainers() executou sem erro (Docker pode estar disponível)")
		} else {
			t.Logf("stopDockerContainers() retornou erro esperado: %v", err)
		}
	})
}
