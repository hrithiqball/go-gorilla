package main

import (
	"local_my_api/internal/db"
	"local_my_api/internal/handler"
	"local_my_api/internal/middlewares"
	"local_my_api/internal/repositories"
	"local_my_api/internal/routes"
	"local_my_api/internal/services"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	db.Connect()
	defer db.Close()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	// RUN_MIGRATION := os.Getenv("RUN_MIGRATION")

	// if RUN_MIGRATION == "true" {
	// 	err := db.DB.AutoMigrate(&models.User{})
	// 	if err != nil {
	// 		log.Fatalf("Error applying migrations: %v", err)
	// 	}
	// 	fmt.Println("Migrations applied successfully!")
	// }

	db.Migrations()

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

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
