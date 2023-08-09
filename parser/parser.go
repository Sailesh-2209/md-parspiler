package parser

import (
	"md_parspiler/ast"
	"md_parspiler/tokens"
)

func Parse(tokensList []tokens.Token) []ast.ASTNode {
	var astList []ast.ASTNode

	numTokens := len(tokensList)

	for i := 0; i < numTokens; {
		currToken := tokensList[i]

		var node ast.ASTNode

		switch currToken.TokenType {
		case tokens.HASH1, tokens.HASH2, tokens.HASH3, tokens.HASH4, tokens.HASH5, tokens.HASH6:
			i++
			words := consumeLine(&i, tokensList)
			node = ast.Heading{
				Level:   len(currToken.Value),
				Heading: words,
			}
		default:
			i++
		}

		astList = append(astList, node)
	}

	return astList
}

func consumeLine(i *int, tokensList []tokens.Token) []ast.Word {
	wordsList := []ast.Word{}

	stop := false

	for *i < len(tokensList) && !stop {
		currToken := tokensList[*i]

		switch currToken.TokenType {
		case tokens.NEW_LINE:
			stop = true
			*i++
		case tokens.SPACE:
			// spaces are ignored. The final words list is joined
			// together by single space characters
			*i++
		case tokens.UNDERSCORE, tokens.STAR1, tokens.STAR2, tokens.STAR3, tokens.TICK:
			words := consumeStylizedWords(i, tokensList)
			wordsList = append(wordsList, words...)
		default:
			word := ast.Word{
				WordType: ast.WORD_REGULAR,
				Word:     currToken.Value,
			}
			wordsList = append(wordsList, word)
			*i++
		}
	}

	return wordsList
}

func consumeStylizedWords(i *int, tokensList []tokens.Token) []ast.Word {
	match := tokensList[*i]
	wordsList := []ast.Word{
		{
			WordType: ast.WORD_REGULAR,
			Word:     tokensList[*i].Value,
		},
	}
	*i++

	for ; *i < len(tokensList) && tokensList[*i].TokenType != tokens.NEW_LINE; *i++ {
		if tokensList[*i].TokenType == match.TokenType && tokensList[*i-1].TokenType != tokens.SPACE {
			var wordType ast.WordType

			switch match.TokenType {
			case tokens.UNDERSCORE, tokens.STAR1:
				wordType = ast.WORD_ITALICS
			case tokens.STAR2:
				wordType = ast.WORD_BOLD
			case tokens.STAR3:
				wordType = ast.WORD_BOLD_ITALICS
			case tokens.TICK:
				wordType = ast.WORD_CODE
			}

			wordsList = wordsList[1:]
			for j := 0; j < len(wordsList); j++ {
				wordsList[j].WordType = wordType
			}
			*i++
			return wordsList
		}
		word := ast.Word{
			WordType: ast.WORD_REGULAR,
			Word:     tokensList[*i].Value,
		}
		wordsList = append(wordsList, word)
	}

	return wordsList
}
