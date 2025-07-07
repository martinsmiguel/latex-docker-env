# 🔄 Guia de Migração

Guia para migrar de versões anteriores do LaTeX Docker Environment.

## Migração da v1.x para v2.0

### 🆕 O que há de novo na v2.0

- **CLI moderna em Go** (`ltx`) com melhor performance
- **Suporte Windows nativo** sem necessidade de WSL2
- **File watching otimizado** para desenvolvimento
- **Sistema de configuração robusto**
- **Melhor tratamento de erros**

### 🔄 Mudanças Principais

#### CLI Command Changes

| v1.x (Bash) | v2.0 (Go) | Status |
|-------------|-----------|--------|
| `./bin/latex-cli` | `./bin/ltx` | ✅ Recomendado |
| `latex-cli setup` | `ltx setup` | ✅ Funcional |
| `latex-cli init` | `ltx init` | ✅ Melhorado |
| `latex-cli build` | `ltx build` | ✅ Mais rápido |
| `latex-cli watch` | `ltx watch` | ✅ Otimizado |

#### Configuração

- **v1.x**: Configuração via variáveis de ambiente
- **v2.0**: Arquivo `config/latex-cli.conf` estruturado

#### Estrutura de Arquivos

A estrutura básica permanece a mesma:
```
├── src/          # Seus arquivos LaTeX
├── dist/         # PDFs compilados
├── config/       # Configurações
└── bin/          # CLIs (agora com ltx e latex-cli)
```

## 📋 Checklist de Migração

### 1. Backup do Projeto Atual
```bash
# Faça backup dos seus arquivos LaTeX
cp -r src/ src-backup/
cp -r dist/ dist-backup/
```

### 2. Atualizar o Repositório
```bash
# Pull das últimas atualizações
git pull origin main

# Ou clonar nova versão
git clone https://github.com/martinsmiguel/latex-docker-env.git latex-docker-env-v2
```

### 3. Migrar Configurações

#### Configuração Antiga (v1.x)
```bash
# Variáveis no .env ou shell
export LATEX_AUTHOR="João Silva"
export LATEX_TEMPLATE="book"
```

#### Configuração Nova (v2.0)
```ini
# config/latex-cli.conf
[default]
author = João Silva
template = book
output_dir = dist
```

### 4. Testar Nova CLI
```bash
# Usar nova CLI
./bin/ltx setup
./bin/ltx status

# CLI legada ainda funciona
./bin/latex-cli setup
```

### 5. Migrar Scripts/Automação

#### Antes (v1.x)
```bash
#!/bin/bash
./bin/latex-cli setup
./bin/latex-cli init --title "Documento"
./bin/latex-cli build
```

#### Depois (v2.0)
```bash
#!/bin/bash
./bin/ltx setup
./bin/ltx init --title "Documento" --author "Nome"
./bin/ltx build
```

## 🔧 Resolução de Problemas

### CLI Legacy vs Moderna

**Problema**: Comandos antigos não funcionam
```bash
# ❌ Não funciona mais
latex-cli setup

# ✅ Use o path completo ou nova CLI
./bin/latex-cli setup  # Legacy
./bin/ltx setup        # Moderna (recomendado)
```

### Configurações Não Aplicadas

**Problema**: Configurações personalizadas não funcionam

**Solução**: Migrar para o novo formato
```bash
# Verificar configuração atual
./bin/ltx status --verbose

# Editar configuração
nano config/latex-cli.conf
```

### Performance de Watch Mode

**Problema**: Watch mode muito lento

**Solução**: Usar nova CLI com debounce otimizado
```bash
# ❌ CLI legada (mais lenta)
./bin/latex-cli watch

# ✅ CLI moderna (otimizada)
./bin/ltx watch --debounce 500
```

### Docker Issues

**Problema**: Container não encontrado após atualização

**Solução**: Recriar ambiente
```bash
./bin/ltx clean
./bin/ltx setup
```

## 🚀 Aproveitando Recursos v2.0

### 1. Windows Nativo
```powershell
# Agora funciona diretamente no PowerShell
.\bin\ltx.exe setup
.\bin\ltx.exe init --title "Documento Windows"
```

### 2. Configuração Avançada
```ini
# config/latex-cli.conf
[default]
author = Seu Nome
template = book
output_dir = dist

[build]
clean_before = true
verbose = false

[watch]
debounce = 500
ignore_patterns = ["*.tmp", "*.aux"]
```

### 3. Better Error Handling
```bash
# A nova CLI fornece erros mais claros
./bin/ltx build --verbose
# Agora mostra exatamente onde LaTeX falhou
```

### 4. Autocompletion Melhorado
```bash
# Setup autocompletion para nova CLI
./bin/ltx completion bash > /etc/bash_completion.d/ltx
./bin/ltx completion zsh > /usr/local/share/zsh/site-functions/_ltx
```

## 🔄 Rollback (se necessário)

Se precisar voltar para v1.x temporariamente:

```bash
# 1. Usar CLI legada
./bin/latex-cli setup

# 2. Ou checkout da versão anterior
git checkout v1.4.0  # última versão estável v1.x

# 3. Para projetos novos, clonar versão específica
git clone -b v1.4.0 https://github.com/martinsmiguel/latex-docker-env.git
```

## 📞 Suporte

### Problemas na Migração

1. **Verifique o [FAQ](faq.md)** primeiro
2. **Abra uma [issue](https://github.com/martinsmiguel/latex-docker-env/issues)** com:
   - Versão anterior que estava usando
   - Comando específico que falhou
   - Logs completos (`./bin/ltx logs`)
   - Sistema operacional

### Dúvidas

- **[Discussions](https://github.com/martinsmiguel/latex-docker-env/discussions)** para perguntas gerais
- **[CLI Reference](cli-reference.md)** para comandos específicos

---

**✅ Migração concluída!** Agora você pode aproveitar todos os recursos da v2.0.
