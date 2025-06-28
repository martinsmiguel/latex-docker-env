#!/bin/bash
#
# Configurações da CLI
#

# Configurações padrão
readonly DEFAULT_CONFIG="${PROJECT_ROOT}/config/latex-cli.conf"

# Variáveis de configuração global
CONFIG_LOADED=false
LATEX_ENGINE="pdflatex"
OUTPUT_DIR="dist"
SOURCE_DIR="src"
CONTAINER_NAME="latex-env"
VERBOSE=false
QUIET=false

# Carrega arquivo de configuração
load_config() {
    if [[ "$CONFIG_LOADED" == "true" ]]; then
        return 0
    fi

    # Arquivo de configuração global do projeto
    if [[ -f "$DEFAULT_CONFIG" ]]; then
        log_debug "Carregando configuração de: $DEFAULT_CONFIG"
        # shellcheck source=/dev/null
        source "$DEFAULT_CONFIG"
    fi

    # Arquivo de configuração local do usuário
    local user_config="${HOME}/.latex-cli.conf"
    if [[ -f "$user_config" ]]; then
        log_debug "Carregando configuração do usuário: $user_config"
        # shellcheck source=/dev/null
        source "$user_config"
    fi

    # Variáveis de ambiente sobrescrevem configurações de arquivo
    LATEX_ENGINE="${LATEX_ENGINE:-$LATEX_ENGINE}"
    OUTPUT_DIR="${LATEX_OUTPUT_DIR:-$OUTPUT_DIR}"
    SOURCE_DIR="${LATEX_SOURCE_DIR:-$SOURCE_DIR}"
    CONTAINER_NAME="${LATEX_CONTAINER_NAME:-$CONTAINER_NAME}"
    VERBOSE="${LATEX_VERBOSE:-$VERBOSE}"
    QUIET="${LATEX_QUIET:-$QUIET}"

    CONFIG_LOADED=true
    log_debug "Configuração carregada com sucesso"
}

# Retorna uma configuração específica
get_config() {
    local key="$1"
    local default="${2:-}"

    case "$key" in
        "latex_engine") echo "$LATEX_ENGINE" ;;
        "output_dir") echo "$OUTPUT_DIR" ;;
        "source_dir") echo "$SOURCE_DIR" ;;
        "container_name") echo "$CONTAINER_NAME" ;;
        "verbose") echo "$VERBOSE" ;;
        "quiet") echo "$QUIET" ;;
        *) echo "$default" ;;
    esac
}

# Define uma configuração
set_config() {
    local key="$1"
    local value="$2"

    case "$key" in
        "latex_engine") LATEX_ENGINE="$value" ;;
        "output_dir") OUTPUT_DIR="$value" ;;
        "source_dir") SOURCE_DIR="$value" ;;
        "container_name") CONTAINER_NAME="$value" ;;
        "verbose") VERBOSE="$value" ;;
        "quiet") QUIET="$value" ;;
        *) log_error "Configuração desconhecida: $key" ;;
    esac
}

# Cria arquivo de configuração padrão
create_default_config() {
    local config_file="$1"

    cat > "$config_file" << 'EOF'
# Configuração do LaTeX CLI
# Este arquivo é carregado automaticamente pela CLI

# Engine LaTeX a usar (pdflatex, xelatex, lualatex)
LATEX_ENGINE="pdflatex"

# Diretório de saída para arquivos compilados
OUTPUT_DIR="dist"

# Diretório fonte dos arquivos LaTeX
SOURCE_DIR="src"

# Nome do container Docker
CONTAINER_NAME="latex-env"

# Logs verbosos (true/false)
VERBOSE=false

# Modo silencioso (true/false)
QUIET=false

# Configurações do Docker Compose
COMPOSE_FILE="config/docker/docker-compose.yml"

# Configurações de template
TEMPLATES_DIR="config/templates"
EOF

    log_info "Arquivo de configuração criado: $config_file"
}
