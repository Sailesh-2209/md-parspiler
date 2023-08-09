package ast

import (
	"fmt"
)

type ASTNode interface {
	Render() string
}

type Word struct {
	WordType WordType
	Word     string
}

type Heading struct {
	Level   int
	Heading []Word
}

func (word Word) Render() string {
	switch word.WordType {
	case WORD_REGULAR:
		return word.Word
	case WORD_BOLD:
		return fmt.Sprintf("<b>%s</b>", word.Word)
	case WORD_ITALICS:
		return fmt.Sprintf("<i>%s</i>", word.Word)
	case WORD_BOLD_ITALICS:
		return fmt.Sprintf("<b><i>%s<i></b>", word.Word)
	case WORD_CODE:
		return fmt.Sprintf("<span style=\"font-family: monospace\">%s</span>", word.Word)
	}

	return word.Word
}

func (heading Heading) Render() string {
	headingString := ""

	for _, word := range heading.Heading {
		if headingString == "" {
			headingString = word.Render()
		} else {
			headingString = fmt.Sprintf("%s %s", headingString, word.Render())
		}
	}

	return fmt.Sprintf("<h%d>%s</h%d>", heading.Level, headingString, heading.Level)
}

func (word Word) String() string {
	return word.Render()
}

func (heading Heading) String() string {
	return heading.Render()
}
