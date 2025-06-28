#!/bin/bash
#
# Sistema de logging da CLI
#

# Cores para mensagens
readonly RED='\033[0;31m'
readonly GREEN='\033[0;32m'
readonly YELLOW='\033[1;33m'
readonly BLUE='\033[0;34m'
readonly PURPLE='\033[0;35m'
readonly CYAN='\033[0;36m'
readonly NC='\033[0m' # No Color

# Níveis de log
readonly LOG_LEVEL_DEBUG=0
readonly LOG_LEVEL_INFO=1
readonly LOG_LEVEL_WARN=2
readonly LOG_LEVEL_ERROR=3

# Configurações do logger
LOG_LEVEL=${LOG_LEVEL:-$LOG_LEVEL_INFO}
LOG_FILE=""
LOG_TO_FILE=false

# Inicializa o sistema de logging
init_logger() {
    local log_dir="${PROJECT_ROOT}/dist/logs"
    LOG_FILE="${log_dir}/latex-cli.log"

    # Cria diretório de logs se não existir
    if [[ ! -d "$log_dir" ]]; then
        mkdir -p "$log_dir"
    fi

    # Verifica se deve logar em arquivo
    if [[ -w "$log_dir" ]]; then
        LOG_TO_FILE=true
    fi
}

# Função genérica de log
_log() {
    local level="$1"
    local level_num="$2"
    local color="$3"
    local message="$4"

    # Verifica se deve mostrar esta mensagem baseado no nível
    if [[ $level_num -lt $LOG_LEVEL ]]; then
        return 0
    fi

    # Prepara timestamp
    local timestamp
    timestamp=$(date '+%Y-%m-%d %H:%M:%S')

    # Prepara mensagem para terminal
    local terminal_msg
    if [[ "${QUIET:-false}" != "true" ]]; then
        terminal_msg="${color}[${level}]${NC} ${message}"
        echo -e "$terminal_msg" >&2
    fi

    # Log para arquivo se habilitado
    if [[ "$LOG_TO_FILE" == "true" ]]; then
        echo "${timestamp} [${level}] ${message}" >> "$LOG_FILE"
    fi
}

# Funções específicas de log
log_debug() {
    _log "DEBUG" $LOG_LEVEL_DEBUG "$PURPLE" "$1"
}

log_info() {
    _log "INFO" $LOG_LEVEL_INFO "$BLUE" "$1"
}

log_success() {
    _log "SUCCESS" $LOG_LEVEL_INFO "$GREEN" "$1"
}

log_warn() {
    _log "WARN" $LOG_LEVEL_WARN "$YELLOW" "$1"
}

log_error() {
    _log "ERROR" $LOG_LEVEL_ERROR "$RED" "$1"
}

# Função para mostrar progressos
show_progress() {
    local message="$1"
    local current="${2:-}"
    local total="${3:-}"

    if [[ -n "$current" && -n "$total" ]]; then
        local percent=$((current * 100 / total))
        echo -e "${CYAN}[${percent}%]${NC} ${message}" >&2
    else
        echo -e "${CYAN}[•]${NC} ${message}" >&2
    fi
}

# Função para mostrar spinner (útil para operações longas)
show_spinner() {
    local pid=$1
    local message="$2"
    local spin='|/-\'

    while kill -0 $pid 2>/dev/null; do
        for i in $(seq 0 3); do
            printf "\r${CYAN}[%c]${NC} %s" "${spin:i:1}" "$message" >&2
            sleep 0.1
        done
    done
    printf "\r" >&2
}

# Define nível de log baseado em string
set_log_level() {
    case "${1,,}" in
        "debug") LOG_LEVEL=$LOG_LEVEL_DEBUG ;;
        "info") LOG_LEVEL=$LOG_LEVEL_INFO ;;
        "warn"|"warning") LOG_LEVEL=$LOG_LEVEL_WARN ;;
        "error") LOG_LEVEL=$LOG_LEVEL_ERROR ;;
        *) log_warn "Nível de log desconhecido: $1" ;;
    esac
}
