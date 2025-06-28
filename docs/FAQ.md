# ❓ Perguntas Frequentes (FAQ)

## 🚀 Primeiros Passos

### ❔ Preciso instalar LaTeX no meu sistema?

**Não!** Este template usa Docker, então todas as dependências LaTeX ficam isoladas no container. Você só precisa do Docker instalado.

### ❔ Funciona no Windows?

**Sim!** Recomendamos usar WSL2 para a melhor experiência. Também funciona diretamente no PowerShell, mas com algumas limitações.

### ❔ Posso usar sem Docker?

Tecnicamente sim, mas não é recomendado. Você precisaria instalar manualmente:
- TeX Live ou MiKTeX
- latexmk
- Todas as dependências de pacotes

O Docker garante um ambiente consistente e evita conflitos.

## 🛠️ Configuração e Uso

### ❔ Como personalizar o template?

1. **Templates**: Edite arquivos em `config/templates/`
2. **Configurações**: Modifique `config/latex-cli.conf`
3. **VS Code**: Ajuste `config/vscode/settings.json`

### ❔ Posso usar meu editor favorito?

Sim! O template funciona com qualquer editor. As configurações do VS Code são opcionais, mas recomendadas para melhor experiência.

### ❔ Como adicionar novos pacotes LaTeX?

Edite o arquivo `src/preamble.tex` e adicione:
```latex
\usepackage{nomeDoPacote}
```

O container já inclui a maioria dos pacotes populares.

### ❔ Como organizar capítulos?

Organize seus capítulos na pasta `src/chapters/`:
```
src/chapters/
├── introduction.tex
├── methodology.tex
├── results.tex
└── conclusion.tex
```

Inclua-os no `main.tex` com `\input{chapters/nomeDoCaptulo}`

## 🐳 Docker e Containers

### ❔ O container consome muitos recursos?

O container é otimizado para desenvolvimento LaTeX:
- **RAM**: ~500MB em uso normal
- **Disco**: ~2GB (TeX Live completo)
- **CPU**: Apenas durante compilação

### ❔ Como limpar espaço em disco?

```bash
# Limpar arquivos temporários do projeto
./bin/latex-cli clean

# Limpar cache Docker (remove todos os containers/imagens não utilizados)
docker system prune -a
```

### ❔ O container persiste dados?

Sim! Seus arquivos ficam no sistema host. O container apenas fornece o ambiente de compilação.

## 📝 Compilação e Output

### ❔ Onde fica o PDF gerado?

O PDF compilado fica em `dist/main.pdf` (configurável).

### ❔ Como compilar apenas partes do documento?

Use comentários no `main.tex`:
```latex
\input{chapters/introduction}
% \input{chapters/methodology}  % Comentado - não compila
\input{chapters/conclusion}
```

### ❔ Como alterar configurações de compilação?

Edite `config/latex-cli.conf`:
```bash
# Comando de compilação personalizado
LATEX_COMPILE_CMD="pdflatex -interaction=nonstopmode"

# Diretório de output
OUTPUT_DIR="build"
```

## 🔧 Troubleshooting

### ❔ "Permission denied" no Linux/macOS

```bash
chmod +x bin/latex-cli
```

### ❔ Container não inicia

1. Verificar se Docker está rodando:
   ```bash
   docker --version
   ```

2. Recriar ambiente:
   ```bash
   ./bin/latex-cli clean
   ./bin/latex-cli setup
   ```

### ❔ Compilação falha

1. Verificar logs:
   ```bash
   ./bin/latex-cli logs
   ```

2. Verificar sintaxe LaTeX no arquivo que foi modificado

3. Testar compilação manual:
   ```bash
   ./bin/latex-cli shell
   cd /workspace/src
   pdflatex main.tex
   ```

### ❔ Autocompletion não funciona

Verifique se adicionou corretamente ao seu shell:

**Bash**:
```bash
echo "source $(pwd)/config/completions/latex-cli.bash" >> ~/.bashrc
source ~/.bashrc
```

**Zsh**:
```bash
echo "fpath=($(pwd)/config/completions \$fpath)" >> ~/.zshrc
source ~/.zshrc
```

## 🔄 Atualização e Manutenção

### ❔ Como atualizar o template?

