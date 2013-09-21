package main

import (
	"os"
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)

//  STRUCTS  /////////////////////////////////////////////////////////////////

type Item struct {
	Id int
	Path string
}

type Status struct {
	Status string
	List []Item
	Id int
}

//  VARS  ////////////////////////////////////////////////////////////////////

var items []Item

//  PRIVATE FUNCS  ///////////////////////////////////////////////////////////

func respond(w http.ResponseWriter, status string, list []Item, id int) {
	data := Status{
		Status: status,
		List: list,
		Id: id,
	}
	json, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(json)
}

func respondOK(w http.ResponseWriter) {
	respond(w, "ok", nil, 0)
}

func respondFail(w http.ResponseWriter) {
	respond(w, "fail", nil, 0)
}

func addItem(path string) (id int) {
	id = 1
	items = append(items, Item{Id: id, Path: path})
	return
}

//  HANDLE  //////////////////////////////////////////////////////////////////

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotImplemented)
		respondFail(w)
		return
	}
	respondOK(w)
}

func handleList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotImplemented)
		respondFail(w)
		return
	}
	respond(w, "ok", items, 0)
}

func handleNew(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusNotImplemented)
		respondFail(w)
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respondFail(w)
		return
	}
	path := r.PostFormValue("path")

	id := addItem(path)
	respond(w, "ok", nil, id)
}

//  INITIALIZATION  //////////////////////////////////////////////////////////

func createMux() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/tracks", handleList)
	mux.HandleFunc("/tracks/new", handleNew)
	return
}

func initialize() {
	items = []Item{}
}

//  MAIN  ////////////////////////////////////////////////////////////////////

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

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Could not listen on port %d\n", port)
		os.Exit(1)
	}
}
