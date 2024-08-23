package routes

import (
	"local_my_api/handler"

	"github.com/gorilla/mux"
)

func SetupAuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()
	router.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	router.HandleFunc("/register", handler.RegisterHandler).Methods("POST")
	router.HandleFunc("/logout", handler.LogoutHandler).Methods("POST")
	router.HandleFunc("/refresh", handler.RefreshHandler).Methods("POST")
}
