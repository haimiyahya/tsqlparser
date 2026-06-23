package main

import (
	"fmt"
	tsqlparser "github.com/haimiyahya/tsqlparser"
)

func test1() {
	fmt.Println("=== Test 1: Simple ROW_NUMBER with AS alias ===")
	sql := `SELECT *, ROW_NUMBER() OVER (ORDER BY id) AS Rank FROM table1`
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

func test2() {
	fmt.Println("\n=== Test 2: ROW_NUMBER with PARTITION BY and ORDER BY DESC ===")
	sql := `SELECT *, ROW_NUMBER() OVER (PARTITION BY cat ORDER BY id DESC) AS Rank FROM table1`
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

func test3() {
	fmt.Println("\n=== Test 3: ROW_NUMBER in CTE ===")
	sql := `WITH cte AS (SELECT *, ROW_NUMBER() OVER (ORDER BY id) AS Rank FROM table1) SELECT * FROM cte`
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

func test4() {
	fmt.Println("\n=== Test 4: The problematic SQL ===")
	sql := `WITH LatestRecords AS (SELECT *, ROW_NUMBER() OVER (PARTITION BY ExtRefNo ORDER BY VersionID DESC) AS Rank FROM DSM_QuotaAcctTracking) SELECT * FROM LatestRecords`
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

func main() {
	test1()
	test2()
	test3()
	test4()
}
