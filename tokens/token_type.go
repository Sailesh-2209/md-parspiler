package tokens

import (
	"fmt"
)

type TokenType int

const (
	SPACE    TokenType = iota // all tab characters are replaced by spaces by the parser
	NEW_LINE                  // \r and \r\n are replaced by \n

	BANG // matches '!'
	TICK // matches '`'

	LEFT_BRACE    // matches '{'
	RIGHT_BRACE   // matches '}'
	LEFT_BRACKET  // matches '['
	RIGHT_BRACKET // matches ']'
	LEFT_PAREN    // matches '('
	RIGHT_PAREN   // matches ')'

	HASH1 // matches '#'
	HASH2 // matches '##'
	HASH3 // matches '###'
	HASH4 // matches '####'
	HASH5 // matches '#####
	HASH6 // matches '######'

	STAR1 // matches '*'
	STAR2 // matches '**'
	STAR3 // matches '***'

	WORD // matches any word

	EOF // represents the end of file. last in the list of tokens
)

func (tt TokenType) String() string {
	switch tt {
	case SPACE:
		return "SPACE"
	case NEW_LINE:
		return "NEW_LINE"
	case BANG:
		return "BANG"
	case TICK:
		return "TICK"
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case LEFT_BRACKET:
		return "LEFT_BRACKET"
	case RIGHT_BRACKET:
		return "RIGHT_BRACKET"
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case HASH1:
		return "HASH1"
	case HASH2:
		return "HASH2"
	case HASH3:
		return "HASH3"
	case HASH4:
		return "HASH4"
	case HASH5:
		return "HASH5"
	case HASH6:
		return "HASH6"
	case STAR1:
		return "STAR1"
	case STAR2:
		return "STAR2"
	case STAR3:
		return "STAR3"
	case WORD:
		return "WORD"
	case EOF:
		return "EOF"
	default:
		return fmt.Sprint(int(tt))
	}
}
