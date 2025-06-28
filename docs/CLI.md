# LaTeX Template CLI

Uma interface de linha de comando robusta e completa para gerenciar projetos LaTeX dockerizados, desenvolvida seguindo as melhores práticas de Bash scripting.

## Instalação Rápida

```bash
# No diretório do projeto
chmod +x latex-cli

# (Opcional) Adicionar ao PATH
sudo ln -sf "$(pwd)/latex-cli" /usr/local/bin/latex-cli

# (Opcional) Habilitar autocompletion
# Para Bash
echo "source $(pwd)/completions/latex-cli.bash" >> ~/.bashrc

# Para Zsh
echo "fpath=($(pwd)/completions \$fpath)" >> ~/.zshrc
echo "autoload -U compinit && compinit" >> ~/.zshrc
```

## Comandos Disponíveis

### Comandos Principais

#### `setup`
Configura o ambiente inicial do projeto, criando diretórios necessários e copiando templates.

```bash
./latex-cli setup
```

#### `init`
Inicializa um novo documento LaTeX com base nos templates disponíveis.

```bash
# Modo interativo
./latex-cli init

# Modo não-interativo
./latex-cli init --title "Minha Dissertação" --author "João Silva" --non-interactive

# Usando template específico
./latex-cli init --template article --title "Artigo Científico"

# Forçar sobrescrita
./latex-cli init --force
```

**Opções:**
- `--title <título>`: Título do documento
- `--author <autor>`: Nome do autor
- `--template <template>`: Template a usar (padrão: main)
- `--non-interactive`: Não solicita entrada do usuário
- `--force`: Sobrescreve arquivos existentes

#### `build`
Compila o documento LaTeX principal usando o ambiente Docker.

```bash
# Compilação básica
./latex-cli build

# Limpeza antes da compilação
./latex-cli build --clean

# Usando engine específico
./latex-cli build --engine xelatex --verbose

# Saída personalizada
./latex-cli build --output-dir dist
```

**Opções:**
- `--output-dir <dir>`: Diretório de saída (padrão: out)
- `--clean`: Limpa arquivos temporários antes de compilar
- `--engine <engine>`: Engine LaTeX (pdflatex, xelatex, lualatex)
- `--no-bib`: Pula processamento de bibliografia
- `--verbose`: Mostra saída detalhada da compilação

#### `watch`
Compila automaticamente o documento quando arquivos LaTeX são modificados.

```bash
./latex-cli watch
```

Pressione `Ctrl+C` para parar o modo de observação.

#### `clean`
Remove arquivos temporários e de compilação.

```bash
# Limpa apenas arquivos temporários
./latex-cli clean

# Limpa tudo (incluindo PDFs)
./latex-cli clean --all

# Limpa apenas PDFs
./latex-cli clean --pdf

# Limpa cache do LaTeX
./latex-cli clean --cache
```

**Opções:**
- `--all`: Remove tudo (temp, PDF, cache)
- `--temp-only`: Remove apenas arquivos temporários (padrão)
- `--pdf`: Inclui arquivos PDF na limpeza
- `--cache`: Limpa cache do LaTeX
- `--quiet`: Suprime saída

### Comandos de Ambiente

#### `dev`
Inicia o ambiente de desenvolvimento Docker.

```bash
# Inicia o ambiente
./latex-cli dev

# Reconstrói a imagem antes de iniciar
./latex-cli dev --rebuild
```

#### `stop`
Para o ambiente Docker.

```bash
./latex-cli stop
```

#### `restart`
Reinicia o ambiente Docker.

```bash
./latex-cli restart
```

#### `status`
Mostra o status completo do ambiente e projeto.

```bash
./latex-cli status
```

### Comandos de Gerenciamento

#### `logs`
Mostra logs do container Docker.

```bash
# Logs básicos
./latex-cli logs

# Seguir logs em tempo real
./latex-cli logs --follow

# Número específico de linhas
./latex-cli logs --lines 100

# Logs de serviço específico
./latex-cli logs latex-env
```

#### `shell`
Abre um shell interativo no container LaTeX.

```bash
./latex-cli shell
```

#### `exec`
Executa um comando específico no container.

```bash
# Executar comando LaTeX
./latex-cli exec pdflatex --version

# Listar arquivos
./latex-cli exec ls -la tex/

# Executar script personalizado
./latex-cli exec bash -c "cd tex && find . -name '*.tex'"
```

