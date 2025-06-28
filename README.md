# LaTeX Template v2.0

Um template moderno e dockerizado para a criação de documentos LaTeX, com uma arquitetura otimizada e uma CLI robusta.

## Principais Características

- **Ambiente Dockerizado**: Desenvolvimento isolado com todas as dependências LaTeX incluídas.
- **Compilação Automática**: Processo de build otimizado com `latexmk`.
- **Arquitetura Moderna**: Estrutura de projeto organizada e modular.
- **CLI Robusta**: Interface de linha de comando intuitiva, seguindo as melhores práticas.
- **Configuração para VS Code**: Configurações otimizadas para o desenvolvimento em LaTeX.
- **Gerenciamento Automatizado**: Instalação e configuração do ambiente simplificadas.
- **Autocompletion**: Suporte completo para Bash e Zsh.
- **Multiplataforma**: Compatível com Windows, macOS e Linux.

## Pré-requisitos

### Todos os Sistemas

- **Docker**: Versão 20.10+ ([Instalar Docker](https://docs.docker.com/get-docker/))
- **Docker Compose**: Versão 2.0+ (geralmente incluído no Docker Desktop)
- **Git**: Para clonar o repositório.

### Windows

- **Windows 10/11** com WSL2 habilitado.
- **Docker Desktop** com integração WSL2 ativa.
- **Terminal recomendado**: Windows Terminal ou um terminal WSL2.

### macOS

- **macOS 10.15+** (Catalina ou superior).
- **Docker Desktop for Mac**.

### Linux

- **Docker Engine** e **Docker Compose**.
- **Bash 4.0+** ou **Zsh**.

## Guia de Início Rápido

### 1. Clone e Configure o Projeto

```bash
# Clone o repositório
git clone https://github.com/martinsmiguel/latex-template.git
cd latex-template

# Conceda permissão de execução ao script (Linux/macOS)
chmod +x bin/latex-cli

# Execute o script de configuração do ambiente
./bin/latex-cli setup
```

### 2. Inicialize seu Documento

```bash
# Use o modo interativo para configurar seu documento
./bin/latex-cli init

# Ou forneça os dados diretamente via argumentos
./bin/latex-cli init --title "Minha Tese" --author "João Silva"
```

### 3. Compile e Desenvolva

```bash
# Realize uma compilação única do documento
./bin/latex-cli build

# Ative o modo de observação para compilação automática a cada alteração
./bin/latex-cli watch

# Verifique o status do ambiente de desenvolvimento
./bin/latex-cli status
```

## Documentação

Para informações mais detalhadas, consulte os seguintes guias:

- **[Guia de Instalação](docs/INSTALLATION.md)**: Instruções detalhadas por sistema operacional.
- **[Perguntas Frequentes (FAQ)](docs/FAQ.md)**: Soluções para problemas comuns.
- **[Documentação da CLI](docs/CLI.md)**: Referência completa de todos os comandos.
- **[Documentação do Projeto](docs/README.md)**: Detalhes sobre a arquitetura e o desenvolvimento.
- **[Guia de Migração](docs/MIGRATION.md)**: Instruções para atualizar de versões anteriores.

```bash
# Você também pode usar a ajuda integrada da CLI
./bin/latex-cli --help
./bin/latex-cli <comando> --help
```

## Exemplos de Uso

### Documento Simples

```bash
# Crie um novo artigo científico
./bin/latex-cli init --title "Meu Artigo" --author "Seu Nome"
./bin/latex-cli build

# O arquivo final estará em: dist/main.pdf
```

### Tese ou Dissertação

```bash
# Utilize o modo interativo para uma configuração completa
./bin/latex-cli init

# Compile em modo de observação para um fluxo de trabalho contínuo
./bin/latex-cli watch
# Edite os arquivos em src/ e o PDF será atualizado automaticamente.
```

### Fluxo de Desenvolvimento Recomendado

```bash
# 1. Configure o projeto
./bin/latex-cli setup

# 2. Inicialize o seu documento
./bin/latex-cli init --title "Minha Pesquisa"

# 3. Inicie o desenvolvimento com compilação automática em background
./bin/latex-cli watch &

# 4. Edite os arquivos-fonte em src/
# 5. O PDF será atualizado automaticamente em dist/

# 6. Para a versão final, limpe os arquivos temporários e compile
./bin/latex-cli clean
./bin/latex-cli build
```

## Comandos Principais

- `setup`: Configura o ambiente de desenvolvimento inicial.
- `init`: Inicializa um novo documento LaTeX.
- `build`: Compila o documento para gerar o PDF.
- `watch`: Ativa o modo de observação para compilação automática.
- `clean`: Remove os arquivos temporários e de compilação.
- `status`: Exibe o status do ambiente Docker.
- `shell`: Acessa o shell do container Docker.
- `logs`: Mostra os logs do container em tempo real.

## Configuração por Sistema Operacional

### Windows

```powershell
# No PowerShell ou CMD
git clone https://github.com/martinsmiguel/latex-template.git
cd latex-template

# Execute o setup (não é necessário chmod no Windows)
.\bin\latex-cli setup

# Para uso com WSL2, execute os comandos no terminal do WSL
wsl
cd /mnt/c/caminho/para/latex-template
chmod +x bin/latex-cli
./bin/latex-cli setup
```

### macOS

```bash
# No Terminal
git clone https://github.com/martinsmiguel/latex-template.git
cd latex-template

# Configure as permissões e execute o setup
chmod +x bin/latex-cli
./bin/latex-cli setup
```

### Linux

```bash
# No seu terminal
git clone https://github.com/martinsmiguel/latex-template.git
cd latex-template

# Configure as permissões e execute o setup
chmod +x bin/latex-cli
./bin/latex-cli setup

# Opcional: Adicione a CLI ao seu PATH para acesso global
sudo ln -sf "$(pwd)/bin/latex-cli" /usr/local/bin/latex-cli
```

## Configuração do Autocompletion

### Bash (Linux/macOS/WSL)

```bash
# Adicione a seguinte linha ao seu ~/.bashrc
echo "source $(pwd)/config/completions/latex-cli.bash" >> ~/.bashrc
source ~/.bashrc
```

### Zsh (macOS/Linux)

```bash
# Adicione as seguintes linhas ao seu ~/.zshrc
echo "fpath=($(pwd)/config/completions \$fpath)" >> ~/.zshrc
echo "autoload -U compinit && compinit" >> ~/.zshrc
source ~/.zshrc
```

## Solução de Problemas

### Problemas Comuns

#### "Permission denied" (Linux/macOS)
Se encontrar este erro, certifique-se de que o script `latex-cli` tenha permissão de execução.
```bash
chmod +x bin/latex-cli
```

#### Docker não encontrado
Verifique se o Docker está em execução.
```bash
# Verifique as versões do Docker e Docker Compose
docker --version
docker compose version

# Se não estiverem instalados, acesse: https://docs.docker.com/get-docker/
```

#### WSL2 no Windows
Para uma melhor experiência no Windows, o uso do WSL2 é recomendado.
```powershell
# Habilite o WSL2 e defina-o como padrão
wsl --install
wsl --set-default-version 2

# Execute os comandos dentro do ambiente WSL2
wsl
cd /mnt/c/seu/projeto
```

#### O container não inicia
Se o container Docker não iniciar corretamente, verifique os logs para identificar a causa.
```bash
# Verifique os logs do container
./bin/latex-cli logs

# Se necessário, recrie o container
./bin/latex-cli clean
./bin/latex-cli setup
```

### Obter Ajuda

```bash
# Para ajuda geral sobre os comandos
./bin/latex-cli --help

# Para ajuda específica de um comando
./bin/latex-cli build --help
./bin/latex-cli init --help
```

## Estrutura do Projeto

```
latex-template/
├── bin/                    # Scripts executáveis
│   └── latex-cli           # A CLI principal
├── lib/                    # Bibliotecas de scripts da CLI
│   ├── commands/           # Implementação dos comandos individuais
│   ├── core/               # Funcionalidades centrais (Docker, config)
│   └── utils/              # Scripts utilitários
├── config/                 # Arquivos de configuração
│   ├── docker/             # Configurações do Docker e Docker Compose
│   ├── templates/          # Templates de arquivos LaTeX
│   ├── vscode/             # Configurações recomendadas para o VS Code
│   ├── completions/        # Scripts de autocompletion para shells
│   └── latex-cli.conf      # Arquivo de configuração principal da CLI
├── src/                    # Arquivos-fonte do seu documento LaTeX
│   ├── main.tex            # Arquivo principal do documento
│   ├── preamble.tex        # Preâmbulo, pacotes e configurações LaTeX
│   ├── references.bib      # Arquivo de bibliografia
│   └── chapters/           # Diretório para os capítulos do documento
├── dist/                   # Diretório de saída dos arquivos compilados (PDFs)
└── docs/                   # Documentação do projeto
```

## Contribuições

Contribuições são sempre bem-vindas! Siga os passos abaixo:

1.  Faça um fork do projeto.
2.  Crie uma branch para a sua nova feature (`git checkout -b feature/AmazingFeature`).
3.  Faça o commit das suas alterações (`git commit -m 'Add some AmazingFeature'`).
4.  Envie a sua branch para o repositório (`git push origin feature/AmazingFeature`).
5.  Abra um Pull Request.

## Reportar Problemas

Encontrou um bug? Abra uma [issue](https://github.com/martinsmiguel/latex-template/issues) e forneça as seguintes informações:

-   Uma descrição clara do problema.
-   Seu sistema operacional e versão.
-   A versão do Docker que está utilizando.
-   Logs relevantes que possam ajudar (`./bin/latex-cli logs`).

## Suporte

- **Documentação**: Consulte a pasta [docs/](docs/).
- **Issues**: Para bugs e problemas, abra uma [GitHub Issue](https://github.com/martinsmiguel/latex-template/issues).
- **Discussões**: Para dúvidas e sugestões, inicie uma [GitHub Discussion](https://github.com/martinsmiguel/latex-template/discussions).

## Licença

Este projeto está licenciado sob a MIT License. Veja o arquivo [LICENSE](docs/LICENSE) para mais detalhes.

---

**Desenvolvido para simplificar a criação de documentos com LaTeX.**