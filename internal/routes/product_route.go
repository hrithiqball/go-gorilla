package routes

import (
	"local_my_api/internal/handler"
	"local_my_api/internal/middlewares"

	"github.com/gorilla/mux"
)

func SetupProductRoutes(r *mux.Router, h handler.ProductHandler) {
	public := r.PathPrefix("/product").Subrouter()
	public.HandleFunc("/list", h.GetProductListHandler).Methods("GET")
	public.HandleFunc("/{id}", h.GetProductHandler).Methods("GET")

	protected := r.PathPrefix("/product").Subrouter()
	protected.HandleFunc("", h.CreateProductHandler).Methods("POST")
	protected.HandleFunc("/{id}", h.UpdateProductHandler).Methods("PATCH")
	protected.HandleFunc("/{id}", h.DeleteProductHandler).Methods("DELETE")
	protected.Use(middlewares.AuthMiddleware)
}
