//https://play.golang.org/p/4gluMnl-vDj

package main

import (
	"fmt"
)

func main() {
	/* without recursive*/
	/*
		x, y := 0, 1
		if x == 0 {
			fmt.Println(0)
			fmt.Println(1)
		}
		for i := 0; i < 10; i++ {
			fmt.Println(x + y)
			tempi := x
			x, y = y, tempi+y

		}
	*/

	for i := 0; i <= 9; i++ {
		fmt.Println(FibonacciRecursion(i))
	}

}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}

	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
