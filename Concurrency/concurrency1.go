package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//example copied from the url https://www.youtube.com/watch?v=3atNYmqXyV4

func sendrequest(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[%d] %s\n", res.StatusCode, url)
}

//go run main.go google.com github.com tour.golang.org
//time go run mycode/Concurrency/concurrency1.go google.com github.com tour.golang.org
//time to run gitbash as it return the time take by program to run
func main() {
	if len(os.Args) < 2 {
		log.Fatalln("to run the command shouldb be go run mycode/Concurrency/concurrency1.go google.com bing.com github.com awsconsole.com")
	}

	for _, url := range os.Args[1:] {
		sendrequest("https://" + url)
	}
}

//real    0m4.332s
