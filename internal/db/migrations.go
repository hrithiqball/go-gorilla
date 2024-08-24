package db

import (
	"fmt"
	"local_my_api/internal/db/migrations"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
)

var migrationList = []*gormigrate.Migration{
	{
		ID:       "20240825012245_create_user_table",
		Migrate:  migrations.Migrate_create_user_table,
		Rollback: migrations.Rollback_create_user_table,
	},
	{
		ID:       "20240825013337_create_business_table",
		Migrate:  migrations.Migrate_create_business_table,
		Rollback: migrations.Rollback_create_business_table,
	},
	{
		ID:       "20240825014020_create_product_table",
		Migrate:  migrations.Migrate_create_product_table,
		Rollback: migrations.Rollback_create_product_table,
	},
}

func Migrations() error {
	gormigrate := gormigrate.New(DB, gormigrate.DefaultOptions, migrationList)

	if err := gormigrate.Migrate(); err != nil {
		log.Printf("An error occurred while migrating: %v", err)
		return fmt.Errorf("an error occurred while migrating: %v", err)
	}

	fmt.Println("🧩 Migrations applied successfully! ")
	return nil
}
