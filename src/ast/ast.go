// https://www.lysator.liu.se/c/ANSI-C-grammar-y.html#postfix-expression
package ast

import (
	"github.com/alecthomas/participle/v2"
	"io"
	"lazarus-c/src/lexer"
)

type node interface {
	getPos() lexer.Position
}

var Parser = participle.MustBuild[TranslationUnit](
	participle.UseLookahead(1),
	participle.Unquote("String", "Char"),
	participle.Lexer(lexer.Lexer),
	participle.Elide("Whitespace", "Comment"),
)

func ParseString(s string) (*TranslationUnit, error) {
	program, err := Parser.ParseString("", s)
	if err != nil {
		return nil, err
	}
	return program, nil
}

func Parse(r io.Reader) (*TranslationUnit, error) {
	program, err := Parser.Parse("", r)
	if err != nil {
		return nil, err
	}
	return program, nil
}

type TranslationUnit struct {
	Pos                  lexer.Position
	ExternalDeclarations []*ExternalDeclaration `parser:"@@+"`
}

type ExternalDeclaration struct {
	Pos                lexer.Position
	FunctionDefinition *FunctionDefinition `parser:"@@"`
	Declaration        *Declaration        `parser:"| @@"`
}

type FunctionDefinition struct {
	Pos                   lexer.Position
	DeclarationSpecifiers *DeclarationSpecifiers `parser:"@@?"`
	Declarator            *Declarator            `parser:"@@"`
	DeclarationList       *DeclarationList       `parser:"@@?"`
	CompoundStatement     *CompoundStatement     `parser:"@@"`
}

type CompoundStatement struct {
	Pos             lexer.Position
	DeclarationList *DeclarationList `parser:"'{' @@?"`
	StatementList   *StatementList   `parser:"@@? '}'"`
}

type StatementList struct {
	Pos        lexer.Position
	Statements []*Statement `parser:"@@+"`
}

type Statement struct {
	Pos                 lexer.Position
	LabeledStatement    *LabeledStatement    `parser:"@@"`
	CompoundStatement   *CompoundStatement   `parser:"| @@"`
	ExpressionStatement *ExpressionStatement `parser:"| @@"`
	SelectionStatement  *SelectionStatement  `parser:"| @@"`
	IterationStatement  *IterationStatement  `parser:"| @@"`
	JumpStatement       *JumpStatement       `parser:"| @@"`
}

type LabeledStatement struct {
	Pos              lexer.Position
	GotoLabel        *string             `parser:"@Ident"`
	GotoStatement    *Statement          `parser:"':' @@"`
	CaseExpression   *ConstantExpression `parser:"| 'case' @@"`
	CaseStatement    *Statement          `parser:"@@"`
	DefaultStatement *Statement          `parser:"| 'default' @@"`
}

type ExpressionStatement struct {
	Pos        lexer.Position
	Expression *Expression `parser:"@@? ';'"`
}

type SelectionStatement struct {
	Pos              lexer.Position
	IfTest           *Expression `parser:"'if' '(' @@ ')'"`
	IfBody           *Statement  `parser:"@@"`
	ElseBody         *Statement  `parser:"( 'else' @@ )?"`
	SwitchExpression *Expression `parser:"| 'switch' '(' @@ ')'"`
	SwitchBody       *Statement  `parser:"@@"`
}

type IterationStatement struct {
	Pos       lexer.Position
	WhileTest *Expression          `parser:"'while' '(' @@ ')'"`
	WhileBody *Statement           `parser:"@@"`
	DoBody    *Statement           `parser:"| 'do' @@"`
	DoTest    *Expression          `parser:"'while' '(' @@ ')' ';'"`
	ForInit   *ExpressionStatement `parser:"| 'for' '(' @@"`
	ForTest   *ExpressionStatement `parser:"@@"`
	ForUpdate *Expression          `parser:"@@? ')'"`
	ForBody   *Statement           `parser:"@@"`
}

type JumpStatement struct {
	Pos              lexer.Position
	GotoIdent        *string     `parser:"'goto' @Ident ';'"`
	IsContinue       bool        `parser:"| @'continue' ';'"`
	IsBreak          bool        `parser:"| @'break' ';'"`
	IsReturn         bool        `parser:"| @'return'"`
	ReturnExpression *Expression `parser:"@@? ';'"`
}

