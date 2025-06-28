#!/usr/bin/env bats

# Testes do sistema de logging

setup() {
    export LATEX_CLI_TEST_MODE="true"
    export TEST_LOG_DIR="/tmp/latex-cli-log-test-$$"
    export TEST_CLI_LOG_FILE="$TEST_LOG_DIR/test-cli.log"
    mkdir -p "$TEST_LOG_DIR"
}

teardown() {
    if [[ -n "$TEST_LOG_DIR" && -d "$TEST_LOG_DIR" ]]; then
        rm -rf "$TEST_LOG_DIR"
    fi
    unset LATEX_CLI_TEST_MODE
    unset TEST_LOG_DIR
    unset TEST_CLI_LOG_FILE
}

load_cli_functions() {
    # Carrega as constantes necessárias
    RED='\033[0;31m'
    GREEN='\033[0;32m'
    YELLOW='\033[1;33m'
    BLUE='\033[0;34m'
    PURPLE='\033[0;35m'
    CYAN='\033[0;36m'
    NC='\033[0m'

    # Carrega variáveis necessárias
    LOG_LEVEL="${LOG_LEVEL:-INFO}"
    LATEX_CLI_VERBOSE="${LATEX_CLI_VERBOSE:-false}"
    LATEX_CLI_QUIET="${LATEX_CLI_QUIET:-false}"

    # Override das constantes para teste
    LOG_DIR="$TEST_LOG_DIR"
    CLI_LOG_FILE="$TEST_CLI_LOG_FILE"

    # Define as funções de logging diretamente
    write_log() {
        local level="$1"
        local message="$2"
        local timestamp=$(date '+%Y-%m-%d %H:%M:%S')
        mkdir -p "$LOG_DIR"
        echo "[$timestamp] [$level] $message" >> "$CLI_LOG_FILE"
    }

    log_debug() {
        local message="$1"
        write_log "DEBUG" "$message"
        if [[ "$LOG_LEVEL" == "DEBUG" ]] || [[ "${LATEX_CLI_VERBOSE:-false}" == "true" ]]; then
            printf "${PURPLE}DEBUG: %s${NC}\n" "$message" >&2
        fi
    }

    log_info() {
        local message="$1"
        write_log "INFO" "$message"
        if [[ "${LATEX_CLI_QUIET:-false}" != "true" ]]; then
            printf "INFO: %s\n" "$message"
        fi
    }

    log_success() {
        local message="$1"
        write_log "SUCCESS" "$message"
        if [[ "${LATEX_CLI_QUIET:-false}" != "true" ]]; then
            printf "${GREEN}✓ %s${NC}\n" "$message"
        fi
    }

    log_warning() {
        local message="$1"
        write_log "WARNING" "$message"
        if [[ "${LATEX_CLI_QUIET:-false}" != "true" ]]; then
            printf "${YELLOW}⚠ %s${NC}\n" "$message" >&2
        fi
    }

    log_error() {
        local message="$1"
        write_log "ERROR" "$message"
        printf "${RED}✗ ERRO: %s${NC}\n" "$message" >&2
    }

    log_command() {
        printf "${PURPLE}Executando: %s${NC}\n" "$1"
    }
}

@test "write_log should create log file and write entry" {
    load_cli_functions

    write_log "INFO" "test message"

    [ -f "$TEST_CLI_LOG_FILE" ]

    # Verifica se a mensagem foi escrita
    run grep "test message" "$TEST_CLI_LOG_FILE"
    [ "$status" -eq 0 ]
    [[ "$output" =~ "[INFO] test message" ]]
}

@test "log_info should write to both console and file" {
    load_cli_functions

    run log_info "info test message"

    [ "$status" -eq 0 ]
    [[ "$output" =~ "INFO: info test message" ]]

    # Verifica se foi escrito no arquivo
    [ -f "$TEST_CLI_LOG_FILE" ]
    run grep "info test message" "$TEST_CLI_LOG_FILE"
    [ "$status" -eq 0 ]
}

@test "log_error should write to both stderr and file" {
    load_cli_functions

    run log_error "error test message"

    [ "$status" -eq 0 ]
    [[ "$output" =~ "ERRO: error test message" ]]

    # Verifica se foi escrito no arquivo
    [ -f "$TEST_CLI_LOG_FILE" ]
    run grep "error test message" "$TEST_CLI_LOG_FILE"
    [ "$status" -eq 0 ]
}

@test "log_debug should only show with verbose mode" {
    load_cli_functions

    # Sem verbose mode
    export LATEX_CLI_VERBOSE="false"
    run log_debug "debug message"
    [ "$status" -eq 0 ]
    [[ ! "$output" =~ "DEBUG:" ]]

    # Com verbose mode
    export LATEX_CLI_VERBOSE="true"
    run log_debug "debug message verbose"
    [ "$status" -eq 0 ]
    [[ "$output" =~ "DEBUG: debug message verbose" ]]
}

@test "log_debug should always write to file regardless of verbose mode" {
    load_cli_functions

    export LATEX_CLI_VERBOSE="false"
    log_debug "debug message file"

    # Verifica se foi escrito no arquivo mesmo sem verbose
    [ -f "$TEST_CLI_LOG_FILE" ]
    run grep "debug message file" "$TEST_CLI_LOG_FILE"
    [ "$status" -eq 0 ]
}

@test "log levels should be correctly written to file" {
    load_cli_functions

    log_debug "debug message"
    log_info "info message"
    log_warning "warning message"
    log_error "error message"
    log_success "success message"

    [ -f "$TEST_CLI_LOG_FILE" ]

    # Verifica cada nível
    run grep "\[DEBUG\]" "$TEST_CLI_LOG_FILE"
    [ "$status" -eq 0 ]

    run grep "\[INFO\]" "$TEST_CLI_LOG_FILE"
    [ "$status" -eq 0 ]

    run grep "\[WARNING\]" "$TEST_CLI_LOG_FILE"
    [ "$status" -eq 0 ]

    run grep "\[ERROR\]" "$TEST_CLI_LOG_FILE"
    [ "$status" -eq 0 ]

    run grep "\[SUCCESS\]" "$TEST_CLI_LOG_FILE"
    [ "$status" -eq 0 ]
}

@test "log timestamps should be valid format" {
    load_cli_functions

    log_info "timestamp test"

    # Verifica formato de timestamp: [YYYY-MM-DD HH:MM:SS]
    run grep -E "^\[20[0-9]{2}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}\]" "$TEST_CLI_LOG_FILE"
    [ "$status" -eq 0 ]
}
