# Configuração da CLI ltx

## Visão Geral

A CLI `ltx` usa um sistema de configuração hierárquico baseado no Viper, permitindo múltiplas fontes de configuração com precedência bem definida.

## Precedência de Configuração

A configuração é aplicada na seguinte ordem (maior precedência primeiro):

1. **Flags de linha de comando**
2. **Variáveis de ambiente**
3. **Arquivo de configuração**
4. **Valores padrão**

## Arquivo de Configuração

### Localização
Por padrão, a CLI procura por:
- `./config/latex-cli.conf` (formato shell)
- `./config/latex-cli.yaml` (formato YAML)
- `./config/latex-cli.json` (formato JSON)

### Formato Shell (padrão)
```bash
# config/latex-cli.conf

# Engine LaTeX a usar
LATEX_ENGINE=xelatex

# Diretórios
OUTPUT_DIR=dist
SOURCE_DIR=src

# Docker
CONTAINER_NAME=latex-env
LATEX_IMAGE=blang/latex:ubuntu
WATCH_DEBOUNCE=500ms

# VS Code
VSCODE_CONFIG=true
```

### Formato YAML (recomendado para configurações avançadas)
```yaml
# config/latex-cli.yaml
latex:
  engine: xelatex
  packages:
    - amsmath
    - graphicx
    - biblatex

directories:
  output: dist
  source: src
  temp: tmp

docker:
  image: blang/latex:ubuntu
  container_name: latex-env
  memory_limit: 2g

watch:
  debounce: 500ms
  ignore_patterns:
    - "*.log"
    - "*.aux"
    - "*.tmp"

vscode:
  enabled: true
  extensions:
    - James-Yu.latex-workshop
    - ms-vscode-remote.remote-containers
```

### Formato JSON
```json
{
  "latex_engine": "xelatex",
  "output_dir": "dist",
  "source_dir": "src",
  "container_name": "latex-env",
  "latex_image": "blang/latex:ubuntu",
  "watch_debounce": "500ms"
}
```

## Variáveis de Ambiente

Todas as configurações podem ser definidas via variáveis de ambiente com prefixo `LTX_`:

```bash
export LTX_LATEX_ENGINE=pdflatex
export LTX_OUTPUT_DIR=output
export LTX_DOCKER_IMAGE=texlive/texlive:latest
export LTX_VERBOSE=true
```

## Flags de Linha de Comando

### Flags Globais
```bash
--config string     # Arquivo de configuração customizado
--verbose, -v       # Saída detalhada
--help, -h         # Ajuda
--version          # Versão da CLI
```

### Exemplo de Uso
```bash
# Usar arquivo de configuração específico
./bin/ltx --config /path/to/custom.conf setup

# Modo verbose
./bin/ltx --verbose build

# Combinando flags
./bin/ltx --config custom.yaml --verbose watch
```

## Configurações Disponíveis

### LaTeX Engine
```bash
# Opções: pdflatex, xelatex, lualatex
LATEX_ENGINE=xelatex
```

### Diretórios
```bash
OUTPUT_DIR=dist        # Onde salvar PDFs compilados
SOURCE_DIR=src         # Onde estão os arquivos .tex
TEMP_DIR=tmp          # Arquivos temporários
```

### Docker
```bash
LATEX_IMAGE=blang/latex:ubuntu    # Imagem Docker a usar
CONTAINER_NAME=latex-env          # Nome do container
DOCKER_MEMORY=2g                  # Limite de memória
DOCKER_CPUS=2                     # Limite de CPU
```

### File Watching
```bash
WATCH_DEBOUNCE=500ms              # Delay antes de recompilar
WATCH_IGNORE="*.log,*.aux,*.tmp"  # Padrões a ignorar
```

### VS Code
```bash
VSCODE_CONFIG=true                # Configurar VS Code automaticamente
VSCODE_EXTENSIONS=true            # Instalar extensões recomendadas
```

