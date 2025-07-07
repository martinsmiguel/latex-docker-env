package colors

import (
	"fmt"
	"regexp"
	"strings"
)

// Códigos de cores ANSI
const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"

	// Cores de texto
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Gray    = "\033[90m"

	// Cores brilhantes
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"
)

// Padrão regex para identificar tags
var tagPattern = regexp.MustCompile(`\[(INFO|SUCCESS|WARN|WARNING|ERROR|REMOVED|DEBUG|CHANGE|OK)\]`)

// Colorize aplica cores às tags das mensagens
func Colorize(text string) string {
	return tagPattern.ReplaceAllStringFunc(text, func(match string) string {
		tag := strings.Trim(match, "[]")
		switch strings.ToUpper(tag) {
		case "SUCCESS":
			return BrightGreen + Bold + "[" + tag + "]" + Reset
		case "ERROR":
			return BrightRed + Bold + "[" + tag + "]" + Reset
		case "WARN", "WARNING":
			return BrightYellow + Bold + "[" + tag + "]" + Reset
		case "INFO":
			return BrightBlue + Bold + "[" + tag + "]" + Reset
		case "REMOVED":
			return BrightMagenta + Bold + "[" + tag + "]" + Reset
		case "DEBUG":
			return Gray + "[" + tag + "]" + Reset
		case "CHANGE":
			return BrightCyan + Bold + "[" + tag + "]" + Reset
		case "OK":
			return BrightGreen + "[" + tag + "]" + Reset
		default:
			return match
		}
	})
}

// ColorizeByTag aplica cor específica baseada no tipo de tag
func ColorizeByTag(tag, message string) string {
	coloredTag := Colorize("[" + tag + "]")
	return coloredTag + " " + message
}

// Funções de conveniência para imprimir mensagens coloridas
func PrintInfo(message string) {
	fmt.Println(Colorize("[INFO] " + message))
}

func PrintSuccess(message string) {
	fmt.Println(Colorize("[SUCCESS] " + message))
}

func PrintWarn(message string) {
	fmt.Println(Colorize("[WARN] " + message))
}

func PrintError(message string) {
	fmt.Println(Colorize("[ERROR] " + message))
}

func PrintRemoved(message string) {
	fmt.Println(Colorize("[REMOVED] " + message))
}

func PrintDebug(message string) {
	fmt.Println(Colorize("[DEBUG] " + message))
}

func PrintChange(message string) {
	fmt.Println(Colorize("[CHANGE] " + message))
}

// Printf com colorização automática
func Printf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	fmt.Print(Colorize(message))
}

// Println com colorização automática
func Println(message string) {
	fmt.Println(Colorize(message))
}

// Sprintf com colorização automática
func Sprintf(format string, args ...interface{}) string {
	message := fmt.Sprintf(format, args...)
	return Colorize(message)
}
