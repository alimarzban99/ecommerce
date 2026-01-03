package main

import (
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/alimarzban99/ecommerce/config"
	"github.com/alimarzban99/ecommerce/pkg/migration"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	// Load configuration
	config.Load()

	// Get database connection string
	dbURL := migration.GetDatabaseURL()

	// Open database connection
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Create migrate instance
	m, err := migrate.New(
		fmt.Sprintf("file://%s", migration.GetMigrationPath()),
		dbURL,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	// Execute command
	switch command {
	case "up":
		if err := m.Up(); err != nil {
			if err == migrate.ErrNoChange {
				log.Println("✅ No new migrations to run")
				os.Exit(0)
			}
			log.Fatalf("Failed to run migrations: %v", err)
		}
		log.Println("✅ Migrations completed successfully")
		printVersion(m)

	case "down":
		steps := 1
		if len(os.Args) > 2 {
			if os.Args[2] == "-all" {
				steps = -1
			} else {
				var err error
				steps, err = strconv.Atoi(os.Args[2])
				if err != nil {
					log.Fatalf("Invalid steps value: %v", err)
				}
			}
		}

		if steps == -1 {
			if err := m.Down(); err != nil {
				if err == migrate.ErrNoChange {
					log.Println("✅ No migrations to rollback")
					os.Exit(0)
				}
				log.Fatalf("Failed to rollback all migrations: %v", err)
			}
			log.Println("✅ All migrations rolled back successfully")
		} else {
			if err := m.Steps(-steps); err != nil {
				if err == migrate.ErrNoChange {
					log.Println("✅ No migrations to rollback")
					os.Exit(0)
				}
				log.Fatalf("Failed to rollback migrations: %v", err)
			}
			log.Printf("✅ Rolled back %d migration(s) successfully", steps)
		}
		printVersion(m)

	case "version":
		printVersion(m)

	case "status":
		printStatus(m)

	case "force":
		if len(os.Args) < 3 {
			log.Fatalf("Usage: go run cmd/migrate/main.go force <version>")
		}
		version, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid version: %v", err)
		}
		if err := m.Force(version); err != nil {
			log.Fatalf("Failed to force version: %v", err)
		}
		log.Printf("✅ Forced migration version to %d", version)
		printVersion(m)

	case "drop":
		log.Println("⚠️  This will drop the entire database schema_migrations table.")
		log.Println("⚠️  Are you sure? This action cannot be undone.")
		log.Fatalf("Use: migrate -path database/migrations -database \"%s\" drop -f (if you have migrate CLI with postgres support)", dbURL)

	default:
		log.Fatalf("Unknown command: %s", command)
		printUsage()
		os.Exit(1)
	}
}

func printVersion(m *migrate.Migrate) {
	version, dirty, err := m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			log.Println("Current migration version: none")
			return
		}
		log.Fatalf("Failed to get version: %v", err)
	}
	log.Printf("Current migration version: %d (dirty: %v)", version, dirty)
}

func printStatus(m *migrate.Migrate) {
	// Get current version
	currentVersion, dirty, err := m.Version()
	hasVersion := err == nil
	if err != nil && err != migrate.ErrNilVersion {
		log.Fatalf("Failed to get version: %v", err)
	}

	// Get migration files
	migrationsDir := migration.GetMigrationPath()
	migrations, err := getMigrationFiles(migrationsDir)
	if err != nil {
		log.Fatalf("Failed to read migration files: %v", err)
	}

	// Print header
	fmt.Println("\n📊 Migration Status")
	fmt.Println(strings.Repeat("=", 80))

	if !hasVersion {
		fmt.Println("Current Version: none")
		fmt.Println("Dirty: false")
	} else {
		fmt.Printf("Current Version: %d\n", currentVersion)
		fmt.Printf("Dirty: %v\n", dirty)
	}

	if hasVersion && dirty {
		fmt.Println("\n⚠️  WARNING: Database is in a dirty state!")
		fmt.Println("   This means a migration failed partway through.")
		fmt.Println("   You may need to manually fix the database state.")
	}

	// Print migration list
	if len(migrations) > 0 {
		fmt.Println("\nMigrations:")
		fmt.Println(strings.Repeat("-", 80))
		fmt.Printf("%-10s %-60s %-10s\n", "Version", "Name", "Status")
		fmt.Println(strings.Repeat("-", 80))

		for _, mig := range migrations {
			status := "down"
			if hasVersion && mig.Version <= uint(currentVersion) {
				status = "up"
			}
			statusIcon := "⬇️"
			if status == "up" {
				statusIcon = "✅"
			}
			fmt.Printf("%-10d %-60s %-10s %s\n", mig.Version, mig.Name, status, statusIcon)
		}
		fmt.Println(strings.Repeat("-", 80))
	} else {
		fmt.Println("\nNo migration files found.")
	}
	fmt.Println()
}

type migrationFile struct {
	Version uint
	Name    string
}

func getMigrationFiles(dir string) ([]migrationFile, error) {
	var migrations []migrationFile
	migrationMap := make(map[uint]string)

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		filename := d.Name()
		// Match pattern: 000001_name.up.sql or 000001_name.down.sql
		if strings.HasSuffix(filename, ".up.sql") || strings.HasSuffix(filename, ".down.sql") {
			// Extract version number (first 6 digits)
			if len(filename) >= 6 {
				versionStr := filename[:6]
				version, err := strconv.ParseUint(versionStr, 10, 32)
				if err == nil {
					// Extract name (remove version, .up.sql or .down.sql)
					name := filename[7:]
					if strings.HasSuffix(name, ".up.sql") {
						name = strings.TrimSuffix(name, ".up.sql")
					} else if strings.HasSuffix(name, ".down.sql") {
						name = strings.TrimSuffix(name, ".down.sql")
					}
					migrationMap[uint(version)] = name
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Convert map to slice and sort
	for version, name := range migrationMap {
		migrations = append(migrations, migrationFile{
			Version: version,
			Name:    name,
		})
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations, nil
}

func printUsage() {
	fmt.Println("Usage: go run cmd/migrate/main.go <command> [options]")
	fmt.Println("\nCommands:")
	fmt.Println("  up              Run all pending migrations")
	fmt.Println("  down [n|-all]   Rollback n migrations (default: 1) or all if -all")
	fmt.Println("  version         Show current migration version")
	fmt.Println("  status          Show detailed migration status")
	fmt.Println("  force <version> Force migration version (use with caution)")
	fmt.Println("\nExamples:")
	fmt.Println("  go run cmd/migrate/main.go up")
	fmt.Println("  go run cmd/migrate/main.go down")
	fmt.Println("  go run cmd/migrate/main.go down 2")
	fmt.Println("  go run cmd/migrate/main.go down -all")
	fmt.Println("  go run cmd/migrate/main.go version")
	fmt.Println("  go run cmd/migrate/main.go status")
	fmt.Println("  go run cmd/migrate/main.go force 1")
}
