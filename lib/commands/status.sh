#!/bin/bash
#
# Comando: status
# Mostra o status do ambiente LaTeX
#

cmd_status() {
    local verbose=false

    # Parse de argumentos
    while [[ $# -gt 0 ]]; do
        case $1 in
            --verbose|-v)
                verbose=true
                shift
                ;;
            --help|-h)
                show_status_help
                return 0
                ;;
            *)
                log_error "Opção desconhecida: $1"
                show_status_help
                return 1
                ;;
        esac
    done

    echo "=== Status do LaTeX Template CLI ==="
    echo

    # Status da CLI
    show_cli_status
    echo

    # Status do Docker
    docker_status
    echo

    # Status do projeto
    show_project_status
    echo

    # Status detalhado se solicitado
    if [[ "$verbose" == "true" ]]; then
        show_detailed_status
    fi
}

show_status_help() {
    cat << EOF
USAGE:
    latex-cli status [OPTIONS]

DESCRIPTION:
    Mostra informações sobre o status do ambiente LaTeX Template.

    Inclui informações sobre:
    - Versão da CLI e configurações
    - Status do Docker e containers
    - Status do projeto LaTeX
    - Arquivos e estrutura do projeto

OPTIONS:
    -v, --verbose    Mostra informações detalhadas
    -h, --help       Mostra esta ajuda

EXAMPLES:
    latex-cli status
    latex-cli status --verbose

EOF
}

show_cli_status() {
    echo "=== LaTeX CLI ==="
    echo "Versão: $(get_config_version)"
    echo "Diretório do projeto: $PROJECT_ROOT"
    echo "Arquivo de configuração: $(get_config_file_path)"

    # Configurações principais
    echo "Configurações:"
    echo "  Engine LaTeX: $(get_config latex_engine)"
    echo "  Diretório fonte: $(get_config source_dir)"
    echo "  Diretório de saída: $(get_config output_dir)"
    echo "  Container: $(get_config container_name)"
}

show_project_status() {
    echo "=== Status do Projeto ==="

    local source_dir="$(get_config source_dir)"
    local output_dir="$(get_config output_dir)"

    # Verifica se projeto foi inicializado
    if [[ -f "$source_dir/main.tex" ]]; then
        echo "✓ Projeto inicializado"

        # Informações sobre o documento principal
        local title=$(grep "\\title{" "$source_dir/main.tex" 2>/dev/null | sed 's/.*\\title{\(.*\)}.*/\1/' || echo "Não encontrado")
        local author=$(grep "\\author{" "$source_dir/main.tex" 2>/dev/null | sed 's/.*\\author{\(.*\)}.*/\1/' || echo "Não encontrado")

        echo "  Título: $title"
        echo "  Autor: $author"
    else
        echo "✗ Projeto não inicializado"
        echo "  Use: latex-cli init"
        return
    fi

    # Arquivos LaTeX
    local tex_files
    tex_files=($(find_latex_files "$source_dir"))
    echo "  Arquivos LaTeX: ${#tex_files[@]}"

    # Última compilação
    local pdf_file="$output_dir/main.pdf"
    if [[ -f "$pdf_file" ]]; then
        local pdf_date=$(stat -f "%Sm" -t "%Y-%m-%d %H:%M:%S" "$pdf_file" 2>/dev/null || stat -c "%y" "$pdf_file" 2>/dev/null || echo "Desconhecido")
        echo "✓ PDF disponível (compilado em: $pdf_date)"
    else
        echo "✗ PDF não encontrado"
        echo "  Use: latex-cli build"
    fi

    # Estrutura de capítulos
    local chapters_dir="$source_dir/chapters"
    if [[ -d "$chapters_dir" ]]; then
        local chapter_count=$(find "$chapters_dir" -name "*.tex" | wc -l)
        echo "  Capítulos: $chapter_count"
    fi

    # Bibliografia
    if [[ -f "$source_dir/references.bib" ]]; then
        local bib_entries=$(grep -c "^@" "$source_dir/references.bib" 2>/dev/null || echo "0")
        echo "  Referências bibliográficas: $bib_entries"
    fi
}

show_detailed_status() {
    echo "=== Informações Detalhadas ==="

    # Estrutura de arquivos
    echo
    echo "Estrutura do projeto:"
    tree "${PROJECT_ROOT}" -I 'dist|.git|node_modules' -L 3 2>/dev/null || find "${PROJECT_ROOT}" -type f -name "*.tex" -o -name "*.bib" -o -name "*.yml" -o -name "latex-cli" | head -20

    # Espaço em disco
    echo
    echo "Uso do disco:"
    local output_dir="$(get_config output_dir)"
    if [[ -d "$output_dir" ]]; then
        du -sh "$output_dir" 2>/dev/null || echo "Não disponível"
    else
        echo "Diretório de saída não existe"
    fi

    # Logs recentes
    echo
    echo "Logs recentes:"
    local log_file="${PROJECT_ROOT}/dist/logs/latex-cli.log"
    if [[ -f "$log_file" ]]; then
        tail -5 "$log_file" 2>/dev/null || echo "Nenhum log disponível"
    else
        echo "Arquivo de log não encontrado"
    fi
}

get_config_version() {
    echo "$CLI_VERSION"
}

get_config_file_path() {
    local config_file="${PROJECT_ROOT}/config/latex-cli.conf"
    if [[ -f "$config_file" ]]; then
        echo "$config_file"
    else
        echo "Não encontrado"
    fi
}
