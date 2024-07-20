package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to TCP server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	_, err = conn.Write([]byte(name))
	if err != nil {
		fmt.Println("Error sending name:", err)
		return
	}

	go func() {
		for {
			message, err := serverReader.ReadString('\n')
			if err != nil {
				fmt.Println("Disconnected from server")
				os.Exit(0)
			}
			fmt.Print(message)
		}
	}()

	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}
