package migrations

import "gorm.io/gorm"

func Migrate_update_price_to_decimal(tx *gorm.DB) error {
	// up
	if err := tx.Exec(
		`ALTER TABLE products 
        ALTER COLUMN price TYPE NUMERIC USING price::NUMERIC`).Error; err != nil {
		return err
	}

	return nil
}

func Rollback_update_price_to_decimal(tx *gorm.DB) error {
	// down
	if err := tx.Exec(
		`ALTER TABLE products 
        ALTER COLUMN price TYPE INT USING price::INTEGER`).Error; err != nil {
		return err
	}

	return nil
}
