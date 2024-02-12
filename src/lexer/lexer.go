package lexer

import "github.com/alecthomas/participle/v2/lexer"

type Position lexer.Position
func (p Position) String() string {
	return lexer.Position(p).String()
}

var Lexer = lexer.MustSimple([]lexer.SimpleRule{
	{Name: "Whitespace", Pattern: `\s+`},
	{Name: "Comment", Pattern: `//[^\n]*`},

	{Name: "Char", Pattern: `'(\\?\p{Any})+'`},
	{Name: "String", Pattern: `"(\\?\p{Any})*"`},
	{Name: "Ident", Pattern: `[\p{L}_][\p{L}\p{N}_]*`},
	{Name: "Float", Pattern: `[0-9]+\.[0-9]+([eE][+-]?[0-9]+)?|[0-9]+(\.[0-9]+)?[eE][+-]?[0-9]+`},
	{Name: "Int", Pattern: `0[xX][0-9a-fA-F]+|0[oO][0-7]+|0[bB][01]+|[0-9]+`},

	{Name: "ThreeOp", Pattern: `\.\.\.|<<=|>>=`},
	{Name: "TwoOp", Pattern: `((=|!|\+|-|\*|/|%|\||&|<|>|\^)=)|<<|>>|\+\+|--|->|&&|\|\|`},
	{Name: "OneOp", Pattern: `;|{|}|,|:|=|\(|\)|\[|\]|\.|&|!|~|-|\+|\*|/|%|<|>|\^|\||\?`},
})





