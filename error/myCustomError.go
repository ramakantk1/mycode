// This is my example to log or printf the error with more details
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fo, err := os.Open("/tmp/dat/my.txt")
	if err != nil {
		rkError(err, " Not able to open the file /tmp/dat/my.txt")
	} else {
		fmt.Println("The file with name", fo, "exist")
	}

}

// this function return error but since the rkErrorDetails also implement the error interface we can rertun the rkErrorDetails
func rkError(err error, customErrorMessage string) {
	//we can have two different log based on some cofig setting
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("The Error is Occurred in file: %v \n at line number %v \n and the error is %v\n Custom meesage sent by function is %v\n", file, line, err, customErrorMessage)
}
