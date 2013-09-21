package main

import (
	"os"
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	type Status struct {
		Status string
	}
	status := Status{
		Status: "ok",
	}
	data, err := json.Marshal(status)
	if err != nil {
		panic(err)
	}
	w.Write(data)
}

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

	http.HandleFunc("/", handleRoot)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
