package event

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/screen"
)

func CommandToNormal() {
	editor.GOM.CrrMode = "NORMAL"
	screen.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
}

func NormalToCommand() {
	screen.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBar)
	editor.GOM.CommandLine.CrrColumn = 1
	editor.GOM.CrrMode = "COMMAND"
	editor.GOM.CommandLine.CrrCommand = []rune{':'}
}
