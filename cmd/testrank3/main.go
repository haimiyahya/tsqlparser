package main

import (
	"fmt"
	tsqlparser "github.com/haimiyahya/tsqlparser"
)

func test1() {
	fmt.Println("=== Test 1: SELECT * (no other columns) ===")
	sql := `WITH cte AS (SELECT *, ROW_NUMBER() OVER (ORDER BY id) AS Rank FROM table1) SELECT * FROM cte`
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
	fmt.Println("\n=== Test 2: SELECT with specific columns (no *) ===")
	sql := `WITH cte AS (SELECT id, ROW_NUMBER() OVER (ORDER BY id) AS Rank FROM table1) SELECT * FROM cte`
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
	fmt.Println("\n=== Test 3: * followed by one more column ===")
	sql := `WITH cte AS (SELECT *, name, ROW_NUMBER() OVER (ORDER BY id) AS Rank FROM table1) SELECT * FROM cte`
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

func test4() {
	fmt.Println("\n=== Test 4: Window function without AS alias ===")
	sql := `WITH cte AS (SELECT *, ROW_NUMBER() OVER (ORDER BY id) FROM table1) SELECT * FROM cte`
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

func test5() {
	fmt.Println("\n=== Test 5: * followed by window function (simple alias) ===")
	sql := `WITH cte AS (SELECT *, ROW_NUMBER() OVER (ORDER BY id) AS x FROM table1) SELECT * FROM cte`
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
	test4()
	test5()
}
