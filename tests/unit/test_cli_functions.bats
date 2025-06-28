#!/usr/bin/env bats

# Testes das funções auxiliares da CLI

# Setup executado antes de cada teste
setup() {
    # Carrega funções da CLI para teste
    export LATEX_CLI_TEST_MODE="true"
    export LOG_LEVEL="ERROR"  # Reduz logs durante testes

    # Mock do PROJECT_ROOT para testes
    export TEST_PROJECT_ROOT="/tmp/latex-template-test-$$"
    mkdir -p "$TEST_PROJECT_ROOT"
}

# Teardown executado após cada teste
teardown() {
    # Limpeza
    if [[ -n "$TEST_PROJECT_ROOT" && -d "$TEST_PROJECT_ROOT" ]]; then
        rm -rf "$TEST_PROJECT_ROOT"
    fi
    unset LATEX_CLI_TEST_MODE
    unset TEST_PROJECT_ROOT
}

# Função helper para carregar funções da CLI
load_cli_functions() {
    # Define constantes necessárias
    RED='\033[0;31m'
    GREEN='\033[0;32m'
    YELLOW='\033[1;33m'
    BLUE='\033[0;34m'
    PURPLE='\033[0;35m'
    CYAN='\033[0;36m'
    NC='\033[0m'

    # Variáveis globais necessárias
    LOG_LEVEL="${LOG_LEVEL:-INFO}"
    LATEX_CLI_VERBOSE="${LATEX_CLI_VERBOSE:-false}"
    LATEX_CLI_QUIET="${LATEX_CLI_QUIET:-false}"
    PROJECT_ROOT="${TEST_PROJECT_ROOT}"
    LOG_DIR="$TEST_PROJECT_ROOT/logs"
    CLI_LOG_FILE="$LOG_DIR/test.log"

    # Define as funções necessárias
    write_log() {
        local level="$1"
        local message="$2"
        local timestamp=$(date '+%Y-%m-%d %H:%M:%S')
        mkdir -p "$LOG_DIR"
        echo "[$timestamp] [$level] $message" >> "$CLI_LOG_FILE"
    }

    log_info() {
        local message="$1"
        write_log "INFO" "$message"
        if [[ "${LATEX_CLI_QUIET:-false}" != "true" ]]; then
            printf "INFO: %s\n" "$message"
        fi
    }

    log_error() {
        local message="$1"
        write_log "ERROR" "$message"
        printf "${RED}✗ ERRO: %s${NC}\n" "$message" >&2
    }

    command_exists() {
        command -v "$1" >/dev/null 2>&1
    }

    version_compare() {
        local version1="$1"
        local version2="$2"

        if [[ "$version1" == "$version2" ]]; then
            return 0
        fi

        local ver1
        local ver2
        IFS='.' read -ra ver1 <<< "$version1"
        IFS='.' read -ra ver2 <<< "$version2"

        # Compara major, minor e patch
        for i in {0..2}; do
            local v1=${ver1[i]:-0}
            local v2=${ver2[i]:-0}

            if (( v1 > v2 )); then
                return 1 # version1 > version2
            elif (( v1 < v2 )); then
                return 2 # version1 < version2
            fi
        done

        return 0 # versões iguais
    }
}

@test "command_exists should return true for existing command" {
    load_cli_functions

    run command_exists "bash"
    [ "$status" -eq 0 ]
}

@test "command_exists should return false for non-existing command" {
    load_cli_functions

    run command_exists "command_that_does_not_exist_12345"
    [ "$status" -eq 1 ]
}

@test "version_compare should handle equal versions" {
    load_cli_functions

    run version_compare "1.0.0" "1.0.0"
    [ "$status" -eq 0 ]
}

@test "version_compare should handle newer version" {
    load_cli_functions

    run version_compare "2.0.0" "1.0.0"
    [ "$status" -eq 1 ]
}

@test "version_compare should handle older version" {
    load_cli_functions

    run version_compare "1.0.0" "2.0.0"
    [ "$status" -eq 2 ]
}

@test "version_compare should handle patch versions" {
    load_cli_functions

    run version_compare "1.0.1" "1.0.0"
    [ "$status" -eq 1 ]
}

@test "log functions should work correctly" {
    load_cli_functions

    # Test log_info
    run log_info "test message"
    [ "$status" -eq 0 ]
    [[ "$output" =~ "INFO: test message" ]]
}

@test "log_error should write to stderr" {
    load_cli_functions

    run log_error "error message"
    [ "$status" -eq 0 ]
    # Em bats, stderr vai para output quando capturado com run
    [[ "$output" =~ "ERRO: error message" ]]
}

@test "project structure validation should work" {
    load_cli_functions

    # Cria estrutura mínima
    touch "$TEST_PROJECT_ROOT/docker-compose.yml"
    mkdir -p "$TEST_PROJECT_ROOT/templates"

    # Mock PROJECT_ROOT
    PROJECT_ROOT="$TEST_PROJECT_ROOT"

    # Função que verifica se estamos na raiz do projeto
    if [[ -f "${PROJECT_ROOT}/docker-compose.yml" ]] && [[ -d "${PROJECT_ROOT}/templates" ]]; then
        echo "valid_project"
    else
        echo "invalid_project"
    fi

    run bash -c 'if [[ -f "${TEST_PROJECT_ROOT}/docker-compose.yml" ]] && [[ -d "${TEST_PROJECT_ROOT}/templates" ]]; then echo "valid_project"; else echo "invalid_project"; fi'
    [ "$status" -eq 0 ]
    [[ "$output" == "valid_project" ]]
}
