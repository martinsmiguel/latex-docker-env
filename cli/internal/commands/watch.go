package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

var (
	watchDebounce time.Duration = 500 * time.Millisecond
)

var WatchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Monitora arquivos e compila automaticamente",
	Long: `Monitora mudanças nos arquivos LaTeX e recompila automaticamente.

Este comando:
1. Inicia o monitoramento de arquivos .tex, .bib e .cls
2. Recompila automaticamente quando detecta mudanças
3. Usa debouncing para evitar compilações excessivas
4. Mantém logs de todas as compilações`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return watchProject()
	},
}

func watchProject() error {
	fmt.Println(">> Iniciando modo de observação...")

	// Verificar se há compilações em andamento (para watch, pode ser que queiramos parar watch anterior)
	isRunning, err := checkRunningCompilation()
	if err != nil {
		fmt.Printf("[WARN] Erro ao verificar compilações: %v\n", err)
	} else if isRunning {
		fmt.Println("[WARN] Há processos de compilação em andamento!")
		if askUserConfirmation("Deseja encerrar processos anteriores e iniciar novo modo watch?") {
			if err := killRunningCompilation(); err != nil {
				return fmt.Errorf("erro ao encerrar processos: %w", err)
			}
			fmt.Println("[SUCCESS] Processos anteriores encerrados")
		} else {
			return fmt.Errorf("operação cancelada pelo usuário")
		}
	}

	fmt.Println("[INFO] Monitorando mudanças em arquivos LaTeX...")
	fmt.Println("[INFO] Pressione Ctrl+C para parar")

	// Verificar se projeto existe
	sourceDir := "src"
	if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
		return fmt.Errorf("diretório %s não encontrado. Execute 'ltx init' primeiro", sourceDir)
	}

	// Compilação inicial
	fmt.Println("[INFO] Compilação inicial...")
	if err := buildProject(); err != nil {
		fmt.Printf("[WARN] Falha na compilação inicial: %v\n", err)
	}

	// Configurar watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("erro ao criar watcher: %w", err)
	}
	defer watcher.Close()

	// Adicionar diretórios ao watcher
	if err := addWatchPaths(watcher, sourceDir); err != nil {
		return fmt.Errorf("erro ao configurar monitoramento: %w", err)
	}

	// Canal para debouncing
	debounceTimer := time.NewTimer(0)
	<-debounceTimer.C // drain the timer

	// Loop principal
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}

			// Filtrar apenas arquivos relevantes
			if isRelevantFile(event.Name) && (event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create) {
				fmt.Printf("[CHANGE] %s\n", event.Name)

				// Reset do timer de debounce
				debounceTimer.Reset(watchDebounce)
			}

		case <-debounceTimer.C:
			// Compilar após debounce
			fmt.Println("[INFO] Recompilando...")
			start := time.Now()

			if err := buildProject(); err != nil {
				fmt.Printf("[ERROR] Falha na compilação: %v\n", err)
			} else {
				duration := time.Since(start)
				fmt.Printf("[SUCCESS] Recompilação concluída em %v\n", duration.Round(time.Millisecond*100))
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			log.Printf("[ERROR] Erro no watcher: %v", err)
		}
	}
}

func addWatchPaths(watcher *fsnotify.Watcher, sourceDir string) error {
	// Adicionar diretório raiz do projeto
	if err := watcher.Add(sourceDir); err != nil {
		return err
	}

	// Caminhar por todos os subdiretórios
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return watcher.Add(path)
		}

		return nil
	})
}

func isRelevantFile(filename string) bool {
	ext := filepath.Ext(filename)
	relevantExts := []string{".tex", ".bib", ".cls", ".sty"}

	for _, relevantExt := range relevantExts {
		if ext == relevantExt {
			return true
		}
	}

	return false
}
