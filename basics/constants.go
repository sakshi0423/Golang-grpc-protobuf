package basics // Defines the package name; here it's "basics"

// Declare constants (immutable values that don’t change)
const pi = 3.14      // Constant for the value of Pi (lowercase → package-private)
const GRAVITY = 9.81 // Constant for gravity (uppercase → exported, accessible outside the package)

func main() {
	// Constant with explicit type (int)
	const days int = 7 // Number of days in a week

	// Constant block: allows grouping multiple related constants together
	const (
		monday       = 1 // Assigns 1 to monday
		tuesday      = 2 // Assigns 2 to tuesday
		wenesday     = 3 // Assigns 3 to wenesday (typo, should be "wednesday")
		thursday int = 4 // Explicitly typed constant, value 4
	)
}
