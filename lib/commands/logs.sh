#!/bin/bash
#
# Comando: logs
# Mostra logs do container e da CLI
#

cmd_logs() {
    local follow=false
    local lines=50
    local type="all"

    # Parse de argumentos
    while [[ $# -gt 0 ]]; do
        case $1 in
            --follow|-f)
                follow=true
                shift
                ;;
            --lines|-n)
                lines="$2"
                shift 2
                ;;
            --type|-t)
                type="$2"
                shift 2
                ;;
            --help|-h)
                show_logs_help
                return 0
                ;;
            *)
                log_error "Opção desconhecida: $1"
                show_logs_help
                return 1
                ;;
        esac
    done

    case "$type" in
        "docker")
            show_docker_logs_only "$follow" "$lines"
            ;;
        "cli")
            show_cli_logs_only "$follow" "$lines"
            ;;
        "latex")
            show_latex_logs_only "$follow" "$lines"
            ;;
        "all"|*)
            show_all_logs "$follow" "$lines"
            ;;
    esac
}

show_logs_help() {
    cat << EOF
USAGE:
    latex-cli logs [OPTIONS]

DESCRIPTION:
    Mostra logs do sistema LaTeX Template.

    Por padrão, mostra logs de todas as fontes. Use --type para
    filtrar logs específicos.

OPTIONS:
    -f, --follow            Segue logs em tempo real (como tail -f)
    -n, --lines NUM         Número de linhas a mostrar (padrão: 50)
    -t, --type TYPE         Tipo de logs (all, docker, cli, latex)
    -h, --help              Mostra esta ajuda

TIPOS DE LOGS:
    all       Todos os logs (padrão)
    docker    Logs do container Docker
    cli       Logs da CLI (latex-cli.log)
    latex     Logs de compilação LaTeX

EXAMPLES:
    latex-cli logs
    latex-cli logs --follow
    latex-cli logs --type docker --lines 100
    latex-cli logs --type cli --follow

EOF
}

show_docker_logs_only() {
    local follow="$1"
    local lines="$2"

    log_info "=== Logs do Docker ==="

    if ! container_exists; then
        log_warn "Container não existe"
        return 1
    fi

    local container_name="$(get_config container_name)"

    if [[ "$follow" == "true" ]]; then
        docker logs -f --tail "$lines" "$container_name"
    else
        docker logs --tail "$lines" "$container_name"
    fi
}

show_cli_logs_only() {
    local follow="$1"
    local lines="$2"

    log_info "=== Logs da CLI ==="

    local log_file="${PROJECT_ROOT}/dist/logs/latex-cli.log"

    if [[ ! -f "$log_file" ]]; then
        log_warn "Arquivo de log da CLI não encontrado: $log_file"
        return 1
    fi

    if [[ "$follow" == "true" ]]; then
        tail -f -n "$lines" "$log_file"
    else
        tail -n "$lines" "$log_file"
    fi
}

show_latex_logs_only() {
    local follow="$1"
    local lines="$2"

    log_info "=== Logs de Compilação LaTeX ==="

    local output_dir="$(get_config output_dir)"
    local main_log="$output_dir/main.log"

    if [[ ! -f "$main_log" ]]; then
        log_warn "Log de compilação não encontrado: $main_log"
        log_info "Execute 'latex-cli build' para gerar logs"
        return 1
    fi

    if [[ "$follow" == "true" ]]; then
        tail -f -n "$lines" "$main_log"
    else
        tail -n "$lines" "$main_log"
    fi
}

show_all_logs() {
    local follow="$1"
    local lines="$2"

    if [[ "$follow" == "true" ]]; then
        log_info "Seguindo todos os logs (Ctrl+C para parar)..."
        echo

        # Para follow, mostra cada tipo em paralelo
        (show_docker_logs_only true "$lines") &
        local docker_pid=$!

        (show_cli_logs_only true "$lines") &
        local cli_pid=$!

        (show_latex_logs_only true "$lines") &
        local latex_pid=$!

        # Aguarda um dos processos terminar
        wait $docker_pid $cli_pid $latex_pid

        # Termina os outros processos
        kill $docker_pid $cli_pid $latex_pid 2>/dev/null || true
    else
        # Para visualização estática, mostra sequencialmente
        show_docker_logs_only false "$lines"
        echo
        show_cli_logs_only false "$lines"
        echo
        show_latex_logs_only false "$lines"
    fi
}
