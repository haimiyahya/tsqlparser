package main

import (
	"fmt"
	"github.com/haimiyahya/tsqlparser"
	"github.com/haimiyahya/tsqlparser/ast"
)

func printAST(node ast.Node, indent int) {
	prefix := ""
	for i := 0; i < indent; i++ {
		prefix += "  "
	}
	
	switch n := node.(type) {
	case *ast.IfStatement:
		fmt.Printf("%sIfStatement\n", prefix)
		fmt.Printf("%s  Condition:\n", prefix)
		printAST(n.Condition, indent+2)
		fmt.Printf("%s  Consequence:\n", prefix)
		printAST(n.Consequence, indent+2)
	case *ast.InfixExpression:
		fmt.Printf("%sInfixExpression: Operator='%s'\n", prefix, n.Operator)
		fmt.Printf("%s  Left:\n", prefix)
		printAST(n.Left, indent+2)
		fmt.Printf("%s  Right:\n", prefix)
		printAST(n.Right, indent+2)
	case *ast.PrefixExpression:
		fmt.Printf("%sPrefixExpression: Operator='%s'\n", prefix, n.Operator)
		fmt.Printf("%s  Right:\n", prefix)
		printAST(n.Right, indent+2)
	case *ast.LikeExpression:
		fmt.Printf("%sLikeExpression: Not=%v\n", prefix, n.Not)
		fmt.Printf("%s  Expr:\n", prefix)
		printAST(n.Expr, indent+2)
		fmt.Printf("%s  Pattern:\n", prefix)
		printAST(n.Pattern, indent+2)
	case *ast.Identifier:
		fmt.Printf("%sIdentifier: %s\n", prefix, n.Value)
	case *ast.StringLiteral:
		fmt.Printf("%sStringLiteral: %s\n", prefix, n.Value)
	default:
		fmt.Printf("%s%T\n", prefix, node)
	}
}

func main() {
	sql := "IF (@sz_CardPAN NOT LIKE '308383%' AND @sz_CardPAN NOT LIKE '958610001%') SELECT 1"
	program, errors := tsqlparser.Parse(sql)
	if len(errors) > 0 {
		fmt.Printf("Parse errors: %v\n", errors)
		return
	}
	
	for _, stmt := range program.Statements {
		printAST(stmt, 0)
	}
}
