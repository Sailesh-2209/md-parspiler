package main

type TokenType int

const (
	NEW_LINE TokenType = iota
	HASH1
	HASH2
	HASH3
	HASH4
	BANG
	WORD
	LINK
	LINK_TEXT
	NULL
	EOF
)

type Token struct {
	tokenType TokenType
	value     string
	line      int
}
