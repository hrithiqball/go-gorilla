package routes

import (
	"local_my_api/internal/handler"
	"local_my_api/internal/middlewares"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, authHandler *handler.AuthHandler, userHandler handler.UserHandler, businessHandler handler.BusinessHandler, productHandler handler.ProductHandler) {
	rateLimiter := middlewares.NewRateLimiter()
	r.Use(rateLimiter.Middleware)

	SetupAuthRoutes(r, authHandler)
	SetupUserRoutes(r, userHandler)
	SetupBusinessRoutes(r, businessHandler)
	SetupProductRoutes(r, productHandler)
}
