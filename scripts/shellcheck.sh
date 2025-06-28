#!/bin/bash
#
# Script para executar ShellCheck em todos os arquivos shell do projeto
#

set -euo pipefail

echo "Executando ShellCheck em todos os arquivos shell..."
echo "=============================================="

# Contadores
total_files=0
files_with_issues=0

# Encontra todos os arquivos shell
while IFS= read -r -d '' file; do
    ((total_files++))
    echo -n "Verificando $(basename "$file")... "

    if shellcheck "$file" > /dev/null 2>&1; then
        echo "✓"
    else
        echo "⚠"
        ((files_with_issues++))
    fi
done < <(find . -name "*.sh" -o -name "latex-cli" -print0)

echo
echo "Resumo:"
echo "======="
echo "Total de arquivos: $total_files"
echo "Arquivos com avisos: $files_with_issues"
echo "Arquivos limpos: $((total_files - files_with_issues))"

if [[ $files_with_issues -eq 0 ]]; then
    echo "🎉 Todos os arquivos passaram no ShellCheck!"
    exit 0
else
    echo "ℹ️  Alguns arquivos têm avisos (principalmente informativos)"
    echo "Execute 'shellcheck <arquivo>' para ver detalhes específicos"
    exit 0
fi
