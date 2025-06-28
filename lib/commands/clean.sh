#!/bin/bash
#
# Comando: clean
# Limpa arquivos temporários e de build
#

cmd_clean() {
    local all=false
    local logs=false
    local cache=false
    local force=false

    # Parse de argumentos
    while [[ $# -gt 0 ]]; do
        case $1 in
            --all|-a)
                all=true
                shift
                ;;
            --logs|-l)
                logs=true
                shift
                ;;
            --cache|-c)
                cache=true
                shift
                ;;
            --force|-f)
                force=true
                shift
                ;;
            --help|-h)
                show_clean_help
                return 0
                ;;
            *)
                log_error "Opção desconhecida: $1"
                show_clean_help
                return 1
                ;;
        esac
    done

    # Se --all foi especificado, ativa todas as opções
    if [[ "$all" == "true" ]]; then
        logs=true
        cache=true
    fi

    # Se nenhuma opção específica foi dada, limpa arquivos de build padrão
    if [[ "$logs" == "false" && "$cache" == "false" ]]; then
        clean_build_files
    else
        # Executa limpezas específicas
        clean_build_files

        if [[ "$logs" == "true" ]]; then
            clean_logs "$force"
        fi

        if [[ "$cache" == "true" ]]; then
            clean_cache "$force"
        fi
    fi

    log_success "Limpeza concluída"
}

show_clean_help() {
    cat << EOF
USAGE:
    latex-cli clean [OPTIONS]

DESCRIPTION:
    Remove arquivos temporários e de build do projeto LaTeX.

    Por padrão, remove apenas arquivos auxiliares da compilação
    (como .aux, .log, .toc, etc.), mantendo o PDF final.

OPTIONS:
    -a, --all       Remove tudo (build, logs, cache)
    -l, --logs      Remove arquivos de log
    -c, --cache     Remove cache do Docker
    -f, --force     Não pede confirmação
    -h, --help      Mostra esta ajuda

TIPOS DE LIMPEZA:
    build    Arquivos auxiliares (.aux, .log, .toc, etc.)
    logs     Logs da CLI e do LaTeX
    cache    Cache do Docker e TeX Live

EXAMPLES:
    latex-cli clean
    latex-cli clean --all
    latex-cli clean --logs --cache
    latex-cli clean --force

EOF
}

clean_build_files() {
    log_info "Removendo arquivos de build..."

    local output_dir="$(get_config output_dir)"
    local source_dir="$(get_config source_dir)"

    # Extensões de arquivos auxiliares
    local aux_extensions=(
        "aux" "bbl" "blg" "fdb_latexmk" "fls" "log"
        "out" "toc" "lof" "lot" "synctex.gz" "nav"
        "snm" "vrb" "bcf" "run.xml" "auxlock"
        "figlist" "makefile" "figlist-*"
    )

    local removed_count=0

    # Remove arquivos auxiliares do diretório de saída
    for ext in "${aux_extensions[@]}"; do
        while IFS= read -r -d '' file; do
            rm -f "$file"
            ((removed_count++))
            log_debug "Removido: $file"
        done < <(find "$output_dir" -name "*.$ext" -print0 2>/dev/null)
    done

    # Remove arquivos auxiliares do diretório fonte
    for ext in "${aux_extensions[@]}"; do
        while IFS= read -r -d '' file; do
            rm -f "$file"
            ((removed_count++))
            log_debug "Removido: $file"
        done < <(find "$source_dir" -name "*.$ext" -print0 2>/dev/null)
    done

    # Remove diretórios temporários
    local temp_dirs=(
        "$output_dir/_minted-*"
        "$source_dir/_minted-*"
        "$output_dir/.texpadtmp"
        "$source_dir/.texpadtmp"
    )

    for pattern in "${temp_dirs[@]}"; do
        if compgen -G "$pattern" > /dev/null; then
            rm -rf $pattern
            ((removed_count++))
        fi
    done

    log_success "Removidos $removed_count arquivos auxiliares"
}

clean_logs() {
    local force="$1"
    local log_dir="${PROJECT_ROOT}/dist/logs"

    if [[ ! -d "$log_dir" ]]; then
        log_info "Nenhum diretório de logs encontrado"
        return 0
    fi

    local log_count=$(find "$log_dir" -type f | wc -l)

    if [[ $log_count -eq 0 ]]; then
        log_info "Nenhum arquivo de log encontrado"
        return 0
    fi

    if [[ "$force" != "true" ]]; then
        if ! ask_yes_no "Remover $log_count arquivos de log?" "n"; then
            log_info "Limpeza de logs cancelada"
            return 0
        fi
    fi

    log_info "Removendo arquivos de log..."

    rm -rf "$log_dir"/*

    log_success "Logs removidos"
}

clean_cache() {
    local force="$1"

    if [[ "$force" != "true" ]]; then
        if ! ask_yes_no "Limpar cache do Docker? Isso irá forçar rebuild da próxima compilação." "n"; then
            log_info "Limpeza de cache cancelada"
            return 0
        fi
    fi

    log_info "Limpando cache do Docker..."

    # Para o container se estiver executando
    if is_container_running; then
        log_info "Parando container..."
        stop_docker_env
    fi

    # Remove volumes do Docker
    if docker volume ls --filter "name=latex-cache" --format "table {{.Name}}" | grep -q "latex-cache"; then
        docker volume rm latex-cache 2>/dev/null || true
        log_success "Cache do TeX Live removido"
    fi

    # Remove imagens não utilizadas
    docker image prune -f >/dev/null 2>&1 || true

    log_success "Cache do Docker limpo"
}
