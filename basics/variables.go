package main // entry point package

import "fmt" // import fmt for printing

// Global variable (package scope)
var middleName = "Cane"

func main() {
	// Different ways to declare variables:
	// var age int                   // declares 'age' as int, default value 0
	// var name string = "John"      // declares and assigns
	// var name1 = "Jane"            // type inferred

	// Short variable declarations (inside functions only):
	// count := 10                   // int
	// lastName := "Smith"           // string
	// middleName := "Cane"          // would shadow the global 'middleName'
	// middleName := "Mayor"         // error: redeclaration

	fmt.Println(middleName) // prints the global variable "Cane"

	// Default values in Go (zero values):
	// Numeric Types → 0
	// Boolean Types → false
	// String Type → ""
	// Pointers, slices, maps, functions, structs → nil

	// ---- SCOPE
	// fmt.Println(firstName) // not accessible here, defined only in printname()
}

func printname() {
	firstName := "Michael" // local variable (function scope)
	fmt.Println(firstName) // prints "Michael"
}
