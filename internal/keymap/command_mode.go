package keymap

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/event"
)

func InputInCommandMode(eventKey *tcell.EventKey) {
	switch eventKey.Key() {
	default:
		if eventKey.Rune() != 0 {
			event.InsertInCommand(eventKey)
		}
	case tcell.KeyEnter:
		event.RunCommand()
		event.CommandToNormal()
	case tcell.KeyEscape:
		event.CommandToNormal()
		editor.GOM.CommandLine.CrrCommand = []rune{}
	case tcell.KeyLeft:
		editor.GOM.CommandLine.MoveCursor(-1)
	case tcell.KeyRight:
		editor.GOM.CommandLine.MoveCursor(1)

	case tcell.KeyBackspace, tcell.KeyBackspace2:
		event.BackSpaceInCommand()
	}
}
