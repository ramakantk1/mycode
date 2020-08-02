// Example copied from the couse Todd McLeod
// the cource code is shared at https://github.com/GoesToEleven/go-programming
// the sample is picked from go-programming-master\code_samples\006-error-handling\07-custom-errors\05_custom-type
// i have added the func and line number in the effor usint runtime.caller(1) as per the exapmle in logginFuncLine.go

package error

import (
	"fmt"
	"log"
	"math"
	"runtime"
)

type norgateMathError struct {
	filename string
	line     int
	err      error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("The Error is Occurred in file: %v \n at line number %v \n and the error is %v\n", n.filename, n.line, n.err)
}

func main() {
	sq, err := sqrt(10)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("The square root is :", sq)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//Havent checked the Ok as per  _, file, line, ok := runtime.Caller(1) as we dont want to further checked it.

		_, file, line, _ := runtime.Caller(1)
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)

		return 0, norgateMathError{file, line, nme}
	}
	return math.Sqrt(f), nil
}

// see use of structs with error type in standard library:
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go
