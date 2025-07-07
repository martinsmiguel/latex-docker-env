# Padrões de Saída da CLI

## Convenções de Output

Para manter consistência e compatibilidade em diferentes terminais, utilizamos apenas caracteres ASCII padrão:

### Prefixos de Status
```
[OK]      - Operação bem-sucedida
[ERROR]   - Erro crítico
[WARN]    - Aviso, operação continua
[INFO]    - Informação geral
[WIP]     - Work in Progress (desenvolvimento)
[SUCCESS] - Sucesso final de operação complexa
>>        - Indicador de progresso/início de operação
```

### Exemplos de Uso

#### Operação bem-sucedida
```
[OK] Docker verificado
[OK] Imagem LaTeX pronta
[OK] Estrutura de diretórios criada
```

#### Progresso de operação
```
>> Configurando ambiente LaTeX Docker...
>> Verificando imagem LaTeX: blang/latex:ubuntu
>> Baixando imagem blang/latex:ubuntu...
```

#### Avisos e informações
```
[WARN] Configuração VS Code opcional falhou: arquivo não encontrado
[INFO] Configurações VS Code disponíveis em config/vscode/settings.json
```

#### Resultado final
```
[SUCCESS] Ambiente configurado com sucesso!
Execute 'ltx init' para criar seu primeiro documento.
```

#### Comando em desenvolvimento
```
[WIP] Comando init em desenvolvimento...
```

### Rationale

1. **Compatibilidade**: Caracteres ASCII funcionam em qualquer terminal
2. **Legibilidade**: Prefixos claros facilitam parsing e leitura
3. **Consistência**: Padrão único em toda a aplicação
4. **Funcionalidade**: Facilita grep, logs e scripts

### Anti-padrões (evitar)

❌ Emojis: 🔧 ✅ ⚠️ 📦 🚧 🎉  
❌ Símbolos Unicode: ▶ ● ◆ ♦ ✓ ✗  
❌ Cores sem fallback: apenas cores podem não funcionar em todos os terminais  

### Implementação

Cada comando deve seguir este padrão:

```go
fmt.Println(">> Iniciando operação...")
// ... lógica ...
fmt.Println("[OK] Operação concluída")

// Para erros
return fmt.Errorf("[ERROR] Falha na operação: %w", err)

// Para avisos
fmt.Printf("[WARN] Aviso: %s\n", warning)
```
