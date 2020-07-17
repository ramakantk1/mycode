package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

//example copied from the url https://www.youtube.com/watch?v=3atNYmqXyV4

// In the example we sending a request for the first url and wait for the response to come,
//once receive the response, check for the error, print and proceed to the next url
var wg sync.WaitGroup
var mut sync.Mutex

func sendrequest(url string) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	mut.Lock()
	fmt.Printf("[%d] %s\n", res.StatusCode, url)
	defer mut.Unlock()
}

//go run main.go google.com github.com tour.golang.org
//time go run mycode/Concurrency/concurrency1.go google.com github.com tour.golang.org
//time to run gitbash as it return the time take by program to run
func main() {
	if len(os.Args) < 2 {
		log.Fatalln("to run the command shouldb be go run mycode/Concurrency/concurrency1.go google.com bing.com github.com awsconsole.com")
	}

	for _, url := range os.Args[1:] {

		//sendrequest("https://" + url)
		go sendrequest("https://" + url)
		wg.Add(1)

	}
	wg.Wait()
}

//real    0m4.332s
//real    0m3.098s
