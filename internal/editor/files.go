package editor

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jhenriquem/Gom/internal/buffer"
)

func (e *EditorStruct) ScanFile(file *os.File) {
	// Se o arquivo existir na lista de buffers
	Exist := false

	for i, buffer := range e.Buffers {
		if buffer.NameFile == file.Name() {
			e.CrrBuffer = &e.Buffers[i]
			e.CrrBffIndex = i
			Exist = true
		}
	}

	if Exist {
		return
	}

	// Se o arquivo não existir na lista de buffers
	newBuffer := buffer.BufferStruct{NameFile: file.Name()}

	e.Buffers = append(e.Buffers, newBuffer)
	e.CrrBuffer = &e.Buffers[len(e.Buffers)-1]

	e.CrrBffIndex = len(e.Buffers) - 1

	lineIndex := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		scannedLine := scanner.Text()
		e.CrrBuffer.Text = append(e.CrrBuffer.Text, []rune{})
		for _, ch := range scannedLine {
			e.CrrBuffer.Text[lineIndex] = append(e.CrrBuffer.Text[lineIndex], rune(ch))
		}
		lineIndex++
	}
	if lineIndex <= 1 {
		e.CrrBuffer.Text = append(e.CrrBuffer.Text, []rune{})
	}
}

func (e *EditorStruct) WriteFile() {
	file, err := os.Create(e.CrrBuffer.NameFile)
	if err != nil {
	}

	writer := bufio.NewWriter(file)
	for _, line := range e.CrrBuffer.Text {
		linetoWrite := string(line) + "\n"
		writer.WriteString(linetoWrite)
	}
	writer.Flush()
}

func (e *EditorStruct) SaveFile(isNewFile bool) string {
	if e.CrrBuffer.NameFile != "" {
		e.WriteFile()
		if isNewFile {
			return fmt.Sprintf("create %s", e.CrrBuffer.NameFile)
		} else {
			return fmt.Sprintf("save %s", e.CrrBuffer.NameFile)
		}
	} else {
		return "unnamed file"
	}
}
