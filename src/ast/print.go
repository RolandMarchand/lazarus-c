package ast

import (
	"fmt"
	"github.com/xlab/treeprint"
	"lazarus-c/src/lexer"
	"reflect"
)

func init() {
	treeprint.EdgeTypeLink = "|"
	treeprint.EdgeTypeMid = "|-"
	treeprint.EdgeTypeEnd = "`-"
}

func format(n node, tree treeprint.Tree) treeprint.Tree {
	var nodeVal = reflect.ValueOf(n).Elem()
	var nodeType = reflect.TypeOf(n).Elem()

	var fields = reflect.VisibleFields(nodeType)
	for idx := 0; idx < nodeVal.NumField(); idx++ {
		if fields[idx].Type == reflect.TypeOf(lexer.Position{}) {
			continue
		}
		if nodeVal.Field(idx).IsNil() {
			continue
		}
		if fields[idx].Type == reflect.TypeOf(true) {
			tree.AddMetaNode(fields[idx].Name, nodeVal.Field(idx).String())
		} else if fields[idx].Type.Elem() == reflect.TypeOf("") {
			var lexeme = fmt.Sprintf("\"%s\"", nodeVal.Field(idx).Elem().String())
			lexeme = fmt.Sprintf("%s: %s", fields[idx].Name, lexeme)
			tree.AddNode(lexeme)
		} else if nodeVal.Field(idx).Kind() == reflect.Slice {
			var fieldVal = nodeVal.Field(idx)
			for _idx := 0; _idx < fieldVal.Len(); _idx++ {
				var elemVal = fieldVal.Index(_idx)
				if elemVal.Elem().Type() == reflect.TypeOf("") {
					var lexeme = fmt.Sprintf("\"%s\"", elemVal.Elem().String())
					lexeme = fmt.Sprintf("%s[%d]: %s", fields[idx].Name, _idx, lexeme)
					tree.AddNode(lexeme)
				} else {
					var n = elemVal.Interface().(node)
					var branch = tree.AddMetaBranch(n.getPos(), elemVal.Elem().Type().Name())
					format(n, branch)
				}
			}
		} else {
			var n = nodeVal.Field(idx).Interface().(node)
			var branch = tree.AddMetaBranch(n.getPos(), nodeVal.Field(idx).Elem().Type().Name())
			format(nodeVal.Field(idx).Interface().(node), branch)
		}
	}

	return tree
}

func (n *TranslationUnit) String() string {
	var tree = treeprint.NewWithRoot("TranslationUnit")
	tree = format(n, tree)
	return tree.String()
}

func (n *ExternalDeclaration) String() string {
	var tree = treeprint.NewWithRoot("ExternalDeclaration")
	tree = format(n, tree)
	return tree.String()
}

func (n *FunctionDefinition) String() string {
	var tree = treeprint.NewWithRoot("FunctionDefinition")
	tree = format(n, tree)
	return tree.String()
}

func (n *CompoundStatement) String() string {
	var tree = treeprint.NewWithRoot("CompoundStatement")
	tree = format(n, tree)
	return tree.String()
}

func (n *StatementList) String() string {
	var tree = treeprint.NewWithRoot("StatementList")
	tree = format(n, tree)
	return tree.String()
}

func (n *Statement) String() string {
	var tree = treeprint.NewWithRoot("Statement")
	tree = format(n, tree)
	return tree.String()
}

func (n *LabeledStatement) String() string {
	var tree = treeprint.NewWithRoot("LabeledStatement")
	tree = format(n, tree)
	return tree.String()
}

func (n *ExpressionStatement) String() string {
	var tree = treeprint.NewWithRoot("ExpressionStatement")
	tree = format(n, tree)
	return tree.String()
}

func (n *SelectionStatement) String() string {
	var tree = treeprint.NewWithRoot("SelectionStatement")
	tree = format(n, tree)
	return tree.String()
}

func (n *IterationStatement) String() string {
	var tree = treeprint.NewWithRoot("IterationStatement")
	tree = format(n, tree)
	return tree.String()
}

func (n *JumpStatement) String() string {
	var tree = treeprint.NewWithRoot("JumpStatement")
	tree = format(n, tree)
	return tree.String()
}

func (n *DeclarationSpecifiers) String() string {
	var tree = treeprint.NewWithRoot("DeclarationSpecifiers")
	tree = format(n, tree)
	return tree.String()
}

func (n *TypeSpecifier) String() string {
	var tree = treeprint.NewWithRoot("TypeSpecifier")
	tree = format(n, tree)
	return tree.String()
}

func (n *StructOrUnionSpecifier) String() string {
	var tree = treeprint.NewWithRoot("StructOrUnionSpecifier")
	tree = format(n, tree)
	return tree.String()
}

func (n *StructDeclarationList) String() string {
	var tree = treeprint.NewWithRoot("StructDeclarationList")
	tree = format(n, tree)
	return tree.String()
}

func (n *StructDeclaration) String() string {
	var tree = treeprint.NewWithRoot("StructDeclaration")
	tree = format(n, tree)
	return tree.String()
}

func (n *SpecifierQualifierList) String() string {
	var tree = treeprint.NewWithRoot("SpecifierQualifierList")
	tree = format(n, tree)
	return tree.String()
}

func (n *StructDeclaratorList) String() string {
	var tree = treeprint.NewWithRoot("StructDeclaratorList")
	tree = format(n, tree)
	return tree.String()
}

func (n *StructDeclarator) String() string {
	var tree = treeprint.NewWithRoot("StructDeclarator")
	tree = format(n, tree)
	return tree.String()
}

func (n *EnumSpecifier) String() string {
	var tree = treeprint.NewWithRoot("EnumSpecifier")
	tree = format(n, tree)
	return tree.String()
}

