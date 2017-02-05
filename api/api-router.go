package api

import "net/http"

// APIRouter is a router for api
var APIRouter = http.NewServeMux()

func init() {
	APIRouter.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
