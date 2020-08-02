// example copied from https://github.com/golang/go/issues/663
// it log or printf the function and line number at which it is called
/* OutPut
E:/Ramakant/GO/Project/src/mycode/error/logginFuncLine.go:12 line # 9
line # 10
E:/Ramakant/GO/Project/src/mycode/error/logginFuncLine.go:14 line # 11
line # 12
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	log("line # 9")
	fmt.Println("line # 10")
	log("line # 11")
	fmt.Println("line # 12")
}

func log(message string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("some error")
		return
	}
	fmt.Printf("%s:%d %s\n", file, line, message)
}
