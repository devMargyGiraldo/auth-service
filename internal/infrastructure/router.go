package infrastructure

import (
	"auth_service/internal/interfaces/handler"
	"log"
	"net/http"

	_ "auth_service/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(authHandler *handler.AuthHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return r
}

func StartServer(r *mux.Router) {
	log.Println("Server starting on port 8080...")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
