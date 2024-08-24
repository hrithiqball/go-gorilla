package routes

import (
	"local_my_api/internal/handler"
	"local_my_api/internal/middlewares"

	"github.com/gorilla/mux"
)

func SetupUserRoutes(r *mux.Router, h handler.UserHandler) {
	public := r.PathPrefix("/user").Subrouter()
	public.HandleFunc("/list", h.GetUserListHandler).Methods("GET")
	public.HandleFunc("/{id}", h.GetUserHandler).Methods("GET")

	protected := r.PathPrefix("/user").Subrouter()
	protected.HandleFunc("/{id}", h.UpdateUserHandler).Methods("PATCH")
	protected.HandleFunc("/{id}", h.DeleteUserHandler).Methods("DELETE")
	protected.Use(middlewares.AuthMiddleware)
}
