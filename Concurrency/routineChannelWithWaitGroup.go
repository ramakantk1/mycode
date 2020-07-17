// example copied from https://www.youtube.com/watch?v=IvKynoaD7PA

package main

import (
	"fmt"
	"sync"
	"time"
)

// we have two task to downlaod a web page
//read the values from the database and then perform some opertation on the download and database values
func main() {
	// define a channel , make with keyword as chan and data type can be anything
	// this is unbuffered channel means only one value can be assigned and if assigned it need to be read
	//done := make(chan bool)
	wg := sync.WaitGroup()

	//assument that downloading the webpage and database take 2 sec each
	//the below two line can be written as anonymaous function as for go routine
	//fmt.Println("downlaoding the web page")
	//time.Sleep(2 * time.Second)

	go func() {
		wg.Add(1)
		fmt.Println("started downlaoding the web page")
		time.Sleep(2 * time.Second)
		fmt.Println("Completed downlaoding the web page and next statement will update the waitgroup")
		//assigned a value to channel
		//done <- true
		wg.Done()
	}()

	fmt.Println("Read the database values")
	fmt.Println("started waiting for the wait group to complete")
	//time.Sleep(2 * time.Second)
	//if we done wait for channel value then we may conplete the main raoutine and will not able to get the web page download
	//try to comment below line and check the value
	//<-done
	wg.Wait()
	fmt.Println("Merging the data")

	fmt.Println("complete function/operation")
}
