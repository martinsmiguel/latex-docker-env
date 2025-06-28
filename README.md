# 📄 LaTeX Template v2.0

Um template moderno e dockerizado para escrita de documentos LaTeX com arquitetura otimizada e CLI robusta.

## ✨ Características

- 🐳 **Ambiente Dockerizado**: Desenvolvimento isolado com todas as dependências LaTeX incluídas
- 🔄 **Compilação Automática**: Auto-build otimizado com `latexmk`
- 📁 **Arquitetura Moderna**: Estrutura organizada e modular
- 🛠️ **CLI Robusta**: Interface de linha de comando seguindo melhores práticas
- 🔧 **Configuração VS Code**: Settings otimizados para desenvolvimento LaTeX
- 📦 **Gerenciamento Automatizado**: Instalação e configuração simplificadas
- ⚡ **Autocompletion**: Suporte completo para Bash e Zsh
- 🌍 **Multi-plataforma**: Funciona em Windows, macOS e Linux

## � Pré-requisitos

### 🖥️ Todos os Sistemas

- **Docker**: Versão 20.10+ ([Instalar Docker](https://docs.docker.com/get-docker/))
- **Docker Compose**: Versão 2.0+ (incluído no Docker Desktop)
- **Git**: Para clonar o repositório

### 🪟 Windows

- **Windows 10/11** com WSL2 habilitado
- **Docker Desktop** com integração WSL2
- **Terminal** recomendado: Windows Terminal ou WSL2

### 🍎 macOS

- **macOS 10.15+** (Catalina ou superior)
- **Docker Desktop for Mac**

### 🐧 Linux

- **Docker Engine** + **Docker Compose**
- **Bash 4.0+** ou **Zsh**

## 🚀 Início Rápido

### 1. Clone e Configure

```bash
# Clone o repositório
git clone https://github.com/martinsmiguel/latex-template.git
cd latex-template

# Configure permissões (Linux/macOS)
chmod +x bin/latex-cli

# Configure o ambiente
./bin/latex-cli setup
```

### 2. Inicialize seu Documento

```bash
# Inicialização interativa
./bin/latex-cli init

# Ou modo direto
./bin/latex-cli init --title "Minha Tese" --author "João Silva"
```

### 3. Compile e Desenvolva

```bash
# Compilação única
./bin/latex-cli build

# Modo de observação (auto-compilação)
./bin/latex-cli watch

# Verificar status do ambiente
./bin/latex-cli status
```

## 📚 Documentação

Para informações detalhadas, consulte:

- **[🛠️ Guia de Instalação](docs/INSTALLATION.md)** - Instruções detalhadas por sistema operacional
- **[❓ Perguntas Frequentes (FAQ)](docs/FAQ.md)** - Soluções para problemas comuns
- **[📖 Documentação Completa da CLI](docs/CLI.md)** - Referência de todos os comandos
- **[📋 Documentação do Projeto](docs/README.md)** - Arquitetura e desenvolvimento
- **[🔄 Migração de Versões](docs/MIGRATION.md)** - Guia de atualização

```bash
# Ou use a ajuda integrada
./bin/latex-cli --help
./bin/latex-cli <comando> --help
```

## 💡 Exemplos de Uso

### 📖 Documento Simples

```bash
# Criar um artigo científico
./bin/latex-cli init --title "Meu Artigo" --author "Seu Nome"
./bin/latex-cli build

# Arquivo gerado: dist/main.pdf
```

### 📚 Tese/Dissertação

```bash
# Modo interativo para configuração completa
./bin/latex-cli init

# Compilação com observação automática
./bin/latex-cli watch
# Agora edite os arquivos em src/ e veja a compilação automática
```

### 🔄 Fluxo de Desenvolvimento

```bash
# 1. Configurar projeto
./bin/latex-cli setup

# 2. Inicializar documento
./bin/latex-cli init --title "Minha Pesquisa"

# 3. Desenvolvimento com auto-compilação
./bin/latex-cli watch &

# 4. Editar arquivos em src/
# 5. PDF atualizado automaticamente em dist/

# 6. Limpeza final
./bin/latex-cli clean
./bin/latex-cli build
```

## 🛠️ Comandos Principais

- `setup` - Configura o ambiente inicial
- `init` - Inicializa um novo documento
- `build` - Compila o documento
- `watch` - Modo de observação com auto-compilação
- `clean` - Limpa arquivos temporários
- `status` - Mostra status do ambiente
- `shell` - Acessa shell do container
- `logs` - Visualiza logs do container

## 🎯 Configuração por Sistema Operacional

### 🪟 Windows

```powershell
# No PowerShell ou CMD
git clone https://github.com/martinsmiguel/latex-template.git
cd latex-template

# Execute o setup (não precisa chmod no Windows)
.\bin\latex-cli setup

# Para uso com WSL2, use o terminal WSL
wsl
cd /mnt/c/caminho/para/latex-template
chmod +x bin/latex-cli
./bin/latex-cli setup
```

### 🍎 macOS

```bash
# No Terminal
git clone https://github.com/martinsmiguel/latex-template.git
cd latex-template

# Configure permissões e execute setup
chmod +x bin/latex-cli
./bin/latex-cli setup
```

### 🐧 Linux

```bash
# Terminal
git clone https://github.com/martinsmiguel/latex-template.git
cd latex-template

# Configure permissões e execute setup
chmod +x bin/latex-cli
./bin/latex-cli setup

# Opcional: Adicionar ao PATH
sudo ln -sf "$(pwd)/bin/latex-cli" /usr/local/bin/latex-cli
```

## ⚡ Configuração de Autocompletion

### Bash (Linux/macOS/WSL)

```bash
# Adicionar ao ~/.bashrc
echo "source $(pwd)/config/completions/latex-cli.bash" >> ~/.bashrc
source ~/.bashrc
```

### Zsh (macOS/Linux)

```bash
# Adicionar ao ~/.zshrc
echo "fpath=($(pwd)/config/completions \$fpath)" >> ~/.zshrc
echo "autoload -U compinit && compinit" >> ~/.zshrc
source ~/.zshrc
```

## 🔧 Troubleshooting

### ❌ Problemas Comuns

#### "Permission denied" (Linux/macOS)
```bash
chmod +x bin/latex-cli
```

#### Docker não encontrado
```bash
# Verificar se Docker está rodando
docker --version
docker compose version

# Se não estiver instalado, visite: https://docs.docker.com/get-docker/
```

#### WSL2 no Windows
```powershell
# Habilitar WSL2
wsl --install
wsl --set-default-version 2

# Usar dentro do WSL2
wsl
cd /mnt/c/seu/projeto
```

#### Container não inicia
```bash
# Verificar logs
./bin/latex-cli logs

# Recriar container
./bin/latex-cli clean
./bin/latex-cli setup
```

### 🆘 Obter Ajuda

```bash
# Ajuda geral
./bin/latex-cli --help

# Ajuda específica de comando
./bin/latex-cli build --help
./bin/latex-cli init --help
```

## 📂 Estrutura do Projeto

```
latex-template/
├── bin/                    # Executáveis
│   └── latex-cli          # CLI principal
├── lib/                   # Bibliotecas da CLI
│   ├── commands/          # Comandos individuais
│   ├── core/             # Funcionalidades centrais
│   └── utils/            # Utilitários
├── config/               # Configurações
│   ├── docker/           # Configurações Docker
│   ├── templates/        # Templates LaTeX
│   ├── vscode/          # Configurações VS Code
│   ├── completions/     # Autocompletion
│   └── latex-cli.conf   # Configuração principal
├── src/                  # Seus arquivos LaTeX
│   ├── main.tex         # Documento principal
│   ├── preamble.tex     # Configurações e pacotes
│   ├── references.bib   # Bibliografia
│   └── chapters/        # Capítulos do documento
├── dist/                 # Arquivos compilados (PDF)
└── docs/                 # Documentação
```

## 🤝 Contribuindo

Contribuições são bem-vindas! Por favor:

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 🐛 Reportar Problemas

Encontrou um bug? Abra uma [issue](https://github.com/martinsmiguel/latex-template/issues) com:

- Descrição do problema
- Sistema operacional
- Versão do Docker
- Logs relevantes (`./bin/latex-cli logs`)

## 📞 Suporte

- � **Documentação**: [docs/](docs/)
- 🐛 **Issues**: [GitHub Issues](https://github.com/martinsmiguel/latex-template/issues)
- 💬 **Discussões**: [GitHub Discussions](https://github.com/martinsmiguel/latex-template/discussions)

## �📄 Licença

Este projeto está licenciado sob a MIT License - veja o arquivo [LICENSE](docs/LICENSE) para detalhes.

---

**✨ Desenvolvido com ❤️ para a comunidade acadêmica**