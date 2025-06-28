# Testes do LaTeX Template CLI

Este diretório contém os testes automatizados para a CLI do LaTeX Template.

## Estrutura

```
tests/
├── README.md              # Este arquivo
├── run_tests.sh           # Script principal para executar todos os testes
├── unit/                  # Testes unitários
│   ├── test_cli_functions.bats    # Testes das funções da CLI
│   ├── test_logging.bats          # Testes do sistema de logging
│   └── test_validation.bats       # Testes de validação
├── integration/           # Testes de integração
│   ├── test_full_workflow.bats    # Teste do workflow completo
│   └── test_docker_integration.bats # Testes de integração com Docker
└── fixtures/              # Dados de teste
    ├── sample_project/     # Projeto de exemplo para testes
    └── test_configs/       # Configurações de teste
```

## Pré-requisitos

Para executar os testes, você precisa instalar o BATS (Bash Automated Testing System):

### macOS
```bash
brew install bats-core
```

### Ubuntu/Debian
```bash
sudo apt-get install bats
```

### Instalação manual
```bash
git clone https://github.com/bats-core/bats-core.git
cd bats-core
sudo ./install.sh /usr/local
```

## Executando os Testes

### Todos os testes
```bash
./tests/run_tests.sh
```

### Testes unitários apenas
```bash
bats tests/unit/*.bats
```

### Testes de integração apenas
```bash
bats tests/integration/*.bats
```

### Teste específico
```bash
bats tests/unit/test_cli_functions.bats
```

## Escrevendo Novos Testes

### Estrutura básica de um teste BATS

```bash
#!/usr/bin/env bats

# Setup executado antes de cada teste
setup() {
    # Configurações iniciais
    export LATEX_CLI_TEST_MODE="true"
}

# Teardown executado após cada teste
teardown() {
    # Limpeza
    unset LATEX_CLI_TEST_MODE
}

# Teste individual
@test "descrição do teste" {
    # Arrange
    local input="valor_teste"

    # Act
    run ./latex-cli comando --option "$input"

    # Assert
    [ "$status" -eq 0 ]
    [[ "$output" =~ "resultado esperado" ]]
}
```

### Convenções

1. **Nomes de arquivos**: `test_<componente>.bats`
2. **Nomes de testes**: Descrições claras do que está sendo testado
3. **Isolamento**: Cada teste deve ser independente
4. **Limpeza**: Use `teardown` para limpar recursos
5. **Mocking**: Use variáveis de ambiente para simular condições

## CI/CD

Os testes são executados automaticamente em:
- Pull Requests
- Push para a branch main
- Releases

Ver `.github/workflows/test.yml` para detalhes da configuração.

## Debugging

Para executar testes com output detalhado:

```bash
bats --tap tests/unit/test_cli_functions.bats
```

Para testes individuais com debug:

```bash
BATS_DEBUG=1 bats tests/unit/test_cli_functions.bats
```
