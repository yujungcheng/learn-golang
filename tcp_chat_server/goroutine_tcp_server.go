package main

import (
	"io"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleRequest(conn)
	}

	//conn.Write([]byte("Hello World!\n"))
	//conn.Close()
}

func handleRequest(conn net.Conn) {
	io.Copy(conn, conn)
}
