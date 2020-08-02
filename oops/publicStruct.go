package oops

import (
	"fmt"
)

type MyEmployeeStruct struct {
	empID   int
	empName string
}

func (emp MyEmployeeStruct) SearchEmp(empID int) {
	fmt.Println("the employee code is", empID)
}
