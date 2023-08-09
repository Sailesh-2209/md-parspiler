package ast

type WordType int

const (
	WORD_REGULAR WordType = iota
	WORD_ITALICS
	WORD_BOLD
	WORD_BOLD_ITALICS
	WORD_CODE
)
