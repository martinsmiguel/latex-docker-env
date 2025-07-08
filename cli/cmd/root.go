package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/martinsmiguel/latex-docker-env/cli/internal/commands"
)

var (
	cfgFile string
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "ltx",
	Short: "LaTeX Docker Environment CLI",
	Long: `ltx - Uma CLI moderna para desenvolvimento LaTeX usando Docker.

Oferece compilação automática, templates customizáveis e
um ambiente de desenvolvimento isolado e reproduzível.`,
	Version: "2.0.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Flags globais
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "arquivo de configuração (padrão: ./config/latex-cli.conf)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "saída detalhada")

	// Comandos
	rootCmd.AddCommand(commands.SetupCmd)
	rootCmd.AddCommand(commands.InitCmd)
	rootCmd.AddCommand(commands.BuildCmd)
	rootCmd.AddCommand(commands.WatchCmd)
	rootCmd.AddCommand(commands.StatusCmd)
	rootCmd.AddCommand(commands.CleanCmd)
	rootCmd.AddCommand(commands.ShellCmd)
	rootCmd.AddCommand(commands.LogsCmd)
	rootCmd.AddCommand(commands.ResetCmd)
	rootCmd.AddCommand(commands.BackupCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./config")
		viper.SetConfigName("latex-cli")
		viper.SetConfigType("conf")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if verbose {
			fmt.Println("Usando arquivo de configuração:", viper.ConfigFileUsed())
		}
	}
}
