package internal

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s Server) Health(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func CreateRouter(s Server) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/health", s.Health).Methods("GET")

	return r
}