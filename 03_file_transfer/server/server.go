package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Connected to %s\n", conn.RemoteAddr().String())

	// Create a buffer to store file name length
	var fileNameLen int32
	err := binary.Read(conn, binary.LittleEndian, &fileNameLen)
	if err != nil {
		fmt.Println("Error reading file name length:", err)
		return
	}

	// Create a buffer to store file name
	fileName := make([]byte, fileNameLen)
	_, err = io.ReadFull(conn, fileName)
	if err != nil {
		fmt.Println("Error reading file name:", err)
		return
	}

	// Create the file
	f, err := os.Create(filepath.Base(string(fileName)))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	// Copy the file content
	_, err = io.Copy(f, conn)
	if err != nil {
		fmt.Println("Error writing file content:", err)
		return
	}

	fmt.Printf("Received file %s\n", fileName)
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer ln.Close()

	fmt.Println("Server started on :8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
