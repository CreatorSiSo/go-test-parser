package preprocessor

import (
	"fmt"
	"regexp"
)

func minify(input string) string {
	n := regexp.MustCompile(`\r?\n`)
	return n.ReplaceAllString(input, "")
}

// Specification for matching a specific Token
type TokenSpec struct {
	regExp     string
	matchIndex int
	tokenType  TokenType
}

type Token struct {
	Value string    `json:"value,omitempty"`
	Type  TokenType `json:"type,omitempty"`
}

func (token Token) String() string {
	return fmt.Sprintf("%s: \"%s\"", token.Type.String(), token.Value)
}

// TODO: Use Unmarshal; encoding/json package
// TODO: Write comment
func (token Token) Json() string {
	return ""
}

type Tokenizer struct {
	Spec   []TokenSpec
	input  string
	cursor int
}

// TODO: Write comment
func (tokenizer *Tokenizer) NextToken() *Token {
	nextInput := tokenizer.input[tokenizer.cursor:]

	for _, tokenSpec := range tokenizer.Spec {
		matched := regexp.MustCompile(tokenSpec.regExp).FindStringSubmatch(nextInput)
		if matched == nil {
			continue
		}

		returnValue := matched[tokenSpec.matchIndex]
		tokenizer.cursor += len(returnValue)
		return &Token{minify(returnValue), tokenSpec.tokenType}
	}

	return nil
}

func (tokenizer *Tokenizer) SetInput(input string) {
	tokenizer.input = input
	tokenizer.cursor = 0
}

// TODO: Write comment
func (tokenizer Tokenizer) Tokenize(input string) []Token {
	tokenizer.SetInput(input)

	var tokens []Token

	for tokenizer.cursor < len(input) {
		nextToken := tokenizer.NextToken()

		if nextToken == nil {
			break
		}

		tokens = append(tokens, *nextToken)
	}

	return tokens
}
