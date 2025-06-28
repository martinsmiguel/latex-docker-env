# 🛠️ Guia de Instalação

Este guia fornece instruções detalhadas para instalar e configurar o LaTeX Template em diferentes sistemas operacionais.

## 📋 Pré-requisitos Gerais

- **Git**: Para clonar o repositório
- **Docker**: Versão 20.10 ou superior
- **Docker Compose**: Versão 2.0 ou superior
- **Editor de texto**: VS Code recomendado (configurações incluídas)

## 🪟 Windows

### Método 1: WSL2 (Recomendado)

WSL2 oferece a melhor experiência para desenvolvimento Linux no Windows.

#### 1. Instalar WSL2

```powershell
# Execute como Administrador no PowerShell
wsl --install

# Reinicie o computador quando solicitado
```

#### 2. Instalar Docker Desktop

1. Baixe [Docker Desktop](https://docs.docker.com/desktop/windows/install/)
2. Durante a instalação, certifique-se de habilitar a integração WSL2
3. Após a instalação, vá em Settings > Resources > WSL Integration
4. Habilite a integração com sua distribuição WSL2

#### 3. Configurar o Projeto

```bash
# No terminal WSL2
cd /mnt/c/Users/SeuUsuario/Documents  # ou onde preferir
git clone https://github.com/martinsmiguel/latex-template.git
cd latex-template

# Configurar permissões
chmod +x bin/latex-cli

# Executar setup
./bin/latex-cli setup
```

### Método 2: PowerShell/CMD (Limitado)

```powershell
# No PowerShell
git clone https://github.com/martinsmiguel/latex-template.git
cd latex-template

# Executar setup (note o .\ no Windows)
.\bin\latex-cli setup
```

**Nota**: Algumas funcionalidades podem ser limitadas sem WSL2.

## 🍎 macOS

### 1. Instalar Docker Desktop

1. Baixe [Docker Desktop for Mac](https://docs.docker.com/desktop/mac/install/)
2. Instale e inicie o Docker Desktop
3. Aguarde até que o Docker esteja executando (ícone na barra de menu)

### 2. Configurar o Projeto

```bash
# No Terminal
cd ~/Documents  # ou onde preferir
git clone https://github.com/martinsmiguel/latex-template.git
cd latex-template

# Configurar permissões
chmod +x bin/latex-cli

# Executar setup
./bin/latex-cli setup
```

### 3. (Opcional) Instalar via Homebrew

Se você usa Homebrew, pode instalar o Docker via terminal:

```bash
# Instalar Docker
brew install --cask docker

# Iniciar Docker
open /Applications/Docker.app
```

## 🐧 Linux

### Ubuntu/Debian

#### 1. Instalar Docker

```bash
# Atualizar índice de pacotes
sudo apt update

# Instalar dependências
sudo apt install apt-transport-https ca-certificates curl gnupg lsb-release

# Adicionar chave GPG oficial do Docker
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Adicionar repositório
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Instalar Docker
sudo apt update
sudo apt install docker-ce docker-ce-cli containerd.io docker-compose-plugin

# Adicionar usuário ao grupo docker
sudo usermod -aG docker $USER

# Fazer logout e login novamente, ou executar:
newgrp docker
```

#### 2. Configurar o Projeto

```bash
cd ~/Documents  # ou onde preferir
git clone https://github.com/martinsmiguel/latex-template.git
cd latex-template

# Configurar permissões
chmod +x bin/latex-cli

# Executar setup
./bin/latex-cli setup

# (Opcional) Adicionar ao PATH
sudo ln -sf "$(pwd)/bin/latex-cli" /usr/local/bin/latex-cli
```

### Fedora/CentOS/RHEL

```bash
# Instalar Docker
sudo dnf install docker docker-compose

# Iniciar serviço
sudo systemctl start docker
sudo systemctl enable docker

# Adicionar usuário ao grupo
sudo usermod -aG docker $USER
newgrp docker

# Configurar projeto (mesmo processo do Ubuntu)
```

### Arch Linux

```bash
# Instalar Docker
sudo pacman -S docker docker-compose

# Iniciar serviço
sudo systemctl start docker
sudo systemctl enable docker

# Adicionar usuário ao grupo
sudo usermod -aG docker $USER
newgrp docker

# Configurar projeto (mesmo processo do Ubuntu)
```

## ⚡ Configuração de Autocompletion

### Bash

```bash
# Adicionar ao ~/.bashrc
echo "source $(pwd)/config/completions/latex-cli.bash" >> ~/.bashrc

# Recarregar configuração
source ~/.bashrc
```

### Zsh

```bash
# Adicionar ao ~/.zshrc
echo "fpath=($(pwd)/config/completions \$fpath)" >> ~/.zshrc
echo "autoload -U compinit && compinit" >> ~/.zshrc

# Recarregar configuração
source ~/.zshrc
```

### Fish

```bash
# Criar diretório se não existir
mkdir -p ~/.config/fish/completions

# Copiar arquivo de completions
cp config/completions/latex-cli.fish ~/.config/fish/completions/
```

## 🔧 Verificação da Instalação

Após a instalação, execute estes comandos para verificar se tudo está funcionando:

```bash
# Verificar versão do Docker
docker --version

# Verificar Docker Compose
docker compose version

# Verificar CLI
./bin/latex-cli --version

# Verificar status do ambiente
./bin/latex-cli status

# Teste básico
./bin/latex-cli init --title "Teste" --author "Teste" --non-interactive
./bin/latex-cli build
```

Se todos os comandos executarem sem erro, sua instalação está correta!

## 🆘 Solução de Problemas

### Docker não inicia

**Windows/macOS**: Certifique-se de que o Docker Desktop está executando.

**Linux**: Verifique o status do serviço:
```bash
sudo systemctl status docker
sudo systemctl start docker
```

### Permission denied

**Linux/macOS**:
```bash
chmod +x bin/latex-cli
```

**WSL2**: Certifique-se de estar no sistema de arquivos Linux, não no Windows montado (/mnt/c).

### Container não encontrado

```bash
# Limpar e recriar ambiente
./bin/latex-cli clean
./bin/latex-cli setup
```

### Problemas de rede

```bash
# Verificar conectividade Docker
docker run hello-world

# Se falhar, pode ser necessário configurar proxy corporativo
```

## 📞 Suporte

Se encontrar problemas não cobertos neste guia:

1. Verifique a [documentação completa](README.md)
2. Consulte as [issues existentes](https://github.com/martinsmiguel/latex-template/issues)
3. Abra uma nova issue com:
   - Sistema operacional e versão
   - Versão do Docker
   - Comando que falhou
   - Logs completos do erro

## 🔄 Atualizações

Para atualizar o template:

```bash
# Fazer backup de seus arquivos em src/
cp -r src/ src_backup/

# Atualizar repositório
git pull origin main

# Reconfigurar se necessário
./bin/latex-cli setup

# Restaurar seus arquivos
cp -r src_backup/* src/
rm -rf src_backup/
```