type DeclarationSpecifiers struct {
	Pos lexer.Position
	// Typedef should be implemented inside of StorageClassSpecifier in the future
	StorageClassSpecifier *string                `parser:"( ( @'extern' | @'static' | @'auto' | @'register' )"`
	TypeSpecifier         *TypeSpecifier         `parser:"| @@"`
	TypeQualifier         *TypeQualifier         `parser:"| @@ )"`
	DeclarationSpecifiers *DeclarationSpecifiers `parser:"@@?"`
}

type TypeSpecifier struct {
	Pos lexer.Position
	// Custom types created with typedef should be added here.
	TypeSpecifier          *string                 `parser:"( @'void' | @'char' | @'short' | @'int' | @'long' | @'float' | @'double' | @'signed' | @'unsigned' )"`
	StructOrUnionSpecifier *StructOrUnionSpecifier `parser:"| @@"`
	EnumSpecifier          *EnumSpecifier          `parser:"| @@"`
}

type StructOrUnionSpecifier struct {
	Pos                   lexer.Position
	StructOrUnion         *string                `parser:"@'struct' | @'union'"`
	Identifier            *string                `parser:"( ( @Ident"`
	StructDeclarationList *StructDeclarationList `parser:"( '{' @@ '}' )? ) | '{' @@ '}' )"`
}

type StructDeclarationList struct {
	Pos                lexer.Position
	StructDeclarations []*StructDeclaration `parser:"@@+"`
}

type StructDeclaration struct {
	Pos                    lexer.Position
	SpecifierQualifierList *SpecifierQualifierList `parser:"@@"`
	StructDeclaratorList   *StructDeclaratorList   `parser:"@@ ';'"`
}

type SpecifierQualifierList struct {
	Pos           lexer.Position
	TypeSpecifier *TypeSpecifier `parser:"( @@"`
	TypeQualifier *TypeQualifier `parser:"| @@ )+"`
}

type StructDeclaratorList struct {
	Pos               lexer.Position
	StructDeclarators []*StructDeclarator `parser:"@@ ( ',' @@ )*"`
}

type StructDeclarator struct {
	Pos                lexer.Position
	Declarator         *Declarator         `parser:"@@"`
	ConstantExpression *ConstantExpression `parser:"( ':' @@ )? | @@"`
}

type EnumSpecifier struct {
	Pos            lexer.Position
	Identifier     *string         `parser:"'enum' ( ( @Ident"`
	EnumeratorList *EnumeratorList `parser:"( '{' @@ '}' )? ) | '{' @@ '}' )"`
}

type EnumeratorList struct {
	Pos         lexer.Position
	Enumerators []*Enumerator `parser:"@@ ( ',' @@ )*"`
}

type Enumerator struct {
	Pos                lexer.Position
	Identifier         *string             `parser:"@Ident"`
	ConstantExpression *ConstantExpression `parser:"( '=' @@ )?"`
}

type DeclarationList struct {
	Pos          lexer.Position
	Declarations []*Declaration `parser:"@@+"`
}

type Declaration struct {
	Pos                   lexer.Position
	DeclarationSpecifiers *DeclarationSpecifiers `parser:"@@"`
	InitDeclaratorList    *InitDeclaratorList    `parser:"@@? ';'"`
}

type InitDeclaratorList struct {
	Pos             lexer.Position
	InitDeclarators []*InitDeclarator `parser:"@@ ( ',' @@ )*"`
}

type InitDeclarator struct {
	Pos         lexer.Position
	Declarator  *Declarator  `parser:"@@"`
	Initializer *Initializer `parser:"( '=' @@ )?"`
}

type Initializer struct {
	Pos                  lexer.Position
	AssignmentExpression *AssignmentExpression `parser:"@@"`
	InitializerList      *InitializerList      `parser:"| '{' @@ ','? '}'"`
}

type InitializerList struct {
	Pos          lexer.Position
	Initializers []*Initializer `parser:"@@ ( ',' @@ )*"`
}

type Declarator struct {
	Pos               lexer.Position
	Pointer           *Pointer            `parser:"@@?"`
	DirectDeclarators []*DirectDeclarator `parser:"@@"`
}

