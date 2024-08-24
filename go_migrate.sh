#!/bin/bash

if [ $# -lt 1 ]; then
    echo "Usage: $0 --<migration_type>"
    echo "Example: $0 --create_user_table"
    exit 1
fi

TIMESTAMP=$(date +"%Y%m%d%H%M%S")
MIGRATION_TYPE=$1

if [[ ! $MIGRATION_TYPE == --* ]]; then
    echo "Invalid argument. Use format: --<migration_type>"
    exit 1
fi

MIGRATION_NAME=${MIGRATION_TYPE:2}

MIGRATIONS_DIR="internal/db/migrations"

if [ ! -d "$MIGRATIONS_DIR" ]; then
    mkdir -p "$MIGRATIONS_DIR"
fi

MIGRATION_FILE="${MIGRATIONS_DIR}/${TIMESTAMP}_${MIGRATION_NAME}.go"
MIGRATION_FILENAME="${TIMESTAMP}_${MIGRATION_NAME}"

touch "$MIGRATION_FILE"

cat <<EOL > "$MIGRATION_FILE"
package migrations

import "gorm.io/gorm"

func Migrate_${MIGRATION_NAME}(tx *gorm.DB) error {
    // up
    return
}

func Rollback_${MIGRATION_NAME}(tx *gorm.DB) error {
    // down
    return
}
EOL

echo "Migration file created: $MIGRATION_FILE"

if command -v xclip &> /dev/null; then
    echo "$MIGRATION_FILENAME" | xclip -selection clipboard
    echo "File name copied to clipboard using xclip. Open internal/db/migrations.go and paste the file name as ID."
elif command -v xsel &> /dev/null; then
    echo "$MIGRATION_FILENAME" | xsel --clipboard
    echo "File name copied to clipboard using xsel. Open internal/db/migrations.go and paste the file name as ID."
else
    echo "Clipboard utility not found. Please install xclip or xsel."
fi
