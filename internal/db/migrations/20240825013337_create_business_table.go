package migrations

import "gorm.io/gorm"

func Migrate_create_business_table(tx *gorm.DB) error {
	// up
	return tx.Migrator().CreateTable(&Business{})
}

func Rollback_create_business_table(tx *gorm.DB) error {
	// down
	return tx.Migrator().DropTable(&Business{})
}

type Business struct {
	gorm.Model
	ID              string `gorm:"primaryKey" json:"id"`
	Name            string `json:"name"`
	CoverPhoto      string `json:"coverPhoto"`
	ProfilePhoto    string `json:"profilePhoto"`
	Email           string `gorm:"unique" json:"email"`
	Phone           string `gorm:"unique" json:"phone"`
	Address         string `json:"address"`
	Website         string `json:"website"`
	BusinessOwnerID string `json:"businessOwnerID"`
	BusinessOwner   User   `gorm:"foreignKey:BusinessOwnerID" json:"businessOwner"`
}
