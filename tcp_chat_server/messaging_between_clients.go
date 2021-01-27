package main

import (
  "fmt"
  "io"
  "log"
  "net"
)

type Message struct {
  conn net.Conn
  message []byte
}

var connections []net.Conn
var addClient = make(chan net.Conn)
var removeClient = make(chan net.Conn)
var messages = make(chan Message)


func main() {
  server, err := net.Listen("tcp", ":8000")
  if err != nil {
    log.Fatal(err)
  }
  defer server.Close()

  go startChannels()

  for {
    conn, err := server.Accept()
    if err != nil {
      log.Fatal(err)
    }
    addClient <- conn
    go handleRequest(conn)
  }
}

func startChannels() {
  for {
    select {
    case message := <-messages:
      broadcastMessage(&message)
    case newClient := <-addClient:
      connections = append(connections, newClient)
      fmt.Println("total connections: ", len(connections))
    case deadClient := <-removeClient:
      removeConn(deadClient)
      fmt.Println("total connections: ", len(connections))
    }
  }
}

func handleRequest(conn net.Conn) {
  for {
    message := make([]byte, 64)
    _, err := conn.Read(message)
    if err != nil {
      if err == io.EOF {
        removeClient <- conn
        conn.Close()
        return
      }
      log.Fatal(err)
    }

    m := Message {
      conn: conn,
      message: message,
    }
    messages <- m
  }
}

func broadcastMessage(m *Message) {
  for _, conn := range connections {
    # todo: skip blank line
    if string(m.message) == "\n" {
      continue
    }
    fmt.Println(m.message)
    remoteIP := conn.RemoteAddr().String()
    remoteIP += "> "
    byteMessage := append([]byte(remoteIP), m.message...)
    if m.conn != conn {
      _, err := conn.Write(byteMessage)
      //_, err := conn.Write(m.message)
      if err != nil {
        log.Fatal(err)
      }
    } else {
      fmt.Printf(string(byteMessage))
    }
  }
}

func removeConn(conn net.Conn) {
  var i int
  for i = range connections {
    if connections[i] == conn {
      break
    }
  }
  connections = append(connections[:i], connections[i+1:]...)
}
