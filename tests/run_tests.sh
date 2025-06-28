#!/bin/bash
#
# Script para executar todos os testes do LaTeX Template CLI
#

set -euo pipefail

# Cores para output
readonly RED='\033[0;31m'
readonly GREEN='\033[0;32m'
readonly YELLOW='\033[1;33m'
readonly BLUE='\033[0;34m'
readonly NC='\033[0m' # No Color

# Diretório do script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Funções auxiliares
log_info() {
    printf "${BLUE}INFO: %s${NC}\n" "$1"
}

log_success() {
    printf "${GREEN}✓ %s${NC}\n" "$1"
}

log_error() {
    printf "${RED}✗ %s${NC}\n" "$1" >&2
}

log_warning() {
    printf "${YELLOW}⚠ %s${NC}\n" "$1"
}

# Verifica se BATS está instalado
check_bats() {
    if ! command -v bats >/dev/null 2>&1; then
        log_error "BATS não está instalado."
        echo
        echo "Instale o BATS (Bash Automated Testing System):"
        echo
        echo "macOS:"
        echo "  brew install bats-core"
        echo
        echo "Ubuntu/Debian:"
        echo "  sudo apt-get install bats"
        echo
        echo "Manual:"
        echo "  git clone https://github.com/bats-core/bats-core.git"
        echo "  cd bats-core && sudo ./install.sh /usr/local"
        echo
        exit 1
    fi
}

# Executa testes de uma categoria
run_test_category() {
    local category="$1"
    local test_dir="$SCRIPT_DIR/$category"

    if [[ ! -d "$test_dir" ]]; then
        log_warning "Diretório de testes '$category' não encontrado"
        return 0
    fi

    local test_files=("$test_dir"/*.bats)
    if [[ ! -e "${test_files[0]}" ]]; then
        log_warning "Nenhum arquivo de teste encontrado em '$category'"
        return 0
    fi

    log_info "Executando testes de $category..."

    local failed=0
    for test_file in "${test_files[@]}"; do
        local test_name=$(basename "$test_file" .bats)
        printf "  %-30s " "$test_name"

        if bats "$test_file" >/dev/null 2>&1; then
            printf "${GREEN}✓${NC}\n"
        else
            printf "${RED}✗${NC}\n"
            ((failed++))
        fi
    done

    if [[ $failed -eq 0 ]]; then
        log_success "Todos os testes de $category passaram"
    else
        log_error "$failed teste(s) de $category falharam"
    fi

    return $failed
}

# Executa teste específico com output detalhado
run_specific_test() {
    local test_file="$1"

    if [[ ! -f "$test_file" ]]; then
        log_error "Arquivo de teste não encontrado: $test_file"
        exit 1
    fi

    log_info "Executando teste específico: $(basename "$test_file")"
    bats "$test_file"
}

# Função principal
main() {
    cd "$PROJECT_ROOT"

    local test_type=""
    local specific_test=""
    local verbose=false

    # Parse argumentos
    while [[ $# -gt 0 ]]; do
        case "$1" in
            --unit)
                test_type="unit"
                shift
                ;;
            --integration)
                test_type="integration"
                shift
                ;;
            --file)
                if [[ -z "$2" ]]; then
                    log_error "A opção --file requer um arquivo"
                    exit 1
                fi
                specific_test="$2"
                shift 2
                ;;
            --verbose|-v)
                verbose=true
                shift
                ;;
            --help|-h)
                cat << EOF
Uso: $0 [OPCOES]

Executa testes automatizados do LaTeX Template CLI.

Opções:
  --unit                 Executa apenas testes unitários
  --integration          Executa apenas testes de integração
  --file <arquivo>       Executa teste específico
  --verbose, -v          Output detalhado
  --help, -h             Mostra esta ajuda

Exemplos:
  $0                     # Executa todos os testes
  $0 --unit              # Apenas testes unitários
  $0 --file tests/unit/test_cli_functions.bats
  $0 --verbose           # Com output detalhado

EOF
                exit 0
                ;;
            *)
                log_error "Opção desconhecida: $1"
                echo "Use $0 --help para ver opções disponíveis"
                exit 1
                ;;
        esac
    done

    # Configurações
    export LATEX_CLI_TEST_MODE="true"
    if [[ "$verbose" == true ]]; then
        export BATS_VERBOSE="1"
    fi

    log_info "Iniciando testes do LaTeX Template CLI"
    echo

    # Verifica pré-requisitos
    check_bats

    # Executa teste específico se solicitado
    if [[ -n "$specific_test" ]]; then
        run_specific_test "$specific_test"
        return $?
    fi

    local total_failures=0

    # Executa categoria específica ou todas
    if [[ -n "$test_type" ]]; then
        run_test_category "$test_type"
        total_failures=$?
    else
        # Executa todas as categorias
        local categories=("unit" "integration")
        for category in "${categories[@]}"; do
            echo
            run_test_category "$category"
            ((total_failures += $?))
        done
    fi

    echo
    if [[ $total_failures -eq 0 ]]; then
        log_success "Todos os testes passaram! 🎉"
        exit 0
    else
        log_error "Alguns testes falharam. Verifique os detalhes acima."
        exit 1
    fi
}

# Executa função principal
main "$@"
