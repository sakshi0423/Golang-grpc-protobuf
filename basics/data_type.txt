package main

import "fmt"

func main() {

	fmt.Println("Hello" + "World")           // Concatenates two strings and prints "HelloWorld"
	fmt.Println("9 X 10 =", 9*10)            // Performs multiplication (9*10 = 90) and prints it
	fmt.Println("180.18/2.0 = ", 180.18/2.0) // Performs floating-point division and prints result
	fmt.Println(true && false)               // Logical AND → prints false
	fmt.Println(true || false)               // Logical OR → prints true
	fmt.Println(!true)                       // Logical NOT → prints false
}
