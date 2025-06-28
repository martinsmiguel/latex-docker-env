#!/bin/bash
#
# Comando: init
# Inicializa um novo documento LaTeX
#

cmd_init() {
    local title=""
    local author=""
    local template="default"
    local force=false

    # Parse de argumentos
    while [[ $# -gt 0 ]]; do
        case $1 in
            --title|-t)
                title="$2"
                shift 2
                ;;
            --author|-a)
                author="$2"
                shift 2
                ;;
            --template)
                template="$2"
                shift 2
                ;;
            --force|-f)
                force=true
                shift
                ;;
            --help|-h)
                show_init_help
                return 0
                ;;
            *)
                log_error "Opção desconhecida: $1"
                show_init_help
                return 1
                ;;
        esac
    done

    log_info "Inicializando novo documento LaTeX..."

    # Verifica se já existe projeto inicializado
    local source_dir="$(get_config source_dir)"
    if [[ -f "$source_dir/main.tex" && "$force" != "true" ]]; then
        if ! ask_yes_no "Já existe um documento. Deseja sobrescrever?" "n"; then
            log_info "Operação cancelada"
            return 0
        fi
    fi

    # Coleta informações se não foram fornecidas
    if is_empty "$title"; then
        title=$(prompt_with_default "Título do documento" "Meu Documento LaTeX")
    fi

    if is_empty "$author"; then
        author=$(prompt_with_default "Nome do autor" "$(git config user.name 2>/dev/null || echo 'Autor')")
    fi

    # Valida entradas
    if ! validate_title "$title" || ! validate_author "$author"; then
        return 1
    fi

    # Cria estrutura do projeto
    create_project_structure "$title" "$author" "$template"

    log_success "Documento LaTeX inicializado com sucesso!"
    log_info "Arquivos criados em: $source_dir"
    log_info "Para compilar: latex-cli build"
}

show_init_help() {
    cat << EOF
USAGE:
    latex-cli init [OPTIONS]

DESCRIPTION:
    Inicializa um novo documento LaTeX a partir de templates.

    Este comando cria:
    - Arquivo principal main.tex
    - Arquivo de preâmbulo preamble.tex
    - Estrutura de capítulos
    - Arquivo de bibliografia references.bib
    - Configurações do VS Code

OPTIONS:
    -t, --title TITLE       Título do documento
    -a, --author AUTHOR     Nome do autor
    --template TEMPLATE     Template a usar (default: default)
    -f, --force             Sobrescreve arquivos existentes
    -h, --help              Mostra esta ajuda

TEMPLATES DISPONÍVEIS:
    default     Template básico para documentos gerais
    article     Template para artigos científicos
    thesis      Template para teses e dissertações
    report      Template para relatórios técnicos
    book        Template para livros

EXAMPLES:
    latex-cli init
    latex-cli init --title "Minha Tese" --author "João Silva"
    latex-cli init --template thesis --title "Dissertação de Mestrado"

EOF
}

create_project_structure() {
    local title="$1"
    local author="$2"
    local template="$3"

    local source_dir="$(get_config source_dir)"
    local templates_dir="${PROJECT_ROOT}/config/templates"

    # Cria diretório fonte se não existir
    ensure_dir "$source_dir"
    ensure_dir "$source_dir/chapters"
    ensure_dir "$source_dir/images"

    log_info "Criando arquivos do template '$template'..."

    # Escapa strings para LaTeX
    local safe_title="$(escape_latex "$title")"
    local safe_author="$(escape_latex "$author")"

    # Cria main.tex
    create_main_tex "$source_dir" "$safe_title" "$safe_author" "$template"

    # Cria preamble.tex
    create_preamble_tex "$source_dir" "$template"

    # Cria structure de capítulos
    create_chapters_structure "$source_dir" "$template"

    # Cria bibliography
    create_bibliography "$source_dir"

    # Cria configurações do VS Code
    create_vscode_settings
}

create_main_tex() {
    local source_dir="$1"
    local title="$2"
    local author="$3"
    local template="$4"

    local main_file="$source_dir/main.tex"
    local template_file="${PROJECT_ROOT}/config/templates/main.tex.tpl"

    if [[ -f "$template_file" ]]; then
        # Usa template existente e substitui variáveis
        sed -e "s/{{TITLE}}/$title/g" \
            -e "s/{{AUTHOR}}/$author/g" \
            -e "s/{{DATE}}/$(date +'%B %Y')/g" \
            "$template_file" > "$main_file"
    else
        # Cria template básico
        cat > "$main_file" << EOF
\\documentclass[12pt,a4paper]{article}

\\input{preamble}

\\title{$title}
\\author{$author}
\\date{$(date +'%B %Y')}

\\begin{document}

\\maketitle
\\tableofcontents
\\newpage

\\input{chapters/introduction}
\\input{chapters/methodology}
\\input{chapters/results}
\\input{chapters/conclusion}

\\bibliography{references}
\\bibliographystyle{plain}

\\end{document}
EOF
    fi

    log_success "Criado: $main_file"
}

