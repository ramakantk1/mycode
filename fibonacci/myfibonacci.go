//Https://play.golang.org/p/Hh8Owj1CXp4

package main

import (
	"fmt"
)

func main() {

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

}
