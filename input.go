package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
)

type Input interface {}

type ShortAnswerField struct {
	textinput textinput.Model
}
// textinput
func NewShortAnswerField() {}

type LongAnswerField struct {
	textarea textarea.Model
}
// textarea
func NewLongAnswerField() *LongAnswerField{
	ta := textarea.New()
	return &LongAnswerField{ta}
}