#### `update`
Atualiza a imagem Docker para a versão mais recente.

```bash
./latex-cli update
```

### Comandos de Informação

#### `templates`
Lista todos os templates disponíveis.

```bash
./latex-cli templates
```

#### `packages`
Lista pacotes LaTeX instalados no ambiente.

```bash
./latex-cli packages
```

#### `version`
Mostra informações de versão da CLI e do ambiente.

```bash
./latex-cli version
```

## Opções Globais

As seguintes opções podem ser usadas com qualquer comando:

- `-h, --help`: Mostra ajuda
- `-v, --verbose`: Ativa modo detalhado
- `-q, --quiet`: Suprime saída desnecessária
- `--no-docker`: Executa comandos sem Docker (quando possível)
- `--version`: Mostra versão

## Estrutura de Arquivos

A CLI trabalha com a seguinte estrutura de projeto:

```
projeto/
├── latex-cli                 # CLI principal
├── .latex-cli.conf          # Configuração (opcional)
├── docker-compose.yml       # Configuração Docker
├── templates/               # Templates LaTeX
│   ├── main.tex.tpl
│   ├── preamble.tex.tpl
│   └── ...
├── tex/                     # Arquivos LaTeX
│   ├── main.tex
│   ├── preamble.tex
│   ├── references.bib
│   └── chapters/
├── out/                     # Saída da compilação
│   ├── main.pdf
│   └── logs/
├── completions/             # Autocompletion
│   ├── latex-cli.bash
│   └── _latex-cli
└── scripts/                 # Scripts auxiliares (legacy)
```

## Configuração

### Arquivo de Configuração

Crie um arquivo `.latex-cli.conf` no diretório do projeto para personalizar o comportamento:

```bash
# Configurações do Docker
LATEX_CONTAINER_NAME="meu-latex-env"
LATEX_DEFAULT_ENGINE="xelatex"
LATEX_AUTO_OPEN_PDF="true"

# Hooks personalizados
LATEX_POST_BUILD_HOOK="echo 'Compilação concluída!'"
```

### Autocompletion

Para habilitar autocompletion no seu shell:

**Bash:**
```bash
# Adicione ao ~/.bashrc
source /caminho/para/completions/latex-cli.bash
```

**Zsh:**
```bash
# Adicione ao ~/.zshrc
fpath=(/caminho/para/completions $fpath)
autoload -U compinit && compinit
```

## Recursos de Segurança

A CLI implementa várias práticas de segurança:

- ✓ Validação rigorosa de entrada
- ✓ Tratamento de erros com `set -euo pipefail`
- ✓ Variáveis sempre entre aspas duplas
- ✓ Verificação de existência de comandos
- ✓ Códigos de saída apropriados
- ✓ Sanitização de caminhos
- ✓ Prevenção de injeção de comandos

## Solução de Problemas

### Container não inicia
```bash
# Verifique se Docker está rodando
docker info

# Reconstrua a imagem
./latex-cli dev --rebuild

# Verifique logs
./latex-cli logs
```

### Compilação falha
```bash
# Modo verbose para debug
./latex-cli build --verbose

# Limpe arquivos temporários
./latex-cli clean --all

# Verifique logs de compilação
cat out/main.log
```

### Permissões
```bash
# Garanta que a CLI é executável
chmod +x latex-cli

# Verifique permissões dos scripts
find scripts/ -name "*.sh" -exec chmod +x {} \;
```

## Exemplos de Workflow

### Workflow Básico
```bash
# 1. Configurar ambiente
./latex-cli setup

# 2. Inicializar documento
./latex-cli init --title "Minha Tese" --author "Meu Nome"

# 3. Iniciar ambiente
./latex-cli dev

# 4. Compilar
./latex-cli build

# 5. Modo de desenvolvimento
./latex-cli watch
```

### Workflow Avançado
```bash
# Desenvolvimento com limpeza automática
./latex-cli build --clean --engine xelatex --verbose

# Pipeline de CI/CD
./latex-cli setup && \
./latex-cli init --non-interactive --title "Relatório" && \
./latex-cli build --output-dir release && \
./latex-cli clean --temp-only
```

## Contribuindo

Para contribuir com melhorias na CLI:

1. Fork o repositório
2. Crie uma branch para sua feature
3. Siga as práticas de Bash scripting implementadas
4. Teste todos os comandos
5. Submeta um PR

## Licença

Esta CLI segue a mesma licença do projeto LaTeX Template.
