package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func main() {
	debug := os.Getenv("DEBUG")
	requestType := os.Getenv("REQUEST_TYPE")

	// use Timeout in HTTP transaction
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		IdleConnTimeout:       90 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Timeout:   1 * time.Second, // basic timeout. 500 * time.Millisecond for half second
		Transport: tr,
	}
	request, err := http.NewRequest("GET", "https://ifconfig.co", nil)
	if err != nil {
		log.Fatal(err)
	}

	if requestType == "application/json" {
		// set request type
		request.Header.Add("Accept", "application/json")
	}

	if debug == "1" {
		debugResponse, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", debugResponse)
	}

	response, err := client.Do(request)
	defer response.Body.Close()

	if debug == "1" {
		debugResponse, err := httputil.DumpResponse(response, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", debugResponse)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", body)
}
