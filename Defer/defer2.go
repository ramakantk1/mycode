//https://www.geeksforgeeks.org/defer-keyword-in-golang/#:~:text=In%20Go%20language%2C%20defer%20statements,until%20the%20nearby%20functions%20returns.

// Go program to illustrate
// multiple defer statements
package main

import "fmt"

// Functions
func add(a1, a2 int) {
	res := a1 + a2
	fmt.Println("Result: ", res)

}

// Main function
func main() {

	fmt.Println("Start")

	// Multiple defer statements
	// Executes in LIFO order
	defer fmt.Println("End")
	//defer add(34, 56)
	//defer add(10, 10)
	defer fmt.Println(30 + 40)
	defer fmt.Println(10 + 10)
}
