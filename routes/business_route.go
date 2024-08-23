package routes

import (
	"local_my_api/handler"

	"github.com/gorilla/mux"
)

func SetupBusinessRoutes(r *mux.Router) {
	router := r.PathPrefix("/business").Subrouter()
	router.HandleFunc("", handler.CreateBusiness).Methods("POST")
	router.HandleFunc("list", handler.GetBusinessList).Methods("GET")
}
