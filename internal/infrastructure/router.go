package infrastructure

import (
	"github.com/gorilla/mux"
	"net/http"
	httpHandler "github.com/shirocola/user-service/internal/interface/http"  // Alias this import
	"log"
)

// InitRouter initializes the mux router and defines the routes
func InitRouter(userHandler *httpHandler.UserHandler) *mux.Router {  // Use the alias here
	router := mux.NewRouter()

	// Define your routes here
	router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	return router
}

// StartServer starts the HTTP server
func StartServer(router *mux.Router) {
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
