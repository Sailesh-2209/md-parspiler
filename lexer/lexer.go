package lexer

import (
	"md_parspiler/tokens"
	"regexp"
)

func Lex(str string) []tokens.Token {
	str = cleanUp(str)

	var lineNo int = 1
	var tokensList []tokens.Token

	var token tokens.Token

	for i := 0; i < len(str); {
		c := str[i]

		switch c {
		case ' ':
			token = tokens.Token{
				TokenType: tokens.SPACE,
				Value:     " ",
				Line:      lineNo,
			}
			i++
		case '\n':
			token = tokens.Token{
				TokenType: tokens.NEW_LINE,
				Value:     "\\n",
				Line:      lineNo,
			}
			i++
			lineNo++
		case '!':
			token = tokens.Token{
				TokenType: tokens.BANG,
				Value:     "!",
				Line:      lineNo,
			}
			i++
		case '`':
			token = tokens.Token{
				TokenType: tokens.TICK,
				Value:     "`",
				Line:      lineNo,
			}
			i++
		case '_':
			token = tokens.Token{
				TokenType: tokens.UNDERSCORE,
				Value:     "_",
				Line:      lineNo,
			}
			i++
		case '-':
			token = tokens.Token{
				TokenType: tokens.MINUS,
				Value:     "-",
				Line:      lineNo,
			}
			i++
		case '{':
			token = tokens.Token{
				TokenType: tokens.LEFT_BRACE,
				Value:     "{",
				Line:      lineNo,
			}
			i++
		case '}':
			token = tokens.Token{
				TokenType: tokens.RIGHT_BRACE,
				Value:     "}",
				Line:      lineNo,
			}
			i++
		case '[':
			token = tokens.Token{
				TokenType: tokens.LEFT_BRACKET,
				Value:     "[",
				Line:      lineNo,
			}
			i++
		case ']':
			token = tokens.Token{
				TokenType: tokens.RIGHT_BRACKET,
				Value:     "]",
				Line:      lineNo,
			}
			i++
		case '(':
			token = tokens.Token{
				TokenType: tokens.LEFT_PAREN,
				Value:     "(",
				Line:      lineNo,
			}
			i++
		case ')':
			token = tokens.Token{
				TokenType: tokens.RIGHT_PAREN,
				Value:     ")",
				Line:      lineNo,
			}
			i++
		case '#':
			token = matchHash(str, &i, lineNo, tokensList)
		case '*':
			token = matchStar(str, &i, lineNo)
		default:
			token = matchWord(str, &i, lineNo)
		}

		tokensList = append(tokensList, token)
	}

	return tokensList
}

// A utility function that makes it easier to tokenize markdown
func cleanUp(str string) string {
	// Uniformity in new line characters across operating systems
	newLineRegex := regexp.MustCompile("(\r\n)|(\r)")
	str = newLineRegex.ReplaceAllString(str, "\n")

	// Empty heading lines are ignored
	emptyHeadingRegex := regexp.MustCompile(`(?m)^#+$`)
	str = emptyHeadingRegex.ReplaceAllString(str, "")

	// Lines containing only * characters are ignored
	emptyStarRegex := regexp.MustCompile(`(?m)^\*+$`)
	str = emptyStarRegex.ReplaceAllString(str, "")

	// More than three surrounding stars around a
	// word can be reduced to two surrounding stars
	duplicateBoldStarsRegex := regexp.MustCompile(`\*\*\*\*+((\S.*\S)|(\S))\*\*\*\*+`)
	str = duplicateBoldStarsRegex.ReplaceAllString(str, "**$1**")

	// All tabs are replaced by 4 spaces
	tabRegex := regexp.MustCompile(`\t`)
	str = tabRegex.ReplaceAllString(str, "    ")

	return str
}

// Matches consecutive '#' characters followed by a space
func matchHash(str string, i *int, lineNo int, tokensList []tokens.Token) tokens.Token {
	cnt := 0

	for *i < len(str) && str[*i] == '#' {
		*i++
		cnt++
	}

	var tokenType tokens.TokenType
	var value string

	prevToken := previousToken(tokensList)

	if (*i < len(str) && str[*i] == ' ') && (prevToken == tokens.Token{} || prevToken.TokenType == tokens.NEW_LINE) {
		switch cnt {
		case 1:
			tokenType = tokens.HASH1
			value = "#"
		case 2:
			tokenType = tokens.HASH2
			value = "##"
		case 3:
			tokenType = tokens.HASH3
			value = "###"
		case 4:
			tokenType = tokens.HASH4
			value = "####"
		case 5:
			tokenType = tokens.HASH5
			value = "#####"
		case 6:
			tokenType = tokens.HASH6
			value = "######"
		default:
			tokenType = tokens.WORD
			value = str[*i-cnt : *i]
		}
	} else {
		tokenType = tokens.WORD
		value = str[*i-cnt : *i]
	}

	return tokens.Token{
		TokenType: tokenType,
		Value:     value,
		Line:      lineNo,
	}
}

// Matches consecutive '*' characters followed by a non-whitespace character
func matchStar(str string, i *int, lineNo int) tokens.Token {
	cnt := 0

	for *i < len(str) && str[*i] == '*' {
		*i++
		cnt++
	}

	var tokenType tokens.TokenType
	var value string

	switch cnt {
	case 1:
		tokenType = tokens.STAR1
		value = tokenType.String()
	case 2:
		tokenType = tokens.STAR2
		value = tokenType.String()
	case 3:
		tokenType = tokens.STAR3
		value = tokenType.String()
	default:
		tokenType = tokens.WORD
		value = str[*i-cnt : *i]
	}

	return tokens.Token{
		TokenType: tokenType,
		Value:     value,
		Line:      lineNo,
	}
}

// Returns a token of type WORD from current character until next
// non-whitespace character
func matchWord(str string, i *int, lineNo int) tokens.Token {
	start := *i

	flag := true

	for *i < len(str) && flag {
		switch str[*i] {
		case ' ', '\n', '_', '*', '`':
			flag = false
		default:
			*i++
		}
	}

	return tokens.Token{
		TokenType: tokens.WORD,
		Value:     str[start:*i],
		Line:      lineNo,
	}
}

func previousToken(tokensList []tokens.Token) tokens.Token {
	n := len(tokensList)
	if n == 0 {
		return tokens.Token{}
	} else {
		return tokensList[n-1]
	}
}
