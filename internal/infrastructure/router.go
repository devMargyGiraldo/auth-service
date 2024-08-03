package infrastructure

import (
	"auth_service/internal/interfaces/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(authHandler *handler.AuthHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")

	return r
}

func StartServer(r *mux.Router) {
	http.ListenAndServe(":8080", r)
}
