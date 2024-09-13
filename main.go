package main

import (
	"encoding/json"
	"log"
	"net/http"
	"risks-api/risk"
	"risks-api/storage"

	"github.com/gorilla/mux"
)

var db = storage.NewInMemoryStorage()

func GetAllRisks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	risks := db.GetAllRisks()
	b, err := json.Marshal(risks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func GetRisk(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	riskID := v["id"]

	risk, err := db.GetRisk(riskID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(risk)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

type incomingRisk struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func PostNewRisk(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var incoming incomingRisk
	err := decoder.Decode(&incoming)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	risk := risk.NewRisk(incoming.Title, incoming.Description)
	err = db.PutRisk(risk)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(risk)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func fallbackHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/v1/risks", GetAllRisks).Methods("GET")
	r.HandleFunc("/v1/risks/{id}", GetRisk).Methods("GET")
	r.HandleFunc("/v1/risks", PostNewRisk).Methods("POST")
	r.PathPrefix("/").HandlerFunc(fallbackHandler)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
