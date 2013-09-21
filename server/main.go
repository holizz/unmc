package main

import (
	"os"
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)

type Item struct {
	Id int
	Path string
}

type Status struct {
	Status string
	List []Item
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	status := Status{
		Status: "ok",
	}
	data, err := json.Marshal(status)
	if err != nil {
		panic(err)
	}
	w.Write(data)
}

func handleList(w http.ResponseWriter, r *http.Request) {
	status := Status{
		Status: "ok",
		List: []Item{},
	}
	data, err := json.Marshal(status)
	if err != nil {
		panic(err)
	}
	w.Write(data)
}

func createMux() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/tracks", handleList)
	return
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

	mux := createMux()

	http.Handle("/", mux)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
