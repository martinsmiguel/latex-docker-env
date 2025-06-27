#!/bin/bash
# scripts/clean.sh

echo "🧹 Limpando arquivos auxiliares..."

# Remove arquivos auxiliares da compilação
rm -f out/*.aux out/*.log out/*.fdb_latexmk out/*.fls out/*.synctex.gz
rm -f out/*.bbl out/*.blg out/*.out out/*.toc out/*.lof out/*.lot

echo "✅ Arquivos auxiliares removidos!"
echo "📄 PDF mantido em: out/main.pdf"
