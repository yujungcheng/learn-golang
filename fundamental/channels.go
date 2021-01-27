package main

import (
  "fmt"
  "time"
)

func slowFunc(c chan string) {
  //c <- "slowFunc() started"
  time.Sleep(time.Second * 2)
  c <- "slowFunc() finished"
}

func receiver(c chan string) {
  for msg := range c {
    fmt.Println(msg)
  }
}

func pinger(c chan string) {
  t := time.NewTicker(1 * time.Second)
  for {
    c <- "ping"
    <-t.C
  }
}

// Using channels as function arguments
func channelReader(messages <-chan string) {
  msg := <- messages
  fmt.Println(msg)
}
func channelWriter(messages chan<- string) {
  messages <- "Hello world"
}
func channelReaderAndWriter(messages chan string) {
  msg := <- messages
  fmt.Println(msg)
  messages <- "Hello world"
}

func ping1(c chan string) {
  time.Sleep(time.Second * 1)
  c <- "ping on channel1"
}

func ping2(c chan string) {
  time.Sleep(time.Second *2)
  c <- "ping on channel2"
}

func sender(c chan string) {
  t := time.NewTicker(1 * time.Second)
  for {
    c <- "I'm sending a message"
    <- t.C
  }
}

func main() {
  c := make(chan string)
  go slowFunc(c)

  msg := <-c
  fmt.Println(msg)

  // using buffered channel
  msgs := make(chan string, 2)
  msgs <- "hello"
  msgs <- "world"
  close(msgs)
  fmt.Println("Pushed two messages on to channel with no receivers")
  time.Sleep(time.Second * 1)
  receiver(msgs)

  // Blocking and flow control
  c2 := make(chan string)
  go pinger(c2)
  r_msg := <- c2
  fmt.Println(r_msg)

  fmt.Println("- Indefinitely receive")
  /*
  for {
    r_msg = <- c2
    fmt.Println(r_msg)

  }
  */

  fmt.Println("- receive a certain number of messages")
  for i := 0; i < 5; i++ {
    r_msg = <- c2
    fmt.Println(r_msg)
  }

  // employing the select statement
  channel1 := make(chan string)
  channel2 := make(chan string)
  go ping1(channel1)
  go ping2(channel2)
  select {
    case msg1 := <- channel1:
      fmt.Println("Received", msg1)
    case msg2 := <- channel2:
      fmt.Println("Received", msg2)
    case <- time.After(500 * time.Millisecond):
      fmt.Println("no messages received. giving up.")
  }

  // Quitting channels
  messages := make(chan string)
  stop := make(chan bool)
  go sender(messages)
  go func() {
    time.Sleep(time.Second * 2)
    fmt.Println("Time's up!")
    stop <- true
  } ()

  for {
    select {
      case <- stop:
        fmt.Println("main return.")
        return
      case  new_msg := <- messages:
        fmt.Println(new_msg)
    }
  }

}
