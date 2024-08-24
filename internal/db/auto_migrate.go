package db

import (
	"fmt"
	"local_my_api/internal/models"
	"log"
)

func Migrations() {
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Business{},
		&models.Product{},
	}

	for _, model := range modelsToMigrate {
		err := DB.AutoMigrate(model)
		if err != nil {
			log.Fatalf("Error applying migrations for model: %T %v", model, err)
		}
		fmt.Printf("Migrations applied successfully for model: %T\n", model)
	}
}
