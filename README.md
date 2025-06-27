# 📄 LaTeX Template

Um template completo e dockerizado para escrita de documentos LaTeX com VS Code, incluindo automação para compilação e ambiente de desenvolvimento consistente.

## ✨ Características

- 🐳 **Ambiente Dockerizado**: Desenvolvimento isolado com todas as dependências LaTeX incluídas
- 🔄 **Compilação Automática**: Auto-build no VS Code com `latexmk`
- 📁 **Estrutura Organizada**: Templates pré-configurados para diferentes tipos de documento
- 🛠️ **Scripts de Automação**: Inicialização e compilação simplificadas
- 🔧 **Configuração VS Code**: Settings otimizados para desenvolvimento LaTeX
- 📦 **Gerenciamento de Pacotes**: Instalação automática de pacotes LaTeX ausentes

## 🚀 Início Rápido

### 1. Configuração Inicial

```bash
# Clone o repositório
git clone https://github.com/martinsmiguel/latex-template.git
cd latex-template

# Dê permissão de execução ao script e execute
chmod +x start.sh
./start.sh

# Inicie o ambiente Docker
docker-compose up -d
```

### 2. Abra no VS Code com Dev Container

```bash
# Abra o projeto no VS Code
code .

# O VS Code irá detectar o Dev Container e oferecerá para reabrir no container
# Ou use Ctrl+Shift+P -> "Dev Containers: Reopen in Container"
```

### 3. Inicialize seu Documento

```bash
# Execute o script de inicialização para criar um novo documento
./scripts/init_project.sh
```

O script irá solicitar:
- **Título do documento**
- **Nome do autor**

E criará automaticamente:
- `tex/main.tex` personalizado
- Estrutura de capítulos em `tex/chapters/` (introdução, metodologia, resultados)
- `tex/preamble.tex` se não existir

## 📂 Estrutura do Projeto

```
latex-template/
├── .devcontainer/          # Configuração do ambiente Docker
│   ├── devcontainer.json   # Configurações do VS Code Dev Container
│   └── Dockerfile          # Imagem personalizada com TeX Live
├── .vscode/                # Configurações do VS Code
│   ├── settings.json       # Settings do LaTeX Workshop
│   └── tasks.json          # Tasks de compilação e limpeza
├── scripts/                # Scripts de automação
│   ├── compile.sh          # Script de compilação manual
│   ├── clean.sh            # Script de limpeza de arquivos auxiliares
│   ├── init_project.sh     # Inicialização de novos documentos
│   └── latexmk-docker.sh   # Wrapper para latexmk no container
├── templates/              # Templates base
│   ├── main.tex.tpl        # Template principal do documento
│   ├── preamble.tex.tpl    # Template de configurações LaTeX
│   ├── settings.json.tpl   # Template de configurações do VS Code
│   └── references.bib.tpl  # Template para bibliografia
├── tex/                    # Seus arquivos LaTeX (criado automaticamente)
│   ├── main.tex            # Documento principal
│   ├── preamble.tex        # Configurações e pacotes LaTeX
│   ├── chapters/           # Capítulos do documento
│   └── references.bib      # Bibliografia
├── out/                    # Arquivos de saída (PDF, logs)
├── docker-compose.yml      # Configuração do Docker Compose
└── start.sh               # Script de configuração inicial
```

## 🛠️ Scripts Disponíveis

### `./start.sh`
Configura o ambiente inicial:
- Cria diretórios necessários
- Copia templates se não existirem
- Configura permissões dos scripts

### `./scripts/init_project.sh`
Inicializa um novo documento:
- Solicita título e autor
- Gera `main.tex` personalizado
- Cria estrutura de capítulos
- Copia template de preamble se necessário

### `./scripts/latexmk-docker.sh`
Wrapper para executar latexmk no container:
- Permite execução direta do latexmk via Docker
- Facilita integração com editores externos
- Passa todos os argumentos para o container

### `./scripts/compile.sh`
Compila o documento manualmente:
- Verifica dependências necessárias
- Usa `latexmk` para compilação otimizada
- Gera PDF em `out/main.pdf`
- Fornece feedback claro sobre erros

### `./scripts/clean.sh`
Limpa arquivos auxiliares:
- Remove logs e arquivos temporários
- Mantém apenas o PDF final
- Útil para debugging de problemas

## 📝 Como Usar

### Escrevendo seu Documento

1. **Documento Principal**: Edite `tex/main.tex`
2. **Capítulos**: Adicione conteúdo em `tex/chapters/`
3. **Bibliografia**: Adicione referências em `tex/references.bib`

