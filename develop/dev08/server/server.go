package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	listnerTCP, errTCP := net.Listen("tcp", ":8080")

	if errTCP != nil {
		panic(errTCP)
	}

	for {
		connTCP, errTCP := listnerTCP.Accept()
		if errTCP != nil {
			log.Printf("Error accepting connection from client: %s", errTCP)
		} else {
			go processClient(connTCP)
		}
	}
}

func processClient(conn net.Conn) {
	_, err := io.Copy(os.Stdout, conn)
	if err != nil {
		fmt.Println(err)
	}
	conn.Close()
}
