package main

import (
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"github.com/smartbro/web-server-go/api"
)

func main() {
	address := ":8080"
	router := mux.NewRouter().StrictSlash(false)
	router.Handle("/api", api.APIRouter)
	router.Handle("/api/{_dummy:.*}", api.APIRouter)
	router.Handle("/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/"))))

	server := &http.Server{
		Addr:    address,
		Handler: router,
	}
	log.Println("Listening " + address)
	server.ListenAndServe()
}
