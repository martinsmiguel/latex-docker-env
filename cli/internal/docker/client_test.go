package docker

import (
	"context"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name        string
		expectError bool
		testDesc    string
	}{
		{
			name:        "criar cliente docker",
			expectError: false, // Pode falhar se Docker não estiver disponível, mas não deve causar pânico
			testDesc:    "should create docker client without panic",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Verificar se não há pânico ao criar cliente
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("NewClient() caused panic: %v", r)
				}
			}()

			client, err := NewClient()

			// Se não há erro, deve ter um cliente
			if err == nil && client == nil {
				t.Error("NewClient() returned nil client without error")
			}

			// Se há um cliente, deve ser possível fechá-lo
			if client != nil {
				defer func() {
					if closeErr := client.Close(); closeErr != nil {
						t.Logf("Warning: error closing client: %v", closeErr)
					}
				}()
			}

			// Log do resultado para depuração
			if err != nil {
				t.Logf("NewClient() error (expected if Docker not available): %v", err)
			} else {
				t.Logf("NewClient() success")
			}
		})
	}
}

func TestClient_Close(t *testing.T) {
	// Tentar criar um cliente para testar Close
	client, err := NewClient()
	if err != nil {
		t.Skipf("Skipping Close test, cannot create Docker client: %v", err)
	}

	// Testar Close
	err = client.Close()
	if err != nil {
		t.Errorf("Close() error = %v", err)
	}

	// Testar Close duplo (não deve causar pânico)
	err = client.Close()
	// Segundo Close pode retornar erro, mas não deve causar pânico
	t.Logf("Second Close() returned: %v", err)
}

func TestClient_PullImage(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Skipf("Skipping PullImage test, cannot create Docker client: %v", err)
	}
	defer func() {
		if closeErr := client.Close(); closeErr != nil {
			t.Logf("Warning: error closing client: %v", closeErr)
		}
	}()

	tests := []struct {
		name      string
		imageName string
		wantErr   bool
	}{
		{
			name:      "imagem inválida",
			imageName: "invalid/nonexistent:impossible",
			wantErr:   true,
		},
		{
			name:      "nome de imagem vazio",
			imageName: "",
			wantErr:   true,
		},
		// Não testamos pull de imagem real para não depender de conexão de rede
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			// Verificar se não há pânico
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("PullImage() caused panic: %v", r)
				}
			}()

			err := client.PullImage(ctx, tt.imageName)
			if (err != nil) != tt.wantErr {
				t.Errorf("PullImage() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				t.Logf("PullImage() error (expected): %v", err)
			}
		})
	}
}

func TestClient_Build(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Skipf("Skipping Build test, cannot create Docker client: %v", err)
	}
	defer func() {
		if closeErr := client.Close(); closeErr != nil {
			t.Logf("Warning: error closing client: %v", closeErr)
		}
	}()

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "build simples",
			wantErr: false, // Build atual é apenas uma simulação
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			// Verificar se não há pânico
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Build() caused panic: %v", r)
				}
			}()

			err := client.Build(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				t.Logf("Build() error: %v", err)
			}
		})
	}
}

func TestClient_GetContainerStatus(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Skipf("Skipping GetContainerStatus test, cannot create Docker client: %v", err)
	}
	defer func() {
		if closeErr := client.Close(); closeErr != nil {
			t.Logf("Warning: error closing client: %v", closeErr)
		}
	}()

	tests := []struct {
		name          string
		containerName string
		expectStatus  bool // Se deve retornar algum status (mesmo que "not found")
	}{
		{
			name:          "container inexistente",
			containerName: "nonexistent-container-12345",
			expectStatus:  true, // Deve retornar status "not found" ou similar
		},
		{
			name:          "nome vazio",
			containerName: "",
			expectStatus:  true, // Deve retornar algum status
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			// Verificar se não há pânico
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("GetContainerStatus() caused panic: %v", r)
				}
			}()

			status, err := client.GetContainerStatus(ctx, tt.containerName)

			// Não deve causar pânico
			t.Logf("GetContainerStatus(%s) = status: %s, error: %v", tt.containerName, status, err)

			// Status deve ser uma string (mesmo que vazia)
			if tt.expectStatus && status == "" && err == nil {
				t.Errorf("GetContainerStatus() returned empty status without error")
			}
		})
	}
}

func TestClientWorkflow(t *testing.T) {
	// Teste de workflow completo do cliente Docker
	client, err := NewClient()
	if err != nil {
		t.Skipf("Skipping workflow test, cannot create Docker client: %v", err)
	}

	// Verificar se não há pânico durante operações sequenciais
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Docker client workflow caused panic: %v", r)
		}
	}()

	ctx := context.Background()

	// 1. Verificar status de container inexistente
	status, err := client.GetContainerStatus(ctx, "test-workflow-container")
	t.Logf("Initial status check: %s, error: %v", status, err)

	// 2. Tentar pull de imagem inexistente (deve falhar, mas não causar pânico)
	err = client.PullImage(ctx, "invalid/test:latest")
	t.Logf("Pull invalid image result: %v", err)

	// 3. Tentar build com parâmetros inválidos (deve falhar, mas não causar pânico)
	err = client.Build(ctx)
	t.Logf("Build with invalid params result: %v", err)

	// 4. Fechar cliente
	err = client.Close()
	if err != nil {
		t.Errorf("Close() error = %v", err)
	}

	t.Log("Docker client workflow completed without panic")
}

func TestErrorHandling(t *testing.T) {
	// Testar tratamento de erros em cenários extremos
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "new_client_multiple_times",
			test: func(t *testing.T) {
				// Criar múltiplos clientes
				for i := 0; i < 3; i++ {
					client, err := NewClient()
					if err != nil {
						t.Logf("NewClient() iteration %d error: %v", i, err)
						continue
					}

					if closeErr := client.Close(); closeErr != nil {
						t.Logf("Close() iteration %d error: %v", i, closeErr)
					}
				}
			},
		},
		{
			name: "operations_on_closed_client",
			test: func(t *testing.T) {
				client, err := NewClient()
				if err != nil {
					t.Skipf("Cannot create client: %v", err)
				}

				// Fechar o cliente
				if closeErr := client.Close(); closeErr != nil {
					t.Fatalf("Error closing client: %v", closeErr)
				}

				ctx := context.Background()

				// Tentar operações em cliente fechado (não deve causar pânico)
				_, err = client.GetContainerStatus(ctx, "test")
				t.Logf("GetContainerStatus on closed client: %v", err)

				err = client.PullImage(ctx, "test:latest")
				t.Logf("PullImage on closed client: %v", err)

				err = client.Build(ctx)
				t.Logf("Build on closed client: %v", err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Test %s caused panic: %v", tt.name, r)
				}
			}()

			tt.test(t)
		})
	}
}

func TestContextCancellation(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Skipf("Skipping context test, cannot create Docker client: %v", err)
	}
	defer func() {
		if closeErr := client.Close(); closeErr != nil {
			t.Logf("Warning: error closing client: %v", closeErr)
		}
	}()

	// Criar contexto cancelado
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancelar imediatamente

	// Operações com contexto cancelado
	_, err = client.GetContainerStatus(ctx, "test")
	t.Logf("GetContainerStatus with cancelled context: %v", err)

	err = client.PullImage(ctx, "test:latest")
	t.Logf("PullImage with cancelled context: %v", err)

	err = client.Build(ctx)
	t.Logf("Build with cancelled context: %v", err)

	// Todas as operações devem retornar erro de contexto cancelado,
	// mas não devem causar pânico
}
