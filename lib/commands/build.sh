#!/bin/bash
#
# Comando: build
# Compila o documento LaTeX
#

cmd_build() {
    local engine=""
    local output_dir=""
    local clean=false
    local watch=false
    local file=""

    # Parse de argumentos
    while [[ $# -gt 0 ]]; do
        case $1 in
            --engine|-e)
                engine="$2"
                shift 2
                ;;
            --output-dir|-o)
                output_dir="$2"
                shift 2
                ;;
            --clean|-c)
                clean=true
                shift
                ;;
            --watch|-w)
                watch=true
                shift
                ;;
            --file|-f)
                file="$2"
                shift 2
                ;;
            --help|-h)
                show_build_help
                return 0
                ;;
            *)
                log_error "Opção desconhecida: $1"
                show_build_help
                return 1
                ;;
        esac
    done

    # Usa configurações padrão se não especificado
    engine="${engine:-$(get_config latex_engine)}"
    output_dir="${output_dir:-$(get_config output_dir)}"

    log_info "Compilando documento LaTeX..."

    # Verifica se o ambiente Docker está executando
    if ! is_container_running; then
        log_info "Iniciando ambiente Docker..."
        start_docker_env || return 1
    fi

    # Encontra arquivo principal se não especificado
    if is_empty "$file"; then
        file=$(find_main_latex_file) || return 1
    fi

    # Garante que arquivo existe
    if [[ ! -f "$file" ]]; then
        log_error "Arquivo não encontrado: $file"
        return 1
    fi

    # Limpa arquivos antigos se solicitado
    if [[ "$clean" == "true" ]]; then
        clean_latex_files
    fi

    # Compila o documento
    if [[ "$watch" == "true" ]]; then
        build_with_watch "$file" "$engine" "$output_dir"
    else
        build_once "$file" "$engine" "$output_dir"
    fi
}

show_build_help() {
    cat << EOF
USAGE:
    latex-cli build [OPTIONS]

DESCRIPTION:
    Compila o documento LaTeX para PDF.

    O comando utiliza latexmk para compilação otimizada e suporte
    a bibliografia. A compilação ocorre dentro do container Docker.

OPTIONS:
    -e, --engine ENGINE     Engine LaTeX (pdflatex, xelatex, lualatex)
    -o, --output-dir DIR    Diretório de saída (padrão: dist)
    -f, --file FILE         Arquivo LaTeX específico para compilar
    -c, --clean             Limpa arquivos auxiliares antes da compilação
    -w, --watch             Modo de observação (auto-compilação)
    -h, --help              Mostra esta ajuda

ENGINES DISPONÍVEIS:
    pdflatex    Engine padrão, mais rápido
    xelatex     Melhor suporte a Unicode e fontes
    lualatex    Engine moderno com Lua integrado

EXAMPLES:
    latex-cli build
    latex-cli build --engine xelatex
    latex-cli build --clean --output-dir output
    latex-cli build --watch

EOF
}

build_once() {
    local file="$1"
    local engine="$2"
    local output_dir="$3"

    local start_time=$(date +%s)

    log_info "Compilando $file com $engine..."

    # Prepara comando de compilação
    local latexmk_cmd=(
        "latexmk"
        "-$engine"
        "-interaction=nonstopmode"
        "-file-line-error"
        "-synctex=1"
        "-output-directory=$output_dir"
        "$file"
    )

    # Executa compilação no container
    if docker_exec_no_tty "$(get_config container_name)" "${latexmk_cmd[@]}"; then
        local end_time=$(date +%s)
        local elapsed=$(show_elapsed_time "$start_time" "$end_time")

        log_success "Compilação concluída em $elapsed"

        # Mostra localização do PDF
        local pdf_file="$output_dir/$(basename "${file%.tex}.pdf")"
        if [[ -f "$pdf_file" ]]; then
            log_info "PDF gerado: $pdf_file"
        fi
    else
        log_error "Falha na compilação"
        show_compilation_errors "$output_dir"
        return 1
    fi
}

build_with_watch() {
    local file="$1"
    local engine="$2"
    local output_dir="$3"

    log_info "Iniciando modo de observação..."
    log_info "Pressione Ctrl+C para parar"

    # Prepara comando de observação
    local latexmk_cmd=(
        "latexmk"
        "-$engine"
        "-interaction=nonstopmode"
        "-file-line-error"
        "-synctex=1"
        "-output-directory=$output_dir"
        "-pvc"
        "$file"
    )

    # Executa no container com modo interativo
    docker_exec "$(get_config container_name)" "${latexmk_cmd[@]}"
}

clean_latex_files() {
    local output_dir="$(get_config output_dir)"

    log_info "Limpando arquivos auxiliares..."

    # Extensões de arquivos a serem removidos
    local extensions=(
        "aux" "bbl" "blg" "fdb_latexmk" "fls" "log"
        "out" "toc" "lof" "lot" "synctex.gz" "nav"
        "snm" "vrb" "bcf" "run.xml"
    )

    for ext in "${extensions[@]}"; do
        find "$output_dir" -name "*.$ext" -delete 2>/dev/null || true
    done

    log_success "Arquivos auxiliares removidos"
}

show_compilation_errors() {
    local output_dir="$1"
    local log_file="$output_dir/main.log"

    if [[ -f "$log_file" ]]; then
        log_error "Últimos erros de compilação:"
        echo "----------------------------------------"

        # Mostra apenas linhas com erros
        grep -A 2 -B 2 "^!" "$log_file" | tail -20 || true

        echo "----------------------------------------"
        log_info "Log completo disponível em: $log_file"
    fi
}
