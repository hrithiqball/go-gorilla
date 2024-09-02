package migrations

import "gorm.io/gorm"

func Migrate_update_product_table(tx *gorm.DB) error {
	// up
	if err := tx.Exec(`ALTER TABLE products ALTER COLUMN price TYPE NUMERIC(20, 2) USING price::NUMERIC`).Error; err != nil {
		return err
	}
	return nil
}

func Rollback_update_product_table(tx *gorm.DB) error {
	// down
	if err := tx.Exec(`ALTER TABLE products ALTER COLUMN price TYPE BIGINT USING price::BIGINT`).Error; err != nil {
		return err
	}
	return nil
}
