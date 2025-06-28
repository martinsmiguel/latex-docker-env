#!/bin/bash
# Bash completion for latex-cli

_latex_cli_completion() {
    local cur prev words cword
    _init_completion || return

    local commands="setup init build watch clean dev stop restart status logs shell exec update templates packages version"

    # Opções globais
    local global_opts="-h --help -v --verbose -q --quiet --no-docker --version"

    case $cword in
        1)
            # Primeiro argumento: comando
            COMPREPLY=($(compgen -W "$commands $global_opts" -- "$cur"))
            return
            ;;
        *)
            # Argumentos específicos do comando
            case ${words[1]} in
                init)
                    local init_opts="--title --author --template --non-interactive --force --help"
                    COMPREPLY=($(compgen -W "$init_opts" -- "$cur"))
                    ;;
                build)
                    local build_opts="--output-dir --clean --engine --no-bib --verbose --help"
                    COMPREPLY=($(compgen -W "$build_opts" -- "$cur"))
                    ;;
                clean)
                    local clean_opts="--all --temp-only --pdf --cache --quiet --help"
                    COMPREPLY=($(compgen -W "$clean_opts" -- "$cur"))
                    ;;
                dev)
                    local dev_opts="--quiet --rebuild --help"
                    COMPREPLY=($(compgen -W "$dev_opts" -- "$cur"))
                    ;;
                logs)
                    local logs_opts="--follow -f --lines -n --help"
                    COMPREPLY=($(compgen -W "$logs_opts latex-env" -- "$cur"))
                    ;;
                exec)
                    # Para exec, completa comandos do sistema
                    local exec_commands="bash ls cat pdflatex latexmk tlmgr"
                    COMPREPLY=($(compgen -W "$exec_commands" -- "$cur"))
                    ;;
                *)
                    # Para outros comandos, apenas --help
                    COMPREPLY=($(compgen -W "--help" -- "$cur"))
                    ;;
            esac
            ;;
    esac
}

# Registra a função de completion
complete -F _latex_cli_completion latex-cli