### Compilação

A compilação acontece automaticamente quando você salva arquivos `.tex` no VS Code. Alternativamente:

```bash
# Compilação manual
./scripts/compile.sh
```

### Tarefas do VS Code

O projeto inclui tarefas pré-configuradas acessíveis via `Ctrl+Shift+P` → "Tasks: Run Task":

- **Compile LaTeX**: Executa `./scripts/compile.sh`
- **Clean LaTeX**: Executa `./scripts/clean.sh`

### Visualização

O PDF gerado ficará em `out/main.pdf` e pode ser visualizado:
- **No VS Code**: Aba automática do LaTeX Workshop
- **Externamente**: Qualquer visualizador de PDF

## 🔧 Configurações

### VS Code Settings

O template inclui configurações otimizadas em `.vscode/settings.json`:

```json
{
  "latex-workshop.latex.autoBuild.run": "onSave",
  "latex-workshop.latex.outDir": "./out",
  "latex-workshop.view.pdf.viewer": "tab",
  "files.autoSave": "afterDelay",
  "latex-workshop.latex.tools": [
    {
      "name": "latexmk-docker",
      "command": "docker",
      "args": [
        "exec", "latex-env", "latexmk",
        "-pdf", "-f", "-interaction=nonstopmode",
        "-synctex=1", "-outdir=./out", "tex/main.tex"
      ]
    }
  ],
  "latex-workshop.latex.recipes": [
    {
      "name": "latexmk (Docker)",
      "tools": ["latexmk-docker"]
    }
  ],
  "latex-workshop.latex.recipe.default": "latexmk (Docker)"
}
```

### Pacotes LaTeX Incluídos

O ambiente Docker inclui:
- **TeX Live completo**
- **latexmk** para compilação
- **biber** para bibliografia
- **python3-pygments** para syntax highlighting
- Pacotes adicionais: `enumitem`, `fancyhdr`, `hyperref`, `xcolor`

## 🐳 Ambiente Docker

### Características do Container

- **Base**: `texlive/texlive:latest`
- **Usuário**: `latexuser` (não-root para segurança, UID/GID 1001)
- **Pacotes**: TeX Live completo + dependências essenciais
- **Health Check**: Monitora disponibilidade do `pdflatex`
- **Volume persistente**: Cache do TeX Live para melhor performance

### Comandos Docker Úteis

```bash
# Iniciar o ambiente
docker-compose up -d

# Verificar status do container
docker ps

# Acessar terminal do container
docker exec -it latex-env bash

# Ver logs do container
docker logs latex-env

# Parar o ambiente
docker-compose down
```

## 📋 Dependências

### Locais (Host)
- **Docker** e **Docker Compose**
- **VS Code** com extensão **Dev Containers**

### Container (Automáticas)
- **TeX Live** (distribuição completa)
- **LaTeX Workshop** (extensão VS Code)
- **latexmk** e **biber** para compilação e bibliografia
- **python3-pygments** para syntax highlighting (pacote minted)
- **Pacotes adicionais**: `enumitem`, `fancyhdr`, `hyperref`, `xcolor`

## 🚨 Solução de Problemas

### Erro de Permissão nos Scripts

Se você receber "permission denied" ao executar scripts:

```bash
# Dê permissão de execução a todos os scripts
chmod +x start.sh
chmod +x scripts/*.sh

# Ou use o comando individual conforme necessário
chmod +x ./scripts/init_project.sh
chmod +x ./scripts/compile.sh
chmod +x ./scripts/clean.sh
chmod +x ./scripts/latexmk-docker.sh
```

### Pacotes LaTeX Ausentes

O script de compilação tenta instalar automaticamente pacotes ausentes. Se falhar:

```bash
# Entre no container
docker exec -it latex-env bash

# Instale manualmente
tlmgr install nome-do-pacote
```

### Erro de Compilação

1. Verifique os logs em `out/logs/`
2. Use compilação manual: `./scripts/compile.sh`
3. Verifique sintaxe LaTeX no VS Code

### Container não Inicia

```bash
# Reconstruir o container
docker-compose down
docker-compose build --no-cache
docker-compose up -d
```

## 🤝 Contribuindo

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 📞 Suporte

- 🐛 **Issues**: [GitHub Issues](https://github.com/martinsmiguel/latex-template/issues)
- 📧 **Email**: miguelrjmartins.dev@gmail.com
- 📚 **Documentação**: Consulte este README para informações completas

---

**Feito com ❤️ para simplificar a escrita acadêmica em LaTeX**
