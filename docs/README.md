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

## 🏗️ Nova Arquitetura

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
├── docs/                 # Documentação
├── src/                  # Seus arquivos LaTeX
├── dist/                 # Arquivos compilados
└── tests/               # Testes
```

## 🚀 Início Rápido

### 1. Configuração Inicial

```bash
# Clone o repositório
git clone https://github.com/martinsmiguel/latex-template.git
cd latex-template

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

O script irá solicitar:
- **Título do documento**
- **Nome do autor**

## LaTeX Template CLI

Este projeto inclui uma **CLI (Command Line Interface) robusta** que centraliza todas as operações e segue as melhores práticas de desenvolvimento Bash. A CLI substitui e melhora os scripts individuais, oferecendo uma experiência mais consistente e segura.

### Características da CLI

- ✓ **Tratamento de Erros Rigoroso**: `set -euo pipefail` e validação completa
- ✓ **Modularização**: Funções pequenas e específicas para cada tarefa
- ✓ **Validação de Entrada**: Nunca confia em entrada não sanitizada
- ✓ **Mensagens Claras**: Output organizado e legível
- ✓ **Códigos de Saída**: Retorna códigos apropriados para automação
- ✓ **Documentação Integrada**: Help detalhado para todos os comandos
- ✓ **Autocompletion**: Suporte para Bash e Zsh
- ✓ **Configuração Flexível**: Arquivo de configuração opcional

### Comandos Principais

```bash
# Configuração inicial do projeto
./latex-cli setup

# Inicialização de documento (interativo ou não)
./latex-cli init
./latex-cli init --title "Minha Tese" --author "João Silva"

# Compilação com opções avançadas
./latex-cli build
./latex-cli build --engine xelatex --output-dir dist --clean

# Desenvolvimento e observação
./latex-cli watch          # Auto-compilação
./latex-cli dev            # Inicia ambiente
./latex-cli status         # Status completo

# Gerenciamento
./latex-cli clean --all    # Limpeza completa
./latex-cli logs --follow  # Logs em tempo real
./latex-cli shell          # Shell no container
./latex-cli update         # Atualiza ambiente
```

### Documentação Completa

Para informações detalhadas sobre todos os comandos, opções e exemplos de uso, consulte:

**[Documentação Completa da CLI](CLI.md)**

```bash
# Ou use a ajuda integrada
./latex-cli --help
./latex-cli <comando> --help
```

### 🔄 Retrocompatibilidade

Os scripts originais (`start.sh`, `scripts/`) ainda funcionam normalmente para compatibilidade com workflows existentes:

```bash
# Método tradicional (ainda funciona)
./start.sh
./scripts/init_project.sh
./scripts/compile.sh

# Novo método CLI (recomendado)
./latex-cli setup
./latex-cli init
./latex-cli build
```

E criará automaticamente:
- `tex/main.tex` personalizado
- Estrutura de capítulos em `tex/chapters/` (introdução, metodologia, resultados)
- `tex/preamble.tex` se não existir

## 📂 Estrutura do Projeto

```
latex-template/
├── latex-cli               # CLI principal (recomendado)
├── .latex-cli.conf         # Configuração opcional da CLI
├── CLI.md                  # Documentação completa da CLI
├── completions/            # Autocompletion para shells
│   ├── latex-cli.bash      #     Bash completion
│   └── _latex-cli          #     Zsh completion
├── .devcontainer/          # Configuração do ambiente Docker
│   ├── devcontainer.json   # Configurações do VS Code Dev Container
│   └── Dockerfile          # Imagem personalizada com TeX Live
├── .vscode/                # Configurações do VS Code
│   ├── settings.json       # Settings do LaTeX Workshop
│   └── tasks.json          # Tasks de compilação e limpeza
├── scripts/                # Scripts de automação (legacy)
│   ├── compile.sh          # Script de compilação manual
│   ├── clean.sh            # Script de limpeza de arquivos auxiliares
│   ├── init_project.sh     # Inicialização de novos documentos
│   └── latexmk-docker.sh   # Wrapper para latexmk no container
├── templates/              # Templates base
│   ├── main.tex.tpl        # Template principal do documento
│   ├── preamble.tex.tpl    # Template de configurações LaTeX
│   ├── settings.json.tpl   # Template de configurações do VS Code
│   └── references.bib.tpl  # Template para bibliografia
├── tex/                    # Seus arquivos LaTeX (criado automaticamente)
│   ├── main.tex            # Documento principal
│   ├── preamble.tex        # Configurações e pacotes LaTeX
│   ├── chapters/           # Capítulos do documento
│   └── references.bib      # Bibliografia
├── out/                    # Arquivos de saída (PDF, logs)
├── docker-compose.yml      # Configuração do Docker Compose
└── start.sh               # Script de configuração inicial
```

## 🛠️ Scripts Disponíveis

### `./start.sh`
Configura o ambiente inicial:
- Cria diretórios necessários
- Copia templates se não existirem
- Configura permissões dos scripts

### `./scripts/init_project.sh`
Inicializa um novo documento:
- Solicita título e autor
- Gera `main.tex` personalizado
- Cria estrutura de capítulos
- Copia template de preamble se necessário