type Pointer struct {
	Pos               lexer.Position
	TypeQualifierList *TypeQualifierList `parser:"'*' ( @@"`
	Pointer           *Pointer           `parser:"@@? | @@ )"`
}

type TypeQualifierList struct {
	Pos            lexer.Position
	TypeQualifiers []*TypeQualifier `parser:"@@+"`
}

type TypeQualifier struct {
	Pos       lexer.Position
	Qualifier *string `parser:"@'const' | @'volatile'"`
}

type DirectDeclarator struct {
	Pos                lexer.Position
	Identifier         *string               `parser:"( @Ident"`
	Declarator         *Declarator           `parser:"| '(' @@ ')' )"`
	ArrayLength        []*ConstantExpression `parser:"( '[' @@? ']'"`
	ParameterTypeLists []*ParameterTypeList  `parser:"| '(' ( @@"`
	IdentifierLists    []*IdentifierList     `parser:"| @@ )? ')' )*"`
}

type IdentifierList struct {
	Pos         lexer.Position
	Identifiers []*string `parser:"@Ident ( ',' @Ident )*"`
}

type ParameterTypeList struct {
	Pos           lexer.Position
	ParameterList *ParameterList `parser:"@@"`
	Ellipsis      bool           `parser:"( ',' @'...' )?"`
}

type ParameterList struct {
	Pos                   lexer.Position
	ParameterDeclarations []*ParameterDeclaration `parser:"@@ ( ',' @@ )*"`
}

type ParameterDeclaration struct {
	Pos                   lexer.Position
	DeclarationSpecifiers *DeclarationSpecifiers `parser:"@@"`
	Declarator            *Declarator            `parser:"( @@"`
	AbstractDeclarator    *AbstractDeclarator    `parser:"| @@ )?"`
}

type AbstractDeclarator struct {
	Pos                      lexer.Position
	Pointer                  *Pointer                  `parser:"@@"`
	DirectAbstractDeclarator *DirectAbstractDeclarator `parser:"@@? | @@"`
}

type DirectAbstractDeclarator struct {
	Pos                      lexer.Position
	AbstractDeclarator       *AbstractDeclarator       `parser:"'(' @@ ')'"`
	DirectAbstractDeclarator *DirectAbstractDeclarator `parser:"| @@?"`
	ConstantExpression       *ConstantExpression       `parser:"( '[' @@? ']'"`
	ParameterTypeList        *ParameterTypeList        `parser:"| '(' @@? ')' )"`
}

type ConstantExpression struct {
	Pos                   lexer.Position
	ConditionalExpression *ConditionalExpression `parser:"@@"`
}

type ConditionalExpression struct {
	Pos                    lexer.Position
	LogicalOrExpression    *LogicalOrExpression   `parser:"@@"`
	TernaryTrueExpression  *Expression            `parser:"( '?' @@"`
	TernaryFalseExpression *ConditionalExpression `parser:"':' @@ )?"`
}

type LogicalOrExpression struct {
	Pos                   lexer.Position
	LogicalAndExpressions []*LogicalAndExpression `parser:"@@ ( '||' @@ )*"`
}

type LogicalAndExpression struct {
	Pos                    lexer.Position
	InclusiveOrExpressions []*InclusiveOrExpression `parser:"@@ ( '&&' @@ )*"`
}

type InclusiveOrExpression struct {
	Pos                    lexer.Position
	ExclusiveOrExpressions []*ExclusiveOrExpression `parser:"@@ ( '|' @@ )*"`
}

type ExclusiveOrExpression struct {
	Pos            lexer.Position
	AndExpressions []*AndExpression `parser:"@@ ( '^' @@ )*"`
}

type AndExpression struct {
	Pos                 lexer.Position
	EqualityExpressions []*EqualityExpression `parser:"@@ ( '&' @@ )*"`
}

type EqualityExpression struct {
	Pos                       lexer.Position
	HeadRelationalExpression  *RelationalExpression   `parser:"@@"`
	Operators                 []*string               `parser:"( ( @'==' | @'!=' )"`
	TailRelationalExpressions []*RelationalExpression `parser:"@@ )*"`
}

