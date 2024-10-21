package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the TCP server
	conn, err := net.Dial("tcp", "localhost:8080") // Change port if necessary
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter message (type 'STOP' to exit): ")
		text, _ := reader.ReadString('\n')

		// Send message to server
		_, err = fmt.Fprintf(conn, text)
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		// Read response from server
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Response from server: ", message)

		// Exit on receiving "STOP" command
		if text == "STOP\n" {
			fmt.Println("Exiting client...")
			break
		}
	}
}