create_preamble_tex() {
    local source_dir="$1"
    local template="$2"

    local preamble_file="$source_dir/preamble.tex"
    local template_file="${PROJECT_ROOT}/config/templates/preamble.tex.tpl"

    if [[ -f "$template_file" ]]; then
        cp "$template_file" "$preamble_file"
    else
        cat > "$preamble_file" << 'EOF'
% Configurações do documento
\usepackage[utf8]{inputenc}
\usepackage[T1]{fontenc}
\usepackage[brazil]{babel}

% Pacotes essenciais
\usepackage{amsmath,amsfonts,amssymb}
\usepackage{graphicx}
\usepackage{hyperref}
\usepackage{url}
\usepackage{enumerate}
\usepackage{fancyhdr}

% Configurações de página
\usepackage[a4paper,margin=2.5cm]{geometry}
\setlength{\parindent}{1.25cm}
\setlength{\parskip}{0pt}

% Configurações de hyperref
\hypersetup{
    colorlinks=true,
    linkcolor=blue,
    filecolor=magenta,
    urlcolor=cyan,
    citecolor=red
}

% Cabeçalhos e rodapés
\pagestyle{fancy}
\fancyhf{}
\rhead{\thepage}
\lhead{\leftmark}
EOF
    fi

    log_success "Criado: $preamble_file"
}

create_chapters_structure() {
    local source_dir="$1"
    local template="$2"

    local chapters_dir="$source_dir/chapters"

    # Cria capítulos básicos
    local chapters=("introduction" "methodology" "results" "conclusion")

    for chapter in "${chapters[@]}"; do
        local chapter_file="$chapters_dir/${chapter}.tex"

        if [[ ! -f "$chapter_file" ]]; then
            case "$chapter" in
                "introduction")
                    cat > "$chapter_file" << 'EOF'
\section{Introdução}

Este é o capítulo de introdução do documento.

\subsection{Objetivos}

Descreva aqui os objetivos do trabalho.

\subsection{Estrutura do Documento}

Descreva a organização do documento.
EOF
                    ;;
                "methodology")
                    cat > "$chapter_file" << 'EOF'
\section{Metodologia}

Descreva a metodologia utilizada no trabalho.

\subsection{Materiais}

Liste os materiais utilizados.

\subsection{Métodos}

Descreva os métodos aplicados.
EOF
                    ;;
                "results")
                    cat > "$chapter_file" << 'EOF'
\section{Resultados}

Apresente os resultados obtidos.

\subsection{Análise dos Dados}

Faça a análise dos dados coletados.

\subsection{Discussão}

Discuta os resultados encontrados.
EOF
                    ;;
                "conclusion")
                    cat > "$chapter_file" << 'EOF'
\section{Conclusão}

Apresente as conclusões do trabalho.

\subsection{Trabalhos Futuros}

Sugira possíveis trabalhos futuros.
EOF
                    ;;
            esac

            log_success "Criado: $chapter_file"
        fi
    done
}

create_bibliography() {
    local source_dir="$1"
    local bib_file="$source_dir/references.bib"

    if [[ ! -f "$bib_file" ]]; then
        cat > "$bib_file" << 'EOF'
@article{example2023,
    title={An Example Article},
    author={Author, Example},
    journal={Journal of Examples},
    volume={1},
    number={1},
    pages={1--10},
    year={2023},
    publisher={Example Publisher}
}

@book{knuth1984,
    title={The {\TeX}book},
    author={Knuth, Donald E.},
    year={1984},
    publisher={Addison-Wesley}
}
EOF
        log_success "Criado: $bib_file"
    fi
}

create_vscode_settings() {
    local vscode_dir="${PROJECT_ROOT}/config/vscode/vscode"
    ensure_dir "$vscode_dir"

    local settings_file="$vscode_dir/settings.json"
    local template_file="${PROJECT_ROOT}/config/templates/settings.json.tpl"

    if [[ -f "$template_file" ]]; then
        cp "$template_file" "$settings_file"
        log_success "Configurações do VS Code criadas"
    fi
}
