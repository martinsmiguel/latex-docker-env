#!/bin/bash
#
# Comando: watch
# Modo de observação para auto-compilação
#

cmd_watch() {
    local file=""
    local engine=""

    # Parse de argumentos
    while [[ $# -gt 0 ]]; do
        case $1 in
            --file|-f)
                file="$2"
                shift 2
                ;;
            --engine|-e)
                engine="$2"
                shift 2
                ;;
            --help|-h)
                show_watch_help
                return 0
                ;;
            *)
                log_error "Opção desconhecida: $1"
                show_watch_help
                return 1
                ;;
        esac
    done

    log_info "Iniciando modo de observação..."

    # Delega para o comando build com --watch
    source "${LIB_DIR}/commands/build.sh"
    cmd_build --watch ${engine:+--engine "$engine"} ${file:+--file "$file"}
}

show_watch_help() {
    cat << EOF
USAGE:
    latex-cli watch [OPTIONS]

DESCRIPTION:
    Inicia o modo de observação para auto-compilação do documento LaTeX.

    O documento será recompilado automaticamente sempre que arquivos
    LaTeX forem modificados. Pressione Ctrl+C para parar.

OPTIONS:
    -f, --file FILE         Arquivo LaTeX específico para observar
    -e, --engine ENGINE     Engine LaTeX (pdflatex, xelatex, lualatex)
    -h, --help              Mostra esta ajuda

EXAMPLES:
    latex-cli watch
    latex-cli watch --engine xelatex
    latex-cli watch --file src/documento.tex

EOF
}
