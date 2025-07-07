# Arquitetura da CLI ltx

## Visão Geral

A CLI `ltx` segue os princípios de design da arquitetura Go, com separação clara de responsabilidades e estrutura modular.

## Estrutura de Diretórios

```
cli/
├── cmd/                    # Interface de linha de comando (Cobra)
│   └── root.go            # Comando raiz e configuração global
├── internal/              # Código interno da aplicação
│   ├── commands/          # Implementação dos comandos CLI
│   ├── config/            # Gerenciamento de configuração
│   ├── docker/            # Cliente e operações Docker
│   └── utils/             # Utilitários compartilhados
├── pkg/                   # Packages públicos reutilizáveis
│   └── types/             # Tipos e structs compartilhados
├── main.go               # Entry point da aplicação
├── go.mod                # Definição do módulo Go
└── go.sum                # Lock file das dependências
```

## Componentes Principais

### 1. Entry Point (`main.go`)
```go
package main

import "github.com/martinsmiguel/latex-docker-env/cli/cmd"

func main() {
    cmd.Execute()
}
```
- Ponto de entrada simples
- Delega execução para o pacote `cmd`

### 2. Command Interface (`cmd/`)
```go
// root.go
var rootCmd = &cobra.Command{
    Use:   "ltx",
    Short: "LaTeX Docker Environment CLI",
    // ...
}
```
- Baseado no framework Cobra
- Gerencia flags globais e configuração
- Registra todos os subcomandos

### 3. Commands (`internal/commands/`)
```go
// Padrão para todos os comandos
var SetupCmd = &cobra.Command{
    Use:   "setup",
    Short: "Configura o ambiente de desenvolvimento",
    RunE:  runSetup,
}

func runSetup(cmd *cobra.Command, args []string) error {
    // Implementação do comando
}
```
- Um arquivo por comando principal
- Função `RunE` para tratamento de erros
- Validação de argumentos e flags

### 4. Configuration (`internal/config/`)
```go
type Config struct {
    LatexEngine   string `mapstructure:"latex_engine"`
    OutputDir     string `mapstructure:"output_dir"`
    // ...
}
```
- Integração com Viper para configuração
- Suporte a arquivos, env vars e flags
- Configurações padrão definidas

### 5. Docker Integration (`internal/docker/`)
```go
type Client struct {
    cli *client.Client
}

func (c *Client) PullImage(ctx context.Context, imageName string) error {
    // Implementação
}
```
- Wrapper do Docker Client SDK
- Operações específicas para LaTeX
- Gestão de containers e imagens

### 6. Types (`pkg/types/`)
```go
type ProjectInfo struct {
    Title       string
    Author      string
    Type        string
    Language    string
    Bibliography bool
}
```
- Structs compartilhados
- Tipos de configuração
- Interfaces públicas

## Fluxo de Dados

### 1. Inicialização
```
main.go → cmd/root.go → cobra.OnInitialize → config loading
```

### 2. Execução de Comando
```
user input → cobra parsing → command validation → business logic → output
```

### 3. Configuração
```
flags → env vars → config file → defaults (precedência)
```

## Princípios de Design

### 1. Separation of Concerns
- `cmd/`: Interface CLI
- `internal/commands/`: Lógica de negócio
- `internal/docker/`: Integração externa
- `internal/config/`: Gerenciamento de estado

### 2. Dependency Injection
```go
func runSetup(cmd *cobra.Command, args []string) error {
    dockerClient, err := docker.NewClient()
    // Injeção via factory functions
}
```

### 3. Error Handling
```go
// Erros são propagados para o nível CLI
return fmt.Errorf("operação falhou: %w", err)
```

### 4. Configuration Management
- Configuração centralizada via Viper
- Valores padrão sensatos
- Override via flags e env vars

## Padrões de Implementação

### 1. Command Pattern
```go
// Cada comando implementa o padrão
type Command interface {
    Execute(args []string) error
}
```

### 2. Factory Pattern
```go
// Para criação de clientes
func NewDockerClient() (*Client, error) {
    // Implementação
}
```

### 3. Builder Pattern (futuro)
```go
// Para construção de projetos LaTeX
type ProjectBuilder struct {
    title   string
    author  string
    doctype string
}
```

## Tratamento de Erros

### 1. Níveis de Erro
- **Critical**: Falha total da operação
- **Warning**: Operação continua com limitações
- **Info**: Informação para o usuário

### 2. Propagação
```go
if err := operation(); err != nil {
    return fmt.Errorf("[ERROR] Operação falhou: %w", err)
}
```

### 3. User-Friendly Messages
- Mensagens claras para o usuário
- Sugestões de solução quando possível
- Links para documentação relevante

## Testes

### 1. Estrutura
```
internal/commands/setup_test.go
internal/config/config_test.go
```

### 2. Padrões
- Testes unitários para cada package
- Mocks para dependências externas (Docker)
- Testes de integração para fluxos completos

### 3. Coverage
- Objetivo: >80% de cobertura
- Focar em lógica de negócio crítica
- Testes de edge cases

## Extensibilidade

### 1. Novos Comandos
```go
// 1. Criar comando em internal/commands/
var NewCmd = &cobra.Command{...}

// 2. Registrar em cmd/root.go
rootCmd.AddCommand(commands.NewCmd)
```

### 2. Novas Configurações
```go
// 1. Adicionar em pkg/types/types.go
type Config struct {
    NewField string `mapstructure:"new_field"`
}

// 2. Definir padrão em internal/config/config.go
viper.SetDefault("new_field", "default_value")
```

### 3. Nova Integração
```go
// 1. Criar package em internal/newservice/
// 2. Implementar interface consistente
// 3. Injetar nos comandos necessários
```

## Performance

### 1. Build Time
- Binário compilado: ~10-50ms startup
- vs. Bash script: ~200-500ms

### 2. Memory Usage
- Footprint mínimo via single binary
- Garbage collection otimizada

### 3. I/O Operations
- File watching eficiente com fsnotify
- Docker operations otimizadas

## Segurança

### 1. Input Validation
- Validação de todos os inputs do usuário
- Sanitização de paths e argumentos

### 2. Docker Security
- Operações Docker via API oficial
- Não execução de comandos shell arbitrários

### 3. File Operations
- Paths absolutos quando necessário
- Verificação de permissões

Esta arquitetura garante que a CLI seja mantível, extensível e performática.
