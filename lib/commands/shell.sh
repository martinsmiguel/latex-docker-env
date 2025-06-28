#!/bin/bash
#
# Comando: shell
# Abre shell no container LaTeX
#

cmd_shell() {
    local user=""
    local command=""

    # Parse de argumentos
    while [[ $# -gt 0 ]]; do
        case $1 in
            --user|-u)
                user="$2"
                shift 2
                ;;
            --command|-c)
                command="$2"
                shift 2
                ;;
            --help|-h)
                show_shell_help
                return 0
                ;;
            *)
                log_error "Opção desconhecida: $1"
                show_shell_help
                return 1
                ;;
        esac
    done

    # Verifica se container está executando
    if ! is_container_running; then
        log_info "Container não está executando. Iniciando..."
        start_docker_env || return 1
    fi

    local container_name="$(get_config container_name)"

    if [[ -n "$command" ]]; then
        log_info "Executando comando no container: $command"
        docker_exec_no_tty "$container_name" bash -c "$command"
    else
        log_info "Abrindo shell no container $container_name"
        log_info "Digite 'exit' para sair"

        if [[ -n "$user" ]]; then
            docker exec -it --user "$user" "$container_name" bash
        else
            docker_exec "$container_name" bash
        fi
    fi
}

show_shell_help() {
    cat << EOF
USAGE:
    latex-cli shell [OPTIONS]

DESCRIPTION:
    Abre um shell interativo no container LaTeX ou executa um comando específico.

    Útil para:
    - Instalar pacotes LaTeX adicionais
    - Debugging de problemas de compilação
    - Execução de comandos LaTeX diretamente
    - Exploração do ambiente

OPTIONS:
    -u, --user USER         Usuário para executar o shell (padrão: latexuser)
    -c, --command CMD       Executa comando específico em vez de shell interativo
    -h, --help              Mostra esta ajuda

EXAMPLES:
    latex-cli shell
    latex-cli shell --user root
    latex-cli shell --command "tlmgr update --all"
    latex-cli shell --command "pdflatex --version"

COMANDOS ÚTEIS NO SHELL:
    tlmgr install <pacote>  # Instala pacote LaTeX
    tlmgr update --all      # Atualiza todos os pacotes
    kpsewhich <arquivo>     # Localiza arquivo LaTeX
    texdoc <pacote>         # Documentação de pacote

EOF
}
