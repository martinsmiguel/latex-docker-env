# ❓ Perguntas Frequentes (FAQ)

Soluções para problemas comuns e dúvidas sobre o LaTeX Docker Environment.

## 🚀 Instalação e Configuração

### O comando `./bin/ltx` retorna "Permission denied"

**Solução para Linux/macOS:**
```bash
chmod +x bin/ltx
```

**No Windows:** Este erro não deve ocorrer. Use `.\bin\ltx.exe` no PowerShell.

### Docker não encontrado

**Verificar se o Docker está instalado e rodando:**
```bash
docker --version
docker ps
```

**Se não estiver instalado:**
- Windows/macOS: Instale [Docker Desktop](https://docs.docker.com/get-docker/)
- Linux: `sudo apt install docker.io docker-compose-plugin` (Ubuntu/Debian)

**Se não estiver rodando:**
```bash
# Linux
sudo systemctl start docker

# Windows/macOS
# Abrir Docker Desktop
```

### "docker compose" vs "docker-compose"

O projeto usa `docker compose` (comando integrado). Se você tem apenas `docker-compose`:

```bash
# Verificar versão
docker compose version

# Se não funcionar, instalar Docker mais recente ou usar workaround
alias docker-compose='docker compose'
```

## 🐳 Problemas com Docker

### Container não inicia

**Ver logs para diagnóstico:**
```bash
./bin/ltx logs
./bin/ltx status --verbose
```

**Soluções comuns:**
```bash
# Recriar ambiente
./bin/ltx clean
./bin/ltx setup

# Atualizar imagens
./bin/ltx update

# Em último caso, limpeza completa do Docker
docker system prune -a
```

### Compilação LaTeX falha

**Ver logs detalhados:**
```bash
./bin/ltx build --verbose
```

**Problemas comuns:**
- **Pacotes LaTeX em falta**: Adicione ao `preamble.tex`
- **Sintaxe LaTeX inválida**: Verifique arquivos `.tex`
- **Referências quebradas**: Verifique `references.bib`

**Debug manual:**
```bash
# Acessar container para debug
./bin/ltx shell

# Dentro do container
cd /workspace
latexmk -pdf main.tex
```

### Performance lenta

**Para modo watch:**
```bash
# Aumentar debounce time
./bin/ltx watch --debounce 1000
```

**Para builds:**
```bash
# Usar build incremental (padrão)
./bin/ltx build

# Para limpar cache se necessário
./bin/ltx clean && ./bin/ltx build
```

## 🖥️ Problemas Específicos por Sistema

### Windows WSL2

**WSL2 não está instalado:**
```powershell
# Como Administrador
wsl --install
wsl --set-default-version 2
```

**Docker não funciona no WSL2:**
1. Docker Desktop > Settings > Resources > WSL Integration
2. Habilitar sua distribuição WSL2
3. Restart Docker Desktop

**Arquivos no Windows não aparecem no WSL2:**
```bash
# Acessar arquivos do Windows no WSL2
cd /mnt/c/Users/SeuUsuario/Documents/latex-docker-env
```

### macOS

**"Cannot connect to the Docker daemon":**
```bash
# Verificar se Docker Desktop está rodando
open -a Docker

# Aguardar Docker inicializar completamente
```

**Permissões em volumes Docker:**
```bash
# Se arquivos ficam com owner errado
sudo chown -R $(whoami) src/ dist/
```

### Linux

**"Permission denied" para Docker:**
```bash
# Adicionar usuário ao grupo docker
sudo usermod -aG docker $USER

# Logout e login novamente, ou:
newgrp docker
```

**Docker compose não encontrado:**
```bash
# Ubuntu/Debian
sudo apt install docker-compose-plugin

# Fedora/RHEL
sudo dnf install docker-compose

# Arch
sudo pacman -S docker-compose
```

## 📝 Problemas com LaTeX

### Pacotes LaTeX não encontrados

**Adicionar pacotes no `src/preamble.tex`:**
```latex
\usepackage{pacote-necessario}
```

**Para pacotes não incluídos na imagem:**
```bash
# Acessar container
./bin/ltx shell

# Instalar pacote manualmente
tlmgr install nome-do-pacote
```

### Bibliografia não aparece

**Verificar arquivos:**
1. `src/references.bib` existe e tem entradas válidas
2. `src/main.tex` tem `\bibliography{references}`
3. Citações usam `\cite{chave}`

**Forçar recompilação:**
```bash
./bin/ltx clean
./bin/ltx build
```

### Imagens não aparecem

**Verificar paths:**
```latex
% Use paths relativos a src/
\includegraphics{images/figura.png}
```

**Verificar formatos suportados:**
- PDF, PNG, JPG são recomendados
- EPS precisa de conversão em pdflatex

## 🔧 Configuração e Personalização

### Mudar porta do container

**Editar `config/docker/docker-compose.yml`:**
```yaml
ports:
  - "8080:8080"  # mudar primeira porta
```

### Usar template personalizado

**Criar template em `config/templates/`:**
```bash
cp config/templates/main.tex.tpl config/templates/meu-template.tex.tpl
# Editar conforme necessário
```

**Usar o template:**
```bash
./bin/ltx init --template meu-template
```

### Configuração global da CLI

**Editar `config/latex-cli.conf`:**
```ini
[default]
author = Seu Nome Padrão
template = book
output_dir = dist
```

## 🆘 Obtendo Mais Ajuda

### Logs e Debug

**Ver todas as informações de debug:**
```bash
./bin/ltx --debug status --verbose
./bin/ltx --debug build --verbose
```

**Exportar logs para análise:**
```bash
./bin/ltx logs > debug.log
./bin/ltx status --verbose > status.log
```

### Reportar Problemas

**Ao abrir uma issue, inclua:**

1. **Sistema operacional** e versão
2. **Versão do Docker**: `docker --version`
3. **Versão da CLI**: `./bin/ltx --version`
4. **Comando que falhou** e output completo
5. **Logs do container**: `./bin/ltx logs`

**Links úteis:**
- [Issues no GitHub](https://github.com/martinsmiguel/latex-docker-env/issues)
- [Discussões](https://github.com/martinsmiguel/latex-docker-env/discussions)
- [Documentação](.)

### Comandos de Diagnóstico

**Script completo de diagnóstico:**
```bash
#!/bin/bash
echo "=== Sistema ==="
uname -a
echo "=== Docker ==="
docker --version
docker compose version
docker ps
echo "=== LaTeX CLI ==="
./bin/ltx --version
./bin/ltx status --verbose
echo "=== Logs ==="
./bin/ltx logs -n 20
```

**Salvar como `debug.sh` e executar:**
```bash
chmod +x debug.sh
./debug.sh > diagnostico.txt
```
