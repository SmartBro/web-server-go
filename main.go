package main

import (
	"net/http"

	"github.com/smartbro/web-server-go/api"
)

func main() {
	http.ListenAndServe(":8080", api.APIRouter)
}
