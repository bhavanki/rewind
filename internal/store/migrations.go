package store

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*
var dbMigrations embed.FS

func runMigrations(db *sql.DB, dialect string) error {
	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "migrations",
	}
	_, err := migrate.Exec(db, dialect, migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	// fmt.Printf("Ran %d migrations\n", n)
	return nil
}
