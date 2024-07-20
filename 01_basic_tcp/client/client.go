package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// TCP 서버에 연결을 시도합니다.
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to TCP server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 사용자로부터 입력을 받아 서버에 메시지를 보냅니다.
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message to send: ")
		message, _ := reader.ReadString('\n')

		// 서버에 메시지를 보냅니다.
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing message:", err)
			return
		}

		// 서버로부터 응답 메시지를 읽습니다.
		reply, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading reply:", err)
			return
		}
		fmt.Printf("Reply from server: %s", reply)
	}
}
