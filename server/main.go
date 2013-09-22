package main

import (
	"os"
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
	"errors"
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
var next_id int

//  PRIVATE FUNCS  ///////////////////////////////////////////////////////////

func respond(w http.ResponseWriter, status string, list []Item, id int) {
	data := Status{
		Status: status,
		List: list,
		Id: id,
	}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func respondOK(w http.ResponseWriter) {
	respond(w, "ok", nil, 0)
}

func respondFail(w http.ResponseWriter) {
	respond(w, "fail", nil, 0)
}

func addItem(path string) (id int) {
	id = next_id
	items = append(items, Item{Id: next_id, Path: path})
	next_id = next_id + 1
	return
}

func removeItem(id int) (err error) {
	for i := 0; i < len(items); i++ {
		if items[i].Id == id {
			items = append(items[:i], items[i+1:]...)
			return
		}
	}
	err = errors.New("could not find id in item list")
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
	// path := r.PostFormValue("path")
	path := r.FormValue("path")

	id := addItem(path)
	respond(w, "ok", nil, id)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusNotImplemented)
		respondFail(w)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respondFail(w)
		return
	}

	err = removeItem(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respondFail(w)
		return
	}
	respondOK(w)
}

//  INITIALIZATION  //////////////////////////////////////////////////////////

func createMux() (m *mux.Router) {
	m = mux.NewRouter()
	m.HandleFunc("/", handleRoot)
	m.HandleFunc("/tracks", handleList)
	m.HandleFunc("/tracks/new", handleNew)
	m.HandleFunc("/tracks/{id}", handleDelete)
	return
}

func initialize() {
	items = []Item{}
	next_id = 1
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
