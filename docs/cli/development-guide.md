# CLI Development Guide

## Estrutura do Projeto

```
cli/
├── cmd/                    # Comandos Cobra
│   └── root.go            # Comando raiz e configuração
├── internal/              # Código interno da aplicação
│   ├── commands/          # Implementação dos comandos
│   │   ├── setup.go       # Comando setup (implementado)
│   │   ├── basic.go       # Outros comandos (placeholders)
│   │   └── setup_test.go  # Testes do setup
│   ├── config/            # Gerenciamento de configuração
│   │   ├── config.go      # Configurações principais
│   │   └── config_test.go # Testes de configuração
│   ├── docker/            # Cliente Docker
│   │   └── client.go      # Interface com Docker API
│   └── utils/             # Utilitários (vazio por enquanto)
├── pkg/                   # Packages públicos
│   └── types/             # Tipos compartilhados
│       └── types.go       # Structs de configuração
├── main.go                # Entry point
├── go.mod                 # Dependências Go
├── go.sum                 # Lock file das dependências
├── Makefile               # Scripts de build e desenvolvimento
└── .goreleaser.yaml       # Configuração de release
```

## Comandos Implementados

### [OK] Comando `setup`
- Verifica estrutura do projeto
- Testa conexão Docker
- Baixa imagem LaTeX
- Cria diretórios necessários
- Configura VS Code (básico)

### [WIP] Comandos Pendentes
- `init` - Criação de documentos LaTeX
- `build` - Compilação via Docker
- `watch` - File watching com fsnotify
- `status` - Status do container Docker
- `clean` - Limpeza de arquivos temporários
- `shell` - Acesso ao container
- `logs` - Logs do container

## Como Desenvolver

### Setup Inicial
```bash
cd cli/
go mod download
make setup-dev  # Instala ferramentas de desenvolvimento
```

### Desenvolvimento
```bash
make dev        # Roda CLI em modo desenvolvimento
make test       # Executa testes
make lint       # Linting do código
make build      # Build local
```

### Testes
```bash
make test              # Todos os testes
make test-coverage     # Testes com coverage
```

### Build e Release
```bash
make build-all         # Build para todas as plataformas
make install           # Instala no ../bin/
make release-snapshot  # Teste de release
make release           # Release real (requer tag Git)
```

## Próximos Passos

### Prioridade Alta
1. **Comando `init`**: Criar sistema de templates
2. **Comando `build`**: Integração com docker-compose
3. **Comando `watch`**: File watching com fsnotify

### Prioridade Média  
4. **Comando `status`**: Status detalhado do container
5. **Melhorar testes**: Mais cobertura de teste
6. **Documentação**: Adicionar mais docs

### Prioridade Baixa
7. **Comando `shell`**: Acesso interativo ao container
8. **Comando `logs`**: Streaming de logs
9. **Configuração avançada**: YAML config, profiles, etc.

## Arquitetura de Comandos

Cada comando segue o padrão:

```go
var CommandCmd = &cobra.Command{
    Use:   "command",
    Short: "Descrição curta",
    Long:  "Descrição longa...",
    RunE:  runCommand,
}

func runCommand(cmd *cobra.Command, args []string) error {
    // 1. Validar argumentos/flags
    // 2. Carregar configuração
    // 3. Executar lógica do comando
    // 4. Retornar erro ou nil
}
```

## Configuração

A configuração é gerenciada via Viper e suporta:
- Arquivo `config/latex-cli.conf`
- Variáveis de ambiente
- Flags de linha de comando

## Dependências Principais

- **Cobra**: CLI framework
- **Viper**: Gerenciamento de configuração  
- **Docker Client**: Integração com Docker
- **fsnotify**: File watching (futuro)

## Pipeline CI/CD

O GitHub Actions foi configurado para:
- [OK] Testes automáticos
- [OK] Linting
- [OK] Build multiplataforma
- [OK] Release automático com GoReleaser

## Como Contribuir

1. Implemente comando em `internal/commands/`
2. Adicione testes em `*_test.go`
3. Atualize `cmd/root.go` para registrar comando
4. Teste com `make test` e `make build`
5. Commit e push para ativar CI/CD
