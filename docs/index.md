# 📚 Documentação do LaTeX Docker Environment

Bem-vindo à documentação oficial do LaTeX Docker Environment - uma solução moderna e completa para desenvolvimento LaTeX com Docker.

## 🚀 Início Rápido

Novo no projeto? Comece aqui:

1. **[📦 Instalação](installation.md)** - Configure o ambiente em qualquer sistema
2. **[🛠️ CLI Reference](cli-reference.md)** - Aprenda a usar a CLI `ltx`
3. **[❓ FAQ](faq.md)** - Soluções para problemas comuns

## 📖 Guias do Usuário

### Básico
- **[Instalação](installation.md)** - Instruções detalhadas por sistema operacional
- **[CLI Reference](cli-reference.md)** - Documentação completa dos comandos
- **[FAQ](faq.md)** - Perguntas frequentes e troubleshooting

### Avançado
- **[Migração](migration.md)** - Migração de versões anteriores
- **[Configuração](cli/configuration.md)** - Personalização avançada

## 🔧 Para Desenvolvedores

- **[Contribuindo](contributing.md)** - Como contribuir para o projeto
- **[Arquitetura CLI](cli/architecture.md)** - Design interno da CLI Go
- **[Desenvolvimento](cli/development-guide.md)** - Setup de desenvolvimento

## 🏗️ Estrutura da Documentação

```
docs/
├── index.md              # Este arquivo - ponto de entrada
├── installation.md       # Guia de instalação
├── cli-reference.md      # Referência completa da CLI
├── faq.md               # Perguntas frequentes
├── migration.md         # Guia de migração
├── contributing.md      # Como contribuir
└── cli/                 # Documentação técnica da CLI
    ├── architecture.md
    ├── configuration.md
    ├── development-guide.md
    └── ...
```

## 📋 Referência Rápida

### Comandos Essenciais
```bash
./bin/ltx setup          # Configurar ambiente
./bin/ltx init           # Criar documento
./bin/ltx build          # Compilar PDF
./bin/ltx watch          # Modo desenvolvimento
./bin/ltx clean          # Limpar arquivos temporários
```

