#!/bin/bash

# Solicita informações básicas
read -p "Título do documento: " title
read -p "Autor: " author

# Gera main.tex personalizado
sed -e "s/{{TITLE}}/$title/g" -e "s/{{AUTHOR}}/$author/g" \
  templates/main.tex.tpl > tex/main.tex

# Copia preamble se não existir
[ -f tex/preamble.tex ] || cp templates/preamble.tex.tpl tex/preamble.tex

# Cria estrutura de capítulos
mkdir -p tex/chapters
touch tex/chapters/{introducao,metodologia,resultados}.tex

echo "📄 Documento '$title' inicializado!"