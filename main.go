package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func fallbackHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404")
}

func main() {

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(fallbackHandler)

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
