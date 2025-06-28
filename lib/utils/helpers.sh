#!/bin/bash
#
# Funções utilitárias
#

# Verifica se uma string é vazia
is_empty() {
    [[ -z "${1:-}" ]]
}

# Verifica se um arquivo existe e é legível
file_readable() {
    [[ -r "$1" ]]
}

# Verifica se um diretório existe e é acessível
dir_accessible() {
    [[ -d "$1" && -x "$1" ]]
}

# Cria diretório se não existir
ensure_dir() {
    local dir="$1"

    if [[ ! -d "$dir" ]]; then
        log_debug "Criando diretório: $dir"
        mkdir -p "$dir"
    fi
}

# Remove arquivo ou diretório se existir
safe_remove() {
    local path="$1"

    if [[ -e "$path" ]]; then
        log_debug "Removendo: $path"
        rm -rf "$path"
    fi
}

# Copia arquivo preservando permissões
safe_copy() {
    local source="$1"
    local dest="$2"

    if [[ ! -f "$source" ]]; then
        log_error "Arquivo fonte não encontrado: $source"
        return 1
    fi

    log_debug "Copiando $source para $dest"
    cp -p "$source" "$dest"
}

# Verifica se uma versão é compatível
version_compatible() {
    local version="$1"
    local min_version="$2"

    # Comparação simples de versão (funciona para a maioria dos casos)
    if [[ "$(printf '%s\n' "$min_version" "$version" | sort -V | head -n1)" == "$min_version" ]]; then
        return 0
    else
        return 1
    fi
}

# Pergunta sim/não para o usuário
ask_yes_no() {
    local question="$1"
    local default="${2:-}"

    while true; do
        if [[ "$default" == "y" ]]; then
            read -rp "$question [Y/n]: " answer
            answer="${answer:-y}"
        elif [[ "$default" == "n" ]]; then
            read -rp "$question [y/N]: " answer
            answer="${answer:-n}"
        else
            read -rp "$question [y/n]: " answer
        fi

        case "$(echo "$answer" | tr '[:upper:]' '[:lower:]')" in
            y|yes) return 0 ;;
            n|no) return 1 ;;
            *) echo "Por favor, responda 'y' ou 'n'." ;;
        esac
    done
}

# Solicita entrada do usuário com valor padrão
prompt_with_default() {
    local prompt="$1"
    local default="$2"
    local result

    read -rp "$prompt [$default]: " result
    echo "${result:-$default}"
}

# Valida se uma string é um título válido
validate_title() {
    local title="$1"

    if is_empty "$title"; then
        log_error "Título não pode estar vazio"
        return 1
    fi

    if [[ ${#title} -gt 200 ]]; then
        log_error "Título muito longo (máximo 200 caracteres)"
        return 1
    fi

    return 0
}

# Valida se uma string é um nome de autor válido
validate_author() {
    local author="$1"

    if is_empty "$author"; then
        log_error "Nome do autor não pode estar vazio"
        return 1
    fi

    if [[ ${#author} -gt 100 ]]; then
        log_error "Nome do autor muito longo (máximo 100 caracteres)"
        return 1
    fi

    return 0
}

# Escapar caracteres especiais para LaTeX
escape_latex() {
    local text="$1"

    # Escapar caracteres especiais do LaTeX
    text="${text//\\/\\textbackslash}"
    text="${text//&/\\&}"
    text="${text//%/\\%}"
    text="${text//\$/\\$}"
    text="${text//#/\\#}"
    text="${text//^/\\textasciicircum}"
    text="${text//_/\\_}"
    text="${text//\{/\\{}"
    text="${text//\}/\\}"
    text="${text//~/\\textasciitilde}"

    echo "$text"
}

# Converte caminho relativo para absoluto
abs_path() {
    local path="$1"

    if [[ "$path" == /* ]]; then
        echo "$path"
    else
        echo "$(cd "$(dirname "$path")" && pwd)/$(basename "$path")"
    fi
}

# Encontra arquivos LaTeX no diretório
find_latex_files() {
    local dir="$1"

    find "$dir" -name "*.tex" -type f 2>/dev/null | sort
}

# Encontra arquivo principal LaTeX
find_main_latex_file() {
    local source_dir="${1:-$(get_config source_dir)}"

    # Verifica se existe main.tex
    if [[ -f "$source_dir/main.tex" ]]; then
        echo "$source_dir/main.tex"
        return 0
    fi

    # Procura por arquivo com documentclass
    local files
    files=($(find_latex_files "$source_dir"))

    for file in "${files[@]}"; do
        if grep -q "\\documentclass" "$file" 2>/dev/null; then
            echo "$file"
            return 0
        fi
    done

    log_error "Nenhum arquivo principal LaTeX encontrado em $source_dir"
    return 1
}

# Sanitiza nome de arquivo
sanitize_filename() {
    local filename="$1"

    # Remove/substitui caracteres problemáticos
    filename="${filename//[^a-zA-Z0-9._-]/_}"
    filename="${filename#_}"  # Remove _ do início
    filename="${filename%_}"  # Remove _ do final

    echo "$filename"
}

# Mostra tempo decorrido
show_elapsed_time() {
    local start_time="$1"
    local end_time="${2:-$(date +%s)}"
    local elapsed=$((end_time - start_time))

    if [[ $elapsed -lt 60 ]]; then
        echo "${elapsed}s"
    elif [[ $elapsed -lt 3600 ]]; then
        echo "$((elapsed / 60))m $((elapsed % 60))s"
    else
        echo "$((elapsed / 3600))h $(((elapsed % 3600) / 60))m $((elapsed % 60))s"
    fi
}
