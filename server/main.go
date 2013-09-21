package main

import (
	"os"
	"fmt"
	"net"
)

func main() {
	if (len(os.Args) != 2) {
		fmt.Printf("Usage: %s path/to/socket\n", os.Args[0])
		os.Exit(1)
	}

	ln, err := net.Listen("unix", os.Args[1])
	if err != nil {
		panic(err)
	}
	defer unlinkSocket(os.Args[1])

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}

func unlinkSocket(path string) {
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("hi there")
}
