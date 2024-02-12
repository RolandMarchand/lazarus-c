package main

// import (
// 	"fmt"
// 	"github.com/alecthomas/participle/v2"
// 	"github.com/alecthomas/participle/v2/lexer"
// )

// type Lmao struct {
// 	Words []Word `parser:"@@*"`
// }
// type Word struct {
// 	Pos lexer.Position
// 	W string `parser:"@Ident"`
// }

// var program, err = participle.MustBuild[Lmao]().ParseString("source:", `auto break case char const continue default do double else enum extern float for goto if int long register return short signed sizeof static struct switch typedef union unsigned void volatile while`)

// func main() {
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, w := range program.Words {
// 		fmt.Println(w)
// 	}
// }

import (
	"fmt"
	"lazarus-c/src/ast"
)

func main() {
	var myAst, err = ast.ParseString(`
int main()
{
	int i = 10;
}
`)
	if err != nil {
		panic(err)
	}
	fmt.Println(myAst.ExternalDeclarations[0])
}
