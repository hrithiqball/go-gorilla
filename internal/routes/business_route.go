package routes

import (
	"local_my_api/internal/handler"
	"local_my_api/internal/middlewares"

	"github.com/gorilla/mux"
)

func SetupBusinessRoutes(r *mux.Router, h handler.BusinessHandler) {
	public := r.PathPrefix("/business").Subrouter()
	public.HandleFunc("list", h.GetBusinessListHandler).Methods("GET")
	public.HandleFunc("/{id}", h.GetBusinessHandler).Methods("GET")

	protected := r.PathPrefix("/business").Subrouter()
	protected.HandleFunc("", h.CreateBusinessHandler).Methods("POST")
	protected.HandleFunc("/{id}", h.UpdateBusinessHandler).Methods("PATCH")
	protected.HandleFunc("/{id}", h.DeleteBusinessHandler).Methods("DELETE")
	protected.Use(middlewares.AuthMiddleware)
}
