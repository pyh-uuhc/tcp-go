package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Client struct {
	conn net.Conn
	name string
}

var (
	clients     = make(map[net.Conn]Client)
	newClients  = make(chan Client)
	deadClients = make(chan net.Conn)
	messages    = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("TCP server listening on port 8080")

	go manageClients()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func manageClients() {
	for {
		select {
		case client := <-newClients:
			clients[client.conn] = client
			fmt.Println("New client connected:", client.name)

		case conn := <-deadClients:
			delete(clients, conn)
			fmt.Println("Client disconnected")

		case message := <-messages:
			for _, client := range clients {
				client.conn.Write([]byte(message))
			}
		}
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	conn.Write([]byte("Enter your name: "))
	name, _ := reader.ReadString('\n')
	client := Client{conn: conn, name: name}
	newClients <- client

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			deadClients <- conn
			return
		}
		messages <- fmt.Sprintf("%s: %s", client.name, message)
	}
}
