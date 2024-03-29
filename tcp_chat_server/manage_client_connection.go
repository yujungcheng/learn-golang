package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var connections []net.Conn
var addClient = make(chan net.Conn)
var removeClient = make(chan net.Conn)

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
		message := make([]byte, 1024)
		_, err := conn.Read(message)
		if err != nil {
			if err == io.EOF {
				removeClient <- conn
				conn.Close()
				return
			}
			log.Fatal(err)
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
