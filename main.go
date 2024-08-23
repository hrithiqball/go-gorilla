package main

import (
	"fmt"
	"local_my_api/db"
	"local_my_api/middleware"
	"local_my_api/model"
	"local_my_api/routes"
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

	dburl := os.Getenv("DATABASE_URL")
	fmt.Println(dburl)
	secretkey := os.Getenv("SECRET_KEY")
	fmt.Println(secretkey)

	err := db.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Error applying migrations: %v", err)
	}
	fmt.Println("Migrations applied successfully!")

	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)

	routes.SetupRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
