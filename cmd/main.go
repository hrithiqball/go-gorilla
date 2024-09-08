package main

import (
	"context"
	"fmt"
	"local_my_api/internal/db"
	"local_my_api/internal/handler"
	"local_my_api/internal/middlewares"
	"local_my_api/internal/repositories"
	"local_my_api/internal/routes"
	"local_my_api/internal/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var version = "unknown"

func main() {
	loadVersion()
	fmt.Printf("App Version: %s\n", version)

	if err := db.Connect(); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Migrations(); err != nil {
		log.Fatalf("Could not run migrations: %v", err)
	}

	userRepo := repositories.NewUserRepository(db.DB)
	businessRepo := repositories.NewBusinessRepository(db.DB)
	productRepo := repositories.NewProductRepository(db.DB)

	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	businessService := services.NewBusinessService(businessRepo)
	productService := services.NewProductService(productRepo)

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	businessHandler := handler.NewBusinessHandler(businessService)
	productHandler := handler.NewProductHandler(productService, businessService)

	router := mux.NewRouter()
	router.Use(middlewares.LoggingMiddleware)
	routes.SetupRoutes(router, authHandler, userHandler, businessHandler, productHandler)

	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)
	corsRouter := corsOptions(router)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: corsRouter,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", PORT, err)
		}
	}()

	fmt.Printf("🚀 Server is listening on port %s! 🚀 \n", PORT)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("🪽 Server exited gracefully")
}

func loadVersion() {
	data, err := os.ReadFile("version.txt")
	if err != nil {
		version = "development"
	}

	version = string(data)
}
