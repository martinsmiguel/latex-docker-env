# 🛠️ CLI Reference

Documentação completa da CLI `ltx` - a interface moderna em Go para o LaTeX Docker Environment.

## 🚀 Comandos Principais

### `ltx setup`
Configura o ambiente inicial de desenvolvimento.

```bash
ltx setup [flags]

Flags:
  -f, --force     Força reconfiguração mesmo se já configurado
  -q, --quiet     Execução silenciosa
  -h, --help      Ajuda para o comando setup
```

**Exemplo:**
```bash
./bin/ltx setup
./bin/ltx setup --force  # Reconfigure ambiente existente
```

### `ltx init`
Cria um novo documento LaTeX com templates predefinidos.

```bash
ltx init [flags]

Flags:
  -t, --title string      Título do documento
  -a, --author string     Autor do documento
  -T, --template string   Template a usar (article, book, thesis)
  -f, --force            Sobrescrever arquivos existentes
  -i, --interactive      Modo interativo
  -h, --help             Ajuda para o comando init
```

**Exemplos:**
```bash
./bin/ltx init                                    # Modo interativo
./bin/ltx init --title "Meu Artigo" --author "João Silva"
./bin/ltx init --template thesis --interactive
./bin/ltx init --force                           # Sobrescrever projeto existente
```

### `ltx build`
Compila o documento LaTeX para PDF.

```bash
ltx build [flags]

Flags:
  -c, --clean     Limpar arquivos temporários antes de compilar
  -o, --output    Diretório de saída (padrão: dist/)
  -v, --verbose   Output detalhado da compilação
  -h, --help      Ajuda para o comando build
```

**Exemplos:**
```bash
./bin/ltx build                    # Compilação padrão
./bin/ltx build --clean            # Limpar antes de compilar
./bin/ltx build --verbose          # Ver logs detalhados
```

### `ltx watch`
Modo de desenvolvimento com compilação automática.

```bash
ltx watch [flags]

Flags:
  -d, --debounce int   Delay em ms antes de recompilar (padrão: 500)
  -i, --ignore string  Padrões de arquivos para ignorar
  -v, --verbose        Output detalhado
  -h, --help          Ajuda para o comando watch
```

**Exemplos:**
```bash
./bin/ltx watch                    # Modo watch padrão
./bin/ltx watch --verbose          # Com logs detalhados
./bin/ltx watch --debounce 1000    # Aguardar 1s antes de recompilar
```

### `ltx clean`
Remove arquivos temporários e de build.

```bash
ltx clean [flags]

Flags:
  -a, --all       Limpar tudo, incluindo PDFs
  -d, --dry-run   Mostrar o que seria removido sem remover
  -h, --help      Ajuda para o comando clean
```

**Exemplos:**
```bash
./bin/ltx clean                    # Limpar arquivos temporários
./bin/ltx clean --all              # Limpar tudo incluindo PDFs
./bin/ltx clean --dry-run          # Ver o que seria removido
```

## 🔧 Comandos de Ambiente

### `ltx status`
Exibe status do ambiente Docker e do projeto.

```bash
ltx status [flags]

Flags:
  -v, --verbose   Informações detalhadas
  -h, --help      Ajuda para o comando status
```

### `ltx shell`
Acessa o shell do container Docker.

```bash
ltx shell [flags]

Flags:
  -c, --command string  Comando para executar no container
  -h, --help           Ajuda para o comando shell
```

**Exemplos:**
```bash
./bin/ltx shell                    # Abrir shell interativo
./bin/ltx shell --command "ls -la" # Executar comando específico
```

### `ltx logs`
Mostra logs do container em tempo real.

```bash
ltx logs [flags]

Flags:
  -f, --follow    Seguir logs em tempo real
  -n, --lines int Número de linhas para mostrar (padrão: 50)
  -h, --help      Ajuda para o comando logs
```

### `ltx update`
Atualiza o ambiente Docker (pull de imagens).

```bash
ltx update [flags]

Flags:
  -f, --force     Forçar atualização mesmo se atualizado
  -h, --help      Ajuda para o comando update
```

## ⚙️ Flags Globais

Todos os comandos suportam estas flags globais:

```bash
Global Flags:
  --config string   Arquivo de configuração (padrão: config/latex-cli.conf)
  --debug          Habilitar output de debug
  --no-color       Desabilitar cores no output
  --quiet          Modo silencioso
  --version        Mostrar versão
```

## 📝 Exemplos de Uso

### Fluxo Completo - Novo Projeto
```bash
# 1. Configurar ambiente
./bin/ltx setup

# 2. Criar documento
./bin/ltx init --title "Minha Dissertação" --author "João Silva"

# 3. Desenvolver com auto-compilação
./bin/ltx watch

# Em outro terminal, para compilação final:
# 4. Compilar versão final
./bin/ltx clean
./bin/ltx build
```

### Desenvolvimento Rápido
```bash
# Setup one-liner para desenvolvimento
./bin/ltx setup && ./bin/ltx init --interactive && ./bin/ltx watch
```

### Debugging e Troubleshooting
```bash
# Ver status detalhado
./bin/ltx status --verbose

# Ver logs do container
./bin/ltx logs --follow

# Compilar com logs detalhados
./bin/ltx build --verbose

# Acessar shell para debug manual
./bin/ltx shell
```

### Limpeza e Manutenção
```bash
# Limpeza básica
./bin/ltx clean

# Limpeza completa
./bin/ltx clean --all

# Atualizar ambiente
./bin/ltx update
```

## 🔄 Migração da CLI Legada

Se você estava usando a CLI legada (`latex-cli`), migre para a nova CLI:

| CLI Legada | CLI Nova | Notas |
|------------|----------|-------|
| `./bin/latex-cli setup` | `./bin/ltx setup` | Idêntico |
| `./bin/latex-cli init` | `./bin/ltx init` | Suporte melhorado para flags |
| `./bin/latex-cli build` | `./bin/ltx build` | Performance melhorada |
| `./bin/latex-cli watch` | `./bin/ltx watch` | File watching otimizado |

## 📖 Mais Informações

- **Configuração**: Veja [configuration.md](configuration.md)
- **FAQ**: Veja [faq.md](faq.md)
- **Contribuindo**: Veja [contributing.md](contributing.md)
