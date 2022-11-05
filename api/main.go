package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", Welcome).Methods("GET")
	r.HandleFunc("/directory", Video).Methods("GET")
	r.HandleFunc("/directory/list", GayPorn).PathPrefix("l")
	http.ListenAndServe(":8000", r)

}
func Video(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("type") == "porno" {
		w.Write([]byte("1"))

	} else if r.URL.Query().Get("type") == "mult" {
		w.Write([]byte("2"))

	}

}

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))

}

func GayPorn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("l"))

}
