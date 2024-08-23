package routes

import (
	"local_my_api/handler"

	"github.com/gorilla/mux"
)

func SetupProductRoutes(r *mux.Router) {
	router := r.PathPrefix("/products").Subrouter()
	router.HandleFunc("", handler.GetProducts).Methods("GET")
}
