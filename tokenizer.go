package preprocessor

import (
	"regexp"
)

func minify(input string) string {
	n := regexp.MustCompile(`\r?\n`)
	return n.ReplaceAllString(input, "")
}

type TokenSpec struct {
	regExp     string
	matchIndex int
	tokenType  TokenType
}

type Token struct {
	content   string
	tokenType TokenType
}

type Tokenizer struct {
	input  string
	cursor int
	spec   []TokenSpec
}

func (tokenizer *Tokenizer) NextToken() *Token {
	nextInput := tokenizer.input[tokenizer.cursor:]

	for _, tokenSpec := range tokenizer.spec {
		matched := regexp.MustCompile(tokenSpec.regExp).FindStringSubmatch(nextInput)

		if matched == nil {
			continue
		}

		tokenizer.cursor += len(matched[tokenSpec.matchIndex])

		result := minify(matched[tokenSpec.matchIndex])

		return &Token{result, tokenSpec.tokenType}
	}

	return nil
}