### Logging
```bash
LOG_LEVEL=info                    # debug, info, warn, error
LOG_FORMAT=text                   # text, json
LOG_FILE=/path/to/logfile         # Arquivo de log (opcional)
```

## Perfis de Configuração

### Desenvolvimento
```yaml
# config/profiles/dev.yaml
latex:
  engine: xelatex
  draft_mode: true

watch:
  debounce: 100ms
  auto_clean: false

logging:
  level: debug
```

### Produção
```yaml
# config/profiles/prod.yaml
latex:
  engine: xelatex
  draft_mode: false

build:
  clean_first: true
  optimize: true

logging:
  level: warn
```

### Uso de Perfis
```bash
./bin/ltx --config config/profiles/dev.yaml build
```

## Configuração Dinâmica

### Via Comando
```bash
# Definir configuração temporariamente
LTX_LATEX_ENGINE=pdflatex ./bin/ltx build

# Múltiplas configurações
LTX_VERBOSE=true LTX_OUTPUT_DIR=output ./bin/ltx watch
```

### Via Script
```bash
#!/bin/bash
# scripts/build-thesis.sh

export LTX_LATEX_ENGINE=xelatex
export LTX_OUTPUT_DIR=thesis-output
export LTX_SOURCE_DIR=thesis-src

./bin/ltx build
```

## Validação de Configuração

A CLI valida automaticamente:

1. **Engines LaTeX suportados**
2. **Diretórios existentes/criáveis**
3. **Imagens Docker válidas**
4. **Formatos de tempo válidos**

### Exemplo de Erro
```
[ERROR] Configuração inválida: latex_engine 'invalid' não suportado
Engines suportados: pdflatex, xelatex, lualatex
```

## Comandos de Configuração

### Visualizar Configuração Atual
```bash
./bin/ltx config show
```

### Validar Configuração
```bash
./bin/ltx config validate
```

### Gerar Configuração Padrão
```bash
./bin/ltx config init
```

### Configuração Interativa
```bash
./bin/ltx config setup
```

## Migração de Configuração

### Da CLI Bash para Go
A CLI Go pode ler automaticamente configurações da CLI Bash:

```bash
# Migrar configuração existente
./bin/ltx config migrate --from config/latex-cli.conf
```

### Formato de Migração
```bash
# Antes (Bash)
LATEX_ENGINE=xelatex
OUTPUT_DIR=dist

# Depois (YAML)
latex:
  engine: xelatex
directories:
  output: dist
```

## Debugging de Configuração

### Modo Verbose
```bash
./bin/ltx --verbose config show
```

### Saída Esperada
```
[INFO] Carregando configuração...
[INFO] Arquivo encontrado: ./config/latex-cli.conf
[INFO] Variáveis de ambiente: 2 encontradas
[INFO] Configuração final:
  latex_engine: xelatex
  output_dir: dist
  source_dir: src
  container_name: latex-env
```

## Exemplos Práticos

### Configuração Mínima
```bash
# config/minimal.conf
LATEX_ENGINE=pdflatex
OUTPUT_DIR=pdf
```

### Configuração Completa
```yaml
# config/complete.yaml
latex:
  engine: xelatex
  options: "--interaction=nonstopmode"

directories:
  output: dist
  source: src
  temp: tmp

docker:
  image: texlive/texlive:latest
  container: latex-container
  memory: 4g

watch:
  enabled: true
  debounce: 300ms
  ignore: ["*.log", "*.aux", "*.fls", "*.fdb_latexmk"]

build:
  clean_first: false
  bibtex: true
  makeindex: true

vscode:
  enabled: true
  settings:
    "latex-workshop.latex.recipes": [
      {
        "name": "XeLaTeX",
        "tools": ["xelatex", "bibtex", "xelatex", "xelatex"]
      }
    ]
```

Esta flexibilidade permite que a CLI se adapte a diferentes fluxos de trabalho e preferências do usuário.
