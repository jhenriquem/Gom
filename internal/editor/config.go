package editor

import (
	"github.com/jhenriquem/Gom/internal/buffer"
)

type commandLine struct {
	CrrCommand []rune
	CrrColumn  int
}

func (c *commandLine) MoveCursor(interator int) {
	if c.CrrColumn >= 1 && c.CrrColumn <= len(c.CrrCommand) {
		c.CrrColumn += interator
	}
}

type EditorStruct struct {
	CrrMode     string
	CrrBuffer   *buffer.BufferStruct
	Buffers     []buffer.BufferStruct
	Running     bool
	CommandLine commandLine
	CrrBffIndex int
}

var GOM EditorStruct = EditorStruct{
	Running:     true,
	CrrMode:     "NORMAL",
	CrrBffIndex: 0,
	CommandLine: commandLine{
		CrrCommand: []rune{},
		CrrColumn:  1,
	},
}

func Inicializer() {
	GOM.CrrBuffer = GOM.InicializeBuffer()
	GOM.LoadArgsFile()
}
