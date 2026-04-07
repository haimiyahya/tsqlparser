package main

import (
	"fmt"
	"github.com/haimiyahya/tsqlparser"
)

func main() {
	// Simple test: just a NOT LIKE
	sql := "SELECT * FROM t WHERE a NOT LIKE 'test%'"
	program, errors := tsqlparser.Parse(sql)
	if len(errors) > 0 {
		fmt.Printf("Parse errors: %v\n", errors)
		return
	}
	fmt.Println("Simple NOT LIKE:")
	fmt.Println(program.String())
	
	// Now test with AND
	sql2 := "SELECT * FROM t WHERE a NOT LIKE 'test%' AND a NOT LIKE 'other%'"
	program2, errors2 := tsqlparser.Parse(sql2)
	if len(errors2) > 0 {
		fmt.Printf("Parse errors: %v\n", errors2)
		return
	}
	fmt.Println("\nNOT LIKE with AND:")
	fmt.Println(program2.String())
}