```bash
# Backup seus arquivos
cp -r src/ src_backup/

# Atualizar
git pull origin main
./bin/latex-cli setup

# Restaurar arquivos
cp -r src_backup/* src/
```

### ❔ Como migrar projeto existente?

1. Copie seus `.tex` para `src/`
2. Copie bibliografia para `src/references.bib`
3. Ajuste `main.tex` para incluir seus arquivos
4. Execute `./bin/latex-cli build`

### ❔ Como fazer backup do projeto?

```bash
# Backup completo (excluindo Docker)
tar -czf meu-projeto-backup.tar.gz \
  --exclude='.git' \
  --exclude='dist' \
  --exclude='node_modules' \
  .

# Backup apenas fontes
tar -czf fontes-backup.tar.gz src/ config/templates/
```

## 📊 Performance e Otimização

### ❔ Compilação está lenta

1. **Use watch mode**: `./bin/latex-cli watch`
2. **Compile incrementalmente**: latexmk faz cache inteligente
3. **Reduza pacotes**: comente pacotes não utilizados
4. **Use SSD**: Docker funciona melhor em SSD

### ❔ Como acelerar desenvolvimento?

1. **VS Code + LaTeX Workshop**: Configuração incluída
2. **Watch mode**: Auto-compilação em mudanças
3. **Snippets**: Use autocompletion do VS Code
4. **Split view**: Visualize PDF enquanto edita

## 🤝 Colaboração

### ❔ Como trabalhar em equipe?

1. **Git**: Versione apenas `src/` e configurações
2. **Branches**: Um branch por pessoa/feature
3. **.gitignore**: Exclua `dist/` e arquivos temporários
4. **Docker**: Garante ambiente igual para todos

### ❔ Como resolver conflitos?

```bash
# Compilar após merge
git merge origin/main
./bin/latex-cli build

# Se houver problemas
./bin/latex-cli clean
./bin/latex-cli build
```

## 📱 Integração e CI/CD

### ❔ Como automatizar compilação?

Exemplo para GitHub Actions:

```yaml
# .github/workflows/latex.yml
name: Build LaTeX
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Build PDF
        run: |
          chmod +x bin/latex-cli
          ./bin/latex-cli setup
          ./bin/latex-cli build
      - name: Upload PDF
        uses: actions/upload-artifact@v2
        with:
          name: document.pdf
          path: dist/main.pdf
```

### ❔ Como integrar com Overleaf?

1. Exporte seus arquivos `src/` para um ZIP
2. Importe no Overleaf
3. Para sincronizar: use Git bridge do Overleaf (plano pago)

## 🔍 Debugging

### ❔ Como debugar problemas de compilação?

1. **Modo verboso**:
   ```bash
   ./bin/latex-cli build --verbose
   ```

2. **Shell interativo**:
   ```bash
   ./bin/latex-cli shell
   pdflatex -interaction=nonstopmode main.tex
   ```

3. **Logs detalhados**:
   ```bash
   ./bin/latex-cli logs --follow
   ```

### ❔ Como verificar dependências?

```bash
# Entrar no container
./bin/latex-cli shell

# Verificar pacotes instalados
tlmgr list --installed

# Verificar versão TeX
pdflatex --version
```

## 🆘 Suporte

### ❔ Onde obter ajuda?

1. **Documentação**: [docs/README.md](README.md)
2. **Issues**: [GitHub Issues](https://github.com/martinsmiguel/latex-template/issues)
3. **Discussões**: [GitHub Discussions](https://github.com/martinsmiguel/latex-template/discussions)
4. **Wiki**: [GitHub Wiki](https://github.com/martinsmiguel/latex-template/wiki)

### ❔ Como reportar bugs?

Inclua sempre:
- Sistema operacional
- Versão do Docker
- Comando que falhou
- Logs completos
- Arquivo de exemplo (se possível)

### ❔ Como sugerir melhorias?

1. Abra uma [Discussion](https://github.com/martinsmiguel/latex-template/discussions)
2. Descreva o caso de uso
3. Proponha a implementação
4. Considere fazer um Pull Request!

---

**💡 Não encontrou sua pergunta? [Abra uma issue](https://github.com/martinsmiguel/latex-template/issues/new) e ajude a expandir este FAQ!**