### `./scripts/latexmk-docker.sh`
Wrapper para executar latexmk no container:
- Permite execução direta do latexmk via Docker
- Facilita integração com editores externos
- Passa todos os argumentos para o container

### `./scripts/compile.sh`
Compila o documento manualmente:
- Verifica dependências necessárias
- Usa `latexmk` para compilação otimizada
- Gera PDF em `out/main.pdf`
- Fornece feedback claro sobre erros

### `./scripts/clean.sh`
Limpa arquivos auxiliares:
- Remove logs e arquivos temporários
- Mantém apenas o PDF final
- Útil para debugging de problemas

## 📝 Como Usar

### Escrevendo seu Documento

1. **Documento Principal**: Edite `tex/main.tex`
2. **Capítulos**: Adicione conteúdo em `tex/chapters/`
3. **Bibliografia**: Adicione referências em `tex/references.bib`

### Compilação

A compilação acontece automaticamente quando você salva arquivos `.tex` no VS Code. Alternativamente:

```bash
# Compilação manual
./scripts/compile.sh
```

### Tarefas do VS Code

O projeto inclui tarefas pré-configuradas acessíveis via `Ctrl+Shift+P` → "Tasks: Run Task":

- **Compile LaTeX**: Executa `./scripts/compile.sh`
- **Clean LaTeX**: Executa `./scripts/clean.sh`

### Visualização

O PDF gerado ficará em `out/main.pdf` e pode ser visualizado:
- **No VS Code**: Aba automática do LaTeX Workshop
- **Externamente**: Qualquer visualizador de PDF

## 🔧 Configurações

### VS Code Settings

O template inclui configurações otimizadas em `.vscode/settings.json`:

```json
{
  "latex-workshop.latex.autoBuild.run": "onSave",
  "latex-workshop.latex.outDir": "./out",
  "latex-workshop.view.pdf.viewer": "tab",
  "files.autoSave": "afterDelay",
  "latex-workshop.latex.tools": [
    {
      "name": "latexmk-docker",
      "command": "docker",
      "args": [
        "exec", "latex-env", "latexmk",
        "-pdf", "-f", "-interaction=nonstopmode",
        "-synctex=1", "-outdir=./out", "tex/main.tex"
      ]
    }
  ],
  "latex-workshop.latex.recipes": [
    {
      "name": "latexmk (Docker)",
      "tools": ["latexmk-docker"]
    }
  ],
  "latex-workshop.latex.recipe.default": "latexmk (Docker)"
}
```

### Pacotes LaTeX Incluídos

O ambiente Docker inclui:
- **TeX Live completo**
- **latexmk** para compilação
- **biber** para bibliografia
- **python3-pygments** para syntax highlighting
- Pacotes adicionais: `enumitem`, `fancyhdr`, `hyperref`, `xcolor`

## 🐳 Ambiente Docker

### Características do Container

- **Base**: `texlive/texlive:latest`
- **Usuário**: `latexuser` (não-root para segurança, UID/GID 1001)
- **Pacotes**: TeX Live completo + dependências essenciais
- **Health Check**: Monitora disponibilidade do `pdflatex`
- **Volume persistente**: Cache do TeX Live para melhor performance

### Comandos Docker Úteis

```bash
# Iniciar o ambiente
docker-compose up -d

# Verificar status do container
docker ps

# Acessar terminal do container
docker exec -it latex-env bash

# Ver logs do container
docker logs latex-env

# Parar o ambiente
docker-compose down
```

## 📋 Dependências

### Locais (Host)
- **Docker** e **Docker Compose**
- **VS Code** com extensão **Dev Containers**

### Container (Automáticas)
- **TeX Live** (distribuição completa)
- **LaTeX Workshop** (extensão VS Code)
- **latexmk** e **biber** para compilação e bibliografia
- **python3-pygments** para syntax highlighting (pacote minted)
- **Pacotes adicionais**: `enumitem`, `fancyhdr`, `hyperref`, `xcolor`

## 🚨 Solução de Problemas

### Erro de Permissão nos Scripts

Se você receber "permission denied" ao executar scripts:

```bash
# Dê permissão de execução a todos os scripts
chmod +x start.sh
chmod +x scripts/*.sh

# Ou use o comando individual conforme necessário
chmod +x ./scripts/init_project.sh
chmod +x ./scripts/compile.sh
chmod +x ./scripts/clean.sh
chmod +x ./scripts/latexmk-docker.sh
```

### Pacotes LaTeX Ausentes

O script de compilação tenta instalar automaticamente pacotes ausentes. Se falhar:

```bash
# Entre no container
docker exec -it latex-env bash

# Instale manualmente
tlmgr install nome-do-pacote
```

### Erro de Compilação

1. Verifique os logs em `out/logs/`
2. Use compilação manual: `./scripts/compile.sh`
3. Verifique sintaxe LaTeX no VS Code

### Container não Inicia

```bash
# Reconstruir o container
docker-compose down
docker-compose build --no-cache
docker-compose up -d
```

## 🤝 Contribuindo

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 📞 Suporte

- 🐛 **Issues**: [GitHub Issues](https://github.com/martinsmiguel/latex-template/issues)
- 📧 **Email**: miguelrjmartins.dev@gmail.com
- 📚 **Documentação**: Consulte este README para informações completas

---

**Feito com ❤️ para simplificar a escrita acadêmica em LaTeX**
