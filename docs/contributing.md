# 🤝 Contribuindo para o LaTeX Docker Environment

Obrigado pelo interesse em contribuir! Este guia ajudará você a começar.

## 🚀 Como Contribuir

### 🐛 Reportar Bugs

1. Verifique se o bug já não foi reportado nas [issues](https://github.com/martinsmiguel/latex-docker-env/issues)
2. Crie uma nova issue incluindo:
   - Descrição clara do problema
   - Passos para reproduzir
   - Sistema operacional e versões
   - Logs relevantes (`./bin/ltx logs`)

### 💡 Sugerir Melhorias

1. Abra uma [Discussion](https://github.com/martinsmiguel/latex-docker-env/discussions) para discutir a ideia
2. Se aprovada, crie uma issue detalhada
3. Implemente a melhoria seguindo este guia

### 🔧 Contribuir com Código

1. **Fork** o repositório
2. **Clone** sua fork:
   ```bash
   git clone https://github.com/seu-usuario/latex-docker-env.git
   cd latex-docker-env
   ```
3. **Crie uma branch** para sua feature:
   ```bash
   git checkout -b feature/minha-feature
   ```
4. **Faça suas alterações**
5. **Teste** suas mudanças
6. **Commit** com mensagem descritiva:
   ```bash
   git commit -m "feat: adicionar comando xyz"
   ```
7. **Push** para sua branch:
   ```bash
   git push origin feature/minha-feature
   ```
8. **Abra um Pull Request**

## 🏗️ Desenvolvimento

### Ambiente de Desenvolvimento

```bash
# Setup inicial
git clone https://github.com/martinsmiguel/latex-docker-env.git
cd latex-docker-env

# Configurar ambiente
./bin/ltx setup

# Para desenvolvimento da CLI Go
cd cli/
make setup-dev
make dev
```

### Estrutura do Projeto

```
latex-docker-env/
├── bin/                    # Executáveis
├── cli/                    # CLI Go (desenvolvimento principal)
├── lib/                    # CLI Bash legada
├── config/                 # Configurações e templates
├── docs/                   # Documentação
└── tests/                  # Testes
```

### Desenvolvimento da CLI Go

A CLI moderna está em `cli/` e usa Go + Cobra:

```bash
cd cli/

# Instalar dependências
go mod download

# Desenvolvimento com hot reload
make dev

# Executar testes
make test

# Build local
make build

# Linting
make lint
```

### Testando Mudanças

```bash
# Testes automatizados
make test

# Teste manual
./bin/ltx setup
./bin/ltx init --title "Teste"
./bin/ltx build
```

## 📝 Padrões de Código

### Mensagens de Commit

Siga o padrão [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: adicionar novo comando
fix: corrigir bug na compilação
docs: atualizar README
style: formatação de código
refactor: reestruturar função
test: adicionar testes
chore: atualizar dependências
```

### Código Go

- Use `gofmt` para formatação
- Siga as [Go Code Review Guidelines](https://github.com/golang/go/wiki/CodeReviewComments)
- Adicione testes para novas funcionalidades
- Documente funções públicas

### Código Bash

- Use `shellcheck` para validação
- Siga o [Google Shell Style Guide](https://google.github.io/styleguide/shellguide.html)
- Teste em diferentes shells (bash, zsh)

### Documentação

- Use Markdown para documentação
- Inclua exemplos práticos
- Mantenha linguagem clara e concisa
- Adicione emojis para melhor legibilidade

## 🧪 Testes

### Executando Testes

```bash
# Todos os testes
make test

# Testes da CLI Go
cd cli && make test

# Testes da CLI Bash
./tests/run_tests.sh

# Testes de integração
./bin/ltx setup
./tests/integration/test_full_workflow.sh
```

### Adicionando Testes

#### Para CLI Go:
```go
// cli/internal/commands/exemplo_test.go
func TestExemplo(t *testing.T) {
    // Teste aqui
}
```

#### Para CLI Bash:
```bash
# tests/unit/test_novo_comando.bats
@test "novo comando funciona" {
    run ./bin/latex-cli novo-comando
    [ "$status" -eq 0 ]
}
```

## 📚 Documentação

### Atualizando Docs

- **README.md**: Informações principais
- **docs/**: Documentação detalhada
- **cli/**: Documentação específica da CLI Go

### Padrões da Documentação

1. **Use headers consistentes** com emojis
2. **Inclua exemplos práticos** em cada seção
3. **Mantenha links atualizados**
4. **Teste comandos documentados**

## 🔄 Processo de Review

### Pull Requests

1. **Descrição clara** do que foi alterado
2. **Referência issues** relacionadas
3. **Screenshots** se relevante
4. **Testes passando**
5. **Documentação atualizada**

### Critérios de Aprovação

- ✅ Código segue padrões estabelecidos
- ✅ Testes passam
- ✅ Documentação atualizada
- ✅ Não quebra funcionalidades existentes
- ✅ Performance não é degradada

## 🌟 Tipos de Contribuição

### 🔧 Código

- Novos comandos para CLI
- Melhorias de performance
- Correções de bugs
- Refatoração

### 📖 Documentação

- Corrigir typos
- Melhorar exemplos
- Adicionar tutoriais
- Traduzir documentação

### 🧪 Testes

- Adicionar casos de teste
- Melhorar cobertura
- Testes de integração

### 🎨 UX/UI

- Melhorar mensagens da CLI
- Output mais claro
- Melhor tratamento de erros

## 📞 Contato

- **Issues**: Para bugs e features
- **Discussions**: Para dúvidas e ideias
- **Email**: Para questões privadas

## 📄 Licença

Ao contribuir, você concorda que suas contribuições serão licenciadas sob a [MIT License](../LICENSE).

---

**Obrigado por contribuir! 🚀**
