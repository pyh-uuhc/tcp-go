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
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		scanner.Scan()
		text := scanner.Text()

		if text == "" {
			continue
		}

		_, err := fmt.Fprintf(conn, text+"\n")
		if err != nil {
			fmt.Println("Error sending data to server:", err)
			break
		}

		serverScanner := bufio.NewScanner(conn)
		if serverScanner.Scan() {
			fmt.Println("Server response:", serverScanner.Text())
		}

		if text == "EXIT" {
			fmt.Println("Exiting client")
			break
		}
	}
}
