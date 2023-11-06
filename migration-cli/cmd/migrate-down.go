package cmd

import (
	"fmt"
	"log"

	"github.com/bug-free-happiness/migration-cli/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

func init() {
	migrateDownCmd := &cobra.Command{
		Use:   "down",
		Short: "migrate from v2 to v1",
		Long:  "Command to downgrde database from v2 to v1",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate down command")
			db := database.Open()

			dbDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
			if err != nil {
				log.Fatalf("instance error: %v\n", err)
			}

			fileSource, err := (&file.File{}).Open("file://db/migrations")
			if err != nil {
				log.Fatalf("opening file error: %v\n", err)
			}

			m, err := migrate.NewWithInstance("file", fileSource, "mydb", dbDriver)
			if err != nil {
				log.Fatalf("migrate error: %v\n", err)
			}

			if err = m.Down(); err != nil {
				log.Fatalf("migrate up error: %v\n", err)
			}

			fmt.Println("Migrate down done with success")
		},
	}

	migrateCmd.AddCommand(migrateDownCmd)
}
