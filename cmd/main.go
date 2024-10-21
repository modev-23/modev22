package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)

	listener, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	for {
		// Accept client connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			continue
		}

		fmt.Println("Client connected:", conn.RemoteAddr().String())

		// Handle the connection concurrently
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read incoming messages from the client
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client left:", conn.RemoteAddr().String())
			return
		}

		fmt.Printf("Received message from %s: %s", conn.RemoteAddr().String(), message)

		// Echo the message back to the client
		_, err = conn.Write([]byte("Echo: " + message))
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}
	}
}
