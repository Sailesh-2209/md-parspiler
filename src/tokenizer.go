package main

/*
Parses through a string and returns a list of tokens.
*/
func Tokenize(str string) []Token {
	var lineNo int = 1
	var tokens []Token

	for i := 0; i < len(str); {
		c := str[i]

		var token Token

		switch c {
		case '#':
			token = matchHash(&i, str, lineNo)
		case '!':
			token = Token{
				tokenType: BANG,
				value:     "!",
				line:      lineNo,
			}
		case '[':
			token = matchLinkText(&i, str, &lineNo)
		case '(':
			// TODO
		case '\n':
		case '\r':
			lineNo++
			token = Token{
				tokenType: NEW_LINE,
				value:     "\\n",
				line:      lineNo,
			}
		case ' ':
			continue
		}

		if (token != Token{}) {
			tokens = append(tokens, token)
		}
	}

	tokens = append(tokens, Token{
		tokenType: EOF,
		value:     "",
		line:      lineNo,
	})
	return tokens
}

/*
Matches consecutive # elements.
If more than 4 consecutive # elements are present, then a NULL Token is returned.
*/
func matchHash(i *int, str string, lineNo int) Token {
	cnt := 0

	for str[*i] == '#' && *i < len(str) {
		*i++
		cnt++
	}

	var tokenType TokenType

	switch cnt {
	case 1:
		tokenType = HASH1
	case 2:
		tokenType = HASH2
	case 3:
		tokenType = HASH3
	case 4:
		tokenType = HASH4
	default:
		tokenType = NULL
	}

	return Token{
		tokenType: tokenType,
		value:     str[*i-cnt : *i],
		line:      lineNo,
	}
}

/*
This function matches for link text. It is triggered whenever a '[' character is
encountered. It tries to match until it finds a ']' character. If it reaches
end of file before reaching a ']' character, then a NULL Token is returned.
*/
func matchLinkText(i *int, str string, lineNo *int) Token {
	cnt := 1 // counter for bracket matching. this allows nested square brackets
	value := "["
	for cnt != 0 && *i < len(str) {
		switch str[*i] {
		case '[':
			cnt++
			value = value + "["
		case ']':
			cnt--
			value = value + "]"
		case '\n':
		case '\r':
			*lineNo++
			value = value + " "
		default:
			value = value + string(str[*i])
		}
	}

	if cnt != 0 {
		return Token{
			tokenType: NULL,
			value:     value,
			line:      *lineNo,
		}
	} else {
		return Token{
			tokenType: LINK_TEXT,
			value:     value,
			line:      *lineNo,
		}
	}
}
