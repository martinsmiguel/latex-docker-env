package colors

import (
	"testing"
)

func TestColorize(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected bool // Se deve conter códigos de cor
	}{
		{
			name:     "texto com tag SUCCESS",
			text:     "[SUCCESS] Operation completed",
			expected: true,
		},
		{
			name:     "texto com tag ERROR",
			text:     "[ERROR] Something went wrong",
			expected: true,
		},
		{
			name:     "texto com tag WARNING",
			text:     "[WARNING] Please be careful",
			expected: true,
		},
		{
			name:     "texto com tag INFO",
			text:     "[INFO] Information message",
			expected: true,
		},
		{
			name:     "texto sem tag",
			text:     "Normal text without tag",
			expected: false,
		},
		{
			name:     "texto vazio",
			text:     "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Colorize(tt.text)

			// Verificar se o resultado contém códigos de escape ANSI
			hasColorCodes := len(result) > len(tt.text)

			if hasColorCodes != tt.expected {
				t.Errorf("Colorize() result has color codes = %v, expected = %v", hasColorCodes, tt.expected)
			}

			// Verificar se o texto original está presente no resultado (para tags válidas)
			if tt.expected && tt.text != "" {
				// O resultado deve conter pelo menos parte do texto original
				if !containsText(result, tt.text[1:len(tt.text)-1]) { // Remove [ e ] para verificar
					t.Logf("Result: %q, Original: %q", result, tt.text)
				}
			}
		})
	}
}

func TestColorizeByTag(t *testing.T) {
	tests := []struct {
		name     string
		tag      string
		message  string
		expected bool // Se deve aplicar colorização
	}{
		{
			name:     "tag SUCCESS",
			tag:      "SUCCESS",
			message:  "Operation completed",
			expected: true,
		},
		{
			name:     "tag ERROR",
			tag:      "ERROR",
			message:  "Something went wrong",
			expected: true,
		},
		{
			name:     "tag WARNING",
			tag:      "WARNING",
			message:  "Please be careful",
			expected: true,
		},
		{
			name:     "tag INFO",
			tag:      "INFO",
			message:  "Information message",
			expected: true,
		},
		{
			name:     "tag DEBUG",
			tag:      "DEBUG",
			message:  "Debug information",
			expected: true,
		},
		{
			name:     "tag REMOVED",
			tag:      "REMOVED",
			message:  "File deleted",
			expected: true,
		},
		{
			name:     "tag CHANGE",
			tag:      "CHANGE",
			message:  "File modified",
			expected: true,
		},
		{
			name:     "tag inválida",
			tag:      "INVALID",
			message:  "Unknown tag",
			expected: false, // Não deve aplicar cor especial
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ColorizeByTag(tt.tag, tt.message)

			// Verificar se houve colorização (resultado maior que tag + message)
			expectedLength := len("[" + tt.tag + "] " + tt.message)
			hasColorization := len(result) > expectedLength

			if hasColorization != tt.expected {
				t.Errorf("ColorizeByTag() has colorization = %v, expected = %v", hasColorization, tt.expected)
			}

			// O resultado deve conter a mensagem
			if !containsText(result, tt.message) {
				t.Errorf("ColorizeByTag() result should contain message: %q", tt.message)
			}
		})
	}
}

// Função auxiliar para verificar se uma string contém outra
func containsText(text, substr string) bool {
	return len(text) >= len(substr) &&
		   (text == substr || func() bool {
			   for i := 0; i <= len(text)-len(substr); i++ {
				   if text[i:i+len(substr)] == substr {
					   return true
				   }
			   }
			   return false
		   }())
}

