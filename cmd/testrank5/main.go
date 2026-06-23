package main

import (
	"fmt"
	tsqlparser "github.com/haimiyahya/tsqlparser"
)

func main() {
	fmt.Println("=== Debug test ===")
	sql := `WITH cte AS (SELECT ROW_NUMBER() OVER (ORDER BY id DESC) AS Rank FROM table1) SELECT * FROM cte`
	stmt, errs := tsqlparser.Parse(sql)
	if len(errs) > 0 {
		fmt.Println("Errors:")
		for _, e := range errs {
			fmt.Println("  ", e)
		}
	} else {
		fmt.Println("Success!")
		fmt.Printf("Statement type: %T\n", stmt)
	}
}
