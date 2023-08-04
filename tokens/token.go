package tokens

import (
	"fmt"
)

type Token struct {
	tokenType TokenType
	value     string
	line      int
}

func (token Token) String() string {
	return fmt.Sprintf("%3d: %v - %s", token.line, token.tokenType, token.value)
}
