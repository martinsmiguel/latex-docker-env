#!/bin/bash
#
# Comando: setup
# Configura o ambiente inicial do projeto
#

cmd_setup() {
    local force=false

    # Parse de argumentos
    while [[ $# -gt 0 ]]; do
        case $1 in
            --force|-f)
                force=true
                shift
                ;;
            --help|-h)
                show_setup_help
                return 0
                ;;
            *)
                log_error "Opção desconhecida: $1"
                show_setup_help
                return 1
                ;;
        esac
    done

    log_info "Configurando ambiente LaTeX Template..."

    # Verifica dependências
    if ! check_dependencies; then
        return 1
    fi

    # Cria estrutura de diretórios
    create_directory_structure "$force"

    # Cria arquivo de configuração se não existir
    local config_file="${PROJECT_ROOT}/config/latex-cli.conf"
    if [[ ! -f "$config_file" || "$force" == "true" ]]; then
        create_default_config "$config_file"
    fi

    # Configura symlinks
    setup_symlinks

    # Configura autocompletion
    setup_autocompletion

    # Inicia ambiente Docker
    if ask_yes_no "Deseja iniciar o ambiente Docker agora?" "y"; then
        start_docker_env
    fi

    log_success "Configuração concluída com sucesso!"
    log_info "Para começar, use: latex-cli init"
}

show_setup_help() {
    cat << EOF
USAGE:
    latex-cli setup [OPTIONS]

DESCRIPTION:
    Configura o ambiente inicial do projeto LaTeX Template.

    Este comando:
    - Verifica dependências (Docker, Docker Compose)
    - Cria estrutura de diretórios necessária
    - Configura arquivos de configuração
    - Configura autocompletion para a CLI
    - Opcionalmente inicia o ambiente Docker

OPTIONS:
    -f, --force     Força reconfiguração mesmo se arquivos já existirem
    -h, --help      Mostra esta ajuda

EXAMPLES:
    latex-cli setup
    latex-cli setup --force

EOF
}

check_dependencies() {
    log_info "Verificando dependências..."

    local all_ok=true

    if ! check_docker; then
        all_ok=false
    fi

    if ! check_docker_compose; then
        all_ok=false
    fi

    if [[ "$all_ok" == "true" ]]; then
        log_success "Todas as dependências estão disponíveis"
        return 0
    else
        log_error "Algumas dependências estão faltando"
        return 1
    fi
}

create_directory_structure() {
    local force="$1"

    log_info "Criando estrutura de diretórios..."

    local dirs=(
        "$(get_config source_dir)"
        "$(get_config output_dir)"
        "${PROJECT_ROOT}/$(get_config output_dir)/logs"
        "${PROJECT_ROOT}/config/vscode"
        "${PROJECT_ROOT}/docs"
    )

    for dir in "${dirs[@]}"; do
        ensure_dir "$dir"
    done

    # Cria .vscode se não existir
    local vscode_dir="${PROJECT_ROOT}/.vscode"
    if [[ ! -d "$vscode_dir" ]]; then
        ln -sf "config/vscode/vscode" "$vscode_dir"
        log_info "Criado symlink para configurações do VS Code"
    fi
}

setup_symlinks() {
    log_info "Configurando symlinks..."

    # Link para o executável principal
    local bin_link="/usr/local/bin/latex-cli"
    if [[ ! -L "$bin_link" ]] && ask_yes_no "Deseja criar symlink global para latex-cli?" "y"; then
        if sudo ln -sf "${PROJECT_ROOT}/bin/latex-cli" "$bin_link"; then
            log_success "Symlink global criado: $bin_link"
        else
            log_warn "Não foi possível criar symlink global"
        fi
    fi
}

setup_autocompletion() {
    log_info "Configurando autocompletion..."

    local completions_dir="${PROJECT_ROOT}/config/completions"

    # Bash completion
    local bash_completion="/etc/bash_completion.d/latex-cli"
    if [[ -d "/etc/bash_completion.d" ]] && ask_yes_no "Configurar autocompletion para Bash?" "y"; then
        if sudo cp "$completions_dir/latex-cli.bash" "$bash_completion"; then
            log_success "Autocompletion do Bash configurado"
        fi
    fi

    # Zsh completion
    if [[ -n "${ZSH_VERSION:-}" ]] && ask_yes_no "Configurar autocompletion para Zsh?" "y"; then
        local zsh_completion="${HOME}/.oh-my-zsh/completions/_latex-cli"
        if [[ -d "$(dirname "$zsh_completion")" ]]; then
            cp "$completions_dir/_latex-cli" "$zsh_completion"
            log_success "Autocompletion do Zsh configurado"
        fi
    fi
}
