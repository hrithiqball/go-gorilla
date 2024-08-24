package routes

import (
	"local_my_api/internal/handler"

	"github.com/gorilla/mux"
)

func SetupAuthRoutes(r *mux.Router, h *handler.AuthHandler) {
	router := r.PathPrefix("/auth").Subrouter()
	router.HandleFunc("/login", h.LoginHandler).Methods("POST")
	router.HandleFunc("/register", h.RegisterHandler).Methods("POST")
	router.HandleFunc("/logout", h.LogoutHandler).Methods("POST")
	router.HandleFunc("/refresh", h.RefreshHandler).Methods("POST")
}
