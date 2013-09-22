package main

import "os"
import "fmt"
import "net/http"
import "strconv"

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s port\n", os.Args[0])
		os.Exit(1)
	}

	port, err := strconv.Atoi(os.Args[1])
	if err != nil || port < 1 {
		fmt.Printf("Port must be a positive integer\n", os.Args[0])
		os.Exit(1)
	}

	initialize()
	mux := createMux()

	http.Handle("/", mux)

	listen := fmt.Sprintf("localhost:%d", port)

	fmt.Printf("Listening on %s\n", listen)

	err = http.ListenAndServe(listen, nil)
	if err != nil {
		fmt.Printf("Could not listen on %s\n", listen)
		os.Exit(1)
	}
}
