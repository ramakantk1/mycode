// example copied for https://www.linode.com/docs/development/go/go-context/
//have impleted custom as well as default httpclient.
//custom http client give more control over proxies, TLS configuration, keep-alives, compression, and other settings, create a Transport:
// https://golang.org/pkg/net/http/
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	myUrl string
	delay int = 5
	wg    sync.WaitGroup
)

type myData struct {
	r   *http.Response
	err error
}

func connect(ctx context.Context) {
	//call wg.Done when the function is completed
	defer wg.Done()
	// define a buffer channel with size 1 of type myData structure
	data := make(chan myData, 1)

	// we are using custon http call and the reason well explained on the below url
	// https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
	// One more Url which we can refer is belwo, as per it we create transport to have better control on proxies, TLS configuration, keep-alives, compression, and other settings, create a Transport:
	// https://golang.org/pkg/net/http/
	tr := &http.Transport{}
	httpClient := &http.Client{Transport: tr}

	//creating the req object

	req, _ := http.NewRequest("GET", myUrl, nil)

	//calling go routing to fetch the response
	go func() {
		//in smipler request this can also be done as

		res, err := httpClient.Do(req)
		if err != nil {
			fmt.Println(err)
			data <- myData{nil, err}
			return
		} else {
			pack := myData{res, err}
			data <- pack
		}

	}()

	select {
	// if root is done
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-data
		fmt.Println("the request was canceled")
	case ok := <-data:
		err := ok.err
		res := ok.r
		if err != nil {
			fmt.Println("error in response", err.Error)
			return
		}
		defer res.Body.Close()

		realHttpData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("error in parising data", err.Error)
			return
		}
		fmt.Printf("Server Response: %s\n", realHttpData)
	}
	return
}

func simpleConnect(ctx context.Context) {
	//call wg.Done when the function is completed
	defer wg.Done()
	// define a buffer channel with size 1 of type myData structure
	data := make(chan myData, 1)

	// we are using simple http call and the reason well explained on the below url
	// https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
	// One more Url which we can refer is belwo, as per it we create transport to have better control on proxies, TLS configuration, keep-alives, compression, and other settings, create a Transport:
	// https://golang.org/pkg/net/http/

	//calling go routing to fetch the response
	go func() {
		//in smipler request this can also be done as

		res, err := http.Get(myUrl)
		if err != nil {
			fmt.Println(err)
			data <- myData{nil, err}
			return
		} else {
			pack := myData{res, err}
			data <- pack
		}

	}()

	select {
	// if root is done
	case <-ctx.Done():

		fmt.Println("the request was canceled")
	case ok := <-data:
		err := ok.err
		res := ok.r
		if err != nil {
			fmt.Println("error in response", err.Error)
			return
		}
		defer res.Body.Close()

		realHttpData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("error in parising data", err.Error)
			return
		}
		fmt.Printf("Server Response: %s\n", realHttpData)
	}
	return
}
func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need a URL and a delay!")
		return
	}

	myUrl = os.Args[1]
	if len(os.Args) == 3 {
		t, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		delay = t
	}

	fmt.Println("Delay:", delay)
	c := context.Background()
	c, cancel := context.WithTimeout(c, time.Duration(delay)*time.Second)
	defer cancel()

	fmt.Printf("Connecting to %s \n", myUrl)
	wg.Add(1)
	go simpleConnect(c)
	//go connect(c)
	wg.Wait()
	fmt.Println("Exiting...")
}