func (n *EnumeratorList) String() string {
	var tree = treeprint.NewWithRoot("EnumeratorList")
	tree = format(n, tree)
	return tree.String()
}

func (n *Enumerator) String() string {
	var tree = treeprint.NewWithRoot("Enumerator")
	tree = format(n, tree)
	return tree.String()
}

func (n *DeclarationList) String() string {
	var tree = treeprint.NewWithRoot("DeclarationList")
	tree = format(n, tree)
	return tree.String()
}

func (n *Declaration) String() string {
	var tree = treeprint.NewWithRoot("Declaration")
	tree = format(n, tree)
	return tree.String()
}

func (n *InitDeclaratorList) String() string {
	var tree = treeprint.NewWithRoot("InitDeclaratorList")
	tree = format(n, tree)
	return tree.String()
}

func (n *InitDeclarator) String() string {
	var tree = treeprint.NewWithRoot("InitDeclarator")
	tree = format(n, tree)
	return tree.String()
}

func (n *Initializer) String() string {
	var tree = treeprint.NewWithRoot("Initializer")
	tree = format(n, tree)
	return tree.String()
}

func (n *InitializerList) String() string {
	var tree = treeprint.NewWithRoot("InitializerList")
	tree = format(n, tree)
	return tree.String()
}

func (n *Declarator) String() string {
	var tree = treeprint.NewWithRoot("Declarator")
	tree = format(n, tree)
	return tree.String()
}

func (n *Pointer) String() string {
	var tree = treeprint.NewWithRoot("Pointer")
	tree = format(n, tree)
	return tree.String()
}

func (n *TypeQualifierList) String() string {
	var tree = treeprint.NewWithRoot("TypeQualifierList")
	tree = format(n, tree)
	return tree.String()
}

func (n *TypeQualifier) String() string {
	var tree = treeprint.NewWithRoot("TypeQualifier")
	tree = format(n, tree)
	return tree.String()
}

func (n *DirectDeclarator) String() string {
	var tree = treeprint.NewWithRoot("DirectDeclarator")
	tree = format(n, tree)
	return tree.String()
}

func (n *IdentifierList) String() string {
	var tree = treeprint.NewWithRoot("IdentifierList")
	tree = format(n, tree)
	return tree.String()
}

func (n *ParameterTypeList) String() string {
	var tree = treeprint.NewWithRoot("ParameterTypeList")
	tree = format(n, tree)
	return tree.String()
}

func (n *ParameterList) String() string {
	var tree = treeprint.NewWithRoot("ParameterList")
	tree = format(n, tree)
	return tree.String()
}

func (n *ParameterDeclaration) String() string {
	var tree = treeprint.NewWithRoot("ParameterDeclaration")
	tree = format(n, tree)
	return tree.String()
}

func (n *AbstractDeclarator) String() string {
	var tree = treeprint.NewWithRoot("AbstractDeclarator")
	tree = format(n, tree)
	return tree.String()
}

func (n *DirectAbstractDeclarator) String() string {
	var tree = treeprint.NewWithRoot("DirectAbstractDeclarator")
	tree = format(n, tree)
	return tree.String()
}

func (n *ConstantExpression) String() string {
	var tree = treeprint.NewWithRoot("ConstantExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *ConditionalExpression) String() string {
	var tree = treeprint.NewWithRoot("ConditionalExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *LogicalOrExpression) String() string {
	var tree = treeprint.NewWithRoot("LogicalOrExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *LogicalAndExpression) String() string {
	var tree = treeprint.NewWithRoot("LogicalAndExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *InclusiveOrExpression) String() string {
	var tree = treeprint.NewWithRoot("InclusiveOrExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *ExclusiveOrExpression) String() string {
	var tree = treeprint.NewWithRoot("ExclusiveOrExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *AndExpression) String() string {
	var tree = treeprint.NewWithRoot("AndExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *EqualityExpression) String() string {
	var tree = treeprint.NewWithRoot("EqualityExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *RelationalExpression) String() string {
	var tree = treeprint.NewWithRoot("RelationalExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *ShiftExpression) String() string {
	var tree = treeprint.NewWithRoot("ShiftExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *AdditiveExpression) String() string {
	var tree = treeprint.NewWithRoot("AdditiveExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *MultiplicativeExpression) String() string {
	var tree = treeprint.NewWithRoot("MultiplicativeExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *CastExpression) String() string {
	var tree = treeprint.NewWithRoot("CastExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *UnaryExpression) String() string {
	var tree = treeprint.NewWithRoot("UnaryExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *UnaryOperator) String() string {
	var tree = treeprint.NewWithRoot("UnaryOperator")
	tree = format(n, tree)
	return tree.String()
}

func (n *TypeName) String() string {
	var tree = treeprint.NewWithRoot("TypeName")
	tree = format(n, tree)
	return tree.String()
}

func (n *PostfixExpression) String() string {
	var tree = treeprint.NewWithRoot("PostfixExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *ArgumentExpressionList) String() string {
	var tree = treeprint.NewWithRoot("ArgumentExpressionList")
	tree = format(n, tree)
	return tree.String()
}

func (n *PrimaryExpression) String() string {
	var tree = treeprint.NewWithRoot("PrimaryExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *Expression) String() string {
	var tree = treeprint.NewWithRoot("Expression")
	tree = format(n, tree)
	return tree.String()
}

func (n *AssignmentExpression) String() string {
	var tree = treeprint.NewWithRoot("AssignmentExpression")
	tree = format(n, tree)
	return tree.String()
}

func (n *AssignmentOperator) String() string {
	var tree = treeprint.NewWithRoot("AssignmentOperator")
	tree = format(n, tree)
	return tree.String()
}
