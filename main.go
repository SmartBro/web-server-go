package main

import (
	"net/http"

	"log"

	"github.com/smartbro/web-server-go/api"
)

func main() {
	address := ":8080"
	server := &http.Server{
		Addr:    address,
		Handler: api.APIRouter,
	}
	log.Println("Listening " + address)
	server.ListenAndServe()
}
