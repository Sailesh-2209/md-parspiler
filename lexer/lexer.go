package lexer

import (
	"fmt"
	"md_parspiler/tokens"
	"regexp"
)

func Lex(str string) []tokens.Token {
	fmt.Printf("Source before cleanup:\n%s", str)
	str = cleanUp(str)
	fmt.Printf("Source after cleanup:\n%s", str)

	return []tokens.Token{}
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
	duplicateBoldStars := regexp.MustCompile(`\*\*\*\*+((\S.*\S)|(\S))\*\*\*\*+`)
	str = duplicateBoldStars.ReplaceAllString(str, "**$1**")

	return str
}