### Links Úteis
- **[GitHub Repository](https://github.com/martinsmiguel/latex-docker-env)**
- **[Issues](https://github.com/martinsmiguel/latex-docker-env/issues)** - Reportar bugs
- **[Discussions](https://github.com/martinsmiguel/latex-docker-env/discussions)** - Dúvidas e sugestões

## 🎯 O que é o LaTeX Docker Environment?

Uma solução completa que combina:

- 🐳 **Ambiente Docker** isolado com todas as dependências LaTeX
- ⚡ **CLI moderna** em Go com suporte nativo ao Windows
- 🔄 **Compilação automática** com detecção de mudanças
- 🛠️ **Configuração VS Code** otimizada para LaTeX
- 📦 **Setup automatizado** em um comando
- 🌐 **Multiplataforma** - Windows, macOS, Linux

## 🤔 Por que usar?

### ✅ Vantagens
- **Sem instalação complexa**: Apenas Docker necessário
- **Ambiente consistente**: Funciona igual em qualquer sistema
- **Setup rápido**: Projeto rodando em minutos
- **CLI intuitiva**: Comandos simples e claros
- **Multiplataforma**: Windows nativo sem WSL2 obrigatório

### 🆚 Comparação com outras soluções

| Aspecto | Este Projeto | LaTeX Local | Overleaf |
|---------|--------------|-------------|----------|
| **Setup** | ⚡ 1 comando | 😰 Complexo | ✅ Zero |
| **Offline** | ✅ Completo | ✅ Completo | ❌ Internet obrigatória |
| **Versionamento** | ✅ Git nativo | ✅ Git nativo | 💰 Premium |
| **Colaboração** | ✅ Git/GitHub | ✅ Git/GitHub | ✅ Built-in |
| **Dependências** | 🐳 Docker apenas | 📦 Muitas | ❌ Nenhuma |
| **Performance** | ⚡ Local | ⚡ Local | 🌐 Depende da internet |

## 📈 Versões

### v2.0 (Atual)
- ✨ CLI moderna em Go
- 🪟 Suporte Windows nativo
- 🚀 Performance melhorada
- 📦 Setup simplificado

### v1.x (Legada)
- 🐚 CLI em Bash
- 🐧 Foco em Linux/macOS
- 📄 Funcional mas limitada

> **💡 Recomendação**: Use sempre a CLI moderna `ltx` ao invés da legada `latex-cli`

## 🆘 Precisa de Ajuda?

1. **Consulte o [FAQ](faq.md)** - Soluções para 90% dos problemas
2. **Leia a [documentação](cli-reference.md)** - Referência completa
3. **Abra uma [issue](https://github.com/martinsmiguel/latex-docker-env/issues)** - Para bugs
4. **Inicie uma [discussion](https://github.com/martinsmiguel/latex-docker-env/discussions)** - Para dúvidas

---

**📝 Última atualização**: Janeiro 2025
**🔖 Versão**: 2.0
**📄 Licença**: [MIT](../LICENSE)

# 3. Testes
make test

# 4. Build
make build-all
```

## Comparação: CLI Legada vs Moderna

| Aspecto | `latex-cli` (Bash) | `ltx` (Go) |
|---------|-------------------|------------|
| **Compatibilidade** | Linux, macOS, WSL2 | Windows, macOS, Linux |
| **Performance** | ~200-500ms startup | ~10-50ms startup |
| **Distribuição** | Scripts + chmod | Binário único |
| **File Watching** | Polling básico | fsnotify nativo |
| **Configuração** | Arquivo simples | Sistema avançado |
| **Status** | Manutenção | Desenvolvimento ativo |

## Navegação por Tópico

### 🚀 Começando
- [Instalação do Docker](INSTALLATION.md#docker)
- [Primeiro documento](cli/usage-guide.md#fluxo-de-trabalho-típico)
- [Configuração básica](cli/configuration.md#arquivo-de-configuração)

### 📝 Uso Diário
- [Comandos da CLI](cli/usage-guide.md#comandos-principais)
- [File watching](cli/usage-guide.md#desenvolvimento)
- [Solução de problemas](cli/usage-guide.md#solução-de-problemas)

### ⚙️ Configuração Avançada
- [Perfis de configuração](cli/configuration.md#perfis-de-configuração)
- [Integração VS Code](cli/configuration.md#vs-code)
- [Docker customizado](cli/configuration.md#docker)

### 🔨 Desenvolvimento
- [Arquitetura da CLI](cli/architecture.md)
- [Padrões de código](cli/output-standards.md)
- [Sistema de testes](cli/development-guide.md#testes)

### 🐛 Troubleshooting
- [Problemas comuns](FAQ.md)
- [Debug da configuração](cli/configuration.md#debugging-de-configuração)
- [Logs e diagnóstico](cli/usage-guide.md#obtendo-ajuda)

## Status do Projeto

### ✅ Implementado
- ✅ CLI base com Cobra
- ✅ Sistema de configuração
- ✅ Integração Docker
- ✅ Comando `setup`
- ✅ Build multiplataforma
- ✅ CI/CD pipeline
- ✅ Documentação completa

### 🚧 Em Desenvolvimento
- 🚧 Comando `init` (templates)
- 🚧 Comando `build` (compilação)
- 🚧 Comando `watch` (file watching)
- 🚧 Comando `status` (monitoramento)

### 📋 Planejado
- 📋 Interface web (opcional)
- 📋 Plugins de terceiros
- 📋 Package manager LaTeX
- 📋 Cloud integration

## Contribuindo

### Para Usuários
- 🐛 [Reportar bugs](https://github.com/martinsmiguel/latex-docker-env/issues)
- 💡 [Sugerir melhorias](https://github.com/martinsmiguel/latex-docker-env/discussions)
- 📖 [Melhorar documentação](https://github.com/martinsmiguel/latex-docker-env/pulls)

### Para Desenvolvedores
- 🔧 [Guia de contribuição](cli/development-guide.md)
- 🧪 [Escrever testes](cli/development-guide.md#testes)
- 📝 [Implementar comandos](cli/development-guide.md#próximos-passos)

## Links Úteis

- **Repositório**: https://github.com/martinsmiguel/latex-docker-env
- **Issues**: https://github.com/martinsmiguel/latex-docker-env/issues
- **Discussões**: https://github.com/martinsmiguel/latex-docker-env/discussions
- **Releases**: https://github.com/martinsmiguel/latex-docker-env/releases

## Licença

Este projeto está licenciado sob a MIT License. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
