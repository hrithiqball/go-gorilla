package migrations

import "gorm.io/gorm"

func Migrate_create_product_table(tx *gorm.DB) error {
	// up
	return tx.Migrator().CreateTable(&Product{})
}

func Rollback_create_product_table(tx *gorm.DB) error {
	// down
	return tx.Migrator().DropTable(&Product{})
}

type Product struct {
	gorm.Model
	ID          string   `gorm:"primaryKey" json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Stock       int      `json:"stock"`
	BusinessID  string   `json:"businessID"`
	Business    Business `gorm:"foreignKey:BusinessID" json:"business"`
}
