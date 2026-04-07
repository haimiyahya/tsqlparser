package main

import (
	"fmt"
	"github.com/haimiyahya/tsqlparser"
)

func main() {
	sql := "SELECT * FROM t WHERE a NOT LIKE 'test%' AND b > 5"
	program, errors := tsqlparser.Parse(sql)
	if len(errors) > 0 {
		fmt.Printf("Parse errors: %v\n", errors)
		return
	}
	fmt.Println("AST:")
	fmt.Println(program.String())
}
