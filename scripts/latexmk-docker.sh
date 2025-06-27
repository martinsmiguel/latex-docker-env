#!/bin/bash
# scripts/latexmk-docker.sh
# Wrapper script para usar latexmk dentro do container Docker

# Passa todos os argumentos para o latexmk dentro do container
docker exec latex-env latexmk "$@"
