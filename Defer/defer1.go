//geeksforgeeks.org/defer-keyword-in-golang/#:~:text=In%20Go%20language%2C%20defer%20statements,

package main

import "fmt"

// Functions
func mul(a1, a2 int) {

	res := a1 * a2
	fmt.Println("Result: ", res)

}

func show() {
	fmt.Println("Hello!, GeeksforGeeks")
}

// Main function
func main() {

	// Calling mul() function
	// Here mul function behaves
	// like a normal function
	mul(23, 45)

	// Calling mul()function
	// Using defer keyword
	// Here the mul() function
	// is defer function
	defer mul(23, 56)

	// Calling show() function
	show()
}

/* output will be
Result:  1035
Hello!, GeeksforGeeks
Result:  1288
*/
