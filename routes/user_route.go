package routes

import (
	"local_my_api/handler"
	"local_my_api/middleware"

	"github.com/gorilla/mux"
)

func SetupUserRoutes(r *mux.Router) {
	public := r.PathPrefix("/user").Subrouter()
	public.HandleFunc("/list", handler.GetUserListHandler).Methods("GET")
	public.HandleFunc("/{id}", handler.GetUserHandler).Methods("GET")

	protected := r.PathPrefix("/user").Subrouter()
	protected.HandleFunc("/{id}", handler.UpdateUserHandler).Methods("PATCH")
	protected.HandleFunc("/{id}", handler.DeleteUserHandler).Methods("DELETE")
	protected.Use(middleware.AuthMiddleware)
}
