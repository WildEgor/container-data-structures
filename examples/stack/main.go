package main

import (
	"fmt"
	"github.com/WildEgor/container-data-structures/pkg/stack"
)

// TextEditor represent simple editor mechanism
type TextEditor struct {
	text  string
	stack *stack.Stack[string]
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		stack: stack.New[string](),
	}
}

func (te *TextEditor) Type(text string) {
	te.stack.Push(te.text)
	te.text += text
}

func (te *TextEditor) Undo() {
	if !te.stack.Empty() {
		te.text = te.stack.Pop()
	}
}

func main() {
	te := NewTextEditor()
	te.Type("HELLO")
	te.Type("WORLD")
	te.Undo()

	fmt.Println(te.text)
}
