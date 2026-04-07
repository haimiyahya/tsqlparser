package main

import (
	"fmt"
	"github.com/haimiyahya/tsqlparser/lexer"
)

func main() {
	sql := "SELECT * FROM t WHERE a NOT LIKE 'test%' AND a NOT LIKE 'other%'"
	l := lexer.New(sql)
	for tok := l.NextToken(); tok.Type != 0; tok = l.NextToken() {
		fmt.Printf("Token: %-15s Literal: %q\n", tok.Type, tok.Literal)
	}
}