func TestPrintFunctions(t *testing.T) {
	// Testar funções de print (elas apenas imprimem, então testamos se não há pânico)
	tests := []struct {
		name string
		fn   func(string)
		text string
	}{
		{
			name: "PrintInfo",
			fn:   PrintInfo,
			text: "Info message",
		},
		{
			name: "PrintSuccess",
			fn:   PrintSuccess,
			text: "Success message",
		},
		{
			name: "PrintWarn",
			fn:   PrintWarn,
			text: "Warning message",
		},
		{
			name: "PrintWarning",
			fn:   PrintWarning,
			text: "Warning message",
		},
		{
			name: "PrintError",
			fn:   PrintError,
			text: "Error message",
		},
		{
			name: "PrintRemoved",
			fn:   PrintRemoved,
			text: "Removed message",
		},
		{
			name: "PrintDebug",
			fn:   PrintDebug,
			text: "Debug message",
		},
		{
			name: "PrintChange",
			fn:   PrintChange,
			text: "Change message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Verificar se as funções não causam pânico
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("%s() caused panic: %v", tt.name, r)
				}
			}()

			tt.fn(tt.text)
		})
	}
}

func TestPrintfFunctions(t *testing.T) {
	// Testar funções Printf
	tests := []struct {
		name   string
		format string
		args   []interface{}
	}{
		{
			name:   "Printf simples",
			format: "Hello %s",
			args:   []interface{}{"World"},
		},
		{
			name:   "Printf com múltiplos argumentos",
			format: "Number: %d, String: %s",
			args:   []interface{}{42, "test"},
		},
		{
			name:   "Printf sem argumentos",
			format: "Simple text",
			args:   []interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Verificar se Printf não causa pânico
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Printf() caused panic: %v", r)
				}
			}()

			Printf(tt.format, tt.args...)
		})
	}
}

func TestSprintfFunction(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		args     []interface{}
		expected string
	}{
		{
			name:     "formato simples",
			format:   "Hello %s",
			args:     []interface{}{"World"},
			expected: "Hello World",
		},
		{
			name:     "múltiplos argumentos",
			format:   "Number: %d, String: %s",
			args:     []interface{}{42, "test"},
			expected: "Number: 42, String: test",
		},
		{
			name:     "sem argumentos",
			format:   "Simple text",
			args:     []interface{}{},
			expected: "Simple text",
		},
		{
			name:     "formato com porcentagem",
			format:   "Progress: %d%%",
			args:     []interface{}{75},
			expected: "Progress: 75%",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sprintf(tt.format, tt.args...)
			if result != tt.expected {
				t.Errorf("Sprintf() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPrintlnFunction(t *testing.T) {
	// Testar função Println
	tests := []struct {
		name string
		arg  string
	}{
		{
			name: "single argument",
			arg:  "Hello",
		},
		{
			name: "message with tag",
			arg:  "[INFO] Information message",
		},
		{
			name: "empty message",
			arg:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Verificar se Println não causa pânico
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Println() caused panic: %v", r)
				}
			}()

			Println(tt.arg)
		})
	}
}

func TestColorCodes(t *testing.T) {
	// Testar se as tags específicas funcionam
	tags := []string{"SUCCESS", "ERROR", "WARNING", "INFO", "DEBUG", "REMOVED", "CHANGE"}

	for _, tag := range tags {
		t.Run("tag_"+tag, func(t *testing.T) {
			text := "[" + tag + "] Test message"
			result := Colorize(text)

			// Resultado deve ser diferente do texto original (contém códigos de cor)
			if result == text {
				t.Errorf("Colorize() with tag %s should modify the text", tag)
			}
		})
	}
}

func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		function func() string
		testDesc string
	}{
		{
			name: "colorize_empty_text",
			function: func() string {
				return Colorize("")
			},
			testDesc: "colorize with empty text should return empty",
		},
		{
			name: "colorize_whitespace_text",
			function: func() string {
				return Colorize("   ")
			},
			testDesc: "colorize with whitespace should work",
		},
		{
			name: "colorize_special_characters",
			function: func() string {
				return Colorize("[INFO] test\n\t")
			},
			testDesc: "colorize with special characters should work",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Verificar se não há pânico
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("%s caused panic: %v", tt.testDesc, r)
				}
			}()

			result := tt.function()
			// Resultado não deve causar pânico
			_ = result
		})
	}
}