type RelationalExpression struct {
	Pos                  lexer.Position
	HeadShiftExpression  *ShiftExpression   `parser:"@@"`
	Operators            []*string          `parser:"( ( @'<' | @'>' | @'<=' | @'>=' )"`
	TailShiftExpressions []*ShiftExpression `parser:"@@ )*"`
}

type ShiftExpression struct {
	Pos                     lexer.Position
	HeadAdditiveExpression  *AdditiveExpression   `parser:"@@"`
	Operators               []*string             `parser:"( ( @'<<' | @'>>' )"`
	TailAdditiveExpressions []*AdditiveExpression `parser:"@@ )*"`
}

type AdditiveExpression struct {
	Pos                          lexer.Position
	HeadMultiplicativeExpression *MultiplicativeExpression   `parser:"@@"`
	Operators                    []*string                   `parser:"( ( @'+' | @'-' )"`
	TailMultiplicativeExpression []*MultiplicativeExpression `parser:"@@ )*"`
}

type MultiplicativeExpression struct {
	Pos                lexer.Position
	HeadCastExpression *CastExpression   `parser:"@@"`
	Operators          []*string         `parser:"( (@'*' | @'/' | @'%' )"`
	TailCastExpression []*CastExpression `parser:"@@ )*"`
}

type CastExpression struct {
	Pos             lexer.Position
	TypeNames       []*TypeName      `parser:"( '(' @@ ')' )*"`
	UnaryExpression *UnaryExpression `parser:"@@"`
}

type UnaryExpression struct {
	Pos                 lexer.Position
	UnaryOperators      []*string          `parser:"( @'++' | @'--' | @'sizeof' )*"`
	PostfixExpression   *PostfixExpression `parser:"( @@"`
	SizeOfTypeName      *TypeName          `parser:"| 'sizeof' '(' @@ ')'"`
	UnaryOperatorOnCast *UnaryOperator     `parser:"| @@"`
	CastExpression      *CastExpression    `parser:"@@ )"`
}

type UnaryOperator struct {
	Pos      lexer.Position
	Operator *string `parser:"@'&' | @'*' | @'+' | @'-' | @'~' | @'!'"`
}

type TypeName struct {
	Pos                    lexer.Position
	SpecifierQualifierList *SpecifierQualifierList `parser:"@@"`
	AbstractDeclarator     *AbstractDeclarator     `parser:"@@?"`
}

type PostfixExpression struct {
	Pos                    lexer.Position
	PrimaryExpression      *PrimaryExpression        `parser:"@@"`
	ArrayAccessExpression  *Expression               `parser:"( '[' @@ ']'"`
	ArgumentExpressionList []*ArgumentExpressionList `parser:"| '(' @@? ')'"`
	IdentifierAccess       *string                   `parser:"| '.' @Ident"`
	IdentifierPtrAccess    *string                   `parser:"| '->' @Ident"`
	Operator               *string                   `parser:"| @'++' | @'--' )*"`
}

type ArgumentExpressionList struct {
	Pos                   lexer.Position
	AssignmentExpressions []*AssignmentExpression `parser:"@@ ( ',' @@ )*"`
}

type PrimaryExpression struct {
	Pos           lexer.Position
	Identifier    *string     `parser:"@Ident"`
	Int           *string     `parser:"| @Int"`
	Float         *string     `parser:"| @Float"`
	Char          *string     `parser:"| @Char"`
	StringLiteral *string     `parser:"| @String"`
	Expression    *Expression `parser:"| '(' @@ ')'"`
}

type Expression struct {
	Pos                   lexer.Position
	AssignmentExpressions []*AssignmentExpression `parser:"@@ ( ',' @@ )*"`
}

type AssignmentExpression struct {
	Pos                   lexer.Position
	UnaryExpressions      []*UnaryExpression     `parser:"( @@"`
	AssignmentOperators   []*AssignmentOperator  `parser:"@@ )*"`
	ConditionalExpression *ConditionalExpression `parser:"@@"`
}

type AssignmentOperator struct {
	Pos                lexer.Position
	AssignmentOperator *string `parser:"@'=' | @'+=' | @'-=' | @'*=' | @'/=' | @'%=' | @'<<=' | @'>>=' | @'|=' | @'&=' | @'^='"`
}

