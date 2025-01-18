package config

import (
	"context"
	"fmt"

	migrations "fiber-starter/db/migrations"

	"github.com/jackc/pgx/v4/pgxpool"
)

func RunSQLMigrations(db *pgxpool.Pool) error {
	var Migrations = []string{
		migrations.CreateUsersTable,
		migrations.InsertUsersData,
		migrations.CreatePostsTable,
		migrations.InsertPostsData,
	}

	for i, migration := range Migrations {
		fmt.Printf("Running migration %d...\n", i+1)

		_, err := db.Exec(context.Background(), migration)
		if err != nil {
			fmt.Printf("Error executing migration %d: %v\n", i+1, err)
			fmt.Println("Skipping this migration and continuing with the next one...")
			continue
		}

		fmt.Printf("Migration %d applied successfully!\n", i+1)
	}

	return nil
}