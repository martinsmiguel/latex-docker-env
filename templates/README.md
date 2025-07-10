# Templates de Usuário

Este diretório é destinado aos templates que você baixa e utiliza localmente.

## Como funciona

- **Templates do CLI**: Ficam em `cli/templates/` e são versionados no repositório
- **Templates do usuário**: Ficam neste diretório (`templates/`) e **não são versionados**

## Uso

1. Baixe ou crie seus templates personalizados aqui
2. Use o comando `ltx template list` para listá-los
3. Use o comando `ltx init --template <nome>` para criar um projeto baseado em um template

## Estrutura recomendada

```
templates/
├── meu-template/
│   ├── template.yaml      # Configuração do template
│   ├── main.tex          # Arquivo principal
│   └── ...               # Outros arquivos do template
└── outro-template/
    ├── template.yaml
    ├── main.tex
    └── ...
```

> **Nota**: Este diretório está no `.gitignore` para não interferir no versionamento do projeto.
