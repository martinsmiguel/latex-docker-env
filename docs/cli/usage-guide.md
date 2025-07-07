# Guia de Uso da CLI ltx

## Introdução

A CLI `ltx` é uma ferramenta moderna para desenvolvimento LaTeX usando Docker. Este guia mostra como usar todos os comandos disponíveis.

## Instalação

### Pré-requisitos
- Docker Desktop instalado e rodando
- Git para clonar o repositório

### Instalação Básica
```bash
git clone https://github.com/martinsmiguel/latex-docker-env.git
cd latex-docker-env
chmod +x bin/ltx
```

## Comandos Principais

### `ltx setup`
Configura o ambiente de desenvolvimento inicial.

```bash
./bin/ltx setup
```

**O que faz:**
- Verifica se o Docker está rodando
- Baixa a imagem LaTeX necessária
- Cria diretórios do projeto (src/, dist/, tmp/)
- Configura VS Code (opcional)

**Saída esperada:**
```
>> Configurando ambiente LaTeX Docker...
[OK] Docker verificado
>> Verificando imagem LaTeX: blang/latex:ubuntu
[OK] Imagem LaTeX pronta
[OK] Estrutura de diretórios criada
[SUCCESS] Ambiente configurado com sucesso!
```

### `ltx init`
**[WIP]** Inicializa um novo documento LaTeX.

```bash
# Modo interativo
./bin/ltx init

# Com parâmetros
./bin/ltx init --title "Minha Tese" --author "Seu Nome"
```

### `ltx build`
**[WIP]** Compila o documento LaTeX.

```bash
./bin/ltx build
```

### `ltx watch`
**[WIP]** Monitora arquivos e recompila automaticamente.

```bash
./bin/ltx watch
```

### `ltx status`
**[WIP]** Mostra o status do ambiente.

```bash
./bin/ltx status
```

### `ltx clean`
**[WIP]** Remove arquivos temporários.

```bash
./bin/ltx clean
```

## Fluxo de Trabalho Típico

### 1. Configuração Inicial (Uma vez)
```bash
git clone https://github.com/martinsmiguel/latex-docker-env.git
cd latex-docker-env
./bin/ltx setup
```

### 2. Criação de Documento
```bash
./bin/ltx init --title "Meu Artigo"
```

### 3. Desenvolvimento
```bash
# Modo automático (recomendado)
./bin/ltx watch

# Ou compilação manual
./bin/ltx build
```

### 4. Finalização
```bash
./bin/ltx clean
./bin/ltx build
```

## Estrutura de Arquivos

Após o setup, você terá:

```
seu-projeto/
├── src/                    # Arquivos LaTeX fonte
│   ├── main.tex           # Documento principal
│   ├── preamble.tex       # Configurações LaTeX
│   └── chapters/          # Capítulos (para documentos longos)
├── dist/                  # PDFs compilados
├── tmp/                   # Arquivos temporários
└── config/               # Configurações do projeto
```

## Configuração

### Arquivo de Configuração
A CLI lê configurações do arquivo `config/latex-cli.conf`:

```bash
# Engine LaTeX a usar
LATEX_ENGINE=xelatex

# Diretórios
OUTPUT_DIR=dist
SOURCE_DIR=src

# Docker
CONTAINER_NAME=latex-env
LATEX_IMAGE=blang/latex:ubuntu
```

### Flags Globais
```bash
./bin/ltx <comando> --help     # Ajuda para comando específico
./bin/ltx --verbose <comando>  # Saída detalhada
./bin/ltx --config /path/to/config <comando>  # Config customizado
```

## Solução de Problemas

### Docker não encontrado
```
[ERROR] Docker não encontrado ou não está rodando
```
**Solução:** Instale e inicie o Docker Desktop.

### Permissão negada (Linux/macOS)
```
permission denied: ./bin/ltx
```
**Solução:**
```bash
chmod +x bin/ltx
```

### Estrutura do projeto inválida
```
[ERROR] estrutura do projeto inválida: diretório config não encontrado
```
**Solução:** Execute no diretório raiz do latex-docker-env.

## Comparação com CLI Legada

| Funcionalidade | `latex-cli` (Bash) | `ltx` (Go) |
|----------------|---------------------|------------|
| Windows nativo | ❌ (requer WSL2)    | ✅         |
| Performance    | Lenta               | Rápida     |
| File watching  | Básico              | Avançado   |
| Configuração   | Simples             | Robusta    |
| Distribuição   | Scripts             | Binário    |

## Obtendo Ajuda

```bash
./bin/ltx --help                    # Ajuda geral
./bin/ltx <comando> --help          # Ajuda específica
./bin/ltx --version                 # Versão da CLI
```

Para mais informações, consulte:
- [Documentação completa](../README.md)
- [Guia de desenvolvimento](development-guide.md)
- [Issues no GitHub](https://github.com/martinsmiguel/latex-docker-env/issues)
