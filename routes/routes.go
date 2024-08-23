package routes

import "github.com/gorilla/mux"

func SetupRoutes(r *mux.Router) {
	SetupAuthRoutes(r)
	SetupUserRoutes(r)
	SetupBusinessRoutes(r)
	SetupProductRoutes(r)
}
