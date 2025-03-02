package renderer

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/internal/colors"
	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/screen"
)

func CommandLine() {
	width, height := screen.Screen.Size()

	bgStyle := tcell.StyleDefault.Background(colors.ColorBgCommandLine).Foreground(colors.ColorFgCommandLine)

	for x := 0; x < width; x++ {
		char := ' '

		if x < len(editor.GOM.CommandLine.CrrCommand) {
			char = editor.GOM.CommandLine.CrrCommand[x]
		}
		screen.Screen.SetContent(x, height-1, char, nil, bgStyle)
	}

	// Mostrar o cursor na CommandLine apenas no modo de comando
	if editor.GOM.CrrMode == "COMMAND" {
		screen.Screen.ShowCursor(editor.GOM.CommandLine.CrrColumn, height-1)
	}
	// Atualizar a tela
	screen.Screen.Show()
}
