package main

import (
	"fmt"
	ast "github.com/haimiyahya/tsqlparser/ast"
	"github.com/haimiyahya/tsqlparser"
)

func main() {
	tests := []string{
		"CREATE TABLE IF NOT EXISTS TestTable (ID INT, Name VARCHAR(50))",
		"CREATE TABLE TestTable (ID INT, Name VARCHAR(50))",
		"CREATE TABLE IF NOT EXISTS dbo.TestTable (ID INT PRIMARY KEY)",
	}

	for _, sql := range tests {
		fmt.Printf("\n=== SQL: %s ===\n", sql)
		program, errs := tsqlparser.Parse(sql)
		fmt.Printf("Errors: %v (len=%d)\n", errs, len(errs))
		if len(errs) == 0 && len(program.Statements) > 0 {
			stmt := program.Statements[0]
			fmt.Printf("Statement type: %T\n", stmt)

			// Check if it's a CreateTableStatement
			if ct, ok := stmt.(*ast.CreateTableStatement); ok {
				fmt.Printf("IfNotExists: %v\n", ct.IfNotExists)
				fmt.Printf("Name: %s\n", ct.Name.String())
			}
		}
	}
}
