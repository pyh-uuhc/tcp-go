package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// TCP 서버를 시작하고 포트 8080을 수신 대기합니다.
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("TCP server listening on port 8080")

	for {
		// 클라이언트 연결을 수락합니다.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 각 클라이언트 연결을 별도의 고루틴으로 처리합니다.
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		// 클라이언트로부터 메시지를 읽습니다.
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}
		fmt.Printf("Message received: %s", message)

		// 클라이언트에 응답 메시지를 보냅니다.
		_, err = conn.Write([]byte("Message received\n"))
		if err != nil {
			fmt.Println("Error writing message:", err)
			return
		}
	}
}
