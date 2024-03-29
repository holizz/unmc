package main

import "net/http"
import "strconv"
import "encoding/json"
import "github.com/gorilla/mux"

//  HELPER FUNCTIONS  ////////////////////////////////////////////////////////

func respond(w http.ResponseWriter, status string, list []Item, id int) {
	data := Status{
		Status: status,
		List:   list,
		Id:     id,
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

//  HANDLE  //////////////////////////////////////////////////////////////////

func handleRoot(w http.ResponseWriter, r *http.Request) {
	respondOK(w)
}

func handleList(w http.ResponseWriter, r *http.Request) {
	respond(w, "ok", items, 0)
}

func handleNew(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")

	id := addItem(path)
	respond(w, "ok", nil, id)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
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

func handlePlay(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respondFail(w)
		return
	}

	index, err := itemIndex(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respondFail(w)
		return
	}

	err = audioPlay(index)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respondFail(w)
		return
	}
}

//  INITIALIZATION  //////////////////////////////////////////////////////////

func createMux() (m *mux.Router) {
	m = mux.NewRouter()
	m.HandleFunc("/", handleRoot).Methods("GET")
	m.HandleFunc("/tracks", handleList).Methods("GET")
	m.HandleFunc("/tracks/new", handleNew).Methods("PUT")
	m.HandleFunc("/tracks/{id}", handleDelete).Methods("DELETE")
	m.HandleFunc("/control/play", handlePlay).Methods("POST")
	return
}
