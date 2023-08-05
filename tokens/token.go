package tokens

import (
	"fmt"
)

type Token struct {
	TokenType TokenType
	Value     string
	Line      int
}

func (token Token) String() string {
	return fmt.Sprintf("%3d: %v - %s", token.Line, token.TokenType, token.Value)
}
