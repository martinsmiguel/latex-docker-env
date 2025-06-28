#!/bin/bash
#
# Comando: update
# Atualiza o ambiente LaTeX
#

cmd_update() {
    local packages=false
    local cli=false
    local docker=false
    local all=false

    # Parse de argumentos
    while [[ $# -gt 0 ]]; do
        case $1 in
            --packages|-p)
                packages=true
                shift
                ;;
            --cli|-c)
                cli=true
                shift
                ;;
            --docker|-d)
                docker=true
                shift
                ;;
            --all|-a)
                all=true
                shift
                ;;
            --help|-h)
                show_update_help
                return 0
                ;;
            *)
                log_error "Opção desconhecida: $1"
                show_update_help
                return 1
                ;;
        esac
    done

    # Se --all foi especificado, ativa todas as opções
    if [[ "$all" == "true" ]]; then
        packages=true
        docker=true
    fi

    # Se nenhuma opção específica foi dada, atualiza pacotes por padrão
    if [[ "$packages" == "false" && "$cli" == "false" && "$docker" == "false" ]]; then
        packages=true
    fi

    log_info "Atualizando ambiente LaTeX..."

    if [[ "$docker" == "true" ]]; then
        update_docker_environment
    fi

    if [[ "$packages" == "true" ]]; then
        update_latex_packages
    fi

    if [[ "$cli" == "true" ]]; then
        update_cli
    fi

    log_success "Atualização concluída"
}

show_update_help() {
    cat << EOF
USAGE:
    latex-cli update [OPTIONS]

DESCRIPTION:
    Atualiza componentes do ambiente LaTeX Template.

    Por padrão, atualiza apenas os pacotes LaTeX. Use as opções
    para controlar o que deve ser atualizado.

OPTIONS:
    -p, --packages          Atualiza pacotes LaTeX (padrão)
    -d, --docker            Reconstrói ambiente Docker
    -c, --cli               Atualiza a CLI (git pull)
    -a, --all               Atualiza tudo
    -h, --help              Mostra esta ajuda

COMPONENTES:
    packages    Pacotes TeX Live via tlmgr
    docker      Imagem Docker e dependências
    cli         Código da CLI (via git)

EXAMPLES:
    latex-cli update
    latex-cli update --packages
    latex-cli update --docker
    latex-cli update --all

EOF
}

update_latex_packages() {
    log_info "Atualizando pacotes LaTeX..."

    # Verifica se container está executando
    if ! is_container_running; then
        log_info "Iniciando container..."
        start_docker_env || return 1
    fi

    local container_name="$(get_config container_name)"

    # Atualiza o índice de pacotes
    log_info "Atualizando índice de pacotes..."
    if docker_exec_no_tty "$container_name" tlmgr update --self; then
        log_success "Índice de pacotes atualizado"
    else
        log_warn "Falha ao atualizar índice de pacotes"
    fi

    # Atualiza todos os pacotes
    log_info "Atualizando pacotes instalados..."
    if docker_exec_no_tty "$container_name" tlmgr update --all; then
        log_success "Pacotes LaTeX atualizados"
    else
        log_warn "Alguns pacotes podem não ter sido atualizados"
    fi

    # Mostra estatísticas
    show_package_statistics
}

update_docker_environment() {
    log_info "Atualizando ambiente Docker..."

    if ask_yes_no "Isso irá reconstruir o container. Continuar?" "y"; then
        rebuild_docker_env
    else
        log_info "Atualização do Docker cancelada"
    fi
}

update_cli() {
    log_info "Atualizando CLI..."

    # Verifica se estamos em um repositório git
    if [[ ! -d "${PROJECT_ROOT}/.git" ]]; then
        log_error "Não é um repositório git. Não é possível atualizar automaticamente."
        log_info "Para atualizar manualmente, baixe a versão mais recente do GitHub."
        return 1
    fi

    # Salva mudanças locais se houver
    local has_changes=false
    if ! git diff --quiet HEAD 2>/dev/null; then
        has_changes=true
        log_warn "Há mudanças locais não commitadas"

        if ask_yes_no "Deseja fazer stash das mudanças locais?" "y"; then
            git stash push -m "Mudanças locais antes da atualização da CLI"
            log_info "Mudanças salvas em stash"
        else
            log_error "Não é possível atualizar com mudanças locais"
            return 1
        fi
    fi

    # Atualiza do repositório remoto
    log_info "Fazendo pull do repositório..."
    if git pull origin main; then
        log_success "CLI atualizada"

        # Restaura stash se necessário
        if [[ "$has_changes" == "true" ]]; then
            if ask_yes_no "Deseja restaurar suas mudanças locais?" "y"; then
                git stash pop
                log_info "Mudanças locais restauradas"
            fi
        fi

        # Recarrega configurações
        log_info "Recarregando configurações..."
        CONFIG_LOADED=false
        load_config

    else
        log_error "Falha ao atualizar CLI"

        # Restaura stash em caso de erro
        if [[ "$has_changes" == "true" ]]; then
            git stash pop
            log_info "Mudanças locais restauradas"
        fi

        return 1
    fi
}

show_package_statistics() {
    local container_name="$(get_config container_name)"

    log_info "Estatísticas dos pacotes:"

    # Número total de pacotes
    local total_packages
    total_packages=$(docker_exec_no_tty "$container_name" tlmgr list --only-installed | wc -l 2>/dev/null || echo "Desconhecido")
    echo "  Pacotes instalados: $total_packages"

    # Último update
    local last_update
    last_update=$(docker_exec_no_tty "$container_name" tlmgr option repository 2>/dev/null | head -1 || echo "Desconhecido")
    echo "  Repositório: $last_update"

    # Versão do TeX Live
    local texlive_version
    texlive_version=$(docker_exec_no_tty "$container_name" tex --version 2>/dev/null | head -1 || echo "Desconhecido")
    echo "  TeX Live: $texlive_version"
}
