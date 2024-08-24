package migrations

import "gorm.io/gorm"

func Migrate_create_user_table(tx *gorm.DB) error {
	// up
	return tx.Migrator().CreateTable(&User{})
}

func Rollback_create_user_table(tx *gorm.DB) error {
	// down
	return tx.Migrator().DropTable(&User{})
}

type User struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	Name         string
	Email        string `gorm:"unique"`
	PasswordHash string
}
