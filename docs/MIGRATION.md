# 🔄 Migração para Arquitetura v2.0

Este documento descreve a migração da arquitetura legacy para a nova estrutura modular do LaTeX Template.

## ✅ Mudanças Implementadas

### 🏗️ Nova Estrutura de Diretórios

```
ANTES (v1.x):
latex-template/
├── latex-cli              # Script monolítico
├── scripts/               # Scripts separados
├── templates/             # Templates
├── .devcontainer/        # Config Docker
├── .vscode/              # Config VS Code
├── completions/          # Autocompletion
└── tex/                  # Arquivos LaTeX

DEPOIS (v2.0):
latex-template/
├── bin/                  # Executáveis
│   └── latex-cli        # CLI principal
├── lib/                 # Bibliotecas modulares
│   ├── commands/        # Comandos individuais
│   ├── core/           # Funcionalidades centrais
│   └── utils/          # Utilitários
├── config/             # Todas as configurações
│   ├── docker/         # Docker configs
│   ├── templates/      # Templates LaTeX
│   ├── vscode/         # VS Code configs
│   ├── completions/    # Autocompletion
│   └── latex-cli.conf  # Config principal
├── docs/               # Documentação
├── src/                # Arquivos LaTeX do usuário
├── dist/               # Arquivos compilados
└── tests/              # Testes
```

### 🚀 CLI Modularizada

#### Arquivos Core Criados:
- `lib/core/config.sh` - Gerenciamento de configurações
- `lib/core/logger.sh` - Sistema de logging robusto
- `lib/core/docker.sh` - Operações Docker
- `lib/utils/helpers.sh` - Funções utilitárias

#### Comandos Implementados:
- `lib/commands/setup.sh` - Configuração inicial
- `lib/commands/init.sh` - Inicialização de projetos
- `lib/commands/build.sh` - Compilação
- `lib/commands/watch.sh` - Auto-compilação
- `lib/commands/clean.sh` - Limpeza
- `lib/commands/status.sh` - Status do ambiente
- `lib/commands/shell.sh` - Shell no container
- `lib/commands/logs.sh` - Visualização de logs
- `lib/commands/update.sh` - Atualizações

### 📝 Melhorias na Documentação

- **README.md**: Simplificado na raiz, aponta para documentação completa
- **docs/README.md**: Documentação completa da nova arquitetura
- **docs/CLI.md**: Documentação detalhada dos comandos
- **docs/LICENSE**: Licença movida para docs/

### ⚙️ Configurações Centralizadas

- `config/latex-cli.conf` - Configuração principal
- `config/docker/` - Todas as configs Docker
- `config/templates/` - Templates LaTeX organizados
- `config/vscode/` - Configurações VS Code

## 🗑️ Arquivos Removidos

### Scripts Legacy (Removidos):
- `start.sh` - Substituído por `latex-cli setup`
- `scripts/init_project.sh` - Substituído por `latex-cli init`
- `scripts/compile.sh` - Substituído por `latex-cli build`
- `scripts/clean.sh` - Substituído por `latex-cli clean`
- `scripts/latexmk-docker.sh` - Funcionalidade integrada na CLI

### Arquivos Movidos:
- `latex-cli` → `bin/latex-cli` (reescrito modular)
- `.devcontainer/` → `config/docker/devcontainer/`
- `docker-compose.yml` → `config/docker/docker-compose.yml`
- `templates/` → `config/templates/`
- `.vscode/` → `config/vscode/vscode/`
- `completions/` → `config/completions/`
- Documentação → `docs/`

## 🔧 Comandos de Migração

Para usuários existentes, a migração é automática:

```bash
# Configure o novo ambiente
./bin/latex-cli setup

# Migre projeto existente (se houver tex/ antigo)
# A CLI detectará automaticamente e oferecerá migração

# Verifique o status
./bin/latex-cli status
```

## 📊 Comparação de Comandos

| Comando Antigo | Comando Novo | Melhoria |
|----------------|--------------|----------|
| `./latex-cli setup` | `./bin/latex-cli setup` | ✅ Modular |
| `./latex-cli init` | `./bin/latex-cli init` | ✅ Templates múltiplos |
| `./latex-cli build` | `./bin/latex-cli build` | ✅ Engines configuráveis |
| `./latex-cli watch` | `./bin/latex-cli watch` | ✅ Melhor handling |
| `./latex-cli clean` | `./bin/latex-cli clean` | ✅ Opções granulares |
| `./latex-cli status` | `./bin/latex-cli status` | ✅ Info detalhada |
| - | `./bin/latex-cli shell` | 🆕 Novo comando |
| - | `./bin/latex-cli logs` | 🆕 Novo comando |
| - | `./bin/latex-cli update` | 🆕 Novo comando |

## ✨ Benefícios da Nova Arquitetura

### 🔧 Manutenibilidade
- **Modularidade**: Cada comando em arquivo separado
- **Separation of Concerns**: Core, utils e commands separados
- **Reutilização**: Funções compartilhadas entre comandos

### 🚀 Performance
- **Loading sob demanda**: Comandos carregados apenas quando necessários
- **Cache de configuração**: Configurações carregadas uma vez
- **Logs otimizados**: Sistema de logging com níveis

### 🛡️ Robustez
- **Tratamento de erros**: Error handling consistente
- **Validação**: Input validation em todas as funções
- **Rollback**: Operações com possibilidade de rollback

### 📚 Usabilidade
- **Help contextual**: `--help` para cada comando
- **Autocompletion**: Suporte melhorado
- **Status detalhado**: Informações completas sobre o ambiente

## 🔮 Próximos Passos

1. **Testes**: Implementar testes unitários e de integração
2. **CI/CD**: Pipeline de testes automatizados
3. **Plugins**: Sistema de plugins para extensibilidade
4. **Templates**: Mais templates para diferentes tipos de documento

## 💡 Notas de Desenvolvimento

- **Compatibilidade**: Mantém compatibilidade com projetos existentes
- **Configuração**: Suporte a múltiplos arquivos de configuração
- **Logging**: Logs estruturados para debugging
- **Docker**: Melhor gerenciamento do ciclo de vida do container

---

**Migração concluída**: ✅
**Status**: Pronto para uso
**Versão**: 2.0.0
