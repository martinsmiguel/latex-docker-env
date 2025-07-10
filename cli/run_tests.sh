#!/bin/bash

# Script para executar testes do CLI Go

set -e

echo "🧪 Executando Testes do CLI LaTeX Docker"
echo "========================================"

# Ir para o diretório do CLI
cd "$(dirname "$0")"

echo "📁 Diretório atual: $(pwd)"

# Verificar se estamos no diretório correto
if [[ ! -f "go.mod" ]]; then
    echo "❌ Erro: Não encontrado go.mod. Execute este script do diretório cli/"
    exit 1
fi

echo ""
echo "🔍 Verificando dependências..."
go mod tidy

echo ""
echo "🏗️  Compilando projeto..."
go build -o ltx ./main.go

echo ""
echo "🧪 Executando testes unitários..."
echo "--------------------------------"

# Executar testes com verbose e coverage seguindo padrões do Go
echo "📋 Executando go test com race detection e coverage..."
go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

TEST_EXIT_CODE=$?

echo ""
echo "📊 Relatório de Cobertura:"
echo "-------------------------"
go tool cover -func=coverage.out

echo ""
echo "📈 Relatório de cobertura por pacote:"
echo "------------------------------------"
go tool cover -func=coverage.out | grep -E "(total:|statements)"

echo ""
echo "📈 Gerando relatório HTML de cobertura..."
go tool cover -html=coverage.out -o coverage.html
echo "📄 Relatório HTML salvo em: coverage.html"

echo ""
echo "🔍 Executando go vet (análise estática)..."
go vet ./...
VET_EXIT_CODE=$?

echo ""
echo "🔧 Executando golangci-lint..."
if command -v golangci-lint &> /dev/null; then
    golangci-lint run --config .golangci.yml
    LINT_EXIT_CODE=$?
else
    echo "⚠️  golangci-lint não encontrado."
    echo "   Instale com: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
    LINT_EXIT_CODE=0  # Não falhar se não estiver instalado
fi

echo ""
echo "🧹 Verificando formatação de código..."
UNFORMATTED=$(gofmt -l .)
if [ -n "$UNFORMATTED" ]; then
    echo "❌ Arquivos não formatados encontrados:"
    echo "$UNFORMATTED"
    echo "   Execute: go fmt ./..."
    GOFMT_EXIT_CODE=1
else
    echo "✅ Todos os arquivos estão formatados corretamente"
    GOFMT_EXIT_CODE=0
fi

echo ""
echo "🔍 Executando verificações adicionais..."

# Verificar se há imports não utilizados
echo "📦 Verificando imports..."
if command -v goimports &> /dev/null; then
    UNORGANIZED=$(goimports -l .)
    if [ -n "$UNORGANIZED" ]; then
        echo "⚠️  Imports desorganizados encontrados:"
        echo "$UNORGANIZED"
        echo "   Execute: goimports -w ."
    else
        echo "✅ Imports organizados corretamente"
    fi
else
    echo "⚠️  goimports não encontrado. Instale com: go install golang.org/x/tools/cmd/goimports@latest"
fi

# Verificar dependências do módulo
echo "📋 Verificando dependências do módulo..."
go mod verify
MOD_EXIT_CODE=$?

echo ""
echo "📋 Resumo dos Testes:"
echo "--------------------"

# Calcular cobertura total
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}')
echo "📊 Cobertura total: $COVERAGE"

# Status dos testes
if [ $TEST_EXIT_CODE -eq 0 ]; then
    echo "✅ Testes unitários: PASSOU"
else
    echo "❌ Testes unitários: FALHOU"
fi

if [ $VET_EXIT_CODE -eq 0 ]; then
    echo "✅ go vet: PASSOU"
else
    echo "❌ go vet: FALHOU"
fi

if [ $LINT_EXIT_CODE -eq 0 ]; then
    echo "✅ golangci-lint: PASSOU"
else
    echo "❌ golangci-lint: FALHOU"
fi

if [ $GOFMT_EXIT_CODE -eq 0 ]; then
    echo "✅ Formatação: PASSOU"
else
    echo "❌ Formatação: FALHOU"
fi

if [ $MOD_EXIT_CODE -eq 0 ]; then
    echo "✅ Verificação de módulo: PASSOU"
else
    echo "❌ Verificação de módulo: FALHOU"
fi

echo ""
echo "📄 Arquivos gerados:"
echo "  - coverage.out (perfil de cobertura)"
echo "  - coverage.html (relatório HTML)"

echo ""
echo "🚀 Comandos úteis:"
echo "  - Visualizar cobertura: open coverage.html"
echo "  - Executar apenas um pacote: go test -v ./internal/commands"
echo "  - Executar apenas um teste: go test -v -run TestNomeDoTeste ./internal/commands"
echo "  - Benchmark: go test -bench=. ./..."

# Calcular código de saída final
FINAL_EXIT_CODE=0
if [ $TEST_EXIT_CODE -ne 0 ] || [ $VET_EXIT_CODE -ne 0 ] || [ $LINT_EXIT_CODE -ne 0 ] || [ $GOFMT_EXIT_CODE -ne 0 ] || [ $MOD_EXIT_CODE -ne 0 ]; then
    FINAL_EXIT_CODE=1
fi

# Verificar se há testes falhando
if [ $FINAL_EXIT_CODE -eq 0 ]; then
    echo ""
    echo "🎉 Todos os testes e verificações passaram!"
    echo "✨ Código está seguindo os padrões do Go!"
    exit 0
else
    echo ""
    echo "💥 Algumas verificações falharam. Verifique a saída acima."
    echo "🔧 Execute as correções sugeridas antes de fazer commit."
    exit 1
fi
