# LaTeX Docker Environment

> Um ambiente Docker moderno e completo para criação de documentos LaTeX com CLI robusta e multiplataforma.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Docker](https://img.shields.io/badge/Docker-20.10%2B-blue)](https://docs.docker.com/get-docker/)
[![Go](https://img.shields.io/badge/Go-1.19%2B-blue)](https://golang.org/)

## ✨ Características

- 🐳 **Ambiente Dockerizado**: Desenvolvimento isolado com todas as dependências LaTeX
- ⚡ **CLI Moderna**: Interface em Go com suporte nativo ao Windows
- 🔄 **Compilação Automática**: Build otimizado com `latexmk` e watch mode
- 🛠️ **VS Code Integrado**: Configurações otimizadas para desenvolvimento LaTeX
- 📦 **Setup Automatizado**: Instalação e configuração em um comando
- 🌐 **Multiplataforma**: Windows, macOS e Linux

## 🚀 Início Rápido

### Pré-requisitos
- [Docker](https://docs.docker.com/get-docker/) 20.10+
- [Docker Compose](https://docs.docker.com/compose/install/) 2.0+
- Git

### Instalação

```bash
# 1. Clone o repositório
git clone https://github.com/martinsmiguel/latex-docker-env.git
cd latex-docker-env

# 2. Configure o ambiente
./bin/ltx setup

# 3. Crie seu documento
./bin/ltx init --title "Meu Documento" --author "Seu Nome"

# 4. Compile e desenvolva
./bin/ltx build          # Compilação única
./bin/ltx watch          # Modo de observação
```

### Comandos Principais

| Comando | Descrição |
|---------|-----------|
| `ltx setup` | Configura o ambiente inicial |
| `ltx init` | Cria um novo documento LaTeX |
| `ltx build` | Compila o documento |
| `ltx watch` | Compilação automática (modo desenvolvimento) |
| `ltx clean` | Remove arquivos temporários |
| `ltx status` | Status do ambiente Docker |
| `ltx backup` | Cria backup do trabalho atual |
| `ltx reset` | Reseta completamente o ambiente |

## � Workflows Avançados

### Backup e Restauração
```bash
# Criar backup automático (com timestamp)
./bin/ltx backup

# Backup com nome específico
./bin/ltx backup --name "versao-final"

# Backup em local customizado
./bin/ltx backup --custom "../meus-documentos"
```

### Reset do Ambiente
```bash
# Reset com confirmação (recomendado)
./bin/ltx reset

# Reset forçado (sem confirmação)
./bin/ltx reset --force

# Workflow completo: backup + reset + novo projeto
./bin/ltx backup --name "projeto-anterior"
./bin/ltx reset --force
./bin/ltx init --title "Novo Projeto" --author "Seu Nome"
```

## �📖 Documentação

- **[📋 Guia de Instalação](docs/installation.md)** - Instruções detalhadas por SO
- **[🛠️ CLI Reference](docs/cli-reference.md)** - Documentação completa da CLI
- **[❓ FAQ](docs/faq.md)** - Perguntas frequentes e solução de problemas
- **[🔄 Migração](docs/migration.md)** - Migração de versões anteriores
- **[� Contribuindo](docs/contributing.md)** - Como contribuir para o projeto

## 📁 Estrutura do Projeto

```
latex-docker-env/
├── bin/                    # Executáveis da CLI
│   ├── ltx                 # CLI moderna (Go) - recomendada
│   └── latex-cli           # CLI legada (Bash) - compatibilidade
├── cli/                    # Código fonte da CLI Go
├── config/                 # Configurações e templates
├── docs/                   # Documentação
├── src/                    # Seus arquivos LaTeX (criado após init)
├── dist/                   # PDFs compilados (criado após build)
└── LICENSE                 # Licença MIT
```

## 🤝 Contribuindo

Contribuições são bem-vindas! Veja [CONTRIBUTING.md](docs/contributing.md) para detalhes.

## 📄 Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

---

> **💡 Dica**: Use `./bin/ltx --help` para ver todos os comandos disponíveis