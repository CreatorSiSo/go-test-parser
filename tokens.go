package preprocessor

type TokenType int

//go:generate stringer -type=TokenType
const (
	CONTENT TokenType = iota + 1
	EXPRESSION
	COMMENT
)
