package migrations

import "gorm.io/gorm"

func Migrate_add_photo_type_product(tx *gorm.DB) error {
	// up
	if err := tx.Exec(`
        ALTER TABLE products 
        ADD COLUMN photos TEXT[], 
        ADD COLUMN type TEXT,
        ADD COLUMN feature_photo TEXT;
    `).Error; err != nil {
		return err
	}
	return nil
}

func Rollback_add_photo_type_product(tx *gorm.DB) error {
	// down
	if err := tx.Exec(`
        ALTER TABLE products 
        DROP COLUMN IF EXISTS photos, 
        DROP COLUMN IF EXISTS type,
        DROP COLUMN IF EXISTS feature_photo;
    `).Error; err != nil {
		return err
	}
	return nil
}
