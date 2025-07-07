# Documentação da CLI ltx

Esta seção contém toda a documentação relacionada à CLI moderna `ltx` escrita em Go.

## Índice

### Para Usuários
- [**Guia de Uso**](usage-guide.md) - Como usar a CLI ltx no dia a dia
- [**Comandos**](commands.md) - Referência completa de todos os comandos
- [**Configuração**](configuration.md) - Como configurar a CLI

### Para Desenvolvedores
- [**Guia de Desenvolvimento**](development-guide.md) - Como desenvolver e contribuir
- [**Padrões de Saída**](output-standards.md) - Convenções de output da CLI
- [**Arquitetura**](architecture.md) - Estrutura interna do código
- [**Testes**](testing.md) - Como escrever e executar testes

### Referência Técnica
- [**API Docker**](docker-integration.md) - Integração com Docker
- [**File Watching**](file-watching.md) - Sistema de monitoramento de arquivos
- [**Build System**](build-system.md) - Sistema de compilação multiplataforma

## Estrutura da CLI

```
cli/
├── cmd/                    # Comandos Cobra
├── internal/               # Código interno
│   ├── commands/          # Implementação dos comandos
│   ├── config/            # Gerenciamento de configuração
│   ├── docker/            # Cliente Docker
│   └── utils/             # Utilitários
├── pkg/                   # Packages públicos
└── main.go               # Entry point
```

## Links Rápidos

- [Começar a desenvolver](development-guide.md#setup-inicial)
- [Executar testes](development-guide.md#testes)
- [Build multiplataforma](development-guide.md#build-e-release)
- [Padrões de código](output-standards.md)

## Status dos Comandos

| Comando | Status | Descrição |
|---------|--------|-----------|
| `setup` | [OK] | Configuração do ambiente |
| `init`  | [WIP] | Criação de documentos |
| `build` | [WIP] | Compilação LaTeX |
| `watch` | [WIP] | File watching |
| `status`| [WIP] | Status do ambiente |
| `clean` | [WIP] | Limpeza de arquivos |
| `shell` | [WIP] | Acesso ao container |
| `logs`  | [WIP] | Visualização de logs |

## Contribuindo

Para contribuir com a CLI:

1. Leia o [Guia de Desenvolvimento](development-guide.md)
2. Siga os [Padrões de Saída](output-standards.md)
3. Escreva testes para novas funcionalidades
4. Execute `make test` antes de fazer PR
