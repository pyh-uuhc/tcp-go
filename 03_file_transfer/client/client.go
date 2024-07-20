package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run client.go <file_path>")
		return
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Send file name length and file name
	fileName := filepath.Base(filePath)
	fileNameLen := int32(len(fileName))
	err = binary.Write(conn, binary.LittleEndian, fileNameLen)
	if err != nil {
		fmt.Println("Error sending file name length:", err)
		return
	}

	_, err = conn.Write([]byte(fileName))
	if err != nil {
		fmt.Println("Error sending file name:", err)
		return
	}

	// Send file content
	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println("Error sending file content:", err)
		return
	}

	fmt.Println("File sent successfully")
}
