package main

import (
	"fmt"
	"github.com/haimiyahya/tsqlparser"
)

func main() {
	sql := "IF (@sz_CardPAN NOT LIKE '308383%' AND @sz_CardPAN NOT LIKE '958610001%') SELECT 1"
	program, errors := tsqlparser.Parse(sql)
	if len(errors) > 0 {
		fmt.Printf("Parse errors: %v\n", errors)
		return
	}
	fmt.Println("AST String output:")
	fmt.Println(program.String())
}