func (n *TranslationUnit) getPos() lexer.Position {
	return n.Pos
}

func (n *ExternalDeclaration) getPos() lexer.Position {
	return n.Pos
}

func (n *FunctionDefinition) getPos() lexer.Position {
	return n.Pos
}

func (n *CompoundStatement) getPos() lexer.Position {
	return n.Pos
}

func (n *StatementList) getPos() lexer.Position {
	return n.Pos
}

func (n *Statement) getPos() lexer.Position {
	return n.Pos
}

func (n *LabeledStatement) getPos() lexer.Position {
	return n.Pos
}

func (n *ExpressionStatement) getPos() lexer.Position {
	return n.Pos
}

func (n *SelectionStatement) getPos() lexer.Position {
	return n.Pos
}

func (n *IterationStatement) getPos() lexer.Position {
	return n.Pos
}

func (n *JumpStatement) getPos() lexer.Position {
	return n.Pos
}

func (n *DeclarationSpecifiers) getPos() lexer.Position {
	return n.Pos
}

func (n *TypeSpecifier) getPos() lexer.Position {
	return n.Pos
}

func (n *StructOrUnionSpecifier) getPos() lexer.Position {
	return n.Pos
}

func (n *StructDeclarationList) getPos() lexer.Position {
	return n.Pos
}

func (n *StructDeclaration) getPos() lexer.Position {
	return n.Pos
}

func (n *SpecifierQualifierList) getPos() lexer.Position {
	return n.Pos
}

func (n *StructDeclaratorList) getPos() lexer.Position {
	return n.Pos
}

func (n *StructDeclarator) getPos() lexer.Position {
	return n.Pos
}

func (n *EnumSpecifier) getPos() lexer.Position {
	return n.Pos
}

func (n *EnumeratorList) getPos() lexer.Position {
	return n.Pos
}

func (n *Enumerator) getPos() lexer.Position {
	return n.Pos
}

func (n *DeclarationList) getPos() lexer.Position {
	return n.Pos
}

func (n *Declaration) getPos() lexer.Position {
	return n.Pos
}

func (n *InitDeclaratorList) getPos() lexer.Position {
	return n.Pos
}

func (n *InitDeclarator) getPos() lexer.Position {
	return n.Pos
}

func (n *Initializer) getPos() lexer.Position {
	return n.Pos
}

func (n *InitializerList) getPos() lexer.Position {
	return n.Pos
}

func (n *Declarator) getPos() lexer.Position {
	return n.Pos
}

func (n *Pointer) getPos() lexer.Position {
	return n.Pos
}

func (n *TypeQualifierList) getPos() lexer.Position {
	return n.Pos
}

func (n *TypeQualifier) getPos() lexer.Position {
	return n.Pos
}

func (n *DirectDeclarator) getPos() lexer.Position {
	return n.Pos
}

func (n *IdentifierList) getPos() lexer.Position {
	return n.Pos
}

func (n *ParameterTypeList) getPos() lexer.Position {
	return n.Pos
}

func (n *ParameterList) getPos() lexer.Position {
	return n.Pos
}

func (n *ParameterDeclaration) getPos() lexer.Position {
	return n.Pos
}

func (n *AbstractDeclarator) getPos() lexer.Position {
	return n.Pos
}

func (n *DirectAbstractDeclarator) getPos() lexer.Position {
	return n.Pos
}

func (n *ConstantExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *ConditionalExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *LogicalOrExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *LogicalAndExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *InclusiveOrExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *ExclusiveOrExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *AndExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *EqualityExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *RelationalExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *ShiftExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *AdditiveExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *MultiplicativeExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *CastExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *UnaryExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *UnaryOperator) getPos() lexer.Position {
	return n.Pos
}

func (n *TypeName) getPos() lexer.Position {
	return n.Pos
}

func (n *PostfixExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *ArgumentExpressionList) getPos() lexer.Position {
	return n.Pos
}

func (n *PrimaryExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *Expression) getPos() lexer.Position {
	return n.Pos
}

func (n *AssignmentExpression) getPos() lexer.Position {
	return n.Pos
}

func (n *AssignmentOperator) getPos() lexer.Position {
	return n.Pos
}
