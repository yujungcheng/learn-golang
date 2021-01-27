// Concurrency is about dealing with lots of things at
// once. Parallelism is about doing lots of things at once.

// dealing -> to trade or handle alot tasks at once
// doing -> execute , process


package main

import (
  "fmt"
  "time"
  "log"
  "net/http"
)

func slowFunc() {
  time.Sleep(time.Second * 2)
  fmt.Println("sleeper() funished")
}

func responseTime(url string) {
  start := time.Now()
  res, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  elapsed := time.Since(start).Seconds()
  fmt.Printf("%s took %v seconds \n", url, elapsed)
}

func main() {

  // Blocking function call
  slowFunc()
  fmt.Println("I am not shown until slowFunc() completes")

  // Handling Concurrent Operation with goroutines
  go slowFunc()
  fmt.Println("I am now shown straightaway!")

  // showing goroutine concurrent execution
  go slowFunc()
  fmt.Println("I am not shown until slow() completes")
  time.Sleep(time.Second * 3)
  fmt.Println("Main function slepted 3 seconds")

  // demonstrating network latency
  urls := make([]string, 3)
  urls[0] = "https://www.usa.gov/"
  urls[1] = "https://www.gov.uk/"
  urls[2] = "http://www.gouvernement.fr/"
  for _, u := range urls {
    responseTime(u)
  }

  // Using goroutines to manage latency
  for _, u := range urls {
    go responseTime(u)
  }
  time.Sleep(time.Second * 5)
}
