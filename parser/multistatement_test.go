package parser

import (
	"fmt"
	"testing"

	"github.com/haimiyahya/tsqlparser/ast"
	"github.com/haimiyahya/tsqlparser/lexer"
)

// TestMultiStatementBatchWithInsert tests parsing of multi-statement batches
// that include INSERT statements.
//
// BUG: INSERT statements in multi-statement batches are not being parsed correctly.
// For example: "CREATE TABLE ...; INSERT INTO ...; SELECT ..."
// Only the CREATE TABLE and SELECT are parsed, the INSERT is missing.
func TestMultiStatementBatchWithInsert(t *testing.T) {
	input := `CREATE TABLE TestT (ID INT IDENTITY(1,1), Name NVARCHAR(50));
INSERT INTO TestT (Name) VALUES (N'A');
SELECT @@IDENTITY AS LastID`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	if p.Errors() != nil {
		t.Fatalf("Parser errors: %v", p.Errors())
	}

	fmt.Printf("=== Multi-Statement Batch Test ===\n")
	fmt.Printf("Input: %s\n\n", input)
	fmt.Printf("Number of statements parsed: %d\n", len(program.Statements))

	// Expected: 3 statements (CREATE TABLE, INSERT, SELECT)
	expectedStatements := 3
	if len(program.Statements) != expectedStatements {
		t.Errorf("Expected %d statements, got %d", expectedStatements, len(program.Statements))
	}

	// Print each statement type
	for i, stmt := range program.Statements {
		fmt.Printf("Statement %d: %T\n", i+1, stmt)
	}

	// Verify we have the correct types
	statementTypes := make([]string, len(program.Statements))
	for i, stmt := range program.Statements {
		switch stmt.(type) {
		case *ast.CreateTableStatement:
			statementTypes[i] = "CreateTableStatement"
		case *ast.InsertStatement:
			statementTypes[i] = "InsertStatement"
		case *ast.SelectStatement:
			statementTypes[i] = "SelectStatement"
		default:
			statementTypes[i] = fmt.Sprintf("%T", stmt)
		}
	}

	// Check for expected statement types
	expectedTypes := []string{"CreateTableStatement", "InsertStatement", "SelectStatement"}
	for i, expectedType := range expectedTypes {
		if i >= len(statementTypes) {
			t.Errorf("Missing statement %d (%s)", i+1, expectedType)
			continue
		}
		if statementTypes[i] != expectedType {
			t.Errorf("Statement %d: expected %s, got %s", i+1, expectedType, statementTypes[i])
		}
	}
}

// TestSimpleInsertStatement tests that a single INSERT statement parses correctly.
func TestSimpleInsertStatement(t *testing.T) {
	input := `INSERT INTO TestT (Name) VALUES (N'A')`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	if p.Errors() != nil {
		t.Fatalf("Parser errors: %v", p.Errors())
	}

	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
	}

	_, ok := program.Statements[0].(*ast.InsertStatement)
	if !ok {
		t.Errorf("Expected InsertStatement, got %T", program.Statements[0])
	}
}

// TestInsertSelectInBatch tests INSERT followed by SELECT in a batch.
func TestInsertSelectInBatch(t *testing.T) {
	input := `INSERT INTO TestT (Name) VALUES (N'A');
SELECT @@IDENTITY AS LastID`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	if p.Errors() != nil {
		t.Fatalf("Parser errors: %v", p.Errors())
	}

	fmt.Printf("\n=== Insert + Select Batch Test ===\n")
	fmt.Printf("Input: %s\n\n", input)
	fmt.Printf("Number of statements parsed: %d\n", len(program.Statements))

	for i, stmt := range program.Statements {
		fmt.Printf("Statement %d: %T\n", i+1, stmt)
	}

	// Expected: 2 statements (INSERT, SELECT)
	expectedStatements := 2
	if len(program.Statements) != expectedStatements {
		t.Errorf("Expected %d statements, got %d", expectedStatements, len(program.Statements))
	}
}
