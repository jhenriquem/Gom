package utils

import "github.com/jhenriquem/Gom/internal/editor"

func MessageCommand(message string) {
	editor.GOM.CommandLine.CrrCommand = []rune{}
	for _, char := range message {
		editor.GOM.CommandLine.CrrCommand = append(editor.GOM.CommandLine.CrrCommand, char)
	}
}
