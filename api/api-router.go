package api

import (
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

type Company struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Employees   int32  `json:"employees"`
}

// APIRouter is a router for api
var APIRouter = mux.NewRouter().StrictSlash(false)

func init() {
	APIRouter.HandleFunc("/api/companies", getCompanies).Methods("GET")
	APIRouter.HandleFunc("/api/companies", addCompany).Methods("POST")
	APIRouter.HandleFunc("/api/companies/{id}", updateCompany).Methods("PUT")
	APIRouter.HandleFunc("/api/companies/{id}", removeCompany).Methods("DELETE")
}

func getCompanies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	company := &Company{"Facebook", "Social Network", 100500}
	response, err := json.Marshal(company)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func addCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	company := &Company{"Facebook", "Social Network", 100500}
	response, err := json.Marshal(company)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func updateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	company := &Company{"Facebook", "Social Network", 100500}
	response, err := json.Marshal(company)
	if err != nil {
		panic(err)
	}
	w.Write(response)
}

func removeCompany(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
