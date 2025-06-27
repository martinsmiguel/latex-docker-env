#!/bin/bash

# Cria diretórios essenciais não versionados
mkdir -p out/logs
mkdir -p .vscode
mkdir -p tex

# Copia templates se os arquivos não existirem
[ -f tex/main.tex ] || cp templates/main.tex.tpl tex/main.tex
[ -f tex/preamble.tex ] || cp templates/preamble.tex.tpl tex/preamble.tex
[ -f .vscode/settings.json ] || cp templates/settings.json.tpl .vscode/settings.json
[ -f tex/references.bib ] || cp templates/references.bib.tpl tex/references.bib

# Configura permissões
chmod +x scripts/*.sh

echo "✅ Ambiente configurado! Execute:"
echo "code . && docker-compose up -d"