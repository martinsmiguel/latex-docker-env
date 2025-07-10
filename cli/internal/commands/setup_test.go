package commands

import (
	"os"
	"testing"
)

func TestSetupCommand(t *testing.T) {
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
		createDirs  []string
		createFiles []string
		wantErr     bool
	}{
		{
			name:        "estrutura inexistente",
			createDirs:  []string{},
			createFiles: []string{},
			wantErr:     true,
		},
		{
			name:        "estrutura válida",
			createDirs:  []string{"config", "lib", "docs", "config/docker"},
			createFiles: []string{"config/latex-cli.conf", "config/docker/docker-compose.yml"},
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Limpar diretório de teste
			if err := os.RemoveAll(tempDir); err != nil {
				t.Fatalf("Erro ao limpar diretório: %v", err)
			}
			if err := os.MkdirAll(tempDir, 0755); err != nil {
				t.Fatalf("Erro ao recriar diretório: %v", err)
			}
			if err := os.Chdir(tempDir); err != nil {
				t.Fatalf("Erro ao mudar para diretório temporário: %v", err)
			}

			// Criar estrutura de teste
			for _, dir := range tt.createDirs {
				if err := os.MkdirAll(dir, 0755); err != nil {
					t.Fatalf("Erro ao criar diretório %s: %v", dir, err)
				}
			}
			for _, file := range tt.createFiles {
				f, err := os.Create(file)
				if err != nil {
					t.Fatalf("Erro ao criar arquivo %s: %v", file, err)
				}
				if err := f.Close(); err != nil {
					t.Fatalf("Erro ao fechar arquivo %s: %v", file, err)
				}
			}

			// Testar apenas a verificação da estrutura
			err := verifyProjectStructure()
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyProjectStructure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateDirectories(t *testing.T) {
	// Criar diretório temporário para teste
	tempDir := t.TempDir()
	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)

	os.Chdir(tempDir)

	err := createDirectories()
	if err != nil {
		t.Errorf("createDirectories() error = %v", err)
	}

	// Verificar se os diretórios foram criados
	expectedDirs := []string{"src", "dist", "tmp"}
	for _, dir := range expectedDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("Diretório %s não foi criado", dir)
		}
	}
}
