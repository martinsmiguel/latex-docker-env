#!/bin/bash
# scripts/compile.sh

# Verifica se os arquivos necessários existem
if [ ! -f "tex/main.tex" ]; then
    echo "❌ Erro: tex/main.tex não encontrado!"
    echo "Execute primeiro: ./scripts/init_project.sh"
    exit 1
fi

if [ ! -f "tex/preamble.tex" ]; then
    echo "⚠️  preamble.tex não encontrado, copiando template..."
    cp templates/preamble.tex.tpl tex/preamble.tex
fi

# Cria diretório de saída se não existir
mkdir -p out

echo "🔨 Compilando documento LaTeX..."

# Primeira tentativa de compilação usando Docker
if docker exec latex-env latexmk -pdf -f -interaction=nonstopmode -synctex=1 -outdir=./out tex/main.tex; then
    echo "✅ Compilação bem-sucedida!"
    echo "📄 PDF gerado: out/main.pdf"
else
    echo "❌ Falha na compilação. Verifique os logs em out/main.log"
    exit 1
fi