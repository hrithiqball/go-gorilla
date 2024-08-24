package routes

import (
	"local_my_api/internal/handler"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, authHandler *handler.AuthHandler, userHandler handler.UserHandler, businessHandler handler.BusinessHandler, productHandler handler.ProductHandler) {
	SetupAuthRoutes(r, authHandler)
	SetupUserRoutes(r, userHandler)
	SetupBusinessRoutes(r, businessHandler)
	SetupProductRoutes(r, productHandler)
}
