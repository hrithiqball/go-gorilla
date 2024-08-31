package routes

import (
	"local_my_api/internal/handler"
	"local_my_api/internal/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, authHandler *handler.AuthHandler, userHandler handler.UserHandler, businessHandler handler.BusinessHandler, productHandler handler.ProductHandler) {
	rateLimiter := middlewares.NewRateLimiter()
	r.Use(rateLimiter.Middleware)
	r.PathPrefix("/bucket/").Handler(http.StripPrefix("/bucket/", http.FileServer(http.Dir("./public/uploads"))))

	SetupAuthRoutes(r, authHandler)
	SetupUserRoutes(r, userHandler)
	SetupBusinessRoutes(r, businessHandler)
	SetupProductRoutes(r, productHandler)
}
