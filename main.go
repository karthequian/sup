package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/version", VersionIndex)
	router.HandleFunc("/sup", Sup)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

//VersionIndex returns version data
func VersionIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "{\"Version\":\"1.0\"}")
}

//Index returns version data
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

//Sup returns Sup
func Sup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sup!")
}

//TodoShow returns todo info
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoID)
}
