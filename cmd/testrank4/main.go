package main

import (
	"fmt"
	tsqlparser "github.com/haimiyahya/tsqlparser"
)

func test1() {
	fmt.Println("=== Test: Without CTE ===")
	sql := `SELECT *, ROW_NUMBER() OVER (ORDER BY id DESC) AS Rank FROM table1`
	_, errs := tsqlparser.Parse(sql)
	if len(errs) > 0 {
		fmt.Println("Errors:")
		for _, e := range errs {
			fmt.Println("  ", e)
		}
	} else {
		fmt.Println("Success!")
	}
}

func test2() {
	fmt.Println("\n=== Test: Without * ===")
	sql := `WITH cte AS (SELECT ROW_NUMBER() OVER (ORDER BY id DESC) AS Rank FROM table1) SELECT * FROM cte`
	_, errs := tsqlparser.Parse(sql)
	if len(errs) > 0 {
		fmt.Println("Errors:")
		for _, e := range errs {
			fmt.Println("  ", e)
		}
	} else {
		fmt.Println("Success!")
	}
}

func test3() {
	fmt.Println("\n=== Test: Without OVER ===")
	sql := `WITH cte AS (SELECT *, dbo.fn() AS Rank FROM table1) SELECT * FROM cte`
	_, errs := tsqlparser.Parse(sql)
	if len(errs) > 0 {
		fmt.Println("Errors:")
		for _, e := range errs {
			fmt.Println("  ", e)
		}
	} else {
		fmt.Println("Success!")
	}
}

func main() {
	test1()
	test2()
	test3()
}